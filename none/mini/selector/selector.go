// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/mewspring/foo/none/mini"
)

type Selector func (nt mini.NodeType) bool

var (
	Any = func(t mini.NodeType) bool { return true }
	AddInst = func(t mini.NodeType) bool { return t == mini.AddInst }
	Args = func(t mini.NodeType) bool { return t == mini.Args }
	BasicBlock = func(t mini.NodeType) bool { return t == mini.BasicBlock }
	CallInst = func(t mini.NodeType) bool { return t == mini.CallInst }
	FloatKind = func(t mini.NodeType) bool { return t == mini.FloatKind }
	FloatType = func(t mini.NodeType) bool { return t == mini.FloatType }
	FooBar = func(t mini.NodeType) bool { return t == mini.FooBar }
	FooBaz = func(t mini.NodeType) bool { return t == mini.FooBaz }
	FuncBody = func(t mini.NodeType) bool { return t == mini.FuncBody }
	FuncDef = func(t mini.NodeType) bool { return t == mini.FuncDef }
	FuncHeader = func(t mini.NodeType) bool { return t == mini.FuncHeader }
	FuncMetadata = func(t mini.NodeType) bool { return t == mini.FuncMetadata }
	GlobalIdent = func(t mini.NodeType) bool { return t == mini.GlobalIdent }
	InstMetadata = func(t mini.NodeType) bool { return t == mini.InstMetadata }
	IntConst = func(t mini.NodeType) bool { return t == mini.IntConst }
	IntLit = func(t mini.NodeType) bool { return t == mini.IntLit }
	IntType = func(t mini.NodeType) bool { return t == mini.IntType }
	LocalDefInst = func(t mini.NodeType) bool { return t == mini.LocalDefInst }
	LocalIdent = func(t mini.NodeType) bool { return t == mini.LocalIdent }
	Module = func(t mini.NodeType) bool { return t == mini.Module }
	OperandBundles = func(t mini.NodeType) bool { return t == mini.OperandBundles }
	Params = func(t mini.NodeType) bool { return t == mini.Params }
	RetTerm = func(t mini.NodeType) bool { return t == mini.RetTerm }
	TypeValue = func(t mini.NodeType) bool { return t == mini.TypeValue }
	VoidType = func(t mini.NodeType) bool { return t == mini.VoidType }
	ConcreteType = OneOf(mini.ConcreteType...)
	Constant = OneOf(mini.Constant...)
	FirstClassType = OneOf(mini.FirstClassType...)
	Instruction = OneOf(mini.Instruction...)
	Terminator = OneOf(mini.Terminator...)
	TopLevelEntity = OneOf(mini.TopLevelEntity...)
	Type = OneOf(mini.Type...)
	Value = OneOf(mini.Value...)
	ValueInstruction = OneOf(mini.ValueInstruction...)
)

func OneOf(types ...mini.NodeType) Selector {
	if len(types) == 0 {
		return func(mini.NodeType) bool { return false }
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
	return func(t mini.NodeType) bool {
		i := uint(t)/bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
