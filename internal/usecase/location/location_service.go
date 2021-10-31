package location

import (
	"context"

	sv "github.com/core-go/service"
)

type LocationService interface {
	Load(ctx context.Context, id string) (*Location, error)
	Create(ctx context.Context, Location *Location) (int64, error)
	Update(ctx context.Context, Location *Location) (int64, error)
	Patch(ctx context.Context, Location map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewLocationService(repository sv.Repository) LocationService {
	return &locationService{repository: repository}
}

type locationService struct {
	repository sv.Repository
}

func (s *locationService) Load(ctx context.Context, id string) (*Location, error) {
	var Location Location
	ok, err := s.repository.LoadAndDecode(ctx, id, &Location)
	if !ok {
		return nil, err
	} else {
		return &Location, err
	}
}
func (s *locationService) Create(ctx context.Context, Location *Location) (int64, error) {
	return s.repository.Insert(ctx, Location)
}
func (s *locationService) Update(ctx context.Context, Location *Location) (int64, error) {
	return s.repository.Update(ctx, Location)
}
func (s *locationService) Patch(ctx context.Context, Location map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Location)
}
func (s *locationService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
