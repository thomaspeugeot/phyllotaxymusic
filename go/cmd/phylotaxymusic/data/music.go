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

	__MovingLine__000000_Cursor := (&models.MovingLine{}).Stage(stage)

	__NoteInfo__000000_0 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000001_1 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000002_2 := (&models.NoteInfo{}).Stage(stage)
	__NoteInfo__000003_3 := (&models.NoteInfo{}).Stage(stage)

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
	__Axis__000001_Construction_Axis.AngleDegree = 97.682917
	__Axis__000001_Construction_Axis.Length = 223.677161
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -29.903557
	__Axis__000001_Construction_Axis.EndY = 221.669235
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
	__Axis__000002_Initial_Axis.AngleDegree = 82.317083
	__Axis__000002_Initial_Axis.Length = 836.547168
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
	__Axis__000004_Rotated_Axis.Length = 836.547168
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -14.951779
	__Bezier__000005_Growth_Curve_Seed.StartY = 110.834618
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 72.258255
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 122.599400
	__Bezier__000005_Growth_Curve_Seed.EndX = 134.415521
	__Bezier__000005_Growth_Curve_Seed.EndY = 243.836159
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 47.205487
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 232.071377
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
	__Circle__000001_Construction_Circle.CenterX = -14.951779
	__Circle__000001_Construction_Circle.CenterY = 110.834618
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 328.638156
	__Circle__000005_Rotated_Next_Circle.CenterY = 44.333847
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
	__CircleGrid__000007_Second_Voice_Notes_Shift_Right.IsDisplayed = false

	__CircleGrid__000008_Second_Voice_notes.Name = `Second Voice notes`
	__CircleGrid__000008_Second_Voice_notes.IsDisplayed = true

	__FrontCurve__000000_Non_Rotated_0_.Name = `Non Rotated 0 `
	__FrontCurve__000000_Non_Rotated_0_.Path = `M821.64 536.73 C823.54 525.79 824.78 516.78 826.15 506.64 C827.53 496.50 828.67 486.32 829.90 475.88 C831.13 465.44 832.19 454.95 833.56 443.99 C834.92 433.04 836.26 421.98 838.07 410.15 C839.89 398.31 842.17 386.11 844.45 373.00 C846.73 359.89 850.13 345.73 851.74 331.49 C853.35 317.26 854.95 301.98 854.11 287.60 C853.26 273.23 850.69 258.78 846.67 245.23 C842.66 231.67 836.77 218.60 830.03 206.29 C823.29 193.97 815.09 182.28 806.25 171.34 C797.42 160.39 787.46 150.11 777.02 140.60 C766.58 131.08 756.56 122.88 743.59 114.24 C730.63 105.59 714.61 95.95 699.22 88.73 C683.84 81.50 667.63 75.37 651.28 70.90 C634.93 66.43 617.92 63.28 601.11 61.88 C584.29 60.49 567.06 60.70 550.38 62.53 C533.71 64.36 516.97 68.07 501.07 72.86 C485.17 77.65 469.66 84.34 454.98 91.27 C440.30 98.20 426.37 106.47 412.97 114.45 C399.58 122.43 386.96 130.92 374.62 139.16 C362.27 147.40 350.51 155.64 338.90 163.90 C327.29 172.16 316.06 180.32 304.98 188.69 C293.90 197.07 283.06 205.44 272.42 214.17 C261.78 222.89 250.52 232.59 241.12 241.04 C231.71 249.49 224.29 256.72 215.99 264.87 C207.69 273.02 199.61 281.39 191.34 289.94 C183.06 298.49 174.95 307.22 166.37 316.16 C157.78 325.09 149.21 334.16 139.84 343.55 C130.47 352.94 120.56 362.36 110.16 372.50 C99.75 382.64 87.87 392.93 77.41 404.40 C66.96 415.86 55.91 428.17 47.42 441.30 C38.92 454.43 31.81 468.70 26.44 483.17 C21.07 497.64 17.55 512.91 15.22 528.11 C12.89 543.31 12.16 558.94 12.46 574.37 C12.75 589.79 14.41 605.38 16.98 620.67 C19.56 635.95 22.62 649.81 27.90 666.07 C33.17 682.34 40.37 701.53 48.63 718.26 C56.88 734.98 66.55 751.33 77.44 766.42 C88.33 781.51 100.64 795.93 113.96 808.80 C127.28 821.67 142.01 833.50 157.37 843.63 C172.73 853.77 189.38 862.42 206.10 869.63 C222.82 876.83 240.47 882.25 257.67 886.87 C274.87 891.50 292.37 894.54 309.31 897.38 C326.24 900.22 342.92 902.10 359.26 903.92 C375.60 905.73 391.53 907.09 407.35 908.25 C423.17 909.41 438.66 910.32 454.17 910.87 C469.68 911.42 485.01 911.73 500.40 911.55 C515.79 911.37 531.21 911.23 546.51 909.78 C561.82 908.33 577.30 906.44 592.23 902.85 C607.16 899.25 622.05 894.39 636.07 888.23 C650.09 882.07 663.82 874.55 676.36 865.87 C688.89 857.18 700.79 847.07 711.30 836.12 C721.80 825.17 731.23 812.79 739.40 800.17 C747.56 787.54 754.29 773.76 760.30 760.37 C766.30 746.98 770.95 733.09 775.43 719.81 C779.91 706.52 783.54 693.40 787.19 680.68 C790.84 667.96 794.10 655.66 797.33 643.48 C800.55 631.30 803.62 619.49 806.52 607.62 C809.42 595.75 812.22 584.08 814.74 572.26 C817.26 560.45 819.74 547.66 821.64 536.73 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M241.12 241.04 C231.71 249.49 224.29 256.72 215.99 264.87 C207.69 273.02 199.61 281.39 191.34 289.94 C183.06 298.49 174.95 307.22 166.37 316.16 C157.78 325.09 149.21 334.16 139.84 343.55 C130.47 352.94 120.56 362.36 110.16 372.50 C99.75 382.64 87.87 392.93 77.41 404.40 C66.96 415.86 55.91 428.17 47.42 441.30 C38.92 454.43 31.81 468.70 26.44 483.17 C21.07 497.64 17.55 512.91 15.22 528.11 C12.89 543.31 12.16 558.94 12.46 574.37 C12.75 589.79 14.41 605.38 16.98 620.67 C19.56 635.95 22.62 649.81 27.90 666.07 C33.17 682.34 40.37 701.53 48.63 718.26 C56.88 734.98 66.55 751.33 77.44 766.42 C88.33 781.51 100.64 795.93 113.96 808.80 C127.28 821.67 142.01 833.50 157.37 843.63 C172.73 853.77 189.38 862.42 206.10 869.63 C222.82 876.83 240.47 882.25 257.67 886.87 C274.87 891.50 292.37 894.54 309.31 897.38 C326.24 900.22 342.92 902.10 359.26 903.92 C375.60 905.73 391.53 907.09 407.35 908.25 C423.17 909.41 438.66 910.32 454.17 910.87 C469.68 911.42 485.01 911.73 500.40 911.55 C515.79 911.37 532.41 910.68 546.51 909.78 C560.62 908.88 572.14 907.64 585.03 906.15 C597.92 904.67 610.79 902.81 623.88 900.87 C636.97 898.93 650.04 896.68 663.57 894.50 C677.11 892.33 690.69 889.99 705.07 887.83 C719.45 885.67 734.24 883.76 749.83 881.54 C765.42 879.32 782.23 877.86 798.61 874.53 C815.00 871.20 832.33 867.53 848.14 861.58 C863.96 855.64 879.37 847.89 893.48 838.86 C907.60 829.83 920.74 818.98 932.84 807.41 C944.94 795.84 956.01 782.90 966.10 769.45 C976.18 756.00 985.26 741.54 993.36 726.72 C1001.47 711.90 1008.20 698.00 1014.72 680.55 C1021.24 663.10 1028.12 641.87 1032.48 622.02 C1036.84 602.18 1039.77 581.71 1040.87 561.48 C1041.97 541.24 1041.45 520.61 1039.08 500.59 C1036.70 480.56 1032.49 460.46 1026.64 441.31 C1020.79 422.16 1012.90 403.36 1003.98 385.68 C995.05 368.01 984.18 351.16 973.09 335.25 C962.00 319.35 949.60 304.51 937.43 290.24 C925.27 275.96 912.60 262.63 900.10 249.61 C887.59 236.60 875.03 224.24 862.38 212.17 C849.72 200.10 837.08 188.48 824.14 177.20 C811.20 165.92 798.18 155.00 784.76 144.51 C771.33 134.01 757.85 123.53 743.59 114.24 C729.34 104.94 714.61 95.95 699.22 88.73 C683.84 81.50 667.63 75.37 651.28 70.90 C634.93 66.43 617.92 63.28 601.11 61.88 C584.29 60.49 567.06 60.70 550.38 62.53 C533.71 64.36 516.97 68.07 501.07 72.86 C485.17 77.65 469.66 84.34 454.98 91.27 C440.30 98.20 426.37 106.47 412.97 114.45 C399.58 122.43 386.96 130.92 374.62 139.16 C362.27 147.40 350.51 155.64 338.90 163.90 C327.29 172.16 316.06 180.32 304.98 188.69 C293.90 197.07 283.06 205.44 272.42 214.17 C261.78 222.89 250.52 232.59 241.12 241.04 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M546.51 909.78 C560.62 908.88 572.14 907.64 585.03 906.15 C597.92 904.67 610.79 902.81 623.88 900.87 C636.97 898.93 650.04 896.68 663.57 894.50 C677.11 892.33 690.69 889.99 705.07 887.83 C719.45 885.67 734.24 883.76 749.83 881.54 C765.42 879.32 782.23 877.86 798.61 874.53 C815.00 871.20 832.33 867.53 848.14 861.58 C863.96 855.64 879.37 847.89 893.48 838.86 C907.60 829.83 920.74 818.98 932.84 807.41 C944.94 795.84 956.01 782.90 966.10 769.45 C976.18 756.00 985.26 741.54 993.36 726.72 C1001.47 711.90 1008.20 698.00 1014.72 680.55 C1021.24 663.10 1028.12 641.87 1032.48 622.02 C1036.84 602.18 1039.77 581.71 1040.87 561.48 C1041.97 541.24 1041.45 520.61 1039.08 500.59 C1036.70 480.56 1032.49 460.46 1026.64 441.31 C1020.79 422.16 1012.90 403.36 1003.98 385.68 C995.05 368.01 984.18 351.16 973.09 335.25 C962.00 319.35 949.60 304.51 937.43 290.24 C925.27 275.96 912.60 262.63 900.10 249.61 C887.59 236.60 875.03 224.24 862.38 212.17 C849.72 200.10 837.08 188.48 824.14 177.20 C811.20 165.92 798.18 155.00 784.76 144.51 C771.33 134.01 756.44 123.16 743.59 114.24 C730.74 105.32 719.87 98.51 707.65 91.00 C695.42 83.49 682.95 76.35 670.24 69.19 C657.52 62.04 644.59 55.19 631.34 48.06 C618.10 40.93 604.67 33.96 590.74 26.42 C576.81 18.87 562.69 10.93 547.76 2.80 C532.83 -5.33 517.40 -14.75 501.15 -22.38 C484.89 -30.02 467.68 -37.90 450.25 -43.00 C432.81 -48.10 414.57 -51.47 396.56 -52.97 C378.55 -54.47 360.18 -53.87 342.19 -52.00 C324.20 -50.12 306.18 -46.45 288.64 -41.72 C271.10 -36.99 253.73 -30.76 236.94 -23.59 C220.15 -16.42 205.13 -9.13 187.89 1.32 C170.65 11.77 150.55 25.20 133.46 39.14 C116.35 53.09 99.99 68.52 85.24 84.96 C70.46 101.38 56.77 119.14 45.00 137.61 C33.25 156.07 22.98 175.78 14.69 195.80 C6.41 215.81 0.03 236.88 -4.74 257.75 C-9.51 278.61 -12.15 300.10 -13.94 321.01 C-15.73 341.91 -15.73 362.84 -15.48 383.19 C-15.22 403.55 -13.91 423.51 -12.41 443.12 C-10.92 462.73 -8.90 481.88 -6.51 500.86 C-4.13 519.83 -1.39 538.46 1.89 556.98 C5.17 575.51 8.84 593.80 13.17 611.98 C17.50 630.16 21.99 648.36 27.90 666.07 C33.81 683.78 40.37 701.53 48.63 718.26 C56.88 734.98 66.55 751.33 77.44 766.42 C88.33 781.51 100.64 795.93 113.96 808.80 C127.28 821.67 142.01 833.50 157.37 843.63 C172.73 853.77 189.38 862.42 206.10 869.63 C222.82 876.83 240.47 882.25 257.67 886.87 C274.87 891.50 292.37 894.54 309.31 897.38 C326.24 900.22 342.92 902.10 359.26 903.92 C375.60 905.73 391.53 907.09 407.35 908.25 C423.17 909.41 438.66 910.32 454.17 910.87 C469.68 911.42 485.01 911.73 500.40 911.55 C515.79 911.37 532.41 910.68 546.51 909.78 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M743.59 114.24 C730.74 105.32 719.87 98.51 707.65 91.00 C695.42 83.49 682.95 76.35 670.24 69.19 C657.52 62.04 644.59 55.19 631.34 48.06 C618.10 40.93 604.67 33.96 590.74 26.42 C576.81 18.87 562.69 10.93 547.76 2.80 C532.83 -5.33 517.40 -14.75 501.15 -22.38 C484.89 -30.02 467.68 -37.90 450.25 -43.00 C432.81 -48.10 414.57 -51.47 396.56 -52.97 C378.55 -54.47 360.18 -53.87 342.19 -52.00 C324.20 -50.12 306.18 -46.45 288.64 -41.72 C271.10 -36.99 253.73 -30.76 236.94 -23.59 C220.15 -16.42 205.13 -9.13 187.89 1.32 C170.65 11.77 150.55 25.20 133.46 39.14 C116.35 53.09 99.99 68.52 85.24 84.96 C70.46 101.38 56.77 119.14 45.00 137.61 C33.25 156.07 22.98 175.78 14.69 195.80 C6.41 215.81 0.03 236.88 -4.74 257.75 C-9.51 278.61 -12.15 300.10 -13.94 321.01 C-15.73 341.91 -15.73 362.84 -15.48 383.19 C-15.22 403.55 -13.91 423.51 -12.41 443.12 C-10.92 462.73 -8.90 481.88 -6.51 500.86 C-4.13 519.83 -1.39 538.46 1.89 556.98 C5.17 575.51 8.84 593.80 13.17 611.98 C17.50 630.16 22.91 649.66 27.90 666.07 C32.88 682.49 37.63 695.74 43.09 710.48 C48.56 725.22 54.50 739.80 60.68 754.52 C66.87 769.23 73.46 783.80 80.19 798.75 C86.91 813.71 93.92 828.61 101.03 844.23 C108.13 859.85 115.18 875.86 122.79 892.47 C130.40 909.09 137.66 926.97 146.70 943.92 C155.74 960.86 165.42 978.56 177.04 994.16 C188.66 1009.76 202.03 1024.46 216.41 1037.50 C230.79 1050.53 246.80 1062.10 263.32 1072.37 C279.84 1082.65 297.53 1091.49 315.53 1099.13 C333.52 1106.78 352.32 1113.09 371.27 1118.25 C390.22 1123.42 407.74 1127.32 429.22 1130.10 C450.70 1132.87 476.54 1135.09 500.16 1134.92 C523.79 1134.74 547.72 1132.85 570.99 1129.04 C594.26 1125.23 617.54 1119.56 639.76 1112.05 C661.99 1104.53 683.85 1095.01 704.35 1083.95 C724.84 1072.88 744.50 1059.72 762.73 1045.64 C780.96 1031.56 797.89 1015.59 813.74 999.47 C829.59 983.34 844.02 966.01 857.82 948.90 C871.62 931.79 884.30 914.24 896.54 896.80 C908.79 879.36 920.27 861.87 931.29 844.25 C942.30 826.62 952.76 808.98 962.64 791.05 C972.52 773.11 981.88 755.06 990.56 736.64 C999.24 718.22 1007.74 699.65 1014.72 680.55 C1021.71 661.45 1028.12 641.87 1032.48 622.02 C1036.84 602.18 1039.77 581.71 1040.87 561.48 C1041.97 541.24 1041.45 520.61 1039.08 500.59 C1036.70 480.56 1032.49 460.46 1026.64 441.31 C1020.79 422.16 1012.90 403.36 1003.98 385.68 C995.05 368.01 984.18 351.16 973.09 335.25 C962.00 319.35 949.60 304.51 937.43 290.24 C925.27 275.96 912.60 262.63 900.10 249.61 C887.59 236.60 875.03 224.24 862.38 212.17 C849.72 200.10 837.08 188.48 824.14 177.20 C811.20 165.92 798.18 155.00 784.76 144.51 C771.33 134.01 756.44 123.16 743.59 114.24 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M27.90 666.07 C32.88 682.49 37.63 695.74 43.09 710.48 C48.56 725.22 54.50 739.80 60.68 754.52 C66.87 769.23 73.46 783.80 80.19 798.75 C86.91 813.71 93.92 828.61 101.03 844.23 C108.13 859.85 115.18 875.86 122.79 892.47 C130.40 909.09 137.66 926.97 146.70 943.92 C155.74 960.86 165.42 978.56 177.04 994.16 C188.66 1009.76 202.03 1024.46 216.41 1037.50 C230.79 1050.53 246.80 1062.10 263.32 1072.37 C279.84 1082.65 297.53 1091.49 315.53 1099.13 C333.52 1106.78 352.32 1113.09 371.27 1118.25 C390.22 1123.42 407.74 1127.32 429.22 1130.10 C450.70 1132.87 476.54 1135.09 500.16 1134.92 C523.79 1134.74 547.72 1132.85 570.99 1129.04 C594.26 1125.23 617.54 1119.56 639.76 1112.05 C661.99 1104.53 683.85 1095.01 704.35 1083.95 C724.84 1072.88 744.50 1059.72 762.73 1045.64 C780.96 1031.56 797.89 1015.59 813.74 999.47 C829.59 983.34 844.02 966.01 857.82 948.90 C871.62 931.79 884.30 914.24 896.54 896.80 C908.79 879.36 920.27 861.87 931.29 844.25 C942.30 826.62 952.76 808.98 962.64 791.05 C972.52 773.11 981.88 755.06 990.56 736.64 C999.24 718.22 1007.86 697.93 1014.72 680.55 C1021.59 663.17 1026.49 648.64 1031.75 632.36 C1037.02 616.09 1041.75 599.61 1046.31 582.88 C1050.87 566.15 1054.95 549.25 1059.11 531.98 C1063.26 514.71 1067.10 497.29 1071.25 479.26 C1075.39 461.24 1079.75 442.93 1083.99 423.81 C1088.23 404.70 1093.48 384.82 1096.68 364.57 C1099.89 344.32 1102.99 323.10 1103.22 302.34 C1103.44 281.57 1101.67 260.43 1098.02 240.00 C1094.37 219.56 1088.47 199.28 1081.30 179.73 C1074.14 160.18 1065.11 141.04 1055.05 122.70 C1044.98 104.35 1033.40 86.57 1020.94 69.64 C1008.47 52.72 996.40 37.80 980.27 21.16 C964.14 4.52 944.02 -14.56 924.16 -30.20 C904.30 -45.84 883.06 -60.26 861.11 -72.68 C839.16 -85.11 815.99 -96.01 792.45 -104.72 C768.91 -113.44 744.36 -120.28 719.89 -124.96 C695.42 -129.65 670.26 -132.09 645.63 -132.82 C621.00 -133.56 596.18 -131.92 572.11 -129.38 C548.04 -126.85 524.30 -122.39 501.21 -117.62 C478.11 -112.85 455.61 -106.95 433.53 -100.76 C411.46 -94.58 389.95 -87.79 368.75 -80.50 C347.55 -73.22 326.80 -65.48 306.35 -57.06 C285.90 -48.64 265.79 -39.71 246.05 -29.99 C226.31 -20.26 206.65 -10.19 187.89 1.32 C169.13 12.84 150.55 25.20 133.46 39.14 C116.35 53.09 99.99 68.52 85.24 84.96 C70.46 101.38 56.77 119.14 45.00 137.61 C33.25 156.07 22.98 175.78 14.69 195.80 C6.41 215.81 0.03 236.88 -4.74 257.75 C-9.51 278.61 -12.15 300.10 -13.94 321.01 C-15.73 341.91 -15.73 362.84 -15.48 383.19 C-15.22 403.55 -13.91 423.51 -12.41 443.12 C-10.92 462.73 -8.90 481.88 -6.51 500.86 C-4.13 519.83 -1.39 538.46 1.89 556.98 C5.17 575.51 8.84 593.80 13.17 611.98 C17.50 630.16 22.91 649.66 27.90 666.07 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M1014.72 680.55 C1021.59 663.17 1026.49 648.64 1031.75 632.36 C1037.02 616.09 1041.75 599.61 1046.31 582.88 C1050.87 566.15 1054.95 549.25 1059.11 531.98 C1063.26 514.71 1067.10 497.29 1071.25 479.26 C1075.39 461.24 1079.75 442.93 1083.99 423.81 C1088.23 404.70 1093.48 384.82 1096.68 364.57 C1099.89 344.32 1102.99 323.10 1103.22 302.34 C1103.44 281.57 1101.67 260.43 1098.02 240.00 C1094.37 219.56 1088.47 199.28 1081.30 179.73 C1074.14 160.18 1065.11 141.04 1055.05 122.70 C1044.98 104.35 1033.40 86.57 1020.94 69.64 C1008.47 52.72 996.40 37.80 980.27 21.16 C964.14 4.52 944.02 -14.56 924.16 -30.20 C904.30 -45.84 883.06 -60.26 861.11 -72.68 C839.16 -85.11 815.99 -96.01 792.45 -104.72 C768.91 -113.44 744.36 -120.28 719.89 -124.96 C695.42 -129.65 670.26 -132.09 645.63 -132.82 C621.00 -133.56 596.18 -131.92 572.11 -129.38 C548.04 -126.85 524.30 -122.39 501.21 -117.62 C478.11 -112.85 455.61 -106.95 433.53 -100.76 C411.46 -94.58 389.95 -87.79 368.75 -80.50 C347.55 -73.22 326.80 -65.48 306.35 -57.06 C285.90 -48.64 265.79 -39.71 246.05 -29.99 C226.31 -20.26 205.37 -8.83 187.89 1.32 C170.41 11.47 156.51 20.57 141.15 30.93 C125.77 41.31 110.70 52.24 95.64 63.54 C80.54 74.84 65.73 86.60 50.66 98.62 C35.59 110.62 20.72 122.95 5.30 135.58 C-10.09 148.19 -25.76 160.91 -41.77 174.37 C-57.76 187.83 -74.96 201.29 -90.71 216.34 C-106.46 231.39 -122.67 247.30 -136.27 264.67 C-149.87 282.04 -162.08 300.98 -172.31 320.56 C-182.54 340.13 -190.85 361.01 -197.66 382.12 C-204.46 403.22 -209.47 425.19 -213.13 447.18 C-216.79 469.16 -218.83 491.67 -219.59 514.04 C-220.35 536.42 -220.13 556.86 -217.69 581.42 C-215.24 605.99 -211.09 635.23 -204.92 661.44 C-198.75 687.66 -190.65 713.77 -180.66 738.72 C-170.67 763.67 -158.68 788.18 -144.97 811.15 C-131.25 834.12 -115.49 856.23 -98.37 876.54 C-81.26 896.85 -62.12 915.80 -42.26 933.02 C-22.41 950.24 -0.84 965.67 20.75 979.87 C42.34 994.06 64.94 1006.48 87.29 1018.18 C109.63 1029.88 132.28 1040.26 154.82 1050.07 C177.36 1059.88 199.88 1068.79 222.50 1077.04 C245.12 1085.28 267.73 1092.82 290.54 1099.54 C313.35 1106.26 336.24 1112.29 359.36 1117.38 C382.47 1122.48 405.75 1127.18 429.22 1130.10 C452.69 1133.02 476.54 1135.09 500.16 1134.92 C523.79 1134.74 547.72 1132.85 570.99 1129.04 C594.26 1125.23 617.54 1119.56 639.76 1112.05 C661.99 1104.53 683.85 1095.01 704.35 1083.95 C724.84 1072.88 744.50 1059.72 762.73 1045.64 C780.96 1031.56 797.89 1015.59 813.74 999.47 C829.59 983.34 844.02 966.01 857.82 948.90 C871.62 931.79 884.30 914.24 896.54 896.80 C908.79 879.36 920.27 861.87 931.29 844.25 C942.30 826.62 952.76 808.98 962.64 791.05 C972.52 773.11 981.88 755.06 990.56 736.64 C999.24 718.22 1007.86 697.93 1014.72 680.55 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M187.89 1.32 C170.41 11.47 156.51 20.57 141.15 30.93 C125.77 41.31 110.70 52.24 95.64 63.54 C80.54 74.84 65.73 86.60 50.66 98.62 C35.59 110.62 20.72 122.95 5.30 135.58 C-10.09 148.19 -25.76 160.91 -41.77 174.37 C-57.76 187.83 -74.96 201.29 -90.71 216.34 C-106.46 231.39 -122.67 247.30 -136.27 264.67 C-149.87 282.04 -162.08 300.98 -172.31 320.56 C-182.54 340.13 -190.85 361.01 -197.66 382.12 C-204.46 403.22 -209.47 425.19 -213.13 447.18 C-216.79 469.16 -218.83 491.67 -219.59 514.04 C-220.35 536.42 -220.13 556.86 -217.69 581.42 C-215.24 605.99 -211.09 635.23 -204.92 661.44 C-198.75 687.66 -190.65 713.77 -180.66 738.72 C-170.67 763.67 -158.68 788.18 -144.97 811.15 C-131.25 834.12 -115.49 856.23 -98.37 876.54 C-81.26 896.85 -62.12 915.80 -42.26 933.02 C-22.41 950.24 -0.84 965.67 20.75 979.87 C42.34 994.06 64.94 1006.48 87.29 1018.18 C109.63 1029.88 132.28 1040.26 154.82 1050.07 C177.36 1059.88 199.88 1068.79 222.50 1077.04 C245.12 1085.28 267.73 1092.82 290.54 1099.54 C313.35 1106.26 336.24 1112.29 359.36 1117.38 C382.47 1122.48 407.74 1126.91 429.22 1130.10 C450.70 1133.28 468.46 1134.94 488.26 1136.50 C508.06 1138.05 527.95 1138.91 548.02 1139.42 C568.10 1139.93 588.23 1139.81 608.71 1139.54 C629.19 1139.28 649.74 1138.55 670.90 1137.86 C692.06 1137.17 713.52 1136.50 735.68 1135.41 C757.85 1134.32 780.94 1133.92 803.88 1131.33 C826.81 1128.75 850.63 1125.72 873.28 1119.89 C895.93 1114.06 918.45 1106.08 939.77 1096.33 C961.10 1086.57 981.71 1074.55 1001.23 1061.37 C1020.75 1048.19 1039.37 1033.17 1056.90 1017.25 C1074.42 1001.33 1090.96 983.95 1106.37 965.84 C1121.77 947.73 1135.09 930.60 1149.32 908.58 C1163.54 886.55 1179.44 859.54 1191.70 833.71 C1203.96 807.87 1214.57 780.87 1222.89 753.57 C1231.20 726.28 1237.60 698.04 1241.59 669.92 C1245.59 641.79 1247.35 613.03 1246.84 584.83 C1246.34 556.63 1243.30 528.22 1238.54 500.72 C1233.78 473.22 1226.48 446.05 1218.28 419.84 C1210.09 393.63 1199.91 368.18 1189.39 343.48 C1178.87 318.78 1167.20 294.94 1155.17 271.65 C1143.14 248.36 1130.47 225.80 1117.23 203.75 C1103.99 181.70 1090.22 160.24 1075.72 139.35 C1061.22 118.46 1046.13 98.10 1030.23 78.40 C1014.32 58.70 997.95 39.26 980.27 21.16 C962.59 3.06 944.02 -14.56 924.16 -30.20 C904.30 -45.84 883.06 -60.26 861.11 -72.68 C839.16 -85.11 815.99 -96.01 792.45 -104.72 C768.91 -113.44 744.36 -120.28 719.89 -124.96 C695.42 -129.65 670.26 -132.09 645.63 -132.82 C621.00 -133.56 596.18 -131.92 572.11 -129.38 C548.04 -126.85 524.30 -122.39 501.21 -117.62 C478.11 -112.85 455.61 -106.95 433.53 -100.76 C411.46 -94.58 389.95 -87.79 368.75 -80.50 C347.55 -73.22 326.80 -65.48 306.35 -57.06 C285.90 -48.64 265.79 -39.71 246.05 -29.99 C226.31 -20.26 205.37 -8.83 187.89 1.32 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M429.22 1130.10 C450.70 1133.28 468.46 1134.94 488.26 1136.50 C508.06 1138.05 527.95 1138.91 548.02 1139.42 C568.10 1139.93 588.23 1139.81 608.71 1139.54 C629.19 1139.28 649.74 1138.55 670.90 1137.86 C692.06 1137.17 713.52 1136.50 735.68 1135.41 C757.85 1134.32 780.94 1133.92 803.88 1131.33 C826.81 1128.75 850.63 1125.72 873.28 1119.89 C895.93 1114.06 918.45 1106.08 939.77 1096.33 C961.10 1086.57 981.71 1074.55 1001.23 1061.37 C1020.75 1048.19 1039.37 1033.17 1056.90 1017.25 C1074.42 1001.33 1090.96 983.95 1106.37 965.84 C1121.77 947.73 1135.09 930.60 1149.32 908.58 C1163.54 886.55 1179.44 859.54 1191.70 833.71 C1203.96 807.87 1214.57 780.87 1222.89 753.57 C1231.20 726.28 1237.60 698.04 1241.59 669.92 C1245.59 641.79 1247.35 613.03 1246.84 584.83 C1246.34 556.63 1243.30 528.22 1238.54 500.72 C1233.78 473.22 1226.48 446.05 1218.28 419.84 C1210.09 393.63 1199.91 368.18 1189.39 343.48 C1178.87 318.78 1167.20 294.94 1155.17 271.65 C1143.14 248.36 1130.47 225.80 1117.23 203.75 C1103.99 181.70 1090.22 160.24 1075.72 139.35 C1061.22 118.46 1046.13 98.10 1030.23 78.40 C1014.32 58.70 996.13 38.14 980.27 21.16 C964.41 4.17 950.64 -9.03 935.08 -23.48 C919.52 -37.94 903.42 -51.86 886.91 -65.56 C870.40 -79.27 853.42 -92.47 836.05 -105.72 C818.68 -118.97 800.93 -131.83 782.70 -145.06 C764.47 -158.29 745.99 -171.66 726.68 -185.08 C707.37 -198.50 687.73 -212.98 666.86 -225.59 C645.98 -238.20 624.11 -250.96 601.40 -260.73 C578.69 -270.50 554.73 -278.45 530.60 -284.20 C506.47 -289.94 481.48 -293.41 456.62 -295.22 C431.77 -297.04 406.46 -296.79 381.47 -295.09 C356.47 -293.39 331.34 -289.87 306.66 -285.01 C281.98 -280.15 259.67 -274.76 233.38 -265.92 C207.09 -257.09 176.12 -245.25 148.91 -232.02 C121.69 -218.79 95.05 -203.53 70.07 -186.53 C45.07 -169.53 21.00 -150.51 -1.08 -130.00 C-23.21 -109.45 -43.99 -86.88 -62.59 -63.17 C-81.12 -39.38 -97.65 -13.74 -112.10 12.26 C-126.50 38.19 -138.63 65.48 -149.38 92.57 C-160.14 119.63 -168.79 147.39 -176.63 174.78 C-184.48 202.17 -190.80 229.65 -196.43 256.93 C-202.05 284.21 -206.63 311.36 -210.36 338.47 C-214.09 365.58 -216.95 392.59 -218.79 419.60 C-220.64 446.62 -221.60 473.60 -221.41 500.57 C-221.23 527.54 -220.44 554.61 -217.69 581.42 C-214.94 608.23 -211.09 635.23 -204.92 661.44 C-198.75 687.66 -190.65 713.77 -180.66 738.72 C-170.67 763.67 -158.68 788.18 -144.97 811.15 C-131.25 834.12 -115.49 856.23 -98.37 876.54 C-81.26 896.85 -62.12 915.80 -42.26 933.02 C-22.41 950.24 -0.84 965.67 20.75 979.87 C42.34 994.06 64.94 1006.48 87.29 1018.18 C109.63 1029.88 132.28 1040.26 154.82 1050.07 C177.36 1059.88 199.88 1068.79 222.50 1077.04 C245.12 1085.28 267.73 1092.82 290.54 1099.54 C313.35 1106.26 336.24 1112.29 359.36 1117.38 C382.47 1122.48 407.74 1126.91 429.22 1130.10 Z`

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

	__MovingLine__000000_Cursor.Name = `Cursor`
	__MovingLine__000000_Cursor.IsDisplayed = true
	__MovingLine__000000_Cursor.AngleDegree = 90.000000
	__MovingLine__000000_Cursor.Length = 2000.000000
	__MovingLine__000000_Cursor.CenterX = 0.000000
	__MovingLine__000000_Cursor.CenterY = 366.000000
	__MovingLine__000000_Cursor.StartX = 0.000000
	__MovingLine__000000_Cursor.EndX = 400.000000
	__MovingLine__000000_Cursor.DurationSeconds = 16.000000
	__MovingLine__000000_Cursor.IsMoving = false
	__MovingLine__000000_Cursor.Color = ``
	__MovingLine__000000_Cursor.FillOpacity = 0.000000
	__MovingLine__000000_Cursor.Stroke = `lightgray`
	__MovingLine__000000_Cursor.StrokeOpacity = 0.700000
	__MovingLine__000000_Cursor.StrokeWidth = 6.000000
	__MovingLine__000000_Cursor.StrokeDashArray = ``
	__MovingLine__000000_Cursor.StrokeDashArrayWhenSelected = ``
	__MovingLine__000000_Cursor.Transform = ``

	__NoteInfo__000000_0.Name = `0`
	__NoteInfo__000000_0.IsKept = false

	__NoteInfo__000001_1.Name = `1`
	__NoteInfo__000001_1.IsKept = false

	__NoteInfo__000002_2.Name = `2`
	__NoteInfo__000002_2.IsKept = false

	__NoteInfo__000003_3.Name = `3`
	__NoteInfo__000003_3.IsKept = false

	__Parameter__000000_Reference.Name = `Reference`
	__Parameter__000000_Reference.BackendColor = `#F1F1F1`
	__Parameter__000000_Reference.MinuteColor = `#5A5A5A`
	__Parameter__000000_Reference.HourColor = `#1E1E1E`
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 16
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.InsideAngle = 112.000000
	__Parameter__000000_Reference.SideLength = 200.000000
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
	__Parameter__000000_Reference.NbOfBeatsInTheme = 6
	__Parameter__000000_Reference.FirstVoiceShiftX = -0.070000
	__Parameter__000000_Reference.FirstVoiceShiftY = 1.830000
	__Parameter__000000_Reference.PitchDifference = 12
	__Parameter__000000_Reference.BeatsPerSecond = 1.000000
	__Parameter__000000_Reference.Level = 2.200000
	__Parameter__000000_Reference.IsMinor = true
	__Parameter__000000_Reference.ThemeBinaryEncoding = 61447
	__Parameter__000000_Reference.OriginX = 40.000000
	__Parameter__000000_Reference.OriginY = 950.000000
	__Parameter__000000_Reference.SpiralOriginX = 500.000000
	__Parameter__000000_Reference.SpiralOriginY = 500.000000
	__Parameter__000000_Reference.OriginCrossWidth = 800.000000
	__Parameter__000000_Reference.SpiralRadiusRatio = 1.060000
	__Parameter__000000_Reference.ShowSpiralBezierConstruct = false
	__Parameter__000000_Reference.ShowInterpolationPoints = false
	__Parameter__000000_Reference.ActualBeatsTemporalShift = 8
	__Parameter__000000_Reference.IsPlaying = false

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 200.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 200.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 328.638156
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 44.333847
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 200.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 200.000000
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
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = false

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = true

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = true

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 320.801052
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -36.178384
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 369.126999
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 45.402179
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 242.605145
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 385.913652
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 288.215768
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 297.232562
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 212.000000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 320.801052
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 212.000000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -36.178384
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 320.801052
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 422.776756
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -36.178384
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -96.585818
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 100.528132
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 11.337074
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 77.344893
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 221.024695
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 320.801052
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -36.178384
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 62.701507
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -179.179011
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
	__Axis__000001_Construction_Axis.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Axis__000002_Initial_Axis.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Axis__000004_Rotated_Axis.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__AxisGrid__000000_Beat_Lines.Reference = __Axis__000000_Beat_Reference
	__AxisGrid__000000_Beat_Lines.ShapeCategory = __ShapeCategory__000008_8_Score_notation
	__AxisGrid__000001_Construction_Axis_Grid.Reference = __Axis__000001_Construction_Axis
	__AxisGrid__000001_Construction_Axis_Grid.ShapeCategory = __ShapeCategory__000004_4_Construction
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
	__MovingLine__000000_Cursor.ShapeCategory = __ShapeCategory__000008_8_Score_notation
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
	__Parameter__000000_Reference.FirstVoiceShiftRigth = __BezierGrid__000003_First_Voice_Shift_Right
	__Parameter__000000_Reference.SecondVoice = __BezierGrid__000000_2nb_Voice
	__Parameter__000000_Reference.SecondVoiceShiftedRight = __BezierGrid__000001_2nd_voice_shifted_right
	__Parameter__000000_Reference.FirstVoiceNotes = __CircleGrid__000002_First_Voice_notes
	__Parameter__000000_Reference.FirstVoiceNotesShiftedRight = __CircleGrid__000001_First_Voice_note_shifted_right
	__Parameter__000000_Reference.SecondVoiceNotes = __CircleGrid__000008_Second_Voice_notes
	__Parameter__000000_Reference.SecondVoiceNotesShiftedRight = __CircleGrid__000007_Second_Voice_Notes_Shift_Right
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000000_0)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000001_1)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000002_2)
	__Parameter__000000_Reference.NoteInfos = append(__Parameter__000000_Reference.NoteInfos, __NoteInfo__000003_3)
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Horizontal_Axis
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Vertical_Axis
	__Parameter__000000_Reference.SpiralOrigin = __SpiralOrigin__000000_Spiral_Origin
	__Parameter__000000_Reference.Cursor = __MovingLine__000000_Cursor
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
