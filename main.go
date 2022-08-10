package main

import (
	"flag"
	"log"

	tgClient "github.com/ssereduk/telegram-bot/clients/telegram"

	event_consumer "github.com/ssereduk/telegram-bot/consumer/event-consumer"
	"github.com/ssereduk/telegram-bot/events/telegram"
	"github.com/ssereduk/telegram-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	//token = flags.Get(token)

	eventProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Printf("service started")

	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

	//tgClient = telegram.New(token)

	//fetcher = fetcher.New(tgClient)

	//processor = processor.New(tgClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token to acces to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token

}
