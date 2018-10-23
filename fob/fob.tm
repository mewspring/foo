language fob(go);

lang = "fob"
package = "github.com/mewspring/foo/fob"
eventBased = true
eventFields = true

:: lexer

'x86_fastcallcc' : /x86_fastcallcc/
'x86_stdcallcc' : /x86_stdcallcc/
'x86_thiscallcc' : /x86_thiscallcc/

:: parser

input : FuncHeader ;

FuncHeader -> FuncHeader
	: CallingConvopt
;

CallingConv -> CallingConv
	: 'x86_fastcallcc'
	| 'x86_stdcallcc'
	| 'x86_thiscallcc'
;
