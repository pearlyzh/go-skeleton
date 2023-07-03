package grpc_server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go-skeleton/generated/grpc/go_skeleton"
	"go-skeleton/utilityfx/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

var Module = fx.Invoke(StartGrpcServer)

type PingServer struct {
	go_skeleton.UnimplementedPingPongServer
}

func (s *PingServer) Ping(ctx context.Context, request *go_skeleton.PingRequest) (*go_skeleton.PingResponse, error) {
	logger.Ctx(ctx).Info("Ping request: ", zap.String("request", request.GetName()))
	return &go_skeleton.PingResponse{
		Name:      request.Name,
		RequestId: uuid.New().String(),
	}, nil
}

func StartGrpcServer(lifecycle fx.Lifecycle) error {

	port := viper.GetInt("grpc.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	pingServer := &PingServer{}
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
		log.Println("gRPC server Shutted down")
		return nil
	}})
	return nil
}
