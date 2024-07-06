package tree

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
	phylotaxymusic_stack "github.com/thomaspeugeot/phylotaxymusic/go/stack"
)

// Tree, for tree workspace holds the supporting data for performing operation on the weber tree
type Tree struct {
	Stack *phylotaxymusic_stack.Stack

	TreeStack *gongtree_stack.Stack

	NodeTree *gongtree_models.Tree

	// AllDiagramNodes is a pratical way to access all diagrams nodes
	// tree
	AllDiagramNodes []*gongtree_models.Node

	// SVGGenerator is the callback for generating the SVG
	SVGGenerator interface {
		GenerateSVG(parameter *phylotaxymusic_models.Parameter)
	}
}
