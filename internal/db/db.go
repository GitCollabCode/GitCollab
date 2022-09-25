package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type PostgresDriver struct {
	Connection pgx.Conn
}

func ConnectPostgres(connectionUrl string) (*PostgresDriver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, connectionUrl)
	if err != nil {
		return nil, err
	}
	return &PostgresDriver{Connection: *conn}, err
}
 