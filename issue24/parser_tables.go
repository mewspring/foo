// generated by Textmapper; DO NOT EDIT

package issue24

import (
	"fmt"
)

// Symbol represents a set of all terminal and non-terminal symbols of the issue24 language.
type Symbol int

var symbolStr = [...]string{
	"Foo",
	"list_of_','_and_1_elements",
	"ID",
}

func (n Symbol) String() string {
	if n < Symbol(NumTokens) {
		return Token(n).String()
	}
	i := int(n) - int(NumTokens)
	if i < len(symbolStr) {
		return symbolStr[i]
	}
	return fmt.Sprintf("nonterminal(%d)", n)
}

var tmAction = []int32{
	-1, -1, 3, -3, -9, -1, 1, -1, -2,
}

var tmLalr = []int32{
	0, 2, 4, 2, -1, -2, 4, -1, 0, 0, -1, -2,
}

var tmGoto = []int32{
	0, 2, 2, 6, 8, 10, 10, 12, 14, 18,
}

var tmFromTo = []int8{
	7, 8, 1, 2, 5, 2, 0, 1, 4, 5, 0, 7, 3, 4, 1, 3,
	5, 6,
}

var tmRuleLen = []int8{
	3, 3, 0, 1,
}

var tmRuleSymbol = []int32{
	6, 7, 7, 8,
}