package logaro

import "encoding/json"

// Logger represents a JSON logger instance.
// It encapsulates the configuration and functionality of a logging instance.
// - Level: the log level for the logger.
// - Writer: the JSON encoder used for writing log entries.
// - Parent: the parent logger if this logger is a child logger.
// - Children: a list of child loggers created from this logger.
// - EventFields: additional fields associated with each log entry from this logger.
// - Serializer: a function for serializing log entry messages and fields (optional).
type Logger struct {
	Level       string
	Writer      *json.Encoder
	Parent      *Logger
	Children    []*Logger
	EventFields map[string]interface{}
	Serializer  func(data interface{}) interface{}
}

type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Message   string                 `json:"message"`
	Level     string                 `json:"level"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}
