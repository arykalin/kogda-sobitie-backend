package clients

import grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"

type client struct {
	grpcClient grpcModels.ApiServiceClient
}
