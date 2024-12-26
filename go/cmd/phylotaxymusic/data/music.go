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
	__Axis__000001_Construction_Axis.Length = 95.062794
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -12.709012
	__Axis__000001_Construction_Axis.EndY = 94.209425
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
	__Axis__000002_Initial_Axis.Length = 355.532547
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
	__Axis__000004_Rotated_Axis.Length = 355.532547
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
	__AxisGrid__000000_Beat_Lines.IsDisplayed = true

	__AxisGrid__000001_Construction_Axis_Grid.Name = `Construction Axis Grid`
	__AxisGrid__000001_Construction_Axis_Grid.IsDisplayed = false

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
	__Bezier__000005_Growth_Curve_Seed.StartX = -6.354506
	__Bezier__000005_Growth_Curve_Seed.StartY = 47.104712
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 30.709758
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 52.104745
	__Bezier__000005_Growth_Curve_Seed.EndX = 57.126596
	__Bezier__000005_Growth_Curve_Seed.EndY = 103.630367
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 20.062332
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 98.630335
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
	__Circle__000001_Construction_Circle.CenterX = -6.354506
	__Circle__000001_Construction_Circle.CenterY = 47.104712
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 139.671216
	__Circle__000005_Rotated_Next_Circle.CenterY = 18.841885
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M637.12 516.01 C637.93 511.36 638.45 507.53 639.04 503.22 C639.62 498.91 640.11 494.58 640.63 490.15 C641.16 485.71 641.61 481.25 642.19 476.59 C642.77 471.94 643.34 467.24 644.11 462.21 C644.88 457.18 645.86 451.99 646.82 446.42 C647.79 440.85 649.24 434.83 649.93 428.78 C650.61 422.73 651.30 416.24 650.94 410.13 C650.58 404.02 649.49 397.88 647.78 392.12 C646.08 386.36 643.58 380.81 640.71 375.58 C637.85 370.34 634.36 365.38 630.61 360.73 C626.85 356.07 622.62 351.71 618.19 347.67 C613.75 343.63 609.49 340.14 603.98 336.47 C598.47 332.80 591.66 328.71 585.12 325.64 C578.58 322.58 571.69 319.97 564.74 318.08 C557.79 316.18 550.56 314.85 543.41 314.26 C536.26 313.67 528.94 313.76 521.85 314.54 C514.76 315.32 507.64 316.91 500.88 318.94 C494.12 320.98 487.52 323.82 481.28 326.77 C475.03 329.71 469.11 333.23 463.41 336.61 C457.72 340.00 452.35 343.61 447.10 347.11 C441.86 350.61 436.85 354.11 431.92 357.61 C426.99 361.11 422.21 364.58 417.51 368.13 C412.80 371.69 408.19 375.24 403.67 378.94 C399.15 382.65 394.37 386.76 390.38 390.35 C386.39 393.94 383.24 397.00 379.72 400.47 C376.19 403.93 372.76 407.48 369.25 411.11 C365.74 414.75 362.30 418.46 358.65 422.25 C355.01 426.05 351.37 429.90 347.40 433.89 C343.42 437.89 339.21 441.89 334.80 446.21 C330.38 450.52 325.33 454.89 320.89 459.77 C316.45 464.65 311.76 469.88 308.15 475.47 C304.54 481.05 301.52 487.12 299.23 493.28 C296.95 499.43 295.45 505.93 294.45 512.39 C293.46 518.85 293.15 525.50 293.27 532.06 C293.39 538.62 294.09 545.24 295.18 551.74 C296.27 558.24 297.57 564.13 299.80 571.04 C302.04 577.95 305.09 586.11 308.60 593.22 C312.10 600.33 316.21 607.28 320.83 613.69 C325.46 620.11 330.69 626.23 336.35 631.70 C342.00 637.17 348.26 642.20 354.79 646.50 C361.31 650.81 368.39 654.49 375.49 657.55 C382.59 660.61 390.10 662.91 397.41 664.87 C404.72 666.84 412.15 668.13 419.35 669.33 C426.54 670.54 433.63 671.34 440.58 672.11 C447.52 672.87 454.29 673.45 461.01 673.94 C467.74 674.43 474.32 674.82 480.91 675.05 C487.50 675.29 494.02 675.42 500.56 675.34 C507.10 675.26 513.66 675.20 520.16 674.58 C526.67 673.97 533.25 673.16 539.59 671.63 C545.94 670.11 552.27 668.04 558.23 665.42 C564.19 662.80 570.02 659.60 575.35 655.91 C580.68 652.22 585.74 647.92 590.20 643.26 C594.67 638.61 598.68 633.35 602.15 627.98 C605.62 622.61 608.48 616.76 611.03 611.06 C613.59 605.37 615.56 599.47 617.47 593.82 C619.37 588.18 620.91 582.60 622.47 577.19 C624.02 571.78 625.41 566.56 626.78 561.38 C628.15 556.20 629.45 551.18 630.69 546.14 C631.92 541.09 633.11 536.13 634.18 531.11 C635.26 526.09 636.31 520.65 637.12 516.01 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M390.38 390.35 C386.39 393.94 383.24 397.00 379.72 400.47 C376.19 403.93 372.76 407.48 369.25 411.11 C365.74 414.75 362.30 418.46 358.65 422.25 C355.01 426.05 351.37 429.90 347.40 433.89 C343.42 437.89 339.21 441.89 334.80 446.21 C330.38 450.52 325.33 454.89 320.89 459.77 C316.45 464.65 311.76 469.88 308.15 475.47 C304.54 481.05 301.52 487.12 299.23 493.28 C296.95 499.43 295.45 505.93 294.45 512.39 C293.46 518.85 293.15 525.50 293.27 532.06 C293.39 538.62 294.09 545.24 295.18 551.74 C296.27 558.24 297.57 564.13 299.80 571.04 C302.04 577.95 305.09 586.11 308.60 593.22 C312.10 600.33 316.21 607.28 320.83 613.69 C325.46 620.11 330.69 626.23 336.35 631.70 C342.00 637.17 348.26 642.20 354.79 646.50 C361.31 650.81 368.39 654.49 375.49 657.55 C382.59 660.61 390.10 662.91 397.41 664.87 C404.72 666.84 412.15 668.13 419.35 669.33 C426.54 670.54 433.63 671.34 440.58 672.11 C447.52 672.87 454.29 673.45 461.01 673.94 C467.74 674.43 474.32 674.82 480.91 675.05 C487.50 675.29 494.02 675.42 500.56 675.34 C507.10 675.26 514.17 674.97 520.16 674.58 C526.15 674.20 531.05 673.67 536.53 673.04 C542.01 672.41 547.48 671.62 553.04 670.79 C558.61 669.97 564.16 669.01 569.91 668.08 C575.67 667.16 581.44 666.16 587.55 665.24 C593.66 664.32 599.95 663.51 606.58 662.57 C613.21 661.63 620.35 661.00 627.31 659.59 C634.28 658.17 641.65 656.61 648.37 654.08 C655.09 651.55 661.64 648.26 667.64 644.42 C673.64 640.58 679.23 635.97 684.37 631.05 C689.52 626.13 694.22 620.63 698.51 614.91 C702.80 609.19 706.65 603.05 710.10 596.75 C713.55 590.45 716.41 584.54 719.18 577.12 C721.96 569.71 724.88 560.68 726.74 552.24 C728.59 543.81 729.84 535.11 730.31 526.51 C730.78 517.91 730.56 509.14 729.55 500.63 C728.54 492.12 726.75 483.58 724.27 475.44 C721.78 467.30 718.43 459.31 714.64 451.80 C710.85 444.29 706.23 437.13 701.51 430.37 C696.80 423.61 691.53 417.30 686.36 411.24 C681.19 405.17 675.81 399.51 670.49 393.98 C665.18 388.45 659.84 383.20 654.46 378.07 C649.08 372.94 643.71 368.01 638.21 363.22 C632.72 358.42 627.18 353.79 621.48 349.33 C615.77 344.87 610.04 340.42 603.98 336.47 C597.92 332.53 591.66 328.71 585.12 325.64 C578.58 322.58 571.69 319.97 564.74 318.08 C557.79 316.18 550.56 314.85 543.41 314.26 C536.26 313.67 528.94 313.76 521.85 314.54 C514.76 315.32 507.64 316.91 500.88 318.94 C494.12 320.98 487.52 323.82 481.28 326.77 C475.03 329.71 469.11 333.23 463.41 336.61 C457.72 340.00 452.35 343.61 447.10 347.11 C441.86 350.61 436.85 354.11 431.92 357.61 C426.99 361.11 422.21 364.58 417.51 368.13 C412.80 371.69 408.19 375.24 403.67 378.94 C399.15 382.65 394.37 386.76 390.38 390.35 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M520.16 674.58 C526.15 674.20 531.05 673.67 536.53 673.04 C542.01 672.41 547.48 671.62 553.04 670.79 C558.61 669.97 564.16 669.01 569.91 668.08 C575.67 667.16 581.44 666.16 587.55 665.24 C593.66 664.32 599.95 663.51 606.58 662.57 C613.21 661.63 620.35 661.00 627.31 659.59 C634.28 658.17 641.65 656.61 648.37 654.08 C655.09 651.55 661.64 648.26 667.64 644.42 C673.64 640.58 679.23 635.97 684.37 631.05 C689.52 626.13 694.22 620.63 698.51 614.91 C702.80 609.19 706.65 603.05 710.10 596.75 C713.55 590.45 716.41 584.54 719.18 577.12 C721.96 569.71 724.88 560.68 726.74 552.24 C728.59 543.81 729.84 535.11 730.31 526.51 C730.78 517.91 730.56 509.14 729.55 500.63 C728.54 492.12 726.75 483.58 724.27 475.44 C721.78 467.30 718.43 459.31 714.64 451.80 C710.85 444.29 706.23 437.13 701.51 430.37 C696.80 423.61 691.53 417.30 686.36 411.24 C681.19 405.17 675.81 399.51 670.49 393.98 C665.18 388.45 659.84 383.20 654.46 378.07 C649.08 372.94 643.71 368.01 638.21 363.22 C632.72 358.42 627.18 353.79 621.48 349.33 C615.77 344.87 609.44 340.26 603.98 336.47 C598.52 332.69 593.90 329.79 588.70 326.61 C583.51 323.42 578.21 320.39 572.80 317.35 C567.40 314.31 561.90 311.40 556.27 308.38 C550.64 305.35 544.94 302.39 539.02 299.19 C533.09 295.99 527.09 292.62 520.75 289.17 C514.40 285.72 507.84 281.73 500.94 278.49 C494.03 275.26 486.71 271.91 479.30 269.75 C471.89 267.60 464.14 266.17 456.48 265.55 C448.83 264.92 441.01 265.18 433.36 265.99 C425.71 266.79 418.05 268.36 410.59 270.38 C403.12 272.40 395.74 275.05 388.59 278.10 C381.45 281.15 375.05 284.24 367.72 288.67 C360.38 293.10 351.83 298.79 344.56 304.68 C337.30 310.57 330.37 317.07 324.13 324.01 C317.90 330.96 312.12 338.49 307.16 346.33 C302.19 354.17 297.85 362.56 294.33 371.07 C290.82 379.59 288.11 388.55 286.08 397.43 C284.05 406.31 282.92 415.44 282.15 424.34 C281.38 433.23 281.37 442.13 281.47 450.78 C281.57 459.43 282.12 467.92 282.75 476.26 C283.37 484.60 284.23 492.74 285.23 500.81 C286.24 508.88 287.40 516.80 288.79 524.67 C290.17 532.54 291.72 540.32 293.56 548.05 C295.40 555.78 297.30 563.51 299.80 571.04 C302.31 578.57 305.09 586.11 308.60 593.22 C312.10 600.33 316.21 607.28 320.83 613.69 C325.46 620.11 330.69 626.23 336.35 631.70 C342.00 637.17 348.26 642.20 354.79 646.50 C361.31 650.81 368.39 654.49 375.49 657.55 C382.59 660.61 390.10 662.91 397.41 664.87 C404.72 666.84 412.15 668.13 419.35 669.33 C426.54 670.54 433.63 671.34 440.58 672.11 C447.52 672.87 454.29 673.45 461.01 673.94 C467.74 674.43 474.32 674.82 480.91 675.05 C487.50 675.29 494.02 675.42 500.56 675.34 C507.10 675.26 514.17 674.97 520.16 674.58 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M603.98 336.47 C598.52 332.69 593.90 329.79 588.70 326.61 C583.51 323.42 578.21 320.39 572.80 317.35 C567.40 314.31 561.90 311.40 556.27 308.38 C550.64 305.35 544.94 302.39 539.02 299.19 C533.09 295.99 527.09 292.62 520.75 289.17 C514.40 285.72 507.84 281.73 500.94 278.49 C494.03 275.26 486.71 271.91 479.30 269.75 C471.89 267.60 464.14 266.17 456.48 265.55 C448.83 264.92 441.01 265.18 433.36 265.99 C425.71 266.79 418.05 268.36 410.59 270.38 C403.12 272.40 395.74 275.05 388.59 278.10 C381.45 281.15 375.05 284.24 367.72 288.67 C360.38 293.10 351.83 298.79 344.56 304.68 C337.30 310.57 330.37 317.07 324.13 324.01 C317.90 330.96 312.12 338.49 307.16 346.33 C302.19 354.17 297.85 362.56 294.33 371.07 C290.82 379.59 288.11 388.55 286.08 397.43 C284.05 406.31 282.92 415.44 282.15 424.34 C281.38 433.23 281.37 442.13 281.47 450.78 C281.57 459.43 282.12 467.92 282.75 476.26 C283.37 484.60 284.23 492.74 285.23 500.81 C286.24 508.88 287.40 516.80 288.79 524.67 C290.17 532.54 291.72 540.32 293.56 548.05 C295.40 555.78 297.69 564.06 299.80 571.04 C301.92 578.02 303.93 583.65 306.25 589.92 C308.57 596.18 311.09 602.38 313.71 608.63 C316.34 614.89 319.14 621.08 321.99 627.44 C324.85 633.79 327.83 640.13 330.84 646.77 C333.86 653.40 336.85 660.21 340.08 667.27 C343.32 674.33 346.39 681.93 350.23 689.14 C354.07 696.34 358.19 703.86 363.12 710.49 C368.06 717.12 373.74 723.37 379.85 728.91 C385.96 734.45 392.76 739.36 399.78 743.73 C406.80 748.09 414.32 751.85 421.97 755.10 C429.62 758.34 437.61 761.03 445.66 763.22 C453.71 765.41 461.16 767.07 470.29 768.25 C479.42 769.43 490.40 770.37 500.44 770.29 C510.49 770.21 520.66 769.41 530.55 767.78 C540.44 766.16 550.33 763.75 559.78 760.56 C569.23 757.36 578.52 753.31 587.23 748.61 C595.95 743.90 604.30 738.31 612.05 732.32 C619.80 726.33 627.00 719.55 633.74 712.69 C640.47 705.84 646.61 698.47 652.47 691.19 C658.34 683.92 663.73 676.46 668.94 669.05 C674.14 661.63 679.03 654.20 683.71 646.71 C688.39 639.21 692.84 631.72 697.04 624.09 C701.24 616.47 705.22 608.79 708.91 600.96 C712.60 593.14 716.21 585.24 719.18 577.12 C722.16 569.00 724.88 560.68 726.74 552.24 C728.59 543.81 729.84 535.11 730.31 526.51 C730.78 517.91 730.56 509.14 729.55 500.63 C728.54 492.12 726.75 483.58 724.27 475.44 C721.78 467.30 718.43 459.31 714.64 451.80 C710.85 444.29 706.23 437.13 701.51 430.37 C696.80 423.61 691.53 417.30 686.36 411.24 C681.19 405.17 675.81 399.51 670.49 393.98 C665.18 388.45 659.84 383.20 654.46 378.07 C649.08 372.94 643.71 368.01 638.21 363.22 C632.72 358.42 627.18 353.79 621.48 349.33 C615.77 344.87 609.44 340.26 603.98 336.47 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M299.80 571.04 C301.92 578.02 303.93 583.65 306.25 589.92 C308.57 596.18 311.09 602.38 313.71 608.63 C316.34 614.89 319.14 621.08 321.99 627.44 C324.85 633.79 327.83 640.13 330.84 646.77 C333.86 653.40 336.85 660.21 340.08 667.27 C343.32 674.33 346.39 681.93 350.23 689.14 C354.07 696.34 358.19 703.86 363.12 710.49 C368.06 717.12 373.74 723.37 379.85 728.91 C385.96 734.45 392.76 739.36 399.78 743.73 C406.80 748.09 414.32 751.85 421.97 755.10 C429.62 758.34 437.61 761.03 445.66 763.22 C453.71 765.41 461.16 767.07 470.29 768.25 C479.42 769.43 490.40 770.37 500.44 770.29 C510.49 770.21 520.66 769.41 530.55 767.78 C540.44 766.16 550.33 763.75 559.78 760.56 C569.23 757.36 578.52 753.31 587.23 748.61 C595.95 743.90 604.30 738.31 612.05 732.32 C619.80 726.33 627.00 719.55 633.74 712.69 C640.47 705.84 646.61 698.47 652.47 691.19 C658.34 683.92 663.73 676.46 668.94 669.05 C674.14 661.63 679.03 654.20 683.71 646.71 C688.39 639.21 692.84 631.72 697.04 624.09 C701.24 616.47 705.22 608.79 708.91 600.96 C712.60 593.14 716.27 584.51 719.18 577.12 C722.10 569.74 724.19 563.56 726.43 556.64 C728.67 549.72 730.68 542.72 732.62 535.61 C734.56 528.50 736.29 521.31 738.06 513.97 C739.83 506.63 741.46 499.23 743.23 491.57 C744.99 483.90 746.85 476.12 748.65 468.00 C750.45 459.87 752.69 451.42 754.05 442.81 C755.41 434.21 756.74 425.19 756.83 416.36 C756.93 407.54 756.18 398.55 754.63 389.87 C753.08 381.18 750.58 372.56 747.53 364.25 C744.49 355.95 740.65 347.81 736.38 340.02 C732.11 332.22 727.18 324.66 721.89 317.47 C716.59 310.28 711.47 303.94 704.61 296.87 C697.76 289.80 689.21 281.70 680.77 275.05 C672.33 268.41 663.31 262.28 653.98 257.01 C644.65 251.74 634.80 247.11 624.80 243.41 C614.80 239.71 604.36 236.81 593.96 234.83 C583.56 232.85 572.87 231.82 562.40 231.52 C551.93 231.21 541.38 231.92 531.15 233.01 C520.91 234.09 510.82 236.00 501.00 238.04 C491.18 240.07 481.61 242.59 472.23 245.23 C462.84 247.87 453.69 250.77 444.67 253.87 C435.66 256.98 426.83 260.27 418.12 263.86 C409.42 267.44 400.87 271.25 392.47 275.38 C384.07 279.52 375.70 283.79 367.72 288.67 C359.73 293.56 351.83 298.79 344.56 304.68 C337.30 310.57 330.37 317.07 324.13 324.01 C317.90 330.96 312.12 338.49 307.16 346.33 C302.19 354.17 297.85 362.56 294.33 371.07 C290.82 379.59 288.11 388.55 286.08 397.43 C284.05 406.31 282.92 415.44 282.15 424.34 C281.38 433.23 281.37 442.13 281.47 450.78 C281.57 459.43 282.12 467.92 282.75 476.26 C283.37 484.60 284.23 492.74 285.23 500.81 C286.24 508.88 287.40 516.80 288.79 524.67 C290.17 532.54 291.72 540.32 293.56 548.05 C295.40 555.78 297.69 564.06 299.80 571.04 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M719.18 577.12 C722.10 569.74 724.19 563.56 726.43 556.64 C728.67 549.72 730.68 542.72 732.62 535.61 C734.56 528.50 736.29 521.31 738.06 513.97 C739.83 506.63 741.46 499.23 743.23 491.57 C744.99 483.90 746.85 476.12 748.65 468.00 C750.45 459.87 752.69 451.42 754.05 442.81 C755.41 434.21 756.74 425.19 756.83 416.36 C756.93 407.54 756.18 398.55 754.63 389.87 C753.08 381.18 750.58 372.56 747.53 364.25 C744.49 355.95 740.65 347.81 736.38 340.02 C732.11 332.22 727.18 324.66 721.89 317.47 C716.59 310.28 711.47 303.94 704.61 296.87 C697.76 289.80 689.21 281.70 680.77 275.05 C672.33 268.41 663.31 262.28 653.98 257.01 C644.65 251.74 634.80 247.11 624.80 243.41 C614.80 239.71 604.36 236.81 593.96 234.83 C583.56 232.85 572.87 231.82 562.40 231.52 C551.93 231.21 541.38 231.92 531.15 233.01 C520.91 234.09 510.82 236.00 501.00 238.04 C491.18 240.07 481.61 242.59 472.23 245.23 C462.84 247.87 453.69 250.77 444.67 253.87 C435.66 256.98 426.83 260.27 418.12 263.86 C409.42 267.44 400.87 271.25 392.47 275.38 C384.07 279.52 375.16 284.37 367.72 288.67 C360.27 292.98 354.36 296.83 347.83 301.21 C341.29 305.59 334.90 310.20 328.51 314.96 C322.13 319.72 315.88 324.68 309.52 329.77 C303.16 334.85 296.89 340.09 290.38 345.46 C283.87 350.83 277.23 356.26 270.44 362.00 C263.66 367.73 256.36 373.48 249.66 379.89 C242.97 386.30 236.07 393.08 230.28 400.48 C224.49 407.87 219.28 415.93 214.92 424.25 C210.56 432.58 207.01 441.46 204.10 450.43 C201.20 459.41 199.05 468.74 197.48 478.09 C195.92 487.44 195.04 497.00 194.70 506.51 C194.37 516.02 194.45 524.71 195.47 535.15 C196.50 545.59 198.25 558.01 200.87 569.15 C203.48 580.29 206.91 591.38 211.15 601.99 C215.39 612.59 220.48 623.00 226.30 632.76 C232.12 642.52 238.82 651.91 246.09 660.54 C253.36 669.17 261.49 677.22 269.92 684.54 C278.36 691.86 287.52 698.41 296.70 704.44 C305.87 710.47 315.48 715.75 324.97 720.72 C334.47 725.69 344.09 730.09 353.67 734.26 C363.25 738.43 372.82 742.21 382.43 745.71 C392.05 749.22 401.65 752.42 411.35 755.27 C421.04 758.13 430.77 760.69 440.60 762.85 C450.42 765.01 460.32 767.01 470.29 768.25 C480.26 769.49 490.40 770.37 500.44 770.29 C510.49 770.21 520.66 769.41 530.55 767.78 C540.44 766.16 550.33 763.75 559.78 760.56 C569.23 757.36 578.52 753.31 587.23 748.61 C595.95 743.90 604.30 738.31 612.05 732.32 C619.80 726.33 627.00 719.55 633.74 712.69 C640.47 705.84 646.61 698.47 652.47 691.19 C658.34 683.92 663.73 676.46 668.94 669.05 C674.14 661.63 679.03 654.20 683.71 646.71 C688.39 639.21 692.84 631.72 697.04 624.09 C701.24 616.47 705.22 608.79 708.91 600.96 C712.60 593.14 716.27 584.51 719.18 577.12 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M367.72 288.67 C360.27 292.98 354.36 296.83 347.83 301.21 C341.29 305.59 334.90 310.20 328.51 314.96 C322.13 319.72 315.88 324.68 309.52 329.77 C303.16 334.85 296.89 340.09 290.38 345.46 C283.87 350.83 277.23 356.26 270.44 362.00 C263.66 367.73 256.36 373.48 249.66 379.89 C242.97 386.30 236.07 393.08 230.28 400.48 C224.49 407.87 219.28 415.93 214.92 424.25 C210.56 432.58 207.01 441.46 204.10 450.43 C201.20 459.41 199.05 468.74 197.48 478.09 C195.92 487.44 195.04 497.00 194.70 506.51 C194.37 516.02 194.45 524.71 195.47 535.15 C196.50 545.59 198.25 558.01 200.87 569.15 C203.48 580.29 206.91 591.38 211.15 601.99 C215.39 612.59 220.48 623.00 226.30 632.76 C232.12 642.52 238.82 651.91 246.09 660.54 C253.36 669.17 261.49 677.22 269.92 684.54 C278.36 691.86 287.52 698.41 296.70 704.44 C305.87 710.47 315.48 715.75 324.97 720.72 C334.47 725.69 344.09 730.09 353.67 734.26 C363.25 738.43 372.82 742.21 382.43 745.71 C392.05 749.22 401.65 752.42 411.35 755.27 C421.04 758.13 430.77 760.69 440.60 762.85 C450.42 765.01 461.16 766.90 470.29 768.25 C479.42 769.60 486.97 770.30 495.38 770.96 C503.80 771.62 512.25 771.99 520.79 772.20 C529.32 772.41 537.88 772.36 546.58 772.25 C555.28 772.14 564.02 771.82 573.01 771.53 C582.01 771.23 591.13 770.94 600.55 770.48 C609.97 770.02 619.79 769.84 629.53 768.74 C639.28 767.64 649.41 766.36 659.03 763.88 C668.66 761.39 678.24 758.00 687.30 753.86 C696.37 749.71 705.13 744.60 713.43 738.99 C721.72 733.39 729.64 727.00 737.09 720.24 C744.54 713.47 751.57 706.08 758.12 698.38 C764.67 690.68 770.33 683.40 776.38 674.04 C782.43 664.68 789.19 653.19 794.40 642.21 C799.62 631.23 804.13 619.75 807.66 608.15 C811.20 596.54 813.92 584.54 815.62 572.59 C817.32 560.63 818.07 548.41 817.86 536.42 C817.65 524.44 816.36 512.36 814.34 500.67 C812.32 488.98 809.21 477.43 805.73 466.29 C802.25 455.16 797.93 444.34 793.46 433.84 C788.99 423.34 784.03 413.21 778.92 403.31 C773.81 393.42 768.43 383.83 762.80 374.46 C757.18 365.09 751.33 355.97 745.17 347.09 C739.01 338.21 732.60 329.56 725.84 321.19 C719.08 312.82 712.12 304.56 704.61 296.87 C697.10 289.18 689.21 281.70 680.77 275.05 C672.33 268.41 663.31 262.28 653.98 257.01 C644.65 251.74 634.80 247.11 624.80 243.41 C614.80 239.71 604.36 236.81 593.96 234.83 C583.56 232.85 572.87 231.82 562.40 231.52 C551.93 231.21 541.38 231.92 531.15 233.01 C520.91 234.09 510.82 236.00 501.00 238.04 C491.18 240.07 481.61 242.59 472.23 245.23 C462.84 247.87 453.69 250.77 444.67 253.87 C435.66 256.98 426.83 260.27 418.12 263.86 C409.42 267.44 400.87 271.25 392.47 275.38 C384.07 279.52 375.16 284.37 367.72 288.67 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M470.29 768.25 C479.42 769.60 486.97 770.30 495.38 770.96 C503.80 771.62 512.25 771.99 520.79 772.20 C529.32 772.41 537.88 772.36 546.58 772.25 C555.28 772.14 564.02 771.82 573.01 771.53 C582.01 771.23 591.13 770.94 600.55 770.48 C609.97 770.02 619.79 769.84 629.53 768.74 C639.28 767.64 649.41 766.36 659.03 763.88 C668.66 761.39 678.24 758.00 687.30 753.86 C696.37 749.71 705.13 744.60 713.43 738.99 C721.72 733.39 729.64 727.00 737.09 720.24 C744.54 713.47 751.57 706.08 758.12 698.38 C764.67 690.68 770.33 683.40 776.38 674.04 C782.43 664.68 789.19 653.19 794.40 642.21 C799.62 631.23 804.13 619.75 807.66 608.15 C811.20 596.54 813.92 584.54 815.62 572.59 C817.32 560.63 818.07 548.41 817.86 536.42 C817.65 524.44 816.36 512.36 814.34 500.67 C812.32 488.98 809.21 477.43 805.73 466.29 C802.25 455.16 797.93 444.34 793.46 433.84 C788.99 423.34 784.03 413.21 778.92 403.31 C773.81 393.42 768.43 383.83 762.80 374.46 C757.18 365.09 751.33 355.97 745.17 347.09 C739.01 338.21 732.60 329.56 725.84 321.19 C719.08 312.82 711.35 304.09 704.61 296.87 C697.87 289.66 692.02 284.05 685.41 277.91 C678.80 271.77 671.96 265.85 664.94 260.03 C657.93 254.21 650.71 248.60 643.34 242.98 C635.96 237.35 628.41 231.89 620.67 226.27 C612.92 220.65 605.07 214.98 596.87 209.28 C588.67 203.58 580.33 197.43 571.46 192.08 C562.59 186.73 553.30 181.31 543.66 177.17 C534.01 173.03 523.83 169.66 513.58 167.23 C503.33 164.80 492.71 163.34 482.15 162.58 C471.60 161.82 460.85 161.95 450.23 162.69 C439.61 163.43 428.93 164.94 418.45 167.03 C407.96 169.12 398.49 171.43 387.32 175.22 C376.15 179.00 362.99 184.06 351.44 189.72 C339.88 195.38 328.57 201.91 317.97 209.17 C307.36 216.44 297.17 224.56 287.82 233.31 C278.48 242.05 269.73 251.64 261.91 261.66 C254.09 271.68 247.08 282.47 240.92 293.44 C234.75 304.40 229.55 315.96 224.94 327.44 C220.33 338.93 216.61 350.71 213.24 362.35 C209.87 373.98 207.15 385.66 204.74 397.25 C202.32 408.84 200.35 420.38 198.74 431.90 C197.13 443.42 195.90 454.90 195.09 466.38 C194.29 477.86 193.87 489.33 193.93 500.79 C194.00 512.25 194.32 523.75 195.47 535.15 C196.63 546.54 198.25 558.01 200.87 569.15 C203.48 580.29 206.91 591.38 211.15 601.99 C215.39 612.59 220.48 623.00 226.30 632.76 C232.12 642.52 238.82 651.91 246.09 660.54 C253.36 669.17 261.49 677.22 269.92 684.54 C278.36 691.86 287.52 698.41 296.70 704.44 C305.87 710.47 315.48 715.75 324.97 720.72 C334.47 725.69 344.09 730.09 353.67 734.26 C363.25 738.43 372.82 742.21 382.43 745.71 C392.05 749.22 401.65 752.42 411.35 755.27 C421.04 758.13 430.77 760.69 440.60 762.85 C450.42 765.01 461.16 766.90 470.29 768.25 Z`

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
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.InsideAngle = 112.000000
	__Parameter__000000_Reference.SideLength = 85.000000
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
	__Parameter__000000_Reference.NbOfBeatsInTheme = 16
	__Parameter__000000_Reference.FirstVoiceShiftX = -0.070000
	__Parameter__000000_Reference.FirstVoiceShiftY = 1.830000
	__Parameter__000000_Reference.PitchDifference = 12
	__Parameter__000000_Reference.BeatsPerSecond = 6.050000
	__Parameter__000000_Reference.Level = 2.200000
	__Parameter__000000_Reference.IsMinor = true
	__Parameter__000000_Reference.ThemeBinaryEncoding = 50115
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
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 85.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 85.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 139.671216
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 18.841885
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 85.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 85.000000
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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 136.340447
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -15.375813
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 156.878974
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 19.295926
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 103.107187
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 164.013302
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 122.491701
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 126.323839
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 90.100000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 136.340447
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 90.100000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -15.375813
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 136.340447
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 179.680121
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -15.375813
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -41.048972
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 42.724456
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 4.818257
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 32.871580
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 93.935495
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 136.340447
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -15.375813
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 26.648140
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -76.151080
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
