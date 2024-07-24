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
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_Angle := (&models.Field{Name: `Angle`}).Stage(stage)
	__Field__000001_Angle := (&models.Field{Name: `Angle`}).Stage(stage)
	__Field__000002_AxisHandleBorderLength := (&models.Field{Name: `AxisHandleBorderLength`}).Stage(stage)
	__Field__000003_AxisHandleBorderLength := (&models.Field{Name: `AxisHandleBorderLength`}).Stage(stage)
	__Field__000004_Axis_Length := (&models.Field{Name: `Axis_Length`}).Stage(stage)
	__Field__000005_Axis_Length := (&models.Field{Name: `Axis_Length`}).Stage(stage)
	__Field__000006_Axis_StrokeWidth := (&models.Field{Name: `Axis_StrokeWidth`}).Stage(stage)
	__Field__000007_Axis_StrokeWidth := (&models.Field{Name: `Axis_StrokeWidth`}).Stage(stage)
	__Field__000008_CenterX := (&models.Field{Name: `CenterX`}).Stage(stage)
	__Field__000009_CenterX := (&models.Field{Name: `CenterX`}).Stage(stage)
	__Field__000010_CenterY := (&models.Field{Name: `CenterY`}).Stage(stage)
	__Field__000011_CenterY := (&models.Field{Name: `CenterY`}).Stage(stage)
	__Field__000012_DiamondAngle := (&models.Field{Name: `DiamondAngle`}).Stage(stage)
	__Field__000013_DiamondSideLenght := (&models.Field{Name: `DiamondSideLenght`}).Stage(stage)
	__Field__000014_InsideAngle := (&models.Field{Name: `InsideAngle`}).Stage(stage)
	__Field__000015_IsAxisDisplayed := (&models.Field{Name: `IsAxisDisplayed`}).Stage(stage)
	__Field__000016_IsAxisDisplayed := (&models.Field{Name: `IsAxisDisplayed`}).Stage(stage)
	__Field__000017_IsDisplayed := (&models.Field{Name: `IsDisplayed`}).Stage(stage)
	__Field__000018_Length := (&models.Field{Name: `Length`}).Stage(stage)
	__Field__000019_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000020_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000021_M := (&models.Field{Name: `M`}).Stage(stage)
	__Field__000022_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000023_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000024_N := (&models.Field{Name: `N`}).Stage(stage)
	__Field__000025_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000026_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000027_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000028_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000029_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000030_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000031_OriginX := (&models.Field{Name: `OriginX`}).Stage(stage)
	__Field__000032_OriginY := (&models.Field{Name: `OriginY`}).Stage(stage)
	__Field__000033_SideLength := (&models.Field{Name: `SideLength`}).Stage(stage)

	__GongStructShape__000000_Default_Circle := (&models.GongStructShape{Name: `Default-Circle`}).Stage(stage)
	__GongStructShape__000001_Default_CircleGrid := (&models.GongStructShape{Name: `Default-CircleGrid`}).Stage(stage)
	__GongStructShape__000002_Default_HorizontalAxis := (&models.GongStructShape{Name: `Default-HorizontalAxis`}).Stage(stage)
	__GongStructShape__000003_Default_InitialAxis := (&models.GongStructShape{Name: `Default-InitialAxis`}).Stage(stage)
	__GongStructShape__000004_Default_Parameter := (&models.GongStructShape{Name: `Default-Parameter`}).Stage(stage)
	__GongStructShape__000005_Default_Rhombus := (&models.GongStructShape{Name: `Default-Rhombus`}).Stage(stage)
	__GongStructShape__000006_Default_RhombusGrid := (&models.GongStructShape{Name: `Default-RhombusGrid`}).Stage(stage)
	__GongStructShape__000007_Default_VerticalAxis := (&models.GongStructShape{Name: `Default-VerticalAxis`}).Stage(stage)

	__Link__000000_Circles := (&models.Link{Name: `Circles`}).Stage(stage)
	__Link__000001_HorizontalAxis := (&models.Link{Name: `HorizontalAxis`}).Stage(stage)
	__Link__000002_InitialAxis := (&models.Link{Name: `InitialAxis`}).Stage(stage)
	__Link__000003_InitialCircle := (&models.Link{Name: `InitialCircle`}).Stage(stage)
	__Link__000004_InitialCircleGrid := (&models.Link{Name: `InitialCircleGrid`}).Stage(stage)
	__Link__000005_InitialRhombus := (&models.Link{Name: `InitialRhombus`}).Stage(stage)
	__Link__000006_InitialRhombusGrid := (&models.Link{Name: `InitialRhombusGrid`}).Stage(stage)
	__Link__000007_Reference := (&models.Link{Name: `Reference`}).Stage(stage)
	__Link__000008_Reference := (&models.Link{Name: `Reference`}).Stage(stage)
	__Link__000009_Rhombuses := (&models.Link{Name: `Rhombuses`}).Stage(stage)
	__Link__000010_RotatedAxis := (&models.Link{Name: `RotatedAxis`}).Stage(stage)
	__Link__000011_RotatedCircleGrid := (&models.Link{Name: `RotatedCircleGrid`}).Stage(stage)
	__Link__000012_RotatedRhombus := (&models.Link{Name: `RotatedRhombus`}).Stage(stage)
	__Link__000013_RotatedRhombusGrid := (&models.Link{Name: `RotatedRhombusGrid`}).Stage(stage)
	__Link__000014_VerticalAxis := (&models.Link{Name: `VerticalAxis`}).Stage(stage)

	__Position__000000_Pos_Default_Circle := (&models.Position{Name: `Pos-Default-Circle`}).Stage(stage)
	__Position__000001_Pos_Default_CircleGrid := (&models.Position{Name: `Pos-Default-CircleGrid`}).Stage(stage)
	__Position__000002_Pos_Default_HorizontalAxis := (&models.Position{Name: `Pos-Default-HorizontalAxis`}).Stage(stage)
	__Position__000003_Pos_Default_InitialAxis := (&models.Position{Name: `Pos-Default-InitialAxis`}).Stage(stage)
	__Position__000004_Pos_Default_Parameter := (&models.Position{Name: `Pos-Default-Parameter`}).Stage(stage)
	__Position__000005_Pos_Default_Rhombus := (&models.Position{Name: `Pos-Default-Rhombus`}).Stage(stage)
	__Position__000006_Pos_Default_RhombusGrid := (&models.Position{Name: `Pos-Default-RhombusGrid`}).Stage(stage)
	__Position__000007_Pos_Default_VerticalAxis := (&models.Position{Name: `Pos-Default-VerticalAxis`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-CircleGrid and Default-Circle`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-CircleGrid and Default-Circle`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-Circle`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-CircleGrid`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-CircleGrid`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-HorizontalAxis`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`}).Stage(stage)
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`}).Stage(stage)
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Parameter and Default-VerticalAxis`}).Stage(stage)
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`}).Stage(stage)
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Angle.Name = `Angle`

	//gong:ident [ref_models.Rhombus.Angle] comment added to overcome the problem with the comment map association
	__Field__000000_Angle.Identifier = `ref_models.Rhombus.Angle`
	__Field__000000_Angle.FieldTypeAsString = ``
	__Field__000000_Angle.Structname = `Rhombus`
	__Field__000000_Angle.Fieldtypename = `float64`

	__Field__000001_Angle.Name = `Angle`

	//gong:ident [ref_models.Axis.Angle] comment added to overcome the problem with the comment map association
	__Field__000001_Angle.Identifier = `ref_models.Axis.Angle`
	__Field__000001_Angle.FieldTypeAsString = ``
	__Field__000001_Angle.Structname = `Axis`
	__Field__000001_Angle.Fieldtypename = `float64`

	__Field__000002_AxisHandleBorderLength.Name = `AxisHandleBorderLength`

	//gong:ident [ref_models.HorizontalAxis.AxisHandleBorderLength] comment added to overcome the problem with the comment map association
	__Field__000002_AxisHandleBorderLength.Identifier = `ref_models.HorizontalAxis.AxisHandleBorderLength`
	__Field__000002_AxisHandleBorderLength.FieldTypeAsString = ``
	__Field__000002_AxisHandleBorderLength.Structname = `HorizontalAxis`
	__Field__000002_AxisHandleBorderLength.Fieldtypename = `float64`

	__Field__000003_AxisHandleBorderLength.Name = `AxisHandleBorderLength`

	//gong:ident [ref_models.VerticalAxis.AxisHandleBorderLength] comment added to overcome the problem with the comment map association
	__Field__000003_AxisHandleBorderLength.Identifier = `ref_models.VerticalAxis.AxisHandleBorderLength`
	__Field__000003_AxisHandleBorderLength.FieldTypeAsString = ``
	__Field__000003_AxisHandleBorderLength.Structname = `VerticalAxis`
	__Field__000003_AxisHandleBorderLength.Fieldtypename = `float64`

	__Field__000004_Axis_Length.Name = `Axis_Length`

	//gong:ident [ref_models.VerticalAxis.Axis_Length] comment added to overcome the problem with the comment map association
	__Field__000004_Axis_Length.Identifier = `ref_models.VerticalAxis.Axis_Length`
	__Field__000004_Axis_Length.FieldTypeAsString = ``
	__Field__000004_Axis_Length.Structname = `VerticalAxis`
	__Field__000004_Axis_Length.Fieldtypename = `float64`

	__Field__000005_Axis_Length.Name = `Axis_Length`

	//gong:ident [ref_models.HorizontalAxis.Axis_Length] comment added to overcome the problem with the comment map association
	__Field__000005_Axis_Length.Identifier = `ref_models.HorizontalAxis.Axis_Length`
	__Field__000005_Axis_Length.FieldTypeAsString = ``
	__Field__000005_Axis_Length.Structname = `HorizontalAxis`
	__Field__000005_Axis_Length.Fieldtypename = `float64`

	__Field__000006_Axis_StrokeWidth.Name = `Axis_StrokeWidth`

	//gong:ident [ref_models.VerticalAxis.StrokeWidth] comment added to overcome the problem with the comment map association
	__Field__000006_Axis_StrokeWidth.Identifier = `ref_models.VerticalAxis.StrokeWidth`
	__Field__000006_Axis_StrokeWidth.FieldTypeAsString = ``
	__Field__000006_Axis_StrokeWidth.Structname = `VerticalAxis`
	__Field__000006_Axis_StrokeWidth.Fieldtypename = `float64`

	__Field__000007_Axis_StrokeWidth.Name = `Axis_StrokeWidth`

	//gong:ident [ref_models.HorizontalAxis.StrokeWidth] comment added to overcome the problem with the comment map association
	__Field__000007_Axis_StrokeWidth.Identifier = `ref_models.HorizontalAxis.StrokeWidth`
	__Field__000007_Axis_StrokeWidth.FieldTypeAsString = ``
	__Field__000007_Axis_StrokeWidth.Structname = `HorizontalAxis`
	__Field__000007_Axis_StrokeWidth.Fieldtypename = `float64`

	__Field__000008_CenterX.Name = `CenterX`

	//gong:ident [ref_models.Rhombus.CenterX] comment added to overcome the problem with the comment map association
	__Field__000008_CenterX.Identifier = `ref_models.Rhombus.CenterX`
	__Field__000008_CenterX.FieldTypeAsString = ``
	__Field__000008_CenterX.Structname = `Rhombus`
	__Field__000008_CenterX.Fieldtypename = `float64`

	__Field__000009_CenterX.Name = `CenterX`

	//gong:ident [ref_models.Circle.CenterX] comment added to overcome the problem with the comment map association
	__Field__000009_CenterX.Identifier = `ref_models.Circle.CenterX`
	__Field__000009_CenterX.FieldTypeAsString = ``
	__Field__000009_CenterX.Structname = `Circle`
	__Field__000009_CenterX.Fieldtypename = `float64`

	__Field__000010_CenterY.Name = `CenterY`

	//gong:ident [ref_models.Circle.CenterY] comment added to overcome the problem with the comment map association
	__Field__000010_CenterY.Identifier = `ref_models.Circle.CenterY`
	__Field__000010_CenterY.FieldTypeAsString = ``
	__Field__000010_CenterY.Structname = `Circle`
	__Field__000010_CenterY.Fieldtypename = `float64`

	__Field__000011_CenterY.Name = `CenterY`

	//gong:ident [ref_models.Rhombus.CenterY] comment added to overcome the problem with the comment map association
	__Field__000011_CenterY.Identifier = `ref_models.Rhombus.CenterY`
	__Field__000011_CenterY.FieldTypeAsString = ``
	__Field__000011_CenterY.Structname = `Rhombus`
	__Field__000011_CenterY.Fieldtypename = `float64`

	__Field__000012_DiamondAngle.Name = `DiamondAngle`

	//gong:ident [ref_models.Parameter.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000012_DiamondAngle.Identifier = `ref_models.Parameter.InsideAngle`
	__Field__000012_DiamondAngle.FieldTypeAsString = ``
	__Field__000012_DiamondAngle.Structname = `Parameter`
	__Field__000012_DiamondAngle.Fieldtypename = `float64`

	__Field__000013_DiamondSideLenght.Name = `DiamondSideLenght`

	//gong:ident [ref_models.Parameter.SideLength] comment added to overcome the problem with the comment map association
	__Field__000013_DiamondSideLenght.Identifier = `ref_models.Parameter.SideLength`
	__Field__000013_DiamondSideLenght.FieldTypeAsString = ``
	__Field__000013_DiamondSideLenght.Structname = `Parameter`
	__Field__000013_DiamondSideLenght.Fieldtypename = `float64`

	__Field__000014_InsideAngle.Name = `InsideAngle`

	//gong:ident [ref_models.Rhombus.InsideAngle] comment added to overcome the problem with the comment map association
	__Field__000014_InsideAngle.Identifier = `ref_models.Rhombus.InsideAngle`
	__Field__000014_InsideAngle.FieldTypeAsString = ``
	__Field__000014_InsideAngle.Structname = `Rhombus`
	__Field__000014_InsideAngle.Fieldtypename = `float64`

	__Field__000015_IsAxisDisplayed.Name = `IsAxisDisplayed`

	//gong:ident [ref_models.VerticalAxis.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000015_IsAxisDisplayed.Identifier = `ref_models.VerticalAxis.IsDisplayed`
	__Field__000015_IsAxisDisplayed.FieldTypeAsString = ``
	__Field__000015_IsAxisDisplayed.Structname = `VerticalAxis`
	__Field__000015_IsAxisDisplayed.Fieldtypename = `bool`

	__Field__000016_IsAxisDisplayed.Name = `IsAxisDisplayed`

	//gong:ident [ref_models.HorizontalAxis.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000016_IsAxisDisplayed.Identifier = `ref_models.HorizontalAxis.IsDisplayed`
	__Field__000016_IsAxisDisplayed.FieldTypeAsString = ``
	__Field__000016_IsAxisDisplayed.Structname = `HorizontalAxis`
	__Field__000016_IsAxisDisplayed.Fieldtypename = `bool`

	__Field__000017_IsDisplayed.Name = `IsDisplayed`

	//gong:ident [ref_models.Rhombus.IsDisplayed] comment added to overcome the problem with the comment map association
	__Field__000017_IsDisplayed.Identifier = `ref_models.Rhombus.IsDisplayed`
	__Field__000017_IsDisplayed.FieldTypeAsString = ``
	__Field__000017_IsDisplayed.Structname = `Rhombus`
	__Field__000017_IsDisplayed.Fieldtypename = `bool`

	__Field__000018_Length.Name = `Length`

	//gong:ident [ref_models.Axis.Length] comment added to overcome the problem with the comment map association
	__Field__000018_Length.Identifier = `ref_models.Axis.Length`
	__Field__000018_Length.FieldTypeAsString = ``
	__Field__000018_Length.Structname = `Axis`
	__Field__000018_Length.Fieldtypename = `float64`

	__Field__000019_M.Name = `M`

	//gong:ident [ref_models.RhombusGrid.M] comment added to overcome the problem with the comment map association
	__Field__000019_M.Identifier = `ref_models.RhombusGrid.M`
	__Field__000019_M.FieldTypeAsString = ``
	__Field__000019_M.Structname = `RhombusGrid`
	__Field__000019_M.Fieldtypename = `int`

	__Field__000020_M.Name = `M`

	//gong:ident [ref_models.Parameter.M] comment added to overcome the problem with the comment map association
	__Field__000020_M.Identifier = `ref_models.Parameter.M`
	__Field__000020_M.FieldTypeAsString = ``
	__Field__000020_M.Structname = `Parameter`
	__Field__000020_M.Fieldtypename = `int`

	__Field__000021_M.Name = `M`

	//gong:ident [ref_models.CircleGrid.M] comment added to overcome the problem with the comment map association
	__Field__000021_M.Identifier = `ref_models.CircleGrid.M`
	__Field__000021_M.FieldTypeAsString = ``
	__Field__000021_M.Structname = `CircleGrid`
	__Field__000021_M.Fieldtypename = `int`

	__Field__000022_N.Name = `N`

	//gong:ident [ref_models.CircleGrid.N] comment added to overcome the problem with the comment map association
	__Field__000022_N.Identifier = `ref_models.CircleGrid.N`
	__Field__000022_N.FieldTypeAsString = ``
	__Field__000022_N.Structname = `CircleGrid`
	__Field__000022_N.Fieldtypename = `int`

	__Field__000023_N.Name = `N`

	//gong:ident [ref_models.Parameter.N] comment added to overcome the problem with the comment map association
	__Field__000023_N.Identifier = `ref_models.Parameter.N`
	__Field__000023_N.FieldTypeAsString = ``
	__Field__000023_N.Structname = `Parameter`
	__Field__000023_N.Fieldtypename = `int`

	__Field__000024_N.Name = `N`

	//gong:ident [ref_models.RhombusGrid.N] comment added to overcome the problem with the comment map association
	__Field__000024_N.Identifier = `ref_models.RhombusGrid.N`
	__Field__000024_N.FieldTypeAsString = ``
	__Field__000024_N.Structname = `RhombusGrid`
	__Field__000024_N.Fieldtypename = `int`

	__Field__000025_Name.Name = `Name`

	//gong:ident [ref_models.Parameter.Name] comment added to overcome the problem with the comment map association
	__Field__000025_Name.Identifier = `ref_models.Parameter.Name`
	__Field__000025_Name.FieldTypeAsString = ``
	__Field__000025_Name.Structname = `Parameter`
	__Field__000025_Name.Fieldtypename = `string`

	__Field__000026_Name.Name = `Name`

	//gong:ident [ref_models.Axis.Name] comment added to overcome the problem with the comment map association
	__Field__000026_Name.Identifier = `ref_models.Axis.Name`
	__Field__000026_Name.FieldTypeAsString = ``
	__Field__000026_Name.Structname = `Axis`
	__Field__000026_Name.Fieldtypename = `string`

	__Field__000027_Name.Name = `Name`

	//gong:ident [ref_models.HorizontalAxis.Name] comment added to overcome the problem with the comment map association
	__Field__000027_Name.Identifier = `ref_models.HorizontalAxis.Name`
	__Field__000027_Name.FieldTypeAsString = ``
	__Field__000027_Name.Structname = `HorizontalAxis`
	__Field__000027_Name.Fieldtypename = `string`

	__Field__000028_Name.Name = `Name`

	//gong:ident [ref_models.RhombusGrid.Name] comment added to overcome the problem with the comment map association
	__Field__000028_Name.Identifier = `ref_models.RhombusGrid.Name`
	__Field__000028_Name.FieldTypeAsString = ``
	__Field__000028_Name.Structname = `RhombusGrid`
	__Field__000028_Name.Fieldtypename = `string`

	__Field__000029_Name.Name = `Name`

	//gong:ident [ref_models.VerticalAxis.Name] comment added to overcome the problem with the comment map association
	__Field__000029_Name.Identifier = `ref_models.VerticalAxis.Name`
	__Field__000029_Name.FieldTypeAsString = ``
	__Field__000029_Name.Structname = `VerticalAxis`
	__Field__000029_Name.Fieldtypename = `string`

	__Field__000030_Name.Name = `Name`

	//gong:ident [ref_models.Rhombus.Name] comment added to overcome the problem with the comment map association
	__Field__000030_Name.Identifier = `ref_models.Rhombus.Name`
	__Field__000030_Name.FieldTypeAsString = ``
	__Field__000030_Name.Structname = `Rhombus`
	__Field__000030_Name.Fieldtypename = `string`

	__Field__000031_OriginX.Name = `OriginX`

	//gong:ident [ref_models.Parameter.OriginX] comment added to overcome the problem with the comment map association
	__Field__000031_OriginX.Identifier = `ref_models.Parameter.OriginX`
	__Field__000031_OriginX.FieldTypeAsString = ``
	__Field__000031_OriginX.Structname = `Parameter`
	__Field__000031_OriginX.Fieldtypename = `float64`

	__Field__000032_OriginY.Name = `OriginY`

	//gong:ident [ref_models.Parameter.OriginY] comment added to overcome the problem with the comment map association
	__Field__000032_OriginY.Identifier = `ref_models.Parameter.OriginY`
	__Field__000032_OriginY.FieldTypeAsString = ``
	__Field__000032_OriginY.Structname = `Parameter`
	__Field__000032_OriginY.Fieldtypename = `float64`

	__Field__000033_SideLength.Name = `SideLength`

	//gong:ident [ref_models.Rhombus.SideLength] comment added to overcome the problem with the comment map association
	__Field__000033_SideLength.Identifier = `ref_models.Rhombus.SideLength`
	__Field__000033_SideLength.FieldTypeAsString = ``
	__Field__000033_SideLength.Structname = `Rhombus`
	__Field__000033_SideLength.Fieldtypename = `float64`

	__GongStructShape__000000_Default_Circle.Name = `Default-Circle`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000000_Default_Circle.ShowNbInstances = false
	__GongStructShape__000000_Default_Circle.NbInstances = 4
	__GongStructShape__000000_Default_Circle.Width = 240.000000
	__GongStructShape__000000_Default_Circle.Height = 93.000000
	__GongStructShape__000000_Default_Circle.IsSelected = false

	__GongStructShape__000001_Default_CircleGrid.Name = `Default-CircleGrid`

	//gong:ident [ref_models.CircleGrid] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_CircleGrid.Identifier = `ref_models.CircleGrid`
	__GongStructShape__000001_Default_CircleGrid.ShowNbInstances = false
	__GongStructShape__000001_Default_CircleGrid.NbInstances = 4
	__GongStructShape__000001_Default_CircleGrid.Width = 240.000000
	__GongStructShape__000001_Default_CircleGrid.Height = 164.000061
	__GongStructShape__000001_Default_CircleGrid.IsSelected = false

	__GongStructShape__000002_Default_HorizontalAxis.Name = `Default-HorizontalAxis`

	//gong:ident [ref_models.HorizontalAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_HorizontalAxis.Identifier = `ref_models.HorizontalAxis`
	__GongStructShape__000002_Default_HorizontalAxis.ShowNbInstances = false
	__GongStructShape__000002_Default_HorizontalAxis.NbInstances = 1
	__GongStructShape__000002_Default_HorizontalAxis.Width = 331.000000
	__GongStructShape__000002_Default_HorizontalAxis.Height = 138.000000
	__GongStructShape__000002_Default_HorizontalAxis.IsSelected = false

	__GongStructShape__000003_Default_InitialAxis.Name = `Default-InitialAxis`

	//gong:ident [ref_models.Axis] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_InitialAxis.Identifier = `ref_models.Axis`
	__GongStructShape__000003_Default_InitialAxis.ShowNbInstances = false
	__GongStructShape__000003_Default_InitialAxis.NbInstances = 3
	__GongStructShape__000003_Default_InitialAxis.Width = 240.000000
	__GongStructShape__000003_Default_InitialAxis.Height = 214.000000
	__GongStructShape__000003_Default_InitialAxis.IsSelected = false

	__GongStructShape__000004_Default_Parameter.Name = `Default-Parameter`

	//gong:ident [ref_models.Parameter] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_Parameter.Identifier = `ref_models.Parameter`
	__GongStructShape__000004_Default_Parameter.ShowNbInstances = false
	__GongStructShape__000004_Default_Parameter.NbInstances = 1
	__GongStructShape__000004_Default_Parameter.Width = 240.000000
	__GongStructShape__000004_Default_Parameter.Height = 168.000000
	__GongStructShape__000004_Default_Parameter.IsSelected = false

	__GongStructShape__000005_Default_Rhombus.Name = `Default-Rhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_Rhombus.Identifier = `ref_models.Rhombus`
	__GongStructShape__000005_Default_Rhombus.ShowNbInstances = false
	__GongStructShape__000005_Default_Rhombus.NbInstances = 4
	__GongStructShape__000005_Default_Rhombus.Width = 314.000000
	__GongStructShape__000005_Default_Rhombus.Height = 168.000000
	__GongStructShape__000005_Default_Rhombus.IsSelected = false

	__GongStructShape__000006_Default_RhombusGrid.Name = `Default-RhombusGrid`

	//gong:ident [ref_models.RhombusGrid] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_RhombusGrid.Identifier = `ref_models.RhombusGrid`
	__GongStructShape__000006_Default_RhombusGrid.ShowNbInstances = false
	__GongStructShape__000006_Default_RhombusGrid.NbInstances = 3
	__GongStructShape__000006_Default_RhombusGrid.Width = 307.000000
	__GongStructShape__000006_Default_RhombusGrid.Height = 125.000000
	__GongStructShape__000006_Default_RhombusGrid.IsSelected = false

	__GongStructShape__000007_Default_VerticalAxis.Name = `Default-VerticalAxis`

	//gong:ident [ref_models.VerticalAxis] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_Default_VerticalAxis.Identifier = `ref_models.VerticalAxis`
	__GongStructShape__000007_Default_VerticalAxis.ShowNbInstances = false
	__GongStructShape__000007_Default_VerticalAxis.NbInstances = 1
	__GongStructShape__000007_Default_VerticalAxis.Width = 337.000000
	__GongStructShape__000007_Default_VerticalAxis.Height = 138.000000
	__GongStructShape__000007_Default_VerticalAxis.IsSelected = false

	__Link__000000_Circles.Name = `Circles`

	//gong:ident [ref_models.CircleGrid.Circles] comment added to overcome the problem with the comment map association
	__Link__000000_Circles.Identifier = `ref_models.CircleGrid.Circles`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__Link__000000_Circles.Fieldtypename = `ref_models.Circle`
	__Link__000000_Circles.FieldOffsetX = -50.000000
	__Link__000000_Circles.FieldOffsetY = -16.000000
	__Link__000000_Circles.TargetMultiplicity = models.MANY
	__Link__000000_Circles.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Circles.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Circles.SourceMultiplicity = models.MANY
	__Link__000000_Circles.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_Circles.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_Circles.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Circles.StartRatio = 0.500000
	__Link__000000_Circles.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Circles.EndRatio = 0.114521
	__Link__000000_Circles.CornerOffsetRatio = 1.956424

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

	__Link__000002_InitialAxis.Name = `InitialAxis`

	//gong:ident [ref_models.Parameter.InitialAxis] comment added to overcome the problem with the comment map association
	__Link__000002_InitialAxis.Identifier = `ref_models.Parameter.InitialAxis`

	//gong:ident [ref_models.Axis] comment added to overcome the problem with the comment map association
	__Link__000002_InitialAxis.Fieldtypename = `ref_models.Axis`
	__Link__000002_InitialAxis.FieldOffsetX = -50.000000
	__Link__000002_InitialAxis.FieldOffsetY = -16.000000
	__Link__000002_InitialAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_InitialAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_InitialAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_InitialAxis.SourceMultiplicity = models.MANY
	__Link__000002_InitialAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_InitialAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_InitialAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_InitialAxis.StartRatio = 0.500000
	__Link__000002_InitialAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_InitialAxis.EndRatio = 0.309375
	__Link__000002_InitialAxis.CornerOffsetRatio = 1.380000

	__Link__000003_InitialCircle.Name = `InitialCircle`

	//gong:ident [ref_models.Parameter.InitialCircle] comment added to overcome the problem with the comment map association
	__Link__000003_InitialCircle.Identifier = `ref_models.Parameter.InitialCircle`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__Link__000003_InitialCircle.Fieldtypename = `ref_models.Circle`
	__Link__000003_InitialCircle.FieldOffsetX = -50.000000
	__Link__000003_InitialCircle.FieldOffsetY = -16.000000
	__Link__000003_InitialCircle.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_InitialCircle.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_InitialCircle.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_InitialCircle.SourceMultiplicity = models.MANY
	__Link__000003_InitialCircle.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_InitialCircle.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_InitialCircle.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_InitialCircle.StartRatio = 0.500000
	__Link__000003_InitialCircle.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_InitialCircle.EndRatio = 0.500000
	__Link__000003_InitialCircle.CornerOffsetRatio = 1.380000

	__Link__000004_InitialCircleGrid.Name = `InitialCircleGrid`

	//gong:ident [ref_models.Parameter.InitialCircleGrid] comment added to overcome the problem with the comment map association
	__Link__000004_InitialCircleGrid.Identifier = `ref_models.Parameter.InitialCircleGrid`

	//gong:ident [ref_models.CircleGrid] comment added to overcome the problem with the comment map association
	__Link__000004_InitialCircleGrid.Fieldtypename = `ref_models.CircleGrid`
	__Link__000004_InitialCircleGrid.FieldOffsetX = -50.000000
	__Link__000004_InitialCircleGrid.FieldOffsetY = -16.000000
	__Link__000004_InitialCircleGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_InitialCircleGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_InitialCircleGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_InitialCircleGrid.SourceMultiplicity = models.MANY
	__Link__000004_InitialCircleGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_InitialCircleGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_InitialCircleGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_InitialCircleGrid.StartRatio = 0.500000
	__Link__000004_InitialCircleGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_InitialCircleGrid.EndRatio = 0.273614
	__Link__000004_InitialCircleGrid.CornerOffsetRatio = 1.380000

	__Link__000005_InitialRhombus.Name = `InitialRhombus`

	//gong:ident [ref_models.Parameter.InitialRhombus] comment added to overcome the problem with the comment map association
	__Link__000005_InitialRhombus.Identifier = `ref_models.Parameter.InitialRhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000005_InitialRhombus.Fieldtypename = `ref_models.Rhombus`
	__Link__000005_InitialRhombus.FieldOffsetX = -50.000000
	__Link__000005_InitialRhombus.FieldOffsetY = -16.000000
	__Link__000005_InitialRhombus.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_InitialRhombus.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_InitialRhombus.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_InitialRhombus.SourceMultiplicity = models.MANY
	__Link__000005_InitialRhombus.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_InitialRhombus.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_InitialRhombus.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_InitialRhombus.StartRatio = 0.494699
	__Link__000005_InitialRhombus.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_InitialRhombus.EndRatio = 0.500000
	__Link__000005_InitialRhombus.CornerOffsetRatio = 1.379109

	__Link__000006_InitialRhombusGrid.Name = `InitialRhombusGrid`

	//gong:ident [ref_models.Parameter.InitialRhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000006_InitialRhombusGrid.Identifier = `ref_models.Parameter.InitialRhombusGrid`

	//gong:ident [ref_models.RhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000006_InitialRhombusGrid.Fieldtypename = `ref_models.RhombusGrid`
	__Link__000006_InitialRhombusGrid.FieldOffsetX = -50.000000
	__Link__000006_InitialRhombusGrid.FieldOffsetY = -16.000000
	__Link__000006_InitialRhombusGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_InitialRhombusGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_InitialRhombusGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_InitialRhombusGrid.SourceMultiplicity = models.MANY
	__Link__000006_InitialRhombusGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_InitialRhombusGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_InitialRhombusGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_InitialRhombusGrid.StartRatio = 0.500000
	__Link__000006_InitialRhombusGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_InitialRhombusGrid.EndRatio = 0.252316
	__Link__000006_InitialRhombusGrid.CornerOffsetRatio = 1.380000

	__Link__000007_Reference.Name = `Reference`

	//gong:ident [ref_models.RhombusGrid.Reference] comment added to overcome the problem with the comment map association
	__Link__000007_Reference.Identifier = `ref_models.RhombusGrid.Reference`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000007_Reference.Fieldtypename = `ref_models.Rhombus`
	__Link__000007_Reference.FieldOffsetX = -50.000000
	__Link__000007_Reference.FieldOffsetY = -16.000000
	__Link__000007_Reference.TargetMultiplicity = models.ZERO_ONE
	__Link__000007_Reference.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_Reference.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_Reference.SourceMultiplicity = models.MANY
	__Link__000007_Reference.SourceMultiplicityOffsetX = 10.000000
	__Link__000007_Reference.SourceMultiplicityOffsetY = -50.000000
	__Link__000007_Reference.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_Reference.StartRatio = 0.500000
	__Link__000007_Reference.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_Reference.EndRatio = 0.500000
	__Link__000007_Reference.CornerOffsetRatio = 1.380000

	__Link__000008_Reference.Name = `Reference`

	//gong:ident [ref_models.CircleGrid.Reference] comment added to overcome the problem with the comment map association
	__Link__000008_Reference.Identifier = `ref_models.CircleGrid.Reference`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__Link__000008_Reference.Fieldtypename = `ref_models.Circle`
	__Link__000008_Reference.FieldOffsetX = -50.000000
	__Link__000008_Reference.FieldOffsetY = -16.000000
	__Link__000008_Reference.TargetMultiplicity = models.ZERO_ONE
	__Link__000008_Reference.TargetMultiplicityOffsetX = -50.000000
	__Link__000008_Reference.TargetMultiplicityOffsetY = 16.000000
	__Link__000008_Reference.SourceMultiplicity = models.MANY
	__Link__000008_Reference.SourceMultiplicityOffsetX = 10.000000
	__Link__000008_Reference.SourceMultiplicityOffsetY = -50.000000
	__Link__000008_Reference.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_Reference.StartRatio = 0.500000
	__Link__000008_Reference.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_Reference.EndRatio = 0.824199
	__Link__000008_Reference.CornerOffsetRatio = 1.380000

	__Link__000009_Rhombuses.Name = `Rhombuses`

	//gong:ident [ref_models.RhombusGrid.Rhombuses] comment added to overcome the problem with the comment map association
	__Link__000009_Rhombuses.Identifier = `ref_models.RhombusGrid.Rhombuses`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000009_Rhombuses.Fieldtypename = `ref_models.Rhombus`
	__Link__000009_Rhombuses.FieldOffsetX = -50.000000
	__Link__000009_Rhombuses.FieldOffsetY = -16.000000
	__Link__000009_Rhombuses.TargetMultiplicity = models.MANY
	__Link__000009_Rhombuses.TargetMultiplicityOffsetX = -50.000000
	__Link__000009_Rhombuses.TargetMultiplicityOffsetY = 16.000000
	__Link__000009_Rhombuses.SourceMultiplicity = models.MANY
	__Link__000009_Rhombuses.SourceMultiplicityOffsetX = 10.000000
	__Link__000009_Rhombuses.SourceMultiplicityOffsetY = -50.000000
	__Link__000009_Rhombuses.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_Rhombuses.StartRatio = 0.500000
	__Link__000009_Rhombuses.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_Rhombuses.EndRatio = 0.194434
	__Link__000009_Rhombuses.CornerOffsetRatio = 1.966609

	__Link__000010_RotatedAxis.Name = `RotatedAxis`

	//gong:ident [ref_models.Parameter.RotatedAxis] comment added to overcome the problem with the comment map association
	__Link__000010_RotatedAxis.Identifier = `ref_models.Parameter.RotatedAxis`

	//gong:ident [ref_models.Axis] comment added to overcome the problem with the comment map association
	__Link__000010_RotatedAxis.Fieldtypename = `ref_models.Axis`
	__Link__000010_RotatedAxis.FieldOffsetX = -50.000000
	__Link__000010_RotatedAxis.FieldOffsetY = -16.000000
	__Link__000010_RotatedAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000010_RotatedAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000010_RotatedAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000010_RotatedAxis.SourceMultiplicity = models.MANY
	__Link__000010_RotatedAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000010_RotatedAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000010_RotatedAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_RotatedAxis.StartRatio = 0.500000
	__Link__000010_RotatedAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_RotatedAxis.EndRatio = 0.818720
	__Link__000010_RotatedAxis.CornerOffsetRatio = 1.380000

	__Link__000011_RotatedCircleGrid.Name = `RotatedCircleGrid`

	//gong:ident [ref_models.Parameter.RotatedCircleGrid] comment added to overcome the problem with the comment map association
	__Link__000011_RotatedCircleGrid.Identifier = `ref_models.Parameter.RotatedCircleGrid`

	//gong:ident [ref_models.CircleGrid] comment added to overcome the problem with the comment map association
	__Link__000011_RotatedCircleGrid.Fieldtypename = `ref_models.CircleGrid`
	__Link__000011_RotatedCircleGrid.FieldOffsetX = -50.000000
	__Link__000011_RotatedCircleGrid.FieldOffsetY = -16.000000
	__Link__000011_RotatedCircleGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000011_RotatedCircleGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000011_RotatedCircleGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000011_RotatedCircleGrid.SourceMultiplicity = models.MANY
	__Link__000011_RotatedCircleGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000011_RotatedCircleGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000011_RotatedCircleGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_RotatedCircleGrid.StartRatio = 0.500000
	__Link__000011_RotatedCircleGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_RotatedCircleGrid.EndRatio = 0.822394
	__Link__000011_RotatedCircleGrid.CornerOffsetRatio = 1.380000

	__Link__000012_RotatedRhombus.Name = `RotatedRhombus`

	//gong:ident [ref_models.Parameter.RotatedRhombus] comment added to overcome the problem with the comment map association
	__Link__000012_RotatedRhombus.Identifier = `ref_models.Parameter.RotatedRhombus`

	//gong:ident [ref_models.Rhombus] comment added to overcome the problem with the comment map association
	__Link__000012_RotatedRhombus.Fieldtypename = `ref_models.Rhombus`
	__Link__000012_RotatedRhombus.FieldOffsetX = -50.000000
	__Link__000012_RotatedRhombus.FieldOffsetY = -16.000000
	__Link__000012_RotatedRhombus.TargetMultiplicity = models.ZERO_ONE
	__Link__000012_RotatedRhombus.TargetMultiplicityOffsetX = -50.000000
	__Link__000012_RotatedRhombus.TargetMultiplicityOffsetY = 16.000000
	__Link__000012_RotatedRhombus.SourceMultiplicity = models.MANY
	__Link__000012_RotatedRhombus.SourceMultiplicityOffsetX = 10.000000
	__Link__000012_RotatedRhombus.SourceMultiplicityOffsetY = -50.000000
	__Link__000012_RotatedRhombus.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_RotatedRhombus.StartRatio = 0.500000
	__Link__000012_RotatedRhombus.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_RotatedRhombus.EndRatio = 0.500000
	__Link__000012_RotatedRhombus.CornerOffsetRatio = 1.380000

	__Link__000013_RotatedRhombusGrid.Name = `RotatedRhombusGrid`

	//gong:ident [ref_models.Parameter.RotatedRhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000013_RotatedRhombusGrid.Identifier = `ref_models.Parameter.RotatedRhombusGrid`

	//gong:ident [ref_models.RhombusGrid] comment added to overcome the problem with the comment map association
	__Link__000013_RotatedRhombusGrid.Fieldtypename = `ref_models.RhombusGrid`
	__Link__000013_RotatedRhombusGrid.FieldOffsetX = -50.000000
	__Link__000013_RotatedRhombusGrid.FieldOffsetY = -16.000000
	__Link__000013_RotatedRhombusGrid.TargetMultiplicity = models.ZERO_ONE
	__Link__000013_RotatedRhombusGrid.TargetMultiplicityOffsetX = -50.000000
	__Link__000013_RotatedRhombusGrid.TargetMultiplicityOffsetY = 16.000000
	__Link__000013_RotatedRhombusGrid.SourceMultiplicity = models.MANY
	__Link__000013_RotatedRhombusGrid.SourceMultiplicityOffsetX = 10.000000
	__Link__000013_RotatedRhombusGrid.SourceMultiplicityOffsetY = -50.000000
	__Link__000013_RotatedRhombusGrid.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000013_RotatedRhombusGrid.StartRatio = 0.500000
	__Link__000013_RotatedRhombusGrid.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000013_RotatedRhombusGrid.EndRatio = 0.740316
	__Link__000013_RotatedRhombusGrid.CornerOffsetRatio = 1.380000

	__Link__000014_VerticalAxis.Name = `VerticalAxis`

	//gong:ident [ref_models.Parameter.VerticalAxis] comment added to overcome the problem with the comment map association
	__Link__000014_VerticalAxis.Identifier = `ref_models.Parameter.VerticalAxis`

	//gong:ident [ref_models.VerticalAxis] comment added to overcome the problem with the comment map association
	__Link__000014_VerticalAxis.Fieldtypename = `ref_models.VerticalAxis`
	__Link__000014_VerticalAxis.FieldOffsetX = -50.000000
	__Link__000014_VerticalAxis.FieldOffsetY = -16.000000
	__Link__000014_VerticalAxis.TargetMultiplicity = models.ZERO_ONE
	__Link__000014_VerticalAxis.TargetMultiplicityOffsetX = -50.000000
	__Link__000014_VerticalAxis.TargetMultiplicityOffsetY = 16.000000
	__Link__000014_VerticalAxis.SourceMultiplicity = models.MANY
	__Link__000014_VerticalAxis.SourceMultiplicityOffsetX = 10.000000
	__Link__000014_VerticalAxis.SourceMultiplicityOffsetY = -50.000000
	__Link__000014_VerticalAxis.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000014_VerticalAxis.StartRatio = 0.500000
	__Link__000014_VerticalAxis.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000014_VerticalAxis.EndRatio = 0.500000
	__Link__000014_VerticalAxis.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_Circle.X = 524.000031
	__Position__000000_Pos_Default_Circle.Y = 765.000015
	__Position__000000_Pos_Default_Circle.Name = `Pos-Default-Circle`

	__Position__000001_Pos_Default_CircleGrid.X = 527.000031
	__Position__000001_Pos_Default_CircleGrid.Y = 884.000046
	__Position__000001_Pos_Default_CircleGrid.Name = `Pos-Default-CircleGrid`

	__Position__000002_Pos_Default_HorizontalAxis.X = 508.999969
	__Position__000002_Pos_Default_HorizontalAxis.Y = 27.000000
	__Position__000002_Pos_Default_HorizontalAxis.Name = `Pos-Default-HorizontalAxis`

	__Position__000003_Pos_Default_InitialAxis.X = 530.000000
	__Position__000003_Pos_Default_InitialAxis.Y = 1072.666641
	__Position__000003_Pos_Default_InitialAxis.Name = `Pos-Default-InitialAxis`

	__Position__000004_Pos_Default_Parameter.X = 18.000000
	__Position__000004_Pos_Default_Parameter.Y = 35.000000
	__Position__000004_Pos_Default_Parameter.Name = `Pos-Default-Parameter`

	__Position__000005_Pos_Default_Rhombus.X = 516.999969
	__Position__000005_Pos_Default_Rhombus.Y = 426.000000
	__Position__000005_Pos_Default_Rhombus.Name = `Pos-Default-Rhombus`

	__Position__000006_Pos_Default_RhombusGrid.X = 526.000000
	__Position__000006_Pos_Default_RhombusGrid.Y = 613.999893
	__Position__000006_Pos_Default_RhombusGrid.Name = `Pos-Default-RhombusGrid`

	__Position__000007_Pos_Default_VerticalAxis.X = 504.999969
	__Position__000007_Pos_Default_VerticalAxis.Y = 197.000000
	__Position__000007_Pos_Default_VerticalAxis.Name = `Pos-Default-VerticalAxis`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.X = 885.000031
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Y = 871.000031
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Name = `Verticle in class diagram Default in middle between Default-CircleGrid and Default-Circle`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.X = 885.000031
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Y = 871.000031
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle.Name = `Verticle in class diagram Default in middle between Default-CircleGrid and Default-Circle`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle.X = 631.000015
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle.Y = 484.000008
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Circle`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.X = 632.000015
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Y = 543.500023
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-CircleGrid`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.X = 632.000015
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Y = 543.500023
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-CircleGrid`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.X = 623.499985
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Y = 115.000000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-HorizontalAxis`

	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.X = 635.500000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Y = 598.833321
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`

	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.X = 625.500000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Y = 491.333321
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-InitialAxis`

	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.X = 627.499984
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Y = 314.500000
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`

	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.X = 607.999985
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Y = 121.000000
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-Rhombus`

	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.X = 628.500000
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Y = 419.499977
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`

	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.X = 631.500000
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Y = 408.499947
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-RhombusGrid`

	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.X = 621.499985
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.Y = 200.000000
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis.Name = `Verticle in class diagram Default in middle between Default-Parameter and Default-VerticalAxis`

	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.X = 879.499984
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Y = 570.999977
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`

	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.X = 879.499984
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Y = 585.999977
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus.Name = `Verticle in class diagram Default in middle between Default-RhombusGrid and Default-Rhombus`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_Parameter)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_Rhombus)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_HorizontalAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_VerticalAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_RhombusGrid)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_InitialAxis)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_CircleGrid)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Circle)
	__GongStructShape__000000_Default_Circle.Position = __Position__000000_Pos_Default_Circle
	__GongStructShape__000000_Default_Circle.Fields = append(__GongStructShape__000000_Default_Circle.Fields, __Field__000009_CenterX)
	__GongStructShape__000000_Default_Circle.Fields = append(__GongStructShape__000000_Default_Circle.Fields, __Field__000010_CenterY)
	__GongStructShape__000001_Default_CircleGrid.Position = __Position__000001_Pos_Default_CircleGrid
	__GongStructShape__000001_Default_CircleGrid.Fields = append(__GongStructShape__000001_Default_CircleGrid.Fields, __Field__000022_N)
	__GongStructShape__000001_Default_CircleGrid.Fields = append(__GongStructShape__000001_Default_CircleGrid.Fields, __Field__000021_M)
	__GongStructShape__000001_Default_CircleGrid.Links = append(__GongStructShape__000001_Default_CircleGrid.Links, __Link__000008_Reference)
	__GongStructShape__000001_Default_CircleGrid.Links = append(__GongStructShape__000001_Default_CircleGrid.Links, __Link__000000_Circles)
	__GongStructShape__000002_Default_HorizontalAxis.Position = __Position__000002_Pos_Default_HorizontalAxis
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000027_Name)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000016_IsAxisDisplayed)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000002_AxisHandleBorderLength)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000005_Axis_Length)
	__GongStructShape__000002_Default_HorizontalAxis.Fields = append(__GongStructShape__000002_Default_HorizontalAxis.Fields, __Field__000007_Axis_StrokeWidth)
	__GongStructShape__000003_Default_InitialAxis.Position = __Position__000003_Pos_Default_InitialAxis
	__GongStructShape__000003_Default_InitialAxis.Fields = append(__GongStructShape__000003_Default_InitialAxis.Fields, __Field__000026_Name)
	__GongStructShape__000003_Default_InitialAxis.Fields = append(__GongStructShape__000003_Default_InitialAxis.Fields, __Field__000001_Angle)
	__GongStructShape__000003_Default_InitialAxis.Fields = append(__GongStructShape__000003_Default_InitialAxis.Fields, __Field__000018_Length)
	__GongStructShape__000004_Default_Parameter.Position = __Position__000004_Pos_Default_Parameter
	__GongStructShape__000004_Default_Parameter.Fields = append(__GongStructShape__000004_Default_Parameter.Fields, __Field__000025_Name)
	__GongStructShape__000004_Default_Parameter.Fields = append(__GongStructShape__000004_Default_Parameter.Fields, __Field__000023_N)
	__GongStructShape__000004_Default_Parameter.Fields = append(__GongStructShape__000004_Default_Parameter.Fields, __Field__000020_M)
	__GongStructShape__000004_Default_Parameter.Fields = append(__GongStructShape__000004_Default_Parameter.Fields, __Field__000012_DiamondAngle)
	__GongStructShape__000004_Default_Parameter.Fields = append(__GongStructShape__000004_Default_Parameter.Fields, __Field__000031_OriginX)
	__GongStructShape__000004_Default_Parameter.Fields = append(__GongStructShape__000004_Default_Parameter.Fields, __Field__000032_OriginY)
	__GongStructShape__000004_Default_Parameter.Fields = append(__GongStructShape__000004_Default_Parameter.Fields, __Field__000013_DiamondSideLenght)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000005_InitialRhombus)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000001_HorizontalAxis)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000014_VerticalAxis)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000006_InitialRhombusGrid)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000002_InitialAxis)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000003_InitialCircle)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000004_InitialCircleGrid)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000010_RotatedAxis)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000012_RotatedRhombus)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000013_RotatedRhombusGrid)
	__GongStructShape__000004_Default_Parameter.Links = append(__GongStructShape__000004_Default_Parameter.Links, __Link__000011_RotatedCircleGrid)
	__GongStructShape__000005_Default_Rhombus.Position = __Position__000005_Pos_Default_Rhombus
	__GongStructShape__000005_Default_Rhombus.Fields = append(__GongStructShape__000005_Default_Rhombus.Fields, __Field__000030_Name)
	__GongStructShape__000005_Default_Rhombus.Fields = append(__GongStructShape__000005_Default_Rhombus.Fields, __Field__000017_IsDisplayed)
	__GongStructShape__000005_Default_Rhombus.Fields = append(__GongStructShape__000005_Default_Rhombus.Fields, __Field__000008_CenterX)
	__GongStructShape__000005_Default_Rhombus.Fields = append(__GongStructShape__000005_Default_Rhombus.Fields, __Field__000011_CenterY)
	__GongStructShape__000005_Default_Rhombus.Fields = append(__GongStructShape__000005_Default_Rhombus.Fields, __Field__000033_SideLength)
	__GongStructShape__000005_Default_Rhombus.Fields = append(__GongStructShape__000005_Default_Rhombus.Fields, __Field__000000_Angle)
	__GongStructShape__000005_Default_Rhombus.Fields = append(__GongStructShape__000005_Default_Rhombus.Fields, __Field__000014_InsideAngle)
	__GongStructShape__000006_Default_RhombusGrid.Position = __Position__000006_Pos_Default_RhombusGrid
	__GongStructShape__000006_Default_RhombusGrid.Fields = append(__GongStructShape__000006_Default_RhombusGrid.Fields, __Field__000028_Name)
	__GongStructShape__000006_Default_RhombusGrid.Fields = append(__GongStructShape__000006_Default_RhombusGrid.Fields, __Field__000024_N)
	__GongStructShape__000006_Default_RhombusGrid.Fields = append(__GongStructShape__000006_Default_RhombusGrid.Fields, __Field__000019_M)
	__GongStructShape__000006_Default_RhombusGrid.Links = append(__GongStructShape__000006_Default_RhombusGrid.Links, __Link__000007_Reference)
	__GongStructShape__000006_Default_RhombusGrid.Links = append(__GongStructShape__000006_Default_RhombusGrid.Links, __Link__000009_Rhombuses)
	__GongStructShape__000007_Default_VerticalAxis.Position = __Position__000007_Pos_Default_VerticalAxis
	__GongStructShape__000007_Default_VerticalAxis.Fields = append(__GongStructShape__000007_Default_VerticalAxis.Fields, __Field__000029_Name)
	__GongStructShape__000007_Default_VerticalAxis.Fields = append(__GongStructShape__000007_Default_VerticalAxis.Fields, __Field__000015_IsAxisDisplayed)
	__GongStructShape__000007_Default_VerticalAxis.Fields = append(__GongStructShape__000007_Default_VerticalAxis.Fields, __Field__000003_AxisHandleBorderLength)
	__GongStructShape__000007_Default_VerticalAxis.Fields = append(__GongStructShape__000007_Default_VerticalAxis.Fields, __Field__000004_Axis_Length)
	__GongStructShape__000007_Default_VerticalAxis.Fields = append(__GongStructShape__000007_Default_VerticalAxis.Fields, __Field__000006_Axis_StrokeWidth)
	__Link__000000_Circles.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle
	__Link__000001_HorizontalAxis.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_HorizontalAxis
	__Link__000002_InitialAxis.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis
	__Link__000003_InitialCircle.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Circle
	__Link__000004_InitialCircleGrid.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid
	__Link__000005_InitialRhombus.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus
	__Link__000006_InitialRhombusGrid.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid
	__Link__000007_Reference.Middlevertice = __Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus
	__Link__000008_Reference.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_CircleGrid_and_Default_Circle
	__Link__000009_Rhombuses.Middlevertice = __Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_RhombusGrid_and_Default_Rhombus
	__Link__000010_RotatedAxis.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_InitialAxis
	__Link__000011_RotatedCircleGrid.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_CircleGrid
	__Link__000012_RotatedRhombus.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_Rhombus
	__Link__000013_RotatedRhombusGrid.Middlevertice = __Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_RhombusGrid
	__Link__000014_VerticalAxis.Middlevertice = __Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_Parameter_and_Default_VerticalAxis
}
