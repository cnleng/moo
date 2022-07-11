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
)

type Presets struct{}

func (Presets) Setup(c cli.Ctx) {
	builder.DefaultBuilder = bnoop.New(builder.Output(os.Stdout))
	runtime.DefaultRuntime = rnoop.New()
	router.DefaultRouter = memory.New()
}

func (Presets) String() string {
	return "test"
}
