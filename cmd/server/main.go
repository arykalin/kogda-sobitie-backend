package main

import (
	"fmt"

	"github.com/arykalin/kogda-sobitie-backend/internal/db"
	eventController "github.com/arykalin/kogda-sobitie-backend/internal/event_controller"
	"github.com/arykalin/kogda-sobitie-backend/pkg/server"
	"go.uber.org/zap"
)

func main() {
	sLoggerConfig := zap.NewDevelopmentConfig()
	sLoggerConfig.DisableStacktrace = true
	sLoggerConfig.DisableCaller = true
	sLogger, err := sLoggerConfig.Build()
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %v", err))
	}
	logger := sLogger.Sugar()

	dbClient := db.Dbconnect()
	collection := dbClient.Database("year-wheel").Collection("events")

	cntrl := eventController.NewController(dbClient, collection, logger)
	s, err := server.NewService(cntrl, logger)
	if err != nil {
		logger.Fatal("failed to create service", zap.Error(err))
	}
	err = s.Start()
	if err != nil {
		logger.Fatal("failed to start service", zap.Error(err))
	}
}
