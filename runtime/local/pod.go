package local

import (
	"io"
	"sync"

	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/local/driver"
)

const maxRetries = 3

type localPod struct {
	sync.RWMutex
	*runtime.Pod
	retries  int
	output   io.Writer
	driver   driver.Driver
	runnable *driver.Runnable
	process  *driver.Process
}

func (p *localPod) Start() error {
	p.RLock()
	defer p.RUnlock()

	if p.retries > maxRetries {
		return nil
	}
	return p.start()
}

func (p *localPod) start() (err error) {
	p.process, err = p.driver.Fork(p.runnable)
	if err != nil {
		return
	}
	p.Status(runtime.Running, nil)
	if p.output != nil {
		p.stream()
	}
	go p.wait()
	return nil
}

func (p *localPod) stream() {
	go io.Copy(p.output, p.process.Stdout)
	go io.Copy(p.output, p.process.Stderr)
}

func (p *localPod) wait() {
	err := p.driver.Wait(p.process)
	p.Lock()
	p.Status(runtime.Exited, err)
	p.retries++
	p.Unlock()
}

func (p *localPod) Stop() error {
	return p.driver.Kill(p.process)
}
