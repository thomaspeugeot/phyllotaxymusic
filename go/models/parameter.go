package models

import "log"

type Parameter struct {
	Name string
	N    int
	M    int

	// Angle is the angle in degree of the diamond at the origin 0,0
	Angle float64

	DiamondSideLenght float64

	Impl ParameterImplInterface

	InitialRhombus *Rhombus

	// for drawing purpose
	IsHorizontalAxisDisplayed bool
	AxisHandleBorderLength    float64
	OriginX                   float64
	OriginY                   float64
	HorizontalAxis_Length     float64
	VerticalAxis_Length       float64

	Axis_StrokeWidth float64
}

func (parameter *Parameter) OnAfterUpdate(stage *StageStruct, _, frontDiagram *Parameter) {

	log.Println("Diagram, OnAfterUpdate", parameter.Name)
	parameter.Impl.OnUpdated(parameter)

}

type ParameterImplInterface interface {
	OnUpdated(updatedDiagram *Parameter)
}

func (parameter *Parameter) GenerateInitialRhombus() {

}
