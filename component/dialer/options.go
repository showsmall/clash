package dialer

var DefaultOptions []Option

type option struct {
	skipDefault   bool
	interfaceName string
	addrReuse     bool
}

type Option func(opt *option)

func WithInterface(name string) Option {
	return func(opt *option) {
		opt.interfaceName = name
	}
}

func WithAddrReuse(reuse bool) Option {
	return func(opt *option) {
		opt.addrReuse = reuse
	}
}

func WithSkipDefault(skip bool) Option {
	return func(opt *option) {
		opt.skipDefault = skip
	}
}
