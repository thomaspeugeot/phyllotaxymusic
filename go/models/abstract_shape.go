package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type Shape struct {
	IsDisplayed   bool
	ShapeCategory *ShapeCategory

	// Shape needs the parameter to update other stacks
	// it is set up at run time
	parameter *Parameter
}

func (abstractShape *Shape) GetIsDisplayed() bool {
	return abstractShape.IsDisplayed
}

func (abstractShape *Shape) GetShapeCategory() *ShapeCategory {
	return abstractShape.ShapeCategory
}

func (abstractShape *Shape) SetCallbackOn(node *gongtree_models.Node, parameter *Parameter) {
	node.Impl = abstractShape
	abstractShape.parameter = parameter
}

// OnAfterUpdate implements gongtree_models.NodeImplInterface.
func (shape *Shape) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {
	log.Println("Node clicked", frontNode.GetName())

	p := shape.parameter

	if frontNode.IsChecked && !stagedNode.IsChecked {
		shape.IsDisplayed = true
		stagedNode.IsChecked = frontNode.IsChecked
		p.UpdatePhyllotaxyStage()
		p.UpdateAndCommitCursorStage()
		p.UpdateAndCommitSlidersStage()
		p.UpdateAndCommitSVGStage()
		p.UpdateAndCommitToneStage()
		p.treeProxy.UpdateAndCommitTreeStage()
		p.phyllotaxymusicStage.Commit()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		shape.IsDisplayed = false
		stagedNode.IsChecked = frontNode.IsChecked
		p.UpdatePhyllotaxyStage()
		p.UpdateAndCommitCursorStage()
		p.UpdateAndCommitSlidersStage()
		p.UpdateAndCommitSVGStage()
		p.UpdateAndCommitToneStage()
		p.treeProxy.UpdateAndCommitTreeStage()
		p.phyllotaxymusicStage.Commit()
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
