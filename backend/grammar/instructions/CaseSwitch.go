package instructions

import (
	Abstract "grammar/abstract"
)

type CaseSwitch struct {
	Expression   Abstract.Expression
	Instructions []interface{}
	Line         int
	Column       int
}

func NewCaseSwitch(Expression Abstract.Expression, Instructions []interface{}, line int, column int) *CaseSwitch {
	return &CaseSwitch{
		Expression:   Expression,
		Instructions: Instructions,
		Line:         line,
		Column:       column,
	}
}

func (c *CaseSwitch) GetOption() Abstract.Expression {
	if c.Expression == nil {
		return nil
	}
	return c.Expression
}

func (c *CaseSwitch) GetInstructions() []interface{} {
	return c.Instructions
}
