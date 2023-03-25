package errs

import (
	"encoding/json"
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

// stackErrs creates a stack of errors
func stackErrs() error {
	var err = errors.New("msg origin")
	for i := 1; i < 10; i++ {
		//var msg = fmt.Sprintf("error #%s", i)
		var op Op = Op(fmt.Sprintf("op.%d", i))
		if i == 5 {
			err = New(err, op, Kind(i))
		} else {
			err = New(err, op, Kind(i), fmt.Sprintf("err at %d", i))
		}
	}

	return err
}

// “
//
//	go test -bench=Error -benchmem -count 12 | tee builder2.txt
//
// “
// “
//
//	benchstat fmt.txt builder2.txt
//
// “
func BenchmarkError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := stackErrs()
		err.Error()
	}
}

func TestNewWithOps(t *testing.T) {
	e1 := eNew(Op("e.Op"))
	t.Log(e1)
	e2 := sNew(Op("s.Op"))
	t.Log(e2)
	t.Log(New(New(e1, Op("Op.1")), Op("Op.2")))
}

func TestNewWithoutOps(t *testing.T) {
	err1 := sNew("")
	err2 := eNew("")

	t.Log(err1)
	t.Log(err2)
	t.Log(New(New(New(err1))))
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

func TestPrintGeneralUse(t *testing.T) {
	err := errors.New("database error")
	err = New(err, Op("persist.Create"), "unable to create record")
	err = New(err, Op("svc.Create"), KindInternal)
	t.Log(err)
}

func TestStackJSON(t *testing.T) {
	stack := Errors(stackErrs())
	ststr, err := json.Marshal(stack)
	if err != nil {
		t.Fatalf("Expect success, got %v", err)
	}

	t.Log(string(ststr))
}
