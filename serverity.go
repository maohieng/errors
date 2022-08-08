package errs

type level byte

const (
	levelWarn level = 3 << iota
	levelError
	levelPanic
)

// Severity is borrowed from go-kit/log's level
type Severity interface {
	String() string
	levelVal()
}

// SevereError returns the unique value added to error events by Error.
func SevereError() Severity { return errorValue }

// SevereWarn returns the unique value added to error events by Warn.
func SevereWarn() Severity { return warnValue }

// SeverePanic returns the unique value added to error events by Panic.
func SeverePanic() Severity { return panicValue }

type levelValue struct {
	name string
	level
}

var (
	errorValue = &levelValue{level: levelError, name: "error"}
	// warnValue is deprecated
	// Deprecated: no more usage and will be removed in the next release
	warnValue  = &levelValue{level: levelWarn, name: "warn"}
	panicValue = &levelValue{level: levelPanic, name: "panic"}
)

func (v *levelValue) String() string { return v.name }
func (v *levelValue) levelVal()      {}
