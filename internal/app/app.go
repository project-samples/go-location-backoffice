package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/mongo"
	sv "github.com/core-go/service/v10"
	"github.com/teris-io/shortid"

	"go-service/internal/bookable"
	"go-service/internal/event"
	"go-service/internal/location"
)

type ApplicationContext struct {
	HealthHandler   *health.Handler
	LocationHandler *location.LocationHandler
	EventHandler    *event.EventHandler
	BookableHandler *bookable.BookableHandler
}

func NewApp(ctx context.Context, mongoConfig mongo.MongoConfig) (*ApplicationContext, error) {
	db, err := mongo.Setup(ctx, mongoConfig)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg

	mongoChecker := mongo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)
	validator := sv.NewValidator()

	locationService := location.NewLocationService(db)
	locationHandler := location.NewLocationHandler(locationService, Generate, validator.Validate, logError)
	eventService := event.NewEventService(db)
	eventHandler := event.NewEventHandler(eventService, Generate, validator.Validate, logError)
	bookableService := bookable.NewBookableService(db)
	bookableHandler := bookable.NewBookableHandler(bookableService, Generate, validator.Validate, logError)

	return &ApplicationContext{
		HealthHandler:   healthHandler,
		LocationHandler: locationHandler,
		EventHandler:    eventHandler,
		BookableHandler: bookableHandler,
	}, nil
}

var sid *shortid.Shortid

func ShortId() (string, error) {
	if sid == nil {
		s, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			return "", err
		}
		sid = s
	}
	return sid.Generate()
}
func Generate(ctx context.Context) (string, error) {
	return ShortId()
}
