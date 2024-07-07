package models

import (
	"log"

	"github.com/fullstack-lang/gongtree/go/models"
)

type HideableShape struct {
	IsDisplayed bool
}

// OnAfterUpdate implements models.NodeImplInterface.
func (hideableShape *HideableShape) OnAfterUpdate(stage *models.StageStruct, stagedNode *models.Node, frontNode *models.Node) {
	log.Println("Node clicked", frontNode.GetName())

	if frontNode.IsChecked && !stagedNode.IsChecked {
		hideableShape.IsDisplayed = true
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		hideableShape.IsDisplayed = false
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
}
