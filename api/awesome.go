package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Model struct {
	Symbol  string `json:"symbol"`
	Price   string `json:"price"`
	Weather string
}

func Awesome(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		resp *http.Response
		body []byte
	)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	model := &Model{}
	binanceApi := "https://api.binance.com/api/v3/ticker/price?symbols=[%22BTCUSDT%22,%22ETHUSDT%22]"
	weatherApi := "https://wttr.in/?format=1"

	if resp, err = http.Get(binanceApi); err != nil {
		goto ERR
	}
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		goto ERR
	}

	json.Unmarshal(body, &model)

	if resp, err = http.Get(weatherApi); err != nil {
		goto ERR
	}
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		goto ERR
	}

	model.Weather = string(body)

	if body, err = json.Marshal(model); err != nil {
		goto ERR
	}

	w.Write(body)
	return
ERR:
	w.Write([]byte(err.Error()))
}
