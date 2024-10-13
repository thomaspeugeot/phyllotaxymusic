package models

import (
	"fmt"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// fetch the control points of the the first bezier curve and
// convert them to spiral coordinates
func (p *Parameter) ComputeSpiralBezierSeed() {

	sc0 := p.SpiralConstructionCircleGrid.SpiralCircles[0]
	sl0 := p.SpiralConstructionOuterLineGrid.SpiralLines[0]

	sc1 := p.SpiralConstructionCircleGrid.SpiralCircles[1]
	sl1 := p.SpiralConstructionOuterLineGrid.SpiralLines[1]

	sb := p.SpiralBezierSeed

	p.spiralCircleLinesToSpiralBezier(sb, sc0, sl0, sc1, sl1)
}

func (p *Parameter) spiralCircleLinesToSpiralBezier(
	sb *SpiralBezier,
	sc0 *SpiralCircle, sl0 *SpiralLine,
	sc1 *SpiralCircle, sl1 *SpiralLine) {
	sb.StartX, sb.StartY = sc0.CenterX, sc0.CenterY

	dx0 := p.SpiralBezierStrength * (sl0.EndX - sl0.StartX)
	dy0 := p.SpiralBezierStrength * (sl0.EndY - sl0.StartY)

	newEndX0 := -dy0 + sl0.StartX
	newEndY0 := dx0 + sl0.StartY
	sb.ControlPointStartX, sb.ControlPointStartY = newEndX0, newEndY0

	sb.EndX, sb.EndY = sc1.CenterX, sc1.CenterY

	dx1 := p.SpiralBezierStrength * (sl1.EndX - sl1.StartX)
	dy1 := p.SpiralBezierStrength * (sl1.EndY - sl1.StartY)

	newEndX1 := dy1 + sl1.StartX
	newEndY1 := -dx1 + sl1.StartY
	sb.ControlPointEndX, sb.ControlPointEndY = newEndX1, newEndY1
}

func (p *Parameter) ComputeSpiralBezierGrid() {

	p.SpiralBezierGrid.SpiralBeziers = p.SpiralBezierGrid.SpiralBeziers[:0]

	nm1 := len(p.SpiralConstructionCircleGrid.SpiralCircles) - 1
	for i := range nm1 {
		sc0 := p.SpiralConstructionCircleGrid.SpiralCircles[i]
		sl0 := p.SpiralConstructionOuterLineGrid.SpiralLines[i]

		sc1 := p.SpiralConstructionCircleGrid.SpiralCircles[(i+1)%nm1]
		sl1 := p.SpiralConstructionOuterLineGrid.SpiralLines[(i+1)%nm1]

		sb := new(SpiralBezier)

		sb.Stroke = gongsvg_models.Grey.ToString()
		sb.StrokeWidth = 2.0
		sb.StrokeOpacity = 0.8

		p.spiralCircleLinesToSpiralBezier(sb, sc0, sl0, sc1, sl1)

		p.SpiralBezierGrid.SpiralBeziers = append(p.SpiralBezierGrid.SpiralBeziers, sb)
	}

}

func (p *Parameter) ComputeSpiralBezierFullGrid() {

	p.SpiralBezierFullGrid.SpiralBeziers = p.SpiralBezierFullGrid.SpiralBeziers[:0]

	nm1 := len(p.SpiralConstructionCircleGrid.SpiralCircles) - 1

	// for i := range 1 {
	for i := range p.Z - p.ShiftToNearestCircle - 1 {

		// construct the front curve
		// for k := range 1 {
		for k := range p.M + p.N - 1 {

			// pick the ith circle
			sc0 := p.SpiralCircleFullGrid.SpiralCircles[i+k]
			sl0 := p.SpiralConstructionOuterLineFullGrid.SpiralLines[i+k]

			nextIndex := (k + p.ShiftToNearestCircle) % nm1

			sc1 := p.SpiralCircleFullGrid.SpiralCircles[i+nextIndex]
			sl1 := p.SpiralConstructionOuterLineFullGrid.SpiralLines[i+nextIndex]

			sb := new(SpiralBezier)

			sb.Stroke = gongsvg_models.Grey.ToString()
			sb.StrokeWidth = 2.0
			sb.StrokeOpacity = 0.5

			p.spiralCircleLinesToSpiralBezier(sb, sc0, sl0, sc1, sl1)

			p.SpiralBezierFullGrid.SpiralBeziers = append(p.SpiralBezierFullGrid.SpiralBeziers, sb)
		}
	}
}

func (p *Parameter) ComputeFrontCurveStacks(stage *StageStruct) {

	for frontCurve := range *GetGongstructInstancesSet[FrontCurve](stage) {
		frontCurve.Unstage(stage)
	}

	p.FrontCurveStack.FrontCurves = p.FrontCurveStack.FrontCurves[:0]
	p.RotatedFrontCurveStack.FrontCurves = p.RotatedFrontCurveStack.FrontCurves[:0]
	p.FrontCurveStack.SpiralCircles = p.FrontCurveStack.SpiralCircles[:0]

	for idx, bezierGrid := range p.GrowthCurveStack.BezierGrids {

		var xs, ys []float64

		for _, bezier := range bezierGrid.Beziers {

			for i := range p.NbInterpolationPoints {
				length := bezier.EndX - bezier.StartX

				x := bezier.StartX + length*float64(i)/float64(p.NbInterpolationPoints)
				y, _ := bezier.ComputeYFromX(x)

				spiralCircle := new(SpiralCircle)
				spiralCircle.Stroke = gongsvg_models.Green.ToString()
				spiralCircle.StrokeWidth = 1
				spiralCircle.StrokeOpacity = 1

				x_r, y_r := p.convertToSpiralCoords(x, y)

				spiralCircle.CenterX = x_r
				spiralCircle.CenterY = y_r
				spiralCircle.HasBespokeRadius = true
				spiralCircle.BespopkeRadius = 3

				xs = append(xs, p.SpiralOriginX+x_r)
				ys = append(ys, p.SpiralOriginY-y_r)

				if p.ShowInterpolationPoints {
					p.FrontCurveStack.SpiralCircles = append(p.FrontCurveStack.SpiralCircles, spiralCircle)
				}
			}
		}

		if p.FrontCurveStack.IsDisplayed {
			frontCurve := new(FrontCurve)
			str := GenerateSmoothSVGPath(xs, ys, 0, 0, 0, 0, 0)
			frontCurve.Path = str
			frontCurve.Name = fmt.Sprintf("Non Rotated %d ", idx)
			frontCurve.Stage(stage)

			p.FrontCurveStack.FrontCurves = append(p.FrontCurveStack.FrontCurves, frontCurve)
		}

		if idx == 0 {
			frontCurve := new(FrontCurve)
			str := GenerateSmoothSVGPath(xs, ys, 0, 0, p.SpiralOriginX, p.SpiralOriginY, p.HourHandleRotationAngle)
			frontCurve.Path = str
			frontCurve.Name = fmt.Sprintf("Rotated %d ", idx)

			frontCurve.Stage(stage)

			p.RotatedFrontCurveStack.FrontCurves = append(p.FrontCurveStack.FrontCurves, frontCurve)
		}

	}
}
