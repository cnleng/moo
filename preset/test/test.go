package test

import (
	"os"

	"github.com/moobu/moo/builder"
	bnoop "github.com/moobu/moo/builder/noop"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
	"github.com/moobu/moo/router/static"
	"github.com/moobu/moo/runtime"
	rnoop "github.com/moobu/moo/runtime/noop"
	"github.com/moobu/moo/server"
	"github.com/moobu/moo/server/http"
)

type Presets struct{}

func (Presets) Setup(c cli.Ctx) error {
	builder.Default = bnoop.New(builder.Output(os.Stdout))
	runtime.Default = rnoop.New()
	router.Default = static.New()
	server.Default = http.New()
	return nil
}

func (Presets) String() string {
	return "test"
}
