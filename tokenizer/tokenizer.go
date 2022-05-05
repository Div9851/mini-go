package tokenizer

import (
	"bytes"
)

type Tokenizer struct {
	Text []byte
	Rest []byte
	Col  int
	Row  int
}

func New(text []byte) *Tokenizer {
	t := new(Tokenizer)
	t.Text = append(text, 0)
	t.Rest = t.Text
	t.Col = 1
	t.Row = 1
	return t
}

func (t *Tokenizer) Tokenize() []*Token {
	var tokens []*Token
	for {
		tok := t.nextToken()
		tokens = append(tokens, tok)
		if tok.Kind == TK_EOF {
			break
		}
	}
	return tokens
}

func (t *Tokenizer) nextCh() byte {
	ch := t.Rest[0]
	t.Rest = t.Rest[1:]
	t.Col++
	if ch == '\n' {
		t.Row++
		t.Col = 1
	}
	return ch
}

func isLetter(ch byte) bool {
	return ch == '_' || ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (t *Tokenizer) nextToken() *Token {
	for isSpace(t.Rest[0]) {
		t.nextCh()
	}
	if t.Rest[0] == 0 {
		return newToken(TK_EOF, t.Col, t.Row)
	}
	// punctuator
	for _, punct := range puncts {
		sz := len(punct)
		if len(t.Rest) >= sz && bytes.Equal(t.Rest[:sz], []byte(punct)) {
			tok := newPunctToken(t.Rest[:sz], t.Col, t.Row)
			for i := 0; i < sz; i++ {
				t.nextCh()
			}
			return tok
		}
	}
	// number
	if isNumber(t.Rest[0]) {
		num := 0
		pos := 0
		for isNumber(t.Rest[pos]) {
			num *= 10
			num += int(t.Rest[pos] - '0')
			pos++
		}
		tok := newToken(TK_NUM, t.Col, t.Row)
		tok.Num = num
		tok.Text = t.Rest[:pos]
		for i := 0; i < pos; i++ {
			t.nextCh()
		}
		return tok
	}
	// ident
	if isLetter(t.Rest[0]) {
		pos := 0
		for isLetter(t.Rest[pos]) || isNumber(t.Rest[pos]) {
			pos++
		}
		tok := newIdentToken(t.Rest[:pos], t.Col, t.Row)
		for i := 0; i < pos; i++ {
			t.nextCh()
		}
		return tok
	}
	panic("unknown token")
}
