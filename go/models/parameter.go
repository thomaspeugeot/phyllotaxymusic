package models

import (
	"log"
)

type Parameter struct {
	Name string

	BackendColor string
	MinuteColor  string
	HourColor    string

	N int
	M int
	Z int // number of rhombus

	// how many circle to go around for the front curve
	// the front curve goes from one circle to the nearest
	ShiftToNearestCircle int

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

	GrowthCurveSeed *Bezier
	GrowthCurve     *BezierGrid

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

	SpiralRhombusGridSeed *SpiralRhombus
	SpiralRhombusGrid     *SpiralRhombusGrid
	SpiralCircleSeed      *SpiralCircle
	SpiralCircleGrid      *SpiralCircleGrid // only n+m circles
	SpiralCircleFullGrid  *SpiralCircleGrid // all, with construction

	SpiralConstructionOuterLineSeed *SpiralLine
	SpiralConstructionInnerLineSeed *SpiralLine
	SpiralConstructionOuterLineGrid *SpiralLineGrid
	SpiralConstructionInnerLineGrid *SpiralLineGrid
	SpiralConstructionCircleGrid    *SpiralCircleGrid

	SpiralConstructionOuterLineFullGrid *SpiralLineGrid

	SpiralBezierSeed *SpiralBezier

	// the set of bezier for the first front curve
	SpiralBezierGrid *SpiralBezierGrid

	// the set of bezier for the first front curve
	SpiralBezierFullGrid *SpiralBezierGrid

	// adjusting the strength of the bezier control points
	SpiralBezierStrength float64

	// the result of the transformation of a vertical
	// bezier into a suite of spiral circle
	FrontCurveStack       *FrontCurveStack
	NbInterpolationPoints int

	// taking into accound the rotation to align the
	// curve on the 12 th hour
	HourCurve *FrontCurveStack

	HourHandleRotationAngle float64
	HourMarker              *SpiralCircle
	HourHandleDiskDistance  float64
	HourHandleRadius        float64

	MinuteCurve *FrontCurveStack

	MinuteHandleRotationAngle float64
	MinuteMarker              *SpiralCircle
	MinuteHandleDiskDistance  float64
	MinuteHandleRadius        float64

	MinuteOffset float64

	BackendCurve *FrontCurveStack

	BackendHandleRotationAngle float64
	BackendMarker              *SpiralCircle
	BackendHandleDiskDistance  float64
	BackendHandleRadius        float64

	BackendOffset float64

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
	SpiralOrigin   *SpiralOrigin

	// drawing the spiral
	SpiralOriginX     float64
	SpiralOriginY     float64
	OriginCrossWidth  float64
	SpiralRadiusRatio float64

	ShowSpiralBezierConstruct bool
	ShowInterpolationPoints   bool

	// number of "minimal" notes to the shift
	ActualNotesTemporalShift int
}

func (parameter *Parameter) OnAfterUpdate(stage *StageStruct, stagedParameter, backRepoParameter *Parameter) {

	log.Println("Diagram, OnAfterUpdate", parameter.Name)
	parameter.Impl.OnUpdated(backRepoParameter)

}

type ParameterImplInterface interface {
	OnUpdated(updatedDiagram *Parameter)
}
