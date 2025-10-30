package main

import (
	"log"
	"os"

	"github.com/jamrig/chippy/internal/emulator"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must include font path")
	}

	if len(os.Args) < 3 {
		log.Fatal("must include program path")
	}

	e, err := emulator.New(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	e.Start()
}
