package logx

import "testing"

func TestTags_Value(t *testing.T) {
	/*
		tag := &Tags{
			key: "foo",
			val: "bar",
		}
		if v := tag.Value("foo"); v != "bar" {
			t.Errorf("expected %v, got %v", "bar", v)
		}
		if v := tag.Value("bar"); v != nil {
			t.Errorf("expected %v, got %v", nil, v)
		}
		tag = &Tags{
			parent: tag,
			key:    "foo2",
			val:    "bar2",
		}
		if v := tag.Value("foo2"); v != "bar2" {
			t.Errorf("expected %v, got %v", "bar2", v)
		}
		if v := tag.Value("foo"); v != "bar" {
			t.Errorf("expected %v, got %v", "bar", v)
		}
	*/
}

func TestTags_ToMap(t *testing.T) {
	tag := &Tags{
		parent: &Tags{
			key: "foo2",
			val: "bar2",
		},
		key: "foo",
		val: "bar",
	}
	mp := tag.ToMap()
	v, ok := mp["foo"]
	if !ok {
		t.Errorf("expected %v, got nothing", "bar")
	}
	if v != "bar" {
		t.Errorf("expected %v, got %v", "bar", v)
	}
	v, ok = mp["foo2"]
	if !ok {
		t.Errorf("expected %v, got nothing", "bar2")
	}
	if v != "bar2" {
		t.Errorf("expected %v, got %v", "bar2", v)
	}
}
