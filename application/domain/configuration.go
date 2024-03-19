package domain

type Shift struct {
	Manufacturer string `json:"manufacturer"`
	Series       string `json:"series"`
}

type Configuration struct {
	Id           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Shift        Shift  `json:"shift"`
}

func (c Configuration) Apply(solutions []Solution) Configuration {
	return Configuration{}
}
