package tour

import "context"

type TourRepository interface {
	Load(ctx context.Context, id string) (*Tour, error)
	Create(ctx context.Context, tour *Tour) (int64, error)
	Update(ctx context.Context, tour *Tour) (int64, error)
	Patch(ctx context.Context, tour map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
