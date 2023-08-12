package abstract

type Visitor interface {
	visitGetExpression(*Expression)
	visitExecute(*Instruction)
}
