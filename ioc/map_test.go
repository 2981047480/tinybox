package ioc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_ioc = NewIoc()

type test_obj struct {
}

func (t *test_obj) Init() error {
	return nil
}

func (t *test_obj) Name() string {
	return "test_obj"
}

func TestRegister(t *testing.T) {
	test_ioc.Register("test", &test_obj{})
	obj, err := test_ioc.Get("test")
	assert.Nil(t, err)
	assert.NotNil(t, obj)
}

func TestGet(t *testing.T) {
	test_ioc.Register("test", &test_obj{})
	test, err := test_ioc.Get("test")
	assert.Nil(t, err)
	assert.Equal(t, &test_obj{}, test)
}
