package main

import "io"

type strReader string

func (sr *strReader) Read(p []byte) (n int, err error) {
	n = copy(p, *sr)
	(*sr) = (*sr)[n:]
	if len(*sr) == 0 {
		err = io.EOF
	}

	return
}

func newReader(s string) io.Reader {
	reader := strReader(s)
	return &reader
}
