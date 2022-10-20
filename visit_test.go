package ztpl

import (
	"os"
	"testing"
	"text/template/parse"
)

func mktree(t *testing.T, test string) *parse.Tree {
	tree, err := Parse("", test, parse.SkipFuncCheck, "{{", "}}")
	if err != nil {
		t.Fatal(err)
	}
	return tree[""]
}

func TestVisit(t *testing.T) {
	tree := mktree(t, `Hello, {{if and 1 2}}{{"X"}} ASD {{.pipe | (a "a")}}{{end}}
{{- if two -}}bar{{- end -}}`)
	PrintTree(os.Stdout, tree.Root)
}
