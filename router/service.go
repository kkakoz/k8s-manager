package router

import (
	"github.com/labstack/echo"
	"k8s-manager/handler"
)

type serviceRouter struct {
	handler *handler.ServiceHandler
}

func NewServiceRouter(handler *handler.ServiceHandler) *serviceRouter {
	return &serviceRouter{handler: handler}
}

func (u serviceRouter) AddRouter(e *echo.Echo) {
	serviceG := e.Group("/api/services")
	{
		serviceG.GET("", u.handler.List)
		serviceG.POST("", u.handler.Add)
	}
}
