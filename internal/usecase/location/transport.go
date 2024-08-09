package location

import (
	"context"
	"github.com/core-go/core"
	b "github.com/core-go/core/handler/builder"
	v "github.com/core-go/core/validator"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/repository"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type LocationTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewLocationTransport(db *mongo.Database, logError core.Log, tracking b.TrackingConfig, writeLog func(context.Context, string, string, bool, string) error, action *core.ActionConfig) (LocationTransport, error) {
	validator, err := v.NewValidator[*Location]()
	if err != nil {
		return nil, err
	}
	locationMapper := geo.NewMapper[Location]()
	queryLocation := mq.UseQuery[Location, *LocationFilter]()
	searchBuilder := repository.NewSearchBuilder[Location, *LocationFilter](db, "location", queryLocation, search.GetSort, locationMapper.DbToModel)
	locationRepository := repository.NewRepository[Location, string](db, "location", locationMapper)
	locationService := NewLocationService(locationRepository)
	locationHandler := NewLocationHandler(searchBuilder.Search, locationService, logError, validator.Validate, tracking, writeLog, action)
	return locationHandler, nil
}
