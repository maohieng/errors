package errs

import (
	"google.golang.org/grpc/status"
)

// EncodeGRPCError creates a grpc error from the given error.
// If err's type is errs.Error, it will create
// with err's Kind, otherwise KindUnknown will be used.
func EncodeGRPCError(err error) error {
	code := Kinds(err).GRPCCode()
	st := status.New(code, err.Error())
	return st.Err()
}
