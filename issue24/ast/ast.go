// generated by Textmapper; DO NOT EDIT

package ast

import (
	"github.com/mewspring/foo/issue24/selector"
)

// Interfaces.

type Issue24Node interface {
	Issue24Node() *Node
}

// All types implement Issue24Node.
func (n Foo) Issue24Node() *Node { return n.Node }
func (n ID) Issue24Node() *Node { return n.Node }

// Types.

type Foo struct {
	*Node
}

func (n Foo) Src() ID {
	return ID{n.Child(selector.ID)}
}

func (n Foo) Indices() []ID {
	nodes := n.Child(selector.ID).NextAll(selector.ID)
	var ret = make([]ID, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ID{node})
	}
	return ret
}

type ID struct {
	*Node
}


