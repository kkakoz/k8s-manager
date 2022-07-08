package router

import (
	"github.com/labstack/echo"
	"k8s-manager/handler"
)

type UserRouter struct {
	handler *handler.UserHandler
}

func NewUserRouter(handler *handler.UserHandler) *UserRouter {
	return &UserRouter{handler: handler}
}

func (u *UserRouter) AddRouter(e *echo.Echo) {

	e.POST("/api/login", u.handler.Login)

	g := e.Group("/api/users")
	{
		g.GET("/current", u.handler.Current)
		g.POST("/", u.handler.Add)
	}
}
