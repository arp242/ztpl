package parse

import (
	"fmt"
	"io"
	"strings"

	"zgo.at/zstd/zstring"
)

// PrintTree prints the tree to w.
func PrintTree(w io.Writer, node Node) {
	Visit(node, func(n Node, depth int) bool {
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
func Visit(node Node, f func(Node, int) bool) {
	visit(node, f, 0)
}

func visit(node Node, f func(Node, int) bool, depth int) {
	if !f(node, depth) {
		return
	}

	switch n := node.(type) {
	case *ListNode:
		depth++
		visitNodes(n.Nodes, f, depth)
	case *IfNode:
		depth++
		visitPipe(n.Pipe, f, depth)
		if n.List != nil {
			visitNodes(n.List.Nodes, f, depth)
		}
		if n.ElseList != nil {
			visitNodes(n.ElseList.Nodes, f, depth)
		}
	case *CommandNode:
		depth++
		visitNodes(n.Args, f, depth)
	case *ActionNode:
		depth++
		visitPipe(n.Pipe, f, depth)
	case *RangeNode:
		depth++
		visitPipe(n.Pipe, f, depth)
	case *WithNode:
		depth++
		visitPipe(n.Pipe, f, depth)
	case *TemplateNode:
		depth++
		visitPipe(n.Pipe, f, depth)
	case *PipeNode:
		depth++
		visitPipe(n, f, depth)
	}
}

func visitNodes(n []Node, f func(Node, int) bool, depth int) {
	for _, nn := range n {
		visit(nn, f, depth)
	}
}

func visitPipe(n *PipeNode, f func(Node, int) bool, depth int) {
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
