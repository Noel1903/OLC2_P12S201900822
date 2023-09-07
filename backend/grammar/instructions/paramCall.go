package instructions

import (
	Abstract "grammar/abstract"
)

type paramCall struct {
	Internal Abstract.Expression
	External string
}

func NewParamCall(Internal Abstract.Expression, External string) *paramCall {
	return &paramCall{Internal: Internal, External: External}
}
