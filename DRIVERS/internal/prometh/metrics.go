package prometh

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Общие метрики для GRPC сервера
var (
	// GrpcRequestsTotal - общее количество GRPC запросов.
	GrpcRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of GRPC requests.",
		},
		[]string{"method", "status"}, // "status" может быть "success", "failure"
	)

	// GrpcRequestDuration - длительность GRPC запросов.
	GrpcRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds",
			Help:    "GRPC request duration in seconds.",
			Buckets: prometheus.DefBuckets, // Или определите свои бакеты: []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10}
		},
		[]string{"method"},
	)
)

// Метрики, специфичные для сервиса водителей
var (
	// RideAcceptedCounter - количество успешно принятых поездок.
	RideAcceptedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "drivers_ride_accepted_total",
		Help: "Total number of rides accepted by drivers.",
	})

	// RideCompletedCounter - количество успешно завершенных поездок.
	RideCompletedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "drivers_ride_completed_total",
		Help: "Total number of rides completed by drivers.",
	})

	// RideCanceledCounter - количество отмененных поездок.
	RideCanceledCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "drivers_ride_canceled_total",
		Help: "Total number of rides canceled by drivers.",
	})

	// KafkaProducedMessages - количество успешно отправленных сообщений Kafka.
	KafkaProducedMessages = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "drivers_kafka_produced_messages_total",
			Help: "Total number of messages successfully produced to Kafka.",
		},
		[]string{"topic"},
	)

	// KafkaProduceErrors - количество ошибок при отправке сообщений Kafka.
	KafkaProduceErrors = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "drivers_kafka_produce_errors_total",
			Help: "Total number of errors when producing messages to Kafka.",
		},
	)

	// KafkaConsumedMessages - количество успешно потребленных сообщений Kafka.
	KafkaConsumedMessages = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "drivers_kafka_consumed_messages_total",
			Help: "Total number of messages successfully consumed from Kafka.",
		},
		[]string{"topic"},
	)

	// KafkaConsumeErrors - количество ошибок при потреблении сообщений Kafka.
	KafkaConsumeErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "drivers_kafka_consume_errors_total",
			Help: "Total number of errors when consuming messages from Kafka.",
		},
		[]string{"topic"}, // Может быть полезно знать, в каком топике произошла ошибка
	)

	// KafkaRequestTimeouts - количество таймаутов при ожидании ответов по Kafka.
	KafkaRequestTimeouts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "drivers_kafka_request_timeouts_total",
			Help: "Total number of timeouts when waiting for Kafka responses.",
		},
		[]string{"event_type"}, // Полезно знать, для какого типа события произошел таймаут
	)
)
