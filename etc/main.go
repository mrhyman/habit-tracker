package main

import (
	tgClient "etc/clients/telegram"
	event_consumer "etc/consumer/event-consumer"
	"etc/events/telegram"
	"etc/lib/storage/files"
	"flag"
	"log"
)

const (
	host        = "api.telegram.org"
	storagePath = "files-storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(host, mustToken()),
		files.New(storagePath),
	)

	log.Println("Starting telegram bot")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("ALARM! ALARM! ABORTING!", err)
	}

}

func mustToken() string {
	token := flag.String("bot-token", "", "TG API Token")

	flag.Parse()

	if *token == "" {
		log.Fatal("There is no TG token. Fatal Error!")
	}

	return *token
}
