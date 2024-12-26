package main

import (
	"time"

	"github.com/thomaspeugeot/phylotaxymusic/substackcursor/go/models"

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

	__Cursor__000000_Cursor := (&models.Cursor{}).Stage(stage)

	// Setup of values

	__Cursor__000000_Cursor.Name = `Cursor`
	__Cursor__000000_Cursor.AngleDegree = 0.000000
	__Cursor__000000_Cursor.Length = 0.000000
	__Cursor__000000_Cursor.CenterX = 0.000000
	__Cursor__000000_Cursor.CenterY = 0.000000
	__Cursor__000000_Cursor.StartX = 0.000000
	__Cursor__000000_Cursor.EndX = 0.000000
	__Cursor__000000_Cursor.DurationSeconds = 0.000000
	__Cursor__000000_Cursor.IsMoving = false
	__Cursor__000000_Cursor.Color = ``
	__Cursor__000000_Cursor.FillOpacity = 0.000000
	__Cursor__000000_Cursor.Stroke = ``
	__Cursor__000000_Cursor.StrokeOpacity = 0.000000
	__Cursor__000000_Cursor.StrokeWidth = 0.000000
	__Cursor__000000_Cursor.StrokeDashArray = ``
	__Cursor__000000_Cursor.StrokeDashArrayWhenSelected = ``
	__Cursor__000000_Cursor.Transform = ``

	// Setup of pointers
}
