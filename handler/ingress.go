package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/logic"
	"k8s-manager/pkg/mdctx"
)

type IngressHandler struct {
	ingressLogic *logic.IngressLogic
}

func NewIngressHandler(ingressLogic *logic.IngressLogic) *IngressHandler {
	return &IngressHandler{ingressLogic: ingressLogic}
}

func (item *IngressHandler) List(ctx echo.Context) error {
	list, err := item.ingressLogic.List(mdctx.NewCtx(ctx.Request()))
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *IngressHandler) Add(ctx echo.Context) error {
	req := &struct {
		Name string `json:"name"`
	}{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.ingressLogic.Add(mdctx.NewCtx(ctx.Request()), req.Name)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
