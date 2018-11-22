language shiftreduce(go);

lang = "shiftreduce"
package = "github.com/mewspring/foo/shiftreduce"
eventBased = true
eventFields = true
eventAST = true

:: lexer

'if' : /if/
'else' : /else/
'return' : /return/
'(' : /[(]/
')' : /[)]/

int_lit_tok : /[0-9]+/
bool_lit_tok : /true|false/

:: parser

%input Stmt ;

File
	: Stmt*
;

Stmt
	: IfStmt
	| ReturnStmt
;

IfStmt
	: 'if' '(' Expr ')' Stmt %shift 'else'
	| 'if' '(' Expr ')' Stmt 'else' Stmt
;

ReturnStmt
	: 'return' Expr
;

Expr
	: IntLit
	| BoolLit
;

IntLit
	: int_lit_tok
;

BoolLit
	: bool_lit_tok
;
