package repository

import (
	"strconv"

	models "github.com/luminous479/car-rental-system/internal/model"
)

type CarRepository struct {
	cars []models.Car
}

func NewCarRepository() *CarRepository {
	return &CarRepository{
		cars: []models.Car{
			{ID: 1, Brand: "Toyota", Model: "Camry", Year: 2020, DailyRate: 50.0, Available: true},
			{ID: 2, Brand: "Honda", Model: "Civic", Year: 2019, DailyRate: 45.0, Available: false},
			{ID: 3, Brand: "Ford", Model: "Mustang", Year: 2021, DailyRate: 70.0, Available: true},
		},
	}
}

func (r *CarRepository) GetByID(id string) (*models.Car, bool) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, false
	}

	for i := range r.cars {
		if r.cars[i].ID == idInt {
			return &r.cars[i], true
		}
	}

	return nil, false
}
