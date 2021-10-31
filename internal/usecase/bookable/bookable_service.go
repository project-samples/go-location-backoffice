package bookable

import (
	"context"

	sv "github.com/core-go/service"
)

type BookableService interface {
	Load(ctx context.Context, id string) (*Bookable, error)
	Create(ctx context.Context, Bookable *Bookable) (int64, error)
	Update(ctx context.Context, Bookable *Bookable) (int64, error)
	Patch(ctx context.Context, Bookable map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewBookableService(repository sv.Repository) BookableService {
	return &bookableService{repository: repository}
}

type bookableService struct {
	repository sv.Repository
}

func (s *bookableService) Load(ctx context.Context, id string) (*Bookable, error) {
	var Bookable Bookable
	ok, err := s.repository.LoadAndDecode(ctx, id, &Bookable)
	if !ok {
		return nil, err
	} else {
		return &Bookable, err
	}
}
func (s *bookableService) Create(ctx context.Context, Bookable *Bookable) (int64, error) {
	return s.repository.Insert(ctx, Bookable)
}
func (s *bookableService) Update(ctx context.Context, Bookable *Bookable) (int64, error) {
	return s.repository.Update(ctx, Bookable)
}
func (s *bookableService) Patch(ctx context.Context, Bookable map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Bookable)
}
func (s *bookableService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
