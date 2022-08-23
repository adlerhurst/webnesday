package storage

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

type Writer struct{}

type WriterMock struct{}

func (m *WriterMock) Save(ctx context.Context, attended string) error {
	return nil
}

type PubsubWriter struct {
	topic *pubsub.Topic
}

func NewPubsubWriter() *PubsubWriter {
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	client, err := pubsub.NewClient(context.TODO(), projectID)
	if err != nil {
		log.Fatal("unable to connect to pubsub")
	}

	return &PubsubWriter{topic: client.Topic("webnesday")}
}

func (w *PubsubWriter) Save(ctx context.Context, attended string) error {
	data, err := json.Marshal(struct {
		Attended string `json:"attended"`
	}{attended})
	if err != nil {
		return err
	}

	res := w.topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err = res.Get(ctx)
	return err
}
