package builder

type Builder interface {
	Build(*Source, ...BuildOption) (*Bundle, error)
	Release(*Bundle, ...ReleaseOption) error
	Clean(*Bundle, ...CleanOption) error
}

type Source struct {
	Name string
	Dir  string
}

type Bundle struct {
	Type   string
	Binary string
	Sum    string
	Source *Source
}

var Default Builder

func Build(s *Source, opts ...BuildOption) (*Bundle, error) {
	return Default.Build(s, opts...)
}

func Release(b *Bundle, opts ...ReleaseOption) error {
	return Default.Release(b, opts...)
}

func Clean(b *Bundle, opts ...CleanOption) error {
	return Default.Clean(b, opts...)
}
