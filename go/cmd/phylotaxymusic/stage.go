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

	__HorizontalAxis__000000_Initial.Name = `Initial`
	__HorizontalAxis__000000_Initial.IsDisplayed = true
	__HorizontalAxis__000000_Initial.AxisHandleBorderLength = 0.000000
	__HorizontalAxis__000000_Initial.Axis_Length = 600.000000
	__HorizontalAxis__000000_Initial.StrokeWidth = 1.000000

	__InitialAxis__000000_Reference.Name = `Reference`
	__InitialAxis__000000_Reference.IsDisplayed = true
	__InitialAxis__000000_Reference.Angle = 112.736483
	__InitialAxis__000000_Reference.Length = 202.489402
	__InitialAxis__000000_Reference.StrokeWidth = 2.000000

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.N = 2
	__Parameter__000000_Reference.M = 3
	__Parameter__000000_Reference.InsideAngle = 77.000000
	__Parameter__000000_Reference.DiamondSideLenght = 100.000000
	__Parameter__000000_Reference.OriginX = 300.000000
	__Parameter__000000_Reference.OriginY = 700.000000

	__Rhombus__000000_0_0.Name = `0 0`
	__Rhombus__000000_0_0.IsDisplayed = true
	__Rhombus__000000_0_0.CenterX = 0.000000
	__Rhombus__000000_0_0.CenterY = 0.000000
	__Rhombus__000000_0_0.SideLength = 100.000000
	__Rhombus__000000_0_0.Angle = 0.000000
	__Rhombus__000000_0_0.InsideAngle = 77.000000
	__Rhombus__000000_0_0.StrokeWidth = 1.000000

	__Rhombus__000001_0_1.Name = `0 1`
	__Rhombus__000001_0_1.IsDisplayed = true
	__Rhombus__000001_0_1.CenterX = -78.260816
	__Rhombus__000001_0_1.CenterY = 62.251464
	__Rhombus__000001_0_1.SideLength = 100.000000
	__Rhombus__000001_0_1.Angle = 0.000000
	__Rhombus__000001_0_1.InsideAngle = 77.000000
	__Rhombus__000001_0_1.StrokeWidth = 1.000000

	__Rhombus__000002_0_2.Name = `0 2`
	__Rhombus__000002_0_2.IsDisplayed = true
	__Rhombus__000002_0_2.CenterX = -156.521631
	__Rhombus__000002_0_2.CenterY = 124.502927
	__Rhombus__000002_0_2.SideLength = 100.000000
	__Rhombus__000002_0_2.Angle = 0.000000
	__Rhombus__000002_0_2.InsideAngle = 77.000000
	__Rhombus__000002_0_2.StrokeWidth = 1.000000

	__Rhombus__000003_1_0.Name = `1 0`
	__Rhombus__000003_1_0.IsDisplayed = true
	__Rhombus__000003_1_0.CenterX = 78.260816
	__Rhombus__000003_1_0.CenterY = 62.251464
	__Rhombus__000003_1_0.SideLength = 100.000000
	__Rhombus__000003_1_0.Angle = 0.000000
	__Rhombus__000003_1_0.InsideAngle = 77.000000
	__Rhombus__000003_1_0.StrokeWidth = 1.000000

	__Rhombus__000004_1_1.Name = `1 1`
	__Rhombus__000004_1_1.IsDisplayed = true
	__Rhombus__000004_1_1.CenterX = 0.000000
	__Rhombus__000004_1_1.CenterY = 124.502927
	__Rhombus__000004_1_1.SideLength = 100.000000
	__Rhombus__000004_1_1.Angle = 0.000000
	__Rhombus__000004_1_1.InsideAngle = 77.000000
	__Rhombus__000004_1_1.StrokeWidth = 1.000000

	__Rhombus__000005_1_2.Name = `1 2`
	__Rhombus__000005_1_2.IsDisplayed = true
	__Rhombus__000005_1_2.CenterX = -78.260816
	__Rhombus__000005_1_2.CenterY = 186.754391
	__Rhombus__000005_1_2.SideLength = 100.000000
	__Rhombus__000005_1_2.Angle = 0.000000
	__Rhombus__000005_1_2.InsideAngle = 77.000000
	__Rhombus__000005_1_2.StrokeWidth = 1.000000

	__Rhombus__000006_Initial.Name = `Initial`
	__Rhombus__000006_Initial.IsDisplayed = true
	__Rhombus__000006_Initial.CenterX = 0.000000
	__Rhombus__000006_Initial.CenterY = 0.000000
	__Rhombus__000006_Initial.SideLength = 100.000000
	__Rhombus__000006_Initial.Angle = 0.000000
	__Rhombus__000006_Initial.InsideAngle = 77.000000
	__Rhombus__000006_Initial.StrokeWidth = 1.000000

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
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000006_Initial
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000000_Initial
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
