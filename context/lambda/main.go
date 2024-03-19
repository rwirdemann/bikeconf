package main

import "github.com/rwirdemann/bikeage/context"

func main() {
	bikeRepository := context.NewPostgresAdapter()
	sat := context.NewSATAdapter()
	configurationService := context.NewConfigurationService(bikeRepository, sat)
	lambda := context.NewLambdaAdapter(configurationService)
	lambda.Start()
}
