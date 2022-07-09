package conda

import "github.com/mooctl/moo/runtime"

type Store interface {
	Write(*runtime.Pod) error
	Read(name, tag string) ([]*runtime.Pod, error)
	Close() error
}
