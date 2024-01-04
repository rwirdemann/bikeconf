package context

import (
	"github.com/rwirdemann/bikeage/application/domain"
)

type SATAdapter struct {
}

func NewSATAdapter() *SATAdapter {
	return &SATAdapter{}
}

func (S SATAdapter) Check(configuration domain.Configuration) bool {
	// TODO implement me
	panic("implement me")
}

func (S SATAdapter) Resolve(configuration domain.Configuration) []domain.Solution {
	// TODO implement me
	panic("implement me")
}
