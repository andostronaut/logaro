package logaro

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// GenerateLogger creates a new logger instance with the specified configuration.
// It initializes a Logger struct with the provided log level and output writer.
// - level: the log level for the logger.
// - writer: the JSON encoder used for writing log entries.
// Returns the newly created logger instance.
// The logger is initialized with default settings for parent, children, event fields, and serializer.
func GenerateLogger() *Logger {
	return &Logger{
		Level:       "info",
		Writer:      json.NewEncoder(os.Stdout),
		Parent:      nil,
		Children:    make([]*Logger, 0),
		EventFields: make(map[string]interface{}),
		Serializer:  nil,
	}
}

// Log logs a message at the specified level, along with optional additional fields.
// It first checks if the logger is enabled for the given log level.
// If the logger is enabled, it constructs a LogEntry with a timestamp, log message,
// log level, and merged event fields (including additional fields passed as parameters).
// If a serializer is set for the logger, it applies the serializer function to the log entry
// before encoding. The serializer can be used to customize the formatting or modify the log entry.
// Finally, it encodes the log entry using the logger's JSON encoder and writes it to the output.
// If there is an error encoding the log entry, an error message is printed to the standard output.
func (l *Logger) Log(level, message string, fields map[string]interface{}) {
	if l.isEnabled(level) {
		entry := LogEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Message:   message,
			Level:     level,
			Fields:    l.mergeFields(fields),
		}

		if l.Serializer != nil {
			entry = l.serializeEntry(entry)
		}

		err := l.Writer.Encode(entry)
		if err != nil {
			fmt.Println("Error encoding log entry:", err)

			return
		}
	}
}

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

// Child creates a child logger with additional event fields.
// It creates a new logger instance as a child of the current logger.
// The child logger inherits the log level and event fields from the parent logger.
// Additionally, it adds the specified additional event fields to its own fields.
// Returns the newly created child logger that can be used for logging with added context.
// Child loggers provide a hierarchical structure to organize and filter log entries.
func (l *Logger) Child(fields map[string]interface{}) *Logger {
	child := GenerateLogger()
	child.Level = l.Level
	child.Parent = l
	child.EventFields = l.mergeFields(fields)

	l.Children = append(l.Children, child)

	return child
}

// WithFields creates a child logger with additional event fields.
// It creates a new child logger instance based on the current logger.
// The child logger inherits the log level from the parent logger.
// Additionally, it adds the specified additional event fields to its own fields.
// Returns the newly created child logger with the added event fields.
// The child logger can be used to log with the inherited log level and the additional context
// provided by the added event fields, allowing more specific and detailed log entries.
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	child := l.Child(fields)
	child.Serializer = l.Serializer

	return child
}

func (l *Logger) WithSerializers(serializers map[string]func(interface{}) interface{}) *Logger {
	child := l.Child(nil)

	child.Serializer = func(data interface{}) interface{} {
		for key, serializer := range serializers {
			if val, ok := data.(map[string]interface{})[key]; ok {
				data.(map[string]interface{})[key] = serializer(val)
			}
		}

		return data
	}

	return child
}

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

func (l *Logger) serializeEntry(entry LogEntry) LogEntry {
	if l.Serializer != nil {
		entry.Message = l.Serializer(entry.Message).(string)
		entry.Fields = l.Serializer(entry.Fields).(map[string]interface{})
	}

	return entry
}
