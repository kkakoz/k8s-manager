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
	podG := e.Group("/api/deployments")
	{
		podG.GET("", u.handler.List)
		podG.POST("", u.handler.Add)
		podG.DELETE("/:name", u.handler.Delete)
		podG.PUT("/", u.handler.Update)
	}
}
