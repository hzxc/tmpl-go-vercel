package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Model struct {
	Bns     []Bn   `json:"bns"`
	Ip      string `json:"ip"`
	Weather string `json:"weather"`
}

type Bn struct {
	Symbol string `json:"symbol"`

	Price string `json:"price"`
}

func Awesome(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		resp *http.Response
		body []byte
	)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	model := &Model{Bns: []Bn{}}
	model.Ip = r.RemoteAddr
	binanceApi := "https://api.binance.com/api/v3/ticker/price?symbols=[%22BTCUSDT%22,%22ETHUSDT%22]"
	weatherApi := fmt.Sprintf("https://wttr.in/%s?format=4", model.Ip)

	if resp, err = http.Get(binanceApi); err != nil {
		goto ERR
	}
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		goto ERR
	}

	json.Unmarshal(body, &model.Bns)

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
