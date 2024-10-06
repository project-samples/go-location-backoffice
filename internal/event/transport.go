package event

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/core-go/core"
	b "github.com/core-go/core/builder"
	v "github.com/core-go/core/validator"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/repository"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
)

type EventTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewEventTransport(db *mongo.Database, logError core.Log, tracking b.TrackingConfig, writeLog func(context.Context, string, string, bool, string) error, action *core.ActionConfig) (EventTransport, error) {
	validator, err := v.NewValidator[*Event]()
	if err != nil {
		return nil, err
	}
	eventMapper := geo.NewMapper[Event]()
	queryEvent := mq.UseQuery[Event, *EventFilter]()
	searchBuilder := repository.NewSearchBuilder[Event, *EventFilter](db, "event", queryEvent, search.GetSort, eventMapper.DbToModel)
	eventRepository := repository.NewRepository[Event, string](db, "event", eventMapper)
	eventService := NewEventService(eventRepository)
	eventHandler := NewEventHandler(searchBuilder.Search, eventService, logError, validator.Validate, tracking, writeLog, action)
	return eventHandler, nil
}
