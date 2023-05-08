package logaro

import (
	"encoding/json"
	"os"
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
