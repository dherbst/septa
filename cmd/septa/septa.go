package main

import (
	"context"
	"flag"
	"fmt"
)

var funcMap map[string]func(context.Context)

func init() {
	funcMap = map[string]func(context.Context){
		"help": Usage,
		"next": NextToArrive,
	}
}

// Usage prints how to invoke `septa` from the command line.
func Usage(ctx context.Context) {
	fmt.Printf(`
Usage:

septa next [from-station] [to-station]
`)

}

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
