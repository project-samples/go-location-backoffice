package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/mongo"
	"github.com/core-go/service/uuid"
	sv "github.com/core-go/service/v10"

	"go-service/internal/location"
)

type ApplicationContext struct {
	HealthHandler   *health.Handler
	LocationHandler *location.LocationHandler
}

func NewApp(ctx context.Context, mongoConfig mongo.MongoConfig) (*ApplicationContext, error) {
	db, err := mongo.Setup(ctx, mongoConfig)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg

	validator := sv.NewValidator()
	locationService := location.NewLocationService(db)
	locationHandler := location.NewLocationHandler(locationService, uuid.Generate, validator.Validate, logError)

	mongoChecker := mongo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	return &ApplicationContext{
		HealthHandler:   healthHandler,
		LocationHandler: locationHandler,
	}, nil
}
