package location

import (
	"context"
	"github.com/core-go/search"
	sv "github.com/core-go/service"
	"net/http"
	"reflect"
)

type LocationHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewLocationHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service LocationService, status sv.StatusConfig, logError func(context.Context, string), validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error), action *sv.ActionConfig) LocationHandler {
	searchModelType := reflect.TypeOf(LocationFilter{})
	modelType := reflect.TypeOf(Location{})
	params := sv.CreateParams(modelType, &status, logError, validate, action)
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, params.Log)
	return &locationHandler{service: service, SearchHandler: searchHandler, Params: params}
}

type locationHandler struct {
	service LocationService
	*search.SearchHandler
	*sv.Params
}

func (h *locationHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *locationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var location Location
	er1 := sv.Decode(w, r, &location)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &location)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Create(r.Context(), &location)
			sv.AfterCreated(w, r, &location, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}
func (h *locationHandler) Update(w http.ResponseWriter, r *http.Request) {
	var location Location
	er1 := sv.DecodeAndCheckId(w, r, &location, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &location)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Update) {
			result, er3 := h.service.Update(r.Context(), &location)
			sv.HandleResult(w, r, &location, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Update)
		}
	}
}
func (h *locationHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var location Location
	r, json, er1 := sv.BuildMapAndCheckId(w, r, &location, h.Keys, h.Indexes)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &location)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Patch) {
			result, er3 := h.service.Patch(r.Context(), json)
			sv.HandleResult(w, r, json, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Patch)
		}
	}
}
func (h *locationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, h.Log, h.Resource, h.Action.Delete)
	}
}
