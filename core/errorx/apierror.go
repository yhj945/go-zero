package errorx

import "net/http"

type ApiMsg struct {
	Msg string `json:"msg"`
}

// 接口错误信息
type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ApiError) Error() string {
	return e.Msg
}

// 错误信息
func NewApiError(code int, msg string) error {
	return &ApiError{Code: code, Msg: msg}
}

// 错误信息，不包含msg
func NewApiErrorWithoutMsg(code int) error {
	return &ApiError{Code: code, Msg: ""}
}

// 内部错误信息
func NewApiInternalError(msg string) error {
	return &ApiError{Code: http.StatusInternalServerError, Msg: msg}
}

// 请求错误信息
func NewApiBadRequestError(msg string) error {
	return &ApiError{Code: http.StatusBadRequest, Msg: msg}
}

// 未授权错误信息
func NewApiUnauthorizedError(msg string) error {
	return &ApiError{Code: http.StatusUnauthorized, Msg: msg}
}

// 禁止错误信息
func NewApiForbiddenError(msg string) error {
	return &ApiError{Code: http.StatusForbidden, Msg: msg}
}

// 找不到资源错误信息
func NewApiNotFoundError(msg string) error {
	return &ApiError{Code: http.StatusNotFound, Msg: msg}
}

// 网关错误信息
func NewApiBadGatewayError(msg string) error {
	return &ApiError{Code: http.StatusBadGateway, Msg: msg}
}
