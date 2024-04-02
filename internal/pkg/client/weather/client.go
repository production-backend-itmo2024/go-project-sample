package weather

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetCurrentWeather() (*Weather, error) {
	resp, err := http.Get("http://www.7timer.info/bin/api.pl?lon=59.93&lat=30.30&product=astro&output=json")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respModel WeatherApiResponse
	err = json.Unmarshal(body, &respModel)
	if err != nil {
		return nil, err
	}

	if len(respModel.Dataseries) > 0 {
		return &Weather{Temp: respModel.Dataseries[0].Tempt}, err
	}

	return nil, err
}
