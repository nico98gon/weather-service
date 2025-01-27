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

	http.HandleFunc(api+"/localities", weatherHandler.HandleGetLocalities)
	http.HandleFunc(api+"/city-forecast", weatherHandler.HandleGetCityForecast)
	http.HandleFunc(api+"/wave-forecast", weatherHandler.HandleGetWaveForecast)
}