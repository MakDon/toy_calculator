package ToyCalculator

import (
	"fmt"
	"testing"
)

func TestTokenizer(t *testing.T) {
	raw := []byte("3.5 + 4.0 - ( 5+3)")
	targets := []token{{
		t:     tokenTypeNumber,
		value: 3.5,
	}, {
		t: tokenTypePlus,
	}, {
		t:     tokenTypeNumber,
		value: 4,
	}, {
		t: tokenTypeMinus,
	}, {
		t: tokenTypeLPar,
	}, {
		t:     tokenTypeNumber,
		value: 5,
	}, {
		t: tokenTypePlus,
	}, {
		t:     tokenTypeNumber,
		value: 3,
	}, {
		t: tokenTypeRPar,
	}}
	passed, _ := checkTokenization(raw, targets)
	if !passed {
		t.Fail()
	}
}

// TODO: test cases

func checkTokenization(raw []byte, targets []token) (bool, error) {
	tknz := &toyTokenizer{}
	tokens, err := tknz.tokenize(raw)
	if err != nil {
		return false, err
	}
	if len(targets) != len(tokens) {
		return false, fmt.Errorf("tokens len not equal")
	}
	for idx := range tokens {
		if !tokens[idx].isEqual(&targets[idx]) {
			return false, fmt.Errorf("token not equal")
		}
	}
	return true, nil
}
