package utilityfx

import (
	"go-skeleton/utilityfx/config"
	"go-skeleton/utilityfx/logger"
	"go.uber.org/fx"
)

var Utility = fx.Options(config.Module, logger.Module)
