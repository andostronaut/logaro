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
