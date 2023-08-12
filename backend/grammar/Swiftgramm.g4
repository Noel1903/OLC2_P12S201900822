grammar Swiftgramm;
import Swiftlex;

@header {
        import "grammar/expressions"
        import "grammar/instructions"
        import "grammar/symbol"
        import "grammar/abstract"
}
s returns [[]interface{} code]
: block EOF{
        $code = $block.blk
} ;


block returns [[]interface{} blk] 
@init{
        $blk = []interface{}{}
        var listInt []ISentenceContext
}
: instr += sentence+
{
        listInt = localctx.(*BlockContext).GetInstr()
        for _,e := range listInt{
                $blk = append($blk,e.GetInstr())
        }
}
;
sentence returns [abstract.Instruction instr]
        : declare_var {$instr = $declare_var.instr}
;



//Declare variables in Swift
declare_var returns [abstract.Instruction instr]
                : VAR ID COLON datatype ASSIGN expression {
                        fmt.Println("Variable declaration: ");
                        $instr = instructions.NewDeclareWithValue($ID.text,$datatype.td,$expression.p)
                }
                | VAR ID ASSIGN expression
                | VAR ID COLON datatype QUESTION_MARK
;



expression returns [abstract.Expression p]
        :left=expression oper=(MULTIPLICATION|DIVISION) right=expression
        | left=expression oper=(SUMMATION|SUBTRACTION) right=expression
        | left=expression oper=(LESS_THAN|LESS_THAN_EQUAL) right=expression
        | left=expression oper=(GREATER_THAN|GREATER_THAN_EQUAL) right=expression
        | left=expression oper=(EQUAL|NOT_EQUAL) right=expression
        | OPEN_PARENTHESIS expression CLOSE_PARENTHESIS
        | NUMBER{
                value,err := strconv.Atoi($NUMBER.text)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.INT)
        }
        | STRING_LITERAL{
                value := $STRING_LITERAL.text[1:len($STRING_LITERAL.text)-1]
                $p = expressions.NewNative(value,symbol.STRING)
        }
        | CHARACTER_LITERAL{
                value := $CHARACTER_LITERAL.text[1:len($CHARACTER_LITERAL.text)-1]
                $p = expressions.NewNative(value,symbol.CHAR)
        }
        | TRUE{
                value,err := strconv.ParseBool($TRUE.text)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.BOOL) 
        }
        | FALSE{
                value,err := strconv.ParseBool($FALSE.text)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.BOOL) 
        }
        | ID{

        }
;

datatype returns [symbol.TypeData td]
        : INT{
                $td = symbol.INT
        }       
        | FLOAT{
                $td = symbol.FLOAT
        }
        | STRING_LITERAL{
                $td = symbol.STRING
        }
        | BOOL{
                $td = symbol.BOOL
        }
        | CHARACTER_LITERAL{
                $td = symbol.CHAR
        }
;