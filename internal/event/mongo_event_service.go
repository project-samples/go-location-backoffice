package event

import (
	"reflect"

	mgo "github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	"github.com/core-go/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoEventService struct {
	search.SearchService
	service.GenericService
	Mapper mgo.Mapper
}

func NewEventService(db *mongo.Database) *MongoEventService {
	var model Event
	modelType := reflect.TypeOf(model)
	mapper := geo.NewMapper(modelType)
	queryBuilder := query.NewBuilder(modelType)
	searchService, genericService := mgo.NewSearchWriter(db, "event", modelType, queryBuilder.BuildQuery, search.GetSort, mapper)
	return &MongoEventService{SearchService: searchService, GenericService: genericService, Mapper: mapper}
}
