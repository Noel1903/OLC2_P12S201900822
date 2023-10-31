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
        | declare_let {$instr = $declare_let.instr}
        | assign_bl {$instr = $assign_bl.instr}
        | assign_vector {$instr = $assign_vector.instr}
        | print_bl {$instr = $print_bl.instr}
        | if_bl {$instr = $if_bl.instr}
        | increment_bl {$instr = $increment_bl.instr}
        | decrement_bl {$instr = $decrement_bl.instr}
        | while_bl {$instr = $while_bl.instr}
        | for_bl {$instr = $for_bl.instr}
        | guard_bl {$instr = $guard_bl.instr}
        | break_bl {$instr = $break_bl.instr}
        | return_bl {$instr = $return_bl.instr}
        | continue_bl {$instr = $continue_bl.instr}
        | vector_bl {$instr = $vector_bl.instr}
        | function_bl {$instr = $function_bl.instr}
        | call_function {$instr = $call_function.instr}
        | array_functions {$instr = $array_functions.instr}
        | declare_array_bl { $instr = $declare_array_bl.instr}
        | switch_bl {$instr = $switch_bl.instr}
;




switch_bl returns [abstract.Instruction instr]
: SWITCH expression OPEN_kEY cases CLOSE_kEY{
        $instr = instructions.NewSwitch($expression.p,$cases.p,$SWITCH.line,$SWITCH.pos)
}
;


cases returns [[]interface{} p]
: CASE expression COLON block cases{
        value := instructions.NewCaseSwitch($expression.p,$block.blk,$CASE.line,$CASE.pos)
        $p = append([]interface{}{value},$cases.p...)
}
|CASE expression COLON block{
        value := instructions.NewCaseSwitch($expression.p,$block.blk,$CASE.line,$CASE.pos)
        $p = []interface{}{value}
}


| DEFAULT COLON block{
        value := instructions.NewCaseSwitch(nil,$block.blk,$DEFAULT.line,$DEFAULT.pos)
        $p = []interface{}{value}       
}
;

increment_bl returns [abstract.Instruction instr]
: ID INCREMENT expression{
        $instr = instructions.NewIncreDecrem($ID.text,$INCREMENT.text,$expression.p,$ID.line,$ID.pos)
} 
;

decrement_bl returns [abstract.Instruction instr]
: ID DECREMENT expression{
        $instr = instructions.NewIncreDecrem($ID.text,$DECREMENT.text,$expression.p,$ID.line,$ID.pos)
}
;

break_bl returns [abstract.Instruction instr]
: BREAK{
        $instr = instructions.NewBreak("break")
}
;

return_bl returns [abstract.Instruction instr]
: RETURN expression{
        $instr = instructions.NewReturn($expression.p)
}
| RETURN{
        $instr = instructions.NewReturn(expressions.NewNative(nil,symbol.NIL,$RETURN.line,$RETURN.pos))
}
;

continue_bl returns [abstract.Instruction instr]
: CONTINUE{
        $instr = instructions.NewContinue("continue")
}
;

declare_let returns [abstract.Instruction instr]
        : LET ID COLON datatype ASSIGN expression{
                $instr = instructions.NewLet($ID.text,$datatype.td,$expression.p,$ID.line,$ID.pos)
        }
        | LET ID ASSIGN expression{
                $instr = instructions.NewLet($ID.text,symbol.UNDEFINED,$expression.p,$ID.line,$ID.pos)
        }
;

native_function returns [abstract.Expression p]
: STRING OPEN_PARENTHESIS expression CLOSE_PARENTHESIS{
        $p = expressions.NewNativeFunction($STRING.text,$expression.p)
}
| INT OPEN_PARENTHESIS expression CLOSE_PARENTHESIS{
        $p = expressions.NewNativeFunction($INT.text,$expression.p)
}
| FLOAT OPEN_PARENTHESIS expression CLOSE_PARENTHESIS{
        $p = expressions.NewNativeFunction($FLOAT.text,$expression.p)
}
;

