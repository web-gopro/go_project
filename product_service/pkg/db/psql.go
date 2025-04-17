package db

import (
	"context"
	"fmt"
	"gitlab.com/e-market724/back-end/product_service/config"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getConfig(cfg config.PgConfig) *pgxpool.Config {

	const (
		defaultMaxConns          = int32(4)
		defaultMinConns          = int32(0)
		defaultMaxConnLifeTime   = time.Hour
		defaultMinConnIdleTime   = time.Minute * 30
		defaultHealthCheckPeriod = time.Minute
		defaultConnectTimeout    = time.Second * 5
	)

	var databaseUrl string = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	)

	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		fmt.Println("error", err)
		return nil
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnIdleTime = defaultMinConnIdleTime
	dbConfig.MaxConnLifetime = defaultMaxConnLifeTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(context.Context, *pgx.Conn) bool {
		log.Println("before acquiring")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("closed conection pool db")
	}

	return dbConfig
}

func ConnDB(cfg config.PgConfig) (*pgxpool.Pool, error) {

	connPool, err := pgxpool.NewWithConfig(context.Background(), getConfig(cfg))
	if err != nil {
		log.Fatal("error while creating connection to the database")
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("error while acquiring connection from the database")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("could not ping db")
	}

	fmt.Println("connected to db")

	return connPool, nil
}
