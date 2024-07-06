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
	// injection point for docLink to identifiers

	"ref_models.GongsimStackName": ref_models.GongsimStackName,

	"ref_models.GongsvgStackName": ref_models.GongsvgStackName,

	"ref_models.GongtableStackName": ref_models.GongtableStackName,

	"ref_models.GongtreeStackName": ref_models.GongtreeStackName,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.X1": (ref_models.Line{}).X1,

	"ref_models.Line.X2": (ref_models.Line{}).X2,

	"ref_models.Line.Y1": (ref_models.Line{}).Y1,

	"ref_models.Line.Y2": (ref_models.Line{}).Y2,

	"ref_models.Parameter": &(ref_models.Parameter{}),

	"ref_models.Parameter.Angle": (ref_models.Parameter{}).Angle,

	"ref_models.Parameter.DiamondSideLenght": (ref_models.Parameter{}).DiamondSideLenght,

	"ref_models.Parameter.InitialRhombus": (ref_models.Parameter{}).InitialRhombus,

	"ref_models.Parameter.M": (ref_models.Parameter{}).M,

	"ref_models.Parameter.N": (ref_models.Parameter{}).N,

	"ref_models.Parameter.Name": (ref_models.Parameter{}).Name,

	"ref_models.Parameter.OriginX": (ref_models.Parameter{}).OriginX,

	"ref_models.Parameter.OriginY": (ref_models.Parameter{}).OriginY,

	"ref_models.Phylotaxy": ref_models.Phylotaxy,

	"ref_models.Rhombus": &(ref_models.Rhombus{}),

	"ref_models.Rhombus.Angle": (ref_models.Rhombus{}).Angle,

	"ref_models.Rhombus.CenterX": (ref_models.Rhombus{}).CenterX,

	"ref_models.Rhombus.CenterY": (ref_models.Rhombus{}).CenterY,

	"ref_models.Rhombus.InsideAngle": (ref_models.Rhombus{}).InsideAngle,

	"ref_models.Rhombus.Name": (ref_models.Rhombus{}).Name,

	"ref_models.Rhombus.SideLength": (ref_models.Rhombus{}).SideLength,

	"ref_models.StacksNames": ref_models.StacksNames(""),
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_DiamondAngle := (&models.Field{Name: `DiamondAngle`}).Stage(stage)
	__Field__000001_DiamondSideLenght := (&models.Field{Name: `DiamondSideLenght`}).Stage(stage)
	__Field__000002_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000003_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000005_OriginX := (&models.Field{Name: `OriginX`}).Stage(stage)
	__Field__000006_OriginY := (&models.Field{Name: `OriginY`}).Stage(stage)

	__GongStructShape__000000_Default_Parameter := (&models.GongStructShape{Name: `Default-Parameter`}).Stage(stage)

	__Position__000000_Pos_Default_Parameter := (&models.Position{Name: `Pos-Default-Parameter`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_DiamondAngle.Name = `DiamondAngle`

	//gong:ident [ref_models.Parameter.Angle] comment added to overcome the problem with the comment map association
	__Field__000000_DiamondAngle.Identifier = `ref_models.Parameter.Angle`
	__Field__000000_DiamondAngle.FieldTypeAsString = ``
	__Field__000000_DiamondAngle.Structname = `Parameter`
	__Field__000000_DiamondAngle.Fieldtypename = `float64`

	__Field__000001_DiamondSideLenght.Name = `DiamondSideLenght`

	//gong:ident [ref_models.Parameter.DiamondSideLenght] comment added to overcome the problem with the comment map association
	__Field__000001_DiamondSideLenght.Identifier = `ref_models.Parameter.DiamondSideLenght`
	__Field__000001_DiamondSideLenght.FieldTypeAsString = ``
	__Field__000001_DiamondSideLenght.Structname = `Parameter`
	__Field__000001_DiamondSideLenght.Fieldtypename = `float64`

	__Field__000002_M.Name = `M`

	//gong:ident [ref_models.Parameter.M] comment added to overcome the problem with the comment map association
	__Field__000002_M.Identifier = `ref_models.Parameter.M`
	__Field__000002_M.FieldTypeAsString = ``
	__Field__000002_M.Structname = `Parameter`
	__Field__000002_M.Fieldtypename = `int`

	__Field__000003_N.Name = `N`

	//gong:ident [ref_models.Parameter.N] comment added to overcome the problem with the comment map association
	__Field__000003_N.Identifier = `ref_models.Parameter.N`
	__Field__000003_N.FieldTypeAsString = ``
	__Field__000003_N.Structname = `Parameter`
	__Field__000003_N.Fieldtypename = `int`

	__Field__000004_Name.Name = `Name`

	//gong:ident [ref_models.Parameter.Name] comment added to overcome the problem with the comment map association
	__Field__000004_Name.Identifier = `ref_models.Parameter.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `Parameter`
	__Field__000004_Name.Fieldtypename = `string`

	__Field__000005_OriginX.Name = `OriginX`

	//gong:ident [ref_models.Parameter.OriginX] comment added to overcome the problem with the comment map association
	__Field__000005_OriginX.Identifier = `ref_models.Parameter.OriginX`
	__Field__000005_OriginX.FieldTypeAsString = ``
	__Field__000005_OriginX.Structname = `Parameter`
	__Field__000005_OriginX.Fieldtypename = `float64`

	__Field__000006_OriginY.Name = `OriginY`

	//gong:ident [ref_models.Parameter.OriginY] comment added to overcome the problem with the comment map association
	__Field__000006_OriginY.Identifier = `ref_models.Parameter.OriginY`
	__Field__000006_OriginY.FieldTypeAsString = ``
	__Field__000006_OriginY.Structname = `Parameter`
	__Field__000006_OriginY.Fieldtypename = `float64`

	__GongStructShape__000000_Default_Parameter.Name = `Default-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000000_Default_Parameter.ShowNbInstances = false
	__GongStructShape__000000_Default_Parameter.NbInstances = 1
	__GongStructShape__000000_Default_Parameter.Width = 240.000000
	__GongStructShape__000000_Default_Parameter.Height = 168.000000
	__GongStructShape__000000_Default_Parameter.IsSelected = false

	__Position__000000_Pos_Default_Parameter.X = 18.000000
	__Position__000000_Pos_Default_Parameter.Y = 35.000000
	__Position__000000_Pos_Default_Parameter.Name = `Pos-Default-Parameter`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Parameter)
	__GongStructShape__000000_Default_Parameter.Position = __Position__000000_Pos_Default_Parameter
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000004_Name)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000003_N)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000002_M)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000000_DiamondAngle)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000005_OriginX)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000006_OriginY)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000001_DiamondSideLenght)
}
