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
	__NoteInfo__000015_22 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000016_23 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000017_3 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000018_4 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000019_5 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000020_6 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000021_7 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000022_8 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000023_9 := (&models.NoteInfo{}).Stage(stage)

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
	__Axis__000000_Construction_Axis.AngleDegree = 97.682917
	__Axis__000000_Construction_Axis.Length = 201.309445
	__Axis__000000_Construction_Axis.CenterX = 0.000000
	__Axis__000000_Construction_Axis.CenterY = 0.000000
	__Axis__000000_Construction_Axis.EndX = -26.913202
	__Axis__000000_Construction_Axis.EndY = 199.502312
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
	__Axis__000001_Initial_Axis.AngleDegree = 82.317083
	__Axis__000001_Initial_Axis.Length = 752.892451
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
	__Axis__000004_Rotated_Axis.Length = 752.892451
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -13.456601
	__Bezier__000005_Growth_Curve_Seed.StartY = 99.751156
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 65.032429
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 110.339460
	__Bezier__000005_Growth_Curve_Seed.EndX = 120.973968
	__Bezier__000005_Growth_Curve_Seed.EndY = 219.452543
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 42.484939
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 208.864239
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
	__Circle__000001_Construction_Circle.CenterX = -13.456601
	__Circle__000001_Construction_Circle.CenterY = 99.751156
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 295.774340
	__Circle__000005_Rotated_Next_Circle.CenterY = 39.900462
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M789.55 533.12 C791.26 523.28 792.37 515.17 793.61 506.05 C794.85 496.92 795.88 487.76 796.99 478.36 C798.10 468.96 799.05 459.52 800.28 449.66 C801.50 439.80 802.71 429.85 804.34 419.20 C805.98 408.55 808.03 397.56 810.08 385.76 C812.13 373.97 815.20 361.22 816.65 348.41 C818.09 335.60 819.54 321.85 818.78 308.91 C818.02 295.97 815.70 282.97 812.09 270.77 C808.48 258.57 803.18 246.81 797.11 235.73 C791.05 224.64 783.66 214.12 775.71 204.27 C767.76 194.42 758.80 185.17 749.41 176.61 C740.01 168.04 730.99 160.67 719.32 152.89 C707.65 145.11 693.23 136.43 679.39 129.93 C665.54 123.43 650.96 117.91 636.24 113.89 C621.53 109.86 606.22 107.03 591.09 105.78 C575.95 104.53 560.44 104.72 545.43 106.37 C530.43 108.01 515.36 111.36 501.05 115.67 C486.73 119.99 472.78 126.01 459.56 132.25 C446.35 138.48 433.81 145.92 421.75 153.10 C409.69 160.29 398.34 167.92 387.22 175.34 C376.11 182.76 365.52 190.18 355.07 197.61 C344.63 205.03 334.52 212.37 324.54 219.91 C314.57 227.45 304.82 234.98 295.24 242.83 C285.66 250.68 275.54 259.41 267.08 267.01 C258.61 274.61 251.94 281.11 244.47 288.45 C237.01 295.78 229.73 303.32 222.29 311.01 C214.85 318.70 207.54 326.56 199.82 334.61 C192.10 342.65 184.39 350.80 175.96 359.26 C167.53 367.71 158.61 376.19 149.25 385.32 C139.88 394.45 129.19 403.71 119.78 414.03 C110.37 424.36 100.43 435.43 92.78 447.25 C85.14 459.07 78.73 471.92 73.90 484.94 C69.07 497.97 65.89 511.71 63.79 525.40 C61.69 539.07 61.04 553.14 61.30 567.02 C61.56 580.91 63.05 594.94 65.37 608.70 C67.68 622.45 70.44 634.92 75.18 649.56 C79.93 664.20 86.41 681.47 93.84 696.53 C101.27 711.58 109.97 726.29 119.77 739.87 C129.56 753.45 140.64 766.43 152.63 778.01 C164.62 789.59 177.88 800.24 191.70 809.36 C205.52 818.48 220.51 826.27 235.55 832.75 C250.60 839.24 266.49 844.11 281.97 848.27 C297.45 852.43 313.20 855.17 328.44 857.73 C343.68 860.28 358.70 861.98 373.40 863.61 C388.11 865.23 402.44 866.46 416.68 867.50 C430.92 868.55 444.86 869.37 458.82 869.86 C472.78 870.36 486.58 870.64 500.43 870.47 C514.28 870.31 528.15 870.19 541.93 868.88 C555.70 867.57 569.64 865.87 583.07 862.64 C596.51 859.40 609.91 855.03 622.53 849.48 C635.15 843.93 647.50 837.17 658.79 829.35 C670.07 821.54 680.78 812.44 690.24 802.58 C699.69 792.73 708.18 781.59 715.53 770.22 C722.88 758.86 728.93 746.46 734.34 734.40 C739.74 722.35 743.93 709.85 747.96 697.90 C751.99 685.94 755.26 674.13 758.54 662.68 C761.83 651.23 764.77 640.16 767.67 629.20 C770.57 618.24 773.33 607.61 775.94 596.92 C778.56 586.24 781.07 575.74 783.34 565.10 C785.61 554.47 787.84 542.96 789.55 533.12 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M267.08 267.01 C258.61 274.61 251.94 281.11 244.47 288.45 C237.01 295.78 229.73 303.32 222.29 311.01 C214.85 318.70 207.54 326.56 199.82 334.61 C192.10 342.65 184.39 350.80 175.96 359.26 C167.53 367.71 158.61 376.19 149.25 385.32 C139.88 394.45 129.19 403.71 119.78 414.03 C110.37 424.36 100.43 435.43 92.78 447.25 C85.14 459.07 78.73 471.92 73.90 484.94 C69.07 497.97 65.89 511.71 63.79 525.40 C61.69 539.07 61.04 553.14 61.30 567.02 C61.56 580.91 63.05 594.94 65.37 608.70 C67.68 622.45 70.44 634.92 75.18 649.56 C79.93 664.20 86.41 681.47 93.84 696.53 C101.27 711.58 109.97 726.29 119.77 739.87 C129.56 753.45 140.64 766.43 152.63 778.01 C164.62 789.59 177.88 800.24 191.70 809.36 C205.52 818.48 220.51 826.27 235.55 832.75 C250.60 839.24 266.49 844.11 281.97 848.27 C297.45 852.43 313.20 855.17 328.44 857.73 C343.68 860.28 358.70 861.98 373.40 863.61 C388.11 865.23 402.44 866.46 416.68 867.50 C430.92 868.55 444.86 869.37 458.82 869.86 C472.78 870.36 486.58 870.64 500.43 870.47 C514.28 870.31 529.23 869.69 541.93 868.88 C554.62 868.07 564.99 866.95 576.59 865.61 C588.20 864.28 599.78 862.61 611.56 860.86 C623.34 859.11 635.11 857.08 647.28 855.13 C659.46 853.17 671.69 851.06 684.63 849.12 C697.57 847.17 710.89 845.45 724.92 843.46 C738.95 841.46 754.07 840.14 768.82 837.15 C783.57 834.15 799.17 830.84 813.40 825.49 C827.63 820.14 841.50 813.17 854.21 805.04 C866.91 796.92 878.74 787.15 889.63 776.74 C900.52 766.33 910.48 754.67 919.56 742.57 C928.64 730.47 936.81 717.45 944.10 704.11 C951.39 690.78 957.46 678.27 963.33 662.56 C969.20 646.86 975.39 627.75 979.31 609.88 C983.23 592.02 985.87 573.61 986.86 555.39 C987.85 537.18 987.38 518.62 985.25 500.59 C983.12 482.56 979.32 464.47 974.06 447.24 C968.79 430.00 961.69 413.09 953.66 397.18 C945.63 381.27 935.85 366.11 925.87 351.79 C915.88 337.47 904.72 324.12 893.77 311.28 C882.82 298.43 871.43 286.43 860.17 274.72 C848.91 263.01 837.62 251.88 826.22 241.02 C814.83 230.16 803.46 219.70 791.81 209.55 C780.17 199.40 768.45 189.57 756.37 180.13 C744.29 170.68 732.15 161.25 719.32 152.89 C706.49 144.52 693.23 136.43 679.39 129.93 C665.54 123.43 650.96 117.91 636.24 113.89 C621.53 109.86 606.22 107.03 591.09 105.78 C575.95 104.53 560.44 104.72 545.43 106.37 C530.43 108.01 515.36 111.36 501.05 115.67 C486.73 119.99 472.78 126.01 459.56 132.25 C446.35 138.48 433.81 145.92 421.75 153.10 C409.69 160.29 398.34 167.92 387.22 175.34 C376.11 182.76 365.52 190.18 355.07 197.61 C344.63 205.03 334.52 212.37 324.54 219.91 C314.57 227.45 304.82 234.98 295.24 242.83 C285.66 250.68 275.54 259.41 267.08 267.01 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M541.93 868.88 C554.62 868.07 564.99 866.95 576.59 865.61 C588.20 864.28 599.78 862.61 611.56 860.86 C623.34 859.11 635.11 857.08 647.28 855.13 C659.46 853.17 671.69 851.06 684.63 849.12 C697.57 847.17 710.89 845.45 724.92 843.46 C738.95 841.46 754.07 840.14 768.82 837.15 C783.57 834.15 799.17 830.84 813.40 825.49 C827.63 820.14 841.50 813.17 854.21 805.04 C866.91 796.92 878.74 787.15 889.63 776.74 C900.52 766.33 910.48 754.67 919.56 742.57 C928.64 730.47 936.81 717.45 944.10 704.11 C951.39 690.78 957.46 678.27 963.33 662.56 C969.20 646.86 975.39 627.75 979.31 609.88 C983.23 592.02 985.87 573.61 986.86 555.39 C987.85 537.18 987.38 518.62 985.25 500.59 C983.12 482.56 979.32 464.47 974.06 447.24 C968.79 430.00 961.69 413.09 953.66 397.18 C945.63 381.27 935.85 366.11 925.87 351.79 C915.88 337.47 904.72 324.12 893.77 311.28 C882.82 298.43 871.43 286.43 860.17 274.72 C848.91 263.01 837.62 251.88 826.22 241.02 C814.83 230.16 803.46 219.70 791.81 209.55 C780.17 199.40 768.45 189.57 756.37 180.13 C744.29 170.68 730.89 160.91 719.32 152.89 C707.76 144.86 697.98 138.73 686.97 131.97 C675.97 125.22 664.75 118.79 653.30 112.35 C641.86 105.91 630.22 99.75 618.30 93.33 C606.38 86.92 594.30 80.65 581.76 73.86 C569.23 67.07 556.52 59.93 543.08 52.61 C529.64 45.29 515.76 36.82 501.13 29.96 C486.51 23.09 471.01 15.99 455.32 11.41 C439.64 6.82 423.22 3.79 407.01 2.45 C390.81 1.11 374.27 1.65 358.08 3.35 C341.89 5.04 325.67 8.35 309.88 12.62 C294.09 16.89 278.46 22.51 263.35 28.97 C248.23 35.43 234.70 42.00 219.18 51.41 C203.64 60.82 185.54 72.89 170.15 85.42 C154.74 97.95 140.03 111.78 126.78 126.53 C113.53 141.28 101.25 157.26 90.69 173.88 C80.12 190.50 70.89 208.25 63.42 226.28 C55.97 244.31 50.21 263.27 45.91 282.06 C41.61 300.84 39.22 320.18 37.61 339.00 C35.99 357.82 35.98 376.65 36.20 394.97 C36.42 413.29 37.60 431.25 38.94 448.90 C40.28 466.55 42.10 483.79 44.24 500.87 C46.38 517.95 48.84 534.71 51.79 551.38 C54.74 568.05 58.04 584.52 61.94 600.88 C65.84 617.24 69.87 633.62 75.18 649.56 C80.50 665.50 86.41 681.47 93.84 696.53 C101.27 711.58 109.97 726.29 119.77 739.87 C129.56 753.45 140.64 766.43 152.63 778.01 C164.62 789.59 177.88 800.24 191.70 809.36 C205.52 818.48 220.51 826.27 235.55 832.75 C250.60 839.24 266.49 844.11 281.97 848.27 C297.45 852.43 313.20 855.17 328.44 857.73 C343.68 860.28 358.70 861.98 373.40 863.61 C388.11 865.23 402.44 866.46 416.68 867.50 C430.92 868.55 444.86 869.37 458.82 869.86 C472.78 870.36 486.58 870.64 500.43 870.47 C514.28 870.31 529.23 869.69 541.93 868.88 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M719.32 152.89 C707.76 144.86 697.98 138.73 686.97 131.97 C675.97 125.22 664.75 118.79 653.30 112.35 C641.86 105.91 630.22 99.75 618.30 93.33 C606.38 86.92 594.30 80.65 581.76 73.86 C569.23 67.07 556.52 59.93 543.08 52.61 C529.64 45.29 515.76 36.82 501.13 29.96 C486.51 23.09 471.01 15.99 455.32 11.41 C439.64 6.82 423.22 3.79 407.01 2.45 C390.81 1.11 374.27 1.65 358.08 3.35 C341.89 5.04 325.67 8.35 309.88 12.62 C294.09 16.89 278.46 22.51 263.35 28.97 C248.23 35.43 234.70 42.00 219.18 51.41 C203.64 60.82 185.54 72.89 170.15 85.42 C154.74 97.95 140.03 111.78 126.78 126.53 C113.53 141.28 101.25 157.26 90.69 173.88 C80.12 190.50 70.89 208.25 63.42 226.28 C55.97 244.31 50.21 263.27 45.91 282.06 C41.61 300.84 39.22 320.18 37.61 339.00 C35.99 357.82 35.98 376.65 36.20 394.97 C36.42 413.29 37.60 431.25 38.94 448.90 C40.28 466.55 42.10 483.79 44.24 500.87 C46.38 517.95 48.84 534.71 51.79 551.38 C54.74 568.05 58.04 584.52 61.94 600.88 C65.84 617.24 70.70 634.78 75.18 649.56 C79.67 664.33 83.94 676.26 88.86 689.52 C93.77 702.79 99.12 715.92 104.69 729.16 C110.25 742.40 116.19 755.51 122.24 768.97 C128.29 782.43 134.60 795.84 140.99 809.90 C147.38 823.96 153.73 838.36 160.58 853.32 C167.43 868.27 173.95 884.36 182.09 899.61 C190.22 914.87 198.94 930.80 209.39 944.83 C219.85 958.87 231.89 972.10 244.83 983.84 C257.77 995.57 272.18 1005.98 287.04 1015.23 C301.91 1024.47 317.84 1032.43 334.03 1039.31 C350.22 1046.19 367.15 1051.87 384.20 1056.51 C401.26 1061.16 417.02 1064.67 436.36 1067.17 C455.69 1069.67 478.94 1071.67 500.21 1071.51 C521.47 1071.35 543.01 1069.65 563.95 1066.22 C584.89 1062.78 605.84 1057.68 625.85 1050.92 C645.85 1044.16 665.53 1035.59 683.98 1025.63 C702.42 1015.67 720.12 1003.82 736.52 991.15 C752.93 978.48 768.17 964.10 782.43 949.59 C796.70 935.08 809.68 919.48 822.10 904.08 C834.52 888.68 845.94 872.89 856.96 857.19 C867.98 841.49 878.32 825.76 888.23 809.89 C898.15 794.03 907.56 778.15 916.45 762.01 C925.34 745.87 933.77 729.62 941.58 713.04 C949.40 696.47 957.04 679.75 963.33 662.56 C969.62 645.37 975.39 627.75 979.31 609.88 C983.23 592.02 985.87 573.61 986.86 555.39 C987.85 537.18 987.38 518.62 985.25 500.59 C983.12 482.56 979.32 464.47 974.06 447.24 C968.79 430.00 961.69 413.09 953.66 397.18 C945.63 381.27 935.85 366.11 925.87 351.79 C915.88 337.47 904.72 324.12 893.77 311.28 C882.82 298.43 871.43 286.43 860.17 274.72 C848.91 263.01 837.62 251.88 826.22 241.02 C814.83 230.16 803.46 219.70 791.81 209.55 C780.17 199.40 768.45 189.57 756.37 180.13 C744.29 170.68 730.89 160.91 719.32 152.89 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M75.18 649.56 C79.67 664.33 83.94 676.26 88.86 689.52 C93.77 702.79 99.12 715.92 104.69 729.16 C110.25 742.40 116.19 755.51 122.24 768.97 C128.29 782.43 134.60 795.84 140.99 809.90 C147.38 823.96 153.73 838.36 160.58 853.32 C167.43 868.27 173.95 884.36 182.09 899.61 C190.22 914.87 198.94 930.80 209.39 944.83 C219.85 958.87 231.89 972.10 244.83 983.84 C257.77 995.57 272.18 1005.98 287.04 1015.23 C301.91 1024.47 317.84 1032.43 334.03 1039.31 C350.22 1046.19 367.15 1051.87 384.20 1056.51 C401.26 1061.16 417.02 1064.67 436.36 1067.17 C455.69 1069.67 478.94 1071.67 500.21 1071.51 C521.47 1071.35 543.01 1069.65 563.95 1066.22 C584.89 1062.78 605.84 1057.68 625.85 1050.92 C645.85 1044.16 665.53 1035.59 683.98 1025.63 C702.42 1015.67 720.12 1003.82 736.52 991.15 C752.93 978.48 768.17 964.10 782.43 949.59 C796.70 935.08 809.68 919.48 822.10 904.08 C834.52 888.68 845.94 872.89 856.96 857.19 C867.98 841.49 878.32 825.76 888.23 809.89 C898.15 794.03 907.56 778.15 916.45 762.01 C925.34 745.87 933.77 729.62 941.58 713.04 C949.40 696.47 957.15 678.20 963.33 662.56 C969.51 646.92 973.92 633.84 978.66 619.19 C983.39 604.54 987.65 589.71 991.76 574.65 C995.86 559.60 999.53 544.39 1003.28 528.85 C1007.02 513.31 1010.47 497.62 1014.20 481.40 C1017.94 465.17 1021.86 448.70 1025.68 431.49 C1029.50 414.29 1034.22 396.39 1037.10 378.17 C1039.98 359.95 1042.78 340.85 1042.98 322.16 C1043.18 303.47 1041.59 284.45 1038.31 266.05 C1035.02 247.66 1029.71 229.41 1023.27 211.81 C1016.82 194.22 1008.69 176.99 999.63 160.48 C990.58 143.97 980.15 127.96 968.94 112.73 C957.72 97.50 946.86 84.07 932.34 69.09 C917.83 54.12 899.71 36.95 881.84 22.88 C863.97 8.80 844.86 -4.18 825.10 -15.36 C805.35 -26.54 784.49 -36.35 763.31 -44.19 C742.13 -52.03 720.03 -58.19 698.01 -62.40 C675.99 -66.62 653.35 -68.81 631.18 -69.47 C609.01 -70.13 586.68 -68.65 565.01 -66.36 C543.35 -64.08 521.99 -60.06 501.20 -55.76 C480.42 -51.47 460.17 -46.15 440.30 -40.57 C420.43 -35.00 401.07 -28.88 381.99 -22.32 C362.92 -15.76 344.23 -8.78 325.83 -1.20 C307.42 6.39 289.32 14.44 271.55 23.21 C253.77 31.97 236.07 41.04 219.18 51.41 C202.27 61.78 185.54 72.89 170.15 85.42 C154.74 97.95 140.03 111.78 126.78 126.53 C113.53 141.28 101.25 157.26 90.69 173.88 C80.12 190.50 70.89 208.25 63.42 226.28 C55.97 244.31 50.21 263.27 45.91 282.06 C41.61 300.84 39.22 320.18 37.61 339.00 C35.99 357.82 35.98 376.65 36.20 394.97 C36.42 413.29 37.60 431.25 38.94 448.90 C40.28 466.55 42.10 483.79 44.24 500.87 C46.38 517.95 48.84 534.71 51.79 551.38 C54.74 568.05 58.04 584.52 61.94 600.88 C65.84 617.24 70.70 634.78 75.18 649.56 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M963.33 662.56 C969.51 646.92 973.92 633.84 978.66 619.19 C983.39 604.54 987.65 589.71 991.76 574.65 C995.86 559.60 999.53 544.39 1003.28 528.85 C1007.02 513.31 1010.47 497.62 1014.20 481.40 C1017.94 465.17 1021.86 448.70 1025.68 431.49 C1029.50 414.29 1034.22 396.39 1037.10 378.17 C1039.98 359.95 1042.78 340.85 1042.98 322.16 C1043.18 303.47 1041.59 284.45 1038.31 266.05 C1035.02 247.66 1029.71 229.41 1023.27 211.81 C1016.82 194.22 1008.69 176.99 999.63 160.48 C990.58 143.97 980.15 127.96 968.94 112.73 C957.72 97.50 946.86 84.07 932.34 69.09 C917.83 54.12 899.71 36.95 881.84 22.88 C863.97 8.80 844.86 -4.18 825.10 -15.36 C805.35 -26.54 784.49 -36.35 763.31 -44.19 C742.13 -52.03 720.03 -58.19 698.01 -62.40 C675.99 -66.62 653.35 -68.81 631.18 -69.47 C609.01 -70.13 586.68 -68.65 565.01 -66.36 C543.35 -64.08 521.99 -60.06 501.20 -55.76 C480.42 -51.47 460.17 -46.15 440.30 -40.57 C420.43 -35.00 401.07 -28.88 381.99 -22.32 C362.92 -15.76 344.23 -8.78 325.83 -1.20 C307.42 6.39 289.32 14.44 271.55 23.21 C253.77 31.97 234.92 42.27 219.18 51.41 C203.43 60.55 190.91 68.73 177.07 78.05 C163.22 87.37 149.66 97.18 136.11 107.31 C122.55 117.43 109.26 127.98 95.75 138.78 C82.23 149.58 68.89 160.69 55.04 172.07 C41.20 183.44 27.09 194.92 12.68 207.06 C-1.73 219.19 -17.23 231.32 -31.42 244.87 C-45.61 258.42 -60.22 272.74 -72.48 288.38 C-84.73 304.01 -95.73 321.05 -104.95 338.66 C-114.17 356.28 -121.66 375.07 -127.79 394.06 C-133.91 413.05 -138.43 432.81 -141.73 452.60 C-145.02 472.39 -146.87 492.64 -147.55 512.77 C-148.24 532.91 -148.05 551.30 -145.85 573.41 C-143.65 595.52 -139.92 621.82 -134.37 645.42 C-128.81 669.01 -121.53 692.50 -112.54 714.96 C-103.54 737.41 -92.76 759.47 -80.41 780.14 C-68.07 800.81 -53.89 820.71 -38.48 838.99 C-23.08 857.27 -5.85 874.32 12.02 889.82 C29.89 905.32 49.30 919.20 68.73 931.98 C88.16 944.75 108.50 955.93 128.62 966.46 C148.73 976.99 169.11 986.33 189.39 995.16 C209.68 1003.98 229.95 1012.00 250.31 1019.42 C270.67 1026.84 291.01 1033.62 311.54 1039.67 C332.07 1045.73 352.68 1051.15 373.48 1055.73 C394.28 1060.32 415.24 1064.54 436.36 1067.17 C457.48 1069.80 478.94 1071.67 500.21 1071.51 C521.47 1071.35 543.01 1069.65 563.95 1066.22 C584.89 1062.78 605.84 1057.68 625.85 1050.92 C645.85 1044.16 665.53 1035.59 683.98 1025.63 C702.42 1015.67 720.12 1003.82 736.52 991.15 C752.93 978.48 768.17 964.10 782.43 949.59 C796.70 935.08 809.68 919.48 822.10 904.08 C834.52 888.68 845.94 872.89 856.96 857.19 C867.98 841.49 878.32 825.76 888.23 809.89 C898.15 794.03 907.56 778.15 916.45 762.01 C925.34 745.87 933.77 729.62 941.58 713.04 C949.40 696.47 957.15 678.20 963.33 662.56 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M219.18 51.41 C203.43 60.55 190.91 68.73 177.07 78.05 C163.22 87.37 149.66 97.18 136.11 107.31 C122.55 117.43 109.26 127.98 95.75 138.78 C82.23 149.58 68.89 160.69 55.04 172.07 C41.20 183.44 27.09 194.92 12.68 207.06 C-1.73 219.19 -17.23 231.32 -31.42 244.87 C-45.61 258.42 -60.22 272.74 -72.48 288.38 C-84.73 304.01 -95.73 321.05 -104.95 338.66 C-114.17 356.28 -121.66 375.07 -127.79 394.06 C-133.91 413.05 -138.43 432.81 -141.73 452.60 C-145.02 472.39 -146.87 492.64 -147.55 512.77 C-148.24 532.91 -148.05 551.30 -145.85 573.41 C-143.65 595.52 -139.92 621.82 -134.37 645.42 C-128.81 669.01 -121.53 692.50 -112.54 714.96 C-103.54 737.41 -92.76 759.47 -80.41 780.14 C-68.07 800.81 -53.89 820.71 -38.48 838.99 C-23.08 857.27 -5.85 874.32 12.02 889.82 C29.89 905.32 49.30 919.20 68.73 931.98 C88.16 944.75 108.50 955.93 128.62 966.46 C148.73 976.99 169.11 986.33 189.39 995.16 C209.68 1003.98 229.95 1012.00 250.31 1019.42 C270.67 1026.84 291.01 1033.62 311.54 1039.67 C332.07 1045.73 352.68 1051.15 373.48 1055.73 C394.28 1060.32 417.02 1064.31 436.36 1067.17 C455.69 1070.04 471.67 1071.53 489.49 1072.93 C507.31 1074.33 525.21 1075.10 543.28 1075.56 C561.35 1076.02 579.47 1075.91 597.90 1075.67 C616.33 1075.44 634.83 1074.77 653.87 1074.15 C672.92 1073.53 692.23 1072.92 712.18 1071.95 C732.13 1070.97 752.91 1070.60 773.55 1068.28 C794.19 1065.95 815.63 1063.23 836.02 1057.98 C856.40 1052.73 876.67 1045.55 895.86 1036.77 C915.06 1027.99 933.61 1017.17 951.18 1005.31 C968.75 993.44 985.51 979.93 1001.28 965.60 C1017.05 951.27 1031.94 935.62 1045.80 919.32 C1059.67 903.02 1071.66 887.61 1084.46 867.79 C1097.26 847.96 1111.57 823.65 1122.61 800.40 C1133.64 777.15 1143.19 752.85 1150.68 728.28 C1158.16 703.71 1163.92 678.30 1167.51 652.98 C1171.11 627.67 1172.70 601.79 1172.24 576.41 C1171.79 551.03 1169.06 525.46 1164.77 500.71 C1160.49 475.96 1153.91 451.50 1146.54 427.91 C1139.17 404.32 1130.01 381.42 1120.54 359.19 C1111.07 336.96 1100.56 315.50 1089.74 294.54 C1078.92 273.58 1067.51 253.27 1055.60 233.43 C1043.68 213.58 1031.29 194.27 1018.24 175.47 C1005.19 156.67 991.61 138.35 977.30 120.62 C962.98 102.89 948.25 85.38 932.34 69.09 C916.43 52.80 899.71 36.95 881.84 22.88 C863.97 8.80 844.86 -4.18 825.10 -15.36 C805.35 -26.54 784.49 -36.35 763.31 -44.19 C742.13 -52.03 720.03 -58.19 698.01 -62.40 C675.99 -66.62 653.35 -68.81 631.18 -69.47 C609.01 -70.13 586.68 -68.65 565.01 -66.36 C543.35 -64.08 521.99 -60.06 501.20 -55.76 C480.42 -51.47 460.17 -46.15 440.30 -40.57 C420.43 -35.00 401.07 -28.88 381.99 -22.32 C362.92 -15.76 344.23 -8.78 325.83 -1.20 C307.42 6.39 289.32 14.44 271.55 23.21 C253.77 31.97 234.92 42.27 219.18 51.41 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M436.36 1067.17 C455.69 1070.04 471.67 1071.53 489.49 1072.93 C507.31 1074.33 525.21 1075.10 543.28 1075.56 C561.35 1076.02 579.47 1075.91 597.90 1075.67 C616.33 1075.44 634.83 1074.77 653.87 1074.15 C672.92 1073.53 692.23 1072.92 712.18 1071.95 C732.13 1070.97 752.91 1070.60 773.55 1068.28 C794.19 1065.95 815.63 1063.23 836.02 1057.98 C856.40 1052.73 876.67 1045.55 895.86 1036.77 C915.06 1027.99 933.61 1017.17 951.18 1005.31 C968.75 993.44 985.51 979.93 1001.28 965.60 C1017.05 951.27 1031.94 935.62 1045.80 919.32 C1059.67 903.02 1071.66 887.61 1084.46 867.79 C1097.26 847.96 1111.57 823.65 1122.61 800.40 C1133.64 777.15 1143.19 752.85 1150.68 728.28 C1158.16 703.71 1163.92 678.30 1167.51 652.98 C1171.11 627.67 1172.70 601.79 1172.24 576.41 C1171.79 551.03 1169.06 525.46 1164.77 500.71 C1160.49 475.96 1153.91 451.50 1146.54 427.91 C1139.17 404.32 1130.01 381.42 1120.54 359.19 C1111.07 336.96 1100.56 315.50 1089.74 294.54 C1078.92 273.58 1067.51 253.27 1055.60 233.43 C1043.68 213.58 1031.29 194.27 1018.24 175.47 C1005.19 156.67 991.61 138.35 977.30 120.62 C962.98 102.89 946.61 84.38 932.34 69.09 C918.07 53.81 905.68 41.93 891.68 28.92 C877.67 15.91 863.18 3.38 848.32 -8.95 C833.47 -21.29 818.18 -33.17 802.55 -45.09 C786.92 -57.01 770.94 -68.59 754.54 -80.50 C738.13 -92.40 721.50 -104.43 704.13 -116.51 C686.75 -128.59 669.08 -141.63 650.29 -152.97 C631.50 -164.32 611.82 -175.81 591.39 -184.60 C570.95 -193.39 549.39 -200.54 527.68 -205.71 C505.97 -210.89 483.47 -214.00 461.11 -215.63 C438.74 -217.27 415.97 -217.05 393.49 -215.51 C371.00 -213.98 348.38 -210.81 326.18 -206.43 C303.97 -202.05 283.91 -197.20 260.26 -189.24 C236.61 -181.29 208.76 -170.63 184.29 -158.71 C159.82 -146.80 135.88 -133.06 113.44 -117.75 C91.00 -102.45 69.41 -85.34 49.64 -66.90 C29.82 -48.48 11.18 -28.32 -6.37 -6.96 C-23.03 15.51 -37.70 38.47 -50.63 61.71 C-63.59 84.91 -74.52 109.38 -84.21 133.69 C-93.91 157.99 -101.72 182.93 -108.79 207.56 C-115.87 232.18 -121.58 256.90 -126.65 281.44 C-131.72 305.98 -135.85 330.40 -139.21 354.79 C-142.58 379.18 -145.16 403.48 -146.82 427.79 C-148.49 452.10 -149.36 476.38 -149.19 500.65 C-149.03 524.92 -148.32 549.28 -145.85 573.41 C-143.38 597.53 -139.92 621.82 -134.37 645.42 C-128.81 669.01 -121.53 692.50 -112.54 714.96 C-103.54 737.41 -92.76 759.47 -80.41 780.14 C-68.07 800.81 -53.89 820.71 -38.48 838.99 C-23.08 857.27 -5.85 874.32 12.02 889.82 C29.89 905.32 49.30 919.20 68.73 931.98 C88.16 944.75 108.50 955.93 128.62 966.46 C148.73 976.99 169.11 986.33 189.39 995.16 C209.68 1003.98 229.95 1012.00 250.31 1019.42 C270.67 1026.84 291.01 1033.62 311.54 1039.67 C332.07 1045.73 352.68 1051.15 373.48 1055.73 C394.28 1060.32 417.02 1064.31 436.36 1067.17 Z`

	__FrontCurve__000008_Rotated_0_.Name = `Rotated 0 `
	__FrontCurve__000008_Rotated_0_.Path = `M789.71 532.67 C791.42 522.80 792.53 514.68 793.77 505.53 C795.00 496.38 796.02 487.19 797.13 477.77 C798.23 468.35 799.18 458.88 800.39 449.00 C801.61 439.12 802.80 429.15 804.42 418.48 C806.04 407.81 808.08 396.80 810.12 384.99 C812.15 373.17 815.20 360.41 816.63 347.58 C818.06 334.76 819.48 320.99 818.70 308.04 C817.92 295.09 815.58 282.07 811.95 269.86 C808.31 257.65 802.99 245.87 796.90 234.78 C790.81 223.68 783.40 213.15 775.43 203.28 C767.45 193.42 758.47 184.16 749.04 175.59 C739.62 167.01 730.58 159.62 718.88 151.83 C707.18 144.04 692.72 135.35 678.84 128.84 C664.96 122.33 650.34 116.80 635.59 112.77 C620.84 108.73 605.50 105.89 590.32 104.63 C575.15 103.36 559.60 103.54 544.56 105.18 C529.52 106.82 514.42 110.15 500.07 114.45 C485.73 118.75 471.74 124.75 458.49 130.98 C445.24 137.20 432.67 144.62 420.59 151.79 C408.50 158.96 397.12 166.58 385.98 173.98 C374.85 181.39 364.23 188.80 353.77 196.22 C343.30 203.64 333.17 210.97 323.19 218.50 C313.20 226.04 303.43 233.57 293.85 241.41 C284.26 249.26 274.13 257.99 265.66 265.60 C257.20 273.20 250.52 279.71 243.05 287.05 C235.59 294.39 228.32 301.93 220.88 309.63 C213.44 317.34 206.15 325.21 198.43 333.26 C190.72 341.32 183.02 349.49 174.60 357.96 C166.18 366.43 157.28 374.93 147.94 384.08 C138.59 393.22 127.92 402.50 118.53 412.85 C109.14 423.19 99.22 434.30 91.59 446.14 C83.96 457.99 77.57 470.87 72.75 483.92 C67.93 496.97 64.76 510.75 62.67 524.46 C60.58 538.17 59.93 552.27 60.20 566.18 C60.47 580.09 61.97 594.16 64.29 607.94 C66.61 621.73 69.37 634.22 74.13 648.89 C78.88 663.57 85.36 680.88 92.80 695.97 C100.24 711.05 108.95 725.80 118.76 739.42 C128.57 753.03 139.66 766.04 151.66 777.65 C163.66 789.27 176.93 799.94 190.76 809.09 C204.60 818.25 219.60 826.06 234.66 832.57 C249.72 839.08 265.63 843.98 281.12 848.17 C296.62 852.36 312.39 855.12 327.65 857.69 C342.91 860.27 357.95 861.98 372.67 863.63 C387.40 865.28 401.76 866.52 416.02 867.58 C430.29 868.63 444.26 869.47 458.24 869.97 C472.23 870.48 486.05 870.76 499.93 870.61 C513.81 870.45 527.71 870.33 541.51 869.02 C555.31 867.72 569.28 866.02 582.74 862.78 C596.20 859.55 609.63 855.17 622.27 849.62 C634.91 844.07 647.29 837.30 658.60 829.47 C669.91 821.65 680.64 812.54 690.12 802.67 C699.59 792.80 708.10 781.65 715.47 770.27 C722.84 758.89 728.92 746.48 734.34 734.40 C739.76 722.33 743.96 709.81 748.01 697.84 C752.06 685.86 755.34 674.03 758.63 662.56 C761.93 651.08 764.88 639.99 767.79 629.00 C770.70 618.02 773.47 607.36 776.09 596.64 C778.71 585.93 781.24 575.40 783.51 564.74 C785.78 554.08 788.00 542.54 789.71 532.67 Z`

	__FrontCurve__000009_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000009_Rotated_1_.Path = `M403.03 861.68 C415.28 865.56 425.43 868.35 436.84 871.39 C448.25 874.39 459.74 877.08 471.48 879.79 C483.20 882.46 494.99 884.87 507.18 887.52 C519.35 890.13 531.59 892.63 544.50 895.56 C557.36 898.44 570.47 901.69 584.41 904.95 C598.30 908.18 612.97 912.45 627.93 915.03 C642.83 917.59 658.68 920.16 674.01 920.35 C689.29 920.52 704.87 919.03 719.78 916.07 C734.65 913.09 749.34 908.25 763.39 902.49 C777.41 896.70 791.04 889.42 804.01 881.41 C816.95 873.37 829.40 864.16 841.16 854.34 C852.88 844.50 863.16 834.98 874.47 822.43 C885.74 809.85 898.56 794.18 908.83 778.89 C919.06 763.57 928.32 747.26 935.97 730.56 C943.58 713.82 949.98 696.23 954.65 678.56 C959.26 660.86 962.38 642.48 963.83 624.40 C965.23 606.28 964.83 587.78 963.20 569.93 C961.52 552.02 957.96 534.18 953.93 517.09 C949.84 499.95 944.32 483.30 938.83 467.21 C933.28 451.09 927.03 435.62 920.83 420.47 C914.56 405.29 908.07 390.69 901.41 376.27 C894.69 361.85 887.87 347.84 880.70 333.99 C873.47 320.16 866.08 306.61 858.23 293.26 C850.32 279.95 841.35 265.82 833.43 254.01 C825.48 242.23 818.56 232.86 810.70 222.44 C802.83 212.06 794.64 201.90 786.25 191.62 C777.84 181.38 769.17 171.33 760.31 160.90 C751.44 150.52 742.40 140.22 733.09 129.23 C723.78 118.29 714.46 106.96 704.49 95.16 C694.54 83.41 684.59 70.42 673.35 58.62 C662.14 46.88 650.17 34.58 637.11 24.53 C624.08 14.53 609.75 5.68 595.03 -1.53 C580.33 -8.70 564.59 -14.25 548.78 -18.63 C532.97 -22.96 516.53 -25.82 500.15 -27.65 C483.78 -29.43 467.05 -29.93 450.51 -29.47 C433.97 -28.98 418.88 -27.80 400.87 -24.77 C382.88 -21.70 361.50 -17.07 342.52 -11.09 C323.55 -5.04 304.73 2.46 286.97 11.33 C269.22 20.27 251.90 30.71 235.95 42.35 C220.01 54.06 204.86 67.31 191.25 81.42 C177.67 95.60 165.29 111.27 154.33 127.26 C143.39 143.31 134.01 160.56 125.53 177.57 C117.07 194.65 110.10 212.30 103.51 229.53 C96.96 246.83 91.42 264.10 86.13 281.13 C80.88 298.22 76.22 315.05 71.89 331.86 C67.61 348.71 63.73 365.34 60.32 382.08 C56.96 398.83 53.98 415.50 51.57 432.31 C49.23 449.12 46.98 465.98 46.07 482.91 C45.21 499.85 44.87 517.03 46.27 533.90 C47.72 550.77 50.45 567.81 54.60 584.18 C58.81 600.54 64.41 616.83 71.34 632.13 C78.34 647.42 86.84 662.33 96.38 676.00 C105.99 689.68 117.16 702.55 128.82 714.21 C140.55 725.87 153.64 736.34 166.57 745.99 C179.57 755.64 193.31 764.05 206.62 772.12 C220.00 780.18 233.45 787.33 246.62 794.34 C259.85 801.32 272.84 807.78 285.81 814.06 C298.82 820.30 311.60 826.23 324.53 831.89 C337.49 837.50 350.34 842.86 363.42 847.85 C376.52 852.79 390.80 857.74 403.03 861.68 Z`

	__FrontCurve__000010_Rotated_1_.Name = `Rotated 1 `
	__FrontCurve__000010_Rotated_1_.Path = `M368.62 1007.73 C385.88 1012.71 400.17 1016.26 416.23 1020.02 C432.28 1023.58 448.41 1026.70 464.86 1029.69 C481.26 1032.50 497.71 1034.90 514.65 1037.39 C531.49 1039.69 548.36 1041.75 566.01 1044.06 C583.50 1046.20 601.18 1048.57 619.83 1050.77 C638.25 1052.83 657.45 1055.77 676.98 1056.85 C696.25 1057.82 716.56 1058.62 736.25 1056.91 C755.71 1055.09 775.58 1051.41 794.62 1046.16 C813.47 1040.80 832.19 1033.39 850.12 1024.97 C867.89 1016.44 885.27 1006.23 901.84 995.22 C918.27 984.10 934.16 971.63 949.21 958.50 C964.12 945.25 977.21 932.53 991.79 916.03 C1006.19 899.38 1022.57 878.71 1035.90 858.76 C1049.01 838.65 1061.02 817.36 1071.18 795.71 C1081.11 773.90 1089.69 751.05 1096.28 728.20 C1102.65 705.19 1107.34 681.36 1110.12 657.94 C1112.66 634.35 1113.21 610.27 1112.31 586.98 C1111.16 563.49 1107.92 540.05 1104.00 517.46 C1099.81 494.70 1093.99 472.48 1088.01 450.90 C1081.74 429.18 1074.58 408.24 1067.26 387.66 C1059.65 367.00 1051.65 347.06 1043.27 327.34 C1034.61 307.62 1025.68 288.46 1016.20 269.52 C1006.46 250.65 996.40 232.21 985.67 214.04 C974.71 196.00 962.23 176.92 951.20 161.01 C939.99 145.25 930.30 132.71 919.26 118.84 C908.10 105.12 896.54 91.74 884.63 78.31 C872.64 65.06 860.32 52.12 847.64 38.87 C834.92 25.82 821.97 12.99 808.56 -0.47 C795.17 -13.71 781.72 -27.20 767.41 -41.09 C753.18 -54.73 738.85 -69.50 723.11 -82.97 C707.48 -96.19 690.91 -109.87 673.26 -121.13 C655.71 -132.18 636.67 -141.97 617.29 -149.96 C597.98 -157.74 577.49 -163.78 556.99 -168.49 C536.54 -173.00 515.34 -175.90 494.28 -177.62 C473.26 -179.16 451.82 -179.25 430.62 -178.24 C409.48 -177.06 390.19 -175.12 367.17 -171.03 C344.20 -166.72 316.93 -160.39 292.64 -152.65 C268.44 -144.65 244.38 -134.97 221.57 -123.76 C198.84 -112.30 176.56 -99.10 155.87 -84.59 C135.27 -69.81 115.49 -53.28 97.52 -35.80 C79.63 -18.05 63.05 1.39 48.14 21.20 C33.32 41.29 20.20 62.73 8.21 83.97 C-3.67 105.51 -13.87 127.71 -23.48 149.56 C-32.94 171.68 -41.17 193.81 -48.97 215.79 C-56.58 238.00 -63.38 259.96 -69.65 281.99 C-75.70 304.18 -81.13 326.17 -85.91 348.33 C-90.43 370.59 -94.36 392.76 -97.51 415.13 C-100.39 437.53 -102.97 460.01 -104.00 482.54 C-104.76 505.07 -104.80 527.92 -102.90 550.34 C-100.73 572.75 -97.06 595.40 -91.79 617.20 C-86.25 638.99 -79.08 660.75 -70.44 681.30 C-61.52 701.84 -50.86 722.00 -39.04 740.67 C-26.92 759.35 -13.03 777.18 1.45 793.54 C16.23 809.91 32.55 824.97 48.80 838.99 C65.38 852.99 82.81 865.58 99.92 877.63 C117.35 889.63 134.91 900.50 152.33 911.07 C170.03 921.52 187.51 931.21 205.12 940.57 C222.95 949.76 240.58 958.40 258.48 966.62 C276.54 974.61 294.53 982.14 312.86 989.13 C331.29 995.86 351.38 1002.52 368.62 1007.73 Z`

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

	__NoteInfo__000000_0.Name = `0`
	__NoteInfo__000000_0.IsKept = true

	__NoteInfo__000001_1.Name = `1`
	__NoteInfo__000001_1.IsKept = false

	__NoteInfo__000002_10.Name = `10`
	__NoteInfo__000002_10.IsKept = false

	__NoteInfo__000003_11.Name = `11`
	__NoteInfo__000003_11.IsKept = false

	__NoteInfo__000004_12.Name = `12`
	__NoteInfo__000004_12.IsKept = false

	__NoteInfo__000005_13.Name = `13`
	__NoteInfo__000005_13.IsKept = false

	__NoteInfo__000006_14.Name = `14`
	__NoteInfo__000006_14.IsKept = false

	__NoteInfo__000007_15.Name = `15`
	__NoteInfo__000007_15.IsKept = false

	__NoteInfo__000008_16.Name = `16`
	__NoteInfo__000008_16.IsKept = true

	__NoteInfo__000009_17.Name = `17`
	__NoteInfo__000009_17.IsKept = false

	__NoteInfo__000010_18.Name = `18`
	__NoteInfo__000010_18.IsKept = false

	__NoteInfo__000011_19.Name = `19`
	__NoteInfo__000011_19.IsKept = false

	__NoteInfo__000012_2.Name = `2`
	__NoteInfo__000012_2.IsKept = false

	__NoteInfo__000013_20.Name = `20`
	__NoteInfo__000013_20.IsKept = false

	__NoteInfo__000014_21.Name = `21`
	__NoteInfo__000014_21.IsKept = false

	__NoteInfo__000015_22.Name = `22`
	__NoteInfo__000015_22.IsKept = false

	__NoteInfo__000016_23.Name = `23`
	__NoteInfo__000016_23.IsKept = false

	__NoteInfo__000017_3.Name = `3`
	__NoteInfo__000017_3.IsKept = false

	__NoteInfo__000018_4.Name = `4`
	__NoteInfo__000018_4.IsKept = false

	__NoteInfo__000019_5.Name = `5`
	__NoteInfo__000019_5.IsKept = false

	__NoteInfo__000020_6.Name = `6`
	__NoteInfo__000020_6.IsKept = false

	__NoteInfo__000021_7.Name = `7`
	__NoteInfo__000021_7.IsKept = false

	__NoteInfo__000022_8.Name = `8`
	__NoteInfo__000022_8.IsKept = true

	__NoteInfo__000023_9.Name = `9`
	__NoteInfo__000023_9.IsKept = false

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.BackendColor = `#F1F1F1`
	__Parameter__000000_Reference.MinuteColor = `#5A5A5A`
	__Parameter__000000_Reference.HourColor = `#1E1E1E`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.InsideAngle = 112.000000
	__Parameter__000000_Reference.SideLength = 180.000000
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
	__Parameter__000000_Reference.NbMeasureLinesPerCurve = 22
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
	__Parameter__000000_Reference.ActualNotesTemporalShift = 8

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 180.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 180.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 295.774340
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 39.900462
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 180.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 180.000000
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
	__ShapeCategory__000006_6_Spiral_growth.IsExpanded = true

	__ShapeCategory__000007_7_Spiral_Growth_Bezier.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = false

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = true

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = true

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 288.720947
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -32.560546
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 332.214299
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 40.861961
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 218.344631
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 347.322287
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 259.394191
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 267.509306
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
	__SpiralCircle__000000_Backend_Marker.CenterY = 475.200000
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
	__SpiralCircle__000001_Construction_Circle_Spiral.CenterX = 190.800000
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
	__SpiralCircle__000002_Hour_Marker.CenterY = 287.100000
	__SpiralCircle__000002_Hour_Marker.HasBespokeRadius = true
	__SpiralCircle__000002_Hour_Marker.BespopkeRadius = 52.200000
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
	__SpiralCircle__000003_Minute_Marker.CenterY = 473.400000
	__SpiralCircle__000003_Minute_Marker.HasBespokeRadius = true
	__SpiralCircle__000003_Minute_Marker.BespopkeRadius = 28.800000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 288.720947
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 190.800000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -32.560546
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 288.720947
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 380.499080
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -32.560546
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -86.927236
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 90.475319
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 10.203367
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 69.610404
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 198.922225
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 288.720947
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -32.560546
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 56.431356
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -161.261110
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
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000000_0)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000001_1)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000012_2)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000017_3)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000018_4)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000019_5)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000020_6)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000021_7)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000022_8)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000023_9)
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
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000015_22)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000016_23)
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
