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

	__FrontCurveStack__000000_Front_Curve_Stack := (&models.FrontCurveStack{}).Stage(stage)

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
	__NoteInfo__000015_3 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000016_4 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000017_5 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000018_6 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000019_7 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000020_8 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000021_9 := (&models.NoteInfo{}).Stage(stage)

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

	__Axis__000000_Construction_Axis.Name = `Construction Axis`
	__Axis__000000_Construction_Axis.IsDisplayed = false
	__Axis__000000_Construction_Axis.AngleDegree = 97.682917
	__Axis__000000_Construction_Axis.Length = 211.374918
	__Axis__000000_Construction_Axis.CenterX = 0.000000
	__Axis__000000_Construction_Axis.CenterY = 0.000000
	__Axis__000000_Construction_Axis.EndX = -28.258862
	__Axis__000000_Construction_Axis.EndY = 209.477427
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
	__Axis__000001_Initial_Axis.Length = 790.537074
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
	__Axis__000004_Rotated_Axis.Length = 790.537074
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -14.129431
	__Bezier__000005_Growth_Curve_Seed.StartY = 104.738714
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 68.284051
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 115.856433
	__Bezier__000005_Growth_Curve_Seed.EndX = 127.022667
	__Bezier__000005_Growth_Curve_Seed.EndY = 230.425170
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 44.609185
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 219.307451
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
	__Circle__000001_Construction_Circle.CenterX = -14.129431
	__Circle__000001_Construction_Circle.CenterY = 104.738714
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 310.563057
	__Circle__000005_Rotated_Next_Circle.CenterY = 41.895485
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M803.99 534.74 C805.79 524.41 806.95 515.90 808.26 506.31 C809.56 496.73 810.63 487.11 811.80 477.24 C812.96 467.38 813.97 457.46 815.25 447.11 C816.54 436.76 817.81 426.31 819.52 415.12 C821.24 403.94 823.39 392.41 825.55 380.02 C827.70 367.63 830.92 354.25 832.44 340.80 C833.96 327.35 835.47 312.91 834.68 299.32 C833.88 285.73 831.45 272.08 827.65 259.28 C823.86 246.47 818.29 234.12 811.93 222.48 C805.56 210.84 797.80 199.79 789.45 189.45 C781.11 179.10 771.70 169.40 761.83 160.40 C751.96 151.41 742.50 143.66 730.25 135.49 C717.99 127.32 702.85 118.21 688.31 111.39 C673.77 104.56 658.46 98.77 643.01 94.54 C627.56 90.32 611.49 87.34 595.60 86.03 C579.71 84.71 563.42 84.91 547.66 86.64 C531.91 88.37 516.08 91.88 501.06 96.41 C486.03 100.94 471.38 107.26 457.50 113.81 C443.62 120.36 430.46 128.17 417.80 135.71 C405.14 143.25 393.22 151.27 381.55 159.06 C369.88 166.85 358.76 174.64 347.80 182.44 C336.83 190.24 326.21 197.95 315.74 205.86 C305.27 213.78 295.03 221.69 284.97 229.93 C274.91 238.18 264.28 247.34 255.39 255.32 C246.51 263.31 239.49 270.14 231.66 277.84 C223.82 285.54 216.17 293.45 208.36 301.53 C200.54 309.61 192.88 317.86 184.77 326.30 C176.66 334.75 168.56 343.31 159.71 352.19 C150.85 361.06 141.49 369.97 131.66 379.55 C121.83 389.14 110.60 398.86 100.71 409.70 C90.83 420.53 80.40 432.17 72.37 444.58 C64.34 456.98 57.62 470.47 52.54 484.15 C47.47 497.82 44.13 512.26 41.93 526.62 C39.73 540.98 39.04 555.75 39.32 570.33 C39.60 584.91 41.16 599.64 43.59 614.08 C46.02 628.53 48.92 641.62 53.90 656.99 C58.89 672.36 65.69 690.50 73.49 706.31 C81.29 722.11 90.43 737.56 100.72 751.82 C111.01 766.08 122.64 779.71 135.23 791.87 C147.82 804.03 161.74 815.20 176.25 824.78 C190.76 834.36 206.50 842.54 222.30 849.35 C238.10 856.16 254.78 861.27 271.04 865.64 C287.29 870.01 303.83 872.89 319.83 875.57 C335.83 878.26 351.60 880.03 367.04 881.75 C382.48 883.46 397.53 884.74 412.48 885.84 C427.43 886.93 442.07 887.80 456.73 888.32 C471.38 888.84 485.87 889.13 500.42 888.96 C514.96 888.79 529.53 888.66 543.99 887.28 C558.45 885.91 573.09 884.13 587.19 880.73 C601.30 877.34 615.38 872.74 628.62 866.92 C641.87 861.09 654.85 853.99 666.69 845.78 C678.54 837.58 689.78 828.02 699.71 817.67 C709.64 807.33 718.55 795.63 726.27 783.70 C733.99 771.77 740.34 758.75 746.02 746.09 C751.70 733.43 756.09 720.31 760.32 707.75 C764.56 695.20 767.98 682.80 771.43 670.78 C774.88 658.76 777.97 647.14 781.01 635.63 C784.06 624.12 786.96 612.95 789.70 601.74 C792.45 590.52 795.09 579.49 797.47 568.32 C799.85 557.16 802.19 545.08 803.99 534.74 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M255.39 255.32 C246.51 263.31 239.49 270.14 231.66 277.84 C223.82 285.54 216.17 293.45 208.36 301.53 C200.54 309.61 192.88 317.86 184.77 326.30 C176.66 334.75 168.56 343.31 159.71 352.19 C150.85 361.06 141.49 369.97 131.66 379.55 C121.83 389.14 110.60 398.86 100.71 409.70 C90.83 420.53 80.40 432.17 72.37 444.58 C64.34 456.98 57.62 470.47 52.54 484.15 C47.47 497.82 44.13 512.26 41.93 526.62 C39.73 540.98 39.04 555.75 39.32 570.33 C39.60 584.91 41.16 599.64 43.59 614.08 C46.02 628.53 48.92 641.62 53.90 656.99 C58.89 672.36 65.69 690.50 73.49 706.31 C81.29 722.11 90.43 737.56 100.72 751.82 C111.01 766.08 122.64 779.71 135.23 791.87 C147.82 804.03 161.74 815.20 176.25 824.78 C190.76 834.36 206.50 842.54 222.30 849.35 C238.10 856.16 254.78 861.27 271.04 865.64 C287.29 870.01 303.83 872.89 319.83 875.57 C335.83 878.26 351.60 880.03 367.04 881.75 C382.48 883.46 397.53 884.74 412.48 885.84 C427.43 886.93 442.07 887.80 456.73 888.32 C471.38 888.84 485.87 889.13 500.42 888.96 C514.96 888.79 530.66 888.14 543.99 887.28 C557.32 886.43 568.20 885.26 580.39 883.86 C592.57 882.45 604.73 880.70 617.10 878.86 C629.47 877.03 641.83 874.90 654.61 872.85 C667.40 870.79 680.24 868.58 693.83 866.54 C707.41 864.49 721.40 862.69 736.13 860.60 C750.86 858.50 766.74 857.11 782.23 853.97 C797.71 850.83 814.09 847.35 829.03 841.73 C843.98 836.11 858.54 828.79 871.88 820.26 C885.22 811.73 897.64 801.48 909.08 790.54 C920.51 779.61 930.97 767.37 940.50 754.67 C950.03 741.96 958.61 728.29 966.27 714.29 C973.93 700.28 980.29 687.15 986.46 670.66 C992.62 654.17 999.12 634.10 1003.24 615.35 C1007.36 596.59 1010.13 577.26 1011.17 558.13 C1012.21 539.00 1011.71 519.51 1009.47 500.59 C1007.23 481.66 1003.25 462.67 997.72 444.57 C992.19 426.47 984.74 408.71 976.30 392.01 C967.87 375.30 957.60 359.38 947.12 344.35 C936.64 329.32 924.92 315.30 913.42 301.81 C901.92 288.32 889.96 275.72 878.14 263.42 C866.32 251.13 854.45 239.44 842.49 228.04 C830.53 216.63 818.59 205.65 806.36 194.99 C794.14 184.33 781.83 174.01 769.14 164.10 C756.46 154.18 743.72 144.28 730.25 135.49 C716.77 126.71 702.85 118.21 688.31 111.39 C673.77 104.56 658.46 98.77 643.01 94.54 C627.56 90.32 611.49 87.34 595.60 86.03 C579.71 84.71 563.42 84.91 547.66 86.64 C531.91 88.37 516.08 91.88 501.06 96.41 C486.03 100.94 471.38 107.26 457.50 113.81 C443.62 120.36 430.46 128.17 417.80 135.71 C405.14 143.25 393.22 151.27 381.55 159.06 C369.88 166.85 358.76 174.64 347.80 182.44 C336.83 190.24 326.21 197.95 315.74 205.86 C305.27 213.78 295.03 221.69 284.97 229.93 C274.91 238.18 264.28 247.34 255.39 255.32 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M543.99 887.28 C557.32 886.43 568.20 885.26 580.39 883.86 C592.57 882.45 604.73 880.70 617.10 878.86 C629.47 877.03 641.83 874.90 654.61 872.85 C667.40 870.79 680.24 868.58 693.83 866.54 C707.41 864.49 721.40 862.69 736.13 860.60 C750.86 858.50 766.74 857.11 782.23 853.97 C797.71 850.83 814.09 847.35 829.03 841.73 C843.98 836.11 858.54 828.79 871.88 820.26 C885.22 811.73 897.64 801.48 909.08 790.54 C920.51 779.61 930.97 767.37 940.50 754.67 C950.03 741.96 958.61 728.29 966.27 714.29 C973.93 700.28 980.29 687.15 986.46 670.66 C992.62 654.17 999.12 634.10 1003.24 615.35 C1007.36 596.59 1010.13 577.26 1011.17 558.13 C1012.21 539.00 1011.71 519.51 1009.47 500.59 C1007.23 481.66 1003.25 462.67 997.72 444.57 C992.19 426.47 984.74 408.71 976.30 392.01 C967.87 375.30 957.60 359.38 947.12 344.35 C936.64 329.32 924.92 315.30 913.42 301.81 C901.92 288.32 889.96 275.72 878.14 263.42 C866.32 251.13 854.45 239.44 842.49 228.04 C830.53 216.63 818.59 205.65 806.36 194.99 C794.14 184.33 781.83 174.01 769.14 164.10 C756.46 154.18 742.39 143.92 730.25 135.49 C718.10 127.07 707.83 120.63 696.28 113.53 C684.72 106.44 672.94 99.69 660.92 92.93 C648.90 86.17 636.69 79.70 624.17 72.96 C611.65 66.22 598.97 59.64 585.80 52.51 C572.64 45.38 559.30 37.88 545.19 30.20 C531.08 22.51 516.50 13.62 501.14 6.40 C485.78 -0.81 469.51 -8.26 453.04 -13.08 C436.57 -17.89 419.33 -21.08 402.31 -22.49 C385.29 -23.90 367.93 -23.33 350.93 -21.56 C333.94 -19.78 316.90 -16.31 300.33 -11.83 C283.75 -7.35 267.34 -1.46 251.47 5.32 C235.60 12.11 221.40 19.00 205.10 28.88 C188.80 38.76 169.80 51.45 153.64 64.61 C137.46 77.77 122.01 92.32 108.09 107.83 C94.15 123.32 81.25 140.10 70.14 157.56 C59.04 175.01 49.34 193.64 41.50 212.57 C33.67 231.49 27.64 251.40 23.12 271.12 C18.61 290.84 16.11 311.15 14.41 330.90 C12.72 350.66 12.71 370.44 12.95 389.67 C13.18 408.91 14.42 427.77 15.83 446.30 C17.24 464.83 19.15 482.93 21.40 500.86 C23.65 518.80 26.24 536.40 29.34 553.90 C32.44 571.41 35.90 588.70 39.99 605.88 C44.09 623.06 48.32 640.25 53.90 656.99 C59.49 673.73 65.69 690.50 73.49 706.31 C81.29 722.11 90.43 737.56 100.72 751.82 C111.01 766.08 122.64 779.71 135.23 791.87 C147.82 804.03 161.74 815.20 176.25 824.78 C190.76 834.36 206.50 842.54 222.30 849.35 C238.10 856.16 254.78 861.27 271.04 865.64 C287.29 870.01 303.83 872.89 319.83 875.57 C335.83 878.26 351.60 880.03 367.04 881.75 C382.48 883.46 397.53 884.74 412.48 885.84 C427.43 886.93 442.07 887.80 456.73 888.32 C471.38 888.84 485.87 889.13 500.42 888.96 C514.96 888.79 530.66 888.14 543.99 887.28 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M730.25 135.49 C718.10 127.07 707.83 120.63 696.28 113.53 C684.72 106.44 672.94 99.69 660.92 92.93 C648.90 86.17 636.69 79.70 624.17 72.96 C611.65 66.22 598.97 59.64 585.80 52.51 C572.64 45.38 559.30 37.88 545.19 30.20 C531.08 22.51 516.50 13.62 501.14 6.40 C485.78 -0.81 469.51 -8.26 453.04 -13.08 C436.57 -17.89 419.33 -21.08 402.31 -22.49 C385.29 -23.90 367.93 -23.33 350.93 -21.56 C333.94 -19.78 316.90 -16.31 300.33 -11.83 C283.75 -7.35 267.34 -1.46 251.47 5.32 C235.60 12.11 221.40 19.00 205.10 28.88 C188.80 38.76 169.80 51.45 153.64 64.61 C137.46 77.77 122.01 92.32 108.09 107.83 C94.15 123.32 81.25 140.10 70.14 157.56 C59.04 175.01 49.34 193.64 41.50 212.57 C33.67 231.49 27.64 251.40 23.12 271.12 C18.61 290.84 16.11 311.15 14.41 330.90 C12.72 350.66 12.71 370.44 12.95 389.67 C13.18 408.91 14.42 427.77 15.83 446.30 C17.24 464.83 19.15 482.93 21.40 500.86 C23.65 518.80 26.24 536.40 29.34 553.90 C32.44 571.41 35.90 588.70 39.99 605.88 C44.09 623.06 49.19 641.48 53.90 656.99 C58.62 672.50 63.10 685.02 68.26 698.95 C73.43 712.88 79.04 726.67 84.88 740.57 C90.73 754.47 96.96 768.24 103.31 782.37 C109.67 796.50 116.29 810.59 123.00 825.35 C129.71 840.11 136.38 855.24 143.57 870.94 C150.77 886.64 157.62 903.53 166.16 919.55 C174.71 935.57 183.86 952.29 194.83 967.03 C205.81 981.77 218.45 995.66 232.04 1007.98 C245.63 1020.30 260.76 1031.23 276.37 1040.94 C291.98 1050.65 308.70 1059.00 325.70 1066.23 C342.71 1073.46 360.48 1079.42 378.38 1084.30 C396.29 1089.17 412.85 1092.87 433.15 1095.49 C453.45 1098.11 477.86 1100.21 500.19 1100.04 C522.52 1099.87 545.13 1098.09 567.12 1094.49 C589.11 1090.88 611.11 1085.53 632.11 1078.43 C653.11 1071.32 673.78 1062.33 693.15 1051.87 C712.51 1041.41 731.09 1028.98 748.32 1015.67 C765.55 1002.36 781.55 987.27 796.52 972.04 C811.50 956.80 825.13 940.42 838.17 924.25 C851.22 908.08 863.20 891.50 874.77 875.02 C886.34 858.53 897.20 842.01 907.61 825.35 C918.02 808.70 927.90 792.03 937.23 775.08 C946.57 758.13 955.42 741.06 963.62 723.66 C971.83 706.26 979.85 688.71 986.46 670.66 C993.06 652.60 999.12 634.10 1003.24 615.35 C1007.36 596.59 1010.13 577.26 1011.17 558.13 C1012.21 539.00 1011.71 519.51 1009.47 500.59 C1007.23 481.66 1003.25 462.67 997.72 444.57 C992.19 426.47 984.74 408.71 976.30 392.01 C967.87 375.30 957.60 359.38 947.12 344.35 C936.64 329.32 924.92 315.30 913.42 301.81 C901.92 288.32 889.96 275.72 878.14 263.42 C866.32 251.13 854.45 239.44 842.49 228.04 C830.53 216.63 818.59 205.65 806.36 194.99 C794.14 184.33 781.83 174.01 769.14 164.10 C756.46 154.18 742.39 143.92 730.25 135.49 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M53.90 656.99 C58.62 672.50 63.10 685.02 68.26 698.95 C73.43 712.88 79.04 726.67 84.88 740.57 C90.73 754.47 96.96 768.24 103.31 782.37 C109.67 796.50 116.29 810.59 123.00 825.35 C129.71 840.11 136.38 855.24 143.57 870.94 C150.77 886.64 157.62 903.53 166.16 919.55 C174.71 935.57 183.86 952.29 194.83 967.03 C205.81 981.77 218.45 995.66 232.04 1007.98 C245.63 1020.30 260.76 1031.23 276.37 1040.94 C291.98 1050.65 308.70 1059.00 325.70 1066.23 C342.71 1073.46 360.48 1079.42 378.38 1084.30 C396.29 1089.17 412.85 1092.87 433.15 1095.49 C453.45 1098.11 477.86 1100.21 500.19 1100.04 C522.52 1099.87 545.13 1098.09 567.12 1094.49 C589.11 1090.88 611.11 1085.53 632.11 1078.43 C653.11 1071.32 673.78 1062.33 693.15 1051.87 C712.51 1041.41 731.09 1028.98 748.32 1015.67 C765.55 1002.36 781.55 987.27 796.52 972.04 C811.50 956.80 825.13 940.42 838.17 924.25 C851.22 908.08 863.20 891.50 874.77 875.02 C886.34 858.53 897.20 842.01 907.61 825.35 C918.02 808.70 927.90 792.03 937.23 775.08 C946.57 758.13 955.42 741.06 963.62 723.66 C971.83 706.26 979.97 687.08 986.46 670.66 C992.94 654.23 997.58 640.50 1002.55 625.12 C1007.52 609.74 1012.00 594.16 1016.30 578.35 C1020.61 562.54 1024.47 546.58 1028.40 530.26 C1032.33 513.94 1035.95 497.48 1039.87 480.44 C1043.79 463.40 1047.91 446.10 1051.92 428.04 C1055.93 409.97 1060.89 391.18 1063.91 372.05 C1066.94 352.92 1069.88 332.86 1070.09 313.24 C1070.30 293.62 1068.63 273.64 1065.18 254.33 C1061.73 235.02 1056.15 215.85 1049.38 197.38 C1042.61 178.90 1034.08 160.82 1024.57 143.48 C1015.06 126.14 1004.11 109.34 992.34 93.34 C980.56 77.35 969.15 63.25 953.91 47.52 C938.67 31.80 919.65 13.77 900.88 -1.01 C882.12 -15.79 862.05 -29.42 841.31 -41.16 C820.57 -52.89 798.67 -63.20 776.42 -71.43 C754.18 -79.67 730.98 -86.13 707.86 -90.56 C684.73 -94.98 660.96 -97.29 637.68 -97.98 C614.41 -98.67 590.95 -97.12 568.21 -94.72 C545.46 -92.33 523.03 -88.11 501.21 -83.60 C479.38 -79.09 458.12 -73.51 437.26 -67.66 C416.40 -61.81 396.07 -55.39 376.04 -48.50 C356.01 -41.62 336.39 -34.29 317.06 -26.33 C297.74 -18.37 278.74 -9.92 260.08 -0.73 C241.42 8.48 222.84 17.99 205.10 28.88 C187.36 39.77 169.80 51.45 153.64 64.61 C137.46 77.77 122.01 92.32 108.09 107.83 C94.15 123.32 81.25 140.10 70.14 157.56 C59.04 175.01 49.34 193.64 41.50 212.57 C33.67 231.49 27.64 251.40 23.12 271.12 C18.61 290.84 16.11 311.15 14.41 330.90 C12.72 350.66 12.71 370.44 12.95 389.67 C13.18 408.91 14.42 427.77 15.83 446.30 C17.24 464.83 19.15 482.93 21.40 500.86 C23.65 518.80 26.24 536.40 29.34 553.90 C32.44 571.41 35.90 588.70 39.99 605.88 C44.09 623.06 49.19 641.48 53.90 656.99 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M986.46 670.66 C992.94 654.23 997.58 640.50 1002.55 625.12 C1007.52 609.74 1012.00 594.16 1016.30 578.35 C1020.61 562.54 1024.47 546.58 1028.40 530.26 C1032.33 513.94 1035.95 497.48 1039.87 480.44 C1043.79 463.40 1047.91 446.10 1051.92 428.04 C1055.93 409.97 1060.89 391.18 1063.91 372.05 C1066.94 352.92 1069.88 332.86 1070.09 313.24 C1070.30 293.62 1068.63 273.64 1065.18 254.33 C1061.73 235.02 1056.15 215.85 1049.38 197.38 C1042.61 178.90 1034.08 160.82 1024.57 143.48 C1015.06 126.14 1004.11 109.34 992.34 93.34 C980.56 77.35 969.15 63.25 953.91 47.52 C938.67 31.80 919.65 13.77 900.88 -1.01 C882.12 -15.79 862.05 -29.42 841.31 -41.16 C820.57 -52.89 798.67 -63.20 776.42 -71.43 C754.18 -79.67 730.98 -86.13 707.86 -90.56 C684.73 -94.98 660.96 -97.29 637.68 -97.98 C614.41 -98.67 590.95 -97.12 568.21 -94.72 C545.46 -92.33 523.03 -88.11 501.21 -83.60 C479.38 -79.09 458.12 -73.51 437.26 -67.66 C416.40 -61.81 396.07 -55.39 376.04 -48.50 C356.01 -41.62 336.39 -34.29 317.06 -26.33 C297.74 -18.37 278.74 -9.92 260.08 -0.73 C241.42 8.48 221.63 19.28 205.10 28.88 C188.57 38.48 175.43 47.07 160.90 56.86 C146.37 66.65 132.13 76.97 117.89 87.62 C103.65 98.27 89.68 109.36 75.47 120.71 C61.26 132.04 47.24 143.71 32.68 155.65 C18.13 167.59 3.32 179.63 -11.82 192.36 C-26.94 205.09 -43.20 217.82 -58.10 232.04 C-72.99 246.27 -88.33 261.30 -101.19 277.72 C-114.05 294.13 -125.59 312.02 -135.27 330.52 C-144.94 349.01 -152.80 368.75 -159.23 388.69 C-165.66 408.63 -170.40 429.38 -173.86 450.16 C-177.32 470.94 -179.25 492.20 -179.97 513.35 C-180.69 534.49 -180.49 553.80 -178.18 577.02 C-175.87 600.23 -171.95 627.86 -166.12 652.63 C-160.29 677.40 -152.64 702.07 -143.19 725.65 C-133.75 749.23 -122.42 772.39 -109.46 794.09 C-96.50 815.80 -81.61 836.69 -65.43 855.89 C-49.26 875.08 -31.17 892.99 -12.41 909.26 C6.35 925.53 26.73 940.11 47.14 953.53 C67.54 966.94 88.90 978.68 110.02 989.74 C131.13 1000.79 152.54 1010.60 173.84 1019.87 C195.13 1029.14 216.42 1037.56 237.79 1045.35 C259.17 1053.14 280.54 1060.26 302.09 1066.61 C323.64 1072.97 345.28 1078.66 367.12 1083.48 C388.97 1088.29 410.97 1092.73 433.15 1095.49 C455.32 1098.25 477.86 1100.21 500.19 1100.04 C522.52 1099.87 545.13 1098.09 567.12 1094.49 C589.11 1090.88 611.11 1085.53 632.11 1078.43 C653.11 1071.32 673.78 1062.33 693.15 1051.87 C712.51 1041.41 731.09 1028.98 748.32 1015.67 C765.55 1002.36 781.55 987.27 796.52 972.04 C811.50 956.80 825.13 940.42 838.17 924.25 C851.22 908.08 863.20 891.50 874.77 875.02 C886.34 858.53 897.20 842.01 907.61 825.35 C918.02 808.70 927.90 792.03 937.23 775.08 C946.57 758.13 955.42 741.06 963.62 723.66 C971.83 706.26 979.97 687.08 986.46 670.66 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M205.10 28.88 C188.57 38.48 175.43 47.07 160.90 56.86 C146.37 66.65 132.13 76.97 117.89 87.62 C103.65 98.27 89.68 109.36 75.47 120.71 C61.26 132.04 47.24 143.71 32.68 155.65 C18.13 167.59 3.32 179.63 -11.82 192.36 C-26.94 205.09 -43.20 217.82 -58.10 232.04 C-72.99 246.27 -88.33 261.30 -101.19 277.72 C-114.05 294.13 -125.59 312.02 -135.27 330.52 C-144.94 349.01 -152.80 368.75 -159.23 388.69 C-165.66 408.63 -170.40 429.38 -173.86 450.16 C-177.32 470.94 -179.25 492.20 -179.97 513.35 C-180.69 534.49 -180.49 553.80 -178.18 577.02 C-175.87 600.23 -171.95 627.86 -166.12 652.63 C-160.29 677.40 -152.64 702.07 -143.19 725.65 C-133.75 749.23 -122.42 772.39 -109.46 794.09 C-96.50 815.80 -81.61 836.69 -65.43 855.89 C-49.26 875.08 -31.17 892.99 -12.41 909.26 C6.35 925.53 26.73 940.11 47.14 953.53 C67.54 966.94 88.90 978.68 110.02 989.74 C131.13 1000.79 152.54 1010.60 173.84 1019.87 C195.13 1029.14 216.42 1037.56 237.79 1045.35 C259.17 1053.14 280.54 1060.26 302.09 1066.61 C323.64 1072.97 345.28 1078.66 367.12 1083.48 C388.97 1088.29 412.84 1092.48 433.15 1095.49 C453.45 1098.50 470.22 1100.07 488.94 1101.54 C507.65 1103.00 526.44 1103.82 545.42 1104.30 C564.39 1104.78 583.41 1104.66 602.77 1104.41 C622.12 1104.17 641.54 1103.47 661.54 1102.82 C681.53 1102.17 701.81 1101.53 722.76 1100.50 C743.70 1099.48 765.53 1099.10 787.20 1096.65 C808.87 1094.21 831.38 1091.35 852.78 1085.84 C874.19 1080.32 895.47 1072.79 915.62 1063.57 C935.78 1054.35 955.25 1042.99 973.70 1030.53 C992.15 1018.08 1009.75 1003.89 1026.31 988.84 C1042.87 973.79 1058.50 957.37 1073.06 940.25 C1087.61 923.14 1100.20 906.95 1113.64 886.14 C1127.08 865.33 1142.11 839.80 1153.70 815.39 C1165.29 790.98 1175.31 765.46 1183.17 739.66 C1191.03 713.86 1197.08 687.18 1200.85 660.60 C1204.62 634.03 1206.29 606.85 1205.81 580.20 C1205.33 553.55 1202.47 526.70 1197.97 500.71 C1193.47 474.73 1186.57 449.04 1178.83 424.28 C1171.08 399.51 1161.46 375.46 1151.52 352.12 C1141.58 328.78 1130.55 306.24 1119.18 284.24 C1107.82 262.23 1095.84 240.91 1083.33 220.07 C1070.82 199.24 1057.81 178.96 1044.11 159.22 C1030.41 139.47 1016.15 120.24 1001.12 101.62 C986.08 83.00 970.62 64.63 953.91 47.52 C937.20 30.42 919.65 13.77 900.88 -1.01 C882.12 -15.79 862.05 -29.42 841.31 -41.16 C820.57 -52.89 798.67 -63.20 776.42 -71.43 C754.18 -79.67 730.98 -86.13 707.86 -90.56 C684.73 -94.98 660.96 -97.29 637.68 -97.98 C614.41 -98.67 590.95 -97.12 568.21 -94.72 C545.46 -92.33 523.03 -88.11 501.21 -83.60 C479.38 -79.09 458.12 -73.51 437.26 -67.66 C416.40 -61.81 396.07 -55.39 376.04 -48.50 C356.01 -41.62 336.39 -34.29 317.06 -26.33 C297.74 -18.37 278.74 -9.92 260.08 -0.73 C241.42 8.48 221.63 19.28 205.10 28.88 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M433.15 1095.49 C453.45 1098.50 470.22 1100.07 488.94 1101.54 C507.65 1103.00 526.44 1103.82 545.42 1104.30 C564.39 1104.78 583.41 1104.66 602.77 1104.41 C622.12 1104.17 641.54 1103.47 661.54 1102.82 C681.53 1102.17 701.81 1101.53 722.76 1100.50 C743.70 1099.48 765.53 1099.10 787.20 1096.65 C808.87 1094.21 831.38 1091.35 852.78 1085.84 C874.19 1080.32 895.47 1072.79 915.62 1063.57 C935.78 1054.35 955.25 1042.99 973.70 1030.53 C992.15 1018.08 1009.75 1003.89 1026.31 988.84 C1042.87 973.79 1058.50 957.37 1073.06 940.25 C1087.61 923.14 1100.20 906.95 1113.64 886.14 C1127.08 865.33 1142.11 839.80 1153.70 815.39 C1165.29 790.98 1175.31 765.46 1183.17 739.66 C1191.03 713.86 1197.08 687.18 1200.85 660.60 C1204.62 634.03 1206.29 606.85 1205.81 580.20 C1205.33 553.55 1202.47 526.70 1197.97 500.71 C1193.47 474.73 1186.57 449.04 1178.83 424.28 C1171.08 399.51 1161.46 375.46 1151.52 352.12 C1141.58 328.78 1130.55 306.24 1119.18 284.24 C1107.82 262.23 1095.84 240.91 1083.33 220.07 C1070.82 199.24 1057.81 178.96 1044.11 159.22 C1030.41 139.47 1016.15 120.24 1001.12 101.62 C986.08 83.00 968.89 63.57 953.91 47.52 C938.93 31.47 925.91 19.00 911.21 5.34 C896.51 -8.32 881.28 -21.48 865.69 -34.43 C850.09 -47.38 834.04 -59.85 817.63 -72.37 C801.21 -84.89 784.43 -97.05 767.21 -109.55 C749.98 -122.05 732.52 -134.69 714.28 -147.37 C696.03 -160.05 677.48 -173.74 657.75 -185.65 C638.02 -197.57 617.35 -209.63 595.89 -218.86 C574.44 -228.09 551.80 -235.60 529.00 -241.03 C506.20 -246.47 482.58 -249.74 459.09 -251.45 C435.60 -253.17 411.69 -252.94 388.08 -251.33 C364.46 -249.72 340.72 -246.39 317.40 -241.80 C294.08 -237.20 273.00 -232.10 248.17 -223.76 C223.33 -215.41 194.07 -204.22 168.37 -191.71 C142.66 -179.21 117.50 -164.79 93.92 -148.73 C70.32 -132.68 47.60 -114.72 26.76 -95.36 C5.84 -76.01 -13.89 -54.73 -31.68 -32.24 C-49.29 -9.49 -64.82 14.87 -78.39 39.43 C-91.96 63.89 -103.41 89.63 -113.56 115.19 C-123.73 140.74 -131.92 166.94 -139.33 192.82 C-146.75 218.68 -152.74 244.64 -158.05 270.41 C-163.37 296.19 -167.71 321.84 -171.23 347.45 C-174.76 373.07 -177.47 398.58 -179.21 424.11 C-180.96 449.64 -181.87 475.13 -181.70 500.62 C-181.52 526.10 -180.78 551.68 -178.18 577.02 C-175.58 602.35 -171.95 627.86 -166.12 652.63 C-160.29 677.40 -152.64 702.07 -143.19 725.65 C-133.75 749.23 -122.42 772.39 -109.46 794.09 C-96.50 815.80 -81.61 836.69 -65.43 855.89 C-49.26 875.08 -31.17 892.99 -12.41 909.26 C6.35 925.53 26.73 940.11 47.14 953.53 C67.54 966.94 88.90 978.68 110.02 989.74 C131.13 1000.79 152.54 1010.60 173.84 1019.87 C195.13 1029.14 216.42 1037.56 237.79 1045.35 C259.17 1053.14 280.54 1060.26 302.09 1066.61 C323.64 1072.97 345.28 1078.66 367.12 1083.48 C388.97 1088.29 412.84 1092.48 433.15 1095.49 Z`

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

	__NoteInfo__000000_0.Name = `0`
	__NoteInfo__000000_0.IsKept = true

	__NoteInfo__000001_1.Name = `1`
	__NoteInfo__000001_1.IsKept = true

	__NoteInfo__000002_10.Name = `10`
	__NoteInfo__000002_10.IsKept = true

	__NoteInfo__000003_11.Name = `11`
	__NoteInfo__000003_11.IsKept = true

	__NoteInfo__000004_12.Name = `12`
	__NoteInfo__000004_12.IsKept = false

	__NoteInfo__000005_13.Name = `13`
	__NoteInfo__000005_13.IsKept = false

	__NoteInfo__000006_14.Name = `14`
	__NoteInfo__000006_14.IsKept = false

	__NoteInfo__000007_15.Name = `15`
	__NoteInfo__000007_15.IsKept = true

	__NoteInfo__000008_16.Name = `16`
	__NoteInfo__000008_16.IsKept = false

	__NoteInfo__000009_17.Name = `17`
	__NoteInfo__000009_17.IsKept = true

	__NoteInfo__000010_18.Name = `18`
	__NoteInfo__000010_18.IsKept = true

	__NoteInfo__000011_19.Name = `19`
	__NoteInfo__000011_19.IsKept = false

	__NoteInfo__000012_2.Name = `2`
	__NoteInfo__000012_2.IsKept = false

	__NoteInfo__000013_20.Name = `20`
	__NoteInfo__000013_20.IsKept = false

	__NoteInfo__000014_21.Name = `21`
	__NoteInfo__000014_21.IsKept = false

	__NoteInfo__000015_3.Name = `3`
	__NoteInfo__000015_3.IsKept = true

	__NoteInfo__000016_4.Name = `4`
	__NoteInfo__000016_4.IsKept = true

	__NoteInfo__000017_5.Name = `5`
	__NoteInfo__000017_5.IsKept = false

	__NoteInfo__000018_6.Name = `6`
	__NoteInfo__000018_6.IsKept = false

	__NoteInfo__000019_7.Name = `7`
	__NoteInfo__000019_7.IsKept = false

	__NoteInfo__000020_8.Name = `8`
	__NoteInfo__000020_8.IsKept = true

	__NoteInfo__000021_9.Name = `9`
	__NoteInfo__000021_9.IsKept = false

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.BackendColor = `#F1F1F1`
	__Parameter__000000_Reference.MinuteColor = `#5A5A5A`
	__Parameter__000000_Reference.HourColor = `#1E1E1E`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.InsideAngle = 112.000000
	__Parameter__000000_Reference.SideLength = 189.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.440000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.088000
	__Parameter__000000_Reference.NbPitchLines = 169
	__Parameter__000000_Reference.BeatLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbBeatLines = 300
	__Parameter__000000_Reference.NbBeatLinesPerCurve = 22
	__Parameter__000000_Reference.FirstVoiceShiftX = -0.070000
	__Parameter__000000_Reference.FirstVoiceShiftY = 1.740000
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
	__Parameter__000000_Reference.ActualBeatsTemporalShift = 8

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 189.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 189.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 310.563057
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 41.895485
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 189.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 189.000000
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
	__ShapeCategory__000006_6_Spiral_growth.IsExpanded = false

	__ShapeCategory__000007_7_Spiral_Growth_Bezier.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = true

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = true

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = true

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 303.156994
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -34.188573
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 348.825014
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 42.905060
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 229.261862
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 364.688401
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 272.363901
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 280.884772
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 200.340000
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterY = 0.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.HasBespokeRadius = true
	__SpiralCircle__000000_Construction_Circle_Spiral.BespopkeRadius = 14.130000
	__SpiralCircle__000000_Construction_Circle_Spiral.Pitch = 0
	__SpiralCircle__000000_Construction_Circle_Spiral.Color = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.FillOpacity = 0.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.Stroke = `blue`
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeOpacity = 0.800000
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeWidth = 1.000000
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeDashArray = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.Transform = ``
	__SpiralCircle__000000_Construction_Circle_Spiral.ShowName = false
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 303.156994
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 200.340000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -34.188573
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 303.156994
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 399.524034
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -34.188573
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -91.273598
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 94.999085
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 10.713535
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 73.090924
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 208.868336
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 303.156994
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -34.188573
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 59.252924
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -169.324165
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
	__FrontCurveStack__000000_Front_Curve_Stack.ShapeCategory = __ShapeCategory__000007_7_Spiral_Growth_Bezier
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000000_Non_Rotated_0_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000001_Non_Rotated_1_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000002_Non_Rotated_2_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000003_Non_Rotated_3_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000004_Non_Rotated_4_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000005_Non_Rotated_5_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000006_Non_Rotated_6_)
	__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves = append(__FrontCurveStack__000000_Front_Curve_Stack.FrontCurves, __FrontCurve__000007_Non_Rotated_7_)
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
	__Parameter__000000_Reference.BeatLines = __AxisGrid__000001_Measure_Lines
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
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000015_3)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000016_4)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000017_5)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000018_6)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000019_7)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000020_8)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000021_9)
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
	__SpiralCircle__000000_Construction_Circle_Spiral.ShapeCategory = __ShapeCategory__000006_6_Spiral_growth
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
