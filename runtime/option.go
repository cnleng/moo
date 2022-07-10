package runtime

import (
	"io"

	"github.com/moobu/moo/builder"
)

type Options struct{}

type Option func(*Options)

type CreateOptions struct {
	Args      []string
	Env       []string
	Bundle    *builder.Bundle
	Output    io.Writer
	Image     string
	Namespace string
	Replicas  int
}

type CreateOption func(*CreateOptions)

func Args(args ...string) CreateOption {
	return func(o *CreateOptions) {
		o.Args = args
	}
}

func Env(env ...string) CreateOption {
	return func(o *CreateOptions) {
		o.Env = env
	}
}

func Bundle(bundle *builder.Bundle) CreateOption {
	return func(o *CreateOptions) {
		o.Bundle = bundle
	}
}

func Output(w io.Writer) CreateOption {
	return func(o *CreateOptions) {
		o.Output = w
	}
}

func Image(name string) CreateOption {
	return func(o *CreateOptions) {
		o.Image = name
	}
}

func CreateWithNamespace(ns string) CreateOption {
	return func(o *CreateOptions) {
		o.Namespace = ns
	}
}

func Replicas(n int) CreateOption {
	return func(o *CreateOptions) {
		o.Replicas = n
	}
}

type ListOptions struct {
	Name      string
	Tag       string
	Namespace string
}

type ListOption func(*ListOptions)

func ListWithNamespace(ns string) ListOption {
	return func(o *ListOptions) {
		o.Namespace = ns
	}
}

func Name(name string) ListOption {
	return func(o *ListOptions) {
		o.Name = name
	}
}

func Tag(tag string) ListOption {
	return func(o *ListOptions) {
		o.Tag = tag
	}
}

type DeleteOptions struct {
	Namespace string
}

type DeleteOption func(*DeleteOptions)

func DeleteWithNamespace(ns string) DeleteOption {
	return func(o *DeleteOptions) {
		o.Namespace = ns
	}
}
