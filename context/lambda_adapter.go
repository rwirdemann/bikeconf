package context

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rwirdemann/bikeage/ports/in"
	"net/http"
)

type LambdaAdapter struct {
	service in.ConfigurationService
}

func NewLambdaAdapter(service in.ConfigurationService) *LambdaAdapter {
	return &LambdaAdapter{service: service}
}

func (a LambdaAdapter) Start() {
	lambda.Start(
		func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			configurations, _ := a.service.GetAll()
			json := toJSON(configurations)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       string(json),
				Headers: map[string]string{
					"Access-Control-Allow-Origin":      "*",
					"Access-Control-Allow-Credentials": "true",
				},
			}, nil
		})
}
