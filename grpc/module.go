package grpc

import (
	"go-skeleton/grpc/grpc_server"
	"go-skeleton/grpc/grpc_service"
	"go.uber.org/fx"
)

var Module = fx.Options(grpc_server.Module, grpc_service.Module)
