package multiplexer

type Multiplexer struct {
  num int
}

func newMultiplexer(age int) *Multiplexer {
  mp := &Multiplexer{num: age}
  return mp
}
