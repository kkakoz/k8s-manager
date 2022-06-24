package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/logic"
	"k8s-manager/pkg/mdctx"
)

type NsHandler struct {
	nsLogic *logic.NSLogic
}

func NewNsHandler(nsLogic *logic.NSLogic) *NsHandler {
	return &NsHandler{nsLogic: nsLogic}
}

func (item *NsHandler) List(ctx echo.Context) error {
	list, err := item.nsLogic.List(mdctx.NewCtx(ctx.Request()))
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *NsHandler) Add(ctx echo.Context) error {
	req := &struct {
		Name string `json:"name"`
	}{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.nsLogic.Add(mdctx.NewCtx(ctx.Request()), req.Name)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
