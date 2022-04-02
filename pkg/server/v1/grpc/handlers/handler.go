package handlers

import (
	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
)

type handler struct {
	event
	grpcModels.UnimplementedApiServiceServer
}
