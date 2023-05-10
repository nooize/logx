package target

import "github.com/nooize/lwr"

func Сondition(first, second lwr.Target, rule lwr.Rule) lwr.Target {
	return &conditionTarget{
		first:  first,
		second: second,
		rule:   rule,
	}
}

type conditionTarget struct {
	first  lwr.Target
	second lwr.Target
	rule   lwr.Rule
}

func (ct *conditionTarget) Handle(e lwr.Event) error {
	if ct.rule(e) {
		return ct.first.Handle(e)
	}
	return ct.second.Handle(e)
}
