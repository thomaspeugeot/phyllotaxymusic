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
	__Axis__000001_Construction_Axis.Length = 143.394109
	__Axis__000001_Construction_Axis.CenterX = 0.000000
	__Axis__000001_Construction_Axis.CenterY = 0.000000
	__Axis__000001_Construction_Axis.EndX = -19.887065
	__Axis__000001_Construction_Axis.EndY = 142.008363
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
	__Axis__000002_Initial_Axis.Length = 516.965934
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
	__Axis__000004_Rotated_Axis.Length = 516.965934
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
	__Bezier__000005_Growth_Curve_Seed.StartX = -9.943533
	__Bezier__000005_Growth_Curve_Seed.StartY = 71.004181
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartX = 44.524953
	__Bezier__000005_Growth_Curve_Seed.ControlPointStartY = 78.632030
	__Bezier__000005_Growth_Curve_Seed.EndX = 81.517415
	__Bezier__000005_Growth_Curve_Seed.EndY = 156.209199
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndX = 27.048929
	__Bezier__000005_Growth_Curve_Seed.ControlPointEndY = 148.581350
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
	__Circle__000001_Construction_Circle.CenterX = -9.943533
	__Circle__000001_Construction_Circle.CenterY = 71.004181
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
	__Circle__000005_Rotated_Next_Circle.CenterX = 202.808961
	__Circle__000005_Rotated_Next_Circle.CenterY = 28.401673
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
	__FrontCurve__000000_Non_Rotated_0_.Path = `M702.82 525.13 C704.09 518.26 704.90 512.70 705.82 506.37 C706.75 500.04 707.52 493.69 708.36 487.18 C709.20 480.67 709.93 474.12 710.87 467.29 C711.81 460.46 712.72 453.57 714.00 446.17 C715.28 438.78 716.87 431.17 718.55 422.94 C720.22 414.70 722.74 405.77 724.07 396.78 C725.40 387.79 726.78 378.09 726.51 368.99 C726.23 359.89 724.79 350.77 722.44 342.20 C720.09 333.63 716.50 325.37 712.40 317.57 C708.30 309.77 703.26 302.36 697.83 295.40 C692.41 288.44 686.27 281.89 679.83 275.80 C673.38 269.71 667.26 264.49 659.16 258.88 C651.06 253.26 640.93 246.89 631.23 242.11 C621.52 237.33 611.28 233.24 600.92 230.21 C590.56 227.18 579.77 224.98 569.08 223.92 C558.39 222.85 547.41 222.81 536.78 223.82 C526.15 224.83 515.45 227.05 505.28 229.97 C495.11 232.90 485.18 237.05 475.78 241.36 C466.38 245.67 457.45 250.84 448.87 255.82 C440.29 260.80 432.21 266.10 424.30 271.24 C416.39 276.38 408.85 281.51 401.42 286.65 C393.98 291.80 386.78 296.87 379.68 302.09 C372.58 307.31 365.63 312.52 358.81 317.96 C351.98 323.41 344.72 329.50 338.73 334.75 C332.73 339.99 328.08 344.41 322.83 349.44 C317.58 354.47 312.45 359.65 307.21 364.92 C301.97 370.20 296.83 375.60 291.37 381.10 C285.92 386.61 280.48 392.20 274.49 397.96 C268.49 403.73 262.16 409.50 255.40 415.70 C248.64 421.90 240.83 428.14 233.94 435.17 C227.06 442.21 219.73 449.77 214.11 457.91 C208.49 466.04 203.81 474.94 200.22 483.97 C196.64 492.99 194.25 502.56 192.59 512.08 C190.93 521.60 190.29 531.42 190.27 541.11 C190.26 550.81 191.10 560.62 192.50 570.25 C193.91 579.89 195.61 588.56 198.72 598.92 C201.83 609.29 206.14 621.70 211.16 632.47 C216.18 643.24 222.11 653.79 228.84 663.56 C235.57 673.32 243.21 682.67 251.53 691.04 C259.85 699.42 269.08 707.14 278.74 713.78 C288.40 720.42 298.92 726.12 309.50 730.89 C320.07 735.66 331.28 739.28 342.19 742.40 C353.10 745.51 364.23 747.61 374.98 749.58 C385.73 751.55 396.33 752.91 406.71 754.22 C417.08 755.54 427.19 756.57 437.24 757.48 C447.28 758.38 457.12 759.13 466.97 759.64 C476.82 760.16 486.57 760.52 496.35 760.56 C506.14 760.61 515.94 760.68 525.68 759.92 C535.43 759.15 545.29 758.10 554.82 755.96 C564.34 753.82 573.86 750.88 582.83 747.09 C591.80 743.30 600.60 738.64 608.64 733.23 C616.69 727.81 624.34 721.47 631.10 714.58 C637.86 707.70 643.94 699.88 649.20 691.89 C654.47 683.90 658.81 675.16 662.69 666.67 C666.57 658.17 669.57 649.35 672.48 640.93 C675.38 632.51 677.74 624.20 680.12 616.15 C682.50 608.09 684.64 600.32 686.75 592.62 C688.87 584.92 690.89 577.45 692.80 569.95 C694.71 562.45 696.56 555.07 698.23 547.60 C699.90 540.13 701.56 532.01 702.82 525.13 Z`

	__FrontCurve__000001_Non_Rotated_1_.Name = `Non Rotated 1 `
	__FrontCurve__000001_Non_Rotated_1_.Path = `M338.73 334.75 C332.73 339.99 328.08 344.41 322.83 349.44 C317.58 354.47 312.45 359.65 307.21 364.92 C301.97 370.20 296.83 375.60 291.37 381.10 C285.92 386.61 280.48 392.20 274.49 397.96 C268.49 403.73 262.16 409.50 255.40 415.70 C248.64 421.90 240.83 428.14 233.94 435.17 C227.06 442.21 219.73 449.77 214.11 457.91 C208.49 466.04 203.81 474.94 200.22 483.97 C196.64 492.99 194.25 502.56 192.59 512.08 C190.93 521.60 190.29 531.42 190.27 541.11 C190.26 550.81 191.10 560.62 192.50 570.25 C193.91 579.89 195.61 588.56 198.72 598.92 C201.83 609.29 206.14 621.70 211.16 632.47 C216.18 643.24 222.11 653.79 228.84 663.56 C235.57 673.32 243.21 682.67 251.53 691.04 C259.85 699.42 269.08 707.14 278.74 713.78 C288.40 720.42 298.92 726.12 309.50 730.89 C320.07 735.66 331.28 739.28 342.19 742.40 C353.10 745.51 364.23 747.61 374.98 749.58 C385.73 751.55 396.33 752.91 406.71 754.22 C417.08 755.54 427.19 756.57 437.24 757.48 C447.28 758.38 457.12 759.13 466.97 759.64 C476.82 760.16 486.57 760.52 496.35 760.56 C506.14 760.61 516.76 760.34 525.68 759.92 C534.60 759.49 541.76 758.83 549.86 758.03 C557.95 757.23 566.03 756.19 574.26 755.12 C582.48 754.04 590.69 752.76 599.20 751.56 C607.70 750.35 616.23 749.03 625.29 747.89 C634.35 746.75 643.65 745.78 653.54 744.70 C663.42 743.63 674.13 743.15 684.58 741.43 C695.03 739.71 706.15 737.83 716.26 734.40 C726.37 730.98 736.20 726.34 745.23 720.88 C754.26 715.43 762.67 708.78 770.45 701.68 C778.22 694.58 785.36 686.59 791.90 678.28 C798.44 669.97 804.36 661.02 809.68 651.83 C815.01 642.64 819.44 634.10 823.85 623.15 C828.26 612.19 833.03 598.70 836.14 586.12 C839.24 573.53 841.45 560.53 842.47 547.65 C843.49 534.76 843.49 521.60 842.28 508.79 C841.08 495.98 838.69 483.10 835.24 470.81 C831.78 458.51 827.01 446.41 821.55 435.02 C816.10 423.63 809.37 412.75 802.49 402.47 C795.61 392.19 787.88 382.59 780.30 373.35 C772.72 364.11 764.82 355.46 757.02 347.02 C749.22 338.57 741.39 330.55 733.49 322.70 C725.59 314.86 717.70 307.30 709.62 299.95 C701.53 292.61 693.39 285.50 684.98 278.65 C676.57 271.80 668.12 264.97 659.16 258.88 C650.20 252.79 640.93 246.89 631.23 242.11 C621.52 237.33 611.28 233.24 600.92 230.21 C590.56 227.18 579.77 224.98 569.08 223.92 C558.39 222.85 547.41 222.81 536.78 223.82 C526.15 224.83 515.45 227.05 505.28 229.97 C495.11 232.90 485.18 237.05 475.78 241.36 C466.38 245.67 457.45 250.84 448.87 255.82 C440.29 260.80 432.21 266.10 424.30 271.24 C416.39 276.38 408.85 281.51 401.42 286.65 C393.98 291.80 386.78 296.87 379.68 302.09 C372.58 307.31 365.63 312.52 358.81 317.96 C351.98 323.41 344.72 329.50 338.73 334.75 Z`

	__FrontCurve__000002_Non_Rotated_2_.Name = `Non Rotated 2 `
	__FrontCurve__000002_Non_Rotated_2_.Path = `M525.68 759.92 C534.60 759.49 541.76 758.83 549.86 758.03 C557.95 757.23 566.03 756.19 574.26 755.12 C582.48 754.04 590.69 752.76 599.20 751.56 C607.70 750.35 616.23 749.03 625.29 747.89 C634.35 746.75 643.65 745.78 653.54 744.70 C663.42 743.63 674.13 743.15 684.58 741.43 C695.03 739.71 706.15 737.83 716.26 734.40 C726.37 730.98 736.20 726.34 745.23 720.88 C754.26 715.43 762.67 708.78 770.45 701.68 C778.22 694.58 785.36 686.59 791.90 678.28 C798.44 669.97 804.36 661.02 809.68 651.83 C815.01 642.64 819.44 634.10 823.85 623.15 C828.26 612.19 833.03 598.70 836.14 586.12 C839.24 573.53 841.45 560.53 842.47 547.65 C843.49 534.76 843.49 521.60 842.28 508.79 C841.08 495.98 838.69 483.10 835.24 470.81 C831.78 458.51 827.01 446.41 821.55 435.02 C816.10 423.63 809.37 412.75 802.49 402.47 C795.61 392.19 787.88 382.59 780.30 373.35 C772.72 364.11 764.82 355.46 757.02 347.02 C749.22 338.57 741.39 330.55 733.49 322.70 C725.59 314.86 717.70 307.30 709.62 299.95 C701.53 292.61 693.39 285.50 684.98 278.65 C676.57 271.80 667.17 264.67 659.16 258.88 C651.14 253.08 644.47 248.72 636.89 243.86 C629.31 239.01 621.58 234.38 613.69 229.73 C605.80 225.08 597.78 220.63 589.57 215.98 C581.35 211.33 573.03 206.79 564.41 201.83 C555.79 196.86 547.06 191.64 537.82 186.19 C528.59 180.73 519.07 174.33 508.98 169.10 C498.89 163.87 488.17 158.39 477.28 154.79 C466.39 151.19 454.95 148.76 443.64 147.50 C432.32 146.24 420.75 146.35 409.40 147.24 C398.05 148.14 386.65 150.17 375.52 152.86 C364.40 155.55 353.37 159.18 342.67 163.40 C331.97 167.62 322.46 171.88 311.33 178.18 C300.20 184.48 287.04 192.68 275.88 201.21 C264.71 209.75 254.02 219.23 244.36 229.39 C234.69 239.56 225.70 250.63 217.91 262.20 C210.13 273.78 203.25 286.19 197.64 298.84 C192.02 311.49 187.61 324.85 184.23 338.10 C180.85 351.36 178.84 365.04 177.37 378.36 C175.91 391.68 175.60 405.04 175.45 418.03 C175.30 431.01 175.84 443.76 176.48 456.28 C177.13 468.80 178.12 481.03 179.34 493.16 C180.56 505.29 182.01 517.20 183.81 529.04 C185.61 540.89 187.66 552.60 190.15 564.25 C192.63 575.90 195.22 587.55 198.72 598.92 C202.22 610.29 206.14 621.70 211.16 632.47 C216.18 643.24 222.11 653.79 228.84 663.56 C235.57 673.32 243.21 682.67 251.53 691.04 C259.85 699.42 269.08 707.14 278.74 713.78 C288.40 720.42 298.92 726.12 309.50 730.89 C320.07 735.66 331.28 739.28 342.19 742.40 C353.10 745.51 364.23 747.61 374.98 749.58 C385.73 751.55 396.33 752.91 406.71 754.22 C417.08 755.54 427.19 756.57 437.24 757.48 C447.28 758.38 457.12 759.13 466.97 759.64 C476.82 760.16 486.57 760.52 496.35 760.56 C506.14 760.61 516.76 760.34 525.68 759.92 Z`

	__FrontCurve__000003_Non_Rotated_3_.Name = `Non Rotated 3 `
	__FrontCurve__000003_Non_Rotated_3_.Path = `M659.16 258.88 C651.14 253.08 644.47 248.72 636.89 243.86 C629.31 239.01 621.58 234.38 613.69 229.73 C605.80 225.08 597.78 220.63 589.57 215.98 C581.35 211.33 573.03 206.79 564.41 201.83 C555.79 196.86 547.06 191.64 537.82 186.19 C528.59 180.73 519.07 174.33 508.98 169.10 C498.89 163.87 488.17 158.39 477.28 154.79 C466.39 151.19 454.95 148.76 443.64 147.50 C432.32 146.24 420.75 146.35 409.40 147.24 C398.05 148.14 386.65 150.17 375.52 152.86 C364.40 155.55 353.37 159.18 342.67 163.40 C331.97 167.62 322.46 171.88 311.33 178.18 C300.20 184.48 287.04 192.68 275.88 201.21 C264.71 209.75 254.02 219.23 244.36 229.39 C234.69 239.56 225.70 250.63 217.91 262.20 C210.13 273.78 203.25 286.19 197.64 298.84 C192.02 311.49 187.61 324.85 184.23 338.10 C180.85 351.36 178.84 365.04 177.37 378.36 C175.91 391.68 175.60 405.04 175.45 418.03 C175.30 431.01 175.84 443.76 176.48 456.28 C177.13 468.80 178.12 481.03 179.34 493.16 C180.56 505.29 182.01 517.20 183.81 529.04 C185.61 540.89 187.66 552.60 190.15 564.25 C192.63 575.90 195.81 588.46 198.72 598.92 C201.63 609.39 204.40 617.69 207.62 627.03 C210.83 636.37 214.35 645.61 218.01 654.94 C221.67 664.27 225.60 673.51 229.58 683.01 C233.57 692.50 237.74 701.95 241.91 711.90 C246.08 721.85 250.21 732.04 254.61 742.70 C259.02 753.36 263.08 764.90 268.35 775.86 C273.61 786.82 279.24 798.34 286.20 808.46 C293.16 818.59 301.30 828.12 310.10 836.62 C318.89 845.11 328.77 852.68 338.98 859.45 C349.18 866.22 360.16 872.11 371.34 877.26 C382.51 882.41 394.23 886.75 406.05 890.37 C417.87 893.99 428.71 896.79 442.28 898.99 C455.85 901.18 472.38 903.16 487.46 903.53 C502.55 903.91 517.87 903.19 532.80 901.23 C547.73 899.26 562.70 896.12 577.04 891.76 C591.38 887.40 605.53 881.75 618.82 875.07 C632.11 868.39 644.91 860.35 656.80 851.68 C668.70 843.02 679.78 833.11 690.18 823.08 C700.57 813.06 710.07 802.23 719.17 791.55 C728.27 780.87 736.65 769.89 744.76 758.98 C752.87 748.07 760.50 737.12 767.84 726.08 C775.17 715.03 782.14 703.97 788.75 692.71 C795.36 681.45 801.65 670.10 807.50 658.50 C813.35 646.91 819.08 635.21 823.85 623.15 C828.62 611.08 833.03 598.70 836.14 586.12 C839.24 573.53 841.45 560.53 842.47 547.65 C843.49 534.76 843.49 521.60 842.28 508.79 C841.08 495.98 838.69 483.10 835.24 470.81 C831.78 458.51 827.01 446.41 821.55 435.02 C816.10 423.63 809.37 412.75 802.49 402.47 C795.61 392.19 787.88 382.59 780.30 373.35 C772.72 364.11 764.82 355.46 757.02 347.02 C749.22 338.57 741.39 330.55 733.49 322.70 C725.59 314.86 717.70 307.30 709.62 299.95 C701.53 292.61 693.39 285.50 684.98 278.65 C676.57 271.80 667.17 264.67 659.16 258.88 Z`

	__FrontCurve__000004_Non_Rotated_4_.Name = `Non Rotated 4 `
	__FrontCurve__000004_Non_Rotated_4_.Path = `M198.72 598.92 C201.63 609.39 204.40 617.69 207.62 627.03 C210.83 636.37 214.35 645.61 218.01 654.94 C221.67 664.27 225.60 673.51 229.58 683.01 C233.57 692.50 237.74 701.95 241.91 711.90 C246.08 721.85 250.21 732.04 254.61 742.70 C259.02 753.36 263.08 764.90 268.35 775.86 C273.61 786.82 279.24 798.34 286.20 808.46 C293.16 818.59 301.30 828.12 310.10 836.62 C318.89 845.11 328.77 852.68 338.98 859.45 C349.18 866.22 360.16 872.11 371.34 877.26 C382.51 882.41 394.23 886.75 406.05 890.37 C417.87 893.99 428.71 896.79 442.28 898.99 C455.85 901.18 472.38 903.16 487.46 903.53 C502.55 903.91 517.87 903.19 532.80 901.23 C547.73 899.26 562.70 896.12 577.04 891.76 C591.38 887.40 605.53 881.75 618.82 875.07 C632.11 868.39 644.91 860.35 656.80 851.68 C668.70 843.02 679.78 833.11 690.18 823.08 C700.57 813.06 710.07 802.23 719.17 791.55 C728.27 780.87 736.65 769.89 744.76 758.98 C752.87 748.07 760.50 737.12 767.84 726.08 C775.17 715.03 782.14 703.97 788.75 692.71 C795.36 681.45 801.65 670.10 807.50 658.50 C813.35 646.91 819.21 634.04 823.85 623.15 C828.49 612.25 831.77 603.28 835.35 593.14 C838.93 582.99 842.17 572.71 845.32 562.28 C848.47 551.84 851.32 541.29 854.24 530.52 C857.16 519.74 859.86 508.87 862.84 497.61 C865.82 486.36 868.94 474.95 872.10 462.99 C875.26 451.04 879.19 438.61 881.79 425.90 C884.39 413.19 887.01 399.82 887.69 386.72 C888.36 373.61 887.70 360.24 885.85 347.28 C884.00 334.33 880.69 321.44 876.58 308.98 C872.48 296.51 867.19 284.28 861.24 272.52 C855.30 260.75 848.38 249.31 840.91 238.38 C833.44 227.45 826.26 217.87 816.40 206.95 C806.54 196.03 794.07 183.31 781.75 172.87 C769.43 162.43 756.20 152.74 742.46 144.32 C728.73 135.89 714.17 128.41 699.33 122.34 C684.49 116.26 668.96 111.36 653.44 107.85 C637.91 104.35 621.89 102.28 606.17 101.32 C590.46 100.36 574.57 100.93 559.15 102.09 C543.73 103.25 528.48 105.66 513.65 108.26 C498.81 110.87 484.33 114.21 470.12 117.74 C455.91 121.27 442.05 125.18 428.38 129.42 C414.71 133.67 401.31 138.20 388.09 143.19 C374.88 148.17 361.87 153.48 349.07 159.31 C336.28 165.15 323.53 171.20 311.33 178.18 C299.13 185.16 287.04 192.68 275.88 201.21 C264.71 209.75 254.02 219.23 244.36 229.39 C234.69 239.56 225.70 250.63 217.91 262.20 C210.13 273.78 203.25 286.19 197.64 298.84 C192.02 311.49 187.61 324.85 184.23 338.10 C180.85 351.36 178.84 365.04 177.37 378.36 C175.91 391.68 175.60 405.04 175.45 418.03 C175.30 431.01 175.84 443.76 176.48 456.28 C177.13 468.80 178.12 481.03 179.34 493.16 C180.56 505.29 182.01 517.20 183.81 529.04 C185.61 540.89 187.66 552.60 190.15 564.25 C192.63 575.90 195.81 588.46 198.72 598.92 Z`

	__FrontCurve__000005_Non_Rotated_5_.Name = `Non Rotated 5 `
	__FrontCurve__000005_Non_Rotated_5_.Path = `M823.85 623.15 C828.49 612.25 831.77 603.28 835.35 593.14 C838.93 582.99 842.17 572.71 845.32 562.28 C848.47 551.84 851.32 541.29 854.24 530.52 C857.16 519.74 859.86 508.87 862.84 497.61 C865.82 486.36 868.94 474.95 872.10 462.99 C875.26 451.04 879.19 438.61 881.79 425.90 C884.39 413.19 887.01 399.82 887.69 386.72 C888.36 373.61 887.70 360.24 885.85 347.28 C884.00 334.33 880.69 321.44 876.58 308.98 C872.48 296.51 867.19 284.28 861.24 272.52 C855.30 260.75 848.38 249.31 840.91 238.38 C833.44 227.45 826.26 217.87 816.40 206.95 C806.54 196.03 794.07 183.31 781.75 172.87 C769.43 162.43 756.20 152.74 742.46 144.32 C728.73 135.89 714.17 128.41 699.33 122.34 C684.49 116.26 668.96 111.36 653.44 107.85 C637.91 104.35 621.89 102.28 606.17 101.32 C590.46 100.36 574.57 100.93 559.15 102.09 C543.73 103.25 528.48 105.66 513.65 108.26 C498.81 110.87 484.33 114.21 470.12 117.74 C455.91 121.27 442.05 125.18 428.38 129.42 C414.71 133.67 401.31 138.20 388.09 143.19 C374.88 148.17 361.87 153.48 349.07 159.31 C336.28 165.15 322.62 172.09 311.33 178.18 C300.05 184.27 291.22 189.68 281.36 195.88 C271.49 202.08 261.83 208.63 252.16 215.40 C242.50 222.16 233.02 229.23 223.38 236.46 C213.73 243.69 204.22 251.16 194.29 258.78 C184.36 266.39 174.24 274.07 163.79 282.16 C153.35 290.25 142.02 298.25 131.60 307.32 C121.19 316.39 110.38 325.98 101.29 336.58 C92.20 347.19 84.03 358.86 77.08 370.95 C70.14 383.05 64.43 396.03 59.64 409.16 C54.85 422.30 51.17 436.02 48.34 449.78 C45.51 463.55 43.68 477.67 42.65 491.74 C41.62 505.81 41.22 518.56 42.16 534.20 C43.10 549.85 45.00 568.73 48.27 585.63 C51.55 602.52 56.07 619.40 61.83 635.58 C67.59 651.76 74.64 667.71 82.84 682.72 C91.04 697.72 100.58 712.22 111.03 725.59 C121.49 738.97 133.28 751.51 145.58 762.96 C157.87 774.40 171.32 784.73 184.80 794.27 C198.28 803.81 212.45 812.22 226.47 820.18 C240.48 828.13 254.72 835.24 268.89 842.00 C283.06 848.75 297.23 854.93 311.49 860.68 C325.74 866.44 340.01 871.74 354.42 876.53 C368.83 881.32 383.31 885.66 397.95 889.40 C412.60 893.14 427.36 896.63 442.28 898.99 C457.19 901.34 472.38 903.16 487.46 903.53 C502.55 903.91 517.87 903.19 532.80 901.23 C547.73 899.26 562.70 896.12 577.04 891.76 C591.38 887.40 605.53 881.75 618.82 875.07 C632.11 868.39 644.91 860.35 656.80 851.68 C668.70 843.02 679.78 833.11 690.18 823.08 C700.57 813.06 710.07 802.23 719.17 791.55 C728.27 780.87 736.65 769.89 744.76 758.98 C752.87 748.07 760.50 737.12 767.84 726.08 C775.17 715.03 782.14 703.97 788.75 692.71 C795.36 681.45 801.65 670.10 807.50 658.50 C813.35 646.91 819.21 634.04 823.85 623.15 Z`

	__FrontCurve__000006_Non_Rotated_6_.Name = `Non Rotated 6 `
	__FrontCurve__000006_Non_Rotated_6_.Path = `M311.33 178.18 C300.05 184.27 291.22 189.68 281.36 195.88 C271.49 202.08 261.83 208.63 252.16 215.40 C242.50 222.16 233.02 229.23 223.38 236.46 C213.73 243.69 204.22 251.16 194.29 258.78 C184.36 266.39 174.24 274.07 163.79 282.16 C153.35 290.25 142.02 298.25 131.60 307.32 C121.19 316.39 110.38 325.98 101.29 336.58 C92.20 347.19 84.03 358.86 77.08 370.95 C70.14 383.05 64.43 396.03 59.64 409.16 C54.85 422.30 51.17 436.02 48.34 449.78 C45.51 463.55 43.68 477.67 42.65 491.74 C41.62 505.81 41.22 518.56 42.16 534.20 C43.10 549.85 45.00 568.73 48.27 585.63 C51.55 602.52 56.07 619.40 61.83 635.58 C67.59 651.76 74.64 667.71 82.84 682.72 C91.04 697.72 100.58 712.22 111.03 725.59 C121.49 738.97 133.28 751.51 145.58 762.96 C157.87 774.40 171.32 784.73 184.80 794.27 C198.28 803.81 212.45 812.22 226.47 820.18 C240.48 828.13 254.72 835.24 268.89 842.00 C283.06 848.75 297.23 854.93 311.49 860.68 C325.74 866.44 340.01 871.74 354.42 876.53 C368.83 881.32 383.31 885.66 397.95 889.40 C412.60 893.14 428.72 896.51 442.28 898.99 C455.83 901.46 466.87 902.84 479.29 904.23 C491.72 905.62 504.22 906.58 516.84 907.32 C529.46 908.06 542.12 908.41 555.01 908.69 C567.90 908.97 580.83 908.95 594.17 909.03 C607.50 909.10 621.01 909.21 635.03 909.16 C649.04 909.11 663.67 909.63 678.24 908.72 C692.82 907.80 708.01 906.69 722.48 903.67 C736.95 900.65 751.37 896.19 765.07 890.59 C778.76 885.00 792.04 877.91 804.68 870.10 C817.31 862.28 829.42 853.28 840.86 843.69 C852.31 834.11 863.18 823.58 873.36 812.57 C883.54 801.57 892.33 791.25 901.95 777.64 C911.57 764.04 922.54 747.13 931.08 730.96 C939.62 714.78 947.14 697.81 953.19 680.59 C959.24 663.37 964.08 645.49 967.38 627.62 C970.67 609.76 972.56 591.42 972.96 573.39 C973.36 555.35 972.14 537.11 969.79 519.42 C967.43 501.73 963.42 484.17 958.82 467.23 C954.22 450.28 948.31 433.78 942.19 417.75 C936.06 401.72 929.18 386.21 922.07 371.05 C914.95 355.89 907.41 341.17 899.49 326.78 C891.57 312.38 883.30 298.35 874.55 284.66 C865.79 270.98 856.64 257.62 846.95 244.67 C837.26 231.72 827.27 218.92 816.40 206.95 C805.53 194.98 794.07 183.31 781.75 172.87 C769.43 162.43 756.20 152.74 742.46 144.32 C728.73 135.89 714.17 128.41 699.33 122.34 C684.49 116.26 668.96 111.36 653.44 107.85 C637.91 104.35 621.89 102.28 606.17 101.32 C590.46 100.36 574.57 100.93 559.15 102.09 C543.73 103.25 528.48 105.66 513.65 108.26 C498.81 110.87 484.33 114.21 470.12 117.74 C455.91 121.27 442.05 125.18 428.38 129.42 C414.71 133.67 401.31 138.20 388.09 143.19 C374.88 148.17 361.87 153.48 349.07 159.31 C336.28 165.15 322.62 172.09 311.33 178.18 Z`

	__FrontCurve__000007_Non_Rotated_7_.Name = `Non Rotated 7 `
	__FrontCurve__000007_Non_Rotated_7_.Path = `M442.28 898.99 C455.83 901.46 466.87 902.84 479.29 904.23 C491.72 905.62 504.22 906.58 516.84 907.32 C529.46 908.06 542.12 908.41 555.01 908.69 C567.90 908.97 580.83 908.95 594.17 909.03 C607.50 909.10 621.01 909.21 635.03 909.16 C649.04 909.11 663.67 909.63 678.24 908.72 C692.82 907.80 708.01 906.69 722.48 903.67 C736.95 900.65 751.37 896.19 765.07 890.59 C778.76 885.00 792.04 877.91 804.68 870.10 C817.31 862.28 829.42 853.28 840.86 843.69 C852.31 834.11 863.18 823.58 873.36 812.57 C883.54 801.57 892.33 791.25 901.95 777.64 C911.57 764.04 922.54 747.13 931.08 730.96 C939.62 714.78 947.14 697.81 953.19 680.59 C959.24 663.37 964.08 645.49 967.38 627.62 C970.67 609.76 972.56 591.42 972.96 573.39 C973.36 555.35 972.14 537.11 969.79 519.42 C967.43 501.73 963.42 484.17 958.82 467.23 C954.22 450.28 948.31 433.78 942.19 417.75 C936.06 401.72 929.18 386.21 922.07 371.05 C914.95 355.89 907.41 341.17 899.49 326.78 C891.57 312.38 883.30 298.35 874.55 284.66 C865.79 270.98 856.64 257.62 846.95 244.67 C837.26 231.72 826.06 218.09 816.40 206.95 C806.74 195.81 798.44 187.27 788.98 177.82 C779.52 168.37 769.70 159.23 759.64 150.22 C749.57 141.21 739.19 132.52 728.58 123.76 C717.97 115.01 707.11 106.51 695.99 97.70 C684.86 88.90 673.60 80.00 661.83 70.95 C650.06 61.90 638.16 52.04 625.39 43.39 C612.62 34.73 599.23 25.88 585.21 19.01 C571.18 12.14 556.29 6.47 541.25 2.18 C526.20 -2.12 510.54 -4.94 494.94 -6.76 C479.33 -8.57 463.39 -9.08 447.60 -8.68 C431.81 -8.29 415.89 -6.76 400.21 -4.37 C384.53 -1.99 370.47 0.74 353.54 5.62 C336.60 10.50 316.37 17.21 298.62 24.92 C280.86 32.63 263.42 41.66 247.00 51.87 C230.57 62.09 214.69 73.63 200.06 86.19 C185.42 98.76 171.60 112.63 159.16 127.22 C146.72 141.82 135.45 157.63 125.45 173.75 C115.46 189.86 106.89 206.94 99.21 223.92 C91.55 240.91 85.22 258.40 79.43 275.68 C73.64 292.96 68.83 310.33 64.47 327.60 C60.12 344.86 56.44 362.07 53.31 379.28 C50.19 396.49 47.62 413.65 45.72 430.85 C43.82 448.05 42.48 465.26 41.89 482.48 C41.30 499.71 41.10 517.01 42.16 534.20 C43.23 551.40 45.00 568.73 48.27 585.63 C51.55 602.52 56.07 619.40 61.83 635.58 C67.59 651.76 74.64 667.71 82.84 682.72 C91.04 697.72 100.58 712.22 111.03 725.59 C121.49 738.97 133.28 751.51 145.58 762.96 C157.87 774.40 171.32 784.73 184.80 794.27 C198.28 803.81 212.45 812.22 226.47 820.18 C240.48 828.13 254.72 835.24 268.89 842.00 C283.06 848.75 297.23 854.93 311.49 860.68 C325.74 866.44 340.01 871.74 354.42 876.53 C368.83 881.32 383.31 885.66 397.95 889.40 C412.60 893.14 428.72 896.51 442.28 898.99 Z`

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
	__Parameter__000000_Reference.SideLength = 125.000000
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
	__Parameter__000000_Reference.IsMinor = true
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
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 125.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 125.000000
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
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 202.808961
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 28.401673
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 125.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 125.000000
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
	__ShapeCategory__000007_7_Spiral_Growth_Bezier.IsExpanded = true

	__ShapeCategory__000008_8_Score_notation.Name = `8. Score notation`
	__ShapeCategory__000008_8_Score_notation.IsExpanded = true

	__ShapeCategory__000009_9_Composer.Name = `9. Composer`
	__ShapeCategory__000009_9_Composer.IsExpanded = true

	__SpiralBezier__000000_Spiral_Bezier_Seed.Name = `Spiral Bezier Seed`
	__SpiralBezier__000000_Spiral_Bezier_Seed.IsDisplayed = false
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartX = 202.019846
	__SpiralBezier__000000_Spiral_Bezier_Seed.StartY = -24.534335
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartX = 234.957444
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointStartY = 27.072705
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndX = 158.228392
	__SpiralBezier__000000_Spiral_Bezier_Seed.EndY = 241.488670
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndX = 184.971560
	__SpiralBezier__000000_Spiral_Bezier_Seed.ControlPointEndY = 182.488604
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
	__SpiralCircle__000000_Construction_Circle_Spiral.CenterX = 132.500000
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
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartX = 202.019846
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.EndX = 132.500000
	__SpiralLine__000000_Spiral_Contruction_Inner_Line.StartY = -24.534335
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
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartX = 202.019846
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndX = 266.528646
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.StartY = -24.534335
	__SpiralLine__000001_Spiral_Contruction_Outer_Line.EndY = -65.706332
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
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r0 = 61.047275
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r0 = 7.413897
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r1 = 48.691792
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r1 = 138.384409
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r2 = 202.019846
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r2 = -24.534335
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.X_r3 = 39.264931
	__SpiralRhombus__000000_Reference_Spiral_Rhombus.Y_r3 = -111.592819
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
