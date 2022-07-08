package server

type Options struct {
	Addr string
}

type Option func(*Options)

func Addr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}
