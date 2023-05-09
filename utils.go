package logaro

import (
	"bytes"
	"encoding/json"
)

// isEnabled checks if the given log level is enabled based on the logger's configured level.
// It uses a map to associate the log levels with numeric values.
// The function compares the numeric log levels of the given level and the logger's level.
// Returns true if the given level is enabled (its numeric value is greater than or equal to
// the logger's numeric level value), false otherwise.
// The function allows determining if a log entry with a specific level should be logged
// based on the logger's configured log level.
func (l *Logger) isEnabled(level string) bool {
	levels := map[string]int{
		"fatal": 5,
		"error": 4,
		"warn":  3,
		"info":  2,
		"debug": 1,
	}

	return levels[level] >= levels[l.Level]
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
