package proto

import (
	"context"
	"fmt"
	"strconv"

	dErr "xxx/app/domain/error"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	ErrorKey string
)

const (
	ErrorCtxKey ErrorKey = "error"
)

func SetError(src context.Context, err error) context.Context {
	return context.WithValue(src, ErrorCtxKey, err)
}

func GetError(ctx context.Context) (error, error) {
	v := ctx.Value(ErrorCtxKey)

	err, ok := v.(error)

	if !ok {
		return nil, fmt.Errorf("error not found")
	}

	return err, nil
}

func TranslateError(err *dErr.Error) error {
	switch err.Type {
	case "InvalidArgument":
		status := createStatus(codes.InvalidArgument, err)
		return status.Err()
	case "Unauthenticated":
		status := createStatus(codes.Unauthenticated, err)
		return status.Err()
	case "PermissionDenied":
		status := createStatus(codes.PermissionDenied, err)
		return status.Err()
	case "NotFound":
		status := createStatus(codes.NotFound, err)
		return status.Err()
	case "Aborted":
		status := createStatus(codes.Aborted, err)
		return status.Err()
	case "Internal":
		status := createStatus(codes.Internal, err)
		return status.Err()
	case "DeadlineExceeded":
		status := createStatus(codes.DeadlineExceeded, err)
		return status.Err()
	default:
		return status.Errorf(codes.Unknown, err.Error())
	}
}

func createStatus(code codes.Code, err *dErr.Error) *status.Status {
	st := status.New(code, err.Message)
	details := new(errdetails.BadRequest)
	// status code
	status := &errdetails.BadRequest_FieldViolation{
		Field:       "Status",
		Description: strconv.Itoa(err.Code),
	}
	details.FieldViolations = append(details.FieldViolations, status)
	// reason
	reason := &errdetails.BadRequest_FieldViolation{
		Field:       "Reason",
		Description: err.Reason,
	}
	details.FieldViolations = append(details.FieldViolations, reason)
	// error
	er := &errdetails.BadRequest_FieldViolation{
		Field:       "Error",
		Description: err.Err.Error(),
	}
	details.FieldViolations = append(details.FieldViolations, er)

	detail, _ := st.WithDetails(details)
	return detail
}
