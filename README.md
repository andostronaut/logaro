# Logaro

![build](https://github.com/iamando/logaro/workflows/build/badge.svg)
![license](https://img.shields.io/github/license/iamando/logaro?color=success)
![Go version](https://img.shields.io/github/go-mod/go-version/iamando/logaro)
[![GoDoc](https://godoc.org/github.com/iamando/logaro?status.svg)](https://godoc.org/github.com/iamando/logaro)

`Logaro` is a lightweight Go package for JSON-based logging. It provides a simple and flexible logging solution with support for log levels, log entry customization, and hierarchical loggers.

## Features

- JSON-based logging format for easy log consumption and analysis.
- Configurable log levels to control the verbosity of the log output.
- Ability to customize log entries with additional fields.
- Hierarchical loggers to organize and inherit log settings.
- Support for serializers to transform log message and field values.

## Installation

To use Logaro in your Go project, you need to have Go installed and set up. Then, run the following command to install the package:

```bash
go get github.com/iamando/logaro
```

## Usage

Here's a basic example of how to use Logaro:

```go
package main

import "github.com/your-username/logaro"


func main() {
 // Create a logger
 logger := logaro.GenerateLogger()

 // Log a message
 logger.Log("info", "Hello, Logaro!", nil)

 // Log a message with additional fields
 logger.Log("error", "An error occurred", map[string]interface{}{
  "error_code": 500,
  "error_msg":  "Internal Server Error",
 })

 // Create a child logger with additional fields
 childLogger := logger.WithFields(map[string]interface{}{
  "component": "subsystemA",
 })

 // Log from the child logger
 childLogger.Log("debug", "Debug message from subsystemA", nil)
}
```

For more detailed examples and advanced usage, please refer to the [examples](/examples) directory.

## API

### `type Logger`

Represents a logger instance that can be used to log messages at different levels.

#### Methods

- `Log(level string, message string, fields map[string]interface{})`
  Logs a message at the specified log `level`. Additional `fields` can be provided as a map of key-value pairs.

- `Child(fields map[string]interface{}) *Logger`
  Creates a child logger with the specified additional `fields`. The child logger inherits the log settings and fields from its parent.
  ssss
- `WithFields(fields map[string]interface{}) *Logger`
  Creates a child logger with the specified additional `fields`. The child logger inherits the log settings and fields from its parent.

- `WithSerializers(serializers map[string]func(interface{}) interface{}) *Logger`
  Creates a child logger with custom `serializers` for transforming log message and field values. The serializers argument should be a map where the keys represent the fields to be serialized and the values are functions that perform the serialization.

### `type LogEntry`

Represents a log entry containing the log message, log level, and additional fields.

#### Fields

- `Timestamp string`
  The timestamp of the log entry in RFC3339 format.

- `Message string`
  The log message.

- `Level string`
  The log level.

- `Fields map[string]interface{}`
  Additional fields associated with the log entry.

## Documentation

For detailed documentation and API reference, please refer to the [GoDoc](https://godoc.org/github.com/iamando/logaro) page.

## Support

Logaro is an MIT-licensed open source project. It can grow thanks to the sponsors and support.

## License

Logaro is [MIT licensed](LICENSE).
