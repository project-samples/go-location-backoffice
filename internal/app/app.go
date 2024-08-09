package app

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/core-go/health"
	hm "github.com/core-go/health/mongo"
	"github.com/core-go/log/zap"
	"github.com/teris-io/shortid"

	"go-service/internal/bookable"
	"go-service/internal/event"
	"go-service/internal/location"
	"go-service/internal/tour"
)

type ApplicationContext struct {
	Health   *health.Handler
	Location location.LocationTransport
	Event    event.EventTransport
	Bookable bookable.BookableTransport
	Tour     tour.TourTransport
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.Uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.Mongo.Database)
	logError := log.LogError

	mongoChecker := hm.NewHealthChecker(client)
	healthHandler := health.NewHandler(mongoChecker)

	locationHandler, err := location.NewLocationTransport(db, logError, cfg.Tracking, nil, nil)
	if err != nil {
		return nil, err
	}
	eventHandler, err := event.NewEventTransport(db, logError, cfg.Tracking, nil, nil)
	if err != nil {
		return nil, err
	}
	bookableHandler, err := bookable.NewBookableTransport(db, logError, cfg.Tracking, nil, nil)
	if err != nil {
		return nil, err
	}
	tourHandler, err := tour.NewTourTransport(db, logError, cfg.Tracking, nil, nil)

	return &ApplicationContext{
		Health:   healthHandler,
		Location: locationHandler,
		Event:    eventHandler,
		Bookable: bookableHandler,
		Tour:     tourHandler,
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
