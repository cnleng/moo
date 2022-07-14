package noop

import "github.com/moobu/moo/builder/retriever"

type noop struct{}

func (n noop) Retrieve(url string, opts ...retriever.RetrieveOption) (*retriever.Repository, error) {
	return &retriever.Repository{Path: url}, nil
}

func (n noop) String() string {
	return "noop"
}

func New(opts ...retriever.Option) retriever.Retriever {
	return noop{}
}
