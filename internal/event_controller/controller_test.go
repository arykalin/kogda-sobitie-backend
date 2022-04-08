package eventController

import (
	"reflect"
	"testing"

	"github.com/arykalin/kogda-sobitie-backend/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func TestNewController(t *testing.T) {
	sLoggerConfig := zap.NewDevelopmentConfig()
	sLoggerConfig.DisableStacktrace = true
	sLoggerConfig.DisableCaller = true
	sLogger, err := sLoggerConfig.Build()
	if err != nil {
		panic(err)
	}
	logger := sLogger.Sugar()

	client := db.Dbconnect()
	collection := client.Database("golang").Collection("events")
	type args struct {
		dbClient   *mongo.Client
		collection *mongo.Collection
	}
	tests := []struct {
		name string
		args args
		want *controller
	}{
		{"TestNewController", args{client, collection}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.dbClient, tt.args.collection, logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
