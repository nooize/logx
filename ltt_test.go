package ltt

import "testing"

func TestNew(t *testing.T) {
	m := New()
	imp, _ := m.(mux)
}
