package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rwirdemann/bikeage/context"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	bikeRepo := context.NewPostgresAdapter()
	configurationService := context.NewConfigurationService(bikeRepo)
	httpAdapter := context.NewHTTPAdapter(configurationService)
	router.HandleFunc("/configurations", httpAdapter.GetConfigurations()).Methods("GET")
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
