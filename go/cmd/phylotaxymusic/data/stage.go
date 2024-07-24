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

	__Circle__000000_Growing := (&models.Circle{Name: `Growing`}).Stage(stage)
	__Circle__000001_Growing_Seed_Left := (&models.Circle{Name: `Growing Seed Left`}).Stage(stage)
	__Circle__000002_Initial_Circle := (&models.Circle{Name: `Initial Circle`}).Stage(stage)
	__Circle__000003_Next_Circle := (&models.Circle{Name: `Next Circle`}).Stage(stage)

	__CircleGrid__000000_Growing := (&models.CircleGrid{Name: `Growing`}).Stage(stage)
	__CircleGrid__000001_Growing_Left := (&models.CircleGrid{Name: `Growing Left`}).Stage(stage)
	__CircleGrid__000002_Initial := (&models.CircleGrid{Name: `Initial`}).Stage(stage)
	__CircleGrid__000003_Rotated := (&models.CircleGrid{Name: `Rotated`}).Stage(stage)

	__HorizontalAxis__000000_Initial := (&models.HorizontalAxis{Name: `Initial`}).Stage(stage)

	__Parameter__000000_Reference := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__000000_Growing_Rhombus_Grid_Seed := (&models.Rhombus{Name: `Growing Rhombus Grid Seed`}).Stage(stage)
	__Rhombus__000001_Initial := (&models.Rhombus{Name: `Initial`}).Stage(stage)
	__Rhombus__000002_Next_Rhombus := (&models.Rhombus{Name: `Next Rhombus`}).Stage(stage)
	__Rhombus__000003_Rotated := (&models.Rhombus{Name: `Rotated`}).Stage(stage)

	__RhombusGrid__000000_Growing_Rhombus_Grid := (&models.RhombusGrid{Name: `Growing Rhombus Grid`}).Stage(stage)
	__RhombusGrid__000001_Initial := (&models.RhombusGrid{Name: `Initial`}).Stage(stage)
	__RhombusGrid__000002_Rotated := (&models.RhombusGrid{Name: `Rotated`}).Stage(stage)

	__VerticalAxis__000000_Initial := (&models.VerticalAxis{Name: `Initial`}).Stage(stage)

	// Setup of values

	__Axis__000000_Reference.Name = `Reference`
	__Axis__000000_Reference.IsDisplayed = false
	__Axis__000000_Reference.Angle = 0.000000
	__Axis__000000_Reference.Length = 1014.857807
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
	__Axis__000001_Rotated.Angle = 82.599764
	__Axis__000001_Rotated.Length = 1014.857807
	__Axis__000001_Rotated.Color = ``
	__Axis__000001_Rotated.FillOpacity = 0.000000
	__Axis__000001_Rotated.Stroke = `black`
	__Axis__000001_Rotated.StrokeOpacity = 1.000000
	__Axis__000001_Rotated.StrokeWidth = 2.000000
	__Axis__000001_Rotated.StrokeDashArray = ``
	__Axis__000001_Rotated.StrokeDashArrayWhenSelected = ``
	__Axis__000001_Rotated.Transform = ``

	__Circle__000000_Growing.Name = `Growing`
	__Circle__000000_Growing.IsDisplayed = false
	__Circle__000000_Growing.CenterX = 0.000000
	__Circle__000000_Growing.CenterY = 0.000000
	__Circle__000000_Growing.Color = ``
	__Circle__000000_Growing.FillOpacity = 0.000000
	__Circle__000000_Growing.Stroke = `blue`
	__Circle__000000_Growing.StrokeOpacity = 0.700000
	__Circle__000000_Growing.StrokeWidth = 2.000000
	__Circle__000000_Growing.StrokeDashArray = ``
	__Circle__000000_Growing.StrokeDashArrayWhenSelected = ``
	__Circle__000000_Growing.Transform = ``

	__Circle__000001_Growing_Seed_Left.Name = `Growing Seed Left`
	__Circle__000001_Growing_Seed_Left.IsDisplayed = false
	__Circle__000001_Growing_Seed_Left.CenterX = 0.000000
	__Circle__000001_Growing_Seed_Left.CenterY = 0.000000
	__Circle__000001_Growing_Seed_Left.Color = ``
	__Circle__000001_Growing_Seed_Left.FillOpacity = 0.000000
	__Circle__000001_Growing_Seed_Left.Stroke = `green`
	__Circle__000001_Growing_Seed_Left.StrokeOpacity = 0.800000
	__Circle__000001_Growing_Seed_Left.StrokeWidth = 1.800000
	__Circle__000001_Growing_Seed_Left.StrokeDashArray = ``
	__Circle__000001_Growing_Seed_Left.StrokeDashArrayWhenSelected = ``
	__Circle__000001_Growing_Seed_Left.Transform = ``

	__Circle__000002_Initial_Circle.Name = `Initial Circle`
	__Circle__000002_Initial_Circle.IsDisplayed = false
	__Circle__000002_Initial_Circle.CenterX = 0.000000
	__Circle__000002_Initial_Circle.CenterY = 0.000000
	__Circle__000002_Initial_Circle.Color = ``
	__Circle__000002_Initial_Circle.FillOpacity = 0.000000
	__Circle__000002_Initial_Circle.Stroke = `lightblue`
	__Circle__000002_Initial_Circle.StrokeOpacity = 0.800000
	__Circle__000002_Initial_Circle.StrokeWidth = 3.000000
	__Circle__000002_Initial_Circle.StrokeDashArray = ``
	__Circle__000002_Initial_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000002_Initial_Circle.Transform = ``

	__Circle__000003_Next_Circle.Name = `Next Circle`
	__Circle__000003_Next_Circle.IsDisplayed = false
	__Circle__000003_Next_Circle.CenterX = 399.208786
	__Circle__000003_Next_Circle.CenterY = 51.849843
	__Circle__000003_Next_Circle.Color = ``
	__Circle__000003_Next_Circle.FillOpacity = 0.000000
	__Circle__000003_Next_Circle.Stroke = `yellow`
	__Circle__000003_Next_Circle.StrokeOpacity = 0.800000
	__Circle__000003_Next_Circle.StrokeWidth = 3.000000
	__Circle__000003_Next_Circle.StrokeDashArray = ``
	__Circle__000003_Next_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000003_Next_Circle.Transform = ``

	__CircleGrid__000000_Growing.Name = `Growing`
	__CircleGrid__000000_Growing.IsDisplayed = true

	__CircleGrid__000001_Growing_Left.Name = `Growing Left`
	__CircleGrid__000001_Growing_Left.IsDisplayed = true

	__CircleGrid__000002_Initial.Name = `Initial`
	__CircleGrid__000002_Initial.IsDisplayed = false

	__CircleGrid__000003_Rotated.Name = `Rotated`
	__CircleGrid__000003_Rotated.IsDisplayed = true

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
	__Parameter__000000_Reference.Z = 12
	__Parameter__000000_Reference.InsideAngle = 96.000000
	__Parameter__000000_Reference.SideLength = 220.000000
	__Parameter__000000_Reference.OriginX = 300.000000
	__Parameter__000000_Reference.OriginY = 600.000000

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 240.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Angle = -82.599764
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 114.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Color = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.FillOpacity = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Stroke = `green`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeOpacity = 0.800000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeWidth = 2.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeDashArray = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeDashArrayWhenSelected = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Transform = ``

	__Rhombus__000001_Initial.Name = `Initial`
	__Rhombus__000001_Initial.IsDisplayed = false
	__Rhombus__000001_Initial.CenterX = 0.000000
	__Rhombus__000001_Initial.CenterY = 0.000000
	__Rhombus__000001_Initial.SideLength = 240.000000
	__Rhombus__000001_Initial.Angle = 0.000000
	__Rhombus__000001_Initial.InsideAngle = 114.000000
	__Rhombus__000001_Initial.Color = ``
	__Rhombus__000001_Initial.FillOpacity = 0.000000
	__Rhombus__000001_Initial.Stroke = `green`
	__Rhombus__000001_Initial.StrokeOpacity = 0.900000
	__Rhombus__000001_Initial.StrokeWidth = 3.000000
	__Rhombus__000001_Initial.StrokeDashArray = ``
	__Rhombus__000001_Initial.StrokeDashArrayWhenSelected = ``
	__Rhombus__000001_Initial.Transform = ``

	__Rhombus__000002_Next_Rhombus.Name = `Next Rhombus`
	__Rhombus__000002_Next_Rhombus.IsDisplayed = false
	__Rhombus__000002_Next_Rhombus.CenterX = 399.208786
	__Rhombus__000002_Next_Rhombus.CenterY = 51.849843
	__Rhombus__000002_Next_Rhombus.SideLength = 240.000000
	__Rhombus__000002_Next_Rhombus.Angle = -82.599764
	__Rhombus__000002_Next_Rhombus.InsideAngle = 114.000000
	__Rhombus__000002_Next_Rhombus.Color = ``
	__Rhombus__000002_Next_Rhombus.FillOpacity = 0.000000
	__Rhombus__000002_Next_Rhombus.Stroke = `lavender`
	__Rhombus__000002_Next_Rhombus.StrokeOpacity = 0.700000
	__Rhombus__000002_Next_Rhombus.StrokeWidth = 6.000000
	__Rhombus__000002_Next_Rhombus.StrokeDashArray = ``
	__Rhombus__000002_Next_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000002_Next_Rhombus.Transform = ``

	__Rhombus__000003_Rotated.Name = `Rotated`
	__Rhombus__000003_Rotated.IsDisplayed = false
	__Rhombus__000003_Rotated.CenterX = 0.000000
	__Rhombus__000003_Rotated.CenterY = 0.000000
	__Rhombus__000003_Rotated.SideLength = 240.000000
	__Rhombus__000003_Rotated.Angle = -82.599764
	__Rhombus__000003_Rotated.InsideAngle = 114.000000
	__Rhombus__000003_Rotated.Color = ``
	__Rhombus__000003_Rotated.FillOpacity = 0.000000
	__Rhombus__000003_Rotated.Stroke = `black`
	__Rhombus__000003_Rotated.StrokeOpacity = 1.000000
	__Rhombus__000003_Rotated.StrokeWidth = 2.000000
	__Rhombus__000003_Rotated.StrokeDashArray = ``
	__Rhombus__000003_Rotated.StrokeDashArrayWhenSelected = ``
	__Rhombus__000003_Rotated.Transform = ``

	__RhombusGrid__000000_Growing_Rhombus_Grid.Name = `Growing Rhombus Grid`
	__RhombusGrid__000000_Growing_Rhombus_Grid.IsDisplayed = false

	__RhombusGrid__000001_Initial.Name = `Initial`
	__RhombusGrid__000001_Initial.IsDisplayed = false

	__RhombusGrid__000002_Rotated.Name = `Rotated`
	__RhombusGrid__000002_Rotated.IsDisplayed = false

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
	__CircleGrid__000000_Growing.Reference = __Circle__000000_Growing
	__CircleGrid__000002_Initial.Reference = __Circle__000002_Initial_Circle
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000001_Initial
	__Parameter__000000_Reference.InitialCircle = __Circle__000002_Initial_Circle
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000001_Initial
	__Parameter__000000_Reference.InitialCircleGrid = __CircleGrid__000002_Initial
	__Parameter__000000_Reference.InitialAxis = __Axis__000001_Rotated
	__Parameter__000000_Reference.RotatedAxis = __Axis__000000_Reference
	__Parameter__000000_Reference.RotatedRhombus = __Rhombus__000003_Rotated
	__Parameter__000000_Reference.RotatedRhombusGrid = __RhombusGrid__000002_Rotated
	__Parameter__000000_Reference.RotatedCircleGrid = __CircleGrid__000003_Rotated
	__Parameter__000000_Reference.NextRhombus = __Rhombus__000002_Next_Rhombus
	__Parameter__000000_Reference.NextCircle = __Circle__000003_Next_Circle
	__Parameter__000000_Reference.GrowingRhombusGridSeed = __Rhombus__000000_Growing_Rhombus_Grid_Seed
	__Parameter__000000_Reference.GrowingRhombusGrid = __RhombusGrid__000000_Growing_Rhombus_Grid
	__Parameter__000000_Reference.GrowingCircleGridSeed = __Circle__000000_Growing
	__Parameter__000000_Reference.GrowingCircleGrid = __CircleGrid__000000_Growing
	__Parameter__000000_Reference.GrowingCircleGridLeftSeed = __Circle__000001_Growing_Seed_Left
	__Parameter__000000_Reference.GrowingCircleGridLeft = __CircleGrid__000001_Growing_Left
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Initial
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Initial
	__RhombusGrid__000000_Growing_Rhombus_Grid.Reference = __Rhombus__000003_Rotated
	__RhombusGrid__000001_Initial.Reference = __Rhombus__000001_Initial
}
