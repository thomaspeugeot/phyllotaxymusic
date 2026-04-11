// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	phyllotaxymusic_go "github.com/thomaspeugeot/phyllotaxymusic/go"
)

// can be used for
//
//	days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __Gong__Abs
	_ = strings.Clone("")
)

const (
	ProbeTreeSidebarSuffix           = ":sidebar of the probe"
	ProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	ProbeTableSuffix                 = ":table of the probe"
	ProbeNotificationTableSuffix     = ":notification table of the probe"
	ProbeFormSuffix                  = ":form of the probe"
	ProbeSplitSuffix                 = ":probe of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
}

// Stage enables storage of staged instances
type Stage struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between
	// succesive commit
	isInDeltaMode bool

	// gongMarshallingMode set the marshalling mode
	gongMarshallingMode GongMarshallingMode
	// some stages have semantic rules that forbids them to be empty
	// like for git, the commit #0 (genesis commit) cannot be rolled back
	isWithGenesisCommit bool

	// insertion point for definition of arrays registering instances
	Axiss                map[*Axis]struct{}
	Axiss_instance       map[*Axis]*Axis
	Axiss_mapString      map[string]*Axis
	AxisOrder            uint
	Axis_stagedOrder     map[*Axis]uint
	Axiss_reference      map[*Axis]*Axis
	Axiss_referenceOrder map[*Axis]uint

	// insertion point for slice of pointers maps
	OnAfterAxisCreateCallback OnAfterCreateInterface[Axis]
	OnAfterAxisUpdateCallback OnAfterUpdateInterface[Axis]
	OnAfterAxisDeleteCallback OnAfterDeleteInterface[Axis]
	OnAfterAxisReadCallback   OnAfterReadInterface[Axis]

	AxisGrids                map[*AxisGrid]struct{}
	AxisGrids_instance       map[*AxisGrid]*AxisGrid
	AxisGrids_mapString      map[string]*AxisGrid
	AxisGridOrder            uint
	AxisGrid_stagedOrder     map[*AxisGrid]uint
	AxisGrids_reference      map[*AxisGrid]*AxisGrid
	AxisGrids_referenceOrder map[*AxisGrid]uint

	// insertion point for slice of pointers maps
	AxisGrid_Axiss_reverseMap map[*Axis]*AxisGrid

	OnAfterAxisGridCreateCallback OnAfterCreateInterface[AxisGrid]
	OnAfterAxisGridUpdateCallback OnAfterUpdateInterface[AxisGrid]
	OnAfterAxisGridDeleteCallback OnAfterDeleteInterface[AxisGrid]
	OnAfterAxisGridReadCallback   OnAfterReadInterface[AxisGrid]

	Beziers                map[*Bezier]struct{}
	Beziers_instance       map[*Bezier]*Bezier
	Beziers_mapString      map[string]*Bezier
	BezierOrder            uint
	Bezier_stagedOrder     map[*Bezier]uint
	Beziers_reference      map[*Bezier]*Bezier
	Beziers_referenceOrder map[*Bezier]uint

	// insertion point for slice of pointers maps
	OnAfterBezierCreateCallback OnAfterCreateInterface[Bezier]
	OnAfterBezierUpdateCallback OnAfterUpdateInterface[Bezier]
	OnAfterBezierDeleteCallback OnAfterDeleteInterface[Bezier]
	OnAfterBezierReadCallback   OnAfterReadInterface[Bezier]

	BezierGrids                map[*BezierGrid]struct{}
	BezierGrids_instance       map[*BezierGrid]*BezierGrid
	BezierGrids_mapString      map[string]*BezierGrid
	BezierGridOrder            uint
	BezierGrid_stagedOrder     map[*BezierGrid]uint
	BezierGrids_reference      map[*BezierGrid]*BezierGrid
	BezierGrids_referenceOrder map[*BezierGrid]uint

	// insertion point for slice of pointers maps
	BezierGrid_Beziers_reverseMap map[*Bezier]*BezierGrid

	OnAfterBezierGridCreateCallback OnAfterCreateInterface[BezierGrid]
	OnAfterBezierGridUpdateCallback OnAfterUpdateInterface[BezierGrid]
	OnAfterBezierGridDeleteCallback OnAfterDeleteInterface[BezierGrid]
	OnAfterBezierGridReadCallback   OnAfterReadInterface[BezierGrid]

	BezierGridStacks                map[*BezierGridStack]struct{}
	BezierGridStacks_instance       map[*BezierGridStack]*BezierGridStack
	BezierGridStacks_mapString      map[string]*BezierGridStack
	BezierGridStackOrder            uint
	BezierGridStack_stagedOrder     map[*BezierGridStack]uint
	BezierGridStacks_reference      map[*BezierGridStack]*BezierGridStack
	BezierGridStacks_referenceOrder map[*BezierGridStack]uint

	// insertion point for slice of pointers maps
	BezierGridStack_BezierGrids_reverseMap map[*BezierGrid]*BezierGridStack

	OnAfterBezierGridStackCreateCallback OnAfterCreateInterface[BezierGridStack]
	OnAfterBezierGridStackUpdateCallback OnAfterUpdateInterface[BezierGridStack]
	OnAfterBezierGridStackDeleteCallback OnAfterDeleteInterface[BezierGridStack]
	OnAfterBezierGridStackReadCallback   OnAfterReadInterface[BezierGridStack]

	Chapters                map[*Chapter]struct{}
	Chapters_instance       map[*Chapter]*Chapter
	Chapters_mapString      map[string]*Chapter
	ChapterOrder            uint
	Chapter_stagedOrder     map[*Chapter]uint
	Chapters_reference      map[*Chapter]*Chapter
	Chapters_referenceOrder map[*Chapter]uint

	// insertion point for slice of pointers maps
	OnAfterChapterCreateCallback OnAfterCreateInterface[Chapter]
	OnAfterChapterUpdateCallback OnAfterUpdateInterface[Chapter]
	OnAfterChapterDeleteCallback OnAfterDeleteInterface[Chapter]
	OnAfterChapterReadCallback   OnAfterReadInterface[Chapter]

	Circles                map[*Circle]struct{}
	Circles_instance       map[*Circle]*Circle
	Circles_mapString      map[string]*Circle
	CircleOrder            uint
	Circle_stagedOrder     map[*Circle]uint
	Circles_reference      map[*Circle]*Circle
	Circles_referenceOrder map[*Circle]uint

	// insertion point for slice of pointers maps
	OnAfterCircleCreateCallback OnAfterCreateInterface[Circle]
	OnAfterCircleUpdateCallback OnAfterUpdateInterface[Circle]
	OnAfterCircleDeleteCallback OnAfterDeleteInterface[Circle]
	OnAfterCircleReadCallback   OnAfterReadInterface[Circle]

	CircleGrids                map[*CircleGrid]struct{}
	CircleGrids_instance       map[*CircleGrid]*CircleGrid
	CircleGrids_mapString      map[string]*CircleGrid
	CircleGridOrder            uint
	CircleGrid_stagedOrder     map[*CircleGrid]uint
	CircleGrids_reference      map[*CircleGrid]*CircleGrid
	CircleGrids_referenceOrder map[*CircleGrid]uint

	// insertion point for slice of pointers maps
	CircleGrid_Circles_reverseMap map[*Circle]*CircleGrid

	OnAfterCircleGridCreateCallback OnAfterCreateInterface[CircleGrid]
	OnAfterCircleGridUpdateCallback OnAfterUpdateInterface[CircleGrid]
	OnAfterCircleGridDeleteCallback OnAfterDeleteInterface[CircleGrid]
	OnAfterCircleGridReadCallback   OnAfterReadInterface[CircleGrid]

	Contents                map[*Content]struct{}
	Contents_instance       map[*Content]*Content
	Contents_mapString      map[string]*Content
	ContentOrder            uint
	Content_stagedOrder     map[*Content]uint
	Contents_reference      map[*Content]*Content
	Contents_referenceOrder map[*Content]uint

	// insertion point for slice of pointers maps
	Content_Chapters_reverseMap map[*Chapter]*Content

	OnAfterContentCreateCallback OnAfterCreateInterface[Content]
	OnAfterContentUpdateCallback OnAfterUpdateInterface[Content]
	OnAfterContentDeleteCallback OnAfterDeleteInterface[Content]
	OnAfterContentReadCallback   OnAfterReadInterface[Content]

	ExportToMusicxmls                map[*ExportToMusicxml]struct{}
	ExportToMusicxmls_instance       map[*ExportToMusicxml]*ExportToMusicxml
	ExportToMusicxmls_mapString      map[string]*ExportToMusicxml
	ExportToMusicxmlOrder            uint
	ExportToMusicxml_stagedOrder     map[*ExportToMusicxml]uint
	ExportToMusicxmls_reference      map[*ExportToMusicxml]*ExportToMusicxml
	ExportToMusicxmls_referenceOrder map[*ExportToMusicxml]uint

	// insertion point for slice of pointers maps
	OnAfterExportToMusicxmlCreateCallback OnAfterCreateInterface[ExportToMusicxml]
	OnAfterExportToMusicxmlUpdateCallback OnAfterUpdateInterface[ExportToMusicxml]
	OnAfterExportToMusicxmlDeleteCallback OnAfterDeleteInterface[ExportToMusicxml]
	OnAfterExportToMusicxmlReadCallback   OnAfterReadInterface[ExportToMusicxml]

	FrontCurves                map[*FrontCurve]struct{}
	FrontCurves_instance       map[*FrontCurve]*FrontCurve
	FrontCurves_mapString      map[string]*FrontCurve
	FrontCurveOrder            uint
	FrontCurve_stagedOrder     map[*FrontCurve]uint
	FrontCurves_reference      map[*FrontCurve]*FrontCurve
	FrontCurves_referenceOrder map[*FrontCurve]uint

	// insertion point for slice of pointers maps
	OnAfterFrontCurveCreateCallback OnAfterCreateInterface[FrontCurve]
	OnAfterFrontCurveUpdateCallback OnAfterUpdateInterface[FrontCurve]
	OnAfterFrontCurveDeleteCallback OnAfterDeleteInterface[FrontCurve]
	OnAfterFrontCurveReadCallback   OnAfterReadInterface[FrontCurve]

	FrontCurveStacks                map[*FrontCurveStack]struct{}
	FrontCurveStacks_instance       map[*FrontCurveStack]*FrontCurveStack
	FrontCurveStacks_mapString      map[string]*FrontCurveStack
	FrontCurveStackOrder            uint
	FrontCurveStack_stagedOrder     map[*FrontCurveStack]uint
	FrontCurveStacks_reference      map[*FrontCurveStack]*FrontCurveStack
	FrontCurveStacks_referenceOrder map[*FrontCurveStack]uint

	// insertion point for slice of pointers maps
	FrontCurveStack_FrontCurves_reverseMap map[*FrontCurve]*FrontCurveStack

	FrontCurveStack_SpiralCircles_reverseMap map[*SpiralCircle]*FrontCurveStack

	OnAfterFrontCurveStackCreateCallback OnAfterCreateInterface[FrontCurveStack]
	OnAfterFrontCurveStackUpdateCallback OnAfterUpdateInterface[FrontCurveStack]
	OnAfterFrontCurveStackDeleteCallback OnAfterDeleteInterface[FrontCurveStack]
	OnAfterFrontCurveStackReadCallback   OnAfterReadInterface[FrontCurveStack]

	HorizontalAxiss                map[*HorizontalAxis]struct{}
	HorizontalAxiss_instance       map[*HorizontalAxis]*HorizontalAxis
	HorizontalAxiss_mapString      map[string]*HorizontalAxis
	HorizontalAxisOrder            uint
	HorizontalAxis_stagedOrder     map[*HorizontalAxis]uint
	HorizontalAxiss_reference      map[*HorizontalAxis]*HorizontalAxis
	HorizontalAxiss_referenceOrder map[*HorizontalAxis]uint

	// insertion point for slice of pointers maps
	OnAfterHorizontalAxisCreateCallback OnAfterCreateInterface[HorizontalAxis]
	OnAfterHorizontalAxisUpdateCallback OnAfterUpdateInterface[HorizontalAxis]
	OnAfterHorizontalAxisDeleteCallback OnAfterDeleteInterface[HorizontalAxis]
	OnAfterHorizontalAxisReadCallback   OnAfterReadInterface[HorizontalAxis]

	Keys                map[*Key]struct{}
	Keys_instance       map[*Key]*Key
	Keys_mapString      map[string]*Key
	KeyOrder            uint
	Key_stagedOrder     map[*Key]uint
	Keys_reference      map[*Key]*Key
	Keys_referenceOrder map[*Key]uint

	// insertion point for slice of pointers maps
	OnAfterKeyCreateCallback OnAfterCreateInterface[Key]
	OnAfterKeyUpdateCallback OnAfterUpdateInterface[Key]
	OnAfterKeyDeleteCallback OnAfterDeleteInterface[Key]
	OnAfterKeyReadCallback   OnAfterReadInterface[Key]

	Parameters                map[*Parameter]struct{}
	Parameters_instance       map[*Parameter]*Parameter
	Parameters_mapString      map[string]*Parameter
	ParameterOrder            uint
	Parameter_stagedOrder     map[*Parameter]uint
	Parameters_reference      map[*Parameter]*Parameter
	Parameters_referenceOrder map[*Parameter]uint

	// insertion point for slice of pointers maps
	OnAfterParameterCreateCallback OnAfterCreateInterface[Parameter]
	OnAfterParameterUpdateCallback OnAfterUpdateInterface[Parameter]
	OnAfterParameterDeleteCallback OnAfterDeleteInterface[Parameter]
	OnAfterParameterReadCallback   OnAfterReadInterface[Parameter]

	Rhombuss                map[*Rhombus]struct{}
	Rhombuss_instance       map[*Rhombus]*Rhombus
	Rhombuss_mapString      map[string]*Rhombus
	RhombusOrder            uint
	Rhombus_stagedOrder     map[*Rhombus]uint
	Rhombuss_reference      map[*Rhombus]*Rhombus
	Rhombuss_referenceOrder map[*Rhombus]uint

	// insertion point for slice of pointers maps
	OnAfterRhombusCreateCallback OnAfterCreateInterface[Rhombus]
	OnAfterRhombusUpdateCallback OnAfterUpdateInterface[Rhombus]
	OnAfterRhombusDeleteCallback OnAfterDeleteInterface[Rhombus]
	OnAfterRhombusReadCallback   OnAfterReadInterface[Rhombus]

	RhombusGrids                map[*RhombusGrid]struct{}
	RhombusGrids_instance       map[*RhombusGrid]*RhombusGrid
	RhombusGrids_mapString      map[string]*RhombusGrid
	RhombusGridOrder            uint
	RhombusGrid_stagedOrder     map[*RhombusGrid]uint
	RhombusGrids_reference      map[*RhombusGrid]*RhombusGrid
	RhombusGrids_referenceOrder map[*RhombusGrid]uint

	// insertion point for slice of pointers maps
	RhombusGrid_Rhombuses_reverseMap map[*Rhombus]*RhombusGrid

	OnAfterRhombusGridCreateCallback OnAfterCreateInterface[RhombusGrid]
	OnAfterRhombusGridUpdateCallback OnAfterUpdateInterface[RhombusGrid]
	OnAfterRhombusGridDeleteCallback OnAfterDeleteInterface[RhombusGrid]
	OnAfterRhombusGridReadCallback   OnAfterReadInterface[RhombusGrid]

	ShapeCategorys                map[*ShapeCategory]struct{}
	ShapeCategorys_instance       map[*ShapeCategory]*ShapeCategory
	ShapeCategorys_mapString      map[string]*ShapeCategory
	ShapeCategoryOrder            uint
	ShapeCategory_stagedOrder     map[*ShapeCategory]uint
	ShapeCategorys_reference      map[*ShapeCategory]*ShapeCategory
	ShapeCategorys_referenceOrder map[*ShapeCategory]uint

	// insertion point for slice of pointers maps
	OnAfterShapeCategoryCreateCallback OnAfterCreateInterface[ShapeCategory]
	OnAfterShapeCategoryUpdateCallback OnAfterUpdateInterface[ShapeCategory]
	OnAfterShapeCategoryDeleteCallback OnAfterDeleteInterface[ShapeCategory]
	OnAfterShapeCategoryReadCallback   OnAfterReadInterface[ShapeCategory]

	SpiralBeziers                map[*SpiralBezier]struct{}
	SpiralBeziers_instance       map[*SpiralBezier]*SpiralBezier
	SpiralBeziers_mapString      map[string]*SpiralBezier
	SpiralBezierOrder            uint
	SpiralBezier_stagedOrder     map[*SpiralBezier]uint
	SpiralBeziers_reference      map[*SpiralBezier]*SpiralBezier
	SpiralBeziers_referenceOrder map[*SpiralBezier]uint

	// insertion point for slice of pointers maps
	OnAfterSpiralBezierCreateCallback OnAfterCreateInterface[SpiralBezier]
	OnAfterSpiralBezierUpdateCallback OnAfterUpdateInterface[SpiralBezier]
	OnAfterSpiralBezierDeleteCallback OnAfterDeleteInterface[SpiralBezier]
	OnAfterSpiralBezierReadCallback   OnAfterReadInterface[SpiralBezier]

	SpiralBezierGrids                map[*SpiralBezierGrid]struct{}
	SpiralBezierGrids_instance       map[*SpiralBezierGrid]*SpiralBezierGrid
	SpiralBezierGrids_mapString      map[string]*SpiralBezierGrid
	SpiralBezierGridOrder            uint
	SpiralBezierGrid_stagedOrder     map[*SpiralBezierGrid]uint
	SpiralBezierGrids_reference      map[*SpiralBezierGrid]*SpiralBezierGrid
	SpiralBezierGrids_referenceOrder map[*SpiralBezierGrid]uint

	// insertion point for slice of pointers maps
	SpiralBezierGrid_SpiralBeziers_reverseMap map[*SpiralBezier]*SpiralBezierGrid

	OnAfterSpiralBezierGridCreateCallback OnAfterCreateInterface[SpiralBezierGrid]
	OnAfterSpiralBezierGridUpdateCallback OnAfterUpdateInterface[SpiralBezierGrid]
	OnAfterSpiralBezierGridDeleteCallback OnAfterDeleteInterface[SpiralBezierGrid]
	OnAfterSpiralBezierGridReadCallback   OnAfterReadInterface[SpiralBezierGrid]

	SpiralCircles                map[*SpiralCircle]struct{}
	SpiralCircles_instance       map[*SpiralCircle]*SpiralCircle
	SpiralCircles_mapString      map[string]*SpiralCircle
	SpiralCircleOrder            uint
	SpiralCircle_stagedOrder     map[*SpiralCircle]uint
	SpiralCircles_reference      map[*SpiralCircle]*SpiralCircle
	SpiralCircles_referenceOrder map[*SpiralCircle]uint

	// insertion point for slice of pointers maps
	OnAfterSpiralCircleCreateCallback OnAfterCreateInterface[SpiralCircle]
	OnAfterSpiralCircleUpdateCallback OnAfterUpdateInterface[SpiralCircle]
	OnAfterSpiralCircleDeleteCallback OnAfterDeleteInterface[SpiralCircle]
	OnAfterSpiralCircleReadCallback   OnAfterReadInterface[SpiralCircle]

	SpiralCircleGrids                map[*SpiralCircleGrid]struct{}
	SpiralCircleGrids_instance       map[*SpiralCircleGrid]*SpiralCircleGrid
	SpiralCircleGrids_mapString      map[string]*SpiralCircleGrid
	SpiralCircleGridOrder            uint
	SpiralCircleGrid_stagedOrder     map[*SpiralCircleGrid]uint
	SpiralCircleGrids_reference      map[*SpiralCircleGrid]*SpiralCircleGrid
	SpiralCircleGrids_referenceOrder map[*SpiralCircleGrid]uint

	// insertion point for slice of pointers maps
	SpiralCircleGrid_SpiralCircles_reverseMap map[*SpiralCircle]*SpiralCircleGrid

	OnAfterSpiralCircleGridCreateCallback OnAfterCreateInterface[SpiralCircleGrid]
	OnAfterSpiralCircleGridUpdateCallback OnAfterUpdateInterface[SpiralCircleGrid]
	OnAfterSpiralCircleGridDeleteCallback OnAfterDeleteInterface[SpiralCircleGrid]
	OnAfterSpiralCircleGridReadCallback   OnAfterReadInterface[SpiralCircleGrid]

	SpiralLines                map[*SpiralLine]struct{}
	SpiralLines_instance       map[*SpiralLine]*SpiralLine
	SpiralLines_mapString      map[string]*SpiralLine
	SpiralLineOrder            uint
	SpiralLine_stagedOrder     map[*SpiralLine]uint
	SpiralLines_reference      map[*SpiralLine]*SpiralLine
	SpiralLines_referenceOrder map[*SpiralLine]uint

	// insertion point for slice of pointers maps
	OnAfterSpiralLineCreateCallback OnAfterCreateInterface[SpiralLine]
	OnAfterSpiralLineUpdateCallback OnAfterUpdateInterface[SpiralLine]
	OnAfterSpiralLineDeleteCallback OnAfterDeleteInterface[SpiralLine]
	OnAfterSpiralLineReadCallback   OnAfterReadInterface[SpiralLine]

	SpiralLineGrids                map[*SpiralLineGrid]struct{}
	SpiralLineGrids_instance       map[*SpiralLineGrid]*SpiralLineGrid
	SpiralLineGrids_mapString      map[string]*SpiralLineGrid
	SpiralLineGridOrder            uint
	SpiralLineGrid_stagedOrder     map[*SpiralLineGrid]uint
	SpiralLineGrids_reference      map[*SpiralLineGrid]*SpiralLineGrid
	SpiralLineGrids_referenceOrder map[*SpiralLineGrid]uint

	// insertion point for slice of pointers maps
	SpiralLineGrid_SpiralLines_reverseMap map[*SpiralLine]*SpiralLineGrid

	OnAfterSpiralLineGridCreateCallback OnAfterCreateInterface[SpiralLineGrid]
	OnAfterSpiralLineGridUpdateCallback OnAfterUpdateInterface[SpiralLineGrid]
	OnAfterSpiralLineGridDeleteCallback OnAfterDeleteInterface[SpiralLineGrid]
	OnAfterSpiralLineGridReadCallback   OnAfterReadInterface[SpiralLineGrid]

	SpiralOrigins                map[*SpiralOrigin]struct{}
	SpiralOrigins_instance       map[*SpiralOrigin]*SpiralOrigin
	SpiralOrigins_mapString      map[string]*SpiralOrigin
	SpiralOriginOrder            uint
	SpiralOrigin_stagedOrder     map[*SpiralOrigin]uint
	SpiralOrigins_reference      map[*SpiralOrigin]*SpiralOrigin
	SpiralOrigins_referenceOrder map[*SpiralOrigin]uint

	// insertion point for slice of pointers maps
	OnAfterSpiralOriginCreateCallback OnAfterCreateInterface[SpiralOrigin]
	OnAfterSpiralOriginUpdateCallback OnAfterUpdateInterface[SpiralOrigin]
	OnAfterSpiralOriginDeleteCallback OnAfterDeleteInterface[SpiralOrigin]
	OnAfterSpiralOriginReadCallback   OnAfterReadInterface[SpiralOrigin]

	SpiralRhombuss                map[*SpiralRhombus]struct{}
	SpiralRhombuss_instance       map[*SpiralRhombus]*SpiralRhombus
	SpiralRhombuss_mapString      map[string]*SpiralRhombus
	SpiralRhombusOrder            uint
	SpiralRhombus_stagedOrder     map[*SpiralRhombus]uint
	SpiralRhombuss_reference      map[*SpiralRhombus]*SpiralRhombus
	SpiralRhombuss_referenceOrder map[*SpiralRhombus]uint

	// insertion point for slice of pointers maps
	OnAfterSpiralRhombusCreateCallback OnAfterCreateInterface[SpiralRhombus]
	OnAfterSpiralRhombusUpdateCallback OnAfterUpdateInterface[SpiralRhombus]
	OnAfterSpiralRhombusDeleteCallback OnAfterDeleteInterface[SpiralRhombus]
	OnAfterSpiralRhombusReadCallback   OnAfterReadInterface[SpiralRhombus]

	SpiralRhombusGrids                map[*SpiralRhombusGrid]struct{}
	SpiralRhombusGrids_instance       map[*SpiralRhombusGrid]*SpiralRhombusGrid
	SpiralRhombusGrids_mapString      map[string]*SpiralRhombusGrid
	SpiralRhombusGridOrder            uint
	SpiralRhombusGrid_stagedOrder     map[*SpiralRhombusGrid]uint
	SpiralRhombusGrids_reference      map[*SpiralRhombusGrid]*SpiralRhombusGrid
	SpiralRhombusGrids_referenceOrder map[*SpiralRhombusGrid]uint

	// insertion point for slice of pointers maps
	SpiralRhombusGrid_SpiralRhombuses_reverseMap map[*SpiralRhombus]*SpiralRhombusGrid

	OnAfterSpiralRhombusGridCreateCallback OnAfterCreateInterface[SpiralRhombusGrid]
	OnAfterSpiralRhombusGridUpdateCallback OnAfterUpdateInterface[SpiralRhombusGrid]
	OnAfterSpiralRhombusGridDeleteCallback OnAfterDeleteInterface[SpiralRhombusGrid]
	OnAfterSpiralRhombusGridReadCallback   OnAfterReadInterface[SpiralRhombusGrid]

	VerticalAxiss                map[*VerticalAxis]struct{}
	VerticalAxiss_instance       map[*VerticalAxis]*VerticalAxis
	VerticalAxiss_mapString      map[string]*VerticalAxis
	VerticalAxisOrder            uint
	VerticalAxis_stagedOrder     map[*VerticalAxis]uint
	VerticalAxiss_reference      map[*VerticalAxis]*VerticalAxis
	VerticalAxiss_referenceOrder map[*VerticalAxis]uint

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

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

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
	// end of insertion point

	NamedStructs []*NamedStruct

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]ModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF

	forwardCommits  []string
	backwardCommits []string

	// when navigating the commit history
	// navigationMode is set to Navigating
	navigationMode gongStageNavigationMode
	commitsBehind  int // the number of commits the stage is behind the front of the history

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool

	modified bool

	lock sync.RWMutex
}

func (s *Stage) SetGongMarshallingMode(mode GongMarshallingMode) {
	s.gongMarshallingMode = mode
}

func (s *Stage) GetGongMarshallingMode() GongMarshallingMode {
	return s.gongMarshallingMode
}

func (s *Stage) SetIsWithGenesisCommit(isWithGenesisCommit bool) {
	s.isWithGenesisCommit = isWithGenesisCommit
}

func (s *Stage) GetIsWithGenesisCommit() bool {
	return s.isWithGenesisCommit
}

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.beforeCommitHooks = append(s.beforeCommitHooks, hook)
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.afterCommitHooks = append(s.afterCommitHooks, hook)
}

type gongStageNavigationMode string

