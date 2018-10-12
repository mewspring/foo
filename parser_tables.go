// generated by Textmapper; DO NOT EDIT

package foo

import (
	"fmt"
)

// Symbol represents a set of all terminal and non-terminal symbols of the foo language.
type Symbol int

var symbolStr = [...]string{
	"input",
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
	-1, 0, -1, -2,
}

var tmGoto = []int32{
	0, 2, 2, 4, 6,
}

var tmFromTo = []int8{
	2, 3, 0, 1, 0, 2,
}

var tmRuleLen = []int8{
	1,
}

var tmRuleSymbol = []int32{
	3,
}
