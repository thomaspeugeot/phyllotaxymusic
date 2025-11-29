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

	const __write__local_time = "2025-11-29 10:41:48.879999 CET"
	const __write__utc_time__ = "2025-11-29 09:41:48.879999 UTC"

	const __commitId__ = "0000001398"

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
	__Axis__000001_Construction_Axis.Length = 90.000000
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -10.323708
	__Axis__000001_Construction_Axis.EndY = 89.405934
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
	__Axis__000002_Initial_Axis.Length = 392.300905
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
	__Axis__000004_Rotated_Axis.Length = 392.300905
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
	__AxisGrid__000000_Beat_Lines.IsDisplayed = true

	__AxisGrid__000001_Construction_Axis_Grid.Name = `Construction Axis Grid`
	__AxisGrid__000001_Construction_Axis_Grid.IsDisplayed = false

	__AxisGrid__000002_Pitch_Lines.Name = `Pitch Lines`
	__AxisGrid__000002_Pitch_Lines.IsDisplayed = true

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
	__Bezier__000005_Growth_Curve_Seed.StartX = -5.161854
	__Bezier__000005_Growth_Curve_Seed.StartY = 44.702967
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 44.905469
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 50.484244
	__Bezier__000005_Growth_Curve_Seed.EndX = 67.104102
	__Bezier__000005_Growth_Curve_Seed.EndY = 98.346528
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 17.036779
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 92.565251
	__Bezier__000005_Growth_Curve_Seed.Color = ``
	__Bezier__000005_Growth_Curve_Seed.FillOpacity = 0.000000
	__Bezier__000005_Growth_Curve_Seed.Stroke = `grey`
	__Bezier__000005_Growth_Curve_Seed.StrokeOpacity = 0.800000
	__Bezier__000005_Growth_Curve_Seed.StrokeWidth = 6.000000
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
	__Circle__000001_Construction_Circle.CenterX = -5.161854
	__Circle__000001_Construction_Circle.CenterY = 44.702967
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 154.855620
	__Circle__000005_Rotated_Next_Circle.CenterY = 17.881187
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M640.41 512.19 C641.07 507.41 641.46 503.22 641.82 498.67 C642.17 494.12 642.37 489.56 642.52 484.92 C642.68 480.28 642.69 475.63 642.77 470.83 C642.85 466.02 642.76 461.24 643.00 456.11 C643.25 450.97 643.44 445.84 644.22 440.01 C644.99 434.17 646.90 427.60 647.64 421.10 C648.39 414.60 649.35 407.37 648.67 401.02 C647.99 394.66 645.95 388.66 643.56 382.99 C641.18 377.32 637.89 372.02 634.37 367.00 C630.85 361.98 626.76 357.27 622.46 352.86 C618.15 348.45 613.44 344.34 608.54 340.54 C603.65 336.75 598.85 333.43 593.09 330.09 C587.32 326.76 580.54 323.24 573.96 320.54 C567.39 317.84 560.54 315.54 553.64 313.88 C546.74 312.23 539.61 311.03 532.55 310.61 C525.49 310.19 518.27 310.30 511.29 311.34 C504.32 312.38 497.27 314.21 490.71 316.84 C484.14 319.47 477.80 323.38 471.88 327.10 C465.95 330.83 460.52 335.31 455.16 339.18 C449.80 343.04 444.75 346.71 439.72 350.29 C434.69 353.86 429.80 357.17 424.99 360.62 C420.17 364.07 415.44 367.44 410.82 370.99 C406.21 374.54 401.67 378.13 397.28 381.92 C392.89 385.71 388.36 389.89 384.47 393.72 C380.57 397.54 377.35 401.03 373.93 404.86 C370.51 408.69 367.23 412.66 363.95 416.72 C360.67 420.78 357.53 424.95 354.26 429.21 C351.00 433.47 347.90 437.83 344.35 442.28 C340.80 446.73 337.31 451.21 332.97 455.89 C328.62 460.58 322.95 465.17 318.27 470.41 C313.60 475.65 308.33 481.34 304.92 487.35 C301.51 493.36 299.42 499.93 297.83 506.45 C296.25 512.96 295.63 519.74 295.39 526.44 C295.15 533.13 295.59 539.93 296.40 546.61 C297.21 553.30 298.56 560.00 300.27 566.54 C301.98 573.09 303.91 579.16 306.66 585.90 C309.41 592.63 312.92 600.20 316.77 606.94 C320.63 613.68 324.98 620.25 329.80 626.32 C334.62 632.40 339.93 638.23 345.68 643.38 C351.43 648.53 357.68 653.34 364.31 657.24 C370.94 661.13 378.15 664.42 385.48 666.75 C392.81 669.09 400.77 670.24 408.31 671.25 C415.84 672.26 423.48 672.34 430.68 672.83 C437.89 673.32 444.75 673.73 451.54 674.18 C458.33 674.62 464.86 675.16 471.42 675.50 C477.99 675.84 484.44 676.16 490.91 676.23 C497.38 676.30 503.82 676.25 510.26 675.92 C516.71 675.59 523.17 675.16 529.56 674.25 C535.96 673.34 542.40 672.20 548.62 670.47 C554.84 668.73 561.04 666.55 566.90 663.86 C572.76 661.18 578.52 658.02 583.78 654.34 C589.04 650.65 594.12 646.46 598.45 641.74 C602.79 637.01 606.70 631.65 609.78 625.97 C612.87 620.30 614.99 613.79 616.95 607.70 C618.90 601.61 620.01 595.25 621.48 589.45 C622.96 583.64 624.33 578.20 625.78 572.87 C627.23 567.53 628.76 562.51 630.18 557.43 C631.60 552.36 633.03 547.42 634.30 542.42 C635.58 537.41 636.79 532.44 637.80 527.40 C638.82 522.37 639.74 516.98 640.41 512.19 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M384.47 393.72 C380.57 397.54 377.35 401.03 373.93 404.86 C370.51 408.69 367.23 412.66 363.95 416.72 C360.67 420.78 357.53 424.95 354.26 429.21 C351.00 433.47 347.90 437.83 344.35 442.28 C340.80 446.73 337.31 451.21 332.97 455.89 C328.62 460.58 322.95 465.17 318.27 470.41 C313.60 475.65 308.33 481.34 304.92 487.35 C301.51 493.36 299.42 499.93 297.83 506.45 C296.25 512.96 295.63 519.74 295.39 526.44 C295.15 533.13 295.59 539.93 296.40 546.61 C297.21 553.30 298.56 560.00 300.27 566.54 C301.98 573.09 303.91 579.16 306.66 585.90 C309.41 592.63 312.92 600.20 316.77 606.94 C320.63 613.68 324.98 620.25 329.80 626.32 C334.62 632.40 339.93 638.23 345.68 643.38 C351.43 648.53 357.68 653.34 364.31 657.24 C370.94 661.13 378.15 664.42 385.48 666.75 C392.81 669.09 400.77 670.24 408.31 671.25 C415.84 672.26 423.48 672.34 430.68 672.83 C437.89 673.32 444.75 673.73 451.54 674.18 C458.33 674.62 464.86 675.16 471.42 675.50 C477.99 675.84 484.44 676.16 490.91 676.23 C497.38 676.30 503.82 676.25 510.26 675.92 C516.71 675.59 523.54 674.98 529.56 674.25 C535.59 673.53 540.79 672.63 546.39 671.55 C551.99 670.48 557.57 669.19 563.18 667.79 C568.79 666.40 574.36 664.80 580.07 663.19 C585.77 661.58 591.40 659.76 597.42 658.13 C603.43 656.49 609.40 654.77 616.16 653.39 C622.92 652.01 630.65 651.50 637.98 649.84 C645.30 648.18 653.36 646.51 660.11 643.43 C666.87 640.34 672.92 635.98 678.51 631.35 C684.10 626.71 689.07 621.25 693.66 615.62 C698.26 609.98 702.36 603.83 706.07 597.53 C709.78 591.23 713.05 584.58 715.92 577.82 C718.79 571.05 721.16 564.56 723.28 556.95 C725.39 549.35 727.43 540.52 728.62 532.18 C729.81 523.84 730.47 515.31 730.42 506.90 C730.38 498.49 729.76 489.96 728.34 481.70 C726.91 473.44 724.86 465.15 721.90 457.36 C718.94 449.56 715.12 441.90 710.58 434.93 C706.04 427.95 700.24 421.53 694.66 415.50 C689.08 409.47 682.81 404.13 677.11 398.75 C671.40 393.37 665.89 388.27 660.43 383.22 C654.96 378.17 649.72 373.23 644.32 368.44 C638.91 363.66 633.55 358.98 628.00 354.52 C622.44 350.06 616.82 345.74 611.00 341.67 C605.18 337.59 599.26 333.61 593.09 330.09 C586.91 326.57 580.54 323.24 573.96 320.54 C567.39 317.84 560.54 315.54 553.64 313.88 C546.74 312.23 539.61 311.03 532.55 310.61 C525.49 310.19 518.27 310.30 511.29 311.34 C504.32 312.38 497.27 314.21 490.71 316.84 C484.14 319.47 477.80 323.38 471.88 327.10 C465.95 330.83 460.52 335.31 455.16 339.18 C449.80 343.04 444.75 346.71 439.72 350.29 C434.69 353.86 429.80 357.17 424.99 360.62 C420.17 364.07 415.44 367.44 410.82 370.99 C406.21 374.54 401.67 378.13 397.28 381.92 C392.89 385.71 388.36 389.89 384.47 393.72 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M529.56 674.25 C535.59 673.53 540.79 672.63 546.39 671.55 C551.99 670.48 557.57 669.19 563.18 667.79 C568.79 666.40 574.36 664.80 580.07 663.19 C585.77 661.58 591.40 659.76 597.42 658.13 C603.43 656.49 609.40 654.77 616.16 653.39 C622.92 652.01 630.65 651.50 637.98 649.84 C645.30 648.18 653.36 646.51 660.11 643.43 C666.87 640.34 672.92 635.98 678.51 631.35 C684.10 626.71 689.07 621.25 693.66 615.62 C698.26 609.98 702.36 603.83 706.07 597.53 C709.78 591.23 713.05 584.58 715.92 577.82 C718.79 571.05 721.16 564.56 723.28 556.95 C725.39 549.35 727.43 540.52 728.62 532.18 C729.81 523.84 730.47 515.31 730.42 506.90 C730.38 498.49 729.76 489.96 728.34 481.70 C726.91 473.44 724.86 465.15 721.90 457.36 C718.94 449.56 715.12 441.90 710.58 434.93 C706.04 427.95 700.24 421.53 694.66 415.50 C689.08 409.47 682.81 404.13 677.11 398.75 C671.40 393.37 665.89 388.27 660.43 383.22 C654.96 378.17 649.72 373.23 644.32 368.44 C638.91 363.66 633.55 358.98 628.00 354.52 C622.44 350.06 616.82 345.74 611.00 341.67 C605.18 337.59 598.82 333.51 593.09 330.09 C587.36 326.67 582.22 323.95 576.61 321.14 C571.01 318.32 565.28 315.72 559.45 313.20 C553.62 310.68 547.69 308.38 541.64 306.03 C535.60 303.68 529.47 301.57 523.16 299.10 C516.84 296.63 510.50 294.29 503.75 291.22 C497.00 288.15 490.02 283.83 482.65 280.69 C475.28 277.55 467.34 273.99 459.55 272.38 C451.77 270.77 443.76 270.63 435.95 271.02 C428.14 271.40 420.32 272.85 412.70 274.69 C405.08 276.54 397.54 279.11 390.23 282.07 C382.92 285.03 375.75 288.57 368.84 292.46 C361.94 296.36 355.63 300.34 348.80 305.43 C341.97 310.53 334.41 316.71 327.87 323.03 C321.33 329.35 315.12 336.17 309.57 343.36 C304.01 350.54 298.86 358.19 294.55 366.14 C290.23 374.09 286.43 382.48 283.67 391.04 C280.92 399.60 278.95 408.60 278.02 417.51 C277.08 426.42 277.51 435.68 278.04 444.51 C278.58 453.33 280.11 462.05 281.22 470.45 C282.32 478.85 283.48 486.92 284.67 494.92 C285.86 502.92 286.98 510.70 288.36 518.45 C289.74 526.20 291.18 533.84 292.95 541.42 C294.71 548.99 296.66 556.50 298.95 563.91 C301.23 571.33 303.69 578.72 306.66 585.90 C309.63 593.07 312.92 600.20 316.77 606.94 C320.63 613.68 324.98 620.25 329.80 626.32 C334.62 632.40 339.93 638.23 345.68 643.38 C351.43 648.53 357.68 653.34 364.31 657.24 C370.94 661.13 378.15 664.42 385.48 666.75 C392.81 669.09 400.77 670.24 408.31 671.25 C415.84 672.26 423.48 672.34 430.68 672.83 C437.89 673.32 444.75 673.73 451.54 674.18 C458.33 674.62 464.86 675.16 471.42 675.50 C477.99 675.84 484.44 676.16 490.91 676.23 C497.38 676.30 503.82 676.25 510.26 675.92 C516.71 675.59 523.54 674.98 529.56 674.25 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M593.09 330.09 C587.36 326.67 582.22 323.95 576.61 321.14 C571.01 318.32 565.28 315.72 559.45 313.20 C553.62 310.68 547.69 308.38 541.64 306.03 C535.60 303.68 529.47 301.57 523.16 299.10 C516.84 296.63 510.50 294.29 503.75 291.22 C497.00 288.15 490.02 283.83 482.65 280.69 C475.28 277.55 467.34 273.99 459.55 272.38 C451.77 270.77 443.76 270.63 435.95 271.02 C428.14 271.40 420.32 272.85 412.70 274.69 C405.08 276.54 397.54 279.11 390.23 282.07 C382.92 285.03 375.75 288.57 368.84 292.46 C361.94 296.36 355.63 300.34 348.80 305.43 C341.97 310.53 334.41 316.71 327.87 323.03 C321.33 329.35 315.12 336.17 309.57 343.36 C304.01 350.54 298.86 358.19 294.55 366.14 C290.23 374.09 286.43 382.48 283.67 391.04 C280.92 399.60 278.95 408.60 278.02 417.51 C277.08 426.42 277.51 435.68 278.04 444.51 C278.58 453.33 280.11 462.05 281.22 470.45 C282.32 478.85 283.48 486.92 284.67 494.92 C285.86 502.92 286.98 510.70 288.36 518.45 C289.74 526.20 291.18 533.84 292.95 541.42 C294.71 548.99 296.66 556.50 298.95 563.91 C301.23 571.33 304.00 579.11 306.66 585.90 C309.32 592.68 311.95 598.46 314.93 604.62 C317.90 610.79 321.13 616.85 324.51 622.89 C327.89 628.93 331.50 634.85 335.20 640.86 C338.90 646.86 342.83 652.71 346.71 658.94 C350.59 665.17 354.58 671.29 358.47 678.22 C362.36 685.15 365.68 693.19 370.05 700.51 C374.41 707.82 379.00 715.80 384.65 722.12 C390.31 728.44 397.07 733.72 404.00 738.42 C410.92 743.13 418.52 746.98 426.21 750.38 C433.90 753.77 441.99 756.51 450.13 758.77 C458.27 761.03 466.67 762.73 475.07 763.95 C483.47 765.18 491.42 765.94 500.55 766.11 C509.67 766.28 520.12 766.01 529.80 764.96 C539.49 763.91 549.22 762.24 558.64 759.83 C568.07 757.43 577.45 754.36 586.35 750.52 C595.25 746.69 604.00 742.15 612.02 736.80 C620.04 731.45 627.69 725.23 634.46 718.42 C641.23 711.61 647.12 703.66 652.63 695.95 C658.15 688.25 662.79 679.99 667.55 672.22 C672.31 664.44 676.80 656.87 681.19 649.32 C685.58 641.76 689.88 634.39 693.91 626.87 C697.94 619.35 701.82 611.86 705.38 604.20 C708.94 596.54 712.28 588.81 715.26 580.94 C718.25 573.06 721.05 565.08 723.28 556.95 C725.50 548.82 727.43 540.52 728.62 532.18 C729.81 523.84 730.47 515.31 730.42 506.90 C730.38 498.49 729.76 489.96 728.34 481.70 C726.91 473.44 724.86 465.15 721.90 457.36 C718.94 449.56 715.12 441.90 710.58 434.93 C706.04 427.95 700.24 421.53 694.66 415.50 C689.08 409.47 682.81 404.13 677.11 398.75 C671.40 393.37 665.89 388.27 660.43 383.22 C654.96 378.17 649.72 373.23 644.32 368.44 C638.91 363.66 633.55 358.98 628.00 354.52 C622.44 350.06 616.82 345.74 611.00 341.67 C605.18 337.59 598.82 333.51 593.09 330.09 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M306.66 585.90 C309.32 592.68 311.95 598.46 314.93 604.62 C317.90 610.79 321.13 616.85 324.51 622.89 C327.89 628.93 331.50 634.85 335.20 640.86 C338.90 646.86 342.83 652.71 346.71 658.94 C350.59 665.17 354.58 671.29 358.47 678.22 C362.36 685.15 365.68 693.19 370.05 700.51 C374.41 707.82 379.00 715.80 384.65 722.12 C390.31 728.44 397.07 733.72 404.00 738.42 C410.92 743.13 418.52 746.98 426.21 750.38 C433.90 753.77 441.99 756.51 450.13 758.77 C458.27 761.03 466.67 762.73 475.07 763.95 C483.47 765.18 491.42 765.94 500.55 766.11 C509.67 766.28 520.12 766.01 529.80 764.96 C539.49 763.91 549.22 762.24 558.64 759.83 C568.07 757.43 577.45 754.36 586.35 750.52 C595.25 746.69 604.00 742.15 612.02 736.80 C620.04 731.45 627.69 725.23 634.46 718.42 C641.23 711.61 647.12 703.66 652.63 695.95 C658.15 688.25 662.79 679.99 667.55 672.22 C672.31 664.44 676.80 656.87 681.19 649.32 C685.58 641.76 689.88 634.39 693.91 626.87 C697.94 619.35 701.82 611.86 705.38 604.20 C708.94 596.54 712.28 588.81 715.26 580.94 C718.25 573.06 721.05 564.54 723.28 556.95 C725.50 549.36 727.08 542.65 728.61 535.39 C730.14 528.12 731.38 520.77 732.46 513.35 C733.54 505.93 734.33 498.46 735.08 490.87 C735.83 483.27 736.28 475.66 736.97 467.79 C737.66 459.92 738.17 452.06 739.22 443.64 C740.28 435.21 742.44 426.26 743.31 417.26 C744.18 408.25 745.26 398.60 744.44 389.60 C743.63 380.61 741.28 371.78 738.43 363.30 C735.57 354.83 731.64 346.63 727.33 338.76 C723.02 330.89 717.97 323.30 712.57 316.08 C707.16 308.85 701.19 301.94 694.90 295.41 C688.61 288.89 682.41 283.04 674.84 276.93 C667.27 270.82 658.29 264.21 649.46 258.74 C640.63 253.27 631.35 248.30 621.85 244.13 C612.35 239.96 602.46 236.36 592.47 233.72 C582.48 231.07 572.16 229.10 561.91 228.25 C551.66 227.41 541.16 227.52 530.99 228.67 C520.83 229.82 510.64 232.45 500.91 235.15 C491.17 237.85 481.78 241.57 472.60 244.86 C463.42 248.16 454.58 251.50 445.83 254.92 C437.09 258.34 428.54 261.71 420.12 265.39 C411.70 269.07 403.42 272.86 395.32 277.01 C387.21 281.15 379.24 285.53 371.49 290.27 C363.74 295.01 356.07 299.97 348.80 305.43 C341.53 310.89 334.41 316.71 327.87 323.03 C321.33 329.35 315.12 336.17 309.57 343.36 C304.01 350.54 298.86 358.19 294.55 366.14 C290.23 374.09 286.43 382.48 283.67 391.04 C280.92 399.60 278.95 408.60 278.02 417.51 C277.08 426.42 277.51 435.68 278.04 444.51 C278.58 453.33 280.11 462.05 281.22 470.45 C282.32 478.85 283.48 486.92 284.67 494.92 C285.86 502.92 286.98 510.70 288.36 518.45 C289.74 526.20 291.18 533.84 292.95 541.42 C294.71 548.99 296.66 556.50 298.95 563.91 C301.23 571.33 304.00 579.11 306.66 585.90 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M723.28 556.95 C725.50 549.36 727.08 542.65 728.61 535.39 C730.14 528.12 731.38 520.77 732.46 513.35 C733.54 505.93 734.33 498.46 735.08 490.87 C735.83 483.27 736.28 475.66 736.97 467.79 C737.66 459.92 738.17 452.06 739.22 443.64 C740.28 435.21 742.44 426.26 743.31 417.26 C744.18 408.25 745.26 398.60 744.44 389.60 C743.63 380.61 741.28 371.78 738.43 363.30 C735.57 354.83 731.64 346.63 727.33 338.76 C723.02 330.89 717.97 323.30 712.57 316.08 C707.16 308.85 701.19 301.94 694.90 295.41 C688.61 288.89 682.41 283.04 674.84 276.93 C667.27 270.82 658.29 264.21 649.46 258.74 C640.63 253.27 631.35 248.30 621.85 244.13 C612.35 239.96 602.46 236.36 592.47 233.72 C582.48 231.07 572.16 229.10 561.91 228.25 C551.66 227.41 541.16 227.52 530.99 228.67 C520.83 229.82 510.64 232.45 500.91 235.15 C491.17 237.85 481.78 241.57 472.60 244.86 C463.42 248.16 454.58 251.50 445.83 254.92 C437.09 258.34 428.54 261.71 420.12 265.39 C411.70 269.07 403.42 272.86 395.32 277.01 C387.21 281.15 379.24 285.53 371.49 290.27 C363.74 295.01 355.71 300.42 348.80 305.43 C341.89 310.44 336.12 315.13 330.03 320.33 C323.94 325.53 318.04 330.99 312.25 336.63 C306.45 342.26 300.88 348.14 295.28 354.14 C289.68 360.15 284.34 366.37 278.66 372.67 C272.99 378.96 267.49 385.38 261.24 391.90 C254.99 398.43 247.56 404.71 241.15 411.83 C234.74 418.95 227.79 426.47 222.77 434.63 C217.75 442.79 214.07 451.77 211.03 460.78 C207.99 469.79 206.01 479.25 204.55 488.69 C203.09 498.12 202.41 507.78 202.27 517.38 C202.13 526.97 202.65 536.67 203.69 546.25 C204.74 555.82 206.14 564.77 208.52 574.84 C210.90 584.91 214.12 596.31 217.97 606.66 C221.81 617.01 226.35 627.24 231.58 636.95 C236.82 646.66 242.74 656.14 249.36 664.91 C255.98 673.69 263.31 682.12 271.30 689.58 C279.28 697.05 288.09 703.90 297.29 709.70 C306.49 715.49 316.64 720.11 326.50 724.36 C336.35 728.61 346.59 731.83 356.41 735.21 C366.23 738.59 375.83 741.70 385.42 744.66 C395.01 747.61 404.43 750.46 413.95 752.94 C423.47 755.42 432.96 757.69 442.53 759.54 C452.11 761.39 461.73 762.94 471.40 764.04 C481.06 765.13 490.81 765.95 500.55 766.11 C510.28 766.26 520.12 766.01 529.80 764.96 C539.49 763.91 549.22 762.24 558.64 759.83 C568.07 757.43 577.45 754.36 586.35 750.52 C595.25 746.69 604.00 742.15 612.02 736.80 C620.04 731.45 627.69 725.23 634.46 718.42 C641.23 711.61 647.12 703.66 652.63 695.95 C658.15 688.25 662.79 679.99 667.55 672.22 C672.31 664.44 676.80 656.87 681.19 649.32 C685.58 641.76 689.88 634.39 693.91 626.87 C697.94 619.35 701.82 611.86 705.38 604.20 C708.94 596.54 712.28 588.81 715.26 580.94 C718.25 573.06 721.05 564.54 723.28 556.95 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M348.80 305.43 C341.89 310.44 336.12 315.13 330.03 320.33 C323.94 325.53 318.04 330.99 312.25 336.63 C306.45 342.26 300.88 348.14 295.28 354.14 C289.68 360.15 284.34 366.37 278.66 372.67 C272.99 378.96 267.49 385.38 261.24 391.90 C254.99 398.43 247.56 404.71 241.15 411.83 C234.74 418.95 227.79 426.47 222.77 434.63 C217.75 442.79 214.07 451.77 211.03 460.78 C207.99 469.79 206.01 479.25 204.55 488.69 C203.09 498.12 202.41 507.78 202.27 517.38 C202.13 526.97 202.65 536.67 203.69 546.25 C204.74 555.82 206.14 564.77 208.52 574.84 C210.90 584.91 214.12 596.31 217.97 606.66 C221.81 617.01 226.35 627.24 231.58 636.95 C236.82 646.66 242.74 656.14 249.36 664.91 C255.98 673.69 263.31 682.12 271.30 689.58 C279.28 697.05 288.09 703.90 297.29 709.70 C306.49 715.49 316.64 720.11 326.50 724.36 C336.35 728.61 346.59 731.83 356.41 735.21 C366.23 738.59 375.83 741.70 385.42 744.66 C395.01 747.61 404.43 750.46 413.95 752.94 C423.47 755.42 432.96 757.69 442.53 759.54 C452.11 761.39 461.73 762.94 471.40 764.04 C481.06 765.13 491.41 765.81 500.55 766.11 C509.68 766.40 517.64 766.21 526.20 765.80 C534.76 765.38 543.33 764.61 551.92 763.60 C560.50 762.60 569.06 761.25 577.71 759.78 C586.35 758.31 594.93 756.48 603.79 754.78 C612.65 753.08 621.45 751.14 630.88 749.56 C640.32 747.98 650.52 747.28 660.42 745.28 C670.32 743.28 680.88 741.26 690.29 737.54 C699.71 733.81 708.56 728.59 716.89 722.94 C725.23 717.28 733.00 710.60 740.31 703.61 C747.63 696.61 754.45 688.93 760.78 680.97 C767.11 673.01 772.96 664.53 778.31 655.82 C783.65 647.12 788.28 638.71 792.84 628.73 C797.39 618.75 802.11 607.10 805.64 595.94 C809.16 584.77 812.03 573.27 814.00 561.74 C815.97 550.21 817.21 538.43 817.44 526.78 C817.67 515.13 817.10 503.30 815.37 491.85 C813.65 480.39 810.87 468.91 807.09 458.03 C803.31 447.15 798.02 436.62 792.70 426.57 C787.39 416.53 781.10 407.07 775.21 397.75 C769.32 388.43 763.40 379.47 757.36 370.65 C751.32 361.84 745.32 353.24 738.98 344.88 C732.64 336.52 726.17 328.36 719.35 320.51 C712.52 312.66 705.44 305.04 698.03 297.78 C690.61 290.51 682.94 283.44 674.84 276.93 C666.75 270.43 658.29 264.21 649.46 258.74 C640.63 253.27 631.35 248.30 621.85 244.13 C612.35 239.96 602.46 236.36 592.47 233.72 C582.48 231.07 572.16 229.10 561.91 228.25 C551.66 227.41 541.16 227.52 530.99 228.67 C520.83 229.82 510.64 232.45 500.91 235.15 C491.17 237.85 481.78 241.57 472.60 244.86 C463.42 248.16 454.58 251.50 445.83 254.92 C437.09 258.34 428.54 261.71 420.12 265.39 C411.70 269.07 403.42 272.86 395.32 277.01 C387.21 281.15 379.24 285.53 371.49 290.27 C363.74 295.01 355.71 300.42 348.80 305.43 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M500.55 766.11 C509.68 766.40 517.64 766.21 526.20 765.80 C534.76 765.38 543.33 764.61 551.92 763.60 C560.50 762.60 569.06 761.25 577.71 759.78 C586.35 758.31 594.93 756.48 603.79 754.78 C612.65 753.08 621.45 751.14 630.88 749.56 C640.32 747.98 650.52 747.28 660.42 745.28 C670.32 743.28 680.88 741.26 690.29 737.54 C699.71 733.81 708.56 728.59 716.89 722.94 C725.23 717.28 733.00 710.60 740.31 703.61 C747.63 696.61 754.45 688.93 760.78 680.97 C767.11 673.01 772.96 664.53 778.31 655.82 C783.65 647.12 788.28 638.71 792.84 628.73 C797.39 618.75 802.11 607.10 805.64 595.94 C809.16 584.77 812.03 573.27 814.00 561.74 C815.97 550.21 817.21 538.43 817.44 526.78 C817.67 515.13 817.10 503.30 815.37 491.85 C813.65 480.39 810.87 468.91 807.09 458.03 C803.31 447.15 798.02 436.62 792.70 426.57 C787.39 416.53 781.10 407.07 775.21 397.75 C769.32 388.43 763.40 379.47 757.36 370.65 C751.32 361.84 745.32 353.24 738.98 344.88 C732.64 336.52 726.17 328.36 719.35 320.51 C712.52 312.66 705.44 305.04 698.03 297.78 C690.61 290.51 682.34 283.16 674.84 276.93 C667.34 270.71 660.51 265.66 653.02 260.42 C645.53 255.18 637.80 250.24 629.91 245.50 C622.01 240.76 613.91 236.33 605.66 231.97 C597.41 227.61 588.98 223.60 580.40 219.34 C571.81 215.08 563.13 211.08 554.16 206.43 C545.18 201.77 536.14 195.98 526.54 191.43 C516.93 186.88 506.79 181.96 496.51 179.12 C486.24 176.29 475.48 175.01 464.86 174.42 C454.25 173.83 443.45 174.43 432.83 175.59 C422.20 176.74 411.55 178.77 401.12 181.35 C390.68 183.93 380.32 187.23 370.23 191.06 C360.13 194.89 350.83 198.92 340.56 204.31 C330.29 209.71 318.81 216.38 308.63 223.42 C298.44 230.46 288.58 238.21 279.44 246.55 C270.30 254.89 261.59 263.90 253.79 273.46 C246.00 283.01 238.76 293.22 232.68 303.87 C226.60 314.51 221.37 325.84 217.33 337.33 C213.28 348.81 210.69 360.97 208.43 372.77 C206.18 384.56 205.10 396.51 203.82 408.12 C202.54 419.72 201.54 431.07 200.76 442.38 C199.99 453.69 199.35 464.82 199.16 475.96 C198.97 487.10 199.04 498.18 199.62 509.22 C200.21 520.26 201.16 531.28 202.64 542.22 C204.13 553.16 205.97 564.10 208.52 574.84 C211.07 585.58 214.12 596.31 217.97 606.66 C221.81 617.01 226.35 627.24 231.58 636.95 C236.82 646.66 242.74 656.14 249.36 664.91 C255.98 673.69 263.31 682.12 271.30 689.58 C279.28 697.05 288.09 703.90 297.29 709.70 C306.49 715.49 316.64 720.11 326.50 724.36 C336.35 728.61 346.59 731.83 356.41 735.21 C366.23 738.59 375.83 741.70 385.42 744.66 C395.01 747.61 404.43 750.46 413.95 752.94 C423.47 755.42 432.96 757.69 442.53 759.54 C452.11 761.39 461.73 762.94 471.40 764.04 C481.06 765.13 491.41 765.81 500.55 766.11 Z`

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
	__HorizontalAxis__000000_Horizontal_Axis.IsDisplayed = true
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
	__Parameter__000000_Reference.SideLength = 90.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.560000
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
	__Parameter__000000_Reference.ThemeBinaryEncoding = 71
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
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 90.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 90.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 154.855620
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 17.881187
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 90.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 90.000000
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
	__ShapeCategory__000000_0_Axes.IsExpanded = false

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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 139.624444
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -11.569611
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 154.703201
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 22.559171
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 92.213155
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 170.394985
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 115.594705
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 139.829900
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 95.400000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 139.624444
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 95.400000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -11.569611
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 139.624444
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 182.285422
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -11.569611
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -30.418057
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 50.523877
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 4.186528
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 33.879335
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 98.687132
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 139.624444
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -11.569611
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 28.073323
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -81.774797
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

