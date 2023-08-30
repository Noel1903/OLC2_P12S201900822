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
;

increment_bl returns [abstract.Instruction instr]
: ID INCREMENT expression{
        $instr = instructions.NewIncreDecrem($ID.text,$INCREMENT.text,$expression.p)
} 
;

decrement_bl returns [abstract.Instruction instr]
: ID DECREMENT expression{
        $instr = instructions.NewIncreDecrem($ID.text,$DECREMENT.text,$expression.p)
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
        $instr = instructions.NewReturn(expressions.NewNative(nil,symbol.NIL))
}
;

continue_bl returns [abstract.Instruction instr]
: CONTINUE{
        $instr = instructions.NewContinue("continue")
}
;

declare_let returns [abstract.Instruction instr]
        : LET ID COLON datatype ASSIGN expression{
                $instr = instructions.NewLet($ID.text,$datatype.td,$expression.p)
        }
        | LET ID ASSIGN expression{
                $instr = instructions.NewLet($ID.text,symbol.UNDEFINED,$expression.p)
        }
;

//Declare variables in Swift
declare_var returns [abstract.Instruction instr]
                : VAR ID COLON datatype ASSIGN expression {
                        $instr = instructions.NewDeclareWithValue($ID.text,$datatype.td,$expression.p)
                }
                | VAR ID ASSIGN expression{
                        $instr = instructions.NewDeclareWithValue($ID.text,symbol.UNDEFINED,$expression.p)
                }
                | VAR ID COLON datatype QUESTION_MARK{
                        $instr = instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL))
                }
;

//block print
print_bl returns[abstract.Instruction instr]
: PRINT OPEN_PARENTHESIS expression CLOSE_PARENTHESIS{
        $instr = instructions.NewPrint($expression.p)
}
;

if_bl returns[abstract.Instruction instr]
: IF expression OPEN_kEY ifblock = block CLOSE_kEY{
        $instr = instructions.NewIf($expression.p,$ifblock.blk,nil)
}
| IF expression OPEN_kEY ifblock = block CLOSE_kEY ELSE OPEN_kEY elseblock = block CLOSE_kEY{
        $instr = instructions.NewIf($expression.p,$ifblock.blk,$elseblock.blk)
}
| IF expression OPEN_kEY ifblock = block CLOSE_kEY else_if {
        $instr = instructions.NewIf($expression.p,$ifblock.blk,$else_if.instr)
}
;

else_if returns[[]interface{} instr]
: ELSE IF expression OPEN_kEY ifblock = block CLOSE_kEY{
        $instr = []interface{}{instructions.NewIf($expression.p,$ifblock.blk,nil)}
}
| ELSE IF expression OPEN_kEY ifblock = block CLOSE_kEY ELSE OPEN_kEY elseblock = block CLOSE_kEY{
        $instr = []interface{}{instructions.NewIf($expression.p,$ifblock.blk,$elseblock.blk)}
}
| ELSE IF expression OPEN_kEY ifblock = block CLOSE_kEY else_if{
        $instr = []interface{}{instructions.NewIf($expression.p,$ifblock.blk,$else_if.instr)}
};

while_bl returns[abstract.Instruction instr]
: WHILE expression OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewWhile($expression.p,$block.blk)
}
;

for_bl returns[abstract.Instruction instr]
: FOR ID IN expression1 = expression RANGE expression2 = expression OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewFor($ID.text,$expression1.p,$expression2.p,$block.blk)
}
;

guard_bl returns[abstract.Instruction instr]
: GUARD expression ELSE OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewGuard($expression.p,$block.blk)
}
;

vector_bl returns[abstract.Instruction instr]
: VAR ID COLON OPEN_BRACKET datatype CLOSE_BRACKET ASSIGN  OPEN_BRACKET array_exp CLOSE_BRACKET{
        
        $instr = instructions.NewVector($ID.text,$datatype.td,$array_exp.p)
}
| VAR ID COLON OPEN_BRACKET datatype CLOSE_BRACKET ASSIGN OPEN_BRACKET CLOSE_BRACKET{
        
        $instr = instructions.NewVector($ID.text,$datatype.td,nil)
}
;

