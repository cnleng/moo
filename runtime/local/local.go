package local

import (
	"errors"
	"sync"

	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/runtime/local/driver"
)

type local struct {
	sync.RWMutex
	options runtime.Options
	driver  driver.Driver
	pods    map[string]map[string]*localPod
	exit    chan struct{}
}

func (l *local) Create(pod *runtime.Pod, opts ...runtime.CreateOption) error {
	l.Lock()
	defer l.Unlock()

	var options runtime.CreateOptions
	for _, o := range opts {
		o(&options)
	}

	key := pod.String()
	ns := options.Namespace
	if _, ok := l.pods[ns]; !ok {
		l.pods[ns] = make(map[string]*localPod)
	}
	if _, ok := l.pods[ns][key]; ok {
		return errors.New("pod already created")
	}

	lpod := &localPod{
		Pod:    pod,
		driver: l.driver,
		output: options.Output,
		runnable: &driver.Runnable{
			Bundle: options.Bundle,
			Env:    options.Env,
			Args:   options.Args,
		},
	}

	if err := lpod.Start(); err != nil {
		return err
	}
	l.pods[ns][key] = lpod
	return nil
}

func (l *local) List(opts ...runtime.ListOption) ([]*runtime.Pod, error) {

	return nil, nil
}

func (l *local) Delete(pod *runtime.Pod, opts ...runtime.DeleteOption) error {
	return nil
}

func (l *local) Start() error {
	return nil
}

func (l *local) Stop() error {
	return nil
}

func New(driver driver.Driver, opts ...runtime.Option) runtime.Runtime {
	var options runtime.Options
	for _, o := range opts {
		o(&options)
	}
	return &local{
		options: options,
		driver:  driver,
		pods:    make(map[string]map[string]*localPod),
		exit:    make(chan struct{}),
	}
}
