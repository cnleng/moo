package raw

import (
	"errors"
	"os"
	"os/exec"
	"syscall"

	"github.com/moobu/moo/runtime/local/driver"
)

type raw struct{}

func (raw) Fork(r *driver.Runnable) (*driver.Process, error) {
	var dir string
	if r.Bundle.Source != nil {
		dir = r.Bundle.Dir
	}
	name := r.Bundle.Entry[0]
	args := append(r.Bundle.Entry[1:], r.Args...)
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	pout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	perr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	pin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return &driver.Process{
		ID:  cmd.Process.Pid,
		Out: pout,
		Err: perr,
		In:  pin,
	}, nil
}

func (raw) Kill(p *driver.Process) error {
	return syscall.Kill(-p.ID, syscall.SIGTERM)
}

func (raw) Wait(p *driver.Process) error {
	proc, err := os.FindProcess(p.ID)
	if err != nil {
		return err
	}
	ps, err := proc.Wait()
	if err != nil {
		return err
	}
	if ps.Success() {
		return nil // the process exited with status 0
	}
	return errors.New(ps.String())
}

func New() driver.Driver {
	return raw{}
}
