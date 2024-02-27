package multiplexer

import (
  "fmt"
)

type Multiplexer struct {
  outputChan chan byte
}

func New() *Multiplexer {
  out := make(chan bye, 0xffff)
  mp := &Multiplexer{
    output: out,
  }
  return mp
}

func (m *Multiplexer) Start() error {
  fmt.Println("Hello")
  _, err = io.Copy(os.Stdout, m)
  return err
}
