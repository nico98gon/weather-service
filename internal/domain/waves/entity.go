package waves

type WaveForecast struct {
	Period        string  `xml:"dia"`
	Agitation     string  `xml:"agitacao"`
	Height        float64 `xml:"altura"`
	Direction     string  `xml:"direcao"`
	WindSpeed     float64 `xml:"vento"`
	WindDirection string  `xml:"vento_dir"`
}

type CityWaveForecast struct {
	Name      string        `xml:"nome"`
	State     string        `xml:"uf"`
	UpdatedAt string        `xml:"atualizacao"`
	Morning   WaveForecast  `xml:"manha"`
	Afternoon WaveForecast  `xml:"tarde"`
	Night     WaveForecast  `xml:"noite"`
}