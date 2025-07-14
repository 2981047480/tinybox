package fileoperation

import (
	"io"
	"os"
)

type FWriter interface {
	io.Writer
	io.Closer
}

func NewFWriter(filename string) (FWriter, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
