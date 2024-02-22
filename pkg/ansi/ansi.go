package ansi

import (
  "fmt"

  "io"
)

type Writer struct {
  writer io.Writer
}

func NewWriter(w io.Writer) *Write {
  return &Writer{writer: w}
}

func (w *Writer) Write(data []byte) (n int, err error) {
  return w.writer.Write(data)
}
