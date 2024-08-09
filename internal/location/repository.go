package location

import "context"

type LocationRepository interface {
	Load(ctx context.Context, id string) (*Location, error)
	Create(ctx context.Context, location *Location) (int64, error)
	Update(ctx context.Context, location *Location) (int64, error)
	Patch(ctx context.Context, location map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
