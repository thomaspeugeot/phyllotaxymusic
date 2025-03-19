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

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_Name := (&models.Field{}).Stage(stage)
	__Field__000001_N := (&models.Field{}).Stage(stage)
	__Field__000002_M := (&models.Field{}).Stage(stage)
	__Field__000003_DiamondAngle := (&models.Field{}).Stage(stage)
	__Field__000004_OriginX := (&models.Field{}).Stage(stage)
	__Field__000005_OriginY := (&models.Field{}).Stage(stage)
	__Field__000006_DiamondSideLenght := (&models.Field{}).Stage(stage)
	__Field__000007_Name := (&models.Field{}).Stage(stage)
	__Field__000008_IsDisplayed := (&models.Field{}).Stage(stage)
	__Field__000009_CenterX := (&models.Field{}).Stage(stage)
	__Field__000010_CenterY := (&models.Field{}).Stage(stage)
	__Field__000011_SideLength := (&models.Field{}).Stage(stage)
	__Field__000012_Angle := (&models.Field{}).Stage(stage)
	__Field__000013_InsideAngle := (&models.Field{}).Stage(stage)
	__Field__000014_Name := (&models.Field{}).Stage(stage)
	__Field__000015_IsAxisDisplayed := (&models.Field{}).Stage(stage)
	__Field__000016_AxisHandleBorderLength := (&models.Field{}).Stage(stage)
	__Field__000017_Axis_Length := (&models.Field{}).Stage(stage)
	__Field__000018_Axis_StrokeWidth := (&models.Field{}).Stage(stage)
	__Field__000019_Name := (&models.Field{}).Stage(stage)
	__Field__000020_IsAxisDisplayed := (&models.Field{}).Stage(stage)
	__Field__000021_AxisHandleBorderLength := (&models.Field{}).Stage(stage)
	__Field__000022_Axis_Length := (&models.Field{}).Stage(stage)
	__Field__000023_Axis_StrokeWidth := (&models.Field{}).Stage(stage)
	__Field__000024_Name := (&models.Field{}).Stage(stage)
	__Field__000025_N := (&models.Field{}).Stage(stage)
	__Field__000026_M := (&models.Field{}).Stage(stage)
	__Field__000027_Name := (&models.Field{}).Stage(stage)
	__Field__000028_Angle := (&models.Field{}).Stage(stage)
	__Field__000029_Length := (&models.Field{}).Stage(stage)
	__Field__000030_N := (&models.Field{}).Stage(stage)
	__Field__000031_M := (&models.Field{}).Stage(stage)
	__Field__000032_CenterX := (&models.Field{}).Stage(stage)
	__Field__000033_CenterY := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_Parameter := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Rhombus := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_HorizontalAxis := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_VerticalAxis := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_RhombusGrid := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Default_InitialAxis := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000006_Default_CircleGrid := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000007_Default_Circle := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_InitialRhombus := (&models.Link{}).Stage(stage)
	__Link__000001_HorizontalAxis := (&models.Link{}).Stage(stage)
	__Link__000002_VerticalAxis := (&models.Link{}).Stage(stage)
	__Link__000003_InitialRhombusGrid := (&models.Link{}).Stage(stage)
	__Link__000004_InitialAxis := (&models.Link{}).Stage(stage)
	__Link__000005_InitialCircle := (&models.Link{}).Stage(stage)
	__Link__000006_InitialCircleGrid := (&models.Link{}).Stage(stage)
	__Link__000007_RotatedAxis := (&models.Link{}).Stage(stage)
	__Link__000008_RotatedRhombus := (&models.Link{}).Stage(stage)
	__Link__000009_RotatedRhombusGrid := (&models.Link{}).Stage(stage)
	__Link__000010_RotatedCircleGrid := (&models.Link{}).Stage(stage)
	__Link__000011_Reference := (&models.Link{}).Stage(stage)
	__Link__000012_Rhombuses := (&models.Link{}).Stage(stage)
	__Link__000013_Reference := (&models.Link{}).Stage(stage)
	__Link__000014_Circles := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_Parameter := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Rhombus := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_HorizontalAxis := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_VerticalAxis := (&models.Position{}).Stage(stage)
	__Position__000004_Pos_Default_RhombusGrid := (&models.Position{}).Stage(stage)
	__Position__000005_Pos_Default_InitialAxis := (&models.Position{}).Stage(stage)
	__Position__000006_Pos_Default_CircleGrid := (&models.Position{}).Stage(stage)
	__Position__000007_Pos_Default_Circle := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis := (&models.Vertice{}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis := (&models.Vertice{}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid := (&models.Vertice{}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis := (&models.Vertice{}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle := (&models.Vertice{}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid := (&models.Vertice{}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis := (&models.Vertice{}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus := (&models.Vertice{}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid := (&models.Vertice{}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid := (&models.Vertice{}).Stage(stage)
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus := (&models.Vertice{}).Stage(stage)
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus := (&models.Vertice{}).Stage(stage)
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle := (&models.Vertice{}).Stage(stage)
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.Parameter.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.Parameter.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Parameter`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_N.Name = `N`

	//gong:ident [ref_models.Parameter.N] comment added to overcome the problem with the comment map association
	__Field__000001_N.Identifier = `ref_models.Parameter.N`
	__Field__000001_N.FieldTypeAsString = ``
	__Field__000001_N.Structname = `Parameter`
	__Field__000001_N.Fieldtypename = `int`

	__Field__000002_M.Name = `M`

	//gong:ident [ref_models.Parameter.M] comment added to overcome the problem with the comment map association
	__Field__000002_M.Identifier = `ref_models.Parameter.M`
	__Field__000002_M.FieldTypeAsString = ``
	__Field__000002_M.Structname = `Parameter`
	__Field__000002_M.Fieldtypename = `int`

	__Field__000003_DiamondAngle.Name = `DiamondAngle`

	//gong:ident [ref_models.Parameter.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000003_DiamondAngle.Identifier = `ref_models.Parameter.InsideAngle`
	__Field__000003_DiamondAngle.FieldTypeAsString = ``
	__Field__000003_DiamondAngle.Structname = `Parameter`
	__Field__000003_DiamondAngle.Fieldtypename = `float64`

	__Field__000004_OriginX.Name = `OriginX`

	//gong:ident [ref_models.Parameter.OriginX] comment added to overcome the problem with the comment map association
	__Field__000004_OriginX.Identifier = `ref_models.Parameter.OriginX`
	__Field__000004_OriginX.FieldTypeAsString = ``
	__Field__000004_OriginX.Structname = `Parameter`
	__Field__000004_OriginX.Fieldtypename = `float64`

	__Field__000005_OriginY.Name = `OriginY`

	//gong:ident [ref_models.Parameter.OriginY] comment added to overcome the problem with the comment map association
	__Field__000005_OriginY.Identifier = `ref_models.Parameter.OriginY`
	__Field__000005_OriginY.FieldTypeAsString = ``
	__Field__000005_OriginY.Structname = `Parameter`
	__Field__000005_OriginY.Fieldtypename = `float64`

	__Field__000006_DiamondSideLenght.Name = `DiamondSideLenght`

	//gong:ident [ref_models.Parameter.SideLength] comment added to overcome the problem with the comment map association
	__Field__000006_DiamondSideLenght.Identifier = `ref_models.Parameter.SideLength`
	__Field__000006_DiamondSideLenght.FieldTypeAsString = ``
	__Field__000006_DiamondSideLenght.Structname = `Parameter`
	__Field__000006_DiamondSideLenght.Fieldtypename = `float64`

	__Field__000007_Name.Name = `Name`

	//gong:ident [ref_models.Rhombus.Name] comment added to overcome the problem with the comment map association
	__Field__000007_Name.Identifier = `ref_models.Rhombus.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = `Rhombus`
	__Field__000007_Name.Fieldtypename = `string`

	__Field__000008_IsDisplayed.Name = `IsDisplayed`

	//gong:ident [ref_models.Rhombus.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000008_IsDisplayed.Identifier = `ref_models.Rhombus.IsDisplayed`
	__Field__000008_IsDisplayed.FieldTypeAsString = ``
	__Field__000008_IsDisplayed.Structname = `Rhombus`
	__Field__000008_IsDisplayed.Fieldtypename = `bool`

	__Field__000009_CenterX.Name = `CenterX`

	//gong:ident [ref_models.Rhombus.CenterX] comment added to overcome the problem with the comment map association
	__Field__000009_CenterX.Identifier = `ref_models.Rhombus.CenterX`
	__Field__000009_CenterX.FieldTypeAsString = ``
	__Field__000009_CenterX.Structname = `Rhombus`
	__Field__000009_CenterX.Fieldtypename = `float64`

	__Field__000010_CenterY.Name = `CenterY`

	//gong:ident [ref_models.Rhombus.CenterY] comment added to overcome the problem with the comment map association
	__Field__000010_CenterY.Identifier = `ref_models.Rhombus.CenterY`
	__Field__000010_CenterY.FieldTypeAsString = ``
	__Field__000010_CenterY.Structname = `Rhombus`
	__Field__000010_CenterY.Fieldtypename = `float64`

	__Field__000011_SideLength.Name = `SideLength`

	//gong:ident [ref_models.Rhombus.SideLength] comment added to overcome the problem with the comment map association
	__Field__000011_SideLength.Identifier = `ref_models.Rhombus.SideLength`
	__Field__000011_SideLength.FieldTypeAsString = ``
	__Field__000011_SideLength.Structname = `Rhombus`
	__Field__000011_SideLength.Fieldtypename = `float64`

	__Field__000012_Angle.Name = `Angle`

	//gong:ident [ref_models.Rhombus.Angle] comment added to overcome the problem with the comment map association
	__Field__000012_Angle.Identifier = `ref_models.Rhombus.Angle`
	__Field__000012_Angle.FieldTypeAsString = ``
	__Field__000012_Angle.Structname = `Rhombus`
	__Field__000012_Angle.Fieldtypename = `float64`

	__Field__000013_InsideAngle.Name = `InsideAngle`

	//gong:ident [ref_models.Rhombus.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000013_InsideAngle.Identifier = `ref_models.Rhombus.InsideAngle`
	__Field__000013_InsideAngle.FieldTypeAsString = ``
	__Field__000013_InsideAngle.Structname = `Rhombus`
	__Field__000013_InsideAngle.Fieldtypename = `float64`

	__Field__000014_Name.Name = `Name`

	//gong:ident [ref_models.HorizontalAxis.Name] comment added to overcome the problem with the comment map association
	__Field__000014_Name.Identifier = `ref_models.HorizontalAxis.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `HorizontalAxis`
	__Field__000014_Name.Fieldtypename = `string`

	__Field__000015_IsAxisDisplayed.Name = `IsAxisDisplayed`

	//gong:ident [ref_models.HorizontalAxis.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000015_IsAxisDisplayed.Identifier = `ref_models.HorizontalAxis.IsDisplayed`
	__Field__000015_IsAxisDisplayed.FieldTypeAsString = ``
	__Field__000015_IsAxisDisplayed.Structname = `HorizontalAxis`
	__Field__000015_IsAxisDisplayed.Fieldtypename = `bool`

	__Field__000016_AxisHandleBorderLength.Name = `AxisHandleBorderLength`

	//gong:ident [ref_models.HorizontalAxis.AxisHandleBorderLength] comment added to overcome the problem with the comment map association
	__Field__000016_AxisHandleBorderLength.Identifier = `ref_models.HorizontalAxis.AxisHandleBorderLength`
	__Field__000016_AxisHandleBorderLength.FieldTypeAsString = ``
	__Field__000016_AxisHandleBorderLength.Structname = `HorizontalAxis`
	__Field__000016_AxisHandleBorderLength.Fieldtypename = `float64`

	__Field__000017_Axis_Length.Name = `Axis_Length`

	//gong:ident [ref_models.HorizontalAxis.Axis_Length] comment added to overcome the problem with the comment map association
	__Field__000017_Axis_Length.Identifier = `ref_models.HorizontalAxis.Axis_Length`
	__Field__000017_Axis_Length.FieldTypeAsString = ``
	__Field__000017_Axis_Length.Structname = `HorizontalAxis`
	__Field__000017_Axis_Length.Fieldtypename = `float64`

	__Field__000018_Axis_StrokeWidth.Name = `Axis_StrokeWidth`

	//gong:ident [ref_models.HorizontalAxis.StrokeWidth] comment added to overcome the problem with the comment map association
	__Field__000018_Axis_StrokeWidth.Identifier = `ref_models.HorizontalAxis.StrokeWidth`
	__Field__000018_Axis_StrokeWidth.FieldTypeAsString = ``
	__Field__000018_Axis_StrokeWidth.Structname = `HorizontalAxis`
	__Field__000018_Axis_StrokeWidth.Fieldtypename = `float64`

	__Field__000019_Name.Name = `Name`

	//gong:ident [ref_models.VerticalAxis.Name] comment added to overcome the problem with the comment map association
	__Field__000019_Name.Identifier = `ref_models.VerticalAxis.Name`
	__Field__000019_Name.FieldTypeAsString = ``
	__Field__000019_Name.Structname = `VerticalAxis`
	__Field__000019_Name.Fieldtypename = `string`

	__Field__000020_IsAxisDisplayed.Name = `IsAxisDisplayed`

	//gong:ident [ref_models.VerticalAxis.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000020_IsAxisDisplayed.Identifier = `ref_models.VerticalAxis.IsDisplayed`
	__Field__000020_IsAxisDisplayed.FieldTypeAsString = ``
	__Field__000020_IsAxisDisplayed.Structname = `VerticalAxis`
	__Field__000020_IsAxisDisplayed.Fieldtypename = `bool`

	__Field__000021_AxisHandleBorderLength.Name = `AxisHandleBorderLength`

	//gong:ident [ref_models.VerticalAxis.AxisHandleBorderLength] comment added to overcome the problem with the comment map association
	__Field__000021_AxisHandleBorderLength.Identifier = `ref_models.VerticalAxis.AxisHandleBorderLength`
	__Field__000021_AxisHandleBorderLength.FieldTypeAsString = ``
	__Field__000021_AxisHandleBorderLength.Structname = `VerticalAxis`
	__Field__000021_AxisHandleBorderLength.Fieldtypename = `float64`

	__Field__000022_Axis_Length.Name = `Axis_Length`

	//gong:ident [ref_models.VerticalAxis.Axis_Length] comment added to overcome the problem with the comment map association
	__Field__000022_Axis_Length.Identifier = `ref_models.VerticalAxis.Axis_Length`
	__Field__000022_Axis_Length.FieldTypeAsString = ``
	__Field__000022_Axis_Length.Structname = `VerticalAxis`
	__Field__000022_Axis_Length.Fieldtypename = `float64`

	__Field__000023_Axis_StrokeWidth.Name = `Axis_StrokeWidth`

	//gong:ident [ref_models.VerticalAxis.StrokeWidth] comment added to overcome the problem with the comment map association
	__Field__000023_Axis_StrokeWidth.Identifier = `ref_models.VerticalAxis.StrokeWidth`
	__Field__000023_Axis_StrokeWidth.FieldTypeAsString = ``
	__Field__000023_Axis_StrokeWidth.Structname = `VerticalAxis`
	__Field__000023_Axis_StrokeWidth.Fieldtypename = `float64`

	__Field__000024_Name.Name = `Name`

	//gong:ident [ref_models.RhombusGrid.Name] comment added to overcome the problem with the comment map association
	__Field__000024_Name.Identifier = `ref_models.RhombusGrid.Name`
	__Field__000024_Name.FieldTypeAsString = ``
	__Field__000024_Name.Structname = `RhombusGrid`
	__Field__000024_Name.Fieldtypename = `string`

	__Field__000025_N.Name = `N`

	//gong:ident [ref_models.RhombusGrid.N] comment added to overcome the problem with the comment map association
	__Field__000025_N.Identifier = `ref_models.RhombusGrid.N`
	__Field__000025_N.FieldTypeAsString = ``
	__Field__000025_N.Structname = `RhombusGrid`
	__Field__000025_N.Fieldtypename = `int`

	__Field__000026_M.Name = `M`

	//gong:ident [ref_models.RhombusGrid.M] comment added to overcome the problem with the comment map association
	__Field__000026_M.Identifier = `ref_models.RhombusGrid.M`
	__Field__000026_M.FieldTypeAsString = ``
	__Field__000026_M.Structname = `RhombusGrid`
	__Field__000026_M.Fieldtypename = `int`

	__Field__000027_Name.Name = `Name`

	//gong:ident [ref_models.Axis.Name] comment added to overcome the problem with the comment map association
	__Field__000027_Name.Identifier = `ref_models.Axis.Name`
	__Field__000027_Name.FieldTypeAsString = ``
	__Field__000027_Name.Structname = `Axis`
	__Field__000027_Name.Fieldtypename = `string`

	__Field__000028_Angle.Name = `Angle`

	//gong:ident [ref_models.Axis.Angle] comment added to overcome the problem with the comment map association
	__Field__000028_Angle.Identifier = `ref_models.Axis.Angle`
	__Field__000028_Angle.FieldTypeAsString = ``
	__Field__000028_Angle.Structname = `Axis`
	__Field__000028_Angle.Fieldtypename = `float64`

	__Field__000029_Length.Name = `Length`

	//gong:ident [ref_models.Axis.Length] comment added to overcome the problem with the comment map association
	__Field__000029_Length.Identifier = `ref_models.Axis.Length`
	__Field__000029_Length.FieldTypeAsString = ``
	__Field__000029_Length.Structname = `Axis`
	__Field__000029_Length.Fieldtypename = `float64`

	__Field__000030_N.Name = `N`

	//gong:ident [ref_models.CircleGrid.N] comment added to overcome the problem with the comment map association
	__Field__000030_N.Identifier = `ref_models.CircleGrid.N`
	__Field__000030_N.FieldTypeAsString = ``
	__Field__000030_N.Structname = `CircleGrid`
	__Field__000030_N.Fieldtypename = `int`

	__Field__000031_M.Name = `M`

	//gong:ident [ref_models.CircleGrid.M] comment added to overcome the problem with the comment map association
	__Field__000031_M.Identifier = `ref_models.CircleGrid.M`
	__Field__000031_M.FieldTypeAsString = ``
	__Field__000031_M.Structname = `CircleGrid`
	__Field__000031_M.Fieldtypename = `int`

	__Field__000032_CenterX.Name = `CenterX`

	//gong:ident [ref_models.Circle.CenterX] comment added to overcome the problem with the comment map association
	__Field__000032_CenterX.Identifier = `ref_models.Circle.CenterX`
	__Field__000032_CenterX.FieldTypeAsString = ``
	__Field__000032_CenterX.Structname = `Circle`
	__Field__000032_CenterX.Fieldtypename = `float64`

	__Field__000033_CenterY.Name = `CenterY`

	//gong:ident [ref_models.Circle.CenterY] comment added to overcome the problem with the comment map association
	__Field__000033_CenterY.Identifier = `ref_models.Circle.CenterY`
	__Field__000033_CenterY.FieldTypeAsString = ``
	__Field__000033_CenterY.Structname = `Circle`
	__Field__000033_CenterY.Fieldtypename = `float64`

	__GongStructShape__000000_Default_Parameter.Name = `Default-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000000_Default_Parameter.ShowNbInstances = false
	__GongStructShape__000000_Default_Parameter.NbInstances = 1
	__GongStructShape__000000_Default_Parameter.Width = 240.000000
	__GongStructShape__000000_Default_Parameter.Height = 168.000000
	__GongStructShape__000000_Default_Parameter.IsSelected = false

	__GongStructShape__000001_Default_Rhombus.Name = `Default-Rhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Rhombus.Identifier = `ref_models.Rhombus`
	__GongStructShape__000001_Default_Rhombus.ShowNbInstances = false
	__GongStructShape__000001_Default_Rhombus.NbInstances = 4
	__GongStructShape__000001_Default_Rhombus.Width = 314.000000
	__GongStructShape__000001_Default_Rhombus.Height = 168.000000
	__GongStructShape__000001_Default_Rhombus.IsSelected = false

	__GongStructShape__000002_Default_HorizontalAxis.Name = `Default-HorizontalAxis`

	//gong:ident [ref_models.HorizontalAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_HorizontalAxis.Identifier = `ref_models.HorizontalAxis`
	__GongStructShape__000002_Default_HorizontalAxis.ShowNbInstances = false
	__GongStructShape__000002_Default_HorizontalAxis.NbInstances = 1
	__GongStructShape__000002_Default_HorizontalAxis.Width = 331.000000
	__GongStructShape__000002_Default_HorizontalAxis.Height = 138.000000
	__GongStructShape__000002_Default_HorizontalAxis.IsSelected = false

	__GongStructShape__000003_Default_VerticalAxis.Name = `Default-VerticalAxis`

	//gong:ident [ref_models.VerticalAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_VerticalAxis.Identifier = `ref_models.VerticalAxis`
	__GongStructShape__000003_Default_VerticalAxis.ShowNbInstances = false
	__GongStructShape__000003_Default_VerticalAxis.NbInstances = 1
	__GongStructShape__000003_Default_VerticalAxis.Width = 337.000000
	__GongStructShape__000003_Default_VerticalAxis.Height = 138.000000
	__GongStructShape__000003_Default_VerticalAxis.IsSelected = false

	__GongStructShape__000004_Default_RhombusGrid.Name = `Default-RhombusGrid`

	//gong:ident [ref_models.RhombusGrid] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_RhombusGrid.Identifier = `ref_models.RhombusGrid`
	__GongStructShape__000004_Default_RhombusGrid.ShowNbInstances = false
	__GongStructShape__000004_Default_RhombusGrid.NbInstances = 3
	__GongStructShape__000004_Default_RhombusGrid.Width = 307.000000
	__GongStructShape__000004_Default_RhombusGrid.Height = 125.000000
	__GongStructShape__000004_Default_RhombusGrid.IsSelected = false

	__GongStructShape__000005_Default_InitialAxis.Name = `Default-InitialAxis`

	//gong:ident [ref_models.Axis] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_InitialAxis.Identifier = `ref_models.Axis`
	__GongStructShape__000005_Default_InitialAxis.ShowNbInstances = false
	__GongStructShape__000005_Default_InitialAxis.NbInstances = 5
	__GongStructShape__000005_Default_InitialAxis.Width = 240.000000
	__GongStructShape__000005_Default_InitialAxis.Height = 214.000000
	__GongStructShape__000005_Default_InitialAxis.IsSelected = false

	__GongStructShape__000006_Default_CircleGrid.Name = `Default-CircleGrid`

	//gong:ident [ref_models.CircleGrid] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_CircleGrid.Identifier = `ref_models.CircleGrid`
	__GongStructShape__000006_Default_CircleGrid.ShowNbInstances = false
	__GongStructShape__000006_Default_CircleGrid.NbInstances = 9
	__GongStructShape__000006_Default_CircleGrid.Width = 240.000000
	__GongStructShape__000006_Default_CircleGrid.Height = 164.000061
	__GongStructShape__000006_Default_CircleGrid.IsSelected = false

	__GongStructShape__000007_Default_Circle.Name = `Default-Circle`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_Default_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000007_Default_Circle.ShowNbInstances = false
	__GongStructShape__000007_Default_Circle.NbInstances = 6
	__GongStructShape__000007_Default_Circle.Width = 240.000000
	__GongStructShape__000007_Default_Circle.Height = 93.000000
	__GongStructShape__000007_Default_Circle.IsSelected = false

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
	__Link__000000_InitialRhombus.StartRatio = 0.494699
	__Link__000000_InitialRhombus.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_InitialRhombus.EndRatio = 0.500000
	__Link__000000_InitialRhombus.CornerOffsetRatio = 1.379109

	__Link__000001_HorizontalAxis.Name = `HorizontalAxis`

	//gong:ident [ref_models.Parameter.HorizontalAxis] comment added to overcome the problem with the comment map association
	__Link__000001_HorizontalAxis.Identifier = `ref_models.Parameter.HorizontalAxis`

	//gong:ident [ref_models.HorizontalAxis] comment added to overcome the problem with the comment map association
	__Link__000001_HorizontalAxis.Fieldtypename = `ref_models.HorizontalAxis`
	__Link__000001_HorizontalAxis.FieldOffsetX = -50.000000
	__Link__000001_HorizontalAxis.FieldOffsetY = -16.000000
	__Link__000001_HorizontalAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_HorizontalAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_HorizontalAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_HorizontalAxis.SourceMultiplicity = models.MANY
	__Link__000001_HorizontalAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_HorizontalAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_HorizontalAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_HorizontalAxis.StartRatio = 0.500000
	__Link__000001_HorizontalAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_HorizontalAxis.EndRatio = 0.500000
	__Link__000001_HorizontalAxis.CornerOffsetRatio = 1.380000

	__Link__000002_VerticalAxis.Name = `VerticalAxis`

	//gong:ident [ref_models.Parameter.VerticalAxis] comment added to overcome the problem with the comment map association
	__Link__000002_VerticalAxis.Identifier = `ref_models.Parameter.VerticalAxis`

	//gong:ident [ref_models.VerticalAxis] comment added to overcome the problem with the comment map association
	__Link__000002_VerticalAxis.Fieldtypename = `ref_models.VerticalAxis`
	__Link__000002_VerticalAxis.FieldOffsetX = -50.000000
	__Link__000002_VerticalAxis.FieldOffsetY = -16.000000
	__Link__000002_VerticalAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_VerticalAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_VerticalAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_VerticalAxis.SourceMultiplicity = models.MANY
	__Link__000002_VerticalAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_VerticalAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_VerticalAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_VerticalAxis.StartRatio = 0.500000
	__Link__000002_VerticalAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_VerticalAxis.EndRatio = 0.500000
	__Link__000002_VerticalAxis.CornerOffsetRatio = 1.380000

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
	__Link__000003_InitialRhombusGrid.EndRatio = 0.252316
	__Link__000003_InitialRhombusGrid.CornerOffsetRatio = 1.380000

	__Link__000004_InitialAxis.Name = `InitialAxis`

	//gong:ident [ref_models.Parameter.InitialAxis] comment added to overcome the problem with the comment map association
	__Link__000004_InitialAxis.Identifier = `ref_models.Parameter.InitialAxis`

	//gong:ident [ref_models.Axis] comment added to overcome the problem with the comment map association
	__Link__000004_InitialAxis.Fieldtypename = `ref_models.Axis`
	__Link__000004_InitialAxis.FieldOffsetX = -50.000000
	__Link__000004_InitialAxis.FieldOffsetY = -16.000000
	__Link__000004_InitialAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_InitialAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_InitialAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_InitialAxis.SourceMultiplicity = models.MANY
	__Link__000004_InitialAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_InitialAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_InitialAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_InitialAxis.StartRatio = 0.500000
	__Link__000004_InitialAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_InitialAxis.EndRatio = 0.309375
	__Link__000004_InitialAxis.CornerOffsetRatio = 1.380000

	__Link__000005_InitialCircle.Name = `InitialCircle`

	//gong:ident [ref_models.Parameter.InitialCircle] comment added to overcome the problem with the comment map association
	__Link__000005_InitialCircle.Identifier = `ref_models.Parameter.InitialCircle`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__Link__000005_InitialCircle.Fieldtypename = `ref_models.Circle`
	__Link__000005_InitialCircle.FieldOffsetX = -50.000000
	__Link__000005_InitialCircle.FieldOffsetY = -16.000000
	__Link__000005_InitialCircle.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_InitialCircle.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_InitialCircle.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_InitialCircle.SourceMultiplicity = models.MANY
	__Link__000005_InitialCircle.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_InitialCircle.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_InitialCircle.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_InitialCircle.StartRatio = 0.500000
	__Link__000005_InitialCircle.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_InitialCircle.EndRatio = 0.500000
	__Link__000005_InitialCircle.CornerOffsetRatio = 1.380000

	__Link__000006_InitialCircleGrid.Name = `InitialCircleGrid`

	//gong:ident [ref_models.Parameter.InitialCircleGrid] comment added to overcome the problem with the comment map association
	__Link__000006_InitialCircleGrid.Identifier = `ref_models.Parameter.InitialCircleGrid`

	//gong:ident [ref_models.CircleGrid] comment added to overcome the problem with the comment map association
	__Link__000006_InitialCircleGrid.Fieldtypename = `ref_models.CircleGrid`
	__Link__000006_InitialCircleGrid.FieldOffsetX = -50.000000
	__Link__000006_InitialCircleGrid.FieldOffsetY = -16.000000
	__Link__000006_InitialCircleGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_InitialCircleGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_InitialCircleGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_InitialCircleGrid.SourceMultiplicity = models.MANY
	__Link__000006_InitialCircleGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_InitialCircleGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_InitialCircleGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_InitialCircleGrid.StartRatio = 0.500000
	__Link__000006_InitialCircleGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_InitialCircleGrid.EndRatio = 0.273614
	__Link__000006_InitialCircleGrid.CornerOffsetRatio = 1.380000

	__Link__000007_RotatedAxis.Name = `RotatedAxis`

	//gong:ident [ref_models.Parameter.RotatedAxis] comment added to overcome the problem with the comment map association
	__Link__000007_RotatedAxis.Identifier = `ref_models.Parameter.RotatedAxis`

	//gong:ident [ref_models.Axis] comment added to overcome the problem with the comment map association
	__Link__000007_RotatedAxis.Fieldtypename = `ref_models.Axis`
	__Link__000007_RotatedAxis.FieldOffsetX = -50.000000
	__Link__000007_RotatedAxis.FieldOffsetY = -16.000000
	__Link__000007_RotatedAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000007_RotatedAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_RotatedAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_RotatedAxis.SourceMultiplicity = models.MANY
	__Link__000007_RotatedAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000007_RotatedAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000007_RotatedAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_RotatedAxis.StartRatio = 0.500000
	__Link__000007_RotatedAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_RotatedAxis.EndRatio = 0.818720
	__Link__000007_RotatedAxis.CornerOffsetRatio = 1.380000

	__Link__000008_RotatedRhombus.Name = `RotatedRhombus`

	//gong:ident [ref_models.Parameter.RotatedRhombus] comment added to overcome the problem with the comment map association
	__Link__000008_RotatedRhombus.Identifier = `ref_models.Parameter.RotatedRhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000008_RotatedRhombus.Fieldtypename = `ref_models.Rhombus`
	__Link__000008_RotatedRhombus.FieldOffsetX = -50.000000
	__Link__000008_RotatedRhombus.FieldOffsetY = -16.000000
	__Link__000008_RotatedRhombus.TargetMultiplicity = models.ZERO_ONE
	__Link__000008_RotatedRhombus.TargetMultiplicityOffsetX = -50.000000
	__Link__000008_RotatedRhombus.TargetMultiplicityOffsetY = 16.000000
	__Link__000008_RotatedRhombus.SourceMultiplicity = models.MANY
	__Link__000008_RotatedRhombus.SourceMultiplicityOffsetX = 10.000000
	__Link__000008_RotatedRhombus.SourceMultiplicityOffsetY = -50.000000
	__Link__000008_RotatedRhombus.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_RotatedRhombus.StartRatio = 0.500000
	__Link__000008_RotatedRhombus.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_RotatedRhombus.EndRatio = 0.500000
	__Link__000008_RotatedRhombus.CornerOffsetRatio = 1.380000

	__Link__000009_RotatedRhombusGrid.Name = `RotatedRhombusGrid`

	//gong:ident [ref_models.Parameter.RotatedRhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000009_RotatedRhombusGrid.Identifier = `ref_models.Parameter.RotatedRhombusGrid`

	//gong:ident [ref_models.RhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000009_RotatedRhombusGrid.Fieldtypename = `ref_models.RhombusGrid`
	__Link__000009_RotatedRhombusGrid.FieldOffsetX = -50.000000
	__Link__000009_RotatedRhombusGrid.FieldOffsetY = -16.000000
	__Link__000009_RotatedRhombusGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000009_RotatedRhombusGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000009_RotatedRhombusGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000009_RotatedRhombusGrid.SourceMultiplicity = models.MANY
	__Link__000009_RotatedRhombusGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000009_RotatedRhombusGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000009_RotatedRhombusGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_RotatedRhombusGrid.StartRatio = 0.500000
	__Link__000009_RotatedRhombusGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_RotatedRhombusGrid.EndRatio = 0.740316
	__Link__000009_RotatedRhombusGrid.CornerOffsetRatio = 1.380000

	__Link__000010_RotatedCircleGrid.Name = `RotatedCircleGrid`

	//gong:ident [ref_models.Parameter.RotatedCircleGrid] comment added to overcome the problem with the comment map association
	__Link__000010_RotatedCircleGrid.Identifier = `ref_models.Parameter.RotatedCircleGrid`

	//gong:ident [ref_models.CircleGrid] comment added to overcome the problem with the comment map association
	__Link__000010_RotatedCircleGrid.Fieldtypename = `ref_models.CircleGrid`
	__Link__000010_RotatedCircleGrid.FieldOffsetX = -50.000000
	__Link__000010_RotatedCircleGrid.FieldOffsetY = -16.000000
	__Link__000010_RotatedCircleGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000010_RotatedCircleGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000010_RotatedCircleGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000010_RotatedCircleGrid.SourceMultiplicity = models.MANY
	__Link__000010_RotatedCircleGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000010_RotatedCircleGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000010_RotatedCircleGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_RotatedCircleGrid.StartRatio = 0.500000
	__Link__000010_RotatedCircleGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_RotatedCircleGrid.EndRatio = 0.822394
	__Link__000010_RotatedCircleGrid.CornerOffsetRatio = 1.380000

	__Link__000011_Reference.Name = `Reference`

	//gong:ident [ref_models.RhombusGrid.Reference] comment added to overcome the problem with the comment map association
	__Link__000011_Reference.Identifier = `ref_models.RhombusGrid.Reference`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000011_Reference.Fieldtypename = `ref_models.Rhombus`
	__Link__000011_Reference.FieldOffsetX = -50.000000
	__Link__000011_Reference.FieldOffsetY = -16.000000
	__Link__000011_Reference.TargetMultiplicity = models.ZERO_ONE
	__Link__000011_Reference.TargetMultiplicityOffsetX = -50.000000
	__Link__000011_Reference.TargetMultiplicityOffsetY = 16.000000
	__Link__000011_Reference.SourceMultiplicity = models.MANY
	__Link__000011_Reference.SourceMultiplicityOffsetX = 10.000000
	__Link__000011_Reference.SourceMultiplicityOffsetY = -50.000000
	__Link__000011_Reference.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_Reference.StartRatio = 0.500000
	__Link__000011_Reference.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_Reference.EndRatio = 0.500000
	__Link__000011_Reference.CornerOffsetRatio = 1.380000

	__Link__000012_Rhombuses.Name = `Rhombuses`

	//gong:ident [ref_models.RhombusGrid.Rhombuses] comment added to overcome the problem with the comment map association
	__Link__000012_Rhombuses.Identifier = `ref_models.RhombusGrid.Rhombuses`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000012_Rhombuses.Fieldtypename = `ref_models.Rhombus`
	__Link__000012_Rhombuses.FieldOffsetX = -50.000000
	__Link__000012_Rhombuses.FieldOffsetY = -16.000000
	__Link__000012_Rhombuses.TargetMultiplicity = models.MANY
	__Link__000012_Rhombuses.TargetMultiplicityOffsetX = -50.000000
	__Link__000012_Rhombuses.TargetMultiplicityOffsetY = 16.000000
	__Link__000012_Rhombuses.SourceMultiplicity = models.MANY
	__Link__000012_Rhombuses.SourceMultiplicityOffsetX = 10.000000
	__Link__000012_Rhombuses.SourceMultiplicityOffsetY = -50.000000
	__Link__000012_Rhombuses.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_Rhombuses.StartRatio = 0.500000
	__Link__000012_Rhombuses.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_Rhombuses.EndRatio = 0.194434
	__Link__000012_Rhombuses.CornerOffsetRatio = 1.966609

	__Link__000013_Reference.Name = `Reference`

	//gong:ident [ref_models.CircleGrid.Reference] comment added to overcome the problem with the comment map association
	__Link__000013_Reference.Identifier = `ref_models.CircleGrid.Reference`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__Link__000013_Reference.Fieldtypename = `ref_models.Circle`
	__Link__000013_Reference.FieldOffsetX = -50.000000
	__Link__000013_Reference.FieldOffsetY = -16.000000
	__Link__000013_Reference.TargetMultiplicity = models.ZERO_ONE
	__Link__000013_Reference.TargetMultiplicityOffsetX = -50.000000
	__Link__000013_Reference.TargetMultiplicityOffsetY = 16.000000
	__Link__000013_Reference.SourceMultiplicity = models.MANY
	__Link__000013_Reference.SourceMultiplicityOffsetX = 10.000000
	__Link__000013_Reference.SourceMultiplicityOffsetY = -50.000000
	__Link__000013_Reference.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000013_Reference.StartRatio = 0.500000
	__Link__000013_Reference.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000013_Reference.EndRatio = 0.824199
	__Link__000013_Reference.CornerOffsetRatio = 1.380000

	__Link__000014_Circles.Name = `Circles`

	//gong:ident [ref_models.CircleGrid.Circles] comment added to overcome the problem with the comment map association
	__Link__000014_Circles.Identifier = `ref_models.CircleGrid.Circles`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__Link__000014_Circles.Fieldtypename = `ref_models.Circle`
	__Link__000014_Circles.FieldOffsetX = -50.000000
	__Link__000014_Circles.FieldOffsetY = -16.000000
	__Link__000014_Circles.TargetMultiplicity = models.MANY
	__Link__000014_Circles.TargetMultiplicityOffsetX = -50.000000
	__Link__000014_Circles.TargetMultiplicityOffsetY = 16.000000
	__Link__000014_Circles.SourceMultiplicity = models.MANY
	__Link__000014_Circles.SourceMultiplicityOffsetX = 10.000000
	__Link__000014_Circles.SourceMultiplicityOffsetY = -50.000000
	__Link__000014_Circles.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000014_Circles.StartRatio = 0.500000
	__Link__000014_Circles.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000014_Circles.EndRatio = 0.114521
	__Link__000014_Circles.CornerOffsetRatio = 1.956424

	__Position__000000_Pos_Default_Parameter.X = 18.000000
	__Position__000000_Pos_Default_Parameter.Y = 35.000000
	__Position__000000_Pos_Default_Parameter.Name = `Pos-Default-Parameter`

	__Position__000001_Pos_Default_Rhombus.X = 516.999969
	__Position__000001_Pos_Default_Rhombus.Y = 426.000000
	__Position__000001_Pos_Default_Rhombus.Name = `Pos-Default-Rhombus`

	__Position__000002_Pos_Default_HorizontalAxis.X = 508.999969
	__Position__000002_Pos_Default_HorizontalAxis.Y = 27.000000
	__Position__000002_Pos_Default_HorizontalAxis.Name = `Pos-Default-HorizontalAxis`

	__Position__000003_Pos_Default_VerticalAxis.X = 504.999969
	__Position__000003_Pos_Default_VerticalAxis.Y = 197.000000
	__Position__000003_Pos_Default_VerticalAxis.Name = `Pos-Default-VerticalAxis`

	__Position__000004_Pos_Default_RhombusGrid.X = 526.000000
	__Position__000004_Pos_Default_RhombusGrid.Y = 613.999893
	__Position__000004_Pos_Default_RhombusGrid.Name = `Pos-Default-RhombusGrid`

	__Position__000005_Pos_Default_InitialAxis.X = 530.000000
	__Position__000005_Pos_Default_InitialAxis.Y = 1072.666641
	__Position__000005_Pos_Default_InitialAxis.Name = `Pos-Default-InitialAxis`

	__Position__000006_Pos_Default_CircleGrid.X = 527.000031
	__Position__000006_Pos_Default_CircleGrid.Y = 884.000046
	__Position__000006_Pos_Default_CircleGrid.Name = `Pos-Default-CircleGrid`

	__Position__000007_Pos_Default_Circle.X = 524.000031
	__Position__000007_Pos_Default_Circle.Y = 765.000015
	__Position__000007_Pos_Default_Circle.Name = `Pos-Default-Circle`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.X = 607.999985
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Y = 121.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.X = 623.499985
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Y = 115.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-HorizontalAxis`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.X = 621.499985
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.Y = 200.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-VerticalAxis`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.X = 628.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Y = 419.499977
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.X = 625.500000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Y = 491.333321
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle.X = 631.000015
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle.Y = 484.000008
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Circle`

	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.X = 632.000015
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Y = 543.500023
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-CircleGrid`

	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.X = 635.500000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Y = 598.833321
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`

	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.X = 627.499984
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Y = 314.500000
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`

	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.X = 631.500000
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Y = 408.499947
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`

	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.X = 632.000015
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Y = 543.500023
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-CircleGrid`

	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.X = 879.499984
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Y = 570.999977
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`

	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.X = 879.499984
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Y = 585.999977
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`

	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.X = 885.000031
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Y = 871.000031
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Name = `Verticle in class diagram Default in middle between Default-CircleGrid and Default-Circle`

	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.X = 885.000031
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Y = 871.000031
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Name = `Verticle in class diagram Default in middle between Default-CircleGrid and Default-Circle`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Parameter)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Rhombus)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_HorizontalAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_VerticalAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_RhombusGrid)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_InitialAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_CircleGrid)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_Circle)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Parameter.Position = __Position__000000_Pos_Default_Parameter
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000000_Name)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000001_N)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000002_M)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000003_DiamondAngle)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000004_OriginX)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000005_OriginY)
	__GongStructShape__000000_Default_Parameter.Fields = append(__GongStructShape__000000_Default_Parameter.Fields, __Field__000006_DiamondSideLenght)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000000_InitialRhombus)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000001_HorizontalAxis)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000002_VerticalAxis)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000003_InitialRhombusGrid)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000004_InitialAxis)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000005_InitialCircle)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000006_InitialCircleGrid)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000007_RotatedAxis)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000008_RotatedRhombus)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000009_RotatedRhombusGrid)
	__GongStructShape__000000_Default_Parameter.Links = append(__GongStructShape__000000_Default_Parameter.Links, __Link__000010_RotatedCircleGrid)
	__GongStructShape__000001_Default_Rhombus.Position = __Position__000001_Pos_Default_Rhombus
	__GongStructShape__000001_Default_Rhombus.Fields = append(__GongStructShape__000001_Default_Rhombus.Fields, __Field__000007_Name)
	__GongStructShape__000001_Default_Rhombus.Fields = append(__GongStructShape__000001_Default_Rhombus.Fields, __Field__000008_IsDisplayed)
	__GongStructShape__000001_Default_Rhombus.Fields = append(__GongStructShape__000001_Default_Rhombus.Fields, __Field__000009_CenterX)
	__GongStructShape__000001_Default_Rhombus.Fields = append(__GongStructShape__000001_Default_Rhombus.Fields, __Field__000010_CenterY)
	__GongStructShape__000001_Default_Rhombus.Fields = append(__GongStructShape__000001_Default_Rhombus.Fields, __Field__000011_SideLength)
	__GongStructShape__000001_Default_Rhombus.Fields = append(__GongStructShape__000001_Default_Rhombus.Fields, __Field__000012_Angle)
	__GongStructShape__000001_Default_Rhombus.Fields = append(__GongStructShape__000001_Default_Rhombus.Fields, __Field__000013_InsideAngle)
	__GongStructShape__000002_Default_HorizontalAxis.Position = __Position__000002_Pos_Default_HorizontalAxis
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000014_Name)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000015_IsAxisDisplayed)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000016_AxisHandleBorderLength)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000017_Axis_Length)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000018_Axis_StrokeWidth)
	__GongStructShape__000003_Default_VerticalAxis.Position = __Position__000003_Pos_Default_VerticalAxis
	__GongStructShape__000003_Default_VerticalAxis.Fields = append(__GongStructShape__000003_Default_VerticalAxis.Fields, __Field__000019_Name)
	__GongStructShape__000003_Default_VerticalAxis.Fields = append(__GongStructShape__000003_Default_VerticalAxis.Fields, __Field__000020_IsAxisDisplayed)
	__GongStructShape__000003_Default_VerticalAxis.Fields = append(__GongStructShape__000003_Default_VerticalAxis.Fields, __Field__000021_AxisHandleBorderLength)
	__GongStructShape__000003_Default_VerticalAxis.Fields = append(__GongStructShape__000003_Default_VerticalAxis.Fields, __Field__000022_Axis_Length)
	__GongStructShape__000003_Default_VerticalAxis.Fields = append(__GongStructShape__000003_Default_VerticalAxis.Fields, __Field__000023_Axis_StrokeWidth)
	__GongStructShape__000004_Default_RhombusGrid.Position = __Position__000004_Pos_Default_RhombusGrid
	__GongStructShape__000004_Default_RhombusGrid.Fields = append(__GongStructShape__000004_Default_RhombusGrid.Fields, __Field__000024_Name)
	__GongStructShape__000004_Default_RhombusGrid.Fields = append(__GongStructShape__000004_Default_RhombusGrid.Fields, __Field__000025_N)
	__GongStructShape__000004_Default_RhombusGrid.Fields = append(__GongStructShape__000004_Default_RhombusGrid.Fields, __Field__000026_M)
	__GongStructShape__000004_Default_RhombusGrid.Links = append(__GongStructShape__000004_Default_RhombusGrid.Links, __Link__000011_Reference)
	__GongStructShape__000004_Default_RhombusGrid.Links = append(__GongStructShape__000004_Default_RhombusGrid.Links, __Link__000012_Rhombuses)
	__GongStructShape__000005_Default_InitialAxis.Position = __Position__000005_Pos_Default_InitialAxis
	__GongStructShape__000005_Default_InitialAxis.Fields = append(__GongStructShape__000005_Default_InitialAxis.Fields, __Field__000027_Name)
	__GongStructShape__000005_Default_InitialAxis.Fields = append(__GongStructShape__000005_Default_InitialAxis.Fields, __Field__000028_Angle)
	__GongStructShape__000005_Default_InitialAxis.Fields = append(__GongStructShape__000005_Default_InitialAxis.Fields, __Field__000029_Length)
	__GongStructShape__000006_Default_CircleGrid.Position = __Position__000006_Pos_Default_CircleGrid
	__GongStructShape__000006_Default_CircleGrid.Fields = append(__GongStructShape__000006_Default_CircleGrid.Fields, __Field__000030_N)
	__GongStructShape__000006_Default_CircleGrid.Fields = append(__GongStructShape__000006_Default_CircleGrid.Fields, __Field__000031_M)
	__GongStructShape__000006_Default_CircleGrid.Links = append(__GongStructShape__000006_Default_CircleGrid.Links, __Link__000013_Reference)
	__GongStructShape__000006_Default_CircleGrid.Links = append(__GongStructShape__000006_Default_CircleGrid.Links, __Link__000014_Circles)
	__GongStructShape__000007_Default_Circle.Position = __Position__000007_Pos_Default_Circle
	__GongStructShape__000007_Default_Circle.Fields = append(__GongStructShape__000007_Default_Circle.Fields, __Field__000032_CenterX)
	__GongStructShape__000007_Default_Circle.Fields = append(__GongStructShape__000007_Default_Circle.Fields, __Field__000033_CenterY)
	// setup of Link instances pointers
	__Link__000000_InitialRhombus.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus
	__Link__000001_HorizontalAxis.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis
	__Link__000002_VerticalAxis.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis
	__Link__000003_InitialRhombusGrid.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid
	__Link__000004_InitialAxis.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis
	__Link__000005_InitialCircle.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle
	__Link__000006_InitialCircleGrid.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid
	__Link__000007_RotatedAxis.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis
	__Link__000008_RotatedRhombus.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus
	__Link__000009_RotatedRhombusGrid.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid
	__Link__000010_RotatedCircleGrid.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid
	__Link__000011_Reference.Middlevertice = __Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus
	__Link__000012_Rhombuses.Middlevertice = __Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus
	__Link__000013_Reference.Middlevertice = __Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle
	__Link__000014_Circles.Middlevertice = __Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
