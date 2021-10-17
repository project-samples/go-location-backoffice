package event

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/model-builder"
)

type EventHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	Service EventService
}

func NewEventHandler(eventService EventService, generateId func(context.Context) (string, error), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string)) *EventHandler {
	modelType := reflect.TypeOf(Event{})
	searchModelType := reflect.TypeOf(EventFilter{})
	modelBuilder := builder.NewDefaultModelBuilder(generateId, modelType, "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt", "userId")
	searchHandler := search.NewSearchHandler(eventService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewHandler(eventService, modelType, modelBuilder, logError, validate)
	return &EventHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: eventService}
}

func (h *EventHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
