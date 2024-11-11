package stackerrors_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/mnagaa/sandbox/stackerrors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := stackerrors.New("test message")
	assert.Equal(t, "test message", err.Error())
	assert.NotEmpty(t, err.StackTrace())
}

func TestWrap(t *testing.T) {
	originalErr := stackerrors.New("original error")
	wrappedErr := stackerrors.Wrap(originalErr, "wrapped error")

	assert.Equal(t, "wrapped error: original error", wrappedErr.Error())
	assert.Contains(t, wrappedErr.StackTrace(), "Error: wrapped error")
	assert.Contains(t, wrappedErr.StackTrace(), "Error: original error")
}

func TestUnwrap(t *testing.T) {
	// Case 1: Error with no cause (Unwrap should return nil)
	originalErr := stackerrors.New("original error")
	assert.Nil(t, originalErr.Unwrap())
}

func TestIs(t *testing.T) {
	originalErr := stackerrors.New("original error")
	wrappedErr := stackerrors.Wrap(originalErr, "wrapped error")

	assert.True(t, errors.Is(wrappedErr, originalErr))
	assert.False(t, errors.Is(wrappedErr, errors.New("different error")))
}

func TestFormat(t *testing.T) {
	err := stackerrors.New("test error")
	formatted := fmt.Sprintf("%v", err)
	assert.Equal(t, "test error", formatted)

	formattedWithStack := fmt.Sprintf("%+v", err)
	assert.Contains(t, formattedWithStack, "Error Stack Trace:")
	assert.Contains(t, formattedWithStack, "Error: test error")
}

func TestStackTrace(t *testing.T) {
	originalErr := stackerrors.New("original error")
	wrappedErr := stackerrors.Wrap(originalErr, "wrapped error")

	stackTrace := wrappedErr.StackTrace()
	assert.Contains(t, stackTrace, "Error Stack Trace: wrapped error -> original error")
	assert.Contains(t, stackTrace, "└── Error: wrapped error")
	assert.Contains(t, stackTrace, "└── Error: original error")
}

func TestMarshalJSON(t *testing.T) {
	err := stackerrors.New("json error")
	jsonData, jsonErr := json.Marshal(err)

	assert.NoError(t, jsonErr)

	var unmarshalledData map[string]interface{}
	assert.NoError(t, json.Unmarshal(jsonData, &unmarshalledData))

	assert.Equal(t, "json error", unmarshalledData["message"])
	assert.NotEmpty(t, unmarshalledData["frames"])
}

func TestAsWithDifferentType(t *testing.T) {
	originalErr := stackerrors.New("original error")
	wrappedErr := stackerrors.Wrap(originalErr, "wrapped error")

	var differentErr *stackerrors.Error
	assert.True(t, errors.As(wrappedErr, &differentErr))
	assert.Equal(t, wrappedErr, differentErr)
}

func TestUnwrapWithMultipleLevels(t *testing.T) {
	err1 := stackerrors.New("root error")
	err2 := stackerrors.Wrap(err1, "middle error")
	err3 := stackerrors.Wrap(err2, "top error")

	assert.Equal(t, err2, errors.Unwrap(err3))
	assert.Equal(t, err1, errors.Unwrap(err2))
}

func TestStackTraceWithoutWrap(t *testing.T) {
	err := stackerrors.New("standalone error")
	stackTrace := err.StackTrace()

	assert.Contains(t, stackTrace, "Error Stack Trace:")
	assert.Contains(t, stackTrace, "└── Error: standalone error")
}

func TestEqualErrors(t *testing.T) {
	err1 := stackerrors.New("error message")
	err2 := stackerrors.New("error message")
	assert.NotEqual(t, err1, err2)
}

func TestStackTraceFramesUniqueness(t *testing.T) {
	err := stackerrors.New("initial error")
	wrappedErr := stackerrors.Wrap(err, "wrapped once")
	reWrappedErr := stackerrors.Wrap(wrappedErr, "wrapped twice")

	stackTrace := reWrappedErr.StackTrace()
	assert.Contains(t, stackTrace, "Error: wrapped twice")
	assert.Contains(t, stackTrace, "Error: wrapped once")
	assert.Contains(t, stackTrace, "Error: initial error")
}
