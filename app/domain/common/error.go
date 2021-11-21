package common

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	Err     error  `json:"error"`
	Type    string `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Stack   string `json:"stack"`
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) Format() []byte {
	res, _ := json.Marshal(e)
	return res
}

func newError(err error, typ string, code int, msg string, reason string) *Error {
	e := new(Error)
	e.Err = err
	e.Type = typ
	e.Code = code
	e.Message = msg
	e.Reason = reason
	e.Stack = fmt.Sprintf("%+v", errors.WithStack(err))
	return e
}

// 400:クライアントから無効なリクエストが送信されたとき
func NewBadRequestError(err error, reason string) *Error {
	return newError(
		err,
		"InvalidArgument",
		400,
		"BadRequest",
		reason,
	)
}

// 401:認証エラー
func NewUnauthorizedError(err error, reason string) *Error {
	return newError(
		err,
		"Unauthenticated",
		401,
		"Unauthorized",
		reason,
	)
}

// 403:認可エラー
func NewForbiddenError(err error, reason string) *Error {
	return newError(
		err,
		"PermissionDenied",
		403,
		"Forbidden",
		reason,
	)
}

// 404:要素が見つからなかった
func NewNotFoundError(err error, reason string) *Error {
	return newError(
		err,
		"NotFound",
		404,
		"NotFound",
		reason,
	)
}

// 409:コンフリクトが発生
func NewConflictError(err error, reason string) *Error {
	return newError(
		err,
		"Aborted",
		409,
		"Conflict",
		reason,
	)
}

// 500:サーバーエラー
func NewInternalServerError(err error, reason string) *Error {
	return newError(
		err,
		"Internal",
		500,
		"InternalServerError",
		reason,
	)
}

// 504:外部通信タイムアウトエラー
func NewGatewayTimeoutError(err error, reason string) *Error {
	return newError(
		err,
		"DeadlineExceeded",
		504,
		"GatewayTimeout",
		reason,
	)
}