const (
	GongNavigationModeNormal gongStageNavigationMode = "Normal"
	// when the mode is navigating, each commit backward and forward
	// it is possible to go apply the nbCommitsBackward forward commits
	GongNavigationModeNavigating gongStageNavigationMode = "Navigating"
)

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	if len(stage.backwardCommits) == 0 {
		return errors.New("no backward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	if stage.isWithGenesisCommit && stage.commitsBehind >= len(stage.backwardCommits)-1 {
		return errors.New("cannot rollback genesis commit")
	}

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	stage.isApplyingBackwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingBackwardCommit = false
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetForwardCommits() []string {
	return stage.forwardCommits
}

func (stage *Stage) GetBackwardCommits() []string {
	return stage.backwardCommits
}

func (stage *Stage) ApplyForwardCommit() error {
	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.commitsBehind == 0 {
		return errors.New("no more forward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	commitToApply := stage.forwardCommits[len(stage.forwardCommits)-1-stage.commitsBehind+1]

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind--
	stage.isApplyingForwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingForwardCommit = false
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}
	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetCommitsBehind() int {
	return stage.commitsBehind
}

func (stage *Stage) Lock() {
	stage.lock.Lock()
}

func (stage *Stage) Unlock() {
	stage.lock.Unlock()
}

func (stage *Stage) RLock() {
	stage.lock.RLock()
}

func (stage *Stage) RUnlock() {
	stage.lock.RUnlock()
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	newCommitsLen := len(stage.forwardCommits) - stage.GetCommitsBehind()

	stage.forwardCommits = stage.forwardCommits[:newCommitsLen]
	stage.backwardCommits = stage.backwardCommits[:newCommitsLen]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.modified = true
	stage.isSquashing = true

	// insertion point for clear references
	stage.Axiss_reference = make(map[*Axis]*Axis)
	stage.Axiss_instance = make(map[*Axis]*Axis)
	stage.Axiss_referenceOrder = make(map[*Axis]uint)

	stage.AxisGrids_reference = make(map[*AxisGrid]*AxisGrid)
	stage.AxisGrids_instance = make(map[*AxisGrid]*AxisGrid)
	stage.AxisGrids_referenceOrder = make(map[*AxisGrid]uint)

	stage.Beziers_reference = make(map[*Bezier]*Bezier)
	stage.Beziers_instance = make(map[*Bezier]*Bezier)
	stage.Beziers_referenceOrder = make(map[*Bezier]uint)

	stage.BezierGrids_reference = make(map[*BezierGrid]*BezierGrid)
	stage.BezierGrids_instance = make(map[*BezierGrid]*BezierGrid)
	stage.BezierGrids_referenceOrder = make(map[*BezierGrid]uint)

	stage.BezierGridStacks_reference = make(map[*BezierGridStack]*BezierGridStack)
	stage.BezierGridStacks_instance = make(map[*BezierGridStack]*BezierGridStack)
	stage.BezierGridStacks_referenceOrder = make(map[*BezierGridStack]uint)

	stage.Chapters_reference = make(map[*Chapter]*Chapter)
	stage.Chapters_instance = make(map[*Chapter]*Chapter)
	stage.Chapters_referenceOrder = make(map[*Chapter]uint)

	stage.Circles_reference = make(map[*Circle]*Circle)
	stage.Circles_instance = make(map[*Circle]*Circle)
	stage.Circles_referenceOrder = make(map[*Circle]uint)

	stage.CircleGrids_reference = make(map[*CircleGrid]*CircleGrid)
	stage.CircleGrids_instance = make(map[*CircleGrid]*CircleGrid)
	stage.CircleGrids_referenceOrder = make(map[*CircleGrid]uint)

	stage.Contents_reference = make(map[*Content]*Content)
	stage.Contents_instance = make(map[*Content]*Content)
	stage.Contents_referenceOrder = make(map[*Content]uint)

	stage.ExportToMusicxmls_reference = make(map[*ExportToMusicxml]*ExportToMusicxml)
	stage.ExportToMusicxmls_instance = make(map[*ExportToMusicxml]*ExportToMusicxml)
	stage.ExportToMusicxmls_referenceOrder = make(map[*ExportToMusicxml]uint)

	stage.FrontCurves_reference = make(map[*FrontCurve]*FrontCurve)
	stage.FrontCurves_instance = make(map[*FrontCurve]*FrontCurve)
	stage.FrontCurves_referenceOrder = make(map[*FrontCurve]uint)

	stage.FrontCurveStacks_reference = make(map[*FrontCurveStack]*FrontCurveStack)
	stage.FrontCurveStacks_instance = make(map[*FrontCurveStack]*FrontCurveStack)
	stage.FrontCurveStacks_referenceOrder = make(map[*FrontCurveStack]uint)

	stage.HorizontalAxiss_reference = make(map[*HorizontalAxis]*HorizontalAxis)
	stage.HorizontalAxiss_instance = make(map[*HorizontalAxis]*HorizontalAxis)
	stage.HorizontalAxiss_referenceOrder = make(map[*HorizontalAxis]uint)

	stage.Keys_reference = make(map[*Key]*Key)
	stage.Keys_instance = make(map[*Key]*Key)
	stage.Keys_referenceOrder = make(map[*Key]uint)

	stage.Parameters_reference = make(map[*Parameter]*Parameter)
	stage.Parameters_instance = make(map[*Parameter]*Parameter)
	stage.Parameters_referenceOrder = make(map[*Parameter]uint)

	stage.Rhombuss_reference = make(map[*Rhombus]*Rhombus)
	stage.Rhombuss_instance = make(map[*Rhombus]*Rhombus)
	stage.Rhombuss_referenceOrder = make(map[*Rhombus]uint)

	stage.RhombusGrids_reference = make(map[*RhombusGrid]*RhombusGrid)
	stage.RhombusGrids_instance = make(map[*RhombusGrid]*RhombusGrid)
	stage.RhombusGrids_referenceOrder = make(map[*RhombusGrid]uint)

	stage.ShapeCategorys_reference = make(map[*ShapeCategory]*ShapeCategory)
	stage.ShapeCategorys_instance = make(map[*ShapeCategory]*ShapeCategory)
	stage.ShapeCategorys_referenceOrder = make(map[*ShapeCategory]uint)

	stage.SpiralBeziers_reference = make(map[*SpiralBezier]*SpiralBezier)
	stage.SpiralBeziers_instance = make(map[*SpiralBezier]*SpiralBezier)
	stage.SpiralBeziers_referenceOrder = make(map[*SpiralBezier]uint)

	stage.SpiralBezierGrids_reference = make(map[*SpiralBezierGrid]*SpiralBezierGrid)
	stage.SpiralBezierGrids_instance = make(map[*SpiralBezierGrid]*SpiralBezierGrid)
	stage.SpiralBezierGrids_referenceOrder = make(map[*SpiralBezierGrid]uint)

	stage.SpiralCircles_reference = make(map[*SpiralCircle]*SpiralCircle)
	stage.SpiralCircles_instance = make(map[*SpiralCircle]*SpiralCircle)
	stage.SpiralCircles_referenceOrder = make(map[*SpiralCircle]uint)

	stage.SpiralCircleGrids_reference = make(map[*SpiralCircleGrid]*SpiralCircleGrid)
	stage.SpiralCircleGrids_instance = make(map[*SpiralCircleGrid]*SpiralCircleGrid)
	stage.SpiralCircleGrids_referenceOrder = make(map[*SpiralCircleGrid]uint)

	stage.SpiralLines_reference = make(map[*SpiralLine]*SpiralLine)
	stage.SpiralLines_instance = make(map[*SpiralLine]*SpiralLine)
	stage.SpiralLines_referenceOrder = make(map[*SpiralLine]uint)

	stage.SpiralLineGrids_reference = make(map[*SpiralLineGrid]*SpiralLineGrid)
	stage.SpiralLineGrids_instance = make(map[*SpiralLineGrid]*SpiralLineGrid)
	stage.SpiralLineGrids_referenceOrder = make(map[*SpiralLineGrid]uint)

	stage.SpiralOrigins_reference = make(map[*SpiralOrigin]*SpiralOrigin)
	stage.SpiralOrigins_instance = make(map[*SpiralOrigin]*SpiralOrigin)
	stage.SpiralOrigins_referenceOrder = make(map[*SpiralOrigin]uint)

	stage.SpiralRhombuss_reference = make(map[*SpiralRhombus]*SpiralRhombus)
	stage.SpiralRhombuss_instance = make(map[*SpiralRhombus]*SpiralRhombus)
	stage.SpiralRhombuss_referenceOrder = make(map[*SpiralRhombus]uint)

	stage.SpiralRhombusGrids_reference = make(map[*SpiralRhombusGrid]*SpiralRhombusGrid)
	stage.SpiralRhombusGrids_instance = make(map[*SpiralRhombusGrid]*SpiralRhombusGrid)
	stage.SpiralRhombusGrids_referenceOrder = make(map[*SpiralRhombusGrid]uint)

	stage.VerticalAxiss_reference = make(map[*VerticalAxis]*VerticalAxis)
	stage.VerticalAxiss_instance = make(map[*VerticalAxis]*VerticalAxis)
	stage.VerticalAxiss_referenceOrder = make(map[*VerticalAxis]uint)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}

	stage.isSquashing = false
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	var maxAxisOrder uint
	var foundAxis bool
	for _, order := range stage.Axis_stagedOrder {
		if !foundAxis || order > maxAxisOrder {
			maxAxisOrder = order
			foundAxis = true
		}
	}
	if foundAxis {
		stage.AxisOrder = maxAxisOrder + 1
	} else {
		stage.AxisOrder = 0
	}

	var maxAxisGridOrder uint
	var foundAxisGrid bool
	for _, order := range stage.AxisGrid_stagedOrder {
		if !foundAxisGrid || order > maxAxisGridOrder {
			maxAxisGridOrder = order
			foundAxisGrid = true
		}
	}
	if foundAxisGrid {
		stage.AxisGridOrder = maxAxisGridOrder + 1
	} else {
		stage.AxisGridOrder = 0
	}

	var maxBezierOrder uint
	var foundBezier bool
	for _, order := range stage.Bezier_stagedOrder {
		if !foundBezier || order > maxBezierOrder {
			maxBezierOrder = order
			foundBezier = true
		}
	}
	if foundBezier {
		stage.BezierOrder = maxBezierOrder + 1
	} else {
		stage.BezierOrder = 0
	}

	var maxBezierGridOrder uint
	var foundBezierGrid bool
	for _, order := range stage.BezierGrid_stagedOrder {
		if !foundBezierGrid || order > maxBezierGridOrder {
			maxBezierGridOrder = order
			foundBezierGrid = true
		}
	}
	if foundBezierGrid {
		stage.BezierGridOrder = maxBezierGridOrder + 1
	} else {
		stage.BezierGridOrder = 0
	}

	var maxBezierGridStackOrder uint
	var foundBezierGridStack bool
	for _, order := range stage.BezierGridStack_stagedOrder {
		if !foundBezierGridStack || order > maxBezierGridStackOrder {
			maxBezierGridStackOrder = order
			foundBezierGridStack = true
		}
	}
	if foundBezierGridStack {
		stage.BezierGridStackOrder = maxBezierGridStackOrder + 1
	} else {
		stage.BezierGridStackOrder = 0
	}

	var maxChapterOrder uint
	var foundChapter bool
	for _, order := range stage.Chapter_stagedOrder {
		if !foundChapter || order > maxChapterOrder {
			maxChapterOrder = order
			foundChapter = true
		}
	}
	if foundChapter {
		stage.ChapterOrder = maxChapterOrder + 1
	} else {
		stage.ChapterOrder = 0
	}

	var maxCircleOrder uint
	var foundCircle bool
	for _, order := range stage.Circle_stagedOrder {
		if !foundCircle || order > maxCircleOrder {
			maxCircleOrder = order
			foundCircle = true
		}
	}
	if foundCircle {
		stage.CircleOrder = maxCircleOrder + 1
	} else {
		stage.CircleOrder = 0
	}

	var maxCircleGridOrder uint
	var foundCircleGrid bool
	for _, order := range stage.CircleGrid_stagedOrder {
		if !foundCircleGrid || order > maxCircleGridOrder {
			maxCircleGridOrder = order
			foundCircleGrid = true
		}
	}
	if foundCircleGrid {
		stage.CircleGridOrder = maxCircleGridOrder + 1
	} else {
		stage.CircleGridOrder = 0
	}

	var maxContentOrder uint
	var foundContent bool
	for _, order := range stage.Content_stagedOrder {
		if !foundContent || order > maxContentOrder {
			maxContentOrder = order
			foundContent = true
		}
	}
	if foundContent {
		stage.ContentOrder = maxContentOrder + 1
	} else {
		stage.ContentOrder = 0
	}

	var maxExportToMusicxmlOrder uint
	var foundExportToMusicxml bool
	for _, order := range stage.ExportToMusicxml_stagedOrder {
		if !foundExportToMusicxml || order > maxExportToMusicxmlOrder {
			maxExportToMusicxmlOrder = order
			foundExportToMusicxml = true
		}
	}
	if foundExportToMusicxml {
		stage.ExportToMusicxmlOrder = maxExportToMusicxmlOrder + 1
	} else {
		stage.ExportToMusicxmlOrder = 0
	}

	var maxFrontCurveOrder uint
	var foundFrontCurve bool
	for _, order := range stage.FrontCurve_stagedOrder {
		if !foundFrontCurve || order > maxFrontCurveOrder {
			maxFrontCurveOrder = order
			foundFrontCurve = true
		}
	}
	if foundFrontCurve {
		stage.FrontCurveOrder = maxFrontCurveOrder + 1
	} else {
		stage.FrontCurveOrder = 0
	}

	var maxFrontCurveStackOrder uint
	var foundFrontCurveStack bool
	for _, order := range stage.FrontCurveStack_stagedOrder {
		if !foundFrontCurveStack || order > maxFrontCurveStackOrder {
			maxFrontCurveStackOrder = order
			foundFrontCurveStack = true
		}
	}
	if foundFrontCurveStack {
		stage.FrontCurveStackOrder = maxFrontCurveStackOrder + 1
	} else {
		stage.FrontCurveStackOrder = 0
	}

	var maxHorizontalAxisOrder uint
	var foundHorizontalAxis bool
	for _, order := range stage.HorizontalAxis_stagedOrder {
		if !foundHorizontalAxis || order > maxHorizontalAxisOrder {
			maxHorizontalAxisOrder = order
			foundHorizontalAxis = true
		}
	}
	if foundHorizontalAxis {
		stage.HorizontalAxisOrder = maxHorizontalAxisOrder + 1
	} else {
		stage.HorizontalAxisOrder = 0
	}

	var maxKeyOrder uint
	var foundKey bool
	for _, order := range stage.Key_stagedOrder {
		if !foundKey || order > maxKeyOrder {
			maxKeyOrder = order
			foundKey = true
		}
	}
	if foundKey {
		stage.KeyOrder = maxKeyOrder + 1
	} else {
		stage.KeyOrder = 0
	}

	var maxParameterOrder uint
	var foundParameter bool
	for _, order := range stage.Parameter_stagedOrder {
		if !foundParameter || order > maxParameterOrder {
			maxParameterOrder = order
			foundParameter = true
		}
	}
	if foundParameter {
		stage.ParameterOrder = maxParameterOrder + 1
	} else {
		stage.ParameterOrder = 0
	}

	var maxRhombusOrder uint
	var foundRhombus bool
	for _, order := range stage.Rhombus_stagedOrder {
		if !foundRhombus || order > maxRhombusOrder {
			maxRhombusOrder = order
			foundRhombus = true
		}
	}
	if foundRhombus {
		stage.RhombusOrder = maxRhombusOrder + 1
	} else {
		stage.RhombusOrder = 0
	}

	var maxRhombusGridOrder uint
	var foundRhombusGrid bool
	for _, order := range stage.RhombusGrid_stagedOrder {
		if !foundRhombusGrid || order > maxRhombusGridOrder {
			maxRhombusGridOrder = order
			foundRhombusGrid = true
		}
	}
	if foundRhombusGrid {
		stage.RhombusGridOrder = maxRhombusGridOrder + 1
	} else {
		stage.RhombusGridOrder = 0
	}

	var maxShapeCategoryOrder uint
	var foundShapeCategory bool
	for _, order := range stage.ShapeCategory_stagedOrder {
		if !foundShapeCategory || order > maxShapeCategoryOrder {
			maxShapeCategoryOrder = order
			foundShapeCategory = true
		}
	}
	if foundShapeCategory {
		stage.ShapeCategoryOrder = maxShapeCategoryOrder + 1
	} else {
		stage.ShapeCategoryOrder = 0
	}

	var maxSpiralBezierOrder uint
	var foundSpiralBezier bool
	for _, order := range stage.SpiralBezier_stagedOrder {
		if !foundSpiralBezier || order > maxSpiralBezierOrder {
			maxSpiralBezierOrder = order
			foundSpiralBezier = true
		}
	}
	if foundSpiralBezier {
		stage.SpiralBezierOrder = maxSpiralBezierOrder + 1
	} else {
		stage.SpiralBezierOrder = 0
	}

	var maxSpiralBezierGridOrder uint
	var foundSpiralBezierGrid bool
	for _, order := range stage.SpiralBezierGrid_stagedOrder {
		if !foundSpiralBezierGrid || order > maxSpiralBezierGridOrder {
			maxSpiralBezierGridOrder = order
			foundSpiralBezierGrid = true
		}
	}
	if foundSpiralBezierGrid {
		stage.SpiralBezierGridOrder = maxSpiralBezierGridOrder + 1
	} else {
		stage.SpiralBezierGridOrder = 0
	}

	var maxSpiralCircleOrder uint
	var foundSpiralCircle bool
	for _, order := range stage.SpiralCircle_stagedOrder {
		if !foundSpiralCircle || order > maxSpiralCircleOrder {
			maxSpiralCircleOrder = order
			foundSpiralCircle = true
		}
	}
	if foundSpiralCircle {
		stage.SpiralCircleOrder = maxSpiralCircleOrder + 1
	} else {
		stage.SpiralCircleOrder = 0
	}

	var maxSpiralCircleGridOrder uint
	var foundSpiralCircleGrid bool
	for _, order := range stage.SpiralCircleGrid_stagedOrder {
		if !foundSpiralCircleGrid || order > maxSpiralCircleGridOrder {
			maxSpiralCircleGridOrder = order
			foundSpiralCircleGrid = true
		}
	}
	if foundSpiralCircleGrid {
		stage.SpiralCircleGridOrder = maxSpiralCircleGridOrder + 1
	} else {
		stage.SpiralCircleGridOrder = 0
	}

	var maxSpiralLineOrder uint
	var foundSpiralLine bool
	for _, order := range stage.SpiralLine_stagedOrder {
		if !foundSpiralLine || order > maxSpiralLineOrder {
			maxSpiralLineOrder = order
			foundSpiralLine = true
		}
	}
	if foundSpiralLine {
		stage.SpiralLineOrder = maxSpiralLineOrder + 1
	} else {
		stage.SpiralLineOrder = 0
	}

	var maxSpiralLineGridOrder uint
	var foundSpiralLineGrid bool
	for _, order := range stage.SpiralLineGrid_stagedOrder {
		if !foundSpiralLineGrid || order > maxSpiralLineGridOrder {
			maxSpiralLineGridOrder = order
			foundSpiralLineGrid = true
		}
	}
	if foundSpiralLineGrid {
		stage.SpiralLineGridOrder = maxSpiralLineGridOrder + 1
	} else {
		stage.SpiralLineGridOrder = 0
	}

	var maxSpiralOriginOrder uint
	var foundSpiralOrigin bool
	for _, order := range stage.SpiralOrigin_stagedOrder {
		if !foundSpiralOrigin || order > maxSpiralOriginOrder {
			maxSpiralOriginOrder = order
			foundSpiralOrigin = true
		}
	}
	if foundSpiralOrigin {
		stage.SpiralOriginOrder = maxSpiralOriginOrder + 1
	} else {
		stage.SpiralOriginOrder = 0
	}

	var maxSpiralRhombusOrder uint
	var foundSpiralRhombus bool
	for _, order := range stage.SpiralRhombus_stagedOrder {
		if !foundSpiralRhombus || order > maxSpiralRhombusOrder {
			maxSpiralRhombusOrder = order
			foundSpiralRhombus = true
		}
	}
	if foundSpiralRhombus {
		stage.SpiralRhombusOrder = maxSpiralRhombusOrder + 1
	} else {
		stage.SpiralRhombusOrder = 0
	}

	var maxSpiralRhombusGridOrder uint
	var foundSpiralRhombusGrid bool
	for _, order := range stage.SpiralRhombusGrid_stagedOrder {
		if !foundSpiralRhombusGrid || order > maxSpiralRhombusGridOrder {
			maxSpiralRhombusGridOrder = order
			foundSpiralRhombusGrid = true
		}
	}
	if foundSpiralRhombusGrid {
		stage.SpiralRhombusGridOrder = maxSpiralRhombusGridOrder + 1
	} else {
		stage.SpiralRhombusGridOrder = 0
	}

	var maxVerticalAxisOrder uint
	var foundVerticalAxis bool
	for _, order := range stage.VerticalAxis_stagedOrder {
		if !foundVerticalAxis || order > maxVerticalAxisOrder {
			maxVerticalAxisOrder = order
			foundVerticalAxis = true
		}
	}
	if foundVerticalAxis {
		stage.VerticalAxisOrder = maxVerticalAxisOrder + 1
	} else {
		stage.VerticalAxisOrder = 0
	}

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {
	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []string) {
	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

// GetStructInstancesByOrderAuto returns a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Axis:
		tmp := GetStructInstancesByOrder(stage.Axiss, stage.Axis_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Axis implements.
			res = append(res, any(v).(T))
		}
		return res
	case *AxisGrid:
		tmp := GetStructInstancesByOrder(stage.AxisGrids, stage.AxisGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AxisGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Bezier:
		tmp := GetStructInstancesByOrder(stage.Beziers, stage.Bezier_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Bezier implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BezierGrid:
		tmp := GetStructInstancesByOrder(stage.BezierGrids, stage.BezierGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BezierGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *BezierGridStack:
		tmp := GetStructInstancesByOrder(stage.BezierGridStacks, stage.BezierGridStack_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BezierGridStack implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Chapter:
		tmp := GetStructInstancesByOrder(stage.Chapters, stage.Chapter_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Chapter implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Circle:
		tmp := GetStructInstancesByOrder(stage.Circles, stage.Circle_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Circle implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CircleGrid:
		tmp := GetStructInstancesByOrder(stage.CircleGrids, stage.CircleGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CircleGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Content:
		tmp := GetStructInstancesByOrder(stage.Contents, stage.Content_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Content implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ExportToMusicxml:
		tmp := GetStructInstancesByOrder(stage.ExportToMusicxmls, stage.ExportToMusicxml_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ExportToMusicxml implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FrontCurve:
		tmp := GetStructInstancesByOrder(stage.FrontCurves, stage.FrontCurve_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FrontCurve implements.
			res = append(res, any(v).(T))
		}
		return res
	case *FrontCurveStack:
		tmp := GetStructInstancesByOrder(stage.FrontCurveStacks, stage.FrontCurveStack_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *FrontCurveStack implements.
			res = append(res, any(v).(T))
		}
		return res
	case *HorizontalAxis:
		tmp := GetStructInstancesByOrder(stage.HorizontalAxiss, stage.HorizontalAxis_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *HorizontalAxis implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Key:
		tmp := GetStructInstancesByOrder(stage.Keys, stage.Key_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Key implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Parameter:
		tmp := GetStructInstancesByOrder(stage.Parameters, stage.Parameter_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Parameter implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Rhombus:
		tmp := GetStructInstancesByOrder(stage.Rhombuss, stage.Rhombus_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Rhombus implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RhombusGrid:
		tmp := GetStructInstancesByOrder(stage.RhombusGrids, stage.RhombusGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RhombusGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ShapeCategory:
		tmp := GetStructInstancesByOrder(stage.ShapeCategorys, stage.ShapeCategory_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ShapeCategory implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralBezier:
		tmp := GetStructInstancesByOrder(stage.SpiralBeziers, stage.SpiralBezier_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralBezier implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralBezierGrid:
		tmp := GetStructInstancesByOrder(stage.SpiralBezierGrids, stage.SpiralBezierGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralBezierGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralCircle:
		tmp := GetStructInstancesByOrder(stage.SpiralCircles, stage.SpiralCircle_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralCircle implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralCircleGrid:
		tmp := GetStructInstancesByOrder(stage.SpiralCircleGrids, stage.SpiralCircleGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralCircleGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralLine:
		tmp := GetStructInstancesByOrder(stage.SpiralLines, stage.SpiralLine_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralLine implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralLineGrid:
		tmp := GetStructInstancesByOrder(stage.SpiralLineGrids, stage.SpiralLineGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralLineGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralOrigin:
		tmp := GetStructInstancesByOrder(stage.SpiralOrigins, stage.SpiralOrigin_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralOrigin implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralRhombus:
		tmp := GetStructInstancesByOrder(stage.SpiralRhombuss, stage.SpiralRhombus_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralRhombus implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SpiralRhombusGrid:
		tmp := GetStructInstancesByOrder(stage.SpiralRhombusGrids, stage.SpiralRhombusGrid_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SpiralRhombusGrid implements.
			res = append(res, any(v).(T))
		}
		return res
	case *VerticalAxis:
		tmp := GetStructInstancesByOrder(stage.VerticalAxiss, stage.VerticalAxis_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *VerticalAxis implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []T) {
	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {
	switch namedStructName {
	// insertion point for case
	case "Axis":
		res = GetNamedStructInstances(stage.Axiss, stage.Axis_stagedOrder)
	case "AxisGrid":
		res = GetNamedStructInstances(stage.AxisGrids, stage.AxisGrid_stagedOrder)
	case "Bezier":
		res = GetNamedStructInstances(stage.Beziers, stage.Bezier_stagedOrder)
	case "BezierGrid":
		res = GetNamedStructInstances(stage.BezierGrids, stage.BezierGrid_stagedOrder)
	case "BezierGridStack":
		res = GetNamedStructInstances(stage.BezierGridStacks, stage.BezierGridStack_stagedOrder)
	case "Chapter":
		res = GetNamedStructInstances(stage.Chapters, stage.Chapter_stagedOrder)
	case "Circle":
		res = GetNamedStructInstances(stage.Circles, stage.Circle_stagedOrder)
	case "CircleGrid":
		res = GetNamedStructInstances(stage.CircleGrids, stage.CircleGrid_stagedOrder)
	case "Content":
		res = GetNamedStructInstances(stage.Contents, stage.Content_stagedOrder)
	case "ExportToMusicxml":
		res = GetNamedStructInstances(stage.ExportToMusicxmls, stage.ExportToMusicxml_stagedOrder)
	case "FrontCurve":
		res = GetNamedStructInstances(stage.FrontCurves, stage.FrontCurve_stagedOrder)
	case "FrontCurveStack":
		res = GetNamedStructInstances(stage.FrontCurveStacks, stage.FrontCurveStack_stagedOrder)
	case "HorizontalAxis":
		res = GetNamedStructInstances(stage.HorizontalAxiss, stage.HorizontalAxis_stagedOrder)
	case "Key":
		res = GetNamedStructInstances(stage.Keys, stage.Key_stagedOrder)
	case "Parameter":
		res = GetNamedStructInstances(stage.Parameters, stage.Parameter_stagedOrder)
	case "Rhombus":
		res = GetNamedStructInstances(stage.Rhombuss, stage.Rhombus_stagedOrder)
	case "RhombusGrid":
		res = GetNamedStructInstances(stage.RhombusGrids, stage.RhombusGrid_stagedOrder)
	case "ShapeCategory":
		res = GetNamedStructInstances(stage.ShapeCategorys, stage.ShapeCategory_stagedOrder)
	case "SpiralBezier":
		res = GetNamedStructInstances(stage.SpiralBeziers, stage.SpiralBezier_stagedOrder)
	case "SpiralBezierGrid":
		res = GetNamedStructInstances(stage.SpiralBezierGrids, stage.SpiralBezierGrid_stagedOrder)
	case "SpiralCircle":
		res = GetNamedStructInstances(stage.SpiralCircles, stage.SpiralCircle_stagedOrder)
	case "SpiralCircleGrid":
		res = GetNamedStructInstances(stage.SpiralCircleGrids, stage.SpiralCircleGrid_stagedOrder)
	case "SpiralLine":
		res = GetNamedStructInstances(stage.SpiralLines, stage.SpiralLine_stagedOrder)
	case "SpiralLineGrid":
		res = GetNamedStructInstances(stage.SpiralLineGrids, stage.SpiralLineGrid_stagedOrder)
	case "SpiralOrigin":
		res = GetNamedStructInstances(stage.SpiralOrigins, stage.SpiralOrigin_stagedOrder)
	case "SpiralRhombus":
		res = GetNamedStructInstances(stage.SpiralRhombuss, stage.SpiralRhombus_stagedOrder)
	case "SpiralRhombusGrid":
		res = GetNamedStructInstances(stage.SpiralRhombusGrids, stage.SpiralRhombusGrid_stagedOrder)
	case "VerticalAxis":
		res = GetNamedStructInstances(stage.VerticalAxiss, stage.VerticalAxis_stagedOrder)
	}

	return
}

type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
}

func (stage *Stage) GetType() string {
	return "github.com/thomaspeugeot/phyllotaxymusic/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return phyllotaxymusic_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return phyllotaxymusic_go.GoDiagramsDir
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
	CommitChapter(chapter *Chapter)
	CheckoutChapter(chapter *Chapter)
	CommitCircle(circle *Circle)
	CheckoutCircle(circle *Circle)
	CommitCircleGrid(circlegrid *CircleGrid)
	CheckoutCircleGrid(circlegrid *CircleGrid)
	CommitContent(content *Content)
	CheckoutContent(content *Content)
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
		Axiss:           make(map[*Axis]struct{}),
		Axiss_mapString: make(map[string]*Axis),

		AxisGrids:           make(map[*AxisGrid]struct{}),
		AxisGrids_mapString: make(map[string]*AxisGrid),

		Beziers:           make(map[*Bezier]struct{}),
		Beziers_mapString: make(map[string]*Bezier),

		BezierGrids:           make(map[*BezierGrid]struct{}),
		BezierGrids_mapString: make(map[string]*BezierGrid),

		BezierGridStacks:           make(map[*BezierGridStack]struct{}),
		BezierGridStacks_mapString: make(map[string]*BezierGridStack),

		Chapters:           make(map[*Chapter]struct{}),
		Chapters_mapString: make(map[string]*Chapter),

		Circles:           make(map[*Circle]struct{}),
		Circles_mapString: make(map[string]*Circle),

		CircleGrids:           make(map[*CircleGrid]struct{}),
		CircleGrids_mapString: make(map[string]*CircleGrid),

		Contents:           make(map[*Content]struct{}),
		Contents_mapString: make(map[string]*Content),

		ExportToMusicxmls:           make(map[*ExportToMusicxml]struct{}),
		ExportToMusicxmls_mapString: make(map[string]*ExportToMusicxml),

		FrontCurves:           make(map[*FrontCurve]struct{}),
		FrontCurves_mapString: make(map[string]*FrontCurve),

		FrontCurveStacks:           make(map[*FrontCurveStack]struct{}),
		FrontCurveStacks_mapString: make(map[string]*FrontCurveStack),

		HorizontalAxiss:           make(map[*HorizontalAxis]struct{}),
		HorizontalAxiss_mapString: make(map[string]*HorizontalAxis),

		Keys:           make(map[*Key]struct{}),
		Keys_mapString: make(map[string]*Key),

		Parameters:           make(map[*Parameter]struct{}),
		Parameters_mapString: make(map[string]*Parameter),

		Rhombuss:           make(map[*Rhombus]struct{}),
		Rhombuss_mapString: make(map[string]*Rhombus),

		RhombusGrids:           make(map[*RhombusGrid]struct{}),
		RhombusGrids_mapString: make(map[string]*RhombusGrid),

		ShapeCategorys:           make(map[*ShapeCategory]struct{}),
		ShapeCategorys_mapString: make(map[string]*ShapeCategory),

		SpiralBeziers:           make(map[*SpiralBezier]struct{}),
		SpiralBeziers_mapString: make(map[string]*SpiralBezier),

		SpiralBezierGrids:           make(map[*SpiralBezierGrid]struct{}),
		SpiralBezierGrids_mapString: make(map[string]*SpiralBezierGrid),

		SpiralCircles:           make(map[*SpiralCircle]struct{}),
		SpiralCircles_mapString: make(map[string]*SpiralCircle),

		SpiralCircleGrids:           make(map[*SpiralCircleGrid]struct{}),
		SpiralCircleGrids_mapString: make(map[string]*SpiralCircleGrid),

		SpiralLines:           make(map[*SpiralLine]struct{}),
		SpiralLines_mapString: make(map[string]*SpiralLine),

		SpiralLineGrids:           make(map[*SpiralLineGrid]struct{}),
		SpiralLineGrids_mapString: make(map[string]*SpiralLineGrid),

		SpiralOrigins:           make(map[*SpiralOrigin]struct{}),
		SpiralOrigins_mapString: make(map[string]*SpiralOrigin),

		SpiralRhombuss:           make(map[*SpiralRhombus]struct{}),
		SpiralRhombuss_mapString: make(map[string]*SpiralRhombus),

		SpiralRhombusGrids:           make(map[*SpiralRhombusGrid]struct{}),
		SpiralRhombusGrids_mapString: make(map[string]*SpiralRhombusGrid),

		VerticalAxiss:           make(map[*VerticalAxis]struct{}),
		VerticalAxiss_mapString: make(map[string]*VerticalAxis),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Axis_stagedOrder: make(map[*Axis]uint),

		AxisGrid_stagedOrder: make(map[*AxisGrid]uint),

		Bezier_stagedOrder: make(map[*Bezier]uint),

		BezierGrid_stagedOrder: make(map[*BezierGrid]uint),

		BezierGridStack_stagedOrder: make(map[*BezierGridStack]uint),

		Chapter_stagedOrder: make(map[*Chapter]uint),

		Circle_stagedOrder: make(map[*Circle]uint),

		CircleGrid_stagedOrder: make(map[*CircleGrid]uint),

		Content_stagedOrder: make(map[*Content]uint),

		ExportToMusicxml_stagedOrder: make(map[*ExportToMusicxml]uint),

		FrontCurve_stagedOrder: make(map[*FrontCurve]uint),

		FrontCurveStack_stagedOrder: make(map[*FrontCurveStack]uint),

		HorizontalAxis_stagedOrder: make(map[*HorizontalAxis]uint),

		Key_stagedOrder: make(map[*Key]uint),

		Parameter_stagedOrder: make(map[*Parameter]uint),

		Rhombus_stagedOrder: make(map[*Rhombus]uint),

		RhombusGrid_stagedOrder: make(map[*RhombusGrid]uint),

		ShapeCategory_stagedOrder: make(map[*ShapeCategory]uint),

		SpiralBezier_stagedOrder: make(map[*SpiralBezier]uint),

		SpiralBezierGrid_stagedOrder: make(map[*SpiralBezierGrid]uint),

		SpiralCircle_stagedOrder: make(map[*SpiralCircle]uint),

		SpiralCircleGrid_stagedOrder: make(map[*SpiralCircleGrid]uint),

		SpiralLine_stagedOrder: make(map[*SpiralLine]uint),

		SpiralLineGrid_stagedOrder: make(map[*SpiralLineGrid]uint),

		SpiralOrigin_stagedOrder: make(map[*SpiralOrigin]uint),

		SpiralRhombus_stagedOrder: make(map[*SpiralRhombus]uint),

		SpiralRhombusGrid_stagedOrder: make(map[*SpiralRhombusGrid]uint),

		VerticalAxis_stagedOrder: make(map[*VerticalAxis]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Axis": &AxisUnmarshaller{},

			"AxisGrid": &AxisGridUnmarshaller{},

			"Bezier": &BezierUnmarshaller{},

			"BezierGrid": &BezierGridUnmarshaller{},

			"BezierGridStack": &BezierGridStackUnmarshaller{},

			"Chapter": &ChapterUnmarshaller{},

			"Circle": &CircleUnmarshaller{},

			"CircleGrid": &CircleGridUnmarshaller{},

			"Content": &ContentUnmarshaller{},

			"ExportToMusicxml": &ExportToMusicxmlUnmarshaller{},

			"FrontCurve": &FrontCurveUnmarshaller{},

			"FrontCurveStack": &FrontCurveStackUnmarshaller{},

			"HorizontalAxis": &HorizontalAxisUnmarshaller{},

			"Key": &KeyUnmarshaller{},

			"Parameter": &ParameterUnmarshaller{},

			"Rhombus": &RhombusUnmarshaller{},

			"RhombusGrid": &RhombusGridUnmarshaller{},

			"ShapeCategory": &ShapeCategoryUnmarshaller{},

			"SpiralBezier": &SpiralBezierUnmarshaller{},

			"SpiralBezierGrid": &SpiralBezierGridUnmarshaller{},

			"SpiralCircle": &SpiralCircleUnmarshaller{},

			"SpiralCircleGrid": &SpiralCircleGridUnmarshaller{},

			"SpiralLine": &SpiralLineUnmarshaller{},

			"SpiralLineGrid": &SpiralLineGridUnmarshaller{},

			"SpiralOrigin": &SpiralOriginUnmarshaller{},

			"SpiralRhombus": &SpiralRhombusUnmarshaller{},

			"SpiralRhombusGrid": &SpiralRhombusGridUnmarshaller{},

			"VerticalAxis": &VerticalAxisUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Axis"},
			{name: "AxisGrid"},
			{name: "Bezier"},
			{name: "BezierGrid"},
			{name: "BezierGridStack"},
			{name: "Chapter"},
			{name: "Circle"},
			{name: "CircleGrid"},
			{name: "Content"},
			{name: "ExportToMusicxml"},
			{name: "FrontCurve"},
			{name: "FrontCurveStack"},
			{name: "HorizontalAxis"},
			{name: "Key"},
			{name: "Parameter"},
			{name: "Rhombus"},
			{name: "RhombusGrid"},
			{name: "ShapeCategory"},
			{name: "SpiralBezier"},
			{name: "SpiralBezierGrid"},
			{name: "SpiralCircle"},
			{name: "SpiralCircleGrid"},
			{name: "SpiralLine"},
			{name: "SpiralLineGrid"},
			{name: "SpiralOrigin"},
			{name: "SpiralRhombus"},
			{name: "SpiralRhombusGrid"},
			{name: "VerticalAxis"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Axis:
		return stage.Axis_stagedOrder[instance]
	case *AxisGrid:
		return stage.AxisGrid_stagedOrder[instance]
	case *Bezier:
		return stage.Bezier_stagedOrder[instance]
	case *BezierGrid:
		return stage.BezierGrid_stagedOrder[instance]
	case *BezierGridStack:
		return stage.BezierGridStack_stagedOrder[instance]
	case *Chapter:
		return stage.Chapter_stagedOrder[instance]
	case *Circle:
		return stage.Circle_stagedOrder[instance]
	case *CircleGrid:
		return stage.CircleGrid_stagedOrder[instance]
	case *Content:
		return stage.Content_stagedOrder[instance]
	case *ExportToMusicxml:
		return stage.ExportToMusicxml_stagedOrder[instance]
	case *FrontCurve:
		return stage.FrontCurve_stagedOrder[instance]
	case *FrontCurveStack:
		return stage.FrontCurveStack_stagedOrder[instance]
	case *HorizontalAxis:
		return stage.HorizontalAxis_stagedOrder[instance]
	case *Key:
		return stage.Key_stagedOrder[instance]
	case *Parameter:
		return stage.Parameter_stagedOrder[instance]
	case *Rhombus:
		return stage.Rhombus_stagedOrder[instance]
	case *RhombusGrid:
		return stage.RhombusGrid_stagedOrder[instance]
	case *ShapeCategory:
		return stage.ShapeCategory_stagedOrder[instance]
	case *SpiralBezier:
		return stage.SpiralBezier_stagedOrder[instance]
	case *SpiralBezierGrid:
		return stage.SpiralBezierGrid_stagedOrder[instance]
	case *SpiralCircle:
		return stage.SpiralCircle_stagedOrder[instance]
	case *SpiralCircleGrid:
		return stage.SpiralCircleGrid_stagedOrder[instance]
	case *SpiralLine:
		return stage.SpiralLine_stagedOrder[instance]
	case *SpiralLineGrid:
		return stage.SpiralLineGrid_stagedOrder[instance]
	case *SpiralOrigin:
		return stage.SpiralOrigin_stagedOrder[instance]
	case *SpiralRhombus:
		return stage.SpiralRhombus_stagedOrder[instance]
	case *SpiralRhombusGrid:
		return stage.SpiralRhombusGrid_stagedOrder[instance]
	case *VerticalAxis:
		return stage.VerticalAxis_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Axis:
		return stage.Axis_stagedOrder[instance]
	case *AxisGrid:
		return stage.AxisGrid_stagedOrder[instance]
	case *Bezier:
		return stage.Bezier_stagedOrder[instance]
	case *BezierGrid:
		return stage.BezierGrid_stagedOrder[instance]
	case *BezierGridStack:
		return stage.BezierGridStack_stagedOrder[instance]
	case *Chapter:
		return stage.Chapter_stagedOrder[instance]
	case *Circle:
		return stage.Circle_stagedOrder[instance]
	case *CircleGrid:
		return stage.CircleGrid_stagedOrder[instance]
	case *Content:
		return stage.Content_stagedOrder[instance]
	case *ExportToMusicxml:
		return stage.ExportToMusicxml_stagedOrder[instance]
	case *FrontCurve:
		return stage.FrontCurve_stagedOrder[instance]
	case *FrontCurveStack:
		return stage.FrontCurveStack_stagedOrder[instance]
	case *HorizontalAxis:
		return stage.HorizontalAxis_stagedOrder[instance]
	case *Key:
		return stage.Key_stagedOrder[instance]
	case *Parameter:
		return stage.Parameter_stagedOrder[instance]
	case *Rhombus:
		return stage.Rhombus_stagedOrder[instance]
	case *RhombusGrid:
		return stage.RhombusGrid_stagedOrder[instance]
	case *ShapeCategory:
		return stage.ShapeCategory_stagedOrder[instance]
	case *SpiralBezier:
		return stage.SpiralBezier_stagedOrder[instance]
	case *SpiralBezierGrid:
		return stage.SpiralBezierGrid_stagedOrder[instance]
	case *SpiralCircle:
		return stage.SpiralCircle_stagedOrder[instance]
	case *SpiralCircleGrid:
		return stage.SpiralCircleGrid_stagedOrder[instance]
	case *SpiralLine:
		return stage.SpiralLine_stagedOrder[instance]
	case *SpiralLineGrid:
		return stage.SpiralLineGrid_stagedOrder[instance]
	case *SpiralOrigin:
		return stage.SpiralOrigin_stagedOrder[instance]
	case *SpiralRhombus:
		return stage.SpiralRhombus_stagedOrder[instance]
	case *SpiralRhombusGrid:
		return stage.SpiralRhombusGrid_stagedOrder[instance]
	case *VerticalAxis:
		return stage.VerticalAxis_stagedOrder[instance]
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
	tmp2 := stage.beforeCommitHooks
	stage.beforeCommitHooks = nil
	tmp3 := stage.afterCommitHooks
	stage.afterCommitHooks = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.beforeCommitHooks = tmp2
	stage.afterCommitHooks = tmp3
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.isInDeltaMode && stage.navigationMode == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
		stage.ResetHard()
	}

	if stage.IsInDeltaMode() {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
		if stage.probeIF != nil {
			stage.probeIF.RefreshNavigationTree()
		}
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Axis"] = len(stage.Axiss)
	stage.Map_GongStructName_InstancesNb["AxisGrid"] = len(stage.AxisGrids)
	stage.Map_GongStructName_InstancesNb["Bezier"] = len(stage.Beziers)
	stage.Map_GongStructName_InstancesNb["BezierGrid"] = len(stage.BezierGrids)
	stage.Map_GongStructName_InstancesNb["BezierGridStack"] = len(stage.BezierGridStacks)
	stage.Map_GongStructName_InstancesNb["Chapter"] = len(stage.Chapters)
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["CircleGrid"] = len(stage.CircleGrids)
	stage.Map_GongStructName_InstancesNb["Content"] = len(stage.Contents)
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
	stage.ComputeInstancesNb()
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
		stage.Axiss[axis] = struct{}{}
		stage.Axis_stagedOrder[axis] = stage.AxisOrder
		stage.AxisOrder++
	}
	stage.Axiss_mapString[axis.Name] = axis

	return axis
}

// StagePreserveOrder puts axis to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AxisOrder
// - update stage.AxisOrder accordingly
func (axis *Axis) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Axiss[axis]; !ok {
		stage.Axiss[axis] = struct{}{}

		if order > stage.AxisOrder {
			stage.AxisOrder = order
		}
		stage.Axis_stagedOrder[axis] = order
		stage.AxisOrder++
	}
	stage.Axiss_mapString[axis.Name] = axis
}

// Unstage removes axis off the model stage
func (axis *Axis) Unstage(stage *Stage) *Axis {
	delete(stage.Axiss, axis)
	// issue1150
	// delete(stage.Axis_stagedOrder, axis)
	delete(stage.Axiss_mapString, axis.Name)

	return axis
}

// UnstageVoid removes axis off the model stage
func (axis *Axis) UnstageVoid(stage *Stage) {
	delete(stage.Axiss, axis)
	// issue1150
	// delete(stage.Axis_stagedOrder, axis)
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

func (axis *Axis) StageVoid(stage *Stage) {
	axis.Stage(stage)
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

// for satisfaction of GongStruct interface
func (axis *Axis) SetName(name string) {
	axis.Name = name
}

// Stage puts axisgrid to the model stage
func (axisgrid *AxisGrid) Stage(stage *Stage) *AxisGrid {
	if _, ok := stage.AxisGrids[axisgrid]; !ok {
		stage.AxisGrids[axisgrid] = struct{}{}
		stage.AxisGrid_stagedOrder[axisgrid] = stage.AxisGridOrder
		stage.AxisGridOrder++
	}
	stage.AxisGrids_mapString[axisgrid.Name] = axisgrid

	return axisgrid
}

// StagePreserveOrder puts axisgrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AxisGridOrder
// - update stage.AxisGridOrder accordingly
func (axisgrid *AxisGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AxisGrids[axisgrid]; !ok {
		stage.AxisGrids[axisgrid] = struct{}{}

		if order > stage.AxisGridOrder {
			stage.AxisGridOrder = order
		}
		stage.AxisGrid_stagedOrder[axisgrid] = order
		stage.AxisGridOrder++
	}
	stage.AxisGrids_mapString[axisgrid.Name] = axisgrid
}

// Unstage removes axisgrid off the model stage
func (axisgrid *AxisGrid) Unstage(stage *Stage) *AxisGrid {
	delete(stage.AxisGrids, axisgrid)
	// issue1150
	// delete(stage.AxisGrid_stagedOrder, axisgrid)
	delete(stage.AxisGrids_mapString, axisgrid.Name)

	return axisgrid
}

// UnstageVoid removes axisgrid off the model stage
func (axisgrid *AxisGrid) UnstageVoid(stage *Stage) {
	delete(stage.AxisGrids, axisgrid)
	// issue1150
	// delete(stage.AxisGrid_stagedOrder, axisgrid)
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

func (axisgrid *AxisGrid) StageVoid(stage *Stage) {
	axisgrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (axisgrid *AxisGrid) SetName(name string) {
	axisgrid.Name = name
}

// Stage puts bezier to the model stage
func (bezier *Bezier) Stage(stage *Stage) *Bezier {
	if _, ok := stage.Beziers[bezier]; !ok {
		stage.Beziers[bezier] = struct{}{}
		stage.Bezier_stagedOrder[bezier] = stage.BezierOrder
		stage.BezierOrder++
	}
	stage.Beziers_mapString[bezier.Name] = bezier

	return bezier
}

// StagePreserveOrder puts bezier to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BezierOrder
// - update stage.BezierOrder accordingly
func (bezier *Bezier) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Beziers[bezier]; !ok {
		stage.Beziers[bezier] = struct{}{}

		if order > stage.BezierOrder {
			stage.BezierOrder = order
		}
		stage.Bezier_stagedOrder[bezier] = order
		stage.BezierOrder++
	}
	stage.Beziers_mapString[bezier.Name] = bezier
}

// Unstage removes bezier off the model stage
func (bezier *Bezier) Unstage(stage *Stage) *Bezier {
	delete(stage.Beziers, bezier)
	// issue1150
	// delete(stage.Bezier_stagedOrder, bezier)
	delete(stage.Beziers_mapString, bezier.Name)

	return bezier
}

// UnstageVoid removes bezier off the model stage
func (bezier *Bezier) UnstageVoid(stage *Stage) {
	delete(stage.Beziers, bezier)
	// issue1150
	// delete(stage.Bezier_stagedOrder, bezier)
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

func (bezier *Bezier) StageVoid(stage *Stage) {
	bezier.Stage(stage)
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

// for satisfaction of GongStruct interface
func (bezier *Bezier) SetName(name string) {
	bezier.Name = name
}

// Stage puts beziergrid to the model stage
func (beziergrid *BezierGrid) Stage(stage *Stage) *BezierGrid {
	if _, ok := stage.BezierGrids[beziergrid]; !ok {
		stage.BezierGrids[beziergrid] = struct{}{}
		stage.BezierGrid_stagedOrder[beziergrid] = stage.BezierGridOrder
		stage.BezierGridOrder++
	}
	stage.BezierGrids_mapString[beziergrid.Name] = beziergrid

	return beziergrid
}

// StagePreserveOrder puts beziergrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BezierGridOrder
// - update stage.BezierGridOrder accordingly
func (beziergrid *BezierGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BezierGrids[beziergrid]; !ok {
		stage.BezierGrids[beziergrid] = struct{}{}

		if order > stage.BezierGridOrder {
			stage.BezierGridOrder = order
		}
		stage.BezierGrid_stagedOrder[beziergrid] = order
		stage.BezierGridOrder++
	}
	stage.BezierGrids_mapString[beziergrid.Name] = beziergrid
}

// Unstage removes beziergrid off the model stage
func (beziergrid *BezierGrid) Unstage(stage *Stage) *BezierGrid {
	delete(stage.BezierGrids, beziergrid)
	// issue1150
	// delete(stage.BezierGrid_stagedOrder, beziergrid)
	delete(stage.BezierGrids_mapString, beziergrid.Name)

	return beziergrid
}

// UnstageVoid removes beziergrid off the model stage
func (beziergrid *BezierGrid) UnstageVoid(stage *Stage) {
	delete(stage.BezierGrids, beziergrid)
	// issue1150
	// delete(stage.BezierGrid_stagedOrder, beziergrid)
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

func (beziergrid *BezierGrid) StageVoid(stage *Stage) {
	beziergrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (beziergrid *BezierGrid) SetName(name string) {
	beziergrid.Name = name
}

// Stage puts beziergridstack to the model stage
func (beziergridstack *BezierGridStack) Stage(stage *Stage) *BezierGridStack {
	if _, ok := stage.BezierGridStacks[beziergridstack]; !ok {
		stage.BezierGridStacks[beziergridstack] = struct{}{}
		stage.BezierGridStack_stagedOrder[beziergridstack] = stage.BezierGridStackOrder
		stage.BezierGridStackOrder++
	}
	stage.BezierGridStacks_mapString[beziergridstack.Name] = beziergridstack

	return beziergridstack
}

// StagePreserveOrder puts beziergridstack to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BezierGridStackOrder
// - update stage.BezierGridStackOrder accordingly
func (beziergridstack *BezierGridStack) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BezierGridStacks[beziergridstack]; !ok {
		stage.BezierGridStacks[beziergridstack] = struct{}{}

		if order > stage.BezierGridStackOrder {
			stage.BezierGridStackOrder = order
		}
		stage.BezierGridStack_stagedOrder[beziergridstack] = order
		stage.BezierGridStackOrder++
	}
	stage.BezierGridStacks_mapString[beziergridstack.Name] = beziergridstack
}

// Unstage removes beziergridstack off the model stage
func (beziergridstack *BezierGridStack) Unstage(stage *Stage) *BezierGridStack {
	delete(stage.BezierGridStacks, beziergridstack)
	// issue1150
	// delete(stage.BezierGridStack_stagedOrder, beziergridstack)
	delete(stage.BezierGridStacks_mapString, beziergridstack.Name)

	return beziergridstack
}

// UnstageVoid removes beziergridstack off the model stage
func (beziergridstack *BezierGridStack) UnstageVoid(stage *Stage) {
	delete(stage.BezierGridStacks, beziergridstack)
	// issue1150
	// delete(stage.BezierGridStack_stagedOrder, beziergridstack)
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

func (beziergridstack *BezierGridStack) StageVoid(stage *Stage) {
	beziergridstack.Stage(stage)
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

// for satisfaction of GongStruct interface
func (beziergridstack *BezierGridStack) SetName(name string) {
	beziergridstack.Name = name
}

// Stage puts chapter to the model stage
func (chapter *Chapter) Stage(stage *Stage) *Chapter {
	if _, ok := stage.Chapters[chapter]; !ok {
		stage.Chapters[chapter] = struct{}{}
		stage.Chapter_stagedOrder[chapter] = stage.ChapterOrder
		stage.ChapterOrder++
	}
	stage.Chapters_mapString[chapter.Name] = chapter

	return chapter
}

// StagePreserveOrder puts chapter to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ChapterOrder
// - update stage.ChapterOrder accordingly
func (chapter *Chapter) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Chapters[chapter]; !ok {
		stage.Chapters[chapter] = struct{}{}

		if order > stage.ChapterOrder {
			stage.ChapterOrder = order
		}
		stage.Chapter_stagedOrder[chapter] = order
		stage.ChapterOrder++
	}
	stage.Chapters_mapString[chapter.Name] = chapter
}

// Unstage removes chapter off the model stage
func (chapter *Chapter) Unstage(stage *Stage) *Chapter {
	delete(stage.Chapters, chapter)
	// issue1150
	// delete(stage.Chapter_stagedOrder, chapter)
	delete(stage.Chapters_mapString, chapter.Name)

	return chapter
}

// UnstageVoid removes chapter off the model stage
func (chapter *Chapter) UnstageVoid(stage *Stage) {
	delete(stage.Chapters, chapter)
	// issue1150
	// delete(stage.Chapter_stagedOrder, chapter)
	delete(stage.Chapters_mapString, chapter.Name)
}

// commit chapter to the back repo (if it is already staged)
func (chapter *Chapter) Commit(stage *Stage) *Chapter {
	if _, ok := stage.Chapters[chapter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitChapter(chapter)
		}
	}
	return chapter
}

func (chapter *Chapter) CommitVoid(stage *Stage) {
	chapter.Commit(stage)
}

func (chapter *Chapter) StageVoid(stage *Stage) {
	chapter.Stage(stage)
}

// Checkout chapter to the back repo (if it is already staged)
func (chapter *Chapter) Checkout(stage *Stage) *Chapter {
	if _, ok := stage.Chapters[chapter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutChapter(chapter)
		}
	}
	return chapter
}

// for satisfaction of GongStruct interface
func (chapter *Chapter) GetName() (res string) {
	return chapter.Name
}

// for satisfaction of GongStruct interface
func (chapter *Chapter) SetName(name string) {
	chapter.Name = name
}

// Stage puts circle to the model stage
func (circle *Circle) Stage(stage *Stage) *Circle {
	if _, ok := stage.Circles[circle]; !ok {
		stage.Circles[circle] = struct{}{}
		stage.Circle_stagedOrder[circle] = stage.CircleOrder
		stage.CircleOrder++
	}
	stage.Circles_mapString[circle.Name] = circle

	return circle
}

// StagePreserveOrder puts circle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CircleOrder
// - update stage.CircleOrder accordingly
func (circle *Circle) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Circles[circle]; !ok {
		stage.Circles[circle] = struct{}{}

		if order > stage.CircleOrder {
			stage.CircleOrder = order
		}
		stage.Circle_stagedOrder[circle] = order
		stage.CircleOrder++
	}
	stage.Circles_mapString[circle.Name] = circle
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage(stage *Stage) *Circle {
	delete(stage.Circles, circle)
	// issue1150
	// delete(stage.Circle_stagedOrder, circle)
	delete(stage.Circles_mapString, circle.Name)

	return circle
}

// UnstageVoid removes circle off the model stage
func (circle *Circle) UnstageVoid(stage *Stage) {
	delete(stage.Circles, circle)
	// issue1150
	// delete(stage.Circle_stagedOrder, circle)
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

func (circle *Circle) StageVoid(stage *Stage) {
	circle.Stage(stage)
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

// for satisfaction of GongStruct interface
func (circle *Circle) SetName(name string) {
	circle.Name = name
}

// Stage puts circlegrid to the model stage
func (circlegrid *CircleGrid) Stage(stage *Stage) *CircleGrid {
	if _, ok := stage.CircleGrids[circlegrid]; !ok {
		stage.CircleGrids[circlegrid] = struct{}{}
		stage.CircleGrid_stagedOrder[circlegrid] = stage.CircleGridOrder
		stage.CircleGridOrder++
	}
	stage.CircleGrids_mapString[circlegrid.Name] = circlegrid

	return circlegrid
}

// StagePreserveOrder puts circlegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CircleGridOrder
// - update stage.CircleGridOrder accordingly
func (circlegrid *CircleGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CircleGrids[circlegrid]; !ok {
		stage.CircleGrids[circlegrid] = struct{}{}

		if order > stage.CircleGridOrder {
			stage.CircleGridOrder = order
		}
		stage.CircleGrid_stagedOrder[circlegrid] = order
		stage.CircleGridOrder++
	}
	stage.CircleGrids_mapString[circlegrid.Name] = circlegrid
}

// Unstage removes circlegrid off the model stage
func (circlegrid *CircleGrid) Unstage(stage *Stage) *CircleGrid {
	delete(stage.CircleGrids, circlegrid)
	// issue1150
	// delete(stage.CircleGrid_stagedOrder, circlegrid)
	delete(stage.CircleGrids_mapString, circlegrid.Name)

	return circlegrid
}

// UnstageVoid removes circlegrid off the model stage
func (circlegrid *CircleGrid) UnstageVoid(stage *Stage) {
	delete(stage.CircleGrids, circlegrid)
	// issue1150
	// delete(stage.CircleGrid_stagedOrder, circlegrid)
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

func (circlegrid *CircleGrid) StageVoid(stage *Stage) {
	circlegrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (circlegrid *CircleGrid) SetName(name string) {
	circlegrid.Name = name
}

// Stage puts content to the model stage
func (content *Content) Stage(stage *Stage) *Content {
	if _, ok := stage.Contents[content]; !ok {
		stage.Contents[content] = struct{}{}
		stage.Content_stagedOrder[content] = stage.ContentOrder
		stage.ContentOrder++
	}
	stage.Contents_mapString[content.Name] = content

	return content
}

// StagePreserveOrder puts content to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ContentOrder
// - update stage.ContentOrder accordingly
func (content *Content) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Contents[content]; !ok {
		stage.Contents[content] = struct{}{}

		if order > stage.ContentOrder {
			stage.ContentOrder = order
		}
		stage.Content_stagedOrder[content] = order
		stage.ContentOrder++
	}
	stage.Contents_mapString[content.Name] = content
}

// Unstage removes content off the model stage
func (content *Content) Unstage(stage *Stage) *Content {
	delete(stage.Contents, content)
	// issue1150
	// delete(stage.Content_stagedOrder, content)
	delete(stage.Contents_mapString, content.Name)

	return content
}

// UnstageVoid removes content off the model stage
func (content *Content) UnstageVoid(stage *Stage) {
	delete(stage.Contents, content)
	// issue1150
	// delete(stage.Content_stagedOrder, content)
	delete(stage.Contents_mapString, content.Name)
}

// commit content to the back repo (if it is already staged)
func (content *Content) Commit(stage *Stage) *Content {
	if _, ok := stage.Contents[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitContent(content)
		}
	}
	return content
}

func (content *Content) CommitVoid(stage *Stage) {
	content.Commit(stage)
}

func (content *Content) StageVoid(stage *Stage) {
	content.Stage(stage)
}

// Checkout content to the back repo (if it is already staged)
func (content *Content) Checkout(stage *Stage) *Content {
	if _, ok := stage.Contents[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutContent(content)
		}
	}
	return content
}

// for satisfaction of GongStruct interface
func (content *Content) GetName() (res string) {
	return content.Name
}

// for satisfaction of GongStruct interface
func (content *Content) SetName(name string) {
	content.Name = name
}

// Stage puts exporttomusicxml to the model stage
func (exporttomusicxml *ExportToMusicxml) Stage(stage *Stage) *ExportToMusicxml {
	if _, ok := stage.ExportToMusicxmls[exporttomusicxml]; !ok {
		stage.ExportToMusicxmls[exporttomusicxml] = struct{}{}
		stage.ExportToMusicxml_stagedOrder[exporttomusicxml] = stage.ExportToMusicxmlOrder
		stage.ExportToMusicxmlOrder++
	}
	stage.ExportToMusicxmls_mapString[exporttomusicxml.Name] = exporttomusicxml

	return exporttomusicxml
}

// StagePreserveOrder puts exporttomusicxml to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExportToMusicxmlOrder
// - update stage.ExportToMusicxmlOrder accordingly
func (exporttomusicxml *ExportToMusicxml) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ExportToMusicxmls[exporttomusicxml]; !ok {
		stage.ExportToMusicxmls[exporttomusicxml] = struct{}{}

		if order > stage.ExportToMusicxmlOrder {
			stage.ExportToMusicxmlOrder = order
		}
		stage.ExportToMusicxml_stagedOrder[exporttomusicxml] = order
		stage.ExportToMusicxmlOrder++
	}
	stage.ExportToMusicxmls_mapString[exporttomusicxml.Name] = exporttomusicxml
}

// Unstage removes exporttomusicxml off the model stage
func (exporttomusicxml *ExportToMusicxml) Unstage(stage *Stage) *ExportToMusicxml {
	delete(stage.ExportToMusicxmls, exporttomusicxml)
	// issue1150
	// delete(stage.ExportToMusicxml_stagedOrder, exporttomusicxml)
	delete(stage.ExportToMusicxmls_mapString, exporttomusicxml.Name)

	return exporttomusicxml
}

// UnstageVoid removes exporttomusicxml off the model stage
func (exporttomusicxml *ExportToMusicxml) UnstageVoid(stage *Stage) {
	delete(stage.ExportToMusicxmls, exporttomusicxml)
	// issue1150
	// delete(stage.ExportToMusicxml_stagedOrder, exporttomusicxml)
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

func (exporttomusicxml *ExportToMusicxml) StageVoid(stage *Stage) {
	exporttomusicxml.Stage(stage)
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

// for satisfaction of GongStruct interface
func (exporttomusicxml *ExportToMusicxml) SetName(name string) {
	exporttomusicxml.Name = name
}

// Stage puts frontcurve to the model stage
func (frontcurve *FrontCurve) Stage(stage *Stage) *FrontCurve {
	if _, ok := stage.FrontCurves[frontcurve]; !ok {
		stage.FrontCurves[frontcurve] = struct{}{}
		stage.FrontCurve_stagedOrder[frontcurve] = stage.FrontCurveOrder
		stage.FrontCurveOrder++
	}
	stage.FrontCurves_mapString[frontcurve.Name] = frontcurve

	return frontcurve
}

// StagePreserveOrder puts frontcurve to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FrontCurveOrder
// - update stage.FrontCurveOrder accordingly
func (frontcurve *FrontCurve) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FrontCurves[frontcurve]; !ok {
		stage.FrontCurves[frontcurve] = struct{}{}

		if order > stage.FrontCurveOrder {
			stage.FrontCurveOrder = order
		}
		stage.FrontCurve_stagedOrder[frontcurve] = order
		stage.FrontCurveOrder++
	}
	stage.FrontCurves_mapString[frontcurve.Name] = frontcurve
}

// Unstage removes frontcurve off the model stage
func (frontcurve *FrontCurve) Unstage(stage *Stage) *FrontCurve {
	delete(stage.FrontCurves, frontcurve)
	// issue1150
	// delete(stage.FrontCurve_stagedOrder, frontcurve)
	delete(stage.FrontCurves_mapString, frontcurve.Name)

	return frontcurve
}

// UnstageVoid removes frontcurve off the model stage
func (frontcurve *FrontCurve) UnstageVoid(stage *Stage) {
	delete(stage.FrontCurves, frontcurve)
	// issue1150
	// delete(stage.FrontCurve_stagedOrder, frontcurve)
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

func (frontcurve *FrontCurve) StageVoid(stage *Stage) {
	frontcurve.Stage(stage)
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

// for satisfaction of GongStruct interface
func (frontcurve *FrontCurve) SetName(name string) {
	frontcurve.Name = name
}

// Stage puts frontcurvestack to the model stage
func (frontcurvestack *FrontCurveStack) Stage(stage *Stage) *FrontCurveStack {
	if _, ok := stage.FrontCurveStacks[frontcurvestack]; !ok {
		stage.FrontCurveStacks[frontcurvestack] = struct{}{}
		stage.FrontCurveStack_stagedOrder[frontcurvestack] = stage.FrontCurveStackOrder
		stage.FrontCurveStackOrder++
	}
	stage.FrontCurveStacks_mapString[frontcurvestack.Name] = frontcurvestack

	return frontcurvestack
}

// StagePreserveOrder puts frontcurvestack to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FrontCurveStackOrder
// - update stage.FrontCurveStackOrder accordingly
func (frontcurvestack *FrontCurveStack) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.FrontCurveStacks[frontcurvestack]; !ok {
		stage.FrontCurveStacks[frontcurvestack] = struct{}{}

		if order > stage.FrontCurveStackOrder {
			stage.FrontCurveStackOrder = order
		}
		stage.FrontCurveStack_stagedOrder[frontcurvestack] = order
		stage.FrontCurveStackOrder++
	}
	stage.FrontCurveStacks_mapString[frontcurvestack.Name] = frontcurvestack
}

// Unstage removes frontcurvestack off the model stage
func (frontcurvestack *FrontCurveStack) Unstage(stage *Stage) *FrontCurveStack {
	delete(stage.FrontCurveStacks, frontcurvestack)
	// issue1150
	// delete(stage.FrontCurveStack_stagedOrder, frontcurvestack)
	delete(stage.FrontCurveStacks_mapString, frontcurvestack.Name)

	return frontcurvestack
}

// UnstageVoid removes frontcurvestack off the model stage
func (frontcurvestack *FrontCurveStack) UnstageVoid(stage *Stage) {
	delete(stage.FrontCurveStacks, frontcurvestack)
	// issue1150
	// delete(stage.FrontCurveStack_stagedOrder, frontcurvestack)
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

func (frontcurvestack *FrontCurveStack) StageVoid(stage *Stage) {
	frontcurvestack.Stage(stage)
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

// for satisfaction of GongStruct interface
func (frontcurvestack *FrontCurveStack) SetName(name string) {
	frontcurvestack.Name = name
}

// Stage puts horizontalaxis to the model stage
func (horizontalaxis *HorizontalAxis) Stage(stage *Stage) *HorizontalAxis {
	if _, ok := stage.HorizontalAxiss[horizontalaxis]; !ok {
		stage.HorizontalAxiss[horizontalaxis] = struct{}{}
		stage.HorizontalAxis_stagedOrder[horizontalaxis] = stage.HorizontalAxisOrder
		stage.HorizontalAxisOrder++
	}
	stage.HorizontalAxiss_mapString[horizontalaxis.Name] = horizontalaxis

	return horizontalaxis
}

// StagePreserveOrder puts horizontalaxis to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HorizontalAxisOrder
// - update stage.HorizontalAxisOrder accordingly
func (horizontalaxis *HorizontalAxis) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.HorizontalAxiss[horizontalaxis]; !ok {
		stage.HorizontalAxiss[horizontalaxis] = struct{}{}

		if order > stage.HorizontalAxisOrder {
			stage.HorizontalAxisOrder = order
		}
		stage.HorizontalAxis_stagedOrder[horizontalaxis] = order
		stage.HorizontalAxisOrder++
	}
	stage.HorizontalAxiss_mapString[horizontalaxis.Name] = horizontalaxis
}

// Unstage removes horizontalaxis off the model stage
func (horizontalaxis *HorizontalAxis) Unstage(stage *Stage) *HorizontalAxis {
	delete(stage.HorizontalAxiss, horizontalaxis)
	// issue1150
	// delete(stage.HorizontalAxis_stagedOrder, horizontalaxis)
	delete(stage.HorizontalAxiss_mapString, horizontalaxis.Name)

	return horizontalaxis
}

// UnstageVoid removes horizontalaxis off the model stage
func (horizontalaxis *HorizontalAxis) UnstageVoid(stage *Stage) {
	delete(stage.HorizontalAxiss, horizontalaxis)
	// issue1150
	// delete(stage.HorizontalAxis_stagedOrder, horizontalaxis)
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

func (horizontalaxis *HorizontalAxis) StageVoid(stage *Stage) {
	horizontalaxis.Stage(stage)
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

// for satisfaction of GongStruct interface
func (horizontalaxis *HorizontalAxis) SetName(name string) {
	horizontalaxis.Name = name
}

// Stage puts key to the model stage
func (key *Key) Stage(stage *Stage) *Key {
	if _, ok := stage.Keys[key]; !ok {
		stage.Keys[key] = struct{}{}
		stage.Key_stagedOrder[key] = stage.KeyOrder
		stage.KeyOrder++
	}
	stage.Keys_mapString[key.Name] = key

	return key
}

// StagePreserveOrder puts key to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KeyOrder
// - update stage.KeyOrder accordingly
func (key *Key) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Keys[key]; !ok {
		stage.Keys[key] = struct{}{}

		if order > stage.KeyOrder {
			stage.KeyOrder = order
		}
		stage.Key_stagedOrder[key] = order
		stage.KeyOrder++
	}
	stage.Keys_mapString[key.Name] = key
}

// Unstage removes key off the model stage
func (key *Key) Unstage(stage *Stage) *Key {
	delete(stage.Keys, key)
	// issue1150
	// delete(stage.Key_stagedOrder, key)
	delete(stage.Keys_mapString, key.Name)

	return key
}

// UnstageVoid removes key off the model stage
func (key *Key) UnstageVoid(stage *Stage) {
	delete(stage.Keys, key)
	// issue1150
	// delete(stage.Key_stagedOrder, key)
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

func (key *Key) StageVoid(stage *Stage) {
	key.Stage(stage)
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

// for satisfaction of GongStruct interface
func (key *Key) SetName(name string) {
	key.Name = name
}

// Stage puts parameter to the model stage
func (parameter *Parameter) Stage(stage *Stage) *Parameter {
	if _, ok := stage.Parameters[parameter]; !ok {
		stage.Parameters[parameter] = struct{}{}
		stage.Parameter_stagedOrder[parameter] = stage.ParameterOrder
		stage.ParameterOrder++
	}
	stage.Parameters_mapString[parameter.Name] = parameter

	return parameter
}

// StagePreserveOrder puts parameter to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParameterOrder
// - update stage.ParameterOrder accordingly
func (parameter *Parameter) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Parameters[parameter]; !ok {
		stage.Parameters[parameter] = struct{}{}

		if order > stage.ParameterOrder {
			stage.ParameterOrder = order
		}
		stage.Parameter_stagedOrder[parameter] = order
		stage.ParameterOrder++
	}
	stage.Parameters_mapString[parameter.Name] = parameter
}

// Unstage removes parameter off the model stage
func (parameter *Parameter) Unstage(stage *Stage) *Parameter {
	delete(stage.Parameters, parameter)
	// issue1150
	// delete(stage.Parameter_stagedOrder, parameter)
	delete(stage.Parameters_mapString, parameter.Name)

	return parameter
}

// UnstageVoid removes parameter off the model stage
func (parameter *Parameter) UnstageVoid(stage *Stage) {
	delete(stage.Parameters, parameter)
	// issue1150
	// delete(stage.Parameter_stagedOrder, parameter)
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

func (parameter *Parameter) StageVoid(stage *Stage) {
	parameter.Stage(stage)
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

// for satisfaction of GongStruct interface
func (parameter *Parameter) SetName(name string) {
	parameter.Name = name
}

// Stage puts rhombus to the model stage
func (rhombus *Rhombus) Stage(stage *Stage) *Rhombus {
	if _, ok := stage.Rhombuss[rhombus]; !ok {
		stage.Rhombuss[rhombus] = struct{}{}
		stage.Rhombus_stagedOrder[rhombus] = stage.RhombusOrder
		stage.RhombusOrder++
	}
	stage.Rhombuss_mapString[rhombus.Name] = rhombus

	return rhombus
}

// StagePreserveOrder puts rhombus to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RhombusOrder
// - update stage.RhombusOrder accordingly
func (rhombus *Rhombus) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Rhombuss[rhombus]; !ok {
		stage.Rhombuss[rhombus] = struct{}{}

		if order > stage.RhombusOrder {
			stage.RhombusOrder = order
		}
		stage.Rhombus_stagedOrder[rhombus] = order
		stage.RhombusOrder++
	}
	stage.Rhombuss_mapString[rhombus.Name] = rhombus
}

// Unstage removes rhombus off the model stage
func (rhombus *Rhombus) Unstage(stage *Stage) *Rhombus {
	delete(stage.Rhombuss, rhombus)
	// issue1150
	// delete(stage.Rhombus_stagedOrder, rhombus)
	delete(stage.Rhombuss_mapString, rhombus.Name)

	return rhombus
}

// UnstageVoid removes rhombus off the model stage
func (rhombus *Rhombus) UnstageVoid(stage *Stage) {
	delete(stage.Rhombuss, rhombus)
	// issue1150
	// delete(stage.Rhombus_stagedOrder, rhombus)
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

func (rhombus *Rhombus) StageVoid(stage *Stage) {
	rhombus.Stage(stage)
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

// for satisfaction of GongStruct interface
func (rhombus *Rhombus) SetName(name string) {
	rhombus.Name = name
}

// Stage puts rhombusgrid to the model stage
func (rhombusgrid *RhombusGrid) Stage(stage *Stage) *RhombusGrid {
	if _, ok := stage.RhombusGrids[rhombusgrid]; !ok {
		stage.RhombusGrids[rhombusgrid] = struct{}{}
		stage.RhombusGrid_stagedOrder[rhombusgrid] = stage.RhombusGridOrder
		stage.RhombusGridOrder++
	}
	stage.RhombusGrids_mapString[rhombusgrid.Name] = rhombusgrid

	return rhombusgrid
}

// StagePreserveOrder puts rhombusgrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RhombusGridOrder
// - update stage.RhombusGridOrder accordingly
func (rhombusgrid *RhombusGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.RhombusGrids[rhombusgrid]; !ok {
		stage.RhombusGrids[rhombusgrid] = struct{}{}

		if order > stage.RhombusGridOrder {
			stage.RhombusGridOrder = order
		}
		stage.RhombusGrid_stagedOrder[rhombusgrid] = order
		stage.RhombusGridOrder++
	}
	stage.RhombusGrids_mapString[rhombusgrid.Name] = rhombusgrid
}

// Unstage removes rhombusgrid off the model stage
func (rhombusgrid *RhombusGrid) Unstage(stage *Stage) *RhombusGrid {
	delete(stage.RhombusGrids, rhombusgrid)
	// issue1150
	// delete(stage.RhombusGrid_stagedOrder, rhombusgrid)
	delete(stage.RhombusGrids_mapString, rhombusgrid.Name)

	return rhombusgrid
}

// UnstageVoid removes rhombusgrid off the model stage
func (rhombusgrid *RhombusGrid) UnstageVoid(stage *Stage) {
	delete(stage.RhombusGrids, rhombusgrid)
	// issue1150
	// delete(stage.RhombusGrid_stagedOrder, rhombusgrid)
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

func (rhombusgrid *RhombusGrid) StageVoid(stage *Stage) {
	rhombusgrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (rhombusgrid *RhombusGrid) SetName(name string) {
	rhombusgrid.Name = name
}

// Stage puts shapecategory to the model stage
func (shapecategory *ShapeCategory) Stage(stage *Stage) *ShapeCategory {
	if _, ok := stage.ShapeCategorys[shapecategory]; !ok {
		stage.ShapeCategorys[shapecategory] = struct{}{}
		stage.ShapeCategory_stagedOrder[shapecategory] = stage.ShapeCategoryOrder
		stage.ShapeCategoryOrder++
	}
	stage.ShapeCategorys_mapString[shapecategory.Name] = shapecategory

	return shapecategory
}

// StagePreserveOrder puts shapecategory to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ShapeCategoryOrder
// - update stage.ShapeCategoryOrder accordingly
func (shapecategory *ShapeCategory) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ShapeCategorys[shapecategory]; !ok {
		stage.ShapeCategorys[shapecategory] = struct{}{}

		if order > stage.ShapeCategoryOrder {
			stage.ShapeCategoryOrder = order
		}
		stage.ShapeCategory_stagedOrder[shapecategory] = order
		stage.ShapeCategoryOrder++
	}
	stage.ShapeCategorys_mapString[shapecategory.Name] = shapecategory
}

// Unstage removes shapecategory off the model stage
func (shapecategory *ShapeCategory) Unstage(stage *Stage) *ShapeCategory {
	delete(stage.ShapeCategorys, shapecategory)
	// issue1150
	// delete(stage.ShapeCategory_stagedOrder, shapecategory)
	delete(stage.ShapeCategorys_mapString, shapecategory.Name)

	return shapecategory
}

// UnstageVoid removes shapecategory off the model stage
func (shapecategory *ShapeCategory) UnstageVoid(stage *Stage) {
	delete(stage.ShapeCategorys, shapecategory)
	// issue1150
	// delete(stage.ShapeCategory_stagedOrder, shapecategory)
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

func (shapecategory *ShapeCategory) StageVoid(stage *Stage) {
	shapecategory.Stage(stage)
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

// for satisfaction of GongStruct interface
func (shapecategory *ShapeCategory) SetName(name string) {
	shapecategory.Name = name
}

// Stage puts spiralbezier to the model stage
func (spiralbezier *SpiralBezier) Stage(stage *Stage) *SpiralBezier {
	if _, ok := stage.SpiralBeziers[spiralbezier]; !ok {
		stage.SpiralBeziers[spiralbezier] = struct{}{}
		stage.SpiralBezier_stagedOrder[spiralbezier] = stage.SpiralBezierOrder
		stage.SpiralBezierOrder++
	}
	stage.SpiralBeziers_mapString[spiralbezier.Name] = spiralbezier

	return spiralbezier
}

// StagePreserveOrder puts spiralbezier to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralBezierOrder
// - update stage.SpiralBezierOrder accordingly
func (spiralbezier *SpiralBezier) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralBeziers[spiralbezier]; !ok {
		stage.SpiralBeziers[spiralbezier] = struct{}{}

		if order > stage.SpiralBezierOrder {
			stage.SpiralBezierOrder = order
		}
		stage.SpiralBezier_stagedOrder[spiralbezier] = order
		stage.SpiralBezierOrder++
	}
	stage.SpiralBeziers_mapString[spiralbezier.Name] = spiralbezier
}

// Unstage removes spiralbezier off the model stage
func (spiralbezier *SpiralBezier) Unstage(stage *Stage) *SpiralBezier {
	delete(stage.SpiralBeziers, spiralbezier)
	// issue1150
	// delete(stage.SpiralBezier_stagedOrder, spiralbezier)
	delete(stage.SpiralBeziers_mapString, spiralbezier.Name)

	return spiralbezier
}

// UnstageVoid removes spiralbezier off the model stage
func (spiralbezier *SpiralBezier) UnstageVoid(stage *Stage) {
	delete(stage.SpiralBeziers, spiralbezier)
	// issue1150
	// delete(stage.SpiralBezier_stagedOrder, spiralbezier)
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

func (spiralbezier *SpiralBezier) StageVoid(stage *Stage) {
	spiralbezier.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralbezier *SpiralBezier) SetName(name string) {
	spiralbezier.Name = name
}

// Stage puts spiralbeziergrid to the model stage
func (spiralbeziergrid *SpiralBezierGrid) Stage(stage *Stage) *SpiralBezierGrid {
	if _, ok := stage.SpiralBezierGrids[spiralbeziergrid]; !ok {
		stage.SpiralBezierGrids[spiralbeziergrid] = struct{}{}
		stage.SpiralBezierGrid_stagedOrder[spiralbeziergrid] = stage.SpiralBezierGridOrder
		stage.SpiralBezierGridOrder++
	}
	stage.SpiralBezierGrids_mapString[spiralbeziergrid.Name] = spiralbeziergrid

	return spiralbeziergrid
}

// StagePreserveOrder puts spiralbeziergrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralBezierGridOrder
// - update stage.SpiralBezierGridOrder accordingly
func (spiralbeziergrid *SpiralBezierGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralBezierGrids[spiralbeziergrid]; !ok {
		stage.SpiralBezierGrids[spiralbeziergrid] = struct{}{}

		if order > stage.SpiralBezierGridOrder {
			stage.SpiralBezierGridOrder = order
		}
		stage.SpiralBezierGrid_stagedOrder[spiralbeziergrid] = order
		stage.SpiralBezierGridOrder++
	}
	stage.SpiralBezierGrids_mapString[spiralbeziergrid.Name] = spiralbeziergrid
}

// Unstage removes spiralbeziergrid off the model stage
func (spiralbeziergrid *SpiralBezierGrid) Unstage(stage *Stage) *SpiralBezierGrid {
	delete(stage.SpiralBezierGrids, spiralbeziergrid)
	// issue1150
	// delete(stage.SpiralBezierGrid_stagedOrder, spiralbeziergrid)
	delete(stage.SpiralBezierGrids_mapString, spiralbeziergrid.Name)

	return spiralbeziergrid
}

// UnstageVoid removes spiralbeziergrid off the model stage
func (spiralbeziergrid *SpiralBezierGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralBezierGrids, spiralbeziergrid)
	// issue1150
	// delete(stage.SpiralBezierGrid_stagedOrder, spiralbeziergrid)
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

func (spiralbeziergrid *SpiralBezierGrid) StageVoid(stage *Stage) {
	spiralbeziergrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralbeziergrid *SpiralBezierGrid) SetName(name string) {
	spiralbeziergrid.Name = name
}

// Stage puts spiralcircle to the model stage
func (spiralcircle *SpiralCircle) Stage(stage *Stage) *SpiralCircle {
	if _, ok := stage.SpiralCircles[spiralcircle]; !ok {
		stage.SpiralCircles[spiralcircle] = struct{}{}
		stage.SpiralCircle_stagedOrder[spiralcircle] = stage.SpiralCircleOrder
		stage.SpiralCircleOrder++
	}
	stage.SpiralCircles_mapString[spiralcircle.Name] = spiralcircle

	return spiralcircle
}

// StagePreserveOrder puts spiralcircle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralCircleOrder
// - update stage.SpiralCircleOrder accordingly
func (spiralcircle *SpiralCircle) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralCircles[spiralcircle]; !ok {
		stage.SpiralCircles[spiralcircle] = struct{}{}

		if order > stage.SpiralCircleOrder {
			stage.SpiralCircleOrder = order
		}
		stage.SpiralCircle_stagedOrder[spiralcircle] = order
		stage.SpiralCircleOrder++
	}
	stage.SpiralCircles_mapString[spiralcircle.Name] = spiralcircle
}

// Unstage removes spiralcircle off the model stage
func (spiralcircle *SpiralCircle) Unstage(stage *Stage) *SpiralCircle {
	delete(stage.SpiralCircles, spiralcircle)
	// issue1150
	// delete(stage.SpiralCircle_stagedOrder, spiralcircle)
	delete(stage.SpiralCircles_mapString, spiralcircle.Name)

	return spiralcircle
}

// UnstageVoid removes spiralcircle off the model stage
func (spiralcircle *SpiralCircle) UnstageVoid(stage *Stage) {
	delete(stage.SpiralCircles, spiralcircle)
	// issue1150
	// delete(stage.SpiralCircle_stagedOrder, spiralcircle)
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

func (spiralcircle *SpiralCircle) StageVoid(stage *Stage) {
	spiralcircle.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralcircle *SpiralCircle) SetName(name string) {
	spiralcircle.Name = name
}

// Stage puts spiralcirclegrid to the model stage
func (spiralcirclegrid *SpiralCircleGrid) Stage(stage *Stage) *SpiralCircleGrid {
	if _, ok := stage.SpiralCircleGrids[spiralcirclegrid]; !ok {
		stage.SpiralCircleGrids[spiralcirclegrid] = struct{}{}
		stage.SpiralCircleGrid_stagedOrder[spiralcirclegrid] = stage.SpiralCircleGridOrder
		stage.SpiralCircleGridOrder++
	}
	stage.SpiralCircleGrids_mapString[spiralcirclegrid.Name] = spiralcirclegrid

	return spiralcirclegrid
}

// StagePreserveOrder puts spiralcirclegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralCircleGridOrder
// - update stage.SpiralCircleGridOrder accordingly
func (spiralcirclegrid *SpiralCircleGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralCircleGrids[spiralcirclegrid]; !ok {
		stage.SpiralCircleGrids[spiralcirclegrid] = struct{}{}

		if order > stage.SpiralCircleGridOrder {
			stage.SpiralCircleGridOrder = order
		}
		stage.SpiralCircleGrid_stagedOrder[spiralcirclegrid] = order
		stage.SpiralCircleGridOrder++
	}
	stage.SpiralCircleGrids_mapString[spiralcirclegrid.Name] = spiralcirclegrid
}

// Unstage removes spiralcirclegrid off the model stage
func (spiralcirclegrid *SpiralCircleGrid) Unstage(stage *Stage) *SpiralCircleGrid {
	delete(stage.SpiralCircleGrids, spiralcirclegrid)
	// issue1150
	// delete(stage.SpiralCircleGrid_stagedOrder, spiralcirclegrid)
	delete(stage.SpiralCircleGrids_mapString, spiralcirclegrid.Name)

	return spiralcirclegrid
}

// UnstageVoid removes spiralcirclegrid off the model stage
func (spiralcirclegrid *SpiralCircleGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralCircleGrids, spiralcirclegrid)
	// issue1150
	// delete(stage.SpiralCircleGrid_stagedOrder, spiralcirclegrid)
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

func (spiralcirclegrid *SpiralCircleGrid) StageVoid(stage *Stage) {
	spiralcirclegrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralcirclegrid *SpiralCircleGrid) SetName(name string) {
	spiralcirclegrid.Name = name
}

// Stage puts spiralline to the model stage
func (spiralline *SpiralLine) Stage(stage *Stage) *SpiralLine {
	if _, ok := stage.SpiralLines[spiralline]; !ok {
		stage.SpiralLines[spiralline] = struct{}{}
		stage.SpiralLine_stagedOrder[spiralline] = stage.SpiralLineOrder
		stage.SpiralLineOrder++
	}
	stage.SpiralLines_mapString[spiralline.Name] = spiralline

	return spiralline
}

// StagePreserveOrder puts spiralline to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralLineOrder
// - update stage.SpiralLineOrder accordingly
func (spiralline *SpiralLine) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralLines[spiralline]; !ok {
		stage.SpiralLines[spiralline] = struct{}{}

		if order > stage.SpiralLineOrder {
			stage.SpiralLineOrder = order
		}
		stage.SpiralLine_stagedOrder[spiralline] = order
		stage.SpiralLineOrder++
	}
	stage.SpiralLines_mapString[spiralline.Name] = spiralline
}

// Unstage removes spiralline off the model stage
func (spiralline *SpiralLine) Unstage(stage *Stage) *SpiralLine {
	delete(stage.SpiralLines, spiralline)
	// issue1150
	// delete(stage.SpiralLine_stagedOrder, spiralline)
	delete(stage.SpiralLines_mapString, spiralline.Name)

	return spiralline
}

// UnstageVoid removes spiralline off the model stage
func (spiralline *SpiralLine) UnstageVoid(stage *Stage) {
	delete(stage.SpiralLines, spiralline)
	// issue1150
	// delete(stage.SpiralLine_stagedOrder, spiralline)
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

func (spiralline *SpiralLine) StageVoid(stage *Stage) {
	spiralline.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralline *SpiralLine) SetName(name string) {
	spiralline.Name = name
}

// Stage puts spirallinegrid to the model stage
func (spirallinegrid *SpiralLineGrid) Stage(stage *Stage) *SpiralLineGrid {
	if _, ok := stage.SpiralLineGrids[spirallinegrid]; !ok {
		stage.SpiralLineGrids[spirallinegrid] = struct{}{}
		stage.SpiralLineGrid_stagedOrder[spirallinegrid] = stage.SpiralLineGridOrder
		stage.SpiralLineGridOrder++
	}
	stage.SpiralLineGrids_mapString[spirallinegrid.Name] = spirallinegrid

	return spirallinegrid
}

// StagePreserveOrder puts spirallinegrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralLineGridOrder
// - update stage.SpiralLineGridOrder accordingly
func (spirallinegrid *SpiralLineGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralLineGrids[spirallinegrid]; !ok {
		stage.SpiralLineGrids[spirallinegrid] = struct{}{}

		if order > stage.SpiralLineGridOrder {
			stage.SpiralLineGridOrder = order
		}
		stage.SpiralLineGrid_stagedOrder[spirallinegrid] = order
		stage.SpiralLineGridOrder++
	}
	stage.SpiralLineGrids_mapString[spirallinegrid.Name] = spirallinegrid
}

// Unstage removes spirallinegrid off the model stage
func (spirallinegrid *SpiralLineGrid) Unstage(stage *Stage) *SpiralLineGrid {
	delete(stage.SpiralLineGrids, spirallinegrid)
	// issue1150
	// delete(stage.SpiralLineGrid_stagedOrder, spirallinegrid)
	delete(stage.SpiralLineGrids_mapString, spirallinegrid.Name)

	return spirallinegrid
}

// UnstageVoid removes spirallinegrid off the model stage
func (spirallinegrid *SpiralLineGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralLineGrids, spirallinegrid)
	// issue1150
	// delete(stage.SpiralLineGrid_stagedOrder, spirallinegrid)
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

func (spirallinegrid *SpiralLineGrid) StageVoid(stage *Stage) {
	spirallinegrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spirallinegrid *SpiralLineGrid) SetName(name string) {
	spirallinegrid.Name = name
}

// Stage puts spiralorigin to the model stage
func (spiralorigin *SpiralOrigin) Stage(stage *Stage) *SpiralOrigin {
	if _, ok := stage.SpiralOrigins[spiralorigin]; !ok {
		stage.SpiralOrigins[spiralorigin] = struct{}{}
		stage.SpiralOrigin_stagedOrder[spiralorigin] = stage.SpiralOriginOrder
		stage.SpiralOriginOrder++
	}
	stage.SpiralOrigins_mapString[spiralorigin.Name] = spiralorigin

	return spiralorigin
}

// StagePreserveOrder puts spiralorigin to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralOriginOrder
// - update stage.SpiralOriginOrder accordingly
func (spiralorigin *SpiralOrigin) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralOrigins[spiralorigin]; !ok {
		stage.SpiralOrigins[spiralorigin] = struct{}{}

		if order > stage.SpiralOriginOrder {
			stage.SpiralOriginOrder = order
		}
		stage.SpiralOrigin_stagedOrder[spiralorigin] = order
		stage.SpiralOriginOrder++
	}
	stage.SpiralOrigins_mapString[spiralorigin.Name] = spiralorigin
}

// Unstage removes spiralorigin off the model stage
func (spiralorigin *SpiralOrigin) Unstage(stage *Stage) *SpiralOrigin {
	delete(stage.SpiralOrigins, spiralorigin)
	// issue1150
	// delete(stage.SpiralOrigin_stagedOrder, spiralorigin)
	delete(stage.SpiralOrigins_mapString, spiralorigin.Name)

	return spiralorigin
}

// UnstageVoid removes spiralorigin off the model stage
func (spiralorigin *SpiralOrigin) UnstageVoid(stage *Stage) {
	delete(stage.SpiralOrigins, spiralorigin)
	// issue1150
	// delete(stage.SpiralOrigin_stagedOrder, spiralorigin)
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

func (spiralorigin *SpiralOrigin) StageVoid(stage *Stage) {
	spiralorigin.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralorigin *SpiralOrigin) SetName(name string) {
	spiralorigin.Name = name
}

// Stage puts spiralrhombus to the model stage
func (spiralrhombus *SpiralRhombus) Stage(stage *Stage) *SpiralRhombus {
	if _, ok := stage.SpiralRhombuss[spiralrhombus]; !ok {
		stage.SpiralRhombuss[spiralrhombus] = struct{}{}
		stage.SpiralRhombus_stagedOrder[spiralrhombus] = stage.SpiralRhombusOrder
		stage.SpiralRhombusOrder++
	}
	stage.SpiralRhombuss_mapString[spiralrhombus.Name] = spiralrhombus

	return spiralrhombus
}

// StagePreserveOrder puts spiralrhombus to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralRhombusOrder
// - update stage.SpiralRhombusOrder accordingly
func (spiralrhombus *SpiralRhombus) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralRhombuss[spiralrhombus]; !ok {
		stage.SpiralRhombuss[spiralrhombus] = struct{}{}

		if order > stage.SpiralRhombusOrder {
			stage.SpiralRhombusOrder = order
		}
		stage.SpiralRhombus_stagedOrder[spiralrhombus] = order
		stage.SpiralRhombusOrder++
	}
	stage.SpiralRhombuss_mapString[spiralrhombus.Name] = spiralrhombus
}

// Unstage removes spiralrhombus off the model stage
func (spiralrhombus *SpiralRhombus) Unstage(stage *Stage) *SpiralRhombus {
	delete(stage.SpiralRhombuss, spiralrhombus)
	// issue1150
	// delete(stage.SpiralRhombus_stagedOrder, spiralrhombus)
	delete(stage.SpiralRhombuss_mapString, spiralrhombus.Name)

	return spiralrhombus
}

// UnstageVoid removes spiralrhombus off the model stage
func (spiralrhombus *SpiralRhombus) UnstageVoid(stage *Stage) {
	delete(stage.SpiralRhombuss, spiralrhombus)
	// issue1150
	// delete(stage.SpiralRhombus_stagedOrder, spiralrhombus)
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

func (spiralrhombus *SpiralRhombus) StageVoid(stage *Stage) {
	spiralrhombus.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralrhombus *SpiralRhombus) SetName(name string) {
	spiralrhombus.Name = name
}

// Stage puts spiralrhombusgrid to the model stage
func (spiralrhombusgrid *SpiralRhombusGrid) Stage(stage *Stage) *SpiralRhombusGrid {
	if _, ok := stage.SpiralRhombusGrids[spiralrhombusgrid]; !ok {
		stage.SpiralRhombusGrids[spiralrhombusgrid] = struct{}{}
		stage.SpiralRhombusGrid_stagedOrder[spiralrhombusgrid] = stage.SpiralRhombusGridOrder
		stage.SpiralRhombusGridOrder++
	}
	stage.SpiralRhombusGrids_mapString[spiralrhombusgrid.Name] = spiralrhombusgrid

	return spiralrhombusgrid
}

// StagePreserveOrder puts spiralrhombusgrid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpiralRhombusGridOrder
// - update stage.SpiralRhombusGridOrder accordingly
func (spiralrhombusgrid *SpiralRhombusGrid) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SpiralRhombusGrids[spiralrhombusgrid]; !ok {
		stage.SpiralRhombusGrids[spiralrhombusgrid] = struct{}{}

		if order > stage.SpiralRhombusGridOrder {
			stage.SpiralRhombusGridOrder = order
		}
		stage.SpiralRhombusGrid_stagedOrder[spiralrhombusgrid] = order
		stage.SpiralRhombusGridOrder++
	}
	stage.SpiralRhombusGrids_mapString[spiralrhombusgrid.Name] = spiralrhombusgrid
}

// Unstage removes spiralrhombusgrid off the model stage
func (spiralrhombusgrid *SpiralRhombusGrid) Unstage(stage *Stage) *SpiralRhombusGrid {
	delete(stage.SpiralRhombusGrids, spiralrhombusgrid)
	// issue1150
	// delete(stage.SpiralRhombusGrid_stagedOrder, spiralrhombusgrid)
	delete(stage.SpiralRhombusGrids_mapString, spiralrhombusgrid.Name)

	return spiralrhombusgrid
}

// UnstageVoid removes spiralrhombusgrid off the model stage
func (spiralrhombusgrid *SpiralRhombusGrid) UnstageVoid(stage *Stage) {
	delete(stage.SpiralRhombusGrids, spiralrhombusgrid)
	// issue1150
	// delete(stage.SpiralRhombusGrid_stagedOrder, spiralrhombusgrid)
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

func (spiralrhombusgrid *SpiralRhombusGrid) StageVoid(stage *Stage) {
	spiralrhombusgrid.Stage(stage)
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

// for satisfaction of GongStruct interface
func (spiralrhombusgrid *SpiralRhombusGrid) SetName(name string) {
	spiralrhombusgrid.Name = name
}

// Stage puts verticalaxis to the model stage
func (verticalaxis *VerticalAxis) Stage(stage *Stage) *VerticalAxis {
	if _, ok := stage.VerticalAxiss[verticalaxis]; !ok {
		stage.VerticalAxiss[verticalaxis] = struct{}{}
		stage.VerticalAxis_stagedOrder[verticalaxis] = stage.VerticalAxisOrder
		stage.VerticalAxisOrder++
	}
	stage.VerticalAxiss_mapString[verticalaxis.Name] = verticalaxis

	return verticalaxis
}

// StagePreserveOrder puts verticalaxis to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalAxisOrder
// - update stage.VerticalAxisOrder accordingly
func (verticalaxis *VerticalAxis) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.VerticalAxiss[verticalaxis]; !ok {
		stage.VerticalAxiss[verticalaxis] = struct{}{}

		if order > stage.VerticalAxisOrder {
			stage.VerticalAxisOrder = order
		}
		stage.VerticalAxis_stagedOrder[verticalaxis] = order
		stage.VerticalAxisOrder++
	}
	stage.VerticalAxiss_mapString[verticalaxis.Name] = verticalaxis
}

// Unstage removes verticalaxis off the model stage
func (verticalaxis *VerticalAxis) Unstage(stage *Stage) *VerticalAxis {
	delete(stage.VerticalAxiss, verticalaxis)
	// issue1150
	// delete(stage.VerticalAxis_stagedOrder, verticalaxis)
	delete(stage.VerticalAxiss_mapString, verticalaxis.Name)

	return verticalaxis
}

// UnstageVoid removes verticalaxis off the model stage
func (verticalaxis *VerticalAxis) UnstageVoid(stage *Stage) {
	delete(stage.VerticalAxiss, verticalaxis)
	// issue1150
	// delete(stage.VerticalAxis_stagedOrder, verticalaxis)
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

func (verticalaxis *VerticalAxis) StageVoid(stage *Stage) {
	verticalaxis.Stage(stage)
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

// for satisfaction of GongStruct interface
func (verticalaxis *VerticalAxis) SetName(name string) {
	verticalaxis.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAxis(Axis *Axis)
	CreateORMAxisGrid(AxisGrid *AxisGrid)
	CreateORMBezier(Bezier *Bezier)
	CreateORMBezierGrid(BezierGrid *BezierGrid)
	CreateORMBezierGridStack(BezierGridStack *BezierGridStack)
	CreateORMChapter(Chapter *Chapter)
	CreateORMCircle(Circle *Circle)
	CreateORMCircleGrid(CircleGrid *CircleGrid)
	CreateORMContent(Content *Content)
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
	DeleteORMChapter(Chapter *Chapter)
	DeleteORMCircle(Circle *Circle)
	DeleteORMCircleGrid(CircleGrid *CircleGrid)
	DeleteORMContent(Content *Content)
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
	stage.Axiss = make(map[*Axis]struct{})
	stage.Axiss_mapString = make(map[string]*Axis)
	stage.Axis_stagedOrder = make(map[*Axis]uint)
	stage.AxisOrder = 0

	stage.AxisGrids = make(map[*AxisGrid]struct{})
	stage.AxisGrids_mapString = make(map[string]*AxisGrid)
	stage.AxisGrid_stagedOrder = make(map[*AxisGrid]uint)
	stage.AxisGridOrder = 0

	stage.Beziers = make(map[*Bezier]struct{})
	stage.Beziers_mapString = make(map[string]*Bezier)
	stage.Bezier_stagedOrder = make(map[*Bezier]uint)
	stage.BezierOrder = 0

	stage.BezierGrids = make(map[*BezierGrid]struct{})
	stage.BezierGrids_mapString = make(map[string]*BezierGrid)
	stage.BezierGrid_stagedOrder = make(map[*BezierGrid]uint)
	stage.BezierGridOrder = 0

	stage.BezierGridStacks = make(map[*BezierGridStack]struct{})
	stage.BezierGridStacks_mapString = make(map[string]*BezierGridStack)
	stage.BezierGridStack_stagedOrder = make(map[*BezierGridStack]uint)
	stage.BezierGridStackOrder = 0

	stage.Chapters = make(map[*Chapter]struct{})
	stage.Chapters_mapString = make(map[string]*Chapter)
	stage.Chapter_stagedOrder = make(map[*Chapter]uint)
	stage.ChapterOrder = 0

	stage.Circles = make(map[*Circle]struct{})
	stage.Circles_mapString = make(map[string]*Circle)
	stage.Circle_stagedOrder = make(map[*Circle]uint)
	stage.CircleOrder = 0

	stage.CircleGrids = make(map[*CircleGrid]struct{})
	stage.CircleGrids_mapString = make(map[string]*CircleGrid)
	stage.CircleGrid_stagedOrder = make(map[*CircleGrid]uint)
	stage.CircleGridOrder = 0

	stage.Contents = make(map[*Content]struct{})
	stage.Contents_mapString = make(map[string]*Content)
	stage.Content_stagedOrder = make(map[*Content]uint)
	stage.ContentOrder = 0

	stage.ExportToMusicxmls = make(map[*ExportToMusicxml]struct{})
	stage.ExportToMusicxmls_mapString = make(map[string]*ExportToMusicxml)
	stage.ExportToMusicxml_stagedOrder = make(map[*ExportToMusicxml]uint)
	stage.ExportToMusicxmlOrder = 0

	stage.FrontCurves = make(map[*FrontCurve]struct{})
	stage.FrontCurves_mapString = make(map[string]*FrontCurve)
	stage.FrontCurve_stagedOrder = make(map[*FrontCurve]uint)
	stage.FrontCurveOrder = 0

	stage.FrontCurveStacks = make(map[*FrontCurveStack]struct{})
	stage.FrontCurveStacks_mapString = make(map[string]*FrontCurveStack)
	stage.FrontCurveStack_stagedOrder = make(map[*FrontCurveStack]uint)
	stage.FrontCurveStackOrder = 0

	stage.HorizontalAxiss = make(map[*HorizontalAxis]struct{})
	stage.HorizontalAxiss_mapString = make(map[string]*HorizontalAxis)
	stage.HorizontalAxis_stagedOrder = make(map[*HorizontalAxis]uint)
	stage.HorizontalAxisOrder = 0

	stage.Keys = make(map[*Key]struct{})
	stage.Keys_mapString = make(map[string]*Key)
	stage.Key_stagedOrder = make(map[*Key]uint)
	stage.KeyOrder = 0

	stage.Parameters = make(map[*Parameter]struct{})
	stage.Parameters_mapString = make(map[string]*Parameter)
	stage.Parameter_stagedOrder = make(map[*Parameter]uint)
	stage.ParameterOrder = 0

	stage.Rhombuss = make(map[*Rhombus]struct{})
	stage.Rhombuss_mapString = make(map[string]*Rhombus)
	stage.Rhombus_stagedOrder = make(map[*Rhombus]uint)
	stage.RhombusOrder = 0

	stage.RhombusGrids = make(map[*RhombusGrid]struct{})
	stage.RhombusGrids_mapString = make(map[string]*RhombusGrid)
	stage.RhombusGrid_stagedOrder = make(map[*RhombusGrid]uint)
	stage.RhombusGridOrder = 0

	stage.ShapeCategorys = make(map[*ShapeCategory]struct{})
	stage.ShapeCategorys_mapString = make(map[string]*ShapeCategory)
	stage.ShapeCategory_stagedOrder = make(map[*ShapeCategory]uint)
	stage.ShapeCategoryOrder = 0

	stage.SpiralBeziers = make(map[*SpiralBezier]struct{})
	stage.SpiralBeziers_mapString = make(map[string]*SpiralBezier)
	stage.SpiralBezier_stagedOrder = make(map[*SpiralBezier]uint)
	stage.SpiralBezierOrder = 0

	stage.SpiralBezierGrids = make(map[*SpiralBezierGrid]struct{})
	stage.SpiralBezierGrids_mapString = make(map[string]*SpiralBezierGrid)
	stage.SpiralBezierGrid_stagedOrder = make(map[*SpiralBezierGrid]uint)
	stage.SpiralBezierGridOrder = 0

	stage.SpiralCircles = make(map[*SpiralCircle]struct{})
	stage.SpiralCircles_mapString = make(map[string]*SpiralCircle)
	stage.SpiralCircle_stagedOrder = make(map[*SpiralCircle]uint)
	stage.SpiralCircleOrder = 0

	stage.SpiralCircleGrids = make(map[*SpiralCircleGrid]struct{})
	stage.SpiralCircleGrids_mapString = make(map[string]*SpiralCircleGrid)
	stage.SpiralCircleGrid_stagedOrder = make(map[*SpiralCircleGrid]uint)
	stage.SpiralCircleGridOrder = 0

	stage.SpiralLines = make(map[*SpiralLine]struct{})
	stage.SpiralLines_mapString = make(map[string]*SpiralLine)
	stage.SpiralLine_stagedOrder = make(map[*SpiralLine]uint)
	stage.SpiralLineOrder = 0

	stage.SpiralLineGrids = make(map[*SpiralLineGrid]struct{})
	stage.SpiralLineGrids_mapString = make(map[string]*SpiralLineGrid)
	stage.SpiralLineGrid_stagedOrder = make(map[*SpiralLineGrid]uint)
	stage.SpiralLineGridOrder = 0

	stage.SpiralOrigins = make(map[*SpiralOrigin]struct{})
	stage.SpiralOrigins_mapString = make(map[string]*SpiralOrigin)
	stage.SpiralOrigin_stagedOrder = make(map[*SpiralOrigin]uint)
	stage.SpiralOriginOrder = 0

	stage.SpiralRhombuss = make(map[*SpiralRhombus]struct{})
	stage.SpiralRhombuss_mapString = make(map[string]*SpiralRhombus)
	stage.SpiralRhombus_stagedOrder = make(map[*SpiralRhombus]uint)
	stage.SpiralRhombusOrder = 0

	stage.SpiralRhombusGrids = make(map[*SpiralRhombusGrid]struct{})
	stage.SpiralRhombusGrids_mapString = make(map[string]*SpiralRhombusGrid)
	stage.SpiralRhombusGrid_stagedOrder = make(map[*SpiralRhombusGrid]uint)
	stage.SpiralRhombusGridOrder = 0

	stage.VerticalAxiss = make(map[*VerticalAxis]struct{})
	stage.VerticalAxiss_mapString = make(map[string]*VerticalAxis)
	stage.VerticalAxis_stagedOrder = make(map[*VerticalAxis]uint)
	stage.VerticalAxisOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
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

	stage.Chapters = nil
	stage.Chapters_mapString = nil

	stage.Circles = nil
	stage.Circles_mapString = nil

	stage.CircleGrids = nil
	stage.CircleGrids_mapString = nil

	stage.Contents = nil
	stage.Contents_mapString = nil

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

	// end of insertion point for array nil
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

	for chapter := range stage.Chapters {
		chapter.Unstage(stage)
	}

	for circle := range stage.Circles {
		circle.Unstage(stage)
	}

	for circlegrid := range stage.CircleGrids {
		circlegrid.Unstage(stage)
	}

	for content := range stage.Contents {
		content.Unstage(stage)
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

	// end of insertion point for array nil
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface{}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage) (modified bool)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) GongstructIF
	GongGetUUID(stage *Stage) string
}
type PointerToGongstruct interface {
	GongstructIF
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]struct{}) (sortedSlice []T) {
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
	case map[*Chapter]any:
		return any(&stage.Chapters).(*Type)
	case map[*Circle]any:
		return any(&stage.Circles).(*Type)
	case map[*CircleGrid]any:
		return any(&stage.CircleGrids).(*Type)
	case map[*Content]any:
		return any(&stage.Contents).(*Type)
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

// GongGetMap returns the map of staged Gonstruct instance by their name
// Can be usefull if names are unique
func GongGetMap[Type GongstructIF](stage *Stage) map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Axis:
		return any(stage.Axiss_mapString).(map[string]Type)
	case *AxisGrid:
		return any(stage.AxisGrids_mapString).(map[string]Type)
	case *Bezier:
		return any(stage.Beziers_mapString).(map[string]Type)
	case *BezierGrid:
		return any(stage.BezierGrids_mapString).(map[string]Type)
	case *BezierGridStack:
		return any(stage.BezierGridStacks_mapString).(map[string]Type)
	case *Chapter:
		return any(stage.Chapters_mapString).(map[string]Type)
	case *Circle:
		return any(stage.Circles_mapString).(map[string]Type)
	case *CircleGrid:
		return any(stage.CircleGrids_mapString).(map[string]Type)
	case *Content:
		return any(stage.Contents_mapString).(map[string]Type)
	case *ExportToMusicxml:
		return any(stage.ExportToMusicxmls_mapString).(map[string]Type)
	case *FrontCurve:
		return any(stage.FrontCurves_mapString).(map[string]Type)
	case *FrontCurveStack:
		return any(stage.FrontCurveStacks_mapString).(map[string]Type)
	case *HorizontalAxis:
		return any(stage.HorizontalAxiss_mapString).(map[string]Type)
	case *Key:
		return any(stage.Keys_mapString).(map[string]Type)
	case *Parameter:
		return any(stage.Parameters_mapString).(map[string]Type)
	case *Rhombus:
		return any(stage.Rhombuss_mapString).(map[string]Type)
	case *RhombusGrid:
		return any(stage.RhombusGrids_mapString).(map[string]Type)
	case *ShapeCategory:
		return any(stage.ShapeCategorys_mapString).(map[string]Type)
	case *SpiralBezier:
		return any(stage.SpiralBeziers_mapString).(map[string]Type)
	case *SpiralBezierGrid:
		return any(stage.SpiralBezierGrids_mapString).(map[string]Type)
	case *SpiralCircle:
		return any(stage.SpiralCircles_mapString).(map[string]Type)
	case *SpiralCircleGrid:
		return any(stage.SpiralCircleGrids_mapString).(map[string]Type)
	case *SpiralLine:
		return any(stage.SpiralLines_mapString).(map[string]Type)
	case *SpiralLineGrid:
		return any(stage.SpiralLineGrids_mapString).(map[string]Type)
	case *SpiralOrigin:
		return any(stage.SpiralOrigins_mapString).(map[string]Type)
	case *SpiralRhombus:
		return any(stage.SpiralRhombuss_mapString).(map[string]Type)
	case *SpiralRhombusGrid:
		return any(stage.SpiralRhombusGrids_mapString).(map[string]Type)
	case *VerticalAxis:
		return any(stage.VerticalAxiss_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Axis:
		return any(&stage.Axiss).(*map[*Type]struct{})
	case AxisGrid:
		return any(&stage.AxisGrids).(*map[*Type]struct{})
	case Bezier:
		return any(&stage.Beziers).(*map[*Type]struct{})
	case BezierGrid:
		return any(&stage.BezierGrids).(*map[*Type]struct{})
	case BezierGridStack:
		return any(&stage.BezierGridStacks).(*map[*Type]struct{})
	case Chapter:
		return any(&stage.Chapters).(*map[*Type]struct{})
	case Circle:
		return any(&stage.Circles).(*map[*Type]struct{})
	case CircleGrid:
		return any(&stage.CircleGrids).(*map[*Type]struct{})
	case Content:
		return any(&stage.Contents).(*map[*Type]struct{})
	case ExportToMusicxml:
		return any(&stage.ExportToMusicxmls).(*map[*Type]struct{})
	case FrontCurve:
		return any(&stage.FrontCurves).(*map[*Type]struct{})
	case FrontCurveStack:
		return any(&stage.FrontCurveStacks).(*map[*Type]struct{})
	case HorizontalAxis:
		return any(&stage.HorizontalAxiss).(*map[*Type]struct{})
	case Key:
		return any(&stage.Keys).(*map[*Type]struct{})
	case Parameter:
		return any(&stage.Parameters).(*map[*Type]struct{})
	case Rhombus:
		return any(&stage.Rhombuss).(*map[*Type]struct{})
	case RhombusGrid:
		return any(&stage.RhombusGrids).(*map[*Type]struct{})
	case ShapeCategory:
		return any(&stage.ShapeCategorys).(*map[*Type]struct{})
	case SpiralBezier:
		return any(&stage.SpiralBeziers).(*map[*Type]struct{})
	case SpiralBezierGrid:
		return any(&stage.SpiralBezierGrids).(*map[*Type]struct{})
	case SpiralCircle:
		return any(&stage.SpiralCircles).(*map[*Type]struct{})
	case SpiralCircleGrid:
		return any(&stage.SpiralCircleGrids).(*map[*Type]struct{})
	case SpiralLine:
		return any(&stage.SpiralLines).(*map[*Type]struct{})
	case SpiralLineGrid:
		return any(&stage.SpiralLineGrids).(*map[*Type]struct{})
	case SpiralOrigin:
		return any(&stage.SpiralOrigins).(*map[*Type]struct{})
	case SpiralRhombus:
		return any(&stage.SpiralRhombuss).(*map[*Type]struct{})
	case SpiralRhombusGrid:
		return any(&stage.SpiralRhombusGrids).(*map[*Type]struct{})
	case VerticalAxis:
		return any(&stage.VerticalAxiss).(*map[*Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Axis:
		return any(&stage.Axiss).(*map[Type]struct{})
	case *AxisGrid:
		return any(&stage.AxisGrids).(*map[Type]struct{})
	case *Bezier:
		return any(&stage.Beziers).(*map[Type]struct{})
	case *BezierGrid:
		return any(&stage.BezierGrids).(*map[Type]struct{})
	case *BezierGridStack:
		return any(&stage.BezierGridStacks).(*map[Type]struct{})
	case *Chapter:
		return any(&stage.Chapters).(*map[Type]struct{})
	case *Circle:
		return any(&stage.Circles).(*map[Type]struct{})
	case *CircleGrid:
		return any(&stage.CircleGrids).(*map[Type]struct{})
	case *Content:
		return any(&stage.Contents).(*map[Type]struct{})
	case *ExportToMusicxml:
		return any(&stage.ExportToMusicxmls).(*map[Type]struct{})
	case *FrontCurve:
		return any(&stage.FrontCurves).(*map[Type]struct{})
	case *FrontCurveStack:
		return any(&stage.FrontCurveStacks).(*map[Type]struct{})
	case *HorizontalAxis:
		return any(&stage.HorizontalAxiss).(*map[Type]struct{})
	case *Key:
		return any(&stage.Keys).(*map[Type]struct{})
	case *Parameter:
		return any(&stage.Parameters).(*map[Type]struct{})
	case *Rhombus:
		return any(&stage.Rhombuss).(*map[Type]struct{})
	case *RhombusGrid:
		return any(&stage.RhombusGrids).(*map[Type]struct{})
	case *ShapeCategory:
		return any(&stage.ShapeCategorys).(*map[Type]struct{})
	case *SpiralBezier:
		return any(&stage.SpiralBeziers).(*map[Type]struct{})
	case *SpiralBezierGrid:
		return any(&stage.SpiralBezierGrids).(*map[Type]struct{})
	case *SpiralCircle:
		return any(&stage.SpiralCircles).(*map[Type]struct{})
	case *SpiralCircleGrid:
		return any(&stage.SpiralCircleGrids).(*map[Type]struct{})
	case *SpiralLine:
		return any(&stage.SpiralLines).(*map[Type]struct{})
	case *SpiralLineGrid:
		return any(&stage.SpiralLineGrids).(*map[Type]struct{})
	case *SpiralOrigin:
		return any(&stage.SpiralOrigins).(*map[Type]struct{})
	case *SpiralRhombus:
		return any(&stage.SpiralRhombuss).(*map[Type]struct{})
	case *SpiralRhombusGrid:
		return any(&stage.SpiralRhombusGrids).(*map[Type]struct{})
	case *VerticalAxis:
		return any(&stage.VerticalAxiss).(*map[Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
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
	case Chapter:
		return any(&stage.Chapters_mapString).(*map[string]*Type)
	case Circle:
		return any(&stage.Circles_mapString).(*map[string]*Type)
	case CircleGrid:
		return any(&stage.CircleGrids_mapString).(*map[string]*Type)
	case Content:
		return any(&stage.Contents_mapString).(*map[string]*Type)
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
	case Chapter:
		return any(&Chapter{
			// Initialisation of associations
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
	case Content:
		return any(&Content{
			// Initialisation of associations
			// field is initialized with an instance of Chapter with the name of the field
			Chapters: []*Chapter{{Name: "Chapters"}},
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
	// reverse maps of direct associations of Chapter
	case Chapter:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {
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
			res := make(map[*Axis][]*AxisGrid)
			for axisgrid := range stage.AxisGrids {
				for _, axis_ := range axisgrid.Axiss {
					res[axis_] = append(res[axis_], axisgrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*Bezier][]*BezierGrid)
			for beziergrid := range stage.BezierGrids {
				for _, bezier_ := range beziergrid.Beziers {
					res[bezier_] = append(res[bezier_], beziergrid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BezierGridStack
	case BezierGridStack:
		switch fieldname {
		// insertion point for per direct association field
		case "BezierGrids":
			res := make(map[*BezierGrid][]*BezierGridStack)
			for beziergridstack := range stage.BezierGridStacks {
				for _, beziergrid_ := range beziergridstack.BezierGrids {
					res[beziergrid_] = append(res[beziergrid_], beziergridstack)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Chapter
	case Chapter:
		switch fieldname {
		// insertion point for per direct association field
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
			res := make(map[*Circle][]*CircleGrid)
			for circlegrid := range stage.CircleGrids {
				for _, circle_ := range circlegrid.Circles {
					res[circle_] = append(res[circle_], circlegrid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		case "Chapters":
			res := make(map[*Chapter][]*Content)
			for content := range stage.Contents {
				for _, chapter_ := range content.Chapters {
					res[chapter_] = append(res[chapter_], content)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*FrontCurve][]*FrontCurveStack)
			for frontcurvestack := range stage.FrontCurveStacks {
				for _, frontcurve_ := range frontcurvestack.FrontCurves {
					res[frontcurve_] = append(res[frontcurve_], frontcurvestack)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpiralCircles":
			res := make(map[*SpiralCircle][]*FrontCurveStack)
			for frontcurvestack := range stage.FrontCurveStacks {
				for _, spiralcircle_ := range frontcurvestack.SpiralCircles {
					res[spiralcircle_] = append(res[spiralcircle_], frontcurvestack)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*Rhombus][]*RhombusGrid)
			for rhombusgrid := range stage.RhombusGrids {
				for _, rhombus_ := range rhombusgrid.Rhombuses {
					res[rhombus_] = append(res[rhombus_], rhombusgrid)
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
		}
	// reverse maps of direct associations of SpiralBezierGrid
	case SpiralBezierGrid:
		switch fieldname {
		// insertion point for per direct association field
		case "SpiralBeziers":
			res := make(map[*SpiralBezier][]*SpiralBezierGrid)
			for spiralbeziergrid := range stage.SpiralBezierGrids {
				for _, spiralbezier_ := range spiralbeziergrid.SpiralBeziers {
					res[spiralbezier_] = append(res[spiralbezier_], spiralbeziergrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*SpiralCircle][]*SpiralCircleGrid)
			for spiralcirclegrid := range stage.SpiralCircleGrids {
				for _, spiralcircle_ := range spiralcirclegrid.SpiralCircles {
					res[spiralcircle_] = append(res[spiralcircle_], spiralcirclegrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*SpiralLine][]*SpiralLineGrid)
			for spirallinegrid := range stage.SpiralLineGrids {
				for _, spiralline_ := range spirallinegrid.SpiralLines {
					res[spiralline_] = append(res[spiralline_], spirallinegrid)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*SpiralRhombus][]*SpiralRhombusGrid)
			for spiralrhombusgrid := range stage.SpiralRhombusGrids {
				for _, spiralrhombus_ := range spiralrhombusgrid.SpiralRhombuses {
					res[spiralrhombus_] = append(res[spiralrhombus_], spiralrhombusgrid)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalAxis
	case VerticalAxis:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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
	case *Chapter:
		res = "Chapter"
	case *Circle:
		res = "Circle"
	case *CircleGrid:
		res = "CircleGrid"
	case *Content:
		res = "Content"
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *Axis:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AxisGrid"
		rf.Fieldname = "Axiss"
		res = append(res, rf)
	case *AxisGrid:
		var rf ReverseField
		_ = rf
	case *Bezier:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BezierGrid"
		rf.Fieldname = "Beziers"
		res = append(res, rf)
	case *BezierGrid:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BezierGridStack"
		rf.Fieldname = "BezierGrids"
		res = append(res, rf)
	case *BezierGridStack:
		var rf ReverseField
		_ = rf
	case *Chapter:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Content"
		rf.Fieldname = "Chapters"
		res = append(res, rf)
	case *Circle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "CircleGrid"
		rf.Fieldname = "Circles"
		res = append(res, rf)
	case *CircleGrid:
		var rf ReverseField
		_ = rf
	case *Content:
		var rf ReverseField
		_ = rf
	case *ExportToMusicxml:
		var rf ReverseField
		_ = rf
	case *FrontCurve:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FrontCurveStack"
		rf.Fieldname = "FrontCurves"
		res = append(res, rf)
	case *FrontCurveStack:
		var rf ReverseField
		_ = rf
	case *HorizontalAxis:
		var rf ReverseField
		_ = rf
	case *Key:
		var rf ReverseField
		_ = rf
	case *Parameter:
		var rf ReverseField
		_ = rf
	case *Rhombus:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "RhombusGrid"
		rf.Fieldname = "Rhombuses"
		res = append(res, rf)
	case *RhombusGrid:
		var rf ReverseField
		_ = rf
	case *ShapeCategory:
		var rf ReverseField
		_ = rf
	case *SpiralBezier:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SpiralBezierGrid"
		rf.Fieldname = "SpiralBeziers"
		res = append(res, rf)
	case *SpiralBezierGrid:
		var rf ReverseField
		_ = rf
	case *SpiralCircle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FrontCurveStack"
		rf.Fieldname = "SpiralCircles"
		res = append(res, rf)
		rf.GongstructName = "SpiralCircleGrid"
		rf.Fieldname = "SpiralCircles"
		res = append(res, rf)
	case *SpiralCircleGrid:
		var rf ReverseField
		_ = rf
	case *SpiralLine:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SpiralLineGrid"
		rf.Fieldname = "SpiralLines"
		res = append(res, rf)
	case *SpiralLineGrid:
		var rf ReverseField
		_ = rf
	case *SpiralOrigin:
		var rf ReverseField
		_ = rf
	case *SpiralRhombus:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SpiralRhombusGrid"
		rf.Fieldname = "SpiralRhombuses"
		res = append(res, rf)
	case *SpiralRhombusGrid:
		var rf ReverseField
		_ = rf
	case *VerticalAxis:
		var rf ReverseField
		_ = rf
	}
	return
}

// insertion point for get fields header method
func (axis *Axis) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "AngleDegree",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Length",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CenterX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CenterY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (axisgrid *AxisGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Reference",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Axis",
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "Axiss",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Axis",
		},
	}
	return
}

func (bezier *Bezier) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (beziergrid *BezierGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Reference",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bezier",
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "Beziers",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Bezier",
		},
	}
	return
}

func (beziergridstack *BezierGridStack) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "BezierGrids",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BezierGrid",
		},
	}
	return
}

func (chapter *Chapter) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (circle *Circle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "CenterX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CenterY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "HasBespokeRadius",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespopkeRadius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Pitch",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "ShowName",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BeatNb",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (circlegrid *CircleGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Reference",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Circle",
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "Circles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Circle",
		},
	}
	return
}

func (content *Content) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ContentPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OutputPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LayoutPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StaticPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Target",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Target",
		},
		{
			Name:                 "Chapters",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Chapter",
		},
	}
	return
}

func (exporttomusicxml *ExportToMusicxml) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Parameter",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Parameter",
		},
	}
	return
}

func (frontcurve *FrontCurve) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (frontcurvestack *FrontCurveStack) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "FrontCurves",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "FrontCurve",
		},
		{
			Name:                 "SpiralCircles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SpiralCircle",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (horizontalaxis *HorizontalAxis) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "AxisHandleBorderLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Axis_Length",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (key *Key) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (parameter *Parameter) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BackendColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MinuteColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "HourColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "N",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "M",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Z",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "InsideAngle",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ShiftToNearestCircle",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "SideLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "InitialRhombus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rhombus",
		},
		{
			Name:                 "InitialCircle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Circle",
		},
		{
			Name:                 "InitialRhombusGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RhombusGrid",
		},
		{
			Name:                 "InitialCircleGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "InitialAxis",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Axis",
		},
		{
			Name:                 "RotatedAxis",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Axis",
		},
		{
			Name:                 "RotatedRhombus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rhombus",
		},
		{
			Name:                 "RotatedRhombusGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RhombusGrid",
		},
		{
			Name:                 "RotatedCircleGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "NextRhombus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rhombus",
		},
		{
			Name:                 "NextCircle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Circle",
		},
		{
			Name:                 "GrowingRhombusGridSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rhombus",
		},
		{
			Name:                 "GrowingRhombusGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RhombusGrid",
		},
		{
			Name:                 "GrowingCircleGridSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Circle",
		},
		{
			Name:                 "GrowingCircleGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "GrowingCircleGridLeftSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Circle",
		},
		{
			Name:                 "GrowingCircleGridLeft",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "ConstructionAxis",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Axis",
		},
		{
			Name:                 "ConstructionAxisGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AxisGrid",
		},
		{
			Name:                 "ConstructionCircle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Circle",
		},
		{
			Name:                 "ConstructionCircleGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "GrowthCurveSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bezier",
		},
		{
			Name:                 "GrowthCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:                 "GrowthCurveShiftedRightSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bezier",
		},
		{
			Name:                 "GrowthCurveShiftedRight",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:                 "GrowthCurveNextSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bezier",
		},
		{
			Name:                 "GrowthCurveNext",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:                 "GrowthCurveNextShiftedRightSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bezier",
		},
		{
			Name:                 "GrowthCurveNextShiftedRight",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:                 "GrowthCurveStack",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGridStack",
		},
		{
			Name:               "StackWidth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NbShitRight",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "StackHeight",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "BezierControlLengthRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "SpiralRhombusGridSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralRhombus",
		},
		{
			Name:                 "SpiralRhombusGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralRhombusGrid",
		},
		{
			Name:                 "SpiralCircleSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralCircle",
		},
		{
			Name:                 "SpiralCircleGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralCircleGrid",
		},
		{
			Name:                 "SpiralCircleFullGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralCircleGrid",
		},
		{
			Name:                 "SpiralConstructionOuterLineSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralLine",
		},
		{
			Name:                 "SpiralConstructionInnerLineSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralLine",
		},
		{
			Name:                 "SpiralConstructionOuterLineGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralLineGrid",
		},
		{
			Name:                 "SpiralConstructionInnerLineGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralLineGrid",
		},
		{
			Name:                 "SpiralConstructionCircleGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralCircleGrid",
		},
		{
			Name:                 "SpiralConstructionOuterLineFullGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralLineGrid",
		},
		{
			Name:                 "SpiralBezierSeed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralBezier",
		},
		{
			Name:                 "SpiralBezierGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralBezierGrid",
		},
		{
			Name:                 "SpiralBezierFullGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralBezierGrid",
		},
		{
			Name:               "SpiralBezierStrength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "FrontCurveStack",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FrontCurveStack",
		},
		{
			Name:               "NbInterpolationPoints",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Fkey",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Key",
		},
		{
			Name:               "FkeySizeRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "FkeyOriginRelativeX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "FkeyOriginRelativeY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "PitchLines",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AxisGrid",
		},
		{
			Name:               "PitchHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "NbPitchLines",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "BeatLines",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AxisGrid",
		},
		{
			Name:               "BeatLinesHeightRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "NbBeatLines",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NbOfBeatsInTheme",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "FirstVoice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:                 "FirstVoiceShiftedRigth",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:               "FirstVoiceShiftX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "FirstVoiceShiftY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "SecondVoice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:                 "SecondVoiceShiftedRight",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BezierGrid",
		},
		{
			Name:               "PitchDifference",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "BeatsPerSecond",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Level",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "FirstVoiceNotes",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "FirstVoiceNotesShiftedRight",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "SecondVoiceNotes",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:                 "SecondVoiceNotesShiftedRight",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CircleGrid",
		},
		{
			Name:               "IsMinor",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ThemeBinaryEncoding",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "OriginX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "OriginY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "HorizontalAxis",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "HorizontalAxis",
		},
		{
			Name:                 "VerticalAxis",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "VerticalAxis",
		},
		{
			Name:                 "SpiralOrigin",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralOrigin",
		},
		{
			Name:               "SpiralOriginX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "SpiralOriginY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "OriginCrossWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "SpiralRadiusRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ShowSpiralBezierConstruct",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowInterpolationPoints",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ActualBeatsTemporalShift",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "PathToStaticFiles",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PathToGeneratedSVG",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PathToGeneratedScore",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (rhombus *Rhombus) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "CenterX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CenterY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "SideLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "AngleDegree",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "InsideAngle",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (rhombusgrid *RhombusGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Reference",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rhombus",
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "Rhombuses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Rhombus",
		},
	}
	return
}

func (shapecategory *ShapeCategory) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (spiralbezier *SpiralBezier) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointStartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ControlPointEndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "SpiralBeziers",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SpiralBezier",
		},
	}
	return
}

func (spiralcircle *SpiralCircle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "CenterX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "CenterY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "HasBespokeRadius",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespopkeRadius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Pitch",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "ShowName",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BeatNb",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "SpiralRhombusGrid",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SpiralRhombusGrid",
		},
		{
			Name:                 "SpiralCircles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SpiralCircle",
		},
	}
	return
}

func (spiralline *SpiralLine) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StartY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (spirallinegrid *SpiralLineGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "SpiralLines",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SpiralLine",
		},
	}
	return
}

func (spiralorigin *SpiralOrigin) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (spiralrhombus *SpiralRhombus) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "X_r0",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_r0",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "X_r1",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_r1",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "X_r2",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_r2",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "X_r3",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_r3",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:                 "SpiralRhombuses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SpiralRhombus",
		},
	}
	return
}

func (verticalaxis *VerticalAxis) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisplayed",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ShapeCategory",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ShapeCategory",
		},
		{
			Name:               "AxisHandleBorderLength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Axis_Length",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeIntDuration     GongFieldValueType = "GongFieldValueTypeIntDuration"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeDate            GongFieldValueType = "GongFieldValueTypeDate"
	GongFieldValueTypeBasicKind       GongFieldValueType = "GongFieldValueTypeBasicKind"
	GongFieldValueTypePointer         GongFieldValueType = "GongFieldValueTypePointer"
	GongFieldValueTypeSliceOfPointers GongFieldValueType = "GongFieldValueTypeSliceOfPointers"
)

type GongFieldValue struct {
	GongFieldValueType
	valueString string
	valueInt    int
	valueFloat  float64
	valueBool   bool

	// in case of a pointer, the ID of the pointed element
	// in case of a slice of pointers, the IDs, separated by semi columbs
	ids string
}

type GongFieldHeader struct {
	Name string
	GongFieldValueType
	TargetGongstructName string
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

// insertion point for generic get gongstruct field value
func (axis *Axis) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = axis.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", axis.IsDisplayed)
		res.valueBool = axis.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if axis.ShapeCategory != nil {
			res.valueString = axis.ShapeCategory.Name
			res.ids = axis.ShapeCategory.GongGetUUID(stage)
		}
	case "AngleDegree":
		res.valueString = fmt.Sprintf("%f", axis.AngleDegree)
		res.valueFloat = axis.AngleDegree
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Length":
		res.valueString = fmt.Sprintf("%f", axis.Length)
		res.valueFloat = axis.Length
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CenterX":
		res.valueString = fmt.Sprintf("%f", axis.CenterX)
		res.valueFloat = axis.CenterX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CenterY":
		res.valueString = fmt.Sprintf("%f", axis.CenterY)
		res.valueFloat = axis.CenterY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", axis.EndX)
		res.valueFloat = axis.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", axis.EndY)
		res.valueFloat = axis.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = axis.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", axis.FillOpacity)
		res.valueFloat = axis.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = axis.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", axis.StrokeOpacity)
		res.valueFloat = axis.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", axis.StrokeWidth)
		res.valueFloat = axis.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = axis.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = axis.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = axis.Transform
	}
	return
}

func (axisgrid *AxisGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = axisgrid.Name
	case "Reference":
		res.GongFieldValueType = GongFieldValueTypePointer
		if axisgrid.Reference != nil {
			res.valueString = axisgrid.Reference.Name
			res.ids = axisgrid.Reference.GongGetUUID(stage)
		}
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", axisgrid.IsDisplayed)
		res.valueBool = axisgrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if axisgrid.ShapeCategory != nil {
			res.valueString = axisgrid.ShapeCategory.Name
			res.ids = axisgrid.ShapeCategory.GongGetUUID(stage)
		}
	case "Axiss":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range axisgrid.Axiss {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (bezier *Bezier) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bezier.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", bezier.IsDisplayed)
		res.valueBool = bezier.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if bezier.ShapeCategory != nil {
			res.valueString = bezier.ShapeCategory.Name
			res.ids = bezier.ShapeCategory.GongGetUUID(stage)
		}
	case "StartX":
		res.valueString = fmt.Sprintf("%f", bezier.StartX)
		res.valueFloat = bezier.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", bezier.StartY)
		res.valueFloat = bezier.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartX":
		res.valueString = fmt.Sprintf("%f", bezier.ControlPointStartX)
		res.valueFloat = bezier.ControlPointStartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartY":
		res.valueString = fmt.Sprintf("%f", bezier.ControlPointStartY)
		res.valueFloat = bezier.ControlPointStartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", bezier.EndX)
		res.valueFloat = bezier.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", bezier.EndY)
		res.valueFloat = bezier.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndX":
		res.valueString = fmt.Sprintf("%f", bezier.ControlPointEndX)
		res.valueFloat = bezier.ControlPointEndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndY":
		res.valueString = fmt.Sprintf("%f", bezier.ControlPointEndY)
		res.valueFloat = bezier.ControlPointEndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = bezier.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", bezier.FillOpacity)
		res.valueFloat = bezier.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = bezier.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", bezier.StrokeOpacity)
		res.valueFloat = bezier.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", bezier.StrokeWidth)
		res.valueFloat = bezier.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = bezier.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = bezier.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = bezier.Transform
	}
	return
}

func (beziergrid *BezierGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = beziergrid.Name
	case "Reference":
		res.GongFieldValueType = GongFieldValueTypePointer
		if beziergrid.Reference != nil {
			res.valueString = beziergrid.Reference.Name
			res.ids = beziergrid.Reference.GongGetUUID(stage)
		}
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", beziergrid.IsDisplayed)
		res.valueBool = beziergrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if beziergrid.ShapeCategory != nil {
			res.valueString = beziergrid.ShapeCategory.Name
			res.ids = beziergrid.ShapeCategory.GongGetUUID(stage)
		}
	case "Beziers":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range beziergrid.Beziers {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (beziergridstack *BezierGridStack) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = beziergridstack.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", beziergridstack.IsDisplayed)
		res.valueBool = beziergridstack.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if beziergridstack.ShapeCategory != nil {
			res.valueString = beziergridstack.ShapeCategory.Name
			res.ids = beziergridstack.ShapeCategory.GongGetUUID(stage)
		}
	case "BezierGrids":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range beziergridstack.BezierGrids {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (chapter *Chapter) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = chapter.Name
	case "MardownContent":
		res.valueString = chapter.MardownContent
	}
	return
}

func (circle *Circle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = circle.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", circle.IsDisplayed)
		res.valueBool = circle.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circle.ShapeCategory != nil {
			res.valueString = circle.ShapeCategory.Name
			res.ids = circle.ShapeCategory.GongGetUUID(stage)
		}
	case "CenterX":
		res.valueString = fmt.Sprintf("%f", circle.CenterX)
		res.valueFloat = circle.CenterX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CenterY":
		res.valueString = fmt.Sprintf("%f", circle.CenterY)
		res.valueFloat = circle.CenterY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HasBespokeRadius":
		res.valueString = fmt.Sprintf("%t", circle.HasBespokeRadius)
		res.valueBool = circle.HasBespokeRadius
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespopkeRadius":
		res.valueString = fmt.Sprintf("%f", circle.BespopkeRadius)
		res.valueFloat = circle.BespopkeRadius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = circle.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", circle.FillOpacity)
		res.valueFloat = circle.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = circle.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", circle.StrokeOpacity)
		res.valueFloat = circle.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", circle.StrokeWidth)
		res.valueFloat = circle.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = circle.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = circle.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = circle.Transform
	case "Pitch":
		res.valueString = fmt.Sprintf("%d", circle.Pitch)
		res.valueInt = circle.Pitch
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ShowName":
		res.valueString = fmt.Sprintf("%t", circle.ShowName)
		res.valueBool = circle.ShowName
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BeatNb":
		res.valueString = fmt.Sprintf("%d", circle.BeatNb)
		res.valueInt = circle.BeatNb
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (circlegrid *CircleGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = circlegrid.Name
	case "Reference":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circlegrid.Reference != nil {
			res.valueString = circlegrid.Reference.Name
			res.ids = circlegrid.Reference.GongGetUUID(stage)
		}
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", circlegrid.IsDisplayed)
		res.valueBool = circlegrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circlegrid.ShapeCategory != nil {
			res.valueString = circlegrid.ShapeCategory.Name
			res.ids = circlegrid.ShapeCategory.GongGetUUID(stage)
		}
	case "Circles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range circlegrid.Circles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (content *Content) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = content.Name
	case "MardownContent":
		res.valueString = content.MardownContent
	case "ContentPath":
		res.valueString = content.ContentPath
	case "OutputPath":
		res.valueString = content.OutputPath
	case "LayoutPath":
		res.valueString = content.LayoutPath
	case "StaticPath":
		res.valueString = content.StaticPath
	case "Target":
		enum := content.Target
		res.valueString = enum.ToCodeString()
	case "Chapters":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range content.Chapters {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (exporttomusicxml *ExportToMusicxml) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = exporttomusicxml.Name
	case "Parameter":
		res.GongFieldValueType = GongFieldValueTypePointer
		if exporttomusicxml.Parameter != nil {
			res.valueString = exporttomusicxml.Parameter.Name
			res.ids = exporttomusicxml.Parameter.GongGetUUID(stage)
		}
	}
	return
}

func (frontcurve *FrontCurve) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = frontcurve.Name
	case "Path":
		res.valueString = frontcurve.Path
	}
	return
}

func (frontcurvestack *FrontCurveStack) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = frontcurvestack.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", frontcurvestack.IsDisplayed)
		res.valueBool = frontcurvestack.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if frontcurvestack.ShapeCategory != nil {
			res.valueString = frontcurvestack.ShapeCategory.Name
			res.ids = frontcurvestack.ShapeCategory.GongGetUUID(stage)
		}
	case "FrontCurves":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range frontcurvestack.FrontCurves {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SpiralCircles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range frontcurvestack.SpiralCircles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Color":
		res.valueString = frontcurvestack.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", frontcurvestack.FillOpacity)
		res.valueFloat = frontcurvestack.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = frontcurvestack.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", frontcurvestack.StrokeOpacity)
		res.valueFloat = frontcurvestack.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", frontcurvestack.StrokeWidth)
		res.valueFloat = frontcurvestack.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = frontcurvestack.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = frontcurvestack.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = frontcurvestack.Transform
	}
	return
}

func (horizontalaxis *HorizontalAxis) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = horizontalaxis.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", horizontalaxis.IsDisplayed)
		res.valueBool = horizontalaxis.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if horizontalaxis.ShapeCategory != nil {
			res.valueString = horizontalaxis.ShapeCategory.Name
			res.ids = horizontalaxis.ShapeCategory.GongGetUUID(stage)
		}
	case "AxisHandleBorderLength":
		res.valueString = fmt.Sprintf("%f", horizontalaxis.AxisHandleBorderLength)
		res.valueFloat = horizontalaxis.AxisHandleBorderLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Axis_Length":
		res.valueString = fmt.Sprintf("%f", horizontalaxis.Axis_Length)
		res.valueFloat = horizontalaxis.Axis_Length
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = horizontalaxis.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", horizontalaxis.FillOpacity)
		res.valueFloat = horizontalaxis.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = horizontalaxis.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", horizontalaxis.StrokeOpacity)
		res.valueFloat = horizontalaxis.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", horizontalaxis.StrokeWidth)
		res.valueFloat = horizontalaxis.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = horizontalaxis.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = horizontalaxis.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = horizontalaxis.Transform
	}
	return
}

func (key *Key) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = key.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", key.IsDisplayed)
		res.valueBool = key.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if key.ShapeCategory != nil {
			res.valueString = key.ShapeCategory.Name
			res.ids = key.ShapeCategory.GongGetUUID(stage)
		}
	case "Path":
		res.valueString = key.Path
	case "Color":
		res.valueString = key.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", key.FillOpacity)
		res.valueFloat = key.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = key.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", key.StrokeOpacity)
		res.valueFloat = key.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", key.StrokeWidth)
		res.valueFloat = key.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = key.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = key.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = key.Transform
	}
	return
}

func (parameter *Parameter) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = parameter.Name
	case "BackendColor":
		res.valueString = parameter.BackendColor
	case "MinuteColor":
		res.valueString = parameter.MinuteColor
	case "HourColor":
		res.valueString = parameter.HourColor
	case "N":
		res.valueString = fmt.Sprintf("%d", parameter.N)
		res.valueInt = parameter.N
		res.GongFieldValueType = GongFieldValueTypeInt
	case "M":
		res.valueString = fmt.Sprintf("%d", parameter.M)
		res.valueInt = parameter.M
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Z":
		res.valueString = fmt.Sprintf("%d", parameter.Z)
		res.valueInt = parameter.Z
		res.GongFieldValueType = GongFieldValueTypeInt
	case "InsideAngle":
		res.valueString = fmt.Sprintf("%f", parameter.InsideAngle)
		res.valueFloat = parameter.InsideAngle
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ShiftToNearestCircle":
		res.valueString = fmt.Sprintf("%d", parameter.ShiftToNearestCircle)
		res.valueInt = parameter.ShiftToNearestCircle
		res.GongFieldValueType = GongFieldValueTypeInt
	case "SideLength":
		res.valueString = fmt.Sprintf("%f", parameter.SideLength)
		res.valueFloat = parameter.SideLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "InitialRhombus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.InitialRhombus != nil {
			res.valueString = parameter.InitialRhombus.Name
			res.ids = parameter.InitialRhombus.GongGetUUID(stage)
		}
	case "InitialCircle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.InitialCircle != nil {
			res.valueString = parameter.InitialCircle.Name
			res.ids = parameter.InitialCircle.GongGetUUID(stage)
		}
	case "InitialRhombusGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.InitialRhombusGrid != nil {
			res.valueString = parameter.InitialRhombusGrid.Name
			res.ids = parameter.InitialRhombusGrid.GongGetUUID(stage)
		}
	case "InitialCircleGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.InitialCircleGrid != nil {
			res.valueString = parameter.InitialCircleGrid.Name
			res.ids = parameter.InitialCircleGrid.GongGetUUID(stage)
		}
	case "InitialAxis":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.InitialAxis != nil {
			res.valueString = parameter.InitialAxis.Name
			res.ids = parameter.InitialAxis.GongGetUUID(stage)
		}
	case "RotatedAxis":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.RotatedAxis != nil {
			res.valueString = parameter.RotatedAxis.Name
			res.ids = parameter.RotatedAxis.GongGetUUID(stage)
		}
	case "RotatedRhombus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.RotatedRhombus != nil {
			res.valueString = parameter.RotatedRhombus.Name
			res.ids = parameter.RotatedRhombus.GongGetUUID(stage)
		}
	case "RotatedRhombusGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.RotatedRhombusGrid != nil {
			res.valueString = parameter.RotatedRhombusGrid.Name
			res.ids = parameter.RotatedRhombusGrid.GongGetUUID(stage)
		}
	case "RotatedCircleGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.RotatedCircleGrid != nil {
			res.valueString = parameter.RotatedCircleGrid.Name
			res.ids = parameter.RotatedCircleGrid.GongGetUUID(stage)
		}
	case "NextRhombus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.NextRhombus != nil {
			res.valueString = parameter.NextRhombus.Name
			res.ids = parameter.NextRhombus.GongGetUUID(stage)
		}
	case "NextCircle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.NextCircle != nil {
			res.valueString = parameter.NextCircle.Name
			res.ids = parameter.NextCircle.GongGetUUID(stage)
		}
	case "GrowingRhombusGridSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowingRhombusGridSeed != nil {
			res.valueString = parameter.GrowingRhombusGridSeed.Name
			res.ids = parameter.GrowingRhombusGridSeed.GongGetUUID(stage)
		}
	case "GrowingRhombusGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowingRhombusGrid != nil {
			res.valueString = parameter.GrowingRhombusGrid.Name
			res.ids = parameter.GrowingRhombusGrid.GongGetUUID(stage)
		}
	case "GrowingCircleGridSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowingCircleGridSeed != nil {
			res.valueString = parameter.GrowingCircleGridSeed.Name
			res.ids = parameter.GrowingCircleGridSeed.GongGetUUID(stage)
		}
	case "GrowingCircleGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowingCircleGrid != nil {
			res.valueString = parameter.GrowingCircleGrid.Name
			res.ids = parameter.GrowingCircleGrid.GongGetUUID(stage)
		}
	case "GrowingCircleGridLeftSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowingCircleGridLeftSeed != nil {
			res.valueString = parameter.GrowingCircleGridLeftSeed.Name
			res.ids = parameter.GrowingCircleGridLeftSeed.GongGetUUID(stage)
		}
	case "GrowingCircleGridLeft":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowingCircleGridLeft != nil {
			res.valueString = parameter.GrowingCircleGridLeft.Name
			res.ids = parameter.GrowingCircleGridLeft.GongGetUUID(stage)
		}
	case "ConstructionAxis":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.ConstructionAxis != nil {
			res.valueString = parameter.ConstructionAxis.Name
			res.ids = parameter.ConstructionAxis.GongGetUUID(stage)
		}
	case "ConstructionAxisGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.ConstructionAxisGrid != nil {
			res.valueString = parameter.ConstructionAxisGrid.Name
			res.ids = parameter.ConstructionAxisGrid.GongGetUUID(stage)
		}
	case "ConstructionCircle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.ConstructionCircle != nil {
			res.valueString = parameter.ConstructionCircle.Name
			res.ids = parameter.ConstructionCircle.GongGetUUID(stage)
		}
	case "ConstructionCircleGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.ConstructionCircleGrid != nil {
			res.valueString = parameter.ConstructionCircleGrid.Name
			res.ids = parameter.ConstructionCircleGrid.GongGetUUID(stage)
		}
	case "GrowthCurveSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveSeed != nil {
			res.valueString = parameter.GrowthCurveSeed.Name
			res.ids = parameter.GrowthCurveSeed.GongGetUUID(stage)
		}
	case "GrowthCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurve != nil {
			res.valueString = parameter.GrowthCurve.Name
			res.ids = parameter.GrowthCurve.GongGetUUID(stage)
		}
	case "GrowthCurveShiftedRightSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveShiftedRightSeed != nil {
			res.valueString = parameter.GrowthCurveShiftedRightSeed.Name
			res.ids = parameter.GrowthCurveShiftedRightSeed.GongGetUUID(stage)
		}
	case "GrowthCurveShiftedRight":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveShiftedRight != nil {
			res.valueString = parameter.GrowthCurveShiftedRight.Name
			res.ids = parameter.GrowthCurveShiftedRight.GongGetUUID(stage)
		}
	case "GrowthCurveNextSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveNextSeed != nil {
			res.valueString = parameter.GrowthCurveNextSeed.Name
			res.ids = parameter.GrowthCurveNextSeed.GongGetUUID(stage)
		}
	case "GrowthCurveNext":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveNext != nil {
			res.valueString = parameter.GrowthCurveNext.Name
			res.ids = parameter.GrowthCurveNext.GongGetUUID(stage)
		}
	case "GrowthCurveNextShiftedRightSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveNextShiftedRightSeed != nil {
			res.valueString = parameter.GrowthCurveNextShiftedRightSeed.Name
			res.ids = parameter.GrowthCurveNextShiftedRightSeed.GongGetUUID(stage)
		}
	case "GrowthCurveNextShiftedRight":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveNextShiftedRight != nil {
			res.valueString = parameter.GrowthCurveNextShiftedRight.Name
			res.ids = parameter.GrowthCurveNextShiftedRight.GongGetUUID(stage)
		}
	case "GrowthCurveStack":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.GrowthCurveStack != nil {
			res.valueString = parameter.GrowthCurveStack.Name
			res.ids = parameter.GrowthCurveStack.GongGetUUID(stage)
		}
	case "StackWidth":
		res.valueString = fmt.Sprintf("%d", parameter.StackWidth)
		res.valueInt = parameter.StackWidth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NbShitRight":
		res.valueString = fmt.Sprintf("%d", parameter.NbShitRight)
		res.valueInt = parameter.NbShitRight
		res.GongFieldValueType = GongFieldValueTypeInt
	case "StackHeight":
		res.valueString = fmt.Sprintf("%d", parameter.StackHeight)
		res.valueInt = parameter.StackHeight
		res.GongFieldValueType = GongFieldValueTypeInt
	case "BezierControlLengthRatio":
		res.valueString = fmt.Sprintf("%f", parameter.BezierControlLengthRatio)
		res.valueFloat = parameter.BezierControlLengthRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SpiralRhombusGridSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralRhombusGridSeed != nil {
			res.valueString = parameter.SpiralRhombusGridSeed.Name
			res.ids = parameter.SpiralRhombusGridSeed.GongGetUUID(stage)
		}
	case "SpiralRhombusGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralRhombusGrid != nil {
			res.valueString = parameter.SpiralRhombusGrid.Name
			res.ids = parameter.SpiralRhombusGrid.GongGetUUID(stage)
		}
	case "SpiralCircleSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralCircleSeed != nil {
			res.valueString = parameter.SpiralCircleSeed.Name
			res.ids = parameter.SpiralCircleSeed.GongGetUUID(stage)
		}
	case "SpiralCircleGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralCircleGrid != nil {
			res.valueString = parameter.SpiralCircleGrid.Name
			res.ids = parameter.SpiralCircleGrid.GongGetUUID(stage)
		}
	case "SpiralCircleFullGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralCircleFullGrid != nil {
			res.valueString = parameter.SpiralCircleFullGrid.Name
			res.ids = parameter.SpiralCircleFullGrid.GongGetUUID(stage)
		}
	case "SpiralConstructionOuterLineSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralConstructionOuterLineSeed != nil {
			res.valueString = parameter.SpiralConstructionOuterLineSeed.Name
			res.ids = parameter.SpiralConstructionOuterLineSeed.GongGetUUID(stage)
		}
	case "SpiralConstructionInnerLineSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralConstructionInnerLineSeed != nil {
			res.valueString = parameter.SpiralConstructionInnerLineSeed.Name
			res.ids = parameter.SpiralConstructionInnerLineSeed.GongGetUUID(stage)
		}
	case "SpiralConstructionOuterLineGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralConstructionOuterLineGrid != nil {
			res.valueString = parameter.SpiralConstructionOuterLineGrid.Name
			res.ids = parameter.SpiralConstructionOuterLineGrid.GongGetUUID(stage)
		}
	case "SpiralConstructionInnerLineGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralConstructionInnerLineGrid != nil {
			res.valueString = parameter.SpiralConstructionInnerLineGrid.Name
			res.ids = parameter.SpiralConstructionInnerLineGrid.GongGetUUID(stage)
		}
	case "SpiralConstructionCircleGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralConstructionCircleGrid != nil {
			res.valueString = parameter.SpiralConstructionCircleGrid.Name
			res.ids = parameter.SpiralConstructionCircleGrid.GongGetUUID(stage)
		}
	case "SpiralConstructionOuterLineFullGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralConstructionOuterLineFullGrid != nil {
			res.valueString = parameter.SpiralConstructionOuterLineFullGrid.Name
			res.ids = parameter.SpiralConstructionOuterLineFullGrid.GongGetUUID(stage)
		}
	case "SpiralBezierSeed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralBezierSeed != nil {
			res.valueString = parameter.SpiralBezierSeed.Name
			res.ids = parameter.SpiralBezierSeed.GongGetUUID(stage)
		}
	case "SpiralBezierGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralBezierGrid != nil {
			res.valueString = parameter.SpiralBezierGrid.Name
			res.ids = parameter.SpiralBezierGrid.GongGetUUID(stage)
		}
	case "SpiralBezierFullGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralBezierFullGrid != nil {
			res.valueString = parameter.SpiralBezierFullGrid.Name
			res.ids = parameter.SpiralBezierFullGrid.GongGetUUID(stage)
		}
	case "SpiralBezierStrength":
		res.valueString = fmt.Sprintf("%f", parameter.SpiralBezierStrength)
		res.valueFloat = parameter.SpiralBezierStrength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FrontCurveStack":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.FrontCurveStack != nil {
			res.valueString = parameter.FrontCurveStack.Name
			res.ids = parameter.FrontCurveStack.GongGetUUID(stage)
		}
	case "NbInterpolationPoints":
		res.valueString = fmt.Sprintf("%d", parameter.NbInterpolationPoints)
		res.valueInt = parameter.NbInterpolationPoints
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Fkey":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.Fkey != nil {
			res.valueString = parameter.Fkey.Name
			res.ids = parameter.Fkey.GongGetUUID(stage)
		}
	case "FkeySizeRatio":
		res.valueString = fmt.Sprintf("%f", parameter.FkeySizeRatio)
		res.valueFloat = parameter.FkeySizeRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FkeyOriginRelativeX":
		res.valueString = fmt.Sprintf("%f", parameter.FkeyOriginRelativeX)
		res.valueFloat = parameter.FkeyOriginRelativeX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FkeyOriginRelativeY":
		res.valueString = fmt.Sprintf("%f", parameter.FkeyOriginRelativeY)
		res.valueFloat = parameter.FkeyOriginRelativeY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "PitchLines":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.PitchLines != nil {
			res.valueString = parameter.PitchLines.Name
			res.ids = parameter.PitchLines.GongGetUUID(stage)
		}
	case "PitchHeight":
		res.valueString = fmt.Sprintf("%f", parameter.PitchHeight)
		res.valueFloat = parameter.PitchHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "NbPitchLines":
		res.valueString = fmt.Sprintf("%d", parameter.NbPitchLines)
		res.valueInt = parameter.NbPitchLines
		res.GongFieldValueType = GongFieldValueTypeInt
	case "BeatLines":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.BeatLines != nil {
			res.valueString = parameter.BeatLines.Name
			res.ids = parameter.BeatLines.GongGetUUID(stage)
		}
	case "BeatLinesHeightRatio":
		res.valueString = fmt.Sprintf("%f", parameter.BeatLinesHeightRatio)
		res.valueFloat = parameter.BeatLinesHeightRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "NbBeatLines":
		res.valueString = fmt.Sprintf("%d", parameter.NbBeatLines)
		res.valueInt = parameter.NbBeatLines
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NbOfBeatsInTheme":
		res.valueString = fmt.Sprintf("%d", parameter.NbOfBeatsInTheme)
		res.valueInt = parameter.NbOfBeatsInTheme
		res.GongFieldValueType = GongFieldValueTypeInt
	case "FirstVoice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.FirstVoice != nil {
			res.valueString = parameter.FirstVoice.Name
			res.ids = parameter.FirstVoice.GongGetUUID(stage)
		}
	case "FirstVoiceShiftedRigth":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.FirstVoiceShiftedRigth != nil {
			res.valueString = parameter.FirstVoiceShiftedRigth.Name
			res.ids = parameter.FirstVoiceShiftedRigth.GongGetUUID(stage)
		}
	case "FirstVoiceShiftX":
		res.valueString = fmt.Sprintf("%f", parameter.FirstVoiceShiftX)
		res.valueFloat = parameter.FirstVoiceShiftX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FirstVoiceShiftY":
		res.valueString = fmt.Sprintf("%f", parameter.FirstVoiceShiftY)
		res.valueFloat = parameter.FirstVoiceShiftY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SecondVoice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SecondVoice != nil {
			res.valueString = parameter.SecondVoice.Name
			res.ids = parameter.SecondVoice.GongGetUUID(stage)
		}
	case "SecondVoiceShiftedRight":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SecondVoiceShiftedRight != nil {
			res.valueString = parameter.SecondVoiceShiftedRight.Name
			res.ids = parameter.SecondVoiceShiftedRight.GongGetUUID(stage)
		}
	case "PitchDifference":
		res.valueString = fmt.Sprintf("%d", parameter.PitchDifference)
		res.valueInt = parameter.PitchDifference
		res.GongFieldValueType = GongFieldValueTypeInt
	case "BeatsPerSecond":
		res.valueString = fmt.Sprintf("%f", parameter.BeatsPerSecond)
		res.valueFloat = parameter.BeatsPerSecond
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Level":
		res.valueString = fmt.Sprintf("%f", parameter.Level)
		res.valueFloat = parameter.Level
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FirstVoiceNotes":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.FirstVoiceNotes != nil {
			res.valueString = parameter.FirstVoiceNotes.Name
			res.ids = parameter.FirstVoiceNotes.GongGetUUID(stage)
		}
	case "FirstVoiceNotesShiftedRight":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.FirstVoiceNotesShiftedRight != nil {
			res.valueString = parameter.FirstVoiceNotesShiftedRight.Name
			res.ids = parameter.FirstVoiceNotesShiftedRight.GongGetUUID(stage)
		}
	case "SecondVoiceNotes":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SecondVoiceNotes != nil {
			res.valueString = parameter.SecondVoiceNotes.Name
			res.ids = parameter.SecondVoiceNotes.GongGetUUID(stage)
		}
	case "SecondVoiceNotesShiftedRight":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SecondVoiceNotesShiftedRight != nil {
			res.valueString = parameter.SecondVoiceNotesShiftedRight.Name
			res.ids = parameter.SecondVoiceNotesShiftedRight.GongGetUUID(stage)
		}
	case "IsMinor":
		res.valueString = fmt.Sprintf("%t", parameter.IsMinor)
		res.valueBool = parameter.IsMinor
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ThemeBinaryEncoding":
		res.valueString = fmt.Sprintf("%d", parameter.ThemeBinaryEncoding)
		res.valueInt = parameter.ThemeBinaryEncoding
		res.GongFieldValueType = GongFieldValueTypeInt
	case "OriginX":
		res.valueString = fmt.Sprintf("%f", parameter.OriginX)
		res.valueFloat = parameter.OriginX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OriginY":
		res.valueString = fmt.Sprintf("%f", parameter.OriginY)
		res.valueFloat = parameter.OriginY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HorizontalAxis":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.HorizontalAxis != nil {
			res.valueString = parameter.HorizontalAxis.Name
			res.ids = parameter.HorizontalAxis.GongGetUUID(stage)
		}
	case "VerticalAxis":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.VerticalAxis != nil {
			res.valueString = parameter.VerticalAxis.Name
			res.ids = parameter.VerticalAxis.GongGetUUID(stage)
		}
	case "SpiralOrigin":
		res.GongFieldValueType = GongFieldValueTypePointer
		if parameter.SpiralOrigin != nil {
			res.valueString = parameter.SpiralOrigin.Name
			res.ids = parameter.SpiralOrigin.GongGetUUID(stage)
		}
	case "SpiralOriginX":
		res.valueString = fmt.Sprintf("%f", parameter.SpiralOriginX)
		res.valueFloat = parameter.SpiralOriginX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SpiralOriginY":
		res.valueString = fmt.Sprintf("%f", parameter.SpiralOriginY)
		res.valueFloat = parameter.SpiralOriginY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OriginCrossWidth":
		res.valueString = fmt.Sprintf("%f", parameter.OriginCrossWidth)
		res.valueFloat = parameter.OriginCrossWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SpiralRadiusRatio":
		res.valueString = fmt.Sprintf("%f", parameter.SpiralRadiusRatio)
		res.valueFloat = parameter.SpiralRadiusRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ShowSpiralBezierConstruct":
		res.valueString = fmt.Sprintf("%t", parameter.ShowSpiralBezierConstruct)
		res.valueBool = parameter.ShowSpiralBezierConstruct
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowInterpolationPoints":
		res.valueString = fmt.Sprintf("%t", parameter.ShowInterpolationPoints)
		res.valueBool = parameter.ShowInterpolationPoints
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ActualBeatsTemporalShift":
		res.valueString = fmt.Sprintf("%d", parameter.ActualBeatsTemporalShift)
		res.valueInt = parameter.ActualBeatsTemporalShift
		res.GongFieldValueType = GongFieldValueTypeInt
	case "PathToStaticFiles":
		res.valueString = parameter.PathToStaticFiles
	case "PathToGeneratedSVG":
		res.valueString = parameter.PathToGeneratedSVG
	case "PathToGeneratedScore":
		res.valueString = parameter.PathToGeneratedScore
	}
	return
}

func (rhombus *Rhombus) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rhombus.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", rhombus.IsDisplayed)
		res.valueBool = rhombus.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rhombus.ShapeCategory != nil {
			res.valueString = rhombus.ShapeCategory.Name
			res.ids = rhombus.ShapeCategory.GongGetUUID(stage)
		}
	case "CenterX":
		res.valueString = fmt.Sprintf("%f", rhombus.CenterX)
		res.valueFloat = rhombus.CenterX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CenterY":
		res.valueString = fmt.Sprintf("%f", rhombus.CenterY)
		res.valueFloat = rhombus.CenterY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SideLength":
		res.valueString = fmt.Sprintf("%f", rhombus.SideLength)
		res.valueFloat = rhombus.SideLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "AngleDegree":
		res.valueString = fmt.Sprintf("%f", rhombus.AngleDegree)
		res.valueFloat = rhombus.AngleDegree
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "InsideAngle":
		res.valueString = fmt.Sprintf("%f", rhombus.InsideAngle)
		res.valueFloat = rhombus.InsideAngle
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = rhombus.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", rhombus.FillOpacity)
		res.valueFloat = rhombus.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = rhombus.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", rhombus.StrokeOpacity)
		res.valueFloat = rhombus.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", rhombus.StrokeWidth)
		res.valueFloat = rhombus.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = rhombus.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = rhombus.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = rhombus.Transform
	}
	return
}

