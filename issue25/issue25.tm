language issue25(go);

lang = "issue25"
package = "github.com/mewspring/foo/issue25"
eventBased = true
eventFields = true
eventAST = true

:: lexer

id_tok : /[a-z]/
'cs' : /cs/
'delim' : /delim/

whitespace : /[ \t\r\n]+/     (space)

:: parser

%input Foo;

Foo -> Foo
   : 'cs' Bars=ID+ 'delim' Bar=ID
;

ID -> ID
   : id_tok
;
