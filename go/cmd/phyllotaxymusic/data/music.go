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

	const __write__local_time = "2025-06-14 18:50:53.026389 CEST"
	const __write__utc_time__ = "2025-06-14 16:50:53.026389 UTC"

	const __commitId__ = "0000001147"

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
	__Axis__000001_Construction_Axis.AngleDegree = 108.120541
	__Axis__000001_Construction_Axis.Length = 91.356905
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -28.413566
	__Axis__000001_Construction_Axis.EndY = 86.825995
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
	__Axis__000002_Initial_Axis.AngleDegree = 71.879459
	__Axis__000002_Initial_Axis.Length = 734.339718
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
	__Axis__000004_Rotated_Axis.Length = 734.339718
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -14.206783
	__Bezier__000005_Growth_Curve_Seed.StartY = 43.412997
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 93.569063
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 78.682345
	__Bezier__000005_Growth_Curve_Seed.EndX = 45.287078
	__Bezier__000005_Growth_Curve_Seed.EndY = 110.944327
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = -62.488768
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 75.674979
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
	__Circle__000001_Construction_Circle.CenterX = -14.206783
	__Circle__000001_Construction_Circle.CenterY = 43.412997
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 323.216145
	__Circle__000005_Rotated_Next_Circle.CenterY = 9.647333
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M638.57 517.46 C639.54 515.14 640.09 513.67 640.81 511.73 C641.54 509.79 642.23 507.83 642.90 505.84 C643.58 503.85 638.30 502.12 644.85 499.79 C651.41 497.46 677.01 494.37 682.23 491.85 C687.44 489.33 677.86 487.02 676.14 484.66 C674.41 482.31 673.34 479.97 671.87 477.70 C670.40 475.43 669.23 473.15 667.33 471.07 C665.44 468.99 655.85 468.95 660.49 465.24 C665.12 461.53 689.48 453.00 695.16 448.80 C700.84 444.60 694.85 442.97 694.56 440.05 C694.28 437.13 693.91 434.21 693.47 431.29 C693.03 428.38 692.69 426.15 691.90 422.55 C691.11 418.96 689.96 413.96 688.73 409.72 C687.50 405.47 688.14 399.92 684.53 397.08 C680.92 394.24 669.02 396.70 667.06 392.68 C665.10 388.66 673.14 378.08 672.79 372.96 C672.44 367.85 669.95 363.43 664.98 362.01 C660.01 360.60 650.13 363.29 642.96 364.48 C635.79 365.66 626.79 369.65 621.94 369.11 C617.10 368.57 615.18 365.87 613.89 361.26 C612.61 356.66 616.97 344.10 614.21 341.46 C611.45 338.82 601.62 346.05 597.33 345.44 C593.03 344.84 591.48 340.32 588.43 337.84 C585.38 335.36 582.27 332.90 579.04 330.56 C575.81 328.23 572.50 325.95 569.08 323.84 C565.66 321.73 562.71 317.37 558.52 317.89 C554.32 318.41 547.69 328.45 543.90 326.96 C540.12 325.46 539.16 312.33 535.82 308.91 C532.47 305.49 528.02 304.02 523.82 306.44 C519.62 308.85 514.64 317.61 510.61 323.40 C506.58 329.19 503.12 338.26 499.61 341.18 C496.11 344.10 493.27 343.42 489.58 340.91 C485.89 338.41 480.91 326.22 477.49 326.13 C474.07 326.04 472.22 337.96 469.06 340.38 C465.90 342.80 462.07 340.47 458.54 340.65 C455.01 340.82 451.43 341.04 447.87 341.42 C444.30 341.80 440.69 342.26 437.12 342.91 C433.55 343.56 428.80 342.24 426.44 345.31 C424.09 348.38 426.46 360.05 423.02 361.35 C419.58 362.65 410.26 353.52 405.82 353.10 C401.37 352.68 397.45 354.42 396.35 358.82 C395.25 363.21 397.91 372.74 399.22 379.46 C400.53 386.18 404.50 394.89 404.22 399.14 C403.95 403.38 401.69 404.54 397.55 404.94 C393.42 405.33 381.71 399.55 379.41 401.51 C377.10 403.47 384.13 413.11 383.71 416.70 C383.29 420.30 379.12 420.88 376.87 423.08 C374.62 425.28 372.04 427.94 370.20 429.89 C368.36 431.84 367.26 433.10 365.81 434.76 C364.37 436.42 362.94 438.12 361.53 439.85 C360.12 441.58 364.34 445.25 357.34 445.14 C350.33 445.03 325.33 438.64 319.51 439.20 C313.69 439.76 321.71 445.55 322.41 448.51 C323.11 451.47 323.21 454.17 323.71 456.95 C324.21 459.73 324.43 462.42 325.40 465.18 C326.38 467.95 335.26 471.64 329.56 473.53 C323.86 475.41 298.04 474.54 291.19 476.49 C284.35 478.45 289.31 482.31 288.49 485.26 C287.67 488.22 286.93 491.21 286.27 494.22 C285.61 497.23 285.10 499.53 284.51 503.33 C283.93 507.14 283.17 512.44 282.77 517.04 C282.37 521.64 279.76 526.77 282.10 530.94 C284.45 535.10 296.46 537.39 296.84 542.04 C297.22 546.69 285.90 553.78 284.40 558.86 C282.90 563.93 283.65 569.16 287.82 572.49 C291.99 575.83 302.24 577.13 309.43 578.85 C316.61 580.57 326.53 580.36 330.93 582.81 C335.33 585.27 336.21 588.67 335.81 593.60 C335.42 598.54 326.85 608.78 328.56 612.42 C330.27 616.05 342.18 613.11 346.07 615.41 C349.97 617.71 349.88 622.65 351.93 626.23 C353.99 629.81 356.12 633.39 358.40 636.89 C360.69 640.39 363.08 643.87 365.63 647.22 C368.19 650.58 369.48 655.84 373.73 657.03 C377.98 658.21 388.00 651.44 391.13 654.33 C394.27 657.23 390.50 669.91 392.53 674.42 C394.56 678.94 398.34 682.03 403.31 681.43 C408.28 680.82 416.30 674.60 422.36 670.78 C428.42 666.96 435.15 659.87 439.67 658.52 C444.19 657.16 446.77 658.91 449.47 662.65 C452.17 666.40 452.53 679.62 455.87 681.01 C459.21 682.39 465.51 672.03 469.53 670.98 C473.56 669.92 476.46 673.53 480.01 674.69 C483.56 675.85 487.16 676.97 490.82 677.94 C494.48 678.92 498.20 679.81 501.96 680.52 C505.72 681.23 509.86 684.20 513.37 682.19 C516.89 680.18 519.17 668.44 523.04 668.46 C526.92 668.49 532.43 680.37 536.60 682.36 C540.78 684.35 545.26 684.12 548.08 680.39 C550.91 676.65 552.14 666.75 553.57 659.95 C555.01 653.16 554.71 643.53 556.71 639.61 C558.71 635.68 561.42 635.35 565.59 636.42 C569.76 637.48 578.69 647.07 581.73 645.99 C584.77 644.92 581.95 633.27 583.83 629.97 C585.71 626.67 589.96 627.54 593.02 626.20 C596.09 624.87 599.17 623.47 602.20 621.94 C605.23 620.41 608.26 618.79 611.20 617.02 C614.13 615.24 618.90 614.87 619.83 611.30 C620.75 607.73 614.13 597.91 616.72 595.59 C619.32 593.28 631.22 598.50 635.39 597.43 C639.56 596.36 642.46 593.47 641.76 589.15 C641.06 584.84 634.98 577.12 631.20 571.54 C627.41 565.95 620.43 559.50 619.04 555.63 C617.66 551.77 619.25 550.05 622.88 548.35 C626.51 546.65 639.45 547.93 640.82 545.45 C642.18 542.97 632.04 536.79 631.07 533.49 C630.09 530.19 633.71 528.31 634.96 525.64 C636.21 522.97 637.59 519.78 638.57 517.46 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M370.20 429.89 C368.36 431.84 367.26 433.10 365.81 434.76 C364.37 436.42 362.94 438.12 361.53 439.85 C360.12 441.58 364.34 445.25 357.34 445.14 C350.33 445.03 325.33 438.64 319.51 439.20 C313.69 439.76 321.71 445.55 322.41 448.51 C323.11 451.47 323.21 454.17 323.71 456.95 C324.21 459.73 324.43 462.42 325.40 465.18 C326.38 467.95 335.26 471.64 329.56 473.53 C323.86 475.41 298.04 474.54 291.19 476.49 C284.35 478.45 289.31 482.31 288.49 485.26 C287.67 488.22 286.93 491.21 286.27 494.22 C285.61 497.23 285.10 499.53 284.51 503.33 C283.93 507.14 283.17 512.44 282.77 517.04 C282.37 521.64 279.76 526.77 282.10 530.94 C284.45 535.10 296.46 537.39 296.84 542.04 C297.22 546.69 285.90 553.78 284.40 558.86 C282.90 563.93 283.65 569.16 287.82 572.49 C291.99 575.83 302.24 577.13 309.43 578.85 C316.61 580.57 326.53 580.36 330.93 582.81 C335.33 585.27 336.21 588.67 335.81 593.60 C335.42 598.54 326.85 608.78 328.56 612.42 C330.27 616.05 342.18 613.11 346.07 615.41 C349.97 617.71 349.88 622.65 351.93 626.23 C353.99 629.81 356.12 633.39 358.40 636.89 C360.69 640.39 363.08 643.87 365.63 647.22 C368.19 650.58 369.48 655.84 373.73 657.03 C377.98 658.21 388.00 651.44 391.13 654.33 C394.27 657.23 390.50 669.91 392.53 674.42 C394.56 678.94 398.34 682.03 403.31 681.43 C408.28 680.82 416.30 674.60 422.36 670.78 C428.42 666.96 435.15 659.87 439.67 658.52 C444.19 657.16 446.77 658.91 449.47 662.65 C452.17 666.40 452.53 679.62 455.87 681.01 C459.21 682.39 465.51 672.03 469.53 670.98 C473.56 669.92 476.46 673.53 480.01 674.69 C483.56 675.85 487.16 676.97 490.82 677.94 C494.48 678.92 498.20 679.81 501.96 680.52 C505.72 681.23 509.86 684.20 513.37 682.19 C516.89 680.18 519.17 668.44 523.04 668.46 C526.92 668.49 532.43 680.37 536.60 682.36 C540.78 684.35 545.26 684.12 548.08 680.39 C550.91 676.65 552.14 666.75 553.57 659.95 C555.01 653.16 554.71 643.53 556.71 639.61 C558.71 635.68 561.42 635.35 565.59 636.42 C569.76 637.48 578.69 647.07 581.73 645.99 C584.77 644.92 581.95 633.27 583.83 629.97 C585.71 626.67 589.96 627.54 593.02 626.20 C596.09 624.87 599.64 623.19 602.20 621.94 C604.76 620.70 606.32 619.85 608.37 618.75 C610.42 617.64 612.47 616.49 614.51 615.31 C616.56 614.12 614.08 609.06 620.63 611.63 C627.19 614.20 648.17 629.24 653.87 630.75 C659.57 632.26 654.33 623.82 654.85 620.70 C655.37 617.58 656.35 614.92 656.98 612.04 C657.62 609.16 658.48 606.46 658.66 603.42 C658.85 600.37 652.03 593.56 658.08 593.78 C664.14 594.00 687.87 604.18 695.01 604.75 C702.15 605.32 699.02 599.79 700.92 597.22 C702.83 594.64 704.67 592.01 706.44 589.32 C708.22 586.64 709.58 584.59 711.58 581.10 C713.57 577.62 716.30 572.75 718.42 568.42 C720.55 564.09 724.91 560.08 724.31 555.14 C723.71 550.21 713.43 543.47 714.82 538.80 C716.21 534.13 729.36 531.49 732.64 527.12 C735.92 522.74 737.16 517.40 734.52 512.56 C731.88 507.73 722.84 502.55 716.80 498.11 C710.75 493.67 701.45 490.02 698.25 485.92 C695.05 481.81 695.46 478.13 697.61 473.48 C699.77 468.83 711.46 462.24 711.16 458.03 C710.87 453.82 698.68 451.99 695.85 448.22 C693.03 444.45 694.87 439.70 694.21 435.42 C693.55 431.14 692.81 426.83 691.90 422.55 C690.99 418.27 689.96 413.96 688.73 409.72 C687.50 405.47 688.14 399.92 684.53 397.08 C680.92 394.24 669.02 396.70 667.06 392.68 C665.10 388.66 673.14 378.08 672.79 372.96 C672.44 367.85 669.95 363.43 664.98 362.01 C660.01 360.60 650.13 363.29 642.96 364.48 C635.79 365.66 626.79 369.65 621.94 369.11 C617.10 368.57 615.18 365.87 613.89 361.26 C612.61 356.66 616.97 344.10 614.21 341.46 C611.45 338.82 601.62 346.05 597.33 345.44 C593.03 344.84 591.48 340.32 588.43 337.84 C585.38 335.36 582.27 332.90 579.04 330.56 C575.81 328.23 572.50 325.95 569.08 323.84 C565.66 321.73 562.71 317.37 558.52 317.89 C554.32 318.41 547.69 328.45 543.90 326.96 C540.12 325.46 539.16 312.33 535.82 308.91 C532.47 305.49 528.02 304.02 523.82 306.44 C519.62 308.85 514.64 317.61 510.61 323.40 C506.58 329.19 503.12 338.26 499.61 341.18 C496.11 344.10 493.27 343.42 489.58 340.91 C485.89 338.41 480.91 326.22 477.49 326.13 C474.07 326.04 472.22 337.96 469.06 340.38 C465.90 342.80 462.07 340.47 458.54 340.65 C455.01 340.82 451.43 341.04 447.87 341.42 C444.30 341.80 440.69 342.26 437.12 342.91 C433.55 343.56 428.80 342.24 426.44 345.31 C424.09 348.38 426.46 360.05 423.02 361.35 C419.58 362.65 410.26 353.52 405.82 353.10 C401.37 352.68 397.45 354.42 396.35 358.82 C395.25 363.21 397.91 372.74 399.22 379.46 C400.53 386.18 404.50 394.89 404.22 399.14 C403.95 403.38 401.69 404.54 397.55 404.94 C393.42 405.33 381.71 399.55 379.41 401.51 C377.10 403.47 384.13 413.11 383.71 416.70 C383.29 420.30 379.12 420.88 376.87 423.08 C374.62 425.28 372.04 427.94 370.20 429.89 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M602.20 621.94 C604.76 620.70 606.32 619.85 608.37 618.75 C610.42 617.64 612.47 616.49 614.51 615.31 C616.56 614.12 614.08 609.06 620.63 611.63 C627.19 614.20 648.17 629.24 653.87 630.75 C659.57 632.26 654.33 623.82 654.85 620.70 C655.37 617.58 656.35 614.92 656.98 612.04 C657.62 609.16 658.48 606.46 658.66 603.42 C658.85 600.37 652.03 593.56 658.08 593.78 C664.14 594.00 687.87 604.18 695.01 604.75 C702.15 605.32 699.02 599.79 700.92 597.22 C702.83 594.64 704.67 592.01 706.44 589.32 C708.22 586.64 709.58 584.59 711.58 581.10 C713.57 577.62 716.30 572.75 718.42 568.42 C720.55 564.09 724.91 560.08 724.31 555.14 C723.71 550.21 713.43 543.47 714.82 538.80 C716.21 534.13 729.36 531.49 732.64 527.12 C735.92 522.74 737.16 517.40 734.52 512.56 C731.88 507.73 722.84 502.55 716.80 498.11 C710.75 493.67 701.45 490.02 698.25 485.92 C695.05 481.81 695.46 478.13 697.61 473.48 C699.77 468.83 711.46 462.24 711.16 458.03 C710.87 453.82 698.68 451.99 695.85 448.22 C693.03 444.45 694.87 439.70 694.21 435.42 C693.55 431.14 692.81 426.83 691.90 422.55 C690.99 418.27 689.96 413.96 688.73 409.72 C687.50 405.47 688.14 399.92 684.53 397.08 C680.92 394.24 669.02 396.70 667.06 392.68 C665.10 388.66 673.14 378.08 672.79 372.96 C672.44 367.85 669.95 363.43 664.98 362.01 C660.01 360.60 650.13 363.29 642.96 364.48 C635.79 365.66 626.79 369.65 621.94 369.11 C617.10 368.57 615.18 365.87 613.89 361.26 C612.61 356.66 616.97 344.10 614.21 341.46 C611.45 338.82 601.62 346.05 597.33 345.44 C593.03 344.84 591.48 340.32 588.43 337.84 C585.38 335.36 582.27 332.90 579.04 330.56 C575.81 328.23 572.50 325.95 569.08 323.84 C565.66 321.73 562.71 317.37 558.52 317.89 C554.32 318.41 547.69 328.45 543.90 326.96 C540.12 325.46 539.16 312.33 535.82 308.91 C532.47 305.49 528.02 304.02 523.82 306.44 C519.62 308.85 514.64 317.61 510.61 323.40 C506.58 329.19 503.12 338.26 499.61 341.18 C496.11 344.10 493.27 343.42 489.58 340.91 C485.89 338.41 480.91 326.22 477.49 326.13 C474.07 326.04 472.22 337.96 469.06 340.38 C465.90 342.80 462.07 340.47 458.54 340.65 C455.01 340.82 450.86 341.15 447.87 341.42 C444.87 341.69 443.01 341.94 440.57 342.27 C438.12 342.59 435.67 342.96 433.21 343.37 C430.75 343.78 431.07 349.46 425.80 344.72 C420.52 339.97 406.43 318.33 401.56 314.89 C396.69 311.46 398.34 321.31 396.60 324.09 C394.85 326.87 392.85 329.07 391.08 331.59 C389.31 334.11 387.41 336.39 386.01 339.23 C384.61 342.08 388.33 351.01 382.67 348.66 C377.02 346.32 358.59 328.22 352.05 325.15 C345.52 322.09 346.29 328.48 343.47 330.27 C340.66 332.06 337.87 333.93 335.14 335.88 C332.40 337.82 330.31 339.30 327.06 341.94 C323.81 344.58 319.32 348.29 315.63 351.70 C311.93 355.10 306.27 357.40 304.90 362.38 C303.52 367.36 310.50 377.57 307.38 381.58 C304.26 385.59 290.96 383.40 286.20 386.44 C281.44 389.48 278.23 394.18 278.83 399.83 C279.42 405.49 285.85 413.81 289.77 420.35 C293.68 426.89 300.94 433.89 302.34 439.07 C303.75 444.26 301.96 447.73 298.20 451.46 C294.44 455.19 281.09 457.23 279.78 461.46 C278.47 465.68 289.11 472.06 290.34 476.80 C291.56 481.54 288.08 485.49 287.11 489.91 C286.14 494.33 285.24 498.81 284.51 503.33 C283.79 507.86 283.17 512.44 282.77 517.04 C282.37 521.64 279.76 526.77 282.10 530.94 C284.45 535.10 296.46 537.39 296.84 542.04 C297.22 546.69 285.90 553.78 284.40 558.86 C282.90 563.93 283.65 569.16 287.82 572.49 C291.99 575.83 302.24 577.13 309.43 578.85 C316.61 580.57 326.53 580.36 330.93 582.81 C335.33 585.27 336.21 588.67 335.81 593.60 C335.42 598.54 326.85 608.78 328.56 612.42 C330.27 616.05 342.18 613.11 346.07 615.41 C349.97 617.71 349.88 622.65 351.93 626.23 C353.99 629.81 356.12 633.39 358.40 636.89 C360.69 640.39 363.08 643.87 365.63 647.22 C368.19 650.58 369.48 655.84 373.73 657.03 C377.98 658.21 388.00 651.44 391.13 654.33 C394.27 657.23 390.50 669.91 392.53 674.42 C394.56 678.94 398.34 682.03 403.31 681.43 C408.28 680.82 416.30 674.60 422.36 670.78 C428.42 666.96 435.15 659.87 439.67 658.52 C444.19 657.16 446.77 658.91 449.47 662.65 C452.17 666.40 452.53 679.62 455.87 681.01 C459.21 682.39 465.51 672.03 469.53 670.98 C473.56 669.92 476.46 673.53 480.01 674.69 C483.56 675.85 487.16 676.97 490.82 677.94 C494.48 678.92 498.20 679.81 501.96 680.52 C505.72 681.23 509.86 684.20 513.37 682.19 C516.89 680.18 519.17 668.44 523.04 668.46 C526.92 668.49 532.43 680.37 536.60 682.36 C540.78 684.35 545.26 684.12 548.08 680.39 C550.91 676.65 552.14 666.75 553.57 659.95 C555.01 653.16 554.71 643.53 556.71 639.61 C558.71 635.68 561.42 635.35 565.59 636.42 C569.76 637.48 578.69 647.07 581.73 645.99 C584.77 644.92 581.95 633.27 583.83 629.97 C585.71 626.67 589.96 627.54 593.02 626.20 C596.09 624.87 599.64 623.19 602.20 621.94 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M447.87 341.42 C444.87 341.69 443.01 341.94 440.57 342.27 C438.12 342.59 435.67 342.96 433.21 343.37 C430.75 343.78 431.07 349.46 425.80 344.72 C420.52 339.97 406.43 318.33 401.56 314.89 C396.69 311.46 398.34 321.31 396.60 324.09 C394.85 326.87 392.85 329.07 391.08 331.59 C389.31 334.11 387.41 336.39 386.01 339.23 C384.61 342.08 388.33 351.01 382.67 348.66 C377.02 346.32 358.59 328.22 352.05 325.15 C345.52 322.09 346.29 328.48 343.47 330.27 C340.66 332.06 337.87 333.93 335.14 335.88 C332.40 337.82 330.31 339.30 327.06 341.94 C323.81 344.58 319.32 348.29 315.63 351.70 C311.93 355.10 306.27 357.40 304.90 362.38 C303.52 367.36 310.50 377.57 307.38 381.58 C304.26 385.59 290.96 383.40 286.20 386.44 C281.44 389.48 278.23 394.18 278.83 399.83 C279.42 405.49 285.85 413.81 289.77 420.35 C293.68 426.89 300.94 433.89 302.34 439.07 C303.75 444.26 301.96 447.73 298.20 451.46 C294.44 455.19 281.09 457.23 279.78 461.46 C278.47 465.68 289.11 472.06 290.34 476.80 C291.56 481.54 288.08 485.49 287.11 489.91 C286.14 494.33 285.24 498.81 284.51 503.33 C283.79 507.86 283.17 512.44 282.77 517.04 C282.37 521.64 279.76 526.77 282.10 530.94 C284.45 535.10 296.46 537.39 296.84 542.04 C297.22 546.69 285.90 553.78 284.40 558.86 C282.90 563.93 283.65 569.16 287.82 572.49 C291.99 575.83 302.24 577.13 309.43 578.85 C316.61 580.57 326.53 580.36 330.93 582.81 C335.33 585.27 336.21 588.67 335.81 593.60 C335.42 598.54 326.85 608.78 328.56 612.42 C330.27 616.05 342.18 613.11 346.07 615.41 C349.97 617.71 349.88 622.65 351.93 626.23 C353.99 629.81 356.12 633.39 358.40 636.89 C360.69 640.39 363.08 643.87 365.63 647.22 C368.19 650.58 369.48 655.84 373.73 657.03 C377.98 658.21 388.00 651.44 391.13 654.33 C394.27 657.23 390.50 669.91 392.53 674.42 C394.56 678.94 398.34 682.03 403.31 681.43 C408.28 680.82 416.30 674.60 422.36 670.78 C428.42 666.96 435.15 659.87 439.67 658.52 C444.19 657.16 446.77 658.91 449.47 662.65 C452.17 666.40 452.53 679.62 455.87 681.01 C459.21 682.39 465.51 672.03 469.53 670.98 C473.56 669.92 476.46 673.53 480.01 674.69 C483.56 675.85 487.77 677.08 490.82 677.94 C493.87 678.80 495.80 679.26 498.33 679.85 C500.85 680.45 503.40 681.01 505.98 681.53 C508.55 682.05 510.48 676.63 513.77 682.98 C517.07 689.32 522.37 714.61 525.77 719.58 C529.17 724.55 531.39 714.76 534.16 712.79 C536.94 710.83 539.75 709.50 542.45 707.78 C545.15 706.05 547.89 704.61 550.36 702.44 C552.84 700.28 552.79 690.56 557.32 694.79 C561.85 699.01 572.47 722.58 577.54 727.79 C582.62 733.00 584.37 726.72 587.77 726.04 C591.18 725.36 594.57 724.59 597.95 723.73 C601.33 722.87 603.92 722.22 608.06 720.89 C612.20 719.56 617.94 717.64 622.79 715.73 C627.65 713.82 633.92 713.66 637.19 709.42 C640.46 705.19 637.89 693.01 642.42 690.30 C646.95 687.60 658.66 694.40 664.35 693.18 C670.04 691.97 674.89 688.64 676.56 683.02 C678.22 677.41 675.45 667.15 674.35 659.48 C673.24 651.81 669.21 642.48 669.93 636.98 C670.65 631.48 673.71 628.74 678.69 626.49 C683.67 624.23 696.92 627.07 699.80 623.45 C702.67 619.83 695.23 609.81 695.93 604.77 C696.63 599.72 701.40 597.15 704.00 593.20 C706.61 589.26 709.17 585.24 711.58 581.10 C713.98 576.97 716.30 572.75 718.42 568.42 C720.55 564.09 724.91 560.08 724.31 555.14 C723.71 550.21 713.43 543.47 714.82 538.80 C716.21 534.13 729.36 531.49 732.64 527.12 C735.92 522.74 737.16 517.40 734.52 512.56 C731.88 507.73 722.84 502.55 716.80 498.11 C710.75 493.67 701.45 490.02 698.25 485.92 C695.05 481.81 695.46 478.13 697.61 473.48 C699.77 468.83 711.46 462.24 711.16 458.03 C710.87 453.82 698.68 451.99 695.85 448.22 C693.03 444.45 694.87 439.70 694.21 435.42 C693.55 431.14 692.81 426.83 691.90 422.55 C690.99 418.27 689.96 413.96 688.73 409.72 C687.50 405.47 688.14 399.92 684.53 397.08 C680.92 394.24 669.02 396.70 667.06 392.68 C665.10 388.66 673.14 378.08 672.79 372.96 C672.44 367.85 669.95 363.43 664.98 362.01 C660.01 360.60 650.13 363.29 642.96 364.48 C635.79 365.66 626.79 369.65 621.94 369.11 C617.10 368.57 615.18 365.87 613.89 361.26 C612.61 356.66 616.97 344.10 614.21 341.46 C611.45 338.82 601.62 346.05 597.33 345.44 C593.03 344.84 591.48 340.32 588.43 337.84 C585.38 335.36 582.27 332.90 579.04 330.56 C575.81 328.23 572.50 325.95 569.08 323.84 C565.66 321.73 562.71 317.37 558.52 317.89 C554.32 318.41 547.69 328.45 543.90 326.96 C540.12 325.46 539.16 312.33 535.82 308.91 C532.47 305.49 528.02 304.02 523.82 306.44 C519.62 308.85 514.64 317.61 510.61 323.40 C506.58 329.19 503.12 338.26 499.61 341.18 C496.11 344.10 493.27 343.42 489.58 340.91 C485.89 338.41 480.91 326.22 477.49 326.13 C474.07 326.04 472.22 337.96 469.06 340.38 C465.90 342.80 462.07 340.47 458.54 340.65 C455.01 340.82 450.86 341.15 447.87 341.42 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M490.82 677.94 C493.87 678.80 495.80 679.26 498.33 679.85 C500.85 680.45 503.40 681.01 505.98 681.53 C508.55 682.05 510.48 676.63 513.77 682.98 C517.07 689.32 522.37 714.61 525.77 719.58 C529.17 724.55 531.39 714.76 534.16 712.79 C536.94 710.83 539.75 709.50 542.45 707.78 C545.15 706.05 547.89 704.61 550.36 702.44 C552.84 700.28 552.79 690.56 557.32 694.79 C561.85 699.01 572.47 722.58 577.54 727.79 C582.62 733.00 584.37 726.72 587.77 726.04 C591.18 725.36 594.57 724.59 597.95 723.73 C601.33 722.87 603.92 722.22 608.06 720.89 C612.20 719.56 617.94 717.64 622.79 715.73 C627.65 713.82 633.92 713.66 637.19 709.42 C640.46 705.19 637.89 693.01 642.42 690.30 C646.95 687.60 658.66 694.40 664.35 693.18 C670.04 691.97 674.89 688.64 676.56 683.02 C678.22 677.41 675.45 667.15 674.35 659.48 C673.24 651.81 669.21 642.48 669.93 636.98 C670.65 631.48 673.71 628.74 678.69 626.49 C683.67 624.23 696.92 627.07 699.80 623.45 C702.67 619.83 695.23 609.81 695.93 604.77 C696.63 599.72 701.40 597.15 704.00 593.20 C706.61 589.26 709.17 585.24 711.58 581.10 C713.98 576.97 716.30 572.75 718.42 568.42 C720.55 564.09 724.91 560.08 724.31 555.14 C723.71 550.21 713.43 543.47 714.82 538.80 C716.21 534.13 729.36 531.49 732.64 527.12 C735.92 522.74 737.16 517.40 734.52 512.56 C731.88 507.73 722.84 502.55 716.80 498.11 C710.75 493.67 701.45 490.02 698.25 485.92 C695.05 481.81 695.46 478.13 697.61 473.48 C699.77 468.83 711.46 462.24 711.16 458.03 C710.87 453.82 698.68 451.99 695.85 448.22 C693.03 444.45 694.87 439.70 694.21 435.42 C693.55 431.14 692.81 426.83 691.90 422.55 C690.99 418.27 689.96 413.96 688.73 409.72 C687.50 405.47 688.14 399.92 684.53 397.08 C680.92 394.24 669.02 396.70 667.06 392.68 C665.10 388.66 673.14 378.08 672.79 372.96 C672.44 367.85 669.95 363.43 664.98 362.01 C660.01 360.60 650.13 363.29 642.96 364.48 C635.79 365.66 626.79 369.65 621.94 369.11 C617.10 368.57 615.18 365.87 613.89 361.26 C612.61 356.66 616.97 344.10 614.21 341.46 C611.45 338.82 601.62 346.05 597.33 345.44 C593.03 344.84 591.48 340.32 588.43 337.84 C585.38 335.36 581.71 332.56 579.04 330.56 C576.37 328.57 574.65 327.40 572.39 325.87 C570.14 324.33 567.84 322.83 565.51 321.35 C563.18 319.87 559.28 324.16 558.41 317.00 C557.55 309.85 561.78 284.31 560.32 278.40 C558.85 272.48 553.07 280.75 549.63 281.53 C546.19 282.31 542.96 282.49 539.69 283.08 C536.41 283.68 533.20 283.99 529.96 285.08 C526.73 286.17 523.08 295.22 520.28 289.61 C517.49 284.00 516.14 258.14 513.20 251.41 C510.26 244.69 506.19 249.89 502.64 249.27 C499.09 248.65 495.52 248.12 491.92 247.69 C488.33 247.25 485.59 246.91 481.08 246.65 C476.57 246.38 470.30 246.08 464.89 246.10 C459.48 246.13 453.39 244.01 448.59 246.80 C443.80 249.58 441.51 261.92 436.11 262.83 C430.72 263.74 422.13 253.16 416.21 252.27 C410.28 251.37 404.36 252.77 400.56 257.46 C396.76 262.16 395.38 272.81 393.40 280.46 C391.43 288.10 391.57 298.36 388.70 303.32 C385.83 308.27 381.81 309.80 376.18 310.19 C370.55 310.57 359.10 303.18 354.94 305.62 C350.77 308.06 353.85 320.25 351.19 324.81 C348.53 329.38 343.00 330.17 338.98 333.02 C334.95 335.88 330.95 338.83 327.06 341.94 C323.17 345.05 319.32 348.29 315.63 351.70 C311.93 355.10 306.27 357.40 304.90 362.38 C303.52 367.36 310.50 377.57 307.38 381.58 C304.26 385.59 290.96 383.40 286.20 386.44 C281.44 389.48 278.23 394.18 278.83 399.83 C279.42 405.49 285.85 413.81 289.77 420.35 C293.68 426.89 300.94 433.89 302.34 439.07 C303.75 444.26 301.96 447.73 298.20 451.46 C294.44 455.19 281.09 457.23 279.78 461.46 C278.47 465.68 289.11 472.06 290.34 476.80 C291.56 481.54 288.08 485.49 287.11 489.91 C286.14 494.33 285.24 498.81 284.51 503.33 C283.79 507.86 283.17 512.44 282.77 517.04 C282.37 521.64 279.76 526.77 282.10 530.94 C284.45 535.10 296.46 537.39 296.84 542.04 C297.22 546.69 285.90 553.78 284.40 558.86 C282.90 563.93 283.65 569.16 287.82 572.49 C291.99 575.83 302.24 577.13 309.43 578.85 C316.61 580.57 326.53 580.36 330.93 582.81 C335.33 585.27 336.21 588.67 335.81 593.60 C335.42 598.54 326.85 608.78 328.56 612.42 C330.27 616.05 342.18 613.11 346.07 615.41 C349.97 617.71 349.88 622.65 351.93 626.23 C353.99 629.81 356.12 633.39 358.40 636.89 C360.69 640.39 363.08 643.87 365.63 647.22 C368.19 650.58 369.48 655.84 373.73 657.03 C377.98 658.21 388.00 651.44 391.13 654.33 C394.27 657.23 390.50 669.91 392.53 674.42 C394.56 678.94 398.34 682.03 403.31 681.43 C408.28 680.82 416.30 674.60 422.36 670.78 C428.42 666.96 435.15 659.87 439.67 658.52 C444.19 657.16 446.77 658.91 449.47 662.65 C452.17 666.40 452.53 679.62 455.87 681.01 C459.21 682.39 465.51 672.03 469.53 670.98 C473.56 669.92 476.46 673.53 480.01 674.69 C483.56 675.85 487.77 677.08 490.82 677.94 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M579.04 330.56 C576.37 328.57 574.65 327.40 572.39 325.87 C570.14 324.33 567.84 322.83 565.51 321.35 C563.18 319.87 559.28 324.16 558.41 317.00 C557.55 309.85 561.78 284.31 560.32 278.40 C558.85 272.48 553.07 280.75 549.63 281.53 C546.19 282.31 542.96 282.49 539.69 283.08 C536.41 283.68 533.20 283.99 529.96 285.08 C526.73 286.17 523.08 295.22 520.28 289.61 C517.49 284.00 516.14 258.14 513.20 251.41 C510.26 244.69 506.19 249.89 502.64 249.27 C499.09 248.65 495.52 248.12 491.92 247.69 C488.33 247.25 485.59 246.91 481.08 246.65 C476.57 246.38 470.30 246.08 464.89 246.10 C459.48 246.13 453.39 244.01 448.59 246.80 C443.80 249.58 441.51 261.92 436.11 262.83 C430.72 263.74 422.13 253.16 416.21 252.27 C410.28 251.37 404.36 252.77 400.56 257.46 C396.76 262.16 395.38 272.81 393.40 280.46 C391.43 288.10 391.57 298.36 388.70 303.32 C385.83 308.27 381.81 309.80 376.18 310.19 C370.55 310.57 359.10 303.18 354.94 305.62 C350.77 308.06 353.85 320.25 351.19 324.81 C348.53 329.38 343.00 330.17 338.98 333.02 C334.95 335.88 330.95 338.83 327.06 341.94 C323.17 345.05 319.32 348.29 315.63 351.70 C311.93 355.10 306.27 357.40 304.90 362.38 C303.52 367.36 310.50 377.57 307.38 381.58 C304.26 385.59 290.96 383.40 286.20 386.44 C281.44 389.48 278.23 394.18 278.83 399.83 C279.42 405.49 285.85 413.81 289.77 420.35 C293.68 426.89 300.94 433.89 302.34 439.07 C303.75 444.26 301.96 447.73 298.20 451.46 C294.44 455.19 281.09 457.23 279.78 461.46 C278.47 465.68 289.11 472.06 290.34 476.80 C291.56 481.54 288.08 485.49 287.11 489.91 C286.14 494.33 285.24 498.81 284.51 503.33 C283.79 507.86 283.17 512.44 282.77 517.04 C282.37 521.64 279.76 526.77 282.10 530.94 C284.45 535.10 296.46 537.39 296.84 542.04 C297.22 546.69 285.90 553.78 284.40 558.86 C282.90 563.93 283.65 569.16 287.82 572.49 C291.99 575.83 302.24 577.13 309.43 578.85 C316.61 580.57 326.53 580.36 330.93 582.81 C335.33 585.27 336.21 588.67 335.81 593.60 C335.42 598.54 326.85 608.78 328.56 612.42 C330.27 616.05 342.18 613.11 346.07 615.41 C349.97 617.71 349.88 622.65 351.93 626.23 C353.99 629.81 356.53 633.93 358.40 636.89 C360.27 639.85 361.52 641.64 363.15 643.99 C364.78 646.34 366.46 648.67 368.19 650.99 C369.92 653.31 375.23 650.84 373.51 657.90 C371.80 664.96 358.59 687.23 357.90 693.35 C357.20 699.46 365.73 693.98 369.32 694.59 C372.92 695.21 376.10 696.30 379.49 697.03 C382.87 697.75 386.09 698.71 389.62 698.95 C393.14 699.19 399.98 692.18 400.64 698.49 C401.30 704.79 393.19 729.39 393.58 736.77 C393.98 744.16 399.80 740.87 403.00 742.81 C406.20 744.74 409.46 746.59 412.77 748.36 C416.08 750.13 418.61 751.49 422.86 753.44 C427.12 755.39 433.06 758.04 438.30 760.05 C443.55 762.07 448.63 766.32 454.31 765.52 C459.99 764.73 466.84 754.13 472.39 755.29 C477.94 756.45 482.25 769.45 487.63 772.47 C493.01 775.49 499.24 776.38 504.69 773.40 C510.15 770.43 515.54 761.03 520.38 754.63 C525.22 748.24 529.05 738.62 533.73 735.04 C538.42 731.46 542.92 731.48 548.48 733.15 C554.05 734.82 562.18 745.85 567.14 745.06 C572.10 744.27 573.90 731.75 578.24 728.41 C582.57 725.08 588.19 726.31 593.16 725.05 C598.13 723.80 603.12 722.44 608.06 720.89 C613.00 719.34 617.94 717.64 622.79 715.73 C627.65 713.82 633.92 713.66 637.19 709.42 C640.46 705.19 637.89 693.01 642.42 690.30 C646.95 687.60 658.66 694.40 664.35 693.18 C670.04 691.97 674.89 688.64 676.56 683.02 C678.22 677.41 675.45 667.15 674.35 659.48 C673.24 651.81 669.21 642.48 669.93 636.98 C670.65 631.48 673.71 628.74 678.69 626.49 C683.67 624.23 696.92 627.07 699.80 623.45 C702.67 619.83 695.23 609.81 695.93 604.77 C696.63 599.72 701.40 597.15 704.00 593.20 C706.61 589.26 709.17 585.24 711.58 581.10 C713.98 576.97 716.30 572.75 718.42 568.42 C720.55 564.09 724.91 560.08 724.31 555.14 C723.71 550.21 713.43 543.47 714.82 538.80 C716.21 534.13 729.36 531.49 732.64 527.12 C735.92 522.74 737.16 517.40 734.52 512.56 C731.88 507.73 722.84 502.55 716.80 498.11 C710.75 493.67 701.45 490.02 698.25 485.92 C695.05 481.81 695.46 478.13 697.61 473.48 C699.77 468.83 711.46 462.24 711.16 458.03 C710.87 453.82 698.68 451.99 695.85 448.22 C693.03 444.45 694.87 439.70 694.21 435.42 C693.55 431.14 692.81 426.83 691.90 422.55 C690.99 418.27 689.96 413.96 688.73 409.72 C687.50 405.47 688.14 399.92 684.53 397.08 C680.92 394.24 669.02 396.70 667.06 392.68 C665.10 388.66 673.14 378.08 672.79 372.96 C672.44 367.85 669.95 363.43 664.98 362.01 C660.01 360.60 650.13 363.29 642.96 364.48 C635.79 365.66 626.79 369.65 621.94 369.11 C617.10 368.57 615.18 365.87 613.89 361.26 C612.61 356.66 616.97 344.10 614.21 341.46 C611.45 338.82 601.62 346.05 597.33 345.44 C593.03 344.84 591.48 340.32 588.43 337.84 C585.38 335.36 581.71 332.56 579.04 330.56 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M358.40 636.89 C360.27 639.85 361.52 641.64 363.15 643.99 C364.78 646.34 366.46 648.67 368.19 650.99 C369.92 653.31 375.23 650.84 373.51 657.90 C371.80 664.96 358.59 687.23 357.90 693.35 C357.20 699.46 365.73 693.98 369.32 694.59 C372.92 695.21 376.10 696.30 379.49 697.03 C382.87 697.75 386.09 698.71 389.62 698.95 C393.14 699.19 399.98 692.18 400.64 698.49 C401.30 704.79 393.19 729.39 393.58 736.77 C393.98 744.16 399.80 740.87 403.00 742.81 C406.20 744.74 409.46 746.59 412.77 748.36 C416.08 750.13 418.61 751.49 422.86 753.44 C427.12 755.39 433.06 758.04 438.30 760.05 C443.55 762.07 448.63 766.32 454.31 765.52 C459.99 764.73 466.84 754.13 472.39 755.29 C477.94 756.45 482.25 769.45 487.63 772.47 C493.01 775.49 499.24 776.38 504.69 773.40 C510.15 770.43 515.54 761.03 520.38 754.63 C525.22 748.24 529.05 738.62 533.73 735.04 C538.42 731.46 542.92 731.48 548.48 733.15 C554.05 734.82 562.18 745.85 567.14 745.06 C572.10 744.27 573.90 731.75 578.24 728.41 C582.57 725.08 588.19 726.31 593.16 725.05 C598.13 723.80 603.12 722.44 608.06 720.89 C613.00 719.34 617.94 717.64 622.79 715.73 C627.65 713.82 633.92 713.66 637.19 709.42 C640.46 705.19 637.89 693.01 642.42 690.30 C646.95 687.60 658.66 694.40 664.35 693.18 C670.04 691.97 674.89 688.64 676.56 683.02 C678.22 677.41 675.45 667.15 674.35 659.48 C673.24 651.81 669.21 642.48 669.93 636.98 C670.65 631.48 673.71 628.74 678.69 626.49 C683.67 624.23 696.92 627.07 699.80 623.45 C702.67 619.83 695.23 609.81 695.93 604.77 C696.63 599.72 701.40 597.15 704.00 593.20 C706.61 589.26 709.17 585.24 711.58 581.10 C713.98 576.97 716.30 572.75 718.42 568.42 C720.55 564.09 724.91 560.08 724.31 555.14 C723.71 550.21 713.43 543.47 714.82 538.80 C716.21 534.13 729.36 531.49 732.64 527.12 C735.92 522.74 737.16 517.40 734.52 512.56 C731.88 507.73 722.84 502.55 716.80 498.11 C710.75 493.67 701.45 490.02 698.25 485.92 C695.05 481.81 695.46 478.13 697.61 473.48 C699.77 468.83 711.46 462.24 711.16 458.03 C710.87 453.82 698.68 451.99 695.85 448.22 C693.03 444.45 694.87 439.70 694.21 435.42 C693.55 431.14 692.61 426.15 691.90 422.55 C691.19 418.95 690.64 416.73 689.93 413.82 C689.22 410.91 688.46 408.00 687.64 405.08 C686.83 402.17 680.92 402.39 685.04 396.33 C689.16 390.27 709.54 374.26 712.36 368.71 C715.18 363.16 705.17 365.01 701.98 363.01 C698.78 361.00 696.15 358.71 693.19 356.68 C690.24 354.65 687.51 352.47 684.24 350.84 C680.96 349.22 671.95 353.13 673.56 346.91 C675.17 340.70 691.64 320.67 693.90 313.55 C696.15 306.43 689.44 307.26 687.07 304.19 C684.71 301.12 682.26 298.11 679.73 295.15 C677.21 292.19 675.28 289.93 671.91 286.45 C668.54 282.97 663.80 278.17 659.50 274.25 C655.21 270.34 651.88 264.40 646.14 262.95 C640.39 261.50 629.97 268.75 625.05 265.54 C620.14 262.32 620.74 248.54 616.66 243.67 C612.58 238.80 606.94 235.61 600.59 236.30 C594.24 236.99 585.58 243.71 578.55 247.83 C571.51 251.94 564.22 259.44 558.36 261.01 C552.49 262.58 548.12 260.86 543.36 257.23 C538.60 253.59 534.90 240.30 529.80 239.19 C524.70 238.08 518.23 249.07 512.77 250.58 C507.31 252.08 502.34 248.87 497.06 248.21 C491.78 247.56 486.44 247.00 481.08 246.65 C475.72 246.30 470.30 246.08 464.89 246.10 C459.48 246.13 453.39 244.01 448.59 246.80 C443.80 249.58 441.51 261.92 436.11 262.83 C430.72 263.74 422.13 253.16 416.21 252.27 C410.28 251.37 404.36 252.77 400.56 257.46 C396.76 262.16 395.38 272.81 393.40 280.46 C391.43 288.10 391.57 298.36 388.70 303.32 C385.83 308.27 381.81 309.80 376.18 310.19 C370.55 310.57 359.10 303.18 354.94 305.62 C350.77 308.06 353.85 320.25 351.19 324.81 C348.53 329.38 343.00 330.17 338.98 333.02 C334.95 335.88 330.95 338.83 327.06 341.94 C323.17 345.05 319.32 348.29 315.63 351.70 C311.93 355.10 306.27 357.40 304.90 362.38 C303.52 367.36 310.50 377.57 307.38 381.58 C304.26 385.59 290.96 383.40 286.20 386.44 C281.44 389.48 278.23 394.18 278.83 399.83 C279.42 405.49 285.85 413.81 289.77 420.35 C293.68 426.89 300.94 433.89 302.34 439.07 C303.75 444.26 301.96 447.73 298.20 451.46 C294.44 455.19 281.09 457.23 279.78 461.46 C278.47 465.68 289.11 472.06 290.34 476.80 C291.56 481.54 288.08 485.49 287.11 489.91 C286.14 494.33 285.24 498.81 284.51 503.33 C283.79 507.86 283.17 512.44 282.77 517.04 C282.37 521.64 279.76 526.77 282.10 530.94 C284.45 535.10 296.46 537.39 296.84 542.04 C297.22 546.69 285.90 553.78 284.40 558.86 C282.90 563.93 283.65 569.16 287.82 572.49 C291.99 575.83 302.24 577.13 309.43 578.85 C316.61 580.57 326.53 580.36 330.93 582.81 C335.33 585.27 336.21 588.67 335.81 593.60 C335.42 598.54 326.85 608.78 328.56 612.42 C330.27 616.05 342.18 613.11 346.07 615.41 C349.97 617.71 349.88 622.65 351.93 626.23 C353.99 629.81 356.53 633.93 358.40 636.89 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M691.90 422.55 C691.19 418.95 690.64 416.73 689.93 413.82 C689.22 410.91 688.46 408.00 687.64 405.08 C686.83 402.17 680.92 402.39 685.04 396.33 C689.16 390.27 709.54 374.26 712.36 368.71 C715.18 363.16 705.17 365.01 701.98 363.01 C698.78 361.00 696.15 358.71 693.19 356.68 C690.24 354.65 687.51 352.47 684.24 350.84 C680.96 349.22 671.95 353.13 673.56 346.91 C675.17 340.70 691.64 320.67 693.90 313.55 C696.15 306.43 689.44 307.26 687.07 304.19 C684.71 301.12 682.26 298.11 679.73 295.15 C677.21 292.19 675.28 289.93 671.91 286.45 C668.54 282.97 663.80 278.17 659.50 274.25 C655.21 270.34 651.88 264.40 646.14 262.95 C640.39 261.50 629.97 268.75 625.05 265.54 C620.14 262.32 620.74 248.54 616.66 243.67 C612.58 238.80 606.94 235.61 600.59 236.30 C594.24 236.99 585.58 243.71 578.55 247.83 C571.51 251.94 564.22 259.44 558.36 261.01 C552.49 262.58 548.12 260.86 543.36 257.23 C538.60 253.59 534.90 240.30 529.80 239.19 C524.70 238.08 518.23 249.07 512.77 250.58 C507.31 252.08 502.34 248.87 497.06 248.21 C491.78 247.56 486.44 247.00 481.08 246.65 C475.72 246.30 470.30 246.08 464.89 246.10 C459.48 246.13 453.39 244.01 448.59 246.80 C443.80 249.58 441.51 261.92 436.11 262.83 C430.72 263.74 422.13 253.16 416.21 252.27 C410.28 251.37 404.36 252.77 400.56 257.46 C396.76 262.16 395.38 272.81 393.40 280.46 C391.43 288.10 391.57 298.36 388.70 303.32 C385.83 308.27 381.81 309.80 376.18 310.19 C370.55 310.57 359.10 303.18 354.94 305.62 C350.77 308.06 353.85 320.25 351.19 324.81 C348.53 329.38 343.00 330.17 338.98 333.02 C334.95 335.88 330.95 338.83 327.06 341.94 C323.17 345.05 319.32 348.29 315.63 351.70 C311.93 355.10 306.27 357.40 304.90 362.38 C303.52 367.36 310.50 377.57 307.38 381.58 C304.26 385.59 290.96 383.40 286.20 386.44 C281.44 389.48 278.23 394.18 278.83 399.83 C279.42 405.49 285.85 413.81 289.77 420.35 C293.68 426.89 300.94 433.89 302.34 439.07 C303.75 444.26 301.96 447.73 298.20 451.46 C294.44 455.19 281.09 457.23 279.78 461.46 C278.47 465.68 289.11 472.06 290.34 476.80 C291.56 481.54 288.08 485.49 287.11 489.91 C286.14 494.33 285.17 499.56 284.51 503.33 C283.85 507.11 283.55 509.48 283.15 512.59 C282.75 515.69 282.40 518.82 282.10 521.96 C281.80 525.11 287.39 527.20 281.35 531.47 C275.31 535.73 250.48 543.26 245.84 547.52 C241.19 551.79 251.22 553.87 253.49 557.04 C255.76 560.21 257.40 563.44 259.44 566.55 C261.48 569.65 263.26 572.81 265.74 575.65 C268.23 578.50 278.11 578.30 274.37 583.62 C270.64 588.93 248.00 601.63 243.33 607.54 C238.67 613.46 245.28 615.29 246.41 619.13 C247.54 622.97 248.77 626.80 250.10 630.60 C251.43 634.39 252.44 637.30 254.39 641.92 C256.33 646.54 259.10 652.94 261.77 658.32 C264.44 663.71 265.47 670.61 270.41 674.24 C275.34 677.86 287.83 675.10 291.35 680.05 C294.88 684.99 289.40 697.73 291.55 703.90 C293.69 710.07 297.92 715.25 304.23 717.07 C310.54 718.90 321.21 715.97 329.43 714.85 C337.65 713.73 347.34 709.56 353.54 710.38 C359.74 711.19 363.35 714.50 366.61 719.73 C369.88 724.96 368.62 738.78 373.13 741.78 C377.64 744.77 387.88 737.03 393.70 737.72 C399.51 738.41 403.14 743.30 408.00 745.92 C412.86 748.54 417.81 751.08 422.86 753.44 C427.91 755.80 433.06 758.04 438.30 760.05 C443.55 762.07 448.63 766.32 454.31 765.52 C459.99 764.73 466.84 754.13 472.39 755.29 C477.94 756.45 482.25 769.45 487.63 772.47 C493.01 775.49 499.24 776.38 504.69 773.40 C510.15 770.43 515.54 761.03 520.38 754.63 C525.22 748.24 529.05 738.62 533.73 735.04 C538.42 731.46 542.92 731.48 548.48 733.15 C554.05 734.82 562.18 745.85 567.14 745.06 C572.10 744.27 573.90 731.75 578.24 728.41 C582.57 725.08 588.19 726.31 593.16 725.05 C598.13 723.80 603.12 722.44 608.06 720.89 C613.00 719.34 617.94 717.64 622.79 715.73 C627.65 713.82 633.92 713.66 637.19 709.42 C640.46 705.19 637.89 693.01 642.42 690.30 C646.95 687.60 658.66 694.40 664.35 693.18 C670.04 691.97 674.89 688.64 676.56 683.02 C678.22 677.41 675.45 667.15 674.35 659.48 C673.24 651.81 669.21 642.48 669.93 636.98 C670.65 631.48 673.71 628.74 678.69 626.49 C683.67 624.23 696.92 627.07 699.80 623.45 C702.67 619.83 695.23 609.81 695.93 604.77 C696.63 599.72 701.40 597.15 704.00 593.20 C706.61 589.26 709.17 585.24 711.58 581.10 C713.98 576.97 716.30 572.75 718.42 568.42 C720.55 564.09 724.91 560.08 724.31 555.14 C723.71 550.21 713.43 543.47 714.82 538.80 C716.21 534.13 729.36 531.49 732.64 527.12 C735.92 522.74 737.16 517.40 734.52 512.56 C731.88 507.73 722.84 502.55 716.80 498.11 C710.75 493.67 701.45 490.02 698.25 485.92 C695.05 481.81 695.46 478.13 697.61 473.48 C699.77 468.83 711.46 462.24 711.16 458.03 C710.87 453.82 698.68 451.99 695.85 448.22 C693.03 444.45 694.87 439.70 694.21 435.42 C693.55 431.14 692.61 426.15 691.90 422.55 Z`

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
	__Parameter__000000_Reference.N = 7
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.InsideAngle = 119.000000
	__Parameter__000000_Reference.ShiftToNearestCircle = 7
	__Parameter__000000_Reference.SideLength = 90.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 1.260000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.099000
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
	__Parameter__000000_Reference.ActualBeatsTemporalShift = 7
	__Parameter__000000_Reference.PathToStaticFiles = `../../../static`
	__Parameter__000000_Reference.PathToGeneratedSVG = `../../../static/images`
	__Parameter__000000_Reference.PathToGeneratedScore = `../../../static/scores`

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 90.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.AngleDegree = -71.879459
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 119.000000
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
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 119.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 323.216145
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 9.647333
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 90.000000
	__Rhombus__000002_Rotated_Next_Rhombus.AngleDegree = -71.879459
	__Rhombus__000002_Rotated_Next_Rhombus.InsideAngle = 119.000000
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
	__Rhombus__000003_Rotated_Rhombus.AngleDegree = -71.879459
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 119.000000
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
	__ShapeCategory__000001_1_Initial.IsExpanded = true

	__ShapeCategory__000002_2_Rotated.Name = `2. Rotated`
	__ShapeCategory__000002_2_Rotated.IsExpanded = false

	__ShapeCategory__000003_3_Growing.Name = `3. Growing`
	__ShapeCategory__000003_3_Growing.IsExpanded = false

	__ShapeCategory__000004_4_Construction.Name = `4. Construction`
	__ShapeCategory__000004_4_Construction.IsExpanded = false

	__ShapeCategory__000005_5_Vertical_growth.Name = `5. Vertical growth`
	__ShapeCategory__000005_5_Vertical_growth.IsExpanded = false

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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 137.788708
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -16.832115
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 159.416167
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 14.430774
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 191.046278
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 77.969873
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 181.180773
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 38.024530
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 137.788708
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 95.400000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -16.832115
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndY = 0.000000
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 137.788708
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 176.867320
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -16.832115
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -43.866439
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 51.603395
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 6.303813
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 96.531817
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 70.471554
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 137.788708
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -16.832115
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 57.572331
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -42.029786
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

