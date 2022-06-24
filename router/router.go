package router

import (
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"k8s-manager/pkg/echox"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func NewHttp(logger *zap.Logger, pod *podRouter, ns *nsRouter, deployment *deploymentRouter, secret *secretRouter) http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = echox.ErrHandler(logger)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.Debug = true
	pod.AddRouter(e)
	ns.AddRouter(e)
	deployment.AddRouter(e)
	secret.AddRouter(e)
	return e
}

var Provider = fx.Provide(NewHttp, NewPodRouter, NewNsRouter, NewDeploymentRouter, NewSecretRouter)
