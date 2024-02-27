package multiplexer

import (
  "fmt"
  // "os"
  // "io"

  // "github.com/justinschuster/gomux/pkg/ansi" 
  // "github.com/justinschuster/gomux/pkg/pane"
)

// type Multiplexer struct {
  // activePane    pane.Pane
  // closeChan     chan struct{}
  // ouput         chan byte
  // updateChan    <-chan pane.Pane
// }

// func New() *Multiplexer {
  // update := make(chan pane.Pane, 0xff)
  // out := make(chan byte, 0xffff)
  // stdoutWriter := NewChanWriter(out) 
  // terminalPane := pane.newTerminalPane(update)

  // mp := &Multiplexer{
    // activePane: terminalPane,
    // output: out,
    // closeChan:  make(chan struct{}),
  // }
  // return mp
// }

// func (m *Multiplexer) Start() error {
  // fmt.Println("Hello")
  // _, err = io.Copy(os.Stdout, m)
  // return err
// }

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
