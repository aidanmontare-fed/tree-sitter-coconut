package tree_sitter_coconut_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_coconut "github.com/aidanmontare-fed/tree-sitter-coconut/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_coconut.Language())
	if language == nil {
		t.Errorf("Error loading Coconut grammar")
	}
}
