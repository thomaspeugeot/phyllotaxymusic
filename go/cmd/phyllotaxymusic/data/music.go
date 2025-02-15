package main

import (
	"time"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"

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
	__Axis__000001_Construction_Axis.AngleDegree = 100.385138
	__Axis__000001_Construction_Axis.Length = 249.968377
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -45.060303
	__Axis__000001_Construction_Axis.EndY = 245.873460
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
	__Axis__000002_Initial_Axis.AngleDegree = 79.614862
	__Axis__000002_Initial_Axis.Length = 693.339645
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
	__Axis__000004_Rotated_Axis.Length = 693.339645
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -22.530152
	__Bezier__000005_Growth_Curve_Seed.StartY = 122.936730
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 57.536375
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 137.610221
	__Bezier__000005_Growth_Curve_Seed.EndX = 89.101595
	__Bezier__000005_Growth_Curve_Seed.EndY = 270.460806
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 9.035069
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 255.787315
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
	__Circle__000001_Construction_Circle.CenterX = -22.530152
	__Circle__000001_Construction_Circle.CenterY = 122.936730
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 268.323797
	__Circle__000005_Rotated_Next_Circle.CenterY = 49.174692
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

	__ExportToMusicxml__000000_Singloton.Name = `Singloton`

	__FrontCurve__000000_Non_Rotated_0_.Name = `Non Rotated 0 `
	__FrontCurve__000000_Non_Rotated_0_.Path = `M813.23 565.26 C816.17 554.86 817.78 547.84 819.91 538.98 C822.03 530.12 823.93 521.20 825.97 512.10 C828.00 502.99 829.82 493.84 832.13 484.35 C834.44 474.85 836.44 465.37 839.82 455.12 C843.20 444.86 846.37 434.73 852.39 422.82 C858.42 410.92 868.84 397.49 875.98 383.71 C883.12 369.92 891.85 354.20 895.26 340.13 C898.67 326.07 897.83 312.45 896.43 299.30 C895.03 286.16 891.20 273.51 886.87 261.26 C882.54 249.01 876.78 237.17 870.47 225.80 C864.16 214.43 856.83 203.47 849.01 193.02 C841.19 182.57 834.67 174.19 823.53 163.09 C812.40 151.99 796.91 137.37 782.19 126.40 C767.48 115.42 751.63 105.43 735.25 97.26 C718.87 89.08 701.48 82.17 683.92 77.33 C666.36 72.48 647.97 69.29 629.90 68.17 C611.82 67.05 593.24 68.09 575.46 70.62 C557.69 73.15 539.98 78.10 523.25 83.37 C506.51 88.63 490.46 95.63 475.08 102.19 C459.70 108.74 445.18 115.88 430.96 122.70 C416.75 129.52 403.20 136.29 389.80 143.13 C376.39 149.97 363.41 156.68 350.56 163.73 C337.70 170.77 325.11 177.84 312.69 185.38 C300.26 192.92 286.36 201.98 276.02 208.96 C265.68 215.94 259.04 220.98 250.63 227.27 C242.21 233.55 233.96 240.08 225.53 246.69 C217.10 253.30 208.84 260.12 200.04 266.93 C191.23 273.73 182.71 280.76 172.69 287.50 C162.67 294.25 152.94 301.08 139.93 307.40 C126.92 313.73 109.63 318.39 94.65 325.46 C79.66 332.52 62.26 340.06 50.04 349.81 C37.83 359.56 29.19 371.76 21.34 383.98 C13.50 396.20 7.88 409.64 2.96 423.11 C-1.96 436.58 -5.46 450.68 -8.18 464.79 C-10.90 478.90 -12.52 493.36 -13.38 507.76 C-14.23 522.16 -14.72 533.87 -13.31 551.19 C-11.90 568.51 -9.36 591.89 -4.93 611.69 C-0.49 631.48 5.57 651.26 13.31 669.96 C21.05 688.66 30.50 707.03 41.52 723.89 C52.54 740.75 65.38 756.85 79.42 771.11 C93.46 785.37 109.39 798.27 125.77 809.43 C142.14 820.60 160.13 829.82 177.66 838.08 C195.18 846.34 213.41 852.77 230.90 858.99 C248.38 865.22 265.67 870.36 282.56 875.43 C299.46 880.50 315.88 885.09 332.27 889.42 C348.65 893.74 364.70 897.79 380.88 901.39 C397.05 904.98 413.10 908.28 429.32 911.00 C445.54 913.71 461.81 916.42 478.20 917.68 C494.60 918.94 511.29 919.70 527.68 918.54 C544.06 917.38 560.66 914.85 576.52 910.72 C592.38 906.58 608.18 900.94 622.82 893.73 C637.46 886.53 651.67 877.64 664.35 867.48 C677.03 857.33 688.73 845.32 698.91 832.81 C709.10 820.29 717.71 806.09 725.45 792.38 C733.20 778.67 739.35 764.18 745.38 750.54 C751.41 736.91 756.48 723.49 761.61 710.57 C766.75 697.65 771.50 685.29 776.20 673.03 C780.90 660.76 785.45 648.93 789.79 636.99 C794.14 625.04 798.37 613.30 802.28 601.34 C806.18 589.39 810.29 575.65 813.23 565.26 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M276.02 208.96 C265.68 215.94 259.04 220.98 250.63 227.27 C242.21 233.55 233.96 240.08 225.53 246.69 C217.10 253.30 208.84 260.12 200.04 266.93 C191.23 273.73 182.71 280.76 172.69 287.50 C162.67 294.25 152.94 301.08 139.93 307.40 C126.92 313.73 109.63 318.39 94.65 325.46 C79.66 332.52 62.26 340.06 50.04 349.81 C37.83 359.56 29.19 371.76 21.34 383.98 C13.50 396.20 7.88 409.64 2.96 423.11 C-1.96 436.58 -5.46 450.68 -8.18 464.79 C-10.90 478.90 -12.52 493.36 -13.38 507.76 C-14.23 522.16 -14.72 533.87 -13.31 551.19 C-11.90 568.51 -9.36 591.89 -4.93 611.69 C-0.49 631.48 5.57 651.26 13.31 669.96 C21.05 688.66 30.50 707.03 41.52 723.89 C52.54 740.75 65.38 756.85 79.42 771.11 C93.46 785.37 109.39 798.27 125.77 809.43 C142.14 820.60 160.13 829.82 177.66 838.08 C195.18 846.34 213.41 852.77 230.90 858.99 C248.38 865.22 265.67 870.36 282.56 875.43 C299.46 880.50 315.88 885.09 332.27 889.42 C348.65 893.74 364.70 897.79 380.88 901.39 C397.05 904.98 413.10 908.28 429.32 911.00 C445.54 913.71 464.18 916.11 478.20 917.68 C492.23 919.25 501.64 919.73 513.48 920.40 C525.31 921.07 537.18 921.40 549.23 921.72 C561.28 922.04 573.34 922.02 585.78 922.32 C598.22 922.62 610.58 922.51 623.87 923.52 C637.17 924.54 650.29 925.25 665.53 928.42 C680.77 931.58 698.15 938.75 715.33 942.52 C732.50 946.30 751.79 951.24 768.57 951.07 C785.36 950.91 801.01 946.50 816.04 941.53 C831.06 936.55 845.18 929.15 858.74 921.23 C872.30 913.31 885.16 903.94 897.40 894.02 C909.64 884.10 921.23 873.16 932.16 861.73 C943.09 850.30 951.81 840.87 962.97 825.43 C974.13 810.00 988.67 788.73 999.15 769.11 C1009.62 749.50 1018.76 728.77 1025.81 707.73 C1032.85 686.70 1038.28 664.74 1041.43 642.90 C1044.57 621.05 1045.74 598.55 1044.68 576.67 C1043.61 554.78 1040.12 532.66 1035.04 511.59 C1029.96 490.51 1022.27 469.83 1014.20 450.22 C1006.13 430.60 996.21 411.91 986.61 393.90 C977.01 375.89 966.75 358.84 956.62 342.15 C946.50 325.46 936.31 309.46 925.87 293.75 C915.43 278.04 904.96 262.80 894.01 247.90 C883.06 232.99 871.91 218.44 860.16 204.31 C848.41 190.17 836.53 176.07 823.53 163.09 C810.54 150.10 796.91 137.37 782.19 126.40 C767.48 115.42 751.63 105.43 735.25 97.26 C718.87 89.08 701.48 82.17 683.92 77.33 C666.36 72.48 647.97 69.29 629.90 68.17 C611.82 67.05 593.24 68.09 575.46 70.62 C557.69 73.15 539.98 78.10 523.25 83.37 C506.51 88.63 490.46 95.63 475.08 102.19 C459.70 108.74 445.18 115.88 430.96 122.70 C416.75 129.52 403.20 136.29 389.80 143.13 C376.39 149.97 363.41 156.68 350.56 163.73 C337.70 170.77 325.11 177.84 312.69 185.38 C300.26 192.92 286.36 201.98 276.02 208.96 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M478.20 917.68 C492.23 919.25 501.64 919.73 513.48 920.40 C525.31 921.07 537.18 921.40 549.23 921.72 C561.28 922.04 573.34 922.02 585.78 922.32 C598.22 922.62 610.58 922.51 623.87 923.52 C637.17 924.54 650.29 925.25 665.53 928.42 C680.77 931.58 698.15 938.75 715.33 942.52 C732.50 946.30 751.79 951.24 768.57 951.07 C785.36 950.91 801.01 946.50 816.04 941.53 C831.06 936.55 845.18 929.15 858.74 921.23 C872.30 913.31 885.16 903.94 897.40 894.02 C909.64 884.10 921.23 873.16 932.16 861.73 C943.09 850.30 951.81 840.87 962.97 825.43 C974.13 810.00 988.67 788.73 999.15 769.11 C1009.62 749.50 1018.76 728.77 1025.81 707.73 C1032.85 686.70 1038.28 664.74 1041.43 642.90 C1044.57 621.05 1045.74 598.55 1044.68 576.67 C1043.61 554.78 1040.12 532.66 1035.04 511.59 C1029.96 490.51 1022.27 469.83 1014.20 450.22 C1006.13 430.60 996.21 411.91 986.61 393.90 C977.01 375.89 966.75 358.84 956.62 342.15 C946.50 325.46 936.31 309.46 925.87 293.75 C915.43 278.04 904.96 262.80 894.01 247.90 C883.06 232.99 871.91 218.44 860.16 204.31 C848.41 190.17 834.31 174.59 823.53 163.09 C812.75 151.58 805.08 144.38 795.48 135.28 C785.87 126.18 775.98 117.36 765.92 108.48 C755.86 99.60 745.53 91.03 735.10 82.01 C724.67 72.99 713.99 64.38 703.33 54.36 C692.66 44.35 681.90 34.72 671.11 21.91 C660.33 9.09 650.50 -8.11 638.61 -22.54 C626.72 -36.97 613.97 -53.61 599.78 -64.69 C585.59 -75.76 569.49 -82.82 553.48 -88.98 C537.48 -95.14 520.57 -98.80 503.76 -101.65 C486.94 -104.51 469.70 -105.74 452.59 -106.10 C435.48 -106.46 418.19 -105.55 401.10 -103.79 C384.01 -102.02 370.17 -100.38 350.06 -95.52 C329.94 -90.66 302.93 -83.38 280.42 -74.63 C257.90 -65.88 235.71 -55.30 214.96 -43.01 C194.22 -30.71 174.14 -16.53 155.94 -0.83 C137.74 14.89 120.63 32.54 105.67 51.28 C90.67 70.02 77.39 90.59 66.08 111.42 C54.79 132.21 45.84 154.46 38.00 176.19 C30.18 197.92 24.46 220.22 19.06 241.84 C13.67 263.45 9.51 284.84 5.59 305.89 C1.67 326.94 -1.63 347.57 -4.44 368.15 C-7.26 388.73 -9.63 409.02 -11.30 429.37 C-12.97 449.72 -14.15 469.96 -14.48 490.27 C-14.82 510.57 -14.90 530.95 -13.31 551.19 C-11.72 571.43 -9.36 591.89 -4.93 611.69 C-0.49 631.48 5.57 651.26 13.31 669.96 C21.05 688.66 30.50 707.03 41.52 723.89 C52.54 740.75 65.38 756.85 79.42 771.11 C93.46 785.37 109.39 798.27 125.77 809.43 C142.14 820.60 160.13 829.82 177.66 838.08 C195.18 846.34 213.41 852.77 230.90 858.99 C248.38 865.22 265.67 870.36 282.56 875.43 C299.46 880.50 315.88 885.09 332.27 889.42 C348.65 893.74 364.70 897.79 380.88 901.39 C397.05 904.98 413.10 908.28 429.32 911.00 C445.54 913.71 464.18 916.11 478.20 917.68 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M823.53 163.09 C812.75 151.58 805.08 144.38 795.48 135.28 C785.87 126.18 775.98 117.36 765.92 108.48 C755.86 99.60 745.53 91.03 735.10 82.01 C724.67 72.99 713.99 64.38 703.33 54.36 C692.66 44.35 681.90 34.72 671.11 21.91 C660.33 9.09 650.50 -8.11 638.61 -22.54 C626.72 -36.97 613.97 -53.61 599.78 -64.69 C585.59 -75.76 569.49 -82.82 553.48 -88.98 C537.48 -95.14 520.57 -98.80 503.76 -101.65 C486.94 -104.51 469.70 -105.74 452.59 -106.10 C435.48 -106.46 418.19 -105.55 401.10 -103.79 C384.01 -102.02 370.17 -100.38 350.06 -95.52 C329.94 -90.66 302.93 -83.38 280.42 -74.63 C257.90 -65.88 235.71 -55.30 214.96 -43.01 C194.22 -30.71 174.14 -16.53 155.94 -0.83 C137.74 14.89 120.63 32.54 105.67 51.28 C90.67 70.02 77.39 90.59 66.08 111.42 C54.79 132.21 45.84 154.46 38.00 176.19 C30.18 197.92 24.46 220.22 19.06 241.84 C13.67 263.45 9.51 284.84 5.59 305.89 C1.67 326.94 -1.63 347.57 -4.44 368.15 C-7.26 388.73 -9.63 409.02 -11.30 429.37 C-12.97 449.72 -14.15 469.96 -14.48 490.27 C-14.82 510.57 -14.13 533.79 -13.31 551.19 C-12.49 568.59 -11.17 580.14 -9.56 594.65 C-7.95 609.17 -5.87 623.65 -3.66 638.26 C-1.44 652.88 1.23 667.43 3.75 682.36 C6.27 697.29 9.29 712.05 11.45 727.84 C13.60 743.63 16.13 759.17 16.66 777.10 C17.18 795.04 14.17 815.64 14.58 835.45 C14.99 855.26 14.66 877.28 19.10 895.98 C23.53 914.67 32.08 931.54 41.18 947.62 C50.29 963.70 61.75 978.44 73.70 992.46 C85.65 1006.47 99.02 1019.50 112.90 1031.73 C126.77 1043.96 141.64 1055.31 156.95 1065.82 C172.25 1076.33 184.79 1084.66 204.72 1094.80 C224.66 1104.93 251.94 1117.92 276.55 1126.62 C301.17 1135.31 326.78 1142.32 352.41 1146.98 C378.04 1151.63 404.42 1154.33 430.34 1154.56 C456.26 1154.79 482.58 1152.73 507.93 1148.33 C533.28 1143.93 558.53 1136.86 582.44 1128.16 C606.35 1119.46 629.49 1108.00 651.42 1096.11 C673.35 1084.22 694.07 1070.41 714.05 1056.81 C734.03 1043.22 752.90 1028.88 771.31 1014.55 C789.72 1000.21 807.34 985.70 824.49 970.81 C841.65 955.92 858.22 940.86 874.22 925.23 C890.21 909.59 905.69 893.62 920.49 876.99 C935.28 860.35 949.86 843.41 962.97 825.43 C976.08 807.46 988.67 788.73 999.15 769.11 C1009.62 749.50 1018.76 728.77 1025.81 707.73 C1032.85 686.70 1038.28 664.74 1041.43 642.90 C1044.57 621.05 1045.74 598.55 1044.68 576.67 C1043.61 554.78 1040.12 532.66 1035.04 511.59 C1029.96 490.51 1022.27 469.83 1014.20 450.22 C1006.13 430.60 996.21 411.91 986.61 393.90 C977.01 375.89 966.75 358.84 956.62 342.15 C946.50 325.46 936.31 309.46 925.87 293.75 C915.43 278.04 904.96 262.80 894.01 247.90 C883.06 232.99 871.91 218.44 860.16 204.31 C848.41 190.17 834.31 174.59 823.53 163.09 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M-13.31 551.19 C-12.49 568.59 -11.17 580.14 -9.56 594.65 C-7.95 609.17 -5.87 623.65 -3.66 638.26 C-1.44 652.88 1.23 667.43 3.75 682.36 C6.27 697.29 9.29 712.05 11.45 727.84 C13.60 743.63 16.13 759.17 16.66 777.10 C17.18 795.04 14.17 815.64 14.58 835.45 C14.99 855.26 14.66 877.28 19.10 895.98 C23.53 914.67 32.08 931.54 41.18 947.62 C50.29 963.70 61.75 978.44 73.70 992.46 C85.65 1006.47 99.02 1019.50 112.90 1031.73 C126.77 1043.96 141.64 1055.31 156.95 1065.82 C172.25 1076.33 184.79 1084.66 204.72 1094.80 C224.66 1104.93 251.94 1117.92 276.55 1126.62 C301.17 1135.31 326.78 1142.32 352.41 1146.98 C378.04 1151.63 404.42 1154.33 430.34 1154.56 C456.26 1154.79 482.58 1152.73 507.93 1148.33 C533.28 1143.93 558.53 1136.86 582.44 1128.16 C606.35 1119.46 629.49 1108.00 651.42 1096.11 C673.35 1084.22 694.07 1070.41 714.05 1056.81 C734.03 1043.22 752.90 1028.88 771.31 1014.55 C789.72 1000.21 807.34 985.70 824.49 970.81 C841.65 955.92 858.22 940.86 874.22 925.23 C890.21 909.59 905.69 893.62 920.49 876.99 C935.28 860.35 951.26 840.51 962.97 825.43 C974.68 810.36 981.78 799.77 990.72 786.51 C999.67 773.26 1008.19 759.69 1016.64 745.91 C1025.09 732.12 1033.10 718.06 1041.42 703.80 C1049.74 689.54 1057.50 675.04 1066.55 660.33 C1075.60 645.61 1084.13 630.82 1095.69 615.51 C1107.26 600.20 1123.17 585.14 1135.93 568.47 C1148.69 551.79 1163.39 533.90 1172.25 515.44 C1181.12 496.98 1185.61 477.22 1189.12 457.71 C1192.63 438.19 1193.41 418.16 1193.29 398.37 C1193.18 378.57 1191.29 358.61 1188.43 338.94 C1185.58 319.27 1181.34 299.63 1176.17 280.36 C1171.00 261.10 1166.59 245.54 1157.42 223.34 C1148.26 201.13 1135.11 171.47 1121.18 147.14 C1107.24 122.81 1091.34 99.15 1073.80 77.34 C1056.27 55.54 1036.80 34.78 1015.98 16.30 C995.17 -2.19 972.47 -19.15 948.91 -33.59 C925.35 -48.03 900.06 -60.30 874.65 -70.35 C849.24 -80.41 822.55 -87.78 796.45 -93.93 C770.35 -100.08 743.87 -103.88 718.05 -107.25 C692.23 -110.62 666.71 -112.61 641.54 -114.18 C616.36 -115.74 591.63 -116.55 567.01 -116.66 C542.39 -116.76 518.07 -116.24 493.83 -114.80 C469.58 -113.36 445.50 -111.21 421.54 -107.99 C397.58 -104.78 373.58 -101.08 350.06 -95.52 C326.53 -89.96 302.93 -83.38 280.42 -74.63 C257.90 -65.88 235.71 -55.30 214.96 -43.01 C194.22 -30.71 174.14 -16.53 155.94 -0.83 C137.74 14.89 120.63 32.54 105.67 51.28 C90.67 70.02 77.39 90.59 66.08 111.42 C54.79 132.21 45.84 154.46 38.00 176.19 C30.18 197.92 24.46 220.22 19.06 241.84 C13.67 263.45 9.51 284.84 5.59 305.89 C1.67 326.94 -1.63 347.57 -4.44 368.15 C-7.26 388.73 -9.63 409.02 -11.30 429.37 C-12.97 449.72 -14.15 469.96 -14.48 490.27 C-14.82 510.57 -14.13 533.79 -13.31 551.19 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M962.97 825.43 C974.68 810.36 981.78 799.77 990.72 786.51 C999.67 773.26 1008.19 759.69 1016.64 745.91 C1025.09 732.12 1033.10 718.06 1041.42 703.80 C1049.74 689.54 1057.50 675.04 1066.55 660.33 C1075.60 645.61 1084.13 630.82 1095.69 615.51 C1107.26 600.20 1123.17 585.14 1135.93 568.47 C1148.69 551.79 1163.39 533.90 1172.25 515.44 C1181.12 496.98 1185.61 477.22 1189.12 457.71 C1192.63 438.19 1193.41 418.16 1193.29 398.37 C1193.18 378.57 1191.29 358.61 1188.43 338.94 C1185.58 319.27 1181.34 299.63 1176.17 280.36 C1171.00 261.10 1166.59 245.54 1157.42 223.34 C1148.26 201.13 1135.11 171.47 1121.18 147.14 C1107.24 122.81 1091.34 99.15 1073.80 77.34 C1056.27 55.54 1036.80 34.78 1015.98 16.30 C995.17 -2.19 972.47 -19.15 948.91 -33.59 C925.35 -48.03 900.06 -60.30 874.65 -70.35 C849.24 -80.41 822.55 -87.78 796.45 -93.93 C770.35 -100.08 743.87 -103.88 718.05 -107.25 C692.23 -110.62 666.71 -112.61 641.54 -114.18 C616.36 -115.74 591.63 -116.55 567.01 -116.66 C542.39 -116.76 518.07 -116.24 493.83 -114.80 C469.58 -113.36 445.50 -111.21 421.54 -107.99 C397.58 -104.78 370.36 -99.70 350.06 -95.52 C329.75 -91.34 316.42 -87.58 299.70 -82.91 C282.99 -78.24 266.41 -73.00 249.76 -67.50 C233.12 -61.99 216.65 -55.93 199.85 -49.87 C183.05 -43.81 166.54 -37.15 148.96 -31.11 C131.37 -25.07 114.15 -18.58 94.34 -13.64 C74.52 -8.68 51.60 -6.76 30.07 -1.41 C8.42 4.37 -17.10 9.36 -36.78 18.57 C-56.44 27.80 -73.58 40.97 -89.76 54.61 C-105.95 68.23 -120.35 84.08 -133.87 100.35 C-147.40 116.62 -159.63 134.21 -170.90 152.22 C-182.16 170.23 -192.30 189.16 -201.44 208.42 C-210.59 227.68 -217.78 243.39 -225.75 267.79 C-233.72 292.20 -243.64 325.41 -249.26 354.85 C-254.89 384.29 -258.46 414.54 -259.49 444.43 C-260.51 474.33 -259.27 504.71 -255.41 534.23 C-251.55 563.75 -245.15 593.34 -236.35 621.56 C-227.55 649.77 -215.89 677.47 -202.61 703.51 C-189.32 729.56 -173.21 754.38 -156.64 777.83 C-140.07 801.29 -121.56 823.18 -103.19 844.24 C-84.82 865.31 -65.67 885.07 -46.43 904.22 C-27.18 923.37 -7.68 941.60 12.27 959.13 C32.23 976.66 52.45 993.45 73.28 1009.39 C94.11 1025.34 115.37 1040.56 137.27 1054.79 C159.18 1069.03 181.51 1082.83 204.72 1094.80 C227.94 1106.77 251.94 1117.92 276.55 1126.62 C301.17 1135.31 326.78 1142.32 352.41 1146.98 C378.04 1151.63 404.42 1154.33 430.34 1154.56 C456.26 1154.79 482.58 1152.73 507.93 1148.33 C533.28 1143.93 558.53 1136.86 582.44 1128.16 C606.35 1119.46 629.49 1108.00 651.42 1096.11 C673.35 1084.22 694.07 1070.41 714.05 1056.81 C734.03 1043.22 752.90 1028.88 771.31 1014.55 C789.72 1000.21 807.34 985.70 824.49 970.81 C841.65 955.92 858.22 940.86 874.22 925.23 C890.21 909.59 905.69 893.62 920.49 876.99 C935.28 860.35 951.26 840.51 962.97 825.43 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M350.06 -95.52 C329.75 -91.34 316.42 -87.58 299.70 -82.91 C282.99 -78.24 266.41 -73.00 249.76 -67.50 C233.12 -61.99 216.65 -55.93 199.85 -49.87 C183.05 -43.81 166.54 -37.15 148.96 -31.11 C131.37 -25.07 114.15 -18.58 94.34 -13.64 C74.52 -8.68 51.60 -6.76 30.07 -1.41 C8.42 4.37 -17.10 9.36 -36.78 18.57 C-56.44 27.80 -73.58 40.97 -89.76 54.61 C-105.95 68.23 -120.35 84.08 -133.87 100.35 C-147.40 116.62 -159.63 134.21 -170.90 152.22 C-182.16 170.23 -192.30 189.16 -201.44 208.42 C-210.59 227.68 -217.78 243.39 -225.75 267.79 C-233.72 292.20 -243.64 325.41 -249.26 354.85 C-254.89 384.29 -258.46 414.54 -259.49 444.43 C-260.51 474.33 -259.27 504.71 -255.41 534.23 C-251.55 563.75 -245.15 593.34 -236.35 621.56 C-227.55 649.77 -215.89 677.47 -202.61 703.51 C-189.32 729.56 -173.21 754.38 -156.64 777.83 C-140.07 801.29 -121.56 823.18 -103.19 844.24 C-84.82 865.31 -65.67 885.07 -46.43 904.22 C-27.18 923.37 -7.68 941.60 12.27 959.13 C32.23 976.66 52.45 993.45 73.28 1009.39 C94.11 1025.34 115.37 1040.56 137.27 1054.79 C159.18 1069.03 185.11 1083.97 204.72 1094.80 C224.33 1105.62 237.93 1111.88 254.95 1119.74 C271.96 1127.60 289.28 1134.89 306.82 1141.96 C324.36 1149.03 342.15 1155.52 360.20 1162.16 C378.25 1168.80 396.48 1174.75 415.13 1181.82 C433.78 1188.89 452.45 1195.32 472.07 1204.55 C491.70 1213.78 511.74 1227.24 532.90 1237.22 C554.06 1247.20 576.70 1258.79 599.02 1264.42 C621.34 1270.05 644.28 1271.04 666.80 1270.98 C689.31 1270.93 711.92 1267.98 734.11 1264.08 C756.29 1260.17 778.36 1254.37 799.94 1247.55 C821.51 1240.74 842.82 1232.45 863.56 1223.19 C884.29 1213.94 900.98 1206.22 924.35 1192.03 C947.72 1177.84 978.77 1157.90 1003.77 1138.05 C1028.76 1118.21 1052.67 1096.32 1074.33 1072.96 C1095.98 1049.59 1116.16 1024.29 1133.70 997.87 C1151.23 971.46 1166.81 943.25 1179.56 914.47 C1192.31 885.70 1202.46 855.38 1210.20 825.22 C1217.93 795.05 1222.62 763.91 1225.96 733.48 C1229.30 703.06 1230.04 672.50 1230.24 642.66 C1230.44 612.81 1229.08 583.40 1227.14 554.41 C1225.19 525.41 1222.33 496.91 1218.57 468.68 C1214.80 440.44 1210.23 412.58 1204.53 384.99 C1198.83 357.41 1192.23 330.10 1184.38 303.15 C1176.53 276.21 1167.96 249.34 1157.42 223.34 C1146.89 197.33 1135.11 171.47 1121.18 147.14 C1107.24 122.81 1091.34 99.15 1073.80 77.34 C1056.27 55.54 1036.80 34.78 1015.98 16.30 C995.17 -2.19 972.47 -19.15 948.91 -33.59 C925.35 -48.03 900.06 -60.30 874.65 -70.35 C849.24 -80.41 822.55 -87.78 796.45 -93.93 C770.35 -100.08 743.87 -103.88 718.05 -107.25 C692.23 -110.62 666.71 -112.61 641.54 -114.18 C616.36 -115.74 591.63 -116.55 567.01 -116.66 C542.39 -116.76 518.07 -116.24 493.83 -114.80 C469.58 -113.36 445.50 -111.21 421.54 -107.99 C397.58 -104.78 370.36 -99.70 350.06 -95.52 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M204.72 1094.80 C224.33 1105.62 237.93 1111.88 254.95 1119.74 C271.96 1127.60 289.28 1134.89 306.82 1141.96 C324.36 1149.03 342.15 1155.52 360.20 1162.16 C378.25 1168.80 396.48 1174.75 415.13 1181.82 C433.78 1188.89 452.45 1195.32 472.07 1204.55 C491.70 1213.78 511.74 1227.24 532.90 1237.22 C554.06 1247.20 576.70 1258.79 599.02 1264.42 C621.34 1270.05 644.28 1271.04 666.80 1270.98 C689.31 1270.93 711.92 1267.98 734.11 1264.08 C756.29 1260.17 778.36 1254.37 799.94 1247.55 C821.51 1240.74 842.82 1232.45 863.56 1223.19 C884.29 1213.94 900.98 1206.22 924.35 1192.03 C947.72 1177.84 978.77 1157.90 1003.77 1138.05 C1028.76 1118.21 1052.67 1096.32 1074.33 1072.96 C1095.98 1049.59 1116.16 1024.29 1133.70 997.87 C1151.23 971.46 1166.81 943.25 1179.56 914.47 C1192.31 885.70 1202.46 855.38 1210.20 825.22 C1217.93 795.05 1222.62 763.91 1225.96 733.48 C1229.30 703.06 1230.04 672.50 1230.24 642.66 C1230.44 612.81 1229.08 583.40 1227.14 554.41 C1225.19 525.41 1222.33 496.91 1218.57 468.68 C1214.80 440.44 1210.23 412.58 1204.53 384.99 C1198.83 357.41 1192.23 330.10 1184.38 303.15 C1176.53 276.21 1165.84 245.88 1157.42 223.34 C1149.00 200.80 1142.27 186.20 1133.86 167.92 C1125.45 149.63 1116.39 131.61 1106.97 113.64 C1097.55 95.67 1087.50 78.01 1077.32 60.10 C1067.13 42.18 1056.31 24.71 1045.86 6.16 C1035.41 -12.39 1024.48 -30.43 1014.61 -51.19 C1004.75 -71.94 997.28 -96.14 986.66 -118.37 C976.04 -140.60 965.31 -164.98 950.90 -184.57 C936.50 -204.17 918.57 -220.61 900.25 -235.91 C881.93 -251.22 861.64 -264.34 841.01 -276.42 C820.38 -288.49 798.58 -298.99 776.47 -308.35 C754.37 -317.71 731.48 -325.72 708.39 -332.57 C685.30 -339.42 666.54 -344.72 637.93 -349.45 C609.32 -354.18 570.57 -359.59 536.75 -360.94 C502.93 -362.29 468.58 -361.27 435.00 -357.54 C401.42 -353.82 367.69 -347.54 335.27 -338.60 C302.84 -329.65 270.75 -317.94 240.47 -303.86 C210.19 -289.78 180.90 -272.73 153.60 -254.11 C126.29 -235.50 100.68 -214.05 76.63 -192.17 C52.56 -170.30 30.40 -146.52 9.17 -122.84 C-12.13 -99.14 -31.99 -74.59 -51.08 -49.75 C-70.08 -24.76 -87.83 0.50 -104.52 26.12 C-121.20 51.65 -136.94 77.45 -151.55 103.80 C-166.16 130.13 -179.82 156.90 -192.18 184.24 C-204.55 211.56 -216.23 239.36 -225.75 267.79 C-235.26 296.23 -243.64 325.41 -249.26 354.85 C-254.89 384.29 -258.46 414.54 -259.49 444.43 C-260.51 474.33 -259.27 504.71 -255.41 534.23 C-251.55 563.75 -245.15 593.34 -236.35 621.56 C-227.55 649.77 -215.89 677.47 -202.61 703.51 C-189.32 729.56 -173.21 754.38 -156.64 777.83 C-140.07 801.29 -121.56 823.18 -103.19 844.24 C-84.82 865.31 -65.67 885.07 -46.43 904.22 C-27.18 923.37 -7.68 941.60 12.27 959.13 C32.23 976.66 52.45 993.45 73.28 1009.39 C94.11 1025.34 115.37 1040.56 137.27 1054.79 C159.18 1069.03 185.11 1083.97 204.72 1094.80 Z`

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
	__Parameter__000000_Reference.InsideAngle = 95.000000
	__Parameter__000000_Reference.ShiftToNearestCircle = 3
	__Parameter__000000_Reference.SideLength = 185.000000
	__Parameter__000000_Reference.StackWidth = 1
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 8
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.440000
	__Parameter__000000_Reference.SpiralBezierStrength = 0.800000
	__Parameter__000000_Reference.NbInterpolationPoints = 12
	__Parameter__000000_Reference.FkeySizeRatio = 0.001430
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.400000
	__Parameter__000000_Reference.PitchHeight = 0.087000
	__Parameter__000000_Reference.NbPitchLines = 169
	__Parameter__000000_Reference.BeatLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbBeatLines = 300
	__Parameter__000000_Reference.NbOfBeatsInTheme = 16
	__Parameter__000000_Reference.FirstVoiceShiftX = 0.060000
	__Parameter__000000_Reference.FirstVoiceShiftY = 1.490000
	__Parameter__000000_Reference.PitchDifference = 12
	__Parameter__000000_Reference.BeatsPerSecond = 4.000000
	__Parameter__000000_Reference.Level = 2.200000
	__Parameter__000000_Reference.IsMinor = true
	__Parameter__000000_Reference.ThemeBinaryEncoding = 65
	__Parameter__000000_Reference.OriginX = 40.000000
	__Parameter__000000_Reference.OriginY = 950.000000
	__Parameter__000000_Reference.SpiralOriginX = 500.000000
	__Parameter__000000_Reference.SpiralOriginY = 500.000000
	__Parameter__000000_Reference.OriginCrossWidth = 800.000000
	__Parameter__000000_Reference.SpiralRadiusRatio = 1.060000
	__Parameter__000000_Reference.ShowSpiralBezierConstruct = false
	__Parameter__000000_Reference.ShowInterpolationPoints = false
	__Parameter__000000_Reference.ActualBeatsTemporalShift = 6

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 185.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.AngleDegree = -79.614862
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 95.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 185.000000
	__Rhombus__000001_Initial_Rhombus.AngleDegree = -79.106605
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 95.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 268.323797
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 49.174692
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 185.000000
	__Rhombus__000002_Rotated_Next_Rhombus.AngleDegree = -79.614862
	__Rhombus__000002_Rotated_Next_Rhombus.InsideAngle = 95.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 185.000000
	__Rhombus__000003_Rotated_Rhombus.AngleDegree = -79.614862
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 95.000000
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
	__ShapeCategory__000000_0_Axes.IsExpanded = false

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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 312.410029
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -64.687007
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 401.063566
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 9.892139
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 322.551300
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 337.104798
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 320.428523
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 206.795764
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 196.100000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 312.410029
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 196.100000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -64.687007
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 312.410029
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 405.633962
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -64.687007
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -175.503928
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 71.643598
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 14.834383
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 76.707519
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 206.927188
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 312.410029
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -64.687007
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 59.615154
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -160.818606
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
	__ExportToMusicxml__000000_Singloton.Parameter = __Parameter__000000_Reference
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
