package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

var db struct {
	url  string
	ctx  context.Context
	conn *pgx.Conn
}

func Init(dbUser, dbPassword, dbHost, dbPort, dbName string) {
	db.url = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName,
	)
	db.ctx = context.Background()
	connect()
	go reconnectRoutine()
}

func Close() {
	_ = db.conn.Close(db.ctx)
	db.conn = nil
}

func connect() {
	var err error
	for retries := 1; ; retries++ {
		db.conn, err = pgx.Connect(db.ctx, db.url)
		if err != nil {
			log.Printf("#%v: %v", retries, err)
		} else {
			log.Print("Database connection established")
			return
		}
		time.Sleep(time.Second * time.Duration(retries))
	}
}

func reconnectRoutine() {
	var pingRetries int
	for {
		for err := db.conn.Ping(db.ctx); err != nil; pingRetries++ {
			if pingRetries >= 3 {
				log.Printf("Database connection lost!")
				Close()
				connect()
				break
			}
			time.Sleep(time.Second)
		}
		pingRetries = 0
		time.Sleep(time.Second * 5)
	}
}
