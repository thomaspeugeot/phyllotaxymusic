package models

import (
	"log"

	"github.com/fullstack-lang/gongtree/go/models"
)

type HorizontalAxis struct {
	Name string

	IsAxisDisplayed        bool
	AxisHandleBorderLength float64
	Axis_Length            float64

	StrokeWidth float64
}

// OnAfterUpdate implements models.NodeImplInterface.
func (horizontalAxis *HorizontalAxis) OnAfterUpdate(stage *models.StageStruct, stagedNode *models.Node, frontNode *models.Node) {
	log.Println("Node clicked", frontNode.GetName())

	if frontNode.IsChecked && !stagedNode.IsChecked {
		horizontalAxis.IsAxisDisplayed = true
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		horizontalAxis.IsAxisDisplayed = false
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
}
