package context

import (
	"github.com/rwirdemann/bikeage/application/domain"
	"github.com/rwirdemann/bikeage/ports/in"
)

type ConfigurationService struct{}

func NewConfigurationService() in.ConfigurationService {
	return ConfigurationService{}
}

func (ConfigurationService) GetConfigurations() ([]domain.Configuration, error) {
	// TODO implement me
	panic("implement me")
}

func (ConfigurationService) ValidateConfiguration(c domain.Configuration) bool {
	// TODO implement me
	panic("implement me")
}

func (ConfigurationService) ResolveConfiguration(c domain.Configuration) ([]domain.Solution, error) {
	// TODO implement me
	panic("implement me")
}
