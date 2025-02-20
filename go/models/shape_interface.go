package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// ShapeInterface
type ShapeInterface interface {
	GetName() string
	GetIsDisplayed() bool
	GetIsDisplayedPointer() *bool
	GetShapeCategory() *ShapeCategory
	Draw(gongsvgStage *gongsvg_models.StageStruct,
		layer *gongsvg_models.Layer,
		parameter *Parameter)
}
