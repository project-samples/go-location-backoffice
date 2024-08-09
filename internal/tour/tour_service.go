package tour

import "context"

type TourService interface {
	Load(ctx context.Context, id string) (*Tour, error)
	Create(ctx context.Context, tour *Tour) (int64, error)
	Update(ctx context.Context, tour *Tour) (int64, error)
	Patch(ctx context.Context, tour map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewTourService(repository TourRepository) TourService {
	return &TourUseCase{repository: repository}
}

type TourUseCase struct {
	repository TourRepository
}

func (s *TourUseCase) Load(ctx context.Context, id string) (*Tour, error) {
	return s.repository.Load(ctx, id)
}
func (s *TourUseCase) Create(ctx context.Context, Tour *Tour) (int64, error) {
	return s.repository.Create(ctx, Tour)
}
func (s *TourUseCase) Update(ctx context.Context, Tour *Tour) (int64, error) {
	return s.repository.Update(ctx, Tour)
}
func (s *TourUseCase) Patch(ctx context.Context, Tour map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Tour)
}
func (s *TourUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
