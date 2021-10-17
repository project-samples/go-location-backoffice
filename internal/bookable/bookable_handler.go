package bookable

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/model-builder"
)

type BookableHandler struct {
	*sv.GenericHandler
	*search.SearchHandler
	Service BookableService
}

func NewBookableHandler(bookableService BookableService, generateId func(context.Context) (string, error), validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string)) *BookableHandler {
	modelType := reflect.TypeOf(Bookable{})
	searchModelType := reflect.TypeOf(BookableFilter{})
	modelBuilder := builder.NewDefaultModelBuilder(generateId, modelType, "CreatedBy", "CreatedAt", "UpdatedBy", "UpdatedAt", "userId")
	searchHandler := search.NewSearchHandler(bookableService.Search, modelType, searchModelType, logError, nil)
	genericHandler := sv.NewHandler(bookableService, modelType, modelBuilder, logError, validate)
	return &BookableHandler{GenericHandler: genericHandler, SearchHandler: searchHandler, Service: bookableService}
}

func (h *BookableHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sv.JSON(w, http.StatusOK, result)
}
