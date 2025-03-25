package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Message defines the structure for communication
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	Date    string `json:"date"` // Format: RFC3339
}

// Handle incoming message and forward to TypeScript
func handleSend(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST /send")

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	log.Printf("Forwarding message from %s to %s", msg.From, msg.To)

	jsonData, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Failed to encode message", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:5002/log", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error forwarding message to TypeScript:", err)
		http.Error(w, "Failed to forward message", http.StatusBadGateway)
		return
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Println("Error closing response body:", cerr)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "forwarded"}); err != nil {
		log.Println("Error encoding response:", err)
	}
}

// Handle logging message from TypeScript or other clients
func handleLog(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	log.Printf("Message from %s to %s at %s: %s", msg.From, msg.To, msg.Date, msg.Message)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "received"}); err != nil {
		log.Println("Error encoding response:", err)
	}
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
