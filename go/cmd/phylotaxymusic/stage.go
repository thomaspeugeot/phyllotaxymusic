package main

import (
	"time"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Parameter__000000_Reference := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__000000_Initial := (&models.Rhombus{Name: `Initial`}).Stage(stage)

	// Setup of values

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Angle = 76.000000
	__Parameter__000000_Reference.OriginX = 100.000000
	__Parameter__000000_Reference.OriginY = 300.000000
	__Parameter__000000_Reference.DiamondSideLenght = 100.000000

	__Rhombus__000000_Initial.Name = `Initial`
	__Rhombus__000000_Initial.CenterX = 0.000000
	__Rhombus__000000_Initial.CenterY = 0.000000
	__Rhombus__000000_Initial.SideLength = 0.000000
	__Rhombus__000000_Initial.Angle = 0.000000
	__Rhombus__000000_Initial.InsideAngle = 0.000000

	// Setup of pointers
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000000_Initial
}
