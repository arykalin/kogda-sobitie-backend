package eventController

import (
	"reflect"
	"testing"

	"github.com/arykalin/kogda-sobitie-backend/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewController(t *testing.T) {
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
			if got := NewController(tt.args.dbClient, tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
