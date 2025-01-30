package infrastructure

import (
	"net/http"
	"nilus-challenge-backend/internal/domain"
	handler "nilus-challenge-backend/internal/infrastructure/http"
)

func StartHTTPServer() {
	weatherClient := NewWeatherClient()
	weatherService := domain.NewWeatherService(weatherClient)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	api := "/api/v1"

	http.HandleFunc("GET "+api+"/localities", weatherHandler.HandleGetLocalities)
	http.HandleFunc("GET "+api+"/city-forecast", weatherHandler.HandleGetCityForecast)
	http.HandleFunc("GET "+api+"/wave-forecast", weatherHandler.HandleGetWaveForecast)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}