package ansi

import (
  "fmt"

  "io"
)

func NewWriter(w io.Writer) *Write {
  return &Writer{writer: w}
}
