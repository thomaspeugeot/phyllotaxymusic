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

	__Axis__000000_Construction_Axis := (&models.Axis{Name: `Construction Axis`}).Stage(stage)
	__Axis__000001_Initial_Axis := (&models.Axis{Name: `Initial Axis`}).Stage(stage)
	__Axis__000002_Rotated_Axis := (&models.Axis{Name: `Rotated Axis`}).Stage(stage)

	__AxisGrid__000000_Construction_Axis_Grid := (&models.AxisGrid{Name: `Construction Axis Grid`}).Stage(stage)

	__Bezier__000000_Growth_Bezier_Right_Seed := (&models.Bezier{Name: `Growth Bezier Right Seed`}).Stage(stage)
	__Bezier__000001_Growth_Curve_Next_Seed := (&models.Bezier{Name: `Growth Curve Next Seed`}).Stage(stage)
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed := (&models.Bezier{Name: `Growth Curve Next Shift Right Seed`}).Stage(stage)
	__Bezier__000003_Growth_Curve_Seed := (&models.Bezier{Name: `Growth Curve Seed`}).Stage(stage)

	__BezierGrid__000000_Growth_Curve := (&models.BezierGrid{Name: `Growth Curve`}).Stage(stage)
	__BezierGrid__000001_Growth_Curve_Next := (&models.BezierGrid{Name: `Growth Curve Next`}).Stage(stage)
	__BezierGrid__000002_Growth_Curve_Next_Shift_Right := (&models.BezierGrid{Name: `Growth Curve Next Shift Right`}).Stage(stage)
	__BezierGrid__000003_Growth_Curve_Shift_Right := (&models.BezierGrid{Name: `Growth Curve Shift Right`}).Stage(stage)

	__BezierGridStack__000000_The_GrowthCurveStack := (&models.BezierGridStack{Name: `The GrowthCurveStack`}).Stage(stage)

	__Circle__000000_Construction_Circle := (&models.Circle{Name: `Construction Circle`}).Stage(stage)
	__Circle__000001_Growing := (&models.Circle{Name: `Growing`}).Stage(stage)
	__Circle__000002_Growing_Seed_Left := (&models.Circle{Name: `Growing Seed Left`}).Stage(stage)
	__Circle__000003_Initial_Circle := (&models.Circle{Name: `Initial Circle`}).Stage(stage)
	__Circle__000004_Rotated_Next_Circle := (&models.Circle{Name: `Rotated Next Circle`}).Stage(stage)

	__CircleGrid__000000_Construction_Circle_Grid := (&models.CircleGrid{Name: `Construction Circle Grid`}).Stage(stage)
	__CircleGrid__000001_Growing_Circle_Grid := (&models.CircleGrid{Name: `Growing Circle Grid`}).Stage(stage)
	__CircleGrid__000002_Growing_Circle_Grid_Shifted_Left := (&models.CircleGrid{Name: `Growing Circle Grid Shifted Left`}).Stage(stage)
	__CircleGrid__000003_Initial_Circle_Grid := (&models.CircleGrid{Name: `Initial Circle Grid`}).Stage(stage)
	__CircleGrid__000004_Rotated_Circle_Grid := (&models.CircleGrid{Name: `Rotated Circle Grid`}).Stage(stage)

	__HorizontalAxis__000000_Horizontal_Axis := (&models.HorizontalAxis{Name: `Horizontal Axis`}).Stage(stage)

	__Key__000000_F_key := (&models.Key{Name: `F key`}).Stage(stage)

	__Parameter__000000_Reference := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__000000_Growing_Rhombus_Grid_Seed := (&models.Rhombus{Name: `Growing Rhombus Grid Seed`}).Stage(stage)
	__Rhombus__000001_Initial_Rhombus := (&models.Rhombus{Name: `Initial Rhombus`}).Stage(stage)
	__Rhombus__000002_Rotated_Next_Rhombus := (&models.Rhombus{Name: `Rotated Next Rhombus`}).Stage(stage)
	__Rhombus__000003_Rotated_Rhombus := (&models.Rhombus{Name: `Rotated Rhombus`}).Stage(stage)

	__RhombusGrid__000000_Growing_Rhombus_Grid := (&models.RhombusGrid{Name: `Growing Rhombus Grid`}).Stage(stage)
	__RhombusGrid__000001_Initial_Rhombus_Grid := (&models.RhombusGrid{Name: `Initial Rhombus Grid`}).Stage(stage)
	__RhombusGrid__000002_Rotated_Rhombus_Grid := (&models.RhombusGrid{Name: `Rotated Rhombus Grid`}).Stage(stage)

	__ShapeCategory__000000_0_Axes := (&models.ShapeCategory{Name: `0. Axes`}).Stage(stage)
	__ShapeCategory__000001_1_Initial := (&models.ShapeCategory{Name: `1. Initial`}).Stage(stage)
	__ShapeCategory__000002_2_Rotated := (&models.ShapeCategory{Name: `2. Rotated`}).Stage(stage)
	__ShapeCategory__000003_3_Growing := (&models.ShapeCategory{Name: `3. Growing`}).Stage(stage)
	__ShapeCategory__000004_4_Construction := (&models.ShapeCategory{Name: `4. Construction`}).Stage(stage)
	__ShapeCategory__000005_5_Growth := (&models.ShapeCategory{Name: `5. Growth`}).Stage(stage)
	__ShapeCategory__000006_6_Score_notation := (&models.ShapeCategory{Name: `6. Score notation`}).Stage(stage)

	__VerticalAxis__000000_Vertical_Axis := (&models.VerticalAxis{Name: `Vertical Axis`}).Stage(stage)

	// Setup of values

	__Axis__000000_Construction_Axis.Name = `Construction Axis`
	__Axis__000000_Construction_Axis.IsDisplayed = false
	__Axis__000000_Construction_Axis.Angle = 99.862709
	__Axis__000000_Construction_Axis.Length = 367.393056
	__Axis__000000_Construction_Axis.CenterX = 0.000000
	__Axis__000000_Construction_Axis.CenterY = 0.000000
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
	__Axis__000001_Initial_Axis.Angle = 80.137291
	__Axis__000001_Initial_Axis.Length = 1072.443030
	__Axis__000001_Initial_Axis.CenterX = 0.000000
	__Axis__000001_Initial_Axis.CenterY = 0.000000
	__Axis__000001_Initial_Axis.Color = ``
	__Axis__000001_Initial_Axis.FillOpacity = 0.000000
	__Axis__000001_Initial_Axis.Stroke = `black`
	__Axis__000001_Initial_Axis.StrokeOpacity = 1.000000
	__Axis__000001_Initial_Axis.StrokeWidth = 2.000000
	__Axis__000001_Initial_Axis.StrokeDashArray = ``
	__Axis__000001_Initial_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000001_Initial_Axis.Transform = ``

	__Axis__000002_Rotated_Axis.Name = `Rotated Axis`
	__Axis__000002_Rotated_Axis.IsDisplayed = false
	__Axis__000002_Rotated_Axis.Angle = 0.000000
	__Axis__000002_Rotated_Axis.Length = 1072.443030
	__Axis__000002_Rotated_Axis.CenterX = 0.000000
	__Axis__000002_Rotated_Axis.CenterY = 0.000000
	__Axis__000002_Rotated_Axis.Color = ``
	__Axis__000002_Rotated_Axis.FillOpacity = 0.000000
	__Axis__000002_Rotated_Axis.Stroke = `black`
	__Axis__000002_Rotated_Axis.StrokeOpacity = 1.000000
	__Axis__000002_Rotated_Axis.StrokeWidth = 2.000000
	__Axis__000002_Rotated_Axis.StrokeDashArray = ``
	__Axis__000002_Rotated_Axis.StrokeDashArrayWhenSelected = ``
	__Axis__000002_Rotated_Axis.Transform = ``

	__AxisGrid__000000_Construction_Axis_Grid.Name = `Construction Axis Grid`
	__AxisGrid__000000_Construction_Axis_Grid.IsDisplayed = false

	__Bezier__000000_Growth_Bezier_Right_Seed.Name = `Growth Bezier Right Seed`
	__Bezier__000000_Growth_Bezier_Right_Seed.IsDisplayed = false
	__Bezier__000000_Growth_Bezier_Right_Seed.StartX = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.StartY = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.ControlPointStartX = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.ControlPointStartY = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.EndX = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.EndY = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.ControlPointEndX = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.ControlPointEndY = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.Color = ``
	__Bezier__000000_Growth_Bezier_Right_Seed.FillOpacity = 0.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.Stroke = `green`
	__Bezier__000000_Growth_Bezier_Right_Seed.StrokeOpacity = 0.500000
	__Bezier__000000_Growth_Bezier_Right_Seed.StrokeWidth = 6.000000
	__Bezier__000000_Growth_Bezier_Right_Seed.StrokeDashArray = ``
	__Bezier__000000_Growth_Bezier_Right_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000000_Growth_Bezier_Right_Seed.Transform = ``

	__Bezier__000001_Growth_Curve_Next_Seed.Name = `Growth Curve Next Seed`
	__Bezier__000001_Growth_Curve_Next_Seed.IsDisplayed = false
	__Bezier__000001_Growth_Curve_Next_Seed.StartX = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.StartY = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.ControlPointStartX = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.ControlPointStartY = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.EndX = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.EndY = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.ControlPointEndX = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.ControlPointEndY = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.Color = ``
	__Bezier__000001_Growth_Curve_Next_Seed.FillOpacity = 0.000000
	__Bezier__000001_Growth_Curve_Next_Seed.Stroke = `red`
	__Bezier__000001_Growth_Curve_Next_Seed.StrokeOpacity = 0.500000
	__Bezier__000001_Growth_Curve_Next_Seed.StrokeWidth = 6.000000
	__Bezier__000001_Growth_Curve_Next_Seed.StrokeDashArray = ``
	__Bezier__000001_Growth_Curve_Next_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000001_Growth_Curve_Next_Seed.Transform = ``

	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.Name = `Growth Curve Next Shift Right Seed`
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.IsDisplayed = false
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.StartX = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.StartY = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.ControlPointStartX = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.ControlPointStartY = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.EndX = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.EndY = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.ControlPointEndX = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.ControlPointEndY = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.Color = ``
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.FillOpacity = 0.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.Stroke = `red`
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.StrokeOpacity = 0.500000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.StrokeWidth = 6.000000
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.StrokeDashArray = ``
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.Transform = ``

	__Bezier__000003_Growth_Curve_Seed.Name = `Growth Curve Seed`
	__Bezier__000003_Growth_Curve_Seed.IsDisplayed = false
	__Bezier__000003_Growth_Curve_Seed.StartX = -31.464995
	__Bezier__000003_Growth_Curve_Seed.StartY = 180.981680
	__Bezier__000003_Growth_Curve_Seed.ControlPointStartX = 65.086665
	__Bezier__000003_Growth_Curve_Seed.ControlPointStartY = 197.767896
	__Bezier__000003_Growth_Curve_Seed.EndX = 145.265617
	__Bezier__000003_Growth_Curve_Seed.EndY = 398.159696
	__Bezier__000003_Growth_Curve_Seed.ControlPointEndX = 48.713958
	__Bezier__000003_Growth_Curve_Seed.ControlPointEndY = 381.373481
	__Bezier__000003_Growth_Curve_Seed.Color = ``
	__Bezier__000003_Growth_Curve_Seed.FillOpacity = 0.000000
	__Bezier__000003_Growth_Curve_Seed.Stroke = `grey`
	__Bezier__000003_Growth_Curve_Seed.StrokeOpacity = 0.800000
	__Bezier__000003_Growth_Curve_Seed.StrokeWidth = 6.000000
	__Bezier__000003_Growth_Curve_Seed.StrokeDashArray = ``
	__Bezier__000003_Growth_Curve_Seed.StrokeDashArrayWhenSelected = ``
	__Bezier__000003_Growth_Curve_Seed.Transform = ``

	__BezierGrid__000000_Growth_Curve.Name = `Growth Curve`
	__BezierGrid__000000_Growth_Curve.IsDisplayed = true

	__BezierGrid__000001_Growth_Curve_Next.Name = `Growth Curve Next`
	__BezierGrid__000001_Growth_Curve_Next.IsDisplayed = false

	__BezierGrid__000002_Growth_Curve_Next_Shift_Right.Name = `Growth Curve Next Shift Right`
	__BezierGrid__000002_Growth_Curve_Next_Shift_Right.IsDisplayed = false

	__BezierGrid__000003_Growth_Curve_Shift_Right.Name = `Growth Curve Shift Right`
	__BezierGrid__000003_Growth_Curve_Shift_Right.IsDisplayed = true

	__BezierGridStack__000000_The_GrowthCurveStack.Name = `The GrowthCurveStack`
	__BezierGridStack__000000_The_GrowthCurveStack.IsDisplayed = false

	__Circle__000000_Construction_Circle.Name = `Construction Circle`
	__Circle__000000_Construction_Circle.IsDisplayed = false
	__Circle__000000_Construction_Circle.CenterX = -31.464995
	__Circle__000000_Construction_Circle.CenterY = 180.981680
	__Circle__000000_Construction_Circle.HasBespokeRadius = true
	__Circle__000000_Construction_Circle.BespopkeRadius = 20.000000
	__Circle__000000_Construction_Circle.Color = ``
	__Circle__000000_Construction_Circle.FillOpacity = 0.000000
	__Circle__000000_Construction_Circle.Stroke = `blue`
	__Circle__000000_Construction_Circle.StrokeOpacity = 0.800000
	__Circle__000000_Construction_Circle.StrokeWidth = 2.000000
	__Circle__000000_Construction_Circle.StrokeDashArray = ``
	__Circle__000000_Construction_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000000_Construction_Circle.Transform = ``

	__Circle__000001_Growing.Name = `Growing`
	__Circle__000001_Growing.IsDisplayed = false
	__Circle__000001_Growing.CenterX = 0.000000
	__Circle__000001_Growing.CenterY = 0.000000
	__Circle__000001_Growing.HasBespokeRadius = false
	__Circle__000001_Growing.BespopkeRadius = 0.000000
	__Circle__000001_Growing.Color = ``
	__Circle__000001_Growing.FillOpacity = 0.000000
	__Circle__000001_Growing.Stroke = `blue`
	__Circle__000001_Growing.StrokeOpacity = 0.700000
	__Circle__000001_Growing.StrokeWidth = 2.000000
	__Circle__000001_Growing.StrokeDashArray = ``
	__Circle__000001_Growing.StrokeDashArrayWhenSelected = ``
	__Circle__000001_Growing.Transform = ``

	__Circle__000002_Growing_Seed_Left.Name = `Growing Seed Left`
	__Circle__000002_Growing_Seed_Left.IsDisplayed = false
	__Circle__000002_Growing_Seed_Left.CenterX = 0.000000
	__Circle__000002_Growing_Seed_Left.CenterY = 0.000000
	__Circle__000002_Growing_Seed_Left.HasBespokeRadius = false
	__Circle__000002_Growing_Seed_Left.BespopkeRadius = 0.000000
	__Circle__000002_Growing_Seed_Left.Color = ``
	__Circle__000002_Growing_Seed_Left.FillOpacity = 0.000000
	__Circle__000002_Growing_Seed_Left.Stroke = `green`
	__Circle__000002_Growing_Seed_Left.StrokeOpacity = 0.800000
	__Circle__000002_Growing_Seed_Left.StrokeWidth = 1.800000
	__Circle__000002_Growing_Seed_Left.StrokeDashArray = ``
	__Circle__000002_Growing_Seed_Left.StrokeDashArrayWhenSelected = ``
	__Circle__000002_Growing_Seed_Left.Transform = ``

	__Circle__000003_Initial_Circle.Name = `Initial Circle`
	__Circle__000003_Initial_Circle.IsDisplayed = false
	__Circle__000003_Initial_Circle.CenterX = 0.000000
	__Circle__000003_Initial_Circle.CenterY = 0.000000
	__Circle__000003_Initial_Circle.HasBespokeRadius = false
	__Circle__000003_Initial_Circle.BespopkeRadius = 0.000000
	__Circle__000003_Initial_Circle.Color = ``
	__Circle__000003_Initial_Circle.FillOpacity = 0.000000
	__Circle__000003_Initial_Circle.Stroke = `lightblue`
	__Circle__000003_Initial_Circle.StrokeOpacity = 0.800000
	__Circle__000003_Initial_Circle.StrokeWidth = 3.000000
	__Circle__000003_Initial_Circle.StrokeDashArray = ``
	__Circle__000003_Initial_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000003_Initial_Circle.Transform = ``

	__Circle__000004_Rotated_Next_Circle.Name = `Rotated Next Circle`
	__Circle__000004_Rotated_Next_Circle.IsDisplayed = false
	__Circle__000004_Rotated_Next_Circle.CenterX = 416.391214
	__Circle__000004_Rotated_Next_Circle.CenterY = 72.392672
	__Circle__000004_Rotated_Next_Circle.HasBespokeRadius = false
	__Circle__000004_Rotated_Next_Circle.BespopkeRadius = 0.000000
	__Circle__000004_Rotated_Next_Circle.Color = ``
	__Circle__000004_Rotated_Next_Circle.FillOpacity = 0.000000
	__Circle__000004_Rotated_Next_Circle.Stroke = `yellow`
	__Circle__000004_Rotated_Next_Circle.StrokeOpacity = 0.800000
	__Circle__000004_Rotated_Next_Circle.StrokeWidth = 3.000000
	__Circle__000004_Rotated_Next_Circle.StrokeDashArray = ``
	__Circle__000004_Rotated_Next_Circle.StrokeDashArrayWhenSelected = ``
	__Circle__000004_Rotated_Next_Circle.Transform = ``

	__CircleGrid__000000_Construction_Circle_Grid.Name = `Construction Circle Grid`
	__CircleGrid__000000_Construction_Circle_Grid.IsDisplayed = false

	__CircleGrid__000001_Growing_Circle_Grid.Name = `Growing Circle Grid`
	__CircleGrid__000001_Growing_Circle_Grid.IsDisplayed = false

	__CircleGrid__000002_Growing_Circle_Grid_Shifted_Left.Name = `Growing Circle Grid Shifted Left`
	__CircleGrid__000002_Growing_Circle_Grid_Shifted_Left.IsDisplayed = false

	__CircleGrid__000003_Initial_Circle_Grid.Name = `Initial Circle Grid`
	__CircleGrid__000003_Initial_Circle_Grid.IsDisplayed = false

	__CircleGrid__000004_Rotated_Circle_Grid.Name = `Rotated Circle Grid`
	__CircleGrid__000004_Rotated_Circle_Grid.IsDisplayed = false

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
	__Key__000000_F_key.IsDisplayed = true
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
	__Parameter__000000_Reference.N = 3
	__Parameter__000000_Reference.M = 2
	__Parameter__000000_Reference.Z = 9
	__Parameter__000000_Reference.InsideAngle = 98.000000
	__Parameter__000000_Reference.SideLength = 280.000000
	__Parameter__000000_Reference.StackWidth = 3
	__Parameter__000000_Reference.NbShitRight = 2
	__Parameter__000000_Reference.StackHeight = 15
	__Parameter__000000_Reference.BezierControlLengthRatio = 0.350000
	__Parameter__000000_Reference.FkeySizeRatio = 0.001510
	__Parameter__000000_Reference.FkeyOriginRelativeX = 1.500000
	__Parameter__000000_Reference.FkeyOriginRelativeY = -3.600000
	__Parameter__000000_Reference.PitchLinesHeightRatio = 0.000000
	__Parameter__000000_Reference.MeasureLinesHeightRatio = 0.000000
	__Parameter__000000_Reference.OriginX = 200.000000
	__Parameter__000000_Reference.OriginY = 700.000000

	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Name = `Growing Rhombus Grid Seed`
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.IsDisplayed = false
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterX = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.CenterY = 0.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.SideLength = 280.000000
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.Angle = -80.137291
	__Rhombus__000000_Growing_Rhombus_Grid_Seed.InsideAngle = 98.000000
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
	__Rhombus__000001_Initial_Rhombus.SideLength = 280.000000
	__Rhombus__000001_Initial_Rhombus.Angle = 0.000000
	__Rhombus__000001_Initial_Rhombus.InsideAngle = 98.000000
	__Rhombus__000001_Initial_Rhombus.Color = ``
	__Rhombus__000001_Initial_Rhombus.FillOpacity = 0.000000
	__Rhombus__000001_Initial_Rhombus.Stroke = `green`
	__Rhombus__000001_Initial_Rhombus.StrokeOpacity = 0.900000
	__Rhombus__000001_Initial_Rhombus.StrokeWidth = 3.000000
	__Rhombus__000001_Initial_Rhombus.StrokeDashArray = ``
	__Rhombus__000001_Initial_Rhombus.StrokeDashArrayWhenSelected = ``
	__Rhombus__000001_Initial_Rhombus.Transform = ``

	__Rhombus__000002_Rotated_Next_Rhombus.Name = `Rotated Next Rhombus`
	__Rhombus__000002_Rotated_Next_Rhombus.IsDisplayed = false
	__Rhombus__000002_Rotated_Next_Rhombus.CenterX = 416.391214
	__Rhombus__000002_Rotated_Next_Rhombus.CenterY = 72.392672
	__Rhombus__000002_Rotated_Next_Rhombus.SideLength = 280.000000
	__Rhombus__000002_Rotated_Next_Rhombus.Angle = -80.137291
	__Rhombus__000002_Rotated_Next_Rhombus.InsideAngle = 98.000000
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
	__Rhombus__000003_Rotated_Rhombus.SideLength = 280.000000
	__Rhombus__000003_Rotated_Rhombus.Angle = -80.137291
	__Rhombus__000003_Rotated_Rhombus.InsideAngle = 98.000000
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

	__ShapeCategory__000005_5_Growth.Name = `5. Growth`
	__ShapeCategory__000005_5_Growth.IsExpanded = true

	__ShapeCategory__000006_6_Score_notation.Name = `6. Score notation`
	__ShapeCategory__000006_6_Score_notation.IsExpanded = true

	__VerticalAxis__000000_Vertical_Axis.Name = `Vertical Axis`
	__VerticalAxis__000000_Vertical_Axis.IsDisplayed = true
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
	__Axis__000002_Rotated_Axis.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__AxisGrid__000000_Construction_Axis_Grid.Reference = __Axis__000000_Construction_Axis
	__AxisGrid__000000_Construction_Axis_Grid.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Bezier__000000_Growth_Bezier_Right_Seed.ShapeCategory = __ShapeCategory__000005_5_Growth
	__Bezier__000001_Growth_Curve_Next_Seed.ShapeCategory = __ShapeCategory__000005_5_Growth
	__Bezier__000002_Growth_Curve_Next_Shift_Right_Seed.ShapeCategory = __ShapeCategory__000005_5_Growth
	__Bezier__000003_Growth_Curve_Seed.ShapeCategory = __ShapeCategory__000005_5_Growth
	__BezierGrid__000000_Growth_Curve.Reference = __Bezier__000003_Growth_Curve_Seed
	__BezierGrid__000000_Growth_Curve.ShapeCategory = __ShapeCategory__000005_5_Growth
	__BezierGrid__000001_Growth_Curve_Next.Reference = __Bezier__000001_Growth_Curve_Next_Seed
	__BezierGrid__000001_Growth_Curve_Next.ShapeCategory = __ShapeCategory__000005_5_Growth
	__BezierGrid__000002_Growth_Curve_Next_Shift_Right.Reference = __Bezier__000002_Growth_Curve_Next_Shift_Right_Seed
	__BezierGrid__000002_Growth_Curve_Next_Shift_Right.ShapeCategory = __ShapeCategory__000005_5_Growth
	__BezierGrid__000003_Growth_Curve_Shift_Right.Reference = __Bezier__000000_Growth_Bezier_Right_Seed
	__BezierGrid__000003_Growth_Curve_Shift_Right.ShapeCategory = __ShapeCategory__000005_5_Growth
	__BezierGridStack__000000_The_GrowthCurveStack.ShapeCategory = __ShapeCategory__000005_5_Growth
	__Circle__000000_Construction_Circle.ShapeCategory = __ShapeCategory__000004_4_Construction
	__Circle__000003_Initial_Circle.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Circle__000004_Rotated_Next_Circle.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__CircleGrid__000000_Construction_Circle_Grid.Reference = __Circle__000000_Construction_Circle
	__CircleGrid__000000_Construction_Circle_Grid.ShapeCategory = __ShapeCategory__000004_4_Construction
	__CircleGrid__000001_Growing_Circle_Grid.Reference = __Circle__000001_Growing
	__CircleGrid__000001_Growing_Circle_Grid.ShapeCategory = __ShapeCategory__000003_3_Growing
	__CircleGrid__000002_Growing_Circle_Grid_Shifted_Left.ShapeCategory = __ShapeCategory__000003_3_Growing
	__CircleGrid__000003_Initial_Circle_Grid.Reference = __Circle__000003_Initial_Circle
	__CircleGrid__000003_Initial_Circle_Grid.ShapeCategory = __ShapeCategory__000001_1_Initial
	__CircleGrid__000004_Rotated_Circle_Grid.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__HorizontalAxis__000000_Horizontal_Axis.ShapeCategory = __ShapeCategory__000000_0_Axes
	__Key__000000_F_key.ShapeCategory = __ShapeCategory__000006_6_Score_notation
	__Parameter__000000_Reference.InitialRhombus = __Rhombus__000001_Initial_Rhombus
	__Parameter__000000_Reference.InitialCircle = __Circle__000003_Initial_Circle
	__Parameter__000000_Reference.InitialRhombusGrid = __RhombusGrid__000001_Initial_Rhombus_Grid
	__Parameter__000000_Reference.InitialCircleGrid = __CircleGrid__000003_Initial_Circle_Grid
	__Parameter__000000_Reference.InitialAxis = __Axis__000001_Initial_Axis
	__Parameter__000000_Reference.RotatedAxis = __Axis__000002_Rotated_Axis
	__Parameter__000000_Reference.RotatedRhombus = __Rhombus__000003_Rotated_Rhombus
	__Parameter__000000_Reference.RotatedRhombusGrid = __RhombusGrid__000002_Rotated_Rhombus_Grid
	__Parameter__000000_Reference.RotatedCircleGrid = __CircleGrid__000004_Rotated_Circle_Grid
	__Parameter__000000_Reference.NextRhombus = __Rhombus__000002_Rotated_Next_Rhombus
	__Parameter__000000_Reference.NextCircle = __Circle__000004_Rotated_Next_Circle
	__Parameter__000000_Reference.GrowingRhombusGridSeed = __Rhombus__000000_Growing_Rhombus_Grid_Seed
	__Parameter__000000_Reference.GrowingRhombusGrid = __RhombusGrid__000000_Growing_Rhombus_Grid
	__Parameter__000000_Reference.GrowingCircleGridSeed = __Circle__000001_Growing
	__Parameter__000000_Reference.GrowingCircleGrid = __CircleGrid__000001_Growing_Circle_Grid
	__Parameter__000000_Reference.GrowingCircleGridLeftSeed = __Circle__000002_Growing_Seed_Left
	__Parameter__000000_Reference.GrowingCircleGridLeft = __CircleGrid__000002_Growing_Circle_Grid_Shifted_Left
	__Parameter__000000_Reference.ConstructionAxis = __Axis__000000_Construction_Axis
	__Parameter__000000_Reference.ConstructionAxisGrid = __AxisGrid__000000_Construction_Axis_Grid
	__Parameter__000000_Reference.ConstructionCircle = __Circle__000000_Construction_Circle
	__Parameter__000000_Reference.ConstructionCircleGrid = __CircleGrid__000000_Construction_Circle_Grid
	__Parameter__000000_Reference.GrowthCurveSegment = __Bezier__000003_Growth_Curve_Seed
	__Parameter__000000_Reference.GrowthCurve = __BezierGrid__000000_Growth_Curve
	__Parameter__000000_Reference.GrowthCurveShiftedRightSeed = __Bezier__000000_Growth_Bezier_Right_Seed
	__Parameter__000000_Reference.GrowthCurveShiftedRight = __BezierGrid__000003_Growth_Curve_Shift_Right
	__Parameter__000000_Reference.GrowthCurveNextSeed = __Bezier__000001_Growth_Curve_Next_Seed
	__Parameter__000000_Reference.GrowthCurveNext = __BezierGrid__000001_Growth_Curve_Next
	__Parameter__000000_Reference.GrowthCurveNextShiftedRightSeed = __Bezier__000002_Growth_Curve_Next_Shift_Right_Seed
	__Parameter__000000_Reference.GrowthCurveNextShiftedRight = __BezierGrid__000002_Growth_Curve_Next_Shift_Right
	__Parameter__000000_Reference.GrowthCurveStack = __BezierGridStack__000000_The_GrowthCurveStack
	__Parameter__000000_Reference.Fkey = __Key__000000_F_key
	__Parameter__000000_Reference.HorizontalAxis = __HorizontalAxis__000000_Horizontal_Axis
	__Parameter__000000_Reference.VerticalAxis = __VerticalAxis__000000_Vertical_Axis
	__Rhombus__000001_Initial_Rhombus.ShapeCategory = __ShapeCategory__000001_1_Initial
	__Rhombus__000002_Rotated_Next_Rhombus.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__Rhombus__000003_Rotated_Rhombus.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__RhombusGrid__000000_Growing_Rhombus_Grid.Reference = __Rhombus__000003_Rotated_Rhombus
	__RhombusGrid__000000_Growing_Rhombus_Grid.ShapeCategory = __ShapeCategory__000003_3_Growing
	__RhombusGrid__000001_Initial_Rhombus_Grid.Reference = __Rhombus__000001_Initial_Rhombus
	__RhombusGrid__000001_Initial_Rhombus_Grid.ShapeCategory = __ShapeCategory__000001_1_Initial
	__RhombusGrid__000002_Rotated_Rhombus_Grid.ShapeCategory = __ShapeCategory__000002_2_Rotated
	__VerticalAxis__000000_Vertical_Axis.ShapeCategory = __ShapeCategory__000000_0_Axes
}
