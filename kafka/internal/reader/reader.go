package reader

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func New(addr, group, topic string) (*Reader, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{addr},
		GroupID:        group,
		Topic:          topic,
		CommitInterval: 1 * time.Second,
		StartOffset:    kafka.FirstOffset,
	})

	return &Reader{
		group:  group,
		topic:  topic,
		reader: r,
	}, nil
}

type Reader struct {
	group  string
	topic  string
	reader *kafka.Reader
}

func (r *Reader) Close() error {
	return r.reader.Close()
}

func (r *Reader) Read(ctx context.Context) (string, error) {
	message, err := r.reader.ReadMessage(ctx)
	if err != nil {
		return "", fmt.Errorf("group %s pull message from %s failed: %w", r.group, r.topic, err)
	}

	return string(message.Value), nil
}
