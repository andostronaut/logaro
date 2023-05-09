package examples

import (
	"encoding/json"
	"log"
	"os"

	"github.com/iamando/logaro"
)

func main() {
	// Open the log file for writing
	file, err := os.OpenFile("application.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a logger with the file writer
	logger := logaro.GenerateLogger()
	logger.Writer = json.NewEncoder(file)

	// Log messages to the file
	logger.Log("info", "Logging to file", nil)
	logger.Log("warn", "File warning", nil)
	logger.Log("error", "File error", nil)
}
