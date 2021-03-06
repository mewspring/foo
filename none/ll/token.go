// generated by Textmapper; DO NOT EDIT

package ll

import (
	"fmt"
)

// Token is an enum of all terminal symbols of the llvm language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota - 1
	EOI

	INVALID_TOKEN
	COMMENT
	WHITESPACE
	GLOBAL_IDENT_TOK
	LOCAL_IDENT_TOK
	LABEL_IDENT_TOK
	ATTR_GROUP_ID_TOK
	COMDAT_NAME_TOK
	METADATA_NAME_TOK
	METADATA_ID_TOK
	DWARF_TAG_TOK
	DWARF_ATT_ENCODING_TOK
	DI_FLAG_TOK
	DWARF_LANG_TOK
	DWARF_CC_TOK
	CHECKSUM_KIND_TOK
	DWARF_VIRTUALITY_TOK
	DWARF_MACINFO_TOK
	DWARF_OP_TOK
	EMISSION_KIND_TOK
	NAME_TABLE_KIND_TOK
	INT_LIT_TOK
	FLOAT_LIT_TOK
	STRING_LIT_TOK
	INT_TYPE_TOK
	EXCLDIBASICTYPE                // !DIBasicType
	EXCLDICOMPILEUNIT              // !DICompileUnit
	EXCLDICOMPOSITETYPE            // !DICompositeType
	EXCLDIDERIVEDTYPE              // !DIDerivedType
	EXCLDIENUMERATOR               // !DIEnumerator
	EXCLDIEXPRESSION               // !DIExpression
	EXCLDIFILE                     // !DIFile
	EXCLDIGLOBALVARIABLE           // !DIGlobalVariable
	EXCLDIGLOBALVARIABLEEXPRESSION // !DIGlobalVariableExpression
	EXCLDIIMPORTEDENTITY           // !DIImportedEntity
	EXCLDILABEL                    // !DILabel
	EXCLDILEXICALBLOCK             // !DILexicalBlock
	EXCLDILEXICALBLOCKFILE         // !DILexicalBlockFile
	EXCLDILOCALVARIABLE            // !DILocalVariable
	EXCLDILOCATION                 // !DILocation
	EXCLDIMACRO                    // !DIMacro
	EXCLDIMACROFILE                // !DIMacroFile
	EXCLDIMODULE                   // !DIModule
	EXCLDINAMESPACE                // !DINamespace
	EXCLDIOBJCPROPERTY             // !DIObjCProperty
	EXCLDISUBPROGRAM               // !DISubprogram
	EXCLDISUBRANGE                 // !DISubrange
	EXCLDISUBROUTINETYPE           // !DISubroutineType
	EXCLDITEMPLATETYPEPARAMETER    // !DITemplateTypeParameter
	EXCLDITEMPLATEVALUEPARAMETER   // !DITemplateValueParameter
	EXCLGENERICDINODE              // !GenericDINode
	AARCH64_VECTOR_PCS             // aarch64_vector_pcs
	ACQ_REL                        // acq_rel
	ACQUIRE                        // acquire
	ADD                            // add
	ADDRSPACE                      // addrspace
	ADDRSPACECAST                  // addrspacecast
	AFN                            // afn
	ALIAS                          // alias
	ALIGNCOLON                     // align:
	ALIGN                          // align
	ALIGNSTACK                     // alignstack
	ALLOCA                         // alloca
	ALLOCSIZE                      // allocsize
	ALWAYSINLINE                   // alwaysinline
	AMDGPU_CS                      // amdgpu_cs
	AMDGPU_ES                      // amdgpu_es
	AMDGPU_GS                      // amdgpu_gs
	AMDGPU_HS                      // amdgpu_hs
	AMDGPU_KERNEL                  // amdgpu_kernel
	AMDGPU_LS                      // amdgpu_ls
	AMDGPU_PS                      // amdgpu_ps
	AMDGPU_VS                      // amdgpu_vs
	AND                            // and
	ANY                            // any
	ANYREGCC                       // anyregcc
	APPENDING                      // appending
	ARCP                           // arcp
	ARGCOLON                       // arg:
	ARGMEMONLY                     // argmemonly
	ARM_AAPCS_VFPCC                // arm_aapcs_vfpcc
	ARM_AAPCSCC                    // arm_aapcscc
	ARM_APCSCC                     // arm_apcscc
	ASHR                           // ashr
	ASM                            // asm
	ATOMIC                         // atomic
	ATOMICRMW                      // atomicrmw
	ATTRIBUTESCOLON                // attributes:
	ATTRIBUTES                     // attributes
	AVAILABLE_EXTERNALLY           // available_externally
	AVR_INTRCC                     // avr_intrcc
	AVR_SIGNALCC                   // avr_signalcc
	BASETYPECOLON                  // baseType:
	BITCAST                        // bitcast
	BLOCKADDRESS                   // blockaddress
	BR                             // br
	BUILTIN                        // builtin
	BYVAL                          // byval
	CHAR_C                         // c
	CALL                           // call
	CALLER                         // caller
	CATCH                          // catch
	CATCHPAD                       // catchpad
	CATCHRET                       // catchret
	CATCHSWITCH                    // catchswitch
	CCCOLON                        // cc:
	CC                             // cc
	CCC                            // ccc
	CHECKSUMCOLON                  // checksum:
	CHECKSUMKINDCOLON              // checksumkind:
	CLEANUP                        // cleanup
	CLEANUPPAD                     // cleanuppad
	CLEANUPRET                     // cleanupret
	CMPXCHG                        // cmpxchg
	COLD                           // cold
	COLDCC                         // coldcc
	COLUMNCOLON                    // column:
	COMDAT                         // comdat
	COMMON                         // common
	CONFIGMACROSCOLON              // configMacros:
	CONSTANT                       // constant
	CONTAININGTYPECOLON            // containingType:
	CONTRACT                       // contract
	CONVERGENT                     // convergent
	COUNTCOLON                     // count:
	CXX_FAST_TLSCC                 // cxx_fast_tlscc
	DATALAYOUT                     // datalayout
	DEBUGINFOFORPROFILINGCOLON     // debugInfoForProfiling:
	DECLARATIONCOLON               // declaration:
	DECLARE                        // declare
	DEFAULT                        // default
	DEFINE                         // define
	DEREFERENCEABLE_OR_NULL        // dereferenceable_or_null
	DEREFERENCEABLE                // dereferenceable
	DIRECTORYCOLON                 // directory:
	DISCRIMINATORCOLON             // discriminator:
	DISTINCT                       // distinct
	DLLEXPORT                      // dllexport
	DLLIMPORT                      // dllimport
	DOUBLE                         // double
	DSO_LOCAL                      // dso_local
	DSO_PREEMPTABLE                // dso_preemptable
	DWARFADDRESSSPACECOLON         // dwarfAddressSpace:
	DWOIDCOLON                     // dwoId:
	ELEMENTSCOLON                  // elements:
	EMISSIONKINDCOLON              // emissionKind:
	ENCODINGCOLON                  // encoding:
	ENTITYCOLON                    // entity:
	ENUMSCOLON                     // enums:
	EQ                             // eq
	EXACT                          // exact
	EXACTMATCH                     // exactmatch
	EXPORTSYMBOLSCOLON             // exportSymbols:
	EXPRCOLON                      // expr:
	EXTERN_WEAK                    // extern_weak
	EXTERNAL                       // external
	EXTERNALLY_INITIALIZED         // externally_initialized
	EXTRACTELEMENT                 // extractelement
	EXTRACTVALUE                   // extractvalue
	EXTRADATACOLON                 // extraData:
	FADD                           // fadd
	FALSE                          // false
	FAST                           // fast
	FASTCC                         // fastcc
	FCMP                           // fcmp
	FDIV                           // fdiv
	FENCE                          // fence
	FILECOLON                      // file:
	FILENAMECOLON                  // filename:
	FILTER                         // filter
	FLAGSCOLON                     // flags:
	FLOAT                          // float
	FMUL                           // fmul
	FP128                          // fp128
	FPEXT                          // fpext
	FPTOSI                         // fptosi
	FPTOUI                         // fptoui
	FPTRUNC                        // fptrunc
	FREM                           // frem
	FROM                           // from
	FSUB                           // fsub
	GC                             // gc
	GETELEMENTPTR                  // getelementptr
	GETTERCOLON                    // getter:
	GHCCC                          // ghccc
	GLOBAL                         // global
	GLOBALSCOLON                   // globals:
	HALF                           // half
	HEADERCOLON                    // header:
	HHVM_CCC                       // hhvm_ccc
	HHVMCC                         // hhvmcc
	HIDDEN                         // hidden
	ICMP                           // icmp
	IDENTIFIERCOLON                // identifier:
	IFUNC                          // ifunc
	IMPORTSCOLON                   // imports:
	INACCESSIBLEMEM_OR_ARGMEMONLY  // inaccessiblemem_or_argmemonly
	INACCESSIBLEMEMONLY            // inaccessiblememonly
	INALLOCA                       // inalloca
	INBOUNDS                       // inbounds
	INCLUDEPATHCOLON               // includePath:
	INDIRECTBR                     // indirectbr
	INITIALEXEC                    // initialexec
	INLINEDATCOLON                 // inlinedAt:
	INLINEHINT                     // inlinehint
	INRANGE                        // inrange
	INREG                          // inreg
	INSERTELEMENT                  // insertelement
	INSERTVALUE                    // insertvalue
	INTEL_OCL_BICC                 // intel_ocl_bicc
	INTELDIALECT                   // inteldialect
	INTERNAL                       // internal
	INTTOPTR                       // inttoptr
	INVOKE                         // invoke
	ISDEFINITIONCOLON              // isDefinition:
	ISIMPLICITCODECOLON            // isImplicitCode:
	ISLOCALCOLON                   // isLocal:
	ISOPTIMIZEDCOLON               // isOptimized:
	ISUNSIGNEDCOLON                // isUnsigned:
	ISYSROOTCOLON                  // isysroot:
	JUMPTABLE                      // jumptable
	LABEL                          // label
	LANDINGPAD                     // landingpad
	LANGUAGECOLON                  // language:
	LARGEST                        // largest
	LINECOLON                      // line:
	LINKAGENAMECOLON               // linkageName:
	LINKONCE_ODR                   // linkonce_odr
	LINKONCE                       // linkonce
	LOAD                           // load
	LOCAL_UNNAMED_ADDR             // local_unnamed_addr
	LOCALDYNAMIC                   // localdynamic
	LOCALEXEC                      // localexec
	LOWERBOUNDCOLON                // lowerBound:
	LSHR                           // lshr
	MACROSCOLON                    // macros:
	MAX                            // max
	METADATA                       // metadata
	MIN                            // min
	MINSIZE                        // minsize
	MODULE                         // module
	MONOTONIC                      // monotonic
	MSP430_INTRCC                  // msp430_intrcc
	MUL                            // mul
	MUSTTAIL                       // musttail
	NAKED                          // naked
	NAMECOLON                      // name:
	NAMETABLEKINDCOLON             // nameTableKind:
	NAND                           // nand
	NE                             // ne
	NEST                           // nest
	NINF                           // ninf
	NNAN                           // nnan
	NOALIAS                        // noalias
	NOBUILTIN                      // nobuiltin
	NOCAPTURE                      // nocapture
	NOCF_CHECK                     // nocf_check
	NODESCOLON                     // nodes:
	NODUPLICATE                    // noduplicate
	NODUPLICATES                   // noduplicates
	NOIMPLICITFLOAT                // noimplicitfloat
	NOINLINE                       // noinline
	NONE                           // none
	NONLAZYBIND                    // nonlazybind
	NONNULL                        // nonnull
	NORECURSE                      // norecurse
	NOREDZONE                      // noredzone
	NORETURN                       // noreturn
	NOTAIL                         // notail
	NOUNWIND                       // nounwind
	NSW                            // nsw
	NSZ                            // nsz
	NULL                           // null
	NUW                            // nuw
	OEQ                            // oeq
	OFFSETCOLON                    // offset:
	OGE                            // oge
	OGT                            // ogt
	OLE                            // ole
	OLT                            // olt
	ONE                            // one
	OPAQUE                         // opaque
	OPERANDSCOLON                  // operands:
	OPTFORFUZZING                  // optforfuzzing
	OPTNONE                        // optnone
	OPTSIZE                        // optsize
	OR                             // or
	ORD                            // ord
	PERSONALITY                    // personality
	PHI                            // phi
	PPC_FP128                      // ppc_fp128
	PREFIX                         // prefix
	PRESERVE_ALLCC                 // preserve_allcc
	PRESERVE_MOSTCC                // preserve_mostcc
	PRIVATE                        // private
	PRODUCERCOLON                  // producer:
	PROLOGUE                       // prologue
	PROTECTED                      // protected
	PTRTOINT                       // ptrtoint
	PTX_DEVICE                     // ptx_device
	PTX_KERNEL                     // ptx_kernel
	READNONE                       // readnone
	READONLY                       // readonly
	REASSOC                        // reassoc
	RELEASE                        // release
	RESUME                         // resume
	RET                            // ret
	RETAINEDNODESCOLON             // retainedNodes:
	RETAINEDTYPESCOLON             // retainedTypes:
	RETURNED                       // returned
	RETURNS_TWICE                  // returns_twice
	RUNTIMELANGCOLON               // runtimeLang:
	RUNTIMEVERSIONCOLON            // runtimeVersion:
	SAFESTACK                      // safestack
	SAMESIZE                       // samesize
	SANITIZE_ADDRESS               // sanitize_address
	SANITIZE_HWADDRESS             // sanitize_hwaddress
	SANITIZE_MEMORY                // sanitize_memory
	SANITIZE_THREAD                // sanitize_thread
	SCOPECOLON                     // scope:
	SCOPELINECOLON                 // scopeLine:
	SDIV                           // sdiv
	SECTION                        // section
	SELECT                         // select
	SEQ_CST                        // seq_cst
	SETTERCOLON                    // setter:
	SEXT                           // sext
	SGE                            // sge
	SGT                            // sgt
	SHADOWCALLSTACK                // shadowcallstack
	SHL                            // shl
	SHUFFLEVECTOR                  // shufflevector
	SIDEEFFECT                     // sideeffect
	SIGNEXT                        // signext
	SINGLETHREAD                   // singlethread
	SITOFP                         // sitofp
	SIZECOLON                      // size:
	SLE                            // sle
	SLT                            // slt
	SOURCE_FILENAME                // source_filename
	SOURCECOLON                    // source:
	SPECULATABLE                   // speculatable
	SPECULATIVE_LOAD_HARDENING     // speculative_load_hardening
	SPIR_FUNC                      // spir_func
	SPIR_KERNEL                    // spir_kernel
	SPLITDEBUGFILENAMECOLON        // splitDebugFilename:
	SPLITDEBUGINLININGCOLON        // splitDebugInlining:
	SREM                           // srem
	SRET                           // sret
	SSP                            // ssp
	SSPREQ                         // sspreq
	SSPSTRONG                      // sspstrong
	STORE                          // store
	STRICTFP                       // strictfp
	SUB                            // sub
	SWIFTCC                        // swiftcc
	SWIFTERROR                     // swifterror
	SWIFTSELF                      // swiftself
	SWITCH                         // switch
	SYNCSCOPE                      // syncscope
	TAGCOLON                       // tag:
	TAIL                           // tail
	TARGET                         // target
	TEMPLATEPARAMSCOLON            // templateParams:
	THISADJUSTMENTCOLON            // thisAdjustment:
	THREAD_LOCAL                   // thread_local
	THROWNTYPESCOLON               // thrownTypes:
	TO                             // to
	TOKEN                          // token
	TRIPLE                         // triple
	TRUE                           // true
	TRUNC                          // trunc
	TYPECOLON                      // type:
	TYPE                           // type
	TYPESCOLON                     // types:
	UDIV                           // udiv
	UEQ                            // ueq
	UGE                            // uge
	UGT                            // ugt
	UITOFP                         // uitofp
	ULE                            // ule
	ULT                            // ult
	UMAX                           // umax
	UMIN                           // umin
	UNDEF                          // undef
	UNE                            // une
	UNITCOLON                      // unit:
	UNNAMED_ADDR                   // unnamed_addr
	UNO                            // uno
	UNORDERED                      // unordered
	UNREACHABLE                    // unreachable
	UNWIND                         // unwind
	UREM                           // urem
	USELISTORDER_BB                // uselistorder_bb
	USELISTORDER                   // uselistorder
	UWTABLE                        // uwtable
	VA_ARG                         // va_arg
	VALUECOLON                     // value:
	VARCOLON                       // var:
	VIRTUALINDEXCOLON              // virtualIndex:
	VIRTUALITYCOLON                // virtuality:
	VOID                           // void
	VOLATILE                       // volatile
	VTABLEHOLDERCOLON              // vtableHolder:
	WEAK_ODR                       // weak_odr
	WEAK                           // weak
	WEBKIT_JSCC                    // webkit_jscc
	WIN64CC                        // win64cc
	WITHIN                         // within
	WRITEONLY                      // writeonly
	CHAR_X                         // x
	X86_64_SYSVCC                  // x86_64_sysvcc
	X86_FASTCALLCC                 // x86_fastcallcc
	X86_FP80                       // x86_fp80
	X86_INTRCC                     // x86_intrcc
	X86_MMX                        // x86_mmx
	X86_REGCALLCC                  // x86_regcallcc
	X86_STDCALLCC                  // x86_stdcallcc
	X86_THISCALLCC                 // x86_thiscallcc
	X86_VECTORCALLCC               // x86_vectorcallcc
	XCHG                           // xchg
	XOR                            // xor
	ZEROEXT                        // zeroext
	ZEROINITIALIZER                // zeroinitializer
	ZEXT                           // zext
	COMMA
	EXCL
	DOTDOTDOT // ...
	LPAREN
	RPAREN
	LBRACK
	RBRACK
	LBRACE
	RBRACE
	MULT
	LT
	ASSIGN
	GT
	OR1

	NumTokens
)

