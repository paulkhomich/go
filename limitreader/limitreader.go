package main

import "io"

type limReader struct {
	reader io.Reader
	limit  int
}

func (lr *limReader) Read (p []byte) (n int, err error) {
	if lr.limit <= len(p) {
		n, err = lr.reader.Read(p[:lr.limit])
		err = io.EOF
	} else {
		n, err = lr.reader.Read(p)
	}
	lr.limit -= n

	return
}

func limitReader (r io.Reader, n int) io.Reader {
	reader := limReader{r, n}
	return &reader
}
