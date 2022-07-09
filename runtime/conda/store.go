package conda

import "github.com/moobu/moo/runtime"

type Store interface {
	Write(*runtime.Pod) error
	Read(name, tag string) ([]*runtime.Pod, error)
	Close() error
}
