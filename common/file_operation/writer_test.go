package fileoperation

import (
	"testing"
)

func TestNewFWriter(t *testing.T) {
	f, err := NewFWriter("test.txt")
	defer func() {
		err = f.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}
	kbytes, err := f.Write([]byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(kbytes)
}
