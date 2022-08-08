package errs

import "testing"

func errInit() error {
	return SNew("Init error")
}

func TestWithOps(t *testing.T) {
	const op Op = "testing.WithOp1"
	err := SNew("Init error", op)

	const op2 Op = "testing.WithOp2"
	err = New(err, op2)

	t.Log(err.Error())
}

func TestWithoutOps(t *testing.T) {
	err := errInit()

	err = New(err)

	t.Log(err.Error())
}
