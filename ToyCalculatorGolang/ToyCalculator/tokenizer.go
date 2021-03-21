package ToyCalculator

import (
	"errors"
	"fmt"
	"strconv"
)

type tokenizer interface {
	tokenize(raw []byte) []token
}

var nums = map[byte]struct{}{'0': {}, '1': {}, '2': {}, '3': {}, '4': {},
	'5': {}, '6': {}, '7': {}, '8': {}, '9': {}}

type toyTokenizer struct {
}

func (t *toyTokenizer) tokenize(raw []byte) ([]token, error) {
	tokens := make([]token, 0)
	num := token{t: tokenTypeNumber}
	for idx := range raw {
		switch raw[idx] {
		case ' ':
			continue
		case '+':
			tokens = append(tokens, token{t: tokenTypePlus})
		case '-':
			tokens = append(tokens, token{t: tokenTypeMinus})
		case '*':
			tokens = append(tokens, token{t: tokenTypeMultiply})
		case '/':
			tokens = append(tokens, token{t: tokenTypeDivide})
		case '(':
			tokens = append(tokens, token{t: tokenTypeLPar})
		case ')':
			tokens = append(tokens, token{t: tokenTypeRPar})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num.s = append(num.s, raw[idx])
			if idx >= len(raw)-1 || (!isNum(raw[idx+1]) && raw[idx+1] != '.') {
				val, err := strconv.ParseFloat(string(num.s), 64)
				if err != nil {
					return tokens, err
				}
				num.value = val
				tokens = append(tokens, num)
				num = token{t: tokenTypeNumber}
			}
		case '.':
			if num.hasDot {
				return tokens, errors.New("more than 1 decimal points in a number")
			}
			num.s = append(num.s, '.')
			num.hasDot = true
		default:
			return tokens, fmt.Errorf("unknow char %s", string(raw[idx]))
		}
	}
	return tokens, nil
}
func (t token) DebugString() string {
	return fmt.Sprintf("%d, %f", t.t, t.value)
}

func (t *token) isEqual(target *token) bool {
	return t.t == target.t && t.value == target.value
}

func isNum(b byte) bool {
	_, ok := nums[b]
	return ok
}
