package ioc

var DefaultIoc = NewIoc()

func NewIoc() Ioc {
	return &MapIoc{
		Map: make(map[string]Object),
	}
}
