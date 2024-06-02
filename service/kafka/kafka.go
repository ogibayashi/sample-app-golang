package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ogibayashi/sample-app-golang/config"
	"github.com/segmentio/kafka-go"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Writer struct {
	caCert          string
	broker          string
	topic           string
	username        string
	password        string
	topicAutoCreate bool
	writer          *kafkago.Writer
	dialer          *kafkago.Dialer
}

func NewWriter() (*Writer, error) {
	writer := &Writer{
		caCert:          config.GetString("kafka.cacert"),
		broker:          config.GetString("kafka.broker"),
		topic:           config.GetString("kafka.topic"),
		username:        config.GetString("kafka.username"),
		password:        config.GetString("kafka.password"),
		topicAutoCreate: config.GetBool("kafka.topic_auto_create"),
	}
	if writer.caCert == "" {
		return nil, fmt.Errorf("kafka.cacert is required")
	}
	if writer.broker == "" {
		return nil, fmt.Errorf("kafka.broker is required")
	}
	if writer.topic == "" {
		return nil, fmt.Errorf("kafka.topic is required")
	}
	if writer.username == "" {
		return nil, fmt.Errorf("kafka.username is required")
	}
	if writer.password == "" {
		return nil, fmt.Errorf("kafka.password is required")
	}
	err := writer.initialize()
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func (w *Writer) initialize() error {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(w.caCert))

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	mechanism, err := scram.Mechanism(scram.SHA512, w.username, w.password)
	if err != nil {
		return fmt.Errorf("failed to initialize mechanism %w\n", err)
	}
	dialer := &kafkago.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		TLS:           tlsConfig,
		SASLMechanism: mechanism,
	}
	w.dialer = dialer
	p, err := dialer.LookupPartitions(context.Background(), "tcp", w.broker, w.topic)
	if err != nil {
		if errors.Is(err, kafka.UnknownTopicOrPartition) && w.topicAutoCreate {
			if err := w.createTopic(); err != nil {
				return fmt.Errorf("failed to create topic: %w", err)
			}
			log.Printf("Topic created: %s", w.topic)
		} else {
			return fmt.Errorf("failed to lookup partitions %w\n", err)
		}

	}
	log.Printf("partitions: %v", p)
	log.Printf("broker: %v", w.broker)
	w.writer = kafkago.NewWriter(kafkago.WriterConfig{
		Brokers: []string{w.broker},
		Topic:   w.topic,
		Dialer:  dialer,
	})
	return nil
}

func (w *Writer) Write(s string) error {
	return w.writer.WriteMessages(context.Background(), kafka.Message{Value: []byte(s)})
}

func (w *Writer) Close() error {
	return w.writer.Close()
}

func (w *Writer) createTopic() error {
	conn, err := w.dialer.Dial("tcp", w.broker)
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}
	return conn.CreateTopics(kafkago.TopicConfig{
		Topic:         w.topic,
		NumPartitions: 1,
	})

}
