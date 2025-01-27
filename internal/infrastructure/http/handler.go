package handler

import (
	"net/http"
	"nilus-challenge-backend/internal/domain"
	"nilus-challenge-backend/internal/infrastructure/middleware"

	"strconv"
)

type WeatherHandler struct {
	Service *domain.WeatherService
}

func NewWeatherHandler(ws *domain.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		Service: ws,
	}
}

func (wh *WeatherHandler) HandleGetLocalities(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("city")
	if cityName == "" {
		middleware.ErrorResponse(w, http.StatusBadRequest, "Falta el parámetro 'city'")
		return
	}

	localities, err := wh.Service.GetLocalities(cityName)
	if err != nil {
		middleware.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	middleware.SuccessResponse(w, localities)
}

func (wh *WeatherHandler) HandleGetCityForecast(w http.ResponseWriter, r *http.Request) {
	cityID := r.URL.Query().Get("city_id")
	if cityID == "" {
		middleware.ErrorResponse(w, http.StatusBadRequest, "Falta el parámetro 'city_id'")
		return
	}

	forecasts, err := wh.Service.GetCityForecast(cityID)
	if err != nil {
		middleware.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	middleware.SuccessResponse(w, forecasts)
}

func (wh *WeatherHandler) HandleGetWaveForecast(w http.ResponseWriter, r *http.Request) {
	cityID := r.URL.Query().Get("city_id")
	dayStr := r.URL.Query().Get("day")
	if cityID == "" || dayStr == "" {
		middleware.ErrorResponse(w, http.StatusBadRequest, "Faltan los parámetros 'city_id' o 'day'")
		return
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		middleware.ErrorResponse(w, http.StatusBadRequest, "El parámetro 'day' debe ser un número")
		return
	}

	waveForecasts, err := wh.Service.GetWaveForecast(cityID, day)
	if err != nil {
		middleware.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	middleware.SuccessResponse(w, waveForecasts)
}
