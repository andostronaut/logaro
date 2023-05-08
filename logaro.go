package logaro

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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

func (l *Logger) Child(fields map[string]interface{}) *Logger {
	child := GenerateLogger()
	child.Level = l.Level
	child.Parent = l
	child.EventFields = l.mergeFields(fields)

	l.Children = append(l.Children, child)

	return child
}

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
