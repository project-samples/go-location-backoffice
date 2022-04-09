package app

import (
	"context"
	. "github.com/core-go/service"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)

	locationPath := "/locations"
	location := app.LocationHandler
	// r.HandleFunc(locationPath, location.GetAll).Methods(GET)
	r.HandleFunc(locationPath+"/search", location.Search).Methods(GET, POST)
	r.HandleFunc(locationPath+"/{id}", location.Load).Methods(GET)
	r.HandleFunc(locationPath, location.Create).Methods(POST)
	r.HandleFunc(locationPath+"/{id}", location.Update).Methods(PUT)
	r.HandleFunc(locationPath+"/{id}", location.Patch).Methods(PATCH)
	r.HandleFunc(locationPath+"/{id}", location.Delete).Methods(DELETE)

	eventPath := "/events"
	event := app.EventHandler
	// r.HandleFunc(eventPath, event.GetAll).Methods(GET)
	r.HandleFunc(eventPath+"/search", event.Search).Methods(GET, POST)
	r.HandleFunc(eventPath+"/{id}", event.Load).Methods(GET)
	r.HandleFunc(eventPath, event.Create).Methods(POST)
	r.HandleFunc(eventPath+"/{id}", event.Update).Methods(PUT)
	r.HandleFunc(eventPath+"/{id}", event.Patch).Methods(PATCH)
	r.HandleFunc(eventPath+"/{id}", event.Delete).Methods(DELETE)

	bookablePath := "/bookables"
	bookable := app.BookableHandler
	// r.HandleFunc(bookablePath, bookable.GetAll).Methods(GET)
	r.HandleFunc(bookablePath+"/search", bookable.Search).Methods(GET, POST)
	r.HandleFunc(bookablePath+"/{id}", bookable.Load).Methods(GET)
	r.HandleFunc(bookablePath, bookable.Create).Methods(POST)
	r.HandleFunc(bookablePath+"/{id}", bookable.Update).Methods(PUT)
	r.HandleFunc(bookablePath+"/{id}", bookable.Patch).Methods(PATCH)
	r.HandleFunc(bookablePath+"/{id}", bookable.Delete).Methods(DELETE)

	tourPath := "/tours"
	tour := app.BookableHandler
	// r.HandleFunc(tourPath, tour.GetAll).Methods(GET)
	r.HandleFunc(tourPath+"/search", tour.Search).Methods(GET, POST)
	r.HandleFunc(tourPath+"/{id}", tour.Load).Methods(GET)
	r.HandleFunc(tourPath, tour.Create).Methods(POST)
	r.HandleFunc(tourPath+"/{id}", tour.Update).Methods(PUT)
	r.HandleFunc(tourPath+"/{id}", tour.Patch).Methods(PATCH)
	r.HandleFunc(tourPath+"/{id}", tour.Delete).Methods(DELETE)
	return nil
}
