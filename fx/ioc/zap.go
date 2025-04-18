package ioc

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var ZapFxOpt = fx.Provide(
	zap.NewExample,
)
