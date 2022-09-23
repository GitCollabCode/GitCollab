package db

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type PostgresDriver struct {
	Connection pgx.Conn
}

func ConnectPostgres() (*PostgresDriver, error) {
	postgresUrl := os.Getenv("POSTGRES_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, postgresUrl)
	if err != nil {
		return nil, err
	}
	return &PostgresDriver{Connection: *conn}, err
}
