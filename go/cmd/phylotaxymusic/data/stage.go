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

	__Axis__000000_Construction_Axis := (&models.Axis{Name: `Construction Axis`}).Stage(stage)
	__Axis__000001_Reference_Axis := (&models.Axis{Name: `Reference Axis`}).Stage(stage)
	__Axis__000002_Rotated_Axis := (&models.Axis{Name: `Rotated Axis`}).Stage(stage)

	__AxisGrid__000000_Construction_Axis_Grid := (&models.AxisGrid{Name: `Construction Axis Grid`}).Stage(stage)

	__Circle__000000_Construction_Circle := (&models.Circle{Name: `Construction Circle`}).Stage(stage)
	__Circle__000001_Growing := (&models.Circle{Name: `Growing`}).Stage(stage)
	__Circle__000002_Growing_Seed_Left := (&models.Circle{Name: `Growing Seed Left`}).Stage(stage)
	__Circle__000003_Initial_Circle := (&models.Circle{Name: `Initial Circle`}).Stage(stage)
	__Circle__000004_Next_Circle := (&models.Circle{Name: `Next Circle`}).Stage(stage)

	__CircleGrid__000000_Growing_Circle_Grid := (&models.CircleGrid{Name: `Growing Circle Grid`}).Stage(stage)
	__CircleGrid__000001_Growing_Circle_Grid_Shifted_Left := (&models.CircleGrid{Name: `Growing Circle Grid Shifted Left`}).Stage(stage)
	__CircleGrid__000002_Initial_Circle_Grid := (&models.CircleGrid{Name: `Initial Circle Grid`}).Stage(stage)
	__CircleGrid__000003_Rotated_Circle_Grid := (&models.CircleGrid{Name: `Rotated Circle Grid`}).Stage(stage)

	__HorizontalAxis__000000_Horizontal_Axis := (&models.HorizontalAxis{Name: `Horizontal Axis`}).Stage(stage)

	__Parameter__000000_Reference := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__000000_Growing_Rhombus_Grid_Seed := (&models.Rhombus{Name: `Growing Rhombus Grid Seed`}).Stage(stage)
	__Rhombus__000001_Initial_Rhombus := (&models.Rhombus{Name: `Initial Rhombus`}).Stage(stage)
	__Rhombus__000002_Next_Rhombus := (&models.Rhombus{Name: `Next Rhombus`}).Stage(stage)
	__Rhombus__000003_Rotated_Rhombus := (&models.Rhombus{Name: `Rotated Rhombus`}).Stage(stage)

	__RhombusGrid__000000_Growing_Rhombus_Grid := (&models.RhombusGrid{Name: `Growing Rhombus Grid`}).Stage(stage)
	__RhombusGrid__000001_Initial_Rhombus_Grid := (&models.RhombusGrid{Name: `Initial Rhombus Grid`}).Stage(stage)
	__RhombusGrid__000002_Rotated_Rhombus_Grid := (&models.RhombusGrid{Name: `Rotated Rhombus Grid`}).Stage(stage)

	__VerticalAxis__000000_Vertical_Axis := (&models.VerticalAxis{Name: `Vertical Axis`}).Stage(stage)

	// Setup of values

	__Axis__000000_Construction_Axis.Name = `Construction Axis`
	__Axis__000000_Construction_Axis.IsDisplayed = true
	__Axis__000000_Construction_Axis.Angle = 100.208458
	__Axis__000000_Construction_Axis.Length = 173.973958
	__Axis__000000_Construction_Axis.CenterX = 0.000000
	__Axis__000000_Construction_Axis.CenterY = 0.000000
	__Axis__000000_Construction_Axis.Color = ``
	__Axis__000000_Construction_Axis.FillOpacity = 0.000000
	__Axis__000000_Construction_Axis.Stroke = `blue`
	__Axis__000000_Construction_Axis.StrokeOpacity = 0.700000
	__Axis__000000_Construction_Axis.StrokeWidth = 2.000000
	__Axis__000000_Construction_Axis.StrokeDashArray = ``
	__Axis__000000_Construction_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000000_Construction_Axis.Transform = ``

	__Axis__000001_Reference_Axis.Name = `Reference Axis`
	__Axis__000001_Reference_Axis.IsDisplayed = false
	__Axis__000001_Reference_Axis.Angle = 79.791542
	__Axis__000001_Reference_Axis.Length = 490.813990
	__Axis__000001_Reference_Axis.CenterX = 0.000000
	__Axis__000001_Reference_Axis.CenterY = 0.000000
	__Axis__000001_Reference_Axis.Color = ``
	__Axis__000001_Reference_Axis.FillOpacity = 0.000000
	__Axis__000001_Reference_Axis.Stroke = `black`
	__Axis__000001_Reference_Axis.StrokeOpacity = 1.000000
	__Axis__000001_Reference_Axis.StrokeWidth = 2.000000
	__Axis__000001_Reference_Axis.StrokeDashArray = ``
	__Axis__000001_Reference_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000001_Reference_Axis.Transform = ``

	__Axis__000002_Rotated_Axis.Name = `Rotated Axis`
	__Axis__000002_Rotated_Axis.IsDisplayed = false
	__Axis__000002_Rotated_Axis.Angle = 0.000000
	__Axis__000002_Rotated_Axis.Length = 490.813990
	__Axis__000002_Rotated_Axis.CenterX = 0.000000
	__Axis__000002_Rotated_Axis.CenterY = 0.000000
	__Axis__000002_Rotated_Axis.Color = ``
	__Axis__000002_Rotated_Axis.FillOpacity = 0.000000
	__Axis__000002_Rotated_Axis.Stroke = `black`
	__Axis__000002_Rotated_Axis.StrokeOpacity = 1.000000
	__Axis__000002_Rotated_Axis.StrokeWidth = 2.000000
	__Axis__000002_Rotated_Axis.StrokeDashArray = ``
	__Axis__000002_Rotated_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000002_Rotated_Axis.Transform = ``

	__AxisGrid__000000_Construction_Axis_Grid.Name = `Construction Axis Grid`
	__AxisGrid__000000_Construction_Axis_Grid.IsDisplayed = false

	__Circle__000000_Construction_Circle.Name = `Construction Circle`
	__Circle__000000_Construction_Circle.IsDisplayed = true
	__Circle__000000_Construction_Circle.CenterX = -15.416705
	__Circle__000000_Construction_Circle.CenterY = 85.609928
	__Circle__000000_Construction_Circle.HasBespokeRadius = true
	__Circle__000000_Construction_Circle.BespopkeRadius = 20.000000
	__Circle__000000_Construction_Circle.Color = ``
	__Circle__000000_Construction_Circle.FillOpacity = 0.000000
	__Circle__000000_Construction_Circle.Stroke = `blue`
	__Circle__000000_Construction_Circle.StrokeOpacity = 0.800000
	__Circle__000000_Construction_Circle.StrokeWidth = 2.000000
	__Circle__000000_Construction_Circle.StrokeDashArray = ``
	__Circle__000000_Construction_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000000_Construction_Circle.Transform = ``

	__Circle__000001_Growing.Name = `Growing`
	__Circle__000001_Growing.IsDisplayed = false
	__Circle__000001_Growing.CenterX = 0.000000
	__Circle__000001_Growing.CenterY = 0.000000
	__Circle__000001_Growing.HasBespokeRadius = false
	__Circle__000001_Growing.BespopkeRadius = 0.000000
	__Circle__000001_Growing.Color = ``
	__Circle__000001_Growing.FillOpacity = 0.000000
	__Circle__000001_Growing.Stroke = `blue`
	__Circle__000001_Growing.StrokeOpacity = 0.700000
	__Circle__000001_Growing.StrokeWidth = 2.000000
	__Circle__000001_Growing.StrokeDashArray = ``
	__Circle__000001_Growing.StrokeDashArrayWhenSelected = ``
	__Circle__000001_Growing.Transform = ``

	__Circle__000002_Growing_Seed_Left.Name = `Growing Seed Left`
	__Circle__000002_Growing_Seed_Left.IsDisplayed = false
	__Circle__000002_Growing_Seed_Left.CenterX = 0.000000
	__Circle__000002_Growing_Seed_Left.CenterY = 0.000000
	__Circle__000002_Growing_Seed_Left.HasBespokeRadius = false
	__Circle__000002_Growing_Seed_Left.BespopkeRadius = 0.000000
	__Circle__000002_Growing_Seed_Left.Color = ``
	__Circle__000002_Growing_Seed_Left.FillOpacity = 0.000000
	__Circle__000002_Growing_Seed_Left.Stroke = `green`
	__Circle__000002_Growing_Seed_Left.StrokeOpacity = 0.800000
	__Circle__000002_Growing_Seed_Left.StrokeWidth = 1.800000
	__Circle__000002_Growing_Seed_Left.StrokeDashArray = ``
	__Circle__000002_Growing_Seed_Left.StrokeDashArrayWhenSelected = ``
	__Circle__000002_Growing_Seed_Left.Transform = ``

	__Circle__000003_Initial_Circle.Name = `Initial Circle`
	__Circle__000003_Initial_Circle.IsDisplayed = false
	__Circle__000003_Initial_Circle.CenterX = 0.000000
	__Circle__000003_Initial_Circle.CenterY = 0.000000
	__Circle__000003_Initial_Circle.HasBespokeRadius = false
	__Circle__000003_Initial_Circle.BespopkeRadius = 0.000000
	__Circle__000003_Initial_Circle.Color = ``
	__Circle__000003_Initial_Circle.FillOpacity = 0.000000
	__Circle__000003_Initial_Circle.Stroke = `lightblue`
	__Circle__000003_Initial_Circle.StrokeOpacity = 0.800000
	__Circle__000003_Initial_Circle.StrokeWidth = 3.000000
	__Circle__000003_Initial_Circle.StrokeDashArray = ``
	__Circle__000003_Initial_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000003_Initial_Circle.Transform = ``

	__Circle__000004_Next_Circle.Name = `Next Circle`
	__Circle__000004_Next_Circle.IsDisplayed = false
	__Circle__000004_Next_Circle.CenterX = 190.158914
	__Circle__000004_Next_Circle.CenterY = 34.243971
	__Circle__000004_Next_Circle.HasBespokeRadius = false
	__Circle__000004_Next_Circle.BespopkeRadius = 0.000000
	__Circle__000004_Next_Circle.Color = ``
	__Circle__000004_Next_Circle.FillOpacity = 0.000000
	__Circle__000004_Next_Circle.Stroke = `yellow`
	__Circle__000004_Next_Circle.StrokeOpacity = 0.800000
	__Circle__000004_Next_Circle.StrokeWidth = 3.000000
	__Circle__000004_Next_Circle.StrokeDashArray = ``
	__Circle__000004_Next_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000004_Next_Circle.Transform = ``

	__CircleGrid__000000_Growing_Circle_Grid.Name = `Growing Circle Grid`
	__CircleGrid__000000_Growing_Circle_Grid.IsDisplayed = false

	__CircleGrid__000001_Growing_Circle_Grid_Shifted_Left.Name = `Growing Circle Grid Shifted Left`
	__CircleGrid__000001_Growing_Circle_Grid_Shifted_Left.IsDisplayed = false

	__CircleGrid__000002_Initial_Circle_Grid.Name = `Initial Circle Grid`
	__CircleGrid__000002_Initial_Circle_Grid.IsDisplayed = false

	__CircleGrid__000003_Rotated_Circle_Grid.Name = `Rotated Circle Grid`
	__CircleGrid__000003_Rotated_Circle_Grid.IsDisplayed = false

	__HorizontalAxis__000000_Horizontal_Axis.Name = `Horizontal Axis`
	__HorizontalAxis__000000_Horizontal_Axis.IsDisplayed = true
	__HorizontalAxis__000000_Horizontal_Axis.AxisHandleBorderLength = 0.000000
	__HorizontalAxis__000000_Horizontal_Axis.Axis_Length = 600.000000
	__HorizontalAxis__000000_Horizontal_Axis.Color = ``
	__HorizontalAxis__000000_Horizontal_Axis.FillOpacity = 0.000000
	__HorizontalAxis__000000_Horizontal_Axis.Stroke = ``
	__HorizontalAxis__000000_Horizontal_Axis.StrokeOpacity = 0.000000
	__HorizontalAxis__000000_Horizontal_Axis.StrokeWidth = 1.000000
	__HorizontalAxis__000000_Horizontal_Axis.StrokeDashArray = ``
	__HorizontalAxis__000000_Horizontal_Axis.StrokeDashArrayWhenSelected = ``
	__HorizontalAxis__000000_Horizontal_Axis.Transform = ``

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 12
	__Parameter__000000_Reference.InsideAngle = 96.000000
	__Parameter__000000_Reference.SideLength = 130.000000
	__Parameter__000000_Reference.OriginX = 300.000000
	__Parameter__000000_Reference.OriginY = 600.000000

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 130.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Angle = -79.791542
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 96.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Color = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.FillOpacity = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Stroke = `green`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeOpacity = 0.800000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeWidth = 2.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeDashArray = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeDashArrayWhenSelected = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Transform = ``

	__Rhombus__000001_Initial_Rhombus.Name = `Initial Rhombus`
	__Rhombus__000001_Initial_Rhombus.IsDisplayed = false
	__Rhombus__000001_Initial_Rhombus.CenterX = 0.000000
	__Rhombus__000001_Initial_Rhombus.CenterY = 0.000000
	__Rhombus__000001_Initial_Rhombus.SideLength = 130.000000
	__Rhombus__000001_Initial_Rhombus.Angle = 0.000000
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 96.000000
	__Rhombus__000001_Initial_Rhombus.Color = ``
	__Rhombus__000001_Initial_Rhombus.FillOpacity = 0.000000
	__Rhombus__000001_Initial_Rhombus.Stroke = `green`
	__Rhombus__000001_Initial_Rhombus.StrokeOpacity = 0.900000
	__Rhombus__000001_Initial_Rhombus.StrokeWidth = 3.000000
	__Rhombus__000001_Initial_Rhombus.StrokeDashArray = ``
	__Rhombus__000001_Initial_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000001_Initial_Rhombus.Transform = ``

	__Rhombus__000002_Next_Rhombus.Name = `Next Rhombus`
	__Rhombus__000002_Next_Rhombus.IsDisplayed = false
	__Rhombus__000002_Next_Rhombus.CenterX = 190.158914
	__Rhombus__000002_Next_Rhombus.CenterY = 34.243971
	__Rhombus__000002_Next_Rhombus.SideLength = 130.000000
	__Rhombus__000002_Next_Rhombus.Angle = -79.791542
	__Rhombus__000002_Next_Rhombus.InsideAngle = 96.000000
	__Rhombus__000002_Next_Rhombus.Color = ``
	__Rhombus__000002_Next_Rhombus.FillOpacity = 0.000000
	__Rhombus__000002_Next_Rhombus.Stroke = `lavender`
	__Rhombus__000002_Next_Rhombus.StrokeOpacity = 0.700000
	__Rhombus__000002_Next_Rhombus.StrokeWidth = 6.000000
	__Rhombus__000002_Next_Rhombus.StrokeDashArray = ``
	__Rhombus__000002_Next_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000002_Next_Rhombus.Transform = ``

	__Rhombus__000003_Rotated_Rhombus.Name = `Rotated Rhombus`
	__Rhombus__000003_Rotated_Rhombus.IsDisplayed = false
	__Rhombus__000003_Rotated_Rhombus.CenterX = 0.000000
	__Rhombus__000003_Rotated_Rhombus.CenterY = 0.000000
	__Rhombus__000003_Rotated_Rhombus.SideLength = 130.000000
	__Rhombus__000003_Rotated_Rhombus.Angle = -79.791542
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 96.000000
	__Rhombus__000003_Rotated_Rhombus.Color = ``
	__Rhombus__000003_Rotated_Rhombus.FillOpacity = 0.000000
	__Rhombus__000003_Rotated_Rhombus.Stroke = `black`
	__Rhombus__000003_Rotated_Rhombus.StrokeOpacity = 1.000000
	__Rhombus__000003_Rotated_Rhombus.StrokeWidth = 2.000000
	__Rhombus__000003_Rotated_Rhombus.StrokeDashArray = ``
	__Rhombus__000003_Rotated_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000003_Rotated_Rhombus.Transform = ``

	__RhombusGrid__000000_Growing_Rhombus_Grid.Name = `Growing Rhombus Grid`
	__RhombusGrid__000000_Growing_Rhombus_Grid.IsDisplayed = false

	__RhombusGrid__000001_Initial_Rhombus_Grid.Name = `Initial Rhombus Grid`
	__RhombusGrid__000001_Initial_Rhombus_Grid.IsDisplayed = false

	__RhombusGrid__000002_Rotated_Rhombus_Grid.Name = `Rotated Rhombus Grid`
	__RhombusGrid__000002_Rotated_Rhombus_Grid.IsDisplayed = false

	__VerticalAxis__000000_Vertical_Axis.Name = `Vertical Axis`
	__VerticalAxis__000000_Vertical_Axis.IsDisplayed = true
	__VerticalAxis__000000_Vertical_Axis.AxisHandleBorderLength = 0.000000
	__VerticalAxis__000000_Vertical_Axis.Axis_Length = 600.000000
	__VerticalAxis__000000_Vertical_Axis.Color = ``
	__VerticalAxis__000000_Vertical_Axis.FillOpacity = 0.000000
	__VerticalAxis__000000_Vertical_Axis.Stroke = ``
	__VerticalAxis__000000_Vertical_Axis.StrokeOpacity = 0.000000
	__VerticalAxis__000000_Vertical_Axis.StrokeWidth = 1.000000
	__VerticalAxis__000000_Vertical_Axis.StrokeDashArray = ``
	__VerticalAxis__000000_Vertical_Axis.StrokeDashArrayWhenSelected = ``
	__VerticalAxis__000000_Vertical_Axis.Transform = ``

	// Setup of pointers
	__AxisGrid__000000_Construction_Axis_Grid.Reference = __Axis__000000_Construction_Axis
	__CircleGrid__000000_Growing_Circle_Grid.Reference = __Circle__000001_Growing
	__CircleGrid__000002_Initial_Circle_Grid.Reference = __Circle__000003_Initial_Circle
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000001_Initial_Rhombus
	__Parameter__000000_Reference.InitialCircle = __Circle__000003_Initial_Circle
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000001_Initial_Rhombus_Grid
	__Parameter__000000_Reference.InitialCircleGrid = __CircleGrid__000002_Initial_Circle_Grid
	__Parameter__000000_Reference.InitialAxis = __Axis__000001_Reference_Axis
	__Parameter__000000_Reference.RotatedAxis = __Axis__000002_Rotated_Axis
	__Parameter__000000_Reference.RotatedRhombus = __Rhombus__000003_Rotated_Rhombus
	__Parameter__000000_Reference.RotatedRhombusGrid = __RhombusGrid__000002_Rotated_Rhombus_Grid
	__Parameter__000000_Reference.RotatedCircleGrid = __CircleGrid__000003_Rotated_Circle_Grid
	__Parameter__000000_Reference.NextRhombus = __Rhombus__000002_Next_Rhombus
	__Parameter__000000_Reference.NextCircle = __Circle__000004_Next_Circle
	__Parameter__000000_Reference.GrowingRhombusGridSeed = __Rhombus__000000_Growing_Rhombus_Grid_Seed
	__Parameter__000000_Reference.GrowingRhombusGrid = __RhombusGrid__000000_Growing_Rhombus_Grid
	__Parameter__000000_Reference.GrowingCircleGridSeed = __Circle__000001_Growing
	__Parameter__000000_Reference.GrowingCircleGrid = __CircleGrid__000000_Growing_Circle_Grid
	__Parameter__000000_Reference.GrowingCircleGridLeftSeed = __Circle__000002_Growing_Seed_Left
	__Parameter__000000_Reference.GrowingCircleGridLeft = __CircleGrid__000001_Growing_Circle_Grid_Shifted_Left
	__Parameter__000000_Reference.ConstructionAxis = __Axis__000000_Construction_Axis
	__Parameter__000000_Reference.ConstructionAxisGrid = __AxisGrid__000000_Construction_Axis_Grid
	__Parameter__000000_Reference.ConstructionCircle = __Circle__000000_Construction_Circle
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Horizontal_Axis
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Vertical_Axis
	__RhombusGrid__000000_Growing_Rhombus_Grid.Reference = __Rhombus__000003_Rotated_Rhombus
	__RhombusGrid__000001_Initial_Rhombus_Grid.Reference = __Rhombus__000001_Initial_Rhombus
}
