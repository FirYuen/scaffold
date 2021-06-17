package ecode

import "github.com/FirYuen/scaffold/pkg/errno"

var (
	// user errors
	ErrUserNotFound = errno.NewError(20101, "The user was not found.")
	JobError        = errno.NewError(20102, "job error")
)
