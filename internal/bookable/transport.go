package bookable

import (
	"context"
	"github.com/core-go/core"
	b "github.com/core-go/core/builder"
	v "github.com/core-go/core/validator"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/repository"
	"github.com/core-go/search"
	mq "github.com/core-go/search/mongo/query"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type BookableTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewBookableTransport(db *mongo.Database, logError core.Log, tracking b.TrackingConfig, writeLog func(context.Context, string, string, bool, string) error, action *core.ActionConfig) (BookableTransport, error) {
	validator, err := v.NewValidator[*Bookable]()
	if err != nil {
		return nil, err
	}
	bookableMapper := geo.NewMapper[Bookable]()
	queryBookable := mq.UseQuery[Bookable, *BookableFilter]()
	searchBuilder := repository.NewSearchBuilder[Bookable, *BookableFilter](db, "bookable", queryBookable, search.GetSort, bookableMapper.DbToModel)
	bookableRepository := repository.NewRepository[Bookable, string](db, "bookable", bookableMapper)
	bookableService := NewBookableService(bookableRepository)
	bookableHandler := NewBookableHandler(searchBuilder.Search, bookableService, logError, validator.Validate, tracking, writeLog, action)
	return bookableHandler, nil
}
