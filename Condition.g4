grammar Condition;

//操作符
EqualOP:       '==' ;
LargerOp: '>';
LargerEqualOp: '>='  ;
LessOp: '<';
LessEqualOp : '<=';
InOP:       'in';
LeftParen:          '(';
RightParen:         ')';
OrOP: '||';
AndOP: '&&';

//变量
Paramater:                ('model' | 'user')([a-zA-Z_] | [0-9] | [.])*;
Number    : [0-9]+;

//空白字符，抛弃
Whitespace:         [ \t]+ -> skip;
Newline:            ( '\r' '\n'?|'\n')-> skip;


// Rules
start : expression;

expression
   : '('expression ')' # parensExpression
   | left=Paramater  op=('=='|'>'|'>='|'<'|'<='|'in')  right=Number # compare
   | left=Paramater  op=('=='|'>'|'>='|'<'|'<='|'in')  right=Paramater # compare
   | left=expression AndOP right=expression # andCompare
   | left=expression OrOP right=expression # orCompare
   ;

