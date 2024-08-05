package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
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

	__Field__000000_IsSkipped := (&models.Field{Name: `IsSkipped`}).Stage(stage)

	__GongStructShape__000000_Music_NoteInfo := (&models.GongStructShape{Name: `Music-NoteInfo`}).Stage(stage)
	__GongStructShape__000001_Music_Parameter := (&models.GongStructShape{Name: `Music-Parameter`}).Stage(stage)

	__Link__000000_NoteInfos := (&models.Link{Name: `NoteInfos`}).Stage(stage)

	__Position__000000_Pos_Music_NoteInfo := (&models.Position{Name: `Pos-Music-NoteInfo`}).Stage(stage)
	__Position__000001_Pos_Music_Parameter := (&models.Position{Name: `Pos-Music-Parameter`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Music_in_middle_between_Music_Parameter_and_Music_NoteInfo := (&models.Vertice{Name: `Verticle in class diagram Music in middle between Music-Parameter and Music-NoteInfo`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Music.Name = `Music`
	__Classdiagram__000000_Music.IsInDrawMode = false

	__Field__000000_IsSkipped.Name = `IsSkipped`

	//gong:ident [ref_models.NoteInfo.IsSkipped] comment added to overcome the problem with the comment map association
	__Field__000000_IsSkipped.Identifier = `ref_models.NoteInfo.IsSkipped`
	__Field__000000_IsSkipped.FieldTypeAsString = ``
	__Field__000000_IsSkipped.Structname = `NoteInfo`
	__Field__000000_IsSkipped.Fieldtypename = `bool`

	__GongStructShape__000000_Music_NoteInfo.Name = `Music-NoteInfo`

	//gong:ident [ref_models.NoteInfo] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Music_NoteInfo.Identifier = `ref_models.NoteInfo`
	__GongStructShape__000000_Music_NoteInfo.ShowNbInstances = false
	__GongStructShape__000000_Music_NoteInfo.NbInstances = 0
	__GongStructShape__000000_Music_NoteInfo.Width = 240.000000
	__GongStructShape__000000_Music_NoteInfo.Height = 78.000000
	__GongStructShape__000000_Music_NoteInfo.IsSelected = false

	__GongStructShape__000001_Music_Parameter.Name = `Music-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Music_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000001_Music_Parameter.ShowNbInstances = false
	__GongStructShape__000001_Music_Parameter.NbInstances = 1
	__GongStructShape__000001_Music_Parameter.Width = 240.000000
	__GongStructShape__000001_Music_Parameter.Height = 63.000000
	__GongStructShape__000001_Music_Parameter.IsSelected = false

	__Link__000000_NoteInfos.Name = `NoteInfos`

	//gong:ident [ref_models.Parameter.NoteInfos] comment added to overcome the problem with the comment map association
	__Link__000000_NoteInfos.Identifier = `ref_models.Parameter.NoteInfos`

	//gong:ident [ref_models.NoteInfo] comment added to overcome the problem with the comment map association
	__Link__000000_NoteInfos.Fieldtypename = `ref_models.NoteInfo`
	__Link__000000_NoteInfos.FieldOffsetX = -50.000000
	__Link__000000_NoteInfos.FieldOffsetY = -16.000000
	__Link__000000_NoteInfos.TargetMultiplicity = models.MANY
	__Link__000000_NoteInfos.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_NoteInfos.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_NoteInfos.SourceMultiplicity = models.MANY
	__Link__000000_NoteInfos.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_NoteInfos.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_NoteInfos.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_NoteInfos.StartRatio = 0.500000
	__Link__000000_NoteInfos.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_NoteInfos.EndRatio = 0.500000
	__Link__000000_NoteInfos.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Music_NoteInfo.X = 538.000000
	__Position__000000_Pos_Music_NoteInfo.Y = 60.000000
	__Position__000000_Pos_Music_NoteInfo.Name = `Pos-Music-NoteInfo`

	__Position__000001_Pos_Music_Parameter.X = 37.000000
	__Position__000001_Pos_Music_Parameter.Y = 57.000000
	__Position__000001_Pos_Music_Parameter.Name = `Pos-Music-Parameter`

	__Vertice__000000_Verticle_in_class_diagram_Music_in_middle_between_Music_Parameter_and_Music_NoteInfo.X = 636.000000
	__Vertice__000000_Verticle_in_class_diagram_Music_in_middle_between_Music_Parameter_and_Music_NoteInfo.Y = 84.000000
	__Vertice__000000_Verticle_in_class_diagram_Music_in_middle_between_Music_Parameter_and_Music_NoteInfo.Name = `Verticle in class diagram Music in middle between Music-Parameter and Music-NoteInfo`

	// Setup of pointers
	__Classdiagram__000000_Music.GongStructShapes = append(__Classdiagram__000000_Music.GongStructShapes, __GongStructShape__000001_Music_Parameter)
	__Classdiagram__000000_Music.GongStructShapes = append(__Classdiagram__000000_Music.GongStructShapes, __GongStructShape__000000_Music_NoteInfo)
	__GongStructShape__000000_Music_NoteInfo.Position = __Position__000000_Pos_Music_NoteInfo
	__GongStructShape__000000_Music_NoteInfo.Fields = append(__GongStructShape__000000_Music_NoteInfo.Fields, __Field__000000_IsSkipped)
	__GongStructShape__000001_Music_Parameter.Position = __Position__000001_Pos_Music_Parameter
	__GongStructShape__000001_Music_Parameter.Links = append(__GongStructShape__000001_Music_Parameter.Links, __Link__000000_NoteInfos)
	__Link__000000_NoteInfos.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Music_in_middle_between_Music_Parameter_and_Music_NoteInfo
}
