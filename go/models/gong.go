// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const ProbeTreeSidebarSuffix = "-sidebar"
const ProbeTableSuffix = "-table"
const ProbeFormSuffix = "-form"
const ProbeSplitSuffix = "-probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Axiss           map[*Axis]any
	Axiss_mapString map[string]*Axis

	// insertion point for slice of pointers maps
	OnAfterAxisCreateCallback OnAfterCreateInterface[Axis]
	OnAfterAxisUpdateCallback OnAfterUpdateInterface[Axis]
	OnAfterAxisDeleteCallback OnAfterDeleteInterface[Axis]
	OnAfterAxisReadCallback   OnAfterReadInterface[Axis]

	AxisGrids           map[*AxisGrid]any
	AxisGrids_mapString map[string]*AxisGrid

	// insertion point for slice of pointers maps
	AxisGrid_Axiss_reverseMap map[*Axis]*AxisGrid

	OnAfterAxisGridCreateCallback OnAfterCreateInterface[AxisGrid]
	OnAfterAxisGridUpdateCallback OnAfterUpdateInterface[AxisGrid]
	OnAfterAxisGridDeleteCallback OnAfterDeleteInterface[AxisGrid]
	OnAfterAxisGridReadCallback   OnAfterReadInterface[AxisGrid]

	Beziers           map[*Bezier]any
	Beziers_mapString map[string]*Bezier

	// insertion point for slice of pointers maps
	OnAfterBezierCreateCallback OnAfterCreateInterface[Bezier]
	OnAfterBezierUpdateCallback OnAfterUpdateInterface[Bezier]
	OnAfterBezierDeleteCallback OnAfterDeleteInterface[Bezier]
	OnAfterBezierReadCallback   OnAfterReadInterface[Bezier]

	BezierGrids           map[*BezierGrid]any
	BezierGrids_mapString map[string]*BezierGrid

	// insertion point for slice of pointers maps
	BezierGrid_Beziers_reverseMap map[*Bezier]*BezierGrid

	OnAfterBezierGridCreateCallback OnAfterCreateInterface[BezierGrid]
	OnAfterBezierGridUpdateCallback OnAfterUpdateInterface[BezierGrid]
	OnAfterBezierGridDeleteCallback OnAfterDeleteInterface[BezierGrid]
	OnAfterBezierGridReadCallback   OnAfterReadInterface[BezierGrid]

	BezierGridStacks           map[*BezierGridStack]any
	BezierGridStacks_mapString map[string]*BezierGridStack

	// insertion point for slice of pointers maps
	BezierGridStack_BezierGrids_reverseMap map[*BezierGrid]*BezierGridStack

	OnAfterBezierGridStackCreateCallback OnAfterCreateInterface[BezierGridStack]
	OnAfterBezierGridStackUpdateCallback OnAfterUpdateInterface[BezierGridStack]
	OnAfterBezierGridStackDeleteCallback OnAfterDeleteInterface[BezierGridStack]
	OnAfterBezierGridStackReadCallback   OnAfterReadInterface[BezierGridStack]

	Circles           map[*Circle]any
	Circles_mapString map[string]*Circle

	// insertion point for slice of pointers maps
	OnAfterCircleCreateCallback OnAfterCreateInterface[Circle]
	OnAfterCircleUpdateCallback OnAfterUpdateInterface[Circle]
	OnAfterCircleDeleteCallback OnAfterDeleteInterface[Circle]
	OnAfterCircleReadCallback   OnAfterReadInterface[Circle]

	CircleGrids           map[*CircleGrid]any
	CircleGrids_mapString map[string]*CircleGrid

	// insertion point for slice of pointers maps
	CircleGrid_Circles_reverseMap map[*Circle]*CircleGrid

	OnAfterCircleGridCreateCallback OnAfterCreateInterface[CircleGrid]
	OnAfterCircleGridUpdateCallback OnAfterUpdateInterface[CircleGrid]
	OnAfterCircleGridDeleteCallback OnAfterDeleteInterface[CircleGrid]
	OnAfterCircleGridReadCallback   OnAfterReadInterface[CircleGrid]

	ExportToMusicxmls           map[*ExportToMusicxml]any
	ExportToMusicxmls_mapString map[string]*ExportToMusicxml

	// insertion point for slice of pointers maps
	OnAfterExportToMusicxmlCreateCallback OnAfterCreateInterface[ExportToMusicxml]
	OnAfterExportToMusicxmlUpdateCallback OnAfterUpdateInterface[ExportToMusicxml]
	OnAfterExportToMusicxmlDeleteCallback OnAfterDeleteInterface[ExportToMusicxml]
	OnAfterExportToMusicxmlReadCallback   OnAfterReadInterface[ExportToMusicxml]

	FrontCurves           map[*FrontCurve]any
	FrontCurves_mapString map[string]*FrontCurve

	// insertion point for slice of pointers maps
	OnAfterFrontCurveCreateCallback OnAfterCreateInterface[FrontCurve]
	OnAfterFrontCurveUpdateCallback OnAfterUpdateInterface[FrontCurve]
	OnAfterFrontCurveDeleteCallback OnAfterDeleteInterface[FrontCurve]
	OnAfterFrontCurveReadCallback   OnAfterReadInterface[FrontCurve]

	FrontCurveStacks           map[*FrontCurveStack]any
	FrontCurveStacks_mapString map[string]*FrontCurveStack

	// insertion point for slice of pointers maps
	FrontCurveStack_FrontCurves_reverseMap map[*FrontCurve]*FrontCurveStack

	FrontCurveStack_SpiralCircles_reverseMap map[*SpiralCircle]*FrontCurveStack

	OnAfterFrontCurveStackCreateCallback OnAfterCreateInterface[FrontCurveStack]
	OnAfterFrontCurveStackUpdateCallback OnAfterUpdateInterface[FrontCurveStack]
	OnAfterFrontCurveStackDeleteCallback OnAfterDeleteInterface[FrontCurveStack]
	OnAfterFrontCurveStackReadCallback   OnAfterReadInterface[FrontCurveStack]

	HorizontalAxiss           map[*HorizontalAxis]any
	HorizontalAxiss_mapString map[string]*HorizontalAxis

	// insertion point for slice of pointers maps
	OnAfterHorizontalAxisCreateCallback OnAfterCreateInterface[HorizontalAxis]
	OnAfterHorizontalAxisUpdateCallback OnAfterUpdateInterface[HorizontalAxis]
	OnAfterHorizontalAxisDeleteCallback OnAfterDeleteInterface[HorizontalAxis]
	OnAfterHorizontalAxisReadCallback   OnAfterReadInterface[HorizontalAxis]

	Keys           map[*Key]any
	Keys_mapString map[string]*Key

	// insertion point for slice of pointers maps
	OnAfterKeyCreateCallback OnAfterCreateInterface[Key]
	OnAfterKeyUpdateCallback OnAfterUpdateInterface[Key]
	OnAfterKeyDeleteCallback OnAfterDeleteInterface[Key]
	OnAfterKeyReadCallback   OnAfterReadInterface[Key]

	Parameters           map[*Parameter]any
	Parameters_mapString map[string]*Parameter

	// insertion point for slice of pointers maps
	OnAfterParameterCreateCallback OnAfterCreateInterface[Parameter]
	OnAfterParameterUpdateCallback OnAfterUpdateInterface[Parameter]
	OnAfterParameterDeleteCallback OnAfterDeleteInterface[Parameter]
	OnAfterParameterReadCallback   OnAfterReadInterface[Parameter]

	Rhombuss           map[*Rhombus]any
	Rhombuss_mapString map[string]*Rhombus

	// insertion point for slice of pointers maps
	OnAfterRhombusCreateCallback OnAfterCreateInterface[Rhombus]
	OnAfterRhombusUpdateCallback OnAfterUpdateInterface[Rhombus]
	OnAfterRhombusDeleteCallback OnAfterDeleteInterface[Rhombus]
	OnAfterRhombusReadCallback   OnAfterReadInterface[Rhombus]

	RhombusGrids           map[*RhombusGrid]any
	RhombusGrids_mapString map[string]*RhombusGrid

	// insertion point for slice of pointers maps
	RhombusGrid_Rhombuses_reverseMap map[*Rhombus]*RhombusGrid

	OnAfterRhombusGridCreateCallback OnAfterCreateInterface[RhombusGrid]
	OnAfterRhombusGridUpdateCallback OnAfterUpdateInterface[RhombusGrid]
	OnAfterRhombusGridDeleteCallback OnAfterDeleteInterface[RhombusGrid]
	OnAfterRhombusGridReadCallback   OnAfterReadInterface[RhombusGrid]

	ShapeCategorys           map[*ShapeCategory]any
	ShapeCategorys_mapString map[string]*ShapeCategory

	// insertion point for slice of pointers maps
	OnAfterShapeCategoryCreateCallback OnAfterCreateInterface[ShapeCategory]
	OnAfterShapeCategoryUpdateCallback OnAfterUpdateInterface[ShapeCategory]
	OnAfterShapeCategoryDeleteCallback OnAfterDeleteInterface[ShapeCategory]
	OnAfterShapeCategoryReadCallback   OnAfterReadInterface[ShapeCategory]

	SpiralBeziers           map[*SpiralBezier]any
	SpiralBeziers_mapString map[string]*SpiralBezier

	// insertion point for slice of pointers maps
	OnAfterSpiralBezierCreateCallback OnAfterCreateInterface[SpiralBezier]
	OnAfterSpiralBezierUpdateCallback OnAfterUpdateInterface[SpiralBezier]
	OnAfterSpiralBezierDeleteCallback OnAfterDeleteInterface[SpiralBezier]
	OnAfterSpiralBezierReadCallback   OnAfterReadInterface[SpiralBezier]

	SpiralBezierGrids           map[*SpiralBezierGrid]any
	SpiralBezierGrids_mapString map[string]*SpiralBezierGrid

	// insertion point for slice of pointers maps
	SpiralBezierGrid_SpiralBeziers_reverseMap map[*SpiralBezier]*SpiralBezierGrid

	OnAfterSpiralBezierGridCreateCallback OnAfterCreateInterface[SpiralBezierGrid]
	OnAfterSpiralBezierGridUpdateCallback OnAfterUpdateInterface[SpiralBezierGrid]
	OnAfterSpiralBezierGridDeleteCallback OnAfterDeleteInterface[SpiralBezierGrid]
	OnAfterSpiralBezierGridReadCallback   OnAfterReadInterface[SpiralBezierGrid]

	SpiralCircles           map[*SpiralCircle]any
	SpiralCircles_mapString map[string]*SpiralCircle

	// insertion point for slice of pointers maps
	OnAfterSpiralCircleCreateCallback OnAfterCreateInterface[SpiralCircle]
	OnAfterSpiralCircleUpdateCallback OnAfterUpdateInterface[SpiralCircle]
	OnAfterSpiralCircleDeleteCallback OnAfterDeleteInterface[SpiralCircle]
	OnAfterSpiralCircleReadCallback   OnAfterReadInterface[SpiralCircle]

	SpiralCircleGrids           map[*SpiralCircleGrid]any
	SpiralCircleGrids_mapString map[string]*SpiralCircleGrid

	// insertion point for slice of pointers maps
	SpiralCircleGrid_SpiralCircles_reverseMap map[*SpiralCircle]*SpiralCircleGrid

	OnAfterSpiralCircleGridCreateCallback OnAfterCreateInterface[SpiralCircleGrid]
	OnAfterSpiralCircleGridUpdateCallback OnAfterUpdateInterface[SpiralCircleGrid]
	OnAfterSpiralCircleGridDeleteCallback OnAfterDeleteInterface[SpiralCircleGrid]
	OnAfterSpiralCircleGridReadCallback   OnAfterReadInterface[SpiralCircleGrid]

	SpiralLines           map[*SpiralLine]any
	SpiralLines_mapString map[string]*SpiralLine

	// insertion point for slice of pointers maps
	OnAfterSpiralLineCreateCallback OnAfterCreateInterface[SpiralLine]
	OnAfterSpiralLineUpdateCallback OnAfterUpdateInterface[SpiralLine]
	OnAfterSpiralLineDeleteCallback OnAfterDeleteInterface[SpiralLine]
	OnAfterSpiralLineReadCallback   OnAfterReadInterface[SpiralLine]

	SpiralLineGrids           map[*SpiralLineGrid]any
	SpiralLineGrids_mapString map[string]*SpiralLineGrid

	// insertion point for slice of pointers maps
	SpiralLineGrid_SpiralLines_reverseMap map[*SpiralLine]*SpiralLineGrid

	OnAfterSpiralLineGridCreateCallback OnAfterCreateInterface[SpiralLineGrid]
	OnAfterSpiralLineGridUpdateCallback OnAfterUpdateInterface[SpiralLineGrid]
	OnAfterSpiralLineGridDeleteCallback OnAfterDeleteInterface[SpiralLineGrid]
	OnAfterSpiralLineGridReadCallback   OnAfterReadInterface[SpiralLineGrid]

	SpiralOrigins           map[*SpiralOrigin]any
	SpiralOrigins_mapString map[string]*SpiralOrigin

	// insertion point for slice of pointers maps
	OnAfterSpiralOriginCreateCallback OnAfterCreateInterface[SpiralOrigin]
	OnAfterSpiralOriginUpdateCallback OnAfterUpdateInterface[SpiralOrigin]
	OnAfterSpiralOriginDeleteCallback OnAfterDeleteInterface[SpiralOrigin]
	OnAfterSpiralOriginReadCallback   OnAfterReadInterface[SpiralOrigin]

	SpiralRhombuss           map[*SpiralRhombus]any
	SpiralRhombuss_mapString map[string]*SpiralRhombus

	// insertion point for slice of pointers maps
	OnAfterSpiralRhombusCreateCallback OnAfterCreateInterface[SpiralRhombus]
	OnAfterSpiralRhombusUpdateCallback OnAfterUpdateInterface[SpiralRhombus]
	OnAfterSpiralRhombusDeleteCallback OnAfterDeleteInterface[SpiralRhombus]
	OnAfterSpiralRhombusReadCallback   OnAfterReadInterface[SpiralRhombus]

	SpiralRhombusGrids           map[*SpiralRhombusGrid]any
	SpiralRhombusGrids_mapString map[string]*SpiralRhombusGrid

	// insertion point for slice of pointers maps
	SpiralRhombusGrid_SpiralRhombuses_reverseMap map[*SpiralRhombus]*SpiralRhombusGrid

	OnAfterSpiralRhombusGridCreateCallback OnAfterCreateInterface[SpiralRhombusGrid]
	OnAfterSpiralRhombusGridUpdateCallback OnAfterUpdateInterface[SpiralRhombusGrid]
	OnAfterSpiralRhombusGridDeleteCallback OnAfterDeleteInterface[SpiralRhombusGrid]
	OnAfterSpiralRhombusGridReadCallback   OnAfterReadInterface[SpiralRhombusGrid]

	VerticalAxiss           map[*VerticalAxis]any
	VerticalAxiss_mapString map[string]*VerticalAxis

	// insertion point for slice of pointers maps
	OnAfterVerticalAxisCreateCallback OnAfterCreateInterface[VerticalAxis]
	OnAfterVerticalAxisUpdateCallback OnAfterUpdateInterface[VerticalAxis]
	OnAfterVerticalAxisDeleteCallback OnAfterDeleteInterface[VerticalAxis]
	OnAfterVerticalAxisReadCallback   OnAfterReadInterface[VerticalAxis]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	AxisOrder            uint
	AxisMap_Staged_Order map[*Axis]uint

	AxisGridOrder            uint
	AxisGridMap_Staged_Order map[*AxisGrid]uint

	BezierOrder            uint
	BezierMap_Staged_Order map[*Bezier]uint

	BezierGridOrder            uint
	BezierGridMap_Staged_Order map[*BezierGrid]uint

	BezierGridStackOrder            uint
	BezierGridStackMap_Staged_Order map[*BezierGridStack]uint

	CircleOrder            uint
	CircleMap_Staged_Order map[*Circle]uint

	CircleGridOrder            uint
	CircleGridMap_Staged_Order map[*CircleGrid]uint

	ExportToMusicxmlOrder            uint
	ExportToMusicxmlMap_Staged_Order map[*ExportToMusicxml]uint

	FrontCurveOrder            uint
	FrontCurveMap_Staged_Order map[*FrontCurve]uint

	FrontCurveStackOrder            uint
	FrontCurveStackMap_Staged_Order map[*FrontCurveStack]uint

	HorizontalAxisOrder            uint
	HorizontalAxisMap_Staged_Order map[*HorizontalAxis]uint

	KeyOrder            uint
	KeyMap_Staged_Order map[*Key]uint

	ParameterOrder            uint
	ParameterMap_Staged_Order map[*Parameter]uint

	RhombusOrder            uint
	RhombusMap_Staged_Order map[*Rhombus]uint

	RhombusGridOrder            uint
	RhombusGridMap_Staged_Order map[*RhombusGrid]uint

	ShapeCategoryOrder            uint
	ShapeCategoryMap_Staged_Order map[*ShapeCategory]uint

	SpiralBezierOrder            uint
	SpiralBezierMap_Staged_Order map[*SpiralBezier]uint

	SpiralBezierGridOrder            uint
	SpiralBezierGridMap_Staged_Order map[*SpiralBezierGrid]uint

	SpiralCircleOrder            uint
	SpiralCircleMap_Staged_Order map[*SpiralCircle]uint

	SpiralCircleGridOrder            uint
	SpiralCircleGridMap_Staged_Order map[*SpiralCircleGrid]uint

	SpiralLineOrder            uint
	SpiralLineMap_Staged_Order map[*SpiralLine]uint

	SpiralLineGridOrder            uint
	SpiralLineGridMap_Staged_Order map[*SpiralLineGrid]uint

	SpiralOriginOrder            uint
	SpiralOriginMap_Staged_Order map[*SpiralOrigin]uint

	SpiralRhombusOrder            uint
	SpiralRhombusMap_Staged_Order map[*SpiralRhombus]uint

	SpiralRhombusGridOrder            uint
	SpiralRhombusGridMap_Staged_Order map[*SpiralRhombusGrid]uint

	VerticalAxisOrder            uint
	VerticalAxisMap_Staged_Order map[*VerticalAxis]uint

	// end of insertion point
}

func (stage *Stage) GetType() string {
	return "github.com/thomaspeugeot/phyllotaxymusic/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAxis(axis *Axis)
	CheckoutAxis(axis *Axis)
	CommitAxisGrid(axisgrid *AxisGrid)
	CheckoutAxisGrid(axisgrid *AxisGrid)
	CommitBezier(bezier *Bezier)
	CheckoutBezier(bezier *Bezier)
	CommitBezierGrid(beziergrid *BezierGrid)
	CheckoutBezierGrid(beziergrid *BezierGrid)
	CommitBezierGridStack(beziergridstack *BezierGridStack)
	CheckoutBezierGridStack(beziergridstack *BezierGridStack)
	CommitCircle(circle *Circle)
	CheckoutCircle(circle *Circle)
	CommitCircleGrid(circlegrid *CircleGrid)
	CheckoutCircleGrid(circlegrid *CircleGrid)
	CommitExportToMusicxml(exporttomusicxml *ExportToMusicxml)
	CheckoutExportToMusicxml(exporttomusicxml *ExportToMusicxml)
	CommitFrontCurve(frontcurve *FrontCurve)
	CheckoutFrontCurve(frontcurve *FrontCurve)
	CommitFrontCurveStack(frontcurvestack *FrontCurveStack)
	CheckoutFrontCurveStack(frontcurvestack *FrontCurveStack)
	CommitHorizontalAxis(horizontalaxis *HorizontalAxis)
	CheckoutHorizontalAxis(horizontalaxis *HorizontalAxis)
	CommitKey(key *Key)
	CheckoutKey(key *Key)
	CommitParameter(parameter *Parameter)
	CheckoutParameter(parameter *Parameter)
	CommitRhombus(rhombus *Rhombus)
	CheckoutRhombus(rhombus *Rhombus)
	CommitRhombusGrid(rhombusgrid *RhombusGrid)
	CheckoutRhombusGrid(rhombusgrid *RhombusGrid)
	CommitShapeCategory(shapecategory *ShapeCategory)
	CheckoutShapeCategory(shapecategory *ShapeCategory)
	CommitSpiralBezier(spiralbezier *SpiralBezier)
	CheckoutSpiralBezier(spiralbezier *SpiralBezier)
	CommitSpiralBezierGrid(spiralbeziergrid *SpiralBezierGrid)
	CheckoutSpiralBezierGrid(spiralbeziergrid *SpiralBezierGrid)
	CommitSpiralCircle(spiralcircle *SpiralCircle)
	CheckoutSpiralCircle(spiralcircle *SpiralCircle)
	CommitSpiralCircleGrid(spiralcirclegrid *SpiralCircleGrid)
	CheckoutSpiralCircleGrid(spiralcirclegrid *SpiralCircleGrid)
	CommitSpiralLine(spiralline *SpiralLine)
	CheckoutSpiralLine(spiralline *SpiralLine)
	CommitSpiralLineGrid(spirallinegrid *SpiralLineGrid)
	CheckoutSpiralLineGrid(spirallinegrid *SpiralLineGrid)
	CommitSpiralOrigin(spiralorigin *SpiralOrigin)
	CheckoutSpiralOrigin(spiralorigin *SpiralOrigin)
	CommitSpiralRhombus(spiralrhombus *SpiralRhombus)
	CheckoutSpiralRhombus(spiralrhombus *SpiralRhombus)
	CommitSpiralRhombusGrid(spiralrhombusgrid *SpiralRhombusGrid)
	CheckoutSpiralRhombusGrid(spiralrhombusgrid *SpiralRhombusGrid)
	CommitVerticalAxis(verticalaxis *VerticalAxis)
	CheckoutVerticalAxis(verticalaxis *VerticalAxis)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Axiss:           make(map[*Axis]any),
		Axiss_mapString: make(map[string]*Axis),

		AxisGrids:           make(map[*AxisGrid]any),
		AxisGrids_mapString: make(map[string]*AxisGrid),

		Beziers:           make(map[*Bezier]any),
		Beziers_mapString: make(map[string]*Bezier),

		BezierGrids:           make(map[*BezierGrid]any),
		BezierGrids_mapString: make(map[string]*BezierGrid),

		BezierGridStacks:           make(map[*BezierGridStack]any),
		BezierGridStacks_mapString: make(map[string]*BezierGridStack),

		Circles:           make(map[*Circle]any),
		Circles_mapString: make(map[string]*Circle),

		CircleGrids:           make(map[*CircleGrid]any),
		CircleGrids_mapString: make(map[string]*CircleGrid),

		ExportToMusicxmls:           make(map[*ExportToMusicxml]any),
		ExportToMusicxmls_mapString: make(map[string]*ExportToMusicxml),

		FrontCurves:           make(map[*FrontCurve]any),
		FrontCurves_mapString: make(map[string]*FrontCurve),

		FrontCurveStacks:           make(map[*FrontCurveStack]any),
		FrontCurveStacks_mapString: make(map[string]*FrontCurveStack),

		HorizontalAxiss:           make(map[*HorizontalAxis]any),
		HorizontalAxiss_mapString: make(map[string]*HorizontalAxis),

		Keys:           make(map[*Key]any),
		Keys_mapString: make(map[string]*Key),

		Parameters:           make(map[*Parameter]any),
		Parameters_mapString: make(map[string]*Parameter),

		Rhombuss:           make(map[*Rhombus]any),
		Rhombuss_mapString: make(map[string]*Rhombus),

		RhombusGrids:           make(map[*RhombusGrid]any),
		RhombusGrids_mapString: make(map[string]*RhombusGrid),

		ShapeCategorys:           make(map[*ShapeCategory]any),
		ShapeCategorys_mapString: make(map[string]*ShapeCategory),

		SpiralBeziers:           make(map[*SpiralBezier]any),
		SpiralBeziers_mapString: make(map[string]*SpiralBezier),

		SpiralBezierGrids:           make(map[*SpiralBezierGrid]any),
		SpiralBezierGrids_mapString: make(map[string]*SpiralBezierGrid),

		SpiralCircles:           make(map[*SpiralCircle]any),
		SpiralCircles_mapString: make(map[string]*SpiralCircle),

		SpiralCircleGrids:           make(map[*SpiralCircleGrid]any),
		SpiralCircleGrids_mapString: make(map[string]*SpiralCircleGrid),

		SpiralLines:           make(map[*SpiralLine]any),
		SpiralLines_mapString: make(map[string]*SpiralLine),

		SpiralLineGrids:           make(map[*SpiralLineGrid]any),
		SpiralLineGrids_mapString: make(map[string]*SpiralLineGrid),

		SpiralOrigins:           make(map[*SpiralOrigin]any),
		SpiralOrigins_mapString: make(map[string]*SpiralOrigin),

		SpiralRhombuss:           make(map[*SpiralRhombus]any),
		SpiralRhombuss_mapString: make(map[string]*SpiralRhombus),

		SpiralRhombusGrids:           make(map[*SpiralRhombusGrid]any),
		SpiralRhombusGrids_mapString: make(map[string]*SpiralRhombusGrid),

		VerticalAxiss:           make(map[*VerticalAxis]any),
		VerticalAxiss_mapString: make(map[string]*VerticalAxis),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AxisMap_Staged_Order: make(map[*Axis]uint),

		AxisGridMap_Staged_Order: make(map[*AxisGrid]uint),

		BezierMap_Staged_Order: make(map[*Bezier]uint),

		BezierGridMap_Staged_Order: make(map[*BezierGrid]uint),

		BezierGridStackMap_Staged_Order: make(map[*BezierGridStack]uint),

		CircleMap_Staged_Order: make(map[*Circle]uint),

		CircleGridMap_Staged_Order: make(map[*CircleGrid]uint),

		ExportToMusicxmlMap_Staged_Order: make(map[*ExportToMusicxml]uint),

		FrontCurveMap_Staged_Order: make(map[*FrontCurve]uint),

		FrontCurveStackMap_Staged_Order: make(map[*FrontCurveStack]uint),

		HorizontalAxisMap_Staged_Order: make(map[*HorizontalAxis]uint),

		KeyMap_Staged_Order: make(map[*Key]uint),

		ParameterMap_Staged_Order: make(map[*Parameter]uint),

		RhombusMap_Staged_Order: make(map[*Rhombus]uint),

		RhombusGridMap_Staged_Order: make(map[*RhombusGrid]uint),

		ShapeCategoryMap_Staged_Order: make(map[*ShapeCategory]uint),

		SpiralBezierMap_Staged_Order: make(map[*SpiralBezier]uint),

		SpiralBezierGridMap_Staged_Order: make(map[*SpiralBezierGrid]uint),

		SpiralCircleMap_Staged_Order: make(map[*SpiralCircle]uint),

		SpiralCircleGridMap_Staged_Order: make(map[*SpiralCircleGrid]uint),

		SpiralLineMap_Staged_Order: make(map[*SpiralLine]uint),

		SpiralLineGridMap_Staged_Order: make(map[*SpiralLineGrid]uint),

		SpiralOriginMap_Staged_Order: make(map[*SpiralOrigin]uint),

		SpiralRhombusMap_Staged_Order: make(map[*SpiralRhombus]uint),

		SpiralRhombusGridMap_Staged_Order: make(map[*SpiralRhombusGrid]uint),

		VerticalAxisMap_Staged_Order: make(map[*VerticalAxis]uint),

		// end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Axis:
		return stage.AxisMap_Staged_Order[instance]
	case *AxisGrid:
		return stage.AxisGridMap_Staged_Order[instance]
	case *Bezier:
		return stage.BezierMap_Staged_Order[instance]
	case *BezierGrid:
		return stage.BezierGridMap_Staged_Order[instance]
	case *BezierGridStack:
		return stage.BezierGridStackMap_Staged_Order[instance]
	case *Circle:
		return stage.CircleMap_Staged_Order[instance]
	case *CircleGrid:
		return stage.CircleGridMap_Staged_Order[instance]
	case *ExportToMusicxml:
		return stage.ExportToMusicxmlMap_Staged_Order[instance]
	case *FrontCurve:
		return stage.FrontCurveMap_Staged_Order[instance]
	case *FrontCurveStack:
		return stage.FrontCurveStackMap_Staged_Order[instance]
	case *HorizontalAxis:
		return stage.HorizontalAxisMap_Staged_Order[instance]
	case *Key:
		return stage.KeyMap_Staged_Order[instance]
	case *Parameter:
		return stage.ParameterMap_Staged_Order[instance]
	case *Rhombus:
		return stage.RhombusMap_Staged_Order[instance]
	case *RhombusGrid:
		return stage.RhombusGridMap_Staged_Order[instance]
	case *ShapeCategory:
		return stage.ShapeCategoryMap_Staged_Order[instance]
	case *SpiralBezier:
		return stage.SpiralBezierMap_Staged_Order[instance]
	case *SpiralBezierGrid:
		return stage.SpiralBezierGridMap_Staged_Order[instance]
	case *SpiralCircle:
		return stage.SpiralCircleMap_Staged_Order[instance]
	case *SpiralCircleGrid:
		return stage.SpiralCircleGridMap_Staged_Order[instance]
	case *SpiralLine:
		return stage.SpiralLineMap_Staged_Order[instance]
	case *SpiralLineGrid:
		return stage.SpiralLineGridMap_Staged_Order[instance]
	case *SpiralOrigin:
		return stage.SpiralOriginMap_Staged_Order[instance]
	case *SpiralRhombus:
		return stage.SpiralRhombusMap_Staged_Order[instance]
	case *SpiralRhombusGrid:
		return stage.SpiralRhombusGridMap_Staged_Order[instance]
	case *VerticalAxis:
		return stage.VerticalAxisMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Axis"] = len(stage.Axiss)
	stage.Map_GongStructName_InstancesNb["AxisGrid"] = len(stage.AxisGrids)
	stage.Map_GongStructName_InstancesNb["Bezier"] = len(stage.Beziers)
	stage.Map_GongStructName_InstancesNb["BezierGrid"] = len(stage.BezierGrids)
	stage.Map_GongStructName_InstancesNb["BezierGridStack"] = len(stage.BezierGridStacks)
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["CircleGrid"] = len(stage.CircleGrids)
	stage.Map_GongStructName_InstancesNb["ExportToMusicxml"] = len(stage.ExportToMusicxmls)
	stage.Map_GongStructName_InstancesNb["FrontCurve"] = len(stage.FrontCurves)
	stage.Map_GongStructName_InstancesNb["FrontCurveStack"] = len(stage.FrontCurveStacks)
	stage.Map_GongStructName_InstancesNb["HorizontalAxis"] = len(stage.HorizontalAxiss)
	stage.Map_GongStructName_InstancesNb["Key"] = len(stage.Keys)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)
	stage.Map_GongStructName_InstancesNb["Rhombus"] = len(stage.Rhombuss)
	stage.Map_GongStructName_InstancesNb["RhombusGrid"] = len(stage.RhombusGrids)
	stage.Map_GongStructName_InstancesNb["ShapeCategory"] = len(stage.ShapeCategorys)
	stage.Map_GongStructName_InstancesNb["SpiralBezier"] = len(stage.SpiralBeziers)
	stage.Map_GongStructName_InstancesNb["SpiralBezierGrid"] = len(stage.SpiralBezierGrids)
	stage.Map_GongStructName_InstancesNb["SpiralCircle"] = len(stage.SpiralCircles)
	stage.Map_GongStructName_InstancesNb["SpiralCircleGrid"] = len(stage.SpiralCircleGrids)
	stage.Map_GongStructName_InstancesNb["SpiralLine"] = len(stage.SpiralLines)
	stage.Map_GongStructName_InstancesNb["SpiralLineGrid"] = len(stage.SpiralLineGrids)
	stage.Map_GongStructName_InstancesNb["SpiralOrigin"] = len(stage.SpiralOrigins)
	stage.Map_GongStructName_InstancesNb["SpiralRhombus"] = len(stage.SpiralRhombuss)
	stage.Map_GongStructName_InstancesNb["SpiralRhombusGrid"] = len(stage.SpiralRhombusGrids)
	stage.Map_GongStructName_InstancesNb["VerticalAxis"] = len(stage.VerticalAxiss)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Axis"] = len(stage.Axiss)
	stage.Map_GongStructName_InstancesNb["AxisGrid"] = len(stage.AxisGrids)
	stage.Map_GongStructName_InstancesNb["Bezier"] = len(stage.Beziers)
	stage.Map_GongStructName_InstancesNb["BezierGrid"] = len(stage.BezierGrids)
	stage.Map_GongStructName_InstancesNb["BezierGridStack"] = len(stage.BezierGridStacks)
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["CircleGrid"] = len(stage.CircleGrids)
	stage.Map_GongStructName_InstancesNb["ExportToMusicxml"] = len(stage.ExportToMusicxmls)
	stage.Map_GongStructName_InstancesNb["FrontCurve"] = len(stage.FrontCurves)
	stage.Map_GongStructName_InstancesNb["FrontCurveStack"] = len(stage.FrontCurveStacks)
	stage.Map_GongStructName_InstancesNb["HorizontalAxis"] = len(stage.HorizontalAxiss)
	stage.Map_GongStructName_InstancesNb["Key"] = len(stage.Keys)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)
	stage.Map_GongStructName_InstancesNb["Rhombus"] = len(stage.Rhombuss)
	stage.Map_GongStructName_InstancesNb["RhombusGrid"] = len(stage.RhombusGrids)
	stage.Map_GongStructName_InstancesNb["ShapeCategory"] = len(stage.ShapeCategorys)
	stage.Map_GongStructName_InstancesNb["SpiralBezier"] = len(stage.SpiralBeziers)
	stage.Map_GongStructName_InstancesNb["SpiralBezierGrid"] = len(stage.SpiralBezierGrids)
	stage.Map_GongStructName_InstancesNb["SpiralCircle"] = len(stage.SpiralCircles)
	stage.Map_GongStructName_InstancesNb["SpiralCircleGrid"] = len(stage.SpiralCircleGrids)
	stage.Map_GongStructName_InstancesNb["SpiralLine"] = len(stage.SpiralLines)
	stage.Map_GongStructName_InstancesNb["SpiralLineGrid"] = len(stage.SpiralLineGrids)
	stage.Map_GongStructName_InstancesNb["SpiralOrigin"] = len(stage.SpiralOrigins)
	stage.Map_GongStructName_InstancesNb["SpiralRhombus"] = len(stage.SpiralRhombuss)
	stage.Map_GongStructName_InstancesNb["SpiralRhombusGrid"] = len(stage.SpiralRhombusGrids)
	stage.Map_GongStructName_InstancesNb["VerticalAxis"] = len(stage.VerticalAxiss)

}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts axis to the model stage
func (axis *Axis) Stage(stage *Stage) *Axis {

	if _, ok := stage.Axiss[axis]; !ok {
		stage.Axiss[axis] = __member
		stage.AxisMap_Staged_Order[axis] = stage.AxisOrder
		stage.AxisOrder++
	}
	stage.Axiss_mapString[axis.Name] = axis

	return axis
}

