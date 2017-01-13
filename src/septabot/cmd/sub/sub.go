package main

import (
	"context"
	"log"
	"septabot"
)

func main() {
	log.Printf("Getting Suburban Station to Narberth\n")

	septa := septabot.NewSepta(septabot.NewSeptaAPIImpl())

	ctx := context.Background()
	results, err := septa.NextToArrive(ctx, "Suburban%20Station", "Narberth", 5)
	if err != nil {
		log.Printf("Got an error %v\n", err)
	}

	for _, r := range results {
		log.Printf("%v\n", r)
	}
}
