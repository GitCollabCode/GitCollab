package db

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type PostgresDriver struct {
	Log        *logrus.Logger
	Connection pgx.Conn
}

func ConnectPostgres(log *logrus.Logger) (*PostgresDriver, error) {
	postgresUrl := os.Getenv("POSTGRES_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, postgresUrl)
	if err != nil {
		return nil, err
	}

	log.Info("Connected to postgres db")
	return &PostgresDriver{Connection: *conn, Log: log}, err
}
