package errors

import (
	"google.golang.org/grpc/status"
)

// EncodeGRPCError creates a grpc error from the given error.
// If the given error is known as Error, it will create the
// with err's Kind, otherwise KindInternal will be used.
func EncodeGRPCError(err error) error {
	code := Kinds(err).GRPCCode()
	st := status.New(code, err.Error())
	return st.Err()
}
