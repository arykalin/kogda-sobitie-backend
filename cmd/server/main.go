package main

import (
	"fmt"

	"github.com/arykalin/kogda-sobitie-backend/pkg/server"
	"go.uber.org/zap"
)

func main() {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %v", err))
	}
	s, err := server.NewService(zapLogger)
	if err != nil {
		zapLogger.Fatal("failed to create service", zap.Error(err))
	}
	err = s.Start()
	if err != nil {
		zapLogger.Fatal("failed to start service", zap.Error(err))
	}
}
