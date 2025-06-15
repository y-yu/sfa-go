package dfa

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

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/y-yu/sfa-go/utils"
	"os"
)

type CommonNodeAttrs map[string]string

func NewCommonNodeAttrs() CommonNodeAttrs {
	return CommonNodeAttrs{
		"fontname": "meiryo",
		"fontsize": "18",
	}
}

type CommonEdgeAttrs map[string]string

func NewCommonEdgeAttrs() CommonEdgeAttrs {
	return CommonEdgeAttrs{
		"fontname":   "meiryo",
		"fontsize":   "18",
		"len":        "1.5",
		"labelfloat": "false",
	}
}

func DFA2dot(d DFA, name string) {
	const GRAPH_NAME = "DFA"
	g := gographviz.NewGraph()

	// General
	_ = g.SetName(GRAPH_NAME)
	_ = g.SetDir(true)
	_ = g.AddAttr(GRAPH_NAME, "rankdir", "LR")

	// For initial state
	dummyAttrs := NewCommonNodeAttrs()
	dummyAttrs["shape"] = "point"
	_ = g.AddNode(GRAPH_NAME, "\"\"", dummyAttrs)
	_ = g.AddNode(GRAPH_NAME, d.I.String(), NewCommonNodeAttrs())

	initEdgeAttrs := NewCommonEdgeAttrs()
	initEdgeAttrs["len"] = "2"
	_ = g.AddEdge("\"\"", d.I.String(), true, initEdgeAttrs)

	// Make state nodes.
	states := utils.NewSet(d.I)
	for _, v := range d.Rules {
		states.Add(v)
	}
	for q := range states.Iter() {
		attrs := NewCommonNodeAttrs()
		if d.F.Contains(q) {
			attrs["shape"] = "doublecircle"
		}
		_ = g.AddNode(GRAPH_NAME, q.String(), attrs)
	}

	// Make edges from transition rules.
	rules := d.Rules
	for arg, dst := range rules {
		attrs := NewCommonEdgeAttrs()
		attrs["label"] = fmt.Sprintf("\"'%c'\"", arg.C)
		_ = g.AddEdge(arg.From.String(), dst.String(), true, attrs)
	}

	// Output DOT
	file, err := os.Create(fmt.Sprintf("%s.dot", name))
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	_, err = file.Write([]byte(g.String()))
	if err != nil {
		panic(err)
	}
}
