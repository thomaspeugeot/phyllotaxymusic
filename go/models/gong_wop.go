// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Axis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	Angle float64
	Length float64
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

func (from *Axis) CopyBasicFields(to *Axis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.Angle = from.Angle
	to.Length = from.Length
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

type AxisGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *AxisGrid) CopyBasicFields(to *AxisGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type Circle_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	CenterX float64
	CenterY float64
	HasBespokeRadius bool
	BespopkeRadius float64
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
	to.HasBespokeRadius = from.HasBespokeRadius
	to.BespopkeRadius = from.BespopkeRadius
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
	IsDisplayed bool
}

func (from *CircleGrid) CopyBasicFields(to *CircleGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type HorizontalAxis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	AxisHandleBorderLength float64
	Axis_Length float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *HorizontalAxis) CopyBasicFields(to *HorizontalAxis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.AxisHandleBorderLength = from.AxisHandleBorderLength
	to.Axis_Length = from.Axis_Length
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type Parameter_WOP struct {
	// insertion point
	Name string
	N int
	M int
	Z int
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
	to.Z = from.Z
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
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
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
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type RhombusGrid_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *RhombusGrid) CopyBasicFields(to *RhombusGrid) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
}

type VerticalAxis_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
	AxisHandleBorderLength float64
	Axis_Length float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *VerticalAxis) CopyBasicFields(to *VerticalAxis) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
	to.AxisHandleBorderLength = from.AxisHandleBorderLength
	to.Axis_Length = from.Axis_Length
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

