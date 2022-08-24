// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"log"
	"os"

	pgx "github.com/jackc/pgx/v4"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type attended struct {
	Attended string `json:"attended"`
}

// HelloPubSub consumes a Pub/Sub message.
func CountAttended(ctx context.Context, m PubSubMessage) error {
	a := new(attended)

	if err := json.Unmarshal(m.Data, a); err != nil {
		return err
	}
	log.Printf("%#v", a)

	conn, err := pgx.Connect(ctx, os.Getenv("CRDB_CONN"))
	if err != nil {
		log.Printf("unable to connect to crdb: %v", err)
		return err
	}
	_, err = conn.Exec(ctx, "INSERT INTO webnesday (attended, count) VALUES ($1, 1) ON CONFLICT (attended) DO UPDATE SET count = (SELECT count + 1 FROM webnesday WHERE attended = EXCLUDED.attended)", a.Attended)
	log.Println("inserted: ", err)

	return err
}
