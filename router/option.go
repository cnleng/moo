package router

type Options struct {
	Addr string
}

type Option func(*Options)
