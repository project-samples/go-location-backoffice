package bookable

import "context"

type BookableRepository interface {
	Load(ctx context.Context, id string) (*Bookable, error)
	Create(ctx context.Context, bookable *Bookable) (int64, error)
	Update(ctx context.Context, bookable *Bookable) (int64, error)
	Patch(ctx context.Context, bookable map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
