// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/mewspring/foo/baz"
)

type Selector func (nt baz.NodeType) bool

var (
	Any = func(t baz.NodeType) bool { return true }
	ArrayType = func(t baz.NodeType) bool { return t == baz.ArrayType }
	IntType = func(t baz.NodeType) bool { return t == baz.IntType }
	LocalIdent = func(t baz.NodeType) bool { return t == baz.LocalIdent }
	MetadataType = func(t baz.NodeType) bool { return t == baz.MetadataType }
	Module = func(t baz.NodeType) bool { return t == baz.Module }
	OpaqueType = func(t baz.NodeType) bool { return t == baz.OpaqueType }
	TypeDef = func(t baz.NodeType) bool { return t == baz.TypeDef }
	UintLit = func(t baz.NodeType) bool { return t == baz.UintLit }
	VoidType = func(t baz.NodeType) bool { return t == baz.VoidType }
	ConcreteType = OneOf(baz.ConcreteType...)
	FirstClassType = OneOf(baz.FirstClassType...)
	TopLevelEntity = OneOf(baz.TopLevelEntity...)
	Type = OneOf(baz.Type...)
)

func OneOf(types ...baz.NodeType) Selector {
	if len(types) == 0 {
		return func(baz.NodeType) bool { return false }
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
	return func(t baz.NodeType) bool {
		i := uint(t)/bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
