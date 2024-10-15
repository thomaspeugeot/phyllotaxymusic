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
	__FrontCurve__000003_Rotated_0_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000004_Rotated_1_ := (&models.FrontCurve{}).Stage(stage)
	__FrontCurve__000005_Rotated_2_ := (&models.FrontCurve{}).Stage(stage)

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
	__Axis__000000_Construction_Axis.AngleDegree = 100.893395
	__Axis__000000_Construction_Axis.Length = 157.000000
	__Axis__000000_Construction_Axis.CenterX = 0.000000
	__Axis__000000_Construction_Axis.CenterY = 0.000000
	__Axis__000000_Construction_Axis.EndX = -29.670211
	__Axis__000000_Construction_Axis.EndY = 154.170939
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
	__Axis__000001_Initial_Axis.AngleDegree = 79.106605
	__Axis__000001_Initial_Axis.Length = 415.382956
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
	__Axis__000003_Pitch_Line.Length = 4000.000000
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
	__Axis__000004_Rotated_Axis.Length = 415.382956
	__Axis__000004_Rotated_Axis.CenterX = 0.000000
	__Axis__000004_Rotated_Axis.CenterY = 0.000000
	__Axis__000004_Rotated_Axis.EndX = 0.000000
	__Axis__000004_Rotated_Axis.EndY = 0.000000
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
	__AxisGrid__000001_Measure_Lines.IsDisplayed = false

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
	__Bezier__000005_Growth_Curve_Seed.StartX = -14.835106
	__Bezier__000005_Growth_Curve_Seed.StartY = 77.085470
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 34.499595
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 86.579937
	__Bezier__000005_Growth_Curve_Seed.EndX = 103.845739
	__Bezier__000005_Growth_Curve_Seed.EndY = 179.866096
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 54.511038
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 170.371628
	__Bezier__000005_Growth_Curve_Seed.Color = ``
	__Bezier__000005_Growth_Curve_Seed.FillOpacity = 0.000000
	__Bezier__000005_Growth_Curve_Seed.Stroke = `grey`
	__Bezier__000005_Growth_Curve_Seed.StrokeOpacity = 0.800000
	__Bezier__000005_Growth_Curve_Seed.StrokeWidth = 6.000000
	__Bezier__000005_Growth_Curve_Seed.StrokeDashArray = ``
	__Bezier__000005_Growth_Curve_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000005_Growth_Curve_Seed.Transform = ``

	__BezierGrid__000000_2nb_Voice.Name = `2nb Voice`
	__BezierGrid__000000_2nb_Voice.IsDisplayed = false

	__BezierGrid__000001_2nd_voice_shifted_right.Name = `2nd voice shifted right`
	__BezierGrid__000001_2nd_voice_shifted_right.IsDisplayed = false

	__BezierGrid__000002_First_Voice.Name = `First Voice`
	__BezierGrid__000002_First_Voice.IsDisplayed = false

	__BezierGrid__000003_First_Voice_Shift_Right.Name = `First Voice Shift Right`
	__BezierGrid__000003_First_Voice_Shift_Right.IsDisplayed = false

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
	__Circle__000001_Construction_Circle.CenterX = -14.835106
	__Circle__000001_Construction_Circle.CenterY = 77.085470
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 267.031900
	__Circle__000005_Rotated_Next_Circle.CenterY = 51.390313
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

	__FrontCurve__000000_Non_Rotated_0_.Name = `Non Rotated 0 `
	__FrontCurve__000000_Non_Rotated_0_.Path = `M630.26 529.73 C632.67 522.58 634.35 516.93 636.23 510.21 C638.12 503.49 639.94 496.65 641.57 489.39 C643.19 482.14 644.78 474.64 645.99 466.68 C647.20 458.71 648.33 450.37 648.81 441.60 C649.29 432.83 649.55 423.55 648.86 414.06 C648.16 404.57 646.99 394.54 644.64 384.66 C642.28 374.78 639.12 364.49 634.75 354.78 C630.37 345.06 624.90 335.25 618.37 326.38 C611.84 317.51 604.09 308.95 595.56 301.57 C587.03 294.19 577.34 287.49 567.21 282.10 C557.08 276.71 545.99 272.31 534.78 269.23 C523.58 266.15 513.15 263.99 500.00 263.61 C486.85 263.23 470.23 264.04 455.90 266.95 C441.58 269.85 427.20 274.81 414.07 281.05 C400.93 287.29 388.30 295.41 377.09 304.40 C365.89 313.38 355.63 323.94 346.85 334.95 C338.08 345.95 330.55 358.16 324.44 370.43 C318.32 382.70 313.62 395.75 310.17 408.58 C306.73 421.42 304.72 434.62 303.74 447.41 C302.77 460.20 303.14 473.02 304.33 485.34 C305.51 497.65 307.87 509.72 310.83 521.31 C313.79 532.90 317.71 544.08 322.10 554.88 C326.48 565.67 331.61 575.99 337.15 586.07 C342.69 596.15 348.42 606.22 355.36 615.34 C362.31 624.47 370.02 633.52 378.81 640.83 C387.59 648.14 397.78 654.41 408.09 659.19 C418.40 663.98 429.67 667.44 440.67 669.54 C451.68 671.65 463.17 672.30 474.10 671.83 C485.03 671.35 496.01 669.45 506.24 666.69 C516.46 663.94 526.37 659.90 535.45 655.31 C544.52 650.71 553.01 645.07 560.70 639.13 C568.39 633.19 575.32 626.48 581.59 619.67 C587.85 612.86 593.31 605.56 598.27 598.27 C603.24 590.98 607.45 583.45 611.36 575.93 C615.27 568.40 618.57 560.81 621.72 553.11 C624.87 545.41 627.84 536.88 630.26 529.73 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M355.36 615.34 C361.19 623.97 366.17 630.42 372.12 637.82 C378.08 645.22 384.31 652.52 391.08 659.76 C397.85 666.99 404.96 674.20 412.73 681.22 C420.51 688.24 428.73 695.27 437.73 701.87 C446.73 708.46 456.34 715.03 466.72 720.78 C477.10 726.53 488.26 732.06 500.00 736.39 C511.74 740.71 524.34 744.46 537.19 746.71 C550.03 748.96 563.64 750.23 577.08 749.90 C590.53 749.57 604.47 747.91 617.86 744.74 C631.24 741.56 644.76 736.87 657.40 730.87 C670.04 724.87 682.43 717.32 693.69 708.75 C704.96 700.17 714.75 691.74 724.99 679.43 C735.23 667.11 746.99 650.67 755.14 634.84 C763.28 619.02 769.62 601.68 773.87 584.48 C778.11 567.27 780.31 549.14 780.63 531.62 C780.95 514.10 779.15 496.23 775.78 479.33 C772.41 462.43 767.01 445.73 760.40 430.22 C753.80 414.72 745.39 399.85 736.13 386.29 C726.86 372.72 716.11 360.13 704.83 348.83 C693.54 337.53 681.12 327.41 668.42 318.49 C655.72 309.56 642.23 301.90 628.62 295.30 C615.02 288.69 600.93 283.32 586.79 278.86 C572.65 274.40 558.26 271.06 543.80 268.52 C529.33 265.98 514.65 263.88 500.00 263.61 C485.35 263.35 470.23 264.04 455.90 266.95 C441.58 269.85 427.20 274.81 414.07 281.05 C400.93 287.29 388.30 295.41 377.09 304.40 C365.89 313.38 355.63 323.94 346.85 334.95 C338.08 345.95 330.55 358.16 324.44 370.43 C318.32 382.70 313.62 395.75 310.17 408.58 C306.73 421.42 304.72 434.62 303.74 447.41 C302.77 460.20 303.14 473.02 304.33 485.34 C305.51 497.65 307.87 509.72 310.83 521.31 C313.79 532.90 317.71 544.08 322.10 554.88 C326.48 565.67 331.61 575.99 337.15 586.07 C342.69 596.15 349.54 606.72 355.36 615.34 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M500.00 263.61 C486.75 262.74 476.34 262.86 464.32 263.28 C452.30 263.70 440.18 264.60 427.86 266.13 C415.55 267.66 403.07 269.67 390.43 272.48 C377.80 275.29 364.92 278.63 352.05 282.99 C339.17 287.36 326.02 292.39 313.18 298.66 C300.34 304.92 287.30 312.09 275.01 320.57 C262.71 329.06 250.47 338.70 239.42 349.56 C228.38 360.41 217.78 372.59 208.72 385.68 C199.66 398.78 191.51 413.19 185.07 428.12 C178.64 443.05 173.50 459.11 170.12 475.28 C166.73 491.45 164.91 508.42 164.78 525.12 C164.65 541.82 165.40 557.00 169.34 575.47 C173.27 593.94 179.83 616.80 188.40 635.95 C196.97 655.10 208.14 673.72 220.74 690.40 C233.34 707.08 248.22 722.60 263.97 736.03 C279.73 749.45 297.32 761.26 315.26 770.96 C333.21 780.65 352.48 788.42 371.64 794.20 C390.81 799.98 410.76 803.69 430.25 805.61 C449.73 807.54 469.51 807.41 488.56 805.75 C507.61 804.09 526.52 800.51 544.56 795.66 C562.61 790.82 580.17 784.27 596.82 776.70 C613.47 769.13 629.42 760.10 644.48 750.24 C659.54 740.38 673.77 729.32 687.19 717.52 C700.61 705.72 713.67 693.21 724.99 679.43 C736.32 665.65 746.99 650.67 755.14 634.84 C763.28 619.02 769.62 601.68 773.87 584.48 C778.11 567.27 780.31 549.14 780.63 531.62 C780.95 514.10 779.15 496.23 775.78 479.33 C772.41 462.43 767.01 445.73 760.40 430.22 C753.80 414.72 745.39 399.85 736.13 386.29 C726.86 372.72 716.11 360.13 704.83 348.83 C693.54 337.53 681.12 327.41 668.42 318.49 C655.72 309.56 642.23 301.90 628.62 295.30 C615.02 288.69 600.93 283.32 586.79 278.86 C572.65 274.40 558.26 271.06 543.80 268.52 C529.33 265.98 513.25 264.49 500.00 263.61 Z`

	__FrontCurve__000003_Rotated_0_.Name = `Rotated 0 `
	__FrontCurve__000003_Rotated_0_.Path = `M630.26 529.73 C632.67 522.58 634.35 516.93 636.23 510.21 C638.12 503.49 639.94 496.65 641.57 489.39 C643.19 482.14 644.78 474.64 645.99 466.68 C647.20 458.71 648.33 450.37 648.81 441.60 C649.29 432.83 649.55 423.55 648.86 414.06 C648.16 404.57 646.99 394.54 644.64 384.66 C642.28 374.78 639.12 364.49 634.75 354.78 C630.37 345.06 624.90 335.25 618.37 326.38 C611.84 317.51 604.09 308.95 595.56 301.57 C587.03 294.19 577.34 287.49 567.21 282.10 C557.08 276.71 545.99 272.31 534.78 269.23 C523.58 266.15 513.15 263.99 500.00 263.61 C486.85 263.23 470.23 264.04 455.90 266.95 C441.58 269.85 427.20 274.81 414.07 281.05 C400.93 287.29 388.30 295.41 377.09 304.40 C365.89 313.38 355.63 323.94 346.85 334.95 C338.08 345.95 330.55 358.16 324.44 370.43 C318.32 382.70 313.62 395.75 310.17 408.58 C306.73 421.42 304.72 434.62 303.74 447.41 C302.77 460.20 303.14 473.02 304.33 485.34 C305.51 497.65 307.87 509.72 310.83 521.31 C313.79 532.90 317.71 544.08 322.10 554.88 C326.48 565.67 331.61 575.99 337.15 586.07 C342.69 596.15 348.42 606.22 355.36 615.34 C362.31 624.47 370.02 633.52 378.81 640.83 C387.59 648.14 397.78 654.41 408.09 659.19 C418.40 663.98 429.67 667.44 440.67 669.54 C451.68 671.65 463.17 672.30 474.10 671.83 C485.03 671.35 496.01 669.45 506.24 666.69 C516.46 663.94 526.37 659.90 535.45 655.31 C544.52 650.71 553.01 645.07 560.70 639.13 C568.39 633.19 575.32 626.48 581.59 619.67 C587.85 612.86 593.31 605.56 598.27 598.27 C603.24 590.98 607.45 583.45 611.36 575.93 C615.27 568.40 618.57 560.81 621.72 553.11 C624.87 545.41 627.84 536.88 630.26 529.73 Z`

	__FrontCurve__000004_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000004_Rotated_1_.Path = `M680.82 570.97 C685.60 561.14 688.85 553.23 692.46 543.93 C696.06 534.63 699.43 525.12 702.45 515.17 C705.48 505.22 708.30 494.97 710.60 484.22 C712.90 473.46 714.98 462.29 716.25 450.64 C717.53 439.00 718.45 426.81 718.23 414.35 C718.02 401.89 717.19 388.85 714.95 375.90 C712.72 362.95 709.52 349.53 704.82 336.66 C700.12 323.80 694.13 310.76 686.77 298.71 C679.41 286.65 670.58 274.85 660.67 264.34 C650.76 253.83 639.39 244.01 627.30 235.66 C615.21 227.32 601.84 220.02 588.13 214.28 C574.42 208.54 561.61 204.06 545.04 201.21 C528.46 198.36 507.34 196.30 488.67 197.21 C470.00 198.11 450.91 201.45 433.04 206.62 C415.16 211.78 397.52 219.31 381.42 228.22 C365.32 237.13 350.02 248.14 336.42 260.08 C322.83 272.02 310.47 285.69 299.84 299.84 C289.21 313.99 280.10 329.44 272.63 344.98 C265.16 360.53 259.35 376.91 255.00 393.11 C250.65 409.30 247.97 425.91 246.52 442.15 C245.08 458.38 245.19 474.67 246.33 490.51 C247.47 506.35 249.98 521.98 253.34 537.18 C256.71 552.38 261.23 567.21 266.52 581.70 C271.80 596.19 277.60 610.64 285.05 624.10 C292.50 637.56 301.07 650.95 311.23 662.45 C321.39 673.95 333.45 684.42 346.02 693.09 C358.59 701.76 372.61 708.98 386.66 714.45 C400.71 719.93 415.70 723.71 430.31 725.91 C444.93 728.12 459.98 728.55 474.35 727.67 C488.72 726.79 503.05 724.21 516.53 720.61 C530.01 717.01 543.07 711.90 555.22 706.07 C567.36 700.24 578.81 693.17 589.39 685.63 C599.97 678.09 609.71 669.61 618.70 660.83 C627.68 652.05 635.79 642.62 643.30 632.96 C650.81 623.31 657.51 613.23 663.76 602.90 C670.01 592.56 676.04 580.79 680.82 570.97 Z`

	__FrontCurve__000005_Rotated_2_.Name = `Rotated 2 `
	__FrontCurve__000005_Rotated_2_.Path = `M786.09 630.38 C794.45 614.83 800.05 602.16 806.17 587.38 C812.30 572.60 817.89 557.43 822.83 541.68 C827.76 525.93 832.22 509.72 835.78 492.88 C839.35 476.04 842.41 458.62 844.23 440.63 C846.05 422.64 847.21 403.95 846.71 384.95 C846.22 365.96 844.73 346.22 841.24 326.66 C837.76 307.11 832.83 286.97 825.79 267.61 C818.75 248.26 809.86 228.72 799.01 210.53 C788.16 192.35 775.21 174.53 760.69 158.50 C746.16 142.48 729.57 127.40 711.86 114.39 C694.16 101.38 674.62 89.82 654.48 80.45 C634.35 71.08 615.56 63.61 591.04 58.19 C566.51 52.77 535.23 48.10 507.33 47.90 C479.42 47.71 450.73 50.97 423.60 57.01 C396.46 63.06 369.47 72.51 344.53 84.17 C319.59 95.83 295.60 110.61 273.96 126.98 C252.33 143.35 232.30 162.38 214.74 182.39 C197.18 202.40 181.68 224.50 168.61 247.02 C155.53 269.54 144.78 293.55 136.28 317.52 C127.78 341.50 121.69 366.36 117.59 390.86 C113.50 415.37 111.75 440.24 111.72 464.57 C111.70 488.90 113.83 513.17 117.46 536.83 C121.09 560.50 126.62 583.83 133.50 606.58 C140.39 629.33 148.33 652.03 158.76 673.34 C169.19 694.64 181.43 715.82 196.09 734.41 C210.75 752.99 228.24 770.22 246.71 784.85 C265.19 799.48 285.92 812.14 306.95 822.19 C327.99 832.24 350.60 839.91 372.91 845.16 C395.22 850.40 418.41 853.08 440.82 853.64 C463.23 854.19 485.86 852.22 507.39 848.49 C528.91 844.76 550.08 838.69 569.98 831.27 C589.88 823.84 608.98 814.39 626.79 803.94 C644.59 793.49 661.30 781.39 676.80 768.57 C692.30 755.75 706.57 741.66 719.77 727.01 C732.97 712.36 744.94 696.78 755.99 680.68 C767.04 664.57 777.72 645.93 786.09 630.38 Z`

	__FrontCurveStack__000000_Backend_curve.Name = `Backend curve`
	__FrontCurveStack__000000_Backend_curve.IsDisplayed = true
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
	__FrontCurveStack__000002_Hour_Handle.IsDisplayed = true
	__FrontCurveStack__000002_Hour_Handle.Color = ``
	__FrontCurveStack__000002_Hour_Handle.FillOpacity = 0.000000
	__FrontCurveStack__000002_Hour_Handle.Stroke = `green`
	__FrontCurveStack__000002_Hour_Handle.StrokeOpacity = 1.000000
	__FrontCurveStack__000002_Hour_Handle.StrokeWidth = 1.000000
	__FrontCurveStack__000002_Hour_Handle.StrokeDashArray = ``
	__FrontCurveStack__000002_Hour_Handle.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__000002_Hour_Handle.Transform = ``

	__FrontCurveStack__000003_Minute_Handle.Name = `Minute Handle`
	__FrontCurveStack__000003_Minute_Handle.IsDisplayed = true
	__FrontCurveStack__000003_Minute_Handle.Color = ``
	__FrontCurveStack__000003_Minute_Handle.FillOpacity = 0.000000
	__FrontCurveStack__000003_Minute_Handle.Stroke = `green`
	__FrontCurveStack__000003_Minute_Handle.StrokeOpacity = 1.000000
	__FrontCurveStack__000003_Minute_Handle.StrokeWidth = 1.000000
	__FrontCurveStack__000003_Minute_Handle.StrokeDashArray = ``
	__FrontCurveStack__000003_Minute_Handle.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__000003_Minute_Handle.Transform = ``

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
	__Parameter__000000_Reference.N = 2
	__Parameter__000000_Reference.M = 1
	__Parameter__000000_Reference.Z = 7
	__Parameter__000000_Reference.ShiftToNearestCircle = 2
	__Parameter__000000_Reference.InsideAngle = 120.000000
	__Parameter__000000_Reference.SideLength = 157.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 3
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.320000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.HourHandleRotationAngle = 0.000000
	__Parameter__000000_Reference.HourHandleDiskDistance = 1.180000
	__Parameter__000000_Reference.HourHandleRadius = 0.260000
	__Parameter__000000_Reference.MinuteHandleRotationAngle = -120.000000
	__Parameter__000000_Reference.MinuteHandleDiskDistance = 1.671000
	__Parameter__000000_Reference.MinuteHandleRadius = 0.120000
	__Parameter__000000_Reference.MinuteScale = 1.050000
	__Parameter__000000_Reference.BackendHandleRotationAngle = 114.500000
	__Parameter__000000_Reference.BackendHandleDiskDistance = 2.760000
	__Parameter__000000_Reference.BackendHandleRadius = 0.030000
	__Parameter__000000_Reference.BackendScale = 1.330000
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.002000
	__Parameter__000000_Reference.NbPitchLines = 169
	__Parameter__000000_Reference.MeasureLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbMeasureLines = 300
	__Parameter__000000_Reference.NbMeasureLinesPerCurve = 15
	__Parameter__000000_Reference.FirstVoiceShiftX = -0.070000
	__Parameter__000000_Reference.FirstVoiceShiftY = 0.670000
	__Parameter__000000_Reference.PitchDifference = 7
	__Parameter__000000_Reference.Speed = 2.900000
	__Parameter__000000_Reference.Level = 5.600000
	__Parameter__000000_Reference.IsMinor = true
	__Parameter__000000_Reference.OriginX = 100.000000
	__Parameter__000000_Reference.OriginY = 350.000000
	__Parameter__000000_Reference.SpiralOriginX = 500.000000
	__Parameter__000000_Reference.SpiralOriginY = 500.000000
	__Parameter__000000_Reference.OriginCrossWidth = 800.000000
	__Parameter__000000_Reference.SpiralRadiusRatio = 0.360000
	__Parameter__000000_Reference.ShowSpiralBezierConstruct = false
	__Parameter__000000_Reference.ShowInterpolationPoints = false

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 157.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.AngleDegree = -79.106605
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 157.000000
	__Rhombus__000001_Initial_Rhombus.AngleDegree = 0.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 267.031900
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 51.390313
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 157.000000
	__Rhombus__000002_Rotated_Next_Rhombus.AngleDegree = -79.106605
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 157.000000
	__Rhombus__000003_Rotated_Rhombus.AngleDegree = -79.106605
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
	__ShapeCategory__000004_4_Construction.IsExpanded = false

	__ShapeCategory__000005_5_Vertical_growth.Name = `5. Vertical growth`
	__ShapeCategory__000005_5_Vertical_growth.IsExpanded = false

	__ShapeCategory__000006_6_Spiral_growth.Name = `6. Spiral growth`
	__ShapeCategory__000006_6_Spiral_growth.IsExpanded = false

	__ShapeCategory__000007_7_Spiral_Growth_Bezier.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = true

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = false

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = false

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 130.255702
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -29.730014
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 179.603989
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 17.926207
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = -0.000000
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 236.386096
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 55.380866
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 180.582908
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
	__SpiralCircle__000000_Backend_Marker.IsDisplayed = true
	__SpiralCircle__000000_Backend_Marker.CenterX = 0.000000
	__SpiralCircle__000000_Backend_Marker.CenterY = 433.320000
	__SpiralCircle__000000_Backend_Marker.HasBespokeRadius = true
	__SpiralCircle__000000_Backend_Marker.BespopkeRadius = 4.710000
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
	__SpiralCircle__000000_Backend_Marker.Path = `M 504.710000 66.680000 A 4.710000 4.710000 0 1 0 495.290000 66.680000 A 4.710000 4.710000 0 1 0 504.710000 66.680000 Z`

	__SpiralCircle__000001_Construction_Circle_Spiral.Name = `Construction Circle Spiral`
	__SpiralCircle__000001_Construction_Circle_Spiral.IsDisplayed = false
	__SpiralCircle__000001_Construction_Circle_Spiral.CenterX = 56.520000
	__SpiralCircle__000001_Construction_Circle_Spiral.CenterY = 0.000000
	__SpiralCircle__000001_Construction_Circle_Spiral.HasBespokeRadius = false
	__SpiralCircle__000001_Construction_Circle_Spiral.BespopkeRadius = 0.000000
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
	__SpiralCircle__000001_Construction_Circle_Spiral.Path = ``

	__SpiralCircle__000002_Hour_Marker.Name = `Hour Marker`
	__SpiralCircle__000002_Hour_Marker.IsDisplayed = true
	__SpiralCircle__000002_Hour_Marker.CenterX = 0.000000
	__SpiralCircle__000002_Hour_Marker.CenterY = 185.260000
	__SpiralCircle__000002_Hour_Marker.HasBespokeRadius = true
	__SpiralCircle__000002_Hour_Marker.BespopkeRadius = 40.820000
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
	__SpiralCircle__000002_Hour_Marker.Path = `M 540.820000 314.740000 A 40.820000 40.820000 0 1 0 459.180000 314.740000 A 40.820000 40.820000 0 1 0 540.820000 314.740000 Z`

	__SpiralCircle__000003_Minute_Marker.Name = `Minute Marker`
	__SpiralCircle__000003_Minute_Marker.IsDisplayed = true
	__SpiralCircle__000003_Minute_Marker.CenterX = 0.000000
	__SpiralCircle__000003_Minute_Marker.CenterY = 262.347000
	__SpiralCircle__000003_Minute_Marker.HasBespokeRadius = true
	__SpiralCircle__000003_Minute_Marker.BespopkeRadius = 18.840000
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
	__SpiralCircle__000003_Minute_Marker.Path = `M 518.840000 237.653000 A 18.840000 18.840000 0 1 0 481.160000 237.653000 A 18.840000 18.840000 0 1 0 518.840000 237.653000 Z`

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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 130.255702
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 56.520000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -29.730014
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 130.255702
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 189.825977
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -29.730014
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -91.415373
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = -20.049850
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = -4.576248
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = -35.671820
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 74.073297
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 130.255702
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -29.730014
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = -13.374398
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -27.772224
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
	__FrontCurveStack__000000_Backend_curve.FrontCurves = append(__FrontCurveStack__000000_Backend_curve.FrontCurves, __FrontCurve__000005_Rotated_2_)
	__FrontCurveStack__000001_Front_Curve_Stack.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000000_Non_Rotated_0_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000001_Non_Rotated_1_)
	__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000001_Front_Curve_Stack.FrontCurves, __FrontCurve__000002_Non_Rotated_2_)
	__FrontCurveStack__000002_Hour_Handle.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000002_Hour_Handle.FrontCurves = append(__FrontCurveStack__000002_Hour_Handle.FrontCurves, __FrontCurve__000003_Rotated_0_)
	__FrontCurveStack__000003_Minute_Handle.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000003_Minute_Handle.FrontCurves = append(__FrontCurveStack__000003_Minute_Handle.FrontCurves, __FrontCurve__000004_Rotated_1_)
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
