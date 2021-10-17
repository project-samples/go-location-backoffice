package tour

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/model-builder"
)

type TourHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	Service TourService
}

func NewTourHandler(tourService TourService, generateId func(context.Context) (string, error), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string)) *TourHandler {
	modelType := reflect.TypeOf(Tour{})
	searchModelType := reflect.TypeOf(TourFilter{})
	modelBuilder := builder.NewDefaultModelBuilder(generateId, modelType, "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt", "userId")
	searchHandler := search.NewSearchHandler(tourService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewHandler(tourService, modelType, modelBuilder, logError, validate)
	return &TourHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: tourService}
}

func (h *TourHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
