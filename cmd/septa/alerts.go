package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/dherbst/septa"
)

// Alerts optionally takes a route.
func Alerts(ctx context.Context) {
	route := flag.Arg(1)

	fmt.Printf("route=%v\n", route)

	client := septa.NewClient("")
	results, err := client.Alerts(route)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	for _, r := range results {
		fmt.Printf("%v\n", r)
	}
}
