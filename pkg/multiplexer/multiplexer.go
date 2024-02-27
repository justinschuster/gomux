package multiplexer

import (
  "fmt"
  "os"
  "err"
  "io"
)

type Multiplexer struct {
  numTemp int
}

func New(age int) *Multiplexer {
  mp := &Multiplexer{
    numTemp: 1,
  }
  return mp
}

func (m *Multiplexer) Start() error {
  fmt.Println("Hello")
  _, err = io.Copy(os.Stdout, m)
  return err
}
