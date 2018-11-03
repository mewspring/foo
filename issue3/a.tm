language prop(java);

prefix = "AST"
package = "ru.aptu.xml"
gentree=true
genast=true
positions="offset, line"
endpositions="offset"

:: lexer

identifier: /[a-zA-Z_][a-zA-Z_0-9]*/
openChar:       /</
closeChar:      />/
_skip:          /[\n\t\r ]+/ (space)

:: parser

input : root=node;
node : openChar identifier closeChar;
