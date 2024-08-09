package tour

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

func NewTourHandler(
	find search.Search[Tour, *TourFilter],
	tourService TourService,
	logError core.Log,
	validate v.Validate[*Tour],
	tracking b.TrackingConfig,
	writeLog core.WriteLog,
	action *core.ActionConfig,
) *TourHandler {
	tourType := reflect.TypeOf(Tour{})
	builder := b.NewBuilderByConfig[Tour](nil, tracking)
	params := core.CreateParams(tourType, logError, action, writeLog)
	searchHandler := search.NewSearchHandler[Tour, *TourFilter](find, logError, nil)
	return &TourHandler{SearchHandler: searchHandler, service: tourService, validate: validate, builder: builder, Params: params}
}

type TourHandler struct {
	service TourService
	*search.SearchHandler[Tour, *TourFilter]
	*core.Params
	validate v.Validate[*Tour]
	builder  hdl.Builder[Tour]
}

func (h *TourHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		tour, err := h.service.Load(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}
		if tour == nil {
			core.JSON(w, http.StatusNotFound, tour)
		} else {
			core.JSON(w, http.StatusOK, tour)
		}
	}
}
func (h *TourHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tour Tour
	er1 := hdl.Decode(w, r, &tour, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &tour)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &tour)
			if er3 != nil {
				h.Error(r.Context(), er3.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, er3.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("created '%s'", tour.Id))
				core.JSON(w, http.StatusCreated, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", tour.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *TourHandler) Update(w http.ResponseWriter, r *http.Request) {
	tour, err := hdl.DecodeAndCheckId[Tour](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &tour)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &tour)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("%s '%s'", h.Action.Update, tour.Id))
				core.JSON(w, http.StatusOK, tour)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("not found '%s'", tour.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", tour.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *TourHandler) Patch(w http.ResponseWriter, r *http.Request) {
	r, tour, jsonTour, err := hdl.BuildMapAndCheckId[Tour](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &tour)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, err := h.service.Patch(r.Context(), jsonTour)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, true, fmt.Sprintf("%s '%s'", h.Action.Patch, tour.Id))
				core.JSON(w, http.StatusOK, res)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("not found '%s'", tour.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("conflict '%s'", tour.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *TourHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
