package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	t := mustToken()

	//token = flags.Get(token)

	tgClient = telegram.New(tgBotHost, t)

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
