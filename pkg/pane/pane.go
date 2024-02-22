package pane

type Pane interface {
  Exists() bool
  Close()
}
