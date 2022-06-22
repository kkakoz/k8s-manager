package router

import (
	"github.com/labstack/echo"
	"k8s-manager/handler"
)

type nsRouter struct {
	handler *handler.NsHandler
}

func NewNsRouter(handler *handler.NsHandler) *nsRouter {
	return &nsRouter{handler: handler}
}

func (u nsRouter) AddRouter(e *echo.Echo) {
	nsG := e.Group("/api/ns")
	{
		nsG.GET("", u.handler.List)
		nsG.POST("", u.handler.Add)
	}
}
