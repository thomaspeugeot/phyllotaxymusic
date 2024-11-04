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

	__Axis__000000_Construction_Axis := (&models.Axis{}).Stage(stage)
	__Axis__000001_Initial_Axis := (&models.Axis{}).Stage(stage)
	__Axis__000002_Measure := (&models.Axis{}).Stage(stage)
	__Axis__000003_Pitch_Line := (&models.Axis{}).Stage(stage)
	__Axis__000004_Rotated_Axis := (&models.Axis{}).Stage(stage)

	__AxisGrid__000000_Construction_Axis_Grid := (&models.AxisGrid{}).Stage(stage)
	__AxisGrid__000001_Measure_Lines := (&models.AxisGrid{}).Stage(stage)
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
	__FrontCurve__000008_Rotated_0_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000009_Rotated_1_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000010_Rotated_1_ := (&models.FrontCurve{}).Stage(stage)

	__FrontCurveStack__000000_Backend_curve := (&models.FrontCurveStack{}).Stage(stage)
	__FrontCurveStack__000001_Front_Curve_Stack := (&models.FrontCurveStack{}).Stage(stage)
	__FrontCurveStack__000002_Hour_Handle := (&models.FrontCurveStack{}).Stage(stage)
	__FrontCurveStack__000003_Minute_Handle := (&models.FrontCurveStack{}).Stage(stage)

	__HorizontalAxis__000000_Horizontal_Axis := (&models.HorizontalAxis{}).Stage(stage)

	__Key__000000_F_key := (&models.Key{}).Stage(stage)

	__NoteInfo__000000_1 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000001_10 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000002_11 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000003_12 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000004_13 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000005_14 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000006_15 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000007_2 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000008_3 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000009_4 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000010_5 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000011_6 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000012_7 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000013_8 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000014_9 := (&models.NoteInfo{}).Stage(stage)

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

	__SpiralCircle__000000_Backend_Marker := (&models.SpiralCircle{}).Stage(stage)
	__SpiralCircle__000001_Construction_Circle_Spiral := (&models.SpiralCircle{}).Stage(stage)
	__SpiralCircle__000002_Hour_Marker := (&models.SpiralCircle{}).Stage(stage)
	__SpiralCircle__000003_Minute_Marker := (&models.SpiralCircle{}).Stage(stage)

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

	__Axis__000000_Construction_Axis.Name = `Construction Axis`
	__Axis__000000_Construction_Axis.IsDisplayed = false
	__Axis__000000_Construction_Axis.AngleDegree = 98.418226
	__Axis__000000_Construction_Axis.Length = 199.860456
	__Axis__000000_Construction_Axis.CenterX = 0.000000
	__Axis__000000_Construction_Axis.CenterY = 0.000000
	__Axis__000000_Construction_Axis.EndX = -29.259115
	__Axis__000000_Construction_Axis.EndY = 197.707122
	__Axis__000000_Construction_Axis.Color = ``
	__Axis__000000_Construction_Axis.FillOpacity = 0.000000
	__Axis__000000_Construction_Axis.Stroke = `blue`
	__Axis__000000_Construction_Axis.StrokeOpacity = 0.700000
	__Axis__000000_Construction_Axis.StrokeWidth = 2.000000
	__Axis__000000_Construction_Axis.StrokeDashArray = ``
	__Axis__000000_Construction_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000000_Construction_Axis.Transform = ``

	__Axis__000001_Initial_Axis.Name = `Initial Axis`
	__Axis__000001_Initial_Axis.IsDisplayed = false
	__Axis__000001_Initial_Axis.AngleDegree = 81.581774
	__Axis__000001_Initial_Axis.Length = 682.594160
	__Axis__000001_Initial_Axis.CenterX = 0.000000
	__Axis__000001_Initial_Axis.CenterY = 0.000000
	__Axis__000001_Initial_Axis.EndX = 0.000000
	__Axis__000001_Initial_Axis.EndY = 0.000000
	__Axis__000001_Initial_Axis.Color = ``
	__Axis__000001_Initial_Axis.FillOpacity = 0.000000
	__Axis__000001_Initial_Axis.Stroke = `black`
	__Axis__000001_Initial_Axis.StrokeOpacity = 1.000000
	__Axis__000001_Initial_Axis.StrokeWidth = 2.000000
	__Axis__000001_Initial_Axis.StrokeDashArray = ``
	__Axis__000001_Initial_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000001_Initial_Axis.Transform = ``

	__Axis__000002_Measure.Name = `Measure`
	__Axis__000002_Measure.IsDisplayed = false
	__Axis__000002_Measure.AngleDegree = 90.000000
	__Axis__000002_Measure.Length = 2000.000000
	__Axis__000002_Measure.CenterX = 0.000000
	__Axis__000002_Measure.CenterY = 0.000000
	__Axis__000002_Measure.EndX = 0.000000
	__Axis__000002_Measure.EndY = 0.000000
	__Axis__000002_Measure.Color = ``
	__Axis__000002_Measure.FillOpacity = 0.000000
	__Axis__000002_Measure.Stroke = `grey`
	__Axis__000002_Measure.StrokeOpacity = 0.800000
	__Axis__000002_Measure.StrokeWidth = 1.000000
	__Axis__000002_Measure.StrokeDashArray = ``
	__Axis__000002_Measure.StrokeDashArrayWhenSelected = ``
	__Axis__000002_Measure.Transform = ``

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
	__Axis__000004_Rotated_Axis.Length = 682.594160
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

	__AxisGrid__000000_Construction_Axis_Grid.Name = `Construction Axis Grid`
	__AxisGrid__000000_Construction_Axis_Grid.IsDisplayed = false

	__AxisGrid__000001_Measure_Lines.Name = `Measure Lines`
	__AxisGrid__000001_Measure_Lines.IsDisplayed = true

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
	__Bezier__000005_Growth_Curve_Seed.StartX = -14.629557
	__Bezier__000005_Growth_Curve_Seed.StartY = 98.853561
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 58.494015
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 109.675281
	__Bezier__000005_Growth_Curve_Seed.EndX = 104.333806
	__Bezier__000005_Growth_Curve_Seed.EndY = 217.477835
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 31.210234
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 206.656115
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
	__Circle__000001_Construction_Circle.CenterX = -14.629557
	__Circle__000001_Construction_Circle.CenterY = 98.853561
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 267.185841
	__Circle__000005_Rotated_Next_Circle.CenterY = 39.541424
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M775.25 537.75 C777.11 528.45 778.29 521.12 779.66 512.66 C781.03 504.20 782.19 495.71 783.46 487.00 C784.73 478.29 785.84 469.54 787.28 460.41 C788.72 451.27 790.09 442.08 792.10 432.18 C794.12 422.28 796.54 412.13 799.34 401.01 C802.14 389.90 806.41 377.75 808.90 365.50 C811.40 353.25 814.12 339.92 814.32 327.49 C814.51 315.07 812.89 302.64 810.07 290.95 C807.24 279.26 802.65 268.01 797.38 257.33 C792.12 246.66 785.56 236.49 778.48 226.90 C771.41 217.31 763.37 208.25 754.91 199.79 C746.45 191.32 738.55 184.12 727.71 176.11 C716.87 168.11 703.06 158.77 689.86 151.76 C676.66 144.75 662.68 138.67 648.50 134.06 C634.33 129.45 619.51 125.99 604.79 124.11 C590.08 122.23 574.91 121.75 560.20 122.78 C545.48 123.80 530.62 126.53 516.49 130.26 C502.36 134.00 488.52 139.49 475.43 145.19 C462.33 150.90 449.88 157.84 437.92 164.49 C425.96 171.14 414.69 178.23 403.66 185.09 C392.63 191.95 382.11 198.80 371.73 205.65 C361.35 212.51 351.30 219.27 341.38 226.24 C331.45 233.21 321.74 240.17 312.20 247.46 C302.65 254.75 292.41 263.01 284.09 269.98 C275.77 276.96 269.51 282.69 262.31 289.31 C255.10 295.92 248.07 302.74 240.87 309.68 C233.67 316.61 226.62 323.72 219.10 330.94 C211.59 338.16 204.14 345.50 195.80 353.01 C187.46 360.51 178.68 367.98 169.06 375.97 C159.44 383.96 148.08 391.86 138.09 400.96 C128.09 410.05 117.28 419.86 109.08 430.57 C100.87 441.27 94.13 453.13 88.84 465.18 C83.55 477.23 79.97 490.06 77.33 502.86 C74.69 515.66 73.43 528.89 73.00 541.99 C72.56 555.09 73.26 568.38 74.72 581.46 C76.18 594.54 78.00 606.14 81.76 620.48 C85.52 634.81 90.89 652.33 97.30 667.47 C103.71 682.61 111.38 697.50 120.22 711.32 C129.05 725.14 139.18 738.45 150.29 750.40 C161.41 762.36 173.83 773.46 186.91 783.06 C199.99 792.65 214.31 800.97 228.76 807.97 C243.20 814.98 258.59 820.38 273.58 825.07 C288.56 829.76 303.88 833.03 318.67 836.13 C333.45 839.22 348.03 841.47 362.29 843.66 C376.56 845.86 390.45 847.67 404.26 849.29 C418.07 850.91 431.60 852.32 445.15 853.41 C458.71 854.49 472.12 855.36 485.60 855.78 C499.09 856.20 512.59 856.66 526.05 855.95 C539.51 855.24 553.15 854.14 566.35 851.52 C579.55 848.90 592.78 845.16 605.27 840.24 C617.76 835.31 630.04 829.17 641.29 821.95 C652.54 814.72 663.28 806.21 672.78 796.88 C682.27 787.55 690.85 776.90 698.27 765.99 C705.70 755.09 711.84 743.07 717.33 731.42 C722.83 719.77 727.09 707.64 731.23 696.09 C735.37 684.54 738.75 673.14 742.17 662.12 C745.59 651.10 748.69 640.47 751.75 629.95 C754.80 619.43 757.74 609.24 760.53 599.00 C763.32 588.75 766.02 578.68 768.47 568.47 C770.92 558.27 773.38 547.05 775.25 537.75 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M284.09 269.98 C275.77 276.96 269.51 282.69 262.31 289.31 C255.10 295.92 248.07 302.74 240.87 309.68 C233.67 316.61 226.62 323.72 219.10 330.94 C211.59 338.16 204.14 345.50 195.80 353.01 C187.46 360.51 178.68 367.98 169.06 375.97 C159.44 383.96 148.08 391.86 138.09 400.96 C128.09 410.05 117.28 419.86 109.08 430.57 C100.87 441.27 94.13 453.13 88.84 465.18 C83.55 477.23 79.97 490.06 77.33 502.86 C74.69 515.66 73.43 528.89 73.00 541.99 C72.56 555.09 73.26 568.38 74.72 581.46 C76.18 594.54 78.00 606.14 81.76 620.48 C85.52 634.81 90.89 652.33 97.30 667.47 C103.71 682.61 111.38 697.50 120.22 711.32 C129.05 725.14 139.18 738.45 150.29 750.40 C161.41 762.36 173.83 773.46 186.91 783.06 C199.99 792.65 214.31 800.97 228.76 807.97 C243.20 814.98 258.59 820.38 273.58 825.07 C288.56 829.76 303.88 833.03 318.67 836.13 C333.45 839.22 348.03 841.47 362.29 843.66 C376.56 845.86 390.45 847.67 404.26 849.29 C418.07 850.91 431.60 852.32 445.15 853.41 C458.71 854.49 472.12 855.36 485.60 855.78 C499.09 856.20 513.87 856.21 526.05 855.95 C538.23 855.70 547.75 855.05 558.68 854.26 C569.61 853.46 580.52 852.35 591.63 851.20 C602.74 850.04 613.83 848.61 625.33 847.33 C636.83 846.04 648.33 844.58 660.61 843.49 C672.88 842.41 685.46 841.57 698.98 840.81 C712.50 840.06 727.29 840.42 741.74 838.95 C756.20 837.47 771.70 835.91 785.71 831.97 C799.73 828.04 813.31 822.24 825.84 815.34 C838.37 808.45 850.04 799.82 860.91 790.60 C871.79 781.38 881.82 770.92 891.08 760.02 C900.35 749.11 908.81 737.31 916.51 725.17 C924.20 713.03 930.59 701.92 937.24 687.17 C943.88 672.41 951.36 653.87 956.37 636.64 C961.39 619.40 965.18 601.53 967.31 583.76 C969.45 565.98 970.17 547.76 969.19 529.97 C968.21 512.19 965.60 494.22 961.44 477.03 C957.28 459.84 951.26 442.84 944.23 426.82 C937.21 410.80 928.36 395.43 919.29 380.90 C910.23 366.36 899.92 352.75 889.84 339.63 C879.76 326.52 869.23 314.22 858.83 302.21 C848.43 290.20 838.00 278.76 827.45 267.57 C816.90 256.38 806.36 245.58 795.54 235.07 C784.71 224.57 773.79 214.37 762.49 204.55 C751.18 194.72 739.81 184.91 727.71 176.11 C715.60 167.31 703.06 158.77 689.86 151.76 C676.66 144.75 662.68 138.67 648.50 134.06 C634.33 129.45 619.51 125.99 604.79 124.11 C590.08 122.23 574.91 121.75 560.20 122.78 C545.48 123.80 530.62 126.53 516.49 130.26 C502.36 134.00 488.52 139.49 475.43 145.19 C462.33 150.90 449.88 157.84 437.92 164.49 C425.96 171.14 414.69 178.23 403.66 185.09 C392.63 191.95 382.11 198.80 371.73 205.65 C361.35 212.51 351.30 219.27 341.38 226.24 C331.45 233.21 321.74 240.17 312.20 247.46 C302.65 254.75 292.41 263.01 284.09 269.98 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M526.05 855.95 C538.23 855.70 547.75 855.05 558.68 854.26 C569.61 853.46 580.52 852.35 591.63 851.20 C602.74 850.04 613.83 848.61 625.33 847.33 C636.83 846.04 648.33 844.58 660.61 843.49 C672.88 842.41 685.46 841.57 698.98 840.81 C712.50 840.06 727.29 840.42 741.74 838.95 C756.20 837.47 771.70 835.91 785.71 831.97 C799.73 828.04 813.31 822.24 825.84 815.34 C838.37 808.45 850.04 799.82 860.91 790.60 C871.79 781.38 881.82 770.92 891.08 760.02 C900.35 749.11 908.81 737.31 916.51 725.17 C924.20 713.03 930.59 701.92 937.24 687.17 C943.88 672.41 951.36 653.87 956.37 636.64 C961.39 619.40 965.18 601.53 967.31 583.76 C969.45 565.98 970.17 547.76 969.19 529.97 C968.21 512.19 965.60 494.22 961.44 477.03 C957.28 459.84 951.26 442.84 944.23 426.82 C937.21 410.80 928.36 395.43 919.29 380.90 C910.23 366.36 899.92 352.75 889.84 339.63 C879.76 326.52 869.23 314.22 858.83 302.21 C848.43 290.20 838.00 278.76 827.45 267.57 C816.90 256.38 806.36 245.58 795.54 235.07 C784.71 224.57 773.79 214.37 762.49 204.55 C751.18 194.72 738.41 184.38 727.71 176.11 C717.01 167.84 708.32 161.79 698.30 154.91 C688.29 148.03 678.05 141.46 667.62 134.85 C657.18 128.23 646.56 121.90 635.70 115.24 C624.85 108.59 613.85 102.12 602.48 94.92 C591.11 87.72 579.66 80.19 567.51 72.06 C555.36 63.93 542.92 54.18 529.59 46.12 C516.25 38.06 502.04 29.46 487.49 23.70 C472.94 17.94 457.55 13.97 442.29 11.55 C427.04 9.14 411.36 8.68 395.95 9.23 C380.53 9.79 365.00 11.89 349.79 14.88 C334.59 17.87 319.46 22.11 304.73 27.17 C290.01 32.22 277.08 37.26 261.44 45.20 C245.80 53.14 226.88 63.70 210.88 74.82 C194.87 85.94 179.45 98.42 165.41 111.92 C151.37 125.42 138.19 140.22 126.66 155.80 C115.13 171.38 104.82 188.17 96.27 205.37 C87.73 222.57 80.83 240.83 75.41 259.00 C70.00 277.17 66.49 296.03 63.77 314.40 C61.04 332.77 59.94 351.26 59.07 369.22 C58.19 387.18 58.27 404.83 58.51 422.18 C58.75 439.53 59.46 456.48 60.50 473.30 C61.54 490.13 62.91 506.66 64.77 523.13 C66.63 539.60 68.85 555.89 71.68 572.11 C74.51 588.34 77.49 604.58 81.76 620.48 C86.03 636.37 90.89 652.33 97.30 667.47 C103.71 682.61 111.38 697.50 120.22 711.32 C129.05 725.14 139.18 738.45 150.29 750.40 C161.41 762.36 173.83 773.46 186.91 783.06 C199.99 792.65 214.31 800.97 228.76 807.97 C243.20 814.98 258.59 820.38 273.58 825.07 C288.56 829.76 303.88 833.03 318.67 836.13 C333.45 839.22 348.03 841.47 362.29 843.66 C376.56 845.86 390.45 847.67 404.26 849.29 C418.07 850.91 431.60 852.32 445.15 853.41 C458.71 854.49 472.12 855.36 485.60 855.78 C499.09 856.20 513.87 856.21 526.05 855.95 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M727.71 176.11 C717.01 167.84 708.32 161.79 698.30 154.91 C688.29 148.03 678.05 141.46 667.62 134.85 C657.18 128.23 646.56 121.90 635.70 115.24 C624.85 108.59 613.85 102.12 602.48 94.92 C591.11 87.72 579.66 80.19 567.51 72.06 C555.36 63.93 542.92 54.18 529.59 46.12 C516.25 38.06 502.04 29.46 487.49 23.70 C472.94 17.94 457.55 13.97 442.29 11.55 C427.04 9.14 411.36 8.68 395.95 9.23 C380.53 9.79 365.00 11.89 349.79 14.88 C334.59 17.87 319.46 22.11 304.73 27.17 C290.01 32.22 277.08 37.26 261.44 45.20 C245.80 53.14 226.88 63.70 210.88 74.82 C194.87 85.94 179.45 98.42 165.41 111.92 C151.37 125.42 138.19 140.22 126.66 155.80 C115.13 171.38 104.82 188.17 96.27 205.37 C87.73 222.57 80.83 240.83 75.41 259.00 C70.00 277.17 66.49 296.03 63.77 314.40 C61.04 332.77 59.94 351.26 59.07 369.22 C58.19 387.18 58.27 404.83 58.51 422.18 C58.75 439.53 59.46 456.48 60.50 473.30 C61.54 490.13 62.91 506.66 64.77 523.13 C66.63 539.60 68.85 555.89 71.68 572.11 C74.51 588.34 78.31 606.02 81.76 620.48 C85.21 634.94 88.51 646.11 92.39 658.88 C96.27 671.65 100.57 684.31 105.04 697.09 C109.51 709.87 114.35 722.54 119.22 735.57 C124.09 748.60 129.24 761.55 134.26 775.25 C139.27 788.96 144.20 802.96 149.31 817.81 C154.43 832.67 158.80 848.92 164.96 864.39 C171.13 879.86 177.68 896.26 186.31 910.62 C194.93 924.99 205.37 938.46 216.72 950.59 C228.07 962.71 241.00 973.55 254.40 983.38 C267.79 993.22 282.28 1001.88 297.09 1009.60 C311.89 1017.31 327.46 1023.96 343.22 1029.68 C358.98 1035.39 373.21 1039.92 391.67 1043.89 C410.12 1047.86 433.07 1051.87 453.95 1053.49 C474.82 1055.10 496.11 1055.21 516.93 1053.56 C537.76 1051.90 558.75 1048.60 578.92 1043.56 C599.09 1038.52 619.09 1031.66 637.96 1023.29 C656.83 1014.93 675.10 1004.61 692.13 993.35 C709.16 982.10 725.14 969.02 740.16 955.76 C755.17 942.51 768.97 928.06 782.22 913.81 C795.46 899.56 807.73 884.88 819.62 870.28 C831.51 855.67 842.74 841.01 853.56 826.18 C864.38 811.35 874.71 796.48 884.54 781.31 C894.37 766.13 903.75 750.81 912.54 735.12 C921.32 719.43 929.93 703.58 937.24 687.17 C944.54 670.75 951.36 653.87 956.37 636.64 C961.39 619.40 965.18 601.53 967.31 583.76 C969.45 565.98 970.17 547.76 969.19 529.97 C968.21 512.19 965.60 494.22 961.44 477.03 C957.28 459.84 951.26 442.84 944.23 426.82 C937.21 410.80 928.36 395.43 919.29 380.90 C910.23 366.36 899.92 352.75 889.84 339.63 C879.76 326.52 869.23 314.22 858.83 302.21 C848.43 290.20 838.00 278.76 827.45 267.57 C816.90 256.38 806.36 245.58 795.54 235.07 C784.71 224.57 773.79 214.37 762.49 204.55 C751.18 194.72 738.41 184.38 727.71 176.11 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M81.76 620.48 C85.21 634.94 88.51 646.11 92.39 658.88 C96.27 671.65 100.57 684.31 105.04 697.09 C109.51 709.87 114.35 722.54 119.22 735.57 C124.09 748.60 129.24 761.55 134.26 775.25 C139.27 788.96 144.20 802.96 149.31 817.81 C154.43 832.67 158.80 848.92 164.96 864.39 C171.13 879.86 177.68 896.26 186.31 910.62 C194.93 924.99 205.37 938.46 216.72 950.59 C228.07 962.71 241.00 973.55 254.40 983.38 C267.79 993.22 282.28 1001.88 297.09 1009.60 C311.89 1017.31 327.46 1023.96 343.22 1029.68 C358.98 1035.39 373.21 1039.92 391.67 1043.89 C410.12 1047.86 433.07 1051.87 453.95 1053.49 C474.82 1055.10 496.11 1055.21 516.93 1053.56 C537.76 1051.90 558.75 1048.60 578.92 1043.56 C599.09 1038.52 619.09 1031.66 637.96 1023.29 C656.83 1014.93 675.10 1004.61 692.13 993.35 C709.16 982.10 725.14 969.02 740.16 955.76 C755.17 942.51 768.97 928.06 782.22 913.81 C795.46 899.56 807.73 884.88 819.62 870.28 C831.51 855.67 842.74 841.01 853.56 826.18 C864.38 811.35 874.71 796.48 884.54 781.31 C894.37 766.13 903.75 750.81 912.54 735.12 C921.32 719.43 930.25 701.81 937.24 687.17 C944.23 672.52 949.07 660.76 954.49 647.25 C959.91 633.74 964.89 620.04 969.75 606.12 C974.61 592.19 979.06 578.11 983.66 563.73 C988.26 549.35 992.54 534.83 997.34 519.82 C1002.15 504.81 1007.15 489.62 1012.49 473.65 C1017.83 457.69 1024.58 441.10 1029.39 424.00 C1034.19 406.91 1039.23 388.85 1041.32 371.08 C1043.41 353.31 1043.48 335.10 1041.94 317.38 C1040.41 299.67 1036.77 281.97 1032.11 264.78 C1027.44 247.59 1021.12 230.64 1013.93 214.26 C1006.73 197.88 998.22 181.87 988.92 166.50 C979.62 151.12 970.83 137.76 958.15 122.01 C945.47 106.25 929.10 87.43 912.85 71.97 C896.60 56.50 879.03 42.00 860.65 29.24 C842.27 16.47 822.67 4.95 802.56 -4.63 C782.46 -14.21 761.29 -22.21 740.02 -28.23 C718.75 -34.26 696.68 -38.32 674.96 -40.78 C653.23 -43.25 631.14 -43.55 609.68 -43.02 C588.22 -42.49 566.92 -40.19 546.18 -37.61 C525.45 -35.02 505.17 -31.40 485.26 -27.51 C465.34 -23.62 445.89 -19.18 426.68 -14.27 C407.47 -9.36 388.61 -4.03 369.98 1.95 C351.35 7.93 332.99 14.39 314.91 21.59 C296.81 28.80 278.78 36.33 261.44 45.20 C244.10 54.07 226.88 63.70 210.88 74.82 C194.87 85.94 179.45 98.42 165.41 111.92 C151.37 125.42 138.19 140.22 126.66 155.80 C115.13 171.38 104.82 188.17 96.27 205.37 C87.73 222.57 80.83 240.83 75.41 259.00 C70.00 277.17 66.49 296.03 63.77 314.40 C61.04 332.77 59.94 351.26 59.07 369.22 C58.19 387.18 58.27 404.83 58.51 422.18 C58.75 439.53 59.46 456.48 60.50 473.30 C61.54 490.13 62.91 506.66 64.77 523.13 C66.63 539.60 68.85 555.89 71.68 572.11 C74.51 588.34 78.31 606.02 81.76 620.48 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M937.24 687.17 C944.23 672.52 949.07 660.76 954.49 647.25 C959.91 633.74 964.89 620.04 969.75 606.12 C974.61 592.19 979.06 578.11 983.66 563.73 C988.26 549.35 992.54 534.83 997.34 519.82 C1002.15 504.81 1007.15 489.62 1012.49 473.65 C1017.83 457.69 1024.58 441.10 1029.39 424.00 C1034.19 406.91 1039.23 388.85 1041.32 371.08 C1043.41 353.31 1043.48 335.10 1041.94 317.38 C1040.41 299.67 1036.77 281.97 1032.11 264.78 C1027.44 247.59 1021.12 230.64 1013.93 214.26 C1006.73 197.88 998.22 181.87 988.92 166.50 C979.62 151.12 970.83 137.76 958.15 122.01 C945.47 106.25 929.10 87.43 912.85 71.97 C896.60 56.50 879.03 42.00 860.65 29.24 C842.27 16.47 822.67 4.95 802.56 -4.63 C782.46 -14.21 761.29 -22.21 740.02 -28.23 C718.75 -34.26 696.68 -38.32 674.96 -40.78 C653.23 -43.25 631.14 -43.55 609.68 -43.02 C588.22 -42.49 566.92 -40.19 546.18 -37.61 C525.45 -35.02 505.17 -31.40 485.26 -27.51 C465.34 -23.62 445.89 -19.18 426.68 -14.27 C407.47 -9.36 388.61 -4.03 369.98 1.95 C351.35 7.93 332.99 14.39 314.91 21.59 C296.81 28.80 277.29 37.59 261.44 45.20 C245.59 52.81 233.51 59.47 219.78 67.24 C206.04 75.01 192.54 83.26 179.04 91.80 C165.52 100.34 152.24 109.32 138.69 118.47 C125.13 127.62 111.77 137.11 97.71 146.69 C83.66 156.26 69.37 165.87 54.37 175.89 C39.38 185.91 22.94 195.54 7.75 206.81 C-7.43 218.07 -23.35 229.97 -36.73 243.51 C-50.11 257.05 -62.13 272.24 -72.54 288.07 C-82.94 303.90 -91.65 321.07 -99.17 338.50 C-106.69 355.93 -112.73 374.24 -117.67 392.67 C-122.60 411.09 -126.23 430.08 -128.79 449.06 C-131.35 468.05 -132.99 485.02 -133.04 506.58 C-133.08 528.14 -132.14 554.72 -129.06 578.44 C-125.99 602.16 -121.20 625.97 -114.61 648.92 C-108.01 671.86 -99.61 694.60 -89.50 716.11 C-79.40 737.61 -67.39 758.55 -53.99 777.96 C-40.58 797.38 -25.23 815.75 -9.09 832.61 C7.06 849.47 24.93 864.84 42.89 879.12 C60.86 893.39 79.89 906.13 98.71 918.24 C117.53 930.34 136.71 941.29 155.82 951.74 C174.94 962.19 194.09 971.85 213.39 980.93 C232.70 990.00 252.04 998.45 271.64 1006.18 C291.23 1013.91 310.96 1021.03 330.97 1027.31 C350.97 1033.60 371.17 1039.53 391.67 1043.89 C412.16 1048.25 433.07 1051.87 453.95 1053.49 C474.82 1055.10 496.11 1055.21 516.93 1053.56 C537.76 1051.90 558.75 1048.60 578.92 1043.56 C599.09 1038.52 619.09 1031.66 637.96 1023.29 C656.83 1014.93 675.10 1004.61 692.13 993.35 C709.16 982.10 725.14 969.02 740.16 955.76 C755.17 942.51 768.97 928.06 782.22 913.81 C795.46 899.56 807.73 884.88 819.62 870.28 C831.51 855.67 842.74 841.01 853.56 826.18 C864.38 811.35 874.71 796.48 884.54 781.31 C894.37 766.13 903.75 750.81 912.54 735.12 C921.32 719.43 930.25 701.81 937.24 687.17 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M261.44 45.20 C245.59 52.81 233.51 59.47 219.78 67.24 C206.04 75.01 192.54 83.26 179.04 91.80 C165.52 100.34 152.24 109.32 138.69 118.47 C125.13 127.62 111.77 137.11 97.71 146.69 C83.66 156.26 69.37 165.87 54.37 175.89 C39.38 185.91 22.94 195.54 7.75 206.81 C-7.43 218.07 -23.35 229.97 -36.73 243.51 C-50.11 257.05 -62.13 272.24 -72.54 288.07 C-82.94 303.90 -91.65 321.07 -99.17 338.50 C-106.69 355.93 -112.73 374.24 -117.67 392.67 C-122.60 411.09 -126.23 430.08 -128.79 449.06 C-131.35 468.05 -132.99 485.02 -133.04 506.58 C-133.08 528.14 -132.14 554.72 -129.06 578.44 C-125.99 602.16 -121.20 625.97 -114.61 648.92 C-108.01 671.86 -99.61 694.60 -89.50 716.11 C-79.40 737.61 -67.39 758.55 -53.99 777.96 C-40.58 797.38 -25.23 815.75 -9.09 832.61 C7.06 849.47 24.93 864.84 42.89 879.12 C60.86 893.39 79.89 906.13 98.71 918.24 C117.53 930.34 136.71 941.29 155.82 951.74 C174.94 962.19 194.09 971.85 213.39 980.93 C232.70 990.00 252.04 998.45 271.64 1006.18 C291.23 1013.91 310.96 1021.03 330.97 1027.31 C350.97 1033.60 373.26 1039.50 391.67 1043.89 C410.07 1048.28 424.68 1050.88 441.40 1053.66 C458.13 1056.44 474.97 1058.64 492.00 1060.57 C509.02 1062.50 526.13 1063.88 543.54 1065.25 C560.96 1066.61 578.43 1067.52 596.48 1068.74 C614.52 1069.95 632.77 1071.21 651.81 1072.54 C670.85 1073.87 690.77 1076.35 710.74 1076.71 C730.71 1077.08 751.64 1077.36 771.62 1074.73 C791.61 1072.10 811.58 1067.30 830.67 1060.93 C849.76 1054.57 868.36 1046.06 886.17 1036.55 C903.98 1027.04 921.17 1015.88 937.54 1003.90 C953.91 991.91 969.58 978.62 984.38 964.63 C999.19 950.63 1011.88 937.71 1026.36 919.93 C1040.83 902.15 1057.85 879.59 1071.24 857.97 C1084.63 836.35 1096.66 813.52 1106.68 790.21 C1116.69 766.89 1125.10 742.54 1131.32 718.06 C1137.55 693.58 1141.86 668.31 1144.05 643.32 C1146.23 618.34 1146.17 592.91 1144.44 568.15 C1142.72 543.40 1138.64 518.69 1133.70 494.79 C1128.76 470.88 1121.95 447.49 1114.81 424.74 C1107.67 401.99 1099.44 379.91 1090.86 358.28 C1082.29 336.66 1073.09 315.63 1063.34 295.01 C1053.59 274.38 1043.32 254.24 1032.33 234.54 C1021.35 214.83 1009.79 195.55 997.42 176.80 C985.06 158.04 972.24 139.48 958.15 122.01 C944.05 104.54 929.10 87.43 912.85 71.97 C896.60 56.50 879.03 42.00 860.65 29.24 C842.27 16.47 822.67 4.95 802.56 -4.63 C782.46 -14.21 761.29 -22.21 740.02 -28.23 C718.75 -34.26 696.68 -38.32 674.96 -40.78 C653.23 -43.25 631.14 -43.55 609.68 -43.02 C588.22 -42.49 566.92 -40.19 546.18 -37.61 C525.45 -35.02 505.17 -31.40 485.26 -27.51 C465.34 -23.62 445.89 -19.18 426.68 -14.27 C407.47 -9.36 388.61 -4.03 369.98 1.95 C351.35 7.93 332.99 14.39 314.91 21.59 C296.81 28.80 277.29 37.59 261.44 45.20 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M391.67 1043.89 C410.07 1048.28 424.68 1050.88 441.40 1053.66 C458.13 1056.44 474.97 1058.64 492.00 1060.57 C509.02 1062.50 526.13 1063.88 543.54 1065.25 C560.96 1066.61 578.43 1067.52 596.48 1068.74 C614.52 1069.95 632.77 1071.21 651.81 1072.54 C670.85 1073.87 690.77 1076.35 710.74 1076.71 C730.71 1077.08 751.64 1077.36 771.62 1074.73 C791.61 1072.10 811.58 1067.30 830.67 1060.93 C849.76 1054.57 868.36 1046.06 886.17 1036.55 C903.98 1027.04 921.17 1015.88 937.54 1003.90 C953.91 991.91 969.58 978.62 984.38 964.63 C999.19 950.63 1011.88 937.71 1026.36 919.93 C1040.83 902.15 1057.85 879.59 1071.24 857.97 C1084.63 836.35 1096.66 813.52 1106.68 790.21 C1116.69 766.89 1125.10 742.54 1131.32 718.06 C1137.55 693.58 1141.86 668.31 1144.05 643.32 C1146.23 618.34 1146.17 592.91 1144.44 568.15 C1142.72 543.40 1138.64 518.69 1133.70 494.79 C1128.76 470.88 1121.95 447.49 1114.81 424.74 C1107.67 401.99 1099.44 379.91 1090.86 358.28 C1082.29 336.66 1073.09 315.63 1063.34 295.01 C1053.59 274.38 1043.32 254.24 1032.33 234.54 C1021.35 214.83 1009.79 195.55 997.42 176.80 C985.06 158.04 970.50 138.08 958.15 122.01 C945.80 105.94 935.38 93.93 923.32 80.37 C911.26 66.81 898.69 53.65 885.79 40.64 C872.88 27.62 859.53 15.02 845.90 2.29 C832.26 -10.45 818.26 -22.83 803.99 -35.76 C789.72 -48.68 775.28 -61.73 760.26 -75.28 C745.23 -88.83 730.22 -103.77 713.82 -117.06 C697.43 -130.34 680.21 -144.12 661.88 -154.99 C643.55 -165.86 623.84 -174.97 603.83 -182.27 C583.82 -189.56 562.83 -194.83 541.82 -198.76 C520.81 -202.70 499.23 -204.91 477.78 -205.91 C456.32 -206.90 434.58 -206.39 413.08 -204.73 C391.58 -203.07 372.54 -200.90 348.79 -195.94 C325.03 -190.99 296.01 -183.81 270.58 -175.00 C245.15 -166.20 220.02 -155.53 196.20 -143.12 C172.38 -130.71 149.20 -116.40 127.67 -100.55 C106.15 -84.69 85.66 -66.92 67.08 -47.95 C48.52 -28.95 31.54 -7.90 16.00 14.36 C-0.06 36.04 -13.68 58.88 -25.92 81.69 C-38.11 104.50 -48.45 128.18 -58.03 151.60 C-67.62 175.01 -75.86 198.66 -83.46 222.19 C-91.06 245.72 -97.74 269.23 -103.64 292.79 C-109.54 316.36 -114.67 339.91 -118.86 363.59 C-123.06 387.26 -126.47 410.99 -128.83 434.83 C-131.19 458.66 -133.00 482.64 -133.04 506.58 C-133.08 530.51 -132.14 554.72 -129.06 578.44 C-125.99 602.16 -121.20 625.97 -114.61 648.92 C-108.01 671.86 -99.61 694.60 -89.50 716.11 C-79.40 737.61 -67.39 758.55 -53.99 777.96 C-40.58 797.38 -25.23 815.75 -9.09 832.61 C7.06 849.47 24.93 864.84 42.89 879.12 C60.86 893.39 79.89 906.13 98.71 918.24 C117.53 930.34 136.71 941.29 155.82 951.74 C174.94 962.19 194.09 971.85 213.39 980.93 C232.70 990.00 252.04 998.45 271.64 1006.18 C291.23 1013.91 310.96 1021.03 330.97 1027.31 C350.97 1033.60 373.26 1039.50 391.67 1043.89 Z`

	__FrontCurve__000008_Rotated_0_.Name = `Rotated 0 `
	__FrontCurve__000008_Rotated_0_.Path = `M775.42 537.31 C777.28 527.99 778.46 520.64 779.82 512.15 C781.19 503.67 782.34 495.15 783.61 486.42 C784.87 477.69 785.97 468.92 787.40 459.76 C788.84 450.61 790.20 441.39 792.20 431.48 C794.20 421.56 796.61 411.38 799.40 400.25 C802.18 389.12 806.43 376.96 808.91 364.69 C811.39 352.42 814.09 339.07 814.27 326.63 C814.44 314.20 812.80 301.76 809.95 290.05 C807.11 278.35 802.49 267.08 797.20 256.39 C791.92 245.70 785.34 235.53 778.24 225.92 C771.13 216.32 763.07 207.24 754.59 198.77 C746.10 190.29 738.18 183.07 727.31 175.06 C716.44 167.04 702.59 157.69 689.36 150.67 C676.12 143.65 662.11 137.56 647.90 132.93 C633.69 128.31 618.83 124.84 604.08 122.94 C589.33 121.05 574.13 120.56 559.37 121.58 C544.62 122.59 529.73 125.30 515.56 129.02 C501.40 132.74 487.53 138.21 474.40 143.90 C461.27 149.59 448.79 156.52 436.80 163.15 C424.81 169.79 413.51 176.87 402.45 183.72 C391.40 190.57 380.86 197.40 370.45 204.25 C360.05 211.11 349.98 217.86 340.04 224.82 C330.10 231.79 320.38 238.75 310.82 246.04 C301.26 253.33 291.00 261.59 282.68 268.56 C274.36 275.54 268.09 281.28 260.89 287.90 C253.68 294.52 246.65 301.35 239.45 308.29 C232.25 315.24 225.20 322.36 217.70 329.59 C210.19 336.83 202.75 344.18 194.42 351.70 C186.09 359.21 177.32 366.70 167.72 374.71 C158.11 382.72 146.77 390.62 136.80 399.74 C126.82 408.86 116.04 418.69 107.85 429.42 C99.66 440.15 92.93 452.03 87.65 464.11 C82.38 476.19 78.81 489.05 76.18 501.88 C73.55 514.71 72.30 527.96 71.87 541.09 C71.44 554.22 72.15 567.54 73.61 580.65 C75.08 593.76 76.90 605.39 80.67 619.76 C84.44 634.13 89.81 651.69 96.23 666.87 C102.65 682.04 110.33 696.96 119.18 710.82 C128.02 724.67 138.16 738.01 149.28 750.00 C160.41 761.99 172.85 773.13 185.94 782.75 C199.03 792.38 213.37 800.73 227.82 807.76 C242.28 814.79 257.69 820.22 272.69 824.94 C287.69 829.66 303.03 832.95 317.83 836.07 C332.64 839.19 347.24 841.46 361.53 843.67 C375.82 845.89 389.73 847.71 403.57 849.35 C417.40 850.99 430.95 852.41 444.54 853.51 C458.12 854.60 471.56 855.48 485.07 855.91 C498.58 856.34 512.11 856.80 525.60 856.10 C539.09 855.39 552.75 854.29 565.98 851.67 C579.21 849.05 592.46 845.31 604.98 840.38 C617.49 835.45 629.80 829.31 641.08 822.07 C652.35 814.84 663.12 806.32 672.63 796.98 C682.15 787.64 690.75 776.98 698.20 766.06 C705.64 755.13 711.81 743.10 717.32 731.44 C722.83 719.77 727.12 707.62 731.27 696.05 C735.43 684.47 738.82 673.06 742.26 662.01 C745.69 650.97 748.80 640.32 751.87 629.77 C754.94 619.22 757.88 609.01 760.68 598.73 C763.47 588.46 766.18 578.36 768.64 568.13 C771.09 557.89 773.55 546.64 775.42 537.31 Z`

	__FrontCurve__000009_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000009_Rotated_1_.Path = `M395.80 845.09 C407.41 849.20 416.63 852.04 427.25 855.25 C437.87 858.42 448.58 861.30 459.50 864.23 C470.41 867.11 481.40 869.75 492.74 872.68 C504.06 875.57 515.43 878.30 527.44 881.68 C539.41 885.03 551.55 888.73 564.61 892.84 C577.62 896.92 591.41 902.52 605.59 906.29 C619.71 910.03 634.89 914.06 649.52 915.36 C664.10 916.64 678.99 916.01 693.28 914.01 C707.53 911.98 721.64 908.02 735.20 903.24 C748.74 898.44 761.95 892.18 774.61 885.24 C787.24 878.28 799.46 870.19 811.09 861.53 C822.68 852.84 832.69 844.66 844.28 833.17 C855.84 821.66 869.54 806.86 880.49 792.45 C891.39 778.00 901.39 762.52 909.83 746.56 C918.22 730.57 925.47 713.65 930.99 696.57 C936.45 679.45 940.48 661.57 942.80 643.90 C945.06 626.19 945.55 608.00 944.76 590.40 C943.91 572.75 941.15 555.06 937.90 538.11 C934.60 521.12 929.83 504.57 925.11 488.57 C920.34 472.54 914.87 457.15 909.44 442.06 C903.95 426.95 898.26 412.40 892.38 398.02 C886.44 383.64 880.40 369.65 874.00 355.81 C867.54 341.99 860.91 328.43 853.80 315.06 C846.64 301.71 838.31 287.34 831.19 275.67 C824.03 264.02 818.01 255.17 811.03 245.06 C804.03 234.97 796.73 225.08 789.25 215.07 C781.75 205.10 774.00 195.31 766.13 185.11 C758.25 174.95 750.19 164.92 742.02 154.04 C733.84 143.21 725.73 132.03 717.15 120.01 C708.58 108.04 700.33 94.45 690.61 82.09 C680.91 69.78 670.56 56.64 658.89 46.01 C647.24 35.42 634.13 26.18 620.61 18.43 C607.10 10.72 592.48 4.67 577.75 -0.36 C563.03 -5.35 547.62 -8.95 532.22 -11.63 C516.83 -14.26 501.04 -15.71 485.36 -16.28 C469.68 -16.81 455.69 -16.71 438.11 -14.94 C420.54 -13.11 398.95 -9.99 379.91 -5.36 C360.88 -0.67 341.91 5.47 323.90 13.04 C305.90 20.67 288.22 29.82 271.82 40.27 C255.45 50.78 239.74 62.86 225.54 75.93 C211.36 89.06 198.30 103.77 186.66 118.88 C175.04 134.06 164.92 150.55 155.72 166.83 C146.53 183.17 138.81 200.17 131.47 216.74 C124.16 233.38 117.85 250.03 111.78 266.44 C105.75 282.91 100.29 299.13 95.16 315.36 C90.09 331.63 85.41 347.69 81.19 363.89 C77.02 380.12 73.24 396.28 70.03 412.61 C66.88 428.95 63.83 445.35 62.10 461.89 C60.43 478.42 59.26 495.24 59.81 511.83 C60.42 528.43 62.29 545.25 65.59 561.47 C68.95 577.69 73.69 593.92 79.80 609.21 C85.98 624.50 93.67 639.48 102.47 653.26 C111.34 667.04 121.81 680.09 132.83 691.93 C143.92 703.77 156.44 714.46 168.81 724.33 C181.25 734.20 194.48 742.85 207.24 751.15 C220.09 759.45 232.99 766.86 245.62 774.14 C258.31 781.39 270.74 788.13 283.17 794.70 C295.64 801.23 307.88 807.46 320.28 813.42 C332.72 819.33 345.06 825.00 357.64 830.31 C370.24 835.56 384.20 840.92 395.80 845.09 Z`

	__FrontCurve__000010_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000010_Rotated_1_.Path = `M356.71 989.95 C373.29 995.32 386.41 999.01 401.54 1003.08 C416.67 1006.97 431.89 1010.42 447.40 1013.79 C462.87 1016.97 478.39 1019.77 494.37 1022.72 C510.26 1025.49 526.15 1027.97 542.83 1030.94 C559.35 1033.75 575.96 1036.78 593.68 1040.06 C611.16 1043.19 629.42 1047.69 648.13 1050.19 C666.58 1052.59 686.20 1055.08 705.17 1054.73 C723.91 1054.27 743.09 1051.70 761.52 1047.67 C779.77 1043.52 797.93 1037.25 815.42 1030.08 C832.74 1022.80 849.77 1013.88 866.09 1004.21 C882.28 994.42 898.03 983.36 913.05 971.65 C927.94 959.82 940.79 948.75 955.87 933.53 C970.78 918.15 988.44 898.44 1002.75 879.48 C1016.84 860.36 1029.92 840.01 1041.17 819.18 C1052.20 798.18 1061.95 776.07 1069.70 753.86 C1077.23 731.46 1083.15 708.14 1087.09 685.14 C1090.81 661.96 1092.55 638.17 1092.78 615.10 C1092.76 591.84 1090.61 568.49 1087.76 545.99 C1084.65 523.30 1079.85 501.10 1074.90 479.52 C1069.67 457.79 1063.54 436.82 1057.25 416.18 C1050.67 395.45 1043.69 375.43 1036.32 355.60 C1028.65 335.77 1020.72 316.48 1012.21 297.38 C1003.43 278.34 994.30 259.71 984.50 241.33 C974.44 223.08 962.74 203.47 952.67 187.58 C942.44 171.85 933.92 159.87 923.97 146.25 C913.91 132.78 903.45 119.60 892.69 106.37 C881.84 93.29 870.66 80.52 859.19 67.39 C847.68 54.46 835.94 41.78 823.89 28.28 C811.85 15.01 799.83 1.54 787.11 -12.75 C774.46 -26.80 762.04 -42.34 748.01 -56.57 C734.09 -70.56 719.31 -85.30 703.23 -97.37 C687.25 -109.22 669.59 -119.66 651.57 -128.43 C633.62 -137.00 614.39 -143.81 595.12 -149.44 C575.88 -154.88 555.85 -158.81 535.88 -161.67 C515.94 -164.34 495.54 -165.70 475.29 -166.03 C455.08 -166.19 437.06 -165.64 414.42 -163.14 C391.82 -160.42 364.06 -155.97 339.55 -149.92 C315.10 -143.60 290.68 -135.61 267.41 -126.01 C244.21 -116.15 221.33 -104.54 199.97 -91.50 C178.67 -78.18 158.08 -63.06 139.26 -46.87 C120.51 -30.40 102.95 -12.13 87.08 6.59 C71.28 25.62 57.13 46.13 44.13 66.47 C31.21 87.13 19.96 108.54 9.29 129.60 C-1.25 150.96 -10.55 172.36 -19.45 193.63 C-28.16 215.15 -36.06 236.44 -43.45 257.84 C-50.61 279.43 -57.15 300.82 -63.04 322.44 C-68.67 344.18 -73.71 365.86 -77.97 387.77 C-81.96 409.75 -85.62 431.81 -87.76 454.01 C-89.62 476.21 -90.77 498.78 -89.98 521.01 C-88.91 543.25 -86.37 565.80 -82.19 587.57 C-77.74 609.35 -71.68 631.19 -64.07 651.87 C-56.18 672.55 -46.54 692.96 -35.66 711.90 C-24.48 730.86 -11.46 749.06 2.20 765.77 C16.18 782.51 31.78 797.99 47.33 812.42 C63.21 826.84 80.00 839.86 96.46 852.36 C113.25 864.81 130.17 876.14 146.95 887.19 C164.04 898.12 180.90 908.30 197.91 918.17 C215.16 927.86 232.20 937.02 249.54 945.75 C267.07 954.26 284.53 962.30 302.36 969.80 C320.30 977.05 340.17 984.36 356.71 989.95 Z`

	__FrontCurveStack__000000_Backend_curve.Name = `Backend curve`
	__FrontCurveStack__000000_Backend_curve.IsDisplayed = false
	__FrontCurveStack__000000_Backend_curve.Color = ``
	__FrontCurveStack__000000_Backend_curve.FillOpacity = 0.000000
	__FrontCurveStack__000000_Backend_curve.Stroke = `green`
	__FrontCurveStack__000000_Backend_curve.StrokeOpacity = 1.000000
	__FrontCurveStack__000000_Backend_curve.StrokeWidth = 1.000000
	__FrontCurveStack__000000_Backend_curve.StrokeDashArray = ``
	__FrontCurveStack__000000_Backend_curve.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__000000_Backend_curve.Transform = ``

	__FrontCurveStack__000001_Front_Curve_Stack.Name = `Front Curve Stack`
	__FrontCurveStack__000001_Front_Curve_Stack.IsDisplayed = false
	__FrontCurveStack__000001_Front_Curve_Stack.Color = ``
	__FrontCurveStack__000001_Front_Curve_Stack.FillOpacity = 0.000000
	__FrontCurveStack__000001_Front_Curve_Stack.Stroke = `blue`
	__FrontCurveStack__000001_Front_Curve_Stack.StrokeOpacity = 1.000000
	__FrontCurveStack__000001_Front_Curve_Stack.StrokeWidth = 1.000000
	__FrontCurveStack__000001_Front_Curve_Stack.StrokeDashArray = ``
	__FrontCurveStack__000001_Front_Curve_Stack.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__000001_Front_Curve_Stack.Transform = ``

	__FrontCurveStack__000002_Hour_Handle.Name = `Hour Handle`
	__FrontCurveStack__000002_Hour_Handle.IsDisplayed = false
	__FrontCurveStack__000002_Hour_Handle.Color = ``
	__FrontCurveStack__000002_Hour_Handle.FillOpacity = 0.000000
	__FrontCurveStack__000002_Hour_Handle.Stroke = `green`
	__FrontCurveStack__000002_Hour_Handle.StrokeOpacity = 1.000000
	__FrontCurveStack__000002_Hour_Handle.StrokeWidth = 1.000000
	__FrontCurveStack__000002_Hour_Handle.StrokeDashArray = ``
	__FrontCurveStack__000002_Hour_Handle.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__000002_Hour_Handle.Transform = ``

	__FrontCurveStack__000003_Minute_Handle.Name = `Minute Handle`
	__FrontCurveStack__000003_Minute_Handle.IsDisplayed = false
	__FrontCurveStack__000003_Minute_Handle.Color = ``
	__FrontCurveStack__000003_Minute_Handle.FillOpacity = 0.000000
	__FrontCurveStack__000003_Minute_Handle.Stroke = `green`
	__FrontCurveStack__000003_Minute_Handle.StrokeOpacity = 1.000000
	__FrontCurveStack__000003_Minute_Handle.StrokeWidth = 1.000000
	__FrontCurveStack__000003_Minute_Handle.StrokeDashArray = ``
	__FrontCurveStack__000003_Minute_Handle.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__000003_Minute_Handle.Transform = ``

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

	__NoteInfo__000000_1.Name = `1`
	__NoteInfo__000000_1.IsKept = true

	__NoteInfo__000001_10.Name = `10`
	__NoteInfo__000001_10.IsKept = true

	__NoteInfo__000002_11.Name = `11`
	__NoteInfo__000002_11.IsKept = false

	__NoteInfo__000003_12.Name = `12`
	__NoteInfo__000003_12.IsKept = false

	__NoteInfo__000004_13.Name = `13`
	__NoteInfo__000004_13.IsKept = true

	__NoteInfo__000005_14.Name = `14`
	__NoteInfo__000005_14.IsKept = true

	__NoteInfo__000006_15.Name = `15`
	__NoteInfo__000006_15.IsKept = false

	__NoteInfo__000007_2.Name = `2`
	__NoteInfo__000007_2.IsKept = false

	__NoteInfo__000008_3.Name = `3`
	__NoteInfo__000008_3.IsKept = true

	__NoteInfo__000009_4.Name = `4`
	__NoteInfo__000009_4.IsKept = false

	__NoteInfo__000010_5.Name = `5`
	__NoteInfo__000010_5.IsKept = true

	__NoteInfo__000011_6.Name = `6`
	__NoteInfo__000011_6.IsKept = true

	__NoteInfo__000012_7.Name = `7`
	__NoteInfo__000012_7.IsKept = true

	__NoteInfo__000013_8.Name = `8`
	__NoteInfo__000013_8.IsKept = true

	__NoteInfo__000014_9.Name = `9`
	__NoteInfo__000014_9.IsKept = true

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.BackendColor = `#F1F1F1`
	__Parameter__000000_Reference.MinuteColor = `#5A5A5A`
	__Parameter__000000_Reference.HourColor = `#1E1E1E`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.InsideAngle = 107.000000
	__Parameter__000000_Reference.SideLength = 168.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.440000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.HourHandleRotationAngle = 0.000000
	__Parameter__000000_Reference.HourHandleDiskDistance = 1.595000
	__Parameter__000000_Reference.HourHandleRadius = 0.290000
	__Parameter__000000_Reference.MinuteHandleRotationAngle = -120.000000
	__Parameter__000000_Reference.MinuteHandleDiskDistance = 2.630000
	__Parameter__000000_Reference.MinuteHandleRadius = 0.160000
	__Parameter__000000_Reference.MinuteOffset = 44.000000
	__Parameter__000000_Reference.BackendHandleRotationAngle = -120.500000
	__Parameter__000000_Reference.BackendHandleDiskDistance = 2.640000
	__Parameter__000000_Reference.BackendHandleRadius = 0.000000
	__Parameter__000000_Reference.BackendOffset = 194.000000
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.040000
	__Parameter__000000_Reference.NbPitchLines = 169
	__Parameter__000000_Reference.MeasureLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbMeasureLines = 300
	__Parameter__000000_Reference.NbMeasureLinesPerCurve = 15
	__Parameter__000000_Reference.FirstVoiceShiftX = -0.070000
	__Parameter__000000_Reference.FirstVoiceShiftY = 0.100000
	__Parameter__000000_Reference.PitchDifference = 7
	__Parameter__000000_Reference.Speed = 2.900000
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

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 168.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.AngleDegree = -81.581774
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 107.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 168.000000
	__Rhombus__000001_Initial_Rhombus.AngleDegree = -79.106605
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 107.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 267.185841
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 39.541424
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 168.000000
	__Rhombus__000002_Rotated_Next_Rhombus.AngleDegree = -81.581774
	__Rhombus__000002_Rotated_Next_Rhombus.InsideAngle = 107.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 168.000000
	__Rhombus__000003_Rotated_Rhombus.AngleDegree = -81.581774
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 107.000000
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
	__ShapeCategory__000006_6_Spiral_growth.IsExpanded = true

	__ShapeCategory__000007_7_Spiral_Growth_Bezier.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = false

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = true

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = true

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 274.426377
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -37.180112
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 324.674414
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 33.070910
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 226.737717
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 324.123446
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 258.165889
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 237.332459
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

	__SpiralCircle__000000_Backend_Marker.Name = `Backend Marker`
	__SpiralCircle__000000_Backend_Marker.IsDisplayed = false
	__SpiralCircle__000000_Backend_Marker.CenterX = 0.000000
	__SpiralCircle__000000_Backend_Marker.CenterY = 443.520000
	__SpiralCircle__000000_Backend_Marker.HasBespokeRadius = true
	__SpiralCircle__000000_Backend_Marker.BespopkeRadius = 0.000000
	__SpiralCircle__000000_Backend_Marker.Pitch = 0
	__SpiralCircle__000000_Backend_Marker.Color = ``
	__SpiralCircle__000000_Backend_Marker.FillOpacity = 0.000000
	__SpiralCircle__000000_Backend_Marker.Stroke = `green`
	__SpiralCircle__000000_Backend_Marker.StrokeOpacity = 1.000000
	__SpiralCircle__000000_Backend_Marker.StrokeWidth = 1.000000
	__SpiralCircle__000000_Backend_Marker.StrokeDashArray = ``
	__SpiralCircle__000000_Backend_Marker.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__000000_Backend_Marker.Transform = ``
	__SpiralCircle__000000_Backend_Marker.ShowName = false
	__SpiralCircle__000000_Backend_Marker.Path = `M 500.000000 204.320000 A 0.000000 0.000000 0 1 0 500.000000 204.320000 A 0.000000 0.000000 0 1 0 500.000000 204.320000 Z`

	__SpiralCircle__000001_Construction_Circle_Spiral.Name = `Construction Circle Spiral`
	__SpiralCircle__000001_Construction_Circle_Spiral.IsDisplayed = false
	__SpiralCircle__000001_Construction_Circle_Spiral.CenterX = 178.080000
	__SpiralCircle__000001_Construction_Circle_Spiral.CenterY = 0.000000
	__SpiralCircle__000001_Construction_Circle_Spiral.HasBespokeRadius = true
	__SpiralCircle__000001_Construction_Circle_Spiral.BespopkeRadius = 14.130000
	__SpiralCircle__000001_Construction_Circle_Spiral.Pitch = 0
	__SpiralCircle__000001_Construction_Circle_Spiral.Color = ``
	__SpiralCircle__000001_Construction_Circle_Spiral.FillOpacity = 0.000000
	__SpiralCircle__000001_Construction_Circle_Spiral.Stroke = `blue`
	__SpiralCircle__000001_Construction_Circle_Spiral.StrokeOpacity = 0.800000
	__SpiralCircle__000001_Construction_Circle_Spiral.StrokeWidth = 1.000000
	__SpiralCircle__000001_Construction_Circle_Spiral.StrokeDashArray = ``
	__SpiralCircle__000001_Construction_Circle_Spiral.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__000001_Construction_Circle_Spiral.Transform = ``
	__SpiralCircle__000001_Construction_Circle_Spiral.ShowName = false
	__SpiralCircle__000001_Construction_Circle_Spiral.Path = `M 554.450000 500.000000 A 14.130000 14.130000 0 1 0 526.190000 500.000000 A 14.130000 14.130000 0 1 0 554.450000 500.000000 Z`

	__SpiralCircle__000002_Hour_Marker.Name = `Hour Marker`
	__SpiralCircle__000002_Hour_Marker.IsDisplayed = false
	__SpiralCircle__000002_Hour_Marker.CenterX = 0.000000
	__SpiralCircle__000002_Hour_Marker.CenterY = 267.960000
	__SpiralCircle__000002_Hour_Marker.HasBespokeRadius = true
	__SpiralCircle__000002_Hour_Marker.BespopkeRadius = 48.720000
	__SpiralCircle__000002_Hour_Marker.Pitch = 0
	__SpiralCircle__000002_Hour_Marker.Color = ``
	__SpiralCircle__000002_Hour_Marker.FillOpacity = 0.000000
	__SpiralCircle__000002_Hour_Marker.Stroke = `green`
	__SpiralCircle__000002_Hour_Marker.StrokeOpacity = 1.000000
	__SpiralCircle__000002_Hour_Marker.StrokeWidth = 1.000000
	__SpiralCircle__000002_Hour_Marker.StrokeDashArray = ``
	__SpiralCircle__000002_Hour_Marker.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__000002_Hour_Marker.Transform = ``
	__SpiralCircle__000002_Hour_Marker.ShowName = false
	__SpiralCircle__000002_Hour_Marker.Path = `M 532.480000 321.360000 A 32.480000 32.480000 0 1 0 467.520000 321.360000 A 32.480000 32.480000 0 1 0 532.480000 321.360000 Z`

	__SpiralCircle__000003_Minute_Marker.Name = `Minute Marker`
	__SpiralCircle__000003_Minute_Marker.IsDisplayed = false
	__SpiralCircle__000003_Minute_Marker.CenterX = 0.000000
	__SpiralCircle__000003_Minute_Marker.CenterY = 441.840000
	__SpiralCircle__000003_Minute_Marker.HasBespokeRadius = true
	__SpiralCircle__000003_Minute_Marker.BespopkeRadius = 26.880000
	__SpiralCircle__000003_Minute_Marker.Pitch = 0
	__SpiralCircle__000003_Minute_Marker.Color = ``
	__SpiralCircle__000003_Minute_Marker.FillOpacity = 0.000000
	__SpiralCircle__000003_Minute_Marker.Stroke = `green`
	__SpiralCircle__000003_Minute_Marker.StrokeOpacity = 1.000000
	__SpiralCircle__000003_Minute_Marker.StrokeWidth = 1.000000
	__SpiralCircle__000003_Minute_Marker.StrokeDashArray = ``
	__SpiralCircle__000003_Minute_Marker.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__000003_Minute_Marker.Transform = ``
	__SpiralCircle__000003_Minute_Marker.ShowName = false
	__SpiralCircle__000003_Minute_Marker.Path = `M 517.920000 205.440000 A 17.920000 17.920000 0 1 0 482.080000 205.440000 A 17.920000 17.920000 0 1 0 517.920000 205.440000 Z`

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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 274.426377
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 178.080000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -37.180112
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 274.426377
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 362.240155
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -37.180112
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -99.990158
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 78.509172
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 10.636659
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 66.184282
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 186.452528
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 274.426377
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -37.180112
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 52.957032
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -149.189086
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
	__Axis__000000_Construction_Axis.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Axis__000001_Initial_Axis.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Axis__000004_Rotated_Axis.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__AxisGrid__000000_Construction_Axis_Grid.Reference = __Axis__000000_Construction_Axis
	__AxisGrid__000000_Construction_Axis_Grid.ShapeCategory = __ShapeCategory__000004_4_Construction
	__AxisGrid__000001_Measure_Lines.Reference = __Axis__000002_Measure
	__AxisGrid__000001_Measure_Lines.ShapeCategory = __ShapeCategory__000008_8_Score_notation
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
	__FrontCurveStack__000000_Backend_curve.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000000_Backend_curve.FrontCurves = append(__FrontCurveStack__000000_Backend_curve.FrontCurves, __FrontCurve__000010_Rotated_1_)
	__FrontCurveStack__000001_Front_Curve_Stack.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000000_Non_Rotated_0_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000001_Non_Rotated_1_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000002_Non_Rotated_2_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000003_Non_Rotated_3_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000004_Non_Rotated_4_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000005_Non_Rotated_5_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000006_Non_Rotated_6_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000007_Non_Rotated_7_)
	__FrontCurveStack__000002_Hour_Handle.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000002_Hour_Handle.FrontCurves = append(__FrontCurveStack__000002_Hour_Handle.FrontCurves, __FrontCurve__000008_Rotated_0_)
	__FrontCurveStack__000003_Minute_Handle.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000003_Minute_Handle.FrontCurves = append(__FrontCurveStack__000003_Minute_Handle.FrontCurves, __FrontCurve__000009_Rotated_1_)
	__HorizontalAxis__000000_Horizontal_Axis.ShapeCategory = __ShapeCategory__000000_0_Axes
	__Key__000000_F_key.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000001_Initial_Rhombus
	__Parameter__000000_Reference.InitialCircle = __Circle__000004_Initial_Circle
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000001_Initial_Rhombus_Grid
	__Parameter__000000_Reference.InitialCircleGrid = __CircleGrid__000005_Initial_Circle_Grid
	__Parameter__000000_Reference.InitialAxis = __Axis__000001_Initial_Axis
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
	__Parameter__000000_Reference.ConstructionAxis = __Axis__000000_Construction_Axis
	__Parameter__000000_Reference.ConstructionAxisGrid = __AxisGrid__000000_Construction_Axis_Grid
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
	__Parameter__000000_Reference.SpiralCircleSeed = __SpiralCircle__000001_Construction_Circle_Spiral
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
	__Parameter__000000_Reference.FrontCurveStack = __FrontCurveStack__000001_Front_Curve_Stack
	__Parameter__000000_Reference.HourCurve = __FrontCurveStack__000002_Hour_Handle
	__Parameter__000000_Reference.HourMarker = __SpiralCircle__000002_Hour_Marker
	__Parameter__000000_Reference.MinuteCurve = __FrontCurveStack__000003_Minute_Handle
	__Parameter__000000_Reference.MinuteMarker = __SpiralCircle__000003_Minute_Marker
	__Parameter__000000_Reference.BackendCurve = __FrontCurveStack__000000_Backend_curve
	__Parameter__000000_Reference.BackendMarker = __SpiralCircle__000000_Backend_Marker
	__Parameter__000000_Reference.Fkey = __Key__000000_F_key
	__Parameter__000000_Reference.PitchLines = __AxisGrid__000002_Pitch_Lines
	__Parameter__000000_Reference.MeasureLines = __AxisGrid__000001_Measure_Lines
	__Parameter__000000_Reference.FirstVoice = __BezierGrid__000002_First_Voice
	__Parameter__000000_Reference.FirstVoiceShiftRigth = __BezierGrid__000003_First_Voice_Shift_Right
	__Parameter__000000_Reference.SecondVoice = __BezierGrid__000000_2nb_Voice
	__Parameter__000000_Reference.SecondVoiceShiftedRight = __BezierGrid__000001_2nd_voice_shifted_right
	__Parameter__000000_Reference.FirstVoiceNotes = __CircleGrid__000002_First_Voice_notes
	__Parameter__000000_Reference.FirstVoiceNotesShiftedRight = __CircleGrid__000001_First_Voice_note_shifted_right
	__Parameter__000000_Reference.SecondVoiceNotes = __CircleGrid__000008_Second_Voice_notes
	__Parameter__000000_Reference.SecondVoiceNotesShiftedRight = __CircleGrid__000007_Second_Voice_Notes_Shift_Right
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000000_1)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000007_2)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000008_3)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000009_4)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000010_5)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000011_6)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000012_7)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000013_8)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000014_9)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000001_10)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000002_11)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000003_12)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000004_13)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000005_14)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000006_15)
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Horizontal_Axis
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Vertical_Axis
	__Parameter__000000_Reference.SpiralOrigin = __SpiralOrigin__000000_Spiral_Origin
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
	__SpiralCircle__000000_Backend_Marker.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircle__000001_Construction_Circle_Spiral.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
	__SpiralCircle__000002_Hour_Marker.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__SpiralCircle__000003_Minute_Marker.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
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
