package cmd

import (
	"context"

	"github.com/moobu/moo/internal/cli"
)

var cmd = &cli.Cmd{
	Name:    "moo",
	Help:    "A pluggable serverless engine",
	Version: "0.0.1",
}

func Run() error {
	return cmd.RunCtx(context.Background())
}
