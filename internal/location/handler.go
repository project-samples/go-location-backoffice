package location

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/core-go/core"
	hdl "github.com/core-go/core/handler"
	b "github.com/core-go/core/handler/builder"
	v "github.com/core-go/core/validator"
	search "github.com/core-go/search/handler"
)

func NewLocationHandler(
	find search.Search[Location, *LocationFilter],
	locationService LocationService,
	logError core.Log,
	validate v.Validate[*Location],
	tracking b.TrackingConfig,
	writeLog core.WriteLog,
	action *core.ActionConfig,
) *LocationHandler {
	locationType := reflect.TypeOf(Location{})
	builder := b.NewBuilderByConfig[Location](nil, tracking)
	params := core.CreateParams(locationType, logError, action, writeLog)
	searchHandler := search.NewSearchHandler[Location, *LocationFilter](find, logError, nil)
	return &LocationHandler{SearchHandler: searchHandler, service: locationService, validate: validate, builder: builder, Params: params}
}

type LocationHandler struct {
	service LocationService
	*search.SearchHandler[Location, *LocationFilter]
	*core.Params
	validate v.Validate[*Location]
	builder  hdl.Builder[Location]
}

func (h *LocationHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		location, err := h.service.Load(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}
		if location == nil {
			core.JSON(w, http.StatusNotFound, location)
		} else {
			core.JSON(w, http.StatusOK, location)
		}
	}
}
func (h *LocationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var location Location
	er1 := hdl.Decode(w, r, &location, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &location)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &location)
			if er3 != nil {
				h.Error(r.Context(), er3.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, er3.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("created '%s'", location.Id))
				core.JSON(w, http.StatusCreated, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", location.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *LocationHandler) Update(w http.ResponseWriter, r *http.Request) {
	location, err := hdl.DecodeAndCheckId[Location](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &location)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &location)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("%s '%s'", h.Action.Update, location.Id))
				core.JSON(w, http.StatusOK, location)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("not found '%s'", location.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", location.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *LocationHandler) Patch(w http.ResponseWriter, r *http.Request) {
	r, location, jsonLocation, err := hdl.BuildMapAndCheckId[Location](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &location)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, err := h.service.Patch(r.Context(), jsonLocation)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, true, fmt.Sprintf("%s '%s'", h.Action.Patch, location.Id))
				core.JSON(w, http.StatusOK, res)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("not found '%s'", location.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("conflict '%s'", location.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *LocationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		res, err := h.service.Delete(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}

		if res > 0 {
			h.Log(r.Context(), h.Resource, h.Action.Delete, true, fmt.Sprintf("%s '%s'", h.Action.Delete, id))
			core.JSON(w, http.StatusOK, res)
		} else if res == 0 {
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, fmt.Sprintf("not found '%s'", id))
			core.JSON(w, http.StatusNotFound, res)
		} else {
			h.Log(r.Context(), h.Resource, h.Action.Delete, false, fmt.Sprintf("conflict '%s'", id))
			core.JSON(w, http.StatusConflict, res)
		}
	}
}
