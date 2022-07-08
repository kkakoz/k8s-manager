package router

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"k8s-manager/model"
	"k8s-manager/pkg/echox"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func NewHttp(logger *zap.Logger, routers []Router) http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = echox.ErrHandler(logger)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig), withNamespace)
	db := ormx.DB(context.TODO())
	db.AutoMigrate(&model.User{})

	e.Debug = true
	for _, router := range routers {
		router.AddRouter(e)
	}
	return e
}

type Router interface {
	AddRouter(e *echo.Echo)
}

func Routers(pod *podRouter, ns *nsRouter, deployment *deploymentRouter,
	secret *secretRouter, service *serviceRouter, ingress *ingressRouter, user *UserRouter) []Router {
	return []Router{pod, ns, deployment, secret, service, ingress, user}
}

var Provider = fx.Provide(NewHttp, NewPodRouter, NewNsRouter, NewDeploymentRouter, NewSecretRouter, NewServiceRouter, NewIngressRouter, NewUserRouter, Routers)
