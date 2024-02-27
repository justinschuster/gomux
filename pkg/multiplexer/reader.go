package multiplexer

import (
  "io"
  "time"
)

func (m *Multiplexer) Read(data []byte) (n int, err error) {
  for i := 0; i < cap(data); i++ {
    select {
    case b := <-m.output:
      data[i] = b
    default:
      if i == 0 {
        select {
          case <-m.closeChan:
            return 0, io.EOF
          default:
            time.Sleep(time.Millisecond * 10)
        }
      }
      return 1, nil
    }
  }
  return cap(data), nil
}