func (rhombusgrid *RhombusGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rhombusgrid.Name
	case "Reference":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rhombusgrid.Reference != nil {
			res.valueString = rhombusgrid.Reference.Name
			res.ids = rhombusgrid.Reference.GongGetUUID(stage)
		}
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", rhombusgrid.IsDisplayed)
		res.valueBool = rhombusgrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rhombusgrid.ShapeCategory != nil {
			res.valueString = rhombusgrid.ShapeCategory.Name
			res.ids = rhombusgrid.ShapeCategory.GongGetUUID(stage)
		}
	case "Rhombuses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range rhombusgrid.Rhombuses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (shapecategory *ShapeCategory) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = shapecategory.Name
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", shapecategory.IsExpanded)
		res.valueBool = shapecategory.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (spiralbezier *SpiralBezier) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralbezier.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralbezier.IsDisplayed)
		res.valueBool = spiralbezier.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralbezier.ShapeCategory != nil {
			res.valueString = spiralbezier.ShapeCategory.Name
			res.ids = spiralbezier.ShapeCategory.GongGetUUID(stage)
		}
	case "StartX":
		res.valueString = fmt.Sprintf("%f", spiralbezier.StartX)
		res.valueFloat = spiralbezier.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", spiralbezier.StartY)
		res.valueFloat = spiralbezier.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartX":
		res.valueString = fmt.Sprintf("%f", spiralbezier.ControlPointStartX)
		res.valueFloat = spiralbezier.ControlPointStartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointStartY":
		res.valueString = fmt.Sprintf("%f", spiralbezier.ControlPointStartY)
		res.valueFloat = spiralbezier.ControlPointStartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", spiralbezier.EndX)
		res.valueFloat = spiralbezier.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", spiralbezier.EndY)
		res.valueFloat = spiralbezier.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndX":
		res.valueString = fmt.Sprintf("%f", spiralbezier.ControlPointEndX)
		res.valueFloat = spiralbezier.ControlPointEndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ControlPointEndY":
		res.valueString = fmt.Sprintf("%f", spiralbezier.ControlPointEndY)
		res.valueFloat = spiralbezier.ControlPointEndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = spiralbezier.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", spiralbezier.FillOpacity)
		res.valueFloat = spiralbezier.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = spiralbezier.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", spiralbezier.StrokeOpacity)
		res.valueFloat = spiralbezier.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", spiralbezier.StrokeWidth)
		res.valueFloat = spiralbezier.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = spiralbezier.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = spiralbezier.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = spiralbezier.Transform
	}
	return
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralbeziergrid.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralbeziergrid.IsDisplayed)
		res.valueBool = spiralbeziergrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralbeziergrid.ShapeCategory != nil {
			res.valueString = spiralbeziergrid.ShapeCategory.Name
			res.ids = spiralbeziergrid.ShapeCategory.GongGetUUID(stage)
		}
	case "SpiralBeziers":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range spiralbeziergrid.SpiralBeziers {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (spiralcircle *SpiralCircle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralcircle.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralcircle.IsDisplayed)
		res.valueBool = spiralcircle.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralcircle.ShapeCategory != nil {
			res.valueString = spiralcircle.ShapeCategory.Name
			res.ids = spiralcircle.ShapeCategory.GongGetUUID(stage)
		}
	case "CenterX":
		res.valueString = fmt.Sprintf("%f", spiralcircle.CenterX)
		res.valueFloat = spiralcircle.CenterX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CenterY":
		res.valueString = fmt.Sprintf("%f", spiralcircle.CenterY)
		res.valueFloat = spiralcircle.CenterY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "HasBespokeRadius":
		res.valueString = fmt.Sprintf("%t", spiralcircle.HasBespokeRadius)
		res.valueBool = spiralcircle.HasBespokeRadius
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespopkeRadius":
		res.valueString = fmt.Sprintf("%f", spiralcircle.BespopkeRadius)
		res.valueFloat = spiralcircle.BespopkeRadius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = spiralcircle.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", spiralcircle.FillOpacity)
		res.valueFloat = spiralcircle.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = spiralcircle.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", spiralcircle.StrokeOpacity)
		res.valueFloat = spiralcircle.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", spiralcircle.StrokeWidth)
		res.valueFloat = spiralcircle.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = spiralcircle.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = spiralcircle.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = spiralcircle.Transform
	case "Pitch":
		res.valueString = fmt.Sprintf("%d", spiralcircle.Pitch)
		res.valueInt = spiralcircle.Pitch
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ShowName":
		res.valueString = fmt.Sprintf("%t", spiralcircle.ShowName)
		res.valueBool = spiralcircle.ShowName
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BeatNb":
		res.valueString = fmt.Sprintf("%d", spiralcircle.BeatNb)
		res.valueInt = spiralcircle.BeatNb
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Path":
		res.valueString = spiralcircle.Path
	}
	return
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralcirclegrid.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralcirclegrid.IsDisplayed)
		res.valueBool = spiralcirclegrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralcirclegrid.ShapeCategory != nil {
			res.valueString = spiralcirclegrid.ShapeCategory.Name
			res.ids = spiralcirclegrid.ShapeCategory.GongGetUUID(stage)
		}
	case "SpiralRhombusGrid":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralcirclegrid.SpiralRhombusGrid != nil {
			res.valueString = spiralcirclegrid.SpiralRhombusGrid.Name
			res.ids = spiralcirclegrid.SpiralRhombusGrid.GongGetUUID(stage)
		}
	case "SpiralCircles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range spiralcirclegrid.SpiralCircles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (spiralline *SpiralLine) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralline.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralline.IsDisplayed)
		res.valueBool = spiralline.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralline.ShapeCategory != nil {
			res.valueString = spiralline.ShapeCategory.Name
			res.ids = spiralline.ShapeCategory.GongGetUUID(stage)
		}
	case "StartX":
		res.valueString = fmt.Sprintf("%f", spiralline.StartX)
		res.valueFloat = spiralline.StartX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndX":
		res.valueString = fmt.Sprintf("%f", spiralline.EndX)
		res.valueFloat = spiralline.EndX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartY":
		res.valueString = fmt.Sprintf("%f", spiralline.StartY)
		res.valueFloat = spiralline.StartY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndY":
		res.valueString = fmt.Sprintf("%f", spiralline.EndY)
		res.valueFloat = spiralline.EndY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = spiralline.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", spiralline.FillOpacity)
		res.valueFloat = spiralline.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = spiralline.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", spiralline.StrokeOpacity)
		res.valueFloat = spiralline.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", spiralline.StrokeWidth)
		res.valueFloat = spiralline.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = spiralline.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = spiralline.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = spiralline.Transform
	}
	return
}

