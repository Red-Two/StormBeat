package main

import (
  "encoding/json"
  //"flag"
  "log"
  "net/http"
  "strings"
  "time"
)

type weatherData struct {
  Name string `json:"name"`
  Main struct {
    Kelvin float64 `json:"temp"`
  } `json:"main"`
}

func query(){

}

type openWeatherMap struct{}

func (w openWeatherMap) temperature(city string) (float64, error) {
  apiKey := ""
  resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiKey + "&q=" + city)
  if err != nil {
    return 0, err
  }
  defer resp.Body.Close()
  var d struct {
    Main struct {
      Kelvin float64 `json:"temp"`
    } `json:"main"`
  }
  if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
    return 0, err
  }
  log.Printf("openWeatherMap: %s: %.2f", city, d.Main.Kelvin)
  return d.Main.Kelvin, nil
}