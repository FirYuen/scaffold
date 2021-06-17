package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/FirYuen/scaffold/api"
	"github.com/FirYuen/scaffold/internal/handler/v1/job"
)

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用中间件
	g.NoRoute(api.RouteNotFound)
	g.NoMethod(api.RouteNotFound)
	apiV1 := g.Group("/v1")
	apiV1.Use()
	{
		// 用户
		apiV1.GET("/start", job.Get)
	}
	return g
}