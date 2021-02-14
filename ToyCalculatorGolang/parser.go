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
}

type toyParser struct {
	tokens []token
	depth  map[string]int
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
	p.depth = map[string]int{
		"E": 0,
		"T": 0,
		"F": 0,
	}
	p.tokens = tokens
	isMatch, parsedOffset, head := p.parseE(0)
	if !isMatch || parsedOffset != len(tokens) {
		panic("syntax error")
	}
	return head
}

func (p *toyParser) parseE(offset int) (match bool, parsedOffset int, n *node) {
	var isMatch bool
	if offset >= len(p.tokens) || p.depth["E"] >= MaxRecursion {
		return false, 0, nil
	}
	p.depth["E"]++
	val, ok := cache["E"][offset]
	if ok {
		p.depth["E"]--
		return ok, val.offset, val.n
	}

	isMatch, parsedOffset, n = p.parseE(offset)

	if isMatch {
		isMatchSecond, parsedOffsetSecond, nSecond := p.parseE(offset + 2)
		if parsedOffset+1 <= len(p.tokens) && p.tokens[parsedOffset].t == tokenTypePlus &&
			isMatchSecond {
			e := &node{
				t:        "E",
				children: nil,
			}
			e.children = []*node{n, {t: "+"}, nSecond}
			cache["E"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			p.depth["E"]--
			return true, parsedOffsetSecond, e
		} else if parsedOffset+1 <= len(p.tokens) && p.tokens[parsedOffset].t == tokenTypeMinus &&
			isMatchSecond {
			e := &node{
				t:        "E",
				children: nil,
			}
			e.children = []*node{n, {t: "-"}, nSecond}
			cache["E"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			p.depth["E"]--
			return true, parsedOffsetSecond, e
		} else {
			p.depth["E"]--
			return isMatch, parsedOffset, n
		}
	}
	isMatch, parsedOffset, n = p.parseT(offset)
	if isMatch {
		e := &node{
			t:        "E",
			children: nil,
		}
		e.children = []*node{n}
		p.depth["E"]--
		return isMatch, parsedOffset, e

	}

	panic(" syntax error")
}

func (p *toyParser) parseT(offset int) (match bool, parsedOffset int, n *node) {
	var isMatch bool
	if offset >= len(p.tokens) || p.depth["T"] >= MaxRecursion {
		return false, 0, nil
	}
	p.depth["T"]++
	val, ok := cache["T"][offset]
	if ok {
		p.depth["T"]--
		return ok, val.offset, val.n
	}

	isMatch, parsedOffset, n = p.parseT(offset)

	if isMatch {
		isMatchSecond, parsedOffsetSecond, nSecond := p.parseT(offset + 2)
		if parsedOffset+1 <= len(p.tokens) && p.tokens[parsedOffset].t == tokenTypeMultiply &&
			isMatchSecond {
			e := &node{
				t:        "T",
				children: nil,
			}
			e.children = []*node{n, {t: "*"}, nSecond}
			cache["T"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			p.depth["T"]--
			return true, parsedOffsetSecond, e
		} else if parsedOffset+1 <= len(p.tokens) && p.tokens[parsedOffset].t == tokenTypeDivide &&
			isMatchSecond {
			e := &node{
				t:        "T",
				children: nil,
			}
			e.children = []*node{n, {t: "/"}, nSecond}
			cache["T"][offset] = cacheValue{
				offset: parsedOffsetSecond,
				n:      e,
			}
			p.depth["T"]--
			return true, parsedOffsetSecond, e
		} else {
			p.depth["T"]--
			return isMatch, parsedOffset, n
		}
	}
	isMatch, parsedOffset, n = p.parseF(offset)
	if isMatch {
		e := &node{
			t:        "T",
			children: nil,
		}
		e.children = []*node{n}
		p.depth["T"]--
		return isMatch, parsedOffset, e
	}

	panic(" syntax error")
}

func (p *toyParser) parseF(offset int) (match bool, parsedOffset int, n *node) {
	// ( E )
	// num
	if offset >= len(p.tokens) || p.depth["F"] >= MaxRecursion {
		return false, 0, nil
	}
	p.depth["F"]++
	if p.tokens[offset].t == tokenTypeNumber {
		f := &node{
			t:        "num",
			children: nil,
			value:    p.tokens[offset].value,
		}
		p.depth["F"]--
		return true, offset + 1, f
	}
	if p.tokens[offset].t == tokenTypeLPar {
		isMatch, parsedOffset, n := p.parseT(offset + 1)
		if isMatch && p.tokens[parsedOffset].t == tokenTypeRPar {
			f := &node{
				t:        "f",
				children: []*node{{t: "("}, n, {t: ")"}},
			}
			p.depth["F"]--
			return true, parsedOffset + 1, f
		}
	}
	panic(" syntax error")
}
