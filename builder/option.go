package builder

import (
	"context"
	"io"
)

type Options struct {
	Token  string
	Output io.Writer
}

type Option func(*Options)

func Token(token string) Option {
	return func(o *Options) {
		o.Token = token
	}
}

func Output(w io.Writer) Option {
	return func(o *Options) {
		o.Output = w
	}
}

type BuildOptions struct {
	Context context.Context
	Deps    []string
}

type BuildOption func(*BuildOptions)

func Deps(deps ...string) BuildOption {
	return func(o *BuildOptions) {
		o.Deps = deps
	}
}

func BuildContext(c context.Context) BuildOption {
	return func(o *BuildOptions) {
		o.Context = c
	}
}

type ReleaseOptions struct {
	Context context.Context
}

type ReleaseOption func(*ReleaseOptions)

func ReleaseContext(c context.Context) ReleaseOption {
	return func(o *ReleaseOptions) {
		o.Context = c
	}
}

type CleanOptions struct {
	Context context.Context
}

type CleanOption func(*CleanOptions)

func CleanContext(c context.Context) CleanOption {
	return func(o *CleanOptions) {
		o.Context = c
	}
}
