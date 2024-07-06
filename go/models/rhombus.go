package models

import (
	"log"
	"math"

	"github.com/fullstack-lang/gongtree/go/models"
)

// Rhombus is defined by its center, a side length, a direction (with an angle)
// and an opening angle along this direction
type Rhombus struct {
	Name string

	IsDisplayed bool

	CenterX, CenterY float64

	SideLength float64

	// Angle of one arbitrary Axis of the Rhombus to the horizontal plan
	// in degress.
	Angle float64

	// InsideAngle is the angle inside the rhombus in the edge that
	// crosses the axis of the rhombus
	InsideAngle float64

	StrokeWidth float64
}

// OnAfterUpdate implements models.NodeImplInterface.
func (rhombus *Rhombus) OnAfterUpdate(stage *models.StageStruct, stagedNode *models.Node, frontNode *models.Node) {
	log.Println("Node clicked", frontNode.GetName())

	if frontNode.IsChecked && !stagedNode.IsChecked {
		rhombus.IsDisplayed = true
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		rhombus.IsDisplayed = false
		stagedNode.IsChecked = frontNode.IsChecked
		GeneratorSingloton.Generate()
	}
}

func GetRightSideLine(r Rhombus) (line *Line) {
	// Calculate half of the side length based on the angle for better accuracy
	halfSideLength := r.SideLength * math.Cos(math.Pi*r.InsideAngle/180.0) // Radians

	// Calculate starting point coordinates based on center, angle, and half side length
	startX := r.CenterX + halfSideLength*math.Cos(math.Pi*r.Angle/180.0)
	startY := r.CenterY + halfSideLength*math.Sin(math.Pi*r.Angle/180.0)

	// Calculate ending point coordinates by shifting based on half side length again
	endX := startX + halfSideLength*math.Cos((math.Pi*r.Angle+90.0)/180.0) // 90 degrees for right side
	endY := startY + halfSideLength*math.Sin((math.Pi*r.Angle+90.0)/180.0)

	// Create and return the line equation
	return &Line{X1: startX, Y1: startY, X2: endX, Y2: endY}
}
