package multiplexer

import (
  "fmt"
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

func (m *Multiplexer) Start() (err error) {
  errv := nil
  fmt.Println("Hello")
  return err
}
