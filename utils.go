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

// mergeFields merges the event fields from the parent logger with the current logger's event fields
// and the additional fields provided as a parameter.
// It creates a new map to hold the merged fields and copies the parent's event fields into it.
// Then it adds the current logger's event fields and the additional fields to the merged map.
// Returns the merged map of event fields, combining the inherited fields from the parent logger
// with the current logger's fields and the additional fields provided as a parameter.
// The function is used to create a consolidated set of event fields for log entries,
// ensuring that all relevant fields are included in the log context.
func (l *Logger) mergeFields(fields map[string]interface{}) map[string]interface{} {
	mergedFields := make(map[string]interface{})

	if l.Parent != nil {
		mergedFields = l.Parent.mergeFields(l.Parent.EventFields)
	}

	for key, val := range l.EventFields {
		mergedFields[key] = val
	}

	for key, val := range fields {
		mergedFields[key] = val
	}

	return mergedFields
}

// serializeEntry applies the logger's serializer function to the log entry.
// If a serializer is set for the logger, it applies the serializer function to the log message
// and the fields of the entry, allowing custom modification or formatting of the log entry.
// Returns the serialized log entry with the log message and fields modified by the serializer,
// ensuring that the log data is transformed according to the specified serialization logic.
// The function is used to customize the serialization process for specific log entries
// based on the serializer function set for the logger.
func (l *Logger) serializeEntry(entry LogEntry) LogEntry {
	if l.Serializer != nil {
		entry.Message = l.Serializer(entry.Message).(string)
		entry.Fields = l.Serializer(entry.Fields).(map[string]interface{})
	}

	return entry
}

// compareLogEntries compares two log entries for equality.
// It performs a field-level comparison of the log entries, checking the equality of the
// Timestamp, Message, Level, and Fields. Returns true if the log entries are equal,
// and false otherwise. The function is used to validate the correctness of captured
// log entries by comparing them against expected log entries in test cases.
// It helps ensure that the logged information, including the timestamp, log message,
// log level, and additional fields, is matching the expected values.
func compareLogEntries(a, b LogEntry) bool {
	return a.Timestamp == b.Timestamp &&
		a.Message == b.Message &&
		a.Level == b.Level &&

		compareFields(a.Fields, b.Fields)
}

// compareFields compares two maps of log fields for equality.
// It performs a key-value comparison of the fields, checking the equality of both keys and values.
// Returns true if the field maps are equal, and false otherwise. The function is used as a helper
// function in comparing log entries, specifically the fields section, to validate the correctness
// of the captured log entries against expected log entries in test cases. It ensures that the
// field maps contain the same keys with matching values, confirming that the logged fields are
// consistent and accurate in the captured log entry.
func compareFields(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for key, valA := range a {
		valB, ok := b[key]
		if !ok || !compareFieldValues(valA, valB) {
			return false
		}
	}

	return true
}

// compareFieldValues compares two field values for equality.
// It marshals the values to JSON and compares their byte representations.
// Returns true if the field values are equal, and false otherwise. The function is used
// as a helper function in comparing log entries to validate the correctness of the captured
// log entries against expected log entries in test cases. It ensures that the field values
// are consistent and accurate in the captured log entry by comparing their JSON
// representations. If the marshaling fails or the byte representations differ, the function
// returns false, indicating a mismatch between the field values.
func compareFieldValues(a, b interface{}) bool {
	bytesA, errA := json.Marshal(a)
	bytesB, errB := json.Marshal(b)

	if errA != nil || errB != nil {
		return false
	}

	return bytes.Equal(bytesA, bytesB)
}
