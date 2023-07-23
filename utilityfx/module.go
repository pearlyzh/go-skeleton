package utilityfx

import (
	"go-skeleton/utilityfx/config"
	"go-skeleton/utilityfx/logger"
	"go-skeleton/utilityfx/mysql"
	"go.uber.org/fx"
)

var Module = fx.Options(config.Module, logger.Module, mysql.Module)
