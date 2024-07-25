package models

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
)

type Parameter struct {
	Name string
	N    int
	M    int
	Z    int // number of rhombus

	// InsideAngle is the angle in degree of the diamond at the origin 0,0
	InsideAngle float64

	SideLength float64

	Impl ParameterImplInterface

	//
	// Shapes
	//
	Shapes []Shape

	InitialRhombus     *Rhombus
	InitialCircle      *Circle
	InitialRhombusGrid *RhombusGrid
	InitialCircleGrid  *CircleGrid
	InitialAxis        *Axis

	RotatedAxis        *Axis
	RotatedRhombus     *Rhombus
	RotatedRhombusGrid *RhombusGrid
	RotatedCircleGrid  *CircleGrid

	NextRhombus *Rhombus
	NextCircle  *Circle

	GrowingRhombusGridSeed *Rhombus
	GrowingRhombusGrid     *RhombusGrid

	GrowingCircleGridSeed *Circle
	GrowingCircleGrid     *CircleGrid

	GrowingCircleGridLeftSeed *Circle
	GrowingCircleGridLeft     *CircleGrid

	// Axis between initial circle and the N+M circle
	// shifted to the left. The midle of the axis will serve
	// for building the growth line. The growth line with
	// interpolate the middle of the construction axis
	// for the next circles.
	ConstructionAxis       *Axis
	ConstructionAxisGrid   *AxisGrid
	ConstructionCircle     *Circle
	ConstructionCircleGrid *CircleGrid

	GrowthCurveSegment *Bezier
	GrowthCurve        *BezierGrid

	GrowthCurveShiftedRightSeed *Bezier
	GrowthCurveShiftedRight     *BezierGrid

	GrowthCurveNextSeed *Bezier
	GrowthCurveNext     *BezierGrid

	GrowthCurveNextShiftedRightSeed *Bezier
	GrowthCurveNextShiftedRight     *BezierGrid

	// ratio of the length of the control vector to the side length
	BezierControlLengthRatio float64

	// for drawing purpose
	OriginX        float64
	OriginY        float64
	HorizontalAxis *HorizontalAxis
	VerticalAxis   *VerticalAxis
}

func (parameter *Parameter) OnAfterUpdate(stage *StageStruct, _, frontDiagram *Parameter) {

	log.Println("Diagram, OnAfterUpdate", parameter.Name)
	parameter.Impl.OnUpdated(parameter)

}

type ParameterImplInterface interface {
	OnUpdated(updatedDiagram *Parameter)
}

func (p *Parameter) ComputeShapes(stage *StageStruct) {

	if p.Z < p.M+p.N+1 {
		p.Z = p.M + p.N + 1
	}

	p.Shapes = p.Shapes[:0]

	p.Shapes = append(p.Shapes, p.HorizontalAxis)
	p.Shapes = append(p.Shapes, p.VerticalAxis)

	p.ComputeInitialRhombus()
	p.Shapes = append(p.Shapes, p.InitialRhombus)

	p.ComputeInitialCircle()
	p.Shapes = append(p.Shapes, p.InitialCircle)

	p.ComputeInitialRhombusGrid(stage)
	p.Shapes = append(p.Shapes, p.InitialRhombusGrid)

	p.ComputeInitialCircleGrid(stage)
	p.Shapes = append(p.Shapes, p.InitialCircleGrid)

	p.ComputeInitialAxis()
	p.Shapes = append(p.Shapes, p.InitialAxis)

	p.computeRotatedAxis()
	p.Shapes = append(p.Shapes, p.RotatedAxis)

	p.computeRotatedRhombus()
	p.Shapes = append(p.Shapes, p.RotatedRhombus)

	p.computeRotatedRhombusGrid(stage)
	p.Shapes = append(p.Shapes, p.RotatedRhombusGrid)

	p.computeRotatedCircleGrid(stage)
	p.Shapes = append(p.Shapes, p.RotatedCircleGrid)

	p.ComputeNextRhombus()
	p.Shapes = append(p.Shapes, p.NextRhombus)

	p.ComputeNextCircle()
	p.Shapes = append(p.Shapes, p.NextCircle)

	p.ComputeGrowingRhombusGrid()
	p.Shapes = append(p.Shapes, p.GrowingRhombusGrid)

	p.ComputeGrowingCircleGrid()
	p.Shapes = append(p.Shapes, p.GrowingCircleGrid)

	p.ComputeGrowingCircleGridLeft()
	p.Shapes = append(p.Shapes, p.GrowingCircleGridLeft)

	p.computeConstructionAxis()
	p.Shapes = append(p.Shapes, p.ConstructionAxis)

	p.computeConstructionCircle()
	p.Shapes = append(p.Shapes, p.ConstructionCircle)

	p.computeConstructionAxisGrid()
	p.Shapes = append(p.Shapes, p.ConstructionAxisGrid)

	p.computeConstructionCircleGrid()
	p.Shapes = append(p.Shapes, p.ConstructionCircleGrid)

	p.ComputeGrowthCurveSegment()
	p.Shapes = append(p.Shapes, p.GrowthCurveSegment)

	p.ComputeGrowthCurve()
	p.Shapes = append(p.Shapes, p.GrowthCurve)

	p.ComputeGrowthCurveShiftedRight()
	p.Shapes = append(p.Shapes, p.GrowthCurveShiftedRight)

	p.ComputeGrowthCurveNext()
	p.Shapes = append(p.Shapes, p.GrowthCurveNext)

	p.ComputeGrowthCurveNextShiftedRight()
	p.Shapes = append(p.Shapes, p.GrowthCurveNextShiftedRight)
}

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

	angleRad := g.Reference.Angle * math.Pi / 180
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

	p.InitialAxis.Angle = math.Atan2(y, x) * 180 / math.Pi
	p.InitialAxis.Length = math.Sqrt(x*x + y*y)
}

