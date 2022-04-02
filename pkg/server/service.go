package server

import (
	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
	"google.golang.org/grpc"
)

type Service struct{}

func (s *Service) Start() {
	var grpcServer = grpc.NewServer()
	grpcServer.RegisterService(&grpcModels.ApiService_ServiceDesc, grpcHalers)
}

func NewService() *Service {
	return &Service{}
}
