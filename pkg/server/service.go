package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	eventController "github.com/arykalin/kogda-sobitie-backend/internal/event_controller"
	grpcHandler "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/handlers"
	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	grpcServer         *grpc.Server
	ctx                context.Context
	mux                *runtime.ServeMux
	opts               []grpc.DialOption
	grpcServerEndpoint string
	logger             *zap.SugaredLogger
	httpServerEndpoint string
	cfg                config
}

type config struct {
	GrpcServerPort string
	GrpcServerHost string
	HttpServerPort string
}

func (s *Service) Start() error {
	var group errgroup.Group

	lis, err := net.Listen("tcp", s.cfg.GrpcServerPort)
	if err != nil {
		return err
	}
	group.Go(func() error {
		return s.grpcServer.Serve(lis)
	})
	group.Go(func() error {
		return grpcModels.RegisterApiServiceHandlerFromEndpoint(s.ctx, s.mux, s.grpcServerEndpoint, s.opts)
	})

	group.Go(func() error {
		return http.ListenAndServe(s.httpServerEndpoint, s.mux)
	})

	return group.Wait()
}

func NewService(cntrl eventController.Controller, logger *zap.SugaredLogger) (*Service, error) {
	// TODO: make config
	var cfg config
	cfg.GrpcServerPort = ":9090"
	cfg.GrpcServerHost = "localhost"
	cfg.HttpServerPort = ":8081"

	ctx := context.Background()
	var grpcServer = grpc.NewServer()
	grpcServer.RegisterService(&grpcModels.ApiService_ServiceDesc, grpcHandler.NewHandler(cntrl))

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcServerEndpoint := fmt.Sprintf("%s%s", cfg.GrpcServerHost, cfg.GrpcServerPort)
	httpServerEndpoint := cfg.HttpServerPort

	return &Service{
		grpcServer:         grpcServer,
		mux:                mux,
		ctx:                ctx,
		opts:               opts,
		grpcServerEndpoint: grpcServerEndpoint,
		httpServerEndpoint: httpServerEndpoint,
		logger:             logger,
		cfg:                cfg,
	}, nil
}
