package test

import (
	"os"

	"github.com/moobu/moo/builder"
	bnoop "github.com/moobu/moo/builder/noop"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
	"github.com/moobu/moo/router/memory"
	"github.com/moobu/moo/runtime"
	rnoop "github.com/moobu/moo/runtime/noop"
	"github.com/moobu/moo/server"
	"github.com/moobu/moo/server/http"
)

type Presets struct{}

func (Presets) Setup(c cli.Ctx) error {
	builder.DefaultBuilder = bnoop.New(builder.Output(os.Stdout))
	runtime.DefaultRuntime = rnoop.New()
	router.DefaultRouter = memory.New()
	server.DefaultServer = http.New()
	return nil
}

func (Presets) String() string {
	return "test"
}
