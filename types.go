package logaro

import "encoding/json"

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
