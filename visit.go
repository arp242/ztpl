package ztpl

import (
	"fmt"
	"io"
	"strings"
	"text/template/parse"

	"zgo.at/zstd/zstring"
)

// Parse a template and set the mode.
func Parse(
	name, text string,
	mode parse.Mode,
	leftDelim, rightDelim string,
	funcs ...map[string]interface{},
) (map[string]*parse.Tree, error) {
	treeSet := make(map[string]*parse.Tree)
	t := parse.New(name)
	t.Mode = parse.SkipFuncCheck | parse.ParseComments
	_, err := t.Parse(text, "{{", "}}", treeSet, funcs...)
	return treeSet, err
}

// PrintTree prints the tree to w.
func PrintTree(w io.Writer, node parse.Node) {
	Visit(node, func(n parse.Node, depth int) bool {
		fmt.Fprint(w,
			strings.Repeat("    ", depth),
			zstring.AlignLeft(strings.TrimPrefix(fmt.Sprintf("%T", n), "*parse."), 30),
			" ", strings.ReplaceAll(n.String(), "\n", "\\n"), "\n")
		return true
	})
}

// Visit every node and call f.
//
// Traverse in the node if the return value of f is true.
func Visit(node parse.Node, f func(parse.Node, int) bool) {
	visit(node, f, 0)
}

func visit(node parse.Node, f func(parse.Node, int) bool, depth int) {
	if !f(node, depth) {
		return
	}

	switch n := node.(type) {
	case *parse.ListNode:
		depth++
		visitNodes(n.Nodes, f, depth)
	case *parse.IfNode:
		depth++
		visitPipe(n.Pipe, f, depth)
		if n.List != nil {
			visitNodes(n.List.Nodes, f, depth)
		}
		if n.ElseList != nil {
			visitNodes(n.ElseList.Nodes, f, depth)
		}
	case *parse.CommandNode:
		depth++
		visitNodes(n.Args, f, depth)
	case *parse.ActionNode:
		depth++
		visitPipe(n.Pipe, f, depth)
	case *parse.RangeNode:
		depth++
		visitPipe(n.Pipe, f, depth)
		if n.List != nil {
			visitNodes(n.List.Nodes, f, depth)
		}
	case *parse.WithNode:
		depth++
		visitPipe(n.Pipe, f, depth)
	case *parse.TemplateNode:
		depth++
		visitPipe(n.Pipe, f, depth)
	case *parse.PipeNode:
		depth++
		visitPipe(n, f, depth)
	}
}

func visitNodes(n []parse.Node, f func(parse.Node, int) bool, depth int) {
	for _, nn := range n {
		visit(nn, f, depth)
	}
}

func visitPipe(n *parse.PipeNode, f func(parse.Node, int) bool, depth int) {
	if n == nil {
		return
	}
	for _, nn := range n.Decl {
		visit(nn, f, depth)
	}
	for _, nn := range n.Cmds {
		visit(nn, f, depth)
	}
}
