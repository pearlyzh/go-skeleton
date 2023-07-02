package logger

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Invoke(InitLogger)

func InitLogger(lifecycle fx.Lifecycle) {
	InitOverriddenLogger(!viper.GetBool("debug.logger"))
	lifecycle.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		_ = Sync()
		return nil
	}})
}

func Sync() error {
	return globalLogger.logger.Sync()
}
