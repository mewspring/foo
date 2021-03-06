// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/mewspring/foo/none/ll"
)

type Selector func(nt ll.NodeType) bool

var (
	Any                             = func(t ll.NodeType) bool { return true }
	AShrExpr                        = func(t ll.NodeType) bool { return t == ll.AShrExpr }
	AShrInst                        = func(t ll.NodeType) bool { return t == ll.AShrInst }
	AddExpr                         = func(t ll.NodeType) bool { return t == ll.AddExpr }
	AddInst                         = func(t ll.NodeType) bool { return t == ll.AddInst }
	AddrSpace                       = func(t ll.NodeType) bool { return t == ll.AddrSpace }
	AddrSpaceCastExpr               = func(t ll.NodeType) bool { return t == ll.AddrSpaceCastExpr }
	AddrSpaceCastInst               = func(t ll.NodeType) bool { return t == ll.AddrSpaceCastInst }
	AliasDef                        = func(t ll.NodeType) bool { return t == ll.AliasDef }
	AlignField                      = func(t ll.NodeType) bool { return t == ll.AlignField }
	AlignPair                       = func(t ll.NodeType) bool { return t == ll.AlignPair }
	AlignStack                      = func(t ll.NodeType) bool { return t == ll.AlignStack }
	AlignStackPair                  = func(t ll.NodeType) bool { return t == ll.AlignStackPair }
	Alignment                       = func(t ll.NodeType) bool { return t == ll.Alignment }
	AllocSize                       = func(t ll.NodeType) bool { return t == ll.AllocSize }
	AllocaInst                      = func(t ll.NodeType) bool { return t == ll.AllocaInst }
	AndExpr                         = func(t ll.NodeType) bool { return t == ll.AndExpr }
	AndInst                         = func(t ll.NodeType) bool { return t == ll.AndInst }
	Arg                             = func(t ll.NodeType) bool { return t == ll.Arg }
	ArgField                        = func(t ll.NodeType) bool { return t == ll.ArgField }
	Args                            = func(t ll.NodeType) bool { return t == ll.Args }
	ArrayConst                      = func(t ll.NodeType) bool { return t == ll.ArrayConst }
	ArrayType                       = func(t ll.NodeType) bool { return t == ll.ArrayType }
	Atomic                          = func(t ll.NodeType) bool { return t == ll.Atomic }
	AtomicOp                        = func(t ll.NodeType) bool { return t == ll.AtomicOp }
	AtomicOrdering                  = func(t ll.NodeType) bool { return t == ll.AtomicOrdering }
	AtomicRMWInst                   = func(t ll.NodeType) bool { return t == ll.AtomicRMWInst }
	AttrGroupDef                    = func(t ll.NodeType) bool { return t == ll.AttrGroupDef }
	AttrGroupID                     = func(t ll.NodeType) bool { return t == ll.AttrGroupID }
	AttrPair                        = func(t ll.NodeType) bool { return t == ll.AttrPair }
	AttrString                      = func(t ll.NodeType) bool { return t == ll.AttrString }
	AttributesField                 = func(t ll.NodeType) bool { return t == ll.AttributesField }
	BaseTypeField                   = func(t ll.NodeType) bool { return t == ll.BaseTypeField }
	BasicBlock                      = func(t ll.NodeType) bool { return t == ll.BasicBlock }
	BitCastExpr                     = func(t ll.NodeType) bool { return t == ll.BitCastExpr }
	BitCastInst                     = func(t ll.NodeType) bool { return t == ll.BitCastInst }
	BlockAddressConst               = func(t ll.NodeType) bool { return t == ll.BlockAddressConst }
	BoolConst                       = func(t ll.NodeType) bool { return t == ll.BoolConst }
	BoolLit                         = func(t ll.NodeType) bool { return t == ll.BoolLit }
	BrTerm                          = func(t ll.NodeType) bool { return t == ll.BrTerm }
	CCField                         = func(t ll.NodeType) bool { return t == ll.CCField }
	CallInst                        = func(t ll.NodeType) bool { return t == ll.CallInst }
	CallingConv                     = func(t ll.NodeType) bool { return t == ll.CallingConv }
	Case                            = func(t ll.NodeType) bool { return t == ll.Case }
	CatchClause                     = func(t ll.NodeType) bool { return t == ll.CatchClause }
	CatchPadInst                    = func(t ll.NodeType) bool { return t == ll.CatchPadInst }
	CatchRetTerm                    = func(t ll.NodeType) bool { return t == ll.CatchRetTerm }
	CatchSwitchTerm                 = func(t ll.NodeType) bool { return t == ll.CatchSwitchTerm }
	CharArrayConst                  = func(t ll.NodeType) bool { return t == ll.CharArrayConst }
	ChecksumField                   = func(t ll.NodeType) bool { return t == ll.ChecksumField }
	ChecksumKind                    = func(t ll.NodeType) bool { return t == ll.ChecksumKind }
	ChecksumkindField               = func(t ll.NodeType) bool { return t == ll.ChecksumkindField }
	Cleanup                         = func(t ll.NodeType) bool { return t == ll.Cleanup }
	CleanupPadInst                  = func(t ll.NodeType) bool { return t == ll.CleanupPadInst }
	CleanupRetTerm                  = func(t ll.NodeType) bool { return t == ll.CleanupRetTerm }
	CmpXchgInst                     = func(t ll.NodeType) bool { return t == ll.CmpXchgInst }
	ColumnField                     = func(t ll.NodeType) bool { return t == ll.ColumnField }
	Comdat                          = func(t ll.NodeType) bool { return t == ll.Comdat }
	ComdatDef                       = func(t ll.NodeType) bool { return t == ll.ComdatDef }
	ComdatName                      = func(t ll.NodeType) bool { return t == ll.ComdatName }
	CondBrTerm                      = func(t ll.NodeType) bool { return t == ll.CondBrTerm }
	ConfigMacrosField               = func(t ll.NodeType) bool { return t == ll.ConfigMacrosField }
	ContainingTypeField             = func(t ll.NodeType) bool { return t == ll.ContainingTypeField }
	CountField                      = func(t ll.NodeType) bool { return t == ll.CountField }
	DIBasicType                     = func(t ll.NodeType) bool { return t == ll.DIBasicType }
	DICompileUnit                   = func(t ll.NodeType) bool { return t == ll.DICompileUnit }
	DICompositeType                 = func(t ll.NodeType) bool { return t == ll.DICompositeType }
	DIDerivedType                   = func(t ll.NodeType) bool { return t == ll.DIDerivedType }
	DIEnumerator                    = func(t ll.NodeType) bool { return t == ll.DIEnumerator }
	DIExpression                    = func(t ll.NodeType) bool { return t == ll.DIExpression }
	DIFile                          = func(t ll.NodeType) bool { return t == ll.DIFile }
	DIFlag                          = func(t ll.NodeType) bool { return t == ll.DIFlag }
	DIFlags                         = func(t ll.NodeType) bool { return t == ll.DIFlags }
	DIGlobalVariable                = func(t ll.NodeType) bool { return t == ll.DIGlobalVariable }
	DIGlobalVariableExpression      = func(t ll.NodeType) bool { return t == ll.DIGlobalVariableExpression }
	DIImportedEntity                = func(t ll.NodeType) bool { return t == ll.DIImportedEntity }
	DILabel                         = func(t ll.NodeType) bool { return t == ll.DILabel }
	DILexicalBlock                  = func(t ll.NodeType) bool { return t == ll.DILexicalBlock }
	DILexicalBlockFile              = func(t ll.NodeType) bool { return t == ll.DILexicalBlockFile }
	DILocalVariable                 = func(t ll.NodeType) bool { return t == ll.DILocalVariable }
	DILocation                      = func(t ll.NodeType) bool { return t == ll.DILocation }
	DIMacro                         = func(t ll.NodeType) bool { return t == ll.DIMacro }
	DIMacroFile                     = func(t ll.NodeType) bool { return t == ll.DIMacroFile }
	DIModule                        = func(t ll.NodeType) bool { return t == ll.DIModule }
	DINamespace                     = func(t ll.NodeType) bool { return t == ll.DINamespace }
	DIObjCProperty                  = func(t ll.NodeType) bool { return t == ll.DIObjCProperty }
	DISubprogram                    = func(t ll.NodeType) bool { return t == ll.DISubprogram }
	DISubrange                      = func(t ll.NodeType) bool { return t == ll.DISubrange }
	DISubroutineType                = func(t ll.NodeType) bool { return t == ll.DISubroutineType }
	DITemplateTypeParameter         = func(t ll.NodeType) bool { return t == ll.DITemplateTypeParameter }
	DITemplateValueParameter        = func(t ll.NodeType) bool { return t == ll.DITemplateValueParameter }
	DLLStorageClass                 = func(t ll.NodeType) bool { return t == ll.DLLStorageClass }
	DebugInfoForProfilingField      = func(t ll.NodeType) bool { return t == ll.DebugInfoForProfilingField }
	DeclarationField                = func(t ll.NodeType) bool { return t == ll.DeclarationField }
	Dereferenceable                 = func(t ll.NodeType) bool { return t == ll.Dereferenceable }
	DirectoryField                  = func(t ll.NodeType) bool { return t == ll.DirectoryField }
	DiscriminatorField              = func(t ll.NodeType) bool { return t == ll.DiscriminatorField }
	DiscriminatorIntField           = func(t ll.NodeType) bool { return t == ll.DiscriminatorIntField }
	Distinct                        = func(t ll.NodeType) bool { return t == ll.Distinct }
	DwarfAddressSpaceField          = func(t ll.NodeType) bool { return t == ll.DwarfAddressSpaceField }
	DwarfAttEncoding                = func(t ll.NodeType) bool { return t == ll.DwarfAttEncoding }
	DwarfCC                         = func(t ll.NodeType) bool { return t == ll.DwarfCC }
	DwarfLang                       = func(t ll.NodeType) bool { return t == ll.DwarfLang }
	DwarfMacinfo                    = func(t ll.NodeType) bool { return t == ll.DwarfMacinfo }
	DwarfOp                         = func(t ll.NodeType) bool { return t == ll.DwarfOp }
	DwarfTag                        = func(t ll.NodeType) bool { return t == ll.DwarfTag }
	DwarfVirtuality                 = func(t ll.NodeType) bool { return t == ll.DwarfVirtuality }
	DwoIdField                      = func(t ll.NodeType) bool { return t == ll.DwoIdField }
	ElementsField                   = func(t ll.NodeType) bool { return t == ll.ElementsField }
	Ellipsis                        = func(t ll.NodeType) bool { return t == ll.Ellipsis }
	EmissionKind                    = func(t ll.NodeType) bool { return t == ll.EmissionKind }
	EmissionKindField               = func(t ll.NodeType) bool { return t == ll.EmissionKindField }
	EncodingField                   = func(t ll.NodeType) bool { return t == ll.EncodingField }
	EntityField                     = func(t ll.NodeType) bool { return t == ll.EntityField }
	EnumsField                      = func(t ll.NodeType) bool { return t == ll.EnumsField }
	Exact                           = func(t ll.NodeType) bool { return t == ll.Exact }
	ExceptionArg                    = func(t ll.NodeType) bool { return t == ll.ExceptionArg }
	ExceptionScope                  = func(t ll.NodeType) bool { return t == ll.ExceptionScope }
	ExportSymbolsField              = func(t ll.NodeType) bool { return t == ll.ExportSymbolsField }
	ExprField                       = func(t ll.NodeType) bool { return t == ll.ExprField }
	ExternLinkage                   = func(t ll.NodeType) bool { return t == ll.ExternLinkage }
	ExternallyInitialized           = func(t ll.NodeType) bool { return t == ll.ExternallyInitialized }
	ExtraDataField                  = func(t ll.NodeType) bool { return t == ll.ExtraDataField }
	ExtractElementExpr              = func(t ll.NodeType) bool { return t == ll.ExtractElementExpr }
	ExtractElementInst              = func(t ll.NodeType) bool { return t == ll.ExtractElementInst }
	ExtractValueExpr                = func(t ll.NodeType) bool { return t == ll.ExtractValueExpr }
	ExtractValueInst                = func(t ll.NodeType) bool { return t == ll.ExtractValueInst }
	FAddExpr                        = func(t ll.NodeType) bool { return t == ll.FAddExpr }
	FAddInst                        = func(t ll.NodeType) bool { return t == ll.FAddInst }
	FCmpExpr                        = func(t ll.NodeType) bool { return t == ll.FCmpExpr }
	FCmpInst                        = func(t ll.NodeType) bool { return t == ll.FCmpInst }
	FDivExpr                        = func(t ll.NodeType) bool { return t == ll.FDivExpr }
	FDivInst                        = func(t ll.NodeType) bool { return t == ll.FDivInst }
	FMulExpr                        = func(t ll.NodeType) bool { return t == ll.FMulExpr }
	FMulInst                        = func(t ll.NodeType) bool { return t == ll.FMulInst }
	FPExtExpr                       = func(t ll.NodeType) bool { return t == ll.FPExtExpr }
	FPExtInst                       = func(t ll.NodeType) bool { return t == ll.FPExtInst }
	FPToSIExpr                      = func(t ll.NodeType) bool { return t == ll.FPToSIExpr }
	FPToSIInst                      = func(t ll.NodeType) bool { return t == ll.FPToSIInst }
	FPToUIExpr                      = func(t ll.NodeType) bool { return t == ll.FPToUIExpr }
	FPToUIInst                      = func(t ll.NodeType) bool { return t == ll.FPToUIInst }
	FPTruncExpr                     = func(t ll.NodeType) bool { return t == ll.FPTruncExpr }
	FPTruncInst                     = func(t ll.NodeType) bool { return t == ll.FPTruncInst }
	FPred                           = func(t ll.NodeType) bool { return t == ll.FPred }
	FRemExpr                        = func(t ll.NodeType) bool { return t == ll.FRemExpr }
	FRemInst                        = func(t ll.NodeType) bool { return t == ll.FRemInst }
	FSubExpr                        = func(t ll.NodeType) bool { return t == ll.FSubExpr }
	FSubInst                        = func(t ll.NodeType) bool { return t == ll.FSubInst }
	FastMathFlag                    = func(t ll.NodeType) bool { return t == ll.FastMathFlag }
	FenceInst                       = func(t ll.NodeType) bool { return t == ll.FenceInst }
	FileField                       = func(t ll.NodeType) bool { return t == ll.FileField }
	FilenameField                   = func(t ll.NodeType) bool { return t == ll.FilenameField }
	FilterClause                    = func(t ll.NodeType) bool { return t == ll.FilterClause }
	FlagsField                      = func(t ll.NodeType) bool { return t == ll.FlagsField }
	FlagsStringField                = func(t ll.NodeType) bool { return t == ll.FlagsStringField }
	FloatConst                      = func(t ll.NodeType) bool { return t == ll.FloatConst }
	FloatKind                       = func(t ll.NodeType) bool { return t == ll.FloatKind }
	FloatLit                        = func(t ll.NodeType) bool { return t == ll.FloatLit }
	FloatType                       = func(t ll.NodeType) bool { return t == ll.FloatType }
	FuncAttribute                   = func(t ll.NodeType) bool { return t == ll.FuncAttribute }
	FuncBody                        = func(t ll.NodeType) bool { return t == ll.FuncBody }
	FuncDecl                        = func(t ll.NodeType) bool { return t == ll.FuncDecl }
	FuncDef                         = func(t ll.NodeType) bool { return t == ll.FuncDef }
	FuncHeader                      = func(t ll.NodeType) bool { return t == ll.FuncHeader }
	FuncMetadata                    = func(t ll.NodeType) bool { return t == ll.FuncMetadata }
	FuncType                        = func(t ll.NodeType) bool { return t == ll.FuncType }
	GCNode                          = func(t ll.NodeType) bool { return t == ll.GCNode }
	GEPIndex                        = func(t ll.NodeType) bool { return t == ll.GEPIndex }
	GenericDINode                   = func(t ll.NodeType) bool { return t == ll.GenericDINode }
	GetElementPtrExpr               = func(t ll.NodeType) bool { return t == ll.GetElementPtrExpr }
	GetElementPtrInst               = func(t ll.NodeType) bool { return t == ll.GetElementPtrInst }
	GetterField                     = func(t ll.NodeType) bool { return t == ll.GetterField }
	GlobalDecl                      = func(t ll.NodeType) bool { return t == ll.GlobalDecl }
	GlobalDef                       = func(t ll.NodeType) bool { return t == ll.GlobalDef }
	GlobalIdent                     = func(t ll.NodeType) bool { return t == ll.GlobalIdent }
	GlobalsField                    = func(t ll.NodeType) bool { return t == ll.GlobalsField }
	HeaderField                     = func(t ll.NodeType) bool { return t == ll.HeaderField }
	ICmpExpr                        = func(t ll.NodeType) bool { return t == ll.ICmpExpr }
	ICmpInst                        = func(t ll.NodeType) bool { return t == ll.ICmpInst }
	IFuncDef                        = func(t ll.NodeType) bool { return t == ll.IFuncDef }
	IPred                           = func(t ll.NodeType) bool { return t == ll.IPred }
	IdentifierField                 = func(t ll.NodeType) bool { return t == ll.IdentifierField }
	Immutable                       = func(t ll.NodeType) bool { return t == ll.Immutable }
	ImportsField                    = func(t ll.NodeType) bool { return t == ll.ImportsField }
	InAlloca                        = func(t ll.NodeType) bool { return t == ll.InAlloca }
	InBounds                        = func(t ll.NodeType) bool { return t == ll.InBounds }
	InRange                         = func(t ll.NodeType) bool { return t == ll.InRange }
	Inc                             = func(t ll.NodeType) bool { return t == ll.Inc }
	IncludePathField                = func(t ll.NodeType) bool { return t == ll.IncludePathField }
	IndirectBrTerm                  = func(t ll.NodeType) bool { return t == ll.IndirectBrTerm }
	InlineAsm                       = func(t ll.NodeType) bool { return t == ll.InlineAsm }
	InlinedAtField                  = func(t ll.NodeType) bool { return t == ll.InlinedAtField }
	InsertElementExpr               = func(t ll.NodeType) bool { return t == ll.InsertElementExpr }
	InsertElementInst               = func(t ll.NodeType) bool { return t == ll.InsertElementInst }
	InsertValueExpr                 = func(t ll.NodeType) bool { return t == ll.InsertValueExpr }
	InsertValueInst                 = func(t ll.NodeType) bool { return t == ll.InsertValueInst }
	InstMetadata                    = func(t ll.NodeType) bool { return t == ll.InstMetadata }
	IntConst                        = func(t ll.NodeType) bool { return t == ll.IntConst }
	IntLit                          = func(t ll.NodeType) bool { return t == ll.IntLit }
	IntToPtrExpr                    = func(t ll.NodeType) bool { return t == ll.IntToPtrExpr }
	IntToPtrInst                    = func(t ll.NodeType) bool { return t == ll.IntToPtrInst }
	IntType                         = func(t ll.NodeType) bool { return t == ll.IntType }
	IntelDialect                    = func(t ll.NodeType) bool { return t == ll.IntelDialect }
	InvokeTerm                      = func(t ll.NodeType) bool { return t == ll.InvokeTerm }
	IsDefinitionField               = func(t ll.NodeType) bool { return t == ll.IsDefinitionField }
	IsImplicitCodeField             = func(t ll.NodeType) bool { return t == ll.IsImplicitCodeField }
	IsLocalField                    = func(t ll.NodeType) bool { return t == ll.IsLocalField }
	IsOptimizedField                = func(t ll.NodeType) bool { return t == ll.IsOptimizedField }
	IsUnsignedField                 = func(t ll.NodeType) bool { return t == ll.IsUnsignedField }
	IsysrootField                   = func(t ll.NodeType) bool { return t == ll.IsysrootField }
	LShrExpr                        = func(t ll.NodeType) bool { return t == ll.LShrExpr }
	LShrInst                        = func(t ll.NodeType) bool { return t == ll.LShrInst }
	Label                           = func(t ll.NodeType) bool { return t == ll.Label }
	LabelIdent                      = func(t ll.NodeType) bool { return t == ll.LabelIdent }
	LabelType                       = func(t ll.NodeType) bool { return t == ll.LabelType }
	LandingPadInst                  = func(t ll.NodeType) bool { return t == ll.LandingPadInst }
	LanguageField                   = func(t ll.NodeType) bool { return t == ll.LanguageField }
	LineField                       = func(t ll.NodeType) bool { return t == ll.LineField }
	Linkage                         = func(t ll.NodeType) bool { return t == ll.Linkage }
	LinkageNameField                = func(t ll.NodeType) bool { return t == ll.LinkageNameField }
	LoadInst                        = func(t ll.NodeType) bool { return t == ll.LoadInst }
	LocalDefInst                    = func(t ll.NodeType) bool { return t == ll.LocalDefInst }
	LocalDefTerm                    = func(t ll.NodeType) bool { return t == ll.LocalDefTerm }
	LocalIdent                      = func(t ll.NodeType) bool { return t == ll.LocalIdent }
	LowerBoundField                 = func(t ll.NodeType) bool { return t == ll.LowerBoundField }
	MDFields                        = func(t ll.NodeType) bool { return t == ll.MDFields }
	MDString                        = func(t ll.NodeType) bool { return t == ll.MDString }
	MDTuple                         = func(t ll.NodeType) bool { return t == ll.MDTuple }
	MMXType                         = func(t ll.NodeType) bool { return t == ll.MMXType }
	MacrosField                     = func(t ll.NodeType) bool { return t == ll.MacrosField }
	MetadataAttachment              = func(t ll.NodeType) bool { return t == ll.MetadataAttachment }
	MetadataDef                     = func(t ll.NodeType) bool { return t == ll.MetadataDef }
	MetadataID                      = func(t ll.NodeType) bool { return t == ll.MetadataID }
	MetadataName                    = func(t ll.NodeType) bool { return t == ll.MetadataName }
	MetadataType                    = func(t ll.NodeType) bool { return t == ll.MetadataType }
	Module                          = func(t ll.NodeType) bool { return t == ll.Module }
	ModuleAsm                       = func(t ll.NodeType) bool { return t == ll.ModuleAsm }
	MulExpr                         = func(t ll.NodeType) bool { return t == ll.MulExpr }
	MulInst                         = func(t ll.NodeType) bool { return t == ll.MulInst }
	NameField                       = func(t ll.NodeType) bool { return t == ll.NameField }
	NameTableKind                   = func(t ll.NodeType) bool { return t == ll.NameTableKind }
	NameTableKindField              = func(t ll.NodeType) bool { return t == ll.NameTableKindField }
	NamedMetadataDef                = func(t ll.NodeType) bool { return t == ll.NamedMetadataDef }
	NamedType                       = func(t ll.NodeType) bool { return t == ll.NamedType }
	NodesField                      = func(t ll.NodeType) bool { return t == ll.NodesField }
	NoneConst                       = func(t ll.NodeType) bool { return t == ll.NoneConst }
	NullConst                       = func(t ll.NodeType) bool { return t == ll.NullConst }
	NullLit                         = func(t ll.NodeType) bool { return t == ll.NullLit }
	OffsetField                     = func(t ll.NodeType) bool { return t == ll.OffsetField }
	OpaqueType                      = func(t ll.NodeType) bool { return t == ll.OpaqueType }
	OperandBundle                   = func(t ll.NodeType) bool { return t == ll.OperandBundle }
	OperandBundles                  = func(t ll.NodeType) bool { return t == ll.OperandBundles }
	OperandsField                   = func(t ll.NodeType) bool { return t == ll.OperandsField }
	OrExpr                          = func(t ll.NodeType) bool { return t == ll.OrExpr }
	OrInst                          = func(t ll.NodeType) bool { return t == ll.OrInst }
	OverflowFlag                    = func(t ll.NodeType) bool { return t == ll.OverflowFlag }
	PackedStructType                = func(t ll.NodeType) bool { return t == ll.PackedStructType }
	Param                           = func(t ll.NodeType) bool { return t == ll.Param }
	ParamAttribute                  = func(t ll.NodeType) bool { return t == ll.ParamAttribute }
	Params                          = func(t ll.NodeType) bool { return t == ll.Params }
	Personality                     = func(t ll.NodeType) bool { return t == ll.Personality }
	PhiInst                         = func(t ll.NodeType) bool { return t == ll.PhiInst }
	PointerType                     = func(t ll.NodeType) bool { return t == ll.PointerType }
	Preemption                      = func(t ll.NodeType) bool { return t == ll.Preemption }
	Prefix                          = func(t ll.NodeType) bool { return t == ll.Prefix }
	ProducerField                   = func(t ll.NodeType) bool { return t == ll.ProducerField }
	Prologue                        = func(t ll.NodeType) bool { return t == ll.Prologue }
	PtrToIntExpr                    = func(t ll.NodeType) bool { return t == ll.PtrToIntExpr }
	PtrToIntInst                    = func(t ll.NodeType) bool { return t == ll.PtrToIntInst }
	ResumeTerm                      = func(t ll.NodeType) bool { return t == ll.ResumeTerm }
	RetTerm                         = func(t ll.NodeType) bool { return t == ll.RetTerm }
	RetainedNodesField              = func(t ll.NodeType) bool { return t == ll.RetainedNodesField }
	RetainedTypesField              = func(t ll.NodeType) bool { return t == ll.RetainedTypesField }
	ReturnAttribute                 = func(t ll.NodeType) bool { return t == ll.ReturnAttribute }
	RuntimeLangField                = func(t ll.NodeType) bool { return t == ll.RuntimeLangField }
	RuntimeVersionField             = func(t ll.NodeType) bool { return t == ll.RuntimeVersionField }
	SDivExpr                        = func(t ll.NodeType) bool { return t == ll.SDivExpr }
	SDivInst                        = func(t ll.NodeType) bool { return t == ll.SDivInst }
	SExtExpr                        = func(t ll.NodeType) bool { return t == ll.SExtExpr }
	SExtInst                        = func(t ll.NodeType) bool { return t == ll.SExtInst }
	SIToFPExpr                      = func(t ll.NodeType) bool { return t == ll.SIToFPExpr }
	SIToFPInst                      = func(t ll.NodeType) bool { return t == ll.SIToFPInst }
	SRemExpr                        = func(t ll.NodeType) bool { return t == ll.SRemExpr }
	SRemInst                        = func(t ll.NodeType) bool { return t == ll.SRemInst }
	ScopeField                      = func(t ll.NodeType) bool { return t == ll.ScopeField }
	ScopeLineField                  = func(t ll.NodeType) bool { return t == ll.ScopeLineField }
	Section                         = func(t ll.NodeType) bool { return t == ll.Section }
	SelectExpr                      = func(t ll.NodeType) bool { return t == ll.SelectExpr }
	SelectInst                      = func(t ll.NodeType) bool { return t == ll.SelectInst }
	SelectionKind                   = func(t ll.NodeType) bool { return t == ll.SelectionKind }
	SetterField                     = func(t ll.NodeType) bool { return t == ll.SetterField }
	ShlExpr                         = func(t ll.NodeType) bool { return t == ll.ShlExpr }
	ShlInst                         = func(t ll.NodeType) bool { return t == ll.ShlInst }
	ShuffleVectorExpr               = func(t ll.NodeType) bool { return t == ll.ShuffleVectorExpr }
	ShuffleVectorInst               = func(t ll.NodeType) bool { return t == ll.ShuffleVectorInst }
	SideEffect                      = func(t ll.NodeType) bool { return t == ll.SideEffect }
	SizeField                       = func(t ll.NodeType) bool { return t == ll.SizeField }
	SourceField                     = func(t ll.NodeType) bool { return t == ll.SourceField }
	SourceFilename                  = func(t ll.NodeType) bool { return t == ll.SourceFilename }
	SplitDebugFilenameField         = func(t ll.NodeType) bool { return t == ll.SplitDebugFilenameField }
	SplitDebugInliningField         = func(t ll.NodeType) bool { return t == ll.SplitDebugInliningField }
	StackAlignment                  = func(t ll.NodeType) bool { return t == ll.StackAlignment }
	StoreInst                       = func(t ll.NodeType) bool { return t == ll.StoreInst }
	StringLit                       = func(t ll.NodeType) bool { return t == ll.StringLit }
	StructConst                     = func(t ll.NodeType) bool { return t == ll.StructConst }
	StructType                      = func(t ll.NodeType) bool { return t == ll.StructType }
	SubExpr                         = func(t ll.NodeType) bool { return t == ll.SubExpr }
	SubInst                         = func(t ll.NodeType) bool { return t == ll.SubInst }
	SwiftError                      = func(t ll.NodeType) bool { return t == ll.SwiftError }
	SwitchTerm                      = func(t ll.NodeType) bool { return t == ll.SwitchTerm }
	SyncScope                       = func(t ll.NodeType) bool { return t == ll.SyncScope }
	TLSModel                        = func(t ll.NodeType) bool { return t == ll.TLSModel }
	TagField                        = func(t ll.NodeType) bool { return t == ll.TagField }
	Tail                            = func(t ll.NodeType) bool { return t == ll.Tail }
	TargetDataLayout                = func(t ll.NodeType) bool { return t == ll.TargetDataLayout }
	TargetTriple                    = func(t ll.NodeType) bool { return t == ll.TargetTriple }
	TemplateParamsField             = func(t ll.NodeType) bool { return t == ll.TemplateParamsField }
	ThisAdjustmentField             = func(t ll.NodeType) bool { return t == ll.ThisAdjustmentField }
	ThreadLocal                     = func(t ll.NodeType) bool { return t == ll.ThreadLocal }
	ThrownTypesField                = func(t ll.NodeType) bool { return t == ll.ThrownTypesField }
	TokenType                       = func(t ll.NodeType) bool { return t == ll.TokenType }
	TruncExpr                       = func(t ll.NodeType) bool { return t == ll.TruncExpr }
	TruncInst                       = func(t ll.NodeType) bool { return t == ll.TruncInst }
	TypeConst                       = func(t ll.NodeType) bool { return t == ll.TypeConst }
	TypeDef                         = func(t ll.NodeType) bool { return t == ll.TypeDef }
	TypeField                       = func(t ll.NodeType) bool { return t == ll.TypeField }
	TypeMacinfoField                = func(t ll.NodeType) bool { return t == ll.TypeMacinfoField }
	TypeValue                       = func(t ll.NodeType) bool { return t == ll.TypeValue }
	TypesField                      = func(t ll.NodeType) bool { return t == ll.TypesField }
	UDivExpr                        = func(t ll.NodeType) bool { return t == ll.UDivExpr }
	UDivInst                        = func(t ll.NodeType) bool { return t == ll.UDivInst }
	UIToFPExpr                      = func(t ll.NodeType) bool { return t == ll.UIToFPExpr }
	UIToFPInst                      = func(t ll.NodeType) bool { return t == ll.UIToFPInst }
	URemExpr                        = func(t ll.NodeType) bool { return t == ll.URemExpr }
	URemInst                        = func(t ll.NodeType) bool { return t == ll.URemInst }
	UintLit                         = func(t ll.NodeType) bool { return t == ll.UintLit }
	UndefConst                      = func(t ll.NodeType) bool { return t == ll.UndefConst }
	UnitField                       = func(t ll.NodeType) bool { return t == ll.UnitField }
	UnnamedAddr                     = func(t ll.NodeType) bool { return t == ll.UnnamedAddr }
	UnreachableTerm                 = func(t ll.NodeType) bool { return t == ll.UnreachableTerm }
	UnwindTarget                    = func(t ll.NodeType) bool { return t == ll.UnwindTarget }
	UseListOrder                    = func(t ll.NodeType) bool { return t == ll.UseListOrder }
	UseListOrderBB                  = func(t ll.NodeType) bool { return t == ll.UseListOrderBB }
	VAArgInst                       = func(t ll.NodeType) bool { return t == ll.VAArgInst }
	ValueField                      = func(t ll.NodeType) bool { return t == ll.ValueField }
	ValueIntField                   = func(t ll.NodeType) bool { return t == ll.ValueIntField }
	ValueStringField                = func(t ll.NodeType) bool { return t == ll.ValueStringField }
	VarField                        = func(t ll.NodeType) bool { return t == ll.VarField }
	VectorConst                     = func(t ll.NodeType) bool { return t == ll.VectorConst }
	VectorType                      = func(t ll.NodeType) bool { return t == ll.VectorType }
	VirtualIndexField               = func(t ll.NodeType) bool { return t == ll.VirtualIndexField }
	VirtualityField                 = func(t ll.NodeType) bool { return t == ll.VirtualityField }
	Visibility                      = func(t ll.NodeType) bool { return t == ll.Visibility }
	VoidType                        = func(t ll.NodeType) bool { return t == ll.VoidType }
	Volatile                        = func(t ll.NodeType) bool { return t == ll.Volatile }
	VtableHolderField               = func(t ll.NodeType) bool { return t == ll.VtableHolderField }
	Weak                            = func(t ll.NodeType) bool { return t == ll.Weak }
	XorExpr                         = func(t ll.NodeType) bool { return t == ll.XorExpr }
	XorInst                         = func(t ll.NodeType) bool { return t == ll.XorInst }
	ZExtExpr                        = func(t ll.NodeType) bool { return t == ll.ZExtExpr }
	ZExtInst                        = func(t ll.NodeType) bool { return t == ll.ZExtInst }
	ZeroInitializerConst            = func(t ll.NodeType) bool { return t == ll.ZeroInitializerConst }
	Clause                          = OneOf(ll.Clause...)
	ConcreteType                    = OneOf(ll.ConcreteType...)
	Constant                        = OneOf(ll.Constant...)
	ConstantExpr                    = OneOf(ll.ConstantExpr...)
	DIBasicTypeField                = OneOf(ll.DIBasicTypeField...)
	DICompileUnitField              = OneOf(ll.DICompileUnitField...)
	DICompositeTypeField            = OneOf(ll.DICompositeTypeField...)
	DIDerivedTypeField              = OneOf(ll.DIDerivedTypeField...)
	DIEnumeratorField               = OneOf(ll.DIEnumeratorField...)
	DIExpressionField               = OneOf(ll.DIExpressionField...)
	DIFileField                     = OneOf(ll.DIFileField...)
	DIGlobalVariableExpressionField = OneOf(ll.DIGlobalVariableExpressionField...)
	DIGlobalVariableField           = OneOf(ll.DIGlobalVariableField...)
	DIImportedEntityField           = OneOf(ll.DIImportedEntityField...)
	DILabelField                    = OneOf(ll.DILabelField...)
	DILexicalBlockField             = OneOf(ll.DILexicalBlockField...)
	DILexicalBlockFileField         = OneOf(ll.DILexicalBlockFileField...)
	DILocalVariableField            = OneOf(ll.DILocalVariableField...)
	DILocationField                 = OneOf(ll.DILocationField...)
	DIMacroField                    = OneOf(ll.DIMacroField...)
	DIMacroFileField                = OneOf(ll.DIMacroFileField...)
	DIModuleField                   = OneOf(ll.DIModuleField...)
	DINamespaceField                = OneOf(ll.DINamespaceField...)
	DIObjCPropertyField             = OneOf(ll.DIObjCPropertyField...)
	DISubprogramField               = OneOf(ll.DISubprogramField...)
	DISubrangeField                 = OneOf(ll.DISubrangeField...)
	DISubroutineTypeField           = OneOf(ll.DISubroutineTypeField...)
	DITemplateTypeParameterField    = OneOf(ll.DITemplateTypeParameterField...)
	DITemplateValueParameterField   = OneOf(ll.DITemplateValueParameterField...)
	FirstClassType                  = OneOf(ll.FirstClassType...)
	FuncAttr                        = OneOf(ll.FuncAttr...)
	GenericDINodeField              = OneOf(ll.GenericDINodeField...)
	GlobalAttr                      = OneOf(ll.GlobalAttr...)
	IndirectSymbolDef               = OneOf(ll.IndirectSymbolDef...)
	Instruction                     = OneOf(ll.Instruction...)
	MDField                         = OneOf(ll.MDField...)
	MDFieldOrInt                    = OneOf(ll.MDFieldOrInt...)
	MDNode                          = OneOf(ll.MDNode...)
	Metadata                        = OneOf(ll.Metadata...)
	MetadataNode                    = OneOf(ll.MetadataNode...)
	ParamAttr                       = OneOf(ll.ParamAttr...)
	ReturnAttr                      = OneOf(ll.ReturnAttr...)
	SpecializedMDNode               = OneOf(ll.SpecializedMDNode...)
	TargetDef                       = OneOf(ll.TargetDef...)
	Terminator                      = OneOf(ll.Terminator...)
	TopLevelEntity                  = OneOf(ll.TopLevelEntity...)
	Type                            = OneOf(ll.Type...)
	Value                           = OneOf(ll.Value...)
	ValueInstruction                = OneOf(ll.ValueInstruction...)
	ValueTerminator                 = OneOf(ll.ValueTerminator...)
)

func OneOf(types ...ll.NodeType) Selector {
	if len(types) == 0 {
		return func(ll.NodeType) bool { return false }
	}
	const bits = 32
	max := 1
	for _, t := range types {
		if int(t) > max {
			max = int(t)
		}
	}
	size := (max + bits) / bits
	bitarr := make([]uint32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t ll.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
