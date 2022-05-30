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

func NewClient(connCtx context.Context, user, password, host, name string) (*Client, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s/%s", user, password, host, name,
	)
	connection, err := connectWithRetries(connCtx, url)
	if err != nil {
		return nil, err
	}
	return &Client{
		ctx:  context.Background(),
		conn: connection,
	}, nil
}

func (db *Client) Close() {
	db.conn.Close()
	db.ctx = nil
	db.conn = nil
	log.Print("Database connection closed")
}

func connectWithRetries(ctx context.Context, url string) (*pgxpool.Pool, error) {
	for retries := 1; ; retries++ {
		conn, err := pgxpool.Connect(ctx, url)
		if err != nil {
			log.Printf("#%v: %v", retries, err)
			if ctx.Err() != nil {
				return nil, fmt.Errorf("%v: %w", ctx.Err(), err)
			}
		} else {
			log.Print("Database connection established")
			return conn, nil
		}
		time.Sleep(time.Second * time.Duration(retries))
	}
}
