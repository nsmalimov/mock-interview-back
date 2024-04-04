package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"mock-interview-back/app/logger"
)

const (
	dbUrl = "postgres://myuser:mock-interview@31.129.51.12:5432/mock_interview"
)

type Connect struct {
	conn *pgx.Conn
}

func New() (*Connect, error) {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			logger.Error("error when try pgx.Connect, err: %s", err)
		}
	}(conn, context.Background())

	return &Connect{}, nil
}
