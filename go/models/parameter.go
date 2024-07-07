package models

import (
	"fmt"
	"log"
	"math"
)

type Parameter struct {
	Name string
	N    int
	M    int

	// InsideAngle is the angle in degree of the diamond at the origin 0,0
	InsideAngle float64

	DiamondSideLenght float64

	Impl ParameterImplInterface

	InitialRhombus     *Rhombus
	InitialRhombusGrid *RhombusGrid

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

func (p *Parameter) ComputeInitialRhombus() {
	p.InitialRhombus.InsideAngle = p.InsideAngle
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

	for i := range g.N {
		for j := range g.M {
			r := new(Rhombus).Stage(stage)
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
