package stackerrors

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type ErrorFrame struct {
	Message string
	Frames  []string
}

type Error struct {
	msg   string
	stack []ErrorFrame
	cause *Error
}

func New(msg interface{}) *Error {
	var message string

	switch m := msg.(type) {
	case *Error:
		return m
	case error:
		message = m.Error()
	case string:
		message = m
	default:
		message = fmt.Sprintf("%v", m)
	}

	err := &Error{
		msg:   message,
		stack: []ErrorFrame{{Message: message, Frames: []string{getStackTrace(2)}}},
	}
	return err
}

func Wrap(err error, msg string) *Error {
	wrappedErr, ok := err.(*Error)
	if !ok {
		wrappedErr = New(err.Error())
	}

	newErr := &Error{
		msg:   msg,
		stack: append([]ErrorFrame{{Message: msg, Frames: []string{getStackTrace(2)}}}, wrappedErr.stack...),
		cause: wrappedErr,
	}
	return newErr
}

func getStackTrace(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func (err *Error) Error() string {
	if err.cause != nil {
		return fmt.Sprintf("%s: %v", err.msg, err.cause)
	}
	return err.msg
}

func (err *Error) Unwrap() error {
	if err.cause != nil {
		return err.cause
	}
	return nil
}

func (err *Error) Is(target error) bool {
	if targetErr, ok := target.(*Error); ok {
		return err.msg == targetErr.msg
	}
	return false
}

func (err *Error) As(target interface{}) bool {
	if targetErr, ok := target.(*Error); ok {
		*targetErr = *err
		return true
	}
	return false
}

func (err *Error) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('+') {
		fmt.Fprint(f, err.StackTrace())
		return
	}
	fmt.Fprint(f, err.Error())
}

func (err *Error) StackTrace() string {
	var builder strings.Builder
	builder.WriteString("Error Stack Trace: ")

	appendErrorMessages(&builder, err)
	builder.WriteString("\n")

	seenFrames := make(map[*ErrorFrame]map[string]bool)
	appendFilteredStackTrace(&builder, err.stack, seenFrames)
	return builder.String()
}

func appendErrorMessages(builder *strings.Builder, err *Error) {
	builder.WriteString(err.msg)

	if err.cause != nil {
		builder.WriteString(" -> ")
		appendErrorMessages(builder, err.cause)
	}
}

func appendFilteredStackTrace(builder *strings.Builder, frames []ErrorFrame, seenFrames map[*ErrorFrame]map[string]bool) {
	for i := range frames {
		frame := &frames[i]
		if seenFrames[frame] == nil {
			seenFrames[frame] = make(map[string]bool)
		}
		builder.WriteString("└── Error: " + frame.Message + "\n")

		for _, trace := range frame.Frames {
			if !seenFrames[frame][trace] {
				seenFrames[frame][trace] = true
				builder.WriteString("    └── " + trace + "\n")
			}
		}
	}
}

type ErrorJSON struct {
	Message string   `json:"message"`
	Frames  []string `json:"frames"`
}

func (err *Error) MarshalJSON() ([]byte, error) {
	var frames []string
	for _, frame := range err.stack {
		frames = append(frames, frame.Frames...)
	}

	errJSON := &ErrorJSON{
		Message: err.Error(),
		Frames:  frames,
	}

	return json.Marshal(errJSON)
}
