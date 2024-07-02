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

	__Diagram__000000_Reference := (&models.Diagram{Name: `Reference`}).Stage(stage)

	// Setup of values

	__Diagram__000000_Reference.Name = `Reference`
	__Diagram__000000_Reference.N = 3
	__Diagram__000000_Reference.M = 2
	__Diagram__000000_Reference.DiamondAngle = 60.000000
	__Diagram__000000_Reference.OriginX = 100.000000
	__Diagram__000000_Reference.OriginY = 500.000000
	__Diagram__000000_Reference.DiamondSideLenght = 100.000000

	// Setup of pointers
}
