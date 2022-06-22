package main

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"k8s-manager/handler"
	"k8s-manager/k8s"
	"k8s-manager/pkg/app"
	"k8s-manager/pkg/conf"
	"k8s-manager/pkg/logger"
	"k8s-manager/router"
	"k8s-manager/server"
	"log"
	"net/http"
)

func NewApp(handler http.Handler, servers []app.Server) *app.Application {
	return app.NewApplication("k8s_manager", handler, servers)
}

func main() {

	conf.InitConfig()

	var app = new(app.Application)
	fx.New(
		handler.Provider,
		server.Provider,
		router.Provider,
		logger.Provider,
		k8s.Provider,
		fx.Provide(NewApp),
		fx.Supply(viper.GetViper()),
		fx.Populate(&app),
	)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}

type Pod struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	UID    string            `json:"uid"`
}