// Unstage removes axis off the model stage
func (axis *Axis) Unstage(stage *Stage) *Axis {
	delete(stage.Axiss, axis)
	delete(stage.Axiss_mapString, axis.Name)
	return axis
}

// UnstageVoid removes axis off the model stage
func (axis *Axis) UnstageVoid(stage *Stage) {
	delete(stage.Axiss, axis)
	delete(stage.Axiss_mapString, axis.Name)
}

// commit axis to the back repo (if it is already staged)
func (axis *Axis) Commit(stage *Stage) *Axis {
	if _, ok := stage.Axiss[axis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAxis(axis)
		}
	}
	return axis
}

func (axis *Axis) CommitVoid(stage *Stage) {
	axis.Commit(stage)
}

// Checkout axis to the back repo (if it is already staged)
func (axis *Axis) Checkout(stage *Stage) *Axis {
	if _, ok := stage.Axiss[axis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAxis(axis)
		}
	}
	return axis
}

// for satisfaction of GongStruct interface
func (axis *Axis) GetName() (res string) {
	return axis.Name
}

// Stage puts axisgrid to the model stage
func (axisgrid *AxisGrid) Stage(stage *Stage) *AxisGrid {

	if _, ok := stage.AxisGrids[axisgrid]; !ok {
		stage.AxisGrids[axisgrid] = __member
		stage.AxisGridMap_Staged_Order[axisgrid] = stage.AxisGridOrder
		stage.AxisGridOrder++
	}
	stage.AxisGrids_mapString[axisgrid.Name] = axisgrid

	return axisgrid
}

// Unstage removes axisgrid off the model stage
func (axisgrid *AxisGrid) Unstage(stage *Stage) *AxisGrid {
	delete(stage.AxisGrids, axisgrid)
	delete(stage.AxisGrids_mapString, axisgrid.Name)
	return axisgrid
}

// UnstageVoid removes axisgrid off the model stage
func (axisgrid *AxisGrid) UnstageVoid(stage *Stage) {
	delete(stage.AxisGrids, axisgrid)
	delete(stage.AxisGrids_mapString, axisgrid.Name)
}

// commit axisgrid to the back repo (if it is already staged)
func (axisgrid *AxisGrid) Commit(stage *Stage) *AxisGrid {
	if _, ok := stage.AxisGrids[axisgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAxisGrid(axisgrid)
		}
	}
	return axisgrid
}

func (axisgrid *AxisGrid) CommitVoid(stage *Stage) {
	axisgrid.Commit(stage)
}

// Checkout axisgrid to the back repo (if it is already staged)
func (axisgrid *AxisGrid) Checkout(stage *Stage) *AxisGrid {
	if _, ok := stage.AxisGrids[axisgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAxisGrid(axisgrid)
		}
	}
	return axisgrid
}

// for satisfaction of GongStruct interface
func (axisgrid *AxisGrid) GetName() (res string) {
	return axisgrid.Name
}

// Stage puts bezier to the model stage
func (bezier *Bezier) Stage(stage *Stage) *Bezier {

	if _, ok := stage.Beziers[bezier]; !ok {
		stage.Beziers[bezier] = __member
		stage.BezierMap_Staged_Order[bezier] = stage.BezierOrder
		stage.BezierOrder++
	}
	stage.Beziers_mapString[bezier.Name] = bezier

	return bezier
}

// Unstage removes bezier off the model stage
func (bezier *Bezier) Unstage(stage *Stage) *Bezier {
	delete(stage.Beziers, bezier)
	delete(stage.Beziers_mapString, bezier.Name)
	return bezier
}

// UnstageVoid removes bezier off the model stage
func (bezier *Bezier) UnstageVoid(stage *Stage) {
	delete(stage.Beziers, bezier)
	delete(stage.Beziers_mapString, bezier.Name)
}

// commit bezier to the back repo (if it is already staged)
func (bezier *Bezier) Commit(stage *Stage) *Bezier {
	if _, ok := stage.Beziers[bezier]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezier(bezier)
		}
	}
	return bezier
}

func (bezier *Bezier) CommitVoid(stage *Stage) {
	bezier.Commit(stage)
}

// Checkout bezier to the back repo (if it is already staged)
func (bezier *Bezier) Checkout(stage *Stage) *Bezier {
	if _, ok := stage.Beziers[bezier]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBezier(bezier)
		}
	}
	return bezier
}

// for satisfaction of GongStruct interface
func (bezier *Bezier) GetName() (res string) {
	return bezier.Name
}

// Stage puts beziergrid to the model stage
func (beziergrid *BezierGrid) Stage(stage *Stage) *BezierGrid {

	if _, ok := stage.BezierGrids[beziergrid]; !ok {
		stage.BezierGrids[beziergrid] = __member
		stage.BezierGridMap_Staged_Order[beziergrid] = stage.BezierGridOrder
		stage.BezierGridOrder++
	}
	stage.BezierGrids_mapString[beziergrid.Name] = beziergrid

	return beziergrid
}

// Unstage removes beziergrid off the model stage
func (beziergrid *BezierGrid) Unstage(stage *Stage) *BezierGrid {
	delete(stage.BezierGrids, beziergrid)
	delete(stage.BezierGrids_mapString, beziergrid.Name)
	return beziergrid
}

// UnstageVoid removes beziergrid off the model stage
func (beziergrid *BezierGrid) UnstageVoid(stage *Stage) {
	delete(stage.BezierGrids, beziergrid)
	delete(stage.BezierGrids_mapString, beziergrid.Name)
}

// commit beziergrid to the back repo (if it is already staged)
func (beziergrid *BezierGrid) Commit(stage *Stage) *BezierGrid {
	if _, ok := stage.BezierGrids[beziergrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezierGrid(beziergrid)
		}
	}
	return beziergrid
}

func (beziergrid *BezierGrid) CommitVoid(stage *Stage) {
	beziergrid.Commit(stage)
}

// Checkout beziergrid to the back repo (if it is already staged)
func (beziergrid *BezierGrid) Checkout(stage *Stage) *BezierGrid {
	if _, ok := stage.BezierGrids[beziergrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBezierGrid(beziergrid)
		}
	}
	return beziergrid
}

// for satisfaction of GongStruct interface
func (beziergrid *BezierGrid) GetName() (res string) {
	return beziergrid.Name
}

// Stage puts beziergridstack to the model stage
func (beziergridstack *BezierGridStack) Stage(stage *Stage) *BezierGridStack {

	if _, ok := stage.BezierGridStacks[beziergridstack]; !ok {
		stage.BezierGridStacks[beziergridstack] = __member
		stage.BezierGridStackMap_Staged_Order[beziergridstack] = stage.BezierGridStackOrder
		stage.BezierGridStackOrder++
	}
	stage.BezierGridStacks_mapString[beziergridstack.Name] = beziergridstack

	return beziergridstack
}

// Unstage removes beziergridstack off the model stage
func (beziergridstack *BezierGridStack) Unstage(stage *Stage) *BezierGridStack {
	delete(stage.BezierGridStacks, beziergridstack)
	delete(stage.BezierGridStacks_mapString, beziergridstack.Name)
	return beziergridstack
}

// UnstageVoid removes beziergridstack off the model stage
func (beziergridstack *BezierGridStack) UnstageVoid(stage *Stage) {
	delete(stage.BezierGridStacks, beziergridstack)
	delete(stage.BezierGridStacks_mapString, beziergridstack.Name)
}

// commit beziergridstack to the back repo (if it is already staged)
func (beziergridstack *BezierGridStack) Commit(stage *Stage) *BezierGridStack {
	if _, ok := stage.BezierGridStacks[beziergridstack]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezierGridStack(beziergridstack)
		}
	}
	return beziergridstack
}

func (beziergridstack *BezierGridStack) CommitVoid(stage *Stage) {
	beziergridstack.Commit(stage)
}

// Checkout beziergridstack to the back repo (if it is already staged)
func (beziergridstack *BezierGridStack) Checkout(stage *Stage) *BezierGridStack {
	if _, ok := stage.BezierGridStacks[beziergridstack]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBezierGridStack(beziergridstack)
		}
	}
	return beziergridstack
}

// for satisfaction of GongStruct interface
func (beziergridstack *BezierGridStack) GetName() (res string) {
	return beziergridstack.Name
}

// Stage puts circle to the model stage
func (circle *Circle) Stage(stage *Stage) *Circle {

	if _, ok := stage.Circles[circle]; !ok {
		stage.Circles[circle] = __member
		stage.CircleMap_Staged_Order[circle] = stage.CircleOrder
		stage.CircleOrder++
	}
	stage.Circles_mapString[circle.Name] = circle

	return circle
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage(stage *Stage) *Circle {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
	return circle
}

// UnstageVoid removes circle off the model stage
func (circle *Circle) UnstageVoid(stage *Stage) {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
}

// commit circle to the back repo (if it is already staged)
func (circle *Circle) Commit(stage *Stage) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircle(circle)
		}
	}
	return circle
}

func (circle *Circle) CommitVoid(stage *Stage) {
	circle.Commit(stage)
}

// Checkout circle to the back repo (if it is already staged)
func (circle *Circle) Checkout(stage *Stage) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCircle(circle)
		}
	}
	return circle
}

// for satisfaction of GongStruct interface
func (circle *Circle) GetName() (res string) {
	return circle.Name
}

// Stage puts circlegrid to the model stage
func (circlegrid *CircleGrid) Stage(stage *Stage) *CircleGrid {

	if _, ok := stage.CircleGrids[circlegrid]; !ok {
		stage.CircleGrids[circlegrid] = __member
		stage.CircleGridMap_Staged_Order[circlegrid] = stage.CircleGridOrder
		stage.CircleGridOrder++
	}
	stage.CircleGrids_mapString[circlegrid.Name] = circlegrid

	return circlegrid
}

// Unstage removes circlegrid off the model stage
func (circlegrid *CircleGrid) Unstage(stage *Stage) *CircleGrid {
	delete(stage.CircleGrids, circlegrid)
	delete(stage.CircleGrids_mapString, circlegrid.Name)
	return circlegrid
}

// UnstageVoid removes circlegrid off the model stage
func (circlegrid *CircleGrid) UnstageVoid(stage *Stage) {
	delete(stage.CircleGrids, circlegrid)
	delete(stage.CircleGrids_mapString, circlegrid.Name)
}

// commit circlegrid to the back repo (if it is already staged)
func (circlegrid *CircleGrid) Commit(stage *Stage) *CircleGrid {
	if _, ok := stage.CircleGrids[circlegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircleGrid(circlegrid)
		}
	}
	return circlegrid
}

func (circlegrid *CircleGrid) CommitVoid(stage *Stage) {
	circlegrid.Commit(stage)
}

// Checkout circlegrid to the back repo (if it is already staged)
func (circlegrid *CircleGrid) Checkout(stage *Stage) *CircleGrid {
	if _, ok := stage.CircleGrids[circlegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCircleGrid(circlegrid)
		}
	}
	return circlegrid
}

// for satisfaction of GongStruct interface
func (circlegrid *CircleGrid) GetName() (res string) {
	return circlegrid.Name
}

// Stage puts exporttomusicxml to the model stage
func (exporttomusicxml *ExportToMusicxml) Stage(stage *Stage) *ExportToMusicxml {

	if _, ok := stage.ExportToMusicxmls[exporttomusicxml]; !ok {
		stage.ExportToMusicxmls[exporttomusicxml] = __member
		stage.ExportToMusicxmlMap_Staged_Order[exporttomusicxml] = stage.ExportToMusicxmlOrder
		stage.ExportToMusicxmlOrder++
	}
	stage.ExportToMusicxmls_mapString[exporttomusicxml.Name] = exporttomusicxml

	return exporttomusicxml
}

// Unstage removes exporttomusicxml off the model stage
func (exporttomusicxml *ExportToMusicxml) Unstage(stage *Stage) *ExportToMusicxml {
	delete(stage.ExportToMusicxmls, exporttomusicxml)
	delete(stage.ExportToMusicxmls_mapString, exporttomusicxml.Name)
	return exporttomusicxml
}

// UnstageVoid removes exporttomusicxml off the model stage
func (exporttomusicxml *ExportToMusicxml) UnstageVoid(stage *Stage) {
	delete(stage.ExportToMusicxmls, exporttomusicxml)
	delete(stage.ExportToMusicxmls_mapString, exporttomusicxml.Name)
}

// commit exporttomusicxml to the back repo (if it is already staged)
func (exporttomusicxml *ExportToMusicxml) Commit(stage *Stage) *ExportToMusicxml {
	if _, ok := stage.ExportToMusicxmls[exporttomusicxml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitExportToMusicxml(exporttomusicxml)
		}
	}
	return exporttomusicxml
}

func (exporttomusicxml *ExportToMusicxml) CommitVoid(stage *Stage) {
	exporttomusicxml.Commit(stage)
}

// Checkout exporttomusicxml to the back repo (if it is already staged)
func (exporttomusicxml *ExportToMusicxml) Checkout(stage *Stage) *ExportToMusicxml {
	if _, ok := stage.ExportToMusicxmls[exporttomusicxml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutExportToMusicxml(exporttomusicxml)
		}
	}
	return exporttomusicxml
}

// for satisfaction of GongStruct interface
func (exporttomusicxml *ExportToMusicxml) GetName() (res string) {
	return exporttomusicxml.Name
}

// Stage puts frontcurve to the model stage
func (frontcurve *FrontCurve) Stage(stage *Stage) *FrontCurve {

	if _, ok := stage.FrontCurves[frontcurve]; !ok {
		stage.FrontCurves[frontcurve] = __member
		stage.FrontCurveMap_Staged_Order[frontcurve] = stage.FrontCurveOrder
		stage.FrontCurveOrder++
	}
	stage.FrontCurves_mapString[frontcurve.Name] = frontcurve

	return frontcurve
}

// Unstage removes frontcurve off the model stage
func (frontcurve *FrontCurve) Unstage(stage *Stage) *FrontCurve {
	delete(stage.FrontCurves, frontcurve)
	delete(stage.FrontCurves_mapString, frontcurve.Name)
	return frontcurve
}

// UnstageVoid removes frontcurve off the model stage
func (frontcurve *FrontCurve) UnstageVoid(stage *Stage) {
	delete(stage.FrontCurves, frontcurve)
	delete(stage.FrontCurves_mapString, frontcurve.Name)
}

// commit frontcurve to the back repo (if it is already staged)
func (frontcurve *FrontCurve) Commit(stage *Stage) *FrontCurve {
	if _, ok := stage.FrontCurves[frontcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFrontCurve(frontcurve)
		}
	}
	return frontcurve
}

func (frontcurve *FrontCurve) CommitVoid(stage *Stage) {
	frontcurve.Commit(stage)
}

// Checkout frontcurve to the back repo (if it is already staged)
func (frontcurve *FrontCurve) Checkout(stage *Stage) *FrontCurve {
	if _, ok := stage.FrontCurves[frontcurve]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFrontCurve(frontcurve)
		}
	}
	return frontcurve
}

// for satisfaction of GongStruct interface
func (frontcurve *FrontCurve) GetName() (res string) {
	return frontcurve.Name
}

// Stage puts frontcurvestack to the model stage
func (frontcurvestack *FrontCurveStack) Stage(stage *Stage) *FrontCurveStack {

	if _, ok := stage.FrontCurveStacks[frontcurvestack]; !ok {
		stage.FrontCurveStacks[frontcurvestack] = __member
		stage.FrontCurveStackMap_Staged_Order[frontcurvestack] = stage.FrontCurveStackOrder
		stage.FrontCurveStackOrder++
	}
	stage.FrontCurveStacks_mapString[frontcurvestack.Name] = frontcurvestack

	return frontcurvestack
}

// Unstage removes frontcurvestack off the model stage
func (frontcurvestack *FrontCurveStack) Unstage(stage *Stage) *FrontCurveStack {
	delete(stage.FrontCurveStacks, frontcurvestack)
	delete(stage.FrontCurveStacks_mapString, frontcurvestack.Name)
	return frontcurvestack
}

// UnstageVoid removes frontcurvestack off the model stage
func (frontcurvestack *FrontCurveStack) UnstageVoid(stage *Stage) {
	delete(stage.FrontCurveStacks, frontcurvestack)
	delete(stage.FrontCurveStacks_mapString, frontcurvestack.Name)
}

// commit frontcurvestack to the back repo (if it is already staged)
func (frontcurvestack *FrontCurveStack) Commit(stage *Stage) *FrontCurveStack {
	if _, ok := stage.FrontCurveStacks[frontcurvestack]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFrontCurveStack(frontcurvestack)
		}
	}
	return frontcurvestack
}

func (frontcurvestack *FrontCurveStack) CommitVoid(stage *Stage) {
	frontcurvestack.Commit(stage)
}

// Checkout frontcurvestack to the back repo (if it is already staged)
func (frontcurvestack *FrontCurveStack) Checkout(stage *Stage) *FrontCurveStack {
	if _, ok := stage.FrontCurveStacks[frontcurvestack]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFrontCurveStack(frontcurvestack)
		}
	}
	return frontcurvestack
}

// for satisfaction of GongStruct interface
func (frontcurvestack *FrontCurveStack) GetName() (res string) {
	return frontcurvestack.Name
}

// Stage puts horizontalaxis to the model stage
func (horizontalaxis *HorizontalAxis) Stage(stage *Stage) *HorizontalAxis {

	if _, ok := stage.HorizontalAxiss[horizontalaxis]; !ok {
		stage.HorizontalAxiss[horizontalaxis] = __member
		stage.HorizontalAxisMap_Staged_Order[horizontalaxis] = stage.HorizontalAxisOrder
		stage.HorizontalAxisOrder++
	}
	stage.HorizontalAxiss_mapString[horizontalaxis.Name] = horizontalaxis

	return horizontalaxis
}

// Unstage removes horizontalaxis off the model stage
func (horizontalaxis *HorizontalAxis) Unstage(stage *Stage) *HorizontalAxis {
	delete(stage.HorizontalAxiss, horizontalaxis)
	delete(stage.HorizontalAxiss_mapString, horizontalaxis.Name)
	return horizontalaxis
}

// UnstageVoid removes horizontalaxis off the model stage
func (horizontalaxis *HorizontalAxis) UnstageVoid(stage *Stage) {
	delete(stage.HorizontalAxiss, horizontalaxis)
	delete(stage.HorizontalAxiss_mapString, horizontalaxis.Name)
}

// commit horizontalaxis to the back repo (if it is already staged)
func (horizontalaxis *HorizontalAxis) Commit(stage *Stage) *HorizontalAxis {
	if _, ok := stage.HorizontalAxiss[horizontalaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitHorizontalAxis(horizontalaxis)
		}
	}
	return horizontalaxis
}

