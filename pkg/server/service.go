package server

import (
	"context"
	"flag"
	"net"
	"net/http"

	grpcHandler "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/handlers"
	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	grpcServer *grpc.Server
	ctx        context.Context
	mux        *runtime.ServeMux
}

func (s *Service) Start() error {
	lis, err := net.Listen("tcp", ":8842")
	if err != nil {
		return err
	}
	err = s.grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	err = http.ListenAndServe(":8081", s.mux)
	if err != nil {
		return err
	}
	return nil
}

func NewService() (*Service, error) {
	ctx := context.Background()
	var grpcServer = grpc.NewServer()
	grpcServer.RegisterService(&grpcModels.ApiService_ServiceDesc, grpcHandler.NewHandler())

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	var grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	err := grpcModels.RegisterApiServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return nil, err
	}

	return &Service{
		grpcServer: grpcServer,
		mux:        mux,
		ctx:        ctx,
	}, nil
}
