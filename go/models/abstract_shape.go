package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongtree/go/models"
)

type AbstractShape struct {
	IsDisplayed   bool
	ShapeCategory *ShapeCategory
}

func (abstractShape *AbstractShape) GetIsDisplayed() bool {
	return abstractShape.IsDisplayed
}

func (abstractShape *AbstractShape) GetShapeCategory() *ShapeCategory {
	return abstractShape.ShapeCategory
}

// OnAfterUpdate implements models.NodeImplInterface.
func (hideableShape *AbstractShape) OnAfterUpdate(stage *models.StageStruct, stagedNode *models.Node, frontNode *models.Node) {
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

func (s *AbstractShape) DrawIfDisplayed(
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
