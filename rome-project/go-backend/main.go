package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
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

	// Get TS backend URL from environment (default to localhost)
	tsURL := os.Getenv("TS_BACKEND_URL")
	if tsURL == "" {
		tsURL = "http://localhost:5002/log"
	}

	resp, err := http.Post(tsURL, "application/json", bytes.NewBuffer(jsonData))
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
	godotenv.Load()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	addr := ":" + port

	mux := http.NewServeMux()
	mux.HandleFunc("/send", handleSend)
	mux.HandleFunc("/log", handleLog)

	log.Printf("Go backend is running on port %s", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Server error:", err)
	}
}
