package local

import (
	"os"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/builder/conda"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
	"github.com/moobu/moo/router/memory"
	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/local"
	"github.com/moobu/moo/runtime/local/driver/raw"
	"github.com/moobu/moo/server"
	"github.com/moobu/moo/server/http"
)

type Presets struct{}

func (Presets) Setup(c cli.Ctx) error {
	runtime.Default = local.New(raw.New()) // TODO: use conda driver
	builder.Default = conda.New(builder.Output(os.Stdout))
	router.Default = memory.New()
	server.Default = http.New()
	return nil
}

func (Presets) String() string {
	return "local"
}
