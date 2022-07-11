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
	runtime.Default = kubernetes.New()
	builder.Default = podman.New(builder.Output(os.Stdout))
	router.Default = nil // TODO: should we have a router built on kubernetes?
	server.Default = http.New()
	return nil
}

func (Presets) String() string {
	return "kubernetes"
}
