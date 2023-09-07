package symbol

type Type int

const (
	MUTABLE Type = iota
	UNMUTABLE
)

type Symbol struct {
	Line   int
	Column int
	Value  ReturnSymbol
	Type   Type
}
