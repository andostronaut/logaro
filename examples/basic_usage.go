package examples

import "github.com/iamando/logaro"

func init() {
	// Create a logger
	logger := logaro.GenerateLogger()

	// Log messages at different levels
	logger.Log("info", "This is an informational message", nil)
	logger.Log("warn", "This is a warning message", nil)
	logger.Log("error", "This is an error message", nil)

	// Log a message with additional fields
	logger.Log("debug", "Debugging information", map[string]interface{}{
		"user_id":    123,
		"request_id": "abc123",
	})
}
