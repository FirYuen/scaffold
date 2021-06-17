package service

import (
	"github.com/FirYuen/scaffold/internal/dao"
	"github.com/FirYuen/scaffold/internal/models"
	"github.com/FirYuen/scaffold/pkg/conf"
)


const (
	// DefaultLimit 默认分页数
	DefaultLimit = 50

	// MaxID 最大id
	MaxID = 0xffffffffffff

	// DefaultAvatar 默认头像 key
	DefaultAvatar = "default_avatar.png"
)


var JobSvc *Service

type Service struct {
	c   *conf.Config
	dao *dao.Dao
}
// New init service

func New(c *conf.Config) (s *Service) {
	db := model.GetDB()
	s = &Service{
		c:   c,
		dao: dao.New(db),
	}
	JobSvc =s
	return s
}

// Ping service
func (s *Service) Ping() error {
	return nil
}

// Close service
func (s *Service) Close() {
	s.dao.Close()
}
