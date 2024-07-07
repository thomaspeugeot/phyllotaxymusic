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

	__HorizontalAxis__000000_Initial := (&models.HorizontalAxis{Name: `Initial`}).Stage(stage)

	__Parameter__000000_Reference := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__000000_0 := (&models.Rhombus{Name: `0`}).Stage(stage)
	__Rhombus__000001_1 := (&models.Rhombus{Name: `1`}).Stage(stage)
	__Rhombus__000002_Initial := (&models.Rhombus{Name: `Initial`}).Stage(stage)

	__RhombusGrid__000000_Initial := (&models.RhombusGrid{Name: `Initial`}).Stage(stage)

	__VerticalAxis__000000_Initial := (&models.VerticalAxis{Name: `Initial`}).Stage(stage)

	// Setup of values

	__HorizontalAxis__000000_Initial.Name = `Initial`
	__HorizontalAxis__000000_Initial.IsDisplayed = true
	__HorizontalAxis__000000_Initial.AxisHandleBorderLength = 0.000000
	__HorizontalAxis__000000_Initial.Axis_Length = 600.000000
	__HorizontalAxis__000000_Initial.StrokeWidth = 1.000000

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Angle = 90.000000
	__Parameter__000000_Reference.DiamondSideLenght = 100.000000
	__Parameter__000000_Reference.OriginX = 100.000000
	__Parameter__000000_Reference.OriginY = 700.000000

	__Rhombus__000000_0.Name = `0`
	__Rhombus__000000_0.IsDisplayed = true
	__Rhombus__000000_0.CenterX = 0.000000
	__Rhombus__000000_0.CenterY = 0.000000
	__Rhombus__000000_0.SideLength = 100.000000
	__Rhombus__000000_0.Angle = 0.000000
	__Rhombus__000000_0.InsideAngle = 73.740000
	__Rhombus__000000_0.StrokeWidth = 1.000000

	__Rhombus__000001_1.Name = `1`
	__Rhombus__000001_1.IsDisplayed = true
	__Rhombus__000001_1.CenterX = 79.999893
	__Rhombus__000001_1.CenterY = 60.000143
	__Rhombus__000001_1.SideLength = 100.000000
	__Rhombus__000001_1.Angle = 0.000000
	__Rhombus__000001_1.InsideAngle = 73.740000
	__Rhombus__000001_1.StrokeWidth = 1.000000

	__Rhombus__000002_Initial.Name = `Initial`
	__Rhombus__000002_Initial.IsDisplayed = false
	__Rhombus__000002_Initial.CenterX = 0.000000
	__Rhombus__000002_Initial.CenterY = 0.000000
	__Rhombus__000002_Initial.SideLength = 100.000000
	__Rhombus__000002_Initial.Angle = 0.000000
	__Rhombus__000002_Initial.InsideAngle = 73.740000
	__Rhombus__000002_Initial.StrokeWidth = 1.000000

	__RhombusGrid__000000_Initial.Name = `Initial`
	__RhombusGrid__000000_Initial.N = 2
	__RhombusGrid__000000_Initial.M = 3
	__RhombusGrid__000000_Initial.IsDisplayed = true

	__VerticalAxis__000000_Initial.Name = `Initial`
	__VerticalAxis__000000_Initial.IsDisplayed = true
	__VerticalAxis__000000_Initial.AxisHandleBorderLength = 0.000000
	__VerticalAxis__000000_Initial.Axis_Length = 600.000000
	__VerticalAxis__000000_Initial.StrokeWidth = 1.000000

	// Setup of pointers
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000002_Initial
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000000_Initial
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Initial
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Initial
	__RhombusGrid__000000_Initial.Reference = __Rhombus__000002_Initial
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000000_0)
	__RhombusGrid__000000_Initial.Rhombuses = append(__RhombusGrid__000000_Initial.Rhombuses, __Rhombus__000001_1)
}
