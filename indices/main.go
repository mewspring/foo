package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/llir/ll/ast"
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
	m, err := parseFile(llPath)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, entity := range m.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.FuncDef:
			body := entity.Body()
			for _, block := range body.Blocks() {
				for _, inst := range block.Insts() {
					if inst, ok := inst.(*ast.LocalDefInst); ok {
						switch i := inst.Inst().(type) {
						case *ast.GetElementPtrInst:
							fmt.Printf("text: %q\n", i.Text())
							fmt.Printf("elem type: %q\n", i.ElemType().LlvmNode().Text())
							fmt.Printf("src: %q\n", i.Src().LlvmNode().Text())
							fmt.Println("len(indices):", len(i.Indices()))
						}
					}
				}
			}
		}
	}
	return nil
}

// parseFile parses the given LLVM IR assembly file into an LLVM IR module.
func parseFile(path string) (*ast.Module, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return parse(path, content)
}

// parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from content.
func parse(path, content string) (*ast.Module, error) {
	tree, err := ast.Parse(path, content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse %q into AST", path)
	}
	root := ast.ToLlvmNode(tree.Root())
	return root.(*ast.Module), nil
}
