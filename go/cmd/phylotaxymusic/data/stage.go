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
	__FrontCurve__000005_Rotated_1_ := (&models.FrontCurve{}).Stage(stage)

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
	__Axis__000003_Pitch_Line.Length = 415.382956
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M631.02 530.37 C633.45 523.22 635.12 517.56 637.01 510.83 C638.90 504.11 640.73 497.26 642.36 490.00 C643.99 482.74 645.59 475.24 646.80 467.26 C648.02 459.29 649.16 450.94 649.64 442.16 C650.12 433.38 650.39 424.09 649.70 414.60 C649.01 405.10 647.84 395.06 645.49 385.17 C643.15 375.28 639.99 364.98 635.62 355.26 C631.25 345.54 625.78 335.72 619.25 326.85 C612.73 317.98 604.98 309.40 596.45 302.02 C587.93 294.63 578.24 287.94 568.11 282.55 C557.98 277.16 546.88 272.76 535.68 269.68 C524.47 266.60 514.04 264.45 500.88 264.08 C487.73 263.71 471.10 264.53 456.77 267.45 C442.44 270.38 428.04 275.35 414.90 281.61 C401.75 287.87 389.09 296.02 377.87 305.02 C366.65 314.03 356.37 324.61 347.57 335.64 C338.78 346.67 331.23 358.89 325.10 371.18 C318.96 383.47 314.24 396.54 310.78 409.38 C307.31 422.22 305.29 435.44 304.30 448.24 C303.32 461.04 303.68 473.86 304.86 486.18 C306.03 498.51 308.38 510.58 311.34 522.17 C314.30 533.77 318.22 544.95 322.60 555.74 C326.98 566.53 332.10 576.86 337.65 586.94 C343.19 597.01 348.92 607.08 355.86 616.21 C362.81 625.33 370.52 634.38 379.32 641.69 C388.11 648.99 398.30 655.26 408.62 660.04 C418.94 664.82 430.21 668.28 441.22 670.38 C452.23 672.48 463.74 673.13 474.68 672.64 C485.61 672.16 496.61 670.25 506.84 667.49 C517.08 664.73 526.99 660.68 536.08 656.08 C545.17 651.48 553.66 645.82 561.36 639.88 C569.06 633.93 576.00 627.21 582.27 620.40 C588.54 613.58 594.01 606.28 598.98 598.98 C603.95 591.68 608.18 584.15 612.09 576.61 C616.01 569.08 619.32 561.48 622.47 553.77 C625.63 546.07 628.60 537.53 631.02 530.37 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M355.86 616.21 C361.69 624.83 366.67 631.28 372.63 638.68 C378.58 646.08 384.82 653.38 391.59 660.62 C398.36 667.85 405.47 675.06 413.25 682.07 C421.03 689.09 429.26 696.13 438.26 702.72 C447.26 709.31 456.88 715.87 467.27 721.62 C477.65 727.37 488.81 732.90 500.56 737.21 C512.31 741.53 524.91 745.28 537.77 747.52 C550.62 749.77 564.24 751.03 577.69 750.69 C591.15 750.36 605.10 748.69 618.50 745.51 C631.89 742.32 645.42 737.62 658.07 731.61 C670.72 725.60 683.11 718.05 694.39 709.46 C705.66 700.88 715.47 692.44 725.72 680.11 C735.97 667.78 747.74 651.32 755.90 635.49 C764.06 619.65 770.41 602.30 774.66 585.08 C778.92 567.86 781.13 549.72 781.45 532.18 C781.78 514.65 779.99 496.77 776.63 479.86 C773.26 462.95 767.88 446.23 761.28 430.72 C754.67 415.20 746.27 400.33 737.01 386.75 C727.75 373.18 717.01 360.58 705.72 349.27 C694.44 337.97 682.02 327.84 669.32 318.92 C656.62 309.99 643.13 302.33 629.53 295.72 C615.92 289.12 601.83 283.75 587.69 279.29 C573.56 274.83 559.16 271.50 544.69 268.96 C530.23 266.43 515.54 264.33 500.88 264.08 C486.23 263.83 471.10 264.53 456.77 267.45 C442.44 270.38 428.04 275.35 414.90 281.61 C401.75 287.87 389.09 296.02 377.87 305.02 C366.65 314.03 356.37 324.61 347.57 335.64 C338.78 346.67 331.23 358.89 325.10 371.18 C318.96 383.47 314.24 396.54 310.78 409.38 C307.31 422.22 305.29 435.44 304.30 448.24 C303.32 461.04 303.68 473.86 304.86 486.18 C306.03 498.51 308.38 510.58 311.34 522.17 C314.30 533.77 318.22 544.95 322.60 555.74 C326.98 566.53 332.10 576.86 337.65 586.94 C343.19 597.01 350.03 607.58 355.86 616.21 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M500.88 264.08 C487.63 263.21 477.22 263.34 465.19 263.77 C453.16 264.20 441.03 265.11 428.71 266.66 C416.39 268.21 403.90 270.23 391.25 273.05 C378.61 275.88 365.72 279.24 352.83 283.62 C339.94 288.00 326.77 293.06 313.90 299.35 C301.04 305.63 287.98 312.83 275.66 321.33 C263.34 329.84 251.07 339.51 239.99 350.38 C228.91 361.25 218.28 373.45 209.20 386.56 C200.11 399.67 191.93 414.09 185.47 429.04 C179.01 443.98 173.85 460.05 170.45 476.22 C167.05 492.39 165.22 509.37 165.08 526.08 C164.94 542.78 165.69 557.96 169.62 576.43 C173.55 594.90 180.11 617.76 188.68 636.91 C197.25 656.06 208.44 674.67 221.04 691.35 C233.65 708.03 248.54 723.55 264.31 736.97 C280.08 750.39 297.68 762.19 315.64 771.88 C333.60 781.57 352.89 789.34 372.07 795.11 C391.24 800.87 411.21 804.58 430.72 806.49 C450.22 808.41 470.01 808.27 489.08 806.60 C508.15 804.94 527.07 801.34 545.13 796.49 C563.19 791.64 580.76 785.07 597.43 777.49 C614.10 769.91 630.05 760.87 645.13 751.00 C660.20 741.12 674.45 730.06 687.88 718.24 C701.31 706.43 714.39 693.90 725.72 680.11 C737.06 666.32 747.74 651.32 755.90 635.49 C764.06 619.65 770.41 602.30 774.66 585.08 C778.92 567.86 781.13 549.72 781.45 532.18 C781.78 514.65 779.99 496.77 776.63 479.86 C773.26 462.95 767.88 446.23 761.28 430.72 C754.67 415.20 746.27 400.33 737.01 386.75 C727.75 373.18 717.01 360.58 705.72 349.27 C694.44 337.97 682.02 327.84 669.32 318.92 C656.62 309.99 643.13 302.33 629.53 295.72 C615.92 289.12 601.83 283.75 587.69 279.29 C573.56 274.83 559.16 271.50 544.69 268.96 C530.23 266.43 514.14 264.95 500.88 264.08 Z`

	__FrontCurve__000003_Rotated_0_.Name = `Rotated 0 `
	__FrontCurve__000003_Rotated_0_.Path = `M631.23 529.95 C633.66 522.75 635.34 517.06 637.23 510.28 C639.12 503.51 640.94 496.62 642.56 489.32 C644.18 482.01 645.77 474.47 646.97 466.46 C648.16 458.44 649.28 450.05 649.74 441.23 C650.20 432.42 650.44 423.09 649.72 413.56 C649.00 404.03 647.80 393.95 645.42 384.03 C643.03 374.12 639.84 363.79 635.43 354.04 C631.01 344.30 625.50 334.46 618.93 325.56 C612.36 316.66 604.57 308.07 595.99 300.66 C587.42 293.26 577.68 286.55 567.51 281.15 C557.33 275.74 546.18 271.33 534.93 268.24 C523.68 265.15 513.20 263.00 500.00 262.61 C486.80 262.24 470.10 263.05 455.72 265.97 C441.34 268.88 426.90 273.86 413.70 280.12 C400.51 286.39 387.82 294.54 376.56 303.55 C365.31 312.57 355.00 323.17 346.17 334.21 C337.35 345.26 329.78 357.51 323.63 369.83 C317.48 382.16 312.75 395.27 309.27 408.15 C305.80 421.04 303.77 434.30 302.78 447.15 C301.79 460.01 302.15 472.88 303.33 485.26 C304.51 497.64 306.87 509.78 309.83 521.43 C312.81 533.08 316.74 544.32 321.14 555.17 C325.55 566.02 330.69 576.40 336.27 586.54 C341.84 596.67 347.60 606.79 354.58 615.97 C361.57 625.14 369.32 634.24 378.15 641.59 C386.99 648.93 397.23 655.24 407.59 660.06 C417.96 664.87 429.29 668.36 440.34 670.49 C451.41 672.61 462.96 673.28 473.95 672.82 C484.94 672.35 495.99 670.45 506.27 667.69 C516.56 664.94 526.53 660.89 535.67 656.28 C544.81 651.67 553.35 646.01 561.10 640.04 C568.85 634.08 575.84 627.33 582.15 620.49 C588.46 613.65 593.97 606.31 598.98 598.98 C603.99 591.64 608.24 584.07 612.19 576.49 C616.13 568.91 619.46 561.26 622.64 553.51 C625.81 545.75 628.80 537.16 631.23 529.95 Z`

	__FrontCurve__000004_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000004_Rotated_1_.Path = `M694.55 576.35 C699.56 565.73 702.96 557.20 706.69 547.18 C710.36 537.15 713.74 526.93 716.74 516.24 C719.68 505.58 722.38 494.63 724.51 483.18 C726.58 471.76 728.38 459.93 729.35 447.65 C730.28 435.42 730.80 422.67 730.18 409.66 C729.52 396.72 728.21 383.20 725.50 369.81 C722.76 356.48 719.03 342.69 713.83 329.48 C708.61 316.33 702.07 303.02 694.20 290.70 C686.32 278.45 676.95 266.44 666.54 255.73 C656.11 245.09 644.23 235.12 631.65 226.63 C619.06 218.20 605.20 210.80 591.01 204.95 C576.82 199.16 563.57 194.65 546.47 191.71 C529.36 188.83 507.57 186.77 488.31 187.64 C469.06 188.60 449.36 191.98 430.89 197.19 C412.43 202.49 394.17 210.13 377.47 219.16 C360.79 228.29 344.88 239.51 330.69 251.67 C316.54 263.92 303.60 277.92 292.40 292.40 C281.25 306.97 271.61 322.85 263.63 338.84 C255.71 354.91 249.45 371.85 244.67 388.60 C239.97 405.42 236.95 422.67 235.20 439.56 C233.53 456.49 233.45 473.50 234.43 490.06 C235.50 506.65 237.98 523.03 241.36 538.98 C244.83 554.93 249.50 570.51 254.98 585.74 C260.55 600.93 266.70 616.07 274.50 630.19 C282.38 644.27 291.42 658.27 302.03 670.37 C312.70 682.41 325.29 693.45 338.39 702.66 C351.55 711.80 366.18 719.52 380.84 725.46 C395.56 731.33 411.25 735.52 426.56 738.09 C441.92 740.58 457.74 741.31 472.88 740.68 C488.06 739.96 503.24 737.53 517.54 734.03 C531.87 730.45 545.80 725.33 558.80 719.44 C571.80 713.45 584.13 706.19 595.55 698.41 C606.96 690.54 617.52 681.68 627.30 672.48 C637.03 663.19 645.87 653.21 654.07 642.95 C662.21 632.63 669.49 621.85 676.28 610.77 C683.00 599.63 689.47 586.96 694.55 576.35 Z`

	__FrontCurve__000005_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000005_Rotated_1_.Path = `M821.30 622.88 C828.71 605.11 833.70 590.92 838.97 574.26 C843.82 557.63 848.05 540.76 851.58 523.26 C854.71 505.91 857.25 488.26 858.88 469.95 C860.14 451.91 860.79 433.46 860.24 414.46 C859.37 395.83 857.75 376.65 854.62 357.24 C851.23 338.25 846.82 318.67 840.68 299.33 C834.34 280.43 826.61 261.08 817.14 242.53 C807.53 224.42 796.19 206.19 783.32 189.25 C770.36 172.75 755.49 156.59 739.47 142.09 C723.38 128.00 705.43 114.75 686.77 103.35 C668.05 92.36 647.71 82.63 627.10 74.82 C606.45 67.41 587.26 61.71 562.73 57.65 C538.15 54.08 506.92 51.74 479.36 52.90 C451.82 54.67 423.60 59.42 397.06 66.46 C370.58 74.09 344.25 84.66 319.96 96.98 C295.82 109.90 272.54 125.49 251.52 142.29 C230.72 159.66 211.35 179.29 194.31 199.60 C177.57 220.45 162.68 243.04 150.03 265.84 C137.78 289.13 127.61 313.61 119.51 337.93 C111.89 362.65 106.43 388.03 102.81 412.98 C99.74 438.23 98.78 463.65 99.41 488.51 C100.65 513.53 103.82 538.31 108.40 562.52 C113.60 586.72 120.50 610.41 128.71 633.57 C137.50 656.53 147.38 679.28 159.32 700.67 C171.78 721.74 185.98 742.63 201.97 761.04 C218.41 779.07 237.38 795.94 256.91 810.34 C276.88 824.31 298.84 836.56 320.78 846.39 C343.11 855.75 366.81 863.02 389.97 868.07 C413.49 872.58 437.73 874.85 461.04 875.18 C484.63 874.94 508.32 872.46 530.84 868.40 C553.54 863.72 575.81 856.98 596.79 849.01 C617.82 840.39 638.01 830.00 656.91 818.69 C675.73 806.73 693.41 793.36 709.91 779.28 C726.17 764.61 741.14 748.88 755.07 732.56 C768.61 715.74 780.85 698.20 792.17 680.05 C802.98 661.55 813.40 640.54 821.30 622.88 Z`

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
	__Parameter__000000_Reference.BackendColor = `#F1F1F1`
	__Parameter__000000_Reference.MinuteColor = `#5A5A5A`
	__Parameter__000000_Reference.HourColor = `#1E1E1E`
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
	__Parameter__000000_Reference.HourHandleDiskDistance = 1.090000
	__Parameter__000000_Reference.HourHandleRadius = 0.210000
	__Parameter__000000_Reference.MinuteHandleRotationAngle = -120.000000
	__Parameter__000000_Reference.MinuteHandleDiskDistance = 1.775000
	__Parameter__000000_Reference.MinuteHandleRadius = 0.090000
	__Parameter__000000_Reference.MinuteOffset = 24.000000
	__Parameter__000000_Reference.BackendHandleRotationAngle = -120.500000
	__Parameter__000000_Reference.BackendHandleDiskDistance = 2.640000
	__Parameter__000000_Reference.BackendHandleRadius = 0.030000
	__Parameter__000000_Reference.BackendOffset = 159.000000
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
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = false

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
	__SpiralCircle__000000_Backend_Marker.CenterY = 414.480000
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
	__SpiralCircle__000000_Backend_Marker.Path = `M 504.710000 85.520000 A 4.710000 4.710000 0 1 0 495.290000 85.520000 A 4.710000 4.710000 0 1 0 504.710000 85.520000 Z`

	__SpiralCircle__000001_Construction_Circle_Spiral.Name = `Construction Circle Spiral`
	__SpiralCircle__000001_Construction_Circle_Spiral.IsDisplayed = false
	__SpiralCircle__000001_Construction_Circle_Spiral.CenterX = 56.520000
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
	__SpiralCircle__000001_Construction_Circle_Spiral.Path = ``

	__SpiralCircle__000002_Hour_Marker.Name = `Hour Marker`
	__SpiralCircle__000002_Hour_Marker.IsDisplayed = true
	__SpiralCircle__000002_Hour_Marker.CenterX = 0.000000
	__SpiralCircle__000002_Hour_Marker.CenterY = 170.345000
	__SpiralCircle__000002_Hour_Marker.HasBespokeRadius = true
	__SpiralCircle__000002_Hour_Marker.BespopkeRadius = 29.830000
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
	__SpiralCircle__000002_Hour_Marker.Path = `M 529.830000 329.655000 A 29.830000 29.830000 0 1 0 470.170000 329.655000 A 29.830000 29.830000 0 1 0 529.830000 329.655000 Z`

	__SpiralCircle__000003_Minute_Marker.Name = `Minute Marker`
	__SpiralCircle__000003_Minute_Marker.IsDisplayed = true
	__SpiralCircle__000003_Minute_Marker.CenterX = 0.000000
	__SpiralCircle__000003_Minute_Marker.CenterY = 278.675000
	__SpiralCircle__000003_Minute_Marker.HasBespokeRadius = true
	__SpiralCircle__000003_Minute_Marker.BespopkeRadius = 14.130000
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
	__SpiralCircle__000003_Minute_Marker.Path = `M 514.130000 221.325000 A 14.130000 14.130000 0 1 0 485.870000 221.325000 A 14.130000 14.130000 0 1 0 514.130000 221.325000 Z`

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
	__FrontCurveStack__000000_Backend_curve.FrontCurves = append(__FrontCurveStack__000000_Backend_curve.FrontCurves, __FrontCurve__000005_Rotated_1_)
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
