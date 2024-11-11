package stackerrors_test

import (
	"fmt"
	"testing"

	"github.com/mnagaa/stackerrors"
)

func TestSampleCode(t *testing.T) {
	err := stackerrors.New("1st error")
	fmt.Printf("%+v", err)

	err2 := stackerrors.Wrap(err, "2nd error")
	fmt.Printf("%+v", err2)
}
