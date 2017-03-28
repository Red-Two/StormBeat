package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	weather "github.com/red-two/stormbeat/weather"
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
	// counter := 1
	for {
		now := time.Now()
		bt.query(bt.config.City, b.Name)
		bt.lastIndexTime = now
		logp.Info("Event Sent")
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
	}
}

func (bt *Stormbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *Stormbeat) query(city string, beatname string){
	apiKey := bt.config.ApiKey
  d, err := weather.Query(city, apiKey)
  if err != nil {
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type": beatname,
			"city": d.Name,
			"temperature": d.Main.Kelvin,
			"pressure": d.Main.Pressure,
		  "humidity": d.Main.Humidity,
			"windspeed": d.Wind.Speed,
		}
		fmt.Println(event)
		bt.client.PublishEvent(event)
	}
}
