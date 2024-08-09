package bookable

import "context"

type BookableService interface {
	Load(ctx context.Context, id string) (*Bookable, error)
	Create(ctx context.Context, bookable *Bookable) (int64, error)
	Update(ctx context.Context, bookable *Bookable) (int64, error)
	Patch(ctx context.Context, bookable map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewBookableService(repository BookableRepository) BookableService {
	return &BookableUseCase{repository: repository}
}

type BookableUseCase struct {
	repository BookableRepository
}

func (s *BookableUseCase) Load(ctx context.Context, id string) (*Bookable, error) {
	return s.repository.Load(ctx, id)
}
func (s *BookableUseCase) Create(ctx context.Context, Bookable *Bookable) (int64, error) {
	return s.repository.Create(ctx, Bookable)
}
func (s *BookableUseCase) Update(ctx context.Context, Bookable *Bookable) (int64, error) {
	return s.repository.Update(ctx, Bookable)
}
func (s *BookableUseCase) Patch(ctx context.Context, Bookable map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Bookable)
}
func (s *BookableUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
