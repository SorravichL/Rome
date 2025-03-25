package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Message used for communication via API (JSON)
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	Date    string `json:"date"` // RFC3339 format
}

// DBMessage matches Prisma model for inserting into DB
type DBMessage struct {
	Sender    string
	Receiver  string
	Message   string
	Timestamp time.Time
}

var db *sql.DB

func initDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("❌ Failed to ping database:", err)
	}
	log.Println("✅ Connected to database")
}

func insertMessageToDB(m DBMessage) error {
	query := `INSERT INTO "Message" (sender, receiver, message, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, m.Sender, m.Receiver, m.Message, m.Timestamp)
	return err
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST /send")

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	log.Printf("Forwarding message from %s to %s", msg.From, msg.To)

	// Convert date to time.Time
	t, err := time.Parse(time.RFC3339, msg.Date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Save to DB
	dbMsg := DBMessage{
		Sender:    msg.From,
		Receiver:  msg.To,
		Message:   msg.Message,
		Timestamp: t,
	}
	if err := insertMessageToDB(dbMsg); err != nil {
		log.Println("❌ Failed to insert to DB:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	// Forward to TS
	tsURL := os.Getenv("TS_BACKEND_URL")
	if tsURL == "" {
		tsURL = "http://localhost:5002/log"
	}

	jsonData, _ := json.Marshal(msg)
	resp, err := http.Post(tsURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("❌ Failed to forward to TS:", err)
		http.Error(w, "Failed to forward message", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "forwarded"})
}

func handleLog(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	log.Printf("Message from %s to %s at %s: %s", msg.From, msg.To, msg.Date, msg.Message)

	t, err := time.Parse(time.RFC3339, msg.Date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	dbMsg := DBMessage{
		Sender:    msg.From,
		Receiver:  msg.To,
		Message:   msg.Message,
		Timestamp: t,
	}
	if err := insertMessageToDB(dbMsg); err != nil {
		log.Println("❌ Failed to insert to DB:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}

func main() {
	godotenv.Load()
	initDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	addr := ":" + port

	http.HandleFunc("/send", handleSend)
	http.HandleFunc("/log", handleLog)

	log.Printf("Go backend is running on port %s", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
