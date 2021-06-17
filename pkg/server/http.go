package server

import (
	"github.com/FirYuen/scaffold/internal/routers"
	"github.com/FirYuen/scaffold/internal/service"
	"github.com/FirYuen/scaffold/pkg/conf"
	"github.com/FirYuen/scaffold/pkg/transport/http"
)

// NewHttpServer creates a HTTP server
func NewHttpServer(c *conf.Config, svc *service.Service) *http.Server {
	router := routers.NewRouter()

	var opts []http.ServerOption
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	srv := http.NewServer(opts...)
	srv.Handler = router

	return srv
}
