package api

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
	"wm-infoflow-api-go/common/log"
	"wm-infoflow-api-go/conf"
	"wm-infoflow-api-go/initialize"
	"wm-infoflow-api-go/service"
)

type api struct {
	config *conf.Cfg
	cron   bool
}

func New(cfg *conf.Cfg) *api {
	return &api{config: cfg}
}

func (r *api) StartCron() *api {
	r.cron = true
	return r
}

type server interface {
	ListenAndServe() error
}

func (r *api) Start() error {
	//init router
	Router := initialize.Routers()
	s := initServer(r.config.Frontend.Host, Router)
	service.NewServicesWrap(initialize.DefaultRedisClient)
	log.Info("server run success...")
	log.Error(s.ListenAndServe().Error())
	return nil
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
