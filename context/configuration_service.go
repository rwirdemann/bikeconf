package context

import (
	"github.com/rwirdemann/bikeage/application/domain"
	"github.com/rwirdemann/bikeage/ports/in"
	out2 "github.com/rwirdemann/bikeage/ports/out"
)

type ConfigurationService struct{}

func NewConfigurationService(bikeRepo out2.BikeRepository, sat out2.SATResolver) in.ConfigurationService {
	return ConfigurationService{}
}

func (cs ConfigurationService) GetAll() ([]domain.Configuration, error) {
	// TODO implement me
	panic("implement me")
}

func (cs ConfigurationService) ValidateConfiguration(c domain.Configuration) bool {
	// TODO implement me
	panic("implement me")
}

func (cs ConfigurationService) ResolveConfiguration(c domain.Configuration) []domain.Solution {
	// TODO implement me
	panic("implement me")
}
