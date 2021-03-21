package ToyCalculator

type token struct {
	t      int
	s      []byte
	hasDot bool
	value  float64
}

const (
	tokenTypeUnknown = iota

	tokenTypeNumber

	tokenTypePlus
	tokenTypeMinus
	tokenTypeMultiply
	tokenTypeDivide

	tokenTypeLPar
	tokenTypeRPar
)

var tokenTypeToString = map[int]string{
	tokenTypeUnknown: "unknown",

	tokenTypeNumber: "num",

	tokenTypePlus:     "+",
	tokenTypeMinus:    "-",
	tokenTypeMultiply: "*",
	tokenTypeDivide:   "/",

	tokenTypeLPar: "(",
	tokenTypeRPar: ")",
}

type node struct {
	t        string
	children []*node
	value    float64
}

type cacheValue struct {
	offset int
	n      *node
}

//
type cacheType map[string]map[int]cacheValue
