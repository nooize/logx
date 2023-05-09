package target

import "github.com/nooize/ltt"

func Сondition(first, second ltt.Target, rule ltt.Rule) ltt.Target {
	return &conditionTarget{
		first:  first,
		second: second,
		rule:   rule,
	}
}

type conditionTarget struct {
	first  ltt.Target
	second ltt.Target
	rule   ltt.Rule
}

func (ct *conditionTarget) Handle(e ltt.Event) error {
	if ct.rule(e) {
		return ct.first.Handle(e)
	}
	return ct.second.Handle(e)
}
