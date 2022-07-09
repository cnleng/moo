package conda

import "github.com/moobu/moo/runtime"

type conda struct {
	store Store
}

func (c *conda) Create(pod *runtime.Pod, opts ...runtime.CreateOption) error {
	var options runtime.CreateOptions
	for _, o := range opts {
		o(&options)
	}

	return nil
}

func New(opts ...runtime.Option) runtime.Runtime {
	return &conda{}
}
