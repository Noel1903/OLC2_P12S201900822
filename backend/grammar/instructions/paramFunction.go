package instructions

import (
	Abstract "grammar/abstract"
)

type paramFunction struct {
	Internal Abstract.Instruction
	External string
}

func NewParamFunction(Internal Abstract.Instruction, External string) *paramFunction {
	return &paramFunction{Internal: Internal, External: External}
}
