package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "the route not found")
}
