package server

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"k8s-manager/pkg/app"
)

func Server(viper *viper.Viper) []app.Server {
	return []app.Server{}
}

var Provider = fx.Provide(Server)
