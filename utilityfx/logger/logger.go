package logger

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
)

var Module = fx.Invoke(InitLogger)

var debugMode = "logger.debug"

func InitLogger(lifecycle fx.Lifecycle) {
	log.Printf("Start initialising Logger with debug mode %v\n", viper.GetBool(debugMode))
	InitOverriddenLogger(!viper.GetBool(debugMode))
	lifecycle.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		log.Println("Start sync buffering logs!")
		err := Sync()
		if err != nil {
			log.Printf("Sync log failed! %v\n", err)
		}
		log.Println("Finish sync buffering logs!")
		return nil
	}})
	log.Println("Finish initialising Logger!")
}

func Sync() error {
	return globalLogger.logger.Sync()
}