func (horizontalaxis *HorizontalAxis) CommitVoid(stage *Stage) {
	horizontalaxis.Commit(stage)
}

// Checkout horizontalaxis to the back repo (if it is already staged)
func (horizontalaxis *HorizontalAxis) Checkout(stage *Stage) *HorizontalAxis {
	if _, ok := stage.HorizontalAxiss[horizontalaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutHorizontalAxis(horizontalaxis)
		}
	}
	return horizontalaxis
}

// for satisfaction of GongStruct interface
func (horizontalaxis *HorizontalAxis) GetName() (res string) {
	return horizontalaxis.Name
}

// Stage puts key to the model stage
func (key *Key) Stage(stage *Stage) *Key {

	if _, ok := stage.Keys[key]; !ok {
		stage.Keys[key] = __member
		stage.KeyMap_Staged_Order[key] = stage.KeyOrder
		stage.KeyOrder++
	}
	stage.Keys_mapString[key.Name] = key

	return key
}

// Unstage removes key off the model stage
func (key *Key) Unstage(stage *Stage) *Key {
	delete(stage.Keys, key)
	delete(stage.Keys_mapString, key.Name)
	return key
}

// UnstageVoid removes key off the model stage
func (key *Key) UnstageVoid(stage *Stage) {
	delete(stage.Keys, key)
	delete(stage.Keys_mapString, key.Name)
}

// commit key to the back repo (if it is already staged)
func (key *Key) Commit(stage *Stage) *Key {
	if _, ok := stage.Keys[key]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitKey(key)
		}
	}
	return key
}

func (key *Key) CommitVoid(stage *Stage) {
	key.Commit(stage)
}

// Checkout key to the back repo (if it is already staged)
func (key *Key) Checkout(stage *Stage) *Key {
	if _, ok := stage.Keys[key]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutKey(key)
		}
	}
	return key
}

// for satisfaction of GongStruct interface
func (key *Key) GetName() (res string) {
	return key.Name
}

// Stage puts parameter to the model stage
func (parameter *Parameter) Stage(stage *Stage) *Parameter {

	if _, ok := stage.Parameters[parameter]; !ok {
		stage.Parameters[parameter] = __member
		stage.ParameterMap_Staged_Order[parameter] = stage.ParameterOrder
		stage.ParameterOrder++
	}
	stage.Parameters_mapString[parameter.Name] = parameter

	return parameter
}

// Unstage removes parameter off the model stage
func (parameter *Parameter) Unstage(stage *Stage) *Parameter {
	delete(stage.Parameters, parameter)
	delete(stage.Parameters_mapString, parameter.Name)
	return parameter
}

// UnstageVoid removes parameter off the model stage
func (parameter *Parameter) UnstageVoid(stage *Stage) {
	delete(stage.Parameters, parameter)
	delete(stage.Parameters_mapString, parameter.Name)
}

// commit parameter to the back repo (if it is already staged)
func (parameter *Parameter) Commit(stage *Stage) *Parameter {
	if _, ok := stage.Parameters[parameter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParameter(parameter)
		}
	}
	return parameter
}

func (parameter *Parameter) CommitVoid(stage *Stage) {
	parameter.Commit(stage)
}

// Checkout parameter to the back repo (if it is already staged)
func (parameter *Parameter) Checkout(stage *Stage) *Parameter {
	if _, ok := stage.Parameters[parameter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParameter(parameter)
		}
	}
	return parameter
}

// for satisfaction of GongStruct interface
func (parameter *Parameter) GetName() (res string) {
	return parameter.Name
}

// Stage puts rhombus to the model stage
func (rhombus *Rhombus) Stage(stage *Stage) *Rhombus {

	if _, ok := stage.Rhombuss[rhombus]; !ok {
		stage.Rhombuss[rhombus] = __member
		stage.RhombusMap_Staged_Order[rhombus] = stage.RhombusOrder
		stage.RhombusOrder++
	}
	stage.Rhombuss_mapString[rhombus.Name] = rhombus

	return rhombus
}

// Unstage removes rhombus off the model stage
func (rhombus *Rhombus) Unstage(stage *Stage) *Rhombus {
	delete(stage.Rhombuss, rhombus)
	delete(stage.Rhombuss_mapString, rhombus.Name)
	return rhombus
}

// UnstageVoid removes rhombus off the model stage
func (rhombus *Rhombus) UnstageVoid(stage *Stage) {
	delete(stage.Rhombuss, rhombus)
	delete(stage.Rhombuss_mapString, rhombus.Name)
}

// commit rhombus to the back repo (if it is already staged)
func (rhombus *Rhombus) Commit(stage *Stage) *Rhombus {
	if _, ok := stage.Rhombuss[rhombus]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRhombus(rhombus)
		}
	}
	return rhombus
}

func (rhombus *Rhombus) CommitVoid(stage *Stage) {
	rhombus.Commit(stage)
}

// Checkout rhombus to the back repo (if it is already staged)
func (rhombus *Rhombus) Checkout(stage *Stage) *Rhombus {
	if _, ok := stage.Rhombuss[rhombus]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRhombus(rhombus)
		}
	}
	return rhombus
}

// for satisfaction of GongStruct interface
func (rhombus *Rhombus) GetName() (res string) {
	return rhombus.Name
}

// Stage puts rhombusgrid to the model stage
func (rhombusgrid *RhombusGrid) Stage(stage *Stage) *RhombusGrid {

	if _, ok := stage.RhombusGrids[rhombusgrid]; !ok {
		stage.RhombusGrids[rhombusgrid] = __member
		stage.RhombusGridMap_Staged_Order[rhombusgrid] = stage.RhombusGridOrder
		stage.RhombusGridOrder++
	}
	stage.RhombusGrids_mapString[rhombusgrid.Name] = rhombusgrid

	return rhombusgrid
}

// Unstage removes rhombusgrid off the model stage
func (rhombusgrid *RhombusGrid) Unstage(stage *Stage) *RhombusGrid {
	delete(stage.RhombusGrids, rhombusgrid)
	delete(stage.RhombusGrids_mapString, rhombusgrid.Name)
	return rhombusgrid
}

// UnstageVoid removes rhombusgrid off the model stage
func (rhombusgrid *RhombusGrid) UnstageVoid(stage *Stage) {
	delete(stage.RhombusGrids, rhombusgrid)
	delete(stage.RhombusGrids_mapString, rhombusgrid.Name)
}

// commit rhombusgrid to the back repo (if it is already staged)
func (rhombusgrid *RhombusGrid) Commit(stage *Stage) *RhombusGrid {
	if _, ok := stage.RhombusGrids[rhombusgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRhombusGrid(rhombusgrid)
		}
	}
	return rhombusgrid
}

func (rhombusgrid *RhombusGrid) CommitVoid(stage *Stage) {
	rhombusgrid.Commit(stage)
}

// Checkout rhombusgrid to the back repo (if it is already staged)
func (rhombusgrid *RhombusGrid) Checkout(stage *Stage) *RhombusGrid {
	if _, ok := stage.RhombusGrids[rhombusgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRhombusGrid(rhombusgrid)
		}
	}
	return rhombusgrid
}

// for satisfaction of GongStruct interface
func (rhombusgrid *RhombusGrid) GetName() (res string) {
	return rhombusgrid.Name
}

// Stage puts shapecategory to the model stage
func (shapecategory *ShapeCategory) Stage(stage *Stage) *ShapeCategory {

	if _, ok := stage.ShapeCategorys[shapecategory]; !ok {
		stage.ShapeCategorys[shapecategory] = __member
		stage.ShapeCategoryMap_Staged_Order[shapecategory] = stage.ShapeCategoryOrder
		stage.ShapeCategoryOrder++
	}
	stage.ShapeCategorys_mapString[shapecategory.Name] = shapecategory

	return shapecategory
}

// Unstage removes shapecategory off the model stage
func (shapecategory *ShapeCategory) Unstage(stage *Stage) *ShapeCategory {
	delete(stage.ShapeCategorys, shapecategory)
	delete(stage.ShapeCategorys_mapString, shapecategory.Name)
	return shapecategory
}

// UnstageVoid removes shapecategory off the model stage
func (shapecategory *ShapeCategory) UnstageVoid(stage *Stage) {
	delete(stage.ShapeCategorys, shapecategory)
	delete(stage.ShapeCategorys_mapString, shapecategory.Name)
}

// commit shapecategory to the back repo (if it is already staged)
func (shapecategory *ShapeCategory) Commit(stage *Stage) *ShapeCategory {
	if _, ok := stage.ShapeCategorys[shapecategory]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShapeCategory(shapecategory)
		}
	}
	return shapecategory
}

func (shapecategory *ShapeCategory) CommitVoid(stage *Stage) {
	shapecategory.Commit(stage)
}

// Checkout shapecategory to the back repo (if it is already staged)
func (shapecategory *ShapeCategory) Checkout(stage *Stage) *ShapeCategory {
	if _, ok := stage.ShapeCategorys[shapecategory]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutShapeCategory(shapecategory)
		}
	}
	return shapecategory
}

// for satisfaction of GongStruct interface
func (shapecategory *ShapeCategory) GetName() (res string) {
	return shapecategory.Name
}

// Stage puts spiralbezier to the model stage
func (spiralbezier *SpiralBezier) Stage(stage *Stage) *SpiralBezier {

	if _, ok := stage.SpiralBeziers[spiralbezier]; !ok {
		stage.SpiralBeziers[spiralbezier] = __member
		stage.SpiralBezierMap_Staged_Order[spiralbezier] = stage.SpiralBezierOrder
		stage.SpiralBezierOrder++
	}
	stage.SpiralBeziers_mapString[spiralbezier.Name] = spiralbezier

	return spiralbezier
}

// Unstage removes spiralbezier off the model stage
func (spiralbezier *SpiralBezier) Unstage(stage *Stage) *SpiralBezier {
	delete(stage.SpiralBeziers, spiralbezier)
	delete(stage.SpiralBeziers_mapString, spiralbezier.Name)
	return spiralbezier
}

// UnstageVoid removes spiralbezier off the model stage
func (spiralbezier *SpiralBezier) UnstageVoid(stage *Stage) {
	delete(stage.SpiralBeziers, spiralbezier)
	delete(stage.SpiralBeziers_mapString, spiralbezier.Name)
}

// commit spiralbezier to the back repo (if it is already staged)
func (spiralbezier *SpiralBezier) Commit(stage *Stage) *SpiralBezier {
	if _, ok := stage.SpiralBeziers[spiralbezier]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralBezier(spiralbezier)
		}
	}
	return spiralbezier
}

func (spiralbezier *SpiralBezier) CommitVoid(stage *Stage) {
	spiralbezier.Commit(stage)
}

// Checkout spiralbezier to the back repo (if it is already staged)
func (spiralbezier *SpiralBezier) Checkout(stage *Stage) *SpiralBezier {
	if _, ok := stage.SpiralBeziers[spiralbezier]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralBezier(spiralbezier)
		}
	}
	return spiralbezier
}

// for satisfaction of GongStruct interface
func (spiralbezier *SpiralBezier) GetName() (res string) {
	return spiralbezier.Name
}

// Stage puts spiralbeziergrid to the model stage
func (spiralbeziergrid *SpiralBezierGrid) Stage(stage *Stage) *SpiralBezierGrid {

	if _, ok := stage.SpiralBezierGrids[spiralbeziergrid]; !ok {
		stage.SpiralBezierGrids[spiralbeziergrid] = __member
		stage.SpiralBezierGridMap_Staged_Order[spiralbeziergrid] = stage.SpiralBezierGridOrder
		stage.SpiralBezierGridOrder++
	}
	stage.SpiralBezierGrids_mapString[spiralbeziergrid.Name] = spiralbeziergrid

	return spiralbeziergrid
}

// Unstage removes spiralbeziergrid off the model stage
func (spiralbeziergrid *SpiralBezierGrid) Unstage(stage *Stage) *SpiralBezierGrid {
	delete(stage.SpiralBezierGrids, spiralbeziergrid)
	delete(stage.SpiralBezierGrids_mapString, spiralbeziergrid.Name)
	return spiralbeziergrid
}

// UnstageVoid removes spiralbeziergrid off the model stage
func (spiralbeziergrid *SpiralBezierGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralBezierGrids, spiralbeziergrid)
	delete(stage.SpiralBezierGrids_mapString, spiralbeziergrid.Name)
}

// commit spiralbeziergrid to the back repo (if it is already staged)
func (spiralbeziergrid *SpiralBezierGrid) Commit(stage *Stage) *SpiralBezierGrid {
	if _, ok := stage.SpiralBezierGrids[spiralbeziergrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralBezierGrid(spiralbeziergrid)
		}
	}
	return spiralbeziergrid
}

func (spiralbeziergrid *SpiralBezierGrid) CommitVoid(stage *Stage) {
	spiralbeziergrid.Commit(stage)
}

// Checkout spiralbeziergrid to the back repo (if it is already staged)
func (spiralbeziergrid *SpiralBezierGrid) Checkout(stage *Stage) *SpiralBezierGrid {
	if _, ok := stage.SpiralBezierGrids[spiralbeziergrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralBezierGrid(spiralbeziergrid)
		}
	}
	return spiralbeziergrid
}

// for satisfaction of GongStruct interface
func (spiralbeziergrid *SpiralBezierGrid) GetName() (res string) {
	return spiralbeziergrid.Name
}

// Stage puts spiralcircle to the model stage
func (spiralcircle *SpiralCircle) Stage(stage *Stage) *SpiralCircle {

	if _, ok := stage.SpiralCircles[spiralcircle]; !ok {
		stage.SpiralCircles[spiralcircle] = __member
		stage.SpiralCircleMap_Staged_Order[spiralcircle] = stage.SpiralCircleOrder
		stage.SpiralCircleOrder++
	}
	stage.SpiralCircles_mapString[spiralcircle.Name] = spiralcircle

	return spiralcircle
}

// Unstage removes spiralcircle off the model stage
func (spiralcircle *SpiralCircle) Unstage(stage *Stage) *SpiralCircle {
	delete(stage.SpiralCircles, spiralcircle)
	delete(stage.SpiralCircles_mapString, spiralcircle.Name)
	return spiralcircle
}

// UnstageVoid removes spiralcircle off the model stage
func (spiralcircle *SpiralCircle) UnstageVoid(stage *Stage) {
	delete(stage.SpiralCircles, spiralcircle)
	delete(stage.SpiralCircles_mapString, spiralcircle.Name)
}

// commit spiralcircle to the back repo (if it is already staged)
func (spiralcircle *SpiralCircle) Commit(stage *Stage) *SpiralCircle {
	if _, ok := stage.SpiralCircles[spiralcircle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralCircle(spiralcircle)
		}
	}
	return spiralcircle
}

func (spiralcircle *SpiralCircle) CommitVoid(stage *Stage) {
	spiralcircle.Commit(stage)
}

// Checkout spiralcircle to the back repo (if it is already staged)
func (spiralcircle *SpiralCircle) Checkout(stage *Stage) *SpiralCircle {
	if _, ok := stage.SpiralCircles[spiralcircle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralCircle(spiralcircle)
		}
	}
	return spiralcircle
}

// for satisfaction of GongStruct interface
func (spiralcircle *SpiralCircle) GetName() (res string) {
	return spiralcircle.Name
}

// Stage puts spiralcirclegrid to the model stage
func (spiralcirclegrid *SpiralCircleGrid) Stage(stage *Stage) *SpiralCircleGrid {

	if _, ok := stage.SpiralCircleGrids[spiralcirclegrid]; !ok {
		stage.SpiralCircleGrids[spiralcirclegrid] = __member
		stage.SpiralCircleGridMap_Staged_Order[spiralcirclegrid] = stage.SpiralCircleGridOrder
		stage.SpiralCircleGridOrder++
	}
	stage.SpiralCircleGrids_mapString[spiralcirclegrid.Name] = spiralcirclegrid

	return spiralcirclegrid
}

// Unstage removes spiralcirclegrid off the model stage
func (spiralcirclegrid *SpiralCircleGrid) Unstage(stage *Stage) *SpiralCircleGrid {
	delete(stage.SpiralCircleGrids, spiralcirclegrid)
	delete(stage.SpiralCircleGrids_mapString, spiralcirclegrid.Name)
	return spiralcirclegrid
}

// UnstageVoid removes spiralcirclegrid off the model stage
func (spiralcirclegrid *SpiralCircleGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralCircleGrids, spiralcirclegrid)
	delete(stage.SpiralCircleGrids_mapString, spiralcirclegrid.Name)
}

// commit spiralcirclegrid to the back repo (if it is already staged)
func (spiralcirclegrid *SpiralCircleGrid) Commit(stage *Stage) *SpiralCircleGrid {
	if _, ok := stage.SpiralCircleGrids[spiralcirclegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralCircleGrid(spiralcirclegrid)
		}
	}
	return spiralcirclegrid
}

func (spiralcirclegrid *SpiralCircleGrid) CommitVoid(stage *Stage) {
	spiralcirclegrid.Commit(stage)
}

// Checkout spiralcirclegrid to the back repo (if it is already staged)
func (spiralcirclegrid *SpiralCircleGrid) Checkout(stage *Stage) *SpiralCircleGrid {
	if _, ok := stage.SpiralCircleGrids[spiralcirclegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralCircleGrid(spiralcirclegrid)
		}
	}
	return spiralcirclegrid
}

// for satisfaction of GongStruct interface
func (spiralcirclegrid *SpiralCircleGrid) GetName() (res string) {
	return spiralcirclegrid.Name
}

// Stage puts spiralline to the model stage
func (spiralline *SpiralLine) Stage(stage *Stage) *SpiralLine {

	if _, ok := stage.SpiralLines[spiralline]; !ok {
		stage.SpiralLines[spiralline] = __member
		stage.SpiralLineMap_Staged_Order[spiralline] = stage.SpiralLineOrder
		stage.SpiralLineOrder++
	}
	stage.SpiralLines_mapString[spiralline.Name] = spiralline

	return spiralline
}

// Unstage removes spiralline off the model stage
func (spiralline *SpiralLine) Unstage(stage *Stage) *SpiralLine {
	delete(stage.SpiralLines, spiralline)
	delete(stage.SpiralLines_mapString, spiralline.Name)
	return spiralline
}

// UnstageVoid removes spiralline off the model stage
func (spiralline *SpiralLine) UnstageVoid(stage *Stage) {
	delete(stage.SpiralLines, spiralline)
	delete(stage.SpiralLines_mapString, spiralline.Name)
}

// commit spiralline to the back repo (if it is already staged)
func (spiralline *SpiralLine) Commit(stage *Stage) *SpiralLine {
	if _, ok := stage.SpiralLines[spiralline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralLine(spiralline)
		}
	}
	return spiralline
}

func (spiralline *SpiralLine) CommitVoid(stage *Stage) {
	spiralline.Commit(stage)
}

// Checkout spiralline to the back repo (if it is already staged)
func (spiralline *SpiralLine) Checkout(stage *Stage) *SpiralLine {
	if _, ok := stage.SpiralLines[spiralline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralLine(spiralline)
		}
	}
	return spiralline
}

// for satisfaction of GongStruct interface
func (spiralline *SpiralLine) GetName() (res string) {
	return spiralline.Name
}

// Stage puts spirallinegrid to the model stage
func (spirallinegrid *SpiralLineGrid) Stage(stage *Stage) *SpiralLineGrid {

	if _, ok := stage.SpiralLineGrids[spirallinegrid]; !ok {
		stage.SpiralLineGrids[spirallinegrid] = __member
		stage.SpiralLineGridMap_Staged_Order[spirallinegrid] = stage.SpiralLineGridOrder
		stage.SpiralLineGridOrder++
	}
	stage.SpiralLineGrids_mapString[spirallinegrid.Name] = spirallinegrid

	return spirallinegrid
}

// Unstage removes spirallinegrid off the model stage
func (spirallinegrid *SpiralLineGrid) Unstage(stage *Stage) *SpiralLineGrid {
	delete(stage.SpiralLineGrids, spirallinegrid)
	delete(stage.SpiralLineGrids_mapString, spirallinegrid.Name)
	return spirallinegrid
}

// UnstageVoid removes spirallinegrid off the model stage
func (spirallinegrid *SpiralLineGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralLineGrids, spirallinegrid)
	delete(stage.SpiralLineGrids_mapString, spirallinegrid.Name)
}

// commit spirallinegrid to the back repo (if it is already staged)
func (spirallinegrid *SpiralLineGrid) Commit(stage *Stage) *SpiralLineGrid {
	if _, ok := stage.SpiralLineGrids[spirallinegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralLineGrid(spirallinegrid)
		}
	}
	return spirallinegrid
}

func (spirallinegrid *SpiralLineGrid) CommitVoid(stage *Stage) {
	spirallinegrid.Commit(stage)
}

// Checkout spirallinegrid to the back repo (if it is already staged)
func (spirallinegrid *SpiralLineGrid) Checkout(stage *Stage) *SpiralLineGrid {
	if _, ok := stage.SpiralLineGrids[spirallinegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralLineGrid(spirallinegrid)
		}
	}
	return spirallinegrid
}

// for satisfaction of GongStruct interface
func (spirallinegrid *SpiralLineGrid) GetName() (res string) {
	return spirallinegrid.Name
}

// Stage puts spiralorigin to the model stage
func (spiralorigin *SpiralOrigin) Stage(stage *Stage) *SpiralOrigin {

	if _, ok := stage.SpiralOrigins[spiralorigin]; !ok {
		stage.SpiralOrigins[spiralorigin] = __member
		stage.SpiralOriginMap_Staged_Order[spiralorigin] = stage.SpiralOriginOrder
		stage.SpiralOriginOrder++
	}
	stage.SpiralOrigins_mapString[spiralorigin.Name] = spiralorigin

	return spiralorigin
}

// Unstage removes spiralorigin off the model stage
func (spiralorigin *SpiralOrigin) Unstage(stage *Stage) *SpiralOrigin {
	delete(stage.SpiralOrigins, spiralorigin)
	delete(stage.SpiralOrigins_mapString, spiralorigin.Name)
	return spiralorigin
}

// UnstageVoid removes spiralorigin off the model stage
func (spiralorigin *SpiralOrigin) UnstageVoid(stage *Stage) {
	delete(stage.SpiralOrigins, spiralorigin)
	delete(stage.SpiralOrigins_mapString, spiralorigin.Name)
}

// commit spiralorigin to the back repo (if it is already staged)
func (spiralorigin *SpiralOrigin) Commit(stage *Stage) *SpiralOrigin {
	if _, ok := stage.SpiralOrigins[spiralorigin]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralOrigin(spiralorigin)
		}
	}
	return spiralorigin
}

func (spiralorigin *SpiralOrigin) CommitVoid(stage *Stage) {
	spiralorigin.Commit(stage)
}

// Checkout spiralorigin to the back repo (if it is already staged)
func (spiralorigin *SpiralOrigin) Checkout(stage *Stage) *SpiralOrigin {
	if _, ok := stage.SpiralOrigins[spiralorigin]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralOrigin(spiralorigin)
		}
	}
	return spiralorigin
}

// for satisfaction of GongStruct interface
func (spiralorigin *SpiralOrigin) GetName() (res string) {
	return spiralorigin.Name
}

// Stage puts spiralrhombus to the model stage
func (spiralrhombus *SpiralRhombus) Stage(stage *Stage) *SpiralRhombus {

	if _, ok := stage.SpiralRhombuss[spiralrhombus]; !ok {
		stage.SpiralRhombuss[spiralrhombus] = __member
		stage.SpiralRhombusMap_Staged_Order[spiralrhombus] = stage.SpiralRhombusOrder
		stage.SpiralRhombusOrder++
	}
	stage.SpiralRhombuss_mapString[spiralrhombus.Name] = spiralrhombus

	return spiralrhombus
}

// Unstage removes spiralrhombus off the model stage
func (spiralrhombus *SpiralRhombus) Unstage(stage *Stage) *SpiralRhombus {
	delete(stage.SpiralRhombuss, spiralrhombus)
	delete(stage.SpiralRhombuss_mapString, spiralrhombus.Name)
	return spiralrhombus
}

// UnstageVoid removes spiralrhombus off the model stage
func (spiralrhombus *SpiralRhombus) UnstageVoid(stage *Stage) {
	delete(stage.SpiralRhombuss, spiralrhombus)
	delete(stage.SpiralRhombuss_mapString, spiralrhombus.Name)
}

// commit spiralrhombus to the back repo (if it is already staged)
func (spiralrhombus *SpiralRhombus) Commit(stage *Stage) *SpiralRhombus {
	if _, ok := stage.SpiralRhombuss[spiralrhombus]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralRhombus(spiralrhombus)
		}
	}
	return spiralrhombus
}

func (spiralrhombus *SpiralRhombus) CommitVoid(stage *Stage) {
	spiralrhombus.Commit(stage)
}

// Checkout spiralrhombus to the back repo (if it is already staged)
func (spiralrhombus *SpiralRhombus) Checkout(stage *Stage) *SpiralRhombus {
	if _, ok := stage.SpiralRhombuss[spiralrhombus]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralRhombus(spiralrhombus)
		}
	}
	return spiralrhombus
}

// for satisfaction of GongStruct interface
func (spiralrhombus *SpiralRhombus) GetName() (res string) {
	return spiralrhombus.Name
}

// Stage puts spiralrhombusgrid to the model stage
func (spiralrhombusgrid *SpiralRhombusGrid) Stage(stage *Stage) *SpiralRhombusGrid {

	if _, ok := stage.SpiralRhombusGrids[spiralrhombusgrid]; !ok {
		stage.SpiralRhombusGrids[spiralrhombusgrid] = __member
		stage.SpiralRhombusGridMap_Staged_Order[spiralrhombusgrid] = stage.SpiralRhombusGridOrder
		stage.SpiralRhombusGridOrder++
	}
	stage.SpiralRhombusGrids_mapString[spiralrhombusgrid.Name] = spiralrhombusgrid

	return spiralrhombusgrid
}

// Unstage removes spiralrhombusgrid off the model stage
func (spiralrhombusgrid *SpiralRhombusGrid) Unstage(stage *Stage) *SpiralRhombusGrid {
	delete(stage.SpiralRhombusGrids, spiralrhombusgrid)
	delete(stage.SpiralRhombusGrids_mapString, spiralrhombusgrid.Name)
	return spiralrhombusgrid
}

// UnstageVoid removes spiralrhombusgrid off the model stage
func (spiralrhombusgrid *SpiralRhombusGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralRhombusGrids, spiralrhombusgrid)
	delete(stage.SpiralRhombusGrids_mapString, spiralrhombusgrid.Name)
}

// commit spiralrhombusgrid to the back repo (if it is already staged)
func (spiralrhombusgrid *SpiralRhombusGrid) Commit(stage *Stage) *SpiralRhombusGrid {
	if _, ok := stage.SpiralRhombusGrids[spiralrhombusgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSpiralRhombusGrid(spiralrhombusgrid)
		}
	}
	return spiralrhombusgrid
}

