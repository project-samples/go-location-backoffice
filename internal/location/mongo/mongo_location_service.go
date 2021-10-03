package mongo

import (
	"reflect"

	mgo "github.com/core-go/mongo"
	"github.com/core-go/mongo/geo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	"github.com/core-go/service"
	"go.mongodb.org/mongo-driver/mongo"

	. "go-service/internal/location"
)

type MongoLocationService struct {
	search.SearchService
	service.GenericService
	Mapper mgo.Mapper
}

func NewLocationService(db *mongo.Database) *MongoLocationService {
	var model Location
	modelType := reflect.TypeOf(model)
	mapper := geo.NewMapper(modelType)
	queryBuilder := query.NewBuilder(modelType)
	searchService, genericService := mgo.NewSearchWriter(db, "location", modelType, queryBuilder.BuildQuery, search.GetSort, mapper)
	return &MongoLocationService{SearchService: searchService, GenericService: genericService, Mapper: mapper}
}
