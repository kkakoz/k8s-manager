package handler

import (
	"go.uber.org/fx"
	"k8s-manager/logic"
)

var Provider = fx.Options(handlerProvider, logic.Provider)

var handlerProvider = fx.Provide(NewPodHandler, NewNsHandler, NewDeploymentHandler, NewSecretHandler)
