package conda

import (
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/runtime"
)

type pod struct {
	sync.RWMutex
	*runtime.Pod
	args   []string
	env    []string
	output io.Writer
	bundle *builder.Bundle
	proc   *os.Process
}

func (p *pod) start() error {
	bin := p.bundle.Binary
	if len(bin) > 0 {
		bin = "python"
	}
	args := []string{"run", "-n", p.Name, bin}
	args = append(args, p.args...)
	cmd := exec.Command("conda", args...)
	cmd.Dir = p.bundle.Source.Dir
	cmd.Env = p.env
	cmd.Stdout = p.output
	cmd.Stderr = p.output
	if err := cmd.Start(); err != nil {
		return err
	}
	p.proc = cmd.Process
	return nil
}