//Declare variables in Swift
declare_var returns [abstract.Instruction instr]
                : VAR ID COLON datatype ASSIGN expression {
                        $instr = instructions.NewDeclareWithValue($ID.text,$datatype.td,$expression.p,$ID.line,$ID.pos)
                }
                | VAR ID ASSIGN expression{
                        $instr = instructions.NewDeclareWithValue($ID.text,symbol.UNDEFINED,$expression.p,$ID.line,$ID.pos)
                }
                | VAR ID COLON datatype QUESTION_MARK{
                        $instr = instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL,$QUESTION_MARK.line,$QUESTION_MARK.pos),$ID.line,$ID.pos)
                }
;

assign_bl returns [abstract.Instruction instr]
: ID ASSIGN expression{
        $instr = instructions.NewAsignVariable($ID.text,$expression.p,$ID.line,$ID.pos)
}
;

//block print
print_bl returns[abstract.Instruction instr]
: PRINT OPEN_PARENTHESIS list_print CLOSE_PARENTHESIS{
        $instr = instructions.NewPrint($list_print.p)
}
;

list_print returns[[]interface{} p]
: expression COMMA list_print{
        $p = append([]interface{}{$expression.p},$list_print.p...)
}
| expression{
        $p = []interface{}{$expression.p}
}
;
if_bl returns[abstract.Instruction instr]
: IF expression OPEN_kEY ifblock = block CLOSE_kEY{
        $instr = instructions.NewIf($expression.p,$ifblock.blk,nil,$IF.line,$IF.pos)
}
| IF expression OPEN_kEY ifblock = block CLOSE_kEY ELSE OPEN_kEY elseblock = block CLOSE_kEY{
        $instr = instructions.NewIf($expression.p,$ifblock.blk,$elseblock.blk,$IF.line,$IF.pos)
}
| IF expression OPEN_kEY ifblock = block CLOSE_kEY else_if {
        $instr = instructions.NewIf($expression.p,$ifblock.blk,$else_if.instr,$IF.line,$IF.pos)
}
;

else_if returns[[]interface{} instr]
: ELSE IF expression OPEN_kEY ifblock = block CLOSE_kEY{
        $instr = []interface{}{instructions.NewIf($expression.p,$ifblock.blk,nil,$ELSE.line,$ELSE.pos)}
}
| ELSE IF expression OPEN_kEY ifblock = block CLOSE_kEY ELSE OPEN_kEY elseblock = block CLOSE_kEY{
        $instr = []interface{}{instructions.NewIf($expression.p,$ifblock.blk,$elseblock.blk,$ELSE.line,$ELSE.pos)}
}
| ELSE IF expression OPEN_kEY ifblock = block CLOSE_kEY else_if{
        $instr = []interface{}{instructions.NewIf($expression.p,$ifblock.blk,$else_if.instr,$ELSE.line,$ELSE.pos)}
};

while_bl returns[abstract.Instruction instr]
: WHILE expression OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewWhile($expression.p,$block.blk,$WHILE.line,$WHILE.pos)
}
;

for_bl returns[abstract.Instruction instr]
: FOR ID IN expression1 = expression RANGE expression2 = expression OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewFor($ID.text,$expression1.p,$expression2.p,$block.blk,$ID.line,$ID.pos)
}
;

guard_bl returns[abstract.Instruction instr]
: GUARD expression ELSE OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewGuard($expression.p,$block.blk,$GUARD.line,$GUARD.pos)
}
;

vector_bl returns[abstract.Instruction instr]
: VAR ID COLON OPEN_BRACKET datatype CLOSE_BRACKET ASSIGN  OPEN_BRACKET array_exp CLOSE_BRACKET{
        
        $instr = instructions.NewVector($ID.text,$datatype.td,$array_exp.p,$ID.line,$ID.pos)
}
| VAR ID COLON OPEN_BRACKET datatype CLOSE_BRACKET ASSIGN OPEN_BRACKET CLOSE_BRACKET{
        
        $instr = instructions.NewVector($ID.text,$datatype.td,nil,$ID.line,$ID.pos)
}
;

