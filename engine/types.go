package engine

type Resquest struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Resquest
	Items     []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}