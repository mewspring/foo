// generated by Textmapper; DO NOT EDIT

package ast

import (
	"github.com/mewspring/foo/fob"
	"log"
)

func ToFobNode(n Node) FobNode {
	if n == nil {
		return nil
	}
	switch n.Type() {
	case fob.CallingConv:
		return &CallingConv{n}
	case fob.FuncHeader:
		return &FuncHeader{n}
	}
	log.Fatalf("unknown node type %v\n", n.Type())
	return nil
}
