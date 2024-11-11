# stackerrors

`stackerrors` is a Go package designed for advanced error handling and tracing. It allows you to capture error stacks, link error chains, and output detailed stack traces, making complex error flows easier to debug. Additionally, errors can be serialized to JSON format for easier integration with logging or API responses.

## Features

- Captures **error messages and stack traces** for in-depth error handling.
- Supports **wrapping (Wrap)** existing errors with additional context, preserving the stack trace.
- Serializes errors to **JSON format**.
- Extends the `Error` interface, enabling hierarchical error tracing and formatted output for easy viewing.

## Installation

```bash
go get github.com/mnagaa/stackerrors
```

## Usage

### 1. Creating a Basic Error

```go
import "github.com/mnagaa/stackerrors"

func main() {
    err := stackerrors.New("An error occurred")
    fmt.Println(err)
}
```

### 2. Wrapping Errors with Additional Context

You can wrap an existing error with additional context, adding a new stack frame to the trace.

```go
func main() {
    err := stackerrors.New("An error occurred")
    wrappedErr := stackerrors.Wrap(err, "Additional context")
    fmt.Println(wrappedErr)
}
```

### 3. Displaying the Stack Trace

Use `+v` format to view a detailed stack trace, providing insights into the error's origin and propagation.

```go
func main() {
    err := stackerrors.New("An error occurred")
    wrappedErr := stackerrors.Wrap(err, "Additional context")
    fmt.Printf("%+v\n", wrappedErr)
}
```

### 4. Serializing Errors to JSON

Errors can be serialized to JSON, making them easy to use in logs or as part of an API response.

```go
func main() {
    err := stackerrors.New("An error occurred")
    jsonErr, _ := json.Marshal(err)
    fmt.Println(string(jsonErr))
}
```

### 5. Nested Error Chain Display

Errors are displayed hierarchically, making it easy to track the flow of nested error chains.

```go
func main() {
    err := stackerrors.New("Initial error")
    wrappedErr := stackerrors.Wrap(err, "Intermediate context")
    finalErr := stackerrors.Wrap(wrappedErr, "Final context")
    fmt.Printf("%+v\n", finalErr)
}
```

## Methods for the Error Type

- **`Error()`**: Returns the error message.
- **`Unwrap()`**: Returns the original, wrapped error.
- **`Is(target error)`**: Checks if the error matches a specific target error.
- **`As(target interface{})`**: Checks if the error can be cast to a specified type.
- **`StackTrace()`**: Returns the full stack trace along with error messages.

## License

This project is licensed under the MIT License.
