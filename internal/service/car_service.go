package service

import (
	"errors"

	model "github.com/luminous479/car-rental-system/internal/model"
	"github.com/luminous479/car-rental-system/internal/repository"
)

type CarService struct {
	repo *repository.CarRepository
}

func NewCarService(repo *repository.CarRepository) *CarService {
	return &CarService{
		repo: repo,
	}
}

func (s *CarService) GetCar(id string) (*model.Car, error) {

	car, found := s.repo.GetByID(id)

	if !found {
		return nil, errors.New("car not found")
	}

	return car, nil
}