array_exp returns[[]interface{} p]
: expression COMMA array_exp{
        $p = append($array_exp.p, $expression.p)
}
| expression{
        $p = []interface{}{$expression.p}
        
}
;

function_bl returns[abstract.Instruction instr]
: FUNC ID OPEN_PARENTHESIS CLOSE_PARENTHESIS ARROW datatype OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewDeclareFunction($ID.text,$datatype.td,[]interface{}{},$block.blk)
}
| FUNC ID OPEN_PARENTHESIS CLOSE_PARENTHESIS OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewDeclareFunction($ID.text,symbol.NIL,[]interface{}{},$block.blk)
}
| FUNC ID OPEN_PARENTHESIS params CLOSE_PARENTHESIS ARROW datatype OPEN_kEY block CLOSE_kEY{
        $instr = instructions.NewDeclareFunction($ID.text,$datatype.td,$params.p,$block.blk)
}
;

params returns[[]interface{} p]
: ID COLON datatype COMMA params{
        $p = append($params.p,instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL)))
}| ID COLON datatype{
        $p = []interface{}{instructions.NewDeclareWithoutValue($ID.text,$datatype.td,expressions.NewNative(nil,symbol.NIL))}
}
;

call_function returns[abstract.Instruction instr]
: ID OPEN_PARENTHESIS CLOSE_PARENTHESIS{
        $instr = instructions.NewCallFunction($ID.text,[]interface{}{})
}
| ID OPEN_PARENTHESIS list_exp CLOSE_PARENTHESIS{
        $instr = instructions.NewCallFunction($ID.text,$list_exp.p)
}
;

list_exp returns[[]interface{} p]
: expression COMMA list_exp{
        $p = append($list_exp.p,$expression.p)
}
| expression{
        $p = []interface{}{$expression.p}
}
;

call_function_exp returns[abstract.Expression p]
: call_function{
        $p = expressions.NewCallFunctionExp($call_function.instr)
}
;

expression returns [abstract.Expression p]
        :left=expression oper=(MULTIPLICATION|DIVISION) right=expression{
                $p = expressions.NewArithmeticOperations($left.p,$right.p,$oper.text)
        }
        | left=expression oper=(SUMMATION|SUBTRACTION) right=expression{
                $p = expressions.NewArithmeticOperations($left.p,$right.p,$oper.text)
        }
        | left=expression oper=MOD right=expression{
                $p = expressions.NewArithmeticOperations($left.p,$right.p,$oper.text)
        }
        | left=expression oper=(LESS_THAN|LESS_THAN_EQUAL) right=expression{
                $p = expressions.NewRelationalOperations($left.p,$right.p,$oper.text)
        }
        | left=expression oper=(GREATER_THAN|GREATER_THAN_EQUAL) right=expression{
                $p = expressions.NewRelationalOperations($left.p,$right.p,$oper.text)
        }
        | left=expression oper=(EQUAL|NOT_EQUAL) right=expression{
                $p = expressions.NewRelationalOperations($left.p,$right.p,$oper.text)
        }
        | left=expression oper=(AND|OR) right=expression{
                $p = expressions.NewLogicOperations($left.p,$right.p,$oper.text)
        }
        | oper = NOT expression{
                $p = expressions.NewLogicOperations(nil,$expression.p,$oper.text)
        }
        | OPEN_PARENTHESIS expression CLOSE_PARENTHESIS
        | call_function_exp{
                $p = $call_function_exp.p
        }
        | NUMBER{
                value,err := strconv.Atoi($NUMBER.text)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.INT)
        }
        | FLOATT{
                value,err := strconv.ParseFloat($FLOATT.text,64)
                if err != nil{
                        fmt.Println(err)
                }
                $p = expressions.NewNative(value,symbol.FLOAT)
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
                $p = expressions.NewIdentifier($ID.text)
        }
        | NIL{
                $p = expressions.NewNative(nil,symbol.NIL)
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