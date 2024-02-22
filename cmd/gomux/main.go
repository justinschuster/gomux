package main

import (
  "fmt"

  "github.com/justinschuster/gomux/pkg/multiplexer"
)

func main() {
  mp := mulitplexer.New() 
  if err := mp.Start(); err != nil {
    panic(err)
  }

  fmt.Print("\x1bc")
}
