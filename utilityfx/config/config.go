package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"os"
)

// for GIT revision: https://stackoverflow.com/questions/28459102/golang-compile-environment-variable-into-binary
// for inject config values for running: https://gobyexample.com/environment-variables

var Module = fx.Invoke(loadConfigurations)

func loadConfigurations() {
	cp := os.Getenv("CONFIG_PATH")
	if len(cp) == 0 {
		log.Println("CONFIG_PATH is empty! Start using local config file!")
		cp = "config/local.yaml"
	}
	viper.SetConfigFile(cp)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to load configuration file: ", err)
	}
	log.Printf("Use config %s\n", viper.ConfigFileUsed())
}
