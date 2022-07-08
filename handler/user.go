package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/local"
	"k8s-manager/logic"
	"k8s-manager/pkg/errno"
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
	req := &request.LoginReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	token, err := item.userLogic.Login(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, token)
}

func (item *UserHandler) Current(ctx echo.Context) error {
	user, b := local.GetUser(ctx.Request().Context())
	if !b {
		return errno.NewErr(401, 401, "请重新登录")
	}
	return ctx.JSON(200, user)
}

func (item *UserHandler) Add(ctx echo.Context) error {
	req := &request.UserAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	return item.userLogic.Add(mdctx.NewCtx(ctx.Request()), req)
}
