package main

import (
	"time"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Axis__00000000_ := (&models.Axis{Name: `Beat Reference`}).Stage(stage)
	__Axis__00000001_ := (&models.Axis{Name: `Construction Axis`}).Stage(stage)
	__Axis__00000002_ := (&models.Axis{Name: `Initial Axis`}).Stage(stage)
	__Axis__00000003_ := (&models.Axis{Name: `Pitch Line`}).Stage(stage)
	__Axis__00000004_ := (&models.Axis{Name: `Rotated Axis`}).Stage(stage)

	__AxisGrid__00000000_ := (&models.AxisGrid{Name: `Beat Lines`}).Stage(stage)
	__AxisGrid__00000001_ := (&models.AxisGrid{Name: `Construction Axis Grid`}).Stage(stage)
	__AxisGrid__00000002_ := (&models.AxisGrid{Name: `Pitch Lines`}).Stage(stage)

	__Bezier__00000000_ := (&models.Bezier{Name: `2nd voice seed`}).Stage(stage)
	__Bezier__00000001_ := (&models.Bezier{Name: `First Voice seed`}).Stage(stage)
	__Bezier__00000002_ := (&models.Bezier{Name: `Growth Bezier Right Seed`}).Stage(stage)
	__Bezier__00000003_ := (&models.Bezier{Name: `Growth Curve Next Seed`}).Stage(stage)
	__Bezier__00000004_ := (&models.Bezier{Name: `Growth Curve Next Shift Right Seed`}).Stage(stage)
	__Bezier__00000005_ := (&models.Bezier{Name: `Growth Curve Seed`}).Stage(stage)
	__Bezier__00000006_ := (&models.Bezier{Name: `FirstVoiceShiftedRightSeed`}).Stage(stage)

	__BezierGrid__00000000_ := (&models.BezierGrid{Name: `2nb Voice`}).Stage(stage)
	__BezierGrid__00000001_ := (&models.BezierGrid{Name: `2nd voice shifted right`}).Stage(stage)
	__BezierGrid__00000002_ := (&models.BezierGrid{Name: `First Voice`}).Stage(stage)
	__BezierGrid__00000003_ := (&models.BezierGrid{Name: `First Voice Shift Right`}).Stage(stage)
	__BezierGrid__00000004_ := (&models.BezierGrid{Name: `Growth Curve`}).Stage(stage)
	__BezierGrid__00000005_ := (&models.BezierGrid{Name: `Growth Curve Next`}).Stage(stage)
	__BezierGrid__00000006_ := (&models.BezierGrid{Name: `Growth Curve Next Shift Right`}).Stage(stage)
	__BezierGrid__00000007_ := (&models.BezierGrid{Name: `Growth Curve Shift Right`}).Stage(stage)

	__BezierGridStack__00000000_ := (&models.BezierGridStack{Name: `The GrowthCurveStack`}).Stage(stage)

	__Chapter__00000000_ := (&models.Chapter{Name: `Deep into the phyllotaxy growth curve`}).Stage(stage)
	__Chapter__00000001_ := (&models.Chapter{Name: `Composing your own phyllotaxy music`}).Stage(stage)

	__Circle__00000000_ := (&models.Circle{Name: `0`}).Stage(stage)
	__Circle__00000001_ := (&models.Circle{Name: `Construction Circle`}).Stage(stage)
	__Circle__00000002_ := (&models.Circle{Name: `First voice notes seed`}).Stage(stage)
	__Circle__00000003_ := (&models.Circle{Name: `Growing Seed Left`}).Stage(stage)
	__Circle__00000004_ := (&models.Circle{Name: `Initial Circle`}).Stage(stage)
	__Circle__00000005_ := (&models.Circle{Name: `Rotated Next Circle`}).Stage(stage)

	__CircleGrid__00000000_ := (&models.CircleGrid{Name: `Construction Circle Grid`}).Stage(stage)
	__CircleGrid__00000001_ := (&models.CircleGrid{Name: `First Voice note shifted right`}).Stage(stage)
	__CircleGrid__00000002_ := (&models.CircleGrid{Name: `First Voice notes`}).Stage(stage)
	__CircleGrid__00000003_ := (&models.CircleGrid{Name: `Growing Circle Grid`}).Stage(stage)
	__CircleGrid__00000004_ := (&models.CircleGrid{Name: `Growing Circle Grid Shifted Left`}).Stage(stage)
	__CircleGrid__00000005_ := (&models.CircleGrid{Name: `Initial Circle Grid`}).Stage(stage)
	__CircleGrid__00000006_ := (&models.CircleGrid{Name: `Rotated Circle Grid`}).Stage(stage)
	__CircleGrid__00000007_ := (&models.CircleGrid{Name: `Second Voice Notes Shift Right`}).Stage(stage)
	__CircleGrid__00000008_ := (&models.CircleGrid{Name: `Second Voice notes`}).Stage(stage)

	__Content__00000000_ := (&models.Content{Name: `Phyllotaxy Music`}).Stage(stage)

	__ExportToMusicxml__00000000_ := (&models.ExportToMusicxml{Name: `Singloton`}).Stage(stage)

	__FrontCurve__00000008_ := (&models.FrontCurve{Name: `Non Rotated 0 `}).Stage(stage)
	__FrontCurve__00000009_ := (&models.FrontCurve{Name: `Non Rotated 1 `}).Stage(stage)
	__FrontCurve__00000010_ := (&models.FrontCurve{Name: `Non Rotated 2 `}).Stage(stage)
	__FrontCurve__00000011_ := (&models.FrontCurve{Name: `Non Rotated 3 `}).Stage(stage)
	__FrontCurve__00000012_ := (&models.FrontCurve{Name: `Non Rotated 4 `}).Stage(stage)
	__FrontCurve__00000013_ := (&models.FrontCurve{Name: `Non Rotated 5 `}).Stage(stage)
	__FrontCurve__00000014_ := (&models.FrontCurve{Name: `Non Rotated 6 `}).Stage(stage)
	__FrontCurve__00000015_ := (&models.FrontCurve{Name: `Non Rotated 7 `}).Stage(stage)

	__FrontCurveStack__00000000_ := (&models.FrontCurveStack{Name: `Front Curve Stack`}).Stage(stage)

	__HorizontalAxis__00000000_ := (&models.HorizontalAxis{Name: `Horizontal Axis`}).Stage(stage)

	__Key__00000000_ := (&models.Key{Name: `F key`}).Stage(stage)

	__Parameter__00000000_ := (&models.Parameter{Name: `Reference`}).Stage(stage)

	__Rhombus__00000000_ := (&models.Rhombus{Name: `Growing Rhombus Grid Seed`}).Stage(stage)
	__Rhombus__00000001_ := (&models.Rhombus{Name: `Initial Rhombus`}).Stage(stage)
	__Rhombus__00000002_ := (&models.Rhombus{Name: `Rotated Next Rhombus`}).Stage(stage)
	__Rhombus__00000003_ := (&models.Rhombus{Name: `Rotated Rhombus`}).Stage(stage)

	__RhombusGrid__00000000_ := (&models.RhombusGrid{Name: `Growing Rhombus Grid`}).Stage(stage)
	__RhombusGrid__00000001_ := (&models.RhombusGrid{Name: `Initial Rhombus Grid`}).Stage(stage)
	__RhombusGrid__00000002_ := (&models.RhombusGrid{Name: `Rotated Rhombus Grid`}).Stage(stage)

	__ShapeCategory__00000000_ := (&models.ShapeCategory{Name: `0. Axes`}).Stage(stage)
	__ShapeCategory__00000001_ := (&models.ShapeCategory{Name: `1. Initial`}).Stage(stage)
	__ShapeCategory__00000002_ := (&models.ShapeCategory{Name: `2. Rotated`}).Stage(stage)
	__ShapeCategory__00000003_ := (&models.ShapeCategory{Name: `3. Growing`}).Stage(stage)
	__ShapeCategory__00000004_ := (&models.ShapeCategory{Name: `4. Construction`}).Stage(stage)
	__ShapeCategory__00000005_ := (&models.ShapeCategory{Name: `5. Vertical growth`}).Stage(stage)
	__ShapeCategory__00000006_ := (&models.ShapeCategory{Name: `6. Spiral growth`}).Stage(stage)
	__ShapeCategory__00000007_ := (&models.ShapeCategory{Name: `7. Spiral Growth Bezier`}).Stage(stage)
	__ShapeCategory__00000008_ := (&models.ShapeCategory{Name: `8. Score notation`}).Stage(stage)
	__ShapeCategory__00000009_ := (&models.ShapeCategory{Name: `9. Composer`}).Stage(stage)

	__SpiralBezier__00000000_ := (&models.SpiralBezier{Name: `Spiral Bezier Seed`}).Stage(stage)

	__SpiralBezierGrid__00000000_ := (&models.SpiralBezierGrid{Name: `Spiral Bezier Full Grid`}).Stage(stage)
	__SpiralBezierGrid__00000001_ := (&models.SpiralBezierGrid{Name: `Spiral Bezier Grid`}).Stage(stage)

	__SpiralCircle__00000000_ := (&models.SpiralCircle{}).Stage(stage)

	__SpiralCircleGrid__00000000_ := (&models.SpiralCircleGrid{Name: `Brute Spiral Bezier Circle Grid`}).Stage(stage)
	__SpiralCircleGrid__00000001_ := (&models.SpiralCircleGrid{Name: `Construction Circle Spiral Full Grid`}).Stage(stage)
	__SpiralCircleGrid__00000002_ := (&models.SpiralCircleGrid{Name: `Construction Circle Spiral Grid`}).Stage(stage)
	__SpiralCircleGrid__00000003_ := (&models.SpiralCircleGrid{Name: `Spiral Circle Grid`}).Stage(stage)

	__SpiralLine__00000000_ := (&models.SpiralLine{Name: `Spiral Contruction Inner Line`}).Stage(stage)
	__SpiralLine__00000001_ := (&models.SpiralLine{Name: `Spiral Contruction Outer Line`}).Stage(stage)

	__SpiralLineGrid__00000000_ := (&models.SpiralLineGrid{Name: `Spiral Construction Inner Line Grid Spiral`}).Stage(stage)
	__SpiralLineGrid__00000001_ := (&models.SpiralLineGrid{Name: `Spiral Construction Outer Line Full Grid Spiral`}).Stage(stage)
	__SpiralLineGrid__00000002_ := (&models.SpiralLineGrid{Name: `Spiral Construction Outer Line Grid Spiral`}).Stage(stage)

	__SpiralOrigin__00000000_ := (&models.SpiralOrigin{Name: `Spiral Origin`}).Stage(stage)

	__SpiralRhombus__00000000_ := (&models.SpiralRhombus{Name: `Reference Spiral Rhombus`}).Stage(stage)

	__SpiralRhombusGrid__00000000_ := (&models.SpiralRhombusGrid{Name: `Spiral Rhombus Grid`}).Stage(stage)

	__VerticalAxis__00000000_ := (&models.VerticalAxis{Name: `Vertical Axis`}).Stage(stage)

	// insertion point for initialization of values

	__Axis__00000000_.Name = `Beat Reference`
	__Axis__00000000_.IsDisplayed = false
	__Axis__00000000_.AngleDegree = 90.000000
	__Axis__00000000_.Length = 2000.000000
	__Axis__00000000_.CenterX = 0.000000
	__Axis__00000000_.CenterY = 0.000000
	__Axis__00000000_.EndX = 0.000000
	__Axis__00000000_.EndY = 0.000000
	__Axis__00000000_.Color = ``
	__Axis__00000000_.FillOpacity = 0.000000
	__Axis__00000000_.Stroke = `grey`
	__Axis__00000000_.StrokeOpacity = 0.800000
	__Axis__00000000_.StrokeWidth = 1.000000
	__Axis__00000000_.StrokeDashArray = ``
	__Axis__00000000_.StrokeDashArrayWhenSelected = ``
	__Axis__00000000_.Transform = ``

	__Axis__00000001_.Name = `Construction Axis`
	__Axis__00000001_.IsDisplayed = false
	__Axis__00000001_.AngleDegree = 96.586776
	__Axis__00000001_.Length = 90.000000
	__Axis__00000001_.CenterX = 0.000000
	__Axis__00000001_.CenterY = 0.000000
	__Axis__00000001_.EndX = -10.323708
	__Axis__00000001_.EndY = 89.405934
	__Axis__00000001_.Color = ``
	__Axis__00000001_.FillOpacity = 0.000000
	__Axis__00000001_.Stroke = `blue`
	__Axis__00000001_.StrokeOpacity = 0.700000
	__Axis__00000001_.StrokeWidth = 2.000000
	__Axis__00000001_.StrokeDashArray = ``
	__Axis__00000001_.StrokeDashArrayWhenSelected = ``
	__Axis__00000001_.Transform = ``

	__Axis__00000002_.Name = `Initial Axis`
	__Axis__00000002_.IsDisplayed = false
	__Axis__00000002_.AngleDegree = 83.413224
	__Axis__00000002_.Length = 392.300905
	__Axis__00000002_.CenterX = 0.000000
	__Axis__00000002_.CenterY = 0.000000
	__Axis__00000002_.EndX = 0.000000
	__Axis__00000002_.EndY = 0.000000
	__Axis__00000002_.Color = ``
	__Axis__00000002_.FillOpacity = 0.000000
	__Axis__00000002_.Stroke = `black`
	__Axis__00000002_.StrokeOpacity = 1.000000
	__Axis__00000002_.StrokeWidth = 2.000000
	__Axis__00000002_.StrokeDashArray = ``
	__Axis__00000002_.StrokeDashArrayWhenSelected = ``
	__Axis__00000002_.Transform = ``

	__Axis__00000003_.Name = `Pitch Line`
	__Axis__00000003_.IsDisplayed = false
	__Axis__00000003_.AngleDegree = 0.000000
	__Axis__00000003_.Length = 2000.000000
	__Axis__00000003_.CenterX = 0.000000
	__Axis__00000003_.CenterY = 0.000000
	__Axis__00000003_.EndX = 0.000000
	__Axis__00000003_.EndY = 0.000000
	__Axis__00000003_.Color = ``
	__Axis__00000003_.FillOpacity = 0.000000
	__Axis__00000003_.Stroke = `grey`
	__Axis__00000003_.StrokeOpacity = 0.800000
	__Axis__00000003_.StrokeWidth = 1.000000
	__Axis__00000003_.StrokeDashArray = ``
	__Axis__00000003_.StrokeDashArrayWhenSelected = ``
	__Axis__00000003_.Transform = ``

	__Axis__00000004_.Name = `Rotated Axis`
	__Axis__00000004_.IsDisplayed = false
	__Axis__00000004_.AngleDegree = 0.000000
	__Axis__00000004_.Length = 392.300905
	__Axis__00000004_.CenterX = 0.000000
	__Axis__00000004_.CenterY = 0.000000
	__Axis__00000004_.EndX = -29.670211
	__Axis__00000004_.EndY = 154.170939
	__Axis__00000004_.Color = ``
	__Axis__00000004_.FillOpacity = 0.000000
	__Axis__00000004_.Stroke = `black`
	__Axis__00000004_.StrokeOpacity = 1.000000
	__Axis__00000004_.StrokeWidth = 2.000000
	__Axis__00000004_.StrokeDashArray = ``
	__Axis__00000004_.StrokeDashArrayWhenSelected = ``
	__Axis__00000004_.Transform = ``

	__AxisGrid__00000000_.Name = `Beat Lines`
	__AxisGrid__00000000_.IsDisplayed = true

	__AxisGrid__00000001_.Name = `Construction Axis Grid`
	__AxisGrid__00000001_.IsDisplayed = false

	__AxisGrid__00000002_.Name = `Pitch Lines`
	__AxisGrid__00000002_.IsDisplayed = true

	__Bezier__00000000_.Name = `2nd voice seed`
	__Bezier__00000000_.IsDisplayed = false
	__Bezier__00000000_.StartX = 0.000000
	__Bezier__00000000_.StartY = 0.000000
	__Bezier__00000000_.ControlPointStartX = 0.000000
	__Bezier__00000000_.ControlPointStartY = 0.000000
	__Bezier__00000000_.EndX = 0.000000
	__Bezier__00000000_.EndY = 0.000000
	__Bezier__00000000_.ControlPointEndX = 0.000000
	__Bezier__00000000_.ControlPointEndY = 0.000000
	__Bezier__00000000_.Color = ``
	__Bezier__00000000_.FillOpacity = 0.000000
	__Bezier__00000000_.Stroke = `red`
	__Bezier__00000000_.StrokeOpacity = 0.800000
	__Bezier__00000000_.StrokeWidth = 5.000000
	__Bezier__00000000_.StrokeDashArray = ``
	__Bezier__00000000_.StrokeDashArrayWhenSelected = ``
	__Bezier__00000000_.Transform = ``

	__Bezier__00000001_.Name = `First Voice seed`
	__Bezier__00000001_.IsDisplayed = false
	__Bezier__00000001_.StartX = 0.000000
	__Bezier__00000001_.StartY = 0.000000
	__Bezier__00000001_.ControlPointStartX = 0.000000
	__Bezier__00000001_.ControlPointStartY = 0.000000
	__Bezier__00000001_.EndX = 0.000000
	__Bezier__00000001_.EndY = 0.000000
	__Bezier__00000001_.ControlPointEndX = 0.000000
	__Bezier__00000001_.ControlPointEndY = 0.000000
	__Bezier__00000001_.Color = ``
	__Bezier__00000001_.FillOpacity = 0.000000
	__Bezier__00000001_.Stroke = `grey`
	__Bezier__00000001_.StrokeOpacity = 0.800000
	__Bezier__00000001_.StrokeWidth = 6.000000
	__Bezier__00000001_.StrokeDashArray = ``
	__Bezier__00000001_.StrokeDashArrayWhenSelected = ``
	__Bezier__00000001_.Transform = ``

	__Bezier__00000002_.Name = `Growth Bezier Right Seed`
	__Bezier__00000002_.IsDisplayed = false
	__Bezier__00000002_.StartX = 0.000000
	__Bezier__00000002_.StartY = 0.000000
	__Bezier__00000002_.ControlPointStartX = 0.000000
	__Bezier__00000002_.ControlPointStartY = 0.000000
	__Bezier__00000002_.EndX = 0.000000
	__Bezier__00000002_.EndY = 0.000000
	__Bezier__00000002_.ControlPointEndX = 0.000000
	__Bezier__00000002_.ControlPointEndY = 0.000000
	__Bezier__00000002_.Color = ``
	__Bezier__00000002_.FillOpacity = 0.000000
	__Bezier__00000002_.Stroke = `green`
	__Bezier__00000002_.StrokeOpacity = 0.500000
	__Bezier__00000002_.StrokeWidth = 6.000000
	__Bezier__00000002_.StrokeDashArray = ``
	__Bezier__00000002_.StrokeDashArrayWhenSelected = ``
	__Bezier__00000002_.Transform = ``

	__Bezier__00000003_.Name = `Growth Curve Next Seed`
	__Bezier__00000003_.IsDisplayed = false
	__Bezier__00000003_.StartX = 0.000000
	__Bezier__00000003_.StartY = 0.000000
	__Bezier__00000003_.ControlPointStartX = 0.000000
	__Bezier__00000003_.ControlPointStartY = 0.000000
	__Bezier__00000003_.EndX = 0.000000
	__Bezier__00000003_.EndY = 0.000000
	__Bezier__00000003_.ControlPointEndX = 0.000000
	__Bezier__00000003_.ControlPointEndY = 0.000000
	__Bezier__00000003_.Color = ``
	__Bezier__00000003_.FillOpacity = 0.000000
	__Bezier__00000003_.Stroke = `red`
	__Bezier__00000003_.StrokeOpacity = 0.500000
	__Bezier__00000003_.StrokeWidth = 6.000000
	__Bezier__00000003_.StrokeDashArray = ``
	__Bezier__00000003_.StrokeDashArrayWhenSelected = ``
	__Bezier__00000003_.Transform = ``

	__Bezier__00000004_.Name = `Growth Curve Next Shift Right Seed`
	__Bezier__00000004_.IsDisplayed = false
	__Bezier__00000004_.StartX = 0.000000
	__Bezier__00000004_.StartY = 0.000000
	__Bezier__00000004_.ControlPointStartX = 0.000000
	__Bezier__00000004_.ControlPointStartY = 0.000000
	__Bezier__00000004_.EndX = 0.000000
	__Bezier__00000004_.EndY = 0.000000
	__Bezier__00000004_.ControlPointEndX = 0.000000
	__Bezier__00000004_.ControlPointEndY = 0.000000
	__Bezier__00000004_.Color = ``
	__Bezier__00000004_.FillOpacity = 0.000000
	__Bezier__00000004_.Stroke = `red`
	__Bezier__00000004_.StrokeOpacity = 0.500000
	__Bezier__00000004_.StrokeWidth = 6.000000
	__Bezier__00000004_.StrokeDashArray = ``
	__Bezier__00000004_.StrokeDashArrayWhenSelected = ``
	__Bezier__00000004_.Transform = ``

	__Bezier__00000005_.Name = `Growth Curve Seed`
	__Bezier__00000005_.IsDisplayed = false
	__Bezier__00000005_.StartX = -5.161854
	__Bezier__00000005_.StartY = 44.702967
	__Bezier__00000005_.ControlPointStartX = 44.905469
	__Bezier__00000005_.ControlPointStartY = 50.484244
	__Bezier__00000005_.EndX = 67.104102
	__Bezier__00000005_.EndY = 98.346528
	__Bezier__00000005_.ControlPointEndX = 17.036779
	__Bezier__00000005_.ControlPointEndY = 92.565251
	__Bezier__00000005_.Color = ``
	__Bezier__00000005_.FillOpacity = 0.000000
	__Bezier__00000005_.Stroke = `grey`
	__Bezier__00000005_.StrokeOpacity = 0.800000
	__Bezier__00000005_.StrokeWidth = 6.000000
	__Bezier__00000005_.StrokeDashArray = ``
	__Bezier__00000005_.StrokeDashArrayWhenSelected = ``
	__Bezier__00000005_.Transform = ``

	__Bezier__00000006_.Name = `FirstVoiceShiftedRightSeed`
	__Bezier__00000006_.IsDisplayed = false
	__Bezier__00000006_.StartX = 0.000000
	__Bezier__00000006_.StartY = 0.000000
	__Bezier__00000006_.ControlPointStartX = 0.000000
	__Bezier__00000006_.ControlPointStartY = 0.000000
	__Bezier__00000006_.EndX = 0.000000
	__Bezier__00000006_.EndY = 0.000000
	__Bezier__00000006_.ControlPointEndX = 0.000000
	__Bezier__00000006_.ControlPointEndY = 0.000000
	__Bezier__00000006_.Color = ``
	__Bezier__00000006_.FillOpacity = 0.000000
	__Bezier__00000006_.Stroke = `lightgreen`
	__Bezier__00000006_.StrokeOpacity = 0.800000
	__Bezier__00000006_.StrokeWidth = 6.000000
	__Bezier__00000006_.StrokeDashArray = ``
	__Bezier__00000006_.StrokeDashArrayWhenSelected = ``
	__Bezier__00000006_.Transform = ``

	__BezierGrid__00000000_.Name = `2nb Voice`
	__BezierGrid__00000000_.IsDisplayed = true

	__BezierGrid__00000001_.Name = `2nd voice shifted right`
	__BezierGrid__00000001_.IsDisplayed = true

	__BezierGrid__00000002_.Name = `First Voice`
	__BezierGrid__00000002_.IsDisplayed = true

	__BezierGrid__00000003_.Name = `First Voice Shift Right`
	__BezierGrid__00000003_.IsDisplayed = true

	__BezierGrid__00000004_.Name = `Growth Curve`
	__BezierGrid__00000004_.IsDisplayed = false

	__BezierGrid__00000005_.Name = `Growth Curve Next`
	__BezierGrid__00000005_.IsDisplayed = false

	__BezierGrid__00000006_.Name = `Growth Curve Next Shift Right`
	__BezierGrid__00000006_.IsDisplayed = false

	__BezierGrid__00000007_.Name = `Growth Curve Shift Right`
	__BezierGrid__00000007_.IsDisplayed = false

	__BezierGridStack__00000000_.Name = `The GrowthCurveStack`
	__BezierGridStack__00000000_.IsDisplayed = false

	__Chapter__00000000_.Name = `Deep into the phyllotaxy growth curve`
	__Chapter__00000000_.MardownContent = `To be completed`

	__Chapter__00000001_.Name = `Composing your own phyllotaxy music`
	__Chapter__00000001_.MardownContent = `To be completed`

	__Circle__00000000_.Name = `0`
	__Circle__00000000_.IsDisplayed = false
	__Circle__00000000_.CenterX = 0.000000
	__Circle__00000000_.CenterY = 0.000000
	__Circle__00000000_.HasBespokeRadius = false
	__Circle__00000000_.BespopkeRadius = 0.000000
	__Circle__00000000_.Color = ``
	__Circle__00000000_.FillOpacity = 0.000000
	__Circle__00000000_.Stroke = `blue`
	__Circle__00000000_.StrokeOpacity = 0.700000
	__Circle__00000000_.StrokeWidth = 2.000000
	__Circle__00000000_.StrokeDashArray = ``
	__Circle__00000000_.StrokeDashArrayWhenSelected = ``
	__Circle__00000000_.Transform = ``
	__Circle__00000000_.Pitch = 0
	__Circle__00000000_.ShowName = false
	__Circle__00000000_.BeatNb = 0

	__Circle__00000001_.Name = `Construction Circle`
	__Circle__00000001_.IsDisplayed = false
	__Circle__00000001_.CenterX = -5.161854
	__Circle__00000001_.CenterY = 44.702967
	__Circle__00000001_.HasBespokeRadius = true
	__Circle__00000001_.BespopkeRadius = 20.000000
	__Circle__00000001_.Color = ``
	__Circle__00000001_.FillOpacity = 0.000000
	__Circle__00000001_.Stroke = `blue`
	__Circle__00000001_.StrokeOpacity = 0.800000
	__Circle__00000001_.StrokeWidth = 2.000000
	__Circle__00000001_.StrokeDashArray = ``
	__Circle__00000001_.StrokeDashArrayWhenSelected = ``
	__Circle__00000001_.Transform = ``
	__Circle__00000001_.Pitch = 0
	__Circle__00000001_.ShowName = false
	__Circle__00000001_.BeatNb = 0

	__Circle__00000002_.Name = `First voice notes seed`
	__Circle__00000002_.IsDisplayed = false
	__Circle__00000002_.CenterX = 0.000000
	__Circle__00000002_.CenterY = 0.000000
	__Circle__00000002_.HasBespokeRadius = true
	__Circle__00000002_.BespopkeRadius = 10.000000
	__Circle__00000002_.Color = ``
	__Circle__00000002_.FillOpacity = 0.000000
	__Circle__00000002_.Stroke = `grey`
	__Circle__00000002_.StrokeOpacity = 0.800000
	__Circle__00000002_.StrokeWidth = 3.000000
	__Circle__00000002_.StrokeDashArray = ``
	__Circle__00000002_.StrokeDashArrayWhenSelected = ``
	__Circle__00000002_.Transform = ``
	__Circle__00000002_.Pitch = 0
	__Circle__00000002_.ShowName = false
	__Circle__00000002_.BeatNb = 0

	__Circle__00000003_.Name = `Growing Seed Left`
	__Circle__00000003_.IsDisplayed = false
	__Circle__00000003_.CenterX = 0.000000
	__Circle__00000003_.CenterY = 0.000000
	__Circle__00000003_.HasBespokeRadius = false
	__Circle__00000003_.BespopkeRadius = 0.000000
	__Circle__00000003_.Color = ``
	__Circle__00000003_.FillOpacity = 0.000000
	__Circle__00000003_.Stroke = `green`
	__Circle__00000003_.StrokeOpacity = 0.800000
	__Circle__00000003_.StrokeWidth = 1.800000
	__Circle__00000003_.StrokeDashArray = ``
	__Circle__00000003_.StrokeDashArrayWhenSelected = ``
	__Circle__00000003_.Transform = ``
	__Circle__00000003_.Pitch = 0
	__Circle__00000003_.ShowName = false
	__Circle__00000003_.BeatNb = 0

	__Circle__00000004_.Name = `Initial Circle`
	__Circle__00000004_.IsDisplayed = false
	__Circle__00000004_.CenterX = 0.000000
	__Circle__00000004_.CenterY = 0.000000
	__Circle__00000004_.HasBespokeRadius = false
	__Circle__00000004_.BespopkeRadius = 0.000000
	__Circle__00000004_.Color = ``
	__Circle__00000004_.FillOpacity = 0.000000
	__Circle__00000004_.Stroke = `lightblue`
	__Circle__00000004_.StrokeOpacity = 0.800000
	__Circle__00000004_.StrokeWidth = 3.000000
	__Circle__00000004_.StrokeDashArray = ``
	__Circle__00000004_.StrokeDashArrayWhenSelected = ``
	__Circle__00000004_.Transform = ``
	__Circle__00000004_.Pitch = 0
	__Circle__00000004_.ShowName = false
	__Circle__00000004_.BeatNb = 0

	__Circle__00000005_.Name = `Rotated Next Circle`
	__Circle__00000005_.IsDisplayed = false
	__Circle__00000005_.CenterX = 154.855620
	__Circle__00000005_.CenterY = 17.881187
	__Circle__00000005_.HasBespokeRadius = false
	__Circle__00000005_.BespopkeRadius = 0.000000
	__Circle__00000005_.Color = ``
	__Circle__00000005_.FillOpacity = 0.000000
	__Circle__00000005_.Stroke = `yellow`
	__Circle__00000005_.StrokeOpacity = 0.800000
	__Circle__00000005_.StrokeWidth = 3.000000
	__Circle__00000005_.StrokeDashArray = ``
	__Circle__00000005_.StrokeDashArrayWhenSelected = ``
	__Circle__00000005_.Transform = ``
	__Circle__00000005_.Pitch = 0
	__Circle__00000005_.ShowName = false
	__Circle__00000005_.BeatNb = 0

	__CircleGrid__00000000_.Name = `Construction Circle Grid`
	__CircleGrid__00000000_.IsDisplayed = false

	__CircleGrid__00000001_.Name = `First Voice note shifted right`
	__CircleGrid__00000001_.IsDisplayed = true

	__CircleGrid__00000002_.Name = `First Voice notes`
	__CircleGrid__00000002_.IsDisplayed = true

	__CircleGrid__00000003_.Name = `Growing Circle Grid`
	__CircleGrid__00000003_.IsDisplayed = false

	__CircleGrid__00000004_.Name = `Growing Circle Grid Shifted Left`
	__CircleGrid__00000004_.IsDisplayed = false

	__CircleGrid__00000005_.Name = `Initial Circle Grid`
	__CircleGrid__00000005_.IsDisplayed = false

	__CircleGrid__00000006_.Name = `Rotated Circle Grid`
	__CircleGrid__00000006_.IsDisplayed = false

	__CircleGrid__00000007_.Name = `Second Voice Notes Shift Right`
	__CircleGrid__00000007_.IsDisplayed = true

	__CircleGrid__00000008_.Name = `Second Voice notes`
	__CircleGrid__00000008_.IsDisplayed = true

	__Content__00000000_.Name = `Phyllotaxy Music`
	__Content__00000000_.MardownContent = `This site describes phyllotaxy generated music.

Phyllotaxy and music are two concepts far appart but they can be related.

Precisely, phyllotaxy can help to generates a music
theme, a suite of music notes, a melodic material.

### Cannon

A well known music theme is at the start of Bach's 2nd 
fugue in C minor. It starts with the C note.

<img src="./images/bach2ndFugue.png" style="height: 100px; width: auto;">

In this piece, the theme is repeated with two voices that overlaps.

<img src="./images/bach2ndFugue_large.png" style="height: 400px; width: auto;">

The first voice 
starts with a C on the first measure. On the third measure, the theme starts again, but it starts with a G. 
The G of the second voice is seven half tone higher, or a fifth in music parlance, above the
 C of the first voice. A fifth is an a well suited gap between notes to sounds harmoniously. 
The next note, a F# on the second voice, is played along a A on the first voice, also a
third. A third,  And so one.

It is the composer's chore
to have the theme well suited for supporting overlapping itself.

Having a theme that repeats itself harmoniously with a shift in beats and tone is called a cannon:

- Pachelbel's Canon in D
- "Row, Row, Row Your Boat"

Phyllotaxy will help us to compose theme that are suited for cannon with the following idea : 
if a cannon theme is a curve on a plan, it must overlap with itself with an horizontal and a vertical shift. 

### Phyllotaxy

Phyllotaxy means "shape of leaves" 
("phyllo" is greek for leave and "taxy" is greek for shape). 
It is the 
science developed by botanist, 
mathematicians and physicists 
to understand why plants
have the shape they have.

In the office of Stephane Douady,
a "spiral plants" specialist I 
visited, the shelves are full of 
beautiful 
dried specimen that people 
sent to him
from all over the world. 
Each plant has spectacular
spiral arragenements of leaves
that run upward the trunks.

A common spiral plant is the pinecone.

Stephane taught me that when 
the trunk grows, 
the cells
at the stem divide themselves. New cells
 end up building
the trunk but regularly, 
one of them differentiates and
becomes a leave. 
Because this happens on a regular tempo, 
the leaves arrange themselves along 
the trunk
in spiral patterns. 
Stephane pointed out that if you draw
a curve around the trunk to isolate 
one leave cell from
newer leave cells, you end up drawing 
the "front curve"
of the plant growth.

<img src="./images/growthCurveOnPineCone.png" style="height: 300px; width: auto;">

Stephane then compared a new leave cell 
to the last one.
Because leave cells appear regularly, 
a new cell is 
always located with the same distance
 above and the
same angle sideway, 
a rotation near 168 degrees, 
related to the golden ratio. 
Therefore, the "front curve" has 
this interesting geometric feature : 
but on the new leave cell,
it overlaps itself
after an upward translation 
and a rotation. Let's demonstrate this.

Use a grey pen and draw the unfolded the front curve on a piece 
of paper.

<img src="./images/fistVoiceSVGimage.svg" style="height: 100px; width: auto;">

Use a green pen and draw the same curve at the right of the first.

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRight.svg" style="height: 100px; width: auto;">

Draw it again
on a tracing paper with a red pen and put the tracing 
paper above and
shift it a bit to the right and 
a bit to the top.

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightAndSecondVoice.svg" style="height: 125px; width: auto;">

We did it ! The overlapping 
is perfect 
except for one space that have
the shape of an eye (
the place for the leave).

### Generating a theme form the front curve 

Now take take the first drawing and add a canevas
of vertical measure lines and horizontal pitch lines.

For the pitch lines, the spacing between pitch line
follows the scale of your choice. If you
choose a C minor scale, the tone differences between
notes are 1, 1/2, 1, 1, 1/2, 1 and 3/2.

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightWithGrid.svg" style="height: 300px; width: auto;">

Now, along the front curve, when it intersect the 
measure lines, 
draw some notes of your choosing.


<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightWithGridWithNotes.svg" style="height: 300px; width: auto;">

With the second voice note, this gives

<img src="./images/SVGfirstVoiceAndFirstVoiceShiftedRightWithSecondVoiceWithGridWithNotes.svg" style="height: 300px; width: auto;">

<p>Download the generated music score in the musescore musicxml format:
    <a href="./scores/score.musicxml" download="score.musicxml">Download MusicXML Score</a>
 </p>`
	__Content__00000000_.ContentPath = `content`
	__Content__00000000_.OutputPath = `../../../docs`
	__Content__00000000_.LayoutPath = `../../../vendor/github.com/fullstack-lang/gong/lib/ssg/go/defaults/layouts`
	__Content__00000000_.StaticPath = `../../../static`
	__Content__00000000_.Target = models.FILE

	__ExportToMusicxml__00000000_.Name = `Singloton`

	__FrontCurve__00000008_.Name = `Non Rotated 0 `
	__FrontCurve__00000008_.Path = `M640.41 512.19 C641.07 507.41 641.46 503.22 641.82 498.67 C642.17 494.12 642.37 489.56 642.52 484.92 C642.68 480.28 642.69 475.63 642.77 470.83 C642.85 466.02 642.76 461.24 643.00 456.11 C643.25 450.97 643.44 445.84 644.22 440.01 C644.99 434.17 646.90 427.60 647.64 421.10 C648.39 414.60 649.35 407.37 648.67 401.02 C647.99 394.66 645.95 388.66 643.56 382.99 C641.18 377.32 637.89 372.02 634.37 367.00 C630.85 361.98 626.76 357.27 622.46 352.86 C618.15 348.45 613.44 344.34 608.54 340.54 C603.65 336.75 598.85 333.43 593.09 330.09 C587.32 326.76 580.54 323.24 573.96 320.54 C567.39 317.84 560.54 315.54 553.64 313.88 C546.74 312.23 539.61 311.03 532.55 310.61 C525.49 310.19 518.27 310.30 511.29 311.34 C504.32 312.38 497.27 314.21 490.71 316.84 C484.14 319.47 477.80 323.38 471.88 327.10 C465.95 330.83 460.52 335.31 455.16 339.18 C449.80 343.04 444.75 346.71 439.72 350.29 C434.69 353.86 429.80 357.17 424.99 360.62 C420.17 364.07 415.44 367.44 410.82 370.99 C406.21 374.54 401.67 378.13 397.28 381.92 C392.89 385.71 388.36 389.89 384.47 393.72 C380.57 397.54 377.35 401.03 373.93 404.86 C370.51 408.69 367.23 412.66 363.95 416.72 C360.67 420.78 357.53 424.95 354.26 429.21 C351.00 433.47 347.90 437.83 344.35 442.28 C340.80 446.73 337.31 451.21 332.97 455.89 C328.62 460.58 322.95 465.17 318.27 470.41 C313.60 475.65 308.33 481.34 304.92 487.35 C301.51 493.36 299.42 499.93 297.83 506.45 C296.25 512.96 295.63 519.74 295.39 526.44 C295.15 533.13 295.59 539.93 296.40 546.61 C297.21 553.30 298.56 560.00 300.27 566.54 C301.98 573.09 303.91 579.16 306.66 585.90 C309.41 592.63 312.92 600.20 316.77 606.94 C320.63 613.68 324.98 620.25 329.80 626.32 C334.62 632.40 339.93 638.23 345.68 643.38 C351.43 648.53 357.68 653.34 364.31 657.24 C370.94 661.13 378.15 664.42 385.48 666.75 C392.81 669.09 400.77 670.24 408.31 671.25 C415.84 672.26 423.48 672.34 430.68 672.83 C437.89 673.32 444.75 673.73 451.54 674.18 C458.33 674.62 464.86 675.16 471.42 675.50 C477.99 675.84 484.44 676.16 490.91 676.23 C497.38 676.30 503.82 676.25 510.26 675.92 C516.71 675.59 523.17 675.16 529.56 674.25 C535.96 673.34 542.40 672.20 548.62 670.47 C554.84 668.73 561.04 666.55 566.90 663.86 C572.76 661.18 578.52 658.02 583.78 654.34 C589.04 650.65 594.12 646.46 598.45 641.74 C602.79 637.01 606.70 631.65 609.78 625.97 C612.87 620.30 614.99 613.79 616.95 607.70 C618.90 601.61 620.01 595.25 621.48 589.45 C622.96 583.64 624.33 578.20 625.78 572.87 C627.23 567.53 628.76 562.51 630.18 557.43 C631.60 552.36 633.03 547.42 634.30 542.42 C635.58 537.41 636.79 532.44 637.80 527.40 C638.82 522.37 639.74 516.98 640.41 512.19 Z`

	__FrontCurve__00000009_.Name = `Non Rotated 1 `
	__FrontCurve__00000009_.Path = `M384.47 393.72 C380.57 397.54 377.35 401.03 373.93 404.86 C370.51 408.69 367.23 412.66 363.95 416.72 C360.67 420.78 357.53 424.95 354.26 429.21 C351.00 433.47 347.90 437.83 344.35 442.28 C340.80 446.73 337.31 451.21 332.97 455.89 C328.62 460.58 322.95 465.17 318.27 470.41 C313.60 475.65 308.33 481.34 304.92 487.35 C301.51 493.36 299.42 499.93 297.83 506.45 C296.25 512.96 295.63 519.74 295.39 526.44 C295.15 533.13 295.59 539.93 296.40 546.61 C297.21 553.30 298.56 560.00 300.27 566.54 C301.98 573.09 303.91 579.16 306.66 585.90 C309.41 592.63 312.92 600.20 316.77 606.94 C320.63 613.68 324.98 620.25 329.80 626.32 C334.62 632.40 339.93 638.23 345.68 643.38 C351.43 648.53 357.68 653.34 364.31 657.24 C370.94 661.13 378.15 664.42 385.48 666.75 C392.81 669.09 400.77 670.24 408.31 671.25 C415.84 672.26 423.48 672.34 430.68 672.83 C437.89 673.32 444.75 673.73 451.54 674.18 C458.33 674.62 464.86 675.16 471.42 675.50 C477.99 675.84 484.44 676.16 490.91 676.23 C497.38 676.30 503.82 676.25 510.26 675.92 C516.71 675.59 523.54 674.98 529.56 674.25 C535.59 673.53 540.79 672.63 546.39 671.55 C551.99 670.48 557.57 669.19 563.18 667.79 C568.79 666.40 574.36 664.80 580.07 663.19 C585.77 661.58 591.40 659.76 597.42 658.13 C603.43 656.49 609.40 654.77 616.16 653.39 C622.92 652.01 630.65 651.50 637.98 649.84 C645.30 648.18 653.36 646.51 660.11 643.43 C666.87 640.34 672.92 635.98 678.51 631.35 C684.10 626.71 689.07 621.25 693.66 615.62 C698.26 609.98 702.36 603.83 706.07 597.53 C709.78 591.23 713.05 584.58 715.92 577.82 C718.79 571.05 721.16 564.56 723.28 556.95 C725.39 549.35 727.43 540.52 728.62 532.18 C729.81 523.84 730.47 515.31 730.42 506.90 C730.38 498.49 729.76 489.96 728.34 481.70 C726.91 473.44 724.86 465.15 721.90 457.36 C718.94 449.56 715.12 441.90 710.58 434.93 C706.04 427.95 700.24 421.53 694.66 415.50 C689.08 409.47 682.81 404.13 677.11 398.75 C671.40 393.37 665.89 388.27 660.43 383.22 C654.96 378.17 649.72 373.23 644.32 368.44 C638.91 363.66 633.55 358.98 628.00 354.52 C622.44 350.06 616.82 345.74 611.00 341.67 C605.18 337.59 599.26 333.61 593.09 330.09 C586.91 326.57 580.54 323.24 573.96 320.54 C567.39 317.84 560.54 315.54 553.64 313.88 C546.74 312.23 539.61 311.03 532.55 310.61 C525.49 310.19 518.27 310.30 511.29 311.34 C504.32 312.38 497.27 314.21 490.71 316.84 C484.14 319.47 477.80 323.38 471.88 327.10 C465.95 330.83 460.52 335.31 455.16 339.18 C449.80 343.04 444.75 346.71 439.72 350.29 C434.69 353.86 429.80 357.17 424.99 360.62 C420.17 364.07 415.44 367.44 410.82 370.99 C406.21 374.54 401.67 378.13 397.28 381.92 C392.89 385.71 388.36 389.89 384.47 393.72 Z`

	__FrontCurve__00000010_.Name = `Non Rotated 2 `
	__FrontCurve__00000010_.Path = `M529.56 674.25 C535.59 673.53 540.79 672.63 546.39 671.55 C551.99 670.48 557.57 669.19 563.18 667.79 C568.79 666.40 574.36 664.80 580.07 663.19 C585.77 661.58 591.40 659.76 597.42 658.13 C603.43 656.49 609.40 654.77 616.16 653.39 C622.92 652.01 630.65 651.50 637.98 649.84 C645.30 648.18 653.36 646.51 660.11 643.43 C666.87 640.34 672.92 635.98 678.51 631.35 C684.10 626.71 689.07 621.25 693.66 615.62 C698.26 609.98 702.36 603.83 706.07 597.53 C709.78 591.23 713.05 584.58 715.92 577.82 C718.79 571.05 721.16 564.56 723.28 556.95 C725.39 549.35 727.43 540.52 728.62 532.18 C729.81 523.84 730.47 515.31 730.42 506.90 C730.38 498.49 729.76 489.96 728.34 481.70 C726.91 473.44 724.86 465.15 721.90 457.36 C718.94 449.56 715.12 441.90 710.58 434.93 C706.04 427.95 700.24 421.53 694.66 415.50 C689.08 409.47 682.81 404.13 677.11 398.75 C671.40 393.37 665.89 388.27 660.43 383.22 C654.96 378.17 649.72 373.23 644.32 368.44 C638.91 363.66 633.55 358.98 628.00 354.52 C622.44 350.06 616.82 345.74 611.00 341.67 C605.18 337.59 598.82 333.51 593.09 330.09 C587.36 326.67 582.22 323.95 576.61 321.14 C571.01 318.32 565.28 315.72 559.45 313.20 C553.62 310.68 547.69 308.38 541.64 306.03 C535.60 303.68 529.47 301.57 523.16 299.10 C516.84 296.63 510.50 294.29 503.75 291.22 C497.00 288.15 490.02 283.83 482.65 280.69 C475.28 277.55 467.34 273.99 459.55 272.38 C451.77 270.77 443.76 270.63 435.95 271.02 C428.14 271.40 420.32 272.85 412.70 274.69 C405.08 276.54 397.54 279.11 390.23 282.07 C382.92 285.03 375.75 288.57 368.84 292.46 C361.94 296.36 355.63 300.34 348.80 305.43 C341.97 310.53 334.41 316.71 327.87 323.03 C321.33 329.35 315.12 336.17 309.57 343.36 C304.01 350.54 298.86 358.19 294.55 366.14 C290.23 374.09 286.43 382.48 283.67 391.04 C280.92 399.60 278.95 408.60 278.02 417.51 C277.08 426.42 277.51 435.68 278.04 444.51 C278.58 453.33 280.11 462.05 281.22 470.45 C282.32 478.85 283.48 486.92 284.67 494.92 C285.86 502.92 286.98 510.70 288.36 518.45 C289.74 526.20 291.18 533.84 292.95 541.42 C294.71 548.99 296.66 556.50 298.95 563.91 C301.23 571.33 303.69 578.72 306.66 585.90 C309.63 593.07 312.92 600.20 316.77 606.94 C320.63 613.68 324.98 620.25 329.80 626.32 C334.62 632.40 339.93 638.23 345.68 643.38 C351.43 648.53 357.68 653.34 364.31 657.24 C370.94 661.13 378.15 664.42 385.48 666.75 C392.81 669.09 400.77 670.24 408.31 671.25 C415.84 672.26 423.48 672.34 430.68 672.83 C437.89 673.32 444.75 673.73 451.54 674.18 C458.33 674.62 464.86 675.16 471.42 675.50 C477.99 675.84 484.44 676.16 490.91 676.23 C497.38 676.30 503.82 676.25 510.26 675.92 C516.71 675.59 523.54 674.98 529.56 674.25 Z`

	__FrontCurve__00000011_.Name = `Non Rotated 3 `
	__FrontCurve__00000011_.Path = `M593.09 330.09 C587.36 326.67 582.22 323.95 576.61 321.14 C571.01 318.32 565.28 315.72 559.45 313.20 C553.62 310.68 547.69 308.38 541.64 306.03 C535.60 303.68 529.47 301.57 523.16 299.10 C516.84 296.63 510.50 294.29 503.75 291.22 C497.00 288.15 490.02 283.83 482.65 280.69 C475.28 277.55 467.34 273.99 459.55 272.38 C451.77 270.77 443.76 270.63 435.95 271.02 C428.14 271.40 420.32 272.85 412.70 274.69 C405.08 276.54 397.54 279.11 390.23 282.07 C382.92 285.03 375.75 288.57 368.84 292.46 C361.94 296.36 355.63 300.34 348.80 305.43 C341.97 310.53 334.41 316.71 327.87 323.03 C321.33 329.35 315.12 336.17 309.57 343.36 C304.01 350.54 298.86 358.19 294.55 366.14 C290.23 374.09 286.43 382.48 283.67 391.04 C280.92 399.60 278.95 408.60 278.02 417.51 C277.08 426.42 277.51 435.68 278.04 444.51 C278.58 453.33 280.11 462.05 281.22 470.45 C282.32 478.85 283.48 486.92 284.67 494.92 C285.86 502.92 286.98 510.70 288.36 518.45 C289.74 526.20 291.18 533.84 292.95 541.42 C294.71 548.99 296.66 556.50 298.95 563.91 C301.23 571.33 304.00 579.11 306.66 585.90 C309.32 592.68 311.95 598.46 314.93 604.62 C317.90 610.79 321.13 616.85 324.51 622.89 C327.89 628.93 331.50 634.85 335.20 640.86 C338.90 646.86 342.83 652.71 346.71 658.94 C350.59 665.17 354.58 671.29 358.47 678.22 C362.36 685.15 365.68 693.19 370.05 700.51 C374.41 707.82 379.00 715.80 384.65 722.12 C390.31 728.44 397.07 733.72 404.00 738.42 C410.92 743.13 418.52 746.98 426.21 750.38 C433.90 753.77 441.99 756.51 450.13 758.77 C458.27 761.03 466.67 762.73 475.07 763.95 C483.47 765.18 491.42 765.94 500.55 766.11 C509.67 766.28 520.12 766.01 529.80 764.96 C539.49 763.91 549.22 762.24 558.64 759.83 C568.07 757.43 577.45 754.36 586.35 750.52 C595.25 746.69 604.00 742.15 612.02 736.80 C620.04 731.45 627.69 725.23 634.46 718.42 C641.23 711.61 647.12 703.66 652.63 695.95 C658.15 688.25 662.79 679.99 667.55 672.22 C672.31 664.44 676.80 656.87 681.19 649.32 C685.58 641.76 689.88 634.39 693.91 626.87 C697.94 619.35 701.82 611.86 705.38 604.20 C708.94 596.54 712.28 588.81 715.26 580.94 C718.25 573.06 721.05 565.08 723.28 556.95 C725.50 548.82 727.43 540.52 728.62 532.18 C729.81 523.84 730.47 515.31 730.42 506.90 C730.38 498.49 729.76 489.96 728.34 481.70 C726.91 473.44 724.86 465.15 721.90 457.36 C718.94 449.56 715.12 441.90 710.58 434.93 C706.04 427.95 700.24 421.53 694.66 415.50 C689.08 409.47 682.81 404.13 677.11 398.75 C671.40 393.37 665.89 388.27 660.43 383.22 C654.96 378.17 649.72 373.23 644.32 368.44 C638.91 363.66 633.55 358.98 628.00 354.52 C622.44 350.06 616.82 345.74 611.00 341.67 C605.18 337.59 598.82 333.51 593.09 330.09 Z`

	__FrontCurve__00000012_.Name = `Non Rotated 4 `
	__FrontCurve__00000012_.Path = `M306.66 585.90 C309.32 592.68 311.95 598.46 314.93 604.62 C317.90 610.79 321.13 616.85 324.51 622.89 C327.89 628.93 331.50 634.85 335.20 640.86 C338.90 646.86 342.83 652.71 346.71 658.94 C350.59 665.17 354.58 671.29 358.47 678.22 C362.36 685.15 365.68 693.19 370.05 700.51 C374.41 707.82 379.00 715.80 384.65 722.12 C390.31 728.44 397.07 733.72 404.00 738.42 C410.92 743.13 418.52 746.98 426.21 750.38 C433.90 753.77 441.99 756.51 450.13 758.77 C458.27 761.03 466.67 762.73 475.07 763.95 C483.47 765.18 491.42 765.94 500.55 766.11 C509.67 766.28 520.12 766.01 529.80 764.96 C539.49 763.91 549.22 762.24 558.64 759.83 C568.07 757.43 577.45 754.36 586.35 750.52 C595.25 746.69 604.00 742.15 612.02 736.80 C620.04 731.45 627.69 725.23 634.46 718.42 C641.23 711.61 647.12 703.66 652.63 695.95 C658.15 688.25 662.79 679.99 667.55 672.22 C672.31 664.44 676.80 656.87 681.19 649.32 C685.58 641.76 689.88 634.39 693.91 626.87 C697.94 619.35 701.82 611.86 705.38 604.20 C708.94 596.54 712.28 588.81 715.26 580.94 C718.25 573.06 721.05 564.54 723.28 556.95 C725.50 549.36 727.08 542.65 728.61 535.39 C730.14 528.12 731.38 520.77 732.46 513.35 C733.54 505.93 734.33 498.46 735.08 490.87 C735.83 483.27 736.28 475.66 736.97 467.79 C737.66 459.92 738.17 452.06 739.22 443.64 C740.28 435.21 742.44 426.26 743.31 417.26 C744.18 408.25 745.26 398.60 744.44 389.60 C743.63 380.61 741.28 371.78 738.43 363.30 C735.57 354.83 731.64 346.63 727.33 338.76 C723.02 330.89 717.97 323.30 712.57 316.08 C707.16 308.85 701.19 301.94 694.90 295.41 C688.61 288.89 682.41 283.04 674.84 276.93 C667.27 270.82 658.29 264.21 649.46 258.74 C640.63 253.27 631.35 248.30 621.85 244.13 C612.35 239.96 602.46 236.36 592.47 233.72 C582.48 231.07 572.16 229.10 561.91 228.25 C551.66 227.41 541.16 227.52 530.99 228.67 C520.83 229.82 510.64 232.45 500.91 235.15 C491.17 237.85 481.78 241.57 472.60 244.86 C463.42 248.16 454.58 251.50 445.83 254.92 C437.09 258.34 428.54 261.71 420.12 265.39 C411.70 269.07 403.42 272.86 395.32 277.01 C387.21 281.15 379.24 285.53 371.49 290.27 C363.74 295.01 356.07 299.97 348.80 305.43 C341.53 310.89 334.41 316.71 327.87 323.03 C321.33 329.35 315.12 336.17 309.57 343.36 C304.01 350.54 298.86 358.19 294.55 366.14 C290.23 374.09 286.43 382.48 283.67 391.04 C280.92 399.60 278.95 408.60 278.02 417.51 C277.08 426.42 277.51 435.68 278.04 444.51 C278.58 453.33 280.11 462.05 281.22 470.45 C282.32 478.85 283.48 486.92 284.67 494.92 C285.86 502.92 286.98 510.70 288.36 518.45 C289.74 526.20 291.18 533.84 292.95 541.42 C294.71 548.99 296.66 556.50 298.95 563.91 C301.23 571.33 304.00 579.11 306.66 585.90 Z`

	__FrontCurve__00000013_.Name = `Non Rotated 5 `
	__FrontCurve__00000013_.Path = `M723.28 556.95 C725.50 549.36 727.08 542.65 728.61 535.39 C730.14 528.12 731.38 520.77 732.46 513.35 C733.54 505.93 734.33 498.46 735.08 490.87 C735.83 483.27 736.28 475.66 736.97 467.79 C737.66 459.92 738.17 452.06 739.22 443.64 C740.28 435.21 742.44 426.26 743.31 417.26 C744.18 408.25 745.26 398.60 744.44 389.60 C743.63 380.61 741.28 371.78 738.43 363.30 C735.57 354.83 731.64 346.63 727.33 338.76 C723.02 330.89 717.97 323.30 712.57 316.08 C707.16 308.85 701.19 301.94 694.90 295.41 C688.61 288.89 682.41 283.04 674.84 276.93 C667.27 270.82 658.29 264.21 649.46 258.74 C640.63 253.27 631.35 248.30 621.85 244.13 C612.35 239.96 602.46 236.36 592.47 233.72 C582.48 231.07 572.16 229.10 561.91 228.25 C551.66 227.41 541.16 227.52 530.99 228.67 C520.83 229.82 510.64 232.45 500.91 235.15 C491.17 237.85 481.78 241.57 472.60 244.86 C463.42 248.16 454.58 251.50 445.83 254.92 C437.09 258.34 428.54 261.71 420.12 265.39 C411.70 269.07 403.42 272.86 395.32 277.01 C387.21 281.15 379.24 285.53 371.49 290.27 C363.74 295.01 355.71 300.42 348.80 305.43 C341.89 310.44 336.12 315.13 330.03 320.33 C323.94 325.53 318.04 330.99 312.25 336.63 C306.45 342.26 300.88 348.14 295.28 354.14 C289.68 360.15 284.34 366.37 278.66 372.67 C272.99 378.96 267.49 385.38 261.24 391.90 C254.99 398.43 247.56 404.71 241.15 411.83 C234.74 418.95 227.79 426.47 222.77 434.63 C217.75 442.79 214.07 451.77 211.03 460.78 C207.99 469.79 206.01 479.25 204.55 488.69 C203.09 498.12 202.41 507.78 202.27 517.38 C202.13 526.97 202.65 536.67 203.69 546.25 C204.74 555.82 206.14 564.77 208.52 574.84 C210.90 584.91 214.12 596.31 217.97 606.66 C221.81 617.01 226.35 627.24 231.58 636.95 C236.82 646.66 242.74 656.14 249.36 664.91 C255.98 673.69 263.31 682.12 271.30 689.58 C279.28 697.05 288.09 703.90 297.29 709.70 C306.49 715.49 316.64 720.11 326.50 724.36 C336.35 728.61 346.59 731.83 356.41 735.21 C366.23 738.59 375.83 741.70 385.42 744.66 C395.01 747.61 404.43 750.46 413.95 752.94 C423.47 755.42 432.96 757.69 442.53 759.54 C452.11 761.39 461.73 762.94 471.40 764.04 C481.06 765.13 490.81 765.95 500.55 766.11 C510.28 766.26 520.12 766.01 529.80 764.96 C539.49 763.91 549.22 762.24 558.64 759.83 C568.07 757.43 577.45 754.36 586.35 750.52 C595.25 746.69 604.00 742.15 612.02 736.80 C620.04 731.45 627.69 725.23 634.46 718.42 C641.23 711.61 647.12 703.66 652.63 695.95 C658.15 688.25 662.79 679.99 667.55 672.22 C672.31 664.44 676.80 656.87 681.19 649.32 C685.58 641.76 689.88 634.39 693.91 626.87 C697.94 619.35 701.82 611.86 705.38 604.20 C708.94 596.54 712.28 588.81 715.26 580.94 C718.25 573.06 721.05 564.54 723.28 556.95 Z`

	__FrontCurve__00000014_.Name = `Non Rotated 6 `
	__FrontCurve__00000014_.Path = `M348.80 305.43 C341.89 310.44 336.12 315.13 330.03 320.33 C323.94 325.53 318.04 330.99 312.25 336.63 C306.45 342.26 300.88 348.14 295.28 354.14 C289.68 360.15 284.34 366.37 278.66 372.67 C272.99 378.96 267.49 385.38 261.24 391.90 C254.99 398.43 247.56 404.71 241.15 411.83 C234.74 418.95 227.79 426.47 222.77 434.63 C217.75 442.79 214.07 451.77 211.03 460.78 C207.99 469.79 206.01 479.25 204.55 488.69 C203.09 498.12 202.41 507.78 202.27 517.38 C202.13 526.97 202.65 536.67 203.69 546.25 C204.74 555.82 206.14 564.77 208.52 574.84 C210.90 584.91 214.12 596.31 217.97 606.66 C221.81 617.01 226.35 627.24 231.58 636.95 C236.82 646.66 242.74 656.14 249.36 664.91 C255.98 673.69 263.31 682.12 271.30 689.58 C279.28 697.05 288.09 703.90 297.29 709.70 C306.49 715.49 316.64 720.11 326.50 724.36 C336.35 728.61 346.59 731.83 356.41 735.21 C366.23 738.59 375.83 741.70 385.42 744.66 C395.01 747.61 404.43 750.46 413.95 752.94 C423.47 755.42 432.96 757.69 442.53 759.54 C452.11 761.39 461.73 762.94 471.40 764.04 C481.06 765.13 491.41 765.81 500.55 766.11 C509.68 766.40 517.64 766.21 526.20 765.80 C534.76 765.38 543.33 764.61 551.92 763.60 C560.50 762.60 569.06 761.25 577.71 759.78 C586.35 758.31 594.93 756.48 603.79 754.78 C612.65 753.08 621.45 751.14 630.88 749.56 C640.32 747.98 650.52 747.28 660.42 745.28 C670.32 743.28 680.88 741.26 690.29 737.54 C699.71 733.81 708.56 728.59 716.89 722.94 C725.23 717.28 733.00 710.60 740.31 703.61 C747.63 696.61 754.45 688.93 760.78 680.97 C767.11 673.01 772.96 664.53 778.31 655.82 C783.65 647.12 788.28 638.71 792.84 628.73 C797.39 618.75 802.11 607.10 805.64 595.94 C809.16 584.77 812.03 573.27 814.00 561.74 C815.97 550.21 817.21 538.43 817.44 526.78 C817.67 515.13 817.10 503.30 815.37 491.85 C813.65 480.39 810.87 468.91 807.09 458.03 C803.31 447.15 798.02 436.62 792.70 426.57 C787.39 416.53 781.10 407.07 775.21 397.75 C769.32 388.43 763.40 379.47 757.36 370.65 C751.32 361.84 745.32 353.24 738.98 344.88 C732.64 336.52 726.17 328.36 719.35 320.51 C712.52 312.66 705.44 305.04 698.03 297.78 C690.61 290.51 682.94 283.44 674.84 276.93 C666.75 270.43 658.29 264.21 649.46 258.74 C640.63 253.27 631.35 248.30 621.85 244.13 C612.35 239.96 602.46 236.36 592.47 233.72 C582.48 231.07 572.16 229.10 561.91 228.25 C551.66 227.41 541.16 227.52 530.99 228.67 C520.83 229.82 510.64 232.45 500.91 235.15 C491.17 237.85 481.78 241.57 472.60 244.86 C463.42 248.16 454.58 251.50 445.83 254.92 C437.09 258.34 428.54 261.71 420.12 265.39 C411.70 269.07 403.42 272.86 395.32 277.01 C387.21 281.15 379.24 285.53 371.49 290.27 C363.74 295.01 355.71 300.42 348.80 305.43 Z`

	__FrontCurve__00000015_.Name = `Non Rotated 7 `
	__FrontCurve__00000015_.Path = `M500.55 766.11 C509.68 766.40 517.64 766.21 526.20 765.80 C534.76 765.38 543.33 764.61 551.92 763.60 C560.50 762.60 569.06 761.25 577.71 759.78 C586.35 758.31 594.93 756.48 603.79 754.78 C612.65 753.08 621.45 751.14 630.88 749.56 C640.32 747.98 650.52 747.28 660.42 745.28 C670.32 743.28 680.88 741.26 690.29 737.54 C699.71 733.81 708.56 728.59 716.89 722.94 C725.23 717.28 733.00 710.60 740.31 703.61 C747.63 696.61 754.45 688.93 760.78 680.97 C767.11 673.01 772.96 664.53 778.31 655.82 C783.65 647.12 788.28 638.71 792.84 628.73 C797.39 618.75 802.11 607.10 805.64 595.94 C809.16 584.77 812.03 573.27 814.00 561.74 C815.97 550.21 817.21 538.43 817.44 526.78 C817.67 515.13 817.10 503.30 815.37 491.85 C813.65 480.39 810.87 468.91 807.09 458.03 C803.31 447.15 798.02 436.62 792.70 426.57 C787.39 416.53 781.10 407.07 775.21 397.75 C769.32 388.43 763.40 379.47 757.36 370.65 C751.32 361.84 745.32 353.24 738.98 344.88 C732.64 336.52 726.17 328.36 719.35 320.51 C712.52 312.66 705.44 305.04 698.03 297.78 C690.61 290.51 682.34 283.16 674.84 276.93 C667.34 270.71 660.51 265.66 653.02 260.42 C645.53 255.18 637.80 250.24 629.91 245.50 C622.01 240.76 613.91 236.33 605.66 231.97 C597.41 227.61 588.98 223.60 580.40 219.34 C571.81 215.08 563.13 211.08 554.16 206.43 C545.18 201.77 536.14 195.98 526.54 191.43 C516.93 186.88 506.79 181.96 496.51 179.12 C486.24 176.29 475.48 175.01 464.86 174.42 C454.25 173.83 443.45 174.43 432.83 175.59 C422.20 176.74 411.55 178.77 401.12 181.35 C390.68 183.93 380.32 187.23 370.23 191.06 C360.13 194.89 350.83 198.92 340.56 204.31 C330.29 209.71 318.81 216.38 308.63 223.42 C298.44 230.46 288.58 238.21 279.44 246.55 C270.30 254.89 261.59 263.90 253.79 273.46 C246.00 283.01 238.76 293.22 232.68 303.87 C226.60 314.51 221.37 325.84 217.33 337.33 C213.28 348.81 210.69 360.97 208.43 372.77 C206.18 384.56 205.10 396.51 203.82 408.12 C202.54 419.72 201.54 431.07 200.76 442.38 C199.99 453.69 199.35 464.82 199.16 475.96 C198.97 487.10 199.04 498.18 199.62 509.22 C200.21 520.26 201.16 531.28 202.64 542.22 C204.13 553.16 205.97 564.10 208.52 574.84 C211.07 585.58 214.12 596.31 217.97 606.66 C221.81 617.01 226.35 627.24 231.58 636.95 C236.82 646.66 242.74 656.14 249.36 664.91 C255.98 673.69 263.31 682.12 271.30 689.58 C279.28 697.05 288.09 703.90 297.29 709.70 C306.49 715.49 316.64 720.11 326.50 724.36 C336.35 728.61 346.59 731.83 356.41 735.21 C366.23 738.59 375.83 741.70 385.42 744.66 C395.01 747.61 404.43 750.46 413.95 752.94 C423.47 755.42 432.96 757.69 442.53 759.54 C452.11 761.39 461.73 762.94 471.40 764.04 C481.06 765.13 491.41 765.81 500.55 766.11 Z`

	__FrontCurveStack__00000000_.Name = `Front Curve Stack`
	__FrontCurveStack__00000000_.IsDisplayed = false
	__FrontCurveStack__00000000_.Color = ``
	__FrontCurveStack__00000000_.FillOpacity = 0.000000
	__FrontCurveStack__00000000_.Stroke = `blue`
	__FrontCurveStack__00000000_.StrokeOpacity = 1.000000
	__FrontCurveStack__00000000_.StrokeWidth = 1.000000
	__FrontCurveStack__00000000_.StrokeDashArray = ``
	__FrontCurveStack__00000000_.StrokeDashArrayWhenSelected = ``
	__FrontCurveStack__00000000_.Transform = ``

	__HorizontalAxis__00000000_.Name = `Horizontal Axis`
	__HorizontalAxis__00000000_.IsDisplayed = true
	__HorizontalAxis__00000000_.AxisHandleBorderLength = 0.000000
	__HorizontalAxis__00000000_.Axis_Length = 600.000000
	__HorizontalAxis__00000000_.Color = ``
	__HorizontalAxis__00000000_.FillOpacity = 0.000000
	__HorizontalAxis__00000000_.Stroke = ``
	__HorizontalAxis__00000000_.StrokeOpacity = 0.000000
	__HorizontalAxis__00000000_.StrokeWidth = 1.000000
	__HorizontalAxis__00000000_.StrokeDashArray = ``
	__HorizontalAxis__00000000_.StrokeDashArrayWhenSelected = ``
	__HorizontalAxis__00000000_.Transform = ``

	__Key__00000000_.Name = `F key`
	__Key__00000000_.IsDisplayed = false
	__Key__00000000_.Path = `M562 -21c0 89 -65 150 -155 150c7 -44 34 -203 55 -323c71 29 100 102 100 173zM420 -206l-58 329c-59 -14 -104 -63 -104 -124c0 -49 22 -75 61 -99c12 -8 22 -13 22 -22s-9 -13 -17 -13c-80 0 -135 96 -135 166c0 94 62 190 153 217c-7 41 -14 88 -23 142 c-15 -15 -31 -29 -48 -44c-88 -76 -174 -185 -174 -307c0 -151 122 -251 265 -251c19 0 38 2 58 6zM332 822c-8 -31 -11 -65 -11 -102c0 -42 5 -81 11 -121c69 68 146 146 146 250c0 69 -24 118 -39 118c-52 0 -98 -105 -107 -145zM122 -513c0 66 45 123 115 123 c75 0 116 -57 116 -111c0 -64 -47 -104 -94 -111c-3 -1 -5 -2 -5 -4c0 -1 2 -2 3 -3c2 0 23 -5 47 -5c101 0 154 55 154 159c0 53 -11 123 -30 219c-23 -4 -50 -7 -79 -7c-186 0 -349 147 -349 334c0 200 126 321 217 406c21 17 73 70 74 71c-17 112 -22 161 -22 215 c0 84 18 212 82 288c33 39 64 51 71 51c18 0 47 -35 71 -86c16 -36 44 -110 44 -201c0 -159 -73 -284 -179 -395c9 -56 19 -115 29 -175c146 0 253 -102 253 -253c0 -103 -73 -205 -171 -237c6 -39 12 -69 15 -89c10 -57 16 -102 16 -141c0 -63 -14 -129 -68 -167 c-36 -22 -77 -34 -124 -34c-135 0 -186 87 -186 153z`
	__Key__00000000_.Color = `black`
	__Key__00000000_.FillOpacity = 0.300000
	__Key__00000000_.Stroke = `black`
	__Key__00000000_.StrokeOpacity = 0.300000
	__Key__00000000_.StrokeWidth = 0.000000
	__Key__00000000_.StrokeDashArray = ``
	__Key__00000000_.StrokeDashArrayWhenSelected = ``
	__Key__00000000_.Transform = `scale(0.2,-0.2)`

	__Parameter__00000000_.Name = `Reference`
	__Parameter__00000000_.BackendColor = `#F1F1F1`
	__Parameter__00000000_.MinuteColor = `#5A5A5A`
	__Parameter__00000000_.HourColor = `#1E1E1E`
	__Parameter__00000000_.N = 3
	__Parameter__00000000_.M = 2
	__Parameter__00000000_.Z = 16
	__Parameter__00000000_.InsideAngle = 120.000000
	__Parameter__00000000_.ShiftToNearestCircle = 3
	__Parameter__00000000_.SideLength = 90.000000
	__Parameter__00000000_.StackWidth = 1
	__Parameter__00000000_.NbShitRight = 2
	__Parameter__00000000_.StackHeight = 8
	__Parameter__00000000_.BezierControlLengthRatio = 0.560000
	__Parameter__00000000_.SpiralBezierStrength = 0.800000
	__Parameter__00000000_.NbInterpolationPoints = 12
	__Parameter__00000000_.FkeySizeRatio = 0.001430
	__Parameter__00000000_.FkeyOriginRelativeX = 1.500000
	__Parameter__00000000_.FkeyOriginRelativeY = -3.400000
	__Parameter__00000000_.PitchHeight = 0.139000
	__Parameter__00000000_.NbPitchLines = 61
	__Parameter__00000000_.BeatLinesHeightRatio = 0.170000
	__Parameter__00000000_.NbBeatLines = 49
	__Parameter__00000000_.NbOfBeatsInTheme = 16
	__Parameter__00000000_.FirstVoiceShiftX = 0.100000
	__Parameter__00000000_.FirstVoiceShiftY = 2.480000
	__Parameter__00000000_.PitchDifference = 12
	__Parameter__00000000_.BeatsPerSecond = 6.850000
	__Parameter__00000000_.Level = 11.100000
	__Parameter__00000000_.IsMinor = true
	__Parameter__00000000_.ThemeBinaryEncoding = 71
	__Parameter__00000000_.OriginX = 100.000000
	__Parameter__00000000_.OriginY = 950.000000
	__Parameter__00000000_.SpiralOriginX = 500.000000
	__Parameter__00000000_.SpiralOriginY = 500.000000
	__Parameter__00000000_.OriginCrossWidth = 800.000000
	__Parameter__00000000_.SpiralRadiusRatio = 1.060000
	__Parameter__00000000_.ShowSpiralBezierConstruct = false
	__Parameter__00000000_.ShowInterpolationPoints = false
	__Parameter__00000000_.ActualBeatsTemporalShift = 6
	__Parameter__00000000_.PathToStaticFiles = `../../../static`
	__Parameter__00000000_.PathToGeneratedSVG = `../../../static/images`
	__Parameter__00000000_.PathToGeneratedScore = `../../../static/scores`

	__Rhombus__00000000_.Name = `Growing Rhombus Grid Seed`
	__Rhombus__00000000_.IsDisplayed = false
	__Rhombus__00000000_.CenterX = 0.000000
	__Rhombus__00000000_.CenterY = 0.000000
	__Rhombus__00000000_.SideLength = 90.000000
	__Rhombus__00000000_.AngleDegree = -83.413224
	__Rhombus__00000000_.InsideAngle = 120.000000
	__Rhombus__00000000_.Color = ``
	__Rhombus__00000000_.FillOpacity = 0.000000
	__Rhombus__00000000_.Stroke = `green`
	__Rhombus__00000000_.StrokeOpacity = 0.800000
	__Rhombus__00000000_.StrokeWidth = 2.000000
	__Rhombus__00000000_.StrokeDashArray = ``
	__Rhombus__00000000_.StrokeDashArrayWhenSelected = ``
	__Rhombus__00000000_.Transform = ``

	__Rhombus__00000001_.Name = `Initial Rhombus`
	__Rhombus__00000001_.IsDisplayed = false
	__Rhombus__00000001_.CenterX = 0.000000
	__Rhombus__00000001_.CenterY = 0.000000
	__Rhombus__00000001_.SideLength = 90.000000
	__Rhombus__00000001_.AngleDegree = -79.106605
	__Rhombus__00000001_.InsideAngle = 120.000000
	__Rhombus__00000001_.Color = ``
	__Rhombus__00000001_.FillOpacity = 0.000000
	__Rhombus__00000001_.Stroke = `green`
	__Rhombus__00000001_.StrokeOpacity = 0.900000
	__Rhombus__00000001_.StrokeWidth = 1.000000
	__Rhombus__00000001_.StrokeDashArray = ``
	__Rhombus__00000001_.StrokeDashArrayWhenSelected = ``
	__Rhombus__00000001_.Transform = ``

	__Rhombus__00000002_.Name = `Rotated Next Rhombus`
	__Rhombus__00000002_.IsDisplayed = false
	__Rhombus__00000002_.CenterX = 154.855620
	__Rhombus__00000002_.CenterY = 17.881187
	__Rhombus__00000002_.SideLength = 90.000000
	__Rhombus__00000002_.AngleDegree = -83.413224
	__Rhombus__00000002_.InsideAngle = 120.000000
	__Rhombus__00000002_.Color = ``
	__Rhombus__00000002_.FillOpacity = 0.000000
	__Rhombus__00000002_.Stroke = `lavender`
	__Rhombus__00000002_.StrokeOpacity = 0.700000
	__Rhombus__00000002_.StrokeWidth = 6.000000
	__Rhombus__00000002_.StrokeDashArray = ``
	__Rhombus__00000002_.StrokeDashArrayWhenSelected = ``
	__Rhombus__00000002_.Transform = ``

	__Rhombus__00000003_.Name = `Rotated Rhombus`
	__Rhombus__00000003_.IsDisplayed = false
	__Rhombus__00000003_.CenterX = 0.000000
	__Rhombus__00000003_.CenterY = 0.000000
	__Rhombus__00000003_.SideLength = 90.000000
	__Rhombus__00000003_.AngleDegree = -83.413224
	__Rhombus__00000003_.InsideAngle = 120.000000
	__Rhombus__00000003_.Color = ``
	__Rhombus__00000003_.FillOpacity = 0.000000
	__Rhombus__00000003_.Stroke = `black`
	__Rhombus__00000003_.StrokeOpacity = 1.000000
	__Rhombus__00000003_.StrokeWidth = 2.000000
	__Rhombus__00000003_.StrokeDashArray = ``
	__Rhombus__00000003_.StrokeDashArrayWhenSelected = ``
	__Rhombus__00000003_.Transform = ``

	__RhombusGrid__00000000_.Name = `Growing Rhombus Grid`
	__RhombusGrid__00000000_.IsDisplayed = false

	__RhombusGrid__00000001_.Name = `Initial Rhombus Grid`
	__RhombusGrid__00000001_.IsDisplayed = false

	__RhombusGrid__00000002_.Name = `Rotated Rhombus Grid`
	__RhombusGrid__00000002_.IsDisplayed = false

	__ShapeCategory__00000000_.Name = `0. Axes`
	__ShapeCategory__00000000_.IsExpanded = false

	__ShapeCategory__00000001_.Name = `1. Initial`
	__ShapeCategory__00000001_.IsExpanded = false

	__ShapeCategory__00000002_.Name = `2. Rotated`
	__ShapeCategory__00000002_.IsExpanded = false

	__ShapeCategory__00000003_.Name = `3. Growing`
	__ShapeCategory__00000003_.IsExpanded = false

	__ShapeCategory__00000004_.Name = `4. Construction`
	__ShapeCategory__00000004_.IsExpanded = false

	__ShapeCategory__00000005_.Name = `5. Vertical growth`
	__ShapeCategory__00000005_.IsExpanded = false

	__ShapeCategory__00000006_.Name = `6. Spiral growth`
	__ShapeCategory__00000006_.IsExpanded = false

	__ShapeCategory__00000007_.Name = `7. Spiral Growth Bezier`
	__ShapeCategory__00000007_.IsExpanded = false

	__ShapeCategory__00000008_.Name = `8. Score notation`
	__ShapeCategory__00000008_.IsExpanded = true

	__ShapeCategory__00000009_.Name = `9. Composer`
	__ShapeCategory__00000009_.IsExpanded = true

	__SpiralBezier__00000000_.Name = `Spiral Bezier Seed`
	__SpiralBezier__00000000_.IsDisplayed = false
	__SpiralBezier__00000000_.StartX = 139.624444
	__SpiralBezier__00000000_.StartY = -11.569611
	__SpiralBezier__00000000_.ControlPointStartX = 154.703201
	__SpiralBezier__00000000_.ControlPointStartY = 22.559171
	__SpiralBezier__00000000_.EndX = 92.213155
	__SpiralBezier__00000000_.EndY = 170.394985
	__SpiralBezier__00000000_.ControlPointEndX = 115.594705
	__SpiralBezier__00000000_.ControlPointEndY = 139.829900
	__SpiralBezier__00000000_.Color = ``
	__SpiralBezier__00000000_.FillOpacity = 0.000000
	__SpiralBezier__00000000_.Stroke = `green`
	__SpiralBezier__00000000_.StrokeOpacity = 1.000000
	__SpiralBezier__00000000_.StrokeWidth = 1.000000
	__SpiralBezier__00000000_.StrokeDashArray = ``
	__SpiralBezier__00000000_.StrokeDashArrayWhenSelected = ``
	__SpiralBezier__00000000_.Transform = ``

	__SpiralBezierGrid__00000000_.Name = `Spiral Bezier Full Grid`
	__SpiralBezierGrid__00000000_.IsDisplayed = false

	__SpiralBezierGrid__00000001_.Name = `Spiral Bezier Grid`
	__SpiralBezierGrid__00000001_.IsDisplayed = false

	__SpiralCircle__00000000_.Name = `Construction Circle Spiral`
	__SpiralCircle__00000000_.IsDisplayed = false
	__SpiralCircle__00000000_.CenterX = 95.400000
	__SpiralCircle__00000000_.CenterY = 0.000000
	__SpiralCircle__00000000_.HasBespokeRadius = true
	__SpiralCircle__00000000_.BespopkeRadius = 14.130000
	__SpiralCircle__00000000_.Color = ``
	__SpiralCircle__00000000_.FillOpacity = 0.000000
	__SpiralCircle__00000000_.Stroke = `blue`
	__SpiralCircle__00000000_.StrokeOpacity = 0.800000
	__SpiralCircle__00000000_.StrokeWidth = 1.000000
	__SpiralCircle__00000000_.StrokeDashArray = ``
	__SpiralCircle__00000000_.StrokeDashArrayWhenSelected = ``
	__SpiralCircle__00000000_.Transform = ``
	__SpiralCircle__00000000_.Pitch = 0
	__SpiralCircle__00000000_.ShowName = false
	__SpiralCircle__00000000_.BeatNb = 0
	__SpiralCircle__00000000_.Path = `M 657.230000 500.000000 A 14.130000 14.130000 0 1 0 628.970000 500.000000 A 14.130000 14.130000 0 1 0 657.230000 500.000000 Z`

	__SpiralCircleGrid__00000000_.Name = `Brute Spiral Bezier Circle Grid`
	__SpiralCircleGrid__00000000_.IsDisplayed = true

	__SpiralCircleGrid__00000001_.Name = `Construction Circle Spiral Full Grid`
	__SpiralCircleGrid__00000001_.IsDisplayed = false

	__SpiralCircleGrid__00000002_.Name = `Construction Circle Spiral Grid`
	__SpiralCircleGrid__00000002_.IsDisplayed = false

	__SpiralCircleGrid__00000003_.Name = `Spiral Circle Grid`
	__SpiralCircleGrid__00000003_.IsDisplayed = false

	__SpiralLine__00000000_.Name = `Spiral Contruction Inner Line`
	__SpiralLine__00000000_.IsDisplayed = false
	__SpiralLine__00000000_.StartX = 139.624444
	__SpiralLine__00000000_.EndX = 95.400000
	__SpiralLine__00000000_.StartY = -11.569611
	__SpiralLine__00000000_.EndY = -0.000000
	__SpiralLine__00000000_.Color = ``
	__SpiralLine__00000000_.FillOpacity = 0.000000
	__SpiralLine__00000000_.Stroke = `blue`
	__SpiralLine__00000000_.StrokeOpacity = 1.000000
	__SpiralLine__00000000_.StrokeWidth = 1.000000
	__SpiralLine__00000000_.StrokeDashArray = ``
	__SpiralLine__00000000_.StrokeDashArrayWhenSelected = ``
	__SpiralLine__00000000_.Transform = ``

	__SpiralLine__00000001_.Name = `Spiral Contruction Outer Line`
	__SpiralLine__00000001_.IsDisplayed = false
	__SpiralLine__00000001_.StartX = 139.624444
	__SpiralLine__00000001_.EndX = 182.285422
	__SpiralLine__00000001_.StartY = -11.569611
	__SpiralLine__00000001_.EndY = -30.418057
	__SpiralLine__00000001_.Color = ``
	__SpiralLine__00000001_.FillOpacity = 0.000000
	__SpiralLine__00000001_.Stroke = `green`
	__SpiralLine__00000001_.StrokeOpacity = 1.000000
	__SpiralLine__00000001_.StrokeWidth = 1.000000
	__SpiralLine__00000001_.StrokeDashArray = ``
	__SpiralLine__00000001_.StrokeDashArrayWhenSelected = ``
	__SpiralLine__00000001_.Transform = ``

	__SpiralLineGrid__00000000_.Name = `Spiral Construction Inner Line Grid Spiral`
	__SpiralLineGrid__00000000_.IsDisplayed = false

	__SpiralLineGrid__00000001_.Name = `Spiral Construction Outer Line Full Grid Spiral`
	__SpiralLineGrid__00000001_.IsDisplayed = false

	__SpiralLineGrid__00000002_.Name = `Spiral Construction Outer Line Grid Spiral`
	__SpiralLineGrid__00000002_.IsDisplayed = false

	__SpiralOrigin__00000000_.Name = `Spiral Origin`
	__SpiralOrigin__00000000_.IsDisplayed = false
	__SpiralOrigin__00000000_.Color = ``
	__SpiralOrigin__00000000_.FillOpacity = 0.000000
	__SpiralOrigin__00000000_.Stroke = `black`
	__SpiralOrigin__00000000_.StrokeOpacity = 1.000000
	__SpiralOrigin__00000000_.StrokeWidth = 1.000000
	__SpiralOrigin__00000000_.StrokeDashArray = ``
	__SpiralOrigin__00000000_.StrokeDashArrayWhenSelected = ``
	__SpiralOrigin__00000000_.Transform = ``

	__SpiralRhombus__00000000_.Name = `Reference Spiral Rhombus`
	__SpiralRhombus__00000000_.IsDisplayed = false
	__SpiralRhombus__00000000_.X_r0 = 50.523877
	__SpiralRhombus__00000000_.Y_r0 = 4.186528
	__SpiralRhombus__00000000_.X_r1 = 33.879335
	__SpiralRhombus__00000000_.Y_r1 = 98.687132
	__SpiralRhombus__00000000_.X_r2 = 139.624444
	__SpiralRhombus__00000000_.Y_r2 = -11.569611
	__SpiralRhombus__00000000_.X_r3 = 28.073323
	__SpiralRhombus__00000000_.Y_r3 = -81.774797
	__SpiralRhombus__00000000_.Color = `green`
	__SpiralRhombus__00000000_.FillOpacity = 1.000000
	__SpiralRhombus__00000000_.Stroke = `blue`
	__SpiralRhombus__00000000_.StrokeOpacity = 1.000000
	__SpiralRhombus__00000000_.StrokeWidth = 1.000000
	__SpiralRhombus__00000000_.StrokeDashArray = ``
	__SpiralRhombus__00000000_.StrokeDashArrayWhenSelected = ``
	__SpiralRhombus__00000000_.Transform = ``

	__SpiralRhombusGrid__00000000_.Name = `Spiral Rhombus Grid`
	__SpiralRhombusGrid__00000000_.IsDisplayed = false

	__VerticalAxis__00000000_.Name = `Vertical Axis`
	__VerticalAxis__00000000_.IsDisplayed = false
	__VerticalAxis__00000000_.AxisHandleBorderLength = 0.000000
	__VerticalAxis__00000000_.Axis_Length = 600.000000
	__VerticalAxis__00000000_.Color = ``
	__VerticalAxis__00000000_.FillOpacity = 0.000000
	__VerticalAxis__00000000_.Stroke = ``
	__VerticalAxis__00000000_.StrokeOpacity = 0.000000
	__VerticalAxis__00000000_.StrokeWidth = 1.000000
	__VerticalAxis__00000000_.StrokeDashArray = ``
	__VerticalAxis__00000000_.StrokeDashArrayWhenSelected = ``
	__VerticalAxis__00000000_.Transform = ``

	// insertion point for setup of pointers
	__Axis__00000001_.ShapeCategory = __ShapeCategory__00000004_
	__Axis__00000002_.ShapeCategory = __ShapeCategory__00000001_
	__Axis__00000004_.ShapeCategory = __ShapeCategory__00000002_
	__AxisGrid__00000000_.Reference = __Axis__00000000_
	__AxisGrid__00000000_.ShapeCategory = __ShapeCategory__00000008_
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000000_.Axiss = append(__AxisGrid__00000000_.Axiss, __Axis__00000000_)
	__AxisGrid__00000001_.Reference = __Axis__00000001_
	__AxisGrid__00000001_.ShapeCategory = __ShapeCategory__00000004_
	__AxisGrid__00000001_.Axiss = append(__AxisGrid__00000001_.Axiss, __Axis__00000000_)
	__AxisGrid__00000001_.Axiss = append(__AxisGrid__00000001_.Axiss, __Axis__00000000_)
	__AxisGrid__00000001_.Axiss = append(__AxisGrid__00000001_.Axiss, __Axis__00000000_)
	__AxisGrid__00000001_.Axiss = append(__AxisGrid__00000001_.Axiss, __Axis__00000000_)
	__AxisGrid__00000001_.Axiss = append(__AxisGrid__00000001_.Axiss, __Axis__00000000_)
	__AxisGrid__00000001_.Axiss = append(__AxisGrid__00000001_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Reference = __Axis__00000003_
	__AxisGrid__00000002_.ShapeCategory = __ShapeCategory__00000008_
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__AxisGrid__00000002_.Axiss = append(__AxisGrid__00000002_.Axiss, __Axis__00000000_)
	__Bezier__00000002_.ShapeCategory = __ShapeCategory__00000005_
	__Bezier__00000003_.ShapeCategory = __ShapeCategory__00000005_
	__Bezier__00000004_.ShapeCategory = __ShapeCategory__00000005_
	__Bezier__00000005_.ShapeCategory = __ShapeCategory__00000005_
	__BezierGrid__00000000_.Reference = __Bezier__00000000_
	__BezierGrid__00000000_.ShapeCategory = __ShapeCategory__00000009_
	__BezierGrid__00000000_.Beziers = append(__BezierGrid__00000000_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000000_.Beziers = append(__BezierGrid__00000000_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000000_.Beziers = append(__BezierGrid__00000000_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000000_.Beziers = append(__BezierGrid__00000000_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000000_.Beziers = append(__BezierGrid__00000000_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000001_.Reference = __Bezier__00000000_
	__BezierGrid__00000001_.ShapeCategory = __ShapeCategory__00000009_
	__BezierGrid__00000001_.Beziers = append(__BezierGrid__00000001_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000001_.Beziers = append(__BezierGrid__00000001_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000001_.Beziers = append(__BezierGrid__00000001_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000001_.Beziers = append(__BezierGrid__00000001_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000001_.Beziers = append(__BezierGrid__00000001_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000002_.Reference = __Bezier__00000001_
	__BezierGrid__00000002_.ShapeCategory = __ShapeCategory__00000009_
	__BezierGrid__00000002_.Beziers = append(__BezierGrid__00000002_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000002_.Beziers = append(__BezierGrid__00000002_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000002_.Beziers = append(__BezierGrid__00000002_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000002_.Beziers = append(__BezierGrid__00000002_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000002_.Beziers = append(__BezierGrid__00000002_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000003_.Reference = __Bezier__00000006_
	__BezierGrid__00000003_.ShapeCategory = __ShapeCategory__00000009_
	__BezierGrid__00000003_.Beziers = append(__BezierGrid__00000003_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000003_.Beziers = append(__BezierGrid__00000003_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000003_.Beziers = append(__BezierGrid__00000003_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000003_.Beziers = append(__BezierGrid__00000003_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000003_.Beziers = append(__BezierGrid__00000003_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000004_.Reference = __Bezier__00000005_
	__BezierGrid__00000004_.ShapeCategory = __ShapeCategory__00000005_
	__BezierGrid__00000004_.Beziers = append(__BezierGrid__00000004_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000004_.Beziers = append(__BezierGrid__00000004_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000004_.Beziers = append(__BezierGrid__00000004_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000004_.Beziers = append(__BezierGrid__00000004_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000004_.Beziers = append(__BezierGrid__00000004_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000005_.Reference = __Bezier__00000003_
	__BezierGrid__00000005_.ShapeCategory = __ShapeCategory__00000005_
	__BezierGrid__00000005_.Beziers = append(__BezierGrid__00000005_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000005_.Beziers = append(__BezierGrid__00000005_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000005_.Beziers = append(__BezierGrid__00000005_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000005_.Beziers = append(__BezierGrid__00000005_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000005_.Beziers = append(__BezierGrid__00000005_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000006_.Reference = __Bezier__00000004_
	__BezierGrid__00000006_.ShapeCategory = __ShapeCategory__00000005_
	__BezierGrid__00000006_.Beziers = append(__BezierGrid__00000006_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000006_.Beziers = append(__BezierGrid__00000006_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000006_.Beziers = append(__BezierGrid__00000006_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000006_.Beziers = append(__BezierGrid__00000006_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000006_.Beziers = append(__BezierGrid__00000006_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000007_.Reference = __Bezier__00000002_
	__BezierGrid__00000007_.ShapeCategory = __ShapeCategory__00000005_
	__BezierGrid__00000007_.Beziers = append(__BezierGrid__00000007_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000007_.Beziers = append(__BezierGrid__00000007_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000007_.Beziers = append(__BezierGrid__00000007_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000007_.Beziers = append(__BezierGrid__00000007_.Beziers, __Bezier__00000000_)
	__BezierGrid__00000007_.Beziers = append(__BezierGrid__00000007_.Beziers, __Bezier__00000000_)
	__BezierGridStack__00000000_.ShapeCategory = __ShapeCategory__00000005_
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__BezierGridStack__00000000_.BezierGrids = append(__BezierGridStack__00000000_.BezierGrids, __BezierGrid__00000000_)
	__Circle__00000000_.ShapeCategory = __ShapeCategory__00000003_
	__Circle__00000001_.ShapeCategory = __ShapeCategory__00000004_
	__Circle__00000004_.ShapeCategory = __ShapeCategory__00000001_
	__Circle__00000005_.ShapeCategory = __ShapeCategory__00000002_
	__CircleGrid__00000000_.Reference = __Circle__00000001_
	__CircleGrid__00000000_.ShapeCategory = __ShapeCategory__00000004_
	__CircleGrid__00000000_.Circles = append(__CircleGrid__00000000_.Circles, __Circle__00000000_)
	__CircleGrid__00000000_.Circles = append(__CircleGrid__00000000_.Circles, __Circle__00000000_)
	__CircleGrid__00000000_.Circles = append(__CircleGrid__00000000_.Circles, __Circle__00000000_)
	__CircleGrid__00000000_.Circles = append(__CircleGrid__00000000_.Circles, __Circle__00000000_)
	__CircleGrid__00000000_.Circles = append(__CircleGrid__00000000_.Circles, __Circle__00000000_)
	__CircleGrid__00000000_.Circles = append(__CircleGrid__00000000_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Reference = __Circle__00000002_
	__CircleGrid__00000001_.ShapeCategory = __ShapeCategory__00000009_
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000001_.Circles = append(__CircleGrid__00000001_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Reference = __Circle__00000002_
	__CircleGrid__00000002_.ShapeCategory = __ShapeCategory__00000009_
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000002_.Circles = append(__CircleGrid__00000002_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Reference = __Circle__00000000_
	__CircleGrid__00000003_.ShapeCategory = __ShapeCategory__00000003_
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000003_.Circles = append(__CircleGrid__00000003_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.ShapeCategory = __ShapeCategory__00000003_
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000004_.Circles = append(__CircleGrid__00000004_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Reference = __Circle__00000004_
	__CircleGrid__00000005_.ShapeCategory = __ShapeCategory__00000001_
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000005_.Circles = append(__CircleGrid__00000005_.Circles, __Circle__00000000_)
	__CircleGrid__00000006_.ShapeCategory = __ShapeCategory__00000002_
	__CircleGrid__00000006_.Circles = append(__CircleGrid__00000006_.Circles, __Circle__00000000_)
	__CircleGrid__00000006_.Circles = append(__CircleGrid__00000006_.Circles, __Circle__00000000_)
	__CircleGrid__00000006_.Circles = append(__CircleGrid__00000006_.Circles, __Circle__00000000_)
	__CircleGrid__00000006_.Circles = append(__CircleGrid__00000006_.Circles, __Circle__00000000_)
	__CircleGrid__00000006_.Circles = append(__CircleGrid__00000006_.Circles, __Circle__00000000_)
	__CircleGrid__00000006_.Circles = append(__CircleGrid__00000006_.Circles, __Circle__00000000_)
	__CircleGrid__00000006_.Circles = append(__CircleGrid__00000006_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Reference = __Circle__00000002_
	__CircleGrid__00000007_.ShapeCategory = __ShapeCategory__00000009_
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000007_.Circles = append(__CircleGrid__00000007_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Reference = __Circle__00000002_
	__CircleGrid__00000008_.ShapeCategory = __ShapeCategory__00000009_
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__CircleGrid__00000008_.Circles = append(__CircleGrid__00000008_.Circles, __Circle__00000000_)
	__Content__00000000_.Chapters = append(__Content__00000000_.Chapters, __Chapter__00000000_)
	__Content__00000000_.Chapters = append(__Content__00000000_.Chapters, __Chapter__00000001_)
	__ExportToMusicxml__00000000_.Parameter = __Parameter__00000000_
	__FrontCurveStack__00000000_.ShapeCategory = __ShapeCategory__00000007_
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000008_)
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000009_)
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000010_)
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000011_)
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000012_)
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000013_)
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000014_)
	__FrontCurveStack__00000000_.FrontCurves = append(__FrontCurveStack__00000000_.FrontCurves, __FrontCurve__00000015_)
	__HorizontalAxis__00000000_.ShapeCategory = __ShapeCategory__00000000_
	__Key__00000000_.ShapeCategory = __ShapeCategory__00000008_
	__Parameter__00000000_.InitialRhombus = __Rhombus__00000001_
	__Parameter__00000000_.InitialCircle = __Circle__00000004_
	__Parameter__00000000_.InitialRhombusGrid = __RhombusGrid__00000001_
	__Parameter__00000000_.InitialCircleGrid = __CircleGrid__00000005_
	__Parameter__00000000_.InitialAxis = __Axis__00000002_
	__Parameter__00000000_.RotatedAxis = __Axis__00000004_
	__Parameter__00000000_.RotatedRhombus = __Rhombus__00000003_
	__Parameter__00000000_.RotatedRhombusGrid = __RhombusGrid__00000002_
	__Parameter__00000000_.RotatedCircleGrid = __CircleGrid__00000006_
	__Parameter__00000000_.NextRhombus = __Rhombus__00000002_
	__Parameter__00000000_.NextCircle = __Circle__00000005_
	__Parameter__00000000_.GrowingRhombusGridSeed = __Rhombus__00000000_
	__Parameter__00000000_.GrowingRhombusGrid = __RhombusGrid__00000000_
	__Parameter__00000000_.GrowingCircleGridSeed = __Circle__00000000_
	__Parameter__00000000_.GrowingCircleGrid = __CircleGrid__00000003_
	__Parameter__00000000_.GrowingCircleGridLeftSeed = __Circle__00000003_
	__Parameter__00000000_.GrowingCircleGridLeft = __CircleGrid__00000004_
	__Parameter__00000000_.ConstructionAxis = __Axis__00000001_
	__Parameter__00000000_.ConstructionAxisGrid = __AxisGrid__00000001_
	__Parameter__00000000_.ConstructionCircle = __Circle__00000001_
	__Parameter__00000000_.ConstructionCircleGrid = __CircleGrid__00000000_
	__Parameter__00000000_.GrowthCurveSeed = __Bezier__00000005_
	__Parameter__00000000_.GrowthCurve = __BezierGrid__00000004_
	__Parameter__00000000_.GrowthCurveShiftedRightSeed = __Bezier__00000002_
	__Parameter__00000000_.GrowthCurveShiftedRight = __BezierGrid__00000007_
	__Parameter__00000000_.GrowthCurveNextSeed = __Bezier__00000003_
	__Parameter__00000000_.GrowthCurveNext = __BezierGrid__00000005_
	__Parameter__00000000_.GrowthCurveNextShiftedRightSeed = __Bezier__00000004_
	__Parameter__00000000_.GrowthCurveNextShiftedRight = __BezierGrid__00000006_
	__Parameter__00000000_.GrowthCurveStack = __BezierGridStack__00000000_
	__Parameter__00000000_.SpiralRhombusGridSeed = __SpiralRhombus__00000000_
	__Parameter__00000000_.SpiralRhombusGrid = __SpiralRhombusGrid__00000000_
	__Parameter__00000000_.SpiralCircleSeed = __SpiralCircle__00000000_
	__Parameter__00000000_.SpiralCircleGrid = __SpiralCircleGrid__00000003_
	__Parameter__00000000_.SpiralCircleFullGrid = __SpiralCircleGrid__00000001_
	__Parameter__00000000_.SpiralConstructionOuterLineSeed = __SpiralLine__00000001_
	__Parameter__00000000_.SpiralConstructionInnerLineSeed = __SpiralLine__00000000_
	__Parameter__00000000_.SpiralConstructionOuterLineGrid = __SpiralLineGrid__00000002_
	__Parameter__00000000_.SpiralConstructionInnerLineGrid = __SpiralLineGrid__00000000_
	__Parameter__00000000_.SpiralConstructionCircleGrid = __SpiralCircleGrid__00000002_
	__Parameter__00000000_.SpiralConstructionOuterLineFullGrid = __SpiralLineGrid__00000001_
	__Parameter__00000000_.SpiralBezierSeed = __SpiralBezier__00000000_
	__Parameter__00000000_.SpiralBezierGrid = __SpiralBezierGrid__00000001_
	__Parameter__00000000_.SpiralBezierFullGrid = __SpiralBezierGrid__00000000_
	__Parameter__00000000_.FrontCurveStack = __FrontCurveStack__00000000_
	__Parameter__00000000_.Fkey = __Key__00000000_
	__Parameter__00000000_.PitchLines = __AxisGrid__00000002_
	__Parameter__00000000_.BeatLines = __AxisGrid__00000000_
	__Parameter__00000000_.FirstVoice = __BezierGrid__00000002_
	__Parameter__00000000_.FirstVoiceShiftedRigth = __BezierGrid__00000003_
	__Parameter__00000000_.SecondVoice = __BezierGrid__00000000_
	__Parameter__00000000_.SecondVoiceShiftedRight = __BezierGrid__00000001_
	__Parameter__00000000_.FirstVoiceNotes = __CircleGrid__00000002_
	__Parameter__00000000_.FirstVoiceNotesShiftedRight = __CircleGrid__00000001_
	__Parameter__00000000_.SecondVoiceNotes = __CircleGrid__00000008_
	__Parameter__00000000_.SecondVoiceNotesShiftedRight = __CircleGrid__00000007_
	__Parameter__00000000_.HorizontalAxis = __HorizontalAxis__00000000_
	__Parameter__00000000_.VerticalAxis = __VerticalAxis__00000000_
	__Parameter__00000000_.SpiralOrigin = __SpiralOrigin__00000000_
	__Rhombus__00000000_.ShapeCategory = __ShapeCategory__00000003_
	__Rhombus__00000001_.ShapeCategory = __ShapeCategory__00000001_
	__Rhombus__00000002_.ShapeCategory = __ShapeCategory__00000002_
	__Rhombus__00000003_.ShapeCategory = __ShapeCategory__00000002_
	__RhombusGrid__00000000_.Reference = __Rhombus__00000003_
	__RhombusGrid__00000000_.ShapeCategory = __ShapeCategory__00000003_
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000000_.Rhombuses = append(__RhombusGrid__00000000_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Reference = __Rhombus__00000001_
	__RhombusGrid__00000001_.ShapeCategory = __ShapeCategory__00000001_
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000001_.Rhombuses = append(__RhombusGrid__00000001_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000002_.ShapeCategory = __ShapeCategory__00000002_
	__RhombusGrid__00000002_.Rhombuses = append(__RhombusGrid__00000002_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000002_.Rhombuses = append(__RhombusGrid__00000002_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000002_.Rhombuses = append(__RhombusGrid__00000002_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000002_.Rhombuses = append(__RhombusGrid__00000002_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000002_.Rhombuses = append(__RhombusGrid__00000002_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000002_.Rhombuses = append(__RhombusGrid__00000002_.Rhombuses, __Rhombus__00000000_)
	__RhombusGrid__00000002_.Rhombuses = append(__RhombusGrid__00000002_.Rhombuses, __Rhombus__00000000_)
	__SpiralBezier__00000000_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralBezierGrid__00000000_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000000_.SpiralBeziers = append(__SpiralBezierGrid__00000000_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000001_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralBezierGrid__00000001_.SpiralBeziers = append(__SpiralBezierGrid__00000001_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000001_.SpiralBeziers = append(__SpiralBezierGrid__00000001_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000001_.SpiralBeziers = append(__SpiralBezierGrid__00000001_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000001_.SpiralBeziers = append(__SpiralBezierGrid__00000001_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralBezierGrid__00000001_.SpiralBeziers = append(__SpiralBezierGrid__00000001_.SpiralBeziers, __SpiralBezier__00000000_)
	__SpiralCircle__00000000_.ShapeCategory = __ShapeCategory__00000006_
	__SpiralCircleGrid__00000000_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralCircleGrid__00000001_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000001_.SpiralCircles = append(__SpiralCircleGrid__00000001_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000002_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralCircleGrid__00000002_.SpiralCircles = append(__SpiralCircleGrid__00000002_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000002_.SpiralCircles = append(__SpiralCircleGrid__00000002_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000002_.SpiralCircles = append(__SpiralCircleGrid__00000002_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000002_.SpiralCircles = append(__SpiralCircleGrid__00000002_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000002_.SpiralCircles = append(__SpiralCircleGrid__00000002_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000002_.SpiralCircles = append(__SpiralCircleGrid__00000002_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.ShapeCategory = __ShapeCategory__00000006_
	__SpiralCircleGrid__00000003_.SpiralRhombusGrid = __SpiralRhombusGrid__00000000_
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralCircleGrid__00000003_.SpiralCircles = append(__SpiralCircleGrid__00000003_.SpiralCircles, __SpiralCircle__00000000_)
	__SpiralLine__00000000_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralLine__00000001_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralLineGrid__00000000_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralLineGrid__00000000_.SpiralLines = append(__SpiralLineGrid__00000000_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000000_.SpiralLines = append(__SpiralLineGrid__00000000_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000000_.SpiralLines = append(__SpiralLineGrid__00000000_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000000_.SpiralLines = append(__SpiralLineGrid__00000000_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000000_.SpiralLines = append(__SpiralLineGrid__00000000_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000000_.SpiralLines = append(__SpiralLineGrid__00000000_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000001_.SpiralLines = append(__SpiralLineGrid__00000001_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000002_.ShapeCategory = __ShapeCategory__00000007_
	__SpiralLineGrid__00000002_.SpiralLines = append(__SpiralLineGrid__00000002_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000002_.SpiralLines = append(__SpiralLineGrid__00000002_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000002_.SpiralLines = append(__SpiralLineGrid__00000002_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000002_.SpiralLines = append(__SpiralLineGrid__00000002_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000002_.SpiralLines = append(__SpiralLineGrid__00000002_.SpiralLines, __SpiralLine__00000000_)
	__SpiralLineGrid__00000002_.SpiralLines = append(__SpiralLineGrid__00000002_.SpiralLines, __SpiralLine__00000000_)
	__SpiralOrigin__00000000_.ShapeCategory = __ShapeCategory__00000000_
	__SpiralRhombus__00000000_.ShapeCategory = __ShapeCategory__00000006_
	__SpiralRhombusGrid__00000000_.ShapeCategory = __ShapeCategory__00000006_
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__SpiralRhombusGrid__00000000_.SpiralRhombuses = append(__SpiralRhombusGrid__00000000_.SpiralRhombuses, __SpiralRhombus__00000000_)
	__VerticalAxis__00000000_.ShapeCategory = __ShapeCategory__00000000_
}
