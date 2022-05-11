package db_client

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Client struct {
	ctx  context.Context
	conn *pgxpool.Pool
}

func NewClient(ctx context.Context, user, password, host, port, name string) *Client {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", user, password, host, port, name,
	)
	return &Client{
		ctx:  ctx,
		conn: connectWithRetries(ctx, url),
	}
}

func (db *Client) Close() {
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
