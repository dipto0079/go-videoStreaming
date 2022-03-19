package main

import (
	"log"
	"videoStreaming/hrm/storage/postgres"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}
