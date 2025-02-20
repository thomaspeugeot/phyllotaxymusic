package models

import (
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type Shape struct {
	IsDisplayed   bool
	ShapeCategory *ShapeCategory
}

func (shape *Shape) GetIsDisplayed() bool {
	return shape.IsDisplayed
}

func (shape *Shape) GetIsDisplayedPointer() *bool {
	return &shape.IsDisplayed
}

func (shape *Shape) GetShapeCategory() *ShapeCategory {
	return shape.ShapeCategory
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
