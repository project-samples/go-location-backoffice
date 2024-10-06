package tour

import (
	"context"
	"github.com/core-go/core"
	b "github.com/core-go/core/builder"
	v "github.com/core-go/core/validator"
	"github.com/core-go/mongo/repository"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type TourTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewTourTransport(db *mongo.Database, logError core.Log, tracking b.TrackingConfig, writeLog func(context.Context, string, string, bool, string) error, action *core.ActionConfig) (TourTransport, error) {
	validator, err := v.NewValidator[*Tour]()
	if err != nil {
		return nil, err
	}
	queryTour := mq.UseQuery[Tour, *TourFilter]()
	searchBuilder := repository.NewSearchBuilder[Tour, *TourFilter](db, "tour", queryTour, search.GetSort)
	tourRepository := repository.NewRepository[Tour, string](db, "tour")
	tourService := NewTourService(tourRepository)
	tourHandler := NewTourHandler(searchBuilder.Search, tourService, logError, validator.Validate, tracking, writeLog, action)
	return tourHandler, nil
}
