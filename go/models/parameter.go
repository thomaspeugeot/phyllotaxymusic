package models

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
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

	GrowthCurveStack *BezierGridStack
	StackWidth       int
	NbShitRight      int
	StackHeight      int

	// ratio of the length of the control vector to the side length
	BezierControlLengthRatio float64

	SpiralRhombus          *SpiralRhombus
	SpiralRhombusGrid      *SpiralRhombusGrid
	SpiralCircleSeed       *SpiralCircle
	SpiralCircleGrid       *SpiralCircleGrid
	SpiralConstructionAxis *SpiralAxis
	SpiralAxisGrid         *SpiralAxisGrid

	// the score
	Fkey                *Key
	FkeySizeRatio       float64
	FkeyOriginRelativeX float64
	FkeyOriginRelativeY float64

	PitchLines   *AxisGrid
	PitchHeight  float64
	NbPitchLines int

	MeasureLines            *AxisGrid
	MeasureLinesHeightRatio float64
	NbMeasureLines          int
	NbMeasureLinesPerCurve  int

	// Composing
	FirstVoice           *BezierGrid
	FirstVoiceShiftRigth *BezierGrid
	FirstVoiceShiftX     float64
	FirstVoiceShiftY     float64

	SecondVoice             *BezierGrid
	SecondVoiceShiftedRight *BezierGrid
	PitchDifference         int
	Speed                   float64
	Level                   float64

	// interpolating notes
	FirstVoiceNotes              *CircleGrid
	FirstVoiceNotesShiftedRight  *CircleGrid
	SecondVoiceNotes             *CircleGrid
	SecondVoiceNotesShiftedRight *CircleGrid

	//
	// scale of preference
	//
	IsMinor bool

	//
	// Information per note
	//
	NoteInfos []*NoteInfo

	// for drawing purpose
	OriginX        float64
	OriginY        float64
	HorizontalAxis *HorizontalAxis
	VerticalAxis   *VerticalAxis

	// drawing the spiral
	SpiralOriginX       float64
	SpiralOriginY       float64
	SpiralInitialRadius float64
}

func (parameter *Parameter) OnAfterUpdate(stage *StageStruct, stagedParameter, backRepoParameter *Parameter) {

	log.Println("Diagram, OnAfterUpdate", parameter.Name)
	parameter.Impl.OnUpdated(backRepoParameter)

}

type ParameterImplInterface interface {
	OnUpdated(updatedDiagram *Parameter)
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

	r := p.SpiralRhombus.Rhombus
	x_s, y_s := r.getCoordinates()
	x_r, y_r := p.convertToCircleSpaceCoords(x_s, y_s)

	p.SpiralCircleSeed.CenterX = x_r[2]
	p.SpiralCircleSeed.CenterY = y_r[2]
}

func (p *Parameter) ComputeSpiralCircleGrid() {

	p.SpiralCircleGrid.SpiralCircles = p.SpiralCircleGrid.SpiralCircles[:0]

	for idx, r := range p.SpiralRhombusGrid.RhombusGrid.Rhombuses {

		x_s, y_s := r.getCoordinates()
		x_r, y_r := p.convertToCircleSpaceCoords(x_s, y_s)

		sc := new(SpiralCircle)
		sc.Stroke = GenerateColor(idx % len(colors))
		sc.Stroke = gongsvg_models.Black.ToString()
		sc.StrokeOpacity = 0.5
		sc.StrokeWidth = 2

		sc.CenterX = (x_r[0] + x_r[2]) / 2.0
		sc.CenterY = (y_r[0] + y_r[2]) / 2.0

		sc.CenterX = x_r[2]
		sc.CenterY = y_r[2]

		p.SpiralCircleGrid.SpiralCircles = append(p.SpiralCircleGrid.SpiralCircles, sc)
	}
}

func (p *Parameter) computeSpiralConstructionAxis() {

	r := p.SpiralRhombus.Rhombus

	x_s, y_s := r.getCoordinates()
	x_r, y_r := p.convertToCircleSpaceCoords(x_s, y_s)

	p.SpiralConstructionAxis.CenterX = x_r[2]
	p.SpiralConstructionAxis.CenterY = y_r[2]

	x := x_r[0] - x_r[2]
	y := y_r[0] - y_r[2]

	p.SpiralConstructionAxis.Length = math.Sqrt(x*x + y*y)
	p.SpiralConstructionAxis.Angle = math.Atan2(y, x) * 180 / math.Pi
}

func (p *Parameter) computeSpiralConstructionAxisGrid() {

	for i := range p.Z {
		spiralAxis := new(SpiralAxis)
		spiralAxis.Name = fmt.Sprintf("Spiral Axis %d", i)
		spiralAxis.Stroke = GenerateColor(i)
		spiralAxis.StrokeWidth = 1
		spiralAxis.StrokeOpacity = 1

		r := p.SpiralRhombusGrid.RhombusGrid.Rhombuses[i]

		x_s, y_s := r.getCoordinates()
		x_r, y_r := p.convertToCircleSpaceCoords(x_s, y_s)

		spiralAxis.CenterX = x_r[2]
		spiralAxis.CenterY = y_r[2]

		x := x_r[0] - x_r[2]
		y := y_r[0] - y_r[2]

		spiralAxis.Length = math.Sqrt(x*x + y*y)
		spiralAxis.Angle = math.Atan2(y, x) * 180 / math.Pi

		p.SpiralAxisGrid.SpiralAxises = append(p.SpiralAxisGrid.SpiralAxises, spiralAxis)
	}
}
