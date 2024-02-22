package pane

import (
  "github.com/liamg/termutil/pkg/termutil"
)

type TerminaPane struct {
  terminal    *termutil.Terminal
  updateChan  chan<- Pane
  exist       bool
  active      bool
  closeChan   chan struct{}
  startLock   sync.Mutex
  started     bool
}

func NewTerminalPane(updateChan chan<- Pane, term *termutil.Terminal *TerminalPane) {
  return &TerminalPane{
    terminal:   term,
    updateChan: updateChan,
    closeChan:  closeChan,
    exist:      true,
  }

  func (p *TerminalPane) SetActive(target Pane) {
    p.active = p == target
  }
}
