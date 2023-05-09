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
 logger.Log("info", "Hello, Logaro!")

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
 childLogger.Log("debug", "Debug message from subsystemA")
}
```

For more detailed examples and advanced usage, please refer to the [examples](/examples) directory.

## Documentation

For detailed documentation and API reference, please refer to the [GoDoc](https://godoc.org/github.com/iamando/logaro) page.

## Support

Logaro is an MIT-licensed open source project. It can grow thanks to the sponsors and support.

## License

Logaro is [MIT licensed](LICENSE).
