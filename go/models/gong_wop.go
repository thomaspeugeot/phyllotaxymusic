// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
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
	DiamondAngle float64
	OriginX float64
	OriginY float64
	DiamondSideLenght float64
}

func (from *Parameter) CopyBasicFields(to *Parameter) {
	// insertion point
	to.Name = from.Name
	to.N = from.N
	to.M = from.M
	to.DiamondAngle = from.DiamondAngle
	to.OriginX = from.OriginX
	to.OriginY = from.OriginY
	to.DiamondSideLenght = from.DiamondSideLenght
}

type Rhombus_WOP struct {
	// insertion point
	Name string
	CenterX float64
	CenterY float64
	SideLength float64
	Angle float64
	InsideAngle float64
}

func (from *Rhombus) CopyBasicFields(to *Rhombus) {
	// insertion point
	to.Name = from.Name
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.SideLength = from.SideLength
	to.Angle = from.Angle
	to.InsideAngle = from.InsideAngle
}

