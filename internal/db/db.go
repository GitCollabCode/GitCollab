package db

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type PostgresDriver struct {
	Log        *echo.Logger
	Connection pgx.Conn
}

func ConnectPostgres(Log echo.Logger) (*PostgresDriver, error) {
	postgresUrl := os.Getenv("POSTGRES_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, postgresUrl)
	if err != nil {
		return nil, err
	}

	Log.Info("Connected to postgres db")
	return &PostgresDriver{Connection: *conn, Log: &Log}, err
}
