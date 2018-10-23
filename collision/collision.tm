language collision(go);

lang = "collision"
package = "github.com/mewspring/foo/collision"
eventBased = true
eventFields = true

:: lexer

'gc' : /gc/

string_lit : /["][^"]*["]/

:: parser

input : FuncHeader ;

FuncHeader -> FuncHeader
	: GCopt
;

# TODO: Rename GCNode to GC when collision with token 'gc' has been resolved.
#
# GC is defined as an identifier in both listener.go and token.go.
#
#    lalr: 0.014s, text: 0.084s, parser: 8 states, 0KB
#    # github.com/mewspring/foo/collision
#    ./token.go:18:2: GC redeclared in this block
#       previous declaration at ./listener.go:15:2
#    make: *** [Makefile:9: gen] Error 2

GC -> GC
	: 'gc' string_lit
;
