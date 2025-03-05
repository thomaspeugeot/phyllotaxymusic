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
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Music := (&models.Classdiagram{Name: `Music`}).Stage(stage)

	__GongStructShape__000000_Music_NoteInfo := (&models.GongStructShape{Name: `Music-NoteInfo`}).Stage(stage)
	__GongStructShape__000001_Music_Parameter := (&models.GongStructShape{Name: `Music-Parameter`}).Stage(stage)

	__Position__000000_Pos_Music_NoteInfo := (&models.Position{Name: `Pos-Music-NoteInfo`}).Stage(stage)
	__Position__000001_Pos_Music_Parameter := (&models.Position{Name: `Pos-Music-Parameter`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Music.Name = `Music`
	__Classdiagram__000000_Music.IsInDrawMode = false

	__GongStructShape__000000_Music_NoteInfo.Name = `Music-NoteInfo`

	//gong:ident [ref_models.NoteInfo] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Music_NoteInfo.Identifier = `ref_models.NoteInfo`
	__GongStructShape__000000_Music_NoteInfo.ShowNbInstances = false
	__GongStructShape__000000_Music_NoteInfo.NbInstances = 0
	__GongStructShape__000000_Music_NoteInfo.Width = 240.000000
	__GongStructShape__000000_Music_NoteInfo.Height = 63.000000
	__GongStructShape__000000_Music_NoteInfo.IsSelected = false

	__GongStructShape__000001_Music_Parameter.Name = `Music-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Music_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000001_Music_Parameter.ShowNbInstances = false
	__GongStructShape__000001_Music_Parameter.NbInstances = 0
	__GongStructShape__000001_Music_Parameter.Width = 240.000000
	__GongStructShape__000001_Music_Parameter.Height = 63.000000
	__GongStructShape__000001_Music_Parameter.IsSelected = false

	__Position__000000_Pos_Music_NoteInfo.X = 47.000000
	__Position__000000_Pos_Music_NoteInfo.Y = 57.000000
	__Position__000000_Pos_Music_NoteInfo.Name = `Pos-Music-NoteInfo`

	__Position__000001_Pos_Music_Parameter.X = 76.000000
	__Position__000001_Pos_Music_Parameter.Y = 201.000000
	__Position__000001_Pos_Music_Parameter.Name = `Pos-Music-Parameter`

	// Setup of pointers
	__Classdiagram__000000_Music.GongStructShapes = append(__Classdiagram__000000_Music.GongStructShapes, __GongStructShape__000000_Music_NoteInfo)
	__Classdiagram__000000_Music.GongStructShapes = append(__Classdiagram__000000_Music.GongStructShapes, __GongStructShape__000001_Music_Parameter)
	__GongStructShape__000000_Music_NoteInfo.Position = __Position__000000_Pos_Music_NoteInfo
	__GongStructShape__000001_Music_Parameter.Position = __Position__000001_Pos_Music_Parameter
}
