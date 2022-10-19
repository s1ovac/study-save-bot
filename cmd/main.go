package main

import (
	"log"

	flags "github.com/jessevdk/go-flags"
	"github.com/s1ovac/study-save-to-bot/client/telegram"
)

var opts struct {
	Token string `short:"t" long:"token" default:"" description:"Token is key to connect with tgbot"`
}

const HOST = "api.telegram.org"

func main() {
	t := telegram.New(HOST, mustToken())

	//tgclient := telegram.New(token)

	//fetcher := fetcher.New()

	//processor := processor.New()

	//consumer.Start(fetcher, processor)
}

// mustToken exit the app if tg Token is unavailable
func mustToken() string {
	flags.Parse(&opts)
	if opts.Token == "" {
		log.Fatal("your token is unavailable")
	}
	return opts.Token
}
