// Borrow from go-kit/log's level

package errors

type Severity interface {
	String() string
	levelVal()
}

// SevereError returns the unique value added to error events by Error.
func SevereError() Severity { return errorValue }

// SevereWarn returns the unique value added to error events by Warn.
func SevereWarn() Severity { return warnValue }

type level byte

const (
	levelWarn level = 3 << iota
	levelError
)

type levelValue struct {
	name string
	level
}

var (
	errorValue = &levelValue{level: levelError, name: "error"}
	warnValue  = &levelValue{level: levelWarn, name: "warn"}
)

func (v *levelValue) String() string { return v.name }
func (v *levelValue) levelVal()      {}
