package runtime

import "fmt"

type Runtime interface {
	// Create creates a pod containing a process
	Create(*Pod, ...CreateOption) error
	// Delete deletes a pod
	Delete(*Pod, ...DeleteOption) error
	// List lists pods
	List(...ListOption) ([]*Pod, error)
	// Start/Stop starts/stops the runtime daemon,
	// and is used only by the local runtime
	Start() error
	Stop() error
}

type Pod struct {
	Name     string
	Tag      string
	Metadata map[string]string
}

func (p Pod) String() string {
	return fmt.Sprintf("%s:%s", p.Name, p.Tag)
}

func (p *Pod) Status(status Status, err error) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}
	if err != nil {
		p.Metadata["error"] = err.Error()
	}
	p.Metadata["status"] = status.String()
}

type Status int8

func (s Status) String() string {
	return StatusText[s]
}

const (
	Pending Status = iota
	Running
	Exited
)

var StatusText = [...]string{
	"pendding",
	"running",
	"exited",
}

var Default Runtime

func Create(pod *Pod, opts ...CreateOption) error {
	return Default.Create(pod, opts...)
}

func Delete(pod *Pod, opts ...DeleteOption) error {
	return Default.Delete(pod, opts...)
}

func List(opts ...ListOption) ([]*Pod, error) {
	return Default.List(opts...)
}

func Start() error {
	return Default.Start()
}

func Stop() error {
	return Default.Stop()
}
