package router

import (
	"github.com/labstack/echo"
	"k8s-manager/handler"
)

type podRouter struct {
	handler *handler.PodHandler
}

func NewPodRouter(handler *handler.PodHandler) *podRouter {
	return &podRouter{handler: handler}
}

func (u podRouter) AddRouter(e *echo.Echo) {
	podG := e.Group("/api/pods", loginAuth)
	{
		podG.GET("", u.handler.List)
		podG.GET("/:name/logs", u.handler.Logs)
		podG.POST("", u.handler.Add)
		podG.DELETE("/:name", u.handler.Delete)
		podG.PUT("/", u.handler.Apply)
		podG.GET("/:name/terminal", u.handler.Terminal)
	}
}
