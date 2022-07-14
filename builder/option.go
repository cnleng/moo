package builder

import (
	"context"
)

type Options struct {
	Retriever Retriever
}

type Option func(*Options)

type BuildOptions struct {
	Context context.Context
	Ref     string
}

type BuildOption func(*BuildOptions)

func BuildContext(c context.Context) BuildOption {
	return func(o *BuildOptions) {
		o.Context = c
	}
}

func Ref(ref string) BuildOption {
	return func(o *BuildOptions) {
		o.Ref = ref
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
