language bar(go);

lang = "bar"
package = "github.com/mewspring/foo/bar"
eventBased = true

::lexer

'foobar' : /foobar/

::parser

input : 'foobar' ;
