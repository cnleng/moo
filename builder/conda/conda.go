package conda

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/moobu/moo/builder"
)

type conda struct {
	options builder.Options
}

func (c *conda) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	var options builder.BuildOptions
	for _, o := range opts {
		o(&options)
	}

	env := s.Name
	ctx := options.Context
	// create a new conda env for isolating the source
	args := []string{"create", "-y", "-n", env}
	args = append(args, options.Deps...)
	cmd := exec.CommandContext(ctx, "conda", args...)
	cmd.Dir = s.Dir
	cmd.Stdout = c.options.Output
	cmd.Stderr = c.options.Output
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// install dependencies if provided requirements.txt
	bundle := &builder.Bundle{Type: "conda", Source: s}
	require := "requirements.txt"
	_, err := os.Stat(filepath.Join(s.Dir, require))
	if errors.Is(err, os.ErrNotExist) {
		return bundle, nil
	}
	if err != nil {
		return nil, err
	}
	// file exists
	cmd = exec.CommandContext(ctx, "conda", "install", "-y", "-n", env, "-r", require)
	cmd.Dir = s.Dir
	cmd.Stdout = c.options.Output
	cmd.Stderr = c.options.Output
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return bundle, nil
}

func (c *conda) Release(b *builder.Bundle, opts ...builder.ReleaseOption) error {
	return nil
}

func (c *conda) Clean(b *builder.Bundle, opts ...builder.CleanOption) error {
	var options builder.CleanOptions
	for _, o := range opts {
		o(&options)
	}

	env := b.Source.Name
	cmd := exec.CommandContext(options.Context, "conda", "env", "remove", "-n", env, "-y")
	cmd.Stdout = c.options.Output
	cmd.Stderr = c.options.Output
	return cmd.Run()
}

func New(opts ...builder.Option) builder.Builder {
	options := builder.Options{Output: os.Stdout}
	for _, o := range opts {
		o(&options)
	}
	return &conda{options: options}
}
