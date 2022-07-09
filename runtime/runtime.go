package runtime

import "fmt"

type Runtime interface {
	Create(*Pod, ...CreateOption) error
	// Delete(*Pod) error
	// List() ([]*Pod, error)
	// Start() error
	// Stop() error
}

type Pod struct {
	Name     string
	Tag      string
	Status   Status
	Metadata map[string]string
}

func (p Pod) String() string {
	return fmt.Sprintf("%s:%s", p.Name, p.Tag)
}

type Status int8
