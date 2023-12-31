package main

import (
	"go-skeleton/grpc"
	"go-skeleton/repository"
	"go-skeleton/utilityfx"
	"go.uber.org/fx"
	"log"
)

// Can use x.Shutdowner https://stackoverflow.com/questions/65857064/how-do-you-gracefully-exit-a-go-uber-fx-app

var version string

/*
TODO:
 - Add Redis client
 - Add Kafka client for Producers and Consumers
 - Add Monitoring
 - Add Tests
 - Add Tracing
 - Add RateLimiter
*/
func main() {
	log.Println("Starting skeleton app - version " + version)
	app := fx.New(utilityfx.Module, repository.Module, grpc.Module)
	app.Run()
}
