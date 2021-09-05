package app

import (
	"context"
	"github.com/core-go/mongo"
	. "github.com/core-go/service"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, mongoConfig mongo.MongoConfig) error {
	app, err := NewApp(ctx, mongoConfig)
	if err != nil {
		return err
	}

	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)

	locationPath := "/locations"
	r.HandleFunc(locationPath, app.LocationHandler.GetAll).Methods(GET)
	r.HandleFunc(locationPath+"/search", app.LocationHandler.Search).Methods(GET, POST)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Load).Methods(GET)
	r.HandleFunc(locationPath, app.LocationHandler.Create).Methods(POST)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Update).Methods(PUT)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Patch).Methods(PATCH)
	r.HandleFunc(locationPath+"/{id}", app.LocationHandler.Delete).Methods(DELETE)
	return nil
}
