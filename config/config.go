// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	apiKey string `config:"apikey"`
	city string `config:"city"`
}

var DefaultConfig = Config{
	Period: 20 * time.Second,
	apiKey: "",
	city: "NewYork",
}
