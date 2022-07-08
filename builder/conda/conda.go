package conda

import (
	"os/exec"

	"github.com/mooctl/moo/builder"
)

type conda struct {
}

func (c *conda) Build(s *builder.Source) (*builder.Bundle, error) {
	// create a new conda env
	cmd := exec.Command("conda", "create", "-n", s.Name)
	return nil, nil
}

func (c *conda) Release(b *builder.Bundle) error {
	return nil
}

func (c *conda) Clean(b *builder.Bundle) error {
	return nil
}

func New(opts ...builder.Option) builder.Builder {
	return &conda{}
}
