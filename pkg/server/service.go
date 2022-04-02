package server

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	models "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
	"github.com/golang/protobuf/jsonpb"
	_struct "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	eventBus <-chan interface{}
}

func NewService(eventBus <-chan interface{}) *Service {
	return &Service{eventBus: eventBus}
}

func (b *Service) Address(_ context.Context, req *models.AddressRequest) (*models.AddressResponse, error) {
	if req.Address != "Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6" {
		return &models.AddressResponse{}, status.Error(codes.FailedPrecondition, "wallet not found")
	}
	return &models.AddressResponse{
		Balance: map[string]string{
			"BIP": "12345678987654321",
		},
		TransactionsCount: "120",
	}, nil
}

func (b *Service) Subscribe(req *models.SubscribeRequest, stream models.Service_SubscribeServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case event := <-b.eventBus:
			byteData, err := json.Marshal(event)
			if err != nil {
				return err
			}
			var bb bytes.Buffer
			bb.Write(byteData)
			data := &_struct.Struct{Fields: make(map[string]*_struct.Value)}
			if err := (&jsonpb.Unmarshaler{}).Unmarshal(&bb, data); err != nil {
				return err
			}

			if err := stream.Send(&models.SubscribeResponse{
				Query: req.Query,
				Data:  data,
				Events: []*models.SubscribeResponse_Event{
					{
						Key:    "tx.hash",
						Events: []string{"01EFD8EEF507A5BFC4A7D57ECA6F61B96B7CDFF559698639A6733D25E2553539"},
					},
				},
			}); err != nil {
				return err
			}
		case <-time.After(5 * time.Second):
			return nil
		}
	}
}
