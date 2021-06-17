package job

import "github.com/FirYuen/scaffold/pkg/app"

var response = app.NewResponse()

// CreateRequest 开始任务请求
type startRequest struct {
	Model string `json:"model"`
}
// StartJobResponse 开始任务回应
type CreateResponse struct {
	Model string `json:"model"`
}

