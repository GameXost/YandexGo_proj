package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/prometh"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/config"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/repository"
	server "github.com/GameXost/YandexGo_proj/DRIVERS/internal/server"
	"github.com/GameXost/YandexGo_proj/DRIVERS/internal/services"
)

var publicKey *rsa.PublicKey

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Load config
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// 2. Load keys
	publicKey, err := server.LoadPublicKey(cfg.JWT.PublicKeyPath)
	if err != nil {
		log.Fatalf("failed to load public key: %v", err)
	}
	privateKey, err := server.LoadPrivateKey(cfg.JWT.PrivateKeyPath)
	if err != nil {
		log.Fatalf("failed to load private key: %v", err)
	}
	_ = privateKey
	prometh.InitPrometheus(cfg.Prometheus.Port)

	// 3. DB connection
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.SSLMode,
	)
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to create pool: %v", err)
	}
	defer dbpool.Close()
	log.Println("PGX working")

	// 4. REDIS connection
	redisAddr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	log.Println("Redis working")

	// 5. Kafka connection
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: cfg.Kafka.Brokers,
		Topic:   cfg.Kafka.Topics.Rides,
	})
	defer kafkaWriter.Close()
	log.Println("Kafka writer ready")

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Kafka.Brokers,
		Topic:    cfg.Kafka.Topics.Rides,
		GroupID:  "driver-service",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	defer kafkaReader.Close()
	log.Println("Kafka reader ready")

	// Сигналы для graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	repo := repository.NewDriverRepository(dbpool)
	driverService := services.NewDriverService(repo, redisClient, redisClient, kafkaWriter, cfg.Kafka.Topics.UserRequests, cfg.Kafka.Topics.Rides)

	// --- Kafka consumer ---
	go driverService.StartKafkaConsumer(ctx, kafkaReader)

	// server up
	sv := &server.DriverServer{
		Service: driverService,
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(server.AuthInterceptor(publicKey, cfg.Auth.Disabled)),
	)
	pb.RegisterDriversServer(grpcServer, sv)
	grpcListener, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil {
		log.Fatalf("Unable to listen on %s: %v", cfg.Server.Port, err)
	}
	go func() {
		log.Printf("GRPC server listening on %s", cfg.Server.Port)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("Unable to start grpc server: %v", err)
		}
	}()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(customHeaderMatcher),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterDriversHandlerFromEndpoint(ctx, mux, "localhost"+cfg.Server.Port, opts)
	if err != nil {
		log.Fatalf("Unable to register handler: %v", err)
	}
	log.Println("Mux gateway Listening on", cfg.Server.HTTPPort)
	srv := &http.Server{
		Addr:    cfg.Server.HTTPPort,
		Handler: allowCORS(server.JWTHTTPMiddleware(publicKey)(mux)),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Unable to listen on %s: %v", cfg.Server.HTTPPort, err)
		}
	}()

	// Ждём сигнал завершения
	<-sigs
	log.Println("Graceful shutdown started...")
	cancel()
	grpcServer.GracefulStop()
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
	log.Println("Shutdown complete")
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func customHeaderMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case "authorization":
		return "authorization", true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
