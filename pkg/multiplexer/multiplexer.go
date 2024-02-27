package multiplexer

type Multiplexer struct {
  num int
}

func New(age int) *Multiplexer {
  mp := &Multiplexer{num: 1}
  return mp
}
