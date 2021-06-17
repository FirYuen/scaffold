package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/FirYuen/scaffold/internal/service"
	"github.com/FirYuen/scaffold/pkg/app"
	"github.com/FirYuen/scaffold/pkg/conf"
	logger "github.com/FirYuen/scaffold/pkg/log"
	"github.com/FirYuen/scaffold/pkg/server"
)

func main()  {
	cfg, err := conf.Init("config/config.yaml")
	//fmt.Println(cfg)
	if err != nil {
		panic(err)
	}
	logger.Init(&cfg.Logger)

	svc := service.New(cfg)
	fmt.Println(svc)
	gin.SetMode(conf.Conf.App.Mode)
	app := app.New(cfg,
		app.WithName(cfg.App.Name),
		app.WithVersion(cfg.App.Version),
		app.WithLogger(logger.GetLogger()),
		app.Server(
			// init http server
			server.NewHttpServer(conf.Conf, svc),

		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}