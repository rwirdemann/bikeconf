package out

import "github.com/rwirdemann/bikeage/application/domain"

type BikeRepository interface {
	GetAll() ([]domain.Configuration, error)
}
