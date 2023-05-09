package examples

import "github.com/iamando/logaro"

func init() {
	// Create a root logger
	rootLogger := logaro.GenerateLogger()

	// Log a message from the root logger
	rootLogger.Log("info", "Root logger message", nil)

	// Create a child logger with additional fields
	childLogger := rootLogger.WithFields(map[string]interface{}{
		"component": "subsystemA",
	})

	// Log a message from the child logger
	childLogger.Log("info", "Child logger message", nil)

	// Create another child logger
	anotherChildLogger := rootLogger.WithFields(map[string]interface{}{
		"component": "subsystemB",
	})

	// Log a message from the another child logger
	anotherChildLogger.Log("info", "Another child logger message", nil)
}
