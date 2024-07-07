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

	"ref_models.HorizontalAxis": &(ref_models.HorizontalAxis{}),

	"ref_models.HorizontalAxis.AxisHandleBorderLength": (ref_models.HorizontalAxis{}).AxisHandleBorderLength,

	"ref_models.HorizontalAxis.Axis_Length": (ref_models.HorizontalAxis{}).Axis_Length,

	"ref_models.HorizontalAxis.IsDisplayed": (ref_models.HorizontalAxis{}).IsDisplayed,

	"ref_models.HorizontalAxis.IsDisplayed": (ref_models.HorizontalAxis{}).IsDisplayed,

	"ref_models.HorizontalAxis.Name": (ref_models.HorizontalAxis{}).Name,

	"ref_models.HorizontalAxis.StrokeWidth": (ref_models.HorizontalAxis{}).StrokeWidth,

	"ref_models.HorizontalAxis.StrokeWidth": (ref_models.HorizontalAxis{}).StrokeWidth,

	"ref_models.InitialAxis": &(ref_models.InitialAxis{}),

	"ref_models.InitialAxis.Angle": (ref_models.InitialAxis{}).Angle,

	"ref_models.InitialAxis.IsDisplayed": (ref_models.InitialAxis{}).IsDisplayed,

	"ref_models.InitialAxis.Length": (ref_models.InitialAxis{}).Length,

	"ref_models.InitialAxis.Name": (ref_models.InitialAxis{}).Name,

	"ref_models.InitialAxis.StrokeWidth": (ref_models.InitialAxis{}).StrokeWidth,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.X1": (ref_models.Line{}).X1,

	"ref_models.Line.X2": (ref_models.Line{}).X2,

	"ref_models.Line.Y1": (ref_models.Line{}).Y1,

	"ref_models.Line.Y2": (ref_models.Line{}).Y2,

	"ref_models.Parameter": &(ref_models.Parameter{}),

	"ref_models.Parameter.DiamondSideLenght": (ref_models.Parameter{}).SideLength,

	"ref_models.Parameter.HorizontalAxis": (ref_models.Parameter{}).HorizontalAxis,

	"ref_models.Parameter.InitialAxis": (ref_models.Parameter{}).InitialAxis,

	"ref_models.Parameter.InitialRhombus": (ref_models.Parameter{}).InitialRhombus,

	"ref_models.Parameter.InitialRhombusGrid": (ref_models.Parameter{}).InitialRhombusGrid,

	"ref_models.Parameter.InsideAngle": (ref_models.Parameter{}).InsideAngle,

	"ref_models.Parameter.InsideAngle": (ref_models.Parameter{}).InsideAngle,

	"ref_models.Parameter.M": (ref_models.Parameter{}).M,

	"ref_models.Parameter.N": (ref_models.Parameter{}).N,

	"ref_models.Parameter.Name": (ref_models.Parameter{}).Name,

	"ref_models.Parameter.OriginX": (ref_models.Parameter{}).OriginX,

	"ref_models.Parameter.OriginY": (ref_models.Parameter{}).OriginY,

	"ref_models.Parameter.VerticalAxis": (ref_models.Parameter{}).VerticalAxis,

	"ref_models.Phylotaxy": ref_models.Phylotaxy,

	"ref_models.Rhombus": &(ref_models.Rhombus{}),

	"ref_models.Rhombus.Angle": (ref_models.Rhombus{}).Angle,

	"ref_models.Rhombus.CenterX": (ref_models.Rhombus{}).CenterX,

	"ref_models.Rhombus.CenterY": (ref_models.Rhombus{}).CenterY,

	"ref_models.Rhombus.InsideAngle": (ref_models.Rhombus{}).InsideAngle,

	"ref_models.Rhombus.IsDisplayed": (ref_models.Rhombus{}).IsDisplayed,

	"ref_models.Rhombus.Name": (ref_models.Rhombus{}).Name,

	"ref_models.Rhombus.SideLength": (ref_models.Rhombus{}).SideLength,

	"ref_models.Rhombus.StrokeWidth": (ref_models.Rhombus{}).StrokeWidth,

	"ref_models.RhombusGrid": &(ref_models.RhombusGrid{}),

	"ref_models.RhombusGrid.IsDisplayed": (ref_models.RhombusGrid{}).IsDisplayed,

	"ref_models.RhombusGrid.M": (ref_models.RhombusGrid{}).M,

	"ref_models.RhombusGrid.N": (ref_models.RhombusGrid{}).N,

	"ref_models.RhombusGrid.Name": (ref_models.RhombusGrid{}).Name,

	"ref_models.RhombusGrid.Reference": (ref_models.RhombusGrid{}).Reference,

	"ref_models.RhombusGrid.Rhombuses": (ref_models.RhombusGrid{}).Rhombuses,

	"ref_models.Sidebar": ref_models.Sidebar,

	"ref_models.SidebarTree": ref_models.SidebarTree,

	"ref_models.StacksNames": ref_models.StacksNames(""),

	"ref_models.TreeNames": ref_models.TreeNames(""),

	"ref_models.VerticalAxis": &(ref_models.VerticalAxis{}),

	"ref_models.VerticalAxis.AxisHandleBorderLength": (ref_models.VerticalAxis{}).AxisHandleBorderLength,

	"ref_models.VerticalAxis.Axis_Length": (ref_models.VerticalAxis{}).Axis_Length,

	"ref_models.VerticalAxis.IsDisplayed": (ref_models.VerticalAxis{}).IsDisplayed,

	"ref_models.VerticalAxis.IsDisplayed": (ref_models.VerticalAxis{}).IsDisplayed,

	"ref_models.VerticalAxis.Name": (ref_models.VerticalAxis{}).Name,

	"ref_models.VerticalAxis.StrokeWidth": (ref_models.VerticalAxis{}).StrokeWidth,

	"ref_models.VerticalAxis.StrokeWidth": (ref_models.VerticalAxis{}).StrokeWidth,
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_Angle := (&models.Field{Name: `Angle`}).Stage(stage)
	__Field__000001_AxisHandleBorderLength := (&models.Field{Name: `AxisHandleBorderLength`}).Stage(stage)
	__Field__000002_AxisHandleBorderLength := (&models.Field{Name: `AxisHandleBorderLength`}).Stage(stage)
	__Field__000003_Axis_Length := (&models.Field{Name: `Axis_Length`}).Stage(stage)
	__Field__000004_Axis_Length := (&models.Field{Name: `Axis_Length`}).Stage(stage)
	__Field__000005_Axis_StrokeWidth := (&models.Field{Name: `Axis_StrokeWidth`}).Stage(stage)
	__Field__000006_Axis_StrokeWidth := (&models.Field{Name: `Axis_StrokeWidth`}).Stage(stage)
	__Field__000007_CenterX := (&models.Field{Name: `CenterX`}).Stage(stage)
	__Field__000008_CenterY := (&models.Field{Name: `CenterY`}).Stage(stage)
	__Field__000009_DiamondAngle := (&models.Field{Name: `DiamondAngle`}).Stage(stage)
	__Field__000010_DiamondSideLenght := (&models.Field{Name: `DiamondSideLenght`}).Stage(stage)
	__Field__000011_InsideAngle := (&models.Field{Name: `InsideAngle`}).Stage(stage)
	__Field__000012_IsAxisDisplayed := (&models.Field{Name: `IsAxisDisplayed`}).Stage(stage)
	__Field__000013_IsAxisDisplayed := (&models.Field{Name: `IsAxisDisplayed`}).Stage(stage)
	__Field__000014_IsDisplayed := (&models.Field{Name: `IsDisplayed`}).Stage(stage)
	__Field__000015_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000016_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000017_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000018_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000019_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000020_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000021_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000022_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000023_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000024_OriginX := (&models.Field{Name: `OriginX`}).Stage(stage)
	__Field__000025_OriginY := (&models.Field{Name: `OriginY`}).Stage(stage)
	__Field__000026_SideLength := (&models.Field{Name: `SideLength`}).Stage(stage)

	__GongStructShape__000000_Default_HorizontalAxis := (&models.GongStructShape{Name: `Default-HorizontalAxis`}).Stage(stage)
	__GongStructShape__000001_Default_InitialAxis := (&models.GongStructShape{Name: `Default-InitialAxis`}).Stage(stage)
	__GongStructShape__000002_Default_Parameter := (&models.GongStructShape{Name: `Default-Parameter`}).Stage(stage)
	__GongStructShape__000003_Default_Rhombus := (&models.GongStructShape{Name: `Default-Rhombus`}).Stage(stage)
	__GongStructShape__000004_Default_RhombusGrid := (&models.GongStructShape{Name: `Default-RhombusGrid`}).Stage(stage)
	__GongStructShape__000005_Default_VerticalAxis := (&models.GongStructShape{Name: `Default-VerticalAxis`}).Stage(stage)

	__Link__000000_HorizontalAxis := (&models.Link{Name: `HorizontalAxis`}).Stage(stage)
	__Link__000001_InitialAxis := (&models.Link{Name: `InitialAxis`}).Stage(stage)
	__Link__000002_InitialRhombus := (&models.Link{Name: `InitialRhombus`}).Stage(stage)
	__Link__000003_InitialRhombusGrid := (&models.Link{Name: `InitialRhombusGrid`}).Stage(stage)
	__Link__000004_Reference := (&models.Link{Name: `Reference`}).Stage(stage)
	__Link__000005_Rhombuses := (&models.Link{Name: `Rhombuses`}).Stage(stage)
	__Link__000006_VerticalAxis := (&models.Link{Name: `VerticalAxis`}).Stage(stage)

	__Position__000000_Pos_Default_HorizontalAxis := (&models.Position{Name: `Pos-Default-HorizontalAxis`}).Stage(stage)
	__Position__000001_Pos_Default_InitialAxis := (&models.Position{Name: `Pos-Default-InitialAxis`}).Stage(stage)
	__Position__000002_Pos_Default_Parameter := (&models.Position{Name: `Pos-Default-Parameter`}).Stage(stage)
	__Position__000003_Pos_Default_Rhombus := (&models.Position{Name: `Pos-Default-Rhombus`}).Stage(stage)
	__Position__000004_Pos_Default_RhombusGrid := (&models.Position{Name: `Pos-Default-RhombusGrid`}).Stage(stage)
	__Position__000005_Pos_Default_VerticalAxis := (&models.Position{Name: `Pos-Default-VerticalAxis`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-HorizontalAxis`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-VerticalAxis`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`}).Stage(stage)

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

	__Field__000002_AxisHandleBorderLength.Name = `AxisHandleBorderLength`

	//gong:ident [ref_models.VerticalAxis.AxisHandleBorderLength] comment added to overcome the problem with the comment map association
	__Field__000002_AxisHandleBorderLength.Identifier = `ref_models.VerticalAxis.AxisHandleBorderLength`
	__Field__000002_AxisHandleBorderLength.FieldTypeAsString = ``
	__Field__000002_AxisHandleBorderLength.Structname = `VerticalAxis`
	__Field__000002_AxisHandleBorderLength.Fieldtypename = `float64`

	__Field__000003_Axis_Length.Name = `Axis_Length`

	//gong:ident [ref_models.HorizontalAxis.Axis_Length] comment added to overcome the problem with the comment map association
	__Field__000003_Axis_Length.Identifier = `ref_models.HorizontalAxis.Axis_Length`
	__Field__000003_Axis_Length.FieldTypeAsString = ``
	__Field__000003_Axis_Length.Structname = `HorizontalAxis`
	__Field__000003_Axis_Length.Fieldtypename = `float64`

	__Field__000004_Axis_Length.Name = `Axis_Length`

	//gong:ident [ref_models.VerticalAxis.Axis_Length] comment added to overcome the problem with the comment map association
	__Field__000004_Axis_Length.Identifier = `ref_models.VerticalAxis.Axis_Length`
	__Field__000004_Axis_Length.FieldTypeAsString = ``
	__Field__000004_Axis_Length.Structname = `VerticalAxis`
	__Field__000004_Axis_Length.Fieldtypename = `float64`

	__Field__000005_Axis_StrokeWidth.Name = `Axis_StrokeWidth`

	//gong:ident [ref_models.HorizontalAxis.StrokeWidth] comment added to overcome the problem with the comment map association
	__Field__000005_Axis_StrokeWidth.Identifier = `ref_models.HorizontalAxis.StrokeWidth`
	__Field__000005_Axis_StrokeWidth.FieldTypeAsString = ``
	__Field__000005_Axis_StrokeWidth.Structname = `HorizontalAxis`
	__Field__000005_Axis_StrokeWidth.Fieldtypename = `float64`

	__Field__000006_Axis_StrokeWidth.Name = `Axis_StrokeWidth`

	//gong:ident [ref_models.VerticalAxis.StrokeWidth] comment added to overcome the problem with the comment map association
	__Field__000006_Axis_StrokeWidth.Identifier = `ref_models.VerticalAxis.StrokeWidth`
	__Field__000006_Axis_StrokeWidth.FieldTypeAsString = ``
	__Field__000006_Axis_StrokeWidth.Structname = `VerticalAxis`
	__Field__000006_Axis_StrokeWidth.Fieldtypename = `float64`

	__Field__000007_CenterX.Name = `CenterX`

	//gong:ident [ref_models.Rhombus.CenterX] comment added to overcome the problem with the comment map association
	__Field__000007_CenterX.Identifier = `ref_models.Rhombus.CenterX`
	__Field__000007_CenterX.FieldTypeAsString = ``
	__Field__000007_CenterX.Structname = `Rhombus`
	__Field__000007_CenterX.Fieldtypename = `float64`

	__Field__000008_CenterY.Name = `CenterY`

	//gong:ident [ref_models.Rhombus.CenterY] comment added to overcome the problem with the comment map association
	__Field__000008_CenterY.Identifier = `ref_models.Rhombus.CenterY`
	__Field__000008_CenterY.FieldTypeAsString = ``
	__Field__000008_CenterY.Structname = `Rhombus`
	__Field__000008_CenterY.Fieldtypename = `float64`

	__Field__000009_DiamondAngle.Name = `DiamondAngle`

	//gong:ident [ref_models.Parameter.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000009_DiamondAngle.Identifier = `ref_models.Parameter.InsideAngle`
	__Field__000009_DiamondAngle.FieldTypeAsString = ``
	__Field__000009_DiamondAngle.Structname = `Parameter`
	__Field__000009_DiamondAngle.Fieldtypename = `float64`

	__Field__000010_DiamondSideLenght.Name = `DiamondSideLenght`

	//gong:ident [ref_models.Parameter.DiamondSideLenght] comment added to overcome the problem with the comment map association
	__Field__000010_DiamondSideLenght.Identifier = `ref_models.Parameter.DiamondSideLenght`
	__Field__000010_DiamondSideLenght.FieldTypeAsString = ``
	__Field__000010_DiamondSideLenght.Structname = `Parameter`
	__Field__000010_DiamondSideLenght.Fieldtypename = `float64`

	__Field__000011_InsideAngle.Name = `InsideAngle`

	//gong:ident [ref_models.Rhombus.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000011_InsideAngle.Identifier = `ref_models.Rhombus.InsideAngle`
	__Field__000011_InsideAngle.FieldTypeAsString = ``
	__Field__000011_InsideAngle.Structname = `Rhombus`
	__Field__000011_InsideAngle.Fieldtypename = `float64`

	__Field__000012_IsAxisDisplayed.Name = `IsAxisDisplayed`

	//gong:ident [ref_models.VerticalAxis.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000012_IsAxisDisplayed.Identifier = `ref_models.VerticalAxis.IsDisplayed`
	__Field__000012_IsAxisDisplayed.FieldTypeAsString = ``
	__Field__000012_IsAxisDisplayed.Structname = `VerticalAxis`
	__Field__000012_IsAxisDisplayed.Fieldtypename = `bool`

	__Field__000013_IsAxisDisplayed.Name = `IsAxisDisplayed`

	//gong:ident [ref_models.HorizontalAxis.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000013_IsAxisDisplayed.Identifier = `ref_models.HorizontalAxis.IsDisplayed`
	__Field__000013_IsAxisDisplayed.FieldTypeAsString = ``
	__Field__000013_IsAxisDisplayed.Structname = `HorizontalAxis`
	__Field__000013_IsAxisDisplayed.Fieldtypename = `bool`

	__Field__000014_IsDisplayed.Name = `IsDisplayed`

	//gong:ident [ref_models.Rhombus.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000014_IsDisplayed.Identifier = `ref_models.Rhombus.IsDisplayed`
	__Field__000014_IsDisplayed.FieldTypeAsString = ``
	__Field__000014_IsDisplayed.Structname = `Rhombus`
	__Field__000014_IsDisplayed.Fieldtypename = `bool`

	__Field__000015_M.Name = `M`

	//gong:ident [ref_models.Parameter.M] comment added to overcome the problem with the comment map association
	__Field__000015_M.Identifier = `ref_models.Parameter.M`
	__Field__000015_M.FieldTypeAsString = ``
	__Field__000015_M.Structname = `Parameter`
	__Field__000015_M.Fieldtypename = `int`

	__Field__000016_M.Name = `M`

	//gong:ident [ref_models.RhombusGrid.M] comment added to overcome the problem with the comment map association
	__Field__000016_M.Identifier = `ref_models.RhombusGrid.M`
	__Field__000016_M.FieldTypeAsString = ``
	__Field__000016_M.Structname = `RhombusGrid`
	__Field__000016_M.Fieldtypename = `int`

	__Field__000017_N.Name = `N`

	//gong:ident [ref_models.Parameter.N] comment added to overcome the problem with the comment map association
	__Field__000017_N.Identifier = `ref_models.Parameter.N`
	__Field__000017_N.FieldTypeAsString = ``
	__Field__000017_N.Structname = `Parameter`
	__Field__000017_N.Fieldtypename = `int`

	__Field__000018_N.Name = `N`

	//gong:ident [ref_models.RhombusGrid.N] comment added to overcome the problem with the comment map association
	__Field__000018_N.Identifier = `ref_models.RhombusGrid.N`
	__Field__000018_N.FieldTypeAsString = ``
	__Field__000018_N.Structname = `RhombusGrid`
	__Field__000018_N.Fieldtypename = `int`

	__Field__000019_Name.Name = `Name`

	//gong:ident [ref_models.Rhombus.Name] comment added to overcome the problem with the comment map association
	__Field__000019_Name.Identifier = `ref_models.Rhombus.Name`
	__Field__000019_Name.FieldTypeAsString = ``
	__Field__000019_Name.Structname = `Rhombus`
	__Field__000019_Name.Fieldtypename = `string`

	__Field__000020_Name.Name = `Name`

	//gong:ident [ref_models.HorizontalAxis.Name] comment added to overcome the problem with the comment map association
	__Field__000020_Name.Identifier = `ref_models.HorizontalAxis.Name`
	__Field__000020_Name.FieldTypeAsString = ``
	__Field__000020_Name.Structname = `HorizontalAxis`
	__Field__000020_Name.Fieldtypename = `string`

	__Field__000021_Name.Name = `Name`

	//gong:ident [ref_models.RhombusGrid.Name] comment added to overcome the problem with the comment map association
	__Field__000021_Name.Identifier = `ref_models.RhombusGrid.Name`
	__Field__000021_Name.FieldTypeAsString = ``
	__Field__000021_Name.Structname = `RhombusGrid`
	__Field__000021_Name.Fieldtypename = `string`

	__Field__000022_Name.Name = `Name`

	//gong:ident [ref_models.VerticalAxis.Name] comment added to overcome the problem with the comment map association
	__Field__000022_Name.Identifier = `ref_models.VerticalAxis.Name`
	__Field__000022_Name.FieldTypeAsString = ``
	__Field__000022_Name.Structname = `VerticalAxis`
	__Field__000022_Name.Fieldtypename = `string`

	__Field__000023_Name.Name = `Name`

	//gong:ident [ref_models.Parameter.Name] comment added to overcome the problem with the comment map association
	__Field__000023_Name.Identifier = `ref_models.Parameter.Name`
	__Field__000023_Name.FieldTypeAsString = ``
	__Field__000023_Name.Structname = `Parameter`
	__Field__000023_Name.Fieldtypename = `string`

	__Field__000024_OriginX.Name = `OriginX`

	//gong:ident [ref_models.Parameter.OriginX] comment added to overcome the problem with the comment map association
	__Field__000024_OriginX.Identifier = `ref_models.Parameter.OriginX`
	__Field__000024_OriginX.FieldTypeAsString = ``
	__Field__000024_OriginX.Structname = `Parameter`
	__Field__000024_OriginX.Fieldtypename = `float64`

	__Field__000025_OriginY.Name = `OriginY`

	//gong:ident [ref_models.Parameter.OriginY] comment added to overcome the problem with the comment map association
	__Field__000025_OriginY.Identifier = `ref_models.Parameter.OriginY`
	__Field__000025_OriginY.FieldTypeAsString = ``
	__Field__000025_OriginY.Structname = `Parameter`
	__Field__000025_OriginY.Fieldtypename = `float64`

	__Field__000026_SideLength.Name = `SideLength`

	//gong:ident [ref_models.Rhombus.SideLength] comment added to overcome the problem with the comment map association
	__Field__000026_SideLength.Identifier = `ref_models.Rhombus.SideLength`
	__Field__000026_SideLength.FieldTypeAsString = ``
	__Field__000026_SideLength.Structname = `Rhombus`
	__Field__000026_SideLength.Fieldtypename = `float64`

	__GongStructShape__000000_Default_HorizontalAxis.Name = `Default-HorizontalAxis`

	//gong:ident [ref_models.HorizontalAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_HorizontalAxis.Identifier = `ref_models.HorizontalAxis`
	__GongStructShape__000000_Default_HorizontalAxis.ShowNbInstances = false
	__GongStructShape__000000_Default_HorizontalAxis.NbInstances = 1
	__GongStructShape__000000_Default_HorizontalAxis.Width = 331.000000
	__GongStructShape__000000_Default_HorizontalAxis.Height = 138.000000
	__GongStructShape__000000_Default_HorizontalAxis.IsSelected = false

	__GongStructShape__000001_Default_InitialAxis.Name = `Default-InitialAxis`

	//gong:ident [ref_models.InitialAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_InitialAxis.Identifier = `ref_models.InitialAxis`
	__GongStructShape__000001_Default_InitialAxis.ShowNbInstances = false
	__GongStructShape__000001_Default_InitialAxis.NbInstances = 0
	__GongStructShape__000001_Default_InitialAxis.Width = 240.000000
	__GongStructShape__000001_Default_InitialAxis.Height = 63.000000
	__GongStructShape__000001_Default_InitialAxis.IsSelected = false

	__GongStructShape__000002_Default_Parameter.Name = `Default-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000002_Default_Parameter.ShowNbInstances = false
	__GongStructShape__000002_Default_Parameter.NbInstances = 1
	__GongStructShape__000002_Default_Parameter.Width = 240.000000
	__GongStructShape__000002_Default_Parameter.Height = 168.000000
	__GongStructShape__000002_Default_Parameter.IsSelected = false

	__GongStructShape__000003_Default_Rhombus.Name = `Default-Rhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Rhombus.Identifier = `ref_models.Rhombus`
	__GongStructShape__000003_Default_Rhombus.ShowNbInstances = false
	__GongStructShape__000003_Default_Rhombus.NbInstances = 7
	__GongStructShape__000003_Default_Rhombus.Width = 314.000000
	__GongStructShape__000003_Default_Rhombus.Height = 168.000000
	__GongStructShape__000003_Default_Rhombus.IsSelected = false

	__GongStructShape__000004_Default_RhombusGrid.Name = `Default-RhombusGrid`

	//gong:ident [ref_models.RhombusGrid] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_RhombusGrid.Identifier = `ref_models.RhombusGrid`
	__GongStructShape__000004_Default_RhombusGrid.ShowNbInstances = false
	__GongStructShape__000004_Default_RhombusGrid.NbInstances = 0
	__GongStructShape__000004_Default_RhombusGrid.Width = 307.000000
	__GongStructShape__000004_Default_RhombusGrid.Height = 108.000000
	__GongStructShape__000004_Default_RhombusGrid.IsSelected = false

	__GongStructShape__000005_Default_VerticalAxis.Name = `Default-VerticalAxis`

	//gong:ident [ref_models.VerticalAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_VerticalAxis.Identifier = `ref_models.VerticalAxis`
	__GongStructShape__000005_Default_VerticalAxis.ShowNbInstances = false
	__GongStructShape__000005_Default_VerticalAxis.NbInstances = 1
	__GongStructShape__000005_Default_VerticalAxis.Width = 337.000000
	__GongStructShape__000005_Default_VerticalAxis.Height = 138.000000
	__GongStructShape__000005_Default_VerticalAxis.IsSelected = false

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

	__Link__000001_InitialAxis.Name = `InitialAxis`

	//gong:ident [ref_models.Parameter.InitialAxis] comment added to overcome the problem with the comment map association
	__Link__000001_InitialAxis.Identifier = `ref_models.Parameter.InitialAxis`

	//gong:ident [ref_models.InitialAxis] comment added to overcome the problem with the comment map association
	__Link__000001_InitialAxis.Fieldtypename = `ref_models.InitialAxis`
	__Link__000001_InitialAxis.FieldOffsetX = -50.000000
	__Link__000001_InitialAxis.FieldOffsetY = -16.000000
	__Link__000001_InitialAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_InitialAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_InitialAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_InitialAxis.SourceMultiplicity = models.MANY
	__Link__000001_InitialAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_InitialAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_InitialAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_InitialAxis.StartRatio = 0.500000
	__Link__000001_InitialAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_InitialAxis.EndRatio = 0.500000
	__Link__000001_InitialAxis.CornerOffsetRatio = 1.380000

	__Link__000002_InitialRhombus.Name = `InitialRhombus`

	//gong:ident [ref_models.Parameter.InitialRhombus] comment added to overcome the problem with the comment map association
	__Link__000002_InitialRhombus.Identifier = `ref_models.Parameter.InitialRhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000002_InitialRhombus.Fieldtypename = `ref_models.Rhombus`
	__Link__000002_InitialRhombus.FieldOffsetX = -50.000000
	__Link__000002_InitialRhombus.FieldOffsetY = -16.000000
	__Link__000002_InitialRhombus.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_InitialRhombus.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_InitialRhombus.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_InitialRhombus.SourceMultiplicity = models.MANY
	__Link__000002_InitialRhombus.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_InitialRhombus.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_InitialRhombus.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_InitialRhombus.StartRatio = 0.494699
	__Link__000002_InitialRhombus.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_InitialRhombus.EndRatio = 0.500000
	__Link__000002_InitialRhombus.CornerOffsetRatio = 1.379109

	__Link__000003_InitialRhombusGrid.Name = `InitialRhombusGrid`

	//gong:ident [ref_models.Parameter.InitialRhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000003_InitialRhombusGrid.Identifier = `ref_models.Parameter.InitialRhombusGrid`

	//gong:ident [ref_models.RhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000003_InitialRhombusGrid.Fieldtypename = `ref_models.RhombusGrid`
	__Link__000003_InitialRhombusGrid.FieldOffsetX = -50.000000
	__Link__000003_InitialRhombusGrid.FieldOffsetY = -16.000000
	__Link__000003_InitialRhombusGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_InitialRhombusGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_InitialRhombusGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_InitialRhombusGrid.SourceMultiplicity = models.MANY
	__Link__000003_InitialRhombusGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_InitialRhombusGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_InitialRhombusGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_InitialRhombusGrid.StartRatio = 0.500000
	__Link__000003_InitialRhombusGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_InitialRhombusGrid.EndRatio = 0.500000
	__Link__000003_InitialRhombusGrid.CornerOffsetRatio = 1.380000

	__Link__000004_Reference.Name = `Reference`

	//gong:ident [ref_models.RhombusGrid.Reference] comment added to overcome the problem with the comment map association
	__Link__000004_Reference.Identifier = `ref_models.RhombusGrid.Reference`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000004_Reference.Fieldtypename = `ref_models.Rhombus`
	__Link__000004_Reference.FieldOffsetX = -50.000000
	__Link__000004_Reference.FieldOffsetY = -16.000000
	__Link__000004_Reference.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_Reference.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_Reference.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_Reference.SourceMultiplicity = models.MANY
	__Link__000004_Reference.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_Reference.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_Reference.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Reference.StartRatio = 0.500000
	__Link__000004_Reference.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Reference.EndRatio = 0.500000
	__Link__000004_Reference.CornerOffsetRatio = 1.380000

	__Link__000005_Rhombuses.Name = `Rhombuses`

	//gong:ident [ref_models.RhombusGrid.Rhombuses] comment added to overcome the problem with the comment map association
	__Link__000005_Rhombuses.Identifier = `ref_models.RhombusGrid.Rhombuses`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000005_Rhombuses.Fieldtypename = `ref_models.Rhombus`
	__Link__000005_Rhombuses.FieldOffsetX = -50.000000
	__Link__000005_Rhombuses.FieldOffsetY = -16.000000
	__Link__000005_Rhombuses.TargetMultiplicity = models.MANY
	__Link__000005_Rhombuses.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_Rhombuses.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_Rhombuses.SourceMultiplicity = models.MANY
	__Link__000005_Rhombuses.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_Rhombuses.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_Rhombuses.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_Rhombuses.StartRatio = 0.500000
	__Link__000005_Rhombuses.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_Rhombuses.EndRatio = 0.194434
	__Link__000005_Rhombuses.CornerOffsetRatio = 1.966609

	__Link__000006_VerticalAxis.Name = `VerticalAxis`

	//gong:ident [ref_models.Parameter.VerticalAxis] comment added to overcome the problem with the comment map association
	__Link__000006_VerticalAxis.Identifier = `ref_models.Parameter.VerticalAxis`

	//gong:ident [ref_models.VerticalAxis] comment added to overcome the problem with the comment map association
	__Link__000006_VerticalAxis.Fieldtypename = `ref_models.VerticalAxis`
	__Link__000006_VerticalAxis.FieldOffsetX = -50.000000
	__Link__000006_VerticalAxis.FieldOffsetY = -16.000000
	__Link__000006_VerticalAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_VerticalAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_VerticalAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_VerticalAxis.SourceMultiplicity = models.MANY
	__Link__000006_VerticalAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_VerticalAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_VerticalAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_VerticalAxis.StartRatio = 0.500000
	__Link__000006_VerticalAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_VerticalAxis.EndRatio = 0.500000
	__Link__000006_VerticalAxis.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_HorizontalAxis.X = 508.999969
	__Position__000000_Pos_Default_HorizontalAxis.Y = 27.000000
	__Position__000000_Pos_Default_HorizontalAxis.Name = `Pos-Default-HorizontalAxis`

	__Position__000001_Pos_Default_InitialAxis.X = 519.000000
	__Position__000001_Pos_Default_InitialAxis.Y = 878.666641
	__Position__000001_Pos_Default_InitialAxis.Name = `Pos-Default-InitialAxis`

	__Position__000002_Pos_Default_Parameter.X = 18.000000
	__Position__000002_Pos_Default_Parameter.Y = 35.000000
	__Position__000002_Pos_Default_Parameter.Name = `Pos-Default-Parameter`

	__Position__000003_Pos_Default_Rhombus.X = 516.999969
	__Position__000003_Pos_Default_Rhombus.Y = 426.000000
	__Position__000003_Pos_Default_Rhombus.Name = `Pos-Default-Rhombus`

	__Position__000004_Pos_Default_RhombusGrid.X = 517.000000
	__Position__000004_Pos_Default_RhombusGrid.Y = 695.999893
	__Position__000004_Pos_Default_RhombusGrid.Name = `Pos-Default-RhombusGrid`

	__Position__000005_Pos_Default_VerticalAxis.X = 504.999969
	__Position__000005_Pos_Default_VerticalAxis.Y = 197.000000
	__Position__000005_Pos_Default_VerticalAxis.Name = `Pos-Default-VerticalAxis`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.X = 623.499985
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Y = 115.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-HorizontalAxis`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.X = 625.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Y = 491.333321
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.X = 607.999985
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Y = 121.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.X = 628.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Y = 419.499977
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.X = 621.499985
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.Y = 200.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-VerticalAxis`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.X = 879.499984
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Y = 570.999977
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`

	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.X = 879.499984
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Y = 585.999977
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Parameter)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Rhombus)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_HorizontalAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_VerticalAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_RhombusGrid)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_InitialAxis)
	__GongStructShape__000000_Default_HorizontalAxis.Position = __Position__000000_Pos_Default_HorizontalAxis
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000020_Name)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000013_IsAxisDisplayed)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000001_AxisHandleBorderLength)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000003_Axis_Length)
	__GongStructShape__000000_Default_HorizontalAxis.Fields = append(__GongStructShape__000000_Default_HorizontalAxis.Fields, __Field__000005_Axis_StrokeWidth)
	__GongStructShape__000001_Default_InitialAxis.Position = __Position__000001_Pos_Default_InitialAxis
	__GongStructShape__000002_Default_Parameter.Position = __Position__000002_Pos_Default_Parameter
	__GongStructShape__000002_Default_Parameter.Fields = append(__GongStructShape__000002_Default_Parameter.Fields, __Field__000023_Name)
	__GongStructShape__000002_Default_Parameter.Fields = append(__GongStructShape__000002_Default_Parameter.Fields, __Field__000017_N)
	__GongStructShape__000002_Default_Parameter.Fields = append(__GongStructShape__000002_Default_Parameter.Fields, __Field__000015_M)
	__GongStructShape__000002_Default_Parameter.Fields = append(__GongStructShape__000002_Default_Parameter.Fields, __Field__000009_DiamondAngle)
	__GongStructShape__000002_Default_Parameter.Fields = append(__GongStructShape__000002_Default_Parameter.Fields, __Field__000024_OriginX)
	__GongStructShape__000002_Default_Parameter.Fields = append(__GongStructShape__000002_Default_Parameter.Fields, __Field__000025_OriginY)
	__GongStructShape__000002_Default_Parameter.Fields = append(__GongStructShape__000002_Default_Parameter.Fields, __Field__000010_DiamondSideLenght)
	__GongStructShape__000002_Default_Parameter.Links = append(__GongStructShape__000002_Default_Parameter.Links, __Link__000002_InitialRhombus)
	__GongStructShape__000002_Default_Parameter.Links = append(__GongStructShape__000002_Default_Parameter.Links, __Link__000000_HorizontalAxis)
	__GongStructShape__000002_Default_Parameter.Links = append(__GongStructShape__000002_Default_Parameter.Links, __Link__000006_VerticalAxis)
	__GongStructShape__000002_Default_Parameter.Links = append(__GongStructShape__000002_Default_Parameter.Links, __Link__000003_InitialRhombusGrid)
	__GongStructShape__000002_Default_Parameter.Links = append(__GongStructShape__000002_Default_Parameter.Links, __Link__000001_InitialAxis)
	__GongStructShape__000003_Default_Rhombus.Position = __Position__000003_Pos_Default_Rhombus
	__GongStructShape__000003_Default_Rhombus.Fields = append(__GongStructShape__000003_Default_Rhombus.Fields, __Field__000019_Name)
	__GongStructShape__000003_Default_Rhombus.Fields = append(__GongStructShape__000003_Default_Rhombus.Fields, __Field__000014_IsDisplayed)
	__GongStructShape__000003_Default_Rhombus.Fields = append(__GongStructShape__000003_Default_Rhombus.Fields, __Field__000007_CenterX)
	__GongStructShape__000003_Default_Rhombus.Fields = append(__GongStructShape__000003_Default_Rhombus.Fields, __Field__000008_CenterY)
	__GongStructShape__000003_Default_Rhombus.Fields = append(__GongStructShape__000003_Default_Rhombus.Fields, __Field__000026_SideLength)
	__GongStructShape__000003_Default_Rhombus.Fields = append(__GongStructShape__000003_Default_Rhombus.Fields, __Field__000000_Angle)
	__GongStructShape__000003_Default_Rhombus.Fields = append(__GongStructShape__000003_Default_Rhombus.Fields, __Field__000011_InsideAngle)
	__GongStructShape__000004_Default_RhombusGrid.Position = __Position__000004_Pos_Default_RhombusGrid
	__GongStructShape__000004_Default_RhombusGrid.Fields = append(__GongStructShape__000004_Default_RhombusGrid.Fields, __Field__000021_Name)
	__GongStructShape__000004_Default_RhombusGrid.Fields = append(__GongStructShape__000004_Default_RhombusGrid.Fields, __Field__000018_N)
	__GongStructShape__000004_Default_RhombusGrid.Fields = append(__GongStructShape__000004_Default_RhombusGrid.Fields, __Field__000016_M)
	__GongStructShape__000004_Default_RhombusGrid.Links = append(__GongStructShape__000004_Default_RhombusGrid.Links, __Link__000004_Reference)
	__GongStructShape__000004_Default_RhombusGrid.Links = append(__GongStructShape__000004_Default_RhombusGrid.Links, __Link__000005_Rhombuses)
	__GongStructShape__000005_Default_VerticalAxis.Position = __Position__000005_Pos_Default_VerticalAxis
	__GongStructShape__000005_Default_VerticalAxis.Fields = append(__GongStructShape__000005_Default_VerticalAxis.Fields, __Field__000022_Name)
	__GongStructShape__000005_Default_VerticalAxis.Fields = append(__GongStructShape__000005_Default_VerticalAxis.Fields, __Field__000012_IsAxisDisplayed)
	__GongStructShape__000005_Default_VerticalAxis.Fields = append(__GongStructShape__000005_Default_VerticalAxis.Fields, __Field__000002_AxisHandleBorderLength)
	__GongStructShape__000005_Default_VerticalAxis.Fields = append(__GongStructShape__000005_Default_VerticalAxis.Fields, __Field__000004_Axis_Length)
	__GongStructShape__000005_Default_VerticalAxis.Fields = append(__GongStructShape__000005_Default_VerticalAxis.Fields, __Field__000006_Axis_StrokeWidth)
	__Link__000000_HorizontalAxis.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis
	__Link__000001_InitialAxis.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis
	__Link__000002_InitialRhombus.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus
	__Link__000003_InitialRhombusGrid.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid
	__Link__000004_Reference.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus
	__Link__000005_Rhombuses.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus
	__Link__000006_VerticalAxis.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis
}
