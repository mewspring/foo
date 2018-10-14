language baz(go);

lang = "baz"
package = "github.com/mewspring/foo/baz"
eventBased = true
eventFields = true

:: lexer

_letter = /[-A-Za-z$\._]/

_escape_letter = /{_letter}|[\\]/

_hex_digit = /[A-Fa-f0-9]/

comment : /[;][^\r\n]*/               (space)
whitespace : /[\x00 \t\r\n]+/         (space)

_name = /{_letter}({_letter}|[0-9])*/

_escape_name = /{_escape_letter}({_escape_letter}|[0-9])*/

_quoted_name = /{_quoted_string}/

_id = /[0-9]+/

local_ident_tok : /{_local_name}|{_local_id}/

_local_name = /[%]({_name}|{_quoted_name})/

_local_id = /[%]{_id}/

int_lit_tok : /[-]?[0-9]+/

_quoted_string = /["][^"]*["]/

int_type_tok : /i[0-9]+/

'metadata' : /metadata/
'opaque' : /opaque/
'type' : /type/
'void' : /void/
'x' : /x/

'[' : /[\[]/
']' : /[\]]/
'=' : /[=]/

:: parser

input : Module;

Module -> Module
	: TopLevelEntities=TopLevelEntity*
;

%interface TopLevelEntity;

TopLevelEntity -> TopLevelEntity
	: TypeDef
;

TypeDef -> TypeDef
	: LocalIdent '=' 'type' (OpaqueType | Type)
;

LocalIdent -> LocalIdent
	: local_ident_tok
;

%interface Type;

Type -> Type
	: VoidType
	| FirstClassType
;

%interface FirstClassType;

FirstClassType -> FirstClassType
	: ConcreteType
	| MetadataType
;

%interface ConcreteType;

ConcreteType -> ConcreteType
	: IntType
	| ArrayType
;

VoidType -> VoidType
	: 'void'
;

IntType -> IntType
	: int_type_tok
;

MetadataType -> MetadataType
	: 'metadata'
;

ArrayType -> ArrayType
	: '[' UintLit 'x' Type ']'
;

OpaqueType -> OpaqueType
	: 'opaque'
;

UintLit -> UintLit
	: int_lit_tok
;
