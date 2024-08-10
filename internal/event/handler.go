package event

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

func NewEventHandler(
	find search.Search[Event, *EventFilter],
	eventService EventService,
	logError core.Log,
	validate v.Validate[*Event],
	tracking b.TrackingConfig,
	writeLog core.WriteLog,
	action *core.ActionConfig,
) *EventHandler {
	eventType := reflect.TypeOf(Event{})
	builder := b.NewBuilderByConfig[Event](nil, tracking)
	params := core.CreateParams(eventType, logError, action, writeLog)
	searchHandler := search.NewSearchHandler[Event, *EventFilter](find, logError, nil)
	return &EventHandler{SearchHandler: searchHandler, service: eventService, validate: validate, builder: builder, Params: params}
}

type EventHandler struct {
	service EventService
	*search.SearchHandler[Event, *EventFilter]
	*core.Params
	validate v.Validate[*Event]
	builder  hdl.Builder[Event]
}

func (h *EventHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := core.GetRequiredParam(w, r)
	if len(id) > 0 {
		event, err := h.service.Load(r.Context(), id)
		if err != nil {
			h.Error(r.Context(), err.Error())
			http.Error(w, core.InternalServerError, http.StatusInternalServerError)
			return
		}
		if event == nil {
			core.JSON(w, http.StatusNotFound, event)
		} else {
			core.JSON(w, http.StatusOK, event)
		}
	}
}
func (h *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	event, er1 := hdl.Decode[Event](w, r, &event, h.builder.Create)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &event)
		if !core.HasError(w, r, errors, er2, h.Error, h.Log, h.Resource, h.Action.Create) {
			res, er3 := h.service.Create(r.Context(), &event)
			if er3 != nil {
				h.Error(r.Context(), er3.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, er3.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("created '%s'", event.Id))
				core.JSON(w, http.StatusCreated, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", event.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *EventHandler) Update(w http.ResponseWriter, r *http.Request) {
	event, err := hdl.DecodeAndCheckId[Event](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &event)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Update) {
			res, err := h.service.Update(r.Context(), &event)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Update, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, true, fmt.Sprintf("%s '%s'", h.Action.Update, event.Id))
				core.JSON(w, http.StatusOK, event)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("not found '%s'", event.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Update, false, fmt.Sprintf("conflict '%s'", event.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *EventHandler) Patch(w http.ResponseWriter, r *http.Request) {
	r, event, jsonEvent, err := hdl.BuildMapAndCheckId[Event](w, r, h.Keys, h.Indexes, h.builder.Update)
	if err == nil {
		errors, err := h.validate(r.Context(), &event)
		if !core.HasError(w, r, errors, err, h.Error, h.Log, h.Resource, h.Action.Patch) {
			res, err := h.service.Patch(r.Context(), jsonEvent)
			if err != nil {
				h.Error(r.Context(), err.Error())
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, err.Error())
				http.Error(w, core.InternalServerError, http.StatusInternalServerError)
				return
			}

			if res > 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, true, fmt.Sprintf("%s '%s'", h.Action.Patch, event.Id))
				core.JSON(w, http.StatusOK, res)
			} else if res == 0 {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("not found '%s'", event.Id))
				core.JSON(w, http.StatusNotFound, res)
			} else {
				h.Log(r.Context(), h.Resource, h.Action.Patch, false, fmt.Sprintf("conflict '%s'", event.Id))
				core.JSON(w, http.StatusConflict, res)
			}
		}
	}
}
func (h *EventHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
