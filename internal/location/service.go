package location

import "context"

type LocationService interface {
	Load(ctx context.Context, id string) (*Location, error)
	Create(ctx context.Context, location *Location) (int64, error)
	Update(ctx context.Context, location *Location) (int64, error)
	Patch(ctx context.Context, location map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewLocationService(repository LocationRepository) LocationService {
	return &LocationUseCase{repository: repository}
}

type LocationUseCase struct {
	repository LocationRepository
}

func (s *LocationUseCase) Load(ctx context.Context, id string) (*Location, error) {
	return s.repository.Load(ctx, id)
}
func (s *LocationUseCase) Create(ctx context.Context, Location *Location) (int64, error) {
	return s.repository.Create(ctx, Location)
}
func (s *LocationUseCase) Update(ctx context.Context, Location *Location) (int64, error) {
	return s.repository.Update(ctx, Location)
}
func (s *LocationUseCase) Patch(ctx context.Context, Location map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Location)
}
func (s *LocationUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
