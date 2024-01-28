package errorx

const successCode = 0

const successMsg = "接口请求成功"

const defaultCode = 1001

// 自定义错误代码的错误信息
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 代码错误
func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

// 默认错误代码
func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

// 取消错误代码
func NewCodeCanceledError(msg string) error {
	return &CodeError{Code: 1002, Msg: msg}
}

// 无效参数错误代码
func NewCodeInvalidArgumentError(msg string) error {
	return &CodeError{Code: 1003, Msg: msg}
}

// 找不到资源错误代码
func NewCodeNotFoundError(msg string) error {
	return &CodeError{Code: 1005, Msg: msg}
}

// 已存在错误代码
func NewCodeAlreadyExistsError(msg string) error {
	return &CodeError{Code: 1006, Msg: msg}
}

// 中止错误代码
func NewCodeAbortedError(msg string) error {
	return &CodeError{Code: 1009, Msg: msg}
}

// 内部错误代码
func NewCodeInternalError(msg string) error {
	return &CodeError{Code: 1010, Msg: msg}
}

// 未授权错误代码
func NewCodeUnavailableError(msg string) error {
	return &CodeError{Code: 1012, Msg: msg}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
