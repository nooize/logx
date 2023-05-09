package ltt

type Target interface {
	Handle(Event) error
}
