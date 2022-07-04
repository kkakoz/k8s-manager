package router

import (
	"github.com/labstack/echo"
	"k8s-manager/handler"
)

type deploymentRouter struct {
	handler *handler.DeploymentHandler
}

func NewDeploymentRouter(handler *handler.DeploymentHandler) *deploymentRouter {
	return &deploymentRouter{handler: handler}
}

func (u deploymentRouter) AddRouter(e *echo.Echo) {
	group := e.Group("/api/deployments")
	{
		group.GET("", u.handler.List)
		group.POST("", u.handler.Add)
		group.DELETE("/:name", u.handler.Delete)
		group.PUT("/", u.handler.Update)
		group.PUT("/restart", u.handler.Restart)
	}
}
