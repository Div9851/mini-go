package tokenizer

import "bytes"

func Tokenize(input []byte) []*Token {
	var tokens []*Token
	for {
		newToken, rest := nextToken(input)
		tokens = append(tokens, newToken)
		if newToken.Kind == TK_EOF {
			break
		}
		input = rest
	}
	return tokens
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func skipSpaces(input []byte) []byte {
	for len(input) > 0 && isSpace(input[0]) {
		input = input[1:]
	}
	return input
}

func nextToken(input []byte) (*Token, []byte) {
	input = skipSpaces(input)
	if len(input) == 0 {
		return newToken(TK_EOF), nil
	}
	// punctuator
	for _, punct := range puncts {
		if len(input) >= len(punct) && bytes.Equal(input[0:len(punct)], []byte(punct)) {
			newToken := newToken(TK_PUNCT)
			newToken.Text = input[0:len(punct)]
			return newToken, input[len(punct):]
		}
	}
	// number
	if isDigit(input[0]) {
		num := 0
		pos := 0
		for pos < len(input) && isDigit(input[pos]) {
			num *= 10
			num += int(input[pos] - '0')
			pos++
		}
		newToken := newToken(TK_NUM)
		newToken.Num = num
		newToken.Text = input[0:pos]
		return newToken, input[pos:]
	}
	// Unknown token
	return nil, nil
}
