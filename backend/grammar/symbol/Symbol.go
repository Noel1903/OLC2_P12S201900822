package symbol

type Type int

const (
	MUTABLE Type = iota
	UNMUTABLE
)

type Symbol struct {
	Line     int
	Column   int
	Value    ReturnSymbol
	Type     Type
	Position int
	IsGlobal bool
	InHeap   bool
	TypeArr  TypeData
}

func (s *Symbol) GetPos() int {
	return s.Position
}

func (s *Symbol) GetIsGlobal() bool {
	return s.IsGlobal
}

func (s *Symbol) SetPos(pos int) {
	s.Position = pos
}