func (spirallinegrid *SpiralLineGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spirallinegrid.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spirallinegrid.IsDisplayed)
		res.valueBool = spirallinegrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spirallinegrid.ShapeCategory != nil {
			res.valueString = spirallinegrid.ShapeCategory.Name
			res.ids = spirallinegrid.ShapeCategory.GongGetUUID(stage)
		}
	case "SpiralLines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range spirallinegrid.SpiralLines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (spiralorigin *SpiralOrigin) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralorigin.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralorigin.IsDisplayed)
		res.valueBool = spiralorigin.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralorigin.ShapeCategory != nil {
			res.valueString = spiralorigin.ShapeCategory.Name
			res.ids = spiralorigin.ShapeCategory.GongGetUUID(stage)
		}
	case "Color":
		res.valueString = spiralorigin.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", spiralorigin.FillOpacity)
		res.valueFloat = spiralorigin.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = spiralorigin.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", spiralorigin.StrokeOpacity)
		res.valueFloat = spiralorigin.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", spiralorigin.StrokeWidth)
		res.valueFloat = spiralorigin.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = spiralorigin.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = spiralorigin.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = spiralorigin.Transform
	}
	return
}

func (spiralrhombus *SpiralRhombus) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralrhombus.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralrhombus.IsDisplayed)
		res.valueBool = spiralrhombus.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralrhombus.ShapeCategory != nil {
			res.valueString = spiralrhombus.ShapeCategory.Name
			res.ids = spiralrhombus.ShapeCategory.GongGetUUID(stage)
		}
	case "X_r0":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.X_r0)
		res.valueFloat = spiralrhombus.X_r0
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_r0":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.Y_r0)
		res.valueFloat = spiralrhombus.Y_r0
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "X_r1":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.X_r1)
		res.valueFloat = spiralrhombus.X_r1
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_r1":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.Y_r1)
		res.valueFloat = spiralrhombus.Y_r1
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "X_r2":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.X_r2)
		res.valueFloat = spiralrhombus.X_r2
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_r2":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.Y_r2)
		res.valueFloat = spiralrhombus.Y_r2
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "X_r3":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.X_r3)
		res.valueFloat = spiralrhombus.X_r3
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_r3":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.Y_r3)
		res.valueFloat = spiralrhombus.Y_r3
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = spiralrhombus.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.FillOpacity)
		res.valueFloat = spiralrhombus.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = spiralrhombus.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.StrokeOpacity)
		res.valueFloat = spiralrhombus.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", spiralrhombus.StrokeWidth)
		res.valueFloat = spiralrhombus.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = spiralrhombus.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = spiralrhombus.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = spiralrhombus.Transform
	}
	return
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = spiralrhombusgrid.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", spiralrhombusgrid.IsDisplayed)
		res.valueBool = spiralrhombusgrid.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if spiralrhombusgrid.ShapeCategory != nil {
			res.valueString = spiralrhombusgrid.ShapeCategory.Name
			res.ids = spiralrhombusgrid.ShapeCategory.GongGetUUID(stage)
		}
	case "SpiralRhombuses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range spiralrhombusgrid.SpiralRhombuses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (verticalaxis *VerticalAxis) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalaxis.Name
	case "IsDisplayed":
		res.valueString = fmt.Sprintf("%t", verticalaxis.IsDisplayed)
		res.valueBool = verticalaxis.IsDisplayed
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShapeCategory":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalaxis.ShapeCategory != nil {
			res.valueString = verticalaxis.ShapeCategory.Name
			res.ids = verticalaxis.ShapeCategory.GongGetUUID(stage)
		}
	case "AxisHandleBorderLength":
		res.valueString = fmt.Sprintf("%f", verticalaxis.AxisHandleBorderLength)
		res.valueFloat = verticalaxis.AxisHandleBorderLength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Axis_Length":
		res.valueString = fmt.Sprintf("%f", verticalaxis.Axis_Length)
		res.valueFloat = verticalaxis.Axis_Length
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = verticalaxis.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", verticalaxis.FillOpacity)
		res.valueFloat = verticalaxis.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = verticalaxis.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", verticalaxis.StrokeOpacity)
		res.valueFloat = verticalaxis.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", verticalaxis.StrokeWidth)
		res.valueFloat = verticalaxis.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = verticalaxis.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = verticalaxis.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = verticalaxis.Transform
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (axis *Axis) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		axis.Name = value.GetValueString()
	case "IsDisplayed":
		axis.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			axis.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					axis.ShapeCategory = __instance__
					break
				}
			}
		}
	case "AngleDegree":
		axis.AngleDegree = value.GetValueFloat()
	case "Length":
		axis.Length = value.GetValueFloat()
	case "CenterX":
		axis.CenterX = value.GetValueFloat()
	case "CenterY":
		axis.CenterY = value.GetValueFloat()
	case "EndX":
		axis.EndX = value.GetValueFloat()
	case "EndY":
		axis.EndY = value.GetValueFloat()
	case "Color":
		axis.Color = value.GetValueString()
	case "FillOpacity":
		axis.FillOpacity = value.GetValueFloat()
	case "Stroke":
		axis.Stroke = value.GetValueString()
	case "StrokeOpacity":
		axis.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		axis.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		axis.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		axis.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		axis.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (axisgrid *AxisGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		axisgrid.Name = value.GetValueString()
	case "Reference":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			axisgrid.Reference = nil
			for __instance__ := range stage.Axiss {
				if stage.Axis_stagedOrder[__instance__] == uint(id) {
					axisgrid.Reference = __instance__
					break
				}
			}
		}
	case "IsDisplayed":
		axisgrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			axisgrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					axisgrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "Axiss":
		axisgrid.Axiss = make([]*Axis, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Axiss {
					if stage.Axis_stagedOrder[__instance__] == uint(id) {
						axisgrid.Axiss = append(axisgrid.Axiss, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (bezier *Bezier) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bezier.Name = value.GetValueString()
	case "IsDisplayed":
		bezier.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			bezier.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					bezier.ShapeCategory = __instance__
					break
				}
			}
		}
	case "StartX":
		bezier.StartX = value.GetValueFloat()
	case "StartY":
		bezier.StartY = value.GetValueFloat()
	case "ControlPointStartX":
		bezier.ControlPointStartX = value.GetValueFloat()
	case "ControlPointStartY":
		bezier.ControlPointStartY = value.GetValueFloat()
	case "EndX":
		bezier.EndX = value.GetValueFloat()
	case "EndY":
		bezier.EndY = value.GetValueFloat()
	case "ControlPointEndX":
		bezier.ControlPointEndX = value.GetValueFloat()
	case "ControlPointEndY":
		bezier.ControlPointEndY = value.GetValueFloat()
	case "Color":
		bezier.Color = value.GetValueString()
	case "FillOpacity":
		bezier.FillOpacity = value.GetValueFloat()
	case "Stroke":
		bezier.Stroke = value.GetValueString()
	case "StrokeOpacity":
		bezier.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		bezier.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		bezier.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		bezier.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		bezier.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (beziergrid *BezierGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		beziergrid.Name = value.GetValueString()
	case "Reference":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			beziergrid.Reference = nil
			for __instance__ := range stage.Beziers {
				if stage.Bezier_stagedOrder[__instance__] == uint(id) {
					beziergrid.Reference = __instance__
					break
				}
			}
		}
	case "IsDisplayed":
		beziergrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			beziergrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					beziergrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "Beziers":
		beziergrid.Beziers = make([]*Bezier, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Beziers {
					if stage.Bezier_stagedOrder[__instance__] == uint(id) {
						beziergrid.Beziers = append(beziergrid.Beziers, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (beziergridstack *BezierGridStack) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		beziergridstack.Name = value.GetValueString()
	case "IsDisplayed":
		beziergridstack.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			beziergridstack.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					beziergridstack.ShapeCategory = __instance__
					break
				}
			}
		}
	case "BezierGrids":
		beziergridstack.BezierGrids = make([]*BezierGrid, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.BezierGrids {
					if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
						beziergridstack.BezierGrids = append(beziergridstack.BezierGrids, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (chapter *Chapter) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		chapter.Name = value.GetValueString()
	case "MardownContent":
		chapter.MardownContent = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (circle *Circle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		circle.Name = value.GetValueString()
	case "IsDisplayed":
		circle.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			circle.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					circle.ShapeCategory = __instance__
					break
				}
			}
		}
	case "CenterX":
		circle.CenterX = value.GetValueFloat()
	case "CenterY":
		circle.CenterY = value.GetValueFloat()
	case "HasBespokeRadius":
		circle.HasBespokeRadius = value.GetValueBool()
	case "BespopkeRadius":
		circle.BespopkeRadius = value.GetValueFloat()
	case "Color":
		circle.Color = value.GetValueString()
	case "FillOpacity":
		circle.FillOpacity = value.GetValueFloat()
	case "Stroke":
		circle.Stroke = value.GetValueString()
	case "StrokeOpacity":
		circle.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		circle.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		circle.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		circle.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		circle.Transform = value.GetValueString()
	case "Pitch":
		circle.Pitch = int(value.GetValueInt())
	case "ShowName":
		circle.ShowName = value.GetValueBool()
	case "BeatNb":
		circle.BeatNb = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (circlegrid *CircleGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		circlegrid.Name = value.GetValueString()
	case "Reference":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			circlegrid.Reference = nil
			for __instance__ := range stage.Circles {
				if stage.Circle_stagedOrder[__instance__] == uint(id) {
					circlegrid.Reference = __instance__
					break
				}
			}
		}
	case "IsDisplayed":
		circlegrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			circlegrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					circlegrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "Circles":
		circlegrid.Circles = make([]*Circle, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Circles {
					if stage.Circle_stagedOrder[__instance__] == uint(id) {
						circlegrid.Circles = append(circlegrid.Circles, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (content *Content) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		content.Name = value.GetValueString()
	case "MardownContent":
		content.MardownContent = value.GetValueString()
	case "ContentPath":
		content.ContentPath = value.GetValueString()
	case "OutputPath":
		content.OutputPath = value.GetValueString()
	case "LayoutPath":
		content.LayoutPath = value.GetValueString()
	case "StaticPath":
		content.StaticPath = value.GetValueString()
	case "Target":
		content.Target.FromCodeString(value.GetValueString())
	case "Chapters":
		content.Chapters = make([]*Chapter, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Chapters {
					if stage.Chapter_stagedOrder[__instance__] == uint(id) {
						content.Chapters = append(content.Chapters, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (exporttomusicxml *ExportToMusicxml) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		exporttomusicxml.Name = value.GetValueString()
	case "Parameter":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			exporttomusicxml.Parameter = nil
			for __instance__ := range stage.Parameters {
				if stage.Parameter_stagedOrder[__instance__] == uint(id) {
					exporttomusicxml.Parameter = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (frontcurve *FrontCurve) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		frontcurve.Name = value.GetValueString()
	case "Path":
		frontcurve.Path = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (frontcurvestack *FrontCurveStack) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		frontcurvestack.Name = value.GetValueString()
	case "IsDisplayed":
		frontcurvestack.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			frontcurvestack.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					frontcurvestack.ShapeCategory = __instance__
					break
				}
			}
		}
	case "FrontCurves":
		frontcurvestack.FrontCurves = make([]*FrontCurve, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.FrontCurves {
					if stage.FrontCurve_stagedOrder[__instance__] == uint(id) {
						frontcurvestack.FrontCurves = append(frontcurvestack.FrontCurves, __instance__)
						break
					}
				}
			}
		}
	case "SpiralCircles":
		frontcurvestack.SpiralCircles = make([]*SpiralCircle, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SpiralCircles {
					if stage.SpiralCircle_stagedOrder[__instance__] == uint(id) {
						frontcurvestack.SpiralCircles = append(frontcurvestack.SpiralCircles, __instance__)
						break
					}
				}
			}
		}
	case "Color":
		frontcurvestack.Color = value.GetValueString()
	case "FillOpacity":
		frontcurvestack.FillOpacity = value.GetValueFloat()
	case "Stroke":
		frontcurvestack.Stroke = value.GetValueString()
	case "StrokeOpacity":
		frontcurvestack.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		frontcurvestack.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		frontcurvestack.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		frontcurvestack.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		frontcurvestack.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (horizontalaxis *HorizontalAxis) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		horizontalaxis.Name = value.GetValueString()
	case "IsDisplayed":
		horizontalaxis.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			horizontalaxis.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					horizontalaxis.ShapeCategory = __instance__
					break
				}
			}
		}
	case "AxisHandleBorderLength":
		horizontalaxis.AxisHandleBorderLength = value.GetValueFloat()
	case "Axis_Length":
		horizontalaxis.Axis_Length = value.GetValueFloat()
	case "Color":
		horizontalaxis.Color = value.GetValueString()
	case "FillOpacity":
		horizontalaxis.FillOpacity = value.GetValueFloat()
	case "Stroke":
		horizontalaxis.Stroke = value.GetValueString()
	case "StrokeOpacity":
		horizontalaxis.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		horizontalaxis.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		horizontalaxis.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		horizontalaxis.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		horizontalaxis.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (key *Key) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		key.Name = value.GetValueString()
	case "IsDisplayed":
		key.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			key.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					key.ShapeCategory = __instance__
					break
				}
			}
		}
	case "Path":
		key.Path = value.GetValueString()
	case "Color":
		key.Color = value.GetValueString()
	case "FillOpacity":
		key.FillOpacity = value.GetValueFloat()
	case "Stroke":
		key.Stroke = value.GetValueString()
	case "StrokeOpacity":
		key.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		key.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		key.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		key.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		key.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (parameter *Parameter) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		parameter.Name = value.GetValueString()
	case "BackendColor":
		parameter.BackendColor = value.GetValueString()
	case "MinuteColor":
		parameter.MinuteColor = value.GetValueString()
	case "HourColor":
		parameter.HourColor = value.GetValueString()
	case "N":
		parameter.N = int(value.GetValueInt())
	case "M":
		parameter.M = int(value.GetValueInt())
	case "Z":
		parameter.Z = int(value.GetValueInt())
	case "InsideAngle":
		parameter.InsideAngle = value.GetValueFloat()
	case "ShiftToNearestCircle":
		parameter.ShiftToNearestCircle = int(value.GetValueInt())
	case "SideLength":
		parameter.SideLength = value.GetValueFloat()
	case "InitialRhombus":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.InitialRhombus = nil
			for __instance__ := range stage.Rhombuss {
				if stage.Rhombus_stagedOrder[__instance__] == uint(id) {
					parameter.InitialRhombus = __instance__
					break
				}
			}
		}
	case "InitialCircle":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.InitialCircle = nil
			for __instance__ := range stage.Circles {
				if stage.Circle_stagedOrder[__instance__] == uint(id) {
					parameter.InitialCircle = __instance__
					break
				}
			}
		}
	case "InitialRhombusGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.InitialRhombusGrid = nil
			for __instance__ := range stage.RhombusGrids {
				if stage.RhombusGrid_stagedOrder[__instance__] == uint(id) {
					parameter.InitialRhombusGrid = __instance__
					break
				}
			}
		}
	case "InitialCircleGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.InitialCircleGrid = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.InitialCircleGrid = __instance__
					break
				}
			}
		}
	case "InitialAxis":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.InitialAxis = nil
			for __instance__ := range stage.Axiss {
				if stage.Axis_stagedOrder[__instance__] == uint(id) {
					parameter.InitialAxis = __instance__
					break
				}
			}
		}
	case "RotatedAxis":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.RotatedAxis = nil
			for __instance__ := range stage.Axiss {
				if stage.Axis_stagedOrder[__instance__] == uint(id) {
					parameter.RotatedAxis = __instance__
					break
				}
			}
		}
	case "RotatedRhombus":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.RotatedRhombus = nil
			for __instance__ := range stage.Rhombuss {
				if stage.Rhombus_stagedOrder[__instance__] == uint(id) {
					parameter.RotatedRhombus = __instance__
					break
				}
			}
		}
	case "RotatedRhombusGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.RotatedRhombusGrid = nil
			for __instance__ := range stage.RhombusGrids {
				if stage.RhombusGrid_stagedOrder[__instance__] == uint(id) {
					parameter.RotatedRhombusGrid = __instance__
					break
				}
			}
		}
	case "RotatedCircleGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.RotatedCircleGrid = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.RotatedCircleGrid = __instance__
					break
				}
			}
		}
	case "NextRhombus":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.NextRhombus = nil
			for __instance__ := range stage.Rhombuss {
				if stage.Rhombus_stagedOrder[__instance__] == uint(id) {
					parameter.NextRhombus = __instance__
					break
				}
			}
		}
	case "NextCircle":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.NextCircle = nil
			for __instance__ := range stage.Circles {
				if stage.Circle_stagedOrder[__instance__] == uint(id) {
					parameter.NextCircle = __instance__
					break
				}
			}
		}
	case "GrowingRhombusGridSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowingRhombusGridSeed = nil
			for __instance__ := range stage.Rhombuss {
				if stage.Rhombus_stagedOrder[__instance__] == uint(id) {
					parameter.GrowingRhombusGridSeed = __instance__
					break
				}
			}
		}
	case "GrowingRhombusGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowingRhombusGrid = nil
			for __instance__ := range stage.RhombusGrids {
				if stage.RhombusGrid_stagedOrder[__instance__] == uint(id) {
					parameter.GrowingRhombusGrid = __instance__
					break
				}
			}
		}
	case "GrowingCircleGridSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowingCircleGridSeed = nil
			for __instance__ := range stage.Circles {
				if stage.Circle_stagedOrder[__instance__] == uint(id) {
					parameter.GrowingCircleGridSeed = __instance__
					break
				}
			}
		}
	case "GrowingCircleGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowingCircleGrid = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.GrowingCircleGrid = __instance__
					break
				}
			}
		}
	case "GrowingCircleGridLeftSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowingCircleGridLeftSeed = nil
			for __instance__ := range stage.Circles {
				if stage.Circle_stagedOrder[__instance__] == uint(id) {
					parameter.GrowingCircleGridLeftSeed = __instance__
					break
				}
			}
		}
	case "GrowingCircleGridLeft":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowingCircleGridLeft = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.GrowingCircleGridLeft = __instance__
					break
				}
			}
		}
	case "ConstructionAxis":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.ConstructionAxis = nil
			for __instance__ := range stage.Axiss {
				if stage.Axis_stagedOrder[__instance__] == uint(id) {
					parameter.ConstructionAxis = __instance__
					break
				}
			}
		}
	case "ConstructionAxisGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.ConstructionAxisGrid = nil
			for __instance__ := range stage.AxisGrids {
				if stage.AxisGrid_stagedOrder[__instance__] == uint(id) {
					parameter.ConstructionAxisGrid = __instance__
					break
				}
			}
		}
	case "ConstructionCircle":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.ConstructionCircle = nil
			for __instance__ := range stage.Circles {
				if stage.Circle_stagedOrder[__instance__] == uint(id) {
					parameter.ConstructionCircle = __instance__
					break
				}
			}
		}
	case "ConstructionCircleGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.ConstructionCircleGrid = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.ConstructionCircleGrid = __instance__
					break
				}
			}
		}
	case "GrowthCurveSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveSeed = nil
			for __instance__ := range stage.Beziers {
				if stage.Bezier_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveSeed = __instance__
					break
				}
			}
		}
	case "GrowthCurve":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurve = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurve = __instance__
					break
				}
			}
		}
	case "GrowthCurveShiftedRightSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveShiftedRightSeed = nil
			for __instance__ := range stage.Beziers {
				if stage.Bezier_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveShiftedRightSeed = __instance__
					break
				}
			}
		}
	case "GrowthCurveShiftedRight":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveShiftedRight = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveShiftedRight = __instance__
					break
				}
			}
		}
	case "GrowthCurveNextSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveNextSeed = nil
			for __instance__ := range stage.Beziers {
				if stage.Bezier_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveNextSeed = __instance__
					break
				}
			}
		}
	case "GrowthCurveNext":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveNext = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveNext = __instance__
					break
				}
			}
		}
	case "GrowthCurveNextShiftedRightSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveNextShiftedRightSeed = nil
			for __instance__ := range stage.Beziers {
				if stage.Bezier_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveNextShiftedRightSeed = __instance__
					break
				}
			}
		}
	case "GrowthCurveNextShiftedRight":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveNextShiftedRight = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveNextShiftedRight = __instance__
					break
				}
			}
		}
	case "GrowthCurveStack":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.GrowthCurveStack = nil
			for __instance__ := range stage.BezierGridStacks {
				if stage.BezierGridStack_stagedOrder[__instance__] == uint(id) {
					parameter.GrowthCurveStack = __instance__
					break
				}
			}
		}
	case "StackWidth":
		parameter.StackWidth = int(value.GetValueInt())
	case "NbShitRight":
		parameter.NbShitRight = int(value.GetValueInt())
	case "StackHeight":
		parameter.StackHeight = int(value.GetValueInt())
	case "BezierControlLengthRatio":
		parameter.BezierControlLengthRatio = value.GetValueFloat()
	case "SpiralRhombusGridSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralRhombusGridSeed = nil
			for __instance__ := range stage.SpiralRhombuss {
				if stage.SpiralRhombus_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralRhombusGridSeed = __instance__
					break
				}
			}
		}
	case "SpiralRhombusGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralRhombusGrid = nil
			for __instance__ := range stage.SpiralRhombusGrids {
				if stage.SpiralRhombusGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralRhombusGrid = __instance__
					break
				}
			}
		}
	case "SpiralCircleSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralCircleSeed = nil
			for __instance__ := range stage.SpiralCircles {
				if stage.SpiralCircle_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralCircleSeed = __instance__
					break
				}
			}
		}
	case "SpiralCircleGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralCircleGrid = nil
			for __instance__ := range stage.SpiralCircleGrids {
				if stage.SpiralCircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralCircleGrid = __instance__
					break
				}
			}
		}
	case "SpiralCircleFullGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralCircleFullGrid = nil
			for __instance__ := range stage.SpiralCircleGrids {
				if stage.SpiralCircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralCircleFullGrid = __instance__
					break
				}
			}
		}
	case "SpiralConstructionOuterLineSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralConstructionOuterLineSeed = nil
			for __instance__ := range stage.SpiralLines {
				if stage.SpiralLine_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralConstructionOuterLineSeed = __instance__
					break
				}
			}
		}
	case "SpiralConstructionInnerLineSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralConstructionInnerLineSeed = nil
			for __instance__ := range stage.SpiralLines {
				if stage.SpiralLine_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralConstructionInnerLineSeed = __instance__
					break
				}
			}
		}
	case "SpiralConstructionOuterLineGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralConstructionOuterLineGrid = nil
			for __instance__ := range stage.SpiralLineGrids {
				if stage.SpiralLineGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralConstructionOuterLineGrid = __instance__
					break
				}
			}
		}
	case "SpiralConstructionInnerLineGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralConstructionInnerLineGrid = nil
			for __instance__ := range stage.SpiralLineGrids {
				if stage.SpiralLineGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralConstructionInnerLineGrid = __instance__
					break
				}
			}
		}
	case "SpiralConstructionCircleGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralConstructionCircleGrid = nil
			for __instance__ := range stage.SpiralCircleGrids {
				if stage.SpiralCircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralConstructionCircleGrid = __instance__
					break
				}
			}
		}
	case "SpiralConstructionOuterLineFullGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralConstructionOuterLineFullGrid = nil
			for __instance__ := range stage.SpiralLineGrids {
				if stage.SpiralLineGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralConstructionOuterLineFullGrid = __instance__
					break
				}
			}
		}
	case "SpiralBezierSeed":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralBezierSeed = nil
			for __instance__ := range stage.SpiralBeziers {
				if stage.SpiralBezier_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralBezierSeed = __instance__
					break
				}
			}
		}
	case "SpiralBezierGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralBezierGrid = nil
			for __instance__ := range stage.SpiralBezierGrids {
				if stage.SpiralBezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralBezierGrid = __instance__
					break
				}
			}
		}
	case "SpiralBezierFullGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralBezierFullGrid = nil
			for __instance__ := range stage.SpiralBezierGrids {
				if stage.SpiralBezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralBezierFullGrid = __instance__
					break
				}
			}
		}
	case "SpiralBezierStrength":
		parameter.SpiralBezierStrength = value.GetValueFloat()
	case "FrontCurveStack":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.FrontCurveStack = nil
			for __instance__ := range stage.FrontCurveStacks {
				if stage.FrontCurveStack_stagedOrder[__instance__] == uint(id) {
					parameter.FrontCurveStack = __instance__
					break
				}
			}
		}
	case "NbInterpolationPoints":
		parameter.NbInterpolationPoints = int(value.GetValueInt())
	case "Fkey":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.Fkey = nil
			for __instance__ := range stage.Keys {
				if stage.Key_stagedOrder[__instance__] == uint(id) {
					parameter.Fkey = __instance__
					break
				}
			}
		}
	case "FkeySizeRatio":
		parameter.FkeySizeRatio = value.GetValueFloat()
	case "FkeyOriginRelativeX":
		parameter.FkeyOriginRelativeX = value.GetValueFloat()
	case "FkeyOriginRelativeY":
		parameter.FkeyOriginRelativeY = value.GetValueFloat()
	case "PitchLines":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.PitchLines = nil
			for __instance__ := range stage.AxisGrids {
				if stage.AxisGrid_stagedOrder[__instance__] == uint(id) {
					parameter.PitchLines = __instance__
					break
				}
			}
		}
	case "PitchHeight":
		parameter.PitchHeight = value.GetValueFloat()
	case "NbPitchLines":
		parameter.NbPitchLines = int(value.GetValueInt())
	case "BeatLines":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.BeatLines = nil
			for __instance__ := range stage.AxisGrids {
				if stage.AxisGrid_stagedOrder[__instance__] == uint(id) {
					parameter.BeatLines = __instance__
					break
				}
			}
		}
	case "BeatLinesHeightRatio":
		parameter.BeatLinesHeightRatio = value.GetValueFloat()
	case "NbBeatLines":
		parameter.NbBeatLines = int(value.GetValueInt())
	case "NbOfBeatsInTheme":
		parameter.NbOfBeatsInTheme = int(value.GetValueInt())
	case "FirstVoice":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.FirstVoice = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.FirstVoice = __instance__
					break
				}
			}
		}
	case "FirstVoiceShiftedRigth":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.FirstVoiceShiftedRigth = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.FirstVoiceShiftedRigth = __instance__
					break
				}
			}
		}
	case "FirstVoiceShiftX":
		parameter.FirstVoiceShiftX = value.GetValueFloat()
	case "FirstVoiceShiftY":
		parameter.FirstVoiceShiftY = value.GetValueFloat()
	case "SecondVoice":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SecondVoice = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SecondVoice = __instance__
					break
				}
			}
		}
	case "SecondVoiceShiftedRight":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SecondVoiceShiftedRight = nil
			for __instance__ := range stage.BezierGrids {
				if stage.BezierGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SecondVoiceShiftedRight = __instance__
					break
				}
			}
		}
	case "PitchDifference":
		parameter.PitchDifference = int(value.GetValueInt())
	case "BeatsPerSecond":
		parameter.BeatsPerSecond = value.GetValueFloat()
	case "Level":
		parameter.Level = value.GetValueFloat()
	case "FirstVoiceNotes":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.FirstVoiceNotes = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.FirstVoiceNotes = __instance__
					break
				}
			}
		}
	case "FirstVoiceNotesShiftedRight":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.FirstVoiceNotesShiftedRight = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.FirstVoiceNotesShiftedRight = __instance__
					break
				}
			}
		}
	case "SecondVoiceNotes":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SecondVoiceNotes = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SecondVoiceNotes = __instance__
					break
				}
			}
		}
	case "SecondVoiceNotesShiftedRight":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SecondVoiceNotesShiftedRight = nil
			for __instance__ := range stage.CircleGrids {
				if stage.CircleGrid_stagedOrder[__instance__] == uint(id) {
					parameter.SecondVoiceNotesShiftedRight = __instance__
					break
				}
			}
		}
	case "IsMinor":
		parameter.IsMinor = value.GetValueBool()
	case "ThemeBinaryEncoding":
		parameter.ThemeBinaryEncoding = int(value.GetValueInt())
	case "OriginX":
		parameter.OriginX = value.GetValueFloat()
	case "OriginY":
		parameter.OriginY = value.GetValueFloat()
	case "HorizontalAxis":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.HorizontalAxis = nil
			for __instance__ := range stage.HorizontalAxiss {
				if stage.HorizontalAxis_stagedOrder[__instance__] == uint(id) {
					parameter.HorizontalAxis = __instance__
					break
				}
			}
		}
	case "VerticalAxis":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.VerticalAxis = nil
			for __instance__ := range stage.VerticalAxiss {
				if stage.VerticalAxis_stagedOrder[__instance__] == uint(id) {
					parameter.VerticalAxis = __instance__
					break
				}
			}
		}
	case "SpiralOrigin":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			parameter.SpiralOrigin = nil
			for __instance__ := range stage.SpiralOrigins {
				if stage.SpiralOrigin_stagedOrder[__instance__] == uint(id) {
					parameter.SpiralOrigin = __instance__
					break
				}
			}
		}
	case "SpiralOriginX":
		parameter.SpiralOriginX = value.GetValueFloat()
	case "SpiralOriginY":
		parameter.SpiralOriginY = value.GetValueFloat()
	case "OriginCrossWidth":
		parameter.OriginCrossWidth = value.GetValueFloat()
	case "SpiralRadiusRatio":
		parameter.SpiralRadiusRatio = value.GetValueFloat()
	case "ShowSpiralBezierConstruct":
		parameter.ShowSpiralBezierConstruct = value.GetValueBool()
	case "ShowInterpolationPoints":
		parameter.ShowInterpolationPoints = value.GetValueBool()
	case "ActualBeatsTemporalShift":
		parameter.ActualBeatsTemporalShift = int(value.GetValueInt())
	case "PathToStaticFiles":
		parameter.PathToStaticFiles = value.GetValueString()
	case "PathToGeneratedSVG":
		parameter.PathToGeneratedSVG = value.GetValueString()
	case "PathToGeneratedScore":
		parameter.PathToGeneratedScore = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rhombus *Rhombus) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rhombus.Name = value.GetValueString()
	case "IsDisplayed":
		rhombus.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rhombus.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					rhombus.ShapeCategory = __instance__
					break
				}
			}
		}
	case "CenterX":
		rhombus.CenterX = value.GetValueFloat()
	case "CenterY":
		rhombus.CenterY = value.GetValueFloat()
	case "SideLength":
		rhombus.SideLength = value.GetValueFloat()
	case "AngleDegree":
		rhombus.AngleDegree = value.GetValueFloat()
	case "InsideAngle":
		rhombus.InsideAngle = value.GetValueFloat()
	case "Color":
		rhombus.Color = value.GetValueString()
	case "FillOpacity":
		rhombus.FillOpacity = value.GetValueFloat()
	case "Stroke":
		rhombus.Stroke = value.GetValueString()
	case "StrokeOpacity":
		rhombus.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		rhombus.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		rhombus.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		rhombus.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		rhombus.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rhombusgrid *RhombusGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rhombusgrid.Name = value.GetValueString()
	case "Reference":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rhombusgrid.Reference = nil
			for __instance__ := range stage.Rhombuss {
				if stage.Rhombus_stagedOrder[__instance__] == uint(id) {
					rhombusgrid.Reference = __instance__
					break
				}
			}
		}
	case "IsDisplayed":
		rhombusgrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rhombusgrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					rhombusgrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "Rhombuses":
		rhombusgrid.Rhombuses = make([]*Rhombus, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Rhombuss {
					if stage.Rhombus_stagedOrder[__instance__] == uint(id) {
						rhombusgrid.Rhombuses = append(rhombusgrid.Rhombuses, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (shapecategory *ShapeCategory) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		shapecategory.Name = value.GetValueString()
	case "IsExpanded":
		shapecategory.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralbezier *SpiralBezier) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralbezier.Name = value.GetValueString()
	case "IsDisplayed":
		spiralbezier.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralbezier.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralbezier.ShapeCategory = __instance__
					break
				}
			}
		}
	case "StartX":
		spiralbezier.StartX = value.GetValueFloat()
	case "StartY":
		spiralbezier.StartY = value.GetValueFloat()
	case "ControlPointStartX":
		spiralbezier.ControlPointStartX = value.GetValueFloat()
	case "ControlPointStartY":
		spiralbezier.ControlPointStartY = value.GetValueFloat()
	case "EndX":
		spiralbezier.EndX = value.GetValueFloat()
	case "EndY":
		spiralbezier.EndY = value.GetValueFloat()
	case "ControlPointEndX":
		spiralbezier.ControlPointEndX = value.GetValueFloat()
	case "ControlPointEndY":
		spiralbezier.ControlPointEndY = value.GetValueFloat()
	case "Color":
		spiralbezier.Color = value.GetValueString()
	case "FillOpacity":
		spiralbezier.FillOpacity = value.GetValueFloat()
	case "Stroke":
		spiralbezier.Stroke = value.GetValueString()
	case "StrokeOpacity":
		spiralbezier.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		spiralbezier.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		spiralbezier.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		spiralbezier.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		spiralbezier.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralbeziergrid *SpiralBezierGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralbeziergrid.Name = value.GetValueString()
	case "IsDisplayed":
		spiralbeziergrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralbeziergrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralbeziergrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "SpiralBeziers":
		spiralbeziergrid.SpiralBeziers = make([]*SpiralBezier, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SpiralBeziers {
					if stage.SpiralBezier_stagedOrder[__instance__] == uint(id) {
						spiralbeziergrid.SpiralBeziers = append(spiralbeziergrid.SpiralBeziers, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralcircle *SpiralCircle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralcircle.Name = value.GetValueString()
	case "IsDisplayed":
		spiralcircle.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralcircle.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralcircle.ShapeCategory = __instance__
					break
				}
			}
		}
	case "CenterX":
		spiralcircle.CenterX = value.GetValueFloat()
	case "CenterY":
		spiralcircle.CenterY = value.GetValueFloat()
	case "HasBespokeRadius":
		spiralcircle.HasBespokeRadius = value.GetValueBool()
	case "BespopkeRadius":
		spiralcircle.BespopkeRadius = value.GetValueFloat()
	case "Color":
		spiralcircle.Color = value.GetValueString()
	case "FillOpacity":
		spiralcircle.FillOpacity = value.GetValueFloat()
	case "Stroke":
		spiralcircle.Stroke = value.GetValueString()
	case "StrokeOpacity":
		spiralcircle.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		spiralcircle.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		spiralcircle.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		spiralcircle.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		spiralcircle.Transform = value.GetValueString()
	case "Pitch":
		spiralcircle.Pitch = int(value.GetValueInt())
	case "ShowName":
		spiralcircle.ShowName = value.GetValueBool()
	case "BeatNb":
		spiralcircle.BeatNb = int(value.GetValueInt())
	case "Path":
		spiralcircle.Path = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralcirclegrid *SpiralCircleGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralcirclegrid.Name = value.GetValueString()
	case "IsDisplayed":
		spiralcirclegrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralcirclegrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralcirclegrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "SpiralRhombusGrid":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralcirclegrid.SpiralRhombusGrid = nil
			for __instance__ := range stage.SpiralRhombusGrids {
				if stage.SpiralRhombusGrid_stagedOrder[__instance__] == uint(id) {
					spiralcirclegrid.SpiralRhombusGrid = __instance__
					break
				}
			}
		}
	case "SpiralCircles":
		spiralcirclegrid.SpiralCircles = make([]*SpiralCircle, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SpiralCircles {
					if stage.SpiralCircle_stagedOrder[__instance__] == uint(id) {
						spiralcirclegrid.SpiralCircles = append(spiralcirclegrid.SpiralCircles, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralline *SpiralLine) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralline.Name = value.GetValueString()
	case "IsDisplayed":
		spiralline.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralline.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralline.ShapeCategory = __instance__
					break
				}
			}
		}
	case "StartX":
		spiralline.StartX = value.GetValueFloat()
	case "EndX":
		spiralline.EndX = value.GetValueFloat()
	case "StartY":
		spiralline.StartY = value.GetValueFloat()
	case "EndY":
		spiralline.EndY = value.GetValueFloat()
	case "Color":
		spiralline.Color = value.GetValueString()
	case "FillOpacity":
		spiralline.FillOpacity = value.GetValueFloat()
	case "Stroke":
		spiralline.Stroke = value.GetValueString()
	case "StrokeOpacity":
		spiralline.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		spiralline.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		spiralline.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		spiralline.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		spiralline.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spirallinegrid *SpiralLineGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spirallinegrid.Name = value.GetValueString()
	case "IsDisplayed":
		spirallinegrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spirallinegrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spirallinegrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "SpiralLines":
		spirallinegrid.SpiralLines = make([]*SpiralLine, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SpiralLines {
					if stage.SpiralLine_stagedOrder[__instance__] == uint(id) {
						spirallinegrid.SpiralLines = append(spirallinegrid.SpiralLines, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralorigin *SpiralOrigin) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralorigin.Name = value.GetValueString()
	case "IsDisplayed":
		spiralorigin.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralorigin.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralorigin.ShapeCategory = __instance__
					break
				}
			}
		}
	case "Color":
		spiralorigin.Color = value.GetValueString()
	case "FillOpacity":
		spiralorigin.FillOpacity = value.GetValueFloat()
	case "Stroke":
		spiralorigin.Stroke = value.GetValueString()
	case "StrokeOpacity":
		spiralorigin.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		spiralorigin.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		spiralorigin.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		spiralorigin.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		spiralorigin.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralrhombus *SpiralRhombus) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralrhombus.Name = value.GetValueString()
	case "IsDisplayed":
		spiralrhombus.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralrhombus.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralrhombus.ShapeCategory = __instance__
					break
				}
			}
		}
	case "X_r0":
		spiralrhombus.X_r0 = value.GetValueFloat()
	case "Y_r0":
		spiralrhombus.Y_r0 = value.GetValueFloat()
	case "X_r1":
		spiralrhombus.X_r1 = value.GetValueFloat()
	case "Y_r1":
		spiralrhombus.Y_r1 = value.GetValueFloat()
	case "X_r2":
		spiralrhombus.X_r2 = value.GetValueFloat()
	case "Y_r2":
		spiralrhombus.Y_r2 = value.GetValueFloat()
	case "X_r3":
		spiralrhombus.X_r3 = value.GetValueFloat()
	case "Y_r3":
		spiralrhombus.Y_r3 = value.GetValueFloat()
	case "Color":
		spiralrhombus.Color = value.GetValueString()
	case "FillOpacity":
		spiralrhombus.FillOpacity = value.GetValueFloat()
	case "Stroke":
		spiralrhombus.Stroke = value.GetValueString()
	case "StrokeOpacity":
		spiralrhombus.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		spiralrhombus.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		spiralrhombus.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		spiralrhombus.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		spiralrhombus.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		spiralrhombusgrid.Name = value.GetValueString()
	case "IsDisplayed":
		spiralrhombusgrid.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			spiralrhombusgrid.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					spiralrhombusgrid.ShapeCategory = __instance__
					break
				}
			}
		}
	case "SpiralRhombuses":
		spiralrhombusgrid.SpiralRhombuses = make([]*SpiralRhombus, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SpiralRhombuss {
					if stage.SpiralRhombus_stagedOrder[__instance__] == uint(id) {
						spiralrhombusgrid.SpiralRhombuses = append(spiralrhombusgrid.SpiralRhombuses, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (verticalaxis *VerticalAxis) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		verticalaxis.Name = value.GetValueString()
	case "IsDisplayed":
		verticalaxis.IsDisplayed = value.GetValueBool()
	case "ShapeCategory":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			verticalaxis.ShapeCategory = nil
			for __instance__ := range stage.ShapeCategorys {
				if stage.ShapeCategory_stagedOrder[__instance__] == uint(id) {
					verticalaxis.ShapeCategory = __instance__
					break
				}
			}
		}
	case "AxisHandleBorderLength":
		verticalaxis.AxisHandleBorderLength = value.GetValueFloat()
	case "Axis_Length":
		verticalaxis.Axis_Length = value.GetValueFloat()
	case "Color":
		verticalaxis.Color = value.GetValueString()
	case "FillOpacity":
		verticalaxis.FillOpacity = value.GetValueFloat()
	case "Stroke":
		verticalaxis.Stroke = value.GetValueString()
	case "StrokeOpacity":
		verticalaxis.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		verticalaxis.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		verticalaxis.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		verticalaxis.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		verticalaxis.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (axis *Axis) GongGetGongstructName() string {
	return "Axis"
}

func (axisgrid *AxisGrid) GongGetGongstructName() string {
	return "AxisGrid"
}

func (bezier *Bezier) GongGetGongstructName() string {
	return "Bezier"
}

func (beziergrid *BezierGrid) GongGetGongstructName() string {
	return "BezierGrid"
}

func (beziergridstack *BezierGridStack) GongGetGongstructName() string {
	return "BezierGridStack"
}

func (chapter *Chapter) GongGetGongstructName() string {
	return "Chapter"
}

func (circle *Circle) GongGetGongstructName() string {
	return "Circle"
}

func (circlegrid *CircleGrid) GongGetGongstructName() string {
	return "CircleGrid"
}

func (content *Content) GongGetGongstructName() string {
	return "Content"
}

func (exporttomusicxml *ExportToMusicxml) GongGetGongstructName() string {
	return "ExportToMusicxml"
}

func (frontcurve *FrontCurve) GongGetGongstructName() string {
	return "FrontCurve"
}

func (frontcurvestack *FrontCurveStack) GongGetGongstructName() string {
	return "FrontCurveStack"
}

func (horizontalaxis *HorizontalAxis) GongGetGongstructName() string {
	return "HorizontalAxis"
}

func (key *Key) GongGetGongstructName() string {
	return "Key"
}

func (parameter *Parameter) GongGetGongstructName() string {
	return "Parameter"
}

func (rhombus *Rhombus) GongGetGongstructName() string {
	return "Rhombus"
}

func (rhombusgrid *RhombusGrid) GongGetGongstructName() string {
	return "RhombusGrid"
}

func (shapecategory *ShapeCategory) GongGetGongstructName() string {
	return "ShapeCategory"
}

func (spiralbezier *SpiralBezier) GongGetGongstructName() string {
	return "SpiralBezier"
}

func (spiralbeziergrid *SpiralBezierGrid) GongGetGongstructName() string {
	return "SpiralBezierGrid"
}

func (spiralcircle *SpiralCircle) GongGetGongstructName() string {
	return "SpiralCircle"
}

func (spiralcirclegrid *SpiralCircleGrid) GongGetGongstructName() string {
	return "SpiralCircleGrid"
}

func (spiralline *SpiralLine) GongGetGongstructName() string {
	return "SpiralLine"
}

func (spirallinegrid *SpiralLineGrid) GongGetGongstructName() string {
	return "SpiralLineGrid"
}

func (spiralorigin *SpiralOrigin) GongGetGongstructName() string {
	return "SpiralOrigin"
}

func (spiralrhombus *SpiralRhombus) GongGetGongstructName() string {
	return "SpiralRhombus"
}

func (spiralrhombusgrid *SpiralRhombusGrid) GongGetGongstructName() string {
	return "SpiralRhombusGrid"
}

func (verticalaxis *VerticalAxis) GongGetGongstructName() string {
	return "VerticalAxis"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.Axiss_mapString = make(map[string]*Axis)
	for axis := range stage.Axiss {
		stage.Axiss_mapString[axis.Name] = axis
	}

	stage.AxisGrids_mapString = make(map[string]*AxisGrid)
	for axisgrid := range stage.AxisGrids {
		stage.AxisGrids_mapString[axisgrid.Name] = axisgrid
	}

	stage.Beziers_mapString = make(map[string]*Bezier)
	for bezier := range stage.Beziers {
		stage.Beziers_mapString[bezier.Name] = bezier
	}

	stage.BezierGrids_mapString = make(map[string]*BezierGrid)
	for beziergrid := range stage.BezierGrids {
		stage.BezierGrids_mapString[beziergrid.Name] = beziergrid
	}

	stage.BezierGridStacks_mapString = make(map[string]*BezierGridStack)
	for beziergridstack := range stage.BezierGridStacks {
		stage.BezierGridStacks_mapString[beziergridstack.Name] = beziergridstack
	}

	stage.Chapters_mapString = make(map[string]*Chapter)
	for chapter := range stage.Chapters {
		stage.Chapters_mapString[chapter.Name] = chapter
	}

	stage.Circles_mapString = make(map[string]*Circle)
	for circle := range stage.Circles {
		stage.Circles_mapString[circle.Name] = circle
	}

	stage.CircleGrids_mapString = make(map[string]*CircleGrid)
	for circlegrid := range stage.CircleGrids {
		stage.CircleGrids_mapString[circlegrid.Name] = circlegrid
	}

	stage.Contents_mapString = make(map[string]*Content)
	for content := range stage.Contents {
		stage.Contents_mapString[content.Name] = content
	}

	stage.ExportToMusicxmls_mapString = make(map[string]*ExportToMusicxml)
	for exporttomusicxml := range stage.ExportToMusicxmls {
		stage.ExportToMusicxmls_mapString[exporttomusicxml.Name] = exporttomusicxml
	}

	stage.FrontCurves_mapString = make(map[string]*FrontCurve)
	for frontcurve := range stage.FrontCurves {
		stage.FrontCurves_mapString[frontcurve.Name] = frontcurve
	}

	stage.FrontCurveStacks_mapString = make(map[string]*FrontCurveStack)
	for frontcurvestack := range stage.FrontCurveStacks {
		stage.FrontCurveStacks_mapString[frontcurvestack.Name] = frontcurvestack
	}

	stage.HorizontalAxiss_mapString = make(map[string]*HorizontalAxis)
	for horizontalaxis := range stage.HorizontalAxiss {
		stage.HorizontalAxiss_mapString[horizontalaxis.Name] = horizontalaxis
	}

	stage.Keys_mapString = make(map[string]*Key)
	for key := range stage.Keys {
		stage.Keys_mapString[key.Name] = key
	}

	stage.Parameters_mapString = make(map[string]*Parameter)
	for parameter := range stage.Parameters {
		stage.Parameters_mapString[parameter.Name] = parameter
	}

	stage.Rhombuss_mapString = make(map[string]*Rhombus)
	for rhombus := range stage.Rhombuss {
		stage.Rhombuss_mapString[rhombus.Name] = rhombus
	}

	stage.RhombusGrids_mapString = make(map[string]*RhombusGrid)
	for rhombusgrid := range stage.RhombusGrids {
		stage.RhombusGrids_mapString[rhombusgrid.Name] = rhombusgrid
	}

	stage.ShapeCategorys_mapString = make(map[string]*ShapeCategory)
	for shapecategory := range stage.ShapeCategorys {
		stage.ShapeCategorys_mapString[shapecategory.Name] = shapecategory
	}

	stage.SpiralBeziers_mapString = make(map[string]*SpiralBezier)
	for spiralbezier := range stage.SpiralBeziers {
		stage.SpiralBeziers_mapString[spiralbezier.Name] = spiralbezier
	}

	stage.SpiralBezierGrids_mapString = make(map[string]*SpiralBezierGrid)
	for spiralbeziergrid := range stage.SpiralBezierGrids {
		stage.SpiralBezierGrids_mapString[spiralbeziergrid.Name] = spiralbeziergrid
	}

	stage.SpiralCircles_mapString = make(map[string]*SpiralCircle)
	for spiralcircle := range stage.SpiralCircles {
		stage.SpiralCircles_mapString[spiralcircle.Name] = spiralcircle
	}

	stage.SpiralCircleGrids_mapString = make(map[string]*SpiralCircleGrid)
	for spiralcirclegrid := range stage.SpiralCircleGrids {
		stage.SpiralCircleGrids_mapString[spiralcirclegrid.Name] = spiralcirclegrid
	}

	stage.SpiralLines_mapString = make(map[string]*SpiralLine)
	for spiralline := range stage.SpiralLines {
		stage.SpiralLines_mapString[spiralline.Name] = spiralline
	}

	stage.SpiralLineGrids_mapString = make(map[string]*SpiralLineGrid)
	for spirallinegrid := range stage.SpiralLineGrids {
		stage.SpiralLineGrids_mapString[spirallinegrid.Name] = spirallinegrid
	}

	stage.SpiralOrigins_mapString = make(map[string]*SpiralOrigin)
	for spiralorigin := range stage.SpiralOrigins {
		stage.SpiralOrigins_mapString[spiralorigin.Name] = spiralorigin
	}

	stage.SpiralRhombuss_mapString = make(map[string]*SpiralRhombus)
	for spiralrhombus := range stage.SpiralRhombuss {
		stage.SpiralRhombuss_mapString[spiralrhombus.Name] = spiralrhombus
	}

	stage.SpiralRhombusGrids_mapString = make(map[string]*SpiralRhombusGrid)
	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		stage.SpiralRhombusGrids_mapString[spiralrhombusgrid.Name] = spiralrhombusgrid
	}

	stage.VerticalAxiss_mapString = make(map[string]*VerticalAxis)
	for verticalaxis := range stage.VerticalAxiss {
		stage.VerticalAxiss_mapString[verticalaxis.Name] = verticalaxis
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
