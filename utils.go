package logaro

import (
	"bytes"
	"encoding/json"
)

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
