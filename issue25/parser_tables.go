// generated by Textmapper; DO NOT EDIT

package issue25

import (
	"fmt"
)

// Symbol represents a set of all terminal and non-terminal symbols of the issue25 language.
type Symbol int

var symbolStr = [...]string{
	"Foo",
	"ID_list",
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
	-1, -1, 3, -1, 2, -1, 1, 0, -1, -2,
}

var tmGoto = []int32{
	0, 2, 2, 8, 10, 12, 12, 14, 16, 22,
}

var tmFromTo = []int8{
	8, 9, 1, 2, 3, 2, 5, 2, 0, 1, 3, 5, 0, 8, 1, 3,
	1, 4, 3, 6, 5, 7,
}

var tmRuleLen = []int8{
	4, 2, 1, 1,
}

var tmRuleSymbol = []int32{
	6, 7, 7, 8,
}
