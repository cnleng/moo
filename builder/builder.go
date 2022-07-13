package builder

type Builder interface {
	// Build turns the source into a bundle which's then
	// delivered to the runtime that runs the bundle in
	// one or several pods. If Local is unset, we should
	// ask the retriever for the source using Remote.
	Build(*Source, ...BuildOption) (*Bundle, error)
	// Clean cleans up a bundle and the dependencies created
	// during its build.
	Clean(*Bundle, ...CleanOption) error
}

type Retriever interface {
	// Retrieve retrieves a source normally a git repository
	// from remote platforms like GitHub.
	Retrieve(string) (*Source, error)
}

type Source struct {
	Name   string
	Remote string
	Local  string
}

type Bundle struct {
	Type   string
	Binary string
	Ref    string
	Source *Source
}

var Default Builder

func Build(s *Source, opts ...BuildOption) (*Bundle, error) {
	return Default.Build(s, opts...)
}

func Clean(b *Bundle, opts ...CleanOption) error {
	return Default.Clean(b, opts...)
}
