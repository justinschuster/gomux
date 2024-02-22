package ansi

import (
  "io"
)

type Writer struct {
  writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
  return &Writer{writer: w}
}

func (w *Writer) Write(data []byte) (n int, err error) {
  return w.writer.Write(data)
}
