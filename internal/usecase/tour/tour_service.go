package tour

import (
	"context"

	sv "github.com/core-go/service"
)

type TourService interface {
	Load(ctx context.Context, id string) (*Tour, error)
	Create(ctx context.Context, Tour *Tour) (int64, error)
	Update(ctx context.Context, Tour *Tour) (int64, error)
	Patch(ctx context.Context, Tour map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewTourService(repository sv.Repository) TourService {
	return &tourService{repository: repository}
}

type tourService struct {
	repository sv.Repository
}

func (s *tourService) Load(ctx context.Context, id string) (*Tour, error) {
	var Tour Tour
	ok, err := s.repository.LoadAndDecode(ctx, id, &Tour)
	if !ok {
		return nil, err
	} else {
		return &Tour, err
	}
}
func (s *tourService) Create(ctx context.Context, Tour *Tour) (int64, error) {
	return s.repository.Insert(ctx, Tour)
}
func (s *tourService) Update(ctx context.Context, Tour *Tour) (int64, error) {
	return s.repository.Update(ctx, Tour)
}
func (s *tourService) Patch(ctx context.Context, Tour map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Tour)
}
func (s *tourService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
