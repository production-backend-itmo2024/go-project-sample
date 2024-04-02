package weather

type Weather struct {
	Temp int64
}

type DataseriesApiResponse struct {
	Tempt int64 `json:"temp2m"`
}

type WeatherApiResponse struct {
	Dataseries []DataseriesApiResponse `json:"dataseries"`
}
