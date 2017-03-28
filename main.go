package main

import (
  //"os"
	"fmt"
	//"github.com/elastic/beats/libbeat/beat"
	//"github.com/red-two/stormbeat/beater"
	weather "github.com/red-two/stormbeat/weather"
)

func main() {
	fmt.Println(weather.Query("Rochester", ""))
	// err := beat.Run("stormbeat", "", beater.New)
	// if err != nil {
	// 	os.Exit(1)
	// }
}
