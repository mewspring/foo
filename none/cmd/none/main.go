package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mewspring/foo/none"
	"github.com/mewspring/foo/none/ll/ast"
)

func main() {
	flag.Parse()
	for _, llPath := range flag.Args() {
		fmt.Printf("=== [ %v ] =======================\n", llPath)
		fmt.Println()
		fileStart := time.Now()
		parseStart := time.Now()
		module, err := none.ParseFile(llPath)
		if err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		fmt.Println("parsing into AST took:", time.Since(parseStart))
		fmt.Println()
		for _, entry := range module.TopLevelEntities() {
			f, ok := entry.(*ast.FuncDef)
			if !ok {
				continue
			}
			for _, block := range f.Body().Blocks() {
				for _, inst := range block.Insts() {
					i, ok := inst.(*ast.CallInst)
					if !ok {
						continue
					}
					fmt.Println("inst:", i.Text())
					fmt.Println("typ:", i.Typ())
				}
			}
		}
		fmt.Printf("total time for file %q: %v\n", llPath, time.Since(fileStart))
	}
}
