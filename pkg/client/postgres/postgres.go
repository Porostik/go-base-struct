package postgres

import (
	"architecture/internal/config"
	"architecture/pkg/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...interface{}) pgx.Rows
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, config config.PostgresConfig) (*pgxpool.Pool, error) {
	var (
		pool *pgxpool.Pool
		err  error
	)

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.Username, config.Password, config.Host, config.Port, config.Database,
	)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)

		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("Postgres connection error", err)
	}

	return pool, nil
}
