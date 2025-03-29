package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/thomaspeugeot/phyllotaxymusic/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_Music := (&models.Classdiagram{}).Stage(stage)

	__GongStructShape__000000_Music_Parameter := (&models.GongStructShape{}).Stage(stage)

	__Position__000000_Pos_Music_Parameter := (&models.Position{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Music.Name = `Music`
	__Classdiagram__000000_Music.IsInDrawMode = false

	__GongStructShape__000000_Music_Parameter.Name = `Music-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Music_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000000_Music_Parameter.ShowNbInstances = false
	__GongStructShape__000000_Music_Parameter.NbInstances = 1
	__GongStructShape__000000_Music_Parameter.Width = 240.000000
	__GongStructShape__000000_Music_Parameter.Height = 63.000000
	__GongStructShape__000000_Music_Parameter.IsSelected = false

	__Position__000000_Pos_Music_Parameter.X = 76.000000
	__Position__000000_Pos_Music_Parameter.Y = 201.000000
	__Position__000000_Pos_Music_Parameter.Name = `Pos-Music-Parameter`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Music.GongStructShapes = append(__Classdiagram__000000_Music.GongStructShapes, __GongStructShape__000000_Music_Parameter)
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Music_Parameter.Position = __Position__000000_Pos_Music_Parameter
	// setup of Position instances pointers
}
