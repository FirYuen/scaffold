package service

import "context"

// 用于触发编译期的接口的合理性检查机制
var _ service = (*Service)(nil)

type service interface {
	StartJob(ctx context.Context)error
}