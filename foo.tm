language foo(go);

lang = "foo"
package = "github.com/mewspring/foo"

::lexer

'foobar' : /foobar/

::parser

input : 'foobar' ;
