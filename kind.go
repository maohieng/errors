package errs

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Kind groups all errors into smaller categories.
// Can be predefined codes (http / grpc).
// Or it can be your own defined codes.
type Kind uint32

//type TransportCode interface {
//	// HTTPCode converts its implementation type into HTTP status code.
//	HTTPCode() int
//	// GRPCCode converts its implementation type into GRPC status code.
//	GRPCCode() codes.Code
//}

// All constant values must be the same as grpc's Code
const (
	KindUnknown       Kind = 2
	KindBadRequest    Kind = 3
	KindNotFound      Kind = 5
	KindAlreadyExists Kind = 6
	KindNotAllowed    Kind = 7
	KindUnauthorized  Kind = 16
	KindInternal      Kind = 13
)

func (k Kind) HTTPCode() int {
	switch k {
	case KindBadRequest, KindUnknown, KindAlreadyExists:
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

// KindOfGrpcErr returns a Kind from the given gRPC code
// if it is a Status error.
// KindOk if err is nil, KindUnknown otherwise.
// See status.Code for more info.
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
