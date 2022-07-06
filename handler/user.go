package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/logic"
	"k8s-manager/pkg/mdctx"
	"k8s-manager/request"
)

type UserHandler struct {
	userLogic *logic.UserLogic
}

func NewUserHandler(userLogic *logic.UserLogic) *UserHandler {
	return &UserHandler{userLogic: userLogic}
}

func (item *UserHandler) Login(ctx echo.Context) error {
	auth := &request.LoginReq{}
	err := ctx.Bind(auth)
	if err != nil {
		return err
	}
	token, err := item.userLogic.Login(mdctx.NewCtx(ctx.Request()), auth)
	if err != nil {
		return err
	}
	return ctx.JSON(200, token)
}

func (item *UserHandler) Current(ctx echo.Context) error {
	ctx.Request().Header.Get("Authorization")
	token, err := item.userLogic.Current(mdctx.NewCtx(ctx.Request()), "")
	if err != nil {
		return err
	}
	return ctx.JSON(200, token)
}
