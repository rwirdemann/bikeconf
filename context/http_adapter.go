package context

import (
	"github.com/rwirdemann/bikeage/application/domain"
	"github.com/rwirdemann/bikeage/ports/in"
	"net/http"
)

type HTTPAdapter struct {
	service in.ConfigurationService
}

func NewHTTPAdapter(service in.ConfigurationService) *HTTPAdapter {
	return &HTTPAdapter{service: service}
}

func (a HTTPAdapter) GetConfigurations() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		configurations, err := a.service.GetConfigurations()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		json := toJSON(configurations)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(json)
	}
}

func toJSON(configurations []domain.Configuration) []byte {
	return nil
}
