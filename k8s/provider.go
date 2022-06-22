package k8s

import "go.uber.org/fx"

var Provider = fx.Provide(NewK8sClientSet)
