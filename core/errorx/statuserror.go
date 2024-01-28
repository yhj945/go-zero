package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 内部错误
func NewInternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}

// 参数无效错误
func NewInvalidArgumentError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

// 找不到资源错误
func NewNotFoundError(msg string) error {
	return status.Error(codes.NotFound, msg)
}

// 已存在错误
func NewAlreadyExistsError(msg string) error {
	return status.Error(codes.AlreadyExists, msg)
}

// 未授权错误
func NewUnauthenticatedError(msg string) error {
	return status.Error(codes.Unauthenticated, msg)
}

// 资源耗尽错误
func NewResourceExhaustedError(msg string) error {
	return status.Error(codes.ResourceExhausted, msg)
}
