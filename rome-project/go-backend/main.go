package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go-backend/db"
	"go-backend/types"

	"github.com/joho/godotenv"
)

func handleSend(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST /send")

	var msg types.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	log.Printf("Forwarding message from %s to %s", msg.From, msg.To)

	t, err := time.Parse(time.RFC3339, msg.Date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	dbMsg := types.DBMessage{
		Sender:    msg.From,
		Receiver:  msg.To,
		Message:   msg.Message,
		Timestamp: t,
	}

	if db.IsConnected() {
		if err := db.InsertMessage(dbMsg); err != nil {
			log.Println("❌ Failed to insert to DB:", err)
		}
	} else {
		log.Println("⚠️ Skipped DB insert: not connected")
	}

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
	var msg types.Message
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

	dbMsg := types.DBMessage{
		Sender:    msg.From,
		Receiver:  msg.To,
		Message:   msg.Message,
		Timestamp: t,
	}

	if db.IsConnected() {
		if err := db.InsertMessage(dbMsg); err != nil {
			log.Println("❌ Failed to insert to DB:", err)
		}
	} else {
		log.Println("⚠️ Skipped DB insert: not connected")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}

func main() {
	godotenv.Load()
	db.Init()

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