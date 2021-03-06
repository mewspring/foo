// generated by Textmapper; DO NOT EDIT

package collision

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	FuncHeader NodeType = iota + 1 // GC?
	GC
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"FuncHeader",
	"GC",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

var ruleNodeType = [...]NodeType{
	0,          // input : FuncHeader
	FuncHeader, // FuncHeader : GCopt
	GC,         // GC : 'gc' string_lit
	0,          // GCopt : GC
	0,          // GCopt :
}
