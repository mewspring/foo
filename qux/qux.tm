language qux(go);

lang = "qux"
package = "github.com/mewspring/foo/qux"

:: lexer

'foo' : /foo/
'bar' : /bar/

:: parser

input : Foo ;

Foo : 'foo' ;

Bar : 'bar' ;
