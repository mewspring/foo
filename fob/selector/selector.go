// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/mewspring/foo/fob"
)

type Selector func(nt fob.NodeType) bool

var (
	Any         = func(t fob.NodeType) bool { return true }
	CallingConv = func(t fob.NodeType) bool { return t == fob.CallingConv }
	FuncHeader  = func(t fob.NodeType) bool { return t == fob.FuncHeader }
)

func OneOf(types ...fob.NodeType) Selector {
	if len(types) == 0 {
		return func(fob.NodeType) bool { return false }
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
	return func(t fob.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}