package noop

import "github.com/moobu/moo/runtime"

type noop struct{}

func (noop) Create(pod *runtime.Pod, opts ...runtime.CreateOption) error {
	return nil
}

func (noop) Delete(pod *runtime.Pod, opts ...runtime.DeleteOption) error {
	return nil
}

func (noop) List(opts ...runtime.ListOption) ([]*runtime.Pod, error) {
	return nil, nil
}

func (noop) Start() error {
	return nil
}

func (noop) Stop() error {
	return nil
}

func New(opts ...runtime.Option) runtime.Runtime {
	return noop{}
}
