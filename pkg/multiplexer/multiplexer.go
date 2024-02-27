package multiplexer

type Multiplexer struct {
  num int
}

func New(age int) *Multiplexer {
  mp := &Multiplexer{num: age}
  return mp
}
