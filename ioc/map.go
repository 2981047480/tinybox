package ioc

import "fmt"

type MapIoc struct {
	Map map[string]Object
}

func (m *MapIoc) Init() error {
	for _, obj := range m.Map {
		if err := obj.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (m *MapIoc) Register(name string, obj Object) error {
	m.Map[name] = obj
	return nil
}

func (m *MapIoc) Get(name string) (interface{}, error) {
	if obj, ok := m.Map[name]; ok {
		return obj, nil
	}
	return nil, fmt.Errorf("obj %s not found", name)
}
