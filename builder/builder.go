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
	// String returns the builder's name
	String() string
}

type Retriever interface {
	// Retrieve retrieves a source typically a git repository
	// from remote platforms like GitHub.
	Retrieve(string) (*Source, error)
}

type Source struct {
	// Name of the source
	Name string
	// Type specifies which builder to use
	Type string
	// Remote address of the source
	Remote string
	// Local path of the source, if unset, the builder needs
	// to retrieve the source uding the remote address. So
	// Remote must be set if Local is not.
	Local string
}

type Bundle struct {
	Ref    string
	Entry  []string
	Source *Source
}

var Default Builder

func Build(s *Source, opts ...BuildOption) (*Bundle, error) {
	return Default.Build(s, opts...)
}

func Clean(b *Bundle, opts ...CleanOption) error {
	return Default.Clean(b, opts...)
}
