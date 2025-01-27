package domain

import (
	"fmt"
	"nilus-challenge-backend/internal/domain/forecast"
	"nilus-challenge-backend/internal/domain/locality"
	"nilus-challenge-backend/internal/domain/waves"
)

type WeatherClient interface {
	GetLocalities(cityName string) ([]locality.Locality, error)
	GetCityForecast(cityID string) ([]forecast.Forecast, error)
	GetWaveForecast(cityID string, day int) ([]waves.WaveForecast, error)
}

type WeatherService struct {
	WeatherClient WeatherClient
}

func NewWeatherService(wc WeatherClient) *WeatherService {
	return &WeatherService{
		WeatherClient: wc,
	}
}

func (ws *WeatherService) GetLocalities(cityName string) ([]locality.Locality, error) {
	if cityName == "" {
		return nil, fmt.Errorf("el nombre de la ciudad no puede estar vacío")
	}

	localities, err := ws.WeatherClient.GetLocalities(cityName)
	if err != nil {
		return nil, fmt.Errorf("error al obtener localidades desde la API: %w", err)
	}

	return localities, nil
}

func (ws *WeatherService) GetCityForecast(cityID string) ([]forecast.Forecast, error) {
	if cityID == "" {
		return nil, fmt.Errorf("el ID de la ciudad no puede estar vacío")
	}

	forecasts, err := ws.WeatherClient.GetCityForecast(cityID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener previsión desde la API: %w", err)
	}

	return forecasts, nil
}

func (ws *WeatherService) GetWaveForecast(cityID string, day int) ([]waves.WaveForecast, error) {
	if cityID == "" {
		return nil, fmt.Errorf("el ID de la ciudad no puede estar vacío")
	}

	waveForecasts, err := ws.WeatherClient.GetWaveForecast(cityID, day)
	if err != nil {
		return nil, fmt.Errorf("error al obtener previsión de olas desde la API: %w", err)
	}

	return waveForecasts, nil
}
