package main

import (
	"flag"
	"log"
)

var version string

func main() {
	var info = flag.Bool("v", false, "Print version")
	flag.Parse()

	if *info {
		log.Printf("Current version: %s\n", version)
		return
	}
}
