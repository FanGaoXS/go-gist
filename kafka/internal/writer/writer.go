package writer

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func New(addr, topic string) (*Writer, error) {
	w := &kafka.Writer{
		Addr:  kafka.TCP(addr),
		Topic: topic,
		//Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	return &Writer{
		writer: w,
		topic:  topic,
	}, nil
}

type Writer struct {
	topic  string
	writer *kafka.Writer
}

func (w *Writer) Close() error {
	return w.writer.Close()
}

func (w *Writer) Write(ctx context.Context, msg string) error {
	m := kafka.Message{Value: []byte(msg)}
	if err := w.writer.WriteMessages(ctx, m); err != nil {
		return fmt.Errorf("push mesage to %s failed: %w", w.topic, err)
	}

	return nil
}
