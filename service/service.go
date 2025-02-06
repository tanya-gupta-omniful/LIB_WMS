package service

import (
	"WMS/internal/domain"
	"WMS/repo"
	"context"
)

type Service interface {
	FetchHubs(ctx context.Context) []domain.Hub
	FetchHubByID(ctx context.Context, id int) (domain.Hub, error)
	FetchHubByTenantID(ctx context.Context, id int)(domain.Hub, error)
}

type service struct {
	repo repo.Repository
}

// NewService is the constructor function to create a new instance of ConcreteService.
func NewService(r repo.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FetchHubs(ctx context.Context) []domain.Hub {
	return s.repo.GetAllHubs(ctx)
}

func (s *service) FetchHubByID(ctx context.Context, id int) (domain.Hub, error) { // Implementing the new method
	hub, err := s.repo.GetHubByID(ctx, id)
	if err != nil {
		return domain.Hub{}, err // Return an empty Hub and the error
	}
	return hub, nil
}
func (s *service)FetchHubByTenantID(ctx context.Context, id int)(domain.Hub, error){
	hub, err := s.repo.GetHubByTenantID(ctx,id)
	if err != nil {
		return domain.Hub{}, err // Return an empty Hub and the error
	}
	return hub, nil
	

	
}