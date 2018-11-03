package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mewspring/foo/issue24/ast"
	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	for _, llPath := range flag.Args() {
		if err := foo(llPath); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

func foo(llPath string) error {
	foo, err := parseFile(llPath)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("Src:", foo.Src().Text())
	fmt.Println("len(Indices):", len(foo.Indices()))
	return nil
}

// parseFile parses the given LLVM IR assembly file into an LLVM IR module.
func parseFile(path string) (*ast.Foo, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return parse(path, content)
}

// parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from content.
func parse(path, content string) (*ast.Foo, error) {
	tree, err := ast.Parse(path, content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse %q into AST", path)
	}
	root := ast.ToIssue24Node(tree.Root())
	return root.(*ast.Foo), nil
}
