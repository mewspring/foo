language llvm(go);

lang = "llvm"
package = "github.com/mewspring/foo/none/mini"
eventBased = true
eventFields = true

:: lexer

_ascii_letter_upper = /[A-Z]/

_ascii_letter_lower = /[a-z]/

_ascii_letter = /{_ascii_letter_upper}|{_ascii_letter_lower}/

_letter = /{_ascii_letter}|[-$\._]/

_escape_letter = /{_letter}|[\\]/

_decimal_digit = /[0-9]/

_name = /{_letter}({_letter}|{_decimal_digit})*/

_escape_name = /{_escape_letter}({_escape_letter}|{_decimal_digit})*/

_quoted_name = /{_quoted_string}/

_id = /{_decimals}/

global_ident_tok : /{_global_name}|{_global_id}/

_global_name = /[@]({_name}|{_quoted_name})/

_global_id = /[@]{_id}/

local_ident_tok : /{_local_name}|{_local_id}/

_local_name = /[%]({_name}|{_quoted_name})/

_local_id = /[%]{_id}/

int_type_tok : /i[0-9]+/

_decimals = /{_decimal_digit}+/

int_lit_tok : /[-]?[0-9]+/

_quoted_string = /["][^"]*["]/

'(' : /[(]/
')' : /[)]/
'{' : /[{]/
'}' : /[}]/
',' : /[,]/
'=' : /[=]/
'add' : /add/
'call' : /call/
'define' : /define/
'double' : /double/
'float' : /float/
'fp128' : /fp128/
'half' : /half/
'ppc_fp128' : /ppc_fp128/
'ret' : /ret/
'void' : /void/
'x86_fp80' : /x86_fp80/

'foobar' : /foobar/
'foobaz' : /foobaz/

:: parser

%input Module;

GlobalIdent -> GlobalIdent
	: global_ident_tok
;

LocalIdent -> LocalIdent
	: local_ident_tok
;

IntLit -> IntLit
	: int_lit_tok
;

Module -> Module
	: TopLevelEntities=TopLevelEntity*
;

%interface TopLevelEntity;

TopLevelEntity -> TopLevelEntity
	: FuncDef
	| FooBaz
;

FooBaz -> FooBaz
	: 'foobaz'
;

FuncDef -> FuncDef
	: 'define' Header=FuncHeader Metadata=FuncMetadata Body=FuncBody
;

FuncHeader -> FuncHeader
	: RetType=Type Name=GlobalIdent '(' Params ')'
;

FuncBody -> FuncBody
	: '{' Blocks=BasicBlock+ '}'
;

%interface Type;

Type -> Type
	: VoidType
	| FirstClassType
;

%interface FirstClassType;

FirstClassType -> FirstClassType
	: ConcreteType
;

%interface ConcreteType;

ConcreteType -> ConcreteType
	: IntType
	| FloatType
;

VoidType -> VoidType
	: 'void'
;

IntType -> IntType
	: int_type_tok
;

FloatType -> FloatType
	: FloatKind
;

FloatKind -> FloatKind
	: 'half'
	| 'float'
	| 'double'
	| 'x86_fp80'
	| 'fp128'
	| 'ppc_fp128'
;

%interface Value;

Value -> Value
	: Constant
	| LocalIdent
;

%interface Constant;

Constant -> Constant
	: IntConst
	| GlobalIdent
;

IntConst -> IntConst
	: IntLit
;

BasicBlock -> BasicBlock
	: Insts=Instruction* Term=Terminator
;

%interface Instruction;

Instruction -> Instruction
	: LocalDefInst
	| ValueInstruction
;

LocalDefInst -> LocalDefInst
	: Name=LocalIdent '=' Inst=ValueInstruction
;

%interface ValueInstruction;

ValueInstruction -> ValueInstruction
	: AddInst
	| CallInst
;

AddInst -> AddInst
	: 'add' X=TypeValue ',' Y=Value InstMetadata
;

CallInst -> CallInst
	: 'call' Typ=Type Callee=Value '(' Args ')' OperandBundles InstMetadata
;

%interface Terminator;

Terminator -> Terminator
	: RetTerm
	| FooBar
;

FooBar -> FooBar
	: 'foobar'
;

RetTerm -> RetTerm
	: 'ret' XTyp=VoidType InstMetadata
;

Args -> Args
	: /* empty */
;

FuncMetadata -> FuncMetadata
	: /* empty */
;

InstMetadata -> InstMetadata
   : /* empty */
;

OperandBundles -> OperandBundles
	: /* empty */
;

Params -> Params
	: /* empty */
;

TypeValue -> TypeValue
	: Typ=FirstClassType Val=Value
;
