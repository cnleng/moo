package retriever

type Options struct{}

type Option func(*Options)

type RetrieveOptions struct {
	Ref string
}

type RetrieveOption func(RetrieveOptions)

func Ref(ref string) RetrieveOption {
	return func(o RetrieveOptions) {
		o.Ref = ref
	}
}
