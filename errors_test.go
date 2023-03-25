package errs

import (
	"errors"
	"fmt"
	"testing"
)

// sNew returns an error by getting default op as func name
// if op is not provided.
func sNew(op Op) error {
	return SNew("error of string", op)
}

// eNew returns an error by getting default op as func name
// if op is not provided.
func eNew(op Op) error {
	return New(errors.New("error of new"), op)
}

// stackErrs creates a stack of 10 errors
func stackErrs() error {
	var err = SNew("error origin", Op("op origin"))
	for i := 1; i < 10; i++ {
		//var msg = fmt.Sprintf("error #%s", i)
		var op Op = Op(fmt.Sprintf("op #%d", i))
		err = New(err, op, Kind(i))
	}

	return err
}

func TestNewWithOps(t *testing.T) {
	e1 := eNew(Op("eOp"))
	t.Log(e1.Error())
	e2 := sNew(Op("sOp"))
	t.Log(e2.Error())
	t.Log(New(New(e1, Op("Op1")), Op("0p2")).Error())
}

func TestNewWithoutOps(t *testing.T) {
	err := SNew("root error msg")
	err = New(err)

	t.Log(err.Error())
	t.Log(New(New(New(err))).Error())
}

func TestOps(t *testing.T) {
	t.Log(Ops(stackErrs()))
}

func TestUnwrap(t *testing.T) {
	t.Log(Unwrap(stackErrs()))
}

func TestKinds(t *testing.T) {
	t.Log(Kinds(stackErrs()))
}

func TestPrintError(t *testing.T) {
	t.Log(stackErrs())
}

// “
//
//	go test -bench=Prime -count 12 | tee fmt.txt
//
// “
// “
//
//	benchstat fmt.txt builder.txt
//
// “
func BenchmarkError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := stackErrs()
		err.Error()
	}
}
