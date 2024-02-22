package multiplexer

import (
  // "fmt"
  
  "github.com/justinschuster/gomux/pane/terminal"
)

type Multiplexer struct {
  activePane pane.Pane
  closeChan     chan struct{}
  ouput         chan byte
  stdoutWriter  *ansi.stdoutWriter
  updateChan    <-chan pane.pane
}

func New() *Multiplexer {
  update := make(chan pane.Pane, 0xff)
  out := make(chan byte, 0xffff)
  stdoutWriter := NewChanWriter(out) 
  terminalPane := pane.newTerminalPane(update)

  mp := &Multiplexer{
    activePane: terminalPane,
    closeChan:  make(chan struct{}),
    stdoutWriter: ansi.NewWriter(stdoutWriter),
  }
  return mp
}
