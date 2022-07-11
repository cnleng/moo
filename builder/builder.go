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
	Source *Source
}

var Default Builder

func Build(s *Source) (*Bundle, error) {
	return Default.Build(s)
}

func Release(b *Bundle) error {
	return Default.Release(b)
}

func Clean(b *Bundle) error {
	return Default.Clean(b)
}
