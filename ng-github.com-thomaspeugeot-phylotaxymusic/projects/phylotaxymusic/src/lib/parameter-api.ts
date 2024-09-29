// insertion point for imports
import { RhombusAPI } from './rhombus-api'
import { CircleAPI } from './circle-api'
import { RhombusGridAPI } from './rhombusgrid-api'
import { CircleGridAPI } from './circlegrid-api'
import { AxisAPI } from './axis-api'
import { AxisGridAPI } from './axisgrid-api'
import { BezierAPI } from './bezier-api'
import { BezierGridAPI } from './beziergrid-api'
import { BezierGridStackAPI } from './beziergridstack-api'
import { SpiralRhombusAPI } from './spiralrhombus-api'
import { SpiralRhombusGridAPI } from './spiralrhombusgrid-api'
import { SpiralCircleAPI } from './spiralcircle-api'
import { SpiralCircleGridAPI } from './spiralcirclegrid-api'
import { SpiralAxisAPI } from './spiralaxis-api'
import { KeyAPI } from './key-api'
import { NoteInfoAPI } from './noteinfo-api'
import { HorizontalAxisAPI } from './horizontalaxis-api'
import { VerticalAxisAPI } from './verticalaxis-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ParameterAPI {

	static GONGSTRUCT_NAME = "Parameter"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	N: number = 0
	M: number = 0
	Z: number = 0
	InsideAngle: number = 0
	SideLength: number = 0
	StackWidth: number = 0
	NbShitRight: number = 0
	StackHeight: number = 0
	BezierControlLengthRatio: number = 0
	FkeySizeRatio: number = 0
	FkeyOriginRelativeX: number = 0
	FkeyOriginRelativeY: number = 0
	PitchHeight: number = 0
	NbPitchLines: number = 0
	MeasureLinesHeightRatio: number = 0
	NbMeasureLines: number = 0
	NbMeasureLinesPerCurve: number = 0
	FirstVoiceShiftX: number = 0
	FirstVoiceShiftY: number = 0
	PitchDifference: number = 0
	Speed: number = 0
	Level: number = 0
	IsMinor: boolean = false
	OriginX: number = 0
	OriginY: number = 0
	SpiralOriginX: number = 0
	SpiralOriginY: number = 0
	SpiralInitialRadius: number = 0

	// insertion point for other decls

	ParameterPointersEncoding: ParameterPointersEncoding = new ParameterPointersEncoding
}

