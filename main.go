package main

import (
	"erik-craigo-code-challenge/config"
	"erik-craigo-code-challenge/server"
	"flag"
	"fmt"
	"os"
)

// function to subscribe to events

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
