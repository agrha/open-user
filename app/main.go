package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("No command specified")
		return
	}

	if args[0] == "serve" {
		if err := serveApi(); err != nil {
			log.Fatalf("Error serving API: %v", err)
			return
		}
		return
	}
}
