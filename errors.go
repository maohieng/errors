// Package errors is inspired from a presentation https://youtu.be/4WIhhzTTd0Y
package errs

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// Op is a unique string describing a
// method or a function.
// Multiple operations can construct a
// friendly stack trace.
type Op string

type Error struct {
	// operation where the error occurs
	Op Op
	// category of errors
	Kind Kind
	// the wrapped error, must not nil
	Err error
	// level of error
	Sev Severity
	//... application specific data
}

func (err *Error) Error() string {
	// this is fast implementation compare below 👇
	// return fmt.Sprintf("%v, %s", Ops(err), Unwrap(err.Err).Error())

	// this new implementation improves
	// - 49% performance
	// - reduce 44% allocation
	var sb strings.Builder
	sb.WriteRune('[')
	sb.WriteString(string(err.Op))

	var e error = err.Err
	for {
		sube, ok := e.(*Error)
		if !ok {
			sb.WriteString("], ")
			sb.WriteString(e.Error())
			break
		}
		e = sube.Err
		sb.WriteRune(' ')
		sb.WriteString(string(sube.Op))
	}

	return sb.String()
}

// New creates an error of Error.
// New is the same as E.
// It's used to make sure an error is provided to avoid nil pointer panic.
// If there's an error or string provided in args, the err will be replaced
// by the args error or an error from the args string.
func New(err error, args ...interface{}) error {
	e := &Error{
		Sev: SevereError(), // default severity
	}
	e.Err = err

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
		//case string:
		//	e.Err = errors.New(arg)
		default:
			panic(fmt.Sprintf("bad call to E. unsupported %v", arg))
		}
	}

	// create invoked frame for no op provided
	if e.Op == "" {
		pc := make([]uintptr, 10) // at least 1 entry needed
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		if len(frame.Function) > 0 {
			spsl := strings.Split(frame.Function, "/")
			lastSp := spsl[len(spsl)-1]
			if lastSp != "errs.SNew" {
				e.Op = Op(lastSp)
			}
		}
	}

	return e
}

func SNew(msg string, args ...interface{}) error {
	var e *Error = New(errors.New(msg), args...).(*Error)

	// create invoked frame for no op provided
	if e.Op == "" {
		pc := make([]uintptr, 10) // at least 1 entry needed
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		if len(frame.Function) > 0 {
			spsl := strings.Split(frame.Function, "/")
			lastSp := spsl[len(spsl)-1]
			e.Op = Op(lastSp)
		}
	}

	return e
}

// E creates an error of Error from args that must be type of
// Op, error, Kind, level.Value or a string of error
//
// Prefer using New or SNew to avoid missing an error providing which
// is required.
// Deprecated: This func is no longer maintained,
// and will remove in the next release.
// func E(args ...interface{}) error {
// 	e := &Error{
// 		Sev: SevereError(), // default severity
// 	}
// 	for _, arg := range args {
// 		switch arg := arg.(type) {
// 		case Op:
// 			e.Op = arg
// 		case error:
// 			e.Err = arg
// 		case Kind:
// 			e.Kind = arg
// 		case Severity:
// 			e.Sev = arg
// 		case string:
// 			e.Err = errors.New(arg)
// 		default:
// 			panic(fmt.Sprintf("bad call to E. unsupported %v", arg))
// 		}
// 	}

// 	return e
// }

// Ops returns the "stack" of operations
// for each generated error.
func Ops(err error) []Op {
	e, ok := err.(*Error)
	if !ok {
		return []Op{}
	}

	res := []Op{e.Op}

	subErr, ok := e.Err.(*Error)

	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}

// Kinds unwraps the last stack error's Kind.
func Kinds(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return KindOfGrpcErr(err)
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
