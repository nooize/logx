package lwr

type Tags struct {
	parent *Tags
	key    string
	val    any
}

func (t *Tags) Value(key string) any {
	return nil
}

func (t *Tags) ToMap() map[string]interface{} {
	out := make(map[string]interface{})
	if t == nil {
		return out
	}
	entry := t
	for entry != nil {
		if _, ok := out[entry.key]; !ok {
			out[entry.key] = entry.val
		}
		entry = entry.parent
	}
	return out
}
