package errs

import (
	"errors"
	"os"
	"testing"
)

var errs error

func TestMain(m *testing.M) {
	os.Exit(func() int {
		const op Op = "Op1"
		errs = SNew("Init error", op)

		const op2 Op = "Op2"
		errs = New(errs, op2)
		return m.Run()
	}())
}

func TestWithOps(t *testing.T) {
	t.Log(errs.Error())
}

func TestWithoutOps(t *testing.T) {
	err := SNew("root error msg")
	err = New(err)
	t.Log(err.Error())
}

func TestOps(t *testing.T) {
	t.Log(Ops(errs))
	t.Log(Ops(errors.New("error from errors package")))
}
