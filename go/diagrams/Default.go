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

	"ref_models.Sidebar": ref_models.Sidebar,

	"ref_models.SidebarTree": ref_models.SidebarTree,

	"ref_models.StacksNames": ref_models.StacksNames(""),

	"ref_models.TreeNames": ref_models.TreeNames(""),
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_Angle := (&models.Field{Name: `Angle`}).Stage(stage)
	__Field__000001_CenterX := (&models.Field{Name: `CenterX`}).Stage(stage)
	__Field__000002_CenterY := (&models.Field{Name: `CenterY`}).Stage(stage)
	__Field__000003_DiamondAngle := (&models.Field{Name: `DiamondAngle`}).Stage(stage)
	__Field__000004_DiamondSideLenght := (&models.Field{Name: `DiamondSideLenght`}).Stage(stage)
	__Field__000005_InsideAngle := (&models.Field{Name: `InsideAngle`}).Stage(stage)
	__Field__000006_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000007_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000008_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000009_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000010_OriginX := (&models.Field{Name: `OriginX`}).Stage(stage)
	__Field__000011_OriginY := (&models.Field{Name: `OriginY`}).Stage(stage)
	__Field__000012_SideLength := (&models.Field{Name: `SideLength`}).Stage(stage)

	__GongStructShape__000000_Default_Line := (&models.GongStructShape{Name: `Default-Line`}).Stage(stage)
	__GongStructShape__000001_Default_Parameter := (&models.GongStructShape{Name: `Default-Parameter`}).Stage(stage)
	__GongStructShape__000002_Default_Rhombus := (&models.GongStructShape{Name: `Default-Rhombus`}).Stage(stage)

	__Link__000000_InitialRhombus := (&models.Link{Name: `InitialRhombus`}).Stage(stage)

	__Position__000000_Pos_Default_Line := (&models.Position{Name: `Pos-Default-Line`}).Stage(stage)
	__Position__000001_Pos_Default_Parameter := (&models.Position{Name: `Pos-Default-Parameter`}).Stage(stage)
	__Position__000002_Pos_Default_Rhombus := (&models.Position{Name: `Pos-Default-Rhombus`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Angle.Name = `Angle`

	//gong:ident [ref_models.Rhombus.Angle] comment added to overcome the problem with the comment map association
	__Field__000000_Angle.Identifier = `ref_models.Rhombus.Angle`
	__Field__000000_Angle.FieldTypeAsString = ``
	__Field__000000_Angle.Structname = `Rhombus`
	__Field__000000_Angle.Fieldtypename = `float64`

	__Field__000001_CenterX.Name = `CenterX`

	//gong:ident [ref_models.Rhombus.CenterX] comment added to overcome the problem with the comment map association
	__Field__000001_CenterX.Identifier = `ref_models.Rhombus.CenterX`
	__Field__000001_CenterX.FieldTypeAsString = ``
	__Field__000001_CenterX.Structname = `Rhombus`
	__Field__000001_CenterX.Fieldtypename = `float64`

	__Field__000002_CenterY.Name = `CenterY`

	//gong:ident [ref_models.Rhombus.CenterY] comment added to overcome the problem with the comment map association
	__Field__000002_CenterY.Identifier = `ref_models.Rhombus.CenterY`
	__Field__000002_CenterY.FieldTypeAsString = ``
	__Field__000002_CenterY.Structname = `Rhombus`
	__Field__000002_CenterY.Fieldtypename = `float64`

	__Field__000003_DiamondAngle.Name = `DiamondAngle`

	//gong:ident [ref_models.Parameter.Angle] comment added to overcome the problem with the comment map association
	__Field__000003_DiamondAngle.Identifier = `ref_models.Parameter.Angle`
	__Field__000003_DiamondAngle.FieldTypeAsString = ``
	__Field__000003_DiamondAngle.Structname = `Parameter`
	__Field__000003_DiamondAngle.Fieldtypename = `float64`

	__Field__000004_DiamondSideLenght.Name = `DiamondSideLenght`

	//gong:ident [ref_models.Parameter.DiamondSideLenght] comment added to overcome the problem with the comment map association
	__Field__000004_DiamondSideLenght.Identifier = `ref_models.Parameter.DiamondSideLenght`
	__Field__000004_DiamondSideLenght.FieldTypeAsString = ``
	__Field__000004_DiamondSideLenght.Structname = `Parameter`
	__Field__000004_DiamondSideLenght.Fieldtypename = `float64`

	__Field__000005_InsideAngle.Name = `InsideAngle`

	//gong:ident [ref_models.Rhombus.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000005_InsideAngle.Identifier = `ref_models.Rhombus.InsideAngle`
	__Field__000005_InsideAngle.FieldTypeAsString = ``
	__Field__000005_InsideAngle.Structname = `Rhombus`
	__Field__000005_InsideAngle.Fieldtypename = `float64`

	__Field__000006_M.Name = `M`

	//gong:ident [ref_models.Parameter.M] comment added to overcome the problem with the comment map association
	__Field__000006_M.Identifier = `ref_models.Parameter.M`
	__Field__000006_M.FieldTypeAsString = ``
	__Field__000006_M.Structname = `Parameter`
	__Field__000006_M.Fieldtypename = `int`

	__Field__000007_N.Name = `N`

	//gong:ident [ref_models.Parameter.N] comment added to overcome the problem with the comment map association
	__Field__000007_N.Identifier = `ref_models.Parameter.N`
	__Field__000007_N.FieldTypeAsString = ``
	__Field__000007_N.Structname = `Parameter`
	__Field__000007_N.Fieldtypename = `int`

	__Field__000008_Name.Name = `Name`

	//gong:ident [ref_models.Parameter.Name] comment added to overcome the problem with the comment map association
	__Field__000008_Name.Identifier = `ref_models.Parameter.Name`
	__Field__000008_Name.FieldTypeAsString = ``
	__Field__000008_Name.Structname = `Parameter`
	__Field__000008_Name.Fieldtypename = `string`

	__Field__000009_Name.Name = `Name`

	//gong:ident [ref_models.Rhombus.Name] comment added to overcome the problem with the comment map association
	__Field__000009_Name.Identifier = `ref_models.Rhombus.Name`
	__Field__000009_Name.FieldTypeAsString = ``
	__Field__000009_Name.Structname = `Rhombus`
	__Field__000009_Name.Fieldtypename = `string`

	__Field__000010_OriginX.Name = `OriginX`

	//gong:ident [ref_models.Parameter.OriginX] comment added to overcome the problem with the comment map association
	__Field__000010_OriginX.Identifier = `ref_models.Parameter.OriginX`
	__Field__000010_OriginX.FieldTypeAsString = ``
	__Field__000010_OriginX.Structname = `Parameter`
	__Field__000010_OriginX.Fieldtypename = `float64`

	__Field__000011_OriginY.Name = `OriginY`

	//gong:ident [ref_models.Parameter.OriginY] comment added to overcome the problem with the comment map association
	__Field__000011_OriginY.Identifier = `ref_models.Parameter.OriginY`
	__Field__000011_OriginY.FieldTypeAsString = ``
	__Field__000011_OriginY.Structname = `Parameter`
	__Field__000011_OriginY.Fieldtypename = `float64`

	__Field__000012_SideLength.Name = `SideLength`

	//gong:ident [ref_models.Rhombus.SideLength] comment added to overcome the problem with the comment map association
	__Field__000012_SideLength.Identifier = `ref_models.Rhombus.SideLength`
	__Field__000012_SideLength.FieldTypeAsString = ``
	__Field__000012_SideLength.Structname = `Rhombus`
	__Field__000012_SideLength.Fieldtypename = `float64`

	__GongStructShape__000000_Default_Line.Name = `Default-Line`

	//gong:ident [ref_models.Line] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Line.Identifier = `ref_models.Line`
	__GongStructShape__000000_Default_Line.ShowNbInstances = false
	__GongStructShape__000000_Default_Line.NbInstances = 0
	__GongStructShape__000000_Default_Line.Width = 240.000000
	__GongStructShape__000000_Default_Line.Height = 63.000000
	__GongStructShape__000000_Default_Line.IsSelected = false

	__GongStructShape__000001_Default_Parameter.Name = `Default-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000001_Default_Parameter.ShowNbInstances = false
	__GongStructShape__000001_Default_Parameter.NbInstances = 1
	__GongStructShape__000001_Default_Parameter.Width = 240.000000
	__GongStructShape__000001_Default_Parameter.Height = 168.000000
	__GongStructShape__000001_Default_Parameter.IsSelected = false

	__GongStructShape__000002_Default_Rhombus.Name = `Default-Rhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Rhombus.Identifier = `ref_models.Rhombus`
	__GongStructShape__000002_Default_Rhombus.ShowNbInstances = false
	__GongStructShape__000002_Default_Rhombus.NbInstances = 0
	__GongStructShape__000002_Default_Rhombus.Width = 240.000000
	__GongStructShape__000002_Default_Rhombus.Height = 153.000000
	__GongStructShape__000002_Default_Rhombus.IsSelected = false

	__Link__000000_InitialRhombus.Name = `InitialRhombus`

	//gong:ident [ref_models.Parameter.InitialRhombus] comment added to overcome the problem with the comment map association
	__Link__000000_InitialRhombus.Identifier = `ref_models.Parameter.InitialRhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000000_InitialRhombus.Fieldtypename = `ref_models.Rhombus`
	__Link__000000_InitialRhombus.FieldOffsetX = -50.000000
	__Link__000000_InitialRhombus.FieldOffsetY = -16.000000
	__Link__000000_InitialRhombus.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_InitialRhombus.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_InitialRhombus.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_InitialRhombus.SourceMultiplicity = models.MANY
	__Link__000000_InitialRhombus.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_InitialRhombus.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_InitialRhombus.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_InitialRhombus.StartRatio = 0.500000
	__Link__000000_InitialRhombus.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_InitialRhombus.EndRatio = 0.500000
	__Link__000000_InitialRhombus.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_Line.X = 887.999969
	__Position__000000_Pos_Default_Line.Y = 33.000000
	__Position__000000_Pos_Default_Line.Name = `Pos-Default-Line`

	__Position__000001_Pos_Default_Parameter.X = 18.000000
	__Position__000001_Pos_Default_Parameter.Y = 35.000000
	__Position__000001_Pos_Default_Parameter.Name = `Pos-Default-Parameter`

	__Position__000002_Pos_Default_Rhombus.X = 485.999969
	__Position__000002_Pos_Default_Rhombus.Y = 49.000000
	__Position__000002_Pos_Default_Rhombus.Name = `Pos-Default-Rhombus`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.X = 607.999985
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Y = 121.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Parameter)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Rhombus)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Line)
	__GongStructShape__000000_Default_Line.Position = __Position__000000_Pos_Default_Line
	__GongStructShape__000001_Default_Parameter.Position = __Position__000001_Pos_Default_Parameter
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000008_Name)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000007_N)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000006_M)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000003_DiamondAngle)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000010_OriginX)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000011_OriginY)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000004_DiamondSideLenght)
	__GongStructShape__000001_Default_Parameter.Links = append(__GongStructShape__000001_Default_Parameter.Links, __Link__000000_InitialRhombus)
	__GongStructShape__000002_Default_Rhombus.Position = __Position__000002_Pos_Default_Rhombus
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000009_Name)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000001_CenterX)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000002_CenterY)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000012_SideLength)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000000_Angle)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000005_InsideAngle)
	__Link__000000_InitialRhombus.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus
}
