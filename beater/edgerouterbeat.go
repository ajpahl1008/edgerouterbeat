package beater

import (
	"fmt"
	"github.com/elastic/beats/libbeat/logp"
	"log"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"

	"github.com/ajpahl1008/edgerouterbeat/config"
)

// Edgerouterbeat configuration.
type Edgerouterbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of edgerouterbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}
	bt := &Edgerouterbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts edgerouterbeat.
func (bt *Edgerouterbeat) Run(b *beat.Beat) error {

	log.Printf("edgerouterbeat is running! Hit CTRL-C to stop it.")
	var err error

	log.Println("Refresh Rate: ", bt.config.Period)

	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				CollectEdgeStats().Interfaces[0].InterfaceName : CollectEdgeStats().Interfaces[0],
				CollectEdgeStats().Interfaces[1].InterfaceName : CollectEdgeStats().Interfaces[1],
				CollectEdgeStats().Interfaces[2].InterfaceName : CollectEdgeStats().Interfaces[2],
				CollectEdgeStats().Interfaces[3].InterfaceName : CollectEdgeStats().Interfaces[3],
				CollectEdgeStats().Interfaces[4].InterfaceName : CollectEdgeStats().Interfaces[4],
				CollectEdgeStats().Interfaces[5].InterfaceName : CollectEdgeStats().Interfaces[5],
				CollectEdgeStats().Interfaces[6].InterfaceName : CollectEdgeStats().Interfaces[6],
				CollectEdgeStats().Interfaces[7].InterfaceName : CollectEdgeStats().Interfaces[7],
				CollectEdgeStats().Interfaces[8].InterfaceName : CollectEdgeStats().Interfaces[8],
				"type":    b.Info.Name,
				"counter": counter,
			},
		}
		bt.client.Publish(event)
		logp.NewLogger("Event Sent")
		counter++
	}
}

// Stop stops edgerouterbeat.
func (bt *Edgerouterbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
