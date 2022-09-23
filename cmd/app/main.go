package main

import (
	"log"
	"tg_bot/pkg/postgres"
)

func main() {
	err := postgres.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
