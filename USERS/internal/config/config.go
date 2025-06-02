package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port     string `yaml:"port"`
	HTTPPort string `yaml:"http_port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type KafkaTopics struct {
	Rides               string `yaml:"rides"`
	UserRequests        string `yaml:"user_requests"`
	UserNotifications   string `yaml:"user_notifications"`
	DriverRequests      string `yaml:"driver_requests"`
	DriverNotifications string `yaml:"driver_notifications"`
}

type KafkaConfig struct {
	Brokers []string    `yaml:"brokers"`
	Topics  KafkaTopics `yaml:"topics"`
}

type JWTConfig struct {
	SecretKey      string `yaml:"secret_key"`
	PublicKeyPath  string `yaml:"public_key_path"`
	PrivateKeyPath string `yaml:"private_key_path"`
}

type LoggingConfig struct {
	Level string `yaml:"level"`
}

type GRPCConfig struct {
	MaxConnectionAge string `yaml:"max_connection_age"`
	KeepaliveTime    string `yaml:"keepalive_time"`
}

type AuthConfig struct {
	Disabled bool `yaml:"disabled"`
}

type PrometheusConfig struct {
	Port string `yaml:"port"`
}

type SwaggerConfig struct {
	File string `yaml:"file"`
}

type Config struct {
	Server     ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	Redis      RedisConfig      `yaml:"redis"`
	Kafka      KafkaConfig      `yaml:"kafka"`
	JWT        JWTConfig        `yaml:"jwt"`
	Logging    LoggingConfig    `yaml:"logging"`
	GRPC       GRPCConfig       `yaml:"grpc"`
	Auth       AuthConfig       `yaml:"auth"`
	Prometheus PrometheusConfig `yaml:"prometheus"`
	Swagger    SwaggerConfig    `yaml:"swagger"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
