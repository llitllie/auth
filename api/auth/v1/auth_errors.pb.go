// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsAuthFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AUTH_FAILED.String() && e.Code == 400
}

func ErrorAuthFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_AUTH_FAILED.String(), fmt.Sprintf(format, args...))
}