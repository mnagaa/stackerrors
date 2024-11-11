package stackerrors

import (
	"testing"
)

func TestAs(t *testing.T) {
	var targetErr *Error
	customErr := &Error{msg: "custom error"}
	wrappedErr := Wrap(customErr, "wrapped error")

	if !As(wrappedErr, &targetErr) {
		t.Error("expected As to succeed in finding the target error type (CustomError), but it did not")
		return
	}
}

func TestIs(t *testing.T) {
	rootErr := New("root error")
	wrappedErr := Wrap(rootErr, "wrapped error")

	if !Is(wrappedErr, rootErr) {
		t.Error("expected Is to find the root error in wrapped error, but it did not")
		return
	}

	otherErr := New("another error")
	if Is(wrappedErr, otherErr) {
		t.Error("expected Is not to match with unrelated error, but it did")
		return
	}
}
