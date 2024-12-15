package main

import (
	"time"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Axis__000000_Beat_Reference := (&models.Axis{}).Stage(stage)
	__Axis__000001_Construction_Axis := (&models.Axis{}).Stage(stage)
	__Axis__000002_Initial_Axis := (&models.Axis{}).Stage(stage)
	__Axis__000003_Pitch_Line := (&models.Axis{}).Stage(stage)
	__Axis__000004_Rotated_Axis := (&models.Axis{}).Stage(stage)

	__AxisGrid__000000_Beat_Lines := (&models.AxisGrid{}).Stage(stage)
	__AxisGrid__000001_Construction_Axis_Grid := (&models.AxisGrid{}).Stage(stage)
	__AxisGrid__000002_Pitch_Lines := (&models.AxisGrid{}).Stage(stage)

	__Bezier__000000_2nd_voice_seed := (&models.Bezier{}).Stage(stage)
	__Bezier__000001_First_Voice_seed := (&models.Bezier{}).Stage(stage)
	__Bezier__000002_Growth_Bezier_Right_Seed := (&models.Bezier{}).Stage(stage)
	__Bezier__000003_Growth_Curve_Next_Seed := (&models.Bezier{}).Stage(stage)
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed := (&models.Bezier{}).Stage(stage)
	__Bezier__000005_Growth_Curve_Seed := (&models.Bezier{}).Stage(stage)

	__BezierGrid__000000_2nb_Voice := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000001_2nd_voice_shifted_right := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000002_First_Voice := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000003_First_Voice_Shift_Right := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000004_Growth_Curve := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000005_Growth_Curve_Next := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000007_Growth_Curve_Shift_Right := (&models.BezierGrid{}).Stage(stage)

	__BezierGridStack__000000_The_GrowthCurveStack := (&models.BezierGridStack{}).Stage(stage)

	__Circle__000000_0 := (&models.Circle{}).Stage(stage)
	__Circle__000001_Construction_Circle := (&models.Circle{}).Stage(stage)
	__Circle__000002_First_voice_notes_seed := (&models.Circle{}).Stage(stage)
	__Circle__000003_Growing_Seed_Left := (&models.Circle{}).Stage(stage)
	__Circle__000004_Initial_Circle := (&models.Circle{}).Stage(stage)
	__Circle__000005_Rotated_Next_Circle := (&models.Circle{}).Stage(stage)

	__CircleGrid__000000_Construction_Circle_Grid := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000001_First_Voice_note_shifted_right := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000002_First_Voice_notes := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000003_Growing_Circle_Grid := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000004_Growing_Circle_Grid_Shifted_Left := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000005_Initial_Circle_Grid := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000006_Rotated_Circle_Grid := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000007_Second_Voice_Notes_Shift_Right := (&models.CircleGrid{}).Stage(stage)
	__CircleGrid__000008_Second_Voice_notes := (&models.CircleGrid{}).Stage(stage)

	__FrontCurve__000000_Non_Rotated_0_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000001_Non_Rotated_1_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000002_Non_Rotated_2_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000003_Non_Rotated_3_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000004_Non_Rotated_4_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000005_Non_Rotated_5_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000006_Non_Rotated_6_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000007_Non_Rotated_7_ := (&models.FrontCurve{}).Stage(stage)

	__FrontCurveStack__000000_Front_Curve_Stack := (&models.FrontCurveStack{}).Stage(stage)

	__HorizontalAxis__000000_Horizontal_Axis := (&models.HorizontalAxis{}).Stage(stage)

	__Key__000000_F_key := (&models.Key{}).Stage(stage)

	__MovingLine__000000_Cursor := (&models.MovingLine{}).Stage(stage)

	__NoteInfo__000000_0 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000001_1 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000002_10 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000003_11 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000004_12 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000005_13 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000006_14 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000007_15 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000008_16 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000009_17 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000010_18 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000011_19 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000012_2 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000013_20 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000014_21 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000015_3 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000016_4 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000017_5 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000018_6 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000019_7 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000020_8 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000021_9 := (&models.NoteInfo{}).Stage(stage)

	__Parameter__000000_Reference := (&models.Parameter{}).Stage(stage)

	__Rhombus__000000_Growing_Rhombus_Grid_Seed := (&models.Rhombus{}).Stage(stage)
	__Rhombus__000001_Initial_Rhombus := (&models.Rhombus{}).Stage(stage)
	__Rhombus__000002_Rotated_Next_Rhombus := (&models.Rhombus{}).Stage(stage)
	__Rhombus__000003_Rotated_Rhombus := (&models.Rhombus{}).Stage(stage)

	__RhombusGrid__000000_Growing_Rhombus_Grid := (&models.RhombusGrid{}).Stage(stage)
	__RhombusGrid__000001_Initial_Rhombus_Grid := (&models.RhombusGrid{}).Stage(stage)
	__RhombusGrid__000002_Rotated_Rhombus_Grid := (&models.RhombusGrid{}).Stage(stage)

	__ShapeCategory__000000_0_Axes := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000001_1_Initial := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000002_2_Rotated := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000003_3_Growing := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000004_4_Construction := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000005_5_Vertical_growth := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000006_6_Spiral_growth := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000007_7_Spiral_Growth_Bezier := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000008_8_Score_notation := (&models.ShapeCategory{}).Stage(stage)
	__ShapeCategory__000009_9_Composer := (&models.ShapeCategory{}).Stage(stage)

	__SpiralBezier__000000_Spiral_Bezier_Seed := (&models.SpiralBezier{}).Stage(stage)

	__SpiralBezierGrid__000000_Spiral_Bezier_Full_Grid := (&models.SpiralBezierGrid{}).Stage(stage)
	__SpiralBezierGrid__000001_Spiral_Bezier_Grid := (&models.SpiralBezierGrid{}).Stage(stage)

	__SpiralCircle__000000_Construction_Circle_Spiral := (&models.SpiralCircle{}).Stage(stage)

	__SpiralCircleGrid__000000_Brute_Spiral_Bezier_Circle_Grid := (&models.SpiralCircleGrid{}).Stage(stage)
	__SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid := (&models.SpiralCircleGrid{}).Stage(stage)
	__SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid := (&models.SpiralCircleGrid{}).Stage(stage)
	__SpiralCircleGrid__000003_Spiral_Circle_Grid := (&models.SpiralCircleGrid{}).Stage(stage)

	__SpiralLine__000000_Spiral_Contruction_Inner_Line := (&models.SpiralLine{}).Stage(stage)
	__SpiralLine__000001_Spiral_Contruction_Outer_Line := (&models.SpiralLine{}).Stage(stage)

	__SpiralLineGrid__000000_Spiral_Construction_Inner_Line_Grid_Spiral := (&models.SpiralLineGrid{}).Stage(stage)
	__SpiralLineGrid__000001_Spiral_Construction_Outer_Line_Full_Grid_Spiral := (&models.SpiralLineGrid{}).Stage(stage)
	__SpiralLineGrid__000002_Spiral_Construction_Outer_Line_Grid_Spiral := (&models.SpiralLineGrid{}).Stage(stage)

	__SpiralOrigin__000000_Spiral_Origin := (&models.SpiralOrigin{}).Stage(stage)

	__SpiralRhombus__000000_Reference_Spiral_Rhombus := (&models.SpiralRhombus{}).Stage(stage)

	__SpiralRhombusGrid__000000_Spiral_Rhombus_Grid := (&models.SpiralRhombusGrid{}).Stage(stage)

	__VerticalAxis__000000_Vertical_Axis := (&models.VerticalAxis{}).Stage(stage)

	// Setup of values

	__Axis__000000_Beat_Reference.Name = `Beat Reference`
	__Axis__000000_Beat_Reference.IsDisplayed = false
	__Axis__000000_Beat_Reference.AngleDegree = 90.000000
	__Axis__000000_Beat_Reference.Length = 2000.000000
	__Axis__000000_Beat_Reference.CenterX = 0.000000
	__Axis__000000_Beat_Reference.CenterY = 0.000000
	__Axis__000000_Beat_Reference.EndX = 0.000000
	__Axis__000000_Beat_Reference.EndY = 0.000000
	__Axis__000000_Beat_Reference.Color = ``
	__Axis__000000_Beat_Reference.FillOpacity = 0.000000
	__Axis__000000_Beat_Reference.Stroke = `grey`
	__Axis__000000_Beat_Reference.StrokeOpacity = 0.800000
	__Axis__000000_Beat_Reference.StrokeWidth = 1.000000
	__Axis__000000_Beat_Reference.StrokeDashArray = ``
	__Axis__000000_Beat_Reference.StrokeDashArrayWhenSelected = ``
	__Axis__000000_Beat_Reference.Transform = ``

	__Axis__000001_Construction_Axis.Name = `Construction Axis`
	__Axis__000001_Construction_Axis.IsDisplayed = false
	__Axis__000001_Construction_Axis.AngleDegree = 97.682917
	__Axis__000001_Construction_Axis.Length = 114.075352
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -15.250814
	__Axis__000001_Construction_Axis.EndY = 113.051310
	__Axis__000001_Construction_Axis.Color = ``
	__Axis__000001_Construction_Axis.FillOpacity = 0.000000
	__Axis__000001_Construction_Axis.Stroke = `blue`
	__Axis__000001_Construction_Axis.StrokeOpacity = 0.700000
	__Axis__000001_Construction_Axis.StrokeWidth = 2.000000
	__Axis__000001_Construction_Axis.StrokeDashArray = ``
	__Axis__000001_Construction_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000001_Construction_Axis.Transform = ``

	__Axis__000002_Initial_Axis.Name = `Initial Axis`
	__Axis__000002_Initial_Axis.IsDisplayed = false
	__Axis__000002_Initial_Axis.AngleDegree = 82.317083
	__Axis__000002_Initial_Axis.Length = 426.639056
	__Axis__000002_Initial_Axis.CenterX = 0.000000
	__Axis__000002_Initial_Axis.CenterY = 0.000000
	__Axis__000002_Initial_Axis.EndX = 0.000000
	__Axis__000002_Initial_Axis.EndY = 0.000000
	__Axis__000002_Initial_Axis.Color = ``
	__Axis__000002_Initial_Axis.FillOpacity = 0.000000
	__Axis__000002_Initial_Axis.Stroke = `black`
	__Axis__000002_Initial_Axis.StrokeOpacity = 1.000000
	__Axis__000002_Initial_Axis.StrokeWidth = 2.000000
	__Axis__000002_Initial_Axis.StrokeDashArray = ``
	__Axis__000002_Initial_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000002_Initial_Axis.Transform = ``

	__Axis__000003_Pitch_Line.Name = `Pitch Line`
	__Axis__000003_Pitch_Line.IsDisplayed = false
	__Axis__000003_Pitch_Line.AngleDegree = 0.000000
	__Axis__000003_Pitch_Line.Length = 2000.000000
	__Axis__000003_Pitch_Line.CenterX = 0.000000
	__Axis__000003_Pitch_Line.CenterY = 0.000000
	__Axis__000003_Pitch_Line.EndX = 0.000000
	__Axis__000003_Pitch_Line.EndY = 0.000000
	__Axis__000003_Pitch_Line.Color = ``
	__Axis__000003_Pitch_Line.FillOpacity = 0.000000
	__Axis__000003_Pitch_Line.Stroke = `grey`
	__Axis__000003_Pitch_Line.StrokeOpacity = 0.800000
	__Axis__000003_Pitch_Line.StrokeWidth = 1.000000
	__Axis__000003_Pitch_Line.StrokeDashArray = ``
	__Axis__000003_Pitch_Line.StrokeDashArrayWhenSelected = ``
	__Axis__000003_Pitch_Line.Transform = ``

	__Axis__000004_Rotated_Axis.Name = `Rotated Axis`
	__Axis__000004_Rotated_Axis.IsDisplayed = false
	__Axis__000004_Rotated_Axis.AngleDegree = 0.000000
	__Axis__000004_Rotated_Axis.Length = 426.639056
	__Axis__000004_Rotated_Axis.CenterX = 0.000000
	__Axis__000004_Rotated_Axis.CenterY = 0.000000
	__Axis__000004_Rotated_Axis.EndX = -29.670211
	__Axis__000004_Rotated_Axis.EndY = 154.170939
	__Axis__000004_Rotated_Axis.Color = ``
	__Axis__000004_Rotated_Axis.FillOpacity = 0.000000
	__Axis__000004_Rotated_Axis.Stroke = `black`
	__Axis__000004_Rotated_Axis.StrokeOpacity = 1.000000
	__Axis__000004_Rotated_Axis.StrokeWidth = 2.000000
	__Axis__000004_Rotated_Axis.StrokeDashArray = ``
	__Axis__000004_Rotated_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000004_Rotated_Axis.Transform = ``

	__AxisGrid__000000_Beat_Lines.Name = `Beat Lines`
	__AxisGrid__000000_Beat_Lines.IsDisplayed = false

	__AxisGrid__000001_Construction_Axis_Grid.Name = `Construction Axis Grid`
	__AxisGrid__000001_Construction_Axis_Grid.IsDisplayed = false

	__AxisGrid__000002_Pitch_Lines.Name = `Pitch Lines`
	__AxisGrid__000002_Pitch_Lines.IsDisplayed = false

	__Bezier__000000_2nd_voice_seed.Name = `2nd voice seed`
	__Bezier__000000_2nd_voice_seed.IsDisplayed = false
	__Bezier__000000_2nd_voice_seed.StartX = 0.000000
	__Bezier__000000_2nd_voice_seed.StartY = 0.000000
	__Bezier__000000_2nd_voice_seed.ControlPointStartX = 0.000000
	__Bezier__000000_2nd_voice_seed.ControlPointStartY = 0.000000
	__Bezier__000000_2nd_voice_seed.EndX = 0.000000
	__Bezier__000000_2nd_voice_seed.EndY = 0.000000
	__Bezier__000000_2nd_voice_seed.ControlPointEndX = 0.000000
	__Bezier__000000_2nd_voice_seed.ControlPointEndY = 0.000000
	__Bezier__000000_2nd_voice_seed.Color = ``
	__Bezier__000000_2nd_voice_seed.FillOpacity = 0.000000
	__Bezier__000000_2nd_voice_seed.Stroke = `red`
	__Bezier__000000_2nd_voice_seed.StrokeOpacity = 0.800000
	__Bezier__000000_2nd_voice_seed.StrokeWidth = 5.000000
	__Bezier__000000_2nd_voice_seed.StrokeDashArray = ``
	__Bezier__000000_2nd_voice_seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000000_2nd_voice_seed.Transform = ``

	__Bezier__000001_First_Voice_seed.Name = `First Voice seed`
	__Bezier__000001_First_Voice_seed.IsDisplayed = false
	__Bezier__000001_First_Voice_seed.StartX = 0.000000
	__Bezier__000001_First_Voice_seed.StartY = 0.000000
	__Bezier__000001_First_Voice_seed.ControlPointStartX = 0.000000
	__Bezier__000001_First_Voice_seed.ControlPointStartY = 0.000000
	__Bezier__000001_First_Voice_seed.EndX = 0.000000
	__Bezier__000001_First_Voice_seed.EndY = 0.000000
	__Bezier__000001_First_Voice_seed.ControlPointEndX = 0.000000
	__Bezier__000001_First_Voice_seed.ControlPointEndY = 0.000000
	__Bezier__000001_First_Voice_seed.Color = ``
	__Bezier__000001_First_Voice_seed.FillOpacity = 0.000000
	__Bezier__000001_First_Voice_seed.Stroke = `grey`
	__Bezier__000001_First_Voice_seed.StrokeOpacity = 0.800000
	__Bezier__000001_First_Voice_seed.StrokeWidth = 6.000000
	__Bezier__000001_First_Voice_seed.StrokeDashArray = ``
	__Bezier__000001_First_Voice_seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000001_First_Voice_seed.Transform = ``

	__Bezier__000002_Growth_Bezier_Right_Seed.Name = `Growth Bezier Right Seed`
	__Bezier__000002_Growth_Bezier_Right_Seed.IsDisplayed = false
	__Bezier__000002_Growth_Bezier_Right_Seed.StartX = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.StartY = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.ControlPointStartX = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.ControlPointStartY = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.EndX = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.EndY = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.ControlPointEndX = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.ControlPointEndY = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.Color = ``
	__Bezier__000002_Growth_Bezier_Right_Seed.FillOpacity = 0.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.Stroke = `green`
	__Bezier__000002_Growth_Bezier_Right_Seed.StrokeOpacity = 0.500000
	__Bezier__000002_Growth_Bezier_Right_Seed.StrokeWidth = 6.000000
	__Bezier__000002_Growth_Bezier_Right_Seed.StrokeDashArray = ``
	__Bezier__000002_Growth_Bezier_Right_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000002_Growth_Bezier_Right_Seed.Transform = ``

	__Bezier__000003_Growth_Curve_Next_Seed.Name = `Growth Curve Next Seed`
	__Bezier__000003_Growth_Curve_Next_Seed.IsDisplayed = false
	__Bezier__000003_Growth_Curve_Next_Seed.StartX = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.StartY = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.ControlPointStartX = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.ControlPointStartY = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.EndX = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.EndY = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.ControlPointEndX = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.ControlPointEndY = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.Color = ``
	__Bezier__000003_Growth_Curve_Next_Seed.FillOpacity = 0.000000
	__Bezier__000003_Growth_Curve_Next_Seed.Stroke = `red`
	__Bezier__000003_Growth_Curve_Next_Seed.StrokeOpacity = 0.500000
	__Bezier__000003_Growth_Curve_Next_Seed.StrokeWidth = 6.000000
	__Bezier__000003_Growth_Curve_Next_Seed.StrokeDashArray = ``
	__Bezier__000003_Growth_Curve_Next_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000003_Growth_Curve_Next_Seed.Transform = ``

	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.Name = `Growth Curve Next Shift Right Seed`
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.IsDisplayed = false
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.StartX = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.StartY = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.ControlPointStartX = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.ControlPointStartY = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.EndX = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.EndY = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.ControlPointEndX = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.ControlPointEndY = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.Color = ``
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.FillOpacity = 0.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.Stroke = `red`
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.StrokeOpacity = 0.500000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.StrokeWidth = 6.000000
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.StrokeDashArray = ``
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.Transform = ``

	__Bezier__000005_Growth_Curve_Seed.Name = `Growth Curve Seed`
	__Bezier__000005_Growth_Curve_Seed.IsDisplayed = false
	__Bezier__000005_Growth_Curve_Seed.StartX = -7.625407
	__Bezier__000005_Growth_Curve_Seed.StartY = 56.525655
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 36.851710
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 62.525694
	__Bezier__000005_Growth_Curve_Seed.EndX = 68.551915
	__Bezier__000005_Growth_Curve_Seed.EndY = 124.356441
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 24.074799
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 118.356402
	__Bezier__000005_Growth_Curve_Seed.Color = ``
	__Bezier__000005_Growth_Curve_Seed.FillOpacity = 0.000000
	__Bezier__000005_Growth_Curve_Seed.Stroke = `grey`
	__Bezier__000005_Growth_Curve_Seed.StrokeOpacity = 0.800000
	__Bezier__000005_Growth_Curve_Seed.StrokeWidth = 6.000000
	__Bezier__000005_Growth_Curve_Seed.StrokeDashArray = ``
	__Bezier__000005_Growth_Curve_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000005_Growth_Curve_Seed.Transform = ``

	__BezierGrid__000000_2nb_Voice.Name = `2nb Voice`
	__BezierGrid__000000_2nb_Voice.IsDisplayed = true

	__BezierGrid__000001_2nd_voice_shifted_right.Name = `2nd voice shifted right`
	__BezierGrid__000001_2nd_voice_shifted_right.IsDisplayed = true

	__BezierGrid__000002_First_Voice.Name = `First Voice`
	__BezierGrid__000002_First_Voice.IsDisplayed = true

	__BezierGrid__000003_First_Voice_Shift_Right.Name = `First Voice Shift Right`
	__BezierGrid__000003_First_Voice_Shift_Right.IsDisplayed = true

	__BezierGrid__000004_Growth_Curve.Name = `Growth Curve`
	__BezierGrid__000004_Growth_Curve.IsDisplayed = false

	__BezierGrid__000005_Growth_Curve_Next.Name = `Growth Curve Next`
	__BezierGrid__000005_Growth_Curve_Next.IsDisplayed = false

	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.Name = `Growth Curve Next Shift Right`
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.IsDisplayed = false

	__BezierGrid__000007_Growth_Curve_Shift_Right.Name = `Growth Curve Shift Right`
	__BezierGrid__000007_Growth_Curve_Shift_Right.IsDisplayed = false

	__BezierGridStack__000000_The_GrowthCurveStack.Name = `The GrowthCurveStack`
	__BezierGridStack__000000_The_GrowthCurveStack.IsDisplayed = false

	__Circle__000000_0.Name = `0`
	__Circle__000000_0.IsDisplayed = false
	__Circle__000000_0.CenterX = 0.000000
	__Circle__000000_0.CenterY = 0.000000
	__Circle__000000_0.HasBespokeRadius = false
	__Circle__000000_0.BespopkeRadius = 0.000000
	__Circle__000000_0.Pitch = 0
	__Circle__000000_0.Color = ``
	__Circle__000000_0.FillOpacity = 0.000000
	__Circle__000000_0.Stroke = `blue`
	__Circle__000000_0.StrokeOpacity = 0.700000
	__Circle__000000_0.StrokeWidth = 2.000000
	__Circle__000000_0.StrokeDashArray = ``
	__Circle__000000_0.StrokeDashArrayWhenSelected = ``
	__Circle__000000_0.Transform = ``
	__Circle__000000_0.ShowName = false

	__Circle__000001_Construction_Circle.Name = `Construction Circle`
	__Circle__000001_Construction_Circle.IsDisplayed = false
	__Circle__000001_Construction_Circle.CenterX = -7.625407
	__Circle__000001_Construction_Circle.CenterY = 56.525655
	__Circle__000001_Construction_Circle.HasBespokeRadius = true
	__Circle__000001_Construction_Circle.BespopkeRadius = 20.000000
	__Circle__000001_Construction_Circle.Pitch = 0
	__Circle__000001_Construction_Circle.Color = ``
	__Circle__000001_Construction_Circle.FillOpacity = 0.000000
	__Circle__000001_Construction_Circle.Stroke = `blue`
	__Circle__000001_Construction_Circle.StrokeOpacity = 0.800000
	__Circle__000001_Construction_Circle.StrokeWidth = 2.000000
	__Circle__000001_Construction_Circle.StrokeDashArray = ``
	__Circle__000001_Construction_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000001_Construction_Circle.Transform = ``
	__Circle__000001_Construction_Circle.ShowName = false

	__Circle__000002_First_voice_notes_seed.Name = `First voice notes seed`
	__Circle__000002_First_voice_notes_seed.IsDisplayed = false
	__Circle__000002_First_voice_notes_seed.CenterX = 0.000000
	__Circle__000002_First_voice_notes_seed.CenterY = 0.000000
	__Circle__000002_First_voice_notes_seed.HasBespokeRadius = true
	__Circle__000002_First_voice_notes_seed.BespopkeRadius = 10.000000
	__Circle__000002_First_voice_notes_seed.Pitch = 0
	__Circle__000002_First_voice_notes_seed.Color = ``
	__Circle__000002_First_voice_notes_seed.FillOpacity = 0.000000
	__Circle__000002_First_voice_notes_seed.Stroke = `grey`
	__Circle__000002_First_voice_notes_seed.StrokeOpacity = 0.800000
	__Circle__000002_First_voice_notes_seed.StrokeWidth = 3.000000
	__Circle__000002_First_voice_notes_seed.StrokeDashArray = ``
	__Circle__000002_First_voice_notes_seed.StrokeDashArrayWhenSelected = ``
	__Circle__000002_First_voice_notes_seed.Transform = ``
	__Circle__000002_First_voice_notes_seed.ShowName = false

	__Circle__000003_Growing_Seed_Left.Name = `Growing Seed Left`
	__Circle__000003_Growing_Seed_Left.IsDisplayed = false
	__Circle__000003_Growing_Seed_Left.CenterX = 0.000000
	__Circle__000003_Growing_Seed_Left.CenterY = 0.000000
	__Circle__000003_Growing_Seed_Left.HasBespokeRadius = false
	__Circle__000003_Growing_Seed_Left.BespopkeRadius = 0.000000
	__Circle__000003_Growing_Seed_Left.Pitch = 0
	__Circle__000003_Growing_Seed_Left.Color = ``
	__Circle__000003_Growing_Seed_Left.FillOpacity = 0.000000
	__Circle__000003_Growing_Seed_Left.Stroke = `green`
	__Circle__000003_Growing_Seed_Left.StrokeOpacity = 0.800000
	__Circle__000003_Growing_Seed_Left.StrokeWidth = 1.800000
	__Circle__000003_Growing_Seed_Left.StrokeDashArray = ``
	__Circle__000003_Growing_Seed_Left.StrokeDashArrayWhenSelected = ``
	__Circle__000003_Growing_Seed_Left.Transform = ``
	__Circle__000003_Growing_Seed_Left.ShowName = false

	__Circle__000004_Initial_Circle.Name = `Initial Circle`
	__Circle__000004_Initial_Circle.IsDisplayed = false
	__Circle__000004_Initial_Circle.CenterX = 0.000000
	__Circle__000004_Initial_Circle.CenterY = 0.000000
	__Circle__000004_Initial_Circle.HasBespokeRadius = false
	__Circle__000004_Initial_Circle.BespopkeRadius = 0.000000
	__Circle__000004_Initial_Circle.Pitch = 0
	__Circle__000004_Initial_Circle.Color = ``
	__Circle__000004_Initial_Circle.FillOpacity = 0.000000
	__Circle__000004_Initial_Circle.Stroke = `lightblue`
	__Circle__000004_Initial_Circle.StrokeOpacity = 0.800000
	__Circle__000004_Initial_Circle.StrokeWidth = 3.000000
	__Circle__000004_Initial_Circle.StrokeDashArray = ``
	__Circle__000004_Initial_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000004_Initial_Circle.Transform = ``
	__Circle__000004_Initial_Circle.ShowName = false

	__Circle__000005_Rotated_Next_Circle.Name = `Rotated Next Circle`
	__Circle__000005_Rotated_Next_Circle.IsDisplayed = false
	__Circle__000005_Rotated_Next_Circle.CenterX = 167.605459
	__Circle__000005_Rotated_Next_Circle.CenterY = 22.610262
	__Circle__000005_Rotated_Next_Circle.HasBespokeRadius = false
	__Circle__000005_Rotated_Next_Circle.BespopkeRadius = 0.000000
	__Circle__000005_Rotated_Next_Circle.Pitch = 0
	__Circle__000005_Rotated_Next_Circle.Color = ``
	__Circle__000005_Rotated_Next_Circle.FillOpacity = 0.000000
	__Circle__000005_Rotated_Next_Circle.Stroke = `yellow`
	__Circle__000005_Rotated_Next_Circle.StrokeOpacity = 0.800000
	__Circle__000005_Rotated_Next_Circle.StrokeWidth = 3.000000
	__Circle__000005_Rotated_Next_Circle.StrokeDashArray = ``
	__Circle__000005_Rotated_Next_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000005_Rotated_Next_Circle.Transform = ``
	__Circle__000005_Rotated_Next_Circle.ShowName = false

	__CircleGrid__000000_Construction_Circle_Grid.Name = `Construction Circle Grid`
	__CircleGrid__000000_Construction_Circle_Grid.IsDisplayed = false

	__CircleGrid__000001_First_Voice_note_shifted_right.Name = `First Voice note shifted right`
	__CircleGrid__000001_First_Voice_note_shifted_right.IsDisplayed = true

	__CircleGrid__000002_First_Voice_notes.Name = `First Voice notes`
	__CircleGrid__000002_First_Voice_notes.IsDisplayed = true

	__CircleGrid__000003_Growing_Circle_Grid.Name = `Growing Circle Grid`
	__CircleGrid__000003_Growing_Circle_Grid.IsDisplayed = false

	__CircleGrid__000004_Growing_Circle_Grid_Shifted_Left.Name = `Growing Circle Grid Shifted Left`
	__CircleGrid__000004_Growing_Circle_Grid_Shifted_Left.IsDisplayed = false

	__CircleGrid__000005_Initial_Circle_Grid.Name = `Initial Circle Grid`
	__CircleGrid__000005_Initial_Circle_Grid.IsDisplayed = false

	__CircleGrid__000006_Rotated_Circle_Grid.Name = `Rotated Circle Grid`
	__CircleGrid__000006_Rotated_Circle_Grid.IsDisplayed = false

	__CircleGrid__000007_Second_Voice_Notes_Shift_Right.Name = `Second Voice Notes Shift Right`
	__CircleGrid__000007_Second_Voice_Notes_Shift_Right.IsDisplayed = true

	__CircleGrid__000008_Second_Voice_notes.Name = `Second Voice notes`
	__CircleGrid__000008_Second_Voice_notes.IsDisplayed = true

	__FrontCurve__000000_Non_Rotated_0_.Name = `Non Rotated 0 `
	__FrontCurve__000000_Non_Rotated_0_.Path = `M664.40 519.07 C665.37 513.49 666.00 508.90 666.70 503.72 C667.40 498.55 667.98 493.36 668.61 488.03 C669.24 482.71 669.79 477.36 670.48 471.77 C671.18 466.18 671.86 460.54 672.79 454.51 C673.71 448.48 674.88 442.25 676.04 435.56 C677.21 428.88 678.94 421.66 679.77 414.40 C680.59 407.14 681.41 399.35 680.98 392.01 C680.55 384.68 679.24 377.31 677.19 370.40 C675.15 363.49 672.14 356.83 668.71 350.55 C665.27 344.27 661.09 338.31 656.58 332.73 C652.08 327.15 647.00 321.91 641.67 317.06 C636.35 312.21 631.24 308.03 624.63 303.62 C618.02 299.22 609.84 294.30 602.00 290.62 C594.15 286.94 585.89 283.82 577.55 281.54 C569.20 279.26 560.53 277.66 551.95 276.96 C543.37 276.25 534.58 276.36 526.07 277.30 C517.57 278.24 509.03 280.13 500.91 282.58 C492.80 285.02 484.89 288.44 477.39 291.97 C469.90 295.51 462.79 299.72 455.96 303.79 C449.12 307.86 442.69 312.18 436.39 316.38 C430.09 320.58 424.09 324.78 418.17 328.99 C412.25 333.19 406.52 337.35 400.87 341.61 C395.22 345.88 389.69 350.15 384.27 354.59 C378.84 359.03 373.11 363.97 368.32 368.28 C363.52 372.58 359.74 376.27 355.52 380.42 C351.29 384.57 347.17 388.84 342.96 393.20 C338.74 397.56 334.61 402.01 330.24 406.57 C325.87 411.12 321.50 415.75 316.73 420.54 C311.96 425.33 306.91 430.13 301.60 435.31 C296.30 440.48 290.25 445.73 284.92 451.59 C279.59 457.44 273.96 463.72 269.63 470.42 C265.29 477.12 261.66 484.41 258.92 491.79 C256.18 499.18 254.38 506.97 253.19 514.73 C251.99 522.48 251.62 530.45 251.77 538.32 C251.91 546.19 252.75 554.15 254.06 561.94 C255.37 569.74 256.93 576.80 259.61 585.10 C262.30 593.40 265.96 603.19 270.17 611.72 C274.38 620.25 279.30 628.59 284.85 636.28 C290.40 643.98 296.68 651.33 303.47 657.89 C310.26 664.45 317.77 670.48 325.60 675.65 C333.43 680.82 341.92 685.23 350.45 688.91 C358.97 692.58 367.97 695.34 376.75 697.70 C385.52 700.05 394.44 701.60 403.08 703.05 C411.71 704.50 420.22 705.46 428.56 706.38 C436.89 707.30 445.01 707.99 453.08 708.58 C461.15 709.17 469.05 709.64 476.96 709.92 C484.87 710.20 492.69 710.35 500.54 710.26 C508.39 710.17 516.25 710.10 524.05 709.36 C531.86 708.61 539.76 707.65 547.37 705.82 C554.98 703.98 562.58 701.50 569.73 698.36 C576.88 695.21 583.88 691.38 590.28 686.95 C596.67 682.52 602.74 677.36 608.10 671.77 C613.46 666.19 618.27 659.88 622.44 653.44 C626.60 647.00 630.03 639.97 633.10 633.14 C636.16 626.30 638.53 619.22 640.82 612.44 C643.11 605.67 644.96 598.98 646.82 592.49 C648.68 586.00 650.35 579.73 651.99 573.52 C653.63 567.31 655.20 561.28 656.68 555.22 C658.16 549.17 659.59 543.22 660.88 537.19 C662.16 531.16 663.43 524.64 664.40 519.07 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M368.32 368.28 C363.52 372.58 359.74 376.27 355.52 380.42 C351.29 384.57 347.17 388.84 342.96 393.20 C338.74 397.56 334.61 402.01 330.24 406.57 C325.87 411.12 321.50 415.75 316.73 420.54 C311.96 425.33 306.91 430.13 301.60 435.31 C296.30 440.48 290.25 445.73 284.92 451.59 C279.59 457.44 273.96 463.72 269.63 470.42 C265.29 477.12 261.66 484.41 258.92 491.79 C256.18 499.18 254.38 506.97 253.19 514.73 C251.99 522.48 251.62 530.45 251.77 538.32 C251.91 546.19 252.75 554.15 254.06 561.94 C255.37 569.74 256.93 576.80 259.61 585.10 C262.30 593.40 265.96 603.19 270.17 611.72 C274.38 620.25 279.30 628.59 284.85 636.28 C290.40 643.98 296.68 651.33 303.47 657.89 C310.26 664.45 317.77 670.48 325.60 675.65 C333.43 680.82 341.92 685.23 350.45 688.91 C358.97 692.58 367.97 695.34 376.75 697.70 C385.52 700.05 394.44 701.60 403.08 703.05 C411.71 704.50 420.22 705.46 428.56 706.38 C436.89 707.30 445.01 707.99 453.08 708.58 C461.15 709.17 469.05 709.64 476.96 709.92 C484.87 710.20 492.69 710.35 500.54 710.26 C508.39 710.17 516.86 709.81 524.05 709.36 C531.25 708.90 537.12 708.26 543.70 707.50 C550.27 706.74 556.84 705.80 563.51 704.80 C570.19 703.81 576.86 702.67 583.76 701.56 C590.66 700.45 597.59 699.25 604.92 698.15 C612.26 697.04 619.80 696.07 627.75 694.94 C635.71 693.81 644.28 693.06 652.64 691.36 C660.99 689.66 669.83 687.79 677.90 684.75 C685.96 681.72 693.83 677.77 701.03 673.16 C708.23 668.56 714.93 663.02 721.10 657.12 C727.28 651.22 732.92 644.61 738.07 637.75 C743.21 630.89 747.84 623.52 751.98 615.96 C756.11 608.40 759.55 601.31 762.88 592.41 C766.20 583.51 769.71 572.68 771.94 562.56 C774.16 552.43 775.66 542.00 776.22 531.68 C776.78 521.35 776.52 510.83 775.31 500.62 C774.10 490.41 771.95 480.15 768.97 470.39 C765.99 460.62 761.97 451.03 757.42 442.02 C752.86 433.01 747.32 424.41 741.67 416.30 C736.01 408.19 729.69 400.63 723.48 393.35 C717.28 386.07 710.82 379.27 704.44 372.63 C698.06 366.00 691.66 359.70 685.20 353.54 C678.75 347.39 672.30 341.46 665.71 335.72 C659.11 329.97 652.47 324.40 645.62 319.05 C638.77 313.70 631.90 308.36 624.63 303.62 C617.36 298.88 609.84 294.30 602.00 290.62 C594.15 286.94 585.89 283.82 577.55 281.54 C569.20 279.26 560.53 277.66 551.95 276.96 C543.37 276.25 534.58 276.36 526.07 277.30 C517.57 278.24 509.03 280.13 500.91 282.58 C492.80 285.02 484.89 288.44 477.39 291.97 C469.90 295.51 462.79 299.72 455.96 303.79 C449.12 307.86 442.69 312.18 436.39 316.38 C430.09 320.58 424.09 324.78 418.17 328.99 C412.25 333.19 406.52 337.35 400.87 341.61 C395.22 345.88 389.69 350.15 384.27 354.59 C378.84 359.03 373.11 363.97 368.32 368.28 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M524.05 709.36 C531.25 708.90 537.12 708.26 543.70 707.50 C550.27 706.74 556.84 705.80 563.51 704.80 C570.19 703.81 576.86 702.67 583.76 701.56 C590.66 700.45 597.59 699.25 604.92 698.15 C612.26 697.04 619.80 696.07 627.75 694.94 C635.71 693.81 644.28 693.06 652.64 691.36 C660.99 689.66 669.83 687.79 677.90 684.75 C685.96 681.72 693.83 677.77 701.03 673.16 C708.23 668.56 714.93 663.02 721.10 657.12 C727.28 651.22 732.92 644.61 738.07 637.75 C743.21 630.89 747.84 623.52 751.98 615.96 C756.11 608.40 759.55 601.31 762.88 592.41 C766.20 583.51 769.71 572.68 771.94 562.56 C774.16 552.43 775.66 542.00 776.22 531.68 C776.78 521.35 776.52 510.83 775.31 500.62 C774.10 490.41 771.95 480.15 768.97 470.39 C765.99 460.62 761.97 451.03 757.42 442.02 C752.86 433.01 747.32 424.41 741.67 416.30 C736.01 408.19 729.69 400.63 723.48 393.35 C717.28 386.07 710.82 379.27 704.44 372.63 C698.06 366.00 691.66 359.70 685.20 353.54 C678.75 347.39 672.30 341.46 665.71 335.72 C659.11 329.97 652.47 324.40 645.62 319.05 C638.77 313.70 631.18 308.17 624.63 303.62 C618.07 299.08 612.53 295.60 606.30 291.78 C600.06 287.95 593.70 284.31 587.21 280.67 C580.73 277.02 574.14 273.53 567.38 269.90 C560.62 266.27 553.78 262.72 546.67 258.88 C539.57 255.04 532.37 250.99 524.75 246.85 C517.14 242.71 509.27 237.92 500.98 234.03 C492.69 230.15 483.91 226.13 475.02 223.54 C466.13 220.95 456.83 219.24 447.64 218.49 C438.45 217.74 429.08 218.05 419.90 219.02 C410.72 219.99 401.52 221.87 392.57 224.29 C383.62 226.72 374.75 229.90 366.18 233.56 C357.60 237.22 349.93 240.94 341.12 246.25 C332.32 251.57 322.06 258.39 313.34 265.46 C304.62 272.53 296.30 280.34 288.82 288.68 C281.33 297.01 274.40 306.05 268.44 315.46 C262.48 324.87 257.26 334.93 253.05 345.15 C248.83 355.37 245.57 366.13 243.13 376.78 C240.70 387.43 239.34 398.40 238.42 409.07 C237.50 419.74 237.48 430.42 237.60 440.80 C237.72 451.18 238.38 461.37 239.14 471.37 C239.89 481.38 240.91 491.15 242.12 500.83 C243.33 510.51 244.72 520.01 246.39 529.46 C248.05 538.91 249.91 548.24 252.12 557.51 C254.32 566.79 256.60 576.07 259.61 585.10 C262.62 594.13 265.96 603.19 270.17 611.72 C274.38 620.25 279.30 628.59 284.85 636.28 C290.40 643.98 296.68 651.33 303.47 657.89 C310.26 664.45 317.77 670.48 325.60 675.65 C333.43 680.82 341.92 685.23 350.45 688.91 C358.97 692.58 367.97 695.34 376.75 697.70 C385.52 700.05 394.44 701.60 403.08 703.05 C411.71 704.50 420.22 705.46 428.56 706.38 C436.89 707.30 445.01 707.99 453.08 708.58 C461.15 709.17 469.05 709.64 476.96 709.92 C484.87 710.20 492.69 710.35 500.54 710.26 C508.39 710.17 516.86 709.81 524.05 709.36 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M624.63 303.62 C618.07 299.08 612.53 295.60 606.30 291.78 C600.06 287.95 593.70 284.31 587.21 280.67 C580.73 277.02 574.14 273.53 567.38 269.90 C560.62 266.27 553.78 262.72 546.67 258.88 C539.57 255.04 532.37 250.99 524.75 246.85 C517.14 242.71 509.27 237.92 500.98 234.03 C492.69 230.15 483.91 226.13 475.02 223.54 C466.13 220.95 456.83 219.24 447.64 218.49 C438.45 217.74 429.08 218.05 419.90 219.02 C410.72 219.99 401.52 221.87 392.57 224.29 C383.62 226.72 374.75 229.90 366.18 233.56 C357.60 237.22 349.93 240.94 341.12 246.25 C332.32 251.57 322.06 258.39 313.34 265.46 C304.62 272.53 296.30 280.34 288.82 288.68 C281.33 297.01 274.40 306.05 268.44 315.46 C262.48 324.87 257.26 334.93 253.05 345.15 C248.83 355.37 245.57 366.13 243.13 376.78 C240.70 387.43 239.34 398.40 238.42 409.07 C237.50 419.74 237.48 430.42 237.60 440.80 C237.72 451.18 238.38 461.37 239.14 471.37 C239.89 481.38 240.91 491.15 242.12 500.83 C243.33 510.51 244.72 520.01 246.39 529.46 C248.05 538.91 249.91 548.24 252.12 557.51 C254.32 566.79 257.07 576.73 259.61 585.10 C262.15 593.47 264.57 600.23 267.35 607.75 C270.13 615.27 273.16 622.71 276.31 630.21 C279.46 637.71 282.82 645.15 286.25 652.77 C289.67 660.40 293.25 668.00 296.86 675.97 C300.48 683.93 304.08 692.10 307.96 700.57 C311.84 709.04 315.53 718.16 320.14 726.81 C324.75 735.45 329.68 744.48 335.61 752.43 C341.53 760.39 348.35 767.89 355.68 774.53 C363.02 781.18 371.18 787.08 379.60 792.32 C388.03 797.56 397.05 802.06 406.23 805.96 C415.40 809.86 424.99 813.08 434.66 815.71 C444.32 818.34 453.26 820.33 464.21 821.74 C475.17 823.16 488.35 824.29 500.40 824.20 C512.45 824.10 524.66 823.14 536.52 821.19 C548.39 819.25 560.26 816.35 571.60 812.52 C582.94 808.68 594.09 803.83 604.54 798.18 C615.00 792.54 625.02 785.82 634.32 778.64 C643.62 771.46 652.26 763.31 660.34 755.08 C668.43 746.86 675.79 738.02 682.83 729.29 C689.87 720.56 696.34 711.61 702.58 702.72 C708.83 693.82 714.69 684.90 720.31 675.91 C725.93 666.92 731.26 657.92 736.30 648.77 C741.34 639.62 746.12 630.41 750.55 621.02 C754.98 611.63 759.31 602.15 762.88 592.41 C766.44 582.67 769.71 572.68 771.94 562.56 C774.16 552.43 775.66 542.00 776.22 531.68 C776.78 521.35 776.52 510.83 775.31 500.62 C774.10 490.41 771.95 480.15 768.97 470.39 C765.99 460.62 761.97 451.03 757.42 442.02 C752.86 433.01 747.32 424.41 741.67 416.30 C736.01 408.19 729.69 400.63 723.48 393.35 C717.28 386.07 710.82 379.27 704.44 372.63 C698.06 366.00 691.66 359.70 685.20 353.54 C678.75 347.39 672.30 341.46 665.71 335.72 C659.11 329.97 652.47 324.40 645.62 319.05 C638.77 313.70 631.18 308.17 624.63 303.62 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M259.61 585.10 C262.15 593.47 264.57 600.23 267.35 607.75 C270.13 615.27 273.16 622.71 276.31 630.21 C279.46 637.71 282.82 645.15 286.25 652.77 C289.67 660.40 293.25 668.00 296.86 675.97 C300.48 683.93 304.08 692.10 307.96 700.57 C311.84 709.04 315.53 718.16 320.14 726.81 C324.75 735.45 329.68 744.48 335.61 752.43 C341.53 760.39 348.35 767.89 355.68 774.53 C363.02 781.18 371.18 787.08 379.60 792.32 C388.03 797.56 397.05 802.06 406.23 805.96 C415.40 809.86 424.99 813.08 434.66 815.71 C444.32 818.34 453.26 820.33 464.21 821.74 C475.17 823.16 488.35 824.29 500.40 824.20 C512.45 824.10 524.66 823.14 536.52 821.19 C548.39 819.25 560.26 816.35 571.60 812.52 C582.94 808.68 594.09 803.83 604.54 798.18 C615.00 792.54 625.02 785.82 634.32 778.64 C643.62 771.46 652.26 763.31 660.34 755.08 C668.43 746.86 675.79 738.02 682.83 729.29 C689.87 720.56 696.34 711.61 702.58 702.72 C708.83 693.82 714.69 684.90 720.31 675.91 C725.93 666.92 731.26 657.92 736.30 648.77 C741.34 639.62 746.12 630.41 750.55 621.02 C754.98 611.63 759.37 601.27 762.88 592.41 C766.38 583.55 768.88 576.14 771.56 567.83 C774.25 559.53 776.67 551.12 778.99 542.59 C781.32 534.06 783.40 525.44 785.52 516.63 C787.65 507.82 789.60 498.94 791.72 489.74 C793.84 480.55 796.06 471.21 798.23 461.46 C800.39 451.71 803.07 441.57 804.71 431.24 C806.34 420.92 807.93 410.09 808.04 399.50 C808.16 388.91 807.26 378.13 805.40 367.71 C803.54 357.28 800.53 346.94 796.88 336.97 C793.23 327.00 788.63 317.24 783.50 307.88 C778.37 298.53 772.46 289.46 766.11 280.83 C759.75 272.20 753.60 264.59 745.37 256.10 C737.15 247.62 726.89 237.89 716.76 229.92 C706.64 221.95 695.81 214.59 684.62 208.26 C673.42 201.93 661.60 196.38 649.60 191.94 C637.60 187.50 625.08 184.02 612.60 181.64 C600.12 179.26 587.29 178.02 574.73 177.66 C562.17 177.29 549.51 178.14 537.23 179.44 C524.95 180.75 512.84 183.03 501.06 185.48 C489.27 187.92 477.79 190.94 466.53 194.11 C455.26 197.28 444.29 200.75 433.47 204.48 C422.65 208.20 412.05 212.16 401.61 216.47 C391.17 220.77 380.91 225.33 370.83 230.30 C360.75 235.26 350.71 240.39 341.12 246.25 C331.54 252.11 322.06 258.39 313.34 265.46 C304.62 272.53 296.30 280.34 288.82 288.68 C281.33 297.01 274.40 306.05 268.44 315.46 C262.48 324.87 257.26 334.93 253.05 345.15 C248.83 355.37 245.57 366.13 243.13 376.78 C240.70 387.43 239.34 398.40 238.42 409.07 C237.50 419.74 237.48 430.42 237.60 440.80 C237.72 451.18 238.38 461.37 239.14 471.37 C239.89 481.38 240.91 491.15 242.12 500.83 C243.33 510.51 244.72 520.01 246.39 529.46 C248.05 538.91 249.91 548.24 252.12 557.51 C254.32 566.79 257.07 576.73 259.61 585.10 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M762.88 592.41 C766.38 583.55 768.88 576.14 771.56 567.83 C774.25 559.53 776.67 551.12 778.99 542.59 C781.32 534.06 783.40 525.44 785.52 516.63 C787.65 507.82 789.60 498.94 791.72 489.74 C793.84 480.55 796.06 471.21 798.23 461.46 C800.39 451.71 803.07 441.57 804.71 431.24 C806.34 420.92 807.93 410.09 808.04 399.50 C808.16 388.91 807.26 378.13 805.40 367.71 C803.54 357.28 800.53 346.94 796.88 336.97 C793.23 327.00 788.63 317.24 783.50 307.88 C778.37 298.53 772.46 289.46 766.11 280.83 C759.75 272.20 753.60 264.59 745.37 256.10 C737.15 247.62 726.89 237.89 716.76 229.92 C706.64 221.95 695.81 214.59 684.62 208.26 C673.42 201.93 661.60 196.38 649.60 191.94 C637.60 187.50 625.08 184.02 612.60 181.64 C600.12 179.26 587.29 178.02 574.73 177.66 C562.17 177.29 549.51 178.14 537.23 179.44 C524.95 180.75 512.84 183.03 501.06 185.48 C489.27 187.92 477.79 190.94 466.53 194.11 C455.26 197.28 444.29 200.75 433.47 204.48 C422.65 208.20 412.05 212.16 401.61 216.47 C391.17 220.77 380.91 225.33 370.83 230.30 C360.75 235.26 350.05 241.09 341.12 246.25 C332.20 251.42 325.10 256.04 317.26 261.30 C309.42 266.56 301.74 272.10 294.08 277.81 C286.41 283.52 278.91 289.48 271.28 295.58 C263.65 301.68 256.12 307.97 248.31 314.42 C240.49 320.87 232.52 327.38 224.37 334.26 C216.23 341.15 207.46 348.04 199.42 355.74 C191.39 363.43 183.11 371.57 176.15 380.44 C169.20 389.31 162.96 398.98 157.72 408.97 C152.49 418.96 148.23 429.61 144.74 440.38 C141.26 451.14 138.68 462.35 136.80 473.56 C134.92 484.77 133.86 496.25 133.46 507.66 C133.06 519.07 133.16 529.49 134.40 542.02 C135.63 554.55 137.74 569.45 140.88 582.82 C144.01 596.19 148.13 609.50 153.22 622.22 C158.31 634.94 164.42 647.44 171.41 659.15 C178.40 670.86 186.43 682.14 195.16 692.49 C203.88 702.85 213.64 712.51 223.77 721.29 C233.89 730.07 244.89 737.93 255.90 745.17 C266.91 752.40 278.43 758.74 289.83 764.70 C301.22 770.67 312.77 775.96 324.27 780.96 C335.76 785.96 347.25 790.50 358.78 794.70 C370.32 798.90 381.85 802.74 393.48 806.17 C405.12 809.60 416.79 812.67 428.58 815.27 C440.37 817.86 452.24 820.26 464.21 821.74 C476.18 823.23 488.35 824.29 500.40 824.20 C512.45 824.10 524.66 823.14 536.52 821.19 C548.39 819.25 560.26 816.35 571.60 812.52 C582.94 808.68 594.09 803.83 604.54 798.18 C615.00 792.54 625.02 785.82 634.32 778.64 C643.62 771.46 652.26 763.31 660.34 755.08 C668.43 746.86 675.79 738.02 682.83 729.29 C689.87 720.56 696.34 711.61 702.58 702.72 C708.83 693.82 714.69 684.90 720.31 675.91 C725.93 666.92 731.26 657.92 736.30 648.77 C741.34 639.62 746.12 630.41 750.55 621.02 C754.98 611.63 759.37 601.27 762.88 592.41 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M341.12 246.25 C332.20 251.42 325.10 256.04 317.26 261.30 C309.42 266.56 301.74 272.10 294.08 277.81 C286.41 283.52 278.91 289.48 271.28 295.58 C263.65 301.68 256.12 307.97 248.31 314.42 C240.49 320.87 232.52 327.38 224.37 334.26 C216.23 341.15 207.46 348.04 199.42 355.74 C191.39 363.43 183.11 371.57 176.15 380.44 C169.20 389.31 162.96 398.98 157.72 408.97 C152.49 418.96 148.23 429.61 144.74 440.38 C141.26 451.14 138.68 462.35 136.80 473.56 C134.92 484.77 133.86 496.25 133.46 507.66 C133.06 519.07 133.16 529.49 134.40 542.02 C135.63 554.55 137.74 569.45 140.88 582.82 C144.01 596.19 148.13 609.50 153.22 622.22 C158.31 634.94 164.42 647.44 171.41 659.15 C178.40 670.86 186.43 682.14 195.16 692.49 C203.88 702.85 213.64 712.51 223.77 721.29 C233.89 730.07 244.89 737.93 255.90 745.17 C266.91 752.40 278.43 758.74 289.83 764.70 C301.22 770.67 312.77 775.96 324.27 780.96 C335.76 785.96 347.25 790.50 358.78 794.70 C370.32 798.90 381.85 802.74 393.48 806.17 C405.12 809.60 416.79 812.67 428.58 815.27 C440.37 817.86 453.26 820.12 464.21 821.74 C475.17 823.37 484.23 824.21 494.32 825.00 C504.42 825.79 514.57 826.23 524.81 826.49 C535.05 826.75 545.32 826.68 555.76 826.55 C566.21 826.41 576.69 826.04 587.48 825.68 C598.27 825.33 609.22 824.99 620.52 824.43 C631.83 823.87 643.61 823.67 655.30 822.35 C667.00 821.03 679.15 819.48 690.70 816.51 C702.26 813.53 713.74 809.46 724.62 804.48 C735.50 799.51 746.01 793.37 755.97 786.65 C765.93 779.93 775.43 772.26 784.37 764.14 C793.31 756.02 801.74 747.15 809.60 737.92 C817.46 728.68 824.26 719.94 831.51 708.71 C838.77 697.47 846.88 683.69 853.13 670.52 C859.39 657.34 864.81 643.57 869.05 629.64 C873.29 615.72 876.56 601.31 878.59 586.97 C880.63 572.63 881.54 557.96 881.28 543.57 C881.02 529.19 879.48 514.70 877.05 500.67 C874.63 486.65 870.90 472.79 866.73 459.42 C862.55 446.05 857.36 433.07 852.00 420.48 C846.63 407.88 840.68 395.72 834.55 383.84 C828.42 371.96 821.95 360.46 815.20 349.21 C808.45 337.97 801.43 327.03 794.04 316.37 C786.65 305.72 778.96 295.34 770.84 285.29 C762.73 275.25 754.39 265.33 745.37 256.10 C736.36 246.87 726.89 237.89 716.76 229.92 C706.64 221.95 695.81 214.59 684.62 208.26 C673.42 201.93 661.60 196.38 649.60 191.94 C637.60 187.50 625.08 184.02 612.60 181.64 C600.12 179.26 587.29 178.02 574.73 177.66 C562.17 177.29 549.51 178.14 537.23 179.44 C524.95 180.75 512.84 183.03 501.06 185.48 C489.27 187.92 477.79 190.94 466.53 194.11 C455.26 197.28 444.29 200.75 433.47 204.48 C422.65 208.20 412.05 212.16 401.61 216.47 C391.17 220.77 380.91 225.33 370.83 230.30 C360.75 235.26 350.05 241.09 341.12 246.25 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M464.21 821.74 C475.17 823.37 484.23 824.21 494.32 825.00 C504.42 825.79 514.57 826.23 524.81 826.49 C535.05 826.75 545.32 826.68 555.76 826.55 C566.21 826.41 576.69 826.04 587.48 825.68 C598.27 825.33 609.22 824.99 620.52 824.43 C631.83 823.87 643.61 823.67 655.30 822.35 C667.00 821.03 679.15 819.48 690.70 816.51 C702.26 813.53 713.74 809.46 724.62 804.48 C735.50 799.51 746.01 793.37 755.97 786.65 C765.93 779.93 775.43 772.26 784.37 764.14 C793.31 756.02 801.74 747.15 809.60 737.92 C817.46 728.68 824.26 719.94 831.51 708.71 C838.77 697.47 846.88 683.69 853.13 670.52 C859.39 657.34 864.81 643.57 869.05 629.64 C873.29 615.72 876.56 601.31 878.59 586.97 C880.63 572.63 881.54 557.96 881.28 543.57 C881.02 529.19 879.48 514.70 877.05 500.67 C874.63 486.65 870.90 472.79 866.73 459.42 C862.55 446.05 857.36 433.07 852.00 420.48 C846.63 407.88 840.68 395.72 834.55 383.84 C828.42 371.96 821.95 360.46 815.20 349.21 C808.45 337.97 801.43 327.03 794.04 316.37 C786.65 305.72 778.96 295.34 770.84 285.29 C762.73 275.25 753.46 264.76 745.37 256.10 C737.29 247.44 730.27 240.71 722.34 233.34 C714.40 225.97 706.19 218.88 697.77 211.89 C689.36 204.90 680.70 198.17 671.84 191.42 C662.99 184.67 653.93 178.11 644.64 171.37 C635.35 164.63 625.93 157.81 616.08 150.97 C606.24 144.13 596.23 136.75 585.59 130.33 C574.95 123.90 563.80 117.40 552.23 112.43 C540.65 107.46 528.44 103.41 516.14 100.49 C503.84 97.57 491.10 95.82 478.43 94.91 C465.76 94.00 452.86 94.14 440.12 95.03 C427.38 95.91 414.57 97.73 401.99 100.23 C389.41 102.74 378.04 105.51 364.65 110.05 C351.25 114.59 335.46 120.67 321.60 127.47 C307.73 134.26 294.16 142.10 281.44 150.82 C268.72 159.54 256.48 169.30 245.26 179.80 C234.05 190.31 223.54 201.82 214.15 213.85 C204.76 225.88 196.34 238.84 188.94 252.00 C181.54 265.15 175.29 279.03 169.75 292.81 C164.21 306.59 159.75 320.73 155.70 334.69 C151.66 348.65 148.39 362.66 145.49 376.56 C142.59 390.47 140.23 404.31 138.30 418.14 C136.37 431.96 134.89 445.73 133.93 459.51 C132.97 473.28 132.46 487.04 132.54 500.79 C132.62 514.55 133.01 528.35 134.40 542.02 C135.79 555.69 137.74 569.45 140.88 582.82 C144.01 596.19 148.13 609.50 153.22 622.22 C158.31 634.94 164.42 647.44 171.41 659.15 C178.40 670.86 186.43 682.14 195.16 692.49 C203.88 702.85 213.64 712.51 223.77 721.29 C233.89 730.07 244.89 737.93 255.90 745.17 C266.91 752.40 278.43 758.74 289.83 764.70 C301.22 770.67 312.77 775.96 324.27 780.96 C335.76 785.96 347.25 790.50 358.78 794.70 C370.32 798.90 381.85 802.74 393.48 806.17 C405.12 809.60 416.79 812.67 428.58 815.27 C440.37 817.86 453.26 820.12 464.21 821.74 Z`

	__FrontCurveStack__000000_Front_Curve_Stack.Name = `Front Curve Stack`
	__FrontCurveStack__000000_Front_Curve_Stack.IsDisplayed = false
	__FrontCurveStack__000000_Front_Curve_Stack.Color = ``
	__FrontCurveStack__000000_Front_Curve_Stack.FillOpacity = 0.000000
	__FrontCurveStack__000000_Front_Curve_Stack.Stroke = `blue`
	__FrontCurveStack__000000_Front_Curve_Stack.StrokeOpacity = 1.000000
	__FrontCurveStack__000000_Front_Curve_Stack.StrokeWidth = 1.000000
	__FrontCurveStack__000000_Front_Curve_Stack.StrokeDashArray = ``
	__FrontCurveStack__000000_Front_Curve_Stack.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__000000_Front_Curve_Stack.Transform = ``

	__HorizontalAxis__000000_Horizontal_Axis.Name = `Horizontal Axis`
	__HorizontalAxis__000000_Horizontal_Axis.IsDisplayed = false
	__HorizontalAxis__000000_Horizontal_Axis.AxisHandleBorderLength = 0.000000
	__HorizontalAxis__000000_Horizontal_Axis.Axis_Length = 600.000000
	__HorizontalAxis__000000_Horizontal_Axis.Color = ``
	__HorizontalAxis__000000_Horizontal_Axis.FillOpacity = 0.000000
	__HorizontalAxis__000000_Horizontal_Axis.Stroke = ``
	__HorizontalAxis__000000_Horizontal_Axis.StrokeOpacity = 0.000000
	__HorizontalAxis__000000_Horizontal_Axis.StrokeWidth = 1.000000
	__HorizontalAxis__000000_Horizontal_Axis.StrokeDashArray = ``
	__HorizontalAxis__000000_Horizontal_Axis.StrokeDashArrayWhenSelected = ``
	__HorizontalAxis__000000_Horizontal_Axis.Transform = ``

	__Key__000000_F_key.Name = `F key`
	__Key__000000_F_key.IsDisplayed = true
	__Key__000000_F_key.Path = `M562 -21c0 89 -65 150 -155 150c7 -44 34 -203 55 -323c71 29 100 102 100 173zM420 -206l-58 329c-59 -14 -104 -63 -104 -124c0 -49 22 -75 61 -99c12 -8 22 -13 22 -22s-9 -13 -17 -13c-80 0 -135 96 -135 166c0 94 62 190 153 217c-7 41 -14 88 -23 142 c-15 -15 -31 -29 -48 -44c-88 -76 -174 -185 -174 -307c0 -151 122 -251 265 -251c19 0 38 2 58 6zM332 822c-8 -31 -11 -65 -11 -102c0 -42 5 -81 11 -121c69 68 146 146 146 250c0 69 -24 118 -39 118c-52 0 -98 -105 -107 -145zM122 -513c0 66 45 123 115 123 c75 0 116 -57 116 -111c0 -64 -47 -104 -94 -111c-3 -1 -5 -2 -5 -4c0 -1 2 -2 3 -3c2 0 23 -5 47 -5c101 0 154 55 154 159c0 53 -11 123 -30 219c-23 -4 -50 -7 -79 -7c-186 0 -349 147 -349 334c0 200 126 321 217 406c21 17 73 70 74 71c-17 112 -22 161 -22 215 c0 84 18 212 82 288c33 39 64 51 71 51c18 0 47 -35 71 -86c16 -36 44 -110 44 -201c0 -159 -73 -284 -179 -395c9 -56 19 -115 29 -175c146 0 253 -102 253 -253c0 -103 -73 -205 -171 -237c6 -39 12 -69 15 -89c10 -57 16 -102 16 -141c0 -63 -14 -129 -68 -167 c-36 -22 -77 -34 -124 -34c-135 0 -186 87 -186 153z`
	__Key__000000_F_key.Color = `black`
	__Key__000000_F_key.FillOpacity = 0.300000
	__Key__000000_F_key.Stroke = `black`
	__Key__000000_F_key.StrokeOpacity = 0.300000
	__Key__000000_F_key.StrokeWidth = 0.000000
	__Key__000000_F_key.StrokeDashArray = ``
	__Key__000000_F_key.StrokeDashArrayWhenSelected = ``
	__Key__000000_F_key.Transform = `scale(0.2,-0.2)`

	__MovingLine__000000_Cursor.Name = `Cursor`
	__MovingLine__000000_Cursor.IsDisplayed = true
	__MovingLine__000000_Cursor.AngleDegree = 90.000000
	__MovingLine__000000_Cursor.Length = 2000.000000
	__MovingLine__000000_Cursor.CenterX = 0.000000
	__MovingLine__000000_Cursor.CenterY = 186.660000
	__MovingLine__000000_Cursor.StartX = 0.000000
	__MovingLine__000000_Cursor.EndX = 204.000000
	__MovingLine__000000_Cursor.DurationSeconds = 1.000000
	__MovingLine__000000_Cursor.IsMoving = false
	__MovingLine__000000_Cursor.Color = ``
	__MovingLine__000000_Cursor.FillOpacity = 0.000000
	__MovingLine__000000_Cursor.Stroke = `lightgray`
	__MovingLine__000000_Cursor.StrokeOpacity = 0.700000
	__MovingLine__000000_Cursor.StrokeWidth = 6.000000
	__MovingLine__000000_Cursor.StrokeDashArray = ``
	__MovingLine__000000_Cursor.StrokeDashArrayWhenSelected = ``
	__MovingLine__000000_Cursor.Transform = ``

	__NoteInfo__000000_0.Name = `0`
	__NoteInfo__000000_0.IsKept = true

	__NoteInfo__000001_1.Name = `1`
	__NoteInfo__000001_1.IsKept = true

	__NoteInfo__000002_10.Name = `10`
	__NoteInfo__000002_10.IsKept = true

	__NoteInfo__000003_11.Name = `11`
	__NoteInfo__000003_11.IsKept = true

	__NoteInfo__000004_12.Name = `12`
	__NoteInfo__000004_12.IsKept = false

	__NoteInfo__000005_13.Name = `13`
	__NoteInfo__000005_13.IsKept = false

	__NoteInfo__000006_14.Name = `14`
	__NoteInfo__000006_14.IsKept = false

	__NoteInfo__000007_15.Name = `15`
	__NoteInfo__000007_15.IsKept = true

	__NoteInfo__000008_16.Name = `16`
	__NoteInfo__000008_16.IsKept = false

	__NoteInfo__000009_17.Name = `17`
	__NoteInfo__000009_17.IsKept = false

	__NoteInfo__000010_18.Name = `18`
	__NoteInfo__000010_18.IsKept = true

	__NoteInfo__000011_19.Name = `19`
	__NoteInfo__000011_19.IsKept = true

	__NoteInfo__000012_2.Name = `2`
	__NoteInfo__000012_2.IsKept = true

	__NoteInfo__000013_20.Name = `20`
	__NoteInfo__000013_20.IsKept = false

	__NoteInfo__000014_21.Name = `21`
	__NoteInfo__000014_21.IsKept = false

	__NoteInfo__000015_3.Name = `3`
	__NoteInfo__000015_3.IsKept = true

	__NoteInfo__000016_4.Name = `4`
	__NoteInfo__000016_4.IsKept = true

	__NoteInfo__000017_5.Name = `5`
	__NoteInfo__000017_5.IsKept = false

	__NoteInfo__000018_6.Name = `6`
	__NoteInfo__000018_6.IsKept = false

	__NoteInfo__000019_7.Name = `7`
	__NoteInfo__000019_7.IsKept = false

	__NoteInfo__000020_8.Name = `8`
	__NoteInfo__000020_8.IsKept = true

	__NoteInfo__000021_9.Name = `9`
	__NoteInfo__000021_9.IsKept = false

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.BackendColor = `#F1F1F1`
	__Parameter__000000_Reference.MinuteColor = `#5A5A5A`
	__Parameter__000000_Reference.HourColor = `#1E1E1E`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.InsideAngle = 112.000000
	__Parameter__000000_Reference.SideLength = 102.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.440000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.088000
	__Parameter__000000_Reference.NbPitchLines = 169
	__Parameter__000000_Reference.BeatLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbBeatLines = 300
	__Parameter__000000_Reference.NbBeatLinesPerCurve = 22
	__Parameter__000000_Reference.FirstVoiceShiftX = -0.070000
	__Parameter__000000_Reference.FirstVoiceShiftY = 1.830000
	__Parameter__000000_Reference.PitchDifference = 12
	__Parameter__000000_Reference.Speed = 1.000000
	__Parameter__000000_Reference.Level = 2.200000
	__Parameter__000000_Reference.IsMinor = true
	__Parameter__000000_Reference.OriginX = 40.000000
	__Parameter__000000_Reference.OriginY = 950.000000
	__Parameter__000000_Reference.SpiralOriginX = 500.000000
	__Parameter__000000_Reference.SpiralOriginY = 500.000000
	__Parameter__000000_Reference.OriginCrossWidth = 800.000000
	__Parameter__000000_Reference.SpiralRadiusRatio = 1.060000
	__Parameter__000000_Reference.ShowSpiralBezierConstruct = false
	__Parameter__000000_Reference.ShowInterpolationPoints = false
	__Parameter__000000_Reference.ActualBeatsTemporalShift = 8
	__Parameter__000000_Reference.IsPlaying = false

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 102.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.AngleDegree = -82.317083
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 112.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Color = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.FillOpacity = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Stroke = `green`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeOpacity = 0.800000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeWidth = 2.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeDashArray = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.StrokeDashArrayWhenSelected = ``
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Transform = ``

	__Rhombus__000001_Initial_Rhombus.Name = `Initial Rhombus`
	__Rhombus__000001_Initial_Rhombus.IsDisplayed = false
	__Rhombus__000001_Initial_Rhombus.CenterX = 0.000000
	__Rhombus__000001_Initial_Rhombus.CenterY = 0.000000
	__Rhombus__000001_Initial_Rhombus.SideLength = 102.000000
	__Rhombus__000001_Initial_Rhombus.AngleDegree = -79.106605
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 112.000000
	__Rhombus__000001_Initial_Rhombus.Color = ``
	__Rhombus__000001_Initial_Rhombus.FillOpacity = 0.000000
	__Rhombus__000001_Initial_Rhombus.Stroke = `green`
	__Rhombus__000001_Initial_Rhombus.StrokeOpacity = 0.900000
	__Rhombus__000001_Initial_Rhombus.StrokeWidth = 1.000000
	__Rhombus__000001_Initial_Rhombus.StrokeDashArray = ``
	__Rhombus__000001_Initial_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000001_Initial_Rhombus.Transform = ``

	__Rhombus__000002_Rotated_Next_Rhombus.Name = `Rotated Next Rhombus`
	__Rhombus__000002_Rotated_Next_Rhombus.IsDisplayed = false
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 167.605459
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 22.610262
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 102.000000
	__Rhombus__000002_Rotated_Next_Rhombus.AngleDegree = -82.317083
	__Rhombus__000002_Rotated_Next_Rhombus.InsideAngle = 112.000000
	__Rhombus__000002_Rotated_Next_Rhombus.Color = ``
	__Rhombus__000002_Rotated_Next_Rhombus.FillOpacity = 0.000000
	__Rhombus__000002_Rotated_Next_Rhombus.Stroke = `lavender`
	__Rhombus__000002_Rotated_Next_Rhombus.StrokeOpacity = 0.700000
	__Rhombus__000002_Rotated_Next_Rhombus.StrokeWidth = 6.000000
	__Rhombus__000002_Rotated_Next_Rhombus.StrokeDashArray = ``
	__Rhombus__000002_Rotated_Next_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000002_Rotated_Next_Rhombus.Transform = ``

	__Rhombus__000003_Rotated_Rhombus.Name = `Rotated Rhombus`
	__Rhombus__000003_Rotated_Rhombus.IsDisplayed = false
	__Rhombus__000003_Rotated_Rhombus.CenterX = 0.000000
	__Rhombus__000003_Rotated_Rhombus.CenterY = 0.000000
	__Rhombus__000003_Rotated_Rhombus.SideLength = 102.000000
	__Rhombus__000003_Rotated_Rhombus.AngleDegree = -82.317083
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 112.000000
	__Rhombus__000003_Rotated_Rhombus.Color = ``
	__Rhombus__000003_Rotated_Rhombus.FillOpacity = 0.000000
	__Rhombus__000003_Rotated_Rhombus.Stroke = `black`
	__Rhombus__000003_Rotated_Rhombus.StrokeOpacity = 1.000000
	__Rhombus__000003_Rotated_Rhombus.StrokeWidth = 2.000000
	__Rhombus__000003_Rotated_Rhombus.StrokeDashArray = ``
	__Rhombus__000003_Rotated_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000003_Rotated_Rhombus.Transform = ``

	__RhombusGrid__000000_Growing_Rhombus_Grid.Name = `Growing Rhombus Grid`
	__RhombusGrid__000000_Growing_Rhombus_Grid.IsDisplayed = false

	__RhombusGrid__000001_Initial_Rhombus_Grid.Name = `Initial Rhombus Grid`
	__RhombusGrid__000001_Initial_Rhombus_Grid.IsDisplayed = false

	__RhombusGrid__000002_Rotated_Rhombus_Grid.Name = `Rotated Rhombus Grid`
	__RhombusGrid__000002_Rotated_Rhombus_Grid.IsDisplayed = false

	__ShapeCategory__000000_0_Axes.Name = `0. Axes`
	__ShapeCategory__000000_0_Axes.IsExpanded = true

	__ShapeCategory__000001_1_Initial.Name = `1. Initial`
	__ShapeCategory__000001_1_Initial.IsExpanded = false

	__ShapeCategory__000002_2_Rotated.Name = `2. Rotated`
	__ShapeCategory__000002_2_Rotated.IsExpanded = false

	__ShapeCategory__000003_3_Growing.Name = `3. Growing`
	__ShapeCategory__000003_3_Growing.IsExpanded = false

	__ShapeCategory__000004_4_Construction.Name = `4. Construction`
	__ShapeCategory__000004_4_Construction.IsExpanded = false

	__ShapeCategory__000005_5_Vertical_growth.Name = `5. Vertical growth`
	__ShapeCategory__000005_5_Vertical_growth.IsExpanded = false

	__ShapeCategory__000006_6_Spiral_growth.Name = `6. Spiral growth`
	__ShapeCategory__000006_6_Spiral_growth.IsExpanded = false

	__ShapeCategory__000007_7_Spiral_Growth_Bezier.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = false

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = true

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = true

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 163.608536
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -18.450976
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 188.254769
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 23.155111
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 123.728624
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 196.815963
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 146.990042
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 151.588607
	__SpiralBezier__000000_Spiral_Bezier_Seed.Color = ``
	__SpiralBezier__000000_Spiral_Bezier_Seed.FillOpacity = 0.000000
	__SpiralBezier__000000_Spiral_Bezier_Seed.Stroke = `green`
	__SpiralBezier__000000_Spiral_Bezier_Seed.StrokeOpacity = 1.000000
	__SpiralBezier__000000_Spiral_Bezier_Seed.StrokeWidth = 1.000000
	__SpiralBezier__000000_Spiral_Bezier_Seed.StrokeDashArray = ``
	__SpiralBezier__000000_Spiral_Bezier_Seed.StrokeDashArrayWhenSelected = ``
	__SpiralBezier__000000_Spiral_Bezier_Seed.Transform = ``

	__SpiralBezierGrid__000000_Spiral_Bezier_Full_Grid.Name = `Spiral Bezier Full Grid`
	__SpiralBezierGrid__000000_Spiral_Bezier_Full_Grid.IsDisplayed = false

	__SpiralBezierGrid__000001_Spiral_Bezier_Grid.Name = `Spiral Bezier Grid`
	__SpiralBezierGrid__000001_Spiral_Bezier_Grid.IsDisplayed = false

	__SpiralCircle__000000_Construction_Circle_Spiral.Name = `Construction Circle Spiral`
	__SpiralCircle__000000_Construction_Circle_Spiral.IsDisplayed = false
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 108.120000
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterY = 0.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.HasBespokeRadius = true
	__SpiralCircle__000000_Construction_Circle_Spiral.BespopkeRadius = 14.130000
	__SpiralCircle__000000_Construction_Circle_Spiral.Pitch = 0
	__SpiralCircle__000000_Construction_Circle_Spiral.Color = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.FillOpacity = 0.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.Stroke = `blue`
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeOpacity = 0.800000
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeWidth = 1.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeDashArray = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.Transform = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.ShowName = false
	__SpiralCircle__000000_Construction_Circle_Spiral.Path = `M 704.930000 500.000000 A 14.130000 14.130000 0 1 0 676.670000 500.000000 A 14.130000 14.130000 0 1 0 704.930000 500.000000 Z`

	__SpiralCircleGrid__000000_Brute_Spiral_Bezier_Circle_Grid.Name = `Brute Spiral Bezier Circle Grid`
	__SpiralCircleGrid__000000_Brute_Spiral_Bezier_Circle_Grid.IsDisplayed = true

	__SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid.Name = `Construction Circle Spiral Full Grid`
	__SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid.IsDisplayed = false

	__SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid.Name = `Construction Circle Spiral Grid`
	__SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid.IsDisplayed = false

	__SpiralCircleGrid__000003_Spiral_Circle_Grid.Name = `Spiral Circle Grid`
	__SpiralCircleGrid__000003_Spiral_Circle_Grid.IsDisplayed = false

	__SpiralLine__000000_Spiral_Contruction_Inner_Line.Name = `Spiral Contruction Inner Line`
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.IsDisplayed = false
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 163.608536
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 108.120000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -18.450976
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndY = -0.000000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.Color = ``
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.FillOpacity = 0.000000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.Stroke = `blue`
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StrokeOpacity = 1.000000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StrokeWidth = 1.000000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StrokeDashArray = ``
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StrokeDashArrayWhenSelected = ``
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.Transform = ``

	__SpiralLine__000001_Spiral_Contruction_Outer_Line.Name = `Spiral Contruction Outer Line`
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.IsDisplayed = false
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 163.608536
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 215.616145
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -18.450976
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -49.258767
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.Color = ``
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.FillOpacity = 0.000000
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.Stroke = `green`
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StrokeOpacity = 1.000000
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StrokeWidth = 1.000000
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StrokeDashArray = ``
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StrokeDashArrayWhenSelected = ``
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.Transform = ``

	__SpiralLineGrid__000000_Spiral_Construction_Inner_Line_Grid_Spiral.Name = `Spiral Construction Inner Line Grid Spiral`
	__SpiralLineGrid__000000_Spiral_Construction_Inner_Line_Grid_Spiral.IsDisplayed = false

	__SpiralLineGrid__000001_Spiral_Construction_Outer_Line_Full_Grid_Spiral.Name = `Spiral Construction Outer Line Full Grid Spiral`
	__SpiralLineGrid__000001_Spiral_Construction_Outer_Line_Full_Grid_Spiral.IsDisplayed = false

	__SpiralLineGrid__000002_Spiral_Construction_Outer_Line_Grid_Spiral.Name = `Spiral Construction Outer Line Grid Spiral`
	__SpiralLineGrid__000002_Spiral_Construction_Outer_Line_Grid_Spiral.IsDisplayed = false

	__SpiralOrigin__000000_Spiral_Origin.Name = `Spiral Origin`
	__SpiralOrigin__000000_Spiral_Origin.IsDisplayed = false
	__SpiralOrigin__000000_Spiral_Origin.Color = ``
	__SpiralOrigin__000000_Spiral_Origin.FillOpacity = 0.000000
	__SpiralOrigin__000000_Spiral_Origin.Stroke = `black`
	__SpiralOrigin__000000_Spiral_Origin.StrokeOpacity = 1.000000
	__SpiralOrigin__000000_Spiral_Origin.StrokeWidth = 1.000000
	__SpiralOrigin__000000_Spiral_Origin.StrokeDashArray = ``
	__SpiralOrigin__000000_Spiral_Origin.StrokeDashArrayWhenSelected = ``
	__SpiralOrigin__000000_Spiral_Origin.Transform = ``

	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Name = `Reference Spiral Rhombus`
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.IsDisplayed = false
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 51.269347
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 5.781908
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 39.445895
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 112.722594
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 163.608536
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -18.450976
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 31.977769
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -91.381296
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Color = `green`
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.FillOpacity = 1.000000
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Stroke = `blue`
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.StrokeOpacity = 1.000000
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.StrokeWidth = 1.000000
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.StrokeDashArray = ``
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.StrokeDashArrayWhenSelected = ``
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Transform = ``

	__SpiralRhombusGrid__000000_Spiral_Rhombus_Grid.Name = `Spiral Rhombus Grid`
	__SpiralRhombusGrid__000000_Spiral_Rhombus_Grid.IsDisplayed = false

	__VerticalAxis__000000_Vertical_Axis.Name = `Vertical Axis`
	__VerticalAxis__000000_Vertical_Axis.IsDisplayed = false
	__VerticalAxis__000000_Vertical_Axis.AxisHandleBorderLength = 0.000000
	__VerticalAxis__000000_Vertical_Axis.Axis_Length = 600.000000
	__VerticalAxis__000000_Vertical_Axis.Color = ``
	__VerticalAxis__000000_Vertical_Axis.FillOpacity = 0.000000
	__VerticalAxis__000000_Vertical_Axis.Stroke = ``
	__VerticalAxis__000000_Vertical_Axis.StrokeOpacity = 0.000000
	__VerticalAxis__000000_Vertical_Axis.StrokeWidth = 1.000000
	__VerticalAxis__000000_Vertical_Axis.StrokeDashArray = ``
	__VerticalAxis__000000_Vertical_Axis.StrokeDashArrayWhenSelected = ``
	__VerticalAxis__000000_Vertical_Axis.Transform = ``

	// Setup of pointers
	__Axis__000001_Construction_Axis.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Axis__000002_Initial_Axis.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Axis__000004_Rotated_Axis.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__AxisGrid__000000_Beat_Lines.Reference = __Axis__000000_Beat_Reference
	__AxisGrid__000000_Beat_Lines.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	__AxisGrid__000001_Construction_Axis_Grid.Reference = __Axis__000001_Construction_Axis
	__AxisGrid__000001_Construction_Axis_Grid.ShapeCategory = __ShapeCategory__000004_4_Construction
	__AxisGrid__000002_Pitch_Lines.Reference = __Axis__000003_Pitch_Line
	__AxisGrid__000002_Pitch_Lines.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	__Bezier__000002_Growth_Bezier_Right_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__Bezier__000003_Growth_Curve_Next_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__Bezier__000005_Growth_Curve_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGrid__000000_2nb_Voice.Reference = __Bezier__000000_2nd_voice_seed
	__BezierGrid__000000_2nb_Voice.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000001_2nd_voice_shifted_right.Reference = __Bezier__000000_2nd_voice_seed
	__BezierGrid__000001_2nd_voice_shifted_right.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000002_First_Voice.Reference = __Bezier__000001_First_Voice_seed
	__BezierGrid__000002_First_Voice.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000003_First_Voice_Shift_Right.Reference = __Bezier__000001_First_Voice_seed
	__BezierGrid__000003_First_Voice_Shift_Right.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000004_Growth_Curve.Reference = __Bezier__000005_Growth_Curve_Seed
	__BezierGrid__000004_Growth_Curve.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGrid__000005_Growth_Curve_Next.Reference = __Bezier__000003_Growth_Curve_Next_Seed
	__BezierGrid__000005_Growth_Curve_Next.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.Reference = __Bezier__000004_Growth_Curve_Next_Shift_Right_Seed
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGrid__000007_Growth_Curve_Shift_Right.Reference = __Bezier__000002_Growth_Bezier_Right_Seed
	__BezierGrid__000007_Growth_Curve_Shift_Right.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGridStack__000000_The_GrowthCurveStack.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__Circle__000000_0.ShapeCategory = __ShapeCategory__000003_3_Growing
	__Circle__000001_Construction_Circle.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Circle__000004_Initial_Circle.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Circle__000005_Rotated_Next_Circle.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__CircleGrid__000000_Construction_Circle_Grid.Reference = __Circle__000001_Construction_Circle
	__CircleGrid__000000_Construction_Circle_Grid.ShapeCategory = __ShapeCategory__000004_4_Construction
	__CircleGrid__000001_First_Voice_note_shifted_right.Reference = __Circle__000002_First_voice_notes_seed
	__CircleGrid__000001_First_Voice_note_shifted_right.ShapeCategory = __ShapeCategory__000009_9_Composer
	__CircleGrid__000002_First_Voice_notes.Reference = __Circle__000002_First_voice_notes_seed
	__CircleGrid__000002_First_Voice_notes.ShapeCategory = __ShapeCategory__000009_9_Composer
	__CircleGrid__000003_Growing_Circle_Grid.Reference = __Circle__000000_0
	__CircleGrid__000003_Growing_Circle_Grid.ShapeCategory = __ShapeCategory__000003_3_Growing
	__CircleGrid__000004_Growing_Circle_Grid_Shifted_Left.ShapeCategory = __ShapeCategory__000003_3_Growing
	__CircleGrid__000005_Initial_Circle_Grid.Reference = __Circle__000004_Initial_Circle
	__CircleGrid__000005_Initial_Circle_Grid.ShapeCategory = __ShapeCategory__000001_1_Initial
	__CircleGrid__000006_Rotated_Circle_Grid.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__CircleGrid__000007_Second_Voice_Notes_Shift_Right.Reference = __Circle__000002_First_voice_notes_seed
	__CircleGrid__000007_Second_Voice_Notes_Shift_Right.ShapeCategory = __ShapeCategory__000009_9_Composer
	__CircleGrid__000008_Second_Voice_notes.Reference = __Circle__000002_First_voice_notes_seed
	__CircleGrid__000008_Second_Voice_notes.ShapeCategory = __ShapeCategory__000009_9_Composer
	__FrontCurveStack__000000_Front_Curve_Stack.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000000_Non_Rotated_0_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000001_Non_Rotated_1_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000002_Non_Rotated_2_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000003_Non_Rotated_3_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000004_Non_Rotated_4_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000005_Non_Rotated_5_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000006_Non_Rotated_6_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000007_Non_Rotated_7_)
	__HorizontalAxis__000000_Horizontal_Axis.ShapeCategory = __ShapeCategory__000000_0_Axes
	__Key__000000_F_key.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	__MovingLine__000000_Cursor.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000001_Initial_Rhombus
	__Parameter__000000_Reference.InitialCircle = __Circle__000004_Initial_Circle
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000001_Initial_Rhombus_Grid
	__Parameter__000000_Reference.InitialCircleGrid = __CircleGrid__000005_Initial_Circle_Grid
	__Parameter__000000_Reference.InitialAxis = __Axis__000002_Initial_Axis
	__Parameter__000000_Reference.RotatedAxis = __Axis__000004_Rotated_Axis
	__Parameter__000000_Reference.RotatedRhombus = __Rhombus__000003_Rotated_Rhombus
	__Parameter__000000_Reference.RotatedRhombusGrid = __RhombusGrid__000002_Rotated_Rhombus_Grid
	__Parameter__000000_Reference.RotatedCircleGrid = __CircleGrid__000006_Rotated_Circle_Grid
	__Parameter__000000_Reference.NextRhombus = __Rhombus__000002_Rotated_Next_Rhombus
	__Parameter__000000_Reference.NextCircle = __Circle__000005_Rotated_Next_Circle
	__Parameter__000000_Reference.GrowingRhombusGridSeed = __Rhombus__000000_Growing_Rhombus_Grid_Seed
	__Parameter__000000_Reference.GrowingRhombusGrid = __RhombusGrid__000000_Growing_Rhombus_Grid
	__Parameter__000000_Reference.GrowingCircleGridSeed = __Circle__000000_0
	__Parameter__000000_Reference.GrowingCircleGrid = __CircleGrid__000003_Growing_Circle_Grid
	__Parameter__000000_Reference.GrowingCircleGridLeftSeed = __Circle__000003_Growing_Seed_Left
	__Parameter__000000_Reference.GrowingCircleGridLeft = __CircleGrid__000004_Growing_Circle_Grid_Shifted_Left
	__Parameter__000000_Reference.ConstructionAxis = __Axis__000001_Construction_Axis
	__Parameter__000000_Reference.ConstructionAxisGrid = __AxisGrid__000001_Construction_Axis_Grid
	__Parameter__000000_Reference.ConstructionCircle = __Circle__000001_Construction_Circle
	__Parameter__000000_Reference.ConstructionCircleGrid = __CircleGrid__000000_Construction_Circle_Grid
	__Parameter__000000_Reference.GrowthCurveSeed = __Bezier__000005_Growth_Curve_Seed
	__Parameter__000000_Reference.GrowthCurve = __BezierGrid__000004_Growth_Curve
	__Parameter__000000_Reference.GrowthCurveShiftedRightSeed = __Bezier__000002_Growth_Bezier_Right_Seed
	__Parameter__000000_Reference.GrowthCurveShiftedRight = __BezierGrid__000007_Growth_Curve_Shift_Right
	__Parameter__000000_Reference.GrowthCurveNextSeed = __Bezier__000003_Growth_Curve_Next_Seed
	__Parameter__000000_Reference.GrowthCurveNext = __BezierGrid__000005_Growth_Curve_Next
	__Parameter__000000_Reference.GrowthCurveNextShiftedRightSeed = __Bezier__000004_Growth_Curve_Next_Shift_Right_Seed
	__Parameter__000000_Reference.GrowthCurveNextShiftedRight = __BezierGrid__000006_Growth_Curve_Next_Shift_Right
	__Parameter__000000_Reference.GrowthCurveStack = __BezierGridStack__000000_The_GrowthCurveStack
	__Parameter__000000_Reference.SpiralRhombusGridSeed = __SpiralRhombus__000000_Reference_Spiral_Rhombus
	__Parameter__000000_Reference.SpiralRhombusGrid = __SpiralRhombusGrid__000000_Spiral_Rhombus_Grid
	__Parameter__000000_Reference.SpiralCircleSeed = __SpiralCircle__000000_Construction_Circle_Spiral
	__Parameter__000000_Reference.SpiralCircleGrid = __SpiralCircleGrid__000003_Spiral_Circle_Grid
	__Parameter__000000_Reference.SpiralCircleFullGrid = __SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid
	__Parameter__000000_Reference.SpiralConstructionOuterLineSeed = __SpiralLine__000001_Spiral_Contruction_Outer_Line
	__Parameter__000000_Reference.SpiralConstructionInnerLineSeed = __SpiralLine__000000_Spiral_Contruction_Inner_Line
	__Parameter__000000_Reference.SpiralConstructionOuterLineGrid = __SpiralLineGrid__000002_Spiral_Construction_Outer_Line_Grid_Spiral
	__Parameter__000000_Reference.SpiralConstructionInnerLineGrid = __SpiralLineGrid__000000_Spiral_Construction_Inner_Line_Grid_Spiral
	__Parameter__000000_Reference.SpiralConstructionCircleGrid = __SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid
	__Parameter__000000_Reference.SpiralConstructionOuterLineFullGrid = __SpiralLineGrid__000001_Spiral_Construction_Outer_Line_Full_Grid_Spiral
	__Parameter__000000_Reference.SpiralBezierSeed = __SpiralBezier__000000_Spiral_Bezier_Seed
	__Parameter__000000_Reference.SpiralBezierGrid = __SpiralBezierGrid__000001_Spiral_Bezier_Grid
	__Parameter__000000_Reference.SpiralBezierFullGrid = __SpiralBezierGrid__000000_Spiral_Bezier_Full_Grid
	__Parameter__000000_Reference.FrontCurveStack = __FrontCurveStack__000000_Front_Curve_Stack
	__Parameter__000000_Reference.Fkey = __Key__000000_F_key
	__Parameter__000000_Reference.PitchLines = __AxisGrid__000002_Pitch_Lines
	__Parameter__000000_Reference.BeatLines = __AxisGrid__000000_Beat_Lines
	__Parameter__000000_Reference.FirstVoice = __BezierGrid__000002_First_Voice
	__Parameter__000000_Reference.FirstVoiceShiftRigth = __BezierGrid__000003_First_Voice_Shift_Right
	__Parameter__000000_Reference.SecondVoice = __BezierGrid__000000_2nb_Voice
	__Parameter__000000_Reference.SecondVoiceShiftedRight = __BezierGrid__000001_2nd_voice_shifted_right
	__Parameter__000000_Reference.FirstVoiceNotes = __CircleGrid__000002_First_Voice_notes
	__Parameter__000000_Reference.FirstVoiceNotesShiftedRight = __CircleGrid__000001_First_Voice_note_shifted_right
	__Parameter__000000_Reference.SecondVoiceNotes = __CircleGrid__000008_Second_Voice_notes
	__Parameter__000000_Reference.SecondVoiceNotesShiftedRight = __CircleGrid__000007_Second_Voice_Notes_Shift_Right
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000000_0)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000001_1)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000012_2)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000015_3)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000016_4)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000017_5)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000018_6)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000019_7)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000020_8)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000021_9)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000002_10)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000003_11)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000004_12)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000005_13)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000006_14)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000007_15)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000008_16)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000009_17)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000010_18)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000011_19)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000013_20)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000014_21)
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Horizontal_Axis
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Vertical_Axis
	__Parameter__000000_Reference.SpiralOrigin = __SpiralOrigin__000000_Spiral_Origin
	__Parameter__000000_Reference.Cursor = __MovingLine__000000_Cursor
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.ShapeCategory = __ShapeCategory__000003_3_Growing
	__Rhombus__000001_Initial_Rhombus.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Rhombus__000002_Rotated_Next_Rhombus.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__Rhombus__000003_Rotated_Rhombus.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__RhombusGrid__000000_Growing_Rhombus_Grid.Reference = __Rhombus__000003_Rotated_Rhombus
	__RhombusGrid__000000_Growing_Rhombus_Grid.ShapeCategory = __ShapeCategory__000003_3_Growing
	__RhombusGrid__000001_Initial_Rhombus_Grid.Reference = __Rhombus__000001_Initial_Rhombus
	__RhombusGrid__000001_Initial_Rhombus_Grid.ShapeCategory = __ShapeCategory__000001_1_Initial
	__RhombusGrid__000002_Rotated_Rhombus_Grid.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__SpiralBezier__000000_Spiral_Bezier_Seed.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralBezierGrid__000000_Spiral_Bezier_Full_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralBezierGrid__000001_Spiral_Bezier_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircle__000000_Construction_Circle_Spiral.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	__SpiralCircleGrid__000000_Brute_Spiral_Bezier_Circle_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircleGrid__000003_Spiral_Circle_Grid.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	__SpiralCircleGrid__000003_Spiral_Circle_Grid.SpiralRhombusGrid = __SpiralRhombusGrid__000000_Spiral_Rhombus_Grid
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralLineGrid__000000_Spiral_Construction_Inner_Line_Grid_Spiral.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralLineGrid__000001_Spiral_Construction_Outer_Line_Full_Grid_Spiral.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralLineGrid__000002_Spiral_Construction_Outer_Line_Grid_Spiral.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralOrigin__000000_Spiral_Origin.ShapeCategory = __ShapeCategory__000000_0_Axes
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	__SpiralRhombusGrid__000000_Spiral_Rhombus_Grid.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	__VerticalAxis__000000_Vertical_Axis.ShapeCategory = __ShapeCategory__000000_0_Axes
}
