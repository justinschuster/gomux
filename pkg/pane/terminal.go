package pane

import (
  "github.com/liamg/termutil/pkg/termutil"
)

type TerminalPane struct {
  terminal    *termutil.Terminal
  updateChan  chan<- Pane
  exist       bool
  active      bool
}

func NewTerminalPane(updateChan chan<- Pane, term *termutil.Terminal) *TerminalPane {
  return &TerminalPane{
    terminal:   term,
    updateChan: updateChan,
    exist:      true,
  }
}
 
