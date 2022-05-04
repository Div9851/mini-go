package tokenizer

type TokenKind int

const (
	TK_NUM TokenKind = iota
	TK_PUNCT
	TK_EOF
)

type Token struct {
	Kind TokenKind
	Num  int
	Text []byte
}

var puncts = []string{"+", "-", "*", "/", "%", "(", ")"}

func newToken(kind TokenKind) *Token {
	token := new(Token)
	token.Kind = kind
	return token
}
