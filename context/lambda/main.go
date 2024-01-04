package main

import "github.com/rwirdemann/bikeage/context"

func main() {
	configurationService := context.NewConfigurationService(nil)
	lambda := context.NewLambdaAdapter(configurationService)
	lambda.Start()
}
