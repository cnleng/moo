package local

import (
	"io"
	"sync"

	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/local/driver"
)

const maxRetries = 3

type lpod struct {
	sync.RWMutex
	*runtime.Pod
	output   io.Writer
	driver   driver.Driver
	runnable *driver.Runnable
	process  *driver.Process
	wg       *sync.WaitGroup
	retries  int  // retries the pod has consumed
	running  bool // avoid exceedingly restarting the pod
}

func (p *lpod) restartIfDead() error {
	p.RLock()
	if !p.retry() {
		p.RUnlock()
		return nil
	}
	p.RUnlock()
	return p.start()
}

func (p *lpod) retry() bool {
	if p.running {
		return false
	}
	return p.retries <= maxRetries
}

func (p *lpod) start() (err error) {
	p.Lock()
	defer p.Unlock()

	if !p.retry() {
		return
	}

	p.process, err = p.driver.Fork(p.runnable)
	if err != nil {
		return
	}
	p.Status(runtime.Running, nil)
	if p.output != nil {
		p.stream()
	}
	p.running = true
	p.wg.Add(1)
	go p.wait()
	return nil
}

func (p *lpod) stream() {
	go io.Copy(p.output, p.process.Out)
	go io.Copy(p.output, p.process.Err)
}

func (p *lpod) wait() {
	err := p.driver.Wait(p.process)
	p.Lock()
	defer p.Unlock()
	p.Status(runtime.Exited, err)
	p.retries++
	p.running = false
	p.wg.Done()
}

func (p *lpod) stop() error {
	return p.driver.Kill(p.process)
}
