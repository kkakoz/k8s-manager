package logic

import "go.uber.org/fx"

var Provider = fx.Provide(NewPodLogic, NewNsLogic, NewDeploymentLogic, NewServiceLogic, NewSecretLogic, NewIngressLogic, NewUserLogic)
