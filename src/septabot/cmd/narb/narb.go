package main

import (
	"context"
	"log"
	"septabot"
)

func main() {
	log.Printf("getting Narberth to Suburban Station\n")

	septa := septabot.NewSepta(septabot.NewSeptaAPIImpl())

	ctx := context.Background()
	results, err := septa.NextToArrive(ctx, "Narberth", "Suburban%20Station", 5)
	if err != nil {
		log.Printf("Got an error %v\n", err)
	}

	for _, r := range results {
		log.Printf("%v\n", r)
	}

}
