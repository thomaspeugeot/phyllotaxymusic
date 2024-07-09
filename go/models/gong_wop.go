// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Circle_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	CenterX float64
	CenterY float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Circle) CopyBasicFields(to *Circle) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type CircleGrid_WOP struct {
	// insertion point
	Name string
	N int
	M int
	IsDisplayed bool
}

func (from *CircleGrid) CopyBasicFields(to *CircleGrid) {
	// insertion point
	to.Name = from.Name
	to.N = from.N
	to.M = from.M
	to.IsDisplayed = from.IsDisplayed
}

type HorizontalAxis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	AxisHandleBorderLength float64
	Axis_Length float64
	StrokeWidth float64
}

func (from *HorizontalAxis) CopyBasicFields(to *HorizontalAxis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.AxisHandleBorderLength = from.AxisHandleBorderLength
	to.Axis_Length = from.Axis_Length
	to.StrokeWidth = from.StrokeWidth
}

type InitialAxis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	Angle float64
	Length float64
	StrokeWidth float64
}

func (from *InitialAxis) CopyBasicFields(to *InitialAxis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.Angle = from.Angle
	to.Length = from.Length
	to.StrokeWidth = from.StrokeWidth
}

type Line_WOP struct {
	// insertion point
	Name string
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func (from *Line) CopyBasicFields(to *Line) {
	// insertion point
	to.Name = from.Name
	to.X1 = from.X1
	to.Y1 = from.Y1
	to.X2 = from.X2
	to.Y2 = from.Y2
}

type Parameter_WOP struct {
	// insertion point
	Name string
	N int
	M int
	InsideAngle float64
	SideLength float64
	OriginX float64
	OriginY float64
}

func (from *Parameter) CopyBasicFields(to *Parameter) {
	// insertion point
	to.Name = from.Name
	to.N = from.N
	to.M = from.M
	to.InsideAngle = from.InsideAngle
	to.SideLength = from.SideLength
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
}

type Rhombus_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	CenterX float64
	CenterY float64
	SideLength float64
	Angle float64
	InsideAngle float64
	StrokeWidth float64
}

func (from *Rhombus) CopyBasicFields(to *Rhombus) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.SideLength = from.SideLength
	to.Angle = from.Angle
	to.InsideAngle = from.InsideAngle
	to.StrokeWidth = from.StrokeWidth
}

type RhombusGrid_WOP struct {
	// insertion point
	Name string
	N int
	M int
	IsDisplayed bool
}

func (from *RhombusGrid) CopyBasicFields(to *RhombusGrid) {
	// insertion point
	to.Name = from.Name
	to.N = from.N
	to.M = from.M
	to.IsDisplayed = from.IsDisplayed
}

type VerticalAxis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	AxisHandleBorderLength float64
	Axis_Length float64
	StrokeWidth float64
}

func (from *VerticalAxis) CopyBasicFields(to *VerticalAxis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.AxisHandleBorderLength = from.AxisHandleBorderLength
	to.Axis_Length = from.Axis_Length
	to.StrokeWidth = from.StrokeWidth
}

