// generated by Textmapper; DO NOT EDIT

package ast

import (
	"fmt"
	"github.com/mewspring/foo/issue25"
)

func ToIssue25Node(n *Node) Issue25Node {
	switch n.Type() {
	case issue25.Foo:
		return &Foo{n}
	case issue25.ID:
		return &ID{n}
	}
	panic(fmt.Errorf("ast: unknown node type %v", n.Type()))
	return nil
}