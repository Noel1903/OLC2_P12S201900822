package exceptions

import "strconv"

type Exception struct {
	Msg    string
	Env    string
	Line   int
	Column int
}

func NewException(msg string, env string, line int, column int) *Exception {
	return &Exception{
		Msg:    msg,
		Env:    env,
		Line:   line,
		Column: column,
	}
}

func (e *Exception) GetError() string {
	return "Error: " + e.Msg + " en la linea " + strconv.Itoa(e.Line) + " y columna " + strconv.Itoa(e.Column) + " en el entorno " + e.Env + "\n"
}

func (e *Exception) GetMsg() string {
	return e.Msg
}

func (e *Exception) GetEnv() string {
	return e.Env
}

func (e *Exception) GetLine() int {
	return e.Line
}

func (e *Exception) GetColumn() int {
	return e.Column
}

func (e *Exception) SetMsg(msg string) {
	e.Msg = msg
}

func (e *Exception) SetEnv(env string) {
	e.Env = env
}

func (e *Exception) SetLine(line int) {
	e.Line = line
}

func (e *Exception) SetColumn(column int) {
	e.Column = column
}
