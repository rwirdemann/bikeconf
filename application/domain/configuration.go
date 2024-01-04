package domain

type Configuration struct {
	Id int `json:"id"`
}

func (c Configuration) Apply(solutions []Solution) Configuration {
	return Configuration{}
}
