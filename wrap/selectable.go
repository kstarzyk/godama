package wrap

type Selectable interface {
  Focus() error
  Defocus() error
  HandleKey(key string) error
}
