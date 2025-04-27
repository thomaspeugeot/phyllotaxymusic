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
	__Axis__000001_Construction_Axis.AngleDegree = 97.971943
	__Axis__000001_Construction_Axis.Length = 154.865638
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -21.478030
	__Axis__000001_Construction_Axis.EndY = 153.369032
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
	__Axis__000002_Initial_Axis.AngleDegree = 82.028057
	__Axis__000002_Initial_Axis.Length = 558.323209
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
	__Axis__000004_Rotated_Axis.Length = 558.323209
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -10.739015
	__Bezier__000005_Growth_Curve_Seed.StartY = 76.684516
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 48.086949
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 84.922593
	__Bezier__000005_Growth_Curve_Seed.EndX = 88.038808
	__Bezier__000005_Growth_Curve_Seed.EndY = 168.705935
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 29.212844
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 160.467858
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
	__BezierGrid__000000_2nb_Voice.IsDisplayed = false

	__BezierGrid__000001_2nd_voice_shifted_right.Name = `2nd voice shifted right`
	__BezierGrid__000001_2nd_voice_shifted_right.IsDisplayed = false

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
	__Circle__000001_Construction_Circle.CenterX = -10.739015
	__Circle__000001_Construction_Circle.CenterY = 76.684516
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 219.033677
	__Circle__000005_Rotated_Next_Circle.CenterY = 30.673806
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

	__ExportToMusicxml__000000_Singloton.Name = `Singloton`

	__FrontCurve__000000_Non_Rotated_0_.Name = `Non Rotated 0 `
	__FrontCurve__000000_Non_Rotated_0_.Path = `M718.99 527.09 C720.35 519.67 721.23 513.66 722.23 506.82 C723.23 499.99 724.06 493.13 724.97 486.10 C725.88 479.06 726.66 472.00 727.68 464.62 C728.69 457.24 729.68 449.80 731.06 441.81 C732.45 433.83 734.16 425.61 735.97 416.72 C737.78 407.82 740.50 398.18 741.93 388.47 C743.37 378.76 744.86 368.28 744.57 358.46 C744.27 348.63 742.71 338.78 740.17 329.52 C737.63 320.26 733.76 311.35 729.33 302.92 C724.90 294.50 719.46 286.49 713.60 278.97 C707.73 271.46 701.11 264.38 694.15 257.81 C687.19 251.23 680.58 245.59 671.83 239.53 C663.08 233.46 652.14 226.58 641.66 221.42 C631.18 216.26 620.12 211.84 608.93 208.56 C597.75 205.29 586.09 202.92 574.55 201.77 C563.00 200.62 551.15 200.57 539.66 201.66 C528.18 202.75 516.62 205.15 505.64 208.30 C494.66 211.46 483.94 215.95 473.78 220.60 C463.63 225.25 453.99 230.84 444.73 236.22 C435.46 241.60 426.73 247.32 418.19 252.88 C409.65 258.43 401.51 263.97 393.48 269.52 C385.44 275.08 377.67 280.56 370.00 286.19 C362.33 291.83 354.83 297.46 347.45 303.34 C340.08 309.22 332.25 315.80 325.77 321.47 C319.29 327.14 314.27 331.91 308.60 337.34 C302.92 342.78 297.39 348.36 291.73 354.06 C286.06 359.76 280.51 365.59 274.62 371.54 C268.73 377.48 262.86 383.52 256.38 389.75 C249.91 395.98 243.06 402.20 235.76 408.90 C228.47 415.60 220.02 422.34 212.59 429.93 C205.16 437.53 197.24 445.70 191.17 454.48 C185.10 463.26 180.04 472.87 176.17 482.62 C172.30 492.37 169.72 502.70 167.93 512.99 C166.14 523.27 165.44 533.87 165.43 544.34 C165.41 554.81 166.32 565.41 167.84 575.81 C169.36 586.22 171.19 595.58 174.55 606.77 C177.91 617.97 182.57 631.37 188.00 643.00 C193.42 654.63 199.82 666.03 207.09 676.58 C214.35 687.12 222.61 697.22 231.59 706.26 C240.58 715.30 250.55 723.65 260.98 730.82 C271.42 737.99 282.78 744.15 294.20 749.30 C305.62 754.45 317.72 758.36 329.51 761.73 C341.30 765.09 353.31 767.35 364.93 769.48 C376.54 771.61 387.98 773.08 399.19 774.50 C410.40 775.92 421.32 777.04 432.16 778.01 C443.01 778.99 453.63 779.80 464.28 780.35 C474.92 780.91 485.44 781.30 496.00 781.35 C506.57 781.40 517.16 781.48 527.68 780.65 C538.21 779.82 548.86 778.69 559.15 776.38 C569.43 774.07 579.71 770.89 589.40 766.80 C599.09 762.70 608.59 757.68 617.28 751.83 C625.97 745.98 634.23 739.13 641.53 731.69 C648.83 724.25 655.40 715.81 661.08 707.19 C666.77 698.56 671.46 689.12 675.65 679.94 C679.84 670.77 683.08 661.24 686.22 652.15 C689.36 643.06 691.90 634.08 694.47 625.38 C697.04 616.68 699.35 608.28 701.64 599.97 C703.92 591.65 706.10 583.59 708.16 575.49 C710.23 567.39 712.23 559.42 714.03 551.36 C715.83 543.29 717.62 534.51 718.99 527.09 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M325.77 321.47 C319.29 327.14 314.27 331.91 308.60 337.34 C302.92 342.78 297.39 348.36 291.73 354.06 C286.06 359.76 280.51 365.59 274.62 371.54 C268.73 377.48 262.86 383.52 256.38 389.75 C249.91 395.98 243.06 402.20 235.76 408.90 C228.47 415.60 220.02 422.34 212.59 429.93 C205.16 437.53 197.24 445.70 191.17 454.48 C185.10 463.26 180.04 472.87 176.17 482.62 C172.30 492.37 169.72 502.70 167.93 512.99 C166.14 523.27 165.44 533.87 165.43 544.34 C165.41 554.81 166.32 565.41 167.84 575.81 C169.36 586.22 171.19 595.58 174.55 606.77 C177.91 617.97 182.57 631.37 188.00 643.00 C193.42 654.63 199.82 666.03 207.09 676.58 C214.35 687.12 222.61 697.22 231.59 706.26 C240.58 715.30 250.55 723.65 260.98 730.82 C271.42 737.99 282.78 744.15 294.20 749.30 C305.62 754.45 317.72 758.36 329.51 761.73 C341.30 765.09 353.31 767.35 364.93 769.48 C376.54 771.61 387.98 773.08 399.19 774.50 C410.40 775.92 421.32 777.04 432.16 778.01 C443.01 778.99 453.63 779.80 464.28 780.35 C474.92 780.91 485.44 781.30 496.00 781.35 C506.57 781.40 518.05 781.11 527.68 780.65 C537.31 780.19 545.05 779.48 553.79 778.61 C562.54 777.75 571.26 776.63 580.14 775.47 C589.02 774.30 597.89 772.92 607.08 771.62 C616.26 770.32 625.48 768.90 635.26 767.66 C645.04 766.43 655.09 765.38 665.76 764.22 C676.44 763.06 688.00 762.54 699.29 760.69 C710.58 758.83 722.58 756.80 733.50 753.10 C744.42 749.40 755.04 744.39 764.79 738.49 C774.55 732.60 783.62 725.42 792.02 717.75 C800.42 710.09 808.13 701.46 815.19 692.49 C822.26 683.52 828.65 673.84 834.40 663.92 C840.15 654.00 844.94 644.77 849.70 632.94 C854.46 621.12 859.62 606.54 862.97 592.95 C866.32 579.36 868.70 565.32 869.81 551.40 C870.91 537.49 870.91 523.27 869.60 509.44 C868.30 495.61 865.73 481.70 862.00 468.42 C858.26 455.14 853.11 442.07 847.21 429.77 C841.32 417.47 834.05 405.72 826.63 394.62 C819.20 383.52 810.85 373.15 802.66 363.16 C794.48 353.18 785.95 343.84 777.52 334.72 C769.10 325.61 760.64 316.93 752.11 308.46 C743.58 299.99 735.06 291.82 726.32 283.89 C717.59 275.96 708.79 268.28 699.71 260.88 C690.63 253.49 681.50 246.11 671.83 239.53 C662.15 232.95 652.14 226.58 641.66 221.42 C631.18 216.26 620.12 211.84 608.93 208.56 C597.75 205.29 586.09 202.92 574.55 201.77 C563.00 200.62 551.15 200.57 539.66 201.66 C528.18 202.75 516.62 205.15 505.64 208.30 C494.66 211.46 483.94 215.95 473.78 220.60 C463.63 225.25 453.99 230.84 444.73 236.22 C435.46 241.60 426.73 247.32 418.19 252.88 C409.65 258.43 401.51 263.97 393.48 269.52 C385.44 275.08 377.67 280.56 370.00 286.19 C362.33 291.83 354.83 297.46 347.45 303.34 C340.08 309.22 332.25 315.80 325.77 321.47 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M527.68 780.65 C537.31 780.19 545.05 779.48 553.79 778.61 C562.54 777.75 571.26 776.63 580.14 775.47 C589.02 774.30 597.89 772.92 607.08 771.62 C616.26 770.32 625.48 768.90 635.26 767.66 C645.04 766.43 655.09 765.38 665.76 764.22 C676.44 763.06 688.00 762.54 699.29 760.69 C710.58 758.83 722.58 756.80 733.50 753.10 C744.42 749.40 755.04 744.39 764.79 738.49 C774.55 732.60 783.62 725.42 792.02 717.75 C800.42 710.09 808.13 701.46 815.19 692.49 C822.26 683.52 828.65 673.84 834.40 663.92 C840.15 654.00 844.94 644.77 849.70 632.94 C854.46 621.12 859.62 606.54 862.97 592.95 C866.32 579.36 868.70 565.32 869.81 551.40 C870.91 537.49 870.91 523.27 869.60 509.44 C868.30 495.61 865.73 481.70 862.00 468.42 C858.26 455.14 853.11 442.07 847.21 429.77 C841.32 417.47 834.05 405.72 826.63 394.62 C819.20 383.52 810.85 373.15 802.66 363.16 C794.48 353.18 785.95 343.84 777.52 334.72 C769.10 325.61 760.64 316.93 752.11 308.46 C743.58 299.99 735.06 291.82 726.32 283.89 C717.59 275.96 708.79 268.28 699.71 260.88 C690.63 253.49 680.48 245.79 671.83 239.53 C663.17 233.27 655.96 228.56 647.78 223.31 C639.59 218.07 631.24 213.06 622.72 208.04 C614.20 203.03 605.54 198.22 596.67 193.20 C587.80 188.17 578.81 183.27 569.50 177.91 C560.19 172.55 550.77 166.91 540.79 161.01 C530.81 155.12 520.53 148.21 509.63 142.56 C498.73 136.91 487.16 130.99 475.40 127.10 C463.64 123.21 451.28 120.58 439.07 119.22 C426.85 117.86 414.35 117.98 402.09 118.94 C389.83 119.90 377.52 122.10 365.51 125.00 C353.50 127.91 341.59 131.83 330.03 136.39 C318.48 140.95 308.21 145.55 296.19 152.35 C284.17 159.16 269.95 168.02 257.90 177.24 C245.84 186.46 234.29 196.70 223.85 207.68 C213.41 218.67 203.70 230.62 195.28 243.13 C186.87 255.63 179.44 269.04 173.38 282.70 C167.31 296.36 162.54 310.79 158.89 325.10 C155.24 339.42 153.06 354.19 151.48 368.58 C149.91 382.96 149.57 397.39 149.41 411.41 C149.25 425.44 149.83 439.20 150.53 452.72 C151.23 466.25 152.30 479.46 153.62 492.55 C154.94 505.65 156.50 518.51 158.45 531.31 C160.40 544.10 162.61 556.75 165.29 569.33 C167.98 581.90 170.77 594.49 174.55 606.77 C178.34 619.05 182.57 631.37 188.00 643.00 C193.42 654.63 199.82 666.03 207.09 676.58 C214.35 687.12 222.61 697.22 231.59 706.26 C240.58 715.30 250.55 723.65 260.98 730.82 C271.42 737.99 282.78 744.15 294.20 749.30 C305.62 754.45 317.72 758.36 329.51 761.73 C341.30 765.09 353.31 767.35 364.93 769.48 C376.54 771.61 387.98 773.08 399.19 774.50 C410.40 775.92 421.32 777.04 432.16 778.01 C443.01 778.99 453.63 779.80 464.28 780.35 C474.92 780.91 485.44 781.30 496.00 781.35 C506.57 781.40 518.05 781.11 527.68 780.65 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M671.83 239.53 C663.17 233.27 655.96 228.56 647.78 223.31 C639.59 218.07 631.24 213.06 622.72 208.04 C614.20 203.03 605.54 198.22 596.67 193.20 C587.80 188.17 578.81 183.27 569.50 177.91 C560.19 172.55 550.77 166.91 540.79 161.01 C530.81 155.12 520.53 148.21 509.63 142.56 C498.73 136.91 487.16 130.99 475.40 127.10 C463.64 123.21 451.28 120.58 439.07 119.22 C426.85 117.86 414.35 117.98 402.09 118.94 C389.83 119.90 377.52 122.10 365.51 125.00 C353.50 127.91 341.59 131.83 330.03 136.39 C318.48 140.95 308.21 145.55 296.19 152.35 C284.17 159.16 269.95 168.02 257.90 177.24 C245.84 186.46 234.29 196.70 223.85 207.68 C213.41 218.67 203.70 230.62 195.28 243.13 C186.87 255.63 179.44 269.04 173.38 282.70 C167.31 296.36 162.54 310.79 158.89 325.10 C155.24 339.42 153.06 354.19 151.48 368.58 C149.91 382.96 149.57 397.39 149.41 411.41 C149.25 425.44 149.83 439.20 150.53 452.72 C151.23 466.25 152.30 479.46 153.62 492.55 C154.94 505.65 156.50 518.51 158.45 531.31 C160.40 544.10 162.61 556.75 165.29 569.33 C167.98 581.90 171.41 595.47 174.55 606.77 C177.70 618.07 180.69 627.05 184.16 637.13 C187.64 647.21 191.44 657.20 195.39 667.27 C199.35 677.35 203.59 687.33 207.89 697.58 C212.19 707.83 216.70 718.04 221.21 728.79 C225.71 739.53 230.17 750.53 234.93 762.05 C239.69 773.56 244.08 786.03 249.76 797.87 C255.45 809.70 261.53 822.14 269.04 833.07 C276.56 844.01 285.35 854.30 294.85 863.48 C304.35 872.66 315.02 880.82 326.04 888.14 C337.07 895.46 348.92 901.81 360.99 907.38 C373.06 912.94 385.71 917.62 398.48 921.53 C411.25 925.44 422.95 928.47 437.61 930.84 C452.26 933.21 470.11 935.35 486.41 935.75 C502.70 936.16 519.24 935.38 535.37 933.26 C551.49 931.14 567.67 927.74 583.15 923.04 C598.64 918.33 613.92 912.23 628.27 905.01 C642.63 897.80 656.45 889.11 669.29 879.76 C682.14 870.40 694.11 859.70 705.34 848.87 C716.56 838.05 726.82 826.36 736.65 814.82 C746.47 803.28 755.53 791.43 764.29 779.64 C773.05 767.86 781.29 756.04 789.21 744.11 C797.12 732.18 804.66 720.23 811.80 708.07 C818.94 695.91 825.72 683.65 832.04 671.13 C838.36 658.61 844.54 645.97 849.70 632.94 C854.85 619.91 859.62 606.54 862.97 592.95 C866.32 579.36 868.70 565.32 869.81 551.40 C870.91 537.49 870.91 523.27 869.60 509.44 C868.30 495.61 865.73 481.70 862.00 468.42 C858.26 455.14 853.11 442.07 847.21 429.77 C841.32 417.47 834.05 405.72 826.63 394.62 C819.20 383.52 810.85 373.15 802.66 363.16 C794.48 353.18 785.95 343.84 777.52 334.72 C769.10 325.61 760.64 316.93 752.11 308.46 C743.58 299.99 735.06 291.82 726.32 283.89 C717.59 275.96 708.79 268.28 699.71 260.88 C690.63 253.49 680.48 245.79 671.83 239.53 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M174.55 606.77 C177.70 618.07 180.69 627.05 184.16 637.13 C187.64 647.21 191.44 657.20 195.39 667.27 C199.35 677.35 203.59 687.33 207.89 697.58 C212.19 707.83 216.70 718.04 221.21 728.79 C225.71 739.53 230.17 750.53 234.93 762.05 C239.69 773.56 244.08 786.03 249.76 797.87 C255.45 809.70 261.53 822.14 269.04 833.07 C276.56 844.01 285.35 854.30 294.85 863.48 C304.35 872.66 315.02 880.82 326.04 888.14 C337.07 895.46 348.92 901.81 360.99 907.38 C373.06 912.94 385.71 917.62 398.48 921.53 C411.25 925.44 422.95 928.47 437.61 930.84 C452.26 933.21 470.11 935.35 486.41 935.75 C502.70 936.16 519.24 935.38 535.37 933.26 C551.49 931.14 567.67 927.74 583.15 923.04 C598.64 918.33 613.92 912.23 628.27 905.01 C642.63 897.80 656.45 889.11 669.29 879.76 C682.14 870.40 694.11 859.70 705.34 848.87 C716.56 838.05 726.82 826.36 736.65 814.82 C746.47 803.28 755.53 791.43 764.29 779.64 C773.05 767.86 781.29 756.04 789.21 744.11 C797.12 732.18 804.66 720.23 811.80 708.07 C818.94 695.91 825.72 683.65 832.04 671.13 C838.36 658.61 844.68 644.71 849.70 632.94 C854.71 621.18 858.25 611.49 862.12 600.53 C865.98 589.58 869.49 578.48 872.89 567.21 C876.29 555.93 879.36 544.54 882.51 532.90 C885.67 521.26 888.59 509.52 891.80 497.37 C895.02 485.21 898.39 472.89 901.80 459.98 C905.21 447.07 909.46 433.65 912.27 419.92 C915.08 406.19 917.90 391.75 918.63 377.60 C919.37 363.45 918.65 349.01 916.66 335.02 C914.66 321.02 911.07 307.10 906.64 293.64 C902.21 280.18 896.49 266.97 890.07 254.27 C883.65 241.56 876.19 229.20 868.11 217.40 C860.04 205.60 852.29 195.25 841.64 183.45 C830.99 171.66 817.53 157.92 804.22 146.65 C790.91 135.37 776.62 124.90 761.79 115.81 C746.95 106.71 731.23 98.63 715.21 92.07 C699.18 85.50 682.41 80.21 665.64 76.42 C648.87 72.64 631.57 70.40 614.60 69.36 C597.63 68.32 580.46 68.94 563.81 70.19 C547.16 71.44 530.69 74.03 514.67 76.85 C498.65 79.67 483.01 83.27 467.67 87.08 C452.32 90.89 437.36 95.12 422.59 99.70 C407.83 104.28 393.36 109.18 379.09 114.56 C364.81 119.94 350.76 125.68 336.95 131.98 C323.13 138.28 309.37 144.81 296.19 152.35 C283.02 159.90 269.95 168.02 257.90 177.24 C245.84 186.46 234.29 196.70 223.85 207.68 C213.41 218.67 203.70 230.62 195.28 243.13 C186.87 255.63 179.44 269.04 173.38 282.70 C167.31 296.36 162.54 310.79 158.89 325.10 C155.24 339.42 153.06 354.19 151.48 368.58 C149.91 382.96 149.57 397.39 149.41 411.41 C149.25 425.44 149.83 439.20 150.53 452.72 C151.23 466.25 152.30 479.46 153.62 492.55 C154.94 505.65 156.50 518.51 158.45 531.31 C160.40 544.10 162.61 556.75 165.29 569.33 C167.98 581.90 171.41 595.47 174.55 606.77 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M849.70 632.94 C854.71 621.18 858.25 611.49 862.12 600.53 C865.98 589.58 869.49 578.48 872.89 567.21 C876.29 555.93 879.36 544.54 882.51 532.90 C885.67 521.26 888.59 509.52 891.80 497.37 C895.02 485.21 898.39 472.89 901.80 459.98 C905.21 447.07 909.46 433.65 912.27 419.92 C915.08 406.19 917.90 391.75 918.63 377.60 C919.37 363.45 918.65 349.01 916.66 335.02 C914.66 321.02 911.07 307.10 906.64 293.64 C902.21 280.18 896.49 266.97 890.07 254.27 C883.65 241.56 876.19 229.20 868.11 217.40 C860.04 205.60 852.29 195.25 841.64 183.45 C830.99 171.66 817.53 157.92 804.22 146.65 C790.91 135.37 776.62 124.90 761.79 115.81 C746.95 106.71 731.23 98.63 715.21 92.07 C699.18 85.50 682.41 80.21 665.64 76.42 C648.87 72.64 631.57 70.40 614.60 69.36 C597.63 68.32 580.46 68.94 563.81 70.19 C547.16 71.44 530.69 74.03 514.67 76.85 C498.65 79.67 483.01 83.27 467.67 87.08 C452.32 90.89 437.36 95.12 422.59 99.70 C407.83 104.28 393.36 109.18 379.09 114.56 C364.81 119.94 350.76 125.68 336.95 131.98 C323.13 138.28 308.38 145.77 296.19 152.35 C284.00 158.94 274.47 164.78 263.82 171.48 C253.16 178.18 242.72 185.26 232.29 192.57 C221.85 199.88 211.61 207.51 201.19 215.33 C190.77 223.14 180.49 231.20 169.76 239.43 C159.03 247.66 148.10 255.95 136.81 264.69 C125.53 273.42 113.29 282.06 102.04 291.86 C90.78 301.65 79.11 312.00 69.29 323.45 C59.48 334.90 50.65 347.50 43.15 360.56 C35.66 373.63 29.49 387.64 24.31 401.83 C19.14 416.02 15.18 430.83 12.12 445.69 C9.07 460.55 7.09 475.80 5.98 491.00 C4.87 506.19 4.45 519.96 5.46 536.86 C6.47 553.77 8.53 574.15 12.07 592.40 C15.61 610.65 20.49 628.88 26.71 646.35 C32.93 663.83 40.55 681.06 49.41 697.26 C58.26 713.46 68.57 729.12 79.86 743.57 C91.15 758.01 103.89 771.56 117.17 783.92 C130.45 796.29 144.97 807.44 159.53 817.74 C174.09 828.04 189.40 837.13 204.53 845.72 C219.67 854.31 235.04 862.00 250.35 869.29 C265.65 876.58 280.96 883.26 296.36 889.47 C311.75 895.69 327.16 901.42 342.72 906.59 C358.28 911.76 373.92 916.44 389.74 920.49 C405.55 924.53 421.50 928.30 437.61 930.84 C453.72 933.39 470.11 935.35 486.41 935.75 C502.70 936.16 519.24 935.38 535.37 933.26 C551.49 931.14 567.67 927.74 583.15 923.04 C598.64 918.33 613.92 912.23 628.27 905.01 C642.63 897.80 656.45 889.11 669.29 879.76 C682.14 870.40 694.11 859.70 705.34 848.87 C716.56 838.05 726.82 826.36 736.65 814.82 C746.47 803.28 755.53 791.43 764.29 779.64 C773.05 767.86 781.29 756.04 789.21 744.11 C797.12 732.18 804.66 720.23 811.80 708.07 C818.94 695.91 825.72 683.65 832.04 671.13 C838.36 658.61 844.68 644.71 849.70 632.94 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M296.19 152.35 C284.00 158.94 274.47 164.78 263.82 171.48 C253.16 178.18 242.72 185.26 232.29 192.57 C221.85 199.88 211.61 207.51 201.19 215.33 C190.77 223.14 180.49 231.20 169.76 239.43 C159.03 247.66 148.10 255.95 136.81 264.69 C125.53 273.42 113.29 282.06 102.04 291.86 C90.78 301.65 79.11 312.00 69.29 323.45 C59.48 334.90 50.65 347.50 43.15 360.56 C35.66 373.63 29.49 387.64 24.31 401.83 C19.14 416.02 15.18 430.83 12.12 445.69 C9.07 460.55 7.09 475.80 5.98 491.00 C4.87 506.19 4.45 519.96 5.46 536.86 C6.47 553.77 8.53 574.15 12.07 592.40 C15.61 610.65 20.49 628.88 26.71 646.35 C32.93 663.83 40.55 681.06 49.41 697.26 C58.26 713.46 68.57 729.12 79.86 743.57 C91.15 758.01 103.89 771.56 117.17 783.92 C130.45 796.29 144.97 807.44 159.53 817.74 C174.09 828.04 189.40 837.13 204.53 845.72 C219.67 854.31 235.04 862.00 250.35 869.29 C265.65 876.58 280.96 883.26 296.36 889.47 C311.75 895.69 327.16 901.42 342.72 906.59 C358.28 911.76 373.92 916.44 389.74 920.49 C405.55 924.53 422.97 928.17 437.61 930.84 C452.25 933.51 464.16 935.01 477.59 936.51 C491.01 938.01 504.50 939.04 518.13 939.84 C531.76 940.64 545.44 941.02 559.36 941.32 C573.28 941.63 587.25 941.60 601.65 941.69 C616.05 941.77 630.64 941.88 645.77 941.83 C660.91 941.77 676.71 942.34 692.45 941.35 C708.19 940.37 724.60 939.16 740.23 935.90 C755.85 932.64 771.42 927.83 786.22 921.78 C801.01 915.74 815.35 908.09 828.99 899.65 C842.64 891.21 855.71 881.49 868.08 871.13 C880.44 860.78 892.18 849.41 903.17 837.52 C914.17 825.64 923.66 814.49 934.04 799.80 C944.43 785.11 956.29 766.85 965.51 749.38 C974.73 731.91 982.85 713.58 989.38 694.98 C995.91 676.38 1001.15 657.08 1004.71 637.78 C1008.26 618.49 1010.30 598.68 1010.73 579.20 C1011.17 559.73 1009.85 540.03 1007.31 520.92 C1004.76 501.81 1000.42 482.86 995.46 464.56 C990.49 446.26 984.11 428.43 977.50 411.12 C970.88 393.80 963.45 377.06 955.76 360.68 C948.08 344.31 939.93 328.42 931.38 312.87 C922.83 297.32 913.90 282.16 904.44 267.39 C894.99 252.61 885.11 238.18 874.64 224.19 C864.17 210.20 853.38 196.38 841.64 183.45 C829.90 170.53 817.53 157.92 804.22 146.65 C790.91 135.37 776.62 124.90 761.79 115.81 C746.95 106.71 731.23 98.63 715.21 92.07 C699.18 85.50 682.41 80.21 665.64 76.42 C648.87 72.64 631.57 70.40 614.60 69.36 C597.63 68.32 580.46 68.94 563.81 70.19 C547.16 71.44 530.69 74.03 514.67 76.85 C498.65 79.67 483.01 83.27 467.67 87.08 C452.32 90.89 437.36 95.12 422.59 99.70 C407.83 104.28 393.36 109.18 379.09 114.56 C364.81 119.94 350.76 125.68 336.95 131.98 C323.13 138.28 308.38 145.77 296.19 152.35 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M437.61 930.84 C452.25 933.51 464.16 935.01 477.59 936.51 C491.01 938.01 504.50 939.04 518.13 939.84 C531.76 940.64 545.44 941.02 559.36 941.32 C573.28 941.63 587.25 941.60 601.65 941.69 C616.05 941.77 630.64 941.88 645.77 941.83 C660.91 941.77 676.71 942.34 692.45 941.35 C708.19 940.37 724.60 939.16 740.23 935.90 C755.85 932.64 771.42 927.83 786.22 921.78 C801.01 915.74 815.35 908.09 828.99 899.65 C842.64 891.21 855.71 881.49 868.08 871.13 C880.44 860.78 892.18 849.41 903.17 837.52 C914.17 825.64 923.66 814.49 934.04 799.80 C944.43 785.11 956.29 766.85 965.51 749.38 C974.73 731.91 982.85 713.58 989.38 694.98 C995.91 676.38 1001.15 657.08 1004.71 637.78 C1008.26 618.49 1010.30 598.68 1010.73 579.20 C1011.17 559.73 1009.85 540.03 1007.31 520.92 C1004.76 501.81 1000.42 482.86 995.46 464.56 C990.49 446.26 984.11 428.43 977.50 411.12 C970.88 393.80 963.45 377.06 955.76 360.68 C948.08 344.31 939.93 328.42 931.38 312.87 C922.83 297.32 913.90 282.16 904.44 267.39 C894.99 252.61 885.11 238.18 874.64 224.19 C864.17 210.20 852.08 195.49 841.64 183.45 C831.20 171.42 822.25 162.20 812.03 151.99 C801.81 141.78 791.21 131.92 780.34 122.19 C769.46 112.46 758.25 103.06 746.80 93.61 C735.34 84.16 723.61 74.97 711.59 65.46 C699.58 55.95 687.41 46.34 674.70 36.56 C661.99 26.79 649.14 16.14 635.35 6.79 C621.55 -2.56 607.09 -12.12 591.94 -19.53 C576.80 -26.95 560.71 -33.08 544.46 -37.72 C528.21 -42.36 511.30 -45.42 494.45 -47.38 C477.59 -49.33 460.37 -49.89 443.32 -49.47 C426.27 -49.04 409.07 -47.39 392.14 -44.82 C375.21 -42.26 360.02 -39.31 341.73 -34.05 C323.45 -28.78 301.60 -21.54 282.43 -13.23 C263.26 -4.91 244.43 4.85 226.70 15.87 C208.97 26.90 191.83 39.38 176.03 52.95 C160.22 66.53 145.30 81.53 131.86 97.32 C118.41 113.10 106.21 130.19 95.40 147.61 C84.59 165.02 75.32 183.46 67.03 201.80 C58.74 220.14 51.91 239.02 45.66 257.68 C39.41 276.34 34.21 295.10 29.52 313.74 C24.82 332.39 20.85 350.97 17.48 369.55 C14.11 388.14 11.34 406.67 9.28 425.25 C7.23 443.82 5.79 462.40 5.16 481.01 C4.52 499.61 4.31 518.30 5.46 536.86 C6.61 555.43 8.53 574.15 12.07 592.40 C15.61 610.65 20.49 628.88 26.71 646.35 C32.93 663.83 40.55 681.06 49.41 697.26 C58.26 713.46 68.57 729.12 79.86 743.57 C91.15 758.01 103.89 771.56 117.17 783.92 C130.45 796.29 144.97 807.44 159.53 817.74 C174.09 828.04 189.40 837.13 204.53 845.72 C219.67 854.31 235.04 862.00 250.35 869.29 C265.65 876.58 280.96 883.26 296.36 889.47 C311.75 895.69 327.16 901.42 342.72 906.59 C358.28 911.76 373.92 916.44 389.74 920.49 C405.55 924.53 422.97 928.17 437.61 930.84 Z`

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
	__Parameter__000000_Reference.InsideAngle = 110.000000
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.SideLength = 135.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.440000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.100000
	__Parameter__000000_Reference.NbPitchLines = 61
	__Parameter__000000_Reference.BeatLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbBeatLines = 49
	__Parameter__000000_Reference.NbOfBeatsInTheme = 16
	__Parameter__000000_Reference.FirstVoiceShiftX = 0.060000
	__Parameter__000000_Reference.FirstVoiceShiftY = 1.610000
	__Parameter__000000_Reference.PitchDifference = 12
	__Parameter__000000_Reference.BeatsPerSecond = 10.000000
	__Parameter__000000_Reference.Level = 11.100000
	__Parameter__000000_Reference.IsMinor = true
	__Parameter__000000_Reference.ThemeBinaryEncoding = 28677
	__Parameter__000000_Reference.OriginX = 40.000000
	__Parameter__000000_Reference.OriginY = 850.000000
	__Parameter__000000_Reference.SpiralOriginX = 500.000000
	__Parameter__000000_Reference.SpiralOriginY = 500.000000
	__Parameter__000000_Reference.OriginCrossWidth = 800.000000
	__Parameter__000000_Reference.SpiralRadiusRatio = 1.060000
	__Parameter__000000_Reference.ShowSpiralBezierConstruct = false
	__Parameter__000000_Reference.ShowInterpolationPoints = false
	__Parameter__000000_Reference.ActualBeatsTemporalShift = 6
	__Parameter__000000_Reference.PathToStaticFiles = `../../../static`
	__Parameter__000000_Reference.PathToGeneratedSVG = `../../../static/images`

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 135.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.AngleDegree = -82.028057
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 110.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 135.000000
	__Rhombus__000001_Initial_Rhombus.AngleDegree = -79.106605
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 110.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 219.033677
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 30.673806
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 135.000000
	__Rhombus__000002_Rotated_Next_Rhombus.AngleDegree = -82.028057
	__Rhombus__000002_Rotated_Next_Rhombus.InsideAngle = 110.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 135.000000
	__Rhombus__000003_Rotated_Rhombus.AngleDegree = -82.028057
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 110.000000
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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 218.181434
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -26.497081
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 253.754039
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 29.238522
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 170.886663
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 260.807764
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 199.769285
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 197.087692
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 143.100000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 218.181434
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 143.100000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -26.497081
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 218.181434
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 287.850938
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -26.497081
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -70.962838
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 65.931057
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 8.007008
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 52.587135
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 149.455162
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 218.181434
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -26.497081
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 42.406126
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -120.520245
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
