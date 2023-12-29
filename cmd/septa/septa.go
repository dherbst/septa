package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/dherbst/septa"
)

var funcMap map[string]func(context.Context)

func init() {
	funcMap = map[string]func(context.Context){
		"help":     Usage,
		"next":     NextToArrive,
		"stations": Stations,
		"version":  Version,
		"alerts":   Alerts,
	}
}

// Version prints the version from the septa.GitHash out and exits.
func Version(ctx context.Context) {
	fmt.Printf("Version: %v\n", septa.Version)
}

// Usage prints how to invoke `septa` from the command line.
func Usage(ctx context.Context) {
	fmt.Printf(`
Usage:

septa alerts [route_id]                    ; show alerts
septa stations [pattern]                   ; list stations, or match pattern
septa next [from-station] [to-station]     ; show the next to arrive
septa version                              ; prints the commit version
`)

}

//Usage: alerts|next|stations|version
func main() {
	flag.Parse()

	ctx := context.Background()

	command := flag.Arg(0)
	if command == "" || command == "help" {
		Usage(ctx)
		return
	}

	f := funcMap[command]
	if f == nil {
		fmt.Println("Unknown command")
		Usage(ctx)
		return
	}

	f(ctx)
}
