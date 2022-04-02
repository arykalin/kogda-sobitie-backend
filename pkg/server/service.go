package server

import (
	"context"
	"flag"
	"net"
	"net/http"

	grpcHandler "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/handlers"
	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	grpcServer         *grpc.Server
	ctx                context.Context
	mux                *runtime.ServeMux
	opts               []grpc.DialOption
	grpcServerEndpoint *string
}

func (s *Service) Start() error {
	var group errgroup.Group

	lis, err := net.Listen("tcp", ":8842")
	if err != nil {
		return err
	}
	group.Go(func() error {
		return s.grpcServer.Serve(lis)
	})
	group.Go(func() error {
		return grpcModels.RegisterApiServiceHandlerFromEndpoint(s.ctx, s.mux, *s.grpcServerEndpoint, s.opts)
	})

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
	grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")

	return &Service{
		grpcServer:         grpcServer,
		mux:                mux,
		ctx:                ctx,
		opts:               opts,
		grpcServerEndpoint: grpcServerEndpoint,
	}, nil
}
