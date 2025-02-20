package models

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
