package out

import "github.com/rwirdemann/bikeage/application/domain"

type SATResolver interface {
	Check(configuration domain.Configuration) bool
	Resolve(configuration domain.Configuration) []domain.Solution
}
