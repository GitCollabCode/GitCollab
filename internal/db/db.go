package db

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const POSTGRES_DRIVER_TIMEOUT = 10

type PostgresDriver struct {
	Pool *pgxpool.Pool
	Log  *logrus.Logger
}

var ErrProfiletNotFound = fmt.Errorf("row not found")
var ErrProfiletMultipleRowsAffected = fmt.Errorf("more than one, row was affected in a single row operation")
var ErrProfiletMultipleRowsRetunred = fmt.Errorf("more than one, row was returned when was expected")

func NewPostgresDriver(connectionUrl string, log *logrus.Logger) (*PostgresDriver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), POSTGRES_DRIVER_TIMEOUT*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connectionUrl)
	if err != nil {
		return nil, err
	}

	return &PostgresDriver{Pool: pool, Log: log}, nil
}

func (pd *PostgresDriver) TransactOneRow(sqlStatement string, args ...any) error {
	tx, err := pd.Pool.Begin(context.Background())
	if err != nil {
		pd.Log.Fatal(err)
		return err
	}

	res, err := tx.Exec(context.Background(), sqlStatement, args...)
	if err != nil {
		pd.Log.Errorf("TransactOneRow database EXEC failed: %s", err.Error())
		rollbackErr := tx.Rollback(context.Background())
		if rollbackErr != nil {
			pd.Log.Fatalf("TransactOneRow rollback failed: %s", rollbackErr.Error())
		}
		return err
	}

	if res.RowsAffected() > 1 {
		err = ErrProfiletMultipleRowsAffected
		pd.Log.Errorf("TransactOneRow failed: %s", err.Error())
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		pd.Log.Fatalf("TransactOneRow commit failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *PostgresDriver) QueryRow(sqlStatement string, structure interface{}, args ...any) error {
	rows, err := pd.Pool.Query(context.Background(), sqlStatement, args...)
	if err != nil {
		pd.Log.Errorf("QueryRow Query failed: %s", err.Error())
		return err
	}

	pgxscan.ScanOne(structure, rows)
	if err != nil {
		pd.Log.Errorf("QueryRow ScanOne failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *PostgresDriver) QueryRows(sqlStatement string, structure interface{}, args ...any) error {
	rows, err := pd.Pool.Query(context.Background(), sqlStatement, args...)
	if err != nil {
		pd.Log.Errorf("QueryRows Query failed: %s", err.Error())
		return err
	}

	pgxscan.ScanAll(structure, rows)
	if err != nil {
		pd.Log.Errorf("QueryRows ScanAll failed: %s", err.Error())
		return err
	}

	return nil
}
