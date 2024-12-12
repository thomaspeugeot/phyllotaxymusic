package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongtone/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_Duration := (&models.Field{}).Stage(stage)
	__Field__000001_Info := (&models.Field{}).Stage(stage)
	__Field__000002_Name := (&models.Field{}).Stage(stage)
	__Field__000003_Name := (&models.Field{}).Stage(stage)
	__Field__000004_Start := (&models.Field{}).Stage(stage)
	__Field__000005_Velocity := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_Freqency := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Note := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Frequencies := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_Freqency := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Note := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Note_and_Default_Freqency := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Duration.Name = `Duration`

	//gong:ident [ref_models.Note.Duration] comment added to overcome the problem with the comment map association
	__Field__000000_Duration.Identifier = `ref_models.Note.Duration`
	__Field__000000_Duration.FieldTypeAsString = ``
	__Field__000000_Duration.Structname = `Note`
	__Field__000000_Duration.Fieldtypename = `float64`

	__Field__000001_Info.Name = `Info`

	//gong:ident [ref_models.Note.Info] comment added to overcome the problem with the comment map association
	__Field__000001_Info.Identifier = `ref_models.Note.Info`
	__Field__000001_Info.FieldTypeAsString = ``
	__Field__000001_Info.Structname = `Note`
	__Field__000001_Info.Fieldtypename = `string`

	__Field__000002_Name.Name = `Name`

	//gong:ident [ref_models.Freqency.Name] comment added to overcome the problem with the comment map association
	__Field__000002_Name.Identifier = `ref_models.Freqency.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Freqency`
	__Field__000002_Name.Fieldtypename = `string`

	__Field__000003_Name.Name = `Name`

	//gong:ident [ref_models.Note.Name] comment added to overcome the problem with the comment map association
	__Field__000003_Name.Identifier = `ref_models.Note.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `Note`
	__Field__000003_Name.Fieldtypename = `string`

	__Field__000004_Start.Name = `Start`

	//gong:ident [ref_models.Note.Start] comment added to overcome the problem with the comment map association
	__Field__000004_Start.Identifier = `ref_models.Note.Start`
	__Field__000004_Start.FieldTypeAsString = ``
	__Field__000004_Start.Structname = `Note`
	__Field__000004_Start.Fieldtypename = `float64`

	__Field__000005_Velocity.Name = `Velocity`

	//gong:ident [ref_models.Note.Velocity] comment added to overcome the problem with the comment map association
	__Field__000005_Velocity.Identifier = `ref_models.Note.Velocity`
	__Field__000005_Velocity.FieldTypeAsString = ``
	__Field__000005_Velocity.Structname = `Note`
	__Field__000005_Velocity.Fieldtypename = `float64`

	__GongStructShape__000000_Default_Freqency.Name = `Default-Freqency`

	//gong:ident [ref_models.Freqency] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Freqency.Identifier = `ref_models.Freqency`
	__GongStructShape__000000_Default_Freqency.ShowNbInstances = false
	__GongStructShape__000000_Default_Freqency.NbInstances = 0
	__GongStructShape__000000_Default_Freqency.Width = 240.000000
	__GongStructShape__000000_Default_Freqency.Height = 146.000000
	__GongStructShape__000000_Default_Freqency.IsSelected = false

	__GongStructShape__000001_Default_Note.Name = `Default-Note`

	//gong:ident [ref_models.Note] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Note.Identifier = `ref_models.Note`
	__GongStructShape__000001_Default_Note.ShowNbInstances = false
	__GongStructShape__000001_Default_Note.NbInstances = 0
	__GongStructShape__000001_Default_Note.Width = 240.000000
	__GongStructShape__000001_Default_Note.Height = 138.000000
	__GongStructShape__000001_Default_Note.IsSelected = false

	__Link__000000_Frequencies.Name = `Frequencies`

	//gong:ident [ref_models.Note.Frequencies] comment added to overcome the problem with the comment map association
	__Link__000000_Frequencies.Identifier = `ref_models.Note.Frequencies`

	//gong:ident [ref_models.Freqency] comment added to overcome the problem with the comment map association
	__Link__000000_Frequencies.Fieldtypename = `ref_models.Freqency`
	__Link__000000_Frequencies.FieldOffsetX = -50.000000
	__Link__000000_Frequencies.FieldOffsetY = -16.000000
	__Link__000000_Frequencies.TargetMultiplicity = models.MANY
	__Link__000000_Frequencies.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Frequencies.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Frequencies.SourceMultiplicity = models.MANY
	__Link__000000_Frequencies.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_Frequencies.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_Frequencies.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Frequencies.StartRatio = 0.500000
	__Link__000000_Frequencies.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Frequencies.EndRatio = 0.500000
	__Link__000000_Frequencies.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_Freqency.X = 474.000000
	__Position__000000_Pos_Default_Freqency.Y = 72.000000
	__Position__000000_Pos_Default_Freqency.Name = `Pos-Default-Freqency`

	__Position__000001_Pos_Default_Note.X = 31.000000
	__Position__000001_Pos_Default_Note.Y = 75.000000
	__Position__000001_Pos_Default_Note.Name = `Pos-Default-Note`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Note_and_Default_Freqency.X = 612.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Note_and_Default_Freqency.Y = 105.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Note_and_Default_Freqency.Name = `Verticle in class diagram Default in middle between Default-Note and Default-Freqency`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Note)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Freqency)
	__GongStructShape__000000_Default_Freqency.Position = __Position__000000_Pos_Default_Freqency
	__GongStructShape__000000_Default_Freqency.Fields = append(__GongStructShape__000000_Default_Freqency.Fields, __Field__000002_Name)
	__GongStructShape__000001_Default_Note.Position = __Position__000001_Pos_Default_Note
	__GongStructShape__000001_Default_Note.Fields = append(__GongStructShape__000001_Default_Note.Fields, __Field__000003_Name)
	__GongStructShape__000001_Default_Note.Fields = append(__GongStructShape__000001_Default_Note.Fields, __Field__000004_Start)
	__GongStructShape__000001_Default_Note.Fields = append(__GongStructShape__000001_Default_Note.Fields, __Field__000000_Duration)
	__GongStructShape__000001_Default_Note.Fields = append(__GongStructShape__000001_Default_Note.Fields, __Field__000005_Velocity)
	__GongStructShape__000001_Default_Note.Fields = append(__GongStructShape__000001_Default_Note.Fields, __Field__000001_Info)
	__GongStructShape__000001_Default_Note.Links = append(__GongStructShape__000001_Default_Note.Links, __Link__000000_Frequencies)
	__Link__000000_Frequencies.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Note_and_Default_Freqency
}