array_exp returns[[]interface{} p]
: expression COMMA array_exp{
        $p = append([]interface{}{$expression.p},$array_exp.p...)
}
| expression{
        $p = []interface{}{$expression.p}
        
}
;

array_functions returns[abstract.Instruction instr]
: ID DOT APPEND OPEN_PARENTHESIS expression CLOSE_PARENTHESIS{
        $instr = instructions.NewArrayFunctions($ID.text,$APPEND.text,$expression.p,$ID.line,$ID.pos)
}
| ID DOT REMOVELAST OPEN_PARENTHESIS CLOSE_PARENTHESIS{
        $instr = instructions.NewArrayFunctions($ID.text,$REMOVELAST.text,nil,$ID.line,$ID.pos)
}
| ID DOT REMOVE OPEN_PARENTHESIS AT COLON expression CLOSE_PARENTHESIS{
        $instr = instructions.NewArrayFunctions($ID.text,$REMOVE.text,$expression.p,$ID.line,$ID.pos)
}
;

assign_vector returns[abstract.Instruction instr]
: ID OPEN_BRACKET expression CLOSE_BRACKET ASSIGN expression{
        $instr = instructions.NewAsignVector($ID.text,$expression.p,$expression.p,$ID.line,$ID.pos)
}
;

function_bl returns[abstract.Instruction instr]
: FUNC ID OPEN_PARENTHESIS CLOSE_PARENTHESIS ARROW datatype OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewDeclareFunction($ID.text,$datatype.td,[]interface{}{},$block.blk,$ID.line,$ID.pos)
}
| FUNC ID OPEN_PARENTHESIS CLOSE_PARENTHESIS OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewDeclareFunction($ID.text,symbol.NIL,[]interface{}{},$block.blk,$ID.line,$ID.pos)
}
| FUNC ID OPEN_PARENTHESIS params CLOSE_PARENTHESIS ARROW datatype OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewDeclareFunction($ID.text,$datatype.td,$params.p,$block.blk,$ID.line,$ID.pos)
}
| FUNC ID OPEN_PARENTHESIS params CLOSE_PARENTHESIS OPEN_kEY block CLOSE_kEY{
        //fmt.Println($params.p)
        $instr = instructions.NewDeclareFunction($ID.text,symbol.NIL,$params.p,$block.blk,$ID.line,$ID.pos)
}
;

params returns[[]interface{} p]
: ID COLON datatype COMMA params{
        value:=instructions.NewParamFunction(instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL,$ID.line,$ID.pos),$ID.line,$ID.pos),$ID.text)
        $p = append([]interface{}{value},$params.p...)

}| ID COLON datatype{
        value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL,$ID.line,$ID.pos),$ID.line,$ID.pos),$ID.text)
        $p = []interface{}{value}
}
| extern_params ID COLON datatype COMMA params{
        value:=instructions.NewParamFunction(instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL,$ID.line,$ID.pos),$ID.line,$ID.pos),$extern_params.text)
        $p = append([]interface{}{value},$params.p...)

}| extern_params ID COLON datatype{
        value := instructions.NewParamFunction(instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL,$ID.line,$ID.pos),$ID.line,$ID.pos),$extern_params.text)
        $p = []interface{}{value}
}| ID COLON OPEN_BRACKET datatype CLOSE_BRACKET COMMA params{
        value:=instructions.NewParamFunction(instructions.NewVector($ID.text,$datatype.td,nil,$ID.line,$ID.pos),$ID.text)
        $p = append([]interface{}{value},$params.p...)

}| ID COLON OPEN_BRACKET datatype CLOSE_BRACKET{
        value := instructions.NewParamFunction(instructions.NewVector($ID.text,$datatype.td,nil,$ID.line,$ID.pos),$ID.text)
        $p = []interface{}{value}
}
| extern_params ID COLON OPEN_BRACKET datatype CLOSE_BRACKET COMMA params{
        value:=instructions.NewParamFunction(instructions.NewVector($ID.text,$datatype.td,nil,$ID.line,$ID.pos),$extern_params.text)
        $p = append([]interface{}{value},$params.p...)

}| extern_params ID COLON  OPEN_BRACKET datatype CLOSE_BRACKET{
        value := instructions.NewParamFunction(instructions.NewVector($ID.text,$datatype.td,nil,$ID.line,$ID.pos),$extern_params.text)
        $p = []interface{}{value}
}
;

