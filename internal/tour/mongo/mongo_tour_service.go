package mongo

import (
	"reflect"

	mgo "github.com/core-go/mongo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	"github.com/core-go/service"
	"go.mongodb.org/mongo-driver/mongo"

	. "go-service/internal/tour"
)

type MongoTourService struct {
	search.SearchService
	service.GenericService
}

func NewTourService(db *mongo.Database) *MongoTourService {
	var model Tour
	modelType := reflect.TypeOf(model)
	queryBuilder := query.NewBuilder(modelType)
	searchService, genericService := mgo.NewSearchWriter(db, "tour", modelType, queryBuilder.BuildQuery, search.GetSort)
	return &MongoTourService{SearchService: searchService, GenericService: genericService}
}
