package main

import (
	"context"
	"flag"
	"fmt"
)

// NextToArrive takes a "from" station name and a "to" station name and returns the expected trains.
func NextToArrive(ctx context.Context) {
	from := flag.Arg(1)
	to := flag.Arg(2)
	fmt.Printf("from=%v to=%v\n", from, to)
}
