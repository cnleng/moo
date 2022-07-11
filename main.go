package main

import (
	"log"

	"github.com/moobu/moo/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
