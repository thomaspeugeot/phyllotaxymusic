package models

import "log"

type Parameter struct {
	Name string
	N    int
	M    int

	// DiamondAngle is the angle in degree of the diamond at the origin 0,0
	DiamondAngle float64

	OriginX           float64
	OriginY           float64
	DiamondSideLenght float64

	Impl DiagramImplInterface
}

func (diagram *Parameter) OnAfterUpdate(stage *StageStruct, _, frontDiagram *Parameter) {

	log.Println("Diagram, OnAfterUpdate", diagram.Name)
	diagram.Impl.OnDiagramUpdated(diagram)

}

type DiagramImplInterface interface {
	OnDiagramUpdated(updatedDiagram *Parameter)
}
