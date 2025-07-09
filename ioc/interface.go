package ioc

type Ioc interface {
	Init() error
	Register(name string, obj Object) error
	Get(name string) (interface{}, error)
}

type Object interface {
	Init() error
}
