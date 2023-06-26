package target

import "github.com/nooize/lux"

func Сondition(first, second lux.Target, rule lux.Rule) lux.Target {
	return &conditionTarget{
		first:  first,
		second: second,
		rule:   rule,
	}
}

type conditionTarget struct {
	first  lux.Target
	second lux.Target
	rule   lux.Rule
}

func (ct *conditionTarget) Handle(e lux.Event) error {
	if ct.rule(e) {
		return ct.first.Handle(e)
	}
	return ct.second.Handle(e)
}
