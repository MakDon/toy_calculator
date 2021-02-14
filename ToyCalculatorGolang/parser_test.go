package ToyCalculator

import (
	"fmt"
	"testing"
)

func TestToyParser(t *testing.T) {
	tokens := []token{{t: tokenTypeNumber, value: 3}, {t: tokenTypePlus}, {t: tokenTypeNumber, value: 3.5}}
	parser := toyParser{}
	head := parser.parse(tokens)
	fmt.Print(head)
}
