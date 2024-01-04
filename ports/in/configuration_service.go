package in

import "github.com/rwirdemann/bikeage/application/domain"

type ConfigurationService interface {
	GetAll() ([]domain.Configuration, error)
	ValidateConfiguration(c domain.Configuration) bool
	ResolveConfiguration(c domain.Configuration) []domain.Solution
}
