package application

import (
	"github.com/rwirdemann/bikeage/application/domain"
	"github.com/rwirdemann/bikeage/context"
	"github.com/rwirdemann/bikeage/mocks/github.com/rwirdemann/bikeage/ports/out"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolve(t *testing.T) {
	br := out.NewMockBikeRepository(t)
	sr := out.NewMockSATResolver(t)
	cs := context.NewConfigurationService(br, sr)
	fc := failedConfig()

	assert.False(t, cs.ValidateConfiguration(fc))
	solutions := cs.ResolveConfiguration(fc)
	vc := fc.Apply(solutions)
	assert.True(t, cs.ValidateConfiguration(vc))
}

func failedConfig() domain.Configuration {
	return domain.Configuration{
		Manufacturer: "Rose",
		Model:        "Backroad",
		Shift: domain.Shift{
			Manufacturer: "SRAM",
			Series:       "Red",
		},
	}
}
