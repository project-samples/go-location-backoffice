package event

import "context"

type EventService interface {
	Load(ctx context.Context, id string) (*Event, error)
	Create(ctx context.Context, event *Event) (int64, error)
	Update(ctx context.Context, event *Event) (int64, error)
	Patch(ctx context.Context, event map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewEventService(repository EventRepository) EventService {
	return &EventUseCase{repository: repository}
}

type EventUseCase struct {
	repository EventRepository
}

func (s *EventUseCase) Load(ctx context.Context, id string) (*Event, error) {
	return s.repository.Load(ctx, id)
}
func (s *EventUseCase) Create(ctx context.Context, Event *Event) (int64, error) {
	return s.repository.Create(ctx, Event)
}
func (s *EventUseCase) Update(ctx context.Context, Event *Event) (int64, error) {
	return s.repository.Update(ctx, Event)
}
func (s *EventUseCase) Patch(ctx context.Context, Event map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Event)
}
func (s *EventUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
