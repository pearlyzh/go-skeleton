package grpc_service

import (
	"context"
	"github.com/google/uuid"
	"go-skeleton/generated/grpc/go_skeleton"
	"go-skeleton/repository"
	"go-skeleton/utilityfx/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(newPingService)

func newPingService(repo repository.RequestRepository) PingService {
	return &pingServiceImpl{repo: repo}
}

type PingService interface {
	Ping(ctx context.Context, request *go_skeleton.PingRequest) (*go_skeleton.PingResponse, error)
}

type pingServiceImpl struct {
	repo repository.RequestRepository
}

func (pService *pingServiceImpl) Ping(ctx context.Context, request *go_skeleton.PingRequest) (*go_skeleton.PingResponse, error) {
	logger.Ctx(ctx).Info("Ping request: ", zap.String("request", request.GetName()))
	u := uuid.New().String()
	err := pService.repo.SaveRequest(ctx, &u, request)
	if err != nil {
		logger.Ctx(ctx).Error("Failed saving data!", zap.Error(err))
	}
	return &go_skeleton.PingResponse{
		Name:      request.Name,
		RequestId: u,
	}, nil
}
