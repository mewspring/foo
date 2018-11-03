language issue24(go);

lang = "issue24"
package = "github.com/mewspring/foo/issue24"
eventBased = true
eventFields = true
eventAST = true

:: lexer

id_tok : /[a-z]/
'gep' : /gep/
',' : /,/

whitespace : /[ \t\r\n]+/     (space)

:: parser

%input Foo;

Foo -> Foo
   : 'gep' Src=ID Indices=(',' ID)*
;

ID -> ID
   : id_tok
;
