package models

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type ShapeInterface interface {
	gongtree_models.NodeImplInterface
	GetName() string
	GetIsDisplayed() bool
	GetShapeCategory() *ShapeCategory
	Draw(gongsvgStage *gongsvg_models.StageStruct,
		layer *gongsvg_models.Layer,
		parameter *Parameter)
}
