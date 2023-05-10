package lwr

type Target interface {
	Handle(Event) error
}
