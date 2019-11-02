package main

import "io"

type writerCount struct {
	writer  io.Writer
	counter int64
}

func (w *writerCount) Write(bytes []byte) (int, error) {
	n, err := w.writer.Write(bytes)
	w.counter += int64(n)

	return n, err
}

func countingWirter(w io.Writer) (io.Writer, *int64) {
	wc := writerCount{w, 0}

	return &wc, &wc.counter
}
