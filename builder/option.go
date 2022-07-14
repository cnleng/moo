package builder

type Options struct{}

type Option func(*Options)

type BuildOptions struct {
	Dir string
	Ref string
}

type BuildOption func(*BuildOptions)

func Dir(dir string) BuildOption {
	return func(o *BuildOptions) {
		o.Dir = dir
	}
}

func Ref(ref string) BuildOption {
	return func(o *BuildOptions) {
		o.Ref = ref
	}
}

type CleanOptions struct{}

type CleanOption func(*CleanOptions)
