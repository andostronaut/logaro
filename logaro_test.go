package logaro

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestGenerateLogger(t *testing.T) {
	// Create a buffer to capture the log output
	buf := new(bytes.Buffer)

	// Create a logger with the buffer as the writer
	logger := GenerateLogger()
	logger.Writer = json.NewEncoder(buf)

	// Log an entry
	logger.Log("info", "Test log message", map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	})

	// Expected log entry
	expectedEntry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   "Test log message",
		Level:     "info",
		Fields: map[string]interface{}{
			"key1": "value1",
			"key2": 42,
		},
	}

	// Unmarshal the captured log output
	var capturedEntry LogEntry
	err := json.Unmarshal(buf.Bytes(), &capturedEntry)
	if err != nil {
		t.Fatalf("Error unmarshaling log entry: %s", err)
	}

	// Compare the captured log entry with the expected entry
	if !compareLogEntries(capturedEntry, expectedEntry) {
		t.Errorf("Log entry mismatch.\nExpected: %+v\nCaptured: %+v", expectedEntry, capturedEntry)
	}
}

func compareLogEntries(a, b LogEntry) bool {
	// Compare Timestamp, Message, Level, and Fields
	return a.Timestamp == b.Timestamp &&
		a.Message == b.Message &&
		a.Level == b.Level &&
		compareFields(a.Fields, b.Fields)
}

func compareFields(a, b map[string]interface{}) bool {
	// Compare the lengths of the fields maps
	if len(a) != len(b) {
		return false
	}

	// Compare each key-value pair in the fields maps
	for key, valA := range a {
		valB, ok := b[key]
		if !ok || !compareFieldValues(valA, valB) {
			return false
		}
	}

	return true
}

func compareFieldValues(a, b interface{}) bool {
	// Marshal and compare the JSON representations of the field values
	bytesA, errA := json.Marshal(a)
	bytesB, errB := json.Marshal(b)
	if errA != nil || errB != nil {
		return false
	}

	return bytes.Equal(bytesA, bytesB)
}
