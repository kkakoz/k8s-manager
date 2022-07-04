package router

import (
	"github.com/labstack/echo"
	"k8s-manager/handler"
)

type ingressRouter struct {
	handler *handler.IngressHandler
}

func NewIngressRouter(handler *handler.IngressHandler) *ingressRouter {
	return &ingressRouter{handler: handler}
}

func (u ingressRouter) AddRouter(e *echo.Echo) {
	group := e.Group("/api/ingresses")
	{
		group.GET("", u.handler.List)
		group.POST("", u.handler.Add)
	}
}
