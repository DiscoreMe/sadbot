package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

type Weather struct {
	token string
}

type Data struct {
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed int `json:"speed"`
		Deg   int `json:"deg"`
	} `json:"wind"`
}

func NewWeather(token string) *Weather {
	return &Weather{token: token}
}

func (w *Weather) WeatherByLocation(location string) (*Data, error) {
	resp, err := http.Get(fmt.Sprintf(url, location, w.token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body bytes.Buffer
	if _, err := io.Copy(&body, resp.Body); err != nil {
		return nil, err
	}

	var data Data
	if err := json.Unmarshal(body.Bytes(), &data); err != nil {
		return nil, err
	}

	return &data, err
}
