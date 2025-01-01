package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type Shape struct {
	IsDisplayed   bool
	ShapeCategory *ShapeCategory
}

func (abstractShape *Shape) GetIsDisplayed() bool {
	return abstractShape.IsDisplayed
}

func (abstractShape *Shape) GetShapeCategory() *ShapeCategory {
	return abstractShape.ShapeCategory
}

// OnAfterUpdate implements models.NodeImplInterface.
func (shape *Shape) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {
	log.Println("Node clicked", frontNode.GetName())

	if frontNode.IsChecked && !stagedNode.IsChecked {
		shape.IsDisplayed = true
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		shape.IsDisplayed = false
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
}

func (s *Shape) DrawIfDisplayed(
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
