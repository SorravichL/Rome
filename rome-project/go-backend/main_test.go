package main

import (
	"testing"
	"go-backend/types"
)

func TestMessageStructure(t *testing.T) {
	msg := types.Message{
		From:    "go-service",
		To:      "ts-service",
		Message: "Hi there!",
		Date:    "2025-03-25T15:00:00Z",
	}

	if msg.From == "" || msg.To == "" || msg.Message == "" || msg.Date == "" {
		t.Errorf("Message fields should not be empty")
	}
}
