package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// Kind groups all errors into smaller categories.
// Can be predefined codes (http / grpc).
// Or it can be your own defined codes.
type Kind uint32

type TransportCode interface {
	// HTTPCode converts its implementation type into HTTP status code.
	HTTPCode() int
	// GRPCCode converts its implementation type into GRPC status code.
	GRPCCode() codes.Code
}

// All constant values must be the same as grpc's Code
const (
	KindBadRequest    Kind = 3
	KindNotFound      Kind = 5
	KindUnauthorized  Kind = 16
	KindNotAllowed    Kind = 7
	KindInternal      Kind = 13
	KindUnknown       Kind = 2
	KindAlreadyExists Kind = 6
)

func (k Kind) HTTPCode() int {
	switch k {
	case KindBadRequest:
		return http.StatusBadRequest
	case KindAlreadyExists:
		return http.StatusBadRequest
	case KindNotFound:
		return http.StatusNotFound
	case KindUnauthorized:
		return http.StatusUnauthorized
	case KindNotAllowed:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

func (k Kind) GRPCCode() codes.Code {
	return codes.Code(k)
}

// KindOfHTTPStatus converts HTTP's status code into Kind
func KindOfHTTPStatus(code int) Kind {
	switch code {
	case http.StatusBadRequest:
		return KindBadRequest
	case http.StatusNotFound:
		return KindNotFound
	case http.StatusUnauthorized:
		return KindUnauthorized
	case http.StatusForbidden:
		return KindNotFound
	default:
		return KindInternal
	}
}

func KindOfGRPCCode(code codes.Code) Kind {
	return Kind(code)
}

func KindOfGrpcErr(err error) Kind {
	c := status.Code(err)
	return Kind(c)
}

// KindOfGRPCError returns a Kind representing error if it was produced
// by status.FromError function.
// Otherwise, ok is false and Kind is KindInternal returned.
//func KindOfGRPCError(err error) (Kind, bool) {
//	s, ok := status.FromError(err)
//	if ok {
//		return KindOfGRPC(s.Code()), true
//	}
//
//	return KindInternal, false
//}
