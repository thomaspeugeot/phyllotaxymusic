// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Cursor_WOP struct {
	// insertion point
	Name string
	AngleDegree float64
	Length float64
	CenterX float64
	CenterY float64
	StartX float64
	EndX float64
	DurationSeconds float64
	IsMoving bool
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
	IsPlaying bool
}

func (from *Cursor) CopyBasicFields(to *Cursor) {
	// insertion point
	to.Name = from.Name
	to.AngleDegree = from.AngleDegree
	to.Length = from.Length
	to.CenterX = from.CenterX
	to.CenterY = from.CenterY
	to.StartX = from.StartX
	to.EndX = from.EndX
	to.DurationSeconds = from.DurationSeconds
	to.IsMoving = from.IsMoving
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
	to.IsPlaying = from.IsPlaying
}

