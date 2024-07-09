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

	__Circle__000000_0_0 := (&models.Circle{Name: `0 0`}).Stage(stage)
	__Circle__000001_0_1 := (&models.Circle{Name: `0 1`}).Stage(stage)
	__Circle__000002_0_2 := (&models.Circle{Name: `0 2`}).Stage(stage)
	__Circle__000003_1_0 := (&models.Circle{Name: `1 0`}).Stage(stage)
	__Circle__000004_1_1 := (&models.Circle{Name: `1 1`}).Stage(stage)
	__Circle__000005_1_2 := (&models.Circle{Name: `1 2`}).Stage(stage)
	__Circle__000006_Initial_Circle := (&models.Circle{Name: `Initial Circle`}).Stage(stage)

	__CircleGrid__000000_Initial := (&models.CircleGrid{Name: `Initial`}).Stage(stage)

	__HorizontalAxis__000000_Initial := (&models.HorizontalAxis{Name: `Initial`}).Stage(stage)

	__InitialAxis__000000_Reference := (&models.InitialAxis{Name: `Reference`}).Stage(stage)

	__Parameter__000000_Reference := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__000000_0_0 := (&models.Rhombus{Name: `0 0`}).Stage(stage)
	__Rhombus__000001_0_1 := (&models.Rhombus{Name: `0 1`}).Stage(stage)
	__Rhombus__000002_0_2 := (&models.Rhombus{Name: `0 2`}).Stage(stage)
	__Rhombus__000003_1_0 := (&models.Rhombus{Name: `1 0`}).Stage(stage)
	__Rhombus__000004_1_1 := (&models.Rhombus{Name: `1 1`}).Stage(stage)
	__Rhombus__000005_1_2 := (&models.Rhombus{Name: `1 2`}).Stage(stage)
	__Rhombus__000006_Initial := (&models.Rhombus{Name: `Initial`}).Stage(stage)

	__RhombusGrid__000000_Initial := (&models.RhombusGrid{Name: `Initial`}).Stage(stage)

	__VerticalAxis__000000_Initial := (&models.VerticalAxis{Name: `Initial`}).Stage(stage)

	// Setup of values

	__Circle__000000_0_0.Name = `0 0`
	__Circle__000000_0_0.IsDisplayed = true
	__Circle__000000_0_0.CenterX = 0.000000
	__Circle__000000_0_0.CenterY = 0.000000
	__Circle__000000_0_0.Color = ``
	__Circle__000000_0_0.FillOpacity = 0.000000
	__Circle__000000_0_0.Stroke = `black`
	__Circle__000000_0_0.StrokeOpacity = 0.500000
	__Circle__000000_0_0.StrokeWidth = 1.000000
	__Circle__000000_0_0.StrokeDashArray = ``
	__Circle__000000_0_0.StrokeDashArrayWhenSelected = ``
	__Circle__000000_0_0.Transform = ``

	__Circle__000001_0_1.Name = `0 1`
	__Circle__000001_0_1.IsDisplayed = true
	__Circle__000001_0_1.CenterX = -66.913061
	__Circle__000001_0_1.CenterY = 74.314483
	__Circle__000001_0_1.Color = ``
	__Circle__000001_0_1.FillOpacity = 0.000000
	__Circle__000001_0_1.Stroke = `black`
	__Circle__000001_0_1.StrokeOpacity = 0.500000
	__Circle__000001_0_1.StrokeWidth = 1.000000
	__Circle__000001_0_1.StrokeDashArray = ``
	__Circle__000001_0_1.StrokeDashArrayWhenSelected = ``
	__Circle__000001_0_1.Transform = ``

	__Circle__000002_0_2.Name = `0 2`
	__Circle__000002_0_2.IsDisplayed = true
	__Circle__000002_0_2.CenterX = -133.826121
	__Circle__000002_0_2.CenterY = 148.628965
	__Circle__000002_0_2.Color = ``
	__Circle__000002_0_2.FillOpacity = 0.000000
	__Circle__000002_0_2.Stroke = `black`
	__Circle__000002_0_2.StrokeOpacity = 0.500000
	__Circle__000002_0_2.StrokeWidth = 1.000000
	__Circle__000002_0_2.StrokeDashArray = ``
	__Circle__000002_0_2.StrokeDashArrayWhenSelected = ``
	__Circle__000002_0_2.Transform = ``

	__Circle__000003_1_0.Name = `1 0`
	__Circle__000003_1_0.IsDisplayed = true
	__Circle__000003_1_0.CenterX = 66.913061
	__Circle__000003_1_0.CenterY = 74.314483
	__Circle__000003_1_0.Color = ``
	__Circle__000003_1_0.FillOpacity = 0.000000
	__Circle__000003_1_0.Stroke = `black`
	__Circle__000003_1_0.StrokeOpacity = 0.500000
	__Circle__000003_1_0.StrokeWidth = 1.000000
	__Circle__000003_1_0.StrokeDashArray = ``
	__Circle__000003_1_0.StrokeDashArrayWhenSelected = ``
	__Circle__000003_1_0.Transform = ``

	__Circle__000004_1_1.Name = `1 1`
	__Circle__000004_1_1.IsDisplayed = true
	__Circle__000004_1_1.CenterX = 0.000000
	__Circle__000004_1_1.CenterY = 148.628965
	__Circle__000004_1_1.Color = ``
	__Circle__000004_1_1.FillOpacity = 0.000000
	__Circle__000004_1_1.Stroke = `black`
	__Circle__000004_1_1.StrokeOpacity = 0.500000
	__Circle__000004_1_1.StrokeWidth = 1.000000
	__Circle__000004_1_1.StrokeDashArray = ``
	__Circle__000004_1_1.StrokeDashArrayWhenSelected = ``
	__Circle__000004_1_1.Transform = ``

	__Circle__000005_1_2.Name = `1 2`
	__Circle__000005_1_2.IsDisplayed = true
	__Circle__000005_1_2.CenterX = -66.913061
	__Circle__000005_1_2.CenterY = 222.943448
	__Circle__000005_1_2.Color = ``
	__Circle__000005_1_2.FillOpacity = 0.000000
	__Circle__000005_1_2.Stroke = `black`
	__Circle__000005_1_2.StrokeOpacity = 0.500000
	__Circle__000005_1_2.StrokeWidth = 1.000000
	__Circle__000005_1_2.StrokeDashArray = ``
	__Circle__000005_1_2.StrokeDashArrayWhenSelected = ``
	__Circle__000005_1_2.Transform = ``

	__Circle__000006_Initial_Circle.Name = `Initial Circle`
	__Circle__000006_Initial_Circle.IsDisplayed = true
	__Circle__000006_Initial_Circle.CenterX = 0.000000
	__Circle__000006_Initial_Circle.CenterY = 0.000000
	__Circle__000006_Initial_Circle.Color = ``
	__Circle__000006_Initial_Circle.FillOpacity = 0.000000
	__Circle__000006_Initial_Circle.Stroke = `black`
	__Circle__000006_Initial_Circle.StrokeOpacity = 0.500000
	__Circle__000006_Initial_Circle.StrokeWidth = 1.000000
	__Circle__000006_Initial_Circle.StrokeDashArray = ``
	__Circle__000006_Initial_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000006_Initial_Circle.Transform = ``

	__CircleGrid__000000_Initial.Name = `Initial`
	__CircleGrid__000000_Initial.N = 2
	__CircleGrid__000000_Initial.M = 3
	__CircleGrid__000000_Initial.IsDisplayed = true

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

	__InitialAxis__000000_Reference.Name = `Reference`
	__InitialAxis__000000_Reference.IsDisplayed = true
	__InitialAxis__000000_Reference.Angle = 106.706323
	__InitialAxis__000000_Reference.Length = 232.768423
	__InitialAxis__000000_Reference.Color = ``
	__InitialAxis__000000_Reference.FillOpacity = 0.000000
	__InitialAxis__000000_Reference.Stroke = ``
	__InitialAxis__000000_Reference.StrokeOpacity = 0.000000
	__InitialAxis__000000_Reference.StrokeWidth = 2.000000
	__InitialAxis__000000_Reference.StrokeDashArray = ``
	__InitialAxis__000000_Reference.StrokeDashArrayWhenSelected = ``
	__InitialAxis__000000_Reference.Transform = ``

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.N = 2
	__Parameter__000000_Reference.M = 3
	__Parameter__000000_Reference.InsideAngle = 96.000000
	__Parameter__000000_Reference.SideLength = 100.000000
	__Parameter__000000_Reference.OriginX = 300.000000
	__Parameter__000000_Reference.OriginY = 700.000000

	__Rhombus__000000_0_0.Name = `0 0`
	__Rhombus__000000_0_0.IsDisplayed = true
	__Rhombus__000000_0_0.CenterX = 0.000000
	__Rhombus__000000_0_0.CenterY = 0.000000
	__Rhombus__000000_0_0.SideLength = 100.000000
	__Rhombus__000000_0_0.Angle = 0.000000
	__Rhombus__000000_0_0.InsideAngle = 96.000000
	__Rhombus__000000_0_0.Color = ``
	__Rhombus__000000_0_0.FillOpacity = 0.000000
	__Rhombus__000000_0_0.Stroke = ``
	__Rhombus__000000_0_0.StrokeOpacity = 0.000000
	__Rhombus__000000_0_0.StrokeWidth = 1.000000
	__Rhombus__000000_0_0.StrokeDashArray = ``
	__Rhombus__000000_0_0.StrokeDashArrayWhenSelected = ``
	__Rhombus__000000_0_0.Transform = ``

	__Rhombus__000001_0_1.Name = `0 1`
	__Rhombus__000001_0_1.IsDisplayed = true
	__Rhombus__000001_0_1.CenterX = -66.913061
	__Rhombus__000001_0_1.CenterY = 74.314483
	__Rhombus__000001_0_1.SideLength = 100.000000
	__Rhombus__000001_0_1.Angle = 0.000000
	__Rhombus__000001_0_1.InsideAngle = 96.000000
	__Rhombus__000001_0_1.Color = ``
	__Rhombus__000001_0_1.FillOpacity = 0.000000
	__Rhombus__000001_0_1.Stroke = ``
	__Rhombus__000001_0_1.StrokeOpacity = 0.000000
	__Rhombus__000001_0_1.StrokeWidth = 1.000000
	__Rhombus__000001_0_1.StrokeDashArray = ``
	__Rhombus__000001_0_1.StrokeDashArrayWhenSelected = ``
	__Rhombus__000001_0_1.Transform = ``

	__Rhombus__000002_0_2.Name = `0 2`
	__Rhombus__000002_0_2.IsDisplayed = true
	__Rhombus__000002_0_2.CenterX = -133.826121
	__Rhombus__000002_0_2.CenterY = 148.628965
	__Rhombus__000002_0_2.SideLength = 100.000000
	__Rhombus__000002_0_2.Angle = 0.000000
	__Rhombus__000002_0_2.InsideAngle = 96.000000
	__Rhombus__000002_0_2.Color = ``
	__Rhombus__000002_0_2.FillOpacity = 0.000000
	__Rhombus__000002_0_2.Stroke = ``
	__Rhombus__000002_0_2.StrokeOpacity = 0.000000
	__Rhombus__000002_0_2.StrokeWidth = 1.000000
	__Rhombus__000002_0_2.StrokeDashArray = ``
	__Rhombus__000002_0_2.StrokeDashArrayWhenSelected = ``
	__Rhombus__000002_0_2.Transform = ``

	__Rhombus__000003_1_0.Name = `1 0`
	__Rhombus__000003_1_0.IsDisplayed = true
	__Rhombus__000003_1_0.CenterX = 66.913061
	__Rhombus__000003_1_0.CenterY = 74.314483
	__Rhombus__000003_1_0.SideLength = 100.000000
	__Rhombus__000003_1_0.Angle = 0.000000
	__Rhombus__000003_1_0.InsideAngle = 96.000000
	__Rhombus__000003_1_0.Color = ``
	__Rhombus__000003_1_0.FillOpacity = 0.000000
	__Rhombus__000003_1_0.Stroke = ``
	__Rhombus__000003_1_0.StrokeOpacity = 0.000000
	__Rhombus__000003_1_0.StrokeWidth = 1.000000
	__Rhombus__000003_1_0.StrokeDashArray = ``
	__Rhombus__000003_1_0.StrokeDashArrayWhenSelected = ``
	__Rhombus__000003_1_0.Transform = ``

	__Rhombus__000004_1_1.Name = `1 1`
	__Rhombus__000004_1_1.IsDisplayed = true
	__Rhombus__000004_1_1.CenterX = 0.000000
	__Rhombus__000004_1_1.CenterY = 148.628965
	__Rhombus__000004_1_1.SideLength = 100.000000
	__Rhombus__000004_1_1.Angle = 0.000000
	__Rhombus__000004_1_1.InsideAngle = 96.000000
	__Rhombus__000004_1_1.Color = ``
	__Rhombus__000004_1_1.FillOpacity = 0.000000
	__Rhombus__000004_1_1.Stroke = ``
	__Rhombus__000004_1_1.StrokeOpacity = 0.000000
	__Rhombus__000004_1_1.StrokeWidth = 1.000000
	__Rhombus__000004_1_1.StrokeDashArray = ``
	__Rhombus__000004_1_1.StrokeDashArrayWhenSelected = ``
	__Rhombus__000004_1_1.Transform = ``

	__Rhombus__000005_1_2.Name = `1 2`
	__Rhombus__000005_1_2.IsDisplayed = true
	__Rhombus__000005_1_2.CenterX = -66.913061
	__Rhombus__000005_1_2.CenterY = 222.943448
	__Rhombus__000005_1_2.SideLength = 100.000000
	__Rhombus__000005_1_2.Angle = 0.000000
	__Rhombus__000005_1_2.InsideAngle = 96.000000
	__Rhombus__000005_1_2.Color = ``
	__Rhombus__000005_1_2.FillOpacity = 0.000000
	__Rhombus__000005_1_2.Stroke = ``
	__Rhombus__000005_1_2.StrokeOpacity = 0.000000
	__Rhombus__000005_1_2.StrokeWidth = 1.000000
	__Rhombus__000005_1_2.StrokeDashArray = ``
	__Rhombus__000005_1_2.StrokeDashArrayWhenSelected = ``
	__Rhombus__000005_1_2.Transform = ``

	__Rhombus__000006_Initial.Name = `Initial`
	__Rhombus__000006_Initial.IsDisplayed = true
	__Rhombus__000006_Initial.CenterX = 0.000000
	__Rhombus__000006_Initial.CenterY = 0.000000
	__Rhombus__000006_Initial.SideLength = 100.000000
	__Rhombus__000006_Initial.Angle = 0.000000
	__Rhombus__000006_Initial.InsideAngle = 96.000000
	__Rhombus__000006_Initial.Color = ``
	__Rhombus__000006_Initial.FillOpacity = 0.000000
	__Rhombus__000006_Initial.Stroke = ``
	__Rhombus__000006_Initial.StrokeOpacity = 0.000000
	__Rhombus__000006_Initial.StrokeWidth = 1.000000
	__Rhombus__000006_Initial.StrokeDashArray = ``
	__Rhombus__000006_Initial.StrokeDashArrayWhenSelected = ``
	__Rhombus__000006_Initial.Transform = ``

	__RhombusGrid__000000_Initial.Name = `Initial`
	__RhombusGrid__000000_Initial.N = 2
	__RhombusGrid__000000_Initial.M = 3
	__RhombusGrid__000000_Initial.IsDisplayed = true

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
	__CircleGrid__000000_Initial.Reference = __Circle__000006_Initial_Circle
	__CircleGrid__000000_Initial.Circles = append(__CircleGrid__000000_Initial.Circles, __Circle__000000_0_0)
	__CircleGrid__000000_Initial.Circles = append(__CircleGrid__000000_Initial.Circles, __Circle__000001_0_1)
	__CircleGrid__000000_Initial.Circles = append(__CircleGrid__000000_Initial.Circles, __Circle__000002_0_2)
	__CircleGrid__000000_Initial.Circles = append(__CircleGrid__000000_Initial.Circles, __Circle__000003_1_0)
	__CircleGrid__000000_Initial.Circles = append(__CircleGrid__000000_Initial.Circles, __Circle__000004_1_1)
	__CircleGrid__000000_Initial.Circles = append(__CircleGrid__000000_Initial.Circles, __Circle__000005_1_2)
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000006_Initial
	__Parameter__000000_Reference.InitialCircle = __Circle__000006_Initial_Circle
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000000_Initial
	__Parameter__000000_Reference.InitialCircleGrid = __CircleGrid__000000_Initial
	__Parameter__000000_Reference.InitialAxis = __InitialAxis__000000_Reference
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Initial
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Initial
	__RhombusGrid__000000_Initial.Reference = __Rhombus__000006_Initial
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000000_0_0)
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000001_0_1)
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000002_0_2)
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000003_1_0)
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000004_1_1)
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000005_1_2)
}
