package models

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// ShapeInterface
type ShapeInterface interface {

	// It is call from a node update
	SetCallbackOn(*gongtree_models.Node, *Parameter)
	gongtree_models.NodeImplInterface

	GetName() string
	GetIsDisplayed() bool
	GetIsDisplayedPointer() *bool
	GetShapeCategory() *ShapeCategory
	Draw(gongsvgStage *gongsvg_models.StageStruct,
		layer *gongsvg_models.Layer,
		parameter *Parameter)
}
