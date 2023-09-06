package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
)

var (
	KafkaHost              = os.Getenv("KAFKA_HOST")
	KafkaPort              = os.Getenv("KAFKA_PORT")
	KafkaTopic             = os.Getenv("KAFKA_TOPIC")
	KafkaTopicAuth         = os.Getenv("KAFKA_TOPIC_AUTH")
	KafkaTopicUser         = os.Getenv("KAFKA_TOPIC_USER")
	KafkaTopicProfile      = os.Getenv("KAFKA_TOPIC_PROFILE")
	KafkaConsumerGroup     = os.Getenv("KAFKA_CONSUMER_GROUP")
	KafkaAddressFamily     = os.Getenv("KAFKA_ADDRESS_FAMILY")
	KafkaSessionTimeout, _ = strconv.Atoi(os.Getenv("KAFKA_SESSION_TIMEOUT"))
	KafkaAutoOffsetReset   = os.Getenv("KAFKA_AUTO_OFFSET_RESET")
)

func NewKafkaProducer() *kafka.Producer {
	broker := fmt.Sprintf("%s:%s", KafkaHost, KafkaPort)

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Error on creating kafka producer: %s\n", err.Error())
	}

	return producer
}

func NewKafkaConsumer() *kafka.Consumer {
	broker := fmt.Sprintf("%s:%s", KafkaHost, KafkaPort)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     broker,
		"broker.address.family": KafkaAddressFamily,
		"group.id":              KafkaConsumerGroup,
		"session.timeout.ms":    KafkaSessionTimeout,
		"auto.offset.reset":     KafkaAutoOffsetReset,
	})
	if err != nil {
		log.Fatalf("Error on creating kafka consumer: %s\n", err.Error())
	}

	log.Println("Kafka consumer created")

	return consumer
}
