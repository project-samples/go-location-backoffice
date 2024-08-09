package bookable

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

func NewBookableHandler(
	find search.Search[Bookable, *BookableFilter],
	bookableService BookableService,
	logError core.Log,
	validate v.Validate[*Bookable],
	tracking b.TrackingConfig,
	writeLog core.WriteLog,
	action *core.ActionConfig,
) *BookableHandler {
	bookableType := reflect.TypeOf(Bookable{})
	builder := b.NewBuilderByConfig[Bookable](nil, tracking)
	params := core.CreateParams(bookableType, logError, action, writeLog)
	searchHandler := search.NewSearchHandler[Bookable, *BookableFilter](find, logError, nil)
	return &BookableHandler{SearchHandler: searchHandler, service: bookableService, validate: validate, builder: builder, Params: params}
}

type BookableHandler struct {
	service BookableService
	*search.SearchHandler[Bookable, *BookableFilter]
	*core.Params
	validate v.Validate[*Bookable]
	builder  hdl.Builder[Bookable]
}

func (h *BookableHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		bookable, err := h.service.Load(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}
		if bookable == nil {
			core.JSON(w, http.StatusNotFound, bookable)
		} else {
			core.JSON(w, http.StatusOK, bookable)
		}
	}
}
func (h *BookableHandler) Create(w http.ResponseWriter, r *http.Request) {
	var bookable Bookable
	er1 := hdl.Decode(w, r, &bookable, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &bookable)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &bookable)
			if er3 != nil {
				h.Error(r.Context(), er3.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, er3.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("created '%s'", bookable.Id))
				core.JSON(w, http.StatusCreated, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", bookable.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *BookableHandler) Update(w http.ResponseWriter, r *http.Request) {
	bookable, err := hdl.DecodeAndCheckId[Bookable](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &bookable)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &bookable)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("%s '%s'", h.Action.Update, bookable.Id))
				core.JSON(w, http.StatusOK, bookable)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("not found '%s'", bookable.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", bookable.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *BookableHandler) Patch(w http.ResponseWriter, r *http.Request) {
	r, bookable, jsonBookable, err := hdl.BuildMapAndCheckId[Bookable](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &bookable)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, err := h.service.Patch(r.Context(), jsonBookable)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, true, fmt.Sprintf("%s '%s'", h.Action.Patch, bookable.Id))
				core.JSON(w, http.StatusOK, res)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("not found '%s'", bookable.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("conflict '%s'", bookable.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *BookableHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
