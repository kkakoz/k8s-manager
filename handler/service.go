package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/logic"
	"k8s-manager/pkg/mdctx"
	"k8s-manager/request"
)

type ServiceHandler struct {
	serviceLogic *logic.ServiceLogic
}

func NewServiceHandler(ServiceLogic *logic.ServiceLogic) *ServiceHandler {
	return &ServiceHandler{serviceLogic: ServiceLogic}
}

func (item *ServiceHandler) List(ctx echo.Context) error {
	req := &request.ListReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.serviceLogic.List(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *ServiceHandler) Add(ctx echo.Context) error {
	req := &request.ServiceAddReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.serviceLogic.Add(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
