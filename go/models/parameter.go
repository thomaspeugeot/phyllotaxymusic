package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"
	cursor "github.com/fullstack-lang/gong/lib/cursor/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	slider "github.com/fullstack-lang/gong/lib/slider/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tone "github.com/fullstack-lang/gong/lib/tone/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Parameter struct {
	Name string

	BackendColor string
	MinuteColor  string
	HourColor    string

	N int
	M int
	Z int // number of rhombus

	// InsideAngle is the angle in degree of the diamond at the origin 0,0
	InsideAngle float64

	// how many circle to go around for the front curve
	// the front curve goes from one circle to the nearest
	ShiftToNearestCircle int

	SideLength float64

	//
	// Shapes
	//
	Shapes []ShapeInterface

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

	// the score
	Fkey                *Key
	FkeySizeRatio       float64
	FkeyOriginRelativeX float64
	FkeyOriginRelativeY float64

	PitchLines   *AxisGrid
	PitchHeight  float64
	NbPitchLines int

	BeatLines            *AxisGrid
	BeatLinesHeightRatio float64
	NbBeatLines          int
	NbOfBeatsInTheme     int

	// Composing
	FirstVoice             *BezierGrid
	FirstVoiceShiftedRigth *BezierGrid
	FirstVoiceShiftX       float64
	FirstVoiceShiftY       float64

	SecondVoice             *BezierGrid
	SecondVoiceShiftedRight *BezierGrid
	PitchDifference         int
	BeatsPerSecond          float64
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

	// ThemeBinaryEncoding is the encoding of the theme
	// on a 64 bits integer.
	//
	// 0 means no notes is played
	// 1 means the first note is played
	// 2 means the second note is played
	// 3 means the first and second note are played
	// etc....
	ThemeBinaryEncoding int

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
	ActualBeatsTemporalShift int

	// not persisted fields
	stager *Stager

	cursor *cursor.Cursor

	phyllotaxymusicStage *StageStruct
	svgStage             *svg.StageStruct
	toneStage            *tone.StageStruct
	treeStage            *tree.StageStruct
	cursorStage          *cursor.StageStruct
	loadStage            *load.StageStruct
	sliderStage          *slider.StageStruct
	buttonStage          *button.StageStruct
	splitStage           *split.StageStruct

	Tree *tree.Tree
}

// GetButtonsStage implements models.Target.
func (parameter *Parameter) GetButtonsStage() *button.StageStruct {
	return parameter.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (parameter *Parameter) OnAfterUpdateButton() {
	parameter.UpdatePhyllotaxyStage()
	parameter.GenerateMusicXMLFile()
}

func (parameter *Parameter) SetCursor(cursor *cursor.Cursor) {
	parameter.cursor = cursor
}

func (parameter *Parameter) SetCursorStage(cursorStage *cursor.StageStruct) {
	parameter.cursorStage = cursorStage
}

func (parameter *Parameter) SetLoadStage(loadStage *load.StageStruct) {
	parameter.loadStage = loadStage
}

func (parameter *Parameter) SetSlidersStage(slidersStage *slider.StageStruct) {
	parameter.sliderStage = slidersStage
}

func (parameter *Parameter) SetButtonsStage(buttonsStage *button.StageStruct) {
	parameter.buttonStage = buttonsStage
}

func (parameter *Parameter) SetSplitsStage(splitsStage *split.StageStruct) {
	parameter.splitStage = splitsStage
}

func (parameter *Parameter) SetSvgStage(gongsvgStage *svg.StageStruct) {
	parameter.svgStage = gongsvgStage
}

func (parameter *Parameter) SetToneStage(gongtoneStage *tone.StageStruct) {
	parameter.toneStage = gongtoneStage
}

func (parameter *Parameter) SetPhyllotaxymusicStage(phyllotaxymusicStage *StageStruct) {
	parameter.phyllotaxymusicStage = phyllotaxymusicStage
}
func (parameter *Parameter) CommitPhyllotaxymusicStage() {
	parameter.phyllotaxymusicStage.Commit()
}

// IsNotePlayed checks whether the note at the specified rank is played.
//
// The notes are represented using a 64-bit integer where each bit corresponds
// to a specific note. For example:
// - If ThemeBinaryEncoding = 5 (binary: 101), the 0th and 2nd notes are played.
//
// Parameters:
// - rank (int): The rank of the note to check (0-based index, must be between 0 and 63).
//
// Returns:
// - bool: `true` if the note at the specified rank is played, `false` otherwise.
//
// Notes:
// - If the rank is outside the valid range (0 to 63), the method returns `false`.
func (parameter *Parameter) IsNotePlayed(beatNb int) bool {
	if beatNb < 0 || beatNb > 63 {
		// Rank must be between 0 and 63 for a 64-bit integer
		return false
	}
	// Check if the rank-th note is played
	return parameter.ThemeBinaryEncoding&(1<<beatNb) != 0
}

// ToggleNotePlayed toggles the note at the specified rank.
//
// If the rank is valid (0 <= rank <= 63), it flips the bit at that position
// in parameter.ThemeBinaryEncoding. Otherwise, it does nothing.
func (parameter *Parameter) ToggleNotePlayed(beatNb int) {
	if beatNb < 0 || beatNb > 63 {
		return // Ignore invalid ranks
	}

	// Flip the bit at `rank`
	// log.Println("parameter.ThemeBinaryEncoding, before flip at beat", beatNb, parameter.ThemeBinaryEncoding)
	parameter.ThemeBinaryEncoding ^= 1 << beatNb
	// log.Println("parameter.ThemeBinaryEncoding, after flip at beat", beatNb, parameter.ThemeBinaryEncoding)
}

func (parameter *Parameter) SetGongtreeStage(gongtreeStage *tree.StageStruct) {
	parameter.treeStage = gongtreeStage
}

func (parameter *Parameter) GetGongtreeStage() *tree.StageStruct {
	return parameter.treeStage
}

func (parameter *Parameter) OnAfterUpdateNode() {
	parameter.UpdateAllStages()
}

func (parameter *Parameter) UpdateAllStages() {
	parameter.UpdatePhyllotaxyStage()
	parameter.UpdateAndCommitCursorStage()
	parameter.UpdateAndCommitSVGStage()
	parameter.UpdateAndCommitToneStage()
	parameter.UpdateAndCommitTreeStage()
	parameter.UpdateAndCommitSlidersStage()
	parameter.UpdateAndCommitButtonStage()
	parameter.CommitPhyllotaxymusicStage()
}

func (parameter *Parameter) UpdateAllStagesButSliders() {
	parameter.UpdatePhyllotaxyStage()
	parameter.UpdateAndCommitCursorStage()
	parameter.UpdateAndCommitSVGStage()
	parameter.UpdateAndCommitToneStage()
	parameter.UpdateAndCommitTreeStage()
	parameter.UpdateAndCommitButtonStage()
	parameter.CommitPhyllotaxymusicStage()
}
