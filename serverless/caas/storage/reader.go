package storage

import (
	"context"
	"log"

	"github.com/adlerhurst/webnesday/serverless/caas/handler"
	pgx "github.com/jackc/pgx/v4"
)

type ReaderMock struct{}

func (m *ReaderMock) Get(context.Context) ([]*handler.ResultData, error) {
	return []*handler.ResultData{
		{
			Attended: "hdoor",
			Count:    134,
		},
		{
			Attended: "trea",
			Count:    98,
		},
	}, nil
}

type CRDBReader struct {
	conn *pgx.Conn
}

func NewCRDBReader() *CRDBReader {
	dsn := "postgresql://adlerhurst:3XH4qwkRd5nW-TQAfFqqUg@free-tier11.gcp-us-east1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster%3Dsilvan-1753"
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("unable to connect to crdb: %v", err)
	}
	return &CRDBReader{conn}
}

// CREATE TABLE webnesday (attended STRING PRIMARY KEY, count INT);
func (r *CRDBReader) Get(ctx context.Context) ([]*handler.ResultData, error) {
	rows, err := r.conn.Query(ctx, "SELECT attended, count FROM webnesday ORDER BY count")
	if err != nil {
		return nil, err
	}

	res := make([]*handler.ResultData, 0, 6)

	for rows.Next() {
		row := new(handler.ResultData)
		if err := rows.Scan(row.Attended, row.Count); err != nil {
			rows.Close()
			return nil, err
		}
	}

	rows.Close()

	return res, rows.Err()
}
