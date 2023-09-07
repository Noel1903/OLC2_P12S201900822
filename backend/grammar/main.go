package main

import (
	"fmt"
	"grammar/abstract"
	Exception "grammar/exceptions"
	"grammar/parser"
	"grammar/symbol"
	"reflect"
	"strconv"

	"io/ioutil"
	"os"
	"os/exec"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type TreeShapeListener struct {
	*parser.BaseSwiftgrammListener
	Code []interface{}
}

type Resp struct {
	Output  string
	Flag    bool
	Message string
}

type Message struct {
	Content string `json:"content"`
}

func Interpreter(c *fiber.Ctx) error {
	var message Message
	if err := c.BodyParser(&message); err != nil {
		return err
	}
	code := message.Content
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftlex(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSwiftgrammParser(tokens)
	p.BuildParseTrees = true
	var response = Resp{}
	tree := p.S()
	//fmt.Println(tree.ToStringTree([]string{}, p))
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	table := symbol.NewEnviorement("global", nil)
	var consoleErr = ""
	var ast symbol.AST
	var reportErr = ""
	ast.SetGlobalTable(table)
	for _, instr := range Code {
		//fmt.Println(instr)
		result := instr.(abstract.Instruction).Execute(table, &ast)
		if reflect.TypeOf(result) == reflect.TypeOf(symbol.ReturnSymbol{}) {

			fmt.Println("SI")
			if result.(symbol.ReturnSymbol).Type == symbol.ERROR {
				err := result.(symbol.ReturnSymbol).Value.(*Exception.Exception)
				reportErr += "<tr><td>" + err.Msg + "</td><td>" + err.Env + "</td><td>" + strconv.Itoa(err.Line) + "</td><td>" + strconv.Itoa(err.Column) + "</td></tr>"
				consoleErr += err.GetError()

			}
		}

	}
	var consoleOut = ""
	consoleOut = ast.GetConsoleOut()
	if consoleErr != "" {
		createErrorTable(reportErr)
		response = Resp{
			Output:  consoleErr,
			Flag:    true,
			Message: "Success",
		}
	} else {
		createSymbolTable(ast.GetReportSymbolTable())
		response = Resp{
			Output:  consoleOut,
			Flag:    true,
			Message: "Success",
		}
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func main() {
	//code := "\nprint(\"Hola\"+\"Mundo\")\nprint(!true)\nprint(5!=8)\nprint(\"a\"==\"a\")\nprint(45.10<20.05)"
	//code := "if 10<10{\nprint(\"Hola\")\n} else if 10>10 {\nprint(\"Mundo\")\n}else{\nprint(\"Adios\")\n}"
	//code := "var i:Int = 10\nprint(i)\nlet numero = 85\nprint(numero)\nnumero+=2\nprint(i)"
	//code := "var i:Int = 0\nwhile i<=25{\nprint(i)\ni+=1\nif i<5{\nprint(\"Numero menor a 5\")\n}else{\nbreak\n}\n}"
	//code := "for i in 1...10{\nprint(i)\nif i==5{\nprint(\"Es 5\")\nreturn\n}\n}"
	//code := "var i = 2\nwhile i<=10{\nguard i%2 == 0 else {\ni+=1\ncontinue\n}\nprint(i)\ni+=1\n}"
	//code := "var vector1:[Character] = []\nprint(vector1)"
	//code := "func factorial(x:Int,y:Int)->Int{\nif x == 0{\nreturn y+1\n}else if x>0 && y==0{return factorial(x-1,1)\n}else{factorial(x-1,factorial(x,y-1))\n}\n}\nprint(factorial(3,4))\n"
	//fmt.Println(code)
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/api/source", Interpreter)
	app.Listen(":3000")

}

func createSymbolTable(symboltable string) {
	header := "\ndigraph G{\nfontname=\"Helvetica,Arial,sans-serif\";\nnode [fontname=\"Helvetica,Arial,sans-serif\"];\n"
	header += "a0 [shape=none label=<\n<TABLE border=\"1\" cellspacing=\"0\" cellpadding=\"10\">\n"
	header += "<tr><td>ID</td><td>Tipo de Simbolo</td><td>Tipo de dato</td><td>Ambito</td><td>Linea</td><td>Columna</td></tr>\n"
	header += symboltable
	header += "</TABLE>>];\n}"

	dotFile := "E:\\USAC2023\\SegundoSemestre\\Compiladores2\\OLC2_P12S201900822\\reports\\graphTable.dot"

	// Escribir el contenido en el archivo .dot
	err := ioutil.WriteFile(dotFile, []byte(header), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo .dot:", err)
		return
	}

	// Ruta del archivo de salida .jpg
	outFile := "E:\\USAC2023\\SegundoSemestre\\Compiladores2\\OLC2_P12S201900822\\reports\\graphTable.jpg"

	// Ejecutar el comando "dot" para convertir el archivo .dot en .jpg
	cmd := exec.Command("dot", "-Tjpg", dotFile, "-o", outFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al ejecutar el comando dot:", err)
		return
	}

	fmt.Println("Archivo .jpg generado con éxito en", outFile)

}

func createErrorTable(symboltable string) {
	header := "\ndigraph G{\nfontname=\"Helvetica,Arial,sans-serif\";\nnode [fontname=\"Helvetica,Arial,sans-serif\"];\n"
	header += "a0 [shape=none label=<\n<TABLE border=\"1\" cellspacing=\"0\" cellpadding=\"10\">\n"
	header += "<tr><td>Descripcion</td><td>Ambito</td><td>Linea</td><td>Columna</td></tr>\n"
	header += symboltable
	header += "</TABLE>>];\n}"

	dotFile := "E:\\USAC2023\\SegundoSemestre\\Compiladores2\\OLC2_P12S201900822\\reports\\graphErr.dot"

	// Escribir el contenido en el archivo .dot
	err := ioutil.WriteFile(dotFile, []byte(header), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo .dot:", err)
		return
	}

	// Ruta del archivo de salida .jpg
	outFile := "E:\\USAC2023\\SegundoSemestre\\Compiladores2\\OLC2_P12S201900822\\reports\\graphErr.jpg"

	// Ejecutar el comando "dot" para convertir el archivo .dot en .jpg
	cmd := exec.Command("dot", "-Tjpg", dotFile, "-o", outFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al ejecutar el comando dot:", err)
		return
	}

	fmt.Println("Archivo .jpg generado con éxito en", outFile)

}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}
