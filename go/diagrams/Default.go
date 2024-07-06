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

// Injection point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.GongsimStackName": ref_models.GongsimStackName,

	"ref_models.GongsvgStackName": ref_models.GongsvgStackName,

	"ref_models.GongtableStackName": ref_models.GongtableStackName,

	"ref_models.GongtreeStackName": ref_models.GongtreeStackName,

	"ref_models.HorizontalAxis": &(ref_models.HorizontalAxis{}),

	"ref_models.HorizontalAxis.AxisHandleBorderLength": (ref_models.HorizontalAxis{}).AxisHandleBorderLength,

	"ref_models.HorizontalAxis.Axis_StrokeWidth": (ref_models.HorizontalAxis{}).Axis_StrokeWidth,

	"ref_models.HorizontalAxis.HorizontalAxis_Length": (ref_models.HorizontalAxis{}).HorizontalAxis_Length,

	"ref_models.HorizontalAxis.IsHorizontalAxisDisplayed": (ref_models.HorizontalAxis{}).IsHorizontalAxisDisplayed,

	"ref_models.HorizontalAxis.Name": (ref_models.HorizontalAxis{}).Name,

	"ref_models.HorizontalAxis.OriginX": (ref_models.HorizontalAxis{}).OriginX,

	"ref_models.HorizontalAxis.OriginY": (ref_models.HorizontalAxis{}).OriginY,

	"ref_models.HorizontalAxis.VerticalAxis_Length": (ref_models.HorizontalAxis{}).VerticalAxis_Length,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.X1": (ref_models.Line{}).X1,

	"ref_models.Line.X2": (ref_models.Line{}).X2,

	"ref_models.Line.Y1": (ref_models.Line{}).Y1,

	"ref_models.Line.Y2": (ref_models.Line{}).Y2,

	"ref_models.Parameter": &(ref_models.Parameter{}),

	"ref_models.Parameter.Angle": (ref_models.Parameter{}).Angle,

	"ref_models.Parameter.DiamondSideLenght": (ref_models.Parameter{}).DiamondSideLenght,

	"ref_models.Parameter.HorizontalAxis": (ref_models.Parameter{}).HorizontalAxis,

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
	__Field__000001_AxisHandleBorderLength := (&models.Field{Name: `AxisHandleBorderLength`}).Stage(stage)
	__Field__000002_Axis_StrokeWidth := (&models.Field{Name: `Axis_StrokeWidth`}).Stage(stage)
	__Field__000003_CenterX := (&models.Field{Name: `CenterX`}).Stage(stage)
	__Field__000004_CenterY := (&models.Field{Name: `CenterY`}).Stage(stage)
	__Field__000005_DiamondAngle := (&models.Field{Name: `DiamondAngle`}).Stage(stage)
	__Field__000006_DiamondSideLenght := (&models.Field{Name: `DiamondSideLenght`}).Stage(stage)
	__Field__000007_HorizontalAxis_Length := (&models.Field{Name: `HorizontalAxis_Length`}).Stage(stage)
	__Field__000008_InsideAngle := (&models.Field{Name: `InsideAngle`}).Stage(stage)
	__Field__000009_IsHorizontalAxisDisplayed := (&models.Field{Name: `IsHorizontalAxisDisplayed`}).Stage(stage)
	__Field__000010_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000011_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000014_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000015_OriginX := (&models.Field{Name: `OriginX`}).Stage(stage)
	__Field__000016_OriginX := (&models.Field{Name: `OriginX`}).Stage(stage)
	__Field__000017_OriginY := (&models.Field{Name: `OriginY`}).Stage(stage)
	__Field__000018_OriginY := (&models.Field{Name: `OriginY`}).Stage(stage)
	__Field__000019_SideLength := (&models.Field{Name: `SideLength`}).Stage(stage)
	__Field__000020_VerticalAxis_Length := (&models.Field{Name: `VerticalAxis_Length`}).Stage(stage)

	__GongStructShape__000000_Default_HorizontalAxis := (&models.GongStructShape{Name: `Default-HorizontalAxis`}).Stage(stage)
	__GongStructShape__000001_Default_Parameter := (&models.GongStructShape{Name: `Default-Parameter`}).Stage(stage)
	__GongStructShape__000002_Default_Rhombus := (&models.GongStructShape{Name: `Default-Rhombus`}).Stage(stage)

	__Link__000000_HorizontalAxis := (&models.Link{Name: `HorizontalAxis`}).Stage(stage)
	__Link__000001_InitialRhombus := (&models.Link{Name: `InitialRhombus`}).Stage(stage)

	__Position__000000_Pos_Default_HorizontalAxis := (&models.Position{Name: `Pos-Default-HorizontalAxis`}).Stage(stage)
	__Position__000001_Pos_Default_Parameter := (&models.Position{Name: `Pos-Default-Parameter`}).Stage(stage)
	__Position__000002_Pos_Default_Rhombus := (&models.Position{Name: `Pos-Default-Rhombus`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-HorizontalAxis`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Angle.Name = `Angle`

	//gong:ident [ref_models.Rhombus.Angle] comment added to overcome the problem with the comment map association
	__Field__000000_Angle.Identifier = `ref_models.Rhombus.Angle`
	__Field__000000_Angle.FieldTypeAsString = ``
	__Field__000000_Angle.Structname = `Rhombus`
	__Field__000000_Angle.Fieldtypename = `float64`

	__Field__000001_AxisHandleBorderLength.Name = `AxisHandleBorderLength`

	//gong:ident [ref_models.HorizontalAxis.AxisHandleBorderLength] comment added to overcome the problem with the comment map association
	__Field__000001_AxisHandleBorderLength.Identifier = `ref_models.HorizontalAxis.AxisHandleBorderLength`
	__Field__000001_AxisHandleBorderLength.FieldTypeAsString = ``
	__Field__000001_AxisHandleBorderLength.Structname = `HorizontalAxis`
	__Field__000001_AxisHandleBorderLength.Fieldtypename = `float64`

	__Field__000002_Axis_StrokeWidth.Name = `Axis_StrokeWidth`

	//gong:ident [ref_models.HorizontalAxis.Axis_StrokeWidth] comment added to overcome the problem with the comment map association
	__Field__000002_Axis_StrokeWidth.Identifier = `ref_models.HorizontalAxis.Axis_StrokeWidth`
	__Field__000002_Axis_StrokeWidth.FieldTypeAsString = ``
	__Field__000002_Axis_StrokeWidth.Structname = `HorizontalAxis`
	__Field__000002_Axis_StrokeWidth.Fieldtypename = `float64`

	__Field__000003_CenterX.Name = `CenterX`

	//gong:ident [ref_models.Rhombus.CenterX] comment added to overcome the problem with the comment map association
	__Field__000003_CenterX.Identifier = `ref_models.Rhombus.CenterX`
	__Field__000003_CenterX.FieldTypeAsString = ``
	__Field__000003_CenterX.Structname = `Rhombus`
	__Field__000003_CenterX.Fieldtypename = `float64`

	__Field__000004_CenterY.Name = `CenterY`

	//gong:ident [ref_models.Rhombus.CenterY] comment added to overcome the problem with the comment map association
	__Field__000004_CenterY.Identifier = `ref_models.Rhombus.CenterY`
	__Field__000004_CenterY.FieldTypeAsString = ``
	__Field__000004_CenterY.Structname = `Rhombus`
	__Field__000004_CenterY.Fieldtypename = `float64`

	__Field__000005_DiamondAngle.Name = `DiamondAngle`

	//gong:ident [ref_models.Parameter.Angle] comment added to overcome the problem with the comment map association
	__Field__000005_DiamondAngle.Identifier = `ref_models.Parameter.Angle`
	__Field__000005_DiamondAngle.FieldTypeAsString = ``
	__Field__000005_DiamondAngle.Structname = `Parameter`
	__Field__000005_DiamondAngle.Fieldtypename = `float64`

	__Field__000006_DiamondSideLenght.Name = `DiamondSideLenght`

	//gong:ident [ref_models.Parameter.DiamondSideLenght] comment added to overcome the problem with the comment map association
	__Field__000006_DiamondSideLenght.Identifier = `ref_models.Parameter.DiamondSideLenght`
	__Field__000006_DiamondSideLenght.FieldTypeAsString = ``
	__Field__000006_DiamondSideLenght.Structname = `Parameter`
	__Field__000006_DiamondSideLenght.Fieldtypename = `float64`

	__Field__000007_HorizontalAxis_Length.Name = `HorizontalAxis_Length`

	//gong:ident [ref_models.HorizontalAxis.HorizontalAxis_Length] comment added to overcome the problem with the comment map association
	__Field__000007_HorizontalAxis_Length.Identifier = `ref_models.HorizontalAxis.HorizontalAxis_Length`
	__Field__000007_HorizontalAxis_Length.FieldTypeAsString = ``
	__Field__000007_HorizontalAxis_Length.Structname = `HorizontalAxis`
	__Field__000007_HorizontalAxis_Length.Fieldtypename = `float64`

	__Field__000008_InsideAngle.Name = `InsideAngle`

	//gong:ident [ref_models.Rhombus.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000008_InsideAngle.Identifier = `ref_models.Rhombus.InsideAngle`
	__Field__000008_InsideAngle.FieldTypeAsString = ``
	__Field__000008_InsideAngle.Structname = `Rhombus`
	__Field__000008_InsideAngle.Fieldtypename = `float64`

	__Field__000009_IsHorizontalAxisDisplayed.Name = `IsHorizontalAxisDisplayed`

	//gong:ident [ref_models.HorizontalAxis.IsHorizontalAxisDisplayed] comment added to overcome the problem with the comment map association
	__Field__000009_IsHorizontalAxisDisplayed.Identifier = `ref_models.HorizontalAxis.IsHorizontalAxisDisplayed`
	__Field__000009_IsHorizontalAxisDisplayed.FieldTypeAsString = ``
	__Field__000009_IsHorizontalAxisDisplayed.Structname = `HorizontalAxis`
	__Field__000009_IsHorizontalAxisDisplayed.Fieldtypename = `bool`

	__Field__000010_M.Name = `M`

	//gong:ident [ref_models.Parameter.M] comment added to overcome the problem with the comment map association
	__Field__000010_M.Identifier = `ref_models.Parameter.M`
	__Field__000010_M.FieldTypeAsString = ``
	__Field__000010_M.Structname = `Parameter`
	__Field__000010_M.Fieldtypename = `int`

	__Field__000011_N.Name = `N`

	//gong:ident [ref_models.Parameter.N] comment added to overcome the problem with the comment map association
	__Field__000011_N.Identifier = `ref_models.Parameter.N`
	__Field__000011_N.FieldTypeAsString = ``
	__Field__000011_N.Structname = `Parameter`
	__Field__000011_N.Fieldtypename = `int`

	__Field__000012_Name.Name = `Name`

	//gong:ident [ref_models.HorizontalAxis.Name] comment added to overcome the problem with the comment map association
	__Field__000012_Name.Identifier = `ref_models.HorizontalAxis.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `HorizontalAxis`
	__Field__000012_Name.Fieldtypename = `string`

	__Field__000013_Name.Name = `Name`

	//gong:ident [ref_models.Rhombus.Name] comment added to overcome the problem with the comment map association
	__Field__000013_Name.Identifier = `ref_models.Rhombus.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Rhombus`
	__Field__000013_Name.Fieldtypename = `string`

	__Field__000014_Name.Name = `Name`

	//gong:ident [ref_models.Parameter.Name] comment added to overcome the problem with the comment map association
	__Field__000014_Name.Identifier = `ref_models.Parameter.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `Parameter`
	__Field__000014_Name.Fieldtypename = `string`

	__Field__000015_OriginX.Name = `OriginX`

	//gong:ident [ref_models.Parameter.OriginX] comment added to overcome the problem with the comment map association
	__Field__000015_OriginX.Identifier = `ref_models.Parameter.OriginX`
	__Field__000015_OriginX.FieldTypeAsString = ``
	__Field__000015_OriginX.Structname = `Parameter`
	__Field__000015_OriginX.Fieldtypename = `float64`

	__Field__000016_OriginX.Name = `OriginX`

	//gong:ident [ref_models.HorizontalAxis.OriginX] comment added to overcome the problem with the comment map association
	__Field__000016_OriginX.Identifier = `ref_models.HorizontalAxis.OriginX`
	__Field__000016_OriginX.FieldTypeAsString = ``
	__Field__000016_OriginX.Structname = `HorizontalAxis`
	__Field__000016_OriginX.Fieldtypename = `float64`

	__Field__000017_OriginY.Name = `OriginY`

	//gong:ident [ref_models.HorizontalAxis.OriginY] comment added to overcome the problem with the comment map association
	__Field__000017_OriginY.Identifier = `ref_models.HorizontalAxis.OriginY`
	__Field__000017_OriginY.FieldTypeAsString = ``
	__Field__000017_OriginY.Structname = `HorizontalAxis`
	__Field__000017_OriginY.Fieldtypename = `float64`

	__Field__000018_OriginY.Name = `OriginY`

	//gong:ident [ref_models.Parameter.OriginY] comment added to overcome the problem with the comment map association
	__Field__000018_OriginY.Identifier = `ref_models.Parameter.OriginY`
	__Field__000018_OriginY.FieldTypeAsString = ``
	__Field__000018_OriginY.Structname = `Parameter`
	__Field__000018_OriginY.Fieldtypename = `float64`

	__Field__000019_SideLength.Name = `SideLength`

	//gong:ident [ref_models.Rhombus.SideLength] comment added to overcome the problem with the comment map association
	__Field__000019_SideLength.Identifier = `ref_models.Rhombus.SideLength`
	__Field__000019_SideLength.FieldTypeAsString = ``
	__Field__000019_SideLength.Structname = `Rhombus`
	__Field__000019_SideLength.Fieldtypename = `float64`

	__Field__000020_VerticalAxis_Length.Name = `VerticalAxis_Length`

	//gong:ident [ref_models.HorizontalAxis.VerticalAxis_Length] comment added to overcome the problem with the comment map association
	__Field__000020_VerticalAxis_Length.Identifier = `ref_models.HorizontalAxis.VerticalAxis_Length`
	__Field__000020_VerticalAxis_Length.FieldTypeAsString = ``
	__Field__000020_VerticalAxis_Length.Structname = `HorizontalAxis`
	__Field__000020_VerticalAxis_Length.Fieldtypename = `float64`

	__GongStructShape__000000_Default_HorizontalAxis.Name = `Default-HorizontalAxis`

	//gong:ident [ref_models.HorizontalAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_HorizontalAxis.Identifier = `ref_models.HorizontalAxis`
	__GongStructShape__000000_Default_HorizontalAxis.ShowNbInstances = false
	__GongStructShape__000000_Default_HorizontalAxis.NbInstances = 0
	__GongStructShape__000000_Default_HorizontalAxis.Width = 324.000000
	__GongStructShape__000000_Default_HorizontalAxis.Height = 183.000000
	__GongStructShape__000000_Default_HorizontalAxis.IsSelected = false

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
	__GongStructShape__000002_Default_Rhombus.NbInstances = 1
	__GongStructShape__000002_Default_Rhombus.Width = 240.000000
	__GongStructShape__000002_Default_Rhombus.Height = 153.000000
	__GongStructShape__000002_Default_Rhombus.IsSelected = false

	__Link__000000_HorizontalAxis.Name = `HorizontalAxis`

	//gong:ident [ref_models.Parameter.HorizontalAxis] comment added to overcome the problem with the comment map association
	__Link__000000_HorizontalAxis.Identifier = `ref_models.Parameter.HorizontalAxis`

	//gong:ident [ref_models.HorizontalAxis] comment added to overcome the problem with the comment map association
	__Link__000000_HorizontalAxis.Fieldtypename = `ref_models.HorizontalAxis`
	__Link__000000_HorizontalAxis.FieldOffsetX = -50.000000
	__Link__000000_HorizontalAxis.FieldOffsetY = -16.000000
	__Link__000000_HorizontalAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_HorizontalAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_HorizontalAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_HorizontalAxis.SourceMultiplicity = models.MANY
	__Link__000000_HorizontalAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_HorizontalAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_HorizontalAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_HorizontalAxis.StartRatio = 0.500000
	__Link__000000_HorizontalAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_HorizontalAxis.EndRatio = 0.500000
	__Link__000000_HorizontalAxis.CornerOffsetRatio = 1.380000

	__Link__000001_InitialRhombus.Name = `InitialRhombus`

	//gong:ident [ref_models.Parameter.InitialRhombus] comment added to overcome the problem with the comment map association
	__Link__000001_InitialRhombus.Identifier = `ref_models.Parameter.InitialRhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000001_InitialRhombus.Fieldtypename = `ref_models.Rhombus`
	__Link__000001_InitialRhombus.FieldOffsetX = -50.000000
	__Link__000001_InitialRhombus.FieldOffsetY = -16.000000
	__Link__000001_InitialRhombus.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_InitialRhombus.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_InitialRhombus.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_InitialRhombus.SourceMultiplicity = models.MANY
	__Link__000001_InitialRhombus.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_InitialRhombus.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_InitialRhombus.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_InitialRhombus.StartRatio = 0.553403
	__Link__000001_InitialRhombus.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_InitialRhombus.EndRatio = 0.500000
	__Link__000001_InitialRhombus.CornerOffsetRatio = 1.019642

	__Position__000000_Pos_Default_HorizontalAxis.X = 506.999969
	__Position__000000_Pos_Default_HorizontalAxis.Y = 21.000000
	__Position__000000_Pos_Default_HorizontalAxis.Name = `Pos-Default-HorizontalAxis`

	__Position__000001_Pos_Default_Parameter.X = 18.000000
	__Position__000001_Pos_Default_Parameter.Y = 35.000000
	__Position__000001_Pos_Default_Parameter.Name = `Pos-Default-Parameter`

	__Position__000002_Pos_Default_Rhombus.X = 513.999969
	__Position__000002_Pos_Default_Rhombus.Y = 268.000000
	__Position__000002_Pos_Default_Rhombus.Name = `Pos-Default-Rhombus`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.X = 613.999985
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Y = 120.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-HorizontalAxis`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.X = 607.999985
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Y = 121.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Parameter)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Rhombus)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_HorizontalAxis)
	__GongStructShape__000000_Default_HorizontalAxis.Position = __Position__000000_Pos_Default_HorizontalAxis
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000012_Name)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000009_IsHorizontalAxisDisplayed)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000001_AxisHandleBorderLength)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000016_OriginX)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000017_OriginY)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000007_HorizontalAxis_Length)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000020_VerticalAxis_Length)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000002_Axis_StrokeWidth)
	__GongStructShape__000001_Default_Parameter.Position = __Position__000001_Pos_Default_Parameter
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000014_Name)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000011_N)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000010_M)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000005_DiamondAngle)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000015_OriginX)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000018_OriginY)
	__GongStructShape__000001_Default_Parameter.Fields = append(__GongStructShape__000001_Default_Parameter.Fields, __Field__000006_DiamondSideLenght)
	__GongStructShape__000001_Default_Parameter.Links = append(__GongStructShape__000001_Default_Parameter.Links, __Link__000001_InitialRhombus)
	__GongStructShape__000001_Default_Parameter.Links = append(__GongStructShape__000001_Default_Parameter.Links, __Link__000000_HorizontalAxis)
	__GongStructShape__000002_Default_Rhombus.Position = __Position__000002_Pos_Default_Rhombus
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000013_Name)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000003_CenterX)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000004_CenterY)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000019_SideLength)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000000_Angle)
	__GongStructShape__000002_Default_Rhombus.Fields = append(__GongStructShape__000002_Default_Rhombus.Fields, __Field__000008_InsideAngle)
	__Link__000000_HorizontalAxis.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis
	__Link__000001_InitialRhombus.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus
}
