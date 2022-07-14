package local

import (
	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/builder/auto"
	"github.com/moobu/moo/builder/golang"
	"github.com/moobu/moo/builder/python"
	"github.com/moobu/moo/builder/retriever"
	"github.com/moobu/moo/builder/retriever/git"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
	"github.com/moobu/moo/router/static"
	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/local"
	"github.com/moobu/moo/runtime/local/driver/raw"
	"github.com/moobu/moo/server"
	"github.com/moobu/moo/server/http"
)

type Preset struct{}

func (Preset) Setup(c cli.Ctx) error {
	runtime.Default = local.New(raw.New())
	router.Default = static.New()
	server.Default = http.New()
	builder.Default = retriever.New(git.New(),
		auto.New([]builder.Builder{
			python.New(),
			golang.New(),
		}))
	return nil
}

func (Preset) String() string {
	return "local"
}
