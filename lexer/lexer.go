// Package lexer implements lexer for a simple regular expression.
// MIT License
//
// # Copyright (c) 2019 8ayac
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package lexer

import (
	"github.com/y-yu/sfa-go/token"
)

// Lexer has a slice of symbols to analyze.
type Lexer struct {
	s []rune // string to be analyzed
}

// NewLexer returns a new Lexer.
// This constructor create a sequence of symbols from
// the string given in the argument and hold it.
func NewLexer(s string) *Lexer {
	return &Lexer{
		s: []rune(s),
	}
}

// Scan returns the token list to which converted from
// the symbol slice held in Lexer struct.
func (l *Lexer) Scan() (tokenList []token.Token) {
	for i := 0; i < len(l.s); i++ {
		switch l.s[i] {
		case '\x00':
			tokenList = append(tokenList, token.NewToken(l.s[i], token.EOF))
		case '|':
			tokenList = append(tokenList, token.NewToken(l.s[i], token.UNION))
		case '(':
			tokenList = append(tokenList, token.NewToken(l.s[i], token.LPAREN))
		case ')':
			tokenList = append(tokenList, token.NewToken(l.s[i], token.RPAREN))
		case '*':
			tokenList = append(tokenList, token.NewToken(l.s[i], token.STAR))
		case '+':
			tokenList = append(tokenList, token.NewToken(l.s[i], token.PLUS))
		case '\\':
			tokenList = append(tokenList, token.NewToken(l.s[i+1], token.CHARACTER))
			i++
		default:
			tokenList = append(tokenList, token.NewToken(l.s[i], token.CHARACTER))
		}
	}
	return
}
