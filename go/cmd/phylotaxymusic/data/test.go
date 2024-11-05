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
	__Axis__000000_Construction_Axis.AngleDegree = 100.893395
	__Axis__000000_Construction_Axis.Length = 79.000000
	__Axis__000000_Construction_Axis.CenterX = 0.000000
	__Axis__000000_Construction_Axis.CenterY = 0.000000
	__Axis__000000_Construction_Axis.EndX = -14.929597
	__Axis__000000_Construction_Axis.EndY = 77.576460
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
	__Axis__000001_Initial_Axis.Length = 209.014354
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
	__Axis__000004_Rotated_Axis.Length = 209.014354
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -7.464798
	__Bezier__000005_Growth_Curve_Seed.StartY = 38.788230
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 17.359669
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 43.565701
	__Bezier__000005_Growth_Curve_Seed.EndX = 52.253588
	__Bezier__000005_Growth_Curve_Seed.EndY = 90.505870
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 27.429121
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 85.728399
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
	__Circle__000001_Construction_Circle.CenterX = -7.464798
	__Circle__000001_Construction_Circle.CenterY = 38.788230
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 134.366370
	__Circle__000005_Rotated_Next_Circle.CenterY = 25.858820
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M620.16 627.98 C622.12 621.38 623.25 616.10 624.41 609.97 C625.57 603.84 626.49 597.61 627.11 591.22 C627.72 584.82 628.11 578.29 628.11 571.60 C628.11 564.91 627.87 558.04 627.11 551.07 C626.35 544.09 625.28 536.91 623.56 529.75 C621.83 522.59 619.68 515.24 616.79 508.12 C613.89 500.99 610.41 493.77 606.20 487.01 C601.98 480.26 597.05 473.58 591.50 467.57 C585.94 461.56 579.63 455.85 572.86 450.95 C566.10 446.04 558.63 441.66 550.90 438.14 C543.18 434.62 534.88 431.79 526.52 429.83 C518.16 427.87 510.47 426.53 500.76 426.40 C491.05 426.28 478.85 426.96 468.27 429.07 C457.68 431.18 447.07 434.66 437.26 439.06 C427.45 443.46 417.95 449.14 409.40 455.49 C400.85 461.85 392.90 469.31 385.95 477.19 C379.01 485.07 372.88 493.83 367.76 502.77 C362.63 511.71 358.45 521.29 355.22 530.84 C351.98 540.39 349.74 550.32 348.36 560.08 C346.97 569.83 346.57 579.73 346.90 589.35 C347.23 598.97 348.48 608.55 350.35 617.79 C352.22 627.03 354.91 636.08 358.12 644.79 C361.34 653.49 365.26 661.92 369.65 670.02 C374.04 678.12 378.84 686.12 384.47 693.39 C390.10 700.66 396.42 707.74 403.44 713.65 C410.46 719.56 418.45 724.74 426.61 728.86 C434.77 732.98 443.63 736.17 452.41 738.36 C461.18 740.55 470.37 741.72 479.27 741.98 C488.16 742.25 497.19 741.48 505.77 739.96 C514.35 738.45 522.81 735.95 530.73 732.87 C538.64 729.79 546.25 725.85 553.27 721.49 C560.28 717.12 566.85 712.04 572.83 706.68 C578.82 701.32 584.28 695.41 589.20 689.31 C594.13 683.22 598.49 676.72 602.40 670.10 C606.30 663.48 609.66 656.58 612.62 649.56 C615.58 642.54 618.19 634.58 620.16 627.98 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M384.47 693.39 C389.28 700.18 393.52 705.17 398.53 710.76 C403.54 716.35 408.86 721.76 414.54 726.95 C420.21 732.13 426.21 737.15 432.60 741.87 C438.99 746.60 445.72 751.15 452.88 755.28 C460.04 759.40 467.60 763.32 475.54 766.63 C483.48 769.93 491.88 772.91 500.54 775.09 C509.20 777.26 518.33 778.93 527.51 779.65 C536.70 780.37 546.27 780.38 555.67 779.40 C565.07 778.42 574.68 776.56 583.90 773.77 C593.12 770.98 602.34 767.21 610.98 762.64 C619.62 758.07 628.03 752.53 635.72 746.34 C643.41 740.15 650.08 734.16 657.12 725.50 C664.16 716.85 672.26 705.41 677.97 694.41 C683.68 683.41 688.22 671.44 691.37 659.50 C694.53 647.56 696.35 635.00 696.91 622.77 C697.46 610.54 696.62 598.04 694.68 586.11 C692.73 574.19 689.43 562.32 685.23 551.20 C681.02 540.09 675.57 529.32 669.43 519.40 C663.28 509.48 656.07 500.15 648.36 491.70 C640.65 483.25 632.06 475.55 623.16 468.73 C614.26 461.91 604.69 455.92 594.94 450.78 C585.19 445.63 574.97 441.35 564.67 437.86 C554.36 434.37 543.75 431.73 533.10 429.82 C522.45 427.91 511.57 426.53 500.76 426.40 C489.96 426.28 478.85 426.96 468.27 429.07 C457.68 431.18 447.07 434.66 437.26 439.06 C427.45 443.46 417.95 449.14 409.40 455.49 C400.85 461.85 392.90 469.31 385.95 477.19 C379.01 485.07 372.88 493.83 367.76 502.77 C362.63 511.71 358.45 521.29 355.22 530.84 C351.98 540.39 349.74 550.32 348.36 560.08 C346.97 569.83 346.57 579.73 346.90 589.35 C347.23 598.97 348.48 608.55 350.35 617.79 C352.22 627.03 354.91 636.08 358.12 644.79 C361.34 653.49 365.26 661.92 369.65 670.02 C374.04 678.12 379.66 686.60 384.47 693.39 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M500.76 426.40 C491.00 425.91 483.32 426.24 474.55 426.87 C465.77 427.50 456.95 428.58 448.12 430.17 C439.29 431.76 430.41 433.80 421.57 436.41 C412.73 439.03 403.82 442.11 395.06 445.86 C386.30 449.62 377.50 453.90 369.01 458.93 C360.52 463.96 352.05 469.62 344.14 476.05 C336.23 482.47 328.48 489.64 321.53 497.49 C314.59 505.34 308.03 513.98 302.46 523.14 C296.88 532.29 291.94 542.21 288.08 552.41 C284.22 562.61 281.21 573.45 279.29 584.33 C277.38 595.21 276.45 606.54 276.59 617.69 C276.72 628.83 277.38 638.91 280.10 651.20 C282.82 663.49 287.23 678.66 292.91 691.44 C298.59 704.22 305.91 716.65 314.19 727.87 C322.46 739.10 332.19 749.60 342.54 758.78 C352.90 767.96 364.46 776.13 376.33 782.93 C388.19 789.74 400.96 795.34 413.76 799.61 C426.55 803.88 439.93 806.84 453.08 808.57 C466.24 810.29 479.67 810.68 492.70 809.96 C505.72 809.25 518.74 807.24 531.22 804.29 C543.70 801.34 555.94 797.19 567.57 792.24 C579.19 787.29 590.40 781.28 600.96 774.60 C611.52 767.91 621.55 760.32 630.91 752.13 C640.27 743.95 649.28 735.13 657.12 725.50 C664.96 715.88 672.26 705.41 677.97 694.41 C683.68 683.41 688.22 671.44 691.37 659.50 C694.53 647.56 696.35 635.00 696.91 622.77 C697.46 610.54 696.62 598.04 694.68 586.11 C692.73 574.19 689.43 562.32 685.23 551.20 C681.02 540.09 675.57 529.32 669.43 519.40 C663.28 509.48 656.07 500.15 648.36 491.70 C640.65 483.25 632.06 475.55 623.16 468.73 C614.26 461.91 604.69 455.92 594.94 450.78 C585.19 445.63 574.97 441.35 564.67 437.86 C554.36 434.37 543.75 431.73 533.10 429.82 C522.45 427.91 510.52 426.89 500.76 426.40 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M657.12 725.50 C664.52 717.07 669.67 709.89 675.30 701.53 C680.93 693.17 686.16 684.43 690.90 675.34 C695.64 666.26 699.98 656.80 703.74 647.01 C707.50 637.22 710.85 627.06 713.48 616.60 C716.11 606.13 718.26 595.28 719.53 584.24 C720.80 573.19 721.47 561.76 721.09 550.33 C720.72 538.89 719.57 527.13 717.30 515.62 C715.03 504.11 711.79 492.43 707.47 481.26 C703.15 470.09 697.74 458.97 691.37 448.60 C685.01 438.23 677.52 428.16 669.27 419.01 C661.02 409.87 651.72 401.27 641.86 393.71 C632.00 386.15 622.74 379.77 610.11 373.64 C597.49 367.50 581.21 360.84 566.13 356.90 C551.04 352.97 535.14 350.67 519.60 350.03 C504.07 349.39 488.14 350.48 472.91 353.05 C457.67 355.61 442.47 359.89 428.20 365.41 C413.93 370.93 400.07 378.06 387.32 386.16 C374.56 394.25 362.52 403.77 351.70 413.99 C340.88 424.20 331.00 435.61 322.37 447.45 C313.75 459.29 306.23 472.06 299.94 485.04 C293.65 498.03 288.56 511.67 284.64 525.36 C280.72 539.04 278.02 553.13 276.41 567.13 C274.81 581.13 274.39 595.33 275.01 609.35 C275.62 623.36 277.11 637.52 280.10 651.20 C283.08 664.88 287.23 678.66 292.91 691.44 C298.59 704.22 305.91 716.65 314.19 727.87 C322.46 739.10 332.19 749.60 342.54 758.78 C352.90 767.96 364.46 776.13 376.33 782.93 C388.19 789.74 400.96 795.34 413.76 799.61 C426.55 803.88 439.93 806.84 453.08 808.57 C466.24 810.29 479.67 810.68 492.70 809.96 C505.72 809.25 518.74 807.24 531.22 804.29 C543.70 801.34 555.94 797.19 567.57 792.24 C579.19 787.29 590.40 781.28 600.96 774.60 C611.52 767.91 621.55 760.32 630.91 752.13 C640.27 743.95 649.72 733.94 657.12 725.50 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M280.10 651.20 C282.37 663.65 285.14 673.23 288.64 684.03 C292.13 694.83 296.27 705.53 301.06 716.01 C305.85 726.48 311.27 736.83 317.38 746.87 C323.49 756.91 330.24 766.80 337.72 776.25 C345.20 785.69 353.36 794.94 362.27 803.53 C371.18 812.12 380.83 820.40 391.17 827.79 C401.50 835.17 412.63 842.07 424.28 847.83 C435.92 853.60 448.35 858.63 461.05 862.40 C473.76 866.16 487.13 868.95 500.50 870.43 C513.87 871.91 527.71 872.22 541.29 871.27 C554.86 870.32 568.66 868.11 581.96 864.73 C595.26 861.35 607.10 857.69 621.07 850.99 C635.04 844.29 651.97 834.82 665.76 824.52 C679.55 814.23 692.49 802.13 703.79 789.21 C715.08 776.30 725.17 761.87 733.53 747.01 C741.89 732.15 748.78 716.14 753.96 700.06 C759.14 683.99 762.68 667.15 764.61 650.57 C766.54 634.00 766.77 617.03 765.54 600.62 C764.31 584.20 761.40 567.76 757.23 552.07 C753.06 536.38 747.31 520.99 740.50 506.50 C733.69 492.01 725.45 478.07 716.35 465.13 C707.25 452.19 696.91 439.99 685.88 428.84 C674.85 417.69 662.78 407.43 650.15 398.23 C637.52 389.03 624.12 380.53 610.11 373.64 C596.11 366.75 581.21 360.84 566.13 356.90 C551.04 352.97 535.14 350.67 519.60 350.03 C504.07 349.39 488.14 350.48 472.91 353.05 C457.67 355.61 442.47 359.89 428.20 365.41 C413.93 370.93 400.07 378.06 387.32 386.16 C374.56 394.25 362.52 403.77 351.70 413.99 C340.88 424.20 331.00 435.61 322.37 447.45 C313.75 459.29 306.23 472.06 299.94 485.04 C293.65 498.03 288.56 511.67 284.64 525.36 C280.72 539.04 278.02 553.13 276.41 567.13 C274.81 581.13 274.39 595.33 275.01 609.35 C275.62 623.36 277.83 638.75 280.10 651.20 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M610.11 373.64 C597.68 367.01 587.37 362.82 575.52 358.45 C563.67 354.07 551.45 350.35 539.00 347.37 C526.55 344.40 513.78 342.09 500.83 340.62 C487.88 339.14 474.63 338.36 461.30 338.52 C447.98 338.68 434.38 339.59 420.87 341.59 C407.36 343.59 393.64 346.45 380.25 350.50 C366.86 354.54 353.39 359.60 340.54 365.85 C327.69 372.09 314.97 379.50 303.14 387.96 C291.31 396.43 279.89 406.11 269.58 416.65 C259.26 427.19 249.63 438.89 241.24 451.21 C232.85 463.53 225.40 476.87 219.26 490.57 C213.13 504.27 208.21 516.89 204.43 533.39 C200.65 549.89 197.20 570.80 196.59 589.58 C195.97 608.37 197.44 627.65 200.73 646.11 C204.03 664.56 209.41 683.02 216.34 700.32 C223.27 717.61 232.18 734.45 242.33 749.90 C252.47 765.34 264.38 779.93 277.19 792.97 C290.00 806.01 304.30 817.88 319.17 828.14 C334.03 838.40 350.07 847.26 366.37 854.51 C382.66 861.76 399.80 867.48 416.92 871.64 C434.04 875.81 451.68 878.39 469.08 879.50 C486.48 880.62 504.10 880.16 521.31 878.36 C538.51 876.55 555.69 873.23 572.32 868.67 C588.94 864.10 605.50 858.35 621.07 850.99 C636.64 843.63 651.97 834.82 665.76 824.52 C679.55 814.23 692.49 802.13 703.79 789.21 C715.08 776.30 725.17 761.87 733.53 747.01 C741.89 732.15 748.78 716.14 753.96 700.06 C759.14 683.99 762.68 667.15 764.61 650.57 C766.54 634.00 766.77 617.03 765.54 600.62 C764.31 584.20 761.40 567.76 757.23 552.07 C753.06 536.38 747.31 520.99 740.50 506.50 C733.69 492.01 725.45 478.07 716.35 465.13 C707.25 452.19 696.91 439.99 685.88 428.84 C674.85 417.69 662.78 407.43 650.15 398.23 C637.52 389.03 622.55 380.27 610.11 373.64 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M621.07 850.99 C635.33 844.79 646.01 838.73 657.90 831.47 C669.78 824.20 681.34 816.15 692.38 807.40 C703.42 798.64 714.09 789.14 724.13 778.93 C734.17 768.72 743.79 757.78 752.63 746.15 C761.47 734.52 769.81 722.15 777.15 709.16 C784.50 696.17 791.21 682.42 796.70 668.19 C802.18 653.96 806.83 638.99 810.06 623.79 C813.30 608.58 815.45 592.74 816.12 576.95 C816.79 561.16 816.15 544.93 814.06 529.05 C811.97 513.18 808.44 497.13 803.57 481.71 C798.70 466.29 792.36 450.98 784.83 436.54 C777.30 422.11 769.89 409.43 758.42 395.09 C746.95 380.74 731.54 363.69 716.02 350.46 C700.50 337.22 683.18 325.41 665.31 315.67 C647.44 305.94 628.16 297.96 608.80 292.07 C589.44 286.18 569.13 282.25 549.14 280.32 C529.16 278.39 508.68 278.52 488.89 280.47 C469.10 282.42 449.27 286.43 430.41 292.03 C411.54 297.62 393.04 305.17 375.72 314.04 C358.41 322.91 341.79 333.56 326.51 345.26 C311.23 356.97 296.93 370.21 284.05 384.26 C271.16 398.30 259.46 413.64 249.22 429.54 C238.97 445.44 230.04 462.38 222.58 479.69 C215.11 496.99 208.76 515.07 204.43 533.39 C200.10 551.71 197.20 570.80 196.59 589.58 C195.97 608.37 197.44 627.65 200.73 646.11 C204.03 664.56 209.41 683.02 216.34 700.32 C223.27 717.61 232.18 734.45 242.33 749.90 C252.47 765.34 264.38 779.93 277.19 792.97 C290.00 806.01 304.30 817.88 319.17 828.14 C334.03 838.40 350.07 847.26 366.37 854.51 C382.66 861.76 399.80 867.48 416.92 871.64 C434.04 875.81 451.68 878.39 469.08 879.50 C486.48 880.62 504.10 880.16 521.31 878.36 C538.51 876.55 555.69 873.23 572.32 868.67 C588.94 864.10 606.81 857.19 621.07 850.99 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M204.43 533.39 C200.02 549.80 197.96 563.06 196.12 578.15 C194.28 593.24 193.36 608.59 193.41 623.95 C193.46 639.31 194.43 654.86 196.43 670.31 C198.44 685.77 201.37 701.34 205.42 716.68 C209.47 732.03 214.50 747.42 220.72 762.36 C226.95 777.30 234.25 792.19 242.75 806.33 C251.26 820.48 260.95 834.38 271.77 847.26 C282.59 860.13 294.65 872.48 307.65 883.56 C320.64 894.63 334.86 904.89 349.74 913.69 C364.62 922.50 380.59 930.21 396.92 936.38 C413.24 942.54 430.45 947.40 447.70 950.70 C464.96 953.99 480.64 956.21 500.46 956.14 C520.29 956.08 545.02 954.51 566.63 950.32 C588.25 946.12 609.88 939.48 630.13 930.98 C650.39 922.48 670.13 911.63 688.18 899.31 C706.23 886.99 723.28 872.55 738.45 857.04 C753.61 841.53 767.39 824.22 779.17 806.24 C790.96 788.27 801.06 768.87 809.16 749.21 C817.26 729.55 823.49 708.86 827.78 688.26 C832.07 667.67 834.40 646.45 834.91 625.64 C835.42 604.82 833.96 583.77 830.84 563.38 C827.72 542.98 822.70 522.69 816.18 503.26 C809.67 483.82 801.37 464.79 791.74 446.76 C782.11 428.73 771.04 411.14 758.42 395.09 C745.80 379.03 731.54 363.69 716.02 350.46 C700.50 337.22 683.18 325.41 665.31 315.67 C647.44 305.94 628.16 297.96 608.80 292.07 C589.44 286.18 569.13 282.25 549.14 280.32 C529.16 278.39 508.68 278.52 488.89 280.47 C469.10 282.42 449.27 286.43 430.41 292.03 C411.54 297.62 393.04 305.17 375.72 314.04 C358.41 322.91 341.79 333.56 326.51 345.26 C311.23 356.97 296.93 370.21 284.05 384.26 C271.16 398.30 259.46 413.64 249.22 429.54 C238.97 445.44 230.04 462.38 222.58 479.69 C215.11 496.99 208.84 516.98 204.43 533.39 Z`

	__FrontCurve__000008_Rotated_0_.Name = `Rotated 0 `
	__FrontCurve__000008_Rotated_0_.Path = `M620.43 627.49 C622.40 620.84 623.53 615.52 624.69 609.34 C625.85 603.17 626.77 596.90 627.38 590.45 C627.98 584.01 628.36 577.44 628.35 570.71 C628.33 563.97 628.07 557.06 627.29 550.04 C626.51 543.03 625.41 535.80 623.66 528.61 C621.91 521.41 619.72 514.02 616.80 506.86 C613.87 499.70 610.35 492.45 606.10 485.66 C601.84 478.87 596.87 472.16 591.28 466.12 C585.68 460.08 579.32 454.35 572.51 449.43 C565.70 444.50 558.18 440.10 550.42 436.56 C542.65 433.02 534.30 430.18 525.89 428.21 C517.49 426.24 509.76 424.89 500.00 424.75 C490.24 424.62 477.98 425.30 467.34 427.41 C456.71 429.53 446.05 433.01 436.19 437.42 C426.34 441.84 416.79 447.53 408.20 453.90 C399.61 460.28 391.62 467.77 384.65 475.68 C377.67 483.59 371.51 492.39 366.36 501.37 C361.21 510.35 357.01 519.98 353.76 529.57 C350.51 539.17 348.26 549.16 346.86 558.97 C345.47 568.78 345.07 578.74 345.40 588.41 C345.73 598.09 346.99 607.73 348.87 617.03 C350.75 626.33 353.45 635.44 356.68 644.21 C359.92 652.97 363.86 661.46 368.28 669.62 C372.70 677.77 377.54 685.82 383.20 693.14 C388.87 700.46 395.23 707.58 402.29 713.54 C409.36 719.49 417.39 724.71 425.60 728.86 C433.81 733.01 442.73 736.24 451.55 738.45 C460.38 740.66 469.63 741.85 478.58 742.13 C487.53 742.41 496.61 741.65 505.24 740.14 C513.88 738.62 522.40 736.12 530.36 733.04 C538.33 729.95 545.99 725.99 553.06 721.61 C560.12 717.22 566.73 712.12 572.77 706.73 C578.80 701.34 584.29 695.39 589.26 689.26 C594.22 683.12 598.62 676.59 602.55 669.92 C606.48 663.25 609.87 656.31 612.85 649.24 C615.83 642.16 618.46 634.14 620.43 627.49 Z`

	__FrontCurve__000009_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000009_Rotated_1_.Path = `M679.09 670.29 C683.39 660.42 686.16 652.42 689.04 643.15 C691.81 633.87 694.17 624.40 696.04 614.69 C697.81 605.01 699.16 595.12 699.95 585.02 C700.63 574.96 700.88 564.68 700.44 554.25 C699.91 543.89 698.88 533.29 697.02 522.68 C695.07 512.16 692.51 501.40 689.01 490.88 C685.44 480.45 681.09 469.88 675.79 459.81 C670.44 449.85 664.16 439.93 657.05 430.74 C649.89 421.67 641.74 412.89 632.93 405.02 C624.10 397.27 614.31 390.05 604.09 383.86 C593.85 377.78 582.77 372.44 571.50 368.19 C560.22 364.04 549.80 360.82 536.38 358.62 C522.95 356.56 505.94 355.09 490.86 355.66 C475.78 356.41 460.38 358.83 445.81 362.59 C431.27 366.52 416.84 372.10 403.47 378.74 C390.14 385.55 377.30 393.87 365.67 402.98 C354.12 412.24 343.35 422.81 333.86 433.86 C324.47 445.06 316.10 457.31 308.99 469.77 C302.01 482.37 296.19 495.72 291.58 509.07 C287.10 522.51 283.86 536.43 281.72 550.18 C279.74 563.99 278.97 578.01 279.21 591.74 C279.60 605.50 281.16 619.24 283.61 632.62 C286.22 645.98 289.90 659.17 294.38 671.95 C299.03 684.68 304.42 697.30 310.99 709.12 C317.71 720.85 325.46 732.41 334.27 742.62 C343.21 752.73 353.53 762.13 364.30 770.17 C375.18 778.09 387.17 785.01 399.28 790.56 C411.50 795.99 424.50 800.21 437.34 803.15 C450.25 805.94 463.62 807.44 476.59 807.78 C489.61 807.96 502.77 806.84 515.34 804.74 C527.95 802.47 540.40 798.97 552.16 794.67 C563.92 790.20 575.29 784.62 585.92 778.42 C596.52 772.06 606.55 764.76 615.85 756.97 C625.07 749.03 633.62 740.34 641.46 731.26 C649.20 722.05 656.21 712.26 662.56 702.14 C668.76 691.93 674.65 680.12 679.09 670.29 Z`

	__FrontCurve__000010_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000010_Rotated_1_.Path = `M819.80 722.30 C826.79 704.50 831.34 690.18 835.93 673.60 C840.04 656.99 843.36 640.10 845.84 622.88 C847.82 605.77 849.03 588.40 849.29 570.76 C849.08 553.34 848.07 535.67 845.98 517.85 C843.47 500.36 840.08 482.61 835.48 464.94 C830.53 447.68 824.55 430.19 817.29 413.10 C809.75 396.46 800.99 379.72 791.01 363.74 C780.81 348.23 769.23 332.84 756.63 318.54 C743.84 304.71 729.62 291.30 714.64 279.20 C699.52 267.58 683.04 256.70 666.10 247.25 C649.05 238.28 630.83 230.33 612.44 223.86 C593.98 217.86 577.00 213.28 555.34 209.80 C533.60 206.91 506.13 205.06 481.81 205.91 C457.51 207.50 432.67 211.39 409.09 217.13 C385.63 223.61 362.20 232.37 340.36 242.64 C318.72 253.62 297.66 266.73 278.38 280.96 C259.42 295.88 241.48 312.67 225.43 330.18 C209.80 348.32 195.55 368.00 183.20 388.03 C171.37 408.61 161.16 430.34 152.79 452.11 C145.04 474.31 139.04 497.26 134.77 519.98 C131.19 543.01 129.40 566.40 129.20 589.37 C129.75 612.47 132.03 635.59 135.77 658.15 C140.27 680.67 146.40 702.91 153.87 724.52 C162.06 745.88 171.57 766.99 182.71 786.90 C194.48 806.41 207.91 825.62 222.68 842.89 C238.02 859.69 255.37 875.58 273.29 889.43 C291.72 902.75 311.81 914.72 332.02 924.66 C352.67 933.99 374.55 941.67 396.15 947.41 C418.10 952.48 440.82 955.71 462.91 957.17 C485.25 957.92 507.89 956.77 529.64 954.08 C551.52 950.63 573.26 945.34 593.94 938.75 C614.60 931.39 634.75 922.36 653.74 912.24 C672.54 901.38 690.54 889.06 707.34 875.86 C723.79 861.99 739.22 846.89 753.45 831.08 C767.18 814.71 779.75 797.37 791.14 779.41 C801.90 761.05 812.22 739.95 819.80 722.30 Z`

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
	__FrontCurveStack__000001_Front_Curve_Stack.IsDisplayed = true
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
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.ShiftToNearestCircle = 2
	__Parameter__000000_Reference.InsideAngle = 120.000000
	__Parameter__000000_Reference.SideLength = 79.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.320000
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
	__Parameter__000000_Reference.OriginX = 440.000000
	__Parameter__000000_Reference.OriginY = 350.000000
	__Parameter__000000_Reference.SpiralOriginX = 500.000000
	__Parameter__000000_Reference.SpiralOriginY = 600.000000
	__Parameter__000000_Reference.OriginCrossWidth = 800.000000
	__Parameter__000000_Reference.SpiralRadiusRatio = 1.060000
	__Parameter__000000_Reference.ShowSpiralBezierConstruct = false
	__Parameter__000000_Reference.ShowInterpolationPoints = false

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 79.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 79.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 134.366370
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 25.858820
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 79.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 79.000000
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
	__ShapeCategory__000003_3_Growing.IsExpanded = true

	__ShapeCategory__000004_4_Construction.Name = `4. Construction`
	__ShapeCategory__000004_4_Construction.IsExpanded = false

	__ShapeCategory__000005_5_Vertical_growth.Name = `5. Vertical growth`
	__ShapeCategory__000005_5_Vertical_growth.IsExpanded = false

	__ShapeCategory__000006_6_Spiral_growth.Name = `6. Spiral growth`
	__ShapeCategory__000006_6_Spiral_growth.IsExpanded = true

	__ShapeCategory__000007_7_Spiral_Growth_Bezier.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = true

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = false

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = false

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 119.456191
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -27.265096
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 153.638186
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = -6.557163
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 0.000000
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 174.245870
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 26.757616
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 136.322232
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
	__SpiralCircle__000000_Backend_Marker.CenterY = 208.560000
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
	__SpiralCircle__000001_Construction_Circle_Spiral.CenterX = 83.740000
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
	__SpiralCircle__000001_Construction_Circle_Spiral.Path = `M 597.870000 600.000000 A 14.130000 14.130000 0 1 0 569.610000 600.000000 A 14.130000 14.130000 0 1 0 597.870000 600.000000 Z`

	__SpiralCircle__000002_Hour_Marker.Name = `Hour Marker`
	__SpiralCircle__000002_Hour_Marker.IsDisplayed = false
	__SpiralCircle__000002_Hour_Marker.CenterX = 0.000000
	__SpiralCircle__000002_Hour_Marker.CenterY = 126.005000
	__SpiralCircle__000002_Hour_Marker.HasBespokeRadius = true
	__SpiralCircle__000002_Hour_Marker.BespopkeRadius = 22.910000
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
	__SpiralCircle__000002_Hour_Marker.Path = `M 522.910000 473.995000 A 22.910000 22.910000 0 1 0 477.090000 473.995000 A 22.910000 22.910000 0 1 0 522.910000 473.995000 Z`

	__SpiralCircle__000003_Minute_Marker.Name = `Minute Marker`
	__SpiralCircle__000003_Minute_Marker.IsDisplayed = false
	__SpiralCircle__000003_Minute_Marker.CenterX = 0.000000
	__SpiralCircle__000003_Minute_Marker.CenterY = 207.770000
	__SpiralCircle__000003_Minute_Marker.HasBespokeRadius = true
	__SpiralCircle__000003_Minute_Marker.BespopkeRadius = 12.640000
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
	__SpiralCircle__000003_Minute_Marker.Path = `M 512.640000 392.230000 A 12.640000 12.640000 0 1 0 487.360000 392.230000 A 12.640000 12.640000 0 1 0 512.640000 392.230000 Z`

	__SpiralCircleGrid__000000_Brute_Spiral_Bezier_Circle_Grid.Name = `Brute Spiral Bezier Circle Grid`
	__SpiralCircleGrid__000000_Brute_Spiral_Bezier_Circle_Grid.IsDisplayed = true

	__SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid.Name = `Construction Circle Spiral Full Grid`
	__SpiralCircleGrid__000001_Construction_Circle_Spiral_Full_Grid.IsDisplayed = false

	__SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid.Name = `Construction Circle Spiral Grid`
	__SpiralCircleGrid__000002_Construction_Circle_Spiral_Grid.IsDisplayed = false

	__SpiralCircleGrid__000003_Spiral_Circle_Grid.Name = `Spiral Circle Grid`
	__SpiralCircleGrid__000003_Spiral_Circle_Grid.IsDisplayed = true

	__SpiralLine__000000_Spiral_Contruction_Inner_Line.Name = `Spiral Contruction Inner Line`
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.IsDisplayed = false
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 119.456191
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 83.740000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -27.265096
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 119.456191
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 145.341108
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -27.265096
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -69.992589
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
	__SpiralOrigin__000000_Spiral_Origin.IsDisplayed = true
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 43.824735
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 10.002710
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = -41.943285
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 87.096129
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 119.456191
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -27.265096
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = -30.723564
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -63.798137
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
