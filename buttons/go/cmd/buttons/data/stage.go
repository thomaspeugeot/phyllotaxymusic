package main

import (
	"time"

	"github.com/thomaspeugeot/phyllotaxymusic/buttons/go/models"

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

	__Button__000000_Button_1 := (&models.Button{}).Stage(stage)

	__Group__000000_Group_1 := (&models.Group{}).Stage(stage)
	__Group__000001_Group_2 := (&models.Group{}).Stage(stage)

	__Layout__000000_Layout := (&models.Layout{}).Stage(stage)

	// Setup of values

	__Button__000000_Button_1.Name = `Button 1`
	__Button__000000_Button_1.Label = `Button Label`
	__Button__000000_Button_1.Icon = `music_note`

	__Group__000000_Group_1.Name = `Group 1`
	__Group__000000_Group_1.Percentage = 50.000000

	__Group__000001_Group_2.Name = `Group 2`
	__Group__000001_Group_2.Percentage = 50.000000

	__Layout__000000_Layout.Name = `Layout`

	// Setup of pointers
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000000_Button_1)
	__Layout__000000_Layout.Groups = append(__Layout__000000_Layout.Groups, __Group__000000_Group_1)
	__Layout__000000_Layout.Groups = append(__Layout__000000_Layout.Groups, __Group__000001_Group_2)
}
