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

func TestToyTokenizerAndParser(t *testing.T) {
	tokenRaw := []byte("3.5 + 4.0 - ( 5+3)")
	t2 := &token{
		t:      0,
		s:      nil,
		hasDot: false,
		value:  3,
	}
	_ = t2
	tknzer := toyTokenizer{}
	parser := toyParser{}
	tokens, err := tknzer.tokenize(tokenRaw)
	if err != nil {
		panic(err)
	}
	head := parser.parse(tokens)
	fmt.Print(head)
}
