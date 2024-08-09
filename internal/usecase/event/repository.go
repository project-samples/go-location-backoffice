package event

import "context"

type EventRepository interface {
	Load(ctx context.Context, id string) (*Event, error)
	Create(ctx context.Context, event *Event) (int64, error)
	Update(ctx context.Context, event *Event) (int64, error)
	Patch(ctx context.Context, event map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
