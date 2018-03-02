package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"septabot"
)

func main() {

	if len(os.Args) < 2 {
		log.Printf("Missing command.\nUsage: septa nta <src> <dest>\n")
		os.Exit(1)
	}

	if os.Args[1] != "nta" {
		log.Printf("Unknown command %v\n", os.Args[1])
		os.Exit(2)
	}

	// Need src dest
	if len(os.Args) < 4 {
		log.Printf("Missing command.\nUsage: septa nta <src> <dest>\n")
		os.Exit(3)
	}

	src := os.Args[2]
	dest := os.Args[3]

	septa := septabot.NewSepta(septabot.NewSeptaAPIImpl())
	ctx := context.Background()
	results, err := septa.NextToArrive(ctx, src, dest, 5)
	if err != nil {
		log.Printf("Got error %v\n", err)
		return
	}

	for _, r := range results {
		fmt.Printf("%v\n", r)
	}
}
