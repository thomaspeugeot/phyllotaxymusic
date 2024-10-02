package models

import (
	"log"
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
	SpiralBezierSeed       *SpiralBezier

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
