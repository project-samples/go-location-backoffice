package location

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/model-builder"
)

type LocationHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	Service LocationService
}

func NewLocationHandler(locationService LocationService, generateId func(context.Context) (string, error), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string)) *LocationHandler {
	modelType := reflect.TypeOf(Location{})
	searchModelType := reflect.TypeOf(LocationFilter{})
	modelBuilder := builder.NewDefaultModelBuilder(generateId, modelType, "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt", "userId")
	searchHandler := search.NewSearchHandler(locationService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewHandler(locationService, modelType, modelBuilder, logError, validate)
	return &LocationHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: locationService}
}

func (h *LocationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
