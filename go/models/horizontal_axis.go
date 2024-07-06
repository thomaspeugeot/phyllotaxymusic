package models

import (
	"log"

	"github.com/fullstack-lang/gongtree/go/models"
)

type HorizontalAxis struct {
	Name string

	IsHorizontalAxisDisplayed bool
	AxisHandleBorderLength    float64
	OriginX                   float64
	OriginY                   float64
	HorizontalAxis_Length     float64
	VerticalAxis_Length       float64

	Axis_StrokeWidth float64
}

// OnAfterUpdate implements models.NodeImplInterface.
func (horizontalAxis *HorizontalAxis) OnAfterUpdate(stage *models.StageStruct, stagedNode *models.Node, frontNode *models.Node) {
	log.Println("Node clicked", frontNode.GetName())

	if frontNode.IsChecked && !stagedNode.IsChecked {
		horizontalAxis.IsHorizontalAxisDisplayed = false
		GeneratorSingloton.Generate()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		horizontalAxis.IsHorizontalAxisDisplayed = true
		GeneratorSingloton.Generate()
	}
}
