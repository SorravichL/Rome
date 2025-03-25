package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"go-backend/types"
)

var DB *sql.DB
var connected bool

func Init() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("⚠️ DATABASE_URL not set, skipping DB connection")
		connected = false
		return
	}

	DB, err = sql.Open("postgres", dsn)
	if err != nil || DB.Ping() != nil {
		log.Println("❌ Failed to connect to DB:", err)
		connected = false
		return
	}

	connected = true
	log.Println("✅ Connected to DB")
}

func InsertMessage(m types.DBMessage) error {
	if !connected {
		return nil
	}
	query := `INSERT INTO "Message" (sender, receiver, message, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(query, m.Sender, m.Receiver, m.Message, m.Timestamp)
	return err
}

func IsConnected() bool {
	return connected
}
