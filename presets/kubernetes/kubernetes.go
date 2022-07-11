package kubernetes

import (
	"os"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/builder/podman"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/kubernetes"
)

type Presets struct{}

func (Presets) Setup(c cli.Ctx) {
	runtime.DefaultRuntime = kubernetes.New()
	builder.DefaultBuilder = podman.New(builder.Output(os.Stdout))
	router.DefaultRouter = nil // TODO: should we have a router built on kubernetes?
}

func (Presets) String() string {
	return "kubernetes"
}
