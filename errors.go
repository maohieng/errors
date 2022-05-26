// Package errors is inspired from a presentation https://youtu.be/4WIhhzTTd0Y
package errs

import (
	"errors"
	"fmt"
)

// Op is a unique string describing a
// method or a function.
// Multiple operations can construct a
// friendly stack trace.
type Op string

type Error struct {
	Op   Op       // operation
	Kind Kind     // category of errors
	Err  error    // the wrapped error
	Sev  Severity // level of error
	//... application specific data
}

func (err *Error) Error() string {
	return fmt.Sprintf("%v, %s", Ops(err), Unwrap(err.Err).Error())
	//return UnwrapErrors(err.Err)
}

//func (err *Error) String() string {
//	return fmt.Sprintf("%v, %s", Ops(err), UnwrapErrors(err.Err))
//}

// E creates an error of Error from args that must be type of
// Op, error, Kind, level.Value
func E(args ...interface{}) error {
	e := &Error{
		Sev: SevereError(), // default severity
	}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case Kind:
			e.Kind = arg
		case Severity:
			e.Sev = arg
		case string:
			e.Err = errors.New(arg)
		default:
			panic("bad call to E")
		}
	}

	return e
}

// Ops returns the "stack" of operations
// for each generated error.
func Ops(e *Error) []Op {
	res := []Op{e.Op}

	subErr, ok := e.Err.(*Error)

	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}

// Kinds unwraps the error and returns the first error's Kind.
func Kinds(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return KindInternal
	}

	if e.Kind != 0 {
		return e.Kind
	}

	return Kinds(e.Err)
}

// Unwrap unwraps the original error.
func Unwrap(err error) error {
	e, ok := err.(*Error)
	if !ok {
		return err
	}

	return Unwrap(e.Err)
}
