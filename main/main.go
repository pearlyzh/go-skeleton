package main

import (
	"go-skeleton/utilityfx"
	"go.uber.org/fx"
	"log"
)

// Can use x.Shutdowner https://stackoverflow.com/questions/65857064/how-do-you-gracefully-exit-a-go-uber-fx-app

var version string

func main() {
	log.Println("Starting skeleton app - version " + version)
	app := fx.New(utilityfx.Utility)
	app.Run()
}
