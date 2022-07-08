package builder

import "io"

type Options struct {
	Token  string
	Output io.Writer
}

type Option func(*Options)

func Token(token string) Option {
	return func(o *Options) {
		o.Token = token
	}
}

func Output(w io.Writer) Option {
	return func(o *Options) {
		o.Output = w
	}
}
