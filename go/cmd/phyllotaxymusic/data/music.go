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
	__Axis__000001_Construction_Axis.Length = 212.223281
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -29.432857
	__Axis__000001_Construction_Axis.EndY = 210.172377
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
	__Axis__000002_Initial_Axis.Length = 765.109582
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
	__Axis__000004_Rotated_Axis.Length = 765.109582
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -14.716428
	__Bezier__000005_Growth_Curve_Seed.StartY = 105.086189
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 65.896930
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 116.375405
	__Bezier__000005_Growth_Curve_Seed.EndX = 120.645774
	__Bezier__000005_Growth_Curve_Seed.EndY = 231.189615
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 40.032416
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 219.900399
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
	__Circle__000001_Construction_Circle.CenterX = -14.716428
	__Circle__000001_Construction_Circle.CenterY = 105.086189
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 300.157262
	__Circle__000005_Rotated_Next_Circle.CenterY = 42.034475
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M799.82 536.87 C801.69 526.70 802.90 518.46 804.26 509.10 C805.63 499.74 806.77 490.34 808.01 480.70 C809.25 471.06 810.33 461.38 811.72 451.26 C813.11 441.15 814.47 430.95 816.36 420.01 C818.25 409.07 820.60 397.80 823.08 385.62 C825.56 373.43 829.28 360.21 831.25 346.91 C833.21 333.60 835.25 319.24 834.85 305.78 C834.45 292.32 832.30 278.81 828.82 266.13 C825.34 253.44 820.04 241.22 813.97 229.68 C807.90 218.13 800.44 207.16 792.40 196.86 C784.37 186.55 775.29 176.86 765.75 167.84 C756.21 158.83 747.15 151.10 735.16 142.79 C723.17 134.47 708.18 125.03 693.82 117.96 C679.45 110.88 664.30 104.82 648.97 100.33 C633.64 95.84 617.67 92.59 601.85 91.01 C586.03 89.43 569.79 89.36 554.05 90.85 C538.32 92.34 522.49 95.62 507.45 99.94 C492.41 104.26 477.71 110.41 463.80 116.79 C449.89 123.16 436.69 130.82 424.00 138.19 C411.30 145.56 399.34 153.41 387.64 161.02 C375.94 168.63 364.79 176.23 353.78 183.85 C342.77 191.46 332.12 198.98 321.61 206.71 C311.10 214.44 300.81 222.16 290.71 230.22 C280.60 238.29 269.86 247.31 260.98 255.08 C252.10 262.85 245.21 269.40 237.43 276.85 C229.65 284.30 222.06 291.96 214.30 299.77 C206.53 307.58 198.92 315.57 190.84 323.72 C182.76 331.87 174.71 340.14 165.83 348.67 C156.95 357.21 147.57 365.74 137.56 374.91 C127.55 384.09 115.98 393.32 105.79 403.72 C95.60 414.13 84.75 425.32 76.43 437.34 C68.11 449.37 61.19 462.54 55.88 475.89 C50.57 489.25 47.04 503.40 44.59 517.49 C42.15 531.58 41.20 546.10 41.19 560.45 C41.17 574.79 42.42 589.31 44.51 603.57 C46.59 617.83 49.11 630.65 53.72 646.00 C58.33 661.34 64.72 679.70 72.16 695.64 C79.59 711.59 88.37 727.20 98.33 741.65 C108.29 756.10 119.61 769.95 131.92 782.34 C144.23 794.73 157.90 806.16 172.20 816.00 C186.50 825.83 202.07 834.27 217.73 841.33 C233.38 848.39 249.96 853.75 266.12 858.36 C282.27 862.97 298.74 866.07 314.65 868.99 C330.57 871.91 346.25 873.92 361.61 875.87 C376.96 877.82 391.93 879.35 406.79 880.69 C421.66 882.03 436.22 883.14 450.80 883.90 C465.38 884.66 479.80 885.20 494.28 885.27 C508.76 885.34 523.27 885.45 537.69 884.31 C552.11 883.18 566.71 881.63 580.80 878.46 C594.90 875.30 608.99 870.94 622.26 865.33 C635.54 859.73 648.56 852.84 660.47 844.82 C672.37 836.81 683.69 827.43 693.69 817.24 C703.70 807.04 712.69 795.47 720.48 783.66 C728.27 771.84 734.70 758.89 740.44 746.33 C746.18 733.76 750.63 720.70 754.93 708.24 C759.22 695.78 762.71 683.48 766.23 671.56 C769.76 659.64 772.92 648.13 776.05 636.74 C779.17 625.34 782.16 614.30 784.99 603.19 C787.82 592.09 790.56 581.18 793.03 570.12 C795.50 559.07 797.95 547.04 799.82 536.87 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M260.98 255.08 C252.10 262.85 245.21 269.40 237.43 276.85 C229.65 284.30 222.06 291.96 214.30 299.77 C206.53 307.58 198.92 315.57 190.84 323.72 C182.76 331.87 174.71 340.14 165.83 348.67 C156.95 357.21 147.57 365.74 137.56 374.91 C127.55 384.09 115.98 393.32 105.79 403.72 C95.60 414.13 84.75 425.32 76.43 437.34 C68.11 449.37 61.19 462.54 55.88 475.89 C50.57 489.25 47.04 503.40 44.59 517.49 C42.15 531.58 41.20 546.10 41.19 560.45 C41.17 574.79 42.42 589.31 44.51 603.57 C46.59 617.83 49.11 630.65 53.72 646.00 C58.33 661.34 64.72 679.70 72.16 695.64 C79.59 711.59 88.37 727.20 98.33 741.65 C108.29 756.10 119.61 769.95 131.92 782.34 C144.23 794.73 157.90 806.16 172.20 816.00 C186.50 825.83 202.07 834.27 217.73 841.33 C233.38 848.39 249.96 853.75 266.12 858.36 C282.27 862.97 298.74 866.07 314.65 868.99 C330.57 871.91 346.25 873.92 361.61 875.87 C376.96 877.82 391.93 879.35 406.79 880.69 C421.66 882.03 436.22 883.14 450.80 883.90 C465.38 884.66 479.80 885.20 494.28 885.27 C508.76 885.34 524.49 884.94 537.69 884.31 C550.88 883.69 561.48 882.71 573.47 881.52 C585.45 880.34 597.41 878.81 609.58 877.21 C621.75 875.62 633.90 873.73 646.49 871.95 C659.07 870.17 671.70 868.22 685.10 866.53 C698.50 864.84 712.28 863.41 726.90 861.81 C741.53 860.22 757.37 859.51 772.85 856.97 C788.32 854.43 804.76 851.65 819.72 846.58 C834.68 841.51 849.23 834.64 862.60 826.57 C875.96 818.49 888.40 808.65 899.91 798.15 C911.42 787.64 921.98 775.82 931.66 763.53 C941.34 751.23 950.10 737.98 957.98 724.38 C965.86 710.78 972.41 698.14 978.94 681.93 C985.46 665.73 992.53 645.76 997.12 627.14 C1001.71 608.51 1004.97 589.27 1006.48 570.20 C1008.00 551.13 1007.99 531.65 1006.20 512.70 C1004.42 493.75 1000.89 474.68 995.77 456.48 C990.66 438.29 983.59 420.38 975.52 403.52 C967.44 386.67 957.48 370.56 947.30 355.35 C937.12 340.13 925.67 325.92 914.46 312.24 C903.24 298.56 891.55 285.76 880.00 273.26 C868.46 260.77 856.87 248.88 845.18 237.27 C833.48 225.66 821.81 214.47 809.84 203.60 C797.87 192.73 785.82 182.20 773.37 172.06 C760.92 161.92 748.42 151.80 735.16 142.79 C721.90 133.77 708.18 125.03 693.82 117.96 C679.45 110.88 664.30 104.82 648.97 100.33 C633.64 95.84 617.67 92.59 601.85 91.01 C586.03 89.43 569.79 89.36 554.05 90.85 C538.32 92.34 522.49 95.62 507.45 99.94 C492.41 104.26 477.71 110.41 463.80 116.79 C449.89 123.16 436.69 130.82 424.00 138.19 C411.30 145.56 399.34 153.41 387.64 161.02 C375.94 168.63 364.79 176.23 353.78 183.85 C342.77 191.46 332.12 198.98 321.61 206.71 C311.10 214.44 300.81 222.16 290.71 230.22 C280.60 238.29 269.86 247.31 260.98 255.08 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M537.69 884.31 C550.88 883.69 561.48 882.71 573.47 881.52 C585.45 880.34 597.41 878.81 609.58 877.21 C621.75 875.62 633.90 873.73 646.49 871.95 C659.07 870.17 671.70 868.22 685.10 866.53 C698.50 864.84 712.28 863.41 726.90 861.81 C741.53 860.22 757.37 859.51 772.85 856.97 C788.32 854.43 804.76 851.65 819.72 846.58 C834.68 841.51 849.23 834.64 862.60 826.57 C875.96 818.49 888.40 808.65 899.91 798.15 C911.42 787.64 921.98 775.82 931.66 763.53 C941.34 751.23 950.10 737.98 957.98 724.38 C965.86 710.78 972.41 698.14 978.94 681.93 C985.46 665.73 992.53 645.76 997.12 627.14 C1001.71 608.51 1004.97 589.27 1006.48 570.20 C1008.00 551.13 1007.99 531.65 1006.20 512.70 C1004.42 493.75 1000.89 474.68 995.77 456.48 C990.66 438.29 983.59 420.38 975.52 403.52 C967.44 386.67 957.48 370.56 947.30 355.35 C937.12 340.13 925.67 325.92 914.46 312.24 C903.24 298.56 891.55 285.76 880.00 273.26 C868.46 260.77 856.87 248.88 845.18 237.27 C833.48 225.66 821.81 214.47 809.84 203.60 C797.87 192.73 785.82 182.20 773.37 172.06 C760.92 161.92 747.02 151.37 735.16 142.79 C723.30 134.20 713.42 127.75 702.20 120.56 C690.98 113.37 679.53 106.51 667.86 99.63 C656.19 92.74 644.32 86.15 632.16 79.27 C620.01 72.38 607.69 65.66 594.93 58.30 C582.16 50.95 569.25 43.22 555.58 35.14 C541.91 27.06 527.82 17.58 512.88 9.82 C497.94 2.07 482.08 -6.05 465.97 -11.39 C449.85 -16.73 432.92 -20.35 416.17 -22.22 C399.43 -24.10 382.31 -23.95 365.51 -22.64 C348.72 -21.34 331.85 -18.35 315.40 -14.37 C298.95 -10.40 282.64 -5.04 266.82 1.20 C251.00 7.44 236.94 13.74 220.48 23.07 C204.02 32.41 184.56 44.57 168.05 57.24 C151.52 69.93 135.67 84.01 121.34 99.12 C107.00 114.22 93.63 130.64 82.05 147.80 C70.48 164.94 60.27 183.31 51.94 202.02 C43.61 220.73 37.07 240.49 32.07 260.09 C27.08 279.69 24.11 299.93 21.95 319.63 C19.80 339.34 19.35 359.10 19.14 378.31 C18.94 397.53 19.74 416.38 20.71 434.91 C21.68 453.44 23.15 471.54 24.97 489.49 C26.79 507.43 28.94 525.05 31.62 542.59 C34.29 560.12 37.33 577.45 41.02 594.68 C44.70 611.92 48.53 629.17 53.72 646.00 C58.91 662.83 64.72 679.70 72.16 695.64 C79.59 711.59 88.37 727.20 98.33 741.65 C108.29 756.10 119.61 769.95 131.92 782.34 C144.23 794.73 157.90 806.16 172.20 816.00 C186.50 825.83 202.07 834.27 217.73 841.33 C233.38 848.39 249.96 853.75 266.12 858.36 C282.27 862.97 298.74 866.07 314.65 868.99 C330.57 871.91 346.25 873.92 361.61 875.87 C376.96 877.82 391.93 879.35 406.79 880.69 C421.66 882.03 436.22 883.14 450.80 883.90 C465.38 884.66 479.80 885.20 494.28 885.27 C508.76 885.34 524.49 884.94 537.69 884.31 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M735.16 142.79 C723.30 134.20 713.42 127.75 702.20 120.56 C690.98 113.37 679.53 106.51 667.86 99.63 C656.19 92.74 644.32 86.15 632.16 79.27 C620.01 72.38 607.69 65.66 594.93 58.30 C582.16 50.95 569.25 43.22 555.58 35.14 C541.91 27.06 527.82 17.58 512.88 9.82 C497.94 2.07 482.08 -6.05 465.97 -11.39 C449.85 -16.73 432.92 -20.35 416.17 -22.22 C399.43 -24.10 382.31 -23.95 365.51 -22.64 C348.72 -21.34 331.85 -18.35 315.40 -14.37 C298.95 -10.40 282.64 -5.04 266.82 1.20 C251.00 7.44 236.94 13.74 220.48 23.07 C204.02 32.41 184.56 44.57 168.05 57.24 C151.52 69.93 135.67 84.01 121.34 99.12 C107.00 114.22 93.63 130.64 82.05 147.80 C70.48 164.94 60.27 183.31 51.94 202.02 C43.61 220.73 37.07 240.49 32.07 260.09 C27.08 279.69 24.11 299.93 21.95 319.63 C19.80 339.34 19.35 359.10 19.14 378.31 C18.94 397.53 19.74 416.38 20.71 434.91 C21.68 453.44 23.15 471.54 24.97 489.49 C26.79 507.43 28.94 525.05 31.62 542.59 C34.29 560.12 37.33 577.45 41.02 594.68 C44.70 611.92 49.41 630.51 53.72 646.00 C58.03 661.48 62.14 673.78 66.90 687.60 C71.67 701.41 76.88 715.09 82.30 728.90 C87.72 742.71 93.54 756.39 99.44 770.44 C105.34 784.49 111.52 798.48 117.70 813.20 C123.87 827.93 129.98 843.00 136.50 858.78 C143.03 874.55 149.05 891.64 156.84 907.86 C164.64 924.09 172.97 941.12 183.27 956.11 C193.57 971.10 205.62 985.20 218.64 997.78 C231.66 1010.36 246.28 1021.55 261.39 1031.58 C276.49 1041.61 292.74 1050.32 309.28 1057.94 C325.83 1065.57 343.16 1071.99 360.66 1077.35 C378.15 1082.71 394.19 1086.86 414.27 1090.11 C434.35 1093.36 458.82 1096.29 481.14 1096.84 C503.47 1097.40 526.14 1096.34 548.23 1093.44 C570.33 1090.53 592.49 1085.88 613.71 1079.43 C634.93 1072.98 655.87 1064.62 675.54 1054.73 C695.21 1044.85 714.15 1032.95 731.75 1020.13 C749.35 1007.31 765.75 992.64 781.14 977.81 C796.52 962.98 810.57 946.96 824.04 931.15 C837.50 915.34 849.91 899.10 861.91 882.95 C873.91 866.80 885.21 850.60 896.05 834.26 C906.90 817.91 917.23 801.54 927.01 784.88 C936.79 768.21 946.09 751.41 954.74 734.26 C963.40 717.10 971.87 699.79 978.94 681.93 C986.00 664.08 992.53 645.76 997.12 627.14 C1001.71 608.51 1004.97 589.27 1006.48 570.20 C1008.00 551.13 1007.99 531.65 1006.20 512.70 C1004.42 493.75 1000.89 474.68 995.77 456.48 C990.66 438.29 983.59 420.38 975.52 403.52 C967.44 386.67 957.48 370.56 947.30 355.35 C937.12 340.13 925.67 325.92 914.46 312.24 C903.24 298.56 891.55 285.76 880.00 273.26 C868.46 260.77 856.87 248.88 845.18 237.27 C833.48 225.66 821.81 214.47 809.84 203.60 C797.87 192.73 785.82 182.20 773.37 172.06 C760.92 161.92 747.02 151.37 735.16 142.79 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M53.72 646.00 C58.03 661.48 62.14 673.78 66.90 687.60 C71.67 701.41 76.88 715.09 82.30 728.90 C87.72 742.71 93.54 756.39 99.44 770.44 C105.34 784.49 111.52 798.48 117.70 813.20 C123.87 827.93 129.98 843.00 136.50 858.78 C143.03 874.55 149.05 891.64 156.84 907.86 C164.64 924.09 172.97 941.12 183.27 956.11 C193.57 971.10 205.62 985.20 218.64 997.78 C231.66 1010.36 246.28 1021.55 261.39 1031.58 C276.49 1041.61 292.74 1050.32 309.28 1057.94 C325.83 1065.57 343.16 1071.99 360.66 1077.35 C378.15 1082.71 394.19 1086.86 414.27 1090.11 C434.35 1093.36 458.82 1096.29 481.14 1096.84 C503.47 1097.40 526.14 1096.34 548.23 1093.44 C570.33 1090.53 592.49 1085.88 613.71 1079.43 C634.93 1072.98 655.87 1064.62 675.54 1054.73 C695.21 1044.85 714.15 1032.95 731.75 1020.13 C749.35 1007.31 765.75 992.64 781.14 977.81 C796.52 962.98 810.57 946.96 824.04 931.15 C837.50 915.34 849.91 899.10 861.91 882.95 C873.91 866.80 885.21 850.60 896.05 834.26 C906.90 817.91 917.23 801.54 927.01 784.88 C936.79 768.21 946.09 751.41 954.74 734.26 C963.40 717.10 972.07 698.06 978.94 681.93 C985.80 665.81 990.66 652.54 995.95 637.52 C1001.25 622.51 1006.05 607.30 1010.71 591.86 C1015.36 576.41 1019.58 560.80 1023.90 544.85 C1028.22 528.90 1032.22 512.81 1036.62 496.16 C1041.03 479.50 1045.64 462.61 1050.32 444.93 C1054.99 427.24 1060.81 408.84 1064.66 390.03 C1068.50 371.22 1072.37 351.43 1073.37 332.04 C1074.37 312.65 1073.40 292.86 1070.65 273.69 C1067.91 254.51 1063.00 235.44 1056.93 216.99 C1050.85 198.55 1043.02 180.45 1034.21 163.03 C1025.41 145.62 1015.18 128.69 1004.12 112.51 C993.05 96.34 982.43 82.15 967.83 65.99 C953.24 49.83 934.79 31.00 916.55 15.55 C898.31 0.09 878.72 -14.25 858.39 -26.72 C838.06 -39.19 816.51 -50.27 794.55 -59.27 C772.59 -68.27 749.60 -75.53 726.62 -80.73 C703.64 -85.92 679.93 -89.00 656.67 -90.43 C633.41 -91.86 609.89 -91.02 587.07 -89.32 C564.25 -87.62 541.69 -84.07 519.74 -80.23 C497.78 -76.38 476.36 -71.45 455.34 -66.25 C434.31 -61.04 413.81 -55.26 393.58 -49.00 C373.36 -42.74 353.54 -36.04 333.99 -28.68 C314.44 -21.32 295.20 -13.47 276.28 -4.85 C257.37 3.78 238.52 12.73 220.48 23.07 C202.45 33.42 184.56 44.57 168.05 57.24 C151.52 69.93 135.67 84.01 121.34 99.12 C107.00 114.22 93.63 130.64 82.05 147.80 C70.48 164.94 60.27 183.31 51.94 202.02 C43.61 220.73 37.07 240.49 32.07 260.09 C27.08 279.69 24.11 299.93 21.95 319.63 C19.80 339.34 19.35 359.10 19.14 378.31 C18.94 397.53 19.74 416.38 20.71 434.91 C21.68 453.44 23.15 471.54 24.97 489.49 C26.79 507.43 28.94 525.05 31.62 542.59 C34.29 560.12 37.33 577.45 41.02 594.68 C44.70 611.92 49.41 630.51 53.72 646.00 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M978.94 681.93 C985.80 665.81 990.66 652.54 995.95 637.52 C1001.25 622.51 1006.05 607.30 1010.71 591.86 C1015.36 576.41 1019.58 560.80 1023.90 544.85 C1028.22 528.90 1032.22 512.81 1036.62 496.16 C1041.03 479.50 1045.64 462.61 1050.32 444.93 C1054.99 427.24 1060.81 408.84 1064.66 390.03 C1068.50 371.22 1072.37 351.43 1073.37 332.04 C1074.37 312.65 1073.40 292.86 1070.65 273.69 C1067.91 254.51 1063.00 235.44 1056.93 216.99 C1050.85 198.55 1043.02 180.45 1034.21 163.03 C1025.41 145.62 1015.18 128.69 1004.12 112.51 C993.05 96.34 982.43 82.15 967.83 65.99 C953.24 49.83 934.79 31.00 916.55 15.55 C898.31 0.09 878.72 -14.25 858.39 -26.72 C838.06 -39.19 816.51 -50.27 794.55 -59.27 C772.59 -68.27 749.60 -75.53 726.62 -80.73 C703.64 -85.92 679.93 -89.00 656.67 -90.43 C633.41 -91.86 609.89 -91.02 587.07 -89.32 C564.25 -87.62 541.69 -84.07 519.74 -80.23 C497.78 -76.38 476.36 -71.45 455.34 -66.25 C434.31 -61.04 413.81 -55.26 393.58 -49.00 C373.36 -42.74 353.54 -36.04 333.99 -28.68 C314.44 -21.32 295.20 -13.47 276.28 -4.85 C257.37 3.78 237.17 14.05 220.48 23.07 C203.79 32.11 190.74 40.12 176.15 49.33 C161.56 58.54 147.25 68.28 132.94 78.34 C118.61 88.40 104.55 98.91 90.22 109.66 C75.88 120.40 61.73 131.48 46.97 142.76 C32.21 154.03 17.19 165.37 1.68 177.32 C-13.81 189.26 -30.59 201.06 -46.02 214.45 C-61.43 227.84 -77.42 242.00 -90.85 257.67 C-104.28 273.34 -116.35 290.59 -126.60 308.49 C-136.84 326.38 -145.27 345.58 -152.33 365.02 C-159.40 384.46 -164.81 404.77 -168.98 425.14 C-173.15 445.51 -175.84 466.41 -177.35 487.23 C-178.86 508.06 -179.43 526.93 -178.02 550.10 C-176.62 573.26 -173.80 601.21 -168.94 626.22 C-164.07 651.23 -157.38 676.21 -148.85 700.17 C-140.31 724.12 -129.87 747.73 -117.73 769.94 C-105.59 792.15 -91.47 813.61 -75.99 833.41 C-60.51 853.21 -43.05 871.78 -24.85 888.73 C-6.66 905.67 13.25 920.96 33.20 935.08 C53.15 949.19 74.13 961.65 94.88 973.43 C115.62 985.20 136.68 995.74 157.66 1005.73 C178.63 1015.73 199.61 1024.88 220.71 1033.40 C241.81 1041.92 262.92 1049.77 284.24 1056.86 C305.57 1063.94 327.00 1070.37 348.67 1075.91 C370.34 1081.45 392.19 1086.62 414.27 1090.11 C436.35 1093.59 458.82 1096.29 481.14 1096.84 C503.47 1097.40 526.14 1096.34 548.23 1093.44 C570.33 1090.53 592.49 1085.88 613.71 1079.43 C634.93 1072.98 655.87 1064.62 675.54 1054.73 C695.21 1044.85 714.15 1032.95 731.75 1020.13 C749.35 1007.31 765.75 992.64 781.14 977.81 C796.52 962.98 810.57 946.96 824.04 931.15 C837.50 915.34 849.91 899.10 861.91 882.95 C873.91 866.80 885.21 850.60 896.05 834.26 C906.90 817.91 917.23 801.54 927.01 784.88 C936.79 768.21 946.09 751.41 954.74 734.26 C963.40 717.10 972.07 698.06 978.94 681.93 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M220.48 23.07 C203.79 32.11 190.74 40.12 176.15 49.33 C161.56 58.54 147.25 68.28 132.94 78.34 C118.61 88.40 104.55 98.91 90.22 109.66 C75.88 120.40 61.73 131.48 46.97 142.76 C32.21 154.03 17.19 165.37 1.68 177.32 C-13.81 189.26 -30.59 201.06 -46.02 214.45 C-61.43 227.84 -77.42 242.00 -90.85 257.67 C-104.28 273.34 -116.35 290.59 -126.60 308.49 C-136.84 326.38 -145.27 345.58 -152.33 365.02 C-159.40 384.46 -164.81 404.77 -168.98 425.14 C-173.15 445.51 -175.84 466.41 -177.35 487.23 C-178.86 508.06 -179.43 526.93 -178.02 550.10 C-176.62 573.26 -173.80 601.21 -168.94 626.22 C-164.07 651.23 -157.38 676.21 -148.85 700.17 C-140.31 724.12 -129.87 747.73 -117.73 769.94 C-105.59 792.15 -91.47 813.61 -75.99 833.41 C-60.51 853.21 -43.05 871.78 -24.85 888.73 C-6.66 905.67 13.25 920.96 33.20 935.08 C53.15 949.19 74.13 961.65 94.88 973.43 C115.62 985.20 136.68 995.74 157.66 1005.73 C178.63 1015.73 199.61 1024.88 220.71 1033.40 C241.81 1041.92 262.92 1049.77 284.24 1056.86 C305.57 1063.94 327.00 1070.37 348.67 1075.91 C370.34 1081.45 394.21 1086.45 414.27 1090.11 C434.33 1093.77 450.66 1095.82 469.05 1097.88 C487.45 1099.93 505.94 1101.35 524.62 1102.45 C543.29 1103.55 562.04 1104.06 581.11 1104.48 C600.18 1104.91 619.32 1104.87 639.06 1104.99 C658.79 1105.10 678.79 1105.26 699.53 1105.18 C720.26 1105.11 741.91 1105.89 763.48 1104.54 C785.05 1103.19 807.54 1101.54 828.95 1097.07 C850.37 1092.60 871.70 1086.01 891.97 1077.73 C912.24 1069.45 931.89 1058.97 950.59 1047.40 C969.28 1035.83 987.20 1022.51 1004.14 1008.33 C1021.08 994.14 1037.17 978.57 1052.23 962.28 C1067.30 945.99 1080.30 930.71 1094.53 910.58 C1108.77 890.45 1125.00 865.43 1137.64 841.50 C1150.28 817.56 1161.40 792.45 1170.35 766.96 C1179.30 741.47 1186.47 715.02 1191.34 688.58 C1196.22 662.14 1199.01 635.00 1199.60 608.31 C1200.19 581.62 1198.39 554.63 1194.89 528.45 C1191.40 502.26 1185.46 476.29 1178.65 451.21 C1171.84 426.13 1163.11 401.70 1154.04 377.98 C1144.97 354.25 1134.78 331.30 1124.25 308.86 C1113.71 286.42 1102.55 264.65 1090.83 243.34 C1079.11 222.03 1066.87 201.27 1053.91 181.01 C1040.95 160.76 1027.41 140.99 1013.06 121.82 C998.72 102.65 983.92 83.70 967.83 65.99 C951.75 48.28 934.79 31.00 916.55 15.55 C898.31 0.09 878.72 -14.25 858.39 -26.72 C838.06 -39.19 816.51 -50.27 794.55 -59.27 C772.59 -68.27 749.60 -75.53 726.62 -80.73 C703.64 -85.92 679.93 -89.00 656.67 -90.43 C633.41 -91.86 609.89 -91.02 587.07 -89.32 C564.25 -87.62 541.69 -84.07 519.74 -80.23 C497.78 -76.38 476.36 -71.45 455.34 -66.25 C434.31 -61.04 413.81 -55.26 393.58 -49.00 C373.36 -42.74 353.54 -36.04 333.99 -28.68 C314.44 -21.32 295.20 -13.47 276.28 -4.85 C257.37 3.78 237.17 14.05 220.48 23.07 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M414.27 1090.11 C434.33 1093.77 450.66 1095.82 469.05 1097.88 C487.45 1099.93 505.94 1101.35 524.62 1102.45 C543.29 1103.55 562.04 1104.06 581.11 1104.48 C600.18 1104.91 619.32 1104.87 639.06 1104.99 C658.79 1105.10 678.79 1105.26 699.53 1105.18 C720.26 1105.11 741.91 1105.89 763.48 1104.54 C785.05 1103.19 807.54 1101.54 828.95 1097.07 C850.37 1092.60 871.70 1086.01 891.97 1077.73 C912.24 1069.45 931.89 1058.97 950.59 1047.40 C969.28 1035.83 987.20 1022.51 1004.14 1008.33 C1021.08 994.14 1037.17 978.57 1052.23 962.28 C1067.30 945.99 1080.30 930.71 1094.53 910.58 C1108.77 890.45 1125.00 865.43 1137.64 841.50 C1150.28 817.56 1161.40 792.45 1170.35 766.96 C1179.30 741.47 1186.47 715.02 1191.34 688.58 C1196.22 662.14 1199.01 635.00 1199.60 608.31 C1200.19 581.62 1198.39 554.63 1194.89 528.45 C1191.40 502.26 1185.46 476.29 1178.65 451.21 C1171.84 426.13 1163.11 401.70 1154.04 377.98 C1144.97 354.25 1134.78 331.30 1124.25 308.86 C1113.71 286.42 1102.55 264.65 1090.83 243.34 C1079.11 222.03 1066.87 201.27 1053.91 181.01 C1040.95 160.76 1027.41 140.99 1013.06 121.82 C998.72 102.65 982.14 82.48 967.83 65.99 C953.53 49.50 941.25 36.87 927.25 22.87 C913.24 8.88 898.71 -4.64 883.81 -17.98 C868.91 -31.31 853.54 -44.19 837.84 -57.14 C822.14 -70.10 806.06 -82.69 789.59 -95.73 C773.12 -108.76 756.44 -121.93 739.02 -135.33 C721.60 -148.73 703.98 -163.32 685.07 -176.14 C666.16 -188.95 646.34 -202.06 625.57 -212.23 C604.81 -222.40 582.76 -230.81 560.48 -237.17 C538.21 -243.54 515.03 -247.74 491.92 -250.43 C468.81 -253.13 445.20 -253.91 421.82 -253.34 C398.44 -252.77 374.87 -250.53 351.65 -247.03 C328.43 -243.53 307.60 -239.53 282.52 -232.35 C257.45 -225.17 227.48 -215.30 201.18 -203.96 C174.88 -192.62 149.04 -179.34 124.71 -164.33 C100.38 -149.32 76.85 -132.37 55.16 -113.92 C33.44 -95.48 12.90 -75.11 -5.70 -53.63 C-24.52 -31.93 -41.49 -8.02 -56.16 16.23 C-70.76 40.22 -83.27 65.57 -94.48 90.75 C-105.71 115.90 -114.97 141.80 -123.45 167.39 C-131.94 192.97 -138.99 218.69 -145.38 244.25 C-151.76 269.81 -157.17 295.29 -161.75 320.76 C-166.34 346.24 -170.10 371.64 -172.89 397.10 C-175.68 422.57 -177.63 448.04 -178.49 473.54 C-179.34 499.03 -179.62 524.65 -178.02 550.10 C-176.43 575.55 -173.80 601.21 -168.94 626.22 C-164.07 651.23 -157.38 676.21 -148.85 700.17 C-140.31 724.12 -129.87 747.73 -117.73 769.94 C-105.59 792.15 -91.47 813.61 -75.99 833.41 C-60.51 853.21 -43.05 871.78 -24.85 888.73 C-6.66 905.67 13.25 920.96 33.20 935.08 C53.15 949.19 74.13 961.65 94.88 973.43 C115.62 985.20 136.68 995.74 157.66 1005.73 C178.63 1015.73 199.61 1024.88 220.71 1033.40 C241.81 1041.92 262.92 1049.77 284.24 1056.86 C305.57 1063.94 327.00 1070.37 348.67 1075.91 C370.34 1081.45 394.21 1086.45 414.27 1090.11 Z`

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
	__Parameter__000000_Reference.PitchHeight = 0.050000
	__Parameter__000000_Reference.NbPitchLines = 169
	__Parameter__000000_Reference.BeatLinesHeightRatio = 0.170000
	__Parameter__000000_Reference.NbBeatLines = 300
	__Parameter__000000_Reference.NbOfBeatsInTheme = 16
	__Parameter__000000_Reference.FirstVoiceShiftX = 0.060000
	__Parameter__000000_Reference.FirstVoiceShiftY = 0.610000
	__Parameter__000000_Reference.PitchDifference = 5
	__Parameter__000000_Reference.BeatsPerSecond = 4.000000
	__Parameter__000000_Reference.Level = 2.200000
	__Parameter__000000_Reference.IsMinor = false
	__Parameter__000000_Reference.ThemeBinaryEncoding = 13511
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 185.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 300.157262
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 42.034475
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 185.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 185.000000
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
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 298.989373
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -36.310815
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 347.737017
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 40.067604
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 234.178020
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 357.403231
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 273.757909
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 270.083133
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 298.989373
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 196.100000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -36.310815
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 298.989373
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 394.462397
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -36.310815
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -97.245371
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 90.349968
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 10.972567
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 72.063852
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 204.808926
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 298.989373
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -36.310815
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 58.112099
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -165.157372
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
