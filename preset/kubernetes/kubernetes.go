package kubernetes

import (
	"os"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/builder/podman"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/kubernetes"
	"github.com/moobu/moo/server"
	"github.com/moobu/moo/server/http"
)

type Presets struct{}

func (Presets) Setup(c cli.Ctx) error {
	runtime.DefaultRuntime = kubernetes.New()
	builder.DefaultBuilder = podman.New(builder.Output(os.Stdout))
	router.DefaultRouter = nil // TODO: should we have a router built on kubernetes?
	server.DefaultServer = http.New()
	return nil
}

func (Presets) String() string {
	return "kubernetes"
}
