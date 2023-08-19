package instructions

import (
	Abstract "grammar/abstract"
)

type If struct {
	condition Abstract.Expression
	code      []interface{}
}

func NewIf(condition Abstract.Expression, code []interface{}) If {
	return If{
		condition: condition,
		code:      code,
	}
}
