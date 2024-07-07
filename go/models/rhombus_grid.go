package models

import (
	"log"

	"github.com/fullstack-lang/gongtree/go/models"
)

type RhombusGrid struct {
	Name      string
	Reference *Rhombus
	N         int
	M         int

	IsDisplayed bool

	Rhombuses []*Rhombus
}

// OnAfterUpdate implements models.NodeImplInterface.
func (rhombusgrid *RhombusGrid) OnAfterUpdate(stage *models.StageStruct, stagedNode *models.Node, frontNode *models.Node) {
	log.Println("Node clicked", frontNode.GetName())

	if frontNode.IsChecked && !stagedNode.IsChecked {
		rhombusgrid.IsDisplayed = true
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		rhombusgrid.IsDisplayed = false
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
}
