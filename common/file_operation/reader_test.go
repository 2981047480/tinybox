package fileoperation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFReader(t *testing.T) {
	f, err := NewFReader("test.txt")
	defer func() {
		err = f.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}
	var p = make([]byte, 4)
	kbytes, err := f.Read(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
	t.Log(kbytes)
	assert.Equal(t, "test", string(p))
}
