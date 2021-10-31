package bookable

import (
	"context"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/builder"
	"net/http"
	"reflect"
)

type BookableHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewBookableHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service BookableService, generateId func(context.Context) (string, error), status sv.StatusConfig, logError func(context.Context, string), validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error), tracking builder.TrackingConfig, action *sv.ActionConfig, writeLog func(context.Context, string, string, bool, string) error) BookableHandler {
	searchModelType := reflect.TypeOf(BookableFilter{})
	modelType := reflect.TypeOf(Bookable{})
	builder := builder.NewBuilderWithIdAndConfig(generateId, modelType, tracking)
	patchHandler, params := sv.CreatePatchAndParams(modelType, &status, logError, service.Patch, validate, builder.Patch, action, writeLog)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, params.Log)
	return &bookableHandler{service: service, builder: builder, PatchHandler: patchHandler, SearchHandler: searchHandler, Params: params}
}

type bookableHandler struct {
	service BookableService
	builder sv.Builder
	*sv.PatchHandler
	*search.SearchHandler
	*sv.Params
}

func (h *bookableHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *bookableHandler) Create(w http.ResponseWriter, r *http.Request) {
	var bookable Bookable
	er1 := sv.Decode(w, r, &bookable, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &bookable)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &bookable)
			sv.AfterCreated(w, r, &bookable, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *bookableHandler) Update(w http.ResponseWriter, r *http.Request) {
	var bookable Bookable
	er1 := sv.DecodeAndCheckId(w, r, &bookable, h.Keys, h.Indexes, h.builder.Update)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &bookable)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &bookable)
			sv.HandleResult(w, r, &bookable, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *bookableHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
