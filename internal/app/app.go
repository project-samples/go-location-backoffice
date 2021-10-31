package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	v "github.com/core-go/service/v10"
	"github.com/teris-io/shortid"
	"reflect"

	"go-service/internal/usecase/bookable"
	"go-service/internal/usecase/event"
	"go-service/internal/usecase/location"
	"go-service/internal/usecase/tour"
)

type ApplicationContext struct {
	HealthHandler   *health.Handler
	LocationHandler location.LocationHandler
	EventHandler    *event.EventHandler
	BookableHandler *bookable.BookableHandler
	TourHandler     *tour.TourHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	db, err := mongo.Setup(ctx, root.Mongo)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg
	status := sv.InitializeStatus(root.Status)
	action := sv.InitializeAction(root.Action)
	validator := v.NewValidator()

	mongoChecker := mongo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	locationType := reflect.TypeOf(location.Location{})
	locationMapper := geo.NewMapper(locationType)
	locationQueryBuilder := query.NewBuilder(locationType)
	locationSearchBuilder := mongo.NewSearchBuilder(db, "location", locationQueryBuilder.BuildQuery, search.GetSort, locationMapper.DbToModel)
	locationRepository := mongo.NewRepository(db, "location", locationType, locationMapper)
	locationService := location.NewLocationService(locationRepository)
	locationHandler := location.NewLocationHandler(locationSearchBuilder.Search, locationService, status, logError, validator.Validate, &action)

	eventService := event.NewEventService(db)
	eventHandler := event.NewEventHandler(eventService, Generate, validator.Validate, logError)
	bookableService := bookable.NewBookableService(db)
	bookableHandler := bookable.NewBookableHandler(bookableService, Generate, validator.Validate, logError)
	tourService := tour.NewTourService(db)
	tourHandler := tour.NewTourHandler(tourService, Generate, validator.Validate, logError)
	return &ApplicationContext{
		HealthHandler:   healthHandler,
		LocationHandler: locationHandler,
		EventHandler:    eventHandler,
		BookableHandler: bookableHandler,
		TourHandler:     tourHandler,
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
