package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var db struct {
	url  string
	ctx  context.Context
	conn *pgxpool.Pool
}

func Init(dbUser, dbPassword, dbHost, dbPort, dbName string) {
	db.url = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName,
	)
	db.ctx = context.Background()
	db.conn = connectWithRetries(db.ctx, db.url)
}

func Close() {
	db.conn.Close()
	db.ctx = nil
	db.conn = nil
}

func connectWithRetries(ctx context.Context, url string) *pgxpool.Pool {
	for retries := 1; ; retries++ {
		conn, err := pgxpool.Connect(ctx, url)
		if err != nil {
			log.Printf("#%v: %v", retries, err)
		} else {
			log.Print("Database connection established")
			return conn
		}
		time.Sleep(time.Second * time.Duration(retries))
	}
}
