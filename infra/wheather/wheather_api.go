package wheather

import (
	"encoding/json"
	"fmt"
	"github-sub/daiskoba/di-practice/usecase"
	"io"
	"net/http"
)

type wheatherClient struct {
	place    string
	point    map[string]string
	forecast usecase.Wheather
	endpoint string
}

type weatherData struct {
	PublicTime  string      `json:"publicTime"`
	Description description `json:"description"`
	Forecasts   []forecast  `json:"forecasts"`
}

type description struct {
	PublicTime          string `json:"publicTime"`
	PublicTimeFormatted string `json:"publicTimeFormatted"`
	BodyText            string `json:"bodyText"`
}

type forecast struct {
	Date      string `json:"date"`
	DateLabel string `json:"dateLabel"`
	Telop     string `json:"telop"`
}

func NewWheatherClient() usecase.WheatherClient {
	wc := &wheatherClient{}
	wc.init()
	return wc
}

func (wc *wheatherClient) init() {
	wc.point = make(map[string]string)
	wc.point["tokyo"] = "130010"
	wc.point["osaka"] = "270000"
	wc.endpoint = "https://weather.tsukumijima.net/api/forecast/city/"
}

func (wc *wheatherClient) Forecast(place string) (*usecase.Wheather, error) {
	var err error
	uw := &usecase.Wheather{}

	if !wc.canForecast(place) {
		err = fmt.Errorf("place %s does not forecast", place)
	}

	uw, err = wc.getForecast(place)

	if err != nil {
		err = fmt.Errorf("get forecast faild %s: %s", place, err)
	}

	return uw, err
}

func (wc *wheatherClient) canForecast(place string) bool {
	_, ok := wc.point[place]
	return ok
}

func (wc *wheatherClient) getForecast(place string) (*usecase.Wheather, error) {
	uw := &usecase.Wheather{}
	var err error

	if !wc.canForecast(place) {
		err = fmt.Errorf("place %s does not forecast", place)

		return uw, err
	}

	// Query to Forecase API
	url := wc.endpoint + wc.point[place]
	resp, err := http.Get(url)

	if err != nil {
		err = fmt.Errorf("get forecast failed %s", url)

		return uw, err
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	err = parse(uw, jsonData)

	if err != nil {
		err = fmt.Errorf("parse forecast failed %s", err)
	}

	return uw, err
}

func parse(uw *usecase.Wheather, jsonData []byte) error {
	var err error
	var wd weatherData

	err = json.Unmarshal(jsonData, &wd)

	uw.PublicTime = wd.PublicTime
	uw.Summary = wd.Description.BodyText
	uw.Forecasts = make(map[string]string)

	for _, v := range wd.Forecasts {
		uw.Forecasts[v.Date] = v.Telop
	}

	// fmt.Printf("%v\n", wd.Forecasts)

	if err != nil {
		err = fmt.Errorf("parse forecast failed %s", err)
	}

	return err
}
