package tokenizer

type TokenKind int

const (
	TK_NUM TokenKind = iota
	TK_IDENT
	TK_PUNCT
	TK_EOF
)

type Token struct {
	Kind TokenKind
	Num  int
	Text []byte
	Col  int
	Row  int
}

var puncts = []string{"+", "-", "*", "/", "%", "(", ")", ";", ",", "=", ":="}

func newToken(kind TokenKind, col int, row int) *Token {
	token := new(Token)
	token.Kind = kind
	token.Col = col
	token.Row = row
	return token
}

func newIdentToken(text []byte, col int, row int) *Token {
	token := new(Token)
	token.Kind = TK_IDENT
	token.Text = text
	token.Col = col
	token.Row = row
	return token
}

func newPunctToken(text []byte, col int, row int) *Token {
	token := new(Token)
	token.Kind = TK_PUNCT
	token.Text = text
	token.Col = col
	token.Row = row
	return token
}
