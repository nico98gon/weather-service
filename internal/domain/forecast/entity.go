package forecast

type Forecast struct {
	Dia     string `xml:"dia"`
	Tempo   string `xml:"tempo"`
	Maxima  string `xml:"maxima"`
	Minima  string `xml:"minima"`
	IUV     string `xml:"iuv"`
}