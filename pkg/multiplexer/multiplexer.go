package multiplexer

import (
  "fmt"
)

type Multiplexer struct {
  numTemp int
}

func New() *Multiplexer {
  mp := &Multiplexer{
    numTemp: 1,
  }
  return mp
}

func (m *Multiplexer) Start() error {
  err := error(nil)
  fmt.Println("Hello")
  return err
}
