package builder

type Builder interface {
	Build(*Source) (*Bundle, error)
	Release(*Bundle) error
	Clean(*Bundle) error
}

type Source struct {
	Name string
	Path string
}

type Bundle struct {
}

var DefaultBuilder Builder

func Build(s *Source) (*Bundle, error) {
	return DefaultBuilder.Build(s)
}

func Release(b *Bundle) error {
	return DefaultBuilder.Release(b)
}

func Clean(b *Bundle) error {
	return DefaultBuilder.Clean(b)
}
