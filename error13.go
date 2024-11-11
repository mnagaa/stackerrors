package stackerrors

import (
	"errors"
)

func As(err error, target interface{}) bool { return errors.As(err, target) }

func Is(e error, original error) bool { return errors.Is(e, original) }
