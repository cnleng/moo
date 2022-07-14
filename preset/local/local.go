package local

import (
	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/builder/golang"
	"github.com/moobu/moo/builder/mixed"
	"github.com/moobu/moo/builder/python"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
	"github.com/moobu/moo/router/static"
	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/local"
	"github.com/moobu/moo/runtime/local/driver/raw"
	"github.com/moobu/moo/server"
	"github.com/moobu/moo/server/http"
)

type Presets struct{}

func (Presets) Setup(c cli.Ctx) error {
	runtime.Default = local.New(raw.New())
	router.Default = static.New()
	server.Default = http.New()
	builder.Default = mixed.New([]builder.Builder{
		python.New(),
		golang.New(),
	})
	return nil
}

func (Presets) String() string {
	return "local"
}