func (spiralrhombusgrid *SpiralRhombusGrid) CommitVoid(stage *Stage) {
	spiralrhombusgrid.Commit(stage)
}

// Checkout spiralrhombusgrid to the back repo (if it is already staged)
func (spiralrhombusgrid *SpiralRhombusGrid) Checkout(stage *Stage) *SpiralRhombusGrid {
	if _, ok := stage.SpiralRhombusGrids[spiralrhombusgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSpiralRhombusGrid(spiralrhombusgrid)
		}
	}
	return spiralrhombusgrid
}

// for satisfaction of GongStruct interface
func (spiralrhombusgrid *SpiralRhombusGrid) GetName() (res string) {
	return spiralrhombusgrid.Name
}

// Stage puts verticalaxis to the model stage
func (verticalaxis *VerticalAxis) Stage(stage *Stage) *VerticalAxis {

	if _, ok := stage.VerticalAxiss[verticalaxis]; !ok {
		stage.VerticalAxiss[verticalaxis] = __member
		stage.VerticalAxisMap_Staged_Order[verticalaxis] = stage.VerticalAxisOrder
		stage.VerticalAxisOrder++
	}
	stage.VerticalAxiss_mapString[verticalaxis.Name] = verticalaxis

	return verticalaxis
}

// Unstage removes verticalaxis off the model stage
func (verticalaxis *VerticalAxis) Unstage(stage *Stage) *VerticalAxis {
	delete(stage.VerticalAxiss, verticalaxis)
	delete(stage.VerticalAxiss_mapString, verticalaxis.Name)
	return verticalaxis
}

// UnstageVoid removes verticalaxis off the model stage
func (verticalaxis *VerticalAxis) UnstageVoid(stage *Stage) {
	delete(stage.VerticalAxiss, verticalaxis)
	delete(stage.VerticalAxiss_mapString, verticalaxis.Name)
}

// commit verticalaxis to the back repo (if it is already staged)
func (verticalaxis *VerticalAxis) Commit(stage *Stage) *VerticalAxis {
	if _, ok := stage.VerticalAxiss[verticalaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVerticalAxis(verticalaxis)
		}
	}
	return verticalaxis
}

func (verticalaxis *VerticalAxis) CommitVoid(stage *Stage) {
	verticalaxis.Commit(stage)
}

// Checkout verticalaxis to the back repo (if it is already staged)
func (verticalaxis *VerticalAxis) Checkout(stage *Stage) *VerticalAxis {
	if _, ok := stage.VerticalAxiss[verticalaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVerticalAxis(verticalaxis)
		}
	}
	return verticalaxis
}

