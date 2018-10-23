// copied from inspirer/textmapper

package main

import (
	"io/ioutil"

	"github.com/inspirer/textmapper/tm-go/status"
	"github.com/mewspring/foo/fob"
	"github.com/mewspring/foo/fob/ast"
	"github.com/pkg/errors"
)

func ParseFile(path string) (*ast.FuncHeader, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return Parse(path, content)
}

func Parse(filename, content string) (*ast.FuncHeader, error) {
	var l fob.Lexer
	l.Init(content)
	var p fob.Parser
	b := newBuilder(filename, content)
	//p.Init(b.addError, b.addNode)
	p.Init(b.addNode)
	err := p.Parse(&l)
	if err != nil {
		return nil, err
	}
	if err := b.status.Err(); err != nil {
		return nil, err
	}

	b.file.parsed = b.chunks
	x := ast.ToFobNode(b.file.root())
	return x.(*ast.FuncHeader), nil
}

type builder struct {
	file   *file
	chunks []chunk
	stack  []int
	status status.Status
}

func newBuilder(filename, content string) *builder {
	return &builder{
		file:   newFile(filename, content),
		chunks: []chunk{{offset: -1}},
		stack:  make([]int, 1, 512),
	}
}

func (b *builder) addError(se fob.SyntaxError) bool {
	r := b.file.sourceRange(se.Offset, se.Endoffset)
	b.status.Add(r, "syntax error")
	return true
}

func (b *builder) addNode(t fob.NodeType, offset, endoffset int) {
	if t == fob.FuncHeader {
		offset, endoffset = 0, len(b.file.content)
	}

	index := len(b.chunks)
	start := len(b.stack)
	end := start
	for o := b.chunks[b.stack[start-1]].offset; o >= offset; o = b.chunks[b.stack[start-1]].offset {
		start--
		if o >= endoffset {
			end--
		}
	}
	firstChild := 0
	if start < end {
		firstChild = b.stack[start]
		for _, i := range b.stack[start:end] {
			b.chunks[i].parent = index
		}
	}
	b.chunks[b.stack[start-1]].next = index
	if end == len(b.stack) {
		b.stack = append(b.stack[:start], index)
	} else if start < end {
		b.stack[start] = index
		l := copy(b.stack[start+1:], b.stack[end:])
		b.stack = b.stack[:start+1+l]
	} else {
		b.stack = append(b.stack, 0)
		copy(b.stack[start+1:], b.stack[start:])
		b.stack[start] = index
	}
	b.chunks = append(b.chunks, chunk{
		t:          t,
		offset:     offset,
		endoffset:  endoffset,
		firstChild: firstChild,
	})
}
