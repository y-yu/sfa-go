// Package token provides tokens for parsing the regular expressions.
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
package token

import "fmt"

// Type is integer to identify the type of token.
type Type int

// Each token is identified by a unique integer.
const (
	CHARACTER Type = iota
	UNION
	STAR
	PLUS
	LPAREN
	RPAREN
	EOF
)

func (k Type) String() string {
	switch k {
	case CHARACTER:
		return "CHARACTER"
	case UNION:
		return "UNION"
	case STAR:
		return "STAR"
	case PLUS:
		return "PLUS"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case EOF:
		return "EOF"
	default:
		return ""
	}
}

// Token represents a token.
type Token struct {
	V  rune // token value
	Ty Type // token type
}

func (t Token) String() string {
	return fmt.Sprintf("V -> \x1b[32m%v\x1b[0m\tKind -> \x1b[32m%v\x1b[0m", string(t.V), t.Ty)
}

// NewToken returns a new Token.
func NewToken(value rune, k Type) Token {
	return Token{
		V:  value,
		Ty: k,
	}
}