export class ParameterPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	InitialRhombusID: NullInt64 = new NullInt64 // if pointer is null, InitialRhombus.ID = 0

	InitialCircleID: NullInt64 = new NullInt64 // if pointer is null, InitialCircle.ID = 0

	InitialRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, InitialRhombusGrid.ID = 0

	InitialCircleGridID: NullInt64 = new NullInt64 // if pointer is null, InitialCircleGrid.ID = 0

	InitialAxisID: NullInt64 = new NullInt64 // if pointer is null, InitialAxis.ID = 0

	RotatedAxisID: NullInt64 = new NullInt64 // if pointer is null, RotatedAxis.ID = 0

	RotatedRhombusID: NullInt64 = new NullInt64 // if pointer is null, RotatedRhombus.ID = 0

	RotatedRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, RotatedRhombusGrid.ID = 0

	RotatedCircleGridID: NullInt64 = new NullInt64 // if pointer is null, RotatedCircleGrid.ID = 0

	NextRhombusID: NullInt64 = new NullInt64 // if pointer is null, NextRhombus.ID = 0

	NextCircleID: NullInt64 = new NullInt64 // if pointer is null, NextCircle.ID = 0

	GrowingRhombusGridSeedID: NullInt64 = new NullInt64 // if pointer is null, GrowingRhombusGridSeed.ID = 0

	GrowingRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, GrowingRhombusGrid.ID = 0

	GrowingCircleGridSeedID: NullInt64 = new NullInt64 // if pointer is null, GrowingCircleGridSeed.ID = 0

	GrowingCircleGridID: NullInt64 = new NullInt64 // if pointer is null, GrowingCircleGrid.ID = 0

	GrowingCircleGridLeftSeedID: NullInt64 = new NullInt64 // if pointer is null, GrowingCircleGridLeftSeed.ID = 0

	GrowingCircleGridLeftID: NullInt64 = new NullInt64 // if pointer is null, GrowingCircleGridLeft.ID = 0

	ConstructionAxisID: NullInt64 = new NullInt64 // if pointer is null, ConstructionAxis.ID = 0

	ConstructionAxisGridID: NullInt64 = new NullInt64 // if pointer is null, ConstructionAxisGrid.ID = 0

	ConstructionCircleID: NullInt64 = new NullInt64 // if pointer is null, ConstructionCircle.ID = 0

	ConstructionCircleGridID: NullInt64 = new NullInt64 // if pointer is null, ConstructionCircleGrid.ID = 0

	GrowthCurveSegmentID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveSegment.ID = 0

	GrowthCurveID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurve.ID = 0

	GrowthCurveShiftedRightSeedID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveShiftedRightSeed.ID = 0

	GrowthCurveShiftedRightID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveShiftedRight.ID = 0

	GrowthCurveNextSeedID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveNextSeed.ID = 0

	GrowthCurveNextID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveNext.ID = 0

	GrowthCurveNextShiftedRightSeedID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveNextShiftedRightSeed.ID = 0

	GrowthCurveNextShiftedRightID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveNextShiftedRight.ID = 0

	GrowthCurveStackID: NullInt64 = new NullInt64 // if pointer is null, GrowthCurveStack.ID = 0

	SpiralRhombusID: NullInt64 = new NullInt64 // if pointer is null, SpiralRhombus.ID = 0

	SpiralRhombusGridID: NullInt64 = new NullInt64 // if pointer is null, SpiralRhombusGrid.ID = 0

	SpiralCircleSeedID: NullInt64 = new NullInt64 // if pointer is null, SpiralCircleSeed.ID = 0

	SpiralCircleGridID: NullInt64 = new NullInt64 // if pointer is null, SpiralCircleGrid.ID = 0

	SpiralConstructionAxisID: NullInt64 = new NullInt64 // if pointer is null, SpiralConstructionAxis.ID = 0

	FkeyID: NullInt64 = new NullInt64 // if pointer is null, Fkey.ID = 0

	PitchLinesID: NullInt64 = new NullInt64 // if pointer is null, PitchLines.ID = 0

	MeasureLinesID: NullInt64 = new NullInt64 // if pointer is null, MeasureLines.ID = 0

	FirstVoiceID: NullInt64 = new NullInt64 // if pointer is null, FirstVoice.ID = 0

	FirstVoiceShiftRigthID: NullInt64 = new NullInt64 // if pointer is null, FirstVoiceShiftRigth.ID = 0

	SecondVoiceID: NullInt64 = new NullInt64 // if pointer is null, SecondVoice.ID = 0

	SecondVoiceShiftedRightID: NullInt64 = new NullInt64 // if pointer is null, SecondVoiceShiftedRight.ID = 0

	FirstVoiceNotesID: NullInt64 = new NullInt64 // if pointer is null, FirstVoiceNotes.ID = 0

	FirstVoiceNotesShiftedRightID: NullInt64 = new NullInt64 // if pointer is null, FirstVoiceNotesShiftedRight.ID = 0

	SecondVoiceNotesID: NullInt64 = new NullInt64 // if pointer is null, SecondVoiceNotes.ID = 0

	SecondVoiceNotesShiftedRightID: NullInt64 = new NullInt64 // if pointer is null, SecondVoiceNotesShiftedRight.ID = 0

	NoteInfos: number[] = []
	HorizontalAxisID: NullInt64 = new NullInt64 // if pointer is null, HorizontalAxis.ID = 0

	VerticalAxisID: NullInt64 = new NullInt64 // if pointer is null, VerticalAxis.ID = 0

}
