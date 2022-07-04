package handler

import (
	"github.com/labstack/echo"
	"k8s-manager/logic"
	"k8s-manager/pkg/mdctx"
	"k8s-manager/request"
)

type DeploymentHandler struct {
	deploymentLogic *logic.DeploymentLogic
}

func NewDeploymentHandler(deploymentLogic *logic.DeploymentLogic) *DeploymentHandler {
	return &DeploymentHandler{deploymentLogic: deploymentLogic}
}

func (item *DeploymentHandler) List(ctx echo.Context) error {
	req := &request.ListReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.deploymentLogic.List(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *DeploymentHandler) Add(ctx echo.Context) error {
	req := &request.DeploymentAddReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.deploymentLogic.Add(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *DeploymentHandler) Delete(ctx echo.Context) error {
	req := &request.DeleteReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.deploymentLogic.Delete(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *DeploymentHandler) Update(ctx echo.Context) error {
	req := &request.ApplyReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.deploymentLogic.Update(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *DeploymentHandler) Restart(ctx echo.Context) error {
	req := &request.DeploymentRestartReq{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	err := item.deploymentLogic.Restart(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
