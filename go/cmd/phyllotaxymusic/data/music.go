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
	__Axis__000001_Construction_Axis.AngleDegree = 97.971943
	__Axis__000001_Construction_Axis.Length = 195.015988
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -27.046409
	__Axis__000001_Construction_Axis.EndY = 193.131374
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
	__Axis__000002_Initial_Axis.Length = 703.073670
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
	__Axis__000004_Rotated_Axis.Length = 703.073670
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -13.523204
	__Bezier__000005_Growth_Curve_Seed.StartY = 96.565687
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 60.553936
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 106.939561
	__Bezier__000005_Growth_Curve_Seed.EndX = 110.863684
	__Bezier__000005_Growth_Curve_Seed.EndY = 212.444511
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 36.786544
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 202.070637
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
	__Circle__000001_Construction_Circle.CenterX = -13.523204
	__Circle__000001_Construction_Circle.CenterY = 96.565687
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 275.820186
	__Circle__000005_Rotated_Next_Circle.CenterY = 38.626275
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M775.57 533.93 C777.29 524.59 778.40 517.02 779.65 508.42 C780.91 499.81 781.96 491.18 783.10 482.32 C784.24 473.46 785.23 464.56 786.51 455.27 C787.79 445.97 789.03 436.60 790.77 426.55 C792.51 416.50 794.67 406.14 796.95 394.95 C799.23 383.75 802.65 371.60 804.45 359.37 C806.26 347.15 808.14 333.95 807.77 321.58 C807.39 309.21 805.42 296.80 802.23 285.14 C799.03 273.49 794.16 262.26 788.58 251.65 C783.00 241.04 776.15 230.96 768.76 221.49 C761.38 212.02 753.04 203.11 744.27 194.83 C735.50 186.55 727.18 179.45 716.16 171.81 C705.15 164.17 691.38 155.50 678.17 149.00 C664.97 142.49 651.05 136.93 636.96 132.80 C622.88 128.67 608.20 125.69 593.67 124.24 C579.13 122.78 564.20 122.72 549.74 124.09 C535.28 125.46 520.73 128.48 506.91 132.45 C493.09 136.43 479.58 142.08 466.80 147.94 C454.02 153.79 441.88 160.83 430.22 167.61 C418.55 174.38 407.56 181.59 396.81 188.58 C386.05 195.58 375.80 202.56 365.69 209.56 C355.57 216.55 345.78 223.46 336.12 230.56 C326.46 237.66 317.02 244.75 307.73 252.16 C298.45 259.57 288.57 267.86 280.42 275.00 C272.26 282.14 265.93 288.15 258.78 295.00 C251.64 301.84 244.66 308.88 237.53 316.05 C230.40 323.23 223.40 330.57 215.98 338.06 C208.56 345.55 201.16 353.15 193.00 360.99 C184.85 368.84 176.22 376.68 167.03 385.11 C157.83 393.54 147.20 402.02 137.84 411.59 C128.48 421.15 118.50 431.43 110.86 442.49 C103.21 453.54 96.85 465.64 91.97 477.92 C87.09 490.19 83.85 503.19 81.60 516.14 C79.35 529.09 78.48 542.43 78.46 555.62 C78.45 568.80 79.59 582.15 81.51 595.25 C83.42 608.35 85.74 620.13 89.97 634.24 C94.20 648.34 100.08 665.20 106.91 679.85 C113.74 694.50 121.80 708.86 130.95 722.13 C140.11 735.41 150.51 748.14 161.82 759.52 C173.14 770.91 185.70 781.41 198.84 790.44 C211.98 799.48 226.28 807.23 240.67 813.72 C255.05 820.21 270.29 825.13 285.13 829.37 C299.98 833.61 315.11 836.46 329.73 839.14 C344.36 841.82 358.77 843.67 372.88 845.46 C386.99 847.25 400.74 848.66 414.40 849.89 C428.06 851.12 441.44 852.14 454.84 852.84 C468.24 853.54 481.49 854.03 494.79 854.09 C508.10 854.16 521.43 854.26 534.68 853.21 C547.94 852.17 561.35 850.75 574.31 847.84 C587.26 844.93 600.20 840.92 612.40 835.77 C624.61 830.62 636.57 824.29 647.51 816.92 C658.45 809.56 668.85 800.94 678.04 791.58 C687.24 782.21 695.50 771.58 702.66 760.71 C709.82 749.85 715.73 737.96 721.00 726.41 C726.28 714.86 730.36 702.87 734.31 691.41 C738.26 679.96 741.47 668.66 744.71 657.71 C747.94 646.76 750.85 636.18 753.72 625.71 C756.60 615.24 759.34 605.08 761.94 594.88 C764.55 584.68 767.06 574.65 769.33 564.49 C771.60 554.34 773.85 543.28 775.57 533.93 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M280.42 275.00 C272.26 282.14 265.93 288.15 258.78 295.00 C251.64 301.84 244.66 308.88 237.53 316.05 C230.40 323.23 223.40 330.57 215.98 338.06 C208.56 345.55 201.16 353.15 193.00 360.99 C184.85 368.84 176.22 376.68 167.03 385.11 C157.83 393.54 147.20 402.02 137.84 411.59 C128.48 421.15 118.50 431.43 110.86 442.49 C103.21 453.54 96.85 465.64 91.97 477.92 C87.09 490.19 83.85 503.19 81.60 516.14 C79.35 529.09 78.48 542.43 78.46 555.62 C78.45 568.80 79.59 582.15 81.51 595.25 C83.42 608.35 85.74 620.13 89.97 634.24 C94.20 648.34 100.08 665.20 106.91 679.85 C113.74 694.50 121.80 708.86 130.95 722.13 C140.11 735.41 150.51 748.14 161.82 759.52 C173.14 770.91 185.70 781.41 198.84 790.44 C211.98 799.48 226.28 807.23 240.67 813.72 C255.05 820.21 270.29 825.13 285.13 829.37 C299.98 833.61 315.11 836.46 329.73 839.14 C344.36 841.82 358.77 843.67 372.88 845.46 C386.99 847.25 400.74 848.66 414.40 849.89 C428.06 851.12 441.44 852.14 454.84 852.84 C468.24 853.54 481.49 854.03 494.79 854.09 C508.10 854.16 522.56 853.79 534.68 853.21 C546.81 852.64 556.55 851.74 567.56 850.65 C578.57 849.56 589.56 848.16 600.75 846.69 C611.93 845.22 623.10 843.49 634.66 841.85 C646.23 840.21 657.83 838.42 670.15 836.87 C682.46 835.32 695.12 834.00 708.56 832.54 C722.00 831.07 736.56 830.42 750.78 828.09 C765.00 825.75 780.11 823.19 793.86 818.53 C807.60 813.88 820.98 807.56 833.26 800.14 C845.54 792.73 856.97 783.68 867.55 774.03 C878.12 764.37 887.83 753.51 896.72 742.22 C905.61 730.92 913.66 718.74 920.90 706.24 C928.14 693.75 934.17 682.13 940.17 667.24 C946.16 652.34 952.66 633.99 956.87 616.88 C961.09 599.77 964.09 582.09 965.48 564.56 C966.87 547.04 966.87 529.14 965.22 511.72 C963.58 494.30 960.34 476.78 955.64 460.06 C950.94 443.34 944.45 426.89 937.03 411.40 C929.60 395.91 920.45 381.11 911.10 367.13 C901.75 353.15 891.23 340.09 880.92 327.52 C870.61 314.94 859.87 303.18 849.26 291.70 C838.65 280.22 828.00 269.30 817.26 258.63 C806.51 247.96 795.79 237.67 784.79 227.69 C773.79 217.70 762.71 208.02 751.28 198.71 C739.84 189.39 728.35 180.09 716.16 171.81 C703.98 163.52 691.38 155.50 678.17 149.00 C664.97 142.49 651.05 136.93 636.96 132.80 C622.88 128.67 608.20 125.69 593.67 124.24 C579.13 122.78 564.20 122.72 549.74 124.09 C535.28 125.46 520.73 128.48 506.91 132.45 C493.09 136.43 479.58 142.08 466.80 147.94 C454.02 153.79 441.88 160.83 430.22 167.61 C418.55 174.38 407.56 181.59 396.81 188.58 C386.05 195.58 375.80 202.56 365.69 209.56 C355.57 216.55 345.78 223.46 336.12 230.56 C326.46 237.66 317.02 244.75 307.73 252.16 C298.45 259.57 288.57 267.86 280.42 275.00 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M534.68 853.21 C546.81 852.64 556.55 851.74 567.56 850.65 C578.57 849.56 589.56 848.16 600.75 846.69 C611.93 845.22 623.10 843.49 634.66 841.85 C646.23 840.21 657.83 838.42 670.15 836.87 C682.46 835.32 695.12 834.00 708.56 832.54 C722.00 831.07 736.56 830.42 750.78 828.09 C765.00 825.75 780.11 823.19 793.86 818.53 C807.60 813.88 820.98 807.56 833.26 800.14 C845.54 792.73 856.97 783.68 867.55 774.03 C878.12 764.37 887.83 753.51 896.72 742.22 C905.61 730.92 913.66 718.74 920.90 706.24 C928.14 693.75 934.17 682.13 940.17 667.24 C946.16 652.34 952.66 633.99 956.87 616.88 C961.09 599.77 964.09 582.09 965.48 564.56 C966.87 547.04 966.87 529.14 965.22 511.72 C963.58 494.30 960.34 476.78 955.64 460.06 C950.94 443.34 944.45 426.89 937.03 411.40 C929.60 395.91 920.45 381.11 911.10 367.13 C901.75 353.15 891.23 340.09 880.92 327.52 C870.61 314.94 859.87 303.18 849.26 291.70 C838.65 280.22 828.00 269.30 817.26 258.63 C806.51 247.96 795.79 237.67 784.79 227.69 C773.79 217.70 762.71 208.02 751.28 198.71 C739.84 189.39 727.06 179.70 716.16 171.81 C705.26 163.92 696.18 157.99 685.88 151.38 C675.57 144.78 665.05 138.47 654.32 132.15 C643.60 125.83 632.69 119.77 621.52 113.45 C610.35 107.12 599.03 100.94 587.30 94.19 C575.58 87.43 563.71 80.32 551.15 72.90 C538.58 65.48 525.64 56.77 511.91 49.64 C498.19 42.52 483.61 35.07 468.80 30.16 C453.99 25.26 438.44 21.94 423.05 20.22 C407.67 18.50 391.93 18.64 376.50 19.84 C361.07 21.05 345.57 23.80 330.45 27.46 C315.33 31.11 300.33 36.05 285.79 41.78 C271.25 47.52 258.33 53.32 243.20 61.89 C228.07 70.48 210.18 81.64 195.00 93.28 C179.81 104.92 165.25 117.84 152.09 131.70 C138.92 145.55 126.66 160.63 116.04 176.38 C105.43 192.13 96.05 209.02 88.40 226.22 C80.75 243.42 74.74 261.58 70.14 279.60 C65.55 297.61 62.81 316.22 60.83 334.33 C58.85 352.43 58.43 370.60 58.23 388.25 C58.04 405.91 58.77 423.24 59.66 440.26 C60.55 457.29 61.90 473.92 63.57 490.41 C65.23 506.90 67.21 523.10 69.67 539.21 C72.12 555.32 74.92 571.24 78.30 587.08 C81.68 602.92 85.20 618.77 89.97 634.24 C94.74 649.70 100.08 665.20 106.91 679.85 C113.74 694.50 121.80 708.86 130.95 722.13 C140.11 735.41 150.51 748.14 161.82 759.52 C173.14 770.91 185.70 781.41 198.84 790.44 C211.98 799.48 226.28 807.23 240.67 813.72 C255.05 820.21 270.29 825.13 285.13 829.37 C299.98 833.61 315.11 836.46 329.73 839.14 C344.36 841.82 358.77 843.67 372.88 845.46 C386.99 847.25 400.74 848.66 414.40 849.89 C428.06 851.12 441.44 852.14 454.84 852.84 C468.24 853.54 481.49 854.03 494.79 854.09 C508.10 854.16 522.56 853.79 534.68 853.21 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M716.16 171.81 C705.26 163.92 696.18 157.99 685.88 151.38 C675.57 144.78 665.05 138.47 654.32 132.15 C643.60 125.83 632.69 119.77 621.52 113.45 C610.35 107.12 599.03 100.94 587.30 94.19 C575.58 87.43 563.71 80.32 551.15 72.90 C538.58 65.48 525.64 56.77 511.91 49.64 C498.19 42.52 483.61 35.07 468.80 30.16 C453.99 25.26 438.44 21.94 423.05 20.22 C407.67 18.50 391.93 18.64 376.50 19.84 C361.07 21.05 345.57 23.80 330.45 27.46 C315.33 31.11 300.33 36.05 285.79 41.78 C271.25 47.52 258.33 53.32 243.20 61.89 C228.07 70.48 210.18 81.64 195.00 93.28 C179.81 104.92 165.25 117.84 152.09 131.70 C138.92 145.55 126.66 160.63 116.04 176.38 C105.43 192.13 96.05 209.02 88.40 226.22 C80.75 243.42 74.74 261.58 70.14 279.60 C65.55 297.61 62.81 316.22 60.83 334.33 C58.85 352.43 58.43 370.60 58.23 388.25 C58.04 405.91 58.77 423.24 59.66 440.26 C60.55 457.29 61.90 473.92 63.57 490.41 C65.23 506.90 67.21 523.10 69.67 539.21 C72.12 555.32 74.92 571.24 78.30 587.08 C81.68 602.92 86.01 620.01 89.97 634.24 C93.93 648.47 97.70 659.76 102.08 672.46 C106.46 685.16 111.25 697.73 116.23 710.42 C121.21 723.10 126.55 735.67 131.97 748.58 C137.39 761.50 143.07 774.35 148.75 787.88 C154.42 801.41 160.03 815.26 166.03 829.76 C172.02 844.26 177.55 859.96 184.72 874.87 C191.88 889.77 199.54 905.43 209.00 919.20 C218.46 932.97 229.54 945.93 241.50 957.49 C253.47 969.05 266.90 979.33 280.78 988.55 C294.66 997.76 309.59 1005.77 324.79 1012.78 C339.99 1019.79 355.92 1025.68 372.00 1030.60 C388.08 1035.53 402.82 1039.34 421.27 1042.33 C439.72 1045.31 462.20 1048.01 482.72 1048.52 C503.24 1049.02 524.07 1048.05 544.37 1045.39 C564.68 1042.72 585.04 1038.44 604.54 1032.51 C624.04 1026.58 643.28 1018.90 661.36 1009.82 C679.44 1000.74 696.84 989.80 713.01 978.02 C729.18 966.24 744.26 952.76 758.40 939.13 C772.53 925.50 785.45 910.78 797.82 896.25 C810.19 881.72 821.59 866.80 832.62 851.96 C843.65 837.12 854.03 822.23 864.00 807.21 C873.97 792.19 883.46 777.15 892.44 761.83 C901.43 746.52 909.98 731.08 917.93 715.32 C925.89 699.55 933.67 683.64 940.17 667.24 C946.66 650.83 952.66 633.99 956.87 616.88 C961.09 599.77 964.09 582.09 965.48 564.56 C966.87 547.04 966.87 529.14 965.22 511.72 C963.58 494.30 960.34 476.78 955.64 460.06 C950.94 443.34 944.45 426.89 937.03 411.40 C929.60 395.91 920.45 381.11 911.10 367.13 C901.75 353.15 891.23 340.09 880.92 327.52 C870.61 314.94 859.87 303.18 849.26 291.70 C838.65 280.22 828.00 269.30 817.26 258.63 C806.51 247.96 795.79 237.67 784.79 227.69 C773.79 217.70 762.71 208.02 751.28 198.71 C739.84 189.39 727.06 179.70 716.16 171.81 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M89.97 634.24 C93.93 648.47 97.70 659.76 102.08 672.46 C106.46 685.16 111.25 697.73 116.23 710.42 C121.21 723.10 126.55 735.67 131.97 748.58 C137.39 761.50 143.07 774.35 148.75 787.88 C154.42 801.41 160.03 815.26 166.03 829.76 C172.02 844.26 177.55 859.96 184.72 874.87 C191.88 889.77 199.54 905.43 209.00 919.20 C218.46 932.97 229.54 945.93 241.50 957.49 C253.47 969.05 266.90 979.33 280.78 988.55 C294.66 997.76 309.59 1005.77 324.79 1012.78 C339.99 1019.79 355.92 1025.68 372.00 1030.60 C388.08 1035.53 402.82 1039.34 421.27 1042.33 C439.72 1045.31 462.20 1048.01 482.72 1048.52 C503.24 1049.02 524.07 1048.05 544.37 1045.39 C564.68 1042.72 585.04 1038.44 604.54 1032.51 C624.04 1026.58 643.28 1018.90 661.36 1009.82 C679.44 1000.74 696.84 989.80 713.01 978.02 C729.18 966.24 744.26 952.76 758.40 939.13 C772.53 925.50 785.45 910.78 797.82 896.25 C810.19 881.72 821.59 866.80 832.62 851.96 C843.65 837.12 854.03 822.23 864.00 807.21 C873.97 792.19 883.46 777.15 892.44 761.83 C901.43 746.52 909.98 731.08 917.93 715.32 C925.89 699.55 933.85 682.05 940.17 667.24 C946.48 652.42 950.93 640.22 955.80 626.43 C960.67 612.63 965.08 598.65 969.36 584.46 C973.64 570.27 977.51 555.92 981.48 541.27 C985.45 526.61 989.13 511.82 993.18 496.52 C997.22 481.22 1001.47 465.70 1005.76 449.44 C1010.06 433.19 1015.41 416.28 1018.94 399.00 C1022.47 381.71 1026.03 363.53 1026.95 345.71 C1027.87 327.89 1026.97 309.70 1024.46 292.08 C1021.94 274.46 1017.43 256.93 1011.84 239.99 C1006.26 223.04 999.06 206.40 990.97 190.40 C982.89 174.40 973.49 158.84 963.32 143.98 C953.15 129.11 943.39 116.08 929.98 101.23 C916.57 86.38 899.61 69.07 882.85 54.87 C866.09 40.67 848.10 27.49 829.41 16.03 C810.73 4.57 790.93 -5.60 770.75 -13.87 C750.57 -22.14 729.45 -28.82 708.33 -33.59 C687.21 -38.36 665.43 -41.18 644.05 -42.50 C622.68 -43.81 601.07 -43.04 580.10 -41.47 C559.13 -39.91 538.40 -36.64 518.22 -33.11 C498.05 -29.57 478.37 -25.03 459.05 -20.25 C439.73 -15.46 420.89 -10.14 402.30 -4.39 C383.71 1.37 365.50 7.54 347.53 14.30 C329.56 21.07 311.88 28.30 294.50 36.23 C277.11 44.16 259.78 52.39 243.20 61.89 C226.62 71.41 210.18 81.64 195.00 93.28 C179.81 104.92 165.25 117.84 152.09 131.70 C138.92 145.55 126.66 160.63 116.04 176.38 C105.43 192.13 96.05 209.02 88.40 226.22 C80.75 243.42 74.74 261.58 70.14 279.60 C65.55 297.61 62.81 316.22 60.83 334.33 C58.85 352.43 58.43 370.60 58.23 388.25 C58.04 405.91 58.77 423.24 59.66 440.26 C60.55 457.29 61.90 473.92 63.57 490.41 C65.23 506.90 67.21 523.10 69.67 539.21 C72.12 555.32 74.92 571.24 78.30 587.08 C81.68 602.92 86.01 620.01 89.97 634.24 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M940.17 667.24 C946.48 652.42 950.93 640.22 955.80 626.43 C960.67 612.63 965.08 598.65 969.36 584.46 C973.64 570.27 977.51 555.92 981.48 541.27 C985.45 526.61 989.13 511.82 993.18 496.52 C997.22 481.22 1001.47 465.70 1005.76 449.44 C1010.06 433.19 1015.41 416.28 1018.94 399.00 C1022.47 381.71 1026.03 363.53 1026.95 345.71 C1027.87 327.89 1026.97 309.70 1024.46 292.08 C1021.94 274.46 1017.43 256.93 1011.84 239.99 C1006.26 223.04 999.06 206.40 990.97 190.40 C982.89 174.40 973.49 158.84 963.32 143.98 C953.15 129.11 943.39 116.08 929.98 101.23 C916.57 86.38 899.61 69.07 882.85 54.87 C866.09 40.67 848.10 27.49 829.41 16.03 C810.73 4.57 790.93 -5.60 770.75 -13.87 C750.57 -22.14 729.45 -28.82 708.33 -33.59 C687.21 -38.36 665.43 -41.18 644.05 -42.50 C622.68 -43.81 601.07 -43.04 580.10 -41.47 C559.13 -39.91 538.40 -36.64 518.22 -33.11 C498.05 -29.57 478.37 -25.03 459.05 -20.25 C439.73 -15.46 420.89 -10.14 402.30 -4.39 C383.71 1.37 365.50 7.54 347.53 14.30 C329.56 21.07 311.88 28.30 294.50 36.23 C277.11 44.16 258.54 53.60 243.20 61.89 C227.86 70.20 215.86 77.56 202.45 86.01 C189.03 94.47 175.88 103.40 162.73 112.63 C149.57 121.85 136.66 131.50 123.52 141.36 C110.37 151.21 97.40 161.38 83.85 171.75 C70.31 182.11 56.52 192.55 42.28 203.54 C28.04 214.53 12.62 225.39 -1.56 237.71 C-15.74 250.02 -30.44 263.03 -42.79 277.44 C-55.14 291.84 -66.24 307.70 -75.67 324.14 C-85.09 340.58 -92.85 358.22 -99.34 376.09 C-105.84 393.95 -110.82 412.60 -114.66 431.32 C-118.49 450.03 -120.97 469.24 -122.36 488.38 C-123.75 507.51 -124.27 524.85 -122.99 546.14 C-121.70 567.42 -119.11 593.10 -114.64 616.08 C-110.18 639.06 -104.03 662.02 -96.19 684.03 C-88.35 706.04 -78.75 727.73 -67.59 748.14 C-56.44 768.54 -43.46 788.27 -29.24 806.46 C-15.01 824.66 1.03 841.72 17.75 857.29 C34.47 872.86 52.76 886.91 71.10 899.88 C89.43 912.85 108.71 924.30 127.77 935.12 C146.83 945.94 166.19 955.62 185.46 964.80 C204.73 973.98 224.02 982.39 243.40 990.22 C262.79 998.05 282.19 1005.27 301.78 1011.78 C321.38 1018.29 341.08 1024.19 360.99 1029.28 C380.91 1034.38 400.98 1039.12 421.27 1042.33 C441.56 1045.53 462.20 1048.01 482.72 1048.52 C503.24 1049.02 524.07 1048.05 544.37 1045.39 C564.68 1042.72 585.04 1038.44 604.54 1032.51 C624.04 1026.58 643.28 1018.90 661.36 1009.82 C679.44 1000.74 696.84 989.80 713.01 978.02 C729.18 966.24 744.26 952.76 758.40 939.13 C772.53 925.50 785.45 910.78 797.82 896.25 C810.19 881.72 821.59 866.80 832.62 851.96 C843.65 837.12 854.03 822.23 864.00 807.21 C873.97 792.19 883.46 777.15 892.44 761.83 C901.43 746.52 909.98 731.08 917.93 715.32 C925.89 699.55 933.85 682.05 940.17 667.24 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M243.20 61.89 C227.86 70.20 215.86 77.56 202.45 86.01 C189.03 94.47 175.88 103.40 162.73 112.63 C149.57 121.85 136.66 131.50 123.52 141.36 C110.37 151.21 97.40 161.38 83.85 171.75 C70.31 182.11 56.52 192.55 42.28 203.54 C28.04 214.53 12.62 225.39 -1.56 237.71 C-15.74 250.02 -30.44 263.03 -42.79 277.44 C-55.14 291.84 -66.24 307.70 -75.67 324.14 C-85.09 340.58 -92.85 358.22 -99.34 376.09 C-105.84 393.95 -110.82 412.60 -114.66 431.32 C-118.49 450.03 -120.97 469.24 -122.36 488.38 C-123.75 507.51 -124.27 524.85 -122.99 546.14 C-121.70 567.42 -119.11 593.10 -114.64 616.08 C-110.18 639.06 -104.03 662.02 -96.19 684.03 C-88.35 706.04 -78.75 727.73 -67.59 748.14 C-56.44 768.54 -43.46 788.27 -29.24 806.46 C-15.01 824.66 1.03 841.72 17.75 857.29 C34.47 872.86 52.76 886.91 71.10 899.88 C89.43 912.85 108.71 924.30 127.77 935.12 C146.83 945.94 166.19 955.62 185.46 964.80 C204.73 973.98 224.02 982.39 243.40 990.22 C262.79 998.05 282.19 1005.27 301.78 1011.78 C321.38 1018.29 341.08 1024.19 360.99 1029.28 C380.91 1034.38 402.83 1038.96 421.27 1042.33 C439.71 1045.69 454.71 1047.58 471.61 1049.47 C488.51 1051.36 505.51 1052.66 522.67 1053.67 C539.83 1054.68 557.06 1055.15 574.58 1055.54 C592.11 1055.92 609.70 1055.89 627.83 1056.00 C645.97 1056.10 664.34 1056.25 683.40 1056.18 C702.46 1056.11 722.35 1056.83 742.17 1055.58 C761.99 1054.34 782.66 1052.83 802.33 1048.72 C822.01 1044.61 841.61 1038.55 860.24 1030.95 C878.87 1023.34 896.93 1013.70 914.11 1003.07 C931.29 992.44 947.75 980.20 963.32 967.17 C978.89 954.13 993.67 939.82 1007.51 924.85 C1021.36 909.88 1033.30 895.85 1046.38 877.35 C1059.47 858.85 1074.39 835.86 1086.00 813.86 C1097.61 791.86 1107.83 768.79 1116.06 745.37 C1124.28 721.95 1130.87 697.64 1135.35 673.34 C1139.83 649.04 1142.40 624.10 1142.94 599.58 C1143.48 575.05 1141.83 550.25 1138.62 526.19 C1135.41 502.13 1129.95 478.26 1123.69 455.21 C1117.44 432.17 1109.41 409.72 1101.08 387.92 C1092.75 366.12 1083.38 345.03 1073.70 324.41 C1064.02 303.79 1053.77 283.78 1043.00 264.20 C1032.22 244.62 1020.98 225.53 1009.07 206.92 C997.16 188.31 984.72 170.14 971.54 152.53 C958.36 134.91 944.76 117.50 929.98 101.23 C915.20 84.95 899.61 69.07 882.85 54.87 C866.09 40.67 848.10 27.49 829.41 16.03 C810.73 4.57 790.93 -5.60 770.75 -13.87 C750.57 -22.14 729.45 -28.82 708.33 -33.59 C687.21 -38.36 665.43 -41.18 644.05 -42.50 C622.68 -43.81 601.07 -43.04 580.10 -41.47 C559.13 -39.91 538.40 -36.64 518.22 -33.11 C498.05 -29.57 478.37 -25.03 459.05 -20.25 C439.73 -15.46 420.89 -10.14 402.30 -4.39 C383.71 1.37 365.50 7.54 347.53 14.30 C329.56 21.07 311.88 28.30 294.50 36.23 C277.11 44.16 258.54 53.60 243.20 61.89 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M421.27 1042.33 C439.71 1045.69 454.71 1047.58 471.61 1049.47 C488.51 1051.36 505.51 1052.66 522.67 1053.67 C539.83 1054.68 557.06 1055.15 574.58 1055.54 C592.11 1055.92 609.70 1055.89 627.83 1056.00 C645.97 1056.10 664.34 1056.25 683.40 1056.18 C702.46 1056.11 722.35 1056.83 742.17 1055.58 C761.99 1054.34 782.66 1052.83 802.33 1048.72 C822.01 1044.61 841.61 1038.55 860.24 1030.95 C878.87 1023.34 896.93 1013.70 914.11 1003.07 C931.29 992.44 947.75 980.20 963.32 967.17 C978.89 954.13 993.67 939.82 1007.51 924.85 C1021.36 909.88 1033.30 895.85 1046.38 877.35 C1059.47 858.85 1074.39 835.86 1086.00 813.86 C1097.61 791.86 1107.83 768.79 1116.06 745.37 C1124.28 721.95 1130.87 697.64 1135.35 673.34 C1139.83 649.04 1142.40 624.10 1142.94 599.58 C1143.48 575.05 1141.83 550.25 1138.62 526.19 C1135.41 502.13 1129.95 478.26 1123.69 455.21 C1117.44 432.17 1109.41 409.72 1101.08 387.92 C1092.75 366.12 1083.38 345.03 1073.70 324.41 C1064.02 303.79 1053.77 283.78 1043.00 264.20 C1032.22 244.62 1020.98 225.53 1009.07 206.92 C997.16 188.31 984.72 170.14 971.54 152.53 C958.36 134.91 943.12 116.38 929.98 101.23 C916.84 86.07 905.55 74.46 892.68 61.60 C879.82 48.74 866.46 36.32 852.77 24.07 C839.08 11.81 824.96 -0.02 810.53 -11.92 C796.10 -23.83 781.33 -35.40 766.19 -47.37 C751.06 -59.35 735.74 -71.45 719.73 -83.77 C703.72 -96.08 687.53 -109.49 670.16 -121.26 C652.79 -133.04 634.57 -145.08 615.49 -154.43 C596.41 -163.78 576.15 -171.50 555.69 -177.35 C535.22 -183.20 513.92 -187.05 492.69 -189.53 C471.45 -192.00 449.76 -192.72 428.29 -192.19 C406.81 -191.67 385.15 -189.61 363.82 -186.39 C342.49 -183.17 323.35 -179.49 300.31 -172.88 C277.28 -166.28 249.75 -157.20 225.60 -146.77 C201.45 -136.34 177.73 -124.12 155.40 -110.31 C133.08 -96.50 111.50 -80.89 91.65 -63.89 C71.80 -46.88 53.12 -28.04 36.38 -8.04 C19.41 12.47 3.42 34.34 -10.36 56.18 C-23.98 78.01 -35.59 101.17 -45.97 124.23 C-56.34 147.30 -64.88 171.07 -72.70 194.55 C-80.53 218.04 -87.03 241.66 -92.91 265.14 C-98.79 288.62 -103.77 312.02 -107.99 335.43 C-112.21 358.83 -115.68 382.17 -118.25 405.57 C-120.81 428.96 -122.61 452.36 -123.40 475.79 C-124.19 499.22 -124.45 522.76 -122.99 546.14 C-121.53 569.52 -119.11 593.10 -114.64 616.08 C-110.18 639.06 -104.03 662.02 -96.19 684.03 C-88.35 706.04 -78.75 727.73 -67.59 748.14 C-56.44 768.54 -43.46 788.27 -29.24 806.46 C-15.01 824.66 1.03 841.72 17.75 857.29 C34.47 872.86 52.76 886.91 71.10 899.88 C89.43 912.85 108.71 924.30 127.77 935.12 C146.83 945.94 166.19 955.62 185.46 964.80 C204.73 973.98 224.02 982.39 243.40 990.22 C262.79 998.05 282.19 1005.27 301.78 1011.78 C321.38 1018.29 341.08 1024.19 360.99 1029.28 C380.91 1034.38 402.83 1038.96 421.27 1042.33 Z`

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
	__Parameter__000000_Reference.SideLength = 170.000000
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
	__Parameter__000000_Reference.NbPitchLines = 169
	__Parameter__000000_Reference.BeatLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbBeatLines = 300
	__Parameter__000000_Reference.NbOfBeatsInTheme = 16
	__Parameter__000000_Reference.FirstVoiceShiftX = 0.060000
	__Parameter__000000_Reference.FirstVoiceShiftY = 0.610000
	__Parameter__000000_Reference.PitchDifference = 12
	__Parameter__000000_Reference.BeatsPerSecond = 4.000000
	__Parameter__000000_Reference.Level = 2.200000
	__Parameter__000000_Reference.IsMinor = false
	__Parameter__000000_Reference.ThemeBinaryEncoding = 4101
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
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 170.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 170.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 275.820186
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 38.626275
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 170.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 170.000000
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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 274.746991
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -33.366695
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 319.542124
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 36.818879
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 215.190613
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 328.424591
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 251.561322
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 248.184501
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 180.200000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 274.746991
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 180.200000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -33.366695
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 274.746991
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 362.478959
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -33.366695
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -89.360611
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 83.024294
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 10.082900
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 66.220837
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 188.202797
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 274.746991
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -33.366695
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 53.400307
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -151.766234
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