func (p *Parameter) computeRotatedAxis() {
	p.RotatedAxis.Length = p.InitialAxis.Length
	p.RotatedAxis.Angle = 0
}

func (p *Parameter) computeRotatedRhombus() {
	p.RotatedRhombus.SideLength = p.InitialRhombus.SideLength
	p.RotatedRhombus.InsideAngle = p.InitialRhombus.InsideAngle
	p.RotatedRhombus.Angle = -p.InitialAxis.Angle
}

func (p *Parameter) computeRotatedRhombusGrid(stage *StageStruct) {
	g := p.RotatedRhombusGrid

	// remove the attached rhombus
	for _, r := range g.Rhombuses {
		r.Unstage(stage)
	}
	g.Rhombuses = g.Rhombuses[:0]

	angleRad := -p.InitialAxis.Angle * math.Pi / 180
	cosAngle := math.Cos(angleRad)
	sinAngle := math.Sin(angleRad)

	for _, _r := range p.InitialRhombusGrid.Rhombuses {
		r := new(Rhombus) // .Stage(stage)
		*r = *_r
		r.SideLength = _r.SideLength
		r.InsideAngle = _r.InsideAngle
		r.Angle = -p.InitialAxis.Angle

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

	angleRad := -p.InitialAxis.Angle * math.Pi / 180
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
	p.NextRhombus.Angle = p.RotatedRhombus.Angle
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
	p.GrowingRhombusGridSeed.Angle = p.RotatedRhombus.Angle
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
	log.Println("z", p.Z)
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
	log.Println("z", p.Z)
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
	log.Println("z", p.Z)
}

func (p *Parameter) computeConstructionAxis() {

	// get the N+M circles
	circleNPlusM := p.GrowingCircleGrid.Circles[p.M+p.N]
	x := circleNPlusM.CenterX - p.RotatedAxis.Length
	y := circleNPlusM.CenterY

	p.ConstructionAxis.Length = math.Sqrt(x*x + y*y)
	p.ConstructionAxis.Angle = math.Atan2(y, x) * 180 / math.Pi
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

func (p *Parameter) computeBezier(b *Bezier, startCircle, endCircle *Circle) {
	b.StartX = startCircle.CenterX
	b.StartY = startCircle.CenterY

	b.EndX = endCircle.CenterX
	b.EndY = endCircle.CenterY

	angleRad := p.ConstructionAxis.Angle*math.Pi/180 - math.Pi/2.0

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

func (p *Parameter) ComputeGrowthCurveShiftedRight() {

	g := p.GrowthCurveShiftedRight
	g.Beziers = g.Beziers[:0]

	for _, b := range p.GrowthCurve.Beziers {
		_b := new(Bezier)
		*_b = *p.GrowthCurveShiftedRightSeed
		g.Beziers = append(g.Beziers, _b)

		_b.move(b, p.RotatedAxis.Length, 0)
	}
}

func (p *Parameter) ComputeGrowthCurveNext() {

	g := p.GrowthCurveNext
	g.Beziers = g.Beziers[:0]

	for _, b := range p.GrowthCurve.Beziers {
		_b := new(Bezier)
		*_b = *p.GrowthCurveNextSeed
		g.Beziers = append(g.Beziers, _b)

		_b.move(b, p.NextCircle.CenterX, p.NextCircle.CenterY)

	}
}

func (p *Parameter) ComputeGrowthCurveNextShiftedRight() {

	g := p.GrowthCurveNextShiftedRight
	g.Beziers = g.Beziers[:0]

	for _, b := range p.GrowthCurve.Beziers {
		_b := new(Bezier)
		*_b = *p.GrowthCurveNextShiftedRightSeed
		g.Beziers = append(g.Beziers, _b)

		_b.move(b, p.NextCircle.CenterX+p.RotatedAxis.Length, p.NextCircle.CenterY)

	}
}
