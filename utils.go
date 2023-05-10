package lwr

type blackHoleTarget struct {
}

func (t *blackHoleTarget) Handle(_ Event) error {
	return nil
}
