// Package parser implements function to parse the regular expressions.
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
package parser

import (
	"fmt"
	"log"

	"github.com/y-yu/sfa-go/lexer"
	"github.com/y-yu/sfa-go/node"
	"github.com/y-yu/sfa-go/token"
)

// Parser has a slice of tokens to parse, and now looking token.
type Parser struct {
	tokens []token.Token
	look   token.Token
}

// NewParser returns a new Parser with the tokens to
// parse that were obtained by scanning.
func NewParser(s string) *Parser {
	p := &Parser{
		tokens: lexer.NewLexer(s).Scan(),
	}
	p.move()
	return p
}

// GetAST returns the root node of AST obtained by parsing.
func (psr *Parser) GetAST() node.Node {
	ast := psr.expression()
	return ast
}

// move updates the now looking token to the next token in token slice.
// If token slice is empty, will set token.EOF as now looking token.
func (psr *Parser) move() {
	if len(psr.tokens) == 0 {
		psr.look = token.NewToken('\x00', token.EOF)
	} else {
		psr.look = psr.tokens[0]
		psr.tokens = psr.tokens[1:]
	}
}

// moveWithValidation execute move() with validating whether
// now looking Token type is an expected (or not).
func (psr *Parser) moveWithValidation(expect token.Type) {
	if psr.look.Ty != expect {
		err := fmt.Sprintf("[syntax error] expect:\x1b[31m%s\x1b[0m actual:\x1b[31m%s\x1b[0m", expect, psr.look.Ty)
		log.Fatal(err)
	}
	psr.move()
}

// expression -> subexpr
func (psr *Parser) expression() node.Node {
	nd := psr.subexpr()
	psr.moveWithValidation(token.EOF)
	return nd
}

// subexpr -> subexpr '|' seq | seq
// (
//
//	subexpr  -> seq _subexpr
//	_subexpr -> '|' seq _subexpr | ε
//
// )
func (psr *Parser) subexpr() node.Node {
	nd := psr.seq()
	for {
		if psr.look.Ty == token.UNION {
			psr.moveWithValidation(token.UNION)
			nd2 := psr.seq()
			nd = node.NewUnion(nd, nd2)
		} else {
			break
		}
	}
	return nd
}

// seq -> subseq | ε
func (psr *Parser) seq() node.Node {
	if psr.look.Ty == token.LPAREN || psr.look.Ty == token.CHARACTER {
		return psr.subseq()
	}
	return node.NewCharacter('ε')
}

// subseq -> subseq sufope | sufope
// (
//
//	subseq  -> sufope _subseq
//	_subseq -> sufope _subseq | ε
//
// )
func (psr *Parser) subseq() node.Node {
	nd := psr.sufope()
	if psr.look.Ty == token.LPAREN || psr.look.Ty == token.CHARACTER {
		nd2 := psr.subseq()
		return node.NewConcat(nd, nd2)
	}
	return nd
}

// sufope -> factor ('*'|'+') | factor
func (psr *Parser) sufope() node.Node {
	nd := psr.factor()
	switch psr.look.Ty {
	case token.STAR:
		psr.move()
		return node.NewStar(nd)
	case token.PLUS:
		psr.move()
		return node.NewPlus(nd)
	}
	return nd
}

// factor -> '(' subexpr ')' | CHARACTER
func (psr *Parser) factor() node.Node {
	if psr.look.Ty == token.LPAREN {
		psr.moveWithValidation(token.LPAREN)
		nd := psr.subexpr()
		psr.moveWithValidation(token.RPAREN)
		return nd
	}
	nd := node.NewCharacter(psr.look.V)
	psr.moveWithValidation(token.CHARACTER)
	return nd
}
