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

	const __write__local_time = "2025-09-16 01:29:14.376533 CEST"
	const __write__utc_time__ = "2025-09-15 23:29:14.376533 UTC"

	const __commitId__ = "0000000124"

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
	__Axis__000001_Construction_Axis.Length = 170.000000
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -19.500337
	__Axis__000001_Construction_Axis.EndY = 168.877876
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
	__Axis__000002_Initial_Axis.Length = 741.012820
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
	__Axis__000004_Rotated_Axis.Length = 741.012820
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -9.750169
	__Bezier__000005_Growth_Curve_Seed.StartY = 84.438938
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 66.244875
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 93.214090
	__Bezier__000005_Growth_Curve_Seed.EndX = 126.752193
	__Bezier__000005_Growth_Curve_Seed.EndY = 185.765663
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 50.757149
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 176.990511
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
	__Circle__000001_Construction_Circle.CenterX = -9.750169
	__Circle__000001_Construction_Circle.CenterY = 84.438938
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 292.505061
	__Circle__000005_Rotated_Next_Circle.CenterY = 33.775575
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M764.56 522.42 C765.81 513.36 766.63 505.48 767.47 496.87 C768.30 488.26 768.93 479.63 769.58 470.78 C770.23 461.93 770.73 453.04 771.36 443.78 C772.00 434.52 772.62 425.16 773.40 415.24 C774.18 405.31 775.24 395.06 776.04 384.25 C776.85 373.43 778.17 361.88 778.21 350.34 C778.25 338.81 778.09 326.58 776.28 315.02 C774.46 303.46 771.45 291.84 767.30 280.98 C763.16 270.11 757.65 259.63 751.42 249.84 C745.20 240.05 737.84 230.80 729.97 222.22 C722.10 213.64 713.34 205.66 704.20 198.35 C695.07 191.05 686.03 184.69 675.15 178.40 C664.27 172.11 651.39 165.53 638.91 160.62 C626.44 155.70 613.40 151.67 600.32 148.91 C587.23 146.15 573.71 144.45 560.40 144.07 C547.10 143.70 533.54 144.57 520.47 146.66 C507.40 148.76 494.36 152.30 481.99 156.64 C469.62 160.98 457.62 166.75 446.26 172.70 C434.89 178.65 424.14 185.59 413.78 192.32 C403.43 199.04 393.68 206.14 384.13 213.06 C374.58 219.97 365.47 226.89 356.49 233.82 C347.51 240.76 338.82 247.63 330.26 254.67 C321.71 261.72 313.34 268.77 305.15 276.09 C296.97 283.41 288.49 291.38 281.13 298.62 C273.78 305.86 267.64 312.38 261.03 319.53 C254.42 326.69 248.00 334.03 241.47 341.55 C234.95 349.06 228.56 356.74 221.88 364.64 C215.20 372.53 208.52 380.55 201.42 388.93 C194.32 397.30 186.84 405.78 179.27 414.87 C171.70 423.97 163.32 433.33 156.01 443.49 C148.71 453.64 141.19 464.47 135.44 475.79 C129.69 487.11 124.91 499.21 121.53 511.41 C118.16 523.61 116.15 536.35 115.18 548.97 C114.21 561.59 114.53 574.49 115.71 587.14 C116.90 599.80 119.20 612.53 122.27 624.93 C125.34 637.33 128.92 648.85 134.13 661.55 C139.33 674.26 146.05 688.54 153.50 701.15 C160.96 713.76 169.46 725.99 178.83 737.20 C188.21 748.41 198.63 759.03 209.74 768.41 C220.86 777.79 232.99 786.30 245.53 793.49 C258.06 800.68 271.52 806.67 284.96 811.55 C298.40 816.43 312.48 819.90 326.19 822.79 C339.90 825.67 353.77 827.33 367.21 828.86 C380.65 830.38 393.87 831.18 406.83 831.93 C419.80 832.67 432.44 833.09 444.99 833.34 C457.55 833.58 469.85 833.65 482.15 833.41 C494.45 833.18 506.62 832.75 518.80 831.93 C530.98 831.11 543.18 830.30 555.23 828.49 C567.29 826.68 579.46 824.50 591.14 821.05 C602.82 817.59 614.41 813.16 625.29 807.76 C636.18 802.36 646.78 795.94 656.43 788.64 C666.08 781.35 675.19 772.99 683.22 764.02 C691.24 755.05 698.39 745.01 704.56 734.81 C710.74 724.61 715.77 713.56 720.27 702.81 C724.76 692.06 728.20 680.95 731.52 670.30 C734.83 659.64 737.50 649.12 740.18 638.88 C742.85 628.65 745.24 618.73 747.58 608.89 C749.91 599.05 752.13 589.48 754.19 579.87 C756.25 570.25 758.22 560.78 759.95 551.21 C761.68 541.63 763.31 531.47 764.56 522.42 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M281.13 298.62 C273.78 305.86 267.64 312.38 261.03 319.53 C254.42 326.69 248.00 334.03 241.47 341.55 C234.95 349.06 228.56 356.74 221.88 364.64 C215.20 372.53 208.52 380.55 201.42 388.93 C194.32 397.30 186.84 405.78 179.27 414.87 C171.70 423.97 163.32 433.33 156.01 443.49 C148.71 453.64 141.19 464.47 135.44 475.79 C129.69 487.11 124.91 499.21 121.53 511.41 C118.16 523.61 116.15 536.35 115.18 548.97 C114.21 561.59 114.53 574.49 115.71 587.14 C116.90 599.80 119.20 612.53 122.27 624.93 C125.34 637.33 128.92 648.85 134.13 661.55 C139.33 674.26 146.05 688.54 153.50 701.15 C160.96 713.76 169.46 725.99 178.83 737.20 C188.21 748.41 198.63 759.03 209.74 768.41 C220.86 777.79 232.99 786.30 245.53 793.49 C258.06 800.68 271.52 806.67 284.96 811.55 C298.40 816.43 312.48 819.90 326.19 822.79 C339.90 825.67 353.77 827.33 367.21 828.86 C380.65 830.38 393.87 831.18 406.83 831.93 C419.80 832.67 432.44 833.09 444.99 833.34 C457.55 833.58 469.85 833.65 482.15 833.41 C494.45 833.18 506.62 832.75 518.80 831.93 C530.98 831.11 543.85 829.87 555.23 828.49 C566.61 827.11 576.44 825.50 587.08 823.63 C597.72 821.76 608.32 819.57 619.09 817.26 C629.86 814.95 640.60 812.38 651.68 809.76 C662.76 807.15 673.90 804.39 685.55 801.56 C697.21 798.74 709.20 796.01 721.61 792.82 C734.02 789.62 747.21 786.67 760.00 782.39 C772.78 778.12 786.12 773.40 798.33 767.18 C810.53 760.95 822.42 753.47 833.22 745.03 C844.02 736.60 854.02 726.85 863.10 716.54 C872.17 706.23 880.36 694.89 887.67 683.18 C894.98 671.47 901.40 659.00 906.97 646.30 C912.54 633.60 917.10 621.33 921.09 606.97 C925.07 592.61 928.82 575.91 930.86 560.14 C932.91 544.38 933.83 528.25 933.37 512.40 C932.90 496.55 931.18 480.52 928.08 465.05 C924.99 449.58 920.48 434.17 914.78 419.57 C909.09 404.98 901.86 390.77 893.89 377.48 C885.93 364.18 876.53 351.62 867.00 339.79 C857.48 327.96 847.00 317.01 836.73 306.48 C826.45 295.94 815.84 286.14 805.35 276.59 C794.86 267.03 784.36 257.97 773.78 249.15 C763.20 240.33 752.64 231.85 741.87 223.66 C731.10 215.47 720.27 207.57 709.15 200.03 C698.03 192.48 686.85 184.97 675.15 178.40 C663.44 171.83 651.39 165.53 638.91 160.62 C626.44 155.70 613.40 151.67 600.32 148.91 C587.23 146.15 573.71 144.45 560.40 144.07 C547.10 143.70 533.54 144.57 520.47 146.66 C507.40 148.76 494.36 152.30 481.99 156.64 C469.62 160.98 457.62 166.75 446.26 172.70 C434.89 178.65 424.14 185.59 413.78 192.32 C403.43 199.04 393.68 206.14 384.13 213.06 C374.58 219.97 365.47 226.89 356.49 233.82 C347.51 240.76 338.82 247.63 330.26 254.67 C321.71 261.72 313.34 268.77 305.15 276.09 C296.97 283.41 288.49 291.38 281.13 298.62 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M555.23 828.49 C566.61 827.11 576.44 825.50 587.08 823.63 C597.72 821.76 608.32 819.57 619.09 817.26 C629.86 814.95 640.60 812.38 651.68 809.76 C662.76 807.15 673.90 804.39 685.55 801.56 C697.21 798.74 709.20 796.01 721.61 792.82 C734.02 789.62 747.21 786.67 760.00 782.39 C772.78 778.12 786.12 773.40 798.33 767.18 C810.53 760.95 822.42 753.47 833.22 745.03 C844.02 736.60 854.02 726.85 863.10 716.54 C872.17 706.23 880.36 694.89 887.67 683.18 C894.98 671.47 901.40 659.00 906.97 646.30 C912.54 633.60 917.10 621.33 921.09 606.97 C925.07 592.61 928.82 575.91 930.86 560.14 C932.91 544.38 933.83 528.25 933.37 512.40 C932.90 496.55 931.18 480.52 928.08 465.05 C924.99 449.58 920.48 434.17 914.78 419.57 C909.09 404.98 901.86 390.77 893.89 377.48 C885.93 364.18 876.53 351.62 867.00 339.79 C857.48 327.96 847.00 317.01 836.73 306.48 C826.45 295.94 815.84 286.14 805.35 276.59 C794.86 267.03 784.36 257.97 773.78 249.15 C763.20 240.33 752.64 231.85 741.87 223.66 C731.10 215.47 720.27 207.57 709.15 200.03 C698.03 192.48 685.98 184.86 675.15 178.40 C664.31 171.94 654.66 166.74 644.13 161.25 C633.59 155.77 622.87 150.60 611.93 145.49 C600.99 140.37 589.89 135.52 578.49 130.55 C567.09 125.59 555.54 120.75 543.54 115.69 C531.54 110.64 519.32 105.35 506.49 100.23 C493.67 95.11 480.36 89.37 466.58 84.97 C452.80 80.58 438.30 76.24 423.79 73.85 C409.28 71.46 394.25 70.30 379.52 70.63 C364.78 70.96 349.87 72.86 335.38 75.83 C320.88 78.80 306.47 83.19 292.54 88.46 C278.62 93.72 264.94 100.19 251.83 107.41 C238.71 114.63 226.73 122.11 213.83 131.77 C200.94 141.42 186.68 153.21 174.46 165.31 C162.24 177.40 150.72 190.52 140.52 204.30 C130.32 218.08 121.06 232.82 113.28 247.98 C105.52 263.15 98.95 279.16 93.91 295.29 C88.86 311.42 85.32 328.21 83.00 344.76 C80.67 361.30 80.00 378.17 79.95 394.56 C79.89 410.95 81.20 427.24 82.67 443.09 C84.14 458.93 86.39 474.43 88.77 489.64 C91.16 504.86 93.92 519.71 96.96 534.40 C100.00 549.09 103.29 563.49 107.01 577.78 C110.73 592.07 114.73 606.17 119.25 620.13 C123.77 634.09 128.42 648.05 134.13 661.55 C139.84 675.06 146.05 688.54 153.50 701.15 C160.96 713.76 169.46 725.99 178.83 737.20 C188.21 748.41 198.63 759.03 209.74 768.41 C220.86 777.79 232.99 786.30 245.53 793.49 C258.06 800.68 271.52 806.67 284.96 811.55 C298.40 816.43 312.48 819.90 326.19 822.79 C339.90 825.67 353.77 827.33 367.21 828.86 C380.65 830.38 393.87 831.18 406.83 831.93 C419.80 832.67 432.44 833.09 444.99 833.34 C457.55 833.58 469.85 833.65 482.15 833.41 C494.45 833.18 506.62 832.75 518.80 831.93 C530.98 831.11 543.85 829.87 555.23 828.49 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M675.15 178.40 C664.31 171.94 654.66 166.74 644.13 161.25 C633.59 155.77 622.87 150.60 611.93 145.49 C600.99 140.37 589.89 135.52 578.49 130.55 C567.09 125.59 555.54 120.75 543.54 115.69 C531.54 110.64 519.32 105.35 506.49 100.23 C493.67 95.11 480.36 89.37 466.58 84.97 C452.80 80.58 438.30 76.24 423.79 73.85 C409.28 71.46 394.25 70.30 379.52 70.63 C364.78 70.96 349.87 72.86 335.38 75.83 C320.88 78.80 306.47 83.19 292.54 88.46 C278.62 93.72 264.94 100.19 251.83 107.41 C238.71 114.63 226.73 122.11 213.83 131.77 C200.94 141.42 186.68 153.21 174.46 165.31 C162.24 177.40 150.72 190.52 140.52 204.30 C130.32 218.08 121.06 232.82 113.28 247.98 C105.52 263.15 98.95 279.16 93.91 295.29 C88.86 311.42 85.32 328.21 83.00 344.76 C80.67 361.30 80.00 378.17 79.95 394.56 C79.89 410.95 81.20 427.24 82.67 443.09 C84.14 458.93 86.39 474.43 88.77 489.64 C91.16 504.86 93.92 519.71 96.96 534.40 C100.00 549.09 103.29 563.49 107.01 577.78 C110.73 592.07 114.73 606.17 119.25 620.13 C123.77 634.09 129.08 648.74 134.13 661.55 C139.18 674.37 144.06 685.30 149.55 697.04 C155.03 708.79 160.90 720.37 167.02 732.01 C173.13 743.66 179.59 755.17 186.24 766.91 C192.89 778.66 199.79 790.36 206.94 802.47 C214.09 814.58 221.31 826.98 229.14 839.57 C236.97 852.17 244.84 865.50 253.93 878.02 C263.02 890.55 272.75 903.42 283.70 914.72 C294.64 926.02 306.79 936.62 319.63 945.82 C332.47 955.02 346.43 963.05 360.73 969.92 C375.02 976.80 390.15 982.46 405.42 987.08 C420.69 991.70 436.51 995.16 452.35 997.65 C468.19 1000.13 483.23 1001.65 500.45 1001.96 C517.67 1002.27 537.43 1001.67 555.67 999.50 C573.92 997.33 592.24 993.84 609.91 988.95 C627.58 984.06 645.10 977.76 661.68 970.15 C678.26 962.53 694.41 953.43 709.40 943.25 C724.39 933.08 738.58 921.39 751.63 909.09 C764.68 896.80 776.60 883.17 787.69 869.50 C798.78 855.83 808.71 841.34 818.17 827.05 C827.64 812.77 836.23 798.23 844.50 783.79 C852.76 769.35 860.45 754.93 867.76 740.42 C875.07 725.91 881.95 711.43 888.36 696.75 C894.78 682.08 900.79 667.34 906.24 652.38 C911.69 637.42 916.98 622.34 921.09 606.97 C925.19 591.60 928.82 575.91 930.86 560.14 C932.91 544.38 933.83 528.25 933.37 512.40 C932.90 496.55 931.18 480.52 928.08 465.05 C924.99 449.58 920.48 434.17 914.78 419.57 C909.09 404.98 901.86 390.77 893.89 377.48 C885.93 364.18 876.53 351.62 867.00 339.79 C857.48 327.96 847.00 317.01 836.73 306.48 C826.45 295.94 815.84 286.14 805.35 276.59 C794.86 267.03 784.36 257.97 773.78 249.15 C763.20 240.33 752.64 231.85 741.87 223.66 C731.10 215.47 720.27 207.57 709.15 200.03 C698.03 192.48 685.98 184.86 675.15 178.40 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M134.13 661.55 C139.18 674.37 144.06 685.30 149.55 697.04 C155.03 708.79 160.90 720.37 167.02 732.01 C173.13 743.66 179.59 755.17 186.24 766.91 C192.89 778.66 199.79 790.36 206.94 802.47 C214.09 814.58 221.31 826.98 229.14 839.57 C236.97 852.17 244.84 865.50 253.93 878.02 C263.02 890.55 272.75 903.42 283.70 914.72 C294.64 926.02 306.79 936.62 319.63 945.82 C332.47 955.02 346.43 963.05 360.73 969.92 C375.02 976.80 390.15 982.46 405.42 987.08 C420.69 991.70 436.51 995.16 452.35 997.65 C468.19 1000.13 483.23 1001.65 500.45 1001.96 C517.67 1002.27 537.43 1001.67 555.67 999.50 C573.92 997.33 592.24 993.84 609.91 988.95 C627.58 984.06 645.10 977.76 661.68 970.15 C678.26 962.53 694.41 953.43 709.40 943.25 C724.39 933.08 738.58 921.39 751.63 909.09 C764.68 896.80 776.60 883.17 787.69 869.50 C798.78 855.83 808.71 841.34 818.17 827.05 C827.64 812.77 836.23 798.23 844.50 783.79 C852.76 769.35 860.45 754.93 867.76 740.42 C875.07 725.91 881.95 711.43 888.36 696.75 C894.78 682.08 900.79 667.34 906.24 652.38 C911.69 637.42 916.89 621.32 921.09 606.97 C925.28 592.62 928.33 579.99 931.39 566.28 C934.45 552.56 937.05 538.72 939.45 524.68 C941.85 510.65 943.84 496.50 945.77 482.05 C947.70 467.61 949.36 453.04 951.02 438.00 C952.67 422.97 954.40 407.67 955.71 391.86 C957.02 376.04 958.71 359.63 958.89 343.12 C959.06 326.62 958.90 309.46 956.76 292.82 C954.63 276.18 951.07 259.37 946.06 243.28 C941.05 227.18 934.38 211.34 926.71 196.24 C919.03 181.14 909.91 166.52 900.00 152.67 C890.08 138.81 878.98 125.56 867.23 113.12 C855.49 100.68 843.81 89.56 829.50 78.03 C815.20 66.49 798.17 54.05 781.40 43.90 C764.63 33.75 746.94 24.64 728.89 17.11 C710.84 9.59 692.00 3.32 673.08 -1.27 C654.17 -5.85 634.65 -8.93 615.38 -10.38 C596.10 -11.83 576.50 -11.53 557.44 -9.98 C538.38 -8.43 519.37 -5.06 501.00 -1.07 C482.63 2.92 464.66 8.33 447.20 13.94 C429.74 19.54 412.81 25.96 396.22 32.57 C379.64 39.17 363.52 46.20 347.68 53.58 C331.84 60.96 316.36 68.67 301.17 76.87 C285.98 85.07 271.09 93.64 256.54 102.79 C241.98 111.94 227.51 121.34 213.83 131.77 C200.15 142.19 186.68 153.21 174.46 165.31 C162.24 177.40 150.72 190.52 140.52 204.30 C130.32 218.08 121.06 232.82 113.28 247.98 C105.52 263.15 98.95 279.16 93.91 295.29 C88.86 311.42 85.32 328.21 83.00 344.76 C80.67 361.30 80.00 378.17 79.95 394.56 C79.89 410.95 81.20 427.24 82.67 443.09 C84.14 458.93 86.39 474.43 88.77 489.64 C91.16 504.86 93.92 519.71 96.96 534.40 C100.00 549.09 103.29 563.49 107.01 577.78 C110.73 592.07 114.73 606.17 119.25 620.13 C123.77 634.09 129.08 648.74 134.13 661.55 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M921.09 606.97 C925.28 592.62 928.33 579.99 931.39 566.28 C934.45 552.56 937.05 538.72 939.45 524.68 C941.85 510.65 943.84 496.50 945.77 482.05 C947.70 467.61 949.36 453.04 951.02 438.00 C952.67 422.97 954.40 407.67 955.71 391.86 C957.02 376.04 958.71 359.63 958.89 343.12 C959.06 326.62 958.90 309.46 956.76 292.82 C954.63 276.18 951.07 259.37 946.06 243.28 C941.05 227.18 934.38 211.34 926.71 196.24 C919.03 181.14 909.91 166.52 900.00 152.67 C890.08 138.81 878.98 125.56 867.23 113.12 C855.49 100.68 843.81 89.56 829.50 78.03 C815.20 66.49 798.17 54.05 781.40 43.90 C764.63 33.75 746.94 24.64 728.89 17.11 C710.84 9.59 692.00 3.32 673.08 -1.27 C654.17 -5.85 634.65 -8.93 615.38 -10.38 C596.10 -11.83 576.50 -11.53 557.44 -9.98 C538.38 -8.43 519.37 -5.06 501.00 -1.07 C482.63 2.92 464.66 8.33 447.20 13.94 C429.74 19.54 412.81 25.96 396.22 32.57 C379.64 39.17 363.52 46.20 347.68 53.58 C331.84 60.96 316.36 68.67 301.17 76.87 C285.98 85.07 271.09 93.64 256.54 102.79 C241.98 111.94 226.89 122.26 213.83 131.77 C200.77 141.27 189.82 150.07 178.16 159.80 C166.51 169.53 155.16 179.68 143.90 190.14 C132.63 200.60 121.65 211.42 110.58 222.54 C99.52 233.66 88.63 245.07 77.50 256.86 C66.38 268.65 55.09 280.64 43.86 293.30 C32.63 305.96 20.79 318.88 10.14 332.82 C-0.51 346.76 -11.24 361.43 -20.03 376.92 C-28.83 392.41 -36.55 408.92 -42.63 425.76 C-48.71 442.61 -53.27 460.27 -56.52 477.98 C-59.77 495.69 -61.54 513.91 -62.15 532.00 C-62.77 550.10 -62.02 568.45 -60.21 586.55 C-58.40 604.64 -55.81 621.58 -51.30 640.58 C-46.79 659.59 -40.58 681.13 -33.13 700.59 C-25.68 720.05 -16.79 739.24 -6.59 757.36 C3.61 775.49 15.26 793.08 28.08 809.35 C40.89 825.61 55.16 841.03 70.28 854.97 C85.40 868.90 101.92 881.62 118.82 892.96 C135.71 904.31 153.74 914.15 171.68 923.05 C189.61 931.95 208.15 939.43 226.44 946.39 C244.72 953.35 263.12 959.28 281.39 964.79 C299.66 970.31 317.85 975.14 336.05 979.46 C354.25 983.78 372.39 987.57 390.60 990.71 C408.82 993.86 427.05 996.46 445.35 998.34 C463.66 1000.21 482.06 1001.77 500.45 1001.96 C518.83 1002.16 537.43 1001.67 555.67 999.50 C573.92 997.33 592.24 993.84 609.91 988.95 C627.58 984.06 645.10 977.76 661.68 970.15 C678.26 962.53 694.41 953.43 709.40 943.25 C724.39 933.08 738.58 921.39 751.63 909.09 C764.68 896.80 776.60 883.17 787.69 869.50 C798.78 855.83 808.71 841.34 818.17 827.05 C827.64 812.77 836.23 798.23 844.50 783.79 C852.76 769.35 860.45 754.93 867.76 740.42 C875.07 725.91 881.95 711.43 888.36 696.75 C894.78 682.08 900.79 667.34 906.24 652.38 C911.69 637.42 916.89 621.32 921.09 606.97 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M213.83 131.77 C200.77 141.27 189.82 150.07 178.16 159.80 C166.51 169.53 155.16 179.68 143.90 190.14 C132.63 200.60 121.65 211.42 110.58 222.54 C99.52 233.66 88.63 245.07 77.50 256.86 C66.38 268.65 55.09 280.64 43.86 293.30 C32.63 305.96 20.79 318.88 10.14 332.82 C-0.51 346.76 -11.24 361.43 -20.03 376.92 C-28.83 392.41 -36.55 408.92 -42.63 425.76 C-48.71 442.61 -53.27 460.27 -56.52 477.98 C-59.77 495.69 -61.54 513.91 -62.15 532.00 C-62.77 550.10 -62.02 568.45 -60.21 586.55 C-58.40 604.64 -55.81 621.58 -51.30 640.58 C-46.79 659.59 -40.58 681.13 -33.13 700.59 C-25.68 720.05 -16.79 739.24 -6.59 757.36 C3.61 775.49 15.26 793.08 28.08 809.35 C40.89 825.61 55.16 841.03 70.28 854.97 C85.40 868.90 101.92 881.62 118.82 892.96 C135.71 904.31 153.74 914.15 171.68 923.05 C189.61 931.95 208.15 939.43 226.44 946.39 C244.72 953.35 263.12 959.28 281.39 964.79 C299.66 970.31 317.85 975.14 336.05 979.46 C354.25 983.78 372.39 987.57 390.60 990.71 C408.82 993.86 427.05 996.46 445.35 998.34 C463.66 1000.21 483.18 1001.42 500.45 1001.96 C517.71 1002.51 532.72 1002.24 548.93 1001.62 C565.13 1001.00 581.35 999.80 597.67 998.25 C613.99 996.70 630.31 994.63 646.87 992.33 C663.42 990.02 680.01 987.32 697.00 984.41 C713.98 981.50 731.21 978.50 748.78 974.87 C766.35 971.24 784.55 967.74 802.41 962.66 C820.26 957.58 838.61 951.92 855.91 944.40 C873.20 936.88 890.26 927.85 906.18 917.52 C922.10 907.19 937.30 895.23 951.45 882.42 C965.60 869.61 978.86 855.43 991.09 840.63 C1003.32 825.83 1014.59 809.96 1024.82 793.61 C1035.05 777.26 1043.89 761.40 1052.47 742.56 C1061.06 723.71 1069.89 701.66 1076.35 680.54 C1082.81 659.41 1087.91 637.60 1091.25 615.81 C1094.59 594.02 1096.38 571.74 1096.39 549.79 C1096.39 527.85 1094.65 505.66 1091.27 484.13 C1087.88 462.60 1082.57 441.18 1076.05 420.61 C1069.54 400.05 1061.14 379.99 1052.17 360.74 C1043.20 341.49 1032.81 323.00 1022.23 305.12 C1011.66 287.23 1000.30 270.09 988.73 253.41 C977.16 236.73 965.17 220.65 952.80 205.02 C940.43 189.39 927.72 174.25 914.51 159.61 C901.29 144.97 887.68 130.80 873.51 117.20 C859.35 103.60 844.86 90.24 829.50 78.03 C814.15 65.81 798.17 54.05 781.40 43.90 C764.63 33.75 746.94 24.64 728.89 17.11 C710.84 9.59 692.00 3.32 673.08 -1.27 C654.17 -5.85 634.65 -8.93 615.38 -10.38 C596.10 -11.83 576.50 -11.53 557.44 -9.98 C538.38 -8.43 519.37 -5.06 501.00 -1.07 C482.63 2.92 464.66 8.33 447.20 13.94 C429.74 19.54 412.81 25.96 396.22 32.57 C379.64 39.17 363.52 46.20 347.68 53.58 C331.84 60.96 316.36 68.67 301.17 76.87 C285.98 85.07 271.09 93.64 256.54 102.79 C241.98 111.94 226.89 122.26 213.83 131.77 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M500.45 1001.96 C517.71 1002.51 532.72 1002.24 548.93 1001.62 C565.13 1001.00 581.35 999.80 597.67 998.25 C613.99 996.70 630.31 994.63 646.87 992.33 C663.42 990.02 680.01 987.32 697.00 984.41 C713.98 981.50 731.21 978.50 748.78 974.87 C766.35 971.24 784.55 967.74 802.41 962.66 C820.26 957.58 838.61 951.92 855.91 944.40 C873.20 936.88 890.26 927.85 906.18 917.52 C922.10 907.19 937.30 895.23 951.45 882.42 C965.60 869.61 978.86 855.43 991.09 840.63 C1003.32 825.83 1014.59 809.96 1024.82 793.61 C1035.05 777.26 1043.89 761.40 1052.47 742.56 C1061.06 723.71 1069.89 701.66 1076.35 680.54 C1082.81 659.41 1087.91 637.60 1091.25 615.81 C1094.59 594.02 1096.38 571.74 1096.39 549.79 C1096.39 527.85 1094.65 505.66 1091.27 484.13 C1087.88 462.60 1082.57 441.18 1076.05 420.61 C1069.54 400.05 1061.14 379.99 1052.17 360.74 C1043.20 341.49 1032.81 323.00 1022.23 305.12 C1011.66 287.23 1000.30 270.09 988.73 253.41 C977.16 236.73 965.17 220.65 952.80 205.02 C940.43 189.39 927.72 174.25 914.51 159.61 C901.29 144.97 887.68 130.80 873.51 117.20 C859.35 103.60 843.69 89.79 829.50 78.03 C815.32 66.26 802.48 56.67 788.41 46.62 C774.34 36.56 759.87 27.00 745.09 17.70 C730.30 8.39 715.17 -0.45 699.69 -9.21 C684.21 -17.96 668.44 -26.38 652.22 -34.82 C635.99 -43.27 619.50 -51.74 602.35 -59.88 C585.21 -68.02 567.62 -76.59 549.34 -83.66 C531.06 -90.72 512.03 -97.59 492.68 -102.28 C473.34 -106.97 453.26 -110.26 433.28 -111.79 C413.30 -113.31 392.90 -113.06 372.82 -111.43 C352.73 -109.80 332.52 -106.49 312.75 -102.01 C292.99 -97.52 273.33 -91.55 254.24 -84.51 C235.14 -77.48 217.52 -69.95 198.17 -59.78 C178.83 -49.61 157.21 -36.92 138.15 -23.45 C119.08 -9.97 100.70 5.01 83.73 21.15 C66.69 37.35 50.52 54.94 35.97 73.32 C21.45 91.62 8.26 111.17 -3.10 131.24 C-14.44 151.29 -24.13 172.49 -32.25 193.79 C-40.37 215.08 -46.62 237.18 -51.85 259.00 C-57.09 280.83 -60.66 302.96 -63.66 324.74 C-66.66 346.53 -68.51 368.23 -69.85 389.71 C-71.18 411.19 -71.75 432.47 -71.68 453.63 C-71.60 474.79 -70.89 495.79 -69.40 516.68 C-67.91 537.58 -65.76 558.36 -62.74 579.01 C-59.72 599.66 -56.23 620.32 -51.30 640.58 C-46.37 660.85 -40.58 681.13 -33.13 700.59 C-25.68 720.05 -16.79 739.24 -6.59 757.36 C3.61 775.49 15.26 793.08 28.08 809.35 C40.89 825.61 55.16 841.03 70.28 854.97 C85.40 868.90 101.92 881.62 118.82 892.96 C135.71 904.31 153.74 914.15 171.68 923.05 C189.61 931.95 208.15 939.43 226.44 946.39 C244.72 953.35 263.12 959.28 281.39 964.79 C299.66 970.31 317.85 975.14 336.05 979.46 C354.25 983.78 372.39 987.57 390.60 990.71 C408.82 993.86 427.05 996.46 445.35 998.34 C463.66 1000.21 483.18 1001.42 500.45 1001.96 Z`

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
	__Parameter__000000_Reference.SideLength = 170.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.450000
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
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 170.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 170.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 292.505061
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 33.775575
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 170.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 170.000000
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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 263.735062
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -21.853710
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 292.217157
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 42.611767
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 174.180403
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 321.857195
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 218.345554
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 264.123144
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 180.200000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 263.735062
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 180.200000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -21.853710
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 263.735062
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 344.316908
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -21.853710
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -57.456330
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 95.433990
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 7.907886
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 63.994300
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 186.409028
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 263.735062
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -21.853710
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 53.027389
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -154.463506
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

