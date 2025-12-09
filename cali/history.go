package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const historyFile = "calculator_history.json"

// HistoryEntry represents a single calculation in history
type HistoryEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Operation string    `json:"operation"`
	Result    float64   `json:"result"`
}

type History struct {
	entries []HistoryEntry
}

func NewHistory() *History {
	h := &History{
		entries: make([]HistoryEntry, 0),
	}
	h.loadFromFile()
	return h
}

func (h *History) loadFromFile() {
	data, err := os.ReadFile(historyFile)
	if err != nil {

		return
	}

	err = json.Unmarshal(data, &h.entries)
	if err != nil {
		fmt.Printf("Warning: Could not load history file: %v\n", err)
	}
}

func (h *History) saveToFile() error {
	data, err := json.MarshalIndent(h.entries, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal history: %v", err)
	}

	err = os.WriteFile(historyFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write history file: %v", err)
	}

	return nil
}

func (h *History) Add(operation string, result float64) {
	entry := HistoryEntry{
		Timestamp: time.Now(),
		Operation: operation,
		Result:    result,
	}
	h.entries = append(h.entries, entry)
}

func (h *History) Display() {
	if len(h.entries) == 0 {
		fmt.Println("\n📜 History is empty. No calculations performed yet.")
		return
	}

	fmt.Println("\n=================================")
	fmt.Println("  CALCULATION HISTORY")
	fmt.Println("=================================")

	for i, entry := range h.entries {
		timestamp := entry.Timestamp.Format("15:04:05")
		fmt.Printf("%d. [%s] %s = %.2f\n",
			i+1,
			timestamp,
			entry.Operation,
			entry.Result,
		)
	}
	fmt.Println("=================================")
}

func (h *History) Clear() {
	h.entries = make([]HistoryEntry, 0)
	fmt.Println("✓ History cleared successfully")
}

func (h *History) GetEntries() []HistoryEntry {
	return h.entries
}

func (h *History) Count() int {
	return len(h.entries)
}
