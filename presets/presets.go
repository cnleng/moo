package presets

import (
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/presets/kubernetes"
	"github.com/moobu/moo/presets/local"
	"github.com/moobu/moo/presets/test"
)

type Presets interface {
	Setup(cli.Ctx)
	String() string
}

var presets = map[string]Presets{
	"test":       test.Presets{},
	"local":      local.Presets{},
	"kubernetes": kubernetes.Presets{},
}

func Register(p Presets) {
	presets[p.String()] = p
}

func Deregister(p Presets) {
	delete(presets, p.String())
}
