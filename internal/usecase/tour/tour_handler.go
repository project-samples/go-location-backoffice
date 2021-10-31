package tour

import (
	"context"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"github.com/core-go/service/builder"
	"net/http"
	"reflect"
)

type TourHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewTourHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service TourService, generateId func(context.Context) (string, error), status sv.StatusConfig, logError func(context.Context, string), validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error), tracking builder.TrackingConfig, action *sv.ActionConfig, writeLog func(context.Context, string, string, bool, string) error) TourHandler {
	searchModelType := reflect.TypeOf(TourFilter{})
	modelType := reflect.TypeOf(Tour{})
	builder := builder.NewBuilderWithIdAndConfig(generateId, modelType, tracking)
	patchHandler, params := sv.CreatePatchAndParams(modelType, &status, logError, service.Patch, validate, builder.Patch, action, writeLog)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, params.Log)
	return &tourHandler{service: service, builder: builder, PatchHandler: patchHandler, SearchHandler: searchHandler, Params: params}
}

type tourHandler struct {
	service TourService
	builder sv.Builder
	*sv.PatchHandler
	*search.SearchHandler
	*sv.Params
}

func (h *tourHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *tourHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tour Tour
	er1 := sv.Decode(w, r, &tour, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &tour)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &tour)
			sv.AfterCreated(w, r, &tour, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *tourHandler) Update(w http.ResponseWriter, r *http.Request) {
	var tour Tour
	er1 := sv.DecodeAndCheckId(w, r, &tour, h.Keys, h.Indexes, h.builder.Update)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &tour)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &tour)
			sv.HandleResult(w, r, &tour, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *tourHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