extern_params returns[string p]
: ID{
        $p = $ID.text
}
| UNDERSCORE{
        $p = $UNDERSCORE.text
}
;

call_function returns[abstract.Instruction instr]
: ID OPEN_PARENTHESIS CLOSE_PARENTHESIS{
        $instr = instructions.NewCallFunction($ID.text,[]interface{}{},$ID.line,$ID.pos)
}
| ID OPEN_PARENTHESIS list_exp CLOSE_PARENTHESIS{
        $instr = instructions.NewCallFunction($ID.text,$list_exp.p,$ID.line,$ID.pos)
}
;

list_exp returns[[]interface{} p]
: expression COMMA list_exp{
        value := instructions.NewParamCall($expression.p,"_")
        $p = append([]interface{}{value},$list_exp.p...)
}
| expression{
        value := instructions.NewParamCall($expression.p,"_")
        $p = []interface{}{value}
}
| ID COLON expression COMMA list_exp{
        value := instructions.NewParamCall($expression.p,$ID.text)
        $p = append([]interface{}{value},$list_exp.p...)
}
| ID COLON expression{
        value := instructions.NewParamCall($expression.p,$ID.text)
        $p = []interface{}{value}
}
;

call_function_exp returns[abstract.Expression p]
: call_function{
        $p = expressions.NewCallFunctionExp($call_function.instr)
}
;

declare_array_bl returns [abstract.Instruction instr]
:VAR ID COLON type_matrix ASSIGN OPEN_BRACKET exp_matriz CLOSE_BRACKET{
        $instr = instructions.NewDeclareArray($ID.text,$type_matrix.td,$exp_matriz.p)
}
;

exp_matriz returns [[]interface{} p]
:  expression {
        value := $expression.p
        $p = []interface{}{value}
}
|  expression COMMA exp_matriz{ 
        value := $expression.p
        $p = append([]interface{}{value},$exp_matriz.p...)
}
;

type_matrix returns [symbol.TypeData td]
: OPEN_BRACKET type_matrix CLOSE_BRACKET{
        $td = $type_matrix.td
}| OPEN_BRACKET datatype CLOSE_BRACKET{
        $td = $datatype.td
}
;

definition_matrix returns [[][]interface{} p]
: type_matrix 
;

matriz_pos returns [[]interface{} p]
: OPEN_BRACKET expression CLOSE_BRACKET{
        value := $expression.p
        $p = []interface{}{value}
}
| OPEN_BRACKET expression CLOSE_BRACKET matriz_pos{
        value := $expression.p
        $p = append([]interface{}{value},$matriz_pos.p...)
}
;


