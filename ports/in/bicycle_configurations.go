package in

import "github.com/rwirdemann/bikeage/application/domain"

type BicyleConfigurations interface {
	GetAll() ([]domain.Configuration, error)
}
