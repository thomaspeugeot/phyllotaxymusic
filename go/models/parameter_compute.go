package models

import (
	"cmp"
	"fmt"
	"math"
	"slices"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func (p *Parameter) ComputeInitialRhombus() {
	p.InitialRhombus.InsideAngle = p.InsideAngle
	p.InitialRhombus.SideLength = p.SideLength
	p.InitialRhombus.CenterX = 0
	p.InitialRhombus.CenterY = 0
}

func (p *Parameter) ComputeInitialCircle() {
	p.InitialCircle.CenterX = 0
	p.InitialCircle.CenterY = 0
}

func (p *Parameter) ComputeInitialRhombusGrid(stage *StageStruct) {

	// remove the attached rhombus
	for _, r := range p.InitialRhombusGrid.Rhombuses {
		r.Unstage(stage)
	}
	p.InitialRhombusGrid.Rhombuses = p.InitialRhombusGrid.Rhombuses[:0]

	g := p.InitialRhombusGrid

	angleRad := g.Reference.AngleDegree * math.Pi / 180
	_ = angleRad
	insideAngleRad := g.Reference.InsideAngle * math.Pi / 180
	sinHalfInsideAngle := math.Sin(insideAngleRad / 2)
	cosHalfInsideAngle := math.Cos(insideAngleRad / 2)

	for i := range p.N + 1 {
		for j := range p.M + 1 {
			r := new(Rhombus) // .Stage(stage)
			*r = *g.Reference
			g.Rhombuses = append(g.Rhombuses, r)
			r.Name = fmt.Sprintf("%d %d", i, j)
			r.CenterX +=
				float64(i)*r.SideLength*cosHalfInsideAngle -
					float64(j)*r.SideLength*cosHalfInsideAngle
			r.CenterY +=
				float64(i)*r.SideLength*sinHalfInsideAngle +
					float64(j)*r.SideLength*sinHalfInsideAngle
		}
	}
}

func (p *Parameter) ComputeInitialCircleGrid(stage *StageStruct) {

	// remove the attached rhombus
	for _, c := range p.InitialCircleGrid.Circles {
		c.Unstage(stage)
	}
	p.InitialCircleGrid.Circles = p.InitialCircleGrid.Circles[:0]

	g := p.InitialCircleGrid

	insideAngleRad := p.InsideAngle * math.Pi / 180
	sinHalfInsideAngle := math.Sin(insideAngleRad / 2)
	cosHalfInsideAngle := math.Cos(insideAngleRad / 2)

	for i := range p.N + 1 {
		for j := range p.M + 1 {
			c := new(Circle) // .Stage(stage)
			*c = *g.Reference
			g.Circles = append(g.Circles, c)
			c.Name = fmt.Sprintf("%d %d", i, j)
			c.CenterX +=
				float64(i)*p.SideLength*cosHalfInsideAngle -
					float64(j)*p.SideLength*cosHalfInsideAngle
			c.CenterY +=
				float64(i)*p.SideLength*sinHalfInsideAngle +
					float64(j)*p.SideLength*sinHalfInsideAngle
		}
	}
}

func (p *Parameter) ComputeInitialAxis() {
	insideAngleRad := p.InsideAngle * math.Pi / 180
	sinHalfInsideAngle := math.Sin(insideAngleRad / 2)
	cosHalfInsideAngle := math.Cos(insideAngleRad / 2)
	sideLength := p.InitialRhombus.SideLength

	y := float64(p.N)*sideLength*sinHalfInsideAngle +
		float64(p.M)*sideLength*sinHalfInsideAngle
	x := float64(p.N)*sideLength*cosHalfInsideAngle -
		float64(p.M)*sideLength*cosHalfInsideAngle

	p.InitialAxis.AngleDegree = math.Atan2(y, x) * 180 / math.Pi
	p.InitialAxis.Length = math.Sqrt(x*x + y*y)
}

func (p *Parameter) computeRotatedAxis() {
	p.RotatedAxis.Length = p.InitialAxis.Length
	p.RotatedAxis.AngleDegree = 0
}

func (p *Parameter) computeRotatedRhombus() {
	p.RotatedRhombus.SideLength = p.InitialRhombus.SideLength
	p.RotatedRhombus.InsideAngle = p.InitialRhombus.InsideAngle
	p.RotatedRhombus.AngleDegree = -p.InitialAxis.AngleDegree
}

func (p *Parameter) computeRotatedRhombusGrid(stage *StageStruct) {
	g := p.RotatedRhombusGrid

	// remove the attached rhombus
	for _, r := range g.Rhombuses {
		r.Unstage(stage)
	}
	g.Rhombuses = g.Rhombuses[:0]

	angleRad := -DegreesToRadians(p.InitialAxis.AngleDegree)
	cosAngle := math.Cos(angleRad)
	sinAngle := math.Sin(angleRad)

	for _, _r := range p.InitialRhombusGrid.Rhombuses {
		r := new(Rhombus) // .Stage(stage)
		*r = *_r
		r.SideLength = _r.SideLength
		r.InsideAngle = _r.InsideAngle
		r.AngleDegree = -p.InitialAxis.AngleDegree

		r.Name += " Rotated"

		r.CenterX = _r.CenterX*cosAngle - _r.CenterY*sinAngle
		r.CenterY = _r.CenterX*sinAngle + _r.CenterY*cosAngle

		// keep only rhombus above 0
		if r.CenterY < -0.00001 {
			continue
		}

		g.Rhombuses = append(g.Rhombuses, r)

	}
}

func (p *Parameter) computeRotatedCircleGrid(stage *StageStruct) {
	g := p.RotatedCircleGrid

	// remove the attached Circle
	for _, r := range g.Circles {
		r.Unstage(stage)
	}
	g.Circles = g.Circles[:0]

	angleRad := -DegreesToRadians(p.InitialAxis.AngleDegree)
	cosAngle := math.Cos(angleRad)
	sinAngle := math.Sin(angleRad)

	for _, _c := range p.InitialCircleGrid.Circles {
		c := new(Circle) // .Stage(stage)
		*c = *_c

		c.Name += " Rotated"

		c.CenterX = _c.CenterX*cosAngle - _c.CenterY*sinAngle
		c.CenterY = _c.CenterX*sinAngle + _c.CenterY*cosAngle

		// keep only rhombus above 0
		if c.CenterY < -0.00001 {
			continue
		}
		g.Circles = append(g.Circles, c)

	}
}

func (p *Parameter) ComputeNextRhombus() {

	// parse all rhombus in the rotated rhombus grid
	// and get the first that is not with a non nil ordinate
	slices.SortFunc(p.RotatedRhombusGrid.Rhombuses,
		func(r1, r2 *Rhombus) int {
			if r1.CenterY > r2.CenterY {
				return 1
			} else {
				return -1
			}

		})
	p.NextRhombus.CenterX = p.RotatedRhombusGrid.Rhombuses[2].CenterX
	p.NextRhombus.CenterY = p.RotatedRhombusGrid.Rhombuses[2].CenterY
	p.NextRhombus.SideLength = p.SideLength
	p.NextRhombus.AngleDegree = p.RotatedRhombus.AngleDegree
	p.NextRhombus.InsideAngle = p.InsideAngle
}

func (p *Parameter) ComputeNextCircle() {

	// parse all circle in the rotated circle grid
	// and get the first that is not with a non nil ordinate
	slices.SortFunc(p.RotatedCircleGrid.Circles,
		func(r1, r2 *Circle) int {
			if r1.CenterY > r2.CenterY {
				return 1
			} else {
				return -1
			}

		})
	p.NextCircle.CenterX = p.RotatedCircleGrid.Circles[2].CenterX
	p.NextCircle.CenterY = p.RotatedCircleGrid.Circles[2].CenterY
}

func (p *Parameter) ComputeGrowingRhombusGrid() {

	// configure seed rhombus
	p.GrowingRhombusGridSeed.SideLength = p.RotatedRhombus.SideLength
	p.GrowingRhombusGridSeed.AngleDegree = p.RotatedRhombus.AngleDegree
	p.GrowingRhombusGridSeed.InsideAngle = p.RotatedRhombus.InsideAngle

	p.GrowingRhombusGrid.Rhombuses = p.GrowingRhombusGrid.Rhombuses[:0]
	for i := range p.Z {
		r := new(Rhombus) // .Stage(stage)
		*r = *p.GrowingRhombusGridSeed

		// compute
		x := float64(i) * p.NextRhombus.CenterX
		y := float64(i) * p.NextRhombus.CenterY

		nbRotations := int(x / p.RotatedAxis.Length)
		r.CenterX = x - float64(nbRotations)*p.RotatedAxis.Length
		r.CenterY = y

		p.GrowingRhombusGrid.Rhombuses = append(p.GrowingRhombusGrid.Rhombuses, r)
	}
}

func (p *Parameter) ComputeGrowingCircleGrid() {

	// configure seed circle

	p.GrowingCircleGrid.Circles = p.GrowingCircleGrid.Circles[:0]
	for i := range p.Z {
		r := new(Circle) // .Stage(stage)
		*r = *p.GrowingCircleGridSeed

		// compute
		x := float64(i) * p.NextCircle.CenterX
		y := float64(i) * p.NextCircle.CenterY

		nbRotations := int(x / p.RotatedAxis.Length)
		r.CenterX = x - float64(nbRotations)*p.RotatedAxis.Length
		r.CenterY = y

		p.GrowingCircleGrid.Circles = append(p.GrowingCircleGrid.Circles, r)
	}
	*p.GrowingCircleGridSeed = *p.GrowingCircleGrid.Circles[0]
}

func (p *Parameter) ComputeGrowingCircleGridLeft() {

	// configure seed circle

	p.GrowingCircleGridLeft.Circles = p.GrowingCircleGridLeft.Circles[:0]
	for i := range p.Z {
		r := new(Circle) // .Stage(stage)
		*r = *p.GrowingCircleGridLeftSeed

		r.CenterX = p.GrowingCircleGrid.Circles[i].CenterX - p.RotatedAxis.Length
		r.CenterY = p.GrowingCircleGrid.Circles[i].CenterY

		p.GrowingCircleGridLeft.Circles = append(p.GrowingCircleGridLeft.Circles, r)
	}
}

func (p *Parameter) computeConstructionAxis() {

	// get the N+M circles
	circleNPlusM := p.GrowingCircleGrid.Circles[p.M+p.N]
	x := circleNPlusM.CenterX - p.RotatedAxis.Length
	y := circleNPlusM.CenterY

	p.ConstructionAxis.Length = math.Sqrt(x*x + y*y)
	p.ConstructionAxis.AngleDegree = math.Atan2(y, x) * 180 / math.Pi

	p.ConstructionAxis.EndX = x
	p.ConstructionAxis.EndY = y
}

func (p *Parameter) computeConstructionCircle() {

	circleNPlusM := p.GrowingCircleGrid.Circles[p.M+p.N]
	p.ConstructionCircle.CenterX = (circleNPlusM.CenterX - p.RotatedAxis.Length) / 2.0
	p.ConstructionCircle.CenterY = circleNPlusM.CenterY / 2.0
}

func (p *Parameter) computeConstructionAxisGrid() {

	g := p.ConstructionAxisGrid
	g.Axiss = g.Axiss[:0]

	for i := range p.M + p.N {

		a := new(Axis)
		*a = *p.ConstructionAxis
		g.Axiss = append(g.Axiss, a)

		// apply growing circle coordinates
		c := p.GrowingCircleGrid.Circles[i]
		a.CenterX += c.CenterX
		a.CenterY += c.CenterY
	}

	a := new(Axis)
	*a = *p.ConstructionAxis
	g.Axiss = append(g.Axiss, a)
	a.CenterX += p.RotatedAxis.Length

	slices.SortFunc(g.Axiss, func(c1, c2 *Axis) int {
		return cmp.Compare(c1.CenterX, c2.CenterX)
	})
}

func (p *Parameter) computeConstructionCircleGrid() {

	g := p.ConstructionCircleGrid
	g.Circles = g.Circles[:0]

	for i := range p.M + p.N {

		_c := new(Circle)
		*_c = *p.ConstructionCircle
		g.Circles = append(g.Circles, _c)

		// apply growing circle coordinates
		c := p.GrowingCircleGrid.Circles[i]
		_c.CenterX += c.CenterX
		_c.CenterY += c.CenterY
	}

	c := new(Circle)
	*c = *p.ConstructionCircle
	g.Circles = append(g.Circles, c)
	c.CenterX += p.RotatedAxis.Length

	slices.SortFunc(g.Circles, func(c1, c2 *Circle) int {
		return cmp.Compare(c1.CenterX, c2.CenterX)
	})

	// for _, c := range g.Circles {
	// 	log.Println(c.CenterX)
	// }
}

func (p *Parameter) ComputeGrowthCurveSegment() {

	b := p.GrowthCurveSegment
	_ = b

	p.computeBezier(b, p.ConstructionCircleGrid.Circles[0], p.ConstructionCircleGrid.Circles[1])

}

// The function computeBezier calculates and sets the control points and endpoints
// for a cubic Bezier curve (b) based on two circles (startCircle and endCircle)
// and parameters (p).
// This function essentially constructs a cubic Bezier curve, where the starting and ending
// points are located at the centers of the two given circles, and the control points are computed
// based on the angle and side length provided by the Parameter struct (p).
func (p *Parameter) computeBezier(b *Bezier, startCircle, endCircle *Circle) {
	b.StartX = startCircle.CenterX
	b.StartY = startCircle.CenterY

	b.EndX = endCircle.CenterX

	b.EndY = endCircle.CenterY

	angleRad := p.ConstructionAxis.AngleDegree*math.Pi/180 - math.Pi/2.0

	b.ControlPointStartX = b.StartX +
		p.SideLength*p.BezierControlLengthRatio*math.Cos(angleRad)
	b.ControlPointStartY = b.StartY +
		p.SideLength*p.BezierControlLengthRatio*math.Sin(angleRad)

	b.ControlPointEndX = b.EndX +
		p.SideLength*p.BezierControlLengthRatio*math.Cos(angleRad+math.Pi)
	b.ControlPointEndY = b.EndY +
		p.SideLength*p.BezierControlLengthRatio*math.Sin(angleRad+math.Pi)
}

func (p *Parameter) ComputeGrowthCurve() {

	g := p.GrowthCurve
	g.Beziers = g.Beziers[:0]

	for i := range p.M + p.N {
		_b := new(Bezier)
		*_b = *p.GrowthCurveSegment
		g.Beziers = append(g.Beziers, _b)

		// apply growing bezier coordinates
		c := p.ConstructionCircleGrid

		p.computeBezier(_b, c.Circles[i], c.Circles[i+1])
	}
}

func (p *Parameter) ComputeSpiralRhombusSeed() {

	r := p.GrowingRhombusGridSeed

	x_s, y_s := r.getCoordinates()
	x_r, y_r := p.convertToCircleSpaceCoordsArray(x_s, y_s)
	p.SpiralRhombusGridSeed.GenerateCoordinates(x_r, y_r)

}

// transform all rhombus into spiral rhombus
func (p *Parameter) ComputeSpiralRhombusGrid() {

	p.SpiralRhombusGrid.SpiralRhombuses = p.SpiralRhombusGrid.SpiralRhombuses[:0]

	for idx, r := range p.GrowingRhombusGrid.Rhombuses {

		sr := new(SpiralRhombus)
		x_s, y_s := r.getCoordinates()
		x_r, y_r := p.convertToCircleSpaceCoordsArray(x_s, y_s)
		sr.GenerateCoordinates(x_r, y_r)

		sr.Stroke = GenerateColor(idx % len(colors))
		sr.StrokeWidth = 1
		sr.StrokeOpacity = 1

		p.SpiralRhombusGrid.SpiralRhombuses = append(p.SpiralRhombusGrid.SpiralRhombuses, sr)
	}

}

func (p *Parameter) ComputeFKey() {

}

func (p *Parameter) ComputePitchLines() {

	g := p.PitchLines
	g.Axiss = g.Axiss[0:]

	for i := range p.NbPitchLines {

		if p.IsMinor {
			// remove minor notes
			switch i % 12 {
			case 1, 4, 6, 9, 10:
				continue
			}
		} else {
			// remove minor notes
			switch i % 12 {
			case 1, 3, 6, 8, 10:
				continue
			}
		}

		a := new(Axis)
		*a = *p.PitchLines.Reference

		g.Axiss = append(g.Axiss, a)

		a.CenterY = float64(i) * p.PitchHeight * p.SideLength

		if i%12 == 0 {
			a.StrokeWidth *= 2
		}
	}
}

func (p *Parameter) ComputeMeasureLines() {

	g := p.MeasureLines
	g.Axiss = g.Axiss[0:]

	for i := range p.NbMeasureLines {
		a := new(Axis)
		*a = *p.MeasureLines.Reference

		g.Axiss = append(g.Axiss, a)

		a.CenterX = float64(i) * p.RotatedAxis.Length / float64(p.NbMeasureLinesPerCurve)

		if i%p.NbMeasureLinesPerCurve == 0 {
			a.StrokeWidth *= 2
		}
	}
}

func (p *Parameter) ComputeSpiralCircleSeed() {

	c := p.GrowingCircleGridSeed

	x_r, y_r := p.convertToSpiralCoords(c.CenterX, c.CenterY)

	p.SpiralCircleSeed.CenterX = x_r
	p.SpiralCircleSeed.CenterY = y_r
}

func (p *Parameter) ComputeSpiralCircleGrid() {

	p.SpiralCircleGrid.SpiralCircles = p.SpiralCircleGrid.SpiralCircles[:0]

	for idx, c := range p.GrowingCircleGrid.Circles {

		x_r, y_r := p.convertToSpiralCoords(c.CenterX, c.CenterY)

		sc := new(SpiralCircle)
		sc.Stroke = GenerateColor(idx % len(colors))
		sc.StrokeOpacity = 0.5
		sc.StrokeWidth = 2

		sc.CenterX = x_r
		sc.CenterY = y_r

		p.SpiralCircleGrid.SpiralCircles = append(p.SpiralCircleGrid.SpiralCircles, sc)
	}
}

func (p *Parameter) computeSpiralConstructionAxis() {

	cc := p.ConstructionCircle
	ca := p.ConstructionAxis
	sl := p.SpiralConstructionLine

	p.verticalAxisToSpiralLine(cc, ca, sl)
}

func (p *Parameter) verticalAxisToSpiralLine(cc *Circle, ca *Axis, sl *SpiralLine) {
	sl.StartX,
		sl.StartY =
		p.convertToSpiralCoords(cc.CenterX, cc.CenterY)

	ca_endX, ca_endY :=
		cc.CenterX+ca.Length/2.0*math.Cos(DegreesToRadians(ca.AngleDegree)),
		cc.CenterY+ca.Length/2.0*math.Sin(DegreesToRadians(ca.AngleDegree))

	sl.EndX,
		sl.EndY =
		p.convertToSpiralCoords(ca_endX, ca_endY)
}

func (p *Parameter) computeSpiralConstructionAxisGrid() {

	p.SpiralConstructionLineGrid.SpiralLines =
		p.SpiralConstructionLineGrid.SpiralLines[:0]
	for i, ca := range p.ConstructionAxisGrid.Axiss {

		sl := new(SpiralLine)
		sl.Name = fmt.Sprintf("Spiral Axis %d", i)
		sl.Stroke = GenerateColor(i)
		sl.Stroke = gongsvg_models.Black.ToString()
		sl.StrokeWidth = 1
		sl.StrokeOpacity = 1

		cc := p.ConstructionCircleGrid.Circles[i]

		p.verticalAxisToSpiralLine(cc, ca, sl)

		p.SpiralConstructionLineGrid.SpiralLines =
			append(p.SpiralConstructionLineGrid.SpiralLines, sl)
	}
}

func (p *Parameter) ComputeSpiralConstructionCircleGrid() {
	p.SpiralConstructionCircleGrid.SpiralCircles =
		p.SpiralConstructionCircleGrid.SpiralCircles[:0]

	for idx, c := range p.ConstructionCircleGrid.Circles {

		x_r, y_r := p.convertToSpiralCoords(c.CenterX, c.CenterY)

		sc := new(SpiralCircle)
		sc.Stroke = GenerateColor(idx % len(colors))
		sc.Stroke = gongsvg_models.Black.ToString()
		sc.Stroke = c.Stroke
		sc.BespopkeRadius = c.BespopkeRadius
		sc.HasBespokeRadius = c.HasBespokeRadius
		sc.StrokeOpacity = 0.5
		sc.StrokeWidth = 2

		sc.CenterX = x_r
		sc.CenterY = y_r

		p.SpiralConstructionCircleGrid.SpiralCircles = append(p.SpiralConstructionCircleGrid.SpiralCircles, sc)
	}
}

// fetch the control points of the the first bezier curve and
// convert them to spiral coordinates
func (p *Parameter) ComputeSpiralBezierSeed() {

	b := p.GrowthCurveSegment
	sca := p.SpiralConstructionLine
	sb := p.SpiralBezierSeed

	sb.StartX, sb.StartY =
		p.convertToSpiralCoords(b.StartX, b.StartY)

	sb.EndX, sb.EndY =
		p.convertToSpiralCoords(b.EndX, b.EndY)

	// Calculate the relative position of the end point
	dx := sca.EndX - sca.StartX
	dy := sca.EndY - sca.StartY

	// Apply the rotation transformation (90 degrees counterclockwise)
	newEndX := -dy + sca.StartX
	newEndY := dx + sca.StartY
	sb.ControlPointStartX, sb.ControlPointStartY = newEndX, newEndY

	sb.ControlPointEndX, sb.ControlPointEndY =
		p.convertToSpiralCoords(b.ControlPointEndX, b.ControlPointEndY)
}

func (p *Parameter) ComputeSpiralBezierGrid() {

	p.SpiralBezierGrid.SpiralBeziers = p.SpiralBezierGrid.SpiralBeziers[:0]
	// for _, b := range p.GrowthCurve.Beziers {
	// 	sb := new(SpiralBezier)

	// 	p.ConvertBeizerToSpiralBezier(b, sb)
	// 	sb.Stroke = b.Stroke
	// 	sb.StrokeWidth = b.StrokeWidth
	// 	sb.StrokeOpacity = b.StrokeOpacity
	// 	p.SpiralBezierGrid.SpiralBeziers = append(p.SpiralBezierGrid.SpiralBeziers, sb)
	// }
}