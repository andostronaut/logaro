package examples

import (
	"strings"

	"github.com/iamando/logaro"
)

func init() {
	// Create a logger
	logger := logaro.GenerateLogger()

	// Define a custom serializer function for message values
	messageSerializer := func(value interface{}) interface{} {
		// Convert the message to uppercase
		if str, ok := value.(string); ok {
			return strings.ToUpper(str)
		}
		return value
	}

	// Define a custom serializer function for field values
	fieldSerializer := func(value interface{}) interface{} {
		// Append " - Custom Field" to string values
		if str, ok := value.(string); ok {
			return str + " - Custom Field"
		}
		return value
	}

	// Create a child logger with custom serializers
	childLogger := logger.WithSerializers(map[string]func(interface{}) interface{}{
		"message": messageSerializer,
		"field":   fieldSerializer,
	})

	// Log a message with custom serialization
	childLogger.Log("info", "Custom serialization example", map[string]interface{}{
		"field": "Value",
	})
}
