package main

import (
	"time"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Axis__000000_Reference := (&models.Axis{Name: `Reference`}).Stage(stage)
	__Axis__000001_Rotated := (&models.Axis{Name: `Rotated`}).Stage(stage)

	__Circle__000000_Initial_Circle := (&models.Circle{Name: `Initial Circle`}).Stage(stage)
	__Circle__000001_Next_Circle := (&models.Circle{Name: `Next Circle`}).Stage(stage)

	__CircleGrid__000000_Initial := (&models.CircleGrid{Name: `Initial`}).Stage(stage)
	__CircleGrid__000001_Rotated := (&models.CircleGrid{Name: `Rotated`}).Stage(stage)

	__HorizontalAxis__000000_Initial := (&models.HorizontalAxis{Name: `Initial`}).Stage(stage)

	__Parameter__000000_Reference := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__000000_Initial := (&models.Rhombus{Name: `Initial`}).Stage(stage)
	__Rhombus__000001_Next_Rhombus := (&models.Rhombus{Name: `Next Rhombus`}).Stage(stage)
	__Rhombus__000002_Rotated := (&models.Rhombus{Name: `Rotated`}).Stage(stage)

	__RhombusGrid__000000_Initial := (&models.RhombusGrid{Name: `Initial`}).Stage(stage)
	__RhombusGrid__000001_Rotated := (&models.RhombusGrid{Name: `Rotated`}).Stage(stage)

	__VerticalAxis__000000_Initial := (&models.VerticalAxis{Name: `Initial`}).Stage(stage)

	// Setup of values

	__Axis__000000_Reference.Name = `Reference`
	__Axis__000000_Reference.IsDisplayed = true
	__Axis__000000_Reference.Angle = 0.000000
	__Axis__000000_Reference.Length = 393.636130
	__Axis__000000_Reference.Color = ``
	__Axis__000000_Reference.FillOpacity = 0.000000
	__Axis__000000_Reference.Stroke = `black`
	__Axis__000000_Reference.StrokeOpacity = 1.000000
	__Axis__000000_Reference.StrokeWidth = 2.000000
	__Axis__000000_Reference.StrokeDashArray = ``
	__Axis__000000_Reference.StrokeDashArrayWhenSelected = ``
	__Axis__000000_Reference.Transform = ``

	__Axis__000001_Rotated.Name = `Rotated`
	__Axis__000001_Rotated.IsDisplayed = false
	__Axis__000001_Rotated.Angle = 80.800438
	__Axis__000001_Rotated.Length = 393.636130
	__Axis__000001_Rotated.Color = ``
	__Axis__000001_Rotated.FillOpacity = 0.000000
	__Axis__000001_Rotated.Stroke = `black`
	__Axis__000001_Rotated.StrokeOpacity = 1.000000
	__Axis__000001_Rotated.StrokeWidth = 2.000000
	__Axis__000001_Rotated.StrokeDashArray = ``
	__Axis__000001_Rotated.StrokeDashArrayWhenSelected = ``
	__Axis__000001_Rotated.Transform = ``

	__Circle__000000_Initial_Circle.Name = `Initial Circle`
	__Circle__000000_Initial_Circle.IsDisplayed = false
	__Circle__000000_Initial_Circle.CenterX = 0.000000
	__Circle__000000_Initial_Circle.CenterY = 0.000000
	__Circle__000000_Initial_Circle.Color = ``
	__Circle__000000_Initial_Circle.FillOpacity = 0.000000
	__Circle__000000_Initial_Circle.Stroke = `lightblue`
	__Circle__000000_Initial_Circle.StrokeOpacity = 0.800000
	__Circle__000000_Initial_Circle.StrokeWidth = 3.000000
	__Circle__000000_Initial_Circle.StrokeDashArray = ``
	__Circle__000000_Initial_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000000_Initial_Circle.Transform = ``

	__Circle__000001_Next_Circle.Name = `Next Circle`
	__Circle__000001_Next_Circle.IsDisplayed = true
	__Circle__000001_Next_Circle.CenterX = 153.429982
	__Circle__000001_Next_Circle.CenterY = 24.849030
	__Circle__000001_Next_Circle.Color = ``
	__Circle__000001_Next_Circle.FillOpacity = 0.000000
	__Circle__000001_Next_Circle.Stroke = `yellow`
	__Circle__000001_Next_Circle.StrokeOpacity = 0.800000
	__Circle__000001_Next_Circle.StrokeWidth = 3.000000
	__Circle__000001_Next_Circle.StrokeDashArray = ``
	__Circle__000001_Next_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000001_Next_Circle.Transform = ``

	__CircleGrid__000000_Initial.Name = `Initial`
	__CircleGrid__000000_Initial.IsDisplayed = false

	__CircleGrid__000001_Rotated.Name = `Rotated`
	__CircleGrid__000001_Rotated.IsDisplayed = true

	__HorizontalAxis__000000_Initial.Name = `Initial`
	__HorizontalAxis__000000_Initial.IsDisplayed = true
	__HorizontalAxis__000000_Initial.AxisHandleBorderLength = 0.000000
	__HorizontalAxis__000000_Initial.Axis_Length = 600.000000
	__HorizontalAxis__000000_Initial.Color = ``
	__HorizontalAxis__000000_Initial.FillOpacity = 0.000000
	__HorizontalAxis__000000_Initial.Stroke = ``
	__HorizontalAxis__000000_Initial.StrokeOpacity = 0.000000
	__HorizontalAxis__000000_Initial.StrokeWidth = 1.000000
	__HorizontalAxis__000000_Initial.StrokeDashArray = ``
	__HorizontalAxis__000000_Initial.StrokeDashArrayWhenSelected = ``
	__HorizontalAxis__000000_Initial.Transform = ``

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.InsideAngle = 102.000000
	__Parameter__000000_Reference.SideLength = 100.000000
	__Parameter__000000_Reference.OriginX = 300.000000
	__Parameter__000000_Reference.OriginY = 300.000000

	__Rhombus__000000_Initial.Name = `Initial`
	__Rhombus__000000_Initial.IsDisplayed = false
	__Rhombus__000000_Initial.CenterX = 0.000000
	__Rhombus__000000_Initial.CenterY = 0.000000
	__Rhombus__000000_Initial.SideLength = 100.000000
	__Rhombus__000000_Initial.Angle = 0.000000
	__Rhombus__000000_Initial.InsideAngle = 102.000000
	__Rhombus__000000_Initial.Color = ``
	__Rhombus__000000_Initial.FillOpacity = 0.000000
	__Rhombus__000000_Initial.Stroke = ``
	__Rhombus__000000_Initial.StrokeOpacity = 0.000000
	__Rhombus__000000_Initial.StrokeWidth = 1.000000
	__Rhombus__000000_Initial.StrokeDashArray = ``
	__Rhombus__000000_Initial.StrokeDashArrayWhenSelected = ``
	__Rhombus__000000_Initial.Transform = ``

	__Rhombus__000001_Next_Rhombus.Name = `Next Rhombus`
	__Rhombus__000001_Next_Rhombus.IsDisplayed = true
	__Rhombus__000001_Next_Rhombus.CenterX = 153.429982
	__Rhombus__000001_Next_Rhombus.CenterY = 24.849030
	__Rhombus__000001_Next_Rhombus.SideLength = 100.000000
	__Rhombus__000001_Next_Rhombus.Angle = -80.800438
	__Rhombus__000001_Next_Rhombus.InsideAngle = 102.000000
	__Rhombus__000001_Next_Rhombus.Color = ``
	__Rhombus__000001_Next_Rhombus.FillOpacity = 0.000000
	__Rhombus__000001_Next_Rhombus.Stroke = `lavender`
	__Rhombus__000001_Next_Rhombus.StrokeOpacity = 0.700000
	__Rhombus__000001_Next_Rhombus.StrokeWidth = 3.000000
	__Rhombus__000001_Next_Rhombus.StrokeDashArray = ``
	__Rhombus__000001_Next_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000001_Next_Rhombus.Transform = ``

	__Rhombus__000002_Rotated.Name = `Rotated`
	__Rhombus__000002_Rotated.IsDisplayed = true
	__Rhombus__000002_Rotated.CenterX = 0.000000
	__Rhombus__000002_Rotated.CenterY = 0.000000
	__Rhombus__000002_Rotated.SideLength = 100.000000
	__Rhombus__000002_Rotated.Angle = -80.800438
	__Rhombus__000002_Rotated.InsideAngle = 102.000000
	__Rhombus__000002_Rotated.Color = ``
	__Rhombus__000002_Rotated.FillOpacity = 0.000000
	__Rhombus__000002_Rotated.Stroke = `black`
	__Rhombus__000002_Rotated.StrokeOpacity = 1.000000
	__Rhombus__000002_Rotated.StrokeWidth = 2.000000
	__Rhombus__000002_Rotated.StrokeDashArray = ``
	__Rhombus__000002_Rotated.StrokeDashArrayWhenSelected = ``
	__Rhombus__000002_Rotated.Transform = ``

	__RhombusGrid__000000_Initial.Name = `Initial`
	__RhombusGrid__000000_Initial.IsDisplayed = false

	__RhombusGrid__000001_Rotated.Name = `Rotated`
	__RhombusGrid__000001_Rotated.IsDisplayed = true

	__VerticalAxis__000000_Initial.Name = `Initial`
	__VerticalAxis__000000_Initial.IsDisplayed = true
	__VerticalAxis__000000_Initial.AxisHandleBorderLength = 0.000000
	__VerticalAxis__000000_Initial.Axis_Length = 600.000000
	__VerticalAxis__000000_Initial.Color = ``
	__VerticalAxis__000000_Initial.FillOpacity = 0.000000
	__VerticalAxis__000000_Initial.Stroke = ``
	__VerticalAxis__000000_Initial.StrokeOpacity = 0.000000
	__VerticalAxis__000000_Initial.StrokeWidth = 1.000000
	__VerticalAxis__000000_Initial.StrokeDashArray = ``
	__VerticalAxis__000000_Initial.StrokeDashArrayWhenSelected = ``
	__VerticalAxis__000000_Initial.Transform = ``

	// Setup of pointers
	__CircleGrid__000000_Initial.Reference = __Circle__000000_Initial_Circle
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000000_Initial
	__Parameter__000000_Reference.InitialCircle = __Circle__000000_Initial_Circle
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000000_Initial
	__Parameter__000000_Reference.InitialCircleGrid = __CircleGrid__000000_Initial
	__Parameter__000000_Reference.InitialAxis = __Axis__000001_Rotated
	__Parameter__000000_Reference.RotatedAxis = __Axis__000000_Reference
	__Parameter__000000_Reference.RotatedRhombus = __Rhombus__000002_Rotated
	__Parameter__000000_Reference.RotatedRhombusGrid = __RhombusGrid__000001_Rotated
	__Parameter__000000_Reference.RotatedCircleGrid = __CircleGrid__000001_Rotated
	__Parameter__000000_Reference.NextRhombus = __Rhombus__000001_Next_Rhombus
	__Parameter__000000_Reference.NextCircle = __Circle__000001_Next_Circle
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Initial
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Initial
	__RhombusGrid__000000_Initial.Reference = __Rhombus__000000_Initial
}
