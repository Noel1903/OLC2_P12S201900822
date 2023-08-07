lexer grammar Swiftlex;

//////////////////-------Tokens

//types
INT : 'Int';
FLOAT : 'Float';
STRING : 'String';
BOOL : 'Bool';
CHARACTER: 'Character';


//reserved words
TRUE : 'true';
FALSE : 'false';
NIL : 'nil';
VAR: 'var';
LET: 'let';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
CASE: 'case';
BREAK: 'break';
DEFAULT: 'default';
WHILE: 'while';
FOR: 'for';
IN: 'in';
GUARD: 'guard';
CONTINUE: 'continue';
RETURN: 'return';
FUNC: 'func';
STRUCT: 'struct';
MUTATING: 'mutating';
SELF: 'self';
INOUT: 'inout';

//Regex Primitives
NUMBER : [0-9]+('.'[0-9]+)?;
STRING_LITERAL : '"' (~["])* '"';
ID : ([a-zA-Z_])[a-zA-Z0-9_]*;
CHARACTER_LITERAL : '"' (~[']) '"';

//Symbols
SUMMATION : '+';
SUBTRACTION : '-';
MULTIPLICATION : '*';
DIVISION : '/';
MOD: '%';
QUESTION_MARK : '?';
OR : '||';
AND : '&&';
NOT : '!';
EQUAL : '==';
NOT_EQUAL : '!=';
LESS_THAN : '<';
LESS_THAN_EQUAL : '<=';
GREATER_THAN : '>';
GREATER_THAN_EQUAL : '>=';
ASSIGN : '=';
DOT : '.';
COMMA : ',';
COLON : ':';
SEMICOLON : ';';
OPEN_PARENTHESIS : '(';
CLOSE_PARENTHESIS : ')';
OPEN_kEY : '{';
CLOSE_kEY : '}';


///Ignore 
WHITESPACE : [ \t\r\n]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'*'|'?'|'('|')') // Escape sequence
    ;

