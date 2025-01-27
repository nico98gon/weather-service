package infrastructure

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"nilus-challenge-backend/internal/domain/forecast"
	"nilus-challenge-backend/internal/domain/locality"
	"nilus-challenge-backend/internal/domain/waves"

	"golang.org/x/net/html/charset"
)

type WeatherClient struct {
	BaseURL string
	Client  *http.Client
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{
		BaseURL: "http://servicos.cptec.inpe.br/XML",
		Client:  &http.Client{},
	}
}

func (wc *WeatherClient) GetLocalities(cityName string) ([]locality.Locality, error) {
	url := fmt.Sprintf("%s/listaCidades?city=%s", wc.BaseURL, cityName)
	resp, err := wc.Client.Get(url)
	if err != nil {
			return nil, fmt.Errorf("error al hacer la solicitud: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status de respuesta no válido: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
			return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %w", err)
	}
	
	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = charset.NewReaderLabel
	
	var result struct {
			Cidades []locality.Locality `xml:"cidade"`
	}
	if err := decoder.Decode(&result); err != nil {
			return nil, fmt.Errorf("error al parsear la respuesta: %w", err)
	}

	return result.Cidades, nil
}

func (wc *WeatherClient) GetCityForecast(cityID string) ([]forecast.Forecast, error) {
	url := fmt.Sprintf("%s/cidade/%s/previsao.xml", wc.BaseURL, cityID)
	resp, err := wc.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status de respuesta no válido: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = charset.NewReaderLabel

	var result struct {
		Previsoes []forecast.Forecast `xml:"previsao"`
	}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("error al parsear la respuesta: %w", err)
	}

	return result.Previsoes, nil
}

func (wc *WeatherClient) GetWaveForecast(cityID string, day int) ([]waves.WaveForecast, error) {
	url := fmt.Sprintf("%s/cidade/%s/dia/%d/ondas.xml", wc.BaseURL, cityID, day)
	resp, err := wc.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status de respuesta no válido: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = charset.NewReaderLabel

	var result struct {
		Nome      string              `xml:"nome"`
		UF        string              `xml:"uf"`
		UpdatedAt string              `xml:"atualizacao"`
		Morning   waves.WaveForecast  `xml:"manha"`
		Afternoon waves.WaveForecast  `xml:"tarde"`
		Night     waves.WaveForecast  `xml:"noite"`
	}

	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("error al parsear la respuesta: %w", err)
	}

	waveForecasts := []waves.WaveForecast{
		result.Morning,
		result.Afternoon,
		result.Night,
	}

	return waveForecasts, nil
}