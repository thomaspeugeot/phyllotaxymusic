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

	// Angle is the angle in degree of the diamond at the origin 0,0
	Angle float64

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

func (parameter *Parameter) ComputeInitialRhombus() {

}

func (p *Parameter) ComputeInitialRhombusGrid(stage *StageStruct) {

	// remove the attached rhombus
	for _, r := range p.InitialRhombusGrid.Rhombuses {
		r.Unstage(stage)
	}

	g := p.InitialRhombusGrid

	angleRad := g.Reference.Angle * math.Pi / 180
	_ = angleRad
	insideAngleRad := g.Reference.InsideAngle * math.Pi / 180
	sinInsideAngle := math.Sin(insideAngleRad / 2)
	cosInsideAngle := math.Cos(insideAngleRad / 2)

	for i := range g.N {
		r := new(Rhombus).Stage(stage)
		*r = *g.Reference
		g.Rhombuses = append(g.Rhombuses, r)
		r.Name = fmt.Sprintf("%d", i)
		r.CenterX += float64(i) * r.SideLength * cosInsideAngle
		r.CenterY += float64(i) * r.SideLength * sinInsideAngle
	}
}
