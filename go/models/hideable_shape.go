package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongtree/go/models"
)

type HideableShape struct {
	IsDisplayed bool
}

func (hideableShape *HideableShape) GetIsDisplayed() bool {
	return hideableShape.IsDisplayed
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

func (s *HideableShape) DrawIfDisplayed(
	gongsvgStage *gongsvg_models.StageStruct,
	p *Parameter,
	layer *gongsvg_models.Layer) {
}

type HideableShapeInterface interface {
	DrawIfDisplayed(
		gongsvgStage *gongsvg_models.StageStruct,
		p *Parameter,
		layer *gongsvg_models.Layer)
}