var tokenStr = [...]string{
	"EOI",

	"INVALID_TOKEN",
	"COMMENT",
	"WHITESPACE",
	"GLOBAL_IDENT_TOK",
	"LOCAL_IDENT_TOK",
	"LABEL_IDENT_TOK",
	"ATTR_GROUP_ID_TOK",
	"COMDAT_NAME_TOK",
	"METADATA_NAME_TOK",
	"METADATA_ID_TOK",
	"DWARF_TAG_TOK",
	"DWARF_ATT_ENCODING_TOK",
	"DI_FLAG_TOK",
	"DWARF_LANG_TOK",
	"DWARF_CC_TOK",
	"CHECKSUM_KIND_TOK",
	"DWARF_VIRTUALITY_TOK",
	"DWARF_MACINFO_TOK",
	"DWARF_OP_TOK",
	"EMISSION_KIND_TOK",
	"NAME_TABLE_KIND_TOK",
	"INT_LIT_TOK",
	"FLOAT_LIT_TOK",
	"STRING_LIT_TOK",
	"INT_TYPE_TOK",
	"!DIBasicType",
	"!DICompileUnit",
	"!DICompositeType",
	"!DIDerivedType",
	"!DIEnumerator",
	"!DIExpression",
	"!DIFile",
	"!DIGlobalVariable",
	"!DIGlobalVariableExpression",
	"!DIImportedEntity",
	"!DILabel",
	"!DILexicalBlock",
	"!DILexicalBlockFile",
	"!DILocalVariable",
	"!DILocation",
	"!DIMacro",
	"!DIMacroFile",
	"!DIModule",
	"!DINamespace",
	"!DIObjCProperty",
	"!DISubprogram",
	"!DISubrange",
	"!DISubroutineType",
	"!DITemplateTypeParameter",
	"!DITemplateValueParameter",
	"!GenericDINode",
	"aarch64_vector_pcs",
	"acq_rel",
	"acquire",
	"add",
	"addrspace",
	"addrspacecast",
	"afn",
	"alias",
	"align:",
	"align",
	"alignstack",
	"alloca",
	"allocsize",
	"alwaysinline",
	"amdgpu_cs",
	"amdgpu_es",
	"amdgpu_gs",
	"amdgpu_hs",
	"amdgpu_kernel",
	"amdgpu_ls",
	"amdgpu_ps",
	"amdgpu_vs",
	"and",
	"any",
	"anyregcc",
	"appending",
	"arcp",
	"arg:",
	"argmemonly",
	"arm_aapcs_vfpcc",
	"arm_aapcscc",
	"arm_apcscc",
	"ashr",
	"asm",
	"atomic",
	"atomicrmw",
	"attributes:",
	"attributes",
	"available_externally",
	"avr_intrcc",
	"avr_signalcc",
	"baseType:",
	"bitcast",
	"blockaddress",
	"br",
	"builtin",
	"byval",
	"c",
	"call",
	"caller",
	"catch",
	"catchpad",
	"catchret",
	"catchswitch",
	"cc:",
	"cc",
	"ccc",
	"checksum:",
	"checksumkind:",
	"cleanup",
	"cleanuppad",
	"cleanupret",
	"cmpxchg",
	"cold",
	"coldcc",
	"column:",
	"comdat",
	"common",
	"configMacros:",
	"constant",
	"containingType:",
	"contract",
	"convergent",
	"count:",
	"cxx_fast_tlscc",
	"datalayout",
	"debugInfoForProfiling:",
	"declaration:",
	"declare",
	"default",
	"define",
	"dereferenceable_or_null",
	"dereferenceable",
	"directory:",
	"discriminator:",
	"distinct",
	"dllexport",
	"dllimport",
	"double",
	"dso_local",
	"dso_preemptable",
	"dwarfAddressSpace:",
	"dwoId:",
	"elements:",
	"emissionKind:",
	"encoding:",
	"entity:",
	"enums:",
	"eq",
	"exact",
	"exactmatch",
	"exportSymbols:",
	"expr:",
	"extern_weak",
	"external",
	"externally_initialized",
	"extractelement",
	"extractvalue",
	"extraData:",
	"fadd",
	"false",
	"fast",
	"fastcc",
	"fcmp",
	"fdiv",
	"fence",
	"file:",
	"filename:",
	"filter",
	"flags:",
	"float",
	"fmul",
	"fp128",
	"fpext",
	"fptosi",
	"fptoui",
	"fptrunc",
	"frem",
	"from",
	"fsub",
	"gc",
	"getelementptr",
	"getter:",
	"ghccc",
	"global",
	"globals:",
	"half",
	"header:",
	"hhvm_ccc",
	"hhvmcc",
	"hidden",
	"icmp",
	"identifier:",
	"ifunc",
	"imports:",
	"inaccessiblemem_or_argmemonly",
	"inaccessiblememonly",
	"inalloca",
	"inbounds",
	"includePath:",
	"indirectbr",
	"initialexec",
	"inlinedAt:",
	"inlinehint",
	"inrange",
	"inreg",
	"insertelement",
	"insertvalue",
	"intel_ocl_bicc",
	"inteldialect",
	"internal",
	"inttoptr",
	"invoke",
	"isDefinition:",
	"isImplicitCode:",
	"isLocal:",
	"isOptimized:",
	"isUnsigned:",
	"isysroot:",
	"jumptable",
	"label",
	"landingpad",
	"language:",
	"largest",
	"line:",
	"linkageName:",
	"linkonce_odr",
	"linkonce",
	"load",
	"local_unnamed_addr",
	"localdynamic",
	"localexec",
	"lowerBound:",
	"lshr",
	"macros:",
	"max",
	"metadata",
	"min",
	"minsize",
	"module",
	"monotonic",
	"msp430_intrcc",
	"mul",
	"musttail",
	"naked",
	"name:",
	"nameTableKind:",
	"nand",
	"ne",
	"nest",
	"ninf",
	"nnan",
	"noalias",
	"nobuiltin",
	"nocapture",
	"nocf_check",
	"nodes:",
	"noduplicate",
	"noduplicates",
	"noimplicitfloat",
	"noinline",
	"none",
	"nonlazybind",
	"nonnull",
	"norecurse",
	"noredzone",
	"noreturn",
	"notail",
	"nounwind",
	"nsw",
	"nsz",
	"null",
	"nuw",
	"oeq",
	"offset:",
	"oge",
	"ogt",
	"ole",
	"olt",
	"one",
	"opaque",
	"operands:",
	"optforfuzzing",
	"optnone",
	"optsize",
	"or",
	"ord",
	"personality",
	"phi",
	"ppc_fp128",
	"prefix",
	"preserve_allcc",
	"preserve_mostcc",
	"private",
	"producer:",
	"prologue",
	"protected",
	"ptrtoint",
	"ptx_device",
	"ptx_kernel",
	"readnone",
	"readonly",
	"reassoc",
	"release",
	"resume",
	"ret",
	"retainedNodes:",
	"retainedTypes:",
	"returned",
	"returns_twice",
	"runtimeLang:",
	"runtimeVersion:",
	"safestack",
	"samesize",
	"sanitize_address",
	"sanitize_hwaddress",
	"sanitize_memory",
	"sanitize_thread",
	"scope:",
	"scopeLine:",
	"sdiv",
	"section",
	"select",
	"seq_cst",
	"setter:",
	"sext",
	"sge",
	"sgt",
	"shadowcallstack",
	"shl",
	"shufflevector",
	"sideeffect",
	"signext",
	"singlethread",
	"sitofp",
	"size:",
	"sle",
	"slt",
	"source_filename",
	"source:",
	"speculatable",
	"speculative_load_hardening",
	"spir_func",
	"spir_kernel",
	"splitDebugFilename:",
	"splitDebugInlining:",
	"srem",
	"sret",
	"ssp",
	"sspreq",
	"sspstrong",
	"store",
	"strictfp",
	"sub",
	"swiftcc",
	"swifterror",
	"swiftself",
	"switch",
	"syncscope",
	"tag:",
	"tail",
	"target",
	"templateParams:",
	"thisAdjustment:",
	"thread_local",
	"thrownTypes:",
	"to",
	"token",
	"triple",
	"true",
	"trunc",
	"type:",
	"type",
	"types:",
	"udiv",
	"ueq",
	"uge",
	"ugt",
	"uitofp",
	"ule",
	"ult",
	"umax",
	"umin",
	"undef",
	"une",
	"unit:",
	"unnamed_addr",
	"uno",
	"unordered",
	"unreachable",
	"unwind",
	"urem",
	"uselistorder_bb",
	"uselistorder",
	"uwtable",
	"va_arg",
	"value:",
	"var:",
	"virtualIndex:",
	"virtuality:",
	"void",
	"volatile",
	"vtableHolder:",
	"weak_odr",
	"weak",
	"webkit_jscc",
	"win64cc",
	"within",
	"writeonly",
	"x",
	"x86_64_sysvcc",
	"x86_fastcallcc",
	"x86_fp80",
	"x86_intrcc",
	"x86_mmx",
	"x86_regcallcc",
	"x86_stdcallcc",
	"x86_thiscallcc",
	"x86_vectorcallcc",
	"xchg",
	"xor",
	"zeroext",
	"zeroinitializer",
	"zext",
	"COMMA",
	"EXCL",
	"...",
	"LPAREN",
	"RPAREN",
	"LBRACK",
	"RBRACK",
	"LBRACE",
	"RBRACE",
	"MULT",
	"LT",
	"ASSIGN",
	"GT",
	"OR1",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}
