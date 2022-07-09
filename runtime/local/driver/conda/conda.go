package conda

import (
	"github.com/moobu/moo/runtime/local/driver"
	"github.com/moobu/moo/runtime/local/driver/raw"
)

type conda struct {
	driver driver.Driver
}

func (c conda) Fork(r *driver.Runnable) (*driver.Process, error) {
	bin := r.Bundle.Binary
	args := []string{"run", "-n", r.Bundle.Source.Name, bin}
	args = append(args, r.Args...)
	r.Bundle.Binary = "conda"
	r.Args = args
	return c.driver.Fork(r)
}

func (c conda) Kill(p *driver.Process) error {
	return c.driver.Kill(p)
}

func (c conda) Wait(p *driver.Process) error {
	return c.driver.Wait(p)
}

func New() driver.Driver {
	return &conda{driver: raw.New()}
}
