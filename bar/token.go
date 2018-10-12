// generated by Textmapper; DO NOT EDIT

package bar

import (
	"fmt"
)

// Token is an enum of all terminal symbols of the bar language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota - 1
	EOI

	INVALID_TOKEN
	FOOBAR // foobar

	NumTokens
)

var tokenStr = [...]string{
	"EOI",

	"INVALID_TOKEN",
	"foobar",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}
