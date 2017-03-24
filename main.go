package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/red-two/stormbeat/beater"
)

func main() {
	err := beat.Run("stormbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
