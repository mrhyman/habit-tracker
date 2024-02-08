package event_consumer

import (
	"etc/events"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERROR] Failed consumer job: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Printf("[ERROR] Failed handling events: %s", err.Error())

			continue
		}
	}
}

func (c Consumer) handleEvents(events []events.Event) error {
	//var wg sync.WaitGroup

	for _, event := range events {
		log.Printf("[INFO] Processing event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("[ERROR] Failed processing event: %s", err.Error())
			continue
		}
	}

	return nil
}
