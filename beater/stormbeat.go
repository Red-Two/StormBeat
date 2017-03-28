package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	//"github.com/red-two/stormbeat/weather"
	"github.com/red-two/stormbeat/config"
)

type Stormbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
	lastIndexTime time.Time
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	sb := &Stormbeat{
		done: make(chan struct{}),
		config: config,
	}
	return sb, nil
}

func (bt *Stormbeat) Run(b *beat.Beat) error {
	logp.Info("stormbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		// now := time.Now()
		// bt.query()
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Stormbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

// func (bt *Stormbeat) query(city string, beatname string) (weatherData, error) {
//   config := config.DefaultConfig
// 	apiKey := config.apiKey
//   resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiKey + "&q=" + city)
//   if err != nil {
//     return weatherData{}, err
//   }
//   defer resp.Body.Close()
//   var d weatherData
//   if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
//     return weatherData{}, err
//   }
//   return d, nil
// }
