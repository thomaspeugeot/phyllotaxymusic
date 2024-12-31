package models

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

// TreeProxy, is for construing the gongtree Tree
type TreeProxy struct {
	PhyllotaxyStage *StageStruct

	gongtreeStage *gongtree_models.StageStruct
	NodeTree      *gongtree_models.Tree

	// AllDiagramNodes is a pratical way to access all diagrams nodes
	// tree
	AllDiagramNodes []*gongtree_models.Node
}

func (tree *TreeProxy) SetGongtreeStage(gongtreeStage *gongtree_models.StageStruct) {
	tree.gongtreeStage = gongtreeStage
}
