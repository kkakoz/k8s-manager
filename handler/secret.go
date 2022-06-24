package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/logic"
	"k8s-manager/pkg/mdctx"
	"k8s-manager/request"
)

type SecretHandler struct {
	secretLogic *logic.SecretLogic
}

func NewSecretHandler(secretLogic *logic.SecretLogic) *SecretHandler {
	return &SecretHandler{secretLogic: secretLogic}
}

func (item *SecretHandler) List(ctx echo.Context) error {
	list, err := item.secretLogic.List(mdctx.NewCtx(ctx.Request()))
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *SecretHandler) Add(ctx echo.Context) error {
	req := &request.SecretAddReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.secretLogic.Add(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
