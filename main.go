package main

import (
	"context"
	"log"

	"github.com/moobu/moo/cmd"
)

func main() {
	ctx := context.Background()
	if err := cmd.RunCtx(ctx); err != nil {
		log.Fatal(err)
	}
}