// for satisfaction of GongStruct interface
func (verticalaxis *VerticalAxis) GetName() (res string) {
	return verticalaxis.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAxis(Axis *Axis)
	CreateORMAxisGrid(AxisGrid *AxisGrid)
	CreateORMBezier(Bezier *Bezier)
	CreateORMBezierGrid(BezierGrid *BezierGrid)
	CreateORMBezierGridStack(BezierGridStack *BezierGridStack)
	CreateORMCircle(Circle *Circle)
	CreateORMCircleGrid(CircleGrid *CircleGrid)
	CreateORMExportToMusicxml(ExportToMusicxml *ExportToMusicxml)
	CreateORMFrontCurve(FrontCurve *FrontCurve)
	CreateORMFrontCurveStack(FrontCurveStack *FrontCurveStack)
	CreateORMHorizontalAxis(HorizontalAxis *HorizontalAxis)
	CreateORMKey(Key *Key)
	CreateORMParameter(Parameter *Parameter)
	CreateORMRhombus(Rhombus *Rhombus)
	CreateORMRhombusGrid(RhombusGrid *RhombusGrid)
	CreateORMShapeCategory(ShapeCategory *ShapeCategory)
	CreateORMSpiralBezier(SpiralBezier *SpiralBezier)
	CreateORMSpiralBezierGrid(SpiralBezierGrid *SpiralBezierGrid)
	CreateORMSpiralCircle(SpiralCircle *SpiralCircle)
	CreateORMSpiralCircleGrid(SpiralCircleGrid *SpiralCircleGrid)
	CreateORMSpiralLine(SpiralLine *SpiralLine)
	CreateORMSpiralLineGrid(SpiralLineGrid *SpiralLineGrid)
	CreateORMSpiralOrigin(SpiralOrigin *SpiralOrigin)
	CreateORMSpiralRhombus(SpiralRhombus *SpiralRhombus)
	CreateORMSpiralRhombusGrid(SpiralRhombusGrid *SpiralRhombusGrid)
	CreateORMVerticalAxis(VerticalAxis *VerticalAxis)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAxis(Axis *Axis)
	DeleteORMAxisGrid(AxisGrid *AxisGrid)
	DeleteORMBezier(Bezier *Bezier)
	DeleteORMBezierGrid(BezierGrid *BezierGrid)
	DeleteORMBezierGridStack(BezierGridStack *BezierGridStack)
	DeleteORMCircle(Circle *Circle)
	DeleteORMCircleGrid(CircleGrid *CircleGrid)
	DeleteORMExportToMusicxml(ExportToMusicxml *ExportToMusicxml)
	DeleteORMFrontCurve(FrontCurve *FrontCurve)
	DeleteORMFrontCurveStack(FrontCurveStack *FrontCurveStack)
	DeleteORMHorizontalAxis(HorizontalAxis *HorizontalAxis)
	DeleteORMKey(Key *Key)
	DeleteORMParameter(Parameter *Parameter)
	DeleteORMRhombus(Rhombus *Rhombus)
	DeleteORMRhombusGrid(RhombusGrid *RhombusGrid)
	DeleteORMShapeCategory(ShapeCategory *ShapeCategory)
	DeleteORMSpiralBezier(SpiralBezier *SpiralBezier)
	DeleteORMSpiralBezierGrid(SpiralBezierGrid *SpiralBezierGrid)
	DeleteORMSpiralCircle(SpiralCircle *SpiralCircle)
	DeleteORMSpiralCircleGrid(SpiralCircleGrid *SpiralCircleGrid)
	DeleteORMSpiralLine(SpiralLine *SpiralLine)
	DeleteORMSpiralLineGrid(SpiralLineGrid *SpiralLineGrid)
	DeleteORMSpiralOrigin(SpiralOrigin *SpiralOrigin)
	DeleteORMSpiralRhombus(SpiralRhombus *SpiralRhombus)
	DeleteORMSpiralRhombusGrid(SpiralRhombusGrid *SpiralRhombusGrid)
	DeleteORMVerticalAxis(VerticalAxis *VerticalAxis)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Axiss = make(map[*Axis]any)
	stage.Axiss_mapString = make(map[string]*Axis)
	stage.AxisMap_Staged_Order = make(map[*Axis]uint)
	stage.AxisOrder = 0

	stage.AxisGrids = make(map[*AxisGrid]any)
	stage.AxisGrids_mapString = make(map[string]*AxisGrid)
	stage.AxisGridMap_Staged_Order = make(map[*AxisGrid]uint)
	stage.AxisGridOrder = 0

	stage.Beziers = make(map[*Bezier]any)
	stage.Beziers_mapString = make(map[string]*Bezier)
	stage.BezierMap_Staged_Order = make(map[*Bezier]uint)
	stage.BezierOrder = 0

	stage.BezierGrids = make(map[*BezierGrid]any)
	stage.BezierGrids_mapString = make(map[string]*BezierGrid)
	stage.BezierGridMap_Staged_Order = make(map[*BezierGrid]uint)
	stage.BezierGridOrder = 0

	stage.BezierGridStacks = make(map[*BezierGridStack]any)
	stage.BezierGridStacks_mapString = make(map[string]*BezierGridStack)
	stage.BezierGridStackMap_Staged_Order = make(map[*BezierGridStack]uint)
	stage.BezierGridStackOrder = 0

	stage.Circles = make(map[*Circle]any)
	stage.Circles_mapString = make(map[string]*Circle)
	stage.CircleMap_Staged_Order = make(map[*Circle]uint)
	stage.CircleOrder = 0

	stage.CircleGrids = make(map[*CircleGrid]any)
	stage.CircleGrids_mapString = make(map[string]*CircleGrid)
	stage.CircleGridMap_Staged_Order = make(map[*CircleGrid]uint)
	stage.CircleGridOrder = 0

	stage.ExportToMusicxmls = make(map[*ExportToMusicxml]any)
	stage.ExportToMusicxmls_mapString = make(map[string]*ExportToMusicxml)
	stage.ExportToMusicxmlMap_Staged_Order = make(map[*ExportToMusicxml]uint)
	stage.ExportToMusicxmlOrder = 0

	stage.FrontCurves = make(map[*FrontCurve]any)
	stage.FrontCurves_mapString = make(map[string]*FrontCurve)
	stage.FrontCurveMap_Staged_Order = make(map[*FrontCurve]uint)
	stage.FrontCurveOrder = 0

	stage.FrontCurveStacks = make(map[*FrontCurveStack]any)
	stage.FrontCurveStacks_mapString = make(map[string]*FrontCurveStack)
	stage.FrontCurveStackMap_Staged_Order = make(map[*FrontCurveStack]uint)
	stage.FrontCurveStackOrder = 0

	stage.HorizontalAxiss = make(map[*HorizontalAxis]any)
	stage.HorizontalAxiss_mapString = make(map[string]*HorizontalAxis)
	stage.HorizontalAxisMap_Staged_Order = make(map[*HorizontalAxis]uint)
	stage.HorizontalAxisOrder = 0

	stage.Keys = make(map[*Key]any)
	stage.Keys_mapString = make(map[string]*Key)
	stage.KeyMap_Staged_Order = make(map[*Key]uint)
	stage.KeyOrder = 0

	stage.Parameters = make(map[*Parameter]any)
	stage.Parameters_mapString = make(map[string]*Parameter)
	stage.ParameterMap_Staged_Order = make(map[*Parameter]uint)
	stage.ParameterOrder = 0

	stage.Rhombuss = make(map[*Rhombus]any)
	stage.Rhombuss_mapString = make(map[string]*Rhombus)
	stage.RhombusMap_Staged_Order = make(map[*Rhombus]uint)
	stage.RhombusOrder = 0

	stage.RhombusGrids = make(map[*RhombusGrid]any)
	stage.RhombusGrids_mapString = make(map[string]*RhombusGrid)
	stage.RhombusGridMap_Staged_Order = make(map[*RhombusGrid]uint)
	stage.RhombusGridOrder = 0

	stage.ShapeCategorys = make(map[*ShapeCategory]any)
	stage.ShapeCategorys_mapString = make(map[string]*ShapeCategory)
	stage.ShapeCategoryMap_Staged_Order = make(map[*ShapeCategory]uint)
	stage.ShapeCategoryOrder = 0

	stage.SpiralBeziers = make(map[*SpiralBezier]any)
	stage.SpiralBeziers_mapString = make(map[string]*SpiralBezier)
	stage.SpiralBezierMap_Staged_Order = make(map[*SpiralBezier]uint)
	stage.SpiralBezierOrder = 0

	stage.SpiralBezierGrids = make(map[*SpiralBezierGrid]any)
	stage.SpiralBezierGrids_mapString = make(map[string]*SpiralBezierGrid)
	stage.SpiralBezierGridMap_Staged_Order = make(map[*SpiralBezierGrid]uint)
	stage.SpiralBezierGridOrder = 0

	stage.SpiralCircles = make(map[*SpiralCircle]any)
	stage.SpiralCircles_mapString = make(map[string]*SpiralCircle)
	stage.SpiralCircleMap_Staged_Order = make(map[*SpiralCircle]uint)
	stage.SpiralCircleOrder = 0

	stage.SpiralCircleGrids = make(map[*SpiralCircleGrid]any)
	stage.SpiralCircleGrids_mapString = make(map[string]*SpiralCircleGrid)
	stage.SpiralCircleGridMap_Staged_Order = make(map[*SpiralCircleGrid]uint)
	stage.SpiralCircleGridOrder = 0

	stage.SpiralLines = make(map[*SpiralLine]any)
	stage.SpiralLines_mapString = make(map[string]*SpiralLine)
	stage.SpiralLineMap_Staged_Order = make(map[*SpiralLine]uint)
	stage.SpiralLineOrder = 0

	stage.SpiralLineGrids = make(map[*SpiralLineGrid]any)
	stage.SpiralLineGrids_mapString = make(map[string]*SpiralLineGrid)
	stage.SpiralLineGridMap_Staged_Order = make(map[*SpiralLineGrid]uint)
	stage.SpiralLineGridOrder = 0

	stage.SpiralOrigins = make(map[*SpiralOrigin]any)
	stage.SpiralOrigins_mapString = make(map[string]*SpiralOrigin)
	stage.SpiralOriginMap_Staged_Order = make(map[*SpiralOrigin]uint)
	stage.SpiralOriginOrder = 0

	stage.SpiralRhombuss = make(map[*SpiralRhombus]any)
	stage.SpiralRhombuss_mapString = make(map[string]*SpiralRhombus)
	stage.SpiralRhombusMap_Staged_Order = make(map[*SpiralRhombus]uint)
	stage.SpiralRhombusOrder = 0

	stage.SpiralRhombusGrids = make(map[*SpiralRhombusGrid]any)
	stage.SpiralRhombusGrids_mapString = make(map[string]*SpiralRhombusGrid)
	stage.SpiralRhombusGridMap_Staged_Order = make(map[*SpiralRhombusGrid]uint)
	stage.SpiralRhombusGridOrder = 0

	stage.VerticalAxiss = make(map[*VerticalAxis]any)
	stage.VerticalAxiss_mapString = make(map[string]*VerticalAxis)
	stage.VerticalAxisMap_Staged_Order = make(map[*VerticalAxis]uint)
	stage.VerticalAxisOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Axiss = nil
	stage.Axiss_mapString = nil

	stage.AxisGrids = nil
	stage.AxisGrids_mapString = nil

	stage.Beziers = nil
	stage.Beziers_mapString = nil

	stage.BezierGrids = nil
	stage.BezierGrids_mapString = nil

	stage.BezierGridStacks = nil
	stage.BezierGridStacks_mapString = nil

	stage.Circles = nil
	stage.Circles_mapString = nil

	stage.CircleGrids = nil
	stage.CircleGrids_mapString = nil

	stage.ExportToMusicxmls = nil
	stage.ExportToMusicxmls_mapString = nil

	stage.FrontCurves = nil
	stage.FrontCurves_mapString = nil

	stage.FrontCurveStacks = nil
	stage.FrontCurveStacks_mapString = nil

	stage.HorizontalAxiss = nil
	stage.HorizontalAxiss_mapString = nil

	stage.Keys = nil
	stage.Keys_mapString = nil

	stage.Parameters = nil
	stage.Parameters_mapString = nil

	stage.Rhombuss = nil
	stage.Rhombuss_mapString = nil

	stage.RhombusGrids = nil
	stage.RhombusGrids_mapString = nil

	stage.ShapeCategorys = nil
	stage.ShapeCategorys_mapString = nil

	stage.SpiralBeziers = nil
	stage.SpiralBeziers_mapString = nil

	stage.SpiralBezierGrids = nil
	stage.SpiralBezierGrids_mapString = nil

	stage.SpiralCircles = nil
	stage.SpiralCircles_mapString = nil

	stage.SpiralCircleGrids = nil
	stage.SpiralCircleGrids_mapString = nil

	stage.SpiralLines = nil
	stage.SpiralLines_mapString = nil

	stage.SpiralLineGrids = nil
	stage.SpiralLineGrids_mapString = nil

	stage.SpiralOrigins = nil
	stage.SpiralOrigins_mapString = nil

	stage.SpiralRhombuss = nil
	stage.SpiralRhombuss_mapString = nil

	stage.SpiralRhombusGrids = nil
	stage.SpiralRhombusGrids_mapString = nil

	stage.VerticalAxiss = nil
	stage.VerticalAxiss_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for axis := range stage.Axiss {
		axis.Unstage(stage)
	}

	for axisgrid := range stage.AxisGrids {
		axisgrid.Unstage(stage)
	}

	for bezier := range stage.Beziers {
		bezier.Unstage(stage)
	}

	for beziergrid := range stage.BezierGrids {
		beziergrid.Unstage(stage)
	}

	for beziergridstack := range stage.BezierGridStacks {
		beziergridstack.Unstage(stage)
	}

	for circle := range stage.Circles {
		circle.Unstage(stage)
	}

	for circlegrid := range stage.CircleGrids {
		circlegrid.Unstage(stage)
	}

	for exporttomusicxml := range stage.ExportToMusicxmls {
		exporttomusicxml.Unstage(stage)
	}

	for frontcurve := range stage.FrontCurves {
		frontcurve.Unstage(stage)
	}

	for frontcurvestack := range stage.FrontCurveStacks {
		frontcurvestack.Unstage(stage)
	}

	for horizontalaxis := range stage.HorizontalAxiss {
		horizontalaxis.Unstage(stage)
	}

	for key := range stage.Keys {
		key.Unstage(stage)
	}

	for parameter := range stage.Parameters {
		parameter.Unstage(stage)
	}

	for rhombus := range stage.Rhombuss {
		rhombus.Unstage(stage)
	}

	for rhombusgrid := range stage.RhombusGrids {
		rhombusgrid.Unstage(stage)
	}

	for shapecategory := range stage.ShapeCategorys {
		shapecategory.Unstage(stage)
	}

	for spiralbezier := range stage.SpiralBeziers {
		spiralbezier.Unstage(stage)
	}

	for spiralbeziergrid := range stage.SpiralBezierGrids {
		spiralbeziergrid.Unstage(stage)
	}

	for spiralcircle := range stage.SpiralCircles {
		spiralcircle.Unstage(stage)
	}

	for spiralcirclegrid := range stage.SpiralCircleGrids {
		spiralcirclegrid.Unstage(stage)
	}

	for spiralline := range stage.SpiralLines {
		spiralline.Unstage(stage)
	}

	for spirallinegrid := range stage.SpiralLineGrids {
		spirallinegrid.Unstage(stage)
	}

	for spiralorigin := range stage.SpiralOrigins {
		spiralorigin.Unstage(stage)
	}

	for spiralrhombus := range stage.SpiralRhombuss {
		spiralrhombus.Unstage(stage)
	}

	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		spiralrhombusgrid.Unstage(stage)
	}

	for verticalaxis := range stage.VerticalAxiss {
		verticalaxis.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*Stage)
	UnstageVoid(stage *Stage)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Axis]any:
		return any(&stage.Axiss).(*Type)
	case map[*AxisGrid]any:
		return any(&stage.AxisGrids).(*Type)
	case map[*Bezier]any:
		return any(&stage.Beziers).(*Type)
	case map[*BezierGrid]any:
		return any(&stage.BezierGrids).(*Type)
	case map[*BezierGridStack]any:
		return any(&stage.BezierGridStacks).(*Type)
	case map[*Circle]any:
		return any(&stage.Circles).(*Type)
	case map[*CircleGrid]any:
		return any(&stage.CircleGrids).(*Type)
	case map[*ExportToMusicxml]any:
		return any(&stage.ExportToMusicxmls).(*Type)
	case map[*FrontCurve]any:
		return any(&stage.FrontCurves).(*Type)
	case map[*FrontCurveStack]any:
		return any(&stage.FrontCurveStacks).(*Type)
	case map[*HorizontalAxis]any:
		return any(&stage.HorizontalAxiss).(*Type)
	case map[*Key]any:
		return any(&stage.Keys).(*Type)
	case map[*Parameter]any:
		return any(&stage.Parameters).(*Type)
	case map[*Rhombus]any:
		return any(&stage.Rhombuss).(*Type)
	case map[*RhombusGrid]any:
		return any(&stage.RhombusGrids).(*Type)
	case map[*ShapeCategory]any:
		return any(&stage.ShapeCategorys).(*Type)
	case map[*SpiralBezier]any:
		return any(&stage.SpiralBeziers).(*Type)
	case map[*SpiralBezierGrid]any:
		return any(&stage.SpiralBezierGrids).(*Type)
	case map[*SpiralCircle]any:
		return any(&stage.SpiralCircles).(*Type)
	case map[*SpiralCircleGrid]any:
		return any(&stage.SpiralCircleGrids).(*Type)
	case map[*SpiralLine]any:
		return any(&stage.SpiralLines).(*Type)
	case map[*SpiralLineGrid]any:
		return any(&stage.SpiralLineGrids).(*Type)
	case map[*SpiralOrigin]any:
		return any(&stage.SpiralOrigins).(*Type)
	case map[*SpiralRhombus]any:
		return any(&stage.SpiralRhombuss).(*Type)
	case map[*SpiralRhombusGrid]any:
		return any(&stage.SpiralRhombusGrids).(*Type)
	case map[*VerticalAxis]any:
		return any(&stage.VerticalAxiss).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Axis:
		return any(&stage.Axiss_mapString).(*Type)
	case map[string]*AxisGrid:
		return any(&stage.AxisGrids_mapString).(*Type)
	case map[string]*Bezier:
		return any(&stage.Beziers_mapString).(*Type)
	case map[string]*BezierGrid:
		return any(&stage.BezierGrids_mapString).(*Type)
	case map[string]*BezierGridStack:
		return any(&stage.BezierGridStacks_mapString).(*Type)
	case map[string]*Circle:
		return any(&stage.Circles_mapString).(*Type)
	case map[string]*CircleGrid:
		return any(&stage.CircleGrids_mapString).(*Type)
	case map[string]*ExportToMusicxml:
		return any(&stage.ExportToMusicxmls_mapString).(*Type)
	case map[string]*FrontCurve:
		return any(&stage.FrontCurves_mapString).(*Type)
	case map[string]*FrontCurveStack:
		return any(&stage.FrontCurveStacks_mapString).(*Type)
	case map[string]*HorizontalAxis:
		return any(&stage.HorizontalAxiss_mapString).(*Type)
	case map[string]*Key:
		return any(&stage.Keys_mapString).(*Type)
	case map[string]*Parameter:
		return any(&stage.Parameters_mapString).(*Type)
	case map[string]*Rhombus:
		return any(&stage.Rhombuss_mapString).(*Type)
	case map[string]*RhombusGrid:
		return any(&stage.RhombusGrids_mapString).(*Type)
	case map[string]*ShapeCategory:
		return any(&stage.ShapeCategorys_mapString).(*Type)
	case map[string]*SpiralBezier:
		return any(&stage.SpiralBeziers_mapString).(*Type)
	case map[string]*SpiralBezierGrid:
		return any(&stage.SpiralBezierGrids_mapString).(*Type)
	case map[string]*SpiralCircle:
		return any(&stage.SpiralCircles_mapString).(*Type)
	case map[string]*SpiralCircleGrid:
		return any(&stage.SpiralCircleGrids_mapString).(*Type)
	case map[string]*SpiralLine:
		return any(&stage.SpiralLines_mapString).(*Type)
	case map[string]*SpiralLineGrid:
		return any(&stage.SpiralLineGrids_mapString).(*Type)
	case map[string]*SpiralOrigin:
		return any(&stage.SpiralOrigins_mapString).(*Type)
	case map[string]*SpiralRhombus:
		return any(&stage.SpiralRhombuss_mapString).(*Type)
	case map[string]*SpiralRhombusGrid:
		return any(&stage.SpiralRhombusGrids_mapString).(*Type)
	case map[string]*VerticalAxis:
		return any(&stage.VerticalAxiss_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Axis:
		return any(&stage.Axiss).(*map[*Type]any)
	case AxisGrid:
		return any(&stage.AxisGrids).(*map[*Type]any)
	case Bezier:
		return any(&stage.Beziers).(*map[*Type]any)
	case BezierGrid:
		return any(&stage.BezierGrids).(*map[*Type]any)
	case BezierGridStack:
		return any(&stage.BezierGridStacks).(*map[*Type]any)
	case Circle:
		return any(&stage.Circles).(*map[*Type]any)
	case CircleGrid:
		return any(&stage.CircleGrids).(*map[*Type]any)
	case ExportToMusicxml:
		return any(&stage.ExportToMusicxmls).(*map[*Type]any)
	case FrontCurve:
		return any(&stage.FrontCurves).(*map[*Type]any)
	case FrontCurveStack:
		return any(&stage.FrontCurveStacks).(*map[*Type]any)
	case HorizontalAxis:
		return any(&stage.HorizontalAxiss).(*map[*Type]any)
	case Key:
		return any(&stage.Keys).(*map[*Type]any)
	case Parameter:
		return any(&stage.Parameters).(*map[*Type]any)
	case Rhombus:
		return any(&stage.Rhombuss).(*map[*Type]any)
	case RhombusGrid:
		return any(&stage.RhombusGrids).(*map[*Type]any)
	case ShapeCategory:
		return any(&stage.ShapeCategorys).(*map[*Type]any)
	case SpiralBezier:
		return any(&stage.SpiralBeziers).(*map[*Type]any)
	case SpiralBezierGrid:
		return any(&stage.SpiralBezierGrids).(*map[*Type]any)
	case SpiralCircle:
		return any(&stage.SpiralCircles).(*map[*Type]any)
	case SpiralCircleGrid:
		return any(&stage.SpiralCircleGrids).(*map[*Type]any)
	case SpiralLine:
		return any(&stage.SpiralLines).(*map[*Type]any)
	case SpiralLineGrid:
		return any(&stage.SpiralLineGrids).(*map[*Type]any)
	case SpiralOrigin:
		return any(&stage.SpiralOrigins).(*map[*Type]any)
	case SpiralRhombus:
		return any(&stage.SpiralRhombuss).(*map[*Type]any)
	case SpiralRhombusGrid:
		return any(&stage.SpiralRhombusGrids).(*map[*Type]any)
	case VerticalAxis:
		return any(&stage.VerticalAxiss).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Axis:
		return any(&stage.Axiss).(*map[Type]any)
	case *AxisGrid:
		return any(&stage.AxisGrids).(*map[Type]any)
	case *Bezier:
		return any(&stage.Beziers).(*map[Type]any)
	case *BezierGrid:
		return any(&stage.BezierGrids).(*map[Type]any)
	case *BezierGridStack:
		return any(&stage.BezierGridStacks).(*map[Type]any)
	case *Circle:
		return any(&stage.Circles).(*map[Type]any)
	case *CircleGrid:
		return any(&stage.CircleGrids).(*map[Type]any)
	case *ExportToMusicxml:
		return any(&stage.ExportToMusicxmls).(*map[Type]any)
	case *FrontCurve:
		return any(&stage.FrontCurves).(*map[Type]any)
	case *FrontCurveStack:
		return any(&stage.FrontCurveStacks).(*map[Type]any)
	case *HorizontalAxis:
		return any(&stage.HorizontalAxiss).(*map[Type]any)
	case *Key:
		return any(&stage.Keys).(*map[Type]any)
	case *Parameter:
		return any(&stage.Parameters).(*map[Type]any)
	case *Rhombus:
		return any(&stage.Rhombuss).(*map[Type]any)
	case *RhombusGrid:
		return any(&stage.RhombusGrids).(*map[Type]any)
	case *ShapeCategory:
		return any(&stage.ShapeCategorys).(*map[Type]any)
	case *SpiralBezier:
		return any(&stage.SpiralBeziers).(*map[Type]any)
	case *SpiralBezierGrid:
		return any(&stage.SpiralBezierGrids).(*map[Type]any)
	case *SpiralCircle:
		return any(&stage.SpiralCircles).(*map[Type]any)
	case *SpiralCircleGrid:
		return any(&stage.SpiralCircleGrids).(*map[Type]any)
	case *SpiralLine:
		return any(&stage.SpiralLines).(*map[Type]any)
	case *SpiralLineGrid:
		return any(&stage.SpiralLineGrids).(*map[Type]any)
	case *SpiralOrigin:
		return any(&stage.SpiralOrigins).(*map[Type]any)
	case *SpiralRhombus:
		return any(&stage.SpiralRhombuss).(*map[Type]any)
	case *SpiralRhombusGrid:
		return any(&stage.SpiralRhombusGrids).(*map[Type]any)
	case *VerticalAxis:
		return any(&stage.VerticalAxiss).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Axis:
		return any(&stage.Axiss_mapString).(*map[string]*Type)
	case AxisGrid:
		return any(&stage.AxisGrids_mapString).(*map[string]*Type)
	case Bezier:
		return any(&stage.Beziers_mapString).(*map[string]*Type)
	case BezierGrid:
		return any(&stage.BezierGrids_mapString).(*map[string]*Type)
	case BezierGridStack:
		return any(&stage.BezierGridStacks_mapString).(*map[string]*Type)
	case Circle:
		return any(&stage.Circles_mapString).(*map[string]*Type)
	case CircleGrid:
		return any(&stage.CircleGrids_mapString).(*map[string]*Type)
	case ExportToMusicxml:
		return any(&stage.ExportToMusicxmls_mapString).(*map[string]*Type)
	case FrontCurve:
		return any(&stage.FrontCurves_mapString).(*map[string]*Type)
	case FrontCurveStack:
		return any(&stage.FrontCurveStacks_mapString).(*map[string]*Type)
	case HorizontalAxis:
		return any(&stage.HorizontalAxiss_mapString).(*map[string]*Type)
	case Key:
		return any(&stage.Keys_mapString).(*map[string]*Type)
	case Parameter:
		return any(&stage.Parameters_mapString).(*map[string]*Type)
	case Rhombus:
		return any(&stage.Rhombuss_mapString).(*map[string]*Type)
	case RhombusGrid:
		return any(&stage.RhombusGrids_mapString).(*map[string]*Type)
	case ShapeCategory:
		return any(&stage.ShapeCategorys_mapString).(*map[string]*Type)
	case SpiralBezier:
		return any(&stage.SpiralBeziers_mapString).(*map[string]*Type)
	case SpiralBezierGrid:
		return any(&stage.SpiralBezierGrids_mapString).(*map[string]*Type)
	case SpiralCircle:
		return any(&stage.SpiralCircles_mapString).(*map[string]*Type)
	case SpiralCircleGrid:
		return any(&stage.SpiralCircleGrids_mapString).(*map[string]*Type)
	case SpiralLine:
		return any(&stage.SpiralLines_mapString).(*map[string]*Type)
	case SpiralLineGrid:
		return any(&stage.SpiralLineGrids_mapString).(*map[string]*Type)
	case SpiralOrigin:
		return any(&stage.SpiralOrigins_mapString).(*map[string]*Type)
	case SpiralRhombus:
		return any(&stage.SpiralRhombuss_mapString).(*map[string]*Type)
	case SpiralRhombusGrid:
		return any(&stage.SpiralRhombusGrids_mapString).(*map[string]*Type)
	case VerticalAxis:
		return any(&stage.VerticalAxiss_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Axis:
		return any(&Axis{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case AxisGrid:
		return any(&AxisGrid{
			// Initialisation of associations
			// field is initialized with an instance of Axis with the name of the field
			Reference: &Axis{Name: "Reference"},
			// field is initialized with an instance of Axis with the name of the field
			Axiss: []*Axis{{Name: "Axiss"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case Bezier:
		return any(&Bezier{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case BezierGrid:
		return any(&BezierGrid{
			// Initialisation of associations
			// field is initialized with an instance of Bezier with the name of the field
			Reference: &Bezier{Name: "Reference"},
			// field is initialized with an instance of Bezier with the name of the field
			Beziers: []*Bezier{{Name: "Beziers"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case BezierGridStack:
		return any(&BezierGridStack{
			// Initialisation of associations
			// field is initialized with an instance of BezierGrid with the name of the field
			BezierGrids: []*BezierGrid{{Name: "BezierGrids"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case Circle:
		return any(&Circle{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case CircleGrid:
		return any(&CircleGrid{
			// Initialisation of associations
			// field is initialized with an instance of Circle with the name of the field
			Reference: &Circle{Name: "Reference"},
			// field is initialized with an instance of Circle with the name of the field
			Circles: []*Circle{{Name: "Circles"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case ExportToMusicxml:
		return any(&ExportToMusicxml{
			// Initialisation of associations
			// field is initialized with an instance of Parameter with the name of the field
			Parameter: &Parameter{Name: "Parameter"},
		}).(*Type)
	case FrontCurve:
		return any(&FrontCurve{
			// Initialisation of associations
		}).(*Type)
	case FrontCurveStack:
		return any(&FrontCurveStack{
			// Initialisation of associations
			// field is initialized with an instance of FrontCurve with the name of the field
			FrontCurves: []*FrontCurve{{Name: "FrontCurves"}},
			// field is initialized with an instance of SpiralCircle with the name of the field
			SpiralCircles: []*SpiralCircle{{Circle: Circle{Name: "SpiralCircles"}}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case HorizontalAxis:
		return any(&HorizontalAxis{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case Key:
		return any(&Key{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case Parameter:
		return any(&Parameter{
			// Initialisation of associations
			// field is initialized with an instance of Rhombus with the name of the field
			InitialRhombus: &Rhombus{Name: "InitialRhombus"},
			// field is initialized with an instance of Circle with the name of the field
			InitialCircle: &Circle{Name: "InitialCircle"},
			// field is initialized with an instance of RhombusGrid with the name of the field
			InitialRhombusGrid: &RhombusGrid{Name: "InitialRhombusGrid"},
			// field is initialized with an instance of CircleGrid with the name of the field
			InitialCircleGrid: &CircleGrid{Name: "InitialCircleGrid"},
			// field is initialized with an instance of Axis with the name of the field
			InitialAxis: &Axis{Name: "InitialAxis"},
			// field is initialized with an instance of Axis with the name of the field
			RotatedAxis: &Axis{Name: "RotatedAxis"},
			// field is initialized with an instance of Rhombus with the name of the field
			RotatedRhombus: &Rhombus{Name: "RotatedRhombus"},
			// field is initialized with an instance of RhombusGrid with the name of the field
			RotatedRhombusGrid: &RhombusGrid{Name: "RotatedRhombusGrid"},
			// field is initialized with an instance of CircleGrid with the name of the field
			RotatedCircleGrid: &CircleGrid{Name: "RotatedCircleGrid"},
			// field is initialized with an instance of Rhombus with the name of the field
			NextRhombus: &Rhombus{Name: "NextRhombus"},
			// field is initialized with an instance of Circle with the name of the field
			NextCircle: &Circle{Name: "NextCircle"},
			// field is initialized with an instance of Rhombus with the name of the field
			GrowingRhombusGridSeed: &Rhombus{Name: "GrowingRhombusGridSeed"},
			// field is initialized with an instance of RhombusGrid with the name of the field
			GrowingRhombusGrid: &RhombusGrid{Name: "GrowingRhombusGrid"},
			// field is initialized with an instance of Circle with the name of the field
			GrowingCircleGridSeed: &Circle{Name: "GrowingCircleGridSeed"},
			// field is initialized with an instance of CircleGrid with the name of the field
			GrowingCircleGrid: &CircleGrid{Name: "GrowingCircleGrid"},
			// field is initialized with an instance of Circle with the name of the field
			GrowingCircleGridLeftSeed: &Circle{Name: "GrowingCircleGridLeftSeed"},
			// field is initialized with an instance of CircleGrid with the name of the field
			GrowingCircleGridLeft: &CircleGrid{Name: "GrowingCircleGridLeft"},
			// field is initialized with an instance of Axis with the name of the field
			ConstructionAxis: &Axis{Name: "ConstructionAxis"},
			// field is initialized with an instance of AxisGrid with the name of the field
			ConstructionAxisGrid: &AxisGrid{Name: "ConstructionAxisGrid"},
			// field is initialized with an instance of Circle with the name of the field
			ConstructionCircle: &Circle{Name: "ConstructionCircle"},
			// field is initialized with an instance of CircleGrid with the name of the field
			ConstructionCircleGrid: &CircleGrid{Name: "ConstructionCircleGrid"},
			// field is initialized with an instance of Bezier with the name of the field
			GrowthCurveSeed: &Bezier{Name: "GrowthCurveSeed"},
			// field is initialized with an instance of BezierGrid with the name of the field
			GrowthCurve: &BezierGrid{Name: "GrowthCurve"},
			// field is initialized with an instance of Bezier with the name of the field
			GrowthCurveShiftedRightSeed: &Bezier{Name: "GrowthCurveShiftedRightSeed"},
			// field is initialized with an instance of BezierGrid with the name of the field
			GrowthCurveShiftedRight: &BezierGrid{Name: "GrowthCurveShiftedRight"},
			// field is initialized with an instance of Bezier with the name of the field
			GrowthCurveNextSeed: &Bezier{Name: "GrowthCurveNextSeed"},
			// field is initialized with an instance of BezierGrid with the name of the field
			GrowthCurveNext: &BezierGrid{Name: "GrowthCurveNext"},
			// field is initialized with an instance of Bezier with the name of the field
			GrowthCurveNextShiftedRightSeed: &Bezier{Name: "GrowthCurveNextShiftedRightSeed"},
			// field is initialized with an instance of BezierGrid with the name of the field
			GrowthCurveNextShiftedRight: &BezierGrid{Name: "GrowthCurveNextShiftedRight"},
			// field is initialized with an instance of BezierGridStack with the name of the field
			GrowthCurveStack: &BezierGridStack{Name: "GrowthCurveStack"},
			// field is initialized with an instance of SpiralRhombus with the name of the field
			SpiralRhombusGridSeed: &SpiralRhombus{Name: "SpiralRhombusGridSeed"},
			// field is initialized with an instance of SpiralRhombusGrid with the name of the field
			SpiralRhombusGrid: &SpiralRhombusGrid{Name: "SpiralRhombusGrid"},
			// field is initialized with an instance of SpiralCircle with the name of the field
			SpiralCircleSeed: &SpiralCircle{Circle: Circle{Name: "SpiralCircleSeed"}},
			// field is initialized with an instance of SpiralCircleGrid with the name of the field
			SpiralCircleGrid: &SpiralCircleGrid{Name: "SpiralCircleGrid"},
			// field is initialized with an instance of SpiralCircleGrid with the name of the field
			SpiralCircleFullGrid: &SpiralCircleGrid{Name: "SpiralCircleFullGrid"},
			// field is initialized with an instance of SpiralLine with the name of the field
			SpiralConstructionOuterLineSeed: &SpiralLine{Name: "SpiralConstructionOuterLineSeed"},
			// field is initialized with an instance of SpiralLine with the name of the field
			SpiralConstructionInnerLineSeed: &SpiralLine{Name: "SpiralConstructionInnerLineSeed"},
			// field is initialized with an instance of SpiralLineGrid with the name of the field
			SpiralConstructionOuterLineGrid: &SpiralLineGrid{Name: "SpiralConstructionOuterLineGrid"},
			// field is initialized with an instance of SpiralLineGrid with the name of the field
			SpiralConstructionInnerLineGrid: &SpiralLineGrid{Name: "SpiralConstructionInnerLineGrid"},
			// field is initialized with an instance of SpiralCircleGrid with the name of the field
			SpiralConstructionCircleGrid: &SpiralCircleGrid{Name: "SpiralConstructionCircleGrid"},
			// field is initialized with an instance of SpiralLineGrid with the name of the field
			SpiralConstructionOuterLineFullGrid: &SpiralLineGrid{Name: "SpiralConstructionOuterLineFullGrid"},
			// field is initialized with an instance of SpiralBezier with the name of the field
			SpiralBezierSeed: &SpiralBezier{Name: "SpiralBezierSeed"},
			// field is initialized with an instance of SpiralBezierGrid with the name of the field
			SpiralBezierGrid: &SpiralBezierGrid{Name: "SpiralBezierGrid"},
			// field is initialized with an instance of SpiralBezierGrid with the name of the field
			SpiralBezierFullGrid: &SpiralBezierGrid{Name: "SpiralBezierFullGrid"},
			// field is initialized with an instance of FrontCurveStack with the name of the field
			FrontCurveStack: &FrontCurveStack{Name: "FrontCurveStack"},
			// field is initialized with an instance of Key with the name of the field
			Fkey: &Key{Name: "Fkey"},
			// field is initialized with an instance of AxisGrid with the name of the field
			PitchLines: &AxisGrid{Name: "PitchLines"},
			// field is initialized with an instance of AxisGrid with the name of the field
			BeatLines: &AxisGrid{Name: "BeatLines"},
			// field is initialized with an instance of BezierGrid with the name of the field
			FirstVoice: &BezierGrid{Name: "FirstVoice"},
			// field is initialized with an instance of BezierGrid with the name of the field
			FirstVoiceShiftedRigth: &BezierGrid{Name: "FirstVoiceShiftedRigth"},
			// field is initialized with an instance of BezierGrid with the name of the field
			SecondVoice: &BezierGrid{Name: "SecondVoice"},
			// field is initialized with an instance of BezierGrid with the name of the field
			SecondVoiceShiftedRight: &BezierGrid{Name: "SecondVoiceShiftedRight"},
			// field is initialized with an instance of CircleGrid with the name of the field
			FirstVoiceNotes: &CircleGrid{Name: "FirstVoiceNotes"},
			// field is initialized with an instance of CircleGrid with the name of the field
			FirstVoiceNotesShiftedRight: &CircleGrid{Name: "FirstVoiceNotesShiftedRight"},
			// field is initialized with an instance of CircleGrid with the name of the field
			SecondVoiceNotes: &CircleGrid{Name: "SecondVoiceNotes"},
			// field is initialized with an instance of CircleGrid with the name of the field
			SecondVoiceNotesShiftedRight: &CircleGrid{Name: "SecondVoiceNotesShiftedRight"},
			// field is initialized with an instance of HorizontalAxis with the name of the field
			HorizontalAxis: &HorizontalAxis{Name: "HorizontalAxis"},
			// field is initialized with an instance of VerticalAxis with the name of the field
			VerticalAxis: &VerticalAxis{Name: "VerticalAxis"},
			// field is initialized with an instance of SpiralOrigin with the name of the field
			SpiralOrigin: &SpiralOrigin{Name: "SpiralOrigin"},
		}).(*Type)
	case Rhombus:
		return any(&Rhombus{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case RhombusGrid:
		return any(&RhombusGrid{
			// Initialisation of associations
			// field is initialized with an instance of Rhombus with the name of the field
			Reference: &Rhombus{Name: "Reference"},
			// field is initialized with an instance of Rhombus with the name of the field
			Rhombuses: []*Rhombus{{Name: "Rhombuses"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case ShapeCategory:
		return any(&ShapeCategory{
			// Initialisation of associations
		}).(*Type)
	case SpiralBezier:
		return any(&SpiralBezier{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralBezierGrid:
		return any(&SpiralBezierGrid{
			// Initialisation of associations
			// field is initialized with an instance of SpiralBezier with the name of the field
			SpiralBeziers: []*SpiralBezier{{Name: "SpiralBeziers"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralCircle:
		return any(&SpiralCircle{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralCircleGrid:
		return any(&SpiralCircleGrid{
			// Initialisation of associations
			// field is initialized with an instance of SpiralRhombusGrid with the name of the field
			SpiralRhombusGrid: &SpiralRhombusGrid{Name: "SpiralRhombusGrid"},
			// field is initialized with an instance of SpiralCircle with the name of the field
			SpiralCircles: []*SpiralCircle{{Circle: Circle{Name: "SpiralCircles"}}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralLine:
		return any(&SpiralLine{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralLineGrid:
		return any(&SpiralLineGrid{
			// Initialisation of associations
			// field is initialized with an instance of SpiralLine with the name of the field
			SpiralLines: []*SpiralLine{{Name: "SpiralLines"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralOrigin:
		return any(&SpiralOrigin{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralRhombus:
		return any(&SpiralRhombus{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	case SpiralRhombusGrid:
		return any(&SpiralRhombusGrid{
			// Initialisation of associations
			// field is initialized with an instance of SpiralRhombus with the name of the field
			SpiralRhombuses: []*SpiralRhombus{{Name: "SpiralRhombuses"}},
			// field is initialized with Shape problem with composites

		}).(*Type)
	case VerticalAxis:
		return any(&VerticalAxis{
			// Initialisation of associations
			// field is initialized with Shape problem with composites

		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Axis
	case Axis:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*Axis)
			for axis := range stage.Axiss {
				if axis.ShapeCategory != nil {
					shapecategory_ := axis.ShapeCategory
					var axiss []*Axis
					_, ok := res[shapecategory_]
					if ok {
						axiss = res[shapecategory_]
					} else {
						axiss = make([]*Axis, 0)
					}
					axiss = append(axiss, axis)
					res[shapecategory_] = axiss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AxisGrid
	case AxisGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Reference":
			res := make(map[*Axis][]*AxisGrid)
			for axisgrid := range stage.AxisGrids {
				if axisgrid.Reference != nil {
					axis_ := axisgrid.Reference
					var axisgrids []*AxisGrid
					_, ok := res[axis_]
					if ok {
						axisgrids = res[axis_]
					} else {
						axisgrids = make([]*AxisGrid, 0)
					}
					axisgrids = append(axisgrids, axisgrid)
					res[axis_] = axisgrids
				}
			}
			return any(res).(map[*End][]*Start)
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*AxisGrid)
			for axisgrid := range stage.AxisGrids {
				if axisgrid.ShapeCategory != nil {
					shapecategory_ := axisgrid.ShapeCategory
					var axisgrids []*AxisGrid
					_, ok := res[shapecategory_]
					if ok {
						axisgrids = res[shapecategory_]
					} else {
						axisgrids = make([]*AxisGrid, 0)
					}
					axisgrids = append(axisgrids, axisgrid)
					res[shapecategory_] = axisgrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Bezier
	case Bezier:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*Bezier)
			for bezier := range stage.Beziers {
				if bezier.ShapeCategory != nil {
					shapecategory_ := bezier.ShapeCategory
					var beziers []*Bezier
					_, ok := res[shapecategory_]
					if ok {
						beziers = res[shapecategory_]
					} else {
						beziers = make([]*Bezier, 0)
					}
					beziers = append(beziers, bezier)
					res[shapecategory_] = beziers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BezierGrid
	case BezierGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Reference":
			res := make(map[*Bezier][]*BezierGrid)
			for beziergrid := range stage.BezierGrids {
				if beziergrid.Reference != nil {
					bezier_ := beziergrid.Reference
					var beziergrids []*BezierGrid
					_, ok := res[bezier_]
					if ok {
						beziergrids = res[bezier_]
					} else {
						beziergrids = make([]*BezierGrid, 0)
					}
					beziergrids = append(beziergrids, beziergrid)
					res[bezier_] = beziergrids
				}
			}
			return any(res).(map[*End][]*Start)
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*BezierGrid)
			for beziergrid := range stage.BezierGrids {
				if beziergrid.ShapeCategory != nil {
					shapecategory_ := beziergrid.ShapeCategory
					var beziergrids []*BezierGrid
					_, ok := res[shapecategory_]
					if ok {
						beziergrids = res[shapecategory_]
					} else {
						beziergrids = make([]*BezierGrid, 0)
					}
					beziergrids = append(beziergrids, beziergrid)
					res[shapecategory_] = beziergrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BezierGridStack
	case BezierGridStack:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*BezierGridStack)
			for beziergridstack := range stage.BezierGridStacks {
				if beziergridstack.ShapeCategory != nil {
					shapecategory_ := beziergridstack.ShapeCategory
					var beziergridstacks []*BezierGridStack
					_, ok := res[shapecategory_]
					if ok {
						beziergridstacks = res[shapecategory_]
					} else {
						beziergridstacks = make([]*BezierGridStack, 0)
					}
					beziergridstacks = append(beziergridstacks, beziergridstack)
					res[shapecategory_] = beziergridstacks
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Circle
	case Circle:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*Circle)
			for circle := range stage.Circles {
				if circle.ShapeCategory != nil {
					shapecategory_ := circle.ShapeCategory
					var circles []*Circle
					_, ok := res[shapecategory_]
					if ok {
						circles = res[shapecategory_]
					} else {
						circles = make([]*Circle, 0)
					}
					circles = append(circles, circle)
					res[shapecategory_] = circles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CircleGrid
	case CircleGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Reference":
			res := make(map[*Circle][]*CircleGrid)
			for circlegrid := range stage.CircleGrids {
				if circlegrid.Reference != nil {
					circle_ := circlegrid.Reference
					var circlegrids []*CircleGrid
					_, ok := res[circle_]
					if ok {
						circlegrids = res[circle_]
					} else {
						circlegrids = make([]*CircleGrid, 0)
					}
					circlegrids = append(circlegrids, circlegrid)
					res[circle_] = circlegrids
				}
			}
			return any(res).(map[*End][]*Start)
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*CircleGrid)
			for circlegrid := range stage.CircleGrids {
				if circlegrid.ShapeCategory != nil {
					shapecategory_ := circlegrid.ShapeCategory
					var circlegrids []*CircleGrid
					_, ok := res[shapecategory_]
					if ok {
						circlegrids = res[shapecategory_]
					} else {
						circlegrids = make([]*CircleGrid, 0)
					}
					circlegrids = append(circlegrids, circlegrid)
					res[shapecategory_] = circlegrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ExportToMusicxml
	case ExportToMusicxml:
		switch fieldname {
		// insertion point for per direct association field
		case "Parameter":
			res := make(map[*Parameter][]*ExportToMusicxml)
			for exporttomusicxml := range stage.ExportToMusicxmls {
				if exporttomusicxml.Parameter != nil {
					parameter_ := exporttomusicxml.Parameter
					var exporttomusicxmls []*ExportToMusicxml
					_, ok := res[parameter_]
					if ok {
						exporttomusicxmls = res[parameter_]
					} else {
						exporttomusicxmls = make([]*ExportToMusicxml, 0)
					}
					exporttomusicxmls = append(exporttomusicxmls, exporttomusicxml)
					res[parameter_] = exporttomusicxmls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FrontCurve
	case FrontCurve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FrontCurveStack
	case FrontCurveStack:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*FrontCurveStack)
			for frontcurvestack := range stage.FrontCurveStacks {
				if frontcurvestack.ShapeCategory != nil {
					shapecategory_ := frontcurvestack.ShapeCategory
					var frontcurvestacks []*FrontCurveStack
					_, ok := res[shapecategory_]
					if ok {
						frontcurvestacks = res[shapecategory_]
					} else {
						frontcurvestacks = make([]*FrontCurveStack, 0)
					}
					frontcurvestacks = append(frontcurvestacks, frontcurvestack)
					res[shapecategory_] = frontcurvestacks
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of HorizontalAxis
	case HorizontalAxis:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*HorizontalAxis)
			for horizontalaxis := range stage.HorizontalAxiss {
				if horizontalaxis.ShapeCategory != nil {
					shapecategory_ := horizontalaxis.ShapeCategory
					var horizontalaxiss []*HorizontalAxis
					_, ok := res[shapecategory_]
					if ok {
						horizontalaxiss = res[shapecategory_]
					} else {
						horizontalaxiss = make([]*HorizontalAxis, 0)
					}
					horizontalaxiss = append(horizontalaxiss, horizontalaxis)
					res[shapecategory_] = horizontalaxiss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Key
	case Key:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*Key)
			for key := range stage.Keys {
				if key.ShapeCategory != nil {
					shapecategory_ := key.ShapeCategory
					var keys []*Key
					_, ok := res[shapecategory_]
					if ok {
						keys = res[shapecategory_]
					} else {
						keys = make([]*Key, 0)
					}
					keys = append(keys, key)
					res[shapecategory_] = keys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Parameter
	case Parameter:
		switch fieldname {
		// insertion point for per direct association field
		case "InitialRhombus":
			res := make(map[*Rhombus][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.InitialRhombus != nil {
					rhombus_ := parameter.InitialRhombus
					var parameters []*Parameter
					_, ok := res[rhombus_]
					if ok {
						parameters = res[rhombus_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[rhombus_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "InitialCircle":
			res := make(map[*Circle][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.InitialCircle != nil {
					circle_ := parameter.InitialCircle
					var parameters []*Parameter
					_, ok := res[circle_]
					if ok {
						parameters = res[circle_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circle_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "InitialRhombusGrid":
			res := make(map[*RhombusGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.InitialRhombusGrid != nil {
					rhombusgrid_ := parameter.InitialRhombusGrid
					var parameters []*Parameter
					_, ok := res[rhombusgrid_]
					if ok {
						parameters = res[rhombusgrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[rhombusgrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "InitialCircleGrid":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.InitialCircleGrid != nil {
					circlegrid_ := parameter.InitialCircleGrid
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "InitialAxis":
			res := make(map[*Axis][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.InitialAxis != nil {
					axis_ := parameter.InitialAxis
					var parameters []*Parameter
					_, ok := res[axis_]
					if ok {
						parameters = res[axis_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[axis_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedAxis":
			res := make(map[*Axis][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.RotatedAxis != nil {
					axis_ := parameter.RotatedAxis
					var parameters []*Parameter
					_, ok := res[axis_]
					if ok {
						parameters = res[axis_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[axis_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedRhombus":
			res := make(map[*Rhombus][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.RotatedRhombus != nil {
					rhombus_ := parameter.RotatedRhombus
					var parameters []*Parameter
					_, ok := res[rhombus_]
					if ok {
						parameters = res[rhombus_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[rhombus_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedRhombusGrid":
			res := make(map[*RhombusGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.RotatedRhombusGrid != nil {
					rhombusgrid_ := parameter.RotatedRhombusGrid
					var parameters []*Parameter
					_, ok := res[rhombusgrid_]
					if ok {
						parameters = res[rhombusgrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[rhombusgrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "RotatedCircleGrid":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.RotatedCircleGrid != nil {
					circlegrid_ := parameter.RotatedCircleGrid
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "NextRhombus":
			res := make(map[*Rhombus][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.NextRhombus != nil {
					rhombus_ := parameter.NextRhombus
					var parameters []*Parameter
					_, ok := res[rhombus_]
					if ok {
						parameters = res[rhombus_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[rhombus_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "NextCircle":
			res := make(map[*Circle][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.NextCircle != nil {
					circle_ := parameter.NextCircle
					var parameters []*Parameter
					_, ok := res[circle_]
					if ok {
						parameters = res[circle_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circle_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowingRhombusGridSeed":
			res := make(map[*Rhombus][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowingRhombusGridSeed != nil {
					rhombus_ := parameter.GrowingRhombusGridSeed
					var parameters []*Parameter
					_, ok := res[rhombus_]
					if ok {
						parameters = res[rhombus_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[rhombus_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowingRhombusGrid":
			res := make(map[*RhombusGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowingRhombusGrid != nil {
					rhombusgrid_ := parameter.GrowingRhombusGrid
					var parameters []*Parameter
					_, ok := res[rhombusgrid_]
					if ok {
						parameters = res[rhombusgrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[rhombusgrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowingCircleGridSeed":
			res := make(map[*Circle][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowingCircleGridSeed != nil {
					circle_ := parameter.GrowingCircleGridSeed
					var parameters []*Parameter
					_, ok := res[circle_]
					if ok {
						parameters = res[circle_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circle_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowingCircleGrid":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowingCircleGrid != nil {
					circlegrid_ := parameter.GrowingCircleGrid
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowingCircleGridLeftSeed":
			res := make(map[*Circle][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowingCircleGridLeftSeed != nil {
					circle_ := parameter.GrowingCircleGridLeftSeed
					var parameters []*Parameter
					_, ok := res[circle_]
					if ok {
						parameters = res[circle_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circle_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowingCircleGridLeft":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowingCircleGridLeft != nil {
					circlegrid_ := parameter.GrowingCircleGridLeft
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConstructionAxis":
			res := make(map[*Axis][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.ConstructionAxis != nil {
					axis_ := parameter.ConstructionAxis
					var parameters []*Parameter
					_, ok := res[axis_]
					if ok {
						parameters = res[axis_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[axis_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConstructionAxisGrid":
			res := make(map[*AxisGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.ConstructionAxisGrid != nil {
					axisgrid_ := parameter.ConstructionAxisGrid
					var parameters []*Parameter
					_, ok := res[axisgrid_]
					if ok {
						parameters = res[axisgrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[axisgrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConstructionCircle":
			res := make(map[*Circle][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.ConstructionCircle != nil {
					circle_ := parameter.ConstructionCircle
					var parameters []*Parameter
					_, ok := res[circle_]
					if ok {
						parameters = res[circle_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circle_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConstructionCircleGrid":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.ConstructionCircleGrid != nil {
					circlegrid_ := parameter.ConstructionCircleGrid
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveSeed":
			res := make(map[*Bezier][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveSeed != nil {
					bezier_ := parameter.GrowthCurveSeed
					var parameters []*Parameter
					_, ok := res[bezier_]
					if ok {
						parameters = res[bezier_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[bezier_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurve":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurve != nil {
					beziergrid_ := parameter.GrowthCurve
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveShiftedRightSeed":
			res := make(map[*Bezier][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveShiftedRightSeed != nil {
					bezier_ := parameter.GrowthCurveShiftedRightSeed
					var parameters []*Parameter
					_, ok := res[bezier_]
					if ok {
						parameters = res[bezier_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[bezier_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveShiftedRight":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveShiftedRight != nil {
					beziergrid_ := parameter.GrowthCurveShiftedRight
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveNextSeed":
			res := make(map[*Bezier][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveNextSeed != nil {
					bezier_ := parameter.GrowthCurveNextSeed
					var parameters []*Parameter
					_, ok := res[bezier_]
					if ok {
						parameters = res[bezier_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[bezier_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveNext":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveNext != nil {
					beziergrid_ := parameter.GrowthCurveNext
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveNextShiftedRightSeed":
			res := make(map[*Bezier][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveNextShiftedRightSeed != nil {
					bezier_ := parameter.GrowthCurveNextShiftedRightSeed
					var parameters []*Parameter
					_, ok := res[bezier_]
					if ok {
						parameters = res[bezier_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[bezier_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveNextShiftedRight":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveNextShiftedRight != nil {
					beziergrid_ := parameter.GrowthCurveNextShiftedRight
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "GrowthCurveStack":
			res := make(map[*BezierGridStack][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveStack != nil {
					beziergridstack_ := parameter.GrowthCurveStack
					var parameters []*Parameter
					_, ok := res[beziergridstack_]
					if ok {
						parameters = res[beziergridstack_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergridstack_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralRhombusGridSeed":
			res := make(map[*SpiralRhombus][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralRhombusGridSeed != nil {
					spiralrhombus_ := parameter.SpiralRhombusGridSeed
					var parameters []*Parameter
					_, ok := res[spiralrhombus_]
					if ok {
						parameters = res[spiralrhombus_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralrhombus_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralRhombusGrid":
			res := make(map[*SpiralRhombusGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralRhombusGrid != nil {
					spiralrhombusgrid_ := parameter.SpiralRhombusGrid
					var parameters []*Parameter
					_, ok := res[spiralrhombusgrid_]
					if ok {
						parameters = res[spiralrhombusgrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralrhombusgrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralCircleSeed":
			res := make(map[*SpiralCircle][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralCircleSeed != nil {
					spiralcircle_ := parameter.SpiralCircleSeed
					var parameters []*Parameter
					_, ok := res[spiralcircle_]
					if ok {
						parameters = res[spiralcircle_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralcircle_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralCircleGrid":
			res := make(map[*SpiralCircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralCircleGrid != nil {
					spiralcirclegrid_ := parameter.SpiralCircleGrid
					var parameters []*Parameter
					_, ok := res[spiralcirclegrid_]
					if ok {
						parameters = res[spiralcirclegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralcirclegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralCircleFullGrid":
			res := make(map[*SpiralCircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralCircleFullGrid != nil {
					spiralcirclegrid_ := parameter.SpiralCircleFullGrid
					var parameters []*Parameter
					_, ok := res[spiralcirclegrid_]
					if ok {
						parameters = res[spiralcirclegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralcirclegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralConstructionOuterLineSeed":
			res := make(map[*SpiralLine][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralConstructionOuterLineSeed != nil {
					spiralline_ := parameter.SpiralConstructionOuterLineSeed
					var parameters []*Parameter
					_, ok := res[spiralline_]
					if ok {
						parameters = res[spiralline_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralline_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralConstructionInnerLineSeed":
			res := make(map[*SpiralLine][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralConstructionInnerLineSeed != nil {
					spiralline_ := parameter.SpiralConstructionInnerLineSeed
					var parameters []*Parameter
					_, ok := res[spiralline_]
					if ok {
						parameters = res[spiralline_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralline_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralConstructionOuterLineGrid":
			res := make(map[*SpiralLineGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralConstructionOuterLineGrid != nil {
					spirallinegrid_ := parameter.SpiralConstructionOuterLineGrid
					var parameters []*Parameter
					_, ok := res[spirallinegrid_]
					if ok {
						parameters = res[spirallinegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spirallinegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralConstructionInnerLineGrid":
			res := make(map[*SpiralLineGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralConstructionInnerLineGrid != nil {
					spirallinegrid_ := parameter.SpiralConstructionInnerLineGrid
					var parameters []*Parameter
					_, ok := res[spirallinegrid_]
					if ok {
						parameters = res[spirallinegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spirallinegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralConstructionCircleGrid":
			res := make(map[*SpiralCircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralConstructionCircleGrid != nil {
					spiralcirclegrid_ := parameter.SpiralConstructionCircleGrid
					var parameters []*Parameter
					_, ok := res[spiralcirclegrid_]
					if ok {
						parameters = res[spiralcirclegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralcirclegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralConstructionOuterLineFullGrid":
			res := make(map[*SpiralLineGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralConstructionOuterLineFullGrid != nil {
					spirallinegrid_ := parameter.SpiralConstructionOuterLineFullGrid
					var parameters []*Parameter
					_, ok := res[spirallinegrid_]
					if ok {
						parameters = res[spirallinegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spirallinegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralBezierSeed":
			res := make(map[*SpiralBezier][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralBezierSeed != nil {
					spiralbezier_ := parameter.SpiralBezierSeed
					var parameters []*Parameter
					_, ok := res[spiralbezier_]
					if ok {
						parameters = res[spiralbezier_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralbezier_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralBezierGrid":
			res := make(map[*SpiralBezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralBezierGrid != nil {
					spiralbeziergrid_ := parameter.SpiralBezierGrid
					var parameters []*Parameter
					_, ok := res[spiralbeziergrid_]
					if ok {
						parameters = res[spiralbeziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralbeziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralBezierFullGrid":
			res := make(map[*SpiralBezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralBezierFullGrid != nil {
					spiralbeziergrid_ := parameter.SpiralBezierFullGrid
					var parameters []*Parameter
					_, ok := res[spiralbeziergrid_]
					if ok {
						parameters = res[spiralbeziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralbeziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "FrontCurveStack":
			res := make(map[*FrontCurveStack][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.FrontCurveStack != nil {
					frontcurvestack_ := parameter.FrontCurveStack
					var parameters []*Parameter
					_, ok := res[frontcurvestack_]
					if ok {
						parameters = res[frontcurvestack_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[frontcurvestack_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "Fkey":
			res := make(map[*Key][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.Fkey != nil {
					key_ := parameter.Fkey
					var parameters []*Parameter
					_, ok := res[key_]
					if ok {
						parameters = res[key_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[key_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "PitchLines":
			res := make(map[*AxisGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.PitchLines != nil {
					axisgrid_ := parameter.PitchLines
					var parameters []*Parameter
					_, ok := res[axisgrid_]
					if ok {
						parameters = res[axisgrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[axisgrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "BeatLines":
			res := make(map[*AxisGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.BeatLines != nil {
					axisgrid_ := parameter.BeatLines
					var parameters []*Parameter
					_, ok := res[axisgrid_]
					if ok {
						parameters = res[axisgrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[axisgrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "FirstVoice":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.FirstVoice != nil {
					beziergrid_ := parameter.FirstVoice
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "FirstVoiceShiftedRigth":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.FirstVoiceShiftedRigth != nil {
					beziergrid_ := parameter.FirstVoiceShiftedRigth
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SecondVoice":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SecondVoice != nil {
					beziergrid_ := parameter.SecondVoice
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SecondVoiceShiftedRight":
			res := make(map[*BezierGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SecondVoiceShiftedRight != nil {
					beziergrid_ := parameter.SecondVoiceShiftedRight
					var parameters []*Parameter
					_, ok := res[beziergrid_]
					if ok {
						parameters = res[beziergrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[beziergrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "FirstVoiceNotes":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.FirstVoiceNotes != nil {
					circlegrid_ := parameter.FirstVoiceNotes
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "FirstVoiceNotesShiftedRight":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.FirstVoiceNotesShiftedRight != nil {
					circlegrid_ := parameter.FirstVoiceNotesShiftedRight
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SecondVoiceNotes":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SecondVoiceNotes != nil {
					circlegrid_ := parameter.SecondVoiceNotes
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SecondVoiceNotesShiftedRight":
			res := make(map[*CircleGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SecondVoiceNotesShiftedRight != nil {
					circlegrid_ := parameter.SecondVoiceNotesShiftedRight
					var parameters []*Parameter
					_, ok := res[circlegrid_]
					if ok {
						parameters = res[circlegrid_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[circlegrid_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "HorizontalAxis":
			res := make(map[*HorizontalAxis][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.HorizontalAxis != nil {
					horizontalaxis_ := parameter.HorizontalAxis
					var parameters []*Parameter
					_, ok := res[horizontalaxis_]
					if ok {
						parameters = res[horizontalaxis_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[horizontalaxis_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "VerticalAxis":
			res := make(map[*VerticalAxis][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.VerticalAxis != nil {
					verticalaxis_ := parameter.VerticalAxis
					var parameters []*Parameter
					_, ok := res[verticalaxis_]
					if ok {
						parameters = res[verticalaxis_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[verticalaxis_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralOrigin":
			res := make(map[*SpiralOrigin][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.SpiralOrigin != nil {
					spiralorigin_ := parameter.SpiralOrigin
					var parameters []*Parameter
					_, ok := res[spiralorigin_]
					if ok {
						parameters = res[spiralorigin_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[spiralorigin_] = parameters
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Rhombus
	case Rhombus:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*Rhombus)
			for rhombus := range stage.Rhombuss {
				if rhombus.ShapeCategory != nil {
					shapecategory_ := rhombus.ShapeCategory
					var rhombuss []*Rhombus
					_, ok := res[shapecategory_]
					if ok {
						rhombuss = res[shapecategory_]
					} else {
						rhombuss = make([]*Rhombus, 0)
					}
					rhombuss = append(rhombuss, rhombus)
					res[shapecategory_] = rhombuss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RhombusGrid
	case RhombusGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Reference":
			res := make(map[*Rhombus][]*RhombusGrid)
			for rhombusgrid := range stage.RhombusGrids {
				if rhombusgrid.Reference != nil {
					rhombus_ := rhombusgrid.Reference
					var rhombusgrids []*RhombusGrid
					_, ok := res[rhombus_]
					if ok {
						rhombusgrids = res[rhombus_]
					} else {
						rhombusgrids = make([]*RhombusGrid, 0)
					}
					rhombusgrids = append(rhombusgrids, rhombusgrid)
					res[rhombus_] = rhombusgrids
				}
			}
			return any(res).(map[*End][]*Start)
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*RhombusGrid)
			for rhombusgrid := range stage.RhombusGrids {
				if rhombusgrid.ShapeCategory != nil {
					shapecategory_ := rhombusgrid.ShapeCategory
					var rhombusgrids []*RhombusGrid
					_, ok := res[shapecategory_]
					if ok {
						rhombusgrids = res[shapecategory_]
					} else {
						rhombusgrids = make([]*RhombusGrid, 0)
					}
					rhombusgrids = append(rhombusgrids, rhombusgrid)
					res[shapecategory_] = rhombusgrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ShapeCategory
	case ShapeCategory:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpiralBezier
	case SpiralBezier:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralBezier)
			for spiralbezier := range stage.SpiralBeziers {
				if spiralbezier.ShapeCategory != nil {
					shapecategory_ := spiralbezier.ShapeCategory
					var spiralbeziers []*SpiralBezier
					_, ok := res[shapecategory_]
					if ok {
						spiralbeziers = res[shapecategory_]
					} else {
						spiralbeziers = make([]*SpiralBezier, 0)
					}
					spiralbeziers = append(spiralbeziers, spiralbezier)
					res[shapecategory_] = spiralbeziers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralBezierGrid
	case SpiralBezierGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralBezierGrid)
			for spiralbeziergrid := range stage.SpiralBezierGrids {
				if spiralbeziergrid.ShapeCategory != nil {
					shapecategory_ := spiralbeziergrid.ShapeCategory
					var spiralbeziergrids []*SpiralBezierGrid
					_, ok := res[shapecategory_]
					if ok {
						spiralbeziergrids = res[shapecategory_]
					} else {
						spiralbeziergrids = make([]*SpiralBezierGrid, 0)
					}
					spiralbeziergrids = append(spiralbeziergrids, spiralbeziergrid)
					res[shapecategory_] = spiralbeziergrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralCircle
	case SpiralCircle:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralCircle)
			for spiralcircle := range stage.SpiralCircles {
				if spiralcircle.ShapeCategory != nil {
					shapecategory_ := spiralcircle.ShapeCategory
					var spiralcircles []*SpiralCircle
					_, ok := res[shapecategory_]
					if ok {
						spiralcircles = res[shapecategory_]
					} else {
						spiralcircles = make([]*SpiralCircle, 0)
					}
					spiralcircles = append(spiralcircles, spiralcircle)
					res[shapecategory_] = spiralcircles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralCircleGrid
	case SpiralCircleGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralCircleGrid)
			for spiralcirclegrid := range stage.SpiralCircleGrids {
				if spiralcirclegrid.ShapeCategory != nil {
					shapecategory_ := spiralcirclegrid.ShapeCategory
					var spiralcirclegrids []*SpiralCircleGrid
					_, ok := res[shapecategory_]
					if ok {
						spiralcirclegrids = res[shapecategory_]
					} else {
						spiralcirclegrids = make([]*SpiralCircleGrid, 0)
					}
					spiralcirclegrids = append(spiralcirclegrids, spiralcirclegrid)
					res[shapecategory_] = spiralcirclegrids
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralRhombusGrid":
			res := make(map[*SpiralRhombusGrid][]*SpiralCircleGrid)
			for spiralcirclegrid := range stage.SpiralCircleGrids {
				if spiralcirclegrid.SpiralRhombusGrid != nil {
					spiralrhombusgrid_ := spiralcirclegrid.SpiralRhombusGrid
					var spiralcirclegrids []*SpiralCircleGrid
					_, ok := res[spiralrhombusgrid_]
					if ok {
						spiralcirclegrids = res[spiralrhombusgrid_]
					} else {
						spiralcirclegrids = make([]*SpiralCircleGrid, 0)
					}
					spiralcirclegrids = append(spiralcirclegrids, spiralcirclegrid)
					res[spiralrhombusgrid_] = spiralcirclegrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralLine
	case SpiralLine:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralLine)
			for spiralline := range stage.SpiralLines {
				if spiralline.ShapeCategory != nil {
					shapecategory_ := spiralline.ShapeCategory
					var spirallines []*SpiralLine
					_, ok := res[shapecategory_]
					if ok {
						spirallines = res[shapecategory_]
					} else {
						spirallines = make([]*SpiralLine, 0)
					}
					spirallines = append(spirallines, spiralline)
					res[shapecategory_] = spirallines
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralLineGrid
	case SpiralLineGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralLineGrid)
			for spirallinegrid := range stage.SpiralLineGrids {
				if spirallinegrid.ShapeCategory != nil {
					shapecategory_ := spirallinegrid.ShapeCategory
					var spirallinegrids []*SpiralLineGrid
					_, ok := res[shapecategory_]
					if ok {
						spirallinegrids = res[shapecategory_]
					} else {
						spirallinegrids = make([]*SpiralLineGrid, 0)
					}
					spirallinegrids = append(spirallinegrids, spirallinegrid)
					res[shapecategory_] = spirallinegrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralOrigin
	case SpiralOrigin:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralOrigin)
			for spiralorigin := range stage.SpiralOrigins {
				if spiralorigin.ShapeCategory != nil {
					shapecategory_ := spiralorigin.ShapeCategory
					var spiralorigins []*SpiralOrigin
					_, ok := res[shapecategory_]
					if ok {
						spiralorigins = res[shapecategory_]
					} else {
						spiralorigins = make([]*SpiralOrigin, 0)
					}
					spiralorigins = append(spiralorigins, spiralorigin)
					res[shapecategory_] = spiralorigins
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralRhombus
	case SpiralRhombus:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralRhombus)
			for spiralrhombus := range stage.SpiralRhombuss {
				if spiralrhombus.ShapeCategory != nil {
					shapecategory_ := spiralrhombus.ShapeCategory
					var spiralrhombuss []*SpiralRhombus
					_, ok := res[shapecategory_]
					if ok {
						spiralrhombuss = res[shapecategory_]
					} else {
						spiralrhombuss = make([]*SpiralRhombus, 0)
					}
					spiralrhombuss = append(spiralrhombuss, spiralrhombus)
					res[shapecategory_] = spiralrhombuss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SpiralRhombusGrid
	case SpiralRhombusGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*SpiralRhombusGrid)
			for spiralrhombusgrid := range stage.SpiralRhombusGrids {
				if spiralrhombusgrid.ShapeCategory != nil {
					shapecategory_ := spiralrhombusgrid.ShapeCategory
					var spiralrhombusgrids []*SpiralRhombusGrid
					_, ok := res[shapecategory_]
					if ok {
						spiralrhombusgrids = res[shapecategory_]
					} else {
						spiralrhombusgrids = make([]*SpiralRhombusGrid, 0)
					}
					spiralrhombusgrids = append(spiralrhombusgrids, spiralrhombusgrid)
					res[shapecategory_] = spiralrhombusgrids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalAxis
	case VerticalAxis:
		switch fieldname {
		// insertion point for per direct association field
		case "ShapeCategory":
			res := make(map[*ShapeCategory][]*VerticalAxis)
			for verticalaxis := range stage.VerticalAxiss {
				if verticalaxis.ShapeCategory != nil {
					shapecategory_ := verticalaxis.ShapeCategory
					var verticalaxiss []*VerticalAxis
					_, ok := res[shapecategory_]
					if ok {
						verticalaxiss = res[shapecategory_]
					} else {
						verticalaxiss = make([]*VerticalAxis, 0)
					}
					verticalaxiss = append(verticalaxiss, verticalaxis)
					res[shapecategory_] = verticalaxiss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Axis
	case Axis:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AxisGrid
	case AxisGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Axiss":
			res := make(map[*Axis]*AxisGrid)
			for axisgrid := range stage.AxisGrids {
				for _, axis_ := range axisgrid.Axiss {
					res[axis_] = axisgrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Bezier
	case Bezier:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BezierGrid
	case BezierGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Beziers":
			res := make(map[*Bezier]*BezierGrid)
			for beziergrid := range stage.BezierGrids {
				for _, bezier_ := range beziergrid.Beziers {
					res[bezier_] = beziergrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of BezierGridStack
	case BezierGridStack:
		switch fieldname {
		// insertion point for per direct association field
		case "BezierGrids":
			res := make(map[*BezierGrid]*BezierGridStack)
			for beziergridstack := range stage.BezierGridStacks {
				for _, beziergrid_ := range beziergridstack.BezierGrids {
					res[beziergrid_] = beziergridstack
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Circle
	case Circle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CircleGrid
	case CircleGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Circles":
			res := make(map[*Circle]*CircleGrid)
			for circlegrid := range stage.CircleGrids {
				for _, circle_ := range circlegrid.Circles {
					res[circle_] = circlegrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ExportToMusicxml
	case ExportToMusicxml:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FrontCurve
	case FrontCurve:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FrontCurveStack
	case FrontCurveStack:
		switch fieldname {
		// insertion point for per direct association field
		case "FrontCurves":
			res := make(map[*FrontCurve]*FrontCurveStack)
			for frontcurvestack := range stage.FrontCurveStacks {
				for _, frontcurve_ := range frontcurvestack.FrontCurves {
					res[frontcurve_] = frontcurvestack
				}
			}
			return any(res).(map[*End]*Start)
		case "SpiralCircles":
			res := make(map[*SpiralCircle]*FrontCurveStack)
			for frontcurvestack := range stage.FrontCurveStacks {
				for _, spiralcircle_ := range frontcurvestack.SpiralCircles {
					res[spiralcircle_] = frontcurvestack
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of HorizontalAxis
	case HorizontalAxis:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Key
	case Key:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Parameter
	case Parameter:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Rhombus
	case Rhombus:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RhombusGrid
	case RhombusGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "Rhombuses":
			res := make(map[*Rhombus]*RhombusGrid)
			for rhombusgrid := range stage.RhombusGrids {
				for _, rhombus_ := range rhombusgrid.Rhombuses {
					res[rhombus_] = rhombusgrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ShapeCategory
	case ShapeCategory:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpiralBezier
	case SpiralBezier:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpiralBezierGrid
	case SpiralBezierGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "SpiralBeziers":
			res := make(map[*SpiralBezier]*SpiralBezierGrid)
			for spiralbeziergrid := range stage.SpiralBezierGrids {
				for _, spiralbezier_ := range spiralbeziergrid.SpiralBeziers {
					res[spiralbezier_] = spiralbeziergrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SpiralCircle
	case SpiralCircle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpiralCircleGrid
	case SpiralCircleGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "SpiralCircles":
			res := make(map[*SpiralCircle]*SpiralCircleGrid)
			for spiralcirclegrid := range stage.SpiralCircleGrids {
				for _, spiralcircle_ := range spiralcirclegrid.SpiralCircles {
					res[spiralcircle_] = spiralcirclegrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SpiralLine
	case SpiralLine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpiralLineGrid
	case SpiralLineGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "SpiralLines":
			res := make(map[*SpiralLine]*SpiralLineGrid)
			for spirallinegrid := range stage.SpiralLineGrids {
				for _, spiralline_ := range spirallinegrid.SpiralLines {
					res[spiralline_] = spirallinegrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SpiralOrigin
	case SpiralOrigin:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpiralRhombus
	case SpiralRhombus:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpiralRhombusGrid
	case SpiralRhombusGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "SpiralRhombuses":
			res := make(map[*SpiralRhombus]*SpiralRhombusGrid)
			for spiralrhombusgrid := range stage.SpiralRhombusGrids {
				for _, spiralrhombus_ := range spiralrhombusgrid.SpiralRhombuses {
					res[spiralrhombus_] = spiralrhombusgrid
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of VerticalAxis
	case VerticalAxis:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Axis:
		res = "Axis"
	case AxisGrid:
		res = "AxisGrid"
	case Bezier:
		res = "Bezier"
	case BezierGrid:
		res = "BezierGrid"
	case BezierGridStack:
		res = "BezierGridStack"
	case Circle:
		res = "Circle"
	case CircleGrid:
		res = "CircleGrid"
	case ExportToMusicxml:
		res = "ExportToMusicxml"
	case FrontCurve:
		res = "FrontCurve"
	case FrontCurveStack:
		res = "FrontCurveStack"
	case HorizontalAxis:
		res = "HorizontalAxis"
	case Key:
		res = "Key"
	case Parameter:
		res = "Parameter"
	case Rhombus:
		res = "Rhombus"
	case RhombusGrid:
		res = "RhombusGrid"
	case ShapeCategory:
		res = "ShapeCategory"
	case SpiralBezier:
		res = "SpiralBezier"
	case SpiralBezierGrid:
		res = "SpiralBezierGrid"
	case SpiralCircle:
		res = "SpiralCircle"
	case SpiralCircleGrid:
		res = "SpiralCircleGrid"
	case SpiralLine:
		res = "SpiralLine"
	case SpiralLineGrid:
		res = "SpiralLineGrid"
	case SpiralOrigin:
		res = "SpiralOrigin"
	case SpiralRhombus:
		res = "SpiralRhombus"
	case SpiralRhombusGrid:
		res = "SpiralRhombusGrid"
	case VerticalAxis:
		res = "VerticalAxis"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Axis:
		res = "Axis"
	case *AxisGrid:
		res = "AxisGrid"
	case *Bezier:
		res = "Bezier"
	case *BezierGrid:
		res = "BezierGrid"
	case *BezierGridStack:
		res = "BezierGridStack"
	case *Circle:
		res = "Circle"
	case *CircleGrid:
		res = "CircleGrid"
	case *ExportToMusicxml:
		res = "ExportToMusicxml"
	case *FrontCurve:
		res = "FrontCurve"
	case *FrontCurveStack:
		res = "FrontCurveStack"
	case *HorizontalAxis:
		res = "HorizontalAxis"
	case *Key:
		res = "Key"
	case *Parameter:
		res = "Parameter"
	case *Rhombus:
		res = "Rhombus"
	case *RhombusGrid:
		res = "RhombusGrid"
	case *ShapeCategory:
		res = "ShapeCategory"
	case *SpiralBezier:
		res = "SpiralBezier"
	case *SpiralBezierGrid:
		res = "SpiralBezierGrid"
	case *SpiralCircle:
		res = "SpiralCircle"
	case *SpiralCircleGrid:
		res = "SpiralCircleGrid"
	case *SpiralLine:
		res = "SpiralLine"
	case *SpiralLineGrid:
		res = "SpiralLineGrid"
	case *SpiralOrigin:
		res = "SpiralOrigin"
	case *SpiralRhombus:
		res = "SpiralRhombus"
	case *SpiralRhombusGrid:
		res = "SpiralRhombusGrid"
	case *VerticalAxis:
		res = "VerticalAxis"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Axis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AngleDegree", "Length", "CenterX", "CenterY", "EndX", "EndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case AxisGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Axiss"}
	case Bezier:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "StartY", "ControlPointStartX", "ControlPointStartY", "EndX", "EndY", "ControlPointEndX", "ControlPointEndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case BezierGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Beziers"}
	case BezierGridStack:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "BezierGrids"}
	case Circle:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "HasBespokeRadius", "BespopkeRadius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Pitch", "ShowName", "BeatNb"}
	case CircleGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Circles"}
	case ExportToMusicxml:
		res = []string{"Name", "Parameter"}
	case FrontCurve:
		res = []string{"Name", "Path"}
	case FrontCurveStack:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "FrontCurves", "SpiralCircles", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case HorizontalAxis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AxisHandleBorderLength", "Axis_Length", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case Key:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Path", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case Parameter:
		res = []string{"Name", "BackendColor", "MinuteColor", "HourColor", "N", "M", "Z", "InsideAngle", "ShiftToNearestCircle", "SideLength", "InitialRhombus", "InitialCircle", "InitialRhombusGrid", "InitialCircleGrid", "InitialAxis", "RotatedAxis", "RotatedRhombus", "RotatedRhombusGrid", "RotatedCircleGrid", "NextRhombus", "NextCircle", "GrowingRhombusGridSeed", "GrowingRhombusGrid", "GrowingCircleGridSeed", "GrowingCircleGrid", "GrowingCircleGridLeftSeed", "GrowingCircleGridLeft", "ConstructionAxis", "ConstructionAxisGrid", "ConstructionCircle", "ConstructionCircleGrid", "GrowthCurveSeed", "GrowthCurve", "GrowthCurveShiftedRightSeed", "GrowthCurveShiftedRight", "GrowthCurveNextSeed", "GrowthCurveNext", "GrowthCurveNextShiftedRightSeed", "GrowthCurveNextShiftedRight", "GrowthCurveStack", "StackWidth", "NbShitRight", "StackHeight", "BezierControlLengthRatio", "SpiralRhombusGridSeed", "SpiralRhombusGrid", "SpiralCircleSeed", "SpiralCircleGrid", "SpiralCircleFullGrid", "SpiralConstructionOuterLineSeed", "SpiralConstructionInnerLineSeed", "SpiralConstructionOuterLineGrid", "SpiralConstructionInnerLineGrid", "SpiralConstructionCircleGrid", "SpiralConstructionOuterLineFullGrid", "SpiralBezierSeed", "SpiralBezierGrid", "SpiralBezierFullGrid", "SpiralBezierStrength", "FrontCurveStack", "NbInterpolationPoints", "Fkey", "FkeySizeRatio", "FkeyOriginRelativeX", "FkeyOriginRelativeY", "PitchLines", "PitchHeight", "NbPitchLines", "BeatLines", "BeatLinesHeightRatio", "NbBeatLines", "NbOfBeatsInTheme", "FirstVoice", "FirstVoiceShiftedRigth", "FirstVoiceShiftX", "FirstVoiceShiftY", "SecondVoice", "SecondVoiceShiftedRight", "PitchDifference", "BeatsPerSecond", "Level", "FirstVoiceNotes", "FirstVoiceNotesShiftedRight", "SecondVoiceNotes", "SecondVoiceNotesShiftedRight", "IsMinor", "ThemeBinaryEncoding", "OriginX", "OriginY", "HorizontalAxis", "VerticalAxis", "SpiralOrigin", "SpiralOriginX", "SpiralOriginY", "OriginCrossWidth", "SpiralRadiusRatio", "ShowSpiralBezierConstruct", "ShowInterpolationPoints", "ActualBeatsTemporalShift"}
	case Rhombus:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "SideLength", "AngleDegree", "InsideAngle", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case RhombusGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Rhombuses"}
	case ShapeCategory:
		res = []string{"Name", "IsExpanded"}
	case SpiralBezier:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "StartY", "ControlPointStartX", "ControlPointStartY", "EndX", "EndY", "ControlPointEndX", "ControlPointEndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case SpiralBezierGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralBeziers"}
	case SpiralCircle:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "HasBespokeRadius", "BespopkeRadius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Pitch", "ShowName", "BeatNb", "Path"}
	case SpiralCircleGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralRhombusGrid", "SpiralCircles"}
	case SpiralLine:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "EndX", "StartY", "EndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case SpiralLineGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralLines"}
	case SpiralOrigin:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case SpiralRhombus:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "X_r0", "Y_r0", "X_r1", "Y_r1", "X_r2", "Y_r2", "X_r3", "Y_r3", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case SpiralRhombusGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralRhombuses"}
	case VerticalAxis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AxisHandleBorderLength", "Axis_Length", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Axis:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AxisGrid"
		rf.Fieldname = "Axiss"
		res = append(res, rf)
	case AxisGrid:
		var rf ReverseField
		_ = rf
	case Bezier:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BezierGrid"
		rf.Fieldname = "Beziers"
		res = append(res, rf)
	case BezierGrid:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BezierGridStack"
		rf.Fieldname = "BezierGrids"
		res = append(res, rf)
	case BezierGridStack:
		var rf ReverseField
		_ = rf
	case Circle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "CircleGrid"
		rf.Fieldname = "Circles"
		res = append(res, rf)
	case CircleGrid:
		var rf ReverseField
		_ = rf
	case ExportToMusicxml:
		var rf ReverseField
		_ = rf
	case FrontCurve:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FrontCurveStack"
		rf.Fieldname = "FrontCurves"
		res = append(res, rf)
	case FrontCurveStack:
		var rf ReverseField
		_ = rf
	case HorizontalAxis:
		var rf ReverseField
		_ = rf
	case Key:
		var rf ReverseField
		_ = rf
	case Parameter:
		var rf ReverseField
		_ = rf
	case Rhombus:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RhombusGrid"
		rf.Fieldname = "Rhombuses"
		res = append(res, rf)
	case RhombusGrid:
		var rf ReverseField
		_ = rf
	case ShapeCategory:
		var rf ReverseField
		_ = rf
	case SpiralBezier:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SpiralBezierGrid"
		rf.Fieldname = "SpiralBeziers"
		res = append(res, rf)
	case SpiralBezierGrid:
		var rf ReverseField
		_ = rf
	case SpiralCircle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FrontCurveStack"
		rf.Fieldname = "SpiralCircles"
		res = append(res, rf)
		rf.GongstructName = "SpiralCircleGrid"
		rf.Fieldname = "SpiralCircles"
		res = append(res, rf)
	case SpiralCircleGrid:
		var rf ReverseField
		_ = rf
	case SpiralLine:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SpiralLineGrid"
		rf.Fieldname = "SpiralLines"
		res = append(res, rf)
	case SpiralLineGrid:
		var rf ReverseField
		_ = rf
	case SpiralOrigin:
		var rf ReverseField
		_ = rf
	case SpiralRhombus:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SpiralRhombusGrid"
		rf.Fieldname = "SpiralRhombuses"
		res = append(res, rf)
	case SpiralRhombusGrid:
		var rf ReverseField
		_ = rf
	case VerticalAxis:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Axis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AngleDegree", "Length", "CenterX", "CenterY", "EndX", "EndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *AxisGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Axiss"}
	case *Bezier:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "StartY", "ControlPointStartX", "ControlPointStartY", "EndX", "EndY", "ControlPointEndX", "ControlPointEndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *BezierGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Beziers"}
	case *BezierGridStack:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "BezierGrids"}
	case *Circle:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "HasBespokeRadius", "BespopkeRadius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Pitch", "ShowName", "BeatNb"}
	case *CircleGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Circles"}
	case *ExportToMusicxml:
		res = []string{"Name", "Parameter"}
	case *FrontCurve:
		res = []string{"Name", "Path"}
	case *FrontCurveStack:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "FrontCurves", "SpiralCircles", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *HorizontalAxis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AxisHandleBorderLength", "Axis_Length", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *Key:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Path", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *Parameter:
		res = []string{"Name", "BackendColor", "MinuteColor", "HourColor", "N", "M", "Z", "InsideAngle", "ShiftToNearestCircle", "SideLength", "InitialRhombus", "InitialCircle", "InitialRhombusGrid", "InitialCircleGrid", "InitialAxis", "RotatedAxis", "RotatedRhombus", "RotatedRhombusGrid", "RotatedCircleGrid", "NextRhombus", "NextCircle", "GrowingRhombusGridSeed", "GrowingRhombusGrid", "GrowingCircleGridSeed", "GrowingCircleGrid", "GrowingCircleGridLeftSeed", "GrowingCircleGridLeft", "ConstructionAxis", "ConstructionAxisGrid", "ConstructionCircle", "ConstructionCircleGrid", "GrowthCurveSeed", "GrowthCurve", "GrowthCurveShiftedRightSeed", "GrowthCurveShiftedRight", "GrowthCurveNextSeed", "GrowthCurveNext", "GrowthCurveNextShiftedRightSeed", "GrowthCurveNextShiftedRight", "GrowthCurveStack", "StackWidth", "NbShitRight", "StackHeight", "BezierControlLengthRatio", "SpiralRhombusGridSeed", "SpiralRhombusGrid", "SpiralCircleSeed", "SpiralCircleGrid", "SpiralCircleFullGrid", "SpiralConstructionOuterLineSeed", "SpiralConstructionInnerLineSeed", "SpiralConstructionOuterLineGrid", "SpiralConstructionInnerLineGrid", "SpiralConstructionCircleGrid", "SpiralConstructionOuterLineFullGrid", "SpiralBezierSeed", "SpiralBezierGrid", "SpiralBezierFullGrid", "SpiralBezierStrength", "FrontCurveStack", "NbInterpolationPoints", "Fkey", "FkeySizeRatio", "FkeyOriginRelativeX", "FkeyOriginRelativeY", "PitchLines", "PitchHeight", "NbPitchLines", "BeatLines", "BeatLinesHeightRatio", "NbBeatLines", "NbOfBeatsInTheme", "FirstVoice", "FirstVoiceShiftedRigth", "FirstVoiceShiftX", "FirstVoiceShiftY", "SecondVoice", "SecondVoiceShiftedRight", "PitchDifference", "BeatsPerSecond", "Level", "FirstVoiceNotes", "FirstVoiceNotesShiftedRight", "SecondVoiceNotes", "SecondVoiceNotesShiftedRight", "IsMinor", "ThemeBinaryEncoding", "OriginX", "OriginY", "HorizontalAxis", "VerticalAxis", "SpiralOrigin", "SpiralOriginX", "SpiralOriginY", "OriginCrossWidth", "SpiralRadiusRatio", "ShowSpiralBezierConstruct", "ShowInterpolationPoints", "ActualBeatsTemporalShift"}
	case *Rhombus:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "SideLength", "AngleDegree", "InsideAngle", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *RhombusGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Rhombuses"}
	case *ShapeCategory:
		res = []string{"Name", "IsExpanded"}
	case *SpiralBezier:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "StartY", "ControlPointStartX", "ControlPointStartY", "EndX", "EndY", "ControlPointEndX", "ControlPointEndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *SpiralBezierGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralBeziers"}
	case *SpiralCircle:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "HasBespokeRadius", "BespopkeRadius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "Pitch", "ShowName", "BeatNb", "Path"}
	case *SpiralCircleGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralRhombusGrid", "SpiralCircles"}
	case *SpiralLine:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "EndX", "StartY", "EndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *SpiralLineGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralLines"}
	case *SpiralOrigin:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *SpiralRhombus:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "X_r0", "Y_r0", "X_r1", "Y_r1", "X_r2", "Y_r2", "X_r3", "Y_r3", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *SpiralRhombusGrid:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "SpiralRhombuses"}
	case *VerticalAxis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AxisHandleBorderLength", "Axis_Length", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt    GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat  GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool   GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Axis:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "AngleDegree":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AngleDegree)
			res.valueFloat = inferedInstance.AngleDegree
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Length":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Length)
			res.valueFloat = inferedInstance.Length
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *AxisGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Axiss":
			for idx, __instance__ := range inferedInstance.Axiss {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Bezier:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartY)
			res.valueFloat = inferedInstance.StartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartX)
			res.valueFloat = inferedInstance.ControlPointStartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartY)
			res.valueFloat = inferedInstance.ControlPointStartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndX)
			res.valueFloat = inferedInstance.ControlPointEndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndY)
			res.valueFloat = inferedInstance.ControlPointEndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *BezierGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Beziers":
			for idx, __instance__ := range inferedInstance.Beziers {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *BezierGridStack:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "BezierGrids":
			for idx, __instance__ := range inferedInstance.BezierGrids {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Circle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasBespokeRadius":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasBespokeRadius)
			res.valueBool = inferedInstance.HasBespokeRadius
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BespopkeRadius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BespopkeRadius)
			res.valueFloat = inferedInstance.BespopkeRadius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Pitch":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Pitch)
			res.valueInt = inferedInstance.Pitch
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ShowName":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowName)
			res.valueBool = inferedInstance.ShowName
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BeatNb":
			res.valueString = fmt.Sprintf("%d", inferedInstance.BeatNb)
			res.valueInt = inferedInstance.BeatNb
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case *CircleGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Circles":
			for idx, __instance__ := range inferedInstance.Circles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *ExportToMusicxml:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Parameter":
			if inferedInstance.Parameter != nil {
				res.valueString = inferedInstance.Parameter.Name
			}
		}
	case *FrontCurve:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Path":
			res.valueString = inferedInstance.Path
		}
	case *FrontCurveStack:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "FrontCurves":
			for idx, __instance__ := range inferedInstance.FrontCurves {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SpiralCircles":
			for idx, __instance__ := range inferedInstance.SpiralCircles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *HorizontalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
			res.valueFloat = inferedInstance.AxisHandleBorderLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Axis_Length":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Axis_Length)
			res.valueFloat = inferedInstance.Axis_Length
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *Key:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Path":
			res.valueString = inferedInstance.Path
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *Parameter:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "BackendColor":
			res.valueString = inferedInstance.BackendColor
		case "MinuteColor":
			res.valueString = inferedInstance.MinuteColor
		case "HourColor":
			res.valueString = inferedInstance.HourColor
		case "N":
			res.valueString = fmt.Sprintf("%d", inferedInstance.N)
			res.valueInt = inferedInstance.N
			res.GongFieldValueType = GongFieldValueTypeInt
		case "M":
			res.valueString = fmt.Sprintf("%d", inferedInstance.M)
			res.valueInt = inferedInstance.M
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Z":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Z)
			res.valueInt = inferedInstance.Z
			res.GongFieldValueType = GongFieldValueTypeInt
		case "InsideAngle":
			res.valueString = fmt.Sprintf("%f", inferedInstance.InsideAngle)
			res.valueFloat = inferedInstance.InsideAngle
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ShiftToNearestCircle":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ShiftToNearestCircle)
			res.valueInt = inferedInstance.ShiftToNearestCircle
			res.GongFieldValueType = GongFieldValueTypeInt
		case "SideLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SideLength)
			res.valueFloat = inferedInstance.SideLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "InitialRhombus":
			if inferedInstance.InitialRhombus != nil {
				res.valueString = inferedInstance.InitialRhombus.Name
			}
		case "InitialCircle":
			if inferedInstance.InitialCircle != nil {
				res.valueString = inferedInstance.InitialCircle.Name
			}
		case "InitialRhombusGrid":
			if inferedInstance.InitialRhombusGrid != nil {
				res.valueString = inferedInstance.InitialRhombusGrid.Name
			}
		case "InitialCircleGrid":
			if inferedInstance.InitialCircleGrid != nil {
				res.valueString = inferedInstance.InitialCircleGrid.Name
			}
		case "InitialAxis":
			if inferedInstance.InitialAxis != nil {
				res.valueString = inferedInstance.InitialAxis.Name
			}
		case "RotatedAxis":
			if inferedInstance.RotatedAxis != nil {
				res.valueString = inferedInstance.RotatedAxis.Name
			}
		case "RotatedRhombus":
			if inferedInstance.RotatedRhombus != nil {
				res.valueString = inferedInstance.RotatedRhombus.Name
			}
		case "RotatedRhombusGrid":
			if inferedInstance.RotatedRhombusGrid != nil {
				res.valueString = inferedInstance.RotatedRhombusGrid.Name
			}
		case "RotatedCircleGrid":
			if inferedInstance.RotatedCircleGrid != nil {
				res.valueString = inferedInstance.RotatedCircleGrid.Name
			}
		case "NextRhombus":
			if inferedInstance.NextRhombus != nil {
				res.valueString = inferedInstance.NextRhombus.Name
			}
		case "NextCircle":
			if inferedInstance.NextCircle != nil {
				res.valueString = inferedInstance.NextCircle.Name
			}
		case "GrowingRhombusGridSeed":
			if inferedInstance.GrowingRhombusGridSeed != nil {
				res.valueString = inferedInstance.GrowingRhombusGridSeed.Name
			}
		case "GrowingRhombusGrid":
			if inferedInstance.GrowingRhombusGrid != nil {
				res.valueString = inferedInstance.GrowingRhombusGrid.Name
			}
		case "GrowingCircleGridSeed":
			if inferedInstance.GrowingCircleGridSeed != nil {
				res.valueString = inferedInstance.GrowingCircleGridSeed.Name
			}
		case "GrowingCircleGrid":
			if inferedInstance.GrowingCircleGrid != nil {
				res.valueString = inferedInstance.GrowingCircleGrid.Name
			}
		case "GrowingCircleGridLeftSeed":
			if inferedInstance.GrowingCircleGridLeftSeed != nil {
				res.valueString = inferedInstance.GrowingCircleGridLeftSeed.Name
			}
		case "GrowingCircleGridLeft":
			if inferedInstance.GrowingCircleGridLeft != nil {
				res.valueString = inferedInstance.GrowingCircleGridLeft.Name
			}
		case "ConstructionAxis":
			if inferedInstance.ConstructionAxis != nil {
				res.valueString = inferedInstance.ConstructionAxis.Name
			}
		case "ConstructionAxisGrid":
			if inferedInstance.ConstructionAxisGrid != nil {
				res.valueString = inferedInstance.ConstructionAxisGrid.Name
			}
		case "ConstructionCircle":
			if inferedInstance.ConstructionCircle != nil {
				res.valueString = inferedInstance.ConstructionCircle.Name
			}
		case "ConstructionCircleGrid":
			if inferedInstance.ConstructionCircleGrid != nil {
				res.valueString = inferedInstance.ConstructionCircleGrid.Name
			}
		case "GrowthCurveSeed":
			if inferedInstance.GrowthCurveSeed != nil {
				res.valueString = inferedInstance.GrowthCurveSeed.Name
			}
		case "GrowthCurve":
			if inferedInstance.GrowthCurve != nil {
				res.valueString = inferedInstance.GrowthCurve.Name
			}
		case "GrowthCurveShiftedRightSeed":
			if inferedInstance.GrowthCurveShiftedRightSeed != nil {
				res.valueString = inferedInstance.GrowthCurveShiftedRightSeed.Name
			}
		case "GrowthCurveShiftedRight":
			if inferedInstance.GrowthCurveShiftedRight != nil {
				res.valueString = inferedInstance.GrowthCurveShiftedRight.Name
			}
		case "GrowthCurveNextSeed":
			if inferedInstance.GrowthCurveNextSeed != nil {
				res.valueString = inferedInstance.GrowthCurveNextSeed.Name
			}
		case "GrowthCurveNext":
			if inferedInstance.GrowthCurveNext != nil {
				res.valueString = inferedInstance.GrowthCurveNext.Name
			}
		case "GrowthCurveNextShiftedRightSeed":
			if inferedInstance.GrowthCurveNextShiftedRightSeed != nil {
				res.valueString = inferedInstance.GrowthCurveNextShiftedRightSeed.Name
			}
		case "GrowthCurveNextShiftedRight":
			if inferedInstance.GrowthCurveNextShiftedRight != nil {
				res.valueString = inferedInstance.GrowthCurveNextShiftedRight.Name
			}
		case "GrowthCurveStack":
			if inferedInstance.GrowthCurveStack != nil {
				res.valueString = inferedInstance.GrowthCurveStack.Name
			}
		case "StackWidth":
			res.valueString = fmt.Sprintf("%d", inferedInstance.StackWidth)
			res.valueInt = inferedInstance.StackWidth
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NbShitRight":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbShitRight)
			res.valueInt = inferedInstance.NbShitRight
			res.GongFieldValueType = GongFieldValueTypeInt
		case "StackHeight":
			res.valueString = fmt.Sprintf("%d", inferedInstance.StackHeight)
			res.valueInt = inferedInstance.StackHeight
			res.GongFieldValueType = GongFieldValueTypeInt
		case "BezierControlLengthRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BezierControlLengthRatio)
			res.valueFloat = inferedInstance.BezierControlLengthRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SpiralRhombusGridSeed":
			if inferedInstance.SpiralRhombusGridSeed != nil {
				res.valueString = inferedInstance.SpiralRhombusGridSeed.Name
			}
		case "SpiralRhombusGrid":
			if inferedInstance.SpiralRhombusGrid != nil {
				res.valueString = inferedInstance.SpiralRhombusGrid.Name
			}
		case "SpiralCircleSeed":
			if inferedInstance.SpiralCircleSeed != nil {
				res.valueString = inferedInstance.SpiralCircleSeed.Name
			}
		case "SpiralCircleGrid":
			if inferedInstance.SpiralCircleGrid != nil {
				res.valueString = inferedInstance.SpiralCircleGrid.Name
			}
		case "SpiralCircleFullGrid":
			if inferedInstance.SpiralCircleFullGrid != nil {
				res.valueString = inferedInstance.SpiralCircleFullGrid.Name
			}
		case "SpiralConstructionOuterLineSeed":
			if inferedInstance.SpiralConstructionOuterLineSeed != nil {
				res.valueString = inferedInstance.SpiralConstructionOuterLineSeed.Name
			}
		case "SpiralConstructionInnerLineSeed":
			if inferedInstance.SpiralConstructionInnerLineSeed != nil {
				res.valueString = inferedInstance.SpiralConstructionInnerLineSeed.Name
			}
		case "SpiralConstructionOuterLineGrid":
			if inferedInstance.SpiralConstructionOuterLineGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionOuterLineGrid.Name
			}
		case "SpiralConstructionInnerLineGrid":
			if inferedInstance.SpiralConstructionInnerLineGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionInnerLineGrid.Name
			}
		case "SpiralConstructionCircleGrid":
			if inferedInstance.SpiralConstructionCircleGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionCircleGrid.Name
			}
		case "SpiralConstructionOuterLineFullGrid":
			if inferedInstance.SpiralConstructionOuterLineFullGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionOuterLineFullGrid.Name
			}
		case "SpiralBezierSeed":
			if inferedInstance.SpiralBezierSeed != nil {
				res.valueString = inferedInstance.SpiralBezierSeed.Name
			}
		case "SpiralBezierGrid":
			if inferedInstance.SpiralBezierGrid != nil {
				res.valueString = inferedInstance.SpiralBezierGrid.Name
			}
		case "SpiralBezierFullGrid":
			if inferedInstance.SpiralBezierFullGrid != nil {
				res.valueString = inferedInstance.SpiralBezierFullGrid.Name
			}
		case "SpiralBezierStrength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralBezierStrength)
			res.valueFloat = inferedInstance.SpiralBezierStrength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FrontCurveStack":
			if inferedInstance.FrontCurveStack != nil {
				res.valueString = inferedInstance.FrontCurveStack.Name
			}
		case "NbInterpolationPoints":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbInterpolationPoints)
			res.valueInt = inferedInstance.NbInterpolationPoints
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Fkey":
			if inferedInstance.Fkey != nil {
				res.valueString = inferedInstance.Fkey.Name
			}
		case "FkeySizeRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FkeySizeRatio)
			res.valueFloat = inferedInstance.FkeySizeRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FkeyOriginRelativeX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeX)
			res.valueFloat = inferedInstance.FkeyOriginRelativeX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FkeyOriginRelativeY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeY)
			res.valueFloat = inferedInstance.FkeyOriginRelativeY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "PitchLines":
			if inferedInstance.PitchLines != nil {
				res.valueString = inferedInstance.PitchLines.Name
			}
		case "PitchHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.PitchHeight)
			res.valueFloat = inferedInstance.PitchHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "NbPitchLines":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbPitchLines)
			res.valueInt = inferedInstance.NbPitchLines
			res.GongFieldValueType = GongFieldValueTypeInt
		case "BeatLines":
			if inferedInstance.BeatLines != nil {
				res.valueString = inferedInstance.BeatLines.Name
			}
		case "BeatLinesHeightRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BeatLinesHeightRatio)
			res.valueFloat = inferedInstance.BeatLinesHeightRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "NbBeatLines":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbBeatLines)
			res.valueInt = inferedInstance.NbBeatLines
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NbOfBeatsInTheme":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbOfBeatsInTheme)
			res.valueInt = inferedInstance.NbOfBeatsInTheme
			res.GongFieldValueType = GongFieldValueTypeInt
		case "FirstVoice":
			if inferedInstance.FirstVoice != nil {
				res.valueString = inferedInstance.FirstVoice.Name
			}
		case "FirstVoiceShiftedRigth":
			if inferedInstance.FirstVoiceShiftedRigth != nil {
				res.valueString = inferedInstance.FirstVoiceShiftedRigth.Name
			}
		case "FirstVoiceShiftX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FirstVoiceShiftX)
			res.valueFloat = inferedInstance.FirstVoiceShiftX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FirstVoiceShiftY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FirstVoiceShiftY)
			res.valueFloat = inferedInstance.FirstVoiceShiftY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SecondVoice":
			if inferedInstance.SecondVoice != nil {
				res.valueString = inferedInstance.SecondVoice.Name
			}
		case "SecondVoiceShiftedRight":
			if inferedInstance.SecondVoiceShiftedRight != nil {
				res.valueString = inferedInstance.SecondVoiceShiftedRight.Name
			}
		case "PitchDifference":
			res.valueString = fmt.Sprintf("%d", inferedInstance.PitchDifference)
			res.valueInt = inferedInstance.PitchDifference
			res.GongFieldValueType = GongFieldValueTypeInt
		case "BeatsPerSecond":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BeatsPerSecond)
			res.valueFloat = inferedInstance.BeatsPerSecond
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Level":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Level)
			res.valueFloat = inferedInstance.Level
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FirstVoiceNotes":
			if inferedInstance.FirstVoiceNotes != nil {
				res.valueString = inferedInstance.FirstVoiceNotes.Name
			}
		case "FirstVoiceNotesShiftedRight":
			if inferedInstance.FirstVoiceNotesShiftedRight != nil {
				res.valueString = inferedInstance.FirstVoiceNotesShiftedRight.Name
			}
		case "SecondVoiceNotes":
			if inferedInstance.SecondVoiceNotes != nil {
				res.valueString = inferedInstance.SecondVoiceNotes.Name
			}
		case "SecondVoiceNotesShiftedRight":
			if inferedInstance.SecondVoiceNotesShiftedRight != nil {
				res.valueString = inferedInstance.SecondVoiceNotesShiftedRight.Name
			}
		case "IsMinor":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsMinor)
			res.valueBool = inferedInstance.IsMinor
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ThemeBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ThemeBinaryEncoding)
			res.valueInt = inferedInstance.ThemeBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
		case "OriginX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.OriginX)
			res.valueFloat = inferedInstance.OriginX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "OriginY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.OriginY)
			res.valueFloat = inferedInstance.OriginY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HorizontalAxis":
			if inferedInstance.HorizontalAxis != nil {
				res.valueString = inferedInstance.HorizontalAxis.Name
			}
		case "VerticalAxis":
			if inferedInstance.VerticalAxis != nil {
				res.valueString = inferedInstance.VerticalAxis.Name
			}
		case "SpiralOrigin":
			if inferedInstance.SpiralOrigin != nil {
				res.valueString = inferedInstance.SpiralOrigin.Name
			}
		case "SpiralOriginX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralOriginX)
			res.valueFloat = inferedInstance.SpiralOriginX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SpiralOriginY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralOriginY)
			res.valueFloat = inferedInstance.SpiralOriginY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "OriginCrossWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.OriginCrossWidth)
			res.valueFloat = inferedInstance.OriginCrossWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SpiralRadiusRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralRadiusRatio)
			res.valueFloat = inferedInstance.SpiralRadiusRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ShowSpiralBezierConstruct":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowSpiralBezierConstruct)
			res.valueBool = inferedInstance.ShowSpiralBezierConstruct
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShowInterpolationPoints":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowInterpolationPoints)
			res.valueBool = inferedInstance.ShowInterpolationPoints
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ActualBeatsTemporalShift":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ActualBeatsTemporalShift)
			res.valueInt = inferedInstance.ActualBeatsTemporalShift
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case *Rhombus:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SideLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SideLength)
			res.valueFloat = inferedInstance.SideLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "AngleDegree":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AngleDegree)
			res.valueFloat = inferedInstance.AngleDegree
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "InsideAngle":
			res.valueString = fmt.Sprintf("%f", inferedInstance.InsideAngle)
			res.valueFloat = inferedInstance.InsideAngle
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *RhombusGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Rhombuses":
			for idx, __instance__ := range inferedInstance.Rhombuses {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *ShapeCategory:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *SpiralBezier:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartY)
			res.valueFloat = inferedInstance.StartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartX)
			res.valueFloat = inferedInstance.ControlPointStartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartY)
			res.valueFloat = inferedInstance.ControlPointStartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndX)
			res.valueFloat = inferedInstance.ControlPointEndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndY)
			res.valueFloat = inferedInstance.ControlPointEndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *SpiralBezierGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralBeziers":
			for idx, __instance__ := range inferedInstance.SpiralBeziers {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *SpiralCircle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasBespokeRadius":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasBespokeRadius)
			res.valueBool = inferedInstance.HasBespokeRadius
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BespopkeRadius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BespopkeRadius)
			res.valueFloat = inferedInstance.BespopkeRadius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Pitch":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Pitch)
			res.valueInt = inferedInstance.Pitch
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ShowName":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowName)
			res.valueBool = inferedInstance.ShowName
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BeatNb":
			res.valueString = fmt.Sprintf("%d", inferedInstance.BeatNb)
			res.valueInt = inferedInstance.BeatNb
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Path":
			res.valueString = inferedInstance.Path
		}
	case *SpiralCircleGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralRhombusGrid":
			if inferedInstance.SpiralRhombusGrid != nil {
				res.valueString = inferedInstance.SpiralRhombusGrid.Name
			}
		case "SpiralCircles":
			for idx, __instance__ := range inferedInstance.SpiralCircles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *SpiralLine:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartY)
			res.valueFloat = inferedInstance.StartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *SpiralLineGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralLines":
			for idx, __instance__ := range inferedInstance.SpiralLines {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *SpiralOrigin:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *SpiralRhombus:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "X_r0":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r0)
			res.valueFloat = inferedInstance.X_r0
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r0":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r0)
			res.valueFloat = inferedInstance.Y_r0
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_r1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r1)
			res.valueFloat = inferedInstance.X_r1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r1)
			res.valueFloat = inferedInstance.Y_r1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_r2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r2)
			res.valueFloat = inferedInstance.X_r2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r2)
			res.valueFloat = inferedInstance.Y_r2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_r3":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r3)
			res.valueFloat = inferedInstance.X_r3
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r3":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r3)
			res.valueFloat = inferedInstance.Y_r3
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case *SpiralRhombusGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralRhombuses":
			for idx, __instance__ := range inferedInstance.SpiralRhombuses {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *VerticalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
			res.valueFloat = inferedInstance.AxisHandleBorderLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Axis_Length":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Axis_Length)
			res.valueFloat = inferedInstance.Axis_Length
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Axis:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "AngleDegree":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AngleDegree)
			res.valueFloat = inferedInstance.AngleDegree
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Length":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Length)
			res.valueFloat = inferedInstance.Length
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case AxisGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Axiss":
			for idx, __instance__ := range inferedInstance.Axiss {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Bezier:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartY)
			res.valueFloat = inferedInstance.StartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartX)
			res.valueFloat = inferedInstance.ControlPointStartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartY)
			res.valueFloat = inferedInstance.ControlPointStartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndX)
			res.valueFloat = inferedInstance.ControlPointEndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndY)
			res.valueFloat = inferedInstance.ControlPointEndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case BezierGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Beziers":
			for idx, __instance__ := range inferedInstance.Beziers {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case BezierGridStack:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "BezierGrids":
			for idx, __instance__ := range inferedInstance.BezierGrids {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Circle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasBespokeRadius":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasBespokeRadius)
			res.valueBool = inferedInstance.HasBespokeRadius
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BespopkeRadius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BespopkeRadius)
			res.valueFloat = inferedInstance.BespopkeRadius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Pitch":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Pitch)
			res.valueInt = inferedInstance.Pitch
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ShowName":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowName)
			res.valueBool = inferedInstance.ShowName
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BeatNb":
			res.valueString = fmt.Sprintf("%d", inferedInstance.BeatNb)
			res.valueInt = inferedInstance.BeatNb
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case CircleGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Circles":
			for idx, __instance__ := range inferedInstance.Circles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case ExportToMusicxml:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Parameter":
			if inferedInstance.Parameter != nil {
				res.valueString = inferedInstance.Parameter.Name
			}
		}
	case FrontCurve:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Path":
			res.valueString = inferedInstance.Path
		}
	case FrontCurveStack:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "FrontCurves":
			for idx, __instance__ := range inferedInstance.FrontCurves {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SpiralCircles":
			for idx, __instance__ := range inferedInstance.SpiralCircles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case HorizontalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
			res.valueFloat = inferedInstance.AxisHandleBorderLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Axis_Length":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Axis_Length)
			res.valueFloat = inferedInstance.Axis_Length
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case Key:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Path":
			res.valueString = inferedInstance.Path
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case Parameter:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "BackendColor":
			res.valueString = inferedInstance.BackendColor
		case "MinuteColor":
			res.valueString = inferedInstance.MinuteColor
		case "HourColor":
			res.valueString = inferedInstance.HourColor
		case "N":
			res.valueString = fmt.Sprintf("%d", inferedInstance.N)
			res.valueInt = inferedInstance.N
			res.GongFieldValueType = GongFieldValueTypeInt
		case "M":
			res.valueString = fmt.Sprintf("%d", inferedInstance.M)
			res.valueInt = inferedInstance.M
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Z":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Z)
			res.valueInt = inferedInstance.Z
			res.GongFieldValueType = GongFieldValueTypeInt
		case "InsideAngle":
			res.valueString = fmt.Sprintf("%f", inferedInstance.InsideAngle)
			res.valueFloat = inferedInstance.InsideAngle
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ShiftToNearestCircle":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ShiftToNearestCircle)
			res.valueInt = inferedInstance.ShiftToNearestCircle
			res.GongFieldValueType = GongFieldValueTypeInt
		case "SideLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SideLength)
			res.valueFloat = inferedInstance.SideLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "InitialRhombus":
			if inferedInstance.InitialRhombus != nil {
				res.valueString = inferedInstance.InitialRhombus.Name
			}
		case "InitialCircle":
			if inferedInstance.InitialCircle != nil {
				res.valueString = inferedInstance.InitialCircle.Name
			}
		case "InitialRhombusGrid":
			if inferedInstance.InitialRhombusGrid != nil {
				res.valueString = inferedInstance.InitialRhombusGrid.Name
			}
		case "InitialCircleGrid":
			if inferedInstance.InitialCircleGrid != nil {
				res.valueString = inferedInstance.InitialCircleGrid.Name
			}
		case "InitialAxis":
			if inferedInstance.InitialAxis != nil {
				res.valueString = inferedInstance.InitialAxis.Name
			}
		case "RotatedAxis":
			if inferedInstance.RotatedAxis != nil {
				res.valueString = inferedInstance.RotatedAxis.Name
			}
		case "RotatedRhombus":
			if inferedInstance.RotatedRhombus != nil {
				res.valueString = inferedInstance.RotatedRhombus.Name
			}
		case "RotatedRhombusGrid":
			if inferedInstance.RotatedRhombusGrid != nil {
				res.valueString = inferedInstance.RotatedRhombusGrid.Name
			}
		case "RotatedCircleGrid":
			if inferedInstance.RotatedCircleGrid != nil {
				res.valueString = inferedInstance.RotatedCircleGrid.Name
			}
		case "NextRhombus":
			if inferedInstance.NextRhombus != nil {
				res.valueString = inferedInstance.NextRhombus.Name
			}
		case "NextCircle":
			if inferedInstance.NextCircle != nil {
				res.valueString = inferedInstance.NextCircle.Name
			}
		case "GrowingRhombusGridSeed":
			if inferedInstance.GrowingRhombusGridSeed != nil {
				res.valueString = inferedInstance.GrowingRhombusGridSeed.Name
			}
		case "GrowingRhombusGrid":
			if inferedInstance.GrowingRhombusGrid != nil {
				res.valueString = inferedInstance.GrowingRhombusGrid.Name
			}
		case "GrowingCircleGridSeed":
			if inferedInstance.GrowingCircleGridSeed != nil {
				res.valueString = inferedInstance.GrowingCircleGridSeed.Name
			}
		case "GrowingCircleGrid":
			if inferedInstance.GrowingCircleGrid != nil {
				res.valueString = inferedInstance.GrowingCircleGrid.Name
			}
		case "GrowingCircleGridLeftSeed":
			if inferedInstance.GrowingCircleGridLeftSeed != nil {
				res.valueString = inferedInstance.GrowingCircleGridLeftSeed.Name
			}
		case "GrowingCircleGridLeft":
			if inferedInstance.GrowingCircleGridLeft != nil {
				res.valueString = inferedInstance.GrowingCircleGridLeft.Name
			}
		case "ConstructionAxis":
			if inferedInstance.ConstructionAxis != nil {
				res.valueString = inferedInstance.ConstructionAxis.Name
			}
		case "ConstructionAxisGrid":
			if inferedInstance.ConstructionAxisGrid != nil {
				res.valueString = inferedInstance.ConstructionAxisGrid.Name
			}
		case "ConstructionCircle":
			if inferedInstance.ConstructionCircle != nil {
				res.valueString = inferedInstance.ConstructionCircle.Name
			}
		case "ConstructionCircleGrid":
			if inferedInstance.ConstructionCircleGrid != nil {
				res.valueString = inferedInstance.ConstructionCircleGrid.Name
			}
		case "GrowthCurveSeed":
			if inferedInstance.GrowthCurveSeed != nil {
				res.valueString = inferedInstance.GrowthCurveSeed.Name
			}
		case "GrowthCurve":
			if inferedInstance.GrowthCurve != nil {
				res.valueString = inferedInstance.GrowthCurve.Name
			}
		case "GrowthCurveShiftedRightSeed":
			if inferedInstance.GrowthCurveShiftedRightSeed != nil {
				res.valueString = inferedInstance.GrowthCurveShiftedRightSeed.Name
			}
		case "GrowthCurveShiftedRight":
			if inferedInstance.GrowthCurveShiftedRight != nil {
				res.valueString = inferedInstance.GrowthCurveShiftedRight.Name
			}
		case "GrowthCurveNextSeed":
			if inferedInstance.GrowthCurveNextSeed != nil {
				res.valueString = inferedInstance.GrowthCurveNextSeed.Name
			}
		case "GrowthCurveNext":
			if inferedInstance.GrowthCurveNext != nil {
				res.valueString = inferedInstance.GrowthCurveNext.Name
			}
		case "GrowthCurveNextShiftedRightSeed":
			if inferedInstance.GrowthCurveNextShiftedRightSeed != nil {
				res.valueString = inferedInstance.GrowthCurveNextShiftedRightSeed.Name
			}
		case "GrowthCurveNextShiftedRight":
			if inferedInstance.GrowthCurveNextShiftedRight != nil {
				res.valueString = inferedInstance.GrowthCurveNextShiftedRight.Name
			}
		case "GrowthCurveStack":
			if inferedInstance.GrowthCurveStack != nil {
				res.valueString = inferedInstance.GrowthCurveStack.Name
			}
		case "StackWidth":
			res.valueString = fmt.Sprintf("%d", inferedInstance.StackWidth)
			res.valueInt = inferedInstance.StackWidth
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NbShitRight":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbShitRight)
			res.valueInt = inferedInstance.NbShitRight
			res.GongFieldValueType = GongFieldValueTypeInt
		case "StackHeight":
			res.valueString = fmt.Sprintf("%d", inferedInstance.StackHeight)
			res.valueInt = inferedInstance.StackHeight
			res.GongFieldValueType = GongFieldValueTypeInt
		case "BezierControlLengthRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BezierControlLengthRatio)
			res.valueFloat = inferedInstance.BezierControlLengthRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SpiralRhombusGridSeed":
			if inferedInstance.SpiralRhombusGridSeed != nil {
				res.valueString = inferedInstance.SpiralRhombusGridSeed.Name
			}
		case "SpiralRhombusGrid":
			if inferedInstance.SpiralRhombusGrid != nil {
				res.valueString = inferedInstance.SpiralRhombusGrid.Name
			}
		case "SpiralCircleSeed":
			if inferedInstance.SpiralCircleSeed != nil {
				res.valueString = inferedInstance.SpiralCircleSeed.Name
			}
		case "SpiralCircleGrid":
			if inferedInstance.SpiralCircleGrid != nil {
				res.valueString = inferedInstance.SpiralCircleGrid.Name
			}
		case "SpiralCircleFullGrid":
			if inferedInstance.SpiralCircleFullGrid != nil {
				res.valueString = inferedInstance.SpiralCircleFullGrid.Name
			}
		case "SpiralConstructionOuterLineSeed":
			if inferedInstance.SpiralConstructionOuterLineSeed != nil {
				res.valueString = inferedInstance.SpiralConstructionOuterLineSeed.Name
			}
		case "SpiralConstructionInnerLineSeed":
			if inferedInstance.SpiralConstructionInnerLineSeed != nil {
				res.valueString = inferedInstance.SpiralConstructionInnerLineSeed.Name
			}
		case "SpiralConstructionOuterLineGrid":
			if inferedInstance.SpiralConstructionOuterLineGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionOuterLineGrid.Name
			}
		case "SpiralConstructionInnerLineGrid":
			if inferedInstance.SpiralConstructionInnerLineGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionInnerLineGrid.Name
			}
		case "SpiralConstructionCircleGrid":
			if inferedInstance.SpiralConstructionCircleGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionCircleGrid.Name
			}
		case "SpiralConstructionOuterLineFullGrid":
			if inferedInstance.SpiralConstructionOuterLineFullGrid != nil {
				res.valueString = inferedInstance.SpiralConstructionOuterLineFullGrid.Name
			}
		case "SpiralBezierSeed":
			if inferedInstance.SpiralBezierSeed != nil {
				res.valueString = inferedInstance.SpiralBezierSeed.Name
			}
		case "SpiralBezierGrid":
			if inferedInstance.SpiralBezierGrid != nil {
				res.valueString = inferedInstance.SpiralBezierGrid.Name
			}
		case "SpiralBezierFullGrid":
			if inferedInstance.SpiralBezierFullGrid != nil {
				res.valueString = inferedInstance.SpiralBezierFullGrid.Name
			}
		case "SpiralBezierStrength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralBezierStrength)
			res.valueFloat = inferedInstance.SpiralBezierStrength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FrontCurveStack":
			if inferedInstance.FrontCurveStack != nil {
				res.valueString = inferedInstance.FrontCurveStack.Name
			}
		case "NbInterpolationPoints":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbInterpolationPoints)
			res.valueInt = inferedInstance.NbInterpolationPoints
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Fkey":
			if inferedInstance.Fkey != nil {
				res.valueString = inferedInstance.Fkey.Name
			}
		case "FkeySizeRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FkeySizeRatio)
			res.valueFloat = inferedInstance.FkeySizeRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FkeyOriginRelativeX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeX)
			res.valueFloat = inferedInstance.FkeyOriginRelativeX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FkeyOriginRelativeY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeY)
			res.valueFloat = inferedInstance.FkeyOriginRelativeY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "PitchLines":
			if inferedInstance.PitchLines != nil {
				res.valueString = inferedInstance.PitchLines.Name
			}
		case "PitchHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.PitchHeight)
			res.valueFloat = inferedInstance.PitchHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "NbPitchLines":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbPitchLines)
			res.valueInt = inferedInstance.NbPitchLines
			res.GongFieldValueType = GongFieldValueTypeInt
		case "BeatLines":
			if inferedInstance.BeatLines != nil {
				res.valueString = inferedInstance.BeatLines.Name
			}
		case "BeatLinesHeightRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BeatLinesHeightRatio)
			res.valueFloat = inferedInstance.BeatLinesHeightRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "NbBeatLines":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbBeatLines)
			res.valueInt = inferedInstance.NbBeatLines
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NbOfBeatsInTheme":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbOfBeatsInTheme)
			res.valueInt = inferedInstance.NbOfBeatsInTheme
			res.GongFieldValueType = GongFieldValueTypeInt
		case "FirstVoice":
			if inferedInstance.FirstVoice != nil {
				res.valueString = inferedInstance.FirstVoice.Name
			}
		case "FirstVoiceShiftedRigth":
			if inferedInstance.FirstVoiceShiftedRigth != nil {
				res.valueString = inferedInstance.FirstVoiceShiftedRigth.Name
			}
		case "FirstVoiceShiftX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FirstVoiceShiftX)
			res.valueFloat = inferedInstance.FirstVoiceShiftX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FirstVoiceShiftY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FirstVoiceShiftY)
			res.valueFloat = inferedInstance.FirstVoiceShiftY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SecondVoice":
			if inferedInstance.SecondVoice != nil {
				res.valueString = inferedInstance.SecondVoice.Name
			}
		case "SecondVoiceShiftedRight":
			if inferedInstance.SecondVoiceShiftedRight != nil {
				res.valueString = inferedInstance.SecondVoiceShiftedRight.Name
			}
		case "PitchDifference":
			res.valueString = fmt.Sprintf("%d", inferedInstance.PitchDifference)
			res.valueInt = inferedInstance.PitchDifference
			res.GongFieldValueType = GongFieldValueTypeInt
		case "BeatsPerSecond":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BeatsPerSecond)
			res.valueFloat = inferedInstance.BeatsPerSecond
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Level":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Level)
			res.valueFloat = inferedInstance.Level
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FirstVoiceNotes":
			if inferedInstance.FirstVoiceNotes != nil {
				res.valueString = inferedInstance.FirstVoiceNotes.Name
			}
		case "FirstVoiceNotesShiftedRight":
			if inferedInstance.FirstVoiceNotesShiftedRight != nil {
				res.valueString = inferedInstance.FirstVoiceNotesShiftedRight.Name
			}
		case "SecondVoiceNotes":
			if inferedInstance.SecondVoiceNotes != nil {
				res.valueString = inferedInstance.SecondVoiceNotes.Name
			}
		case "SecondVoiceNotesShiftedRight":
			if inferedInstance.SecondVoiceNotesShiftedRight != nil {
				res.valueString = inferedInstance.SecondVoiceNotesShiftedRight.Name
			}
		case "IsMinor":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsMinor)
			res.valueBool = inferedInstance.IsMinor
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ThemeBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ThemeBinaryEncoding)
			res.valueInt = inferedInstance.ThemeBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
		case "OriginX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.OriginX)
			res.valueFloat = inferedInstance.OriginX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "OriginY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.OriginY)
			res.valueFloat = inferedInstance.OriginY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HorizontalAxis":
			if inferedInstance.HorizontalAxis != nil {
				res.valueString = inferedInstance.HorizontalAxis.Name
			}
		case "VerticalAxis":
			if inferedInstance.VerticalAxis != nil {
				res.valueString = inferedInstance.VerticalAxis.Name
			}
		case "SpiralOrigin":
			if inferedInstance.SpiralOrigin != nil {
				res.valueString = inferedInstance.SpiralOrigin.Name
			}
		case "SpiralOriginX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralOriginX)
			res.valueFloat = inferedInstance.SpiralOriginX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SpiralOriginY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralOriginY)
			res.valueFloat = inferedInstance.SpiralOriginY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "OriginCrossWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.OriginCrossWidth)
			res.valueFloat = inferedInstance.OriginCrossWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SpiralRadiusRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SpiralRadiusRatio)
			res.valueFloat = inferedInstance.SpiralRadiusRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ShowSpiralBezierConstruct":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowSpiralBezierConstruct)
			res.valueBool = inferedInstance.ShowSpiralBezierConstruct
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShowInterpolationPoints":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowInterpolationPoints)
			res.valueBool = inferedInstance.ShowInterpolationPoints
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ActualBeatsTemporalShift":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ActualBeatsTemporalShift)
			res.valueInt = inferedInstance.ActualBeatsTemporalShift
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case Rhombus:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SideLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SideLength)
			res.valueFloat = inferedInstance.SideLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "AngleDegree":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AngleDegree)
			res.valueFloat = inferedInstance.AngleDegree
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "InsideAngle":
			res.valueString = fmt.Sprintf("%f", inferedInstance.InsideAngle)
			res.valueFloat = inferedInstance.InsideAngle
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case RhombusGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res.valueString = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Rhombuses":
			for idx, __instance__ := range inferedInstance.Rhombuses {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case ShapeCategory:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case SpiralBezier:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartY)
			res.valueFloat = inferedInstance.StartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartX)
			res.valueFloat = inferedInstance.ControlPointStartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointStartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointStartY)
			res.valueFloat = inferedInstance.ControlPointStartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndX)
			res.valueFloat = inferedInstance.ControlPointEndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ControlPointEndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ControlPointEndY)
			res.valueFloat = inferedInstance.ControlPointEndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case SpiralBezierGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralBeziers":
			for idx, __instance__ := range inferedInstance.SpiralBeziers {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case SpiralCircle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterX)
			res.valueFloat = inferedInstance.CenterX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CenterY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CenterY)
			res.valueFloat = inferedInstance.CenterY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "HasBespokeRadius":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasBespokeRadius)
			res.valueBool = inferedInstance.HasBespokeRadius
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BespopkeRadius":
			res.valueString = fmt.Sprintf("%f", inferedInstance.BespopkeRadius)
			res.valueFloat = inferedInstance.BespopkeRadius
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		case "Pitch":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Pitch)
			res.valueInt = inferedInstance.Pitch
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ShowName":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowName)
			res.valueBool = inferedInstance.ShowName
			res.GongFieldValueType = GongFieldValueTypeBool
		case "BeatNb":
			res.valueString = fmt.Sprintf("%d", inferedInstance.BeatNb)
			res.valueInt = inferedInstance.BeatNb
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Path":
			res.valueString = inferedInstance.Path
		}
	case SpiralCircleGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralRhombusGrid":
			if inferedInstance.SpiralRhombusGrid != nil {
				res.valueString = inferedInstance.SpiralRhombusGrid.Name
			}
		case "SpiralCircles":
			for idx, __instance__ := range inferedInstance.SpiralCircles {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case SpiralLine:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StartY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartY)
			res.valueFloat = inferedInstance.StartY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndY)
			res.valueFloat = inferedInstance.EndY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case SpiralLineGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralLines":
			for idx, __instance__ := range inferedInstance.SpiralLines {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case SpiralOrigin:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case SpiralRhombus:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "X_r0":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r0)
			res.valueFloat = inferedInstance.X_r0
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r0":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r0)
			res.valueFloat = inferedInstance.Y_r0
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_r1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r1)
			res.valueFloat = inferedInstance.X_r1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r1)
			res.valueFloat = inferedInstance.Y_r1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_r2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r2)
			res.valueFloat = inferedInstance.X_r2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r2)
			res.valueFloat = inferedInstance.Y_r2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "X_r3":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X_r3)
			res.valueFloat = inferedInstance.X_r3
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y_r3":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y_r3)
			res.valueFloat = inferedInstance.Y_r3
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	case SpiralRhombusGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "SpiralRhombuses":
			for idx, __instance__ := range inferedInstance.SpiralRhombuses {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case VerticalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsDisplayed":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
			res.valueBool = inferedInstance.IsDisplayed
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res.valueString = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res.valueString = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
			res.valueFloat = inferedInstance.AxisHandleBorderLength
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Axis_Length":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Axis_Length)
			res.valueFloat = inferedInstance.Axis_Length
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = inferedInstance.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = inferedInstance.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
			res.valueFloat = inferedInstance.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = inferedInstance.Transform
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
