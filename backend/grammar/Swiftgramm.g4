grammar Swiftgramm;
import Swiftlex;

s: block EOF;

block: (sentences)*
;

sentences: sentence sentences
| sentence 
;

sentence: declare_var
        | declare_constant
        | assign_var
;

//block print
print: PRINT OPEN_PARENTHESIS expression CLOSE_PARENTHESIS
;

//Declare variables in Swift
declare_var: VAR ID COLON datatype ASSIGN expression
            | VAR ID ASSIGN expression
            | VAR ID COLON datatype QUESTION_MARK
;

//Declare constants in Swift

declare_constant: LET ID COLON datatype ASSIGN expression
                | LET ID ASSIGN expression
                | LET ID COLON datatype QUESTION_MARK
;

//Assign variables in Swift
assign_var: ID ASSIGN expression
;

//If sentence in Swift
if_sentence: IF expression sentences
            | IF expression sentences ELSE sentences
;
//Switch sentence in Swift
switch_sentence: SWITCH expression OPEN_kEY switch_cases CLOSE_kEY
;

switch_cases: switch_case switch_cases
            | switch_cases
;
switch_case: CASE expression COLON sentences
            | DEFAULT COLON sentences
;
//While sentence in Swift
while_sentence: WHILE expression OPEN_kEY sentences CLOSE_kEY
;
//for sentence in Swift
for_sentence: FOR ID IN expression OPEN_kEY sentences CLOSE_kEY
;

//guard sentence in Swift
guard_sentence: GUARD expression ELSE sentences
;

//break sentence in Swift
break_sentence: BREAK
;

//continue sentence in Swift
continue_sentence: CONTINUE
;

//return sentence in Swift
return_sentence: RETURN expression
                | RETURN
;

expression: left=expression op=(MOD) right=expression
        | left=expression op=(MULTIPLICATION|DIVISION) right=expression
        | left=expression op=(SUMMATION|SUBTRACTION) right=expression
        | left=expression op=(LESS_THAN|LESS_THAN_EQUAL) right=expression
        | left=expression op=(GREATER_THAN|GREATER_THAN_EQUAL) right=expression
        | left=expression op=(EQUAL|NOT_EQUAL) right=expression
        | left=expression op=(AND) right=expression
        | left=expression op=(OR) right=expression
        | op=(NOT) right=expression
        | OPEN_PARENTHESIS expression CLOSE_PARENTHESIS
        | NUMBER
        | STRING_LITERAL
        | CHARACTER_LITERAL
        | ID
;

datatype: INT
        | FLOAT
        | STRING
        | BOOL
;