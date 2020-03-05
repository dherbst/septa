package main

import (
	"context"
	"fmt"

	"github.com/dherbst/septa"
)

// Stations command lists the stations that acceptable for Next To Arrive.
func Stations(ctx context.Context) {
	for _, v := range septa.Stations {
		fmt.Printf("%v\n", v)
	}
}
