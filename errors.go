// Package errors is inspired from a presentation https://youtu.be/4WIhhzTTd0Y
package errs

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Op is a unique string describing a
// method or a function.
// Multiple operations can construct a
// friendly stack trace.
type Op string

type Error struct {
	msg string
	// operation where the error occurs
	op Op
	// category of errors
	kind Kind
	// the wrapped error, must not nil
	err error
	// level of error
	sev Severity
	//... application specific data
}

// New creates an error of Error.
// New is the same as E.
// It's used to make sure an error is provided to avoid nil pointer panic.
// If there's an error or string provided in args, the err will be replaced
// by the args error or an error from the args string.
func New(err error, args ...interface{}) error {
	e := &Error{
		sev: SevereError(), // default severity
	}
	e.err = err

	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.op = arg
		case error:
			e.err = arg
		case Kind:
			e.kind = arg
		case Severity:
			e.sev = arg
		case string:
			e.msg = arg
		default:
			panic(fmt.Sprintf("bad call to E. unsupported %v", arg))
		}
	}

	// create invoked frame for no op provided
	if e.op == "" {
		pc := make([]uintptr, 10) // at least 1 entry needed
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		if len(frame.Function) > 0 {
			spsl := strings.Split(frame.Function, "/")
			lastSp := spsl[len(spsl)-1]
			if lastSp != "errs.SNew" {
				e.op = Op(lastSp)
			}
		}
	}

	return e
}

func SNew(msg string, args ...interface{}) error {
	var e *Error = New(errors.New(msg), args...).(*Error)

	// create invoked frame for no op provided
	if e.op == "" {
		pc := make([]uintptr, 10) // at least 1 entry needed
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		if len(frame.Function) > 0 {
			spsl := strings.Split(frame.Function, "/")
			lastSp := spsl[len(spsl)-1]
			e.op = Op(lastSp)
		}
	}

	return e
}

func (err *Error) Error() string {
	// The goal of error msg is for:
	// 1. descriptive msg to debugger to understand
	//	 where the cause come from
	// 2. minimized for client be able to capture the
	//   last error msg from the response body,
	//   but not recommended as it not support
	//   the localization msg for now.

	// return fmt.Sprintf("%v, %s", Ops(err), Unwrap(err.Err).Error())

	var sb strings.Builder
	sb.WriteString(string(err.op))
	if err.msg != "" {
		sb.WriteRune(' ')
		sb.WriteString(err.msg)
	}

	var k Kind = err.kind

	var e error = err.err
	for {
		sube, ok := e.(*Error)
		if !ok {
			break
		}
		e = sube.err
		sb.WriteRune(':')
		sb.WriteRune(' ')
		sb.WriteString(string(sube.op))
		if sube.msg != "" {
			sb.WriteRune(' ')
			sb.WriteString(sube.msg)
		}
		if k == 0 {
			k = sube.kind
		}
	}

	sb.WriteString(", ")
	sb.WriteString(e.Error())
	if k != 0 {
		sb.WriteString(", ")
		sb.WriteString("code ")
		sb.WriteString(strconv.Itoa(int(k)))
	}

	return sb.String()
}

// Ops returns the "stack" of operations
// for each generated error.
func Ops(err error) []Op {
	e, ok := err.(*Error)
	if !ok {
		return []Op{}
	}

	res := []Op{e.op}

	subErr, ok := e.err.(*Error)

	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}

// Kinds unwraps the top stack error's Kind.
func Kinds(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return KindOfGrpcErr(err)
	}

	if e.kind != 0 {
		return e.kind
	}

	return Kinds(e.err)
}

// Unwrap unwraps the original error.
func Unwrap(err error) error {
	e, ok := err.(*Error)
	if !ok {
		return err
	}

	return Unwrap(e.err)
}
