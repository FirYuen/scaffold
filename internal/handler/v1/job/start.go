package job

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/FirYuen/scaffold/internal/dao"
	"github.com/FirYuen/scaffold/internal/ecode"
	"github.com/FirYuen/scaffold/internal/service"
	"github.com/FirYuen/scaffold/pkg/errno"
	"github.com/FirYuen/scaffold/pkg/log"
)

// Get 获取用户信息
// @Summary 通过用户id获取用户信息
// @Description Get an user by user id
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /users/:id [get]
func Get(c *gin.Context) {
	log.Info("Get function called.")


	err:=service.JobSvc.StartJob(c.Request.Context())
	if errors.Is(err, dao.ErrNotFound) {
		log.Errorf("get user info err: %+v", err)
		response.Error(c, ecode.JobError)
		return
	}
	if err != nil {
		response.Error(c, errno.ErrInternalServerError)
		return
	}

	response.Success(c,"ok")

}
