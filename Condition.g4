grammar Condition;

//操作符
EqualOP:       '=='|'=' ;    
RelationalOP:       '>'|'>='|'<' |'<=' ; 
InOP:       'in';
LeftParen:          '(';
RightParen:         ')';
OrOP: '||';
AndOP: '&&';

//变量
Paramater:                ('model' | 'user') ([a-zA-Z_] | [0-9] | [.])*;
Number    : [0-9]+;

//空白字符，抛弃
Whitespace:         [ \t]+ -> skip;
Newline:            ( '\r' '\n'?|'\n')-> skip;


// Rules
start : expression EOF;


compareOperator
	:	'='
	|	'=='
	|	'>'
	|	'>='
	|	'<='
	|	'in'
	;

expression
   : expression OrOP expression 
   | expression AndOP expression 
   | LeftParen  expression RightParen                           
   | Paramater  compareOperator  Paramater                      
   | Paramater  compareOperator  Number                      
   ;

