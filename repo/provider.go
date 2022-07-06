package repo

import "go.uber.org/fx"

var Provider = fx.Provide(NewUserRepo)
