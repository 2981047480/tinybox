package fileoperation

import (
	"io"
	"os"
)

type FReader interface {
	io.Reader
	io.Closer
}

func NewFReader(filename string) (FReader, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
