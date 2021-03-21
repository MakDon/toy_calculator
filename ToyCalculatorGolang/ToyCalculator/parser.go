package ToyCalculator

var cache cacheType

var MaxRecursion = 5

func init() {
	cache = make(cacheType)
	cache["E"] = make(map[int]cacheValue)
	cache["T"] = make(map[int]cacheValue)
	cache["F"] = make(map[int]cacheValue)
}

type parser interface {
	parse(tokens []token) *node
}

type toyParser struct {
	tokens []token
	depth  map[string]map[int]int
}

// E -> E + E
//   -> E - E
//   -> T
// T -> T * T
//   -> T / T
//   -> F
// F -> ( E )
//   -> num
func (p *toyParser) parse(tokens []token) *node {
	p.depth = map[string]map[int]int{}
	p.tokens = tokens
	isMatch, parsedOffset, head := p.parseE(0)
	if !isMatch || parsedOffset != len(tokens) {
		panic("syntax error")
	}
	return head
}

func (p *toyParser) parseE(offset int) (match bool, parsedOffset int, n *node) {
	var isMatch bool
	if offset >= len(p.tokens) || p.depth["E"][offset] >= MaxRecursion {
		return false, 0, nil
	}
	depthPlus(p.depth, "E", offset)
	val, ok := cache["E"][offset]
	if ok {
		depthMinus(p.depth, "E", offset)
		return ok, val.offset, val.n
	}

	isMatch, parsedOffset, n = p.parseE(offset)

	if isMatch {
		if parsedOffset+1 <= len(p.tokens) && p.tokens[parsedOffset].t == tokenTypePlus {
			isMatchSecond, parsedOffsetSecond, nSecond := p.parseE(offset + 2)
			if !isMatchSecond {
				goto matchT
			}
			e := &node{
				t:        "E",
				children: nil,
			}
			e.children = []*node{n, {t: "+"}, nSecond}
			cache["E"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			depthMinus(p.depth, "E", offset)
			return true, parsedOffsetSecond, e
		} else if parsedOffset+1 <= len(p.tokens) && p.tokens[parsedOffset].t == tokenTypeMinus {
			isMatchSecond, parsedOffsetSecond, nSecond := p.parseE(offset + 2)
			if !isMatchSecond {
				goto matchT
			}
			e := &node{
				t:        "E",
				children: nil,
			}
			e.children = []*node{n, {t: "-"}, nSecond}
			cache["E"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			depthMinus(p.depth, "E", offset)
			return true, parsedOffsetSecond, e
		} else {
			depthMinus(p.depth, "E", offset)
			return isMatch, parsedOffset, n
		}
	}
matchT:
	isMatch, parsedOffset, n = p.parseT(offset)
	if isMatch {
		e := &node{
			t:        "E",
			children: nil,
		}
		e.children = []*node{n}
		depthMinus(p.depth, "E", offset)
		return isMatch, parsedOffset, e

	}

	panic(" syntax error")
}

func (p *toyParser) parseT(offset int) (match bool, parsedOffset int, n *node) {
	var isMatch bool
	if offset >= len(p.tokens) || p.depth["T"][offset] >= MaxRecursion {
		return false, 0, nil
	}
	depthPlus(p.depth, "T", offset)
	val, ok := cache["T"][offset]
	if ok {
		depthMinus(p.depth, "T", offset)
		return ok, val.offset, val.n
	}

	isMatch, parsedOffset, n = p.parseT(offset)

	if isMatch {
		if parsedOffset+1 < len(p.tokens) && p.tokens[parsedOffset].t == tokenTypeMultiply {
			isMatchSecond, parsedOffsetSecond, nSecond := p.parseT(offset + 2)
			if !isMatchSecond {
				goto matchF
			}
			e := &node{
				t:        "T",
				children: nil,
			}
			e.children = []*node{n, {t: "*"}, nSecond}
			cache["T"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			depthMinus(p.depth, "T", offset)
			return true, parsedOffsetSecond, e
		} else if parsedOffset+1 < len(p.tokens) && p.tokens[parsedOffset].t == tokenTypeDivide {
			isMatchSecond, parsedOffsetSecond, nSecond := p.parseT(offset + 2)
			if !isMatchSecond {
				goto matchF
			}
			e := &node{
				t:        "T",
				children: nil,
			}
			e.children = []*node{n, {t: "/"}, nSecond}
			cache["T"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			depthMinus(p.depth, "T", offset)
			return true, parsedOffsetSecond, e
		} else {
			depthMinus(p.depth, "T", offset)
			return isMatch, parsedOffset, n
		}
	}
matchF:
	isMatch, parsedOffset, n = p.parseF(offset)
	if isMatch {
		e := &node{
			t:        "T",
			children: nil,
		}
		e.children = []*node{n}
		depthMinus(p.depth, "T", offset)
		return isMatch, parsedOffset, e
	}

	panic(" syntax error")
}

func (p *toyParser) parseF(offset int) (match bool, parsedOffset int, n *node) {
	// ( E )
	// num
	if offset >= len(p.tokens) || p.depth["F"][offset] >= MaxRecursion {
		return false, 0, nil
	}
	depthPlus(p.depth, "F", offset)
	if p.tokens[offset].t == tokenTypeNumber {
		f := &node{
			t:        "num",
			children: nil,
			value:    p.tokens[offset].value,
		}
		depthMinus(p.depth, "F", offset)
		return true, offset + 1, f
	}
	if p.tokens[offset].t == tokenTypeLPar {
		isMatch, parsedOffset, n := p.parseE(offset + 1)
		if isMatch && p.tokens[parsedOffset].t == tokenTypeRPar {
			f := &node{
				t:        "f",
				children: []*node{{t: "("}, n, {t: ")"}},
			}
			depthMinus(p.depth, "F", offset)
			return true, parsedOffset + 1, f
		}
	}
	panic(" syntax error")
}

func depthPlus(m map[string]map[int]int, t string, offset int) {
	if val, ok := m[t]; ok {
		val[offset] = val[offset] + 1
	} else {
		m[t] = make(map[int]int)
		m[t][offset] = 1
	}
}

func depthMinus(m map[string]map[int]int, t string, offset int) {
	if val, ok := m[t]; ok {
		val[offset] = val[offset] - 1
	} else {
		m[t] = make(map[int]int)
		m[t][offset] = -1
	}

}
