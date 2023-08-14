package grpc_server

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go-skeleton/generated/grpc/go_skeleton"
	"go-skeleton/grpc/grpc_service"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

var Module = fx.Invoke(startGrpcServer)

type PingServer struct {
	go_skeleton.UnimplementedPingPongServer
	pService grpc_service.PingService
}

func (server *PingServer) Ping(ctx context.Context, request *go_skeleton.PingRequest) (*go_skeleton.PingResponse, error) {
	return server.pService.Ping(ctx, request)
}

func startGrpcServer(lifecycle fx.Lifecycle, pService grpc_service.PingService) error {

	port := viper.GetInt("grpc.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	pingServer := &PingServer{
		pService: pService,
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)
	//health check
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	//actual service
	grpcServer.RegisterService(&go_skeleton.PingPong_ServiceDesc, pingServer)

	lifecycle.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go func() {
			log.Println("gRPC server starting on port: ", port)
			err := grpcServer.Serve(lis)
			if err != nil {
				log.Println("grpcServer.Serve has error: ", err.Error())
			}
		}()
		return nil
	}, OnStop: func(c context.Context) error {
		log.Println("gRPC server Shutting down...")
		grpcServer.GracefulStop()
		log.Println("gRPC server shut down!")
		return nil
	}})
	return nil
}