expression returns [abstract.Expression p]
        :left=expression oper=(MULTIPLICATION|DIVISION) right=expression{
                $p = expressions.NewArithmeticOperations($left.p,$right.p,$oper.text,$oper.line,$oper.pos)
        }
        | left=expression oper=(SUMMATION|SUBTRACTION) right=expression{
                $p = expressions.NewArithmeticOperations($left.p,$right.p,$oper.text,$oper.line,$oper.pos)
        }
        | left=expression oper=MOD right=expression{
                $p = expressions.NewArithmeticOperations($left.p,$right.p,$oper.text,$oper.line,$oper.pos)
        }
        | left=expression oper=(LESS_THAN|LESS_THAN_EQUAL) right=expression{
                $p = expressions.NewRelationalOperations($left.p,$right.p,$oper.text,$oper.line,$oper.pos)
        }
        | left=expression oper=(GREATER_THAN|GREATER_THAN_EQUAL) right=expression{
                $p = expressions.NewRelationalOperations($left.p,$right.p,$oper.text,$oper.line,$oper.pos)
        }
        | left=expression oper=(EQUAL|NOT_EQUAL) right=expression{
                $p = expressions.NewRelationalOperations($left.p,$right.p,$oper.text,$oper.line,$oper.pos)
        }
        | left=expression oper=(AND|OR) right=expression{
                $p = expressions.NewLogicOperations($left.p,$right.p,$oper.text,$oper.line,$oper.pos)
        }
        | oper = NOT expression{
                $p = expressions.NewLogicOperations(nil,$expression.p,$oper.text,$oper.line,$oper.pos)
        }
        | oper = '-' expression{
                $p = expressions.NewNegative($expression.p,$oper.line,$oper.pos)
        }
        | OPEN_PARENTHESIS expression CLOSE_PARENTHESIS{
                $p = $expression.p
        }
        | call_function_exp{
                $p = $call_function_exp.p
        }
        | native_function{
                $p = $native_function.p
        }
        | ID OPEN_BRACKET expression CLOSE_BRACKET{
                $p = expressions.NewGetPosVector($ID.text,$expression.p,$ID.line,$ID.pos)
        }
        
        | ID DOT COUNT{
                $p = expressions.NewVectorValues($ID.text,$COUNT.text,$ID.line,$ID.pos)
        }
        | ID DOT ISEMPTY{
                $p = expressions.NewVectorValues($ID.text,$ISEMPTY.text,$ID.line,$ID.pos)
        }
        | OPEN_BRACKET array_exp CLOSE_BRACKET{
                $p = expressions.NewNative($array_exp.p,symbol.ARRAY,$OPEN_BRACKET.line,$OPEN_BRACKET.pos)
        }
        | NUMBER{
                value,err := strconv.Atoi($NUMBER.text)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.INT,$NUMBER.line,$NUMBER.pos)
        }
        | ID matriz_pos {
                $p = expressions.NewGetPosMatrix($ID.text,$matriz_pos.p,$ID.line,$ID.pos)
        }

        | FLOATT{
                value,err := strconv.ParseFloat($FLOATT.text,64)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.FLOAT,$FLOATT.line,$FLOATT.pos)
        }
        | STRING_LITERAL{
                value := $STRING_LITERAL.text[1:len($STRING_LITERAL.text)-1]
                $p = expressions.NewNative(value,symbol.STRING,$STRING_LITERAL.line,$STRING_LITERAL.pos)
        }
        | CHARACTER_LITERAL{
                value := $CHARACTER_LITERAL.text[1:len($CHARACTER_LITERAL.text)-1]
                $p = expressions.NewNative(value,symbol.CHAR,$CHARACTER_LITERAL.line,$CHARACTER_LITERAL.pos)
        }
        | TRUE{
                value,err := strconv.ParseBool($TRUE.text)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.BOOL,$TRUE.line,$TRUE.pos) 
        }
        | FALSE{
                value,err := strconv.ParseBool($FALSE.text)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.BOOL,$FALSE.line,$FALSE.pos) 
        }
        | ID{
                $p = expressions.NewIdentifier($ID.text,$ID.line,$ID.pos)
        }
        | NIL{
                $p = expressions.NewNative(nil,symbol.NIL,$NIL.line,$NIL.pos)
        }
        
;

datatype returns [symbol.TypeData td]
        : INT{
                $td = symbol.INT
        }       
        | FLOAT{
                $td = symbol.FLOAT
        }
        | STRING{
                $td = symbol.STRING
        }
        | BOOL{
                $td = symbol.BOOL
        }
        | CHARACTER{
                $td = symbol.CHAR
        }
;