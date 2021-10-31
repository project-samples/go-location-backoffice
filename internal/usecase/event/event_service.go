package event

import (
	"context"

	sv "github.com/core-go/service"
)

type EventService interface {
	Load(ctx context.Context, id string) (*Event, error)
	Create(ctx context.Context, Event *Event) (int64, error)
	Update(ctx context.Context, Event *Event) (int64, error)
	Patch(ctx context.Context, Event map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewEventService(repository sv.Repository) EventService {
	return &eventService{repository: repository}
}

type eventService struct {
	repository sv.Repository
}

func (s *eventService) Load(ctx context.Context, id string) (*Event, error) {
	var Event Event
	ok, err := s.repository.LoadAndDecode(ctx, id, &Event)
	if !ok {
		return nil, err
	} else {
		return &Event, err
	}
}
func (s *eventService) Create(ctx context.Context, Event *Event) (int64, error) {
	return s.repository.Insert(ctx, Event)
}
func (s *eventService) Update(ctx context.Context, Event *Event) (int64, error) {
	return s.repository.Update(ctx, Event)
}
func (s *eventService) Patch(ctx context.Context, Event map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Event)
}
func (s *eventService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
