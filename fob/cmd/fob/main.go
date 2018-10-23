/*
path: no_input.txt
type:  *ast.CallingConv
value: &{{<nil> 0}}
cc: &ast.CallingConv{}
cc.Node == nil: false
cc.Node: main.node{}
text:  ""

path: input.txt
type:  *ast.CallingConv
value: &{{0xc000094190 1}}
cc: &ast.CallingConv{
    Node: main.node{
        file: &main.file{
            filename: "input.txt",
            content:  "x86_stdcallcc\n",
            lines:    {0, 14},
            parsed:   {
                {t:0, offset:-1, endoffset:0, next:2, firstChild:0, parent:0},
                {t:2, offset:0, endoffset:13, next:0, firstChild:0, parent:2},
                {t:1, offset:0, endoffset:14, next:0, firstChild:1, parent:0},
            },
        },
        index: 1,
    },
}
cc.Node == nil: false
cc.Node: main.node{
    file: &main.file{
        filename: "input.txt",
        content:  "x86_stdcallcc\n",
        lines:    {0, 14},
        parsed:   {
            {t:0, offset:-1, endoffset:0, next:2, firstChild:0, parent:0},
            {t:2, offset:0, endoffset:13, next:0, firstChild:0, parent:2},
            {t:1, offset:0, endoffset:14, next:0, firstChild:1, parent:0},
        },
    },
    index: 1,
}
text:  "x86_stdcallcc"
*/

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kr/pretty"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		fmt.Println("path:", path)
		hdr, err := ParseFile(path)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		cc := hdr.CallingConv()
		fmt.Printf("type:  %T\n", cc)
		fmt.Printf("value: %v\n", cc)
		pretty.Println("cc:", cc)
		fmt.Println("cc.Node == nil:", cc.Node == nil)
		pretty.Println("cc.Node:", cc.Node)
		fmt.Printf("text:  %q\n", cc.Text())
		fmt.Println()
	}
}
