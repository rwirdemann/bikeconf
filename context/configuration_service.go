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
	if c.Manufacturer == "Rose" && c.Model == "Backroad" {
		if c.Shift.Manufacturer == "SRAM" {
			return false
		}
	}

	return true
}

func (cs ConfigurationService) ResolveConfiguration(c domain.Configuration) []domain.Solution {
	return []domain.Solution{}
}
