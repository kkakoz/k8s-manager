package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/logic"
	"k8s-manager/pkg/mdctx"
	"k8s-manager/request"
)

type PodHandler struct {
	podLogic *logic.PodLogic
}

func NewPodHandler(podLogic *logic.PodLogic) *PodHandler {
	return &PodHandler{podLogic: podLogic}
}

func (item *PodHandler) List(ctx echo.Context) error {
	req := &request.ListReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.podLogic.List(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *PodHandler) Add(ctx echo.Context) error {
	req := &request.PodAddReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.podLogic.Add(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *PodHandler) Delete(ctx echo.Context) error {
	req := &request.DeleteReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.podLogic.Delete(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *PodHandler) Apply(ctx echo.Context) error {
	req := &request.ApplyReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.podLogic.Update(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *PodHandler) Logs(ctx echo.Context) error {
	req := &request.PodLogReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	logs, err := item.podLogic.GetLog(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]interface{}{
		"logs": logs,
	})
}

func (item *PodHandler) Terminal(ctx echo.Context) error {
	req := &request.PodTerminalReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.podLogic.Terminal(mdctx.NewCtx(ctx.Request()), ctx, req)
	if err != nil {
		return err
	}
	return nil
}
