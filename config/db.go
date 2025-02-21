package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pgOnce sync.Once
)

var DbPool *pgxpool.Pool

func ConnectDB() {
	var err error = nil
	pgOnce.Do(func() {
		pool, erro := pgxpool.NewWithConfig(context.Background(), Config())
		fmt.Println("Database Initialized.")
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Unable to create db connection pool : %v", err)
		}

		DbPool = pool
	})
}

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(50)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Millisecond * 50
	const defaultHealthCheckPeriod = time.Second
	const defaultConnectTimeout = time.Second * 5

	dbConfig, err := pgxpool.ParseConfig("postgres://botmg3002:BoTMG_3002@localhost:5432/student_engagement")
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
		return nil
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}
