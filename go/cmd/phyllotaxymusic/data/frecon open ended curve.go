package main

import (
	"time"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "2025-09-19 04:04:06.302150 CEST"
	const __write__utc_time__ = "2025-09-19 02:04:06.302150 UTC"

	const __commitId__ = "0000000342"

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
	__Bezier__000006_FirstVoiceShiftedRightSeed := (&models.Bezier{}).Stage(stage)

	__BezierGrid__000000_2nb_Voice := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000001_2nd_voice_shifted_right := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000002_First_Voice := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000003_First_Voice_Shift_Right := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000004_Growth_Curve := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000005_Growth_Curve_Next := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right := (&models.BezierGrid{}).Stage(stage)
	__BezierGrid__000007_Growth_Curve_Shift_Right := (&models.BezierGrid{}).Stage(stage)

	__BezierGridStack__000000_The_GrowthCurveStack := (&models.BezierGridStack{}).Stage(stage)

	__Chapter__000000_Deep_into_the_phyllotaxy_growth_curve := (&models.Chapter{}).Stage(stage)
	__Chapter__000001_Composing_your_own_phyllotaxy_music := (&models.Chapter{}).Stage(stage)

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

	__Content__000000_Phyllotaxy_Music := (&models.Content{}).Stage(stage)

	__ExportToMusicxml__000000_Singloton := (&models.ExportToMusicxml{}).Stage(stage)

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
	__Axis__000001_Construction_Axis.AngleDegree = 96.586776
	__Axis__000001_Construction_Axis.Length = 300.000000
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -34.412360
	__Axis__000001_Construction_Axis.EndY = 298.019780
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
	__Axis__000002_Initial_Axis.AngleDegree = 83.413224
	__Axis__000002_Initial_Axis.Length = 1307.669683
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
	__Axis__000004_Rotated_Axis.Length = 1307.669683
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -17.206180
	__Bezier__000005_Growth_Curve_Seed.StartY = 149.009890
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 119.882919
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 164.839576
	__Bezier__000005_Growth_Curve_Seed.EndX = 223.680341
	__Bezier__000005_Growth_Curve_Seed.EndY = 327.821758
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 86.591242
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 311.992073
	__Bezier__000005_Growth_Curve_Seed.Color = ``
	__Bezier__000005_Growth_Curve_Seed.FillOpacity = 0.000000
	__Bezier__000005_Growth_Curve_Seed.Stroke = `grey`
	__Bezier__000005_Growth_Curve_Seed.StrokeOpacity = 0.800000
	__Bezier__000005_Growth_Curve_Seed.StrokeWidth = 1.000000
	__Bezier__000005_Growth_Curve_Seed.StrokeDashArray = ``
	__Bezier__000005_Growth_Curve_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000005_Growth_Curve_Seed.Transform = ``

	__Bezier__000006_FirstVoiceShiftedRightSeed.Name = `FirstVoiceShiftedRightSeed`
	__Bezier__000006_FirstVoiceShiftedRightSeed.IsDisplayed = false
	__Bezier__000006_FirstVoiceShiftedRightSeed.StartX = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.StartY = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.ControlPointStartX = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.ControlPointStartY = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.EndX = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.EndY = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.ControlPointEndX = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.ControlPointEndY = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.Color = ``
	__Bezier__000006_FirstVoiceShiftedRightSeed.FillOpacity = 0.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.Stroke = `lightgreen`
	__Bezier__000006_FirstVoiceShiftedRightSeed.StrokeOpacity = 0.800000
	__Bezier__000006_FirstVoiceShiftedRightSeed.StrokeWidth = 6.000000
	__Bezier__000006_FirstVoiceShiftedRightSeed.StrokeDashArray = ``
	__Bezier__000006_FirstVoiceShiftedRightSeed.StrokeDashArrayWhenSelected = ``
	__Bezier__000006_FirstVoiceShiftedRightSeed.Transform = ``

	__BezierGrid__000000_2nb_Voice.Name = `2nb Voice`
	__BezierGrid__000000_2nb_Voice.IsDisplayed = false

	__BezierGrid__000001_2nd_voice_shifted_right.Name = `2nd voice shifted right`
	__BezierGrid__000001_2nd_voice_shifted_right.IsDisplayed = false

	__BezierGrid__000002_First_Voice.Name = `First Voice`
	__BezierGrid__000002_First_Voice.IsDisplayed = false

	__BezierGrid__000003_First_Voice_Shift_Right.Name = `First Voice Shift Right`
	__BezierGrid__000003_First_Voice_Shift_Right.IsDisplayed = false

	__BezierGrid__000004_Growth_Curve.Name = `Growth Curve`
	__BezierGrid__000004_Growth_Curve.IsDisplayed = true

	__BezierGrid__000005_Growth_Curve_Next.Name = `Growth Curve Next`
	__BezierGrid__000005_Growth_Curve_Next.IsDisplayed = false

	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.Name = `Growth Curve Next Shift Right`
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.IsDisplayed = false

	__BezierGrid__000007_Growth_Curve_Shift_Right.Name = `Growth Curve Shift Right`
	__BezierGrid__000007_Growth_Curve_Shift_Right.IsDisplayed = false

	__BezierGridStack__000000_The_GrowthCurveStack.Name = `The GrowthCurveStack`
	__BezierGridStack__000000_The_GrowthCurveStack.IsDisplayed = false

	__Chapter__000000_Deep_into_the_phyllotaxy_growth_curve.Name = `Deep into the phyllotaxy growth curve`
	__Chapter__000000_Deep_into_the_phyllotaxy_growth_curve.MardownContent = `To be completed`

	__Chapter__000001_Composing_your_own_phyllotaxy_music.Name = `Composing your own phyllotaxy music`
	__Chapter__000001_Composing_your_own_phyllotaxy_music.MardownContent = `To be completed`

	__Circle__000000_0.Name = `0`
	__Circle__000000_0.IsDisplayed = false
	__Circle__000000_0.CenterX = 0.000000
	__Circle__000000_0.CenterY = 0.000000
	__Circle__000000_0.HasBespokeRadius = false
	__Circle__000000_0.BespopkeRadius = 0.000000
	__Circle__000000_0.Color = ``
	__Circle__000000_0.FillOpacity = 0.000000
	__Circle__000000_0.Stroke = `blue`
	__Circle__000000_0.StrokeOpacity = 0.700000
	__Circle__000000_0.StrokeWidth = 2.000000
	__Circle__000000_0.StrokeDashArray = ``
	__Circle__000000_0.StrokeDashArrayWhenSelected = ``
	__Circle__000000_0.Transform = ``
	__Circle__000000_0.Pitch = 0
	__Circle__000000_0.ShowName = false
	__Circle__000000_0.BeatNb = 0

	__Circle__000001_Construction_Circle.Name = `Construction Circle`
	__Circle__000001_Construction_Circle.IsDisplayed = false
	__Circle__000001_Construction_Circle.CenterX = -17.206180
	__Circle__000001_Construction_Circle.CenterY = 149.009890
	__Circle__000001_Construction_Circle.HasBespokeRadius = true
	__Circle__000001_Construction_Circle.BespopkeRadius = 20.000000
	__Circle__000001_Construction_Circle.Color = ``
	__Circle__000001_Construction_Circle.FillOpacity = 0.000000
	__Circle__000001_Construction_Circle.Stroke = `blue`
	__Circle__000001_Construction_Circle.StrokeOpacity = 0.800000
	__Circle__000001_Construction_Circle.StrokeWidth = 2.000000
	__Circle__000001_Construction_Circle.StrokeDashArray = ``
	__Circle__000001_Construction_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000001_Construction_Circle.Transform = ``
	__Circle__000001_Construction_Circle.Pitch = 0
	__Circle__000001_Construction_Circle.ShowName = false
	__Circle__000001_Construction_Circle.BeatNb = 0

	__Circle__000002_First_voice_notes_seed.Name = `First voice notes seed`
	__Circle__000002_First_voice_notes_seed.IsDisplayed = false
	__Circle__000002_First_voice_notes_seed.CenterX = 0.000000
	__Circle__000002_First_voice_notes_seed.CenterY = 0.000000
	__Circle__000002_First_voice_notes_seed.HasBespokeRadius = true
	__Circle__000002_First_voice_notes_seed.BespopkeRadius = 10.000000
	__Circle__000002_First_voice_notes_seed.Color = ``
	__Circle__000002_First_voice_notes_seed.FillOpacity = 0.000000
	__Circle__000002_First_voice_notes_seed.Stroke = `grey`
	__Circle__000002_First_voice_notes_seed.StrokeOpacity = 0.800000
	__Circle__000002_First_voice_notes_seed.StrokeWidth = 3.000000
	__Circle__000002_First_voice_notes_seed.StrokeDashArray = ``
	__Circle__000002_First_voice_notes_seed.StrokeDashArrayWhenSelected = ``
	__Circle__000002_First_voice_notes_seed.Transform = ``
	__Circle__000002_First_voice_notes_seed.Pitch = 0
	__Circle__000002_First_voice_notes_seed.ShowName = false
	__Circle__000002_First_voice_notes_seed.BeatNb = 0

	__Circle__000003_Growing_Seed_Left.Name = `Growing Seed Left`
	__Circle__000003_Growing_Seed_Left.IsDisplayed = false
	__Circle__000003_Growing_Seed_Left.CenterX = 0.000000
	__Circle__000003_Growing_Seed_Left.CenterY = 0.000000
	__Circle__000003_Growing_Seed_Left.HasBespokeRadius = false
	__Circle__000003_Growing_Seed_Left.BespopkeRadius = 0.000000
	__Circle__000003_Growing_Seed_Left.Color = ``
	__Circle__000003_Growing_Seed_Left.FillOpacity = 0.000000
	__Circle__000003_Growing_Seed_Left.Stroke = `green`
	__Circle__000003_Growing_Seed_Left.StrokeOpacity = 0.800000
	__Circle__000003_Growing_Seed_Left.StrokeWidth = 1.800000
	__Circle__000003_Growing_Seed_Left.StrokeDashArray = ``
	__Circle__000003_Growing_Seed_Left.StrokeDashArrayWhenSelected = ``
	__Circle__000003_Growing_Seed_Left.Transform = ``
	__Circle__000003_Growing_Seed_Left.Pitch = 0
	__Circle__000003_Growing_Seed_Left.ShowName = false
	__Circle__000003_Growing_Seed_Left.BeatNb = 0

	__Circle__000004_Initial_Circle.Name = `Initial Circle`
	__Circle__000004_Initial_Circle.IsDisplayed = false
	__Circle__000004_Initial_Circle.CenterX = 0.000000
	__Circle__000004_Initial_Circle.CenterY = 0.000000
	__Circle__000004_Initial_Circle.HasBespokeRadius = false
	__Circle__000004_Initial_Circle.BespopkeRadius = 0.000000
	__Circle__000004_Initial_Circle.Color = ``
	__Circle__000004_Initial_Circle.FillOpacity = 0.000000
	__Circle__000004_Initial_Circle.Stroke = `lightblue`
	__Circle__000004_Initial_Circle.StrokeOpacity = 0.800000
	__Circle__000004_Initial_Circle.StrokeWidth = 3.000000
	__Circle__000004_Initial_Circle.StrokeDashArray = ``
	__Circle__000004_Initial_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000004_Initial_Circle.Transform = ``
	__Circle__000004_Initial_Circle.Pitch = 0
	__Circle__000004_Initial_Circle.ShowName = false
	__Circle__000004_Initial_Circle.BeatNb = 0

	__Circle__000005_Rotated_Next_Circle.Name = `Rotated Next Circle`
	__Circle__000005_Rotated_Next_Circle.IsDisplayed = false
	__Circle__000005_Rotated_Next_Circle.CenterX = 516.185401
	__Circle__000005_Rotated_Next_Circle.CenterY = 59.603956
	__Circle__000005_Rotated_Next_Circle.HasBespokeRadius = false
	__Circle__000005_Rotated_Next_Circle.BespopkeRadius = 0.000000
	__Circle__000005_Rotated_Next_Circle.Color = ``
	__Circle__000005_Rotated_Next_Circle.FillOpacity = 0.000000
	__Circle__000005_Rotated_Next_Circle.Stroke = `yellow`
	__Circle__000005_Rotated_Next_Circle.StrokeOpacity = 0.800000
	__Circle__000005_Rotated_Next_Circle.StrokeWidth = 3.000000
	__Circle__000005_Rotated_Next_Circle.StrokeDashArray = ``
	__Circle__000005_Rotated_Next_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000005_Rotated_Next_Circle.Transform = ``
	__Circle__000005_Rotated_Next_Circle.Pitch = 0
	__Circle__000005_Rotated_Next_Circle.ShowName = false
	__Circle__000005_Rotated_Next_Circle.BeatNb = 0

	__CircleGrid__000000_Construction_Circle_Grid.Name = `Construction Circle Grid`
	__CircleGrid__000000_Construction_Circle_Grid.IsDisplayed = false

	__CircleGrid__000001_First_Voice_note_shifted_right.Name = `First Voice note shifted right`
	__CircleGrid__000001_First_Voice_note_shifted_right.IsDisplayed = false

	__CircleGrid__000002_First_Voice_notes.Name = `First Voice notes`
	__CircleGrid__000002_First_Voice_notes.IsDisplayed = false

	__CircleGrid__000003_Growing_Circle_Grid.Name = `Growing Circle Grid`
	__CircleGrid__000003_Growing_Circle_Grid.IsDisplayed = false

	__CircleGrid__000004_Growing_Circle_Grid_Shifted_Left.Name = `Growing Circle Grid Shifted Left`
	__CircleGrid__000004_Growing_Circle_Grid_Shifted_Left.IsDisplayed = false

	__CircleGrid__000005_Initial_Circle_Grid.Name = `Initial Circle Grid`
	__CircleGrid__000005_Initial_Circle_Grid.IsDisplayed = false

	__CircleGrid__000006_Rotated_Circle_Grid.Name = `Rotated Circle Grid`
	__CircleGrid__000006_Rotated_Circle_Grid.IsDisplayed = false

	__CircleGrid__000007_Second_Voice_Notes_Shift_Right.Name = `Second Voice Notes Shift Right`
	__CircleGrid__000007_Second_Voice_Notes_Shift_Right.IsDisplayed = false

	__CircleGrid__000008_Second_Voice_notes.Name = `Second Voice notes`
	__CircleGrid__000008_Second_Voice_notes.IsDisplayed = false

	__Content__000000_Phyllotaxy_Music.Name = `Phyllotaxy Music`
	__Content__000000_Phyllotaxy_Music.MardownContent = `This site describes phyllotaxy generated music.

Phyllotaxy and music are two concepts far appart but they can be related.

Precisely, phyllotaxy can help to generates a music
theme, a suite of music notes, a melodic material.

### Cannon

A well known music theme is at the start of Bach's 2nd 
fugue in C minor. It starts with the C note.

<img src="./images/bach2ndFugue.png" style="height: 100px; width: auto;">

In this piece, the theme is repeated with two voices that overlaps.

<img src="./images/bach2ndFugue_large.png" style="height: 400px; width: auto;">

The first voice 
starts with a C on the first measure. On the third measure, the theme starts again, but it starts with a G. 
The G of the second voice is seven half tone higher, or a fifth in music parlance, above the
 C of the first voice. A fifth is an a well suited gap between notes to sounds harmoniously. 
The next note, a F# on the second voice, is played along a A on the first voice, also a
third. A third,  And so one.

It is the composer's chore
to have the theme well suited for supporting overlapping itself.

Having a theme that repeats itself harmoniously with a shift in beats and tone is called a cannon:

- Pachelbel's Canon in D
- "Row, Row, Row Your Boat"

Phyllotaxy will help us to compose theme that are suited for cannon with the following idea : 
if a cannon theme is a curve on a plan, it must overlap with itself with an horizontal and a vertical shift. 

### Phyllotaxy

Phyllotaxy means "shape of leaves" 
("phyllo" is greek for leave and "taxy" is greek for shape). 
It is the 
science developed by botanist, 
mathematicians and physicists 
to understand why plants
have the shape they have.

In the office of Stephane Douady,
a "spiral plants" specialist I 
visited, the shelves are full of 
beautiful 
dried specimen that people 
sent to him
from all over the world. 
Each plant has spectacular
spiral arragenements of leaves
that run upward the trunks.

A common spiral plant is the pinecone.

Stephane taught me that when 
the trunk grows, 
the cells
at the stem divide themselves. New cells
 end up building
the trunk but regularly, 
one of them differentiates and
becomes a leave. 
Because this happens on a regular tempo, 
the leaves arrange themselves along 
the trunk
in spiral patterns. 
Stephane pointed out that if you draw
a curve around the trunk to isolate 
one leave cell from
newer leave cells, you end up drawing 
the "front curve"
of the plant growth.

<img src="./images/growthCurveOnPineCone.png" style="height: 300px; width: auto;">

Stephane then compared a new leave cell 
to the last one.
Because leave cells appear regularly, 
a new cell is 
always located with the same distance
 above and the
same angle sideway, 
a rotation near 168 degrees, 
related to the golden ratio. 
Therefore, the "front curve" has 
this interesting geometric feature : 
but on the new leave cell,
it overlaps itself
after an upward translation 
and a rotation. Let's demonstrate this.

Use a grey pen and draw the unfolded the front curve on a piece 
of paper.

<img src="./images/fistVoiceSVGimage.svg" style="height: 100px; width: auto;">

Use a green pen and draw the same curve at the right of the first.

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRight.svg" style="height: 100px; width: auto;">

Draw it again
on a tracing paper with a red pen and put the tracing 
paper above and
shift it a bit to the right and 
a bit to the top.

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightAndSecondVoice.svg" style="height: 125px; width: auto;">

We did it ! The overlapping 
is perfect 
except for one space that have
the shape of an eye (
the place for the leave).

### Generating a theme form the front curve 

Now take take the first drawing and add a canevas
of vertical measure lines and horizontal pitch lines.

For the pitch lines, the spacing between pitch line
follows the scale of your choice. If you
choose a C minor scale, the tone differences between
notes are 1, 1/2, 1, 1, 1/2, 1 and 3/2.

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightWithGrid.svg" style="height: 300px; width: auto;">

Now, along the front curve, when it intersect the 
measure lines, 
draw some notes of your choosing.


<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightWithGridWithNotes.svg" style="height: 300px; width: auto;">

With the second voice note, this gives

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightWithSecondVoiceWithGridWithNotes.svg" style="height: 300px; width: auto;">

<p>Download the generated music score in the musescore musicxml format:
    <a href="./scores/score.musicxml" download="score.musicxml">Download MusicXML Score</a>
 </p>`
	__Content__000000_Phyllotaxy_Music.ContentPath = `content`
	__Content__000000_Phyllotaxy_Music.OutputPath = `../../../docs`
	__Content__000000_Phyllotaxy_Music.LayoutPath = `../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/layouts`
	__Content__000000_Phyllotaxy_Music.StaticPath = `../../../static`
	__Content__000000_Phyllotaxy_Music.Target = models.FILE

	__ExportToMusicxml__000000_Singloton.Name = `Singloton`

	__FrontCurve__000000_Non_Rotated_0_.Name = `Non Rotated 0 `
	__FrontCurve__000000_Non_Rotated_0_.Path = `M966.29 539.05 C968.50 523.07 969.92 509.15 971.36 493.97 C972.80 478.79 973.84 463.55 974.92 447.95 C976.00 432.36 976.79 416.69 977.83 400.38 C978.88 384.08 979.85 367.61 981.20 350.12 C982.55 332.64 984.41 314.60 985.93 295.48 C987.45 276.36 990.08 255.87 990.33 235.41 C990.58 214.95 990.55 193.19 987.42 172.71 C984.30 152.24 978.93 131.72 971.58 112.56 C964.23 93.39 954.38 74.96 943.33 57.73 C932.28 40.49 919.21 24.23 905.27 9.13 C891.33 -5.97 875.84 -20.01 859.69 -32.87 C843.54 -45.74 827.58 -56.96 808.37 -68.07 C789.17 -79.18 766.45 -90.81 744.44 -99.53 C722.44 -108.25 699.45 -115.44 676.36 -120.39 C653.28 -125.34 629.43 -128.47 605.94 -129.21 C582.45 -129.95 558.51 -128.51 535.43 -124.83 C512.36 -121.16 489.31 -114.89 467.49 -107.16 C445.67 -99.42 424.51 -89.04 404.49 -78.41 C384.47 -67.79 365.58 -55.35 347.37 -43.40 C329.15 -31.45 312.02 -18.91 295.20 -6.70 C278.39 5.50 262.32 17.63 246.48 29.83 C230.64 42.03 215.27 54.09 200.15 66.50 C185.02 78.91 170.21 91.34 155.72 104.28 C141.22 117.23 126.20 131.31 113.18 144.12 C100.15 156.92 89.29 168.46 77.61 181.11 C65.93 193.76 54.61 206.74 43.13 220.02 C31.65 233.30 20.44 246.87 8.72 260.81 C-2.99 274.75 -14.66 288.91 -27.16 303.66 C-39.65 318.40 -52.81 333.32 -66.26 349.30 C-79.72 365.29 -94.80 381.71 -107.88 399.57 C-120.96 417.44 -134.53 436.51 -144.76 456.48 C-155.00 476.45 -163.39 497.85 -169.29 519.39 C-175.20 540.93 -178.58 563.43 -180.19 585.70 C-181.80 607.98 -181.12 630.73 -178.96 653.06 C-176.79 675.40 -172.65 697.84 -167.19 719.72 C-161.73 741.60 -155.39 761.92 -146.19 784.35 C-137.00 806.77 -125.15 831.99 -112.02 854.27 C-98.90 876.55 -83.94 898.18 -67.44 918.01 C-50.94 937.85 -32.62 956.66 -13.03 973.27 C6.55 989.88 27.93 1004.98 50.07 1017.68 C72.21 1030.39 96.00 1040.95 119.79 1049.50 C143.58 1058.05 168.54 1064.01 192.80 1068.97 C217.06 1073.93 241.61 1076.67 265.35 1079.27 C289.10 1081.87 312.41 1083.24 335.27 1084.56 C358.13 1085.87 380.39 1086.68 402.52 1087.18 C424.65 1087.67 446.34 1087.88 468.03 1087.53 C489.72 1087.19 511.18 1086.50 532.67 1085.10 C554.16 1083.69 575.69 1082.27 596.97 1079.10 C618.25 1075.92 639.73 1072.09 660.35 1066.03 C680.96 1059.98 701.45 1052.23 720.68 1042.76 C739.92 1033.29 758.68 1022.03 775.75 1009.21 C792.83 996.39 808.97 981.69 823.13 965.85 C837.30 950.01 849.91 932.25 860.75 914.17 C871.59 896.09 880.35 876.43 888.16 857.36 C895.98 838.29 901.87 818.57 907.64 799.72 C913.41 780.88 918.06 762.31 922.77 744.27 C927.49 726.22 931.75 708.77 935.92 691.45 C940.09 674.14 944.08 657.29 947.77 640.35 C951.47 623.42 955.01 606.72 958.09 589.84 C961.18 572.96 964.08 555.03 966.29 539.05 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M113.18 144.12 C100.15 156.92 89.29 168.46 77.61 181.11 C65.93 193.76 54.61 206.74 43.13 220.02 C31.65 233.30 20.44 246.87 8.72 260.81 C-2.99 274.75 -14.66 288.91 -27.16 303.66 C-39.65 318.40 -52.81 333.32 -66.26 349.30 C-79.72 365.29 -94.80 381.71 -107.88 399.57 C-120.96 417.44 -134.53 436.51 -144.76 456.48 C-155.00 476.45 -163.39 497.85 -169.29 519.39 C-175.20 540.93 -178.58 563.43 -180.19 585.70 C-181.80 607.98 -181.12 630.73 -178.96 653.06 C-176.79 675.40 -172.65 697.84 -167.19 719.72 C-161.73 741.60 -155.39 761.92 -146.19 784.35 C-137.00 806.77 -125.15 831.99 -112.02 854.27 C-98.90 876.55 -83.94 898.18 -67.44 918.01 C-50.94 937.85 -32.62 956.66 -13.03 973.27 C6.55 989.88 27.93 1004.98 50.07 1017.68 C72.21 1030.39 96.00 1040.95 119.79 1049.50 C143.58 1058.05 168.54 1064.01 192.80 1068.97 C217.06 1073.93 241.61 1076.67 265.35 1079.27 C289.10 1081.87 312.41 1083.24 335.27 1084.56 C358.13 1085.87 380.39 1086.68 402.52 1087.18 C424.65 1087.67 446.34 1087.88 468.03 1087.53 C489.72 1087.19 511.18 1086.50 532.67 1085.10 C554.16 1083.69 576.89 1081.53 596.97 1079.10 C617.05 1076.66 634.38 1073.81 653.15 1070.48 C671.92 1067.15 690.61 1063.24 709.58 1059.10 C728.55 1054.97 747.46 1050.35 766.96 1045.67 C786.46 1040.99 806.03 1036.02 826.56 1031.02 C847.10 1026.02 868.21 1021.21 890.17 1015.68 C912.12 1010.14 935.58 1005.18 958.28 997.79 C980.98 990.41 1004.76 982.29 1026.39 971.38 C1048.01 960.46 1069.00 947.22 1088.03 932.30 C1107.07 917.37 1124.63 900.09 1140.58 881.84 C1156.53 863.60 1170.89 843.53 1183.73 822.84 C1196.57 802.15 1207.84 780.12 1217.64 757.70 C1227.43 735.27 1235.47 713.63 1242.49 688.29 C1249.52 662.95 1256.15 633.48 1259.80 605.67 C1263.44 577.86 1265.13 549.40 1264.38 521.43 C1263.62 493.45 1260.67 465.15 1255.26 437.82 C1249.85 410.50 1241.97 383.27 1231.91 357.50 C1221.85 331.73 1209.04 306.64 1194.89 283.19 C1180.74 259.74 1163.95 237.64 1146.99 216.82 C1130.04 196.00 1111.38 176.78 1093.17 158.25 C1074.95 139.72 1056.21 122.47 1037.71 105.62 C1019.20 88.77 1000.74 72.75 982.12 57.15 C963.50 41.55 944.95 26.52 925.99 12.02 C907.03 -2.47 887.95 -16.48 868.35 -29.83 C848.75 -43.18 829.02 -56.45 808.37 -68.07 C787.72 -79.68 766.45 -90.81 744.44 -99.53 C722.44 -108.25 699.45 -115.44 676.36 -120.39 C653.28 -125.34 629.43 -128.47 605.94 -129.21 C582.45 -129.95 558.51 -128.51 535.43 -124.83 C512.36 -121.16 489.31 -114.89 467.49 -107.16 C445.67 -99.42 424.51 -89.04 404.49 -78.41 C384.47 -67.79 365.58 -55.35 347.37 -43.40 C329.15 -31.45 312.02 -18.91 295.20 -6.70 C278.39 5.50 262.32 17.63 246.48 29.83 C230.64 42.03 215.27 54.09 200.15 66.50 C185.02 78.91 170.21 91.34 155.72 104.28 C141.22 117.23 126.20 131.31 113.18 144.12 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M596.97 1079.10 C617.05 1076.66 634.38 1073.81 653.15 1070.48 C671.92 1067.15 690.61 1063.24 709.58 1059.10 C728.55 1054.97 747.46 1050.35 766.96 1045.67 C786.46 1040.99 806.03 1036.02 826.56 1031.02 C847.10 1026.02 868.21 1021.21 890.17 1015.68 C912.12 1010.14 935.58 1005.18 958.28 997.79 C980.98 990.41 1004.76 982.29 1026.39 971.38 C1048.01 960.46 1069.00 947.22 1088.03 932.30 C1107.07 917.37 1124.63 900.09 1140.58 881.84 C1156.53 863.60 1170.89 843.53 1183.73 822.84 C1196.57 802.15 1207.84 780.12 1217.64 757.70 C1227.43 735.27 1235.47 713.63 1242.49 688.29 C1249.52 662.95 1256.15 633.48 1259.80 605.67 C1263.44 577.86 1265.13 549.40 1264.38 521.43 C1263.62 493.45 1260.67 465.15 1255.26 437.82 C1249.85 410.50 1241.97 383.27 1231.91 357.50 C1221.85 331.73 1209.04 306.64 1194.89 283.19 C1180.74 259.74 1163.95 237.64 1146.99 216.82 C1130.04 196.00 1111.38 176.78 1093.17 158.25 C1074.95 139.72 1056.21 122.47 1037.71 105.62 C1019.20 88.77 1000.74 72.75 982.12 57.15 C963.50 41.55 944.95 26.52 925.99 12.02 C907.03 -2.47 887.95 -16.48 868.35 -29.83 C848.75 -43.18 827.50 -56.66 808.37 -68.07 C789.25 -79.48 772.21 -88.65 753.60 -98.30 C735.00 -107.96 716.05 -117.03 696.73 -126.00 C677.40 -134.98 657.79 -143.45 637.66 -152.14 C617.54 -160.84 597.15 -169.26 575.98 -178.17 C554.81 -187.07 533.28 -196.39 510.64 -205.56 C488.01 -214.73 464.53 -225.19 440.17 -233.18 C415.81 -241.18 390.14 -249.17 364.49 -253.53 C338.82 -257.89 312.24 -259.94 286.20 -259.37 C260.16 -258.80 233.84 -255.37 208.25 -250.11 C182.65 -244.85 157.21 -237.08 132.62 -227.81 C108.03 -218.55 83.88 -207.17 60.70 -194.50 C37.50 -181.83 16.30 -168.69 -6.54 -151.75 C-29.41 -134.79 -54.73 -114.02 -76.44 -92.66 C-98.12 -71.26 -118.52 -48.01 -136.54 -23.59 C-154.51 0.80 -170.85 26.86 -184.51 53.69 C-198.17 80.49 -209.73 108.80 -218.55 137.33 C-227.37 165.85 -233.50 195.57 -237.45 224.85 C-241.40 254.12 -242.33 284.00 -242.23 312.99 C-242.13 341.97 -239.59 370.76 -236.85 398.75 C-234.12 426.74 -230.07 454.06 -225.84 480.91 C-221.61 507.76 -216.79 533.94 -211.47 559.86 C-206.15 585.78 -200.41 611.22 -193.90 636.44 C-187.39 661.67 -180.37 686.56 -172.42 711.21 C-164.47 735.87 -156.26 760.51 -146.19 784.35 C-136.13 808.19 -125.15 831.99 -112.02 854.27 C-98.90 876.55 -83.94 898.18 -67.44 918.01 C-50.94 937.85 -32.62 956.66 -13.03 973.27 C6.55 989.88 27.93 1004.98 50.07 1017.68 C72.21 1030.39 96.00 1040.95 119.79 1049.50 C143.58 1058.05 168.54 1064.01 192.80 1068.97 C217.06 1073.93 241.61 1076.67 265.35 1079.27 C289.10 1081.87 312.41 1083.24 335.27 1084.56 C358.13 1085.87 380.39 1086.68 402.52 1087.18 C424.65 1087.67 446.34 1087.88 468.03 1087.53 C489.72 1087.19 511.18 1086.50 532.67 1085.10 C554.16 1083.69 576.89 1081.53 596.97 1079.10 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M808.37 -68.07 C789.25 -79.48 772.21 -88.65 753.60 -98.30 C735.00 -107.96 716.05 -117.03 696.73 -126.00 C677.40 -134.98 657.79 -143.45 637.66 -152.14 C617.54 -160.84 597.15 -169.26 575.98 -178.17 C554.81 -187.07 533.28 -196.39 510.64 -205.56 C488.01 -214.73 464.53 -225.19 440.17 -233.18 C415.81 -241.18 390.14 -249.17 364.49 -253.53 C338.82 -257.89 312.24 -259.94 286.20 -259.37 C260.16 -258.80 233.84 -255.37 208.25 -250.11 C182.65 -244.85 157.21 -237.08 132.62 -227.81 C108.03 -218.55 83.88 -207.17 60.70 -194.50 C37.50 -181.83 16.30 -168.69 -6.54 -151.75 C-29.41 -134.79 -54.73 -114.02 -76.44 -92.66 C-98.12 -71.26 -118.52 -48.01 -136.54 -23.59 C-154.51 0.80 -170.85 26.86 -184.51 53.69 C-198.17 80.49 -209.73 108.80 -218.55 137.33 C-227.37 165.85 -233.50 195.57 -237.45 224.85 C-241.40 254.12 -242.33 284.00 -242.23 312.99 C-242.13 341.97 -239.59 370.76 -236.85 398.75 C-234.12 426.74 -230.07 454.06 -225.84 480.91 C-221.61 507.76 -216.79 533.94 -211.47 559.86 C-206.15 585.78 -200.41 611.22 -193.90 636.44 C-187.39 661.67 -180.37 686.56 -172.42 711.21 C-164.47 735.87 -155.11 761.72 -146.19 784.35 C-137.28 806.97 -128.63 826.25 -118.92 846.96 C-109.21 867.67 -98.78 888.08 -87.93 908.59 C-77.09 929.10 -65.62 949.35 -53.82 970.01 C-42.02 990.68 -29.77 1011.24 -17.14 1032.58 C-4.51 1053.91 8.22 1075.75 21.96 1098.05 C35.71 1120.35 49.39 1144.12 65.32 1166.40 C81.26 1188.68 98.28 1211.68 117.56 1231.73 C136.84 1251.78 158.33 1270.49 181.01 1286.70 C203.69 1302.92 228.40 1316.97 253.67 1329.03 C278.93 1341.10 305.65 1351.00 332.62 1359.09 C359.59 1367.19 387.52 1373.25 415.47 1377.60 C443.42 1381.95 469.95 1384.63 500.34 1385.18 C530.73 1385.73 565.59 1384.69 597.79 1380.91 C629.99 1377.13 662.34 1371.04 693.54 1362.47 C724.74 1353.90 755.69 1342.89 784.98 1329.50 C814.27 1316.12 842.82 1300.12 869.28 1282.17 C895.75 1264.21 920.80 1243.53 943.78 1221.75 C966.76 1199.97 987.68 1175.75 1007.16 1151.50 C1026.63 1127.26 1044.00 1101.55 1060.63 1076.28 C1077.25 1051.02 1092.36 1025.37 1106.93 999.90 C1121.49 974.44 1135.10 949.05 1148.04 923.50 C1160.98 897.95 1173.20 872.45 1184.56 846.60 C1195.93 820.75 1206.59 794.77 1216.25 768.39 C1225.90 742.00 1235.24 715.41 1242.49 688.29 C1249.75 661.17 1256.15 633.48 1259.80 605.67 C1263.44 577.86 1265.13 549.40 1264.38 521.43 C1263.62 493.45 1260.67 465.15 1255.26 437.82 C1249.85 410.50 1241.97 383.27 1231.91 357.50 C1221.85 331.73 1209.04 306.64 1194.89 283.19 C1180.74 259.74 1163.95 237.64 1146.99 216.82 C1130.04 196.00 1111.38 176.78 1093.17 158.25 C1074.95 139.72 1056.21 122.47 1037.71 105.62 C1019.20 88.77 1000.74 72.75 982.12 57.15 C963.50 41.55 944.95 26.52 925.99 12.02 C907.03 -2.47 887.95 -16.48 868.35 -29.83 C848.75 -43.18 827.50 -56.66 808.37 -68.07 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M-146.19 784.35 C-137.28 806.97 -128.63 826.25 -118.92 846.96 C-109.21 867.67 -98.78 888.08 -87.93 908.59 C-77.09 929.10 -65.62 949.35 -53.82 970.01 C-42.02 990.68 -29.77 1011.24 -17.14 1032.58 C-4.51 1053.91 8.22 1075.75 21.96 1098.05 C35.71 1120.35 49.39 1144.12 65.32 1166.40 C81.26 1188.68 98.28 1211.68 117.56 1231.73 C136.84 1251.78 158.33 1270.49 181.01 1286.70 C203.69 1302.92 228.40 1316.97 253.67 1329.03 C278.93 1341.10 305.65 1351.00 332.62 1359.09 C359.59 1367.19 387.52 1373.25 415.47 1377.60 C443.42 1381.95 469.95 1384.63 500.34 1385.18 C530.73 1385.73 565.59 1384.69 597.79 1380.91 C629.99 1377.13 662.34 1371.04 693.54 1362.47 C724.74 1353.90 755.69 1342.89 784.98 1329.50 C814.27 1316.12 842.82 1300.12 869.28 1282.17 C895.75 1264.21 920.80 1243.53 943.78 1221.75 C966.76 1199.97 987.68 1175.75 1007.16 1151.50 C1026.63 1127.26 1044.00 1101.55 1060.63 1076.28 C1077.25 1051.02 1092.36 1025.37 1106.93 999.90 C1121.49 974.44 1135.10 949.05 1148.04 923.50 C1160.98 897.95 1173.20 872.45 1184.56 846.60 C1195.93 820.75 1206.59 794.77 1216.25 768.39 C1225.90 742.00 1235.10 713.61 1242.49 688.29 C1249.89 662.97 1255.26 640.68 1260.62 616.48 C1265.98 592.28 1270.51 567.85 1274.67 543.08 C1278.83 518.32 1282.25 493.36 1285.57 467.89 C1288.88 442.42 1291.69 416.74 1294.58 390.24 C1297.47 363.73 1300.50 336.79 1302.92 308.87 C1305.34 280.95 1308.62 251.93 1309.13 222.73 C1309.63 193.54 1309.62 163.13 1305.94 133.71 C1302.26 104.29 1295.93 74.61 1287.05 46.21 C1278.17 17.81 1266.28 -10.09 1252.65 -36.70 C1239.02 -63.31 1222.84 -89.05 1205.29 -113.46 C1187.73 -137.87 1168.10 -161.22 1147.33 -183.15 C1126.57 -205.08 1105.95 -224.68 1080.69 -245.04 C1055.44 -265.41 1025.39 -287.38 995.81 -305.33 C966.23 -323.29 935.04 -339.42 903.19 -352.77 C871.34 -366.12 838.12 -377.29 804.73 -385.45 C771.33 -393.61 736.88 -399.14 702.84 -401.72 C668.80 -404.30 634.16 -403.75 600.49 -400.93 C566.81 -398.12 533.23 -392.00 500.79 -384.85 C468.35 -377.70 436.67 -367.97 405.86 -358.02 C375.05 -348.08 345.20 -336.77 315.93 -325.17 C286.66 -313.57 258.21 -301.32 230.23 -288.42 C202.24 -275.52 174.88 -262.11 148.03 -247.79 C121.17 -233.47 94.82 -218.51 69.08 -202.51 C43.31 -186.50 17.70 -170.02 -6.54 -151.75 C-30.82 -133.45 -54.73 -114.02 -76.44 -92.66 C-98.12 -71.26 -118.52 -48.01 -136.54 -23.59 C-154.51 0.80 -170.85 26.86 -184.51 53.69 C-198.17 80.49 -209.73 108.80 -218.55 137.33 C-227.37 165.85 -233.50 195.57 -237.45 224.85 C-241.40 254.12 -242.33 284.00 -242.23 312.99 C-242.13 341.97 -239.59 370.76 -236.85 398.75 C-234.12 426.74 -230.07 454.06 -225.84 480.91 C-221.61 507.76 -216.79 533.94 -211.47 559.86 C-206.15 585.78 -200.41 611.22 -193.90 636.44 C-187.39 661.67 -180.37 686.56 -172.42 711.21 C-164.47 735.87 -155.11 761.72 -146.19 784.35 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M1242.49 688.29 C1249.89 662.97 1255.26 640.68 1260.62 616.48 C1265.98 592.28 1270.51 567.85 1274.67 543.08 C1278.83 518.32 1282.25 493.36 1285.57 467.89 C1288.88 442.42 1291.69 416.74 1294.58 390.24 C1297.47 363.73 1300.50 336.79 1302.92 308.87 C1305.34 280.95 1308.62 251.93 1309.13 222.73 C1309.63 193.54 1309.62 163.13 1305.94 133.71 C1302.26 104.29 1295.93 74.61 1287.05 46.21 C1278.17 17.81 1266.28 -10.09 1252.65 -36.70 C1239.02 -63.31 1222.84 -89.05 1205.29 -113.46 C1187.73 -137.87 1168.10 -161.22 1147.33 -183.15 C1126.57 -205.08 1105.95 -224.68 1080.69 -245.04 C1055.44 -265.41 1025.39 -287.38 995.81 -305.33 C966.23 -323.29 935.04 -339.42 903.19 -352.77 C871.34 -366.12 838.12 -377.29 804.73 -385.45 C771.33 -393.61 736.88 -399.14 702.84 -401.72 C668.80 -404.30 634.16 -403.75 600.49 -400.93 C566.81 -398.12 533.23 -392.00 500.79 -384.85 C468.35 -377.70 436.67 -367.97 405.86 -358.02 C375.05 -348.08 345.20 -336.77 315.93 -325.17 C286.66 -313.57 258.21 -301.32 230.23 -288.42 C202.24 -275.52 174.88 -262.11 148.03 -247.79 C121.17 -233.47 94.82 -218.51 69.08 -202.51 C43.31 -186.50 16.59 -168.42 -6.54 -151.75 C-29.71 -135.05 -49.13 -119.53 -69.79 -102.32 C-90.42 -85.07 -110.45 -67.02 -130.26 -48.43 C-150.04 -29.85 -169.26 -10.63 -188.62 9.10 C-207.97 28.80 -226.97 49.02 -246.46 69.86 C-265.94 90.70 -285.73 111.85 -305.55 134.16 C-325.37 156.46 -346.47 179.15 -365.39 203.70 C-384.31 228.25 -403.48 254.10 -419.06 281.45 C-434.64 308.79 -448.20 338.00 -458.87 367.77 C-469.54 397.54 -477.42 428.77 -483.06 460.06 C-488.70 491.34 -491.70 523.53 -492.70 555.48 C-493.70 587.43 -492.32 619.83 -489.08 651.78 C-485.84 683.73 -481.25 713.62 -473.28 747.18 C-465.31 780.73 -454.37 818.75 -441.25 853.13 C-428.14 887.50 -412.51 921.40 -394.56 953.43 C-376.61 985.47 -356.12 1016.58 -333.55 1045.34 C-310.98 1074.10 -285.85 1101.39 -259.16 1126.00 C-232.46 1150.62 -203.27 1173.06 -173.37 1193.03 C-143.48 1213.00 -111.53 1230.22 -79.78 1245.82 C-48.04 1261.42 -15.23 1274.45 17.08 1286.64 C49.39 1298.84 81.84 1309.26 114.07 1318.99 C146.29 1328.72 178.32 1337.32 210.40 1345.00 C242.49 1352.68 274.45 1359.45 306.57 1365.07 C338.69 1370.69 370.84 1375.36 403.13 1378.71 C435.43 1382.06 467.90 1384.81 500.34 1385.18 C532.78 1385.54 565.59 1384.69 597.79 1380.91 C629.99 1377.13 662.34 1371.04 693.54 1362.47 C724.74 1353.90 755.69 1342.89 784.98 1329.50 C814.27 1316.12 842.82 1300.12 869.28 1282.17 C895.75 1264.21 920.80 1243.53 943.78 1221.75 C966.76 1199.97 987.68 1175.75 1007.16 1151.50 C1026.63 1127.26 1044.00 1101.55 1060.63 1076.28 C1077.25 1051.02 1092.36 1025.37 1106.93 999.90 C1121.49 974.44 1135.10 949.05 1148.04 923.50 C1160.98 897.95 1173.20 872.45 1184.56 846.60 C1195.93 820.75 1206.59 794.77 1216.25 768.39 C1225.90 742.00 1235.10 713.61 1242.49 688.29 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M-6.54 -151.75 C-29.71 -135.05 -49.13 -119.53 -69.79 -102.32 C-90.42 -85.07 -110.45 -67.02 -130.26 -48.43 C-150.04 -29.85 -169.26 -10.63 -188.62 9.10 C-207.97 28.80 -226.97 49.02 -246.46 69.86 C-265.94 90.70 -285.73 111.85 -305.55 134.16 C-325.37 156.46 -346.47 179.15 -365.39 203.70 C-384.31 228.25 -403.48 254.10 -419.06 281.45 C-434.64 308.79 -448.20 338.00 -458.87 367.77 C-469.54 397.54 -477.42 428.77 -483.06 460.06 C-488.70 491.34 -491.70 523.53 -492.70 555.48 C-493.70 587.43 -492.32 619.83 -489.08 651.78 C-485.84 683.73 -481.25 713.62 -473.28 747.18 C-465.31 780.73 -454.37 818.75 -441.25 853.13 C-428.14 887.50 -412.51 921.40 -394.56 953.43 C-376.61 985.47 -356.12 1016.58 -333.55 1045.34 C-310.98 1074.10 -285.85 1101.39 -259.16 1126.00 C-232.46 1150.62 -203.27 1173.06 -173.37 1193.03 C-143.48 1213.00 -111.53 1230.22 -79.78 1245.82 C-48.04 1261.42 -15.23 1274.45 17.08 1286.64 C49.39 1298.84 81.84 1309.26 114.07 1318.99 C146.29 1328.72 178.32 1337.32 210.40 1345.00 C242.49 1352.68 274.45 1359.45 306.57 1365.07 C338.69 1370.69 370.84 1375.36 403.13 1378.71 C435.43 1382.06 469.88 1384.21 500.34 1385.18 C530.80 1386.15 557.30 1385.66 585.88 1384.54 C614.47 1383.41 643.07 1381.24 671.85 1378.45 C700.64 1375.66 729.40 1371.92 758.57 1367.78 C787.75 1363.64 816.95 1358.77 846.90 1353.62 C876.84 1348.46 907.19 1343.15 938.23 1336.86 C969.28 1330.57 1001.53 1324.67 1033.15 1315.89 C1064.77 1307.11 1097.35 1297.37 1127.95 1284.18 C1158.54 1270.99 1188.64 1255.01 1216.73 1236.75 C1244.82 1218.49 1271.55 1197.29 1296.47 1174.62 C1321.39 1151.95 1344.71 1126.88 1366.23 1100.72 C1387.76 1074.56 1407.60 1046.52 1425.62 1017.66 C1443.64 988.80 1459.21 960.81 1474.36 927.55 C1489.51 894.30 1505.10 855.41 1516.54 818.15 C1527.97 780.88 1537.02 742.41 1542.98 703.97 C1548.93 665.52 1552.19 626.20 1552.25 587.47 C1552.32 548.73 1549.33 509.56 1543.34 471.56 C1537.36 433.55 1527.95 395.72 1516.34 359.43 C1504.74 323.14 1489.71 287.77 1473.73 253.84 C1457.75 219.91 1439.22 187.37 1420.47 155.85 C1401.72 124.33 1381.65 94.14 1361.23 64.72 C1340.81 35.30 1319.73 6.91 1297.96 -20.69 C1276.18 -48.30 1253.83 -75.06 1230.56 -100.92 C1207.29 -126.79 1183.32 -151.85 1158.34 -175.87 C1133.36 -199.89 1107.78 -223.47 1080.69 -245.04 C1053.61 -266.62 1025.39 -287.38 995.81 -305.33 C966.23 -323.29 935.04 -339.42 903.19 -352.77 C871.34 -366.12 838.12 -377.29 804.73 -385.45 C771.33 -393.61 736.88 -399.14 702.84 -401.72 C668.80 -404.30 634.16 -403.75 600.49 -400.93 C566.81 -398.12 533.23 -392.00 500.79 -384.85 C468.35 -377.70 436.67 -367.97 405.86 -358.02 C375.05 -348.08 345.20 -336.77 315.93 -325.17 C286.66 -313.57 258.21 -301.32 230.23 -288.42 C202.24 -275.52 174.88 -262.11 148.03 -247.79 C121.17 -233.47 94.82 -218.51 69.08 -202.51 C43.31 -186.50 16.59 -168.42 -6.54 -151.75 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M500.34 1385.18 C530.80 1386.15 557.30 1385.66 585.88 1384.54 C614.47 1383.41 643.07 1381.24 671.85 1378.45 C700.64 1375.66 729.40 1371.92 758.57 1367.78 C787.75 1363.64 816.95 1358.77 846.90 1353.62 C876.84 1348.46 907.19 1343.15 938.23 1336.86 C969.28 1330.57 1001.53 1324.67 1033.15 1315.89 C1064.77 1307.11 1097.35 1297.37 1127.95 1284.18 C1158.54 1270.99 1188.64 1255.01 1216.73 1236.75 C1244.82 1218.49 1271.55 1197.29 1296.47 1174.62 C1321.39 1151.95 1344.71 1126.88 1366.23 1100.72 C1387.76 1074.56 1407.60 1046.52 1425.62 1017.66 C1443.64 988.80 1459.21 960.81 1474.36 927.55 C1489.51 894.30 1505.10 855.41 1516.54 818.15 C1527.97 780.88 1537.02 742.41 1542.98 703.97 C1548.93 665.52 1552.19 626.20 1552.25 587.47 C1552.32 548.73 1549.33 509.56 1543.34 471.56 C1537.36 433.55 1527.95 395.72 1516.34 359.43 C1504.74 323.14 1489.71 287.77 1473.73 253.84 C1457.75 219.91 1439.22 187.37 1420.47 155.85 C1401.72 124.33 1381.65 94.14 1361.23 64.72 C1340.81 35.30 1319.73 6.91 1297.96 -20.69 C1276.18 -48.30 1253.83 -75.06 1230.56 -100.92 C1207.29 -126.79 1183.32 -151.85 1158.34 -175.87 C1133.36 -199.89 1105.73 -224.28 1080.69 -245.04 C1055.66 -265.81 1032.98 -282.72 1008.13 -300.43 C983.28 -318.15 957.71 -334.98 931.58 -351.33 C905.45 -367.69 878.71 -383.21 851.37 -398.58 C824.03 -413.95 796.17 -428.68 767.54 -443.55 C738.90 -458.42 709.81 -473.33 679.57 -487.80 C649.32 -502.28 618.33 -517.71 586.07 -530.39 C553.81 -543.07 520.17 -555.49 486.00 -563.88 C451.83 -572.26 416.33 -578.05 381.03 -580.71 C345.73 -583.37 309.68 -582.78 274.20 -579.83 C238.71 -576.87 203.02 -570.95 168.11 -562.98 C133.20 -555.00 98.48 -544.41 64.73 -531.98 C30.97 -519.54 -0.17 -506.26 -34.39 -488.34 C-68.63 -470.42 -106.88 -448.09 -140.65 -424.45 C-174.43 -400.80 -207.00 -374.61 -237.02 -346.45 C-267.04 -318.29 -295.40 -287.67 -320.77 -255.47 C-346.13 -223.26 -369.26 -188.74 -389.20 -153.24 C-409.12 -117.73 -426.14 -80.17 -440.36 -42.44 C-454.58 -4.71 -465.42 34.47 -474.50 73.13 C-483.58 111.78 -489.68 150.96 -494.86 189.48 C-500.04 228.00 -503.24 266.32 -505.58 304.26 C-507.92 342.19 -508.98 379.74 -508.89 417.10 C-508.80 454.45 -507.62 491.51 -505.04 528.41 C-502.47 565.30 -498.72 602.00 -493.43 638.46 C-488.13 674.92 -481.97 711.40 -473.28 747.18 C-464.58 782.95 -454.37 818.75 -441.25 853.13 C-428.14 887.50 -412.51 921.40 -394.56 953.43 C-376.61 985.47 -356.12 1016.58 -333.55 1045.34 C-310.98 1074.10 -285.85 1101.39 -259.16 1126.00 C-232.46 1150.62 -203.27 1173.06 -173.37 1193.03 C-143.48 1213.00 -111.53 1230.22 -79.78 1245.82 C-48.04 1261.42 -15.23 1274.45 17.08 1286.64 C49.39 1298.84 81.84 1309.26 114.07 1318.99 C146.29 1328.72 178.32 1337.32 210.40 1345.00 C242.49 1352.68 274.45 1359.45 306.57 1365.07 C338.69 1370.69 370.84 1375.36 403.13 1378.71 C435.43 1382.06 469.88 1384.21 500.34 1385.18 Z`

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
	__Key__000000_F_key.IsDisplayed = false
	__Key__000000_F_key.Path = `M562 -21c0 89 -65 150 -155 150c7 -44 34 -203 55 -323c71 29 100 102 100 173zM420 -206l-58 329c-59 -14 -104 -63 -104 -124c0 -49 22 -75 61 -99c12 -8 22 -13 22 -22s-9 -13 -17 -13c-80 0 -135 96 -135 166c0 94 62 190 153 217c-7 41 -14 88 -23 142 c-15 -15 -31 -29 -48 -44c-88 -76 -174 -185 -174 -307c0 -151 122 -251 265 -251c19 0 38 2 58 6zM332 822c-8 -31 -11 -65 -11 -102c0 -42 5 -81 11 -121c69 68 146 146 146 250c0 69 -24 118 -39 118c-52 0 -98 -105 -107 -145zM122 -513c0 66 45 123 115 123 c75 0 116 -57 116 -111c0 -64 -47 -104 -94 -111c-3 -1 -5 -2 -5 -4c0 -1 2 -2 3 -3c2 0 23 -5 47 -5c101 0 154 55 154 159c0 53 -11 123 -30 219c-23 -4 -50 -7 -79 -7c-186 0 -349 147 -349 334c0 200 126 321 217 406c21 17 73 70 74 71c-17 112 -22 161 -22 215 c0 84 18 212 82 288c33 39 64 51 71 51c18 0 47 -35 71 -86c16 -36 44 -110 44 -201c0 -159 -73 -284 -179 -395c9 -56 19 -115 29 -175c146 0 253 -102 253 -253c0 -103 -73 -205 -171 -237c6 -39 12 -69 15 -89c10 -57 16 -102 16 -141c0 -63 -14 -129 -68 -167 c-36 -22 -77 -34 -124 -34c-135 0 -186 87 -186 153z`
	__Key__000000_F_key.Color = `black`
	__Key__000000_F_key.FillOpacity = 0.300000
	__Key__000000_F_key.Stroke = `black`
	__Key__000000_F_key.StrokeOpacity = 0.300000
	__Key__000000_F_key.StrokeWidth = 0.000000
	__Key__000000_F_key.StrokeDashArray = ``
	__Key__000000_F_key.StrokeDashArrayWhenSelected = ``
	__Key__000000_F_key.Transform = `scale(0.2,-0.2)`

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.BackendColor = `#F1F1F1`
	__Parameter__000000_Reference.MinuteColor = `#5A5A5A`
	__Parameter__000000_Reference.HourColor = `#1E1E1E`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.InsideAngle = 120.000000
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.SideLength = 300.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.460000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.139000
	__Parameter__000000_Reference.NbPitchLines = 61
	__Parameter__000000_Reference.BeatLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbBeatLines = 49
	__Parameter__000000_Reference.NbOfBeatsInTheme = 16
	__Parameter__000000_Reference.FirstVoiceShiftX = 0.100000
	__Parameter__000000_Reference.FirstVoiceShiftY = 2.480000
	__Parameter__000000_Reference.PitchDifference = 12
	__Parameter__000000_Reference.BeatsPerSecond = 6.850000
	__Parameter__000000_Reference.Level = 11.100000
	__Parameter__000000_Reference.IsMinor = true
	__Parameter__000000_Reference.ThemeBinaryEncoding = 28743
	__Parameter__000000_Reference.OriginX = 100.000000
	__Parameter__000000_Reference.OriginY = 950.000000
	__Parameter__000000_Reference.SpiralOriginX = 500.000000
	__Parameter__000000_Reference.SpiralOriginY = 500.000000
	__Parameter__000000_Reference.OriginCrossWidth = 800.000000
	__Parameter__000000_Reference.SpiralRadiusRatio = 1.060000
	__Parameter__000000_Reference.ShowSpiralBezierConstruct = false
	__Parameter__000000_Reference.ShowInterpolationPoints = false
	__Parameter__000000_Reference.ActualBeatsTemporalShift = 6
	__Parameter__000000_Reference.PathToStaticFiles = `../../../static`
	__Parameter__000000_Reference.PathToGeneratedSVG = `../../../static/images`
	__Parameter__000000_Reference.PathToGeneratedScore = `../../../static/scores`

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 300.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.AngleDegree = -83.413224
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 120.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 300.000000
	__Rhombus__000001_Initial_Rhombus.AngleDegree = -79.106605
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 120.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 516.185401
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 59.603956
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 300.000000
	__Rhombus__000002_Rotated_Next_Rhombus.AngleDegree = -83.413224
	__Rhombus__000002_Rotated_Next_Rhombus.InsideAngle = 120.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 300.000000
	__Rhombus__000003_Rotated_Rhombus.AngleDegree = -83.413224
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 120.000000
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
	__ShapeCategory__000004_4_Construction.IsExpanded = true

	__ShapeCategory__000005_5_Vertical_growth.Name = `5. Vertical growth`
	__ShapeCategory__000005_5_Vertical_growth.IsExpanded = true

	__ShapeCategory__000006_6_Spiral_growth.Name = `6. Spiral growth`
	__ShapeCategory__000006_6_Spiral_growth.IsExpanded = true

	__ShapeCategory__000007_7_Spiral_Growth_Bezier.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = true

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = true

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = true

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 465.414815
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -38.565371
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 515.677336
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 75.197236
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 307.377182
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 567.983284
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 385.315684
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 466.099665
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 318.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterY = 0.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.HasBespokeRadius = true
	__SpiralCircle__000000_Construction_Circle_Spiral.BespopkeRadius = 14.130000
	__SpiralCircle__000000_Construction_Circle_Spiral.Color = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.FillOpacity = 0.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.Stroke = `blue`
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeOpacity = 0.800000
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeWidth = 1.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeDashArray = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.Transform = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.Pitch = 0
	__SpiralCircle__000000_Construction_Circle_Spiral.ShowName = false
	__SpiralCircle__000000_Construction_Circle_Spiral.BeatNb = 0
	__SpiralCircle__000000_Construction_Circle_Spiral.Path = `M 657.230000 500.000000 A 14.130000 14.130000 0 1 0 628.970000 500.000000 A 14.130000 14.130000 0 1 0 657.230000 500.000000 Z`

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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 465.414815
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 318.000000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -38.565371
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 465.414815
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 607.618073
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -38.565371
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -101.393523
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 168.412923
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 13.955093
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 112.931118
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 328.957108
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 465.414815
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -38.565371
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 93.577745
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -272.582658
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
	// setup of Axis instances pointers
	__Axis__000001_Construction_Axis.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Axis__000002_Initial_Axis.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Axis__000004_Rotated_Axis.ShapeCategory = __ShapeCategory__000002_2_Rotated
	// setup of AxisGrid instances pointers
	__AxisGrid__000000_Beat_Lines.Reference = __Axis__000000_Beat_Reference
	__AxisGrid__000000_Beat_Lines.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	__AxisGrid__000001_Construction_Axis_Grid.Reference = __Axis__000001_Construction_Axis
	__AxisGrid__000001_Construction_Axis_Grid.ShapeCategory = __ShapeCategory__000004_4_Construction
	__AxisGrid__000002_Pitch_Lines.Reference = __Axis__000003_Pitch_Line
	__AxisGrid__000002_Pitch_Lines.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	// setup of Bezier instances pointers
	__Bezier__000002_Growth_Bezier_Right_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__Bezier__000003_Growth_Curve_Next_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__Bezier__000004_Growth_Curve_Next_Shift_Right_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__Bezier__000005_Growth_Curve_Seed.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	// setup of BezierGrid instances pointers
	__BezierGrid__000000_2nb_Voice.Reference = __Bezier__000000_2nd_voice_seed
	__BezierGrid__000000_2nb_Voice.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000001_2nd_voice_shifted_right.Reference = __Bezier__000000_2nd_voice_seed
	__BezierGrid__000001_2nd_voice_shifted_right.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000002_First_Voice.Reference = __Bezier__000001_First_Voice_seed
	__BezierGrid__000002_First_Voice.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000003_First_Voice_Shift_Right.Reference = __Bezier__000006_FirstVoiceShiftedRightSeed
	__BezierGrid__000003_First_Voice_Shift_Right.ShapeCategory = __ShapeCategory__000009_9_Composer
	__BezierGrid__000004_Growth_Curve.Reference = __Bezier__000005_Growth_Curve_Seed
	__BezierGrid__000004_Growth_Curve.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGrid__000005_Growth_Curve_Next.Reference = __Bezier__000003_Growth_Curve_Next_Seed
	__BezierGrid__000005_Growth_Curve_Next.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.Reference = __Bezier__000004_Growth_Curve_Next_Shift_Right_Seed
	__BezierGrid__000006_Growth_Curve_Next_Shift_Right.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	__BezierGrid__000007_Growth_Curve_Shift_Right.Reference = __Bezier__000002_Growth_Bezier_Right_Seed
	__BezierGrid__000007_Growth_Curve_Shift_Right.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	// setup of BezierGridStack instances pointers
	__BezierGridStack__000000_The_GrowthCurveStack.ShapeCategory = __ShapeCategory__000005_5_Vertical_growth
	// setup of Chapter instances pointers
	// setup of Circle instances pointers
	__Circle__000000_0.ShapeCategory = __ShapeCategory__000003_3_Growing
	__Circle__000001_Construction_Circle.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Circle__000004_Initial_Circle.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Circle__000005_Rotated_Next_Circle.ShapeCategory = __ShapeCategory__000002_2_Rotated
	// setup of CircleGrid instances pointers
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
	// setup of Content instances pointers
	__Content__000000_Phyllotaxy_Music.Chapters = append(__Content__000000_Phyllotaxy_Music.Chapters, __Chapter__000000_Deep_into_the_phyllotaxy_growth_curve)
	__Content__000000_Phyllotaxy_Music.Chapters = append(__Content__000000_Phyllotaxy_Music.Chapters, __Chapter__000001_Composing_your_own_phyllotaxy_music)
	// setup of ExportToMusicxml instances pointers
	__ExportToMusicxml__000000_Singloton.Parameter = __Parameter__000000_Reference
	// setup of FrontCurve instances pointers
	// setup of FrontCurveStack instances pointers
	__FrontCurveStack__000000_Front_Curve_Stack.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000000_Non_Rotated_0_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000001_Non_Rotated_1_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000002_Non_Rotated_2_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000003_Non_Rotated_3_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000004_Non_Rotated_4_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000005_Non_Rotated_5_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000006_Non_Rotated_6_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000007_Non_Rotated_7_)
	// setup of HorizontalAxis instances pointers
	__HorizontalAxis__000000_Horizontal_Axis.ShapeCategory = __ShapeCategory__000000_0_Axes
	// setup of Key instances pointers
	__Key__000000_F_key.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	// setup of Parameter instances pointers
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
	__Parameter__000000_Reference.FirstVoiceShiftedRigth = __BezierGrid__000003_First_Voice_Shift_Right
	__Parameter__000000_Reference.SecondVoice = __BezierGrid__000000_2nb_Voice
	__Parameter__000000_Reference.SecondVoiceShiftedRight = __BezierGrid__000001_2nd_voice_shifted_right
	__Parameter__000000_Reference.FirstVoiceNotes = __CircleGrid__000002_First_Voice_notes
	__Parameter__000000_Reference.FirstVoiceNotesShiftedRight = __CircleGrid__000001_First_Voice_note_shifted_right
	__Parameter__000000_Reference.SecondVoiceNotes = __CircleGrid__000008_Second_Voice_notes
	__Parameter__000000_Reference.SecondVoiceNotesShiftedRight = __CircleGrid__000007_Second_Voice_Notes_Shift_Right
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Horizontal_Axis
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Vertical_Axis
	__Parameter__000000_Reference.SpiralOrigin = __SpiralOrigin__000000_Spiral_Origin
	// setup of Rhombus instances pointers
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.ShapeCategory = __ShapeCategory__000003_3_Growing
	__Rhombus__000001_Initial_Rhombus.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Rhombus__000002_Rotated_Next_Rhombus.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__Rhombus__000003_Rotated_Rhombus.ShapeCategory = __ShapeCategory__000002_2_Rotated
	// setup of RhombusGrid instances pointers
	__RhombusGrid__000000_Growing_Rhombus_Grid.Reference = __Rhombus__000003_Rotated_Rhombus
	__RhombusGrid__000000_Growing_Rhombus_Grid.ShapeCategory = __ShapeCategory__000003_3_Growing
	__RhombusGrid__000001_Initial_Rhombus_Grid.Reference = __Rhombus__000001_Initial_Rhombus
	__RhombusGrid__000001_Initial_Rhombus_Grid.ShapeCategory = __ShapeCategory__000001_1_Initial
	__RhombusGrid__000002_Rotated_Rhombus_Grid.ShapeCategory = __ShapeCategory__000002_2_Rotated
	// setup of ShapeCategory instances pointers
	// setup of SpiralBezier instances pointers
	__SpiralBezier__000000_Spiral_Bezier_Seed.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	// setup of SpiralBezierGrid instances pointers
	__SpiralBezierGrid__000000_Spiral_Bezier_Full_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralBezierGrid__000001_Spiral_Bezier_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	// setup of SpiralCircle instances pointers
	__SpiralCircle__000000_Construction_Circle_Spiral.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	// setup of SpiralCircleGrid instances pointers
	__SpiralCircleGrid__000000_Brute_Spiral_Bezier_Circle_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircleGrid__000003_Spiral_Circle_Grid.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	__SpiralCircleGrid__000003_Spiral_Circle_Grid.SpiralRhombusGrid = __SpiralRhombusGrid__000000_Spiral_Rhombus_Grid
	// setup of SpiralLine instances pointers
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	// setup of SpiralLineGrid instances pointers
	__SpiralLineGrid__000000_Spiral_Construction_Inner_Line_Grid_Spiral.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralLineGrid__000001_Spiral_Construction_Outer_Line_Full_Grid_Spiral.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralLineGrid__000002_Spiral_Construction_Outer_Line_Grid_Spiral.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	// setup of SpiralOrigin instances pointers
	__SpiralOrigin__000000_Spiral_Origin.ShapeCategory = __ShapeCategory__000000_0_Axes
	// setup of SpiralRhombus instances pointers
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	// setup of SpiralRhombusGrid instances pointers
	__SpiralRhombusGrid__000000_Spiral_Rhombus_Grid.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	// setup of VerticalAxis instances pointers
	__VerticalAxis__000000_Vertical_Axis.ShapeCategory = __ShapeCategory__000000_0_Axes
}

