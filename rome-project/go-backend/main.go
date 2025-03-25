package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// User defines the structure of the user data
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	Date    string `json:"date"` // Use time.Now().Format(time.RFC3339)
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Received POST /send")

	// Decode the message from request body
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("🔄 Forwarding message from %s to %s: %s", msg.From, msg.To, msg.Message)

	// Forward the message to TypeScript
	jsonData, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Failed to encode message", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:5002/log", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending to TS:", err)
		http.Error(w, "Failed to forward message", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Respond back to sender
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "forwarded"})
}

func handleLog(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message", http.StatusBadRequest)
		return
	}

	log.Printf("📨 Message from %s to %s @ %s: %s", msg.From, msg.To, msg.Date, msg.Message)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", handleSend)
	mux.HandleFunc("/log", handleLog)

	log.Println("Go backend is running at http://localhost:5001")
	if err := http.ListenAndServe(":5001", mux); err != nil {
		log.Fatal("Server error:", err)
	}
}