package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dherbst/septa"
)

// NextToArrive takes a "from" station name and a "to" station name and returns the expected trains.
func NextToArrive(ctx context.Context) {
	from := flag.Arg(1)
	to := flag.Arg(2)
	fmt.Printf("from=%v to=%v\n", from, to)

	client := septa.NewClient("")
	results, err := client.NextToArrive(from, to, 5)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	for _, r := range results {
		fmt.Printf("%v\n", r)
	}
	fmt.Printf("\nDone\n\n")
}
