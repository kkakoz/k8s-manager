package router

import (
	"github.com/labstack/echo"
	"k8s-manager/handler"
)

type secretRouter struct {
	handler *handler.SecretHandler
}

func NewSecretRouter(handler *handler.SecretHandler) *secretRouter {
	return &secretRouter{handler: handler}
}

func (u secretRouter) AddRouter(e *echo.Echo) {
	nsG := e.Group("/api/secrets")
	{
		nsG.GET("", u.handler.List)
		nsG.POST("", u.handler.Add)
	}
}
