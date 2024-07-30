// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

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
}

func (stage *StageStruct) GetType() string {
	return "github.com/thomaspeugeot/phylotaxymusic/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
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
	CommitVerticalAxis(verticalaxis *VerticalAxis)
	CheckoutVerticalAxis(verticalaxis *VerticalAxis)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
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

		VerticalAxiss:           make(map[*VerticalAxis]any),
		VerticalAxiss_mapString: make(map[string]*VerticalAxis),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
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
	stage.Map_GongStructName_InstancesNb["HorizontalAxis"] = len(stage.HorizontalAxiss)
	stage.Map_GongStructName_InstancesNb["Key"] = len(stage.Keys)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)
	stage.Map_GongStructName_InstancesNb["Rhombus"] = len(stage.Rhombuss)
	stage.Map_GongStructName_InstancesNb["RhombusGrid"] = len(stage.RhombusGrids)
	stage.Map_GongStructName_InstancesNb["ShapeCategory"] = len(stage.ShapeCategorys)
	stage.Map_GongStructName_InstancesNb["VerticalAxis"] = len(stage.VerticalAxiss)

}

func (stage *StageStruct) Checkout() {
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
	stage.Map_GongStructName_InstancesNb["HorizontalAxis"] = len(stage.HorizontalAxiss)
	stage.Map_GongStructName_InstancesNb["Key"] = len(stage.Keys)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)
	stage.Map_GongStructName_InstancesNb["Rhombus"] = len(stage.Rhombuss)
	stage.Map_GongStructName_InstancesNb["RhombusGrid"] = len(stage.RhombusGrids)
	stage.Map_GongStructName_InstancesNb["ShapeCategory"] = len(stage.ShapeCategorys)
	stage.Map_GongStructName_InstancesNb["VerticalAxis"] = len(stage.VerticalAxiss)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts axis to the model stage
func (axis *Axis) Stage(stage *StageStruct) *Axis {
	stage.Axiss[axis] = __member
	stage.Axiss_mapString[axis.Name] = axis

	return axis
}

// Unstage removes axis off the model stage
func (axis *Axis) Unstage(stage *StageStruct) *Axis {
	delete(stage.Axiss, axis)
	delete(stage.Axiss_mapString, axis.Name)
	return axis
}

// UnstageVoid removes axis off the model stage
func (axis *Axis) UnstageVoid(stage *StageStruct) {
	delete(stage.Axiss, axis)
	delete(stage.Axiss_mapString, axis.Name)
}

// commit axis to the back repo (if it is already staged)
func (axis *Axis) Commit(stage *StageStruct) *Axis {
	if _, ok := stage.Axiss[axis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAxis(axis)
		}
	}
	return axis
}

func (axis *Axis) CommitVoid(stage *StageStruct) {
	axis.Commit(stage)
}

// Checkout axis to the back repo (if it is already staged)
func (axis *Axis) Checkout(stage *StageStruct) *Axis {
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
func (axisgrid *AxisGrid) Stage(stage *StageStruct) *AxisGrid {
	stage.AxisGrids[axisgrid] = __member
	stage.AxisGrids_mapString[axisgrid.Name] = axisgrid

	return axisgrid
}

// Unstage removes axisgrid off the model stage
func (axisgrid *AxisGrid) Unstage(stage *StageStruct) *AxisGrid {
	delete(stage.AxisGrids, axisgrid)
	delete(stage.AxisGrids_mapString, axisgrid.Name)
	return axisgrid
}

// UnstageVoid removes axisgrid off the model stage
func (axisgrid *AxisGrid) UnstageVoid(stage *StageStruct) {
	delete(stage.AxisGrids, axisgrid)
	delete(stage.AxisGrids_mapString, axisgrid.Name)
}

// commit axisgrid to the back repo (if it is already staged)
func (axisgrid *AxisGrid) Commit(stage *StageStruct) *AxisGrid {
	if _, ok := stage.AxisGrids[axisgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAxisGrid(axisgrid)
		}
	}
	return axisgrid
}

func (axisgrid *AxisGrid) CommitVoid(stage *StageStruct) {
	axisgrid.Commit(stage)
}

// Checkout axisgrid to the back repo (if it is already staged)
func (axisgrid *AxisGrid) Checkout(stage *StageStruct) *AxisGrid {
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
func (bezier *Bezier) Stage(stage *StageStruct) *Bezier {
	stage.Beziers[bezier] = __member
	stage.Beziers_mapString[bezier.Name] = bezier

	return bezier
}

// Unstage removes bezier off the model stage
func (bezier *Bezier) Unstage(stage *StageStruct) *Bezier {
	delete(stage.Beziers, bezier)
	delete(stage.Beziers_mapString, bezier.Name)
	return bezier
}

// UnstageVoid removes bezier off the model stage
func (bezier *Bezier) UnstageVoid(stage *StageStruct) {
	delete(stage.Beziers, bezier)
	delete(stage.Beziers_mapString, bezier.Name)
}

// commit bezier to the back repo (if it is already staged)
func (bezier *Bezier) Commit(stage *StageStruct) *Bezier {
	if _, ok := stage.Beziers[bezier]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezier(bezier)
		}
	}
	return bezier
}

func (bezier *Bezier) CommitVoid(stage *StageStruct) {
	bezier.Commit(stage)
}

// Checkout bezier to the back repo (if it is already staged)
func (bezier *Bezier) Checkout(stage *StageStruct) *Bezier {
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
func (beziergrid *BezierGrid) Stage(stage *StageStruct) *BezierGrid {
	stage.BezierGrids[beziergrid] = __member
	stage.BezierGrids_mapString[beziergrid.Name] = beziergrid

	return beziergrid
}

// Unstage removes beziergrid off the model stage
func (beziergrid *BezierGrid) Unstage(stage *StageStruct) *BezierGrid {
	delete(stage.BezierGrids, beziergrid)
	delete(stage.BezierGrids_mapString, beziergrid.Name)
	return beziergrid
}

// UnstageVoid removes beziergrid off the model stage
func (beziergrid *BezierGrid) UnstageVoid(stage *StageStruct) {
	delete(stage.BezierGrids, beziergrid)
	delete(stage.BezierGrids_mapString, beziergrid.Name)
}

// commit beziergrid to the back repo (if it is already staged)
func (beziergrid *BezierGrid) Commit(stage *StageStruct) *BezierGrid {
	if _, ok := stage.BezierGrids[beziergrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezierGrid(beziergrid)
		}
	}
	return beziergrid
}

func (beziergrid *BezierGrid) CommitVoid(stage *StageStruct) {
	beziergrid.Commit(stage)
}

// Checkout beziergrid to the back repo (if it is already staged)
func (beziergrid *BezierGrid) Checkout(stage *StageStruct) *BezierGrid {
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
func (beziergridstack *BezierGridStack) Stage(stage *StageStruct) *BezierGridStack {
	stage.BezierGridStacks[beziergridstack] = __member
	stage.BezierGridStacks_mapString[beziergridstack.Name] = beziergridstack

	return beziergridstack
}

// Unstage removes beziergridstack off the model stage
func (beziergridstack *BezierGridStack) Unstage(stage *StageStruct) *BezierGridStack {
	delete(stage.BezierGridStacks, beziergridstack)
	delete(stage.BezierGridStacks_mapString, beziergridstack.Name)
	return beziergridstack
}

// UnstageVoid removes beziergridstack off the model stage
func (beziergridstack *BezierGridStack) UnstageVoid(stage *StageStruct) {
	delete(stage.BezierGridStacks, beziergridstack)
	delete(stage.BezierGridStacks_mapString, beziergridstack.Name)
}

// commit beziergridstack to the back repo (if it is already staged)
func (beziergridstack *BezierGridStack) Commit(stage *StageStruct) *BezierGridStack {
	if _, ok := stage.BezierGridStacks[beziergridstack]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBezierGridStack(beziergridstack)
		}
	}
	return beziergridstack
}

func (beziergridstack *BezierGridStack) CommitVoid(stage *StageStruct) {
	beziergridstack.Commit(stage)
}

// Checkout beziergridstack to the back repo (if it is already staged)
func (beziergridstack *BezierGridStack) Checkout(stage *StageStruct) *BezierGridStack {
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
func (circle *Circle) Stage(stage *StageStruct) *Circle {
	stage.Circles[circle] = __member
	stage.Circles_mapString[circle.Name] = circle

	return circle
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage(stage *StageStruct) *Circle {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
	return circle
}

// UnstageVoid removes circle off the model stage
func (circle *Circle) UnstageVoid(stage *StageStruct) {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
}

// commit circle to the back repo (if it is already staged)
func (circle *Circle) Commit(stage *StageStruct) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircle(circle)
		}
	}
	return circle
}

func (circle *Circle) CommitVoid(stage *StageStruct) {
	circle.Commit(stage)
}

// Checkout circle to the back repo (if it is already staged)
func (circle *Circle) Checkout(stage *StageStruct) *Circle {
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
func (circlegrid *CircleGrid) Stage(stage *StageStruct) *CircleGrid {
	stage.CircleGrids[circlegrid] = __member
	stage.CircleGrids_mapString[circlegrid.Name] = circlegrid

	return circlegrid
}

// Unstage removes circlegrid off the model stage
func (circlegrid *CircleGrid) Unstage(stage *StageStruct) *CircleGrid {
	delete(stage.CircleGrids, circlegrid)
	delete(stage.CircleGrids_mapString, circlegrid.Name)
	return circlegrid
}

// UnstageVoid removes circlegrid off the model stage
func (circlegrid *CircleGrid) UnstageVoid(stage *StageStruct) {
	delete(stage.CircleGrids, circlegrid)
	delete(stage.CircleGrids_mapString, circlegrid.Name)
}

// commit circlegrid to the back repo (if it is already staged)
func (circlegrid *CircleGrid) Commit(stage *StageStruct) *CircleGrid {
	if _, ok := stage.CircleGrids[circlegrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircleGrid(circlegrid)
		}
	}
	return circlegrid
}

func (circlegrid *CircleGrid) CommitVoid(stage *StageStruct) {
	circlegrid.Commit(stage)
}

// Checkout circlegrid to the back repo (if it is already staged)
func (circlegrid *CircleGrid) Checkout(stage *StageStruct) *CircleGrid {
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

// Stage puts horizontalaxis to the model stage
func (horizontalaxis *HorizontalAxis) Stage(stage *StageStruct) *HorizontalAxis {
	stage.HorizontalAxiss[horizontalaxis] = __member
	stage.HorizontalAxiss_mapString[horizontalaxis.Name] = horizontalaxis

	return horizontalaxis
}

// Unstage removes horizontalaxis off the model stage
func (horizontalaxis *HorizontalAxis) Unstage(stage *StageStruct) *HorizontalAxis {
	delete(stage.HorizontalAxiss, horizontalaxis)
	delete(stage.HorizontalAxiss_mapString, horizontalaxis.Name)
	return horizontalaxis
}

// UnstageVoid removes horizontalaxis off the model stage
func (horizontalaxis *HorizontalAxis) UnstageVoid(stage *StageStruct) {
	delete(stage.HorizontalAxiss, horizontalaxis)
	delete(stage.HorizontalAxiss_mapString, horizontalaxis.Name)
}

// commit horizontalaxis to the back repo (if it is already staged)
func (horizontalaxis *HorizontalAxis) Commit(stage *StageStruct) *HorizontalAxis {
	if _, ok := stage.HorizontalAxiss[horizontalaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitHorizontalAxis(horizontalaxis)
		}
	}
	return horizontalaxis
}

func (horizontalaxis *HorizontalAxis) CommitVoid(stage *StageStruct) {
	horizontalaxis.Commit(stage)
}

// Checkout horizontalaxis to the back repo (if it is already staged)
func (horizontalaxis *HorizontalAxis) Checkout(stage *StageStruct) *HorizontalAxis {
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
func (key *Key) Stage(stage *StageStruct) *Key {
	stage.Keys[key] = __member
	stage.Keys_mapString[key.Name] = key

	return key
}

// Unstage removes key off the model stage
func (key *Key) Unstage(stage *StageStruct) *Key {
	delete(stage.Keys, key)
	delete(stage.Keys_mapString, key.Name)
	return key
}

// UnstageVoid removes key off the model stage
func (key *Key) UnstageVoid(stage *StageStruct) {
	delete(stage.Keys, key)
	delete(stage.Keys_mapString, key.Name)
}

// commit key to the back repo (if it is already staged)
func (key *Key) Commit(stage *StageStruct) *Key {
	if _, ok := stage.Keys[key]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitKey(key)
		}
	}
	return key
}

func (key *Key) CommitVoid(stage *StageStruct) {
	key.Commit(stage)
}

// Checkout key to the back repo (if it is already staged)
func (key *Key) Checkout(stage *StageStruct) *Key {
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
func (parameter *Parameter) Stage(stage *StageStruct) *Parameter {
	stage.Parameters[parameter] = __member
	stage.Parameters_mapString[parameter.Name] = parameter

	return parameter
}

// Unstage removes parameter off the model stage
func (parameter *Parameter) Unstage(stage *StageStruct) *Parameter {
	delete(stage.Parameters, parameter)
	delete(stage.Parameters_mapString, parameter.Name)
	return parameter
}

// UnstageVoid removes parameter off the model stage
func (parameter *Parameter) UnstageVoid(stage *StageStruct) {
	delete(stage.Parameters, parameter)
	delete(stage.Parameters_mapString, parameter.Name)
}

// commit parameter to the back repo (if it is already staged)
func (parameter *Parameter) Commit(stage *StageStruct) *Parameter {
	if _, ok := stage.Parameters[parameter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParameter(parameter)
		}
	}
	return parameter
}

func (parameter *Parameter) CommitVoid(stage *StageStruct) {
	parameter.Commit(stage)
}

// Checkout parameter to the back repo (if it is already staged)
func (parameter *Parameter) Checkout(stage *StageStruct) *Parameter {
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
func (rhombus *Rhombus) Stage(stage *StageStruct) *Rhombus {
	stage.Rhombuss[rhombus] = __member
	stage.Rhombuss_mapString[rhombus.Name] = rhombus

	return rhombus
}

// Unstage removes rhombus off the model stage
func (rhombus *Rhombus) Unstage(stage *StageStruct) *Rhombus {
	delete(stage.Rhombuss, rhombus)
	delete(stage.Rhombuss_mapString, rhombus.Name)
	return rhombus
}

// UnstageVoid removes rhombus off the model stage
func (rhombus *Rhombus) UnstageVoid(stage *StageStruct) {
	delete(stage.Rhombuss, rhombus)
	delete(stage.Rhombuss_mapString, rhombus.Name)
}

// commit rhombus to the back repo (if it is already staged)
func (rhombus *Rhombus) Commit(stage *StageStruct) *Rhombus {
	if _, ok := stage.Rhombuss[rhombus]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRhombus(rhombus)
		}
	}
	return rhombus
}

func (rhombus *Rhombus) CommitVoid(stage *StageStruct) {
	rhombus.Commit(stage)
}

// Checkout rhombus to the back repo (if it is already staged)
func (rhombus *Rhombus) Checkout(stage *StageStruct) *Rhombus {
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
func (rhombusgrid *RhombusGrid) Stage(stage *StageStruct) *RhombusGrid {
	stage.RhombusGrids[rhombusgrid] = __member
	stage.RhombusGrids_mapString[rhombusgrid.Name] = rhombusgrid

	return rhombusgrid
}

// Unstage removes rhombusgrid off the model stage
func (rhombusgrid *RhombusGrid) Unstage(stage *StageStruct) *RhombusGrid {
	delete(stage.RhombusGrids, rhombusgrid)
	delete(stage.RhombusGrids_mapString, rhombusgrid.Name)
	return rhombusgrid
}

// UnstageVoid removes rhombusgrid off the model stage
func (rhombusgrid *RhombusGrid) UnstageVoid(stage *StageStruct) {
	delete(stage.RhombusGrids, rhombusgrid)
	delete(stage.RhombusGrids_mapString, rhombusgrid.Name)
}

// commit rhombusgrid to the back repo (if it is already staged)
func (rhombusgrid *RhombusGrid) Commit(stage *StageStruct) *RhombusGrid {
	if _, ok := stage.RhombusGrids[rhombusgrid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRhombusGrid(rhombusgrid)
		}
	}
	return rhombusgrid
}

func (rhombusgrid *RhombusGrid) CommitVoid(stage *StageStruct) {
	rhombusgrid.Commit(stage)
}

// Checkout rhombusgrid to the back repo (if it is already staged)
func (rhombusgrid *RhombusGrid) Checkout(stage *StageStruct) *RhombusGrid {
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
func (shapecategory *ShapeCategory) Stage(stage *StageStruct) *ShapeCategory {
	stage.ShapeCategorys[shapecategory] = __member
	stage.ShapeCategorys_mapString[shapecategory.Name] = shapecategory

	return shapecategory
}

// Unstage removes shapecategory off the model stage
func (shapecategory *ShapeCategory) Unstage(stage *StageStruct) *ShapeCategory {
	delete(stage.ShapeCategorys, shapecategory)
	delete(stage.ShapeCategorys_mapString, shapecategory.Name)
	return shapecategory
}

// UnstageVoid removes shapecategory off the model stage
func (shapecategory *ShapeCategory) UnstageVoid(stage *StageStruct) {
	delete(stage.ShapeCategorys, shapecategory)
	delete(stage.ShapeCategorys_mapString, shapecategory.Name)
}

// commit shapecategory to the back repo (if it is already staged)
func (shapecategory *ShapeCategory) Commit(stage *StageStruct) *ShapeCategory {
	if _, ok := stage.ShapeCategorys[shapecategory]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitShapeCategory(shapecategory)
		}
	}
	return shapecategory
}

func (shapecategory *ShapeCategory) CommitVoid(stage *StageStruct) {
	shapecategory.Commit(stage)
}

// Checkout shapecategory to the back repo (if it is already staged)
func (shapecategory *ShapeCategory) Checkout(stage *StageStruct) *ShapeCategory {
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

// Stage puts verticalaxis to the model stage
func (verticalaxis *VerticalAxis) Stage(stage *StageStruct) *VerticalAxis {
	stage.VerticalAxiss[verticalaxis] = __member
	stage.VerticalAxiss_mapString[verticalaxis.Name] = verticalaxis

	return verticalaxis
}

// Unstage removes verticalaxis off the model stage
func (verticalaxis *VerticalAxis) Unstage(stage *StageStruct) *VerticalAxis {
	delete(stage.VerticalAxiss, verticalaxis)
	delete(stage.VerticalAxiss_mapString, verticalaxis.Name)
	return verticalaxis
}

// UnstageVoid removes verticalaxis off the model stage
func (verticalaxis *VerticalAxis) UnstageVoid(stage *StageStruct) {
	delete(stage.VerticalAxiss, verticalaxis)
	delete(stage.VerticalAxiss_mapString, verticalaxis.Name)
}

// commit verticalaxis to the back repo (if it is already staged)
func (verticalaxis *VerticalAxis) Commit(stage *StageStruct) *VerticalAxis {
	if _, ok := stage.VerticalAxiss[verticalaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVerticalAxis(verticalaxis)
		}
	}
	return verticalaxis
}

func (verticalaxis *VerticalAxis) CommitVoid(stage *StageStruct) {
	verticalaxis.Commit(stage)
}

// Checkout verticalaxis to the back repo (if it is already staged)
func (verticalaxis *VerticalAxis) Checkout(stage *StageStruct) *VerticalAxis {
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
	CreateORMHorizontalAxis(HorizontalAxis *HorizontalAxis)
	CreateORMKey(Key *Key)
	CreateORMParameter(Parameter *Parameter)
	CreateORMRhombus(Rhombus *Rhombus)
	CreateORMRhombusGrid(RhombusGrid *RhombusGrid)
	CreateORMShapeCategory(ShapeCategory *ShapeCategory)
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
	DeleteORMHorizontalAxis(HorizontalAxis *HorizontalAxis)
	DeleteORMKey(Key *Key)
	DeleteORMParameter(Parameter *Parameter)
	DeleteORMRhombus(Rhombus *Rhombus)
	DeleteORMRhombusGrid(RhombusGrid *RhombusGrid)
	DeleteORMShapeCategory(ShapeCategory *ShapeCategory)
	DeleteORMVerticalAxis(VerticalAxis *VerticalAxis)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Axiss = make(map[*Axis]any)
	stage.Axiss_mapString = make(map[string]*Axis)

	stage.AxisGrids = make(map[*AxisGrid]any)
	stage.AxisGrids_mapString = make(map[string]*AxisGrid)

	stage.Beziers = make(map[*Bezier]any)
	stage.Beziers_mapString = make(map[string]*Bezier)

	stage.BezierGrids = make(map[*BezierGrid]any)
	stage.BezierGrids_mapString = make(map[string]*BezierGrid)

	stage.BezierGridStacks = make(map[*BezierGridStack]any)
	stage.BezierGridStacks_mapString = make(map[string]*BezierGridStack)

	stage.Circles = make(map[*Circle]any)
	stage.Circles_mapString = make(map[string]*Circle)

	stage.CircleGrids = make(map[*CircleGrid]any)
	stage.CircleGrids_mapString = make(map[string]*CircleGrid)

	stage.HorizontalAxiss = make(map[*HorizontalAxis]any)
	stage.HorizontalAxiss_mapString = make(map[string]*HorizontalAxis)

	stage.Keys = make(map[*Key]any)
	stage.Keys_mapString = make(map[string]*Key)

	stage.Parameters = make(map[*Parameter]any)
	stage.Parameters_mapString = make(map[string]*Parameter)

	stage.Rhombuss = make(map[*Rhombus]any)
	stage.Rhombuss_mapString = make(map[string]*Rhombus)

	stage.RhombusGrids = make(map[*RhombusGrid]any)
	stage.RhombusGrids_mapString = make(map[string]*RhombusGrid)

	stage.ShapeCategorys = make(map[*ShapeCategory]any)
	stage.ShapeCategorys_mapString = make(map[string]*ShapeCategory)

	stage.VerticalAxiss = make(map[*VerticalAxis]any)
	stage.VerticalAxiss_mapString = make(map[string]*VerticalAxis)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
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

	stage.VerticalAxiss = nil
	stage.VerticalAxiss_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
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

	for verticalaxis := range stage.VerticalAxiss {
		verticalaxis.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Axis | AxisGrid | Bezier | BezierGrid | BezierGridStack | Circle | CircleGrid | HorizontalAxis | Key | Parameter | Rhombus | RhombusGrid | ShapeCategory | VerticalAxis
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*Axis | *AxisGrid | *Bezier | *BezierGrid | *BezierGridStack | *Circle | *CircleGrid | *HorizontalAxis | *Key | *Parameter | *Rhombus | *RhombusGrid | *ShapeCategory | *VerticalAxis
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*Axis]any |
		map[*AxisGrid]any |
		map[*Bezier]any |
		map[*BezierGrid]any |
		map[*BezierGridStack]any |
		map[*Circle]any |
		map[*CircleGrid]any |
		map[*HorizontalAxis]any |
		map[*Key]any |
		map[*Parameter]any |
		map[*Rhombus]any |
		map[*RhombusGrid]any |
		map[*ShapeCategory]any |
		map[*VerticalAxis]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Axis |
		map[string]*AxisGrid |
		map[string]*Bezier |
		map[string]*BezierGrid |
		map[string]*BezierGridStack |
		map[string]*Circle |
		map[string]*CircleGrid |
		map[string]*HorizontalAxis |
		map[string]*Key |
		map[string]*Parameter |
		map[string]*Rhombus |
		map[string]*RhombusGrid |
		map[string]*ShapeCategory |
		map[string]*VerticalAxis |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
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
	case map[*VerticalAxis]any:
		return any(&stage.VerticalAxiss).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
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
	case map[string]*VerticalAxis:
		return any(&stage.VerticalAxiss_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
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
	case VerticalAxis:
		return any(&stage.VerticalAxiss).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
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
	case *VerticalAxis:
		return any(&stage.VerticalAxiss).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
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
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case AxisGrid:
		return any(&AxisGrid{
			// Initialisation of associations
			// field is initialized with an instance of Axis with the name of the field
			Reference: &Axis{Name: "Reference"},
			// field is initialized with an instance of Axis with the name of the field
			Axiss: []*Axis{{Name: "Axiss"}},
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case Bezier:
		return any(&Bezier{
			// Initialisation of associations
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case BezierGrid:
		return any(&BezierGrid{
			// Initialisation of associations
			// field is initialized with an instance of Bezier with the name of the field
			Reference: &Bezier{Name: "Reference"},
			// field is initialized with an instance of Bezier with the name of the field
			Beziers: []*Bezier{{Name: "Beziers"}},
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case BezierGridStack:
		return any(&BezierGridStack{
			// Initialisation of associations
			// field is initialized with an instance of BezierGrid with the name of the field
			BezierGrids: []*BezierGrid{{Name: "BezierGrids"}},
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case Circle:
		return any(&Circle{
			// Initialisation of associations
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case CircleGrid:
		return any(&CircleGrid{
			// Initialisation of associations
			// field is initialized with an instance of Circle with the name of the field
			Reference: &Circle{Name: "Reference"},
			// field is initialized with an instance of Circle with the name of the field
			Circles: []*Circle{{Name: "Circles"}},
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case HorizontalAxis:
		return any(&HorizontalAxis{
			// Initialisation of associations
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case Key:
		return any(&Key{
			// Initialisation of associations
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
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
			GrowthCurveSegment: &Bezier{Name: "GrowthCurveSegment"},
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
			// field is initialized with an instance of Key with the name of the field
			Fkey: &Key{Name: "Fkey"},
			// field is initialized with an instance of AxisGrid with the name of the field
			PitchLines: &AxisGrid{Name: "PitchLines"},
			// field is initialized with an instance of AxisGrid with the name of the field
			MeasureLines: &AxisGrid{Name: "MeasureLines"},
			// field is initialized with an instance of BezierGrid with the name of the field
			SecondVoice: &BezierGrid{Name: "SecondVoice"},
			// field is initialized with an instance of HorizontalAxis with the name of the field
			HorizontalAxis: &HorizontalAxis{Name: "HorizontalAxis"},
			// field is initialized with an instance of VerticalAxis with the name of the field
			VerticalAxis: &VerticalAxis{Name: "VerticalAxis"},
		}).(*Type)
	case Rhombus:
		return any(&Rhombus{
			// Initialisation of associations
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case RhombusGrid:
		return any(&RhombusGrid{
			// Initialisation of associations
			// field is initialized with an instance of Rhombus with the name of the field
			Reference: &Rhombus{Name: "Reference"},
			// field is initialized with an instance of Rhombus with the name of the field
			Rhombuses: []*Rhombus{{Name: "Rhombuses"}},
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
		}).(*Type)
	case ShapeCategory:
		return any(&ShapeCategory{
			// Initialisation of associations
		}).(*Type)
	case VerticalAxis:
		return any(&VerticalAxis{
			// Initialisation of associations
			// field is initialized with AbstractShape as it is a composite
			AbstractShape: AbstractShape{
				// per field init
				//
				ShapeCategory: &ShapeCategory{Name: "ShapeCategory"},
			},
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

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
		case "GrowthCurveSegment":
			res := make(map[*Bezier][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.GrowthCurveSegment != nil {
					bezier_ := parameter.GrowthCurveSegment
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
		case "MeasureLines":
			res := make(map[*AxisGrid][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.MeasureLines != nil {
					axisgrid_ := parameter.MeasureLines
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

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
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Angle", "Length", "CenterX", "CenterY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case AxisGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Axiss"}
	case Bezier:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "StartY", "ControlPointStartX", "ControlPointStartY", "EndX", "EndY", "ControlPointEndX", "ControlPointEndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case BezierGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Beziers"}
	case BezierGridStack:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "BezierGrids"}
	case Circle:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "HasBespokeRadius", "BespopkeRadius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case CircleGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Circles"}
	case HorizontalAxis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AxisHandleBorderLength", "Axis_Length", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case Key:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Path", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case Parameter:
		res = []string{"Name", "N", "M", "Z", "InsideAngle", "SideLength", "InitialRhombus", "InitialCircle", "InitialRhombusGrid", "InitialCircleGrid", "InitialAxis", "RotatedAxis", "RotatedRhombus", "RotatedRhombusGrid", "RotatedCircleGrid", "NextRhombus", "NextCircle", "GrowingRhombusGridSeed", "GrowingRhombusGrid", "GrowingCircleGridSeed", "GrowingCircleGrid", "GrowingCircleGridLeftSeed", "GrowingCircleGridLeft", "ConstructionAxis", "ConstructionAxisGrid", "ConstructionCircle", "ConstructionCircleGrid", "GrowthCurveSegment", "GrowthCurve", "GrowthCurveShiftedRightSeed", "GrowthCurveShiftedRight", "GrowthCurveNextSeed", "GrowthCurveNext", "GrowthCurveNextShiftedRightSeed", "GrowthCurveNextShiftedRight", "GrowthCurveStack", "StackWidth", "NbShitRight", "StackHeight", "BezierControlLengthRatio", "Fkey", "FkeySizeRatio", "FkeyOriginRelativeX", "FkeyOriginRelativeY", "PitchLines", "PitchLinesHeightRatio", "NbPitchLines", "MeasureLines", "MeasureLinesHeightRatio", "NbMeasureLines", "NbMeasureLinesPerCurve", "SecondVoice", "PitchDifference", "OriginX", "OriginY", "HorizontalAxis", "VerticalAxis"}
	case Rhombus:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "SideLength", "Angle", "InsideAngle", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case RhombusGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Rhombuses"}
	case ShapeCategory:
		res = []string{"Name", "IsExpanded"}
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
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Angle", "Length", "CenterX", "CenterY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *AxisGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Axiss"}
	case *Bezier:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "StartX", "StartY", "ControlPointStartX", "ControlPointStartY", "EndX", "EndY", "ControlPointEndX", "ControlPointEndY", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *BezierGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Beziers"}
	case *BezierGridStack:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "BezierGrids"}
	case *Circle:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "HasBespokeRadius", "BespopkeRadius", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *CircleGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Circles"}
	case *HorizontalAxis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AxisHandleBorderLength", "Axis_Length", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *Key:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "Path", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *Parameter:
		res = []string{"Name", "N", "M", "Z", "InsideAngle", "SideLength", "InitialRhombus", "InitialCircle", "InitialRhombusGrid", "InitialCircleGrid", "InitialAxis", "RotatedAxis", "RotatedRhombus", "RotatedRhombusGrid", "RotatedCircleGrid", "NextRhombus", "NextCircle", "GrowingRhombusGridSeed", "GrowingRhombusGrid", "GrowingCircleGridSeed", "GrowingCircleGrid", "GrowingCircleGridLeftSeed", "GrowingCircleGridLeft", "ConstructionAxis", "ConstructionAxisGrid", "ConstructionCircle", "ConstructionCircleGrid", "GrowthCurveSegment", "GrowthCurve", "GrowthCurveShiftedRightSeed", "GrowthCurveShiftedRight", "GrowthCurveNextSeed", "GrowthCurveNext", "GrowthCurveNextShiftedRightSeed", "GrowthCurveNextShiftedRight", "GrowthCurveStack", "StackWidth", "NbShitRight", "StackHeight", "BezierControlLengthRatio", "Fkey", "FkeySizeRatio", "FkeyOriginRelativeX", "FkeyOriginRelativeY", "PitchLines", "PitchLinesHeightRatio", "NbPitchLines", "MeasureLines", "MeasureLinesHeightRatio", "NbMeasureLines", "NbMeasureLinesPerCurve", "SecondVoice", "PitchDifference", "OriginX", "OriginY", "HorizontalAxis", "VerticalAxis"}
	case *Rhombus:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "CenterX", "CenterY", "SideLength", "Angle", "InsideAngle", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	case *RhombusGrid:
		res = []string{"Name", "Reference", "IsDisplayed", "ShapeCategory", "Rhombuses"}
	case *ShapeCategory:
		res = []string{"Name", "IsExpanded"}
	case *VerticalAxis:
		res = []string{"Name", "IsDisplayed", "ShapeCategory", "AxisHandleBorderLength", "Axis_Length", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Axis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Angle":
			res = fmt.Sprintf("%f", inferedInstance.Angle)
		case "Length":
			res = fmt.Sprintf("%f", inferedInstance.Length)
		case "CenterX":
			res = fmt.Sprintf("%f", inferedInstance.CenterX)
		case "CenterY":
			res = fmt.Sprintf("%f", inferedInstance.CenterY)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case *AxisGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Axiss":
			for idx, __instance__ := range inferedInstance.Axiss {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Bezier:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res = fmt.Sprintf("%f", inferedInstance.StartX)
		case "StartY":
			res = fmt.Sprintf("%f", inferedInstance.StartY)
		case "ControlPointStartX":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointStartX)
		case "ControlPointStartY":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointStartY)
		case "EndX":
			res = fmt.Sprintf("%f", inferedInstance.EndX)
		case "EndY":
			res = fmt.Sprintf("%f", inferedInstance.EndY)
		case "ControlPointEndX":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointEndX)
		case "ControlPointEndY":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointEndY)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case *BezierGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Beziers":
			for idx, __instance__ := range inferedInstance.Beziers {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *BezierGridStack:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "BezierGrids":
			for idx, __instance__ := range inferedInstance.BezierGrids {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Circle:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res = fmt.Sprintf("%f", inferedInstance.CenterX)
		case "CenterY":
			res = fmt.Sprintf("%f", inferedInstance.CenterY)
		case "HasBespokeRadius":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeRadius)
		case "BespopkeRadius":
			res = fmt.Sprintf("%f", inferedInstance.BespopkeRadius)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case *CircleGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Circles":
			for idx, __instance__ := range inferedInstance.Circles {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *HorizontalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case *Key:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Path":
			res = inferedInstance.Path
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case *Parameter:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "N":
			res = fmt.Sprintf("%d", inferedInstance.N)
		case "M":
			res = fmt.Sprintf("%d", inferedInstance.M)
		case "Z":
			res = fmt.Sprintf("%d", inferedInstance.Z)
		case "InsideAngle":
			res = fmt.Sprintf("%f", inferedInstance.InsideAngle)
		case "SideLength":
			res = fmt.Sprintf("%f", inferedInstance.SideLength)
		case "InitialRhombus":
			if inferedInstance.InitialRhombus != nil {
				res = inferedInstance.InitialRhombus.Name
			}
		case "InitialCircle":
			if inferedInstance.InitialCircle != nil {
				res = inferedInstance.InitialCircle.Name
			}
		case "InitialRhombusGrid":
			if inferedInstance.InitialRhombusGrid != nil {
				res = inferedInstance.InitialRhombusGrid.Name
			}
		case "InitialCircleGrid":
			if inferedInstance.InitialCircleGrid != nil {
				res = inferedInstance.InitialCircleGrid.Name
			}
		case "InitialAxis":
			if inferedInstance.InitialAxis != nil {
				res = inferedInstance.InitialAxis.Name
			}
		case "RotatedAxis":
			if inferedInstance.RotatedAxis != nil {
				res = inferedInstance.RotatedAxis.Name
			}
		case "RotatedRhombus":
			if inferedInstance.RotatedRhombus != nil {
				res = inferedInstance.RotatedRhombus.Name
			}
		case "RotatedRhombusGrid":
			if inferedInstance.RotatedRhombusGrid != nil {
				res = inferedInstance.RotatedRhombusGrid.Name
			}
		case "RotatedCircleGrid":
			if inferedInstance.RotatedCircleGrid != nil {
				res = inferedInstance.RotatedCircleGrid.Name
			}
		case "NextRhombus":
			if inferedInstance.NextRhombus != nil {
				res = inferedInstance.NextRhombus.Name
			}
		case "NextCircle":
			if inferedInstance.NextCircle != nil {
				res = inferedInstance.NextCircle.Name
			}
		case "GrowingRhombusGridSeed":
			if inferedInstance.GrowingRhombusGridSeed != nil {
				res = inferedInstance.GrowingRhombusGridSeed.Name
			}
		case "GrowingRhombusGrid":
			if inferedInstance.GrowingRhombusGrid != nil {
				res = inferedInstance.GrowingRhombusGrid.Name
			}
		case "GrowingCircleGridSeed":
			if inferedInstance.GrowingCircleGridSeed != nil {
				res = inferedInstance.GrowingCircleGridSeed.Name
			}
		case "GrowingCircleGrid":
			if inferedInstance.GrowingCircleGrid != nil {
				res = inferedInstance.GrowingCircleGrid.Name
			}
		case "GrowingCircleGridLeftSeed":
			if inferedInstance.GrowingCircleGridLeftSeed != nil {
				res = inferedInstance.GrowingCircleGridLeftSeed.Name
			}
		case "GrowingCircleGridLeft":
			if inferedInstance.GrowingCircleGridLeft != nil {
				res = inferedInstance.GrowingCircleGridLeft.Name
			}
		case "ConstructionAxis":
			if inferedInstance.ConstructionAxis != nil {
				res = inferedInstance.ConstructionAxis.Name
			}
		case "ConstructionAxisGrid":
			if inferedInstance.ConstructionAxisGrid != nil {
				res = inferedInstance.ConstructionAxisGrid.Name
			}
		case "ConstructionCircle":
			if inferedInstance.ConstructionCircle != nil {
				res = inferedInstance.ConstructionCircle.Name
			}
		case "ConstructionCircleGrid":
			if inferedInstance.ConstructionCircleGrid != nil {
				res = inferedInstance.ConstructionCircleGrid.Name
			}
		case "GrowthCurveSegment":
			if inferedInstance.GrowthCurveSegment != nil {
				res = inferedInstance.GrowthCurveSegment.Name
			}
		case "GrowthCurve":
			if inferedInstance.GrowthCurve != nil {
				res = inferedInstance.GrowthCurve.Name
			}
		case "GrowthCurveShiftedRightSeed":
			if inferedInstance.GrowthCurveShiftedRightSeed != nil {
				res = inferedInstance.GrowthCurveShiftedRightSeed.Name
			}
		case "GrowthCurveShiftedRight":
			if inferedInstance.GrowthCurveShiftedRight != nil {
				res = inferedInstance.GrowthCurveShiftedRight.Name
			}
		case "GrowthCurveNextSeed":
			if inferedInstance.GrowthCurveNextSeed != nil {
				res = inferedInstance.GrowthCurveNextSeed.Name
			}
		case "GrowthCurveNext":
			if inferedInstance.GrowthCurveNext != nil {
				res = inferedInstance.GrowthCurveNext.Name
			}
		case "GrowthCurveNextShiftedRightSeed":
			if inferedInstance.GrowthCurveNextShiftedRightSeed != nil {
				res = inferedInstance.GrowthCurveNextShiftedRightSeed.Name
			}
		case "GrowthCurveNextShiftedRight":
			if inferedInstance.GrowthCurveNextShiftedRight != nil {
				res = inferedInstance.GrowthCurveNextShiftedRight.Name
			}
		case "GrowthCurveStack":
			if inferedInstance.GrowthCurveStack != nil {
				res = inferedInstance.GrowthCurveStack.Name
			}
		case "StackWidth":
			res = fmt.Sprintf("%d", inferedInstance.StackWidth)
		case "NbShitRight":
			res = fmt.Sprintf("%d", inferedInstance.NbShitRight)
		case "StackHeight":
			res = fmt.Sprintf("%d", inferedInstance.StackHeight)
		case "BezierControlLengthRatio":
			res = fmt.Sprintf("%f", inferedInstance.BezierControlLengthRatio)
		case "Fkey":
			if inferedInstance.Fkey != nil {
				res = inferedInstance.Fkey.Name
			}
		case "FkeySizeRatio":
			res = fmt.Sprintf("%f", inferedInstance.FkeySizeRatio)
		case "FkeyOriginRelativeX":
			res = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeX)
		case "FkeyOriginRelativeY":
			res = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeY)
		case "PitchLines":
			if inferedInstance.PitchLines != nil {
				res = inferedInstance.PitchLines.Name
			}
		case "PitchLinesHeightRatio":
			res = fmt.Sprintf("%f", inferedInstance.PitchLinesHeightRatio)
		case "NbPitchLines":
			res = fmt.Sprintf("%d", inferedInstance.NbPitchLines)
		case "MeasureLines":
			if inferedInstance.MeasureLines != nil {
				res = inferedInstance.MeasureLines.Name
			}
		case "MeasureLinesHeightRatio":
			res = fmt.Sprintf("%f", inferedInstance.MeasureLinesHeightRatio)
		case "NbMeasureLines":
			res = fmt.Sprintf("%d", inferedInstance.NbMeasureLines)
		case "NbMeasureLinesPerCurve":
			res = fmt.Sprintf("%d", inferedInstance.NbMeasureLinesPerCurve)
		case "SecondVoice":
			if inferedInstance.SecondVoice != nil {
				res = inferedInstance.SecondVoice.Name
			}
		case "PitchDifference":
			res = fmt.Sprintf("%d", inferedInstance.PitchDifference)
		case "OriginX":
			res = fmt.Sprintf("%f", inferedInstance.OriginX)
		case "OriginY":
			res = fmt.Sprintf("%f", inferedInstance.OriginY)
		case "HorizontalAxis":
			if inferedInstance.HorizontalAxis != nil {
				res = inferedInstance.HorizontalAxis.Name
			}
		case "VerticalAxis":
			if inferedInstance.VerticalAxis != nil {
				res = inferedInstance.VerticalAxis.Name
			}
		}
	case *Rhombus:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res = fmt.Sprintf("%f", inferedInstance.CenterX)
		case "CenterY":
			res = fmt.Sprintf("%f", inferedInstance.CenterY)
		case "SideLength":
			res = fmt.Sprintf("%f", inferedInstance.SideLength)
		case "Angle":
			res = fmt.Sprintf("%f", inferedInstance.Angle)
		case "InsideAngle":
			res = fmt.Sprintf("%f", inferedInstance.InsideAngle)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case *RhombusGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Rhombuses":
			for idx, __instance__ := range inferedInstance.Rhombuses {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *ShapeCategory:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsExpanded":
			res = fmt.Sprintf("%t", inferedInstance.IsExpanded)
		}
	case *VerticalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Axis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Angle":
			res = fmt.Sprintf("%f", inferedInstance.Angle)
		case "Length":
			res = fmt.Sprintf("%f", inferedInstance.Length)
		case "CenterX":
			res = fmt.Sprintf("%f", inferedInstance.CenterX)
		case "CenterY":
			res = fmt.Sprintf("%f", inferedInstance.CenterY)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case AxisGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Axiss":
			for idx, __instance__ := range inferedInstance.Axiss {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Bezier:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "StartX":
			res = fmt.Sprintf("%f", inferedInstance.StartX)
		case "StartY":
			res = fmt.Sprintf("%f", inferedInstance.StartY)
		case "ControlPointStartX":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointStartX)
		case "ControlPointStartY":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointStartY)
		case "EndX":
			res = fmt.Sprintf("%f", inferedInstance.EndX)
		case "EndY":
			res = fmt.Sprintf("%f", inferedInstance.EndY)
		case "ControlPointEndX":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointEndX)
		case "ControlPointEndY":
			res = fmt.Sprintf("%f", inferedInstance.ControlPointEndY)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case BezierGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Beziers":
			for idx, __instance__ := range inferedInstance.Beziers {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case BezierGridStack:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "BezierGrids":
			for idx, __instance__ := range inferedInstance.BezierGrids {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Circle:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res = fmt.Sprintf("%f", inferedInstance.CenterX)
		case "CenterY":
			res = fmt.Sprintf("%f", inferedInstance.CenterY)
		case "HasBespokeRadius":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeRadius)
		case "BespopkeRadius":
			res = fmt.Sprintf("%f", inferedInstance.BespopkeRadius)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case CircleGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Circles":
			for idx, __instance__ := range inferedInstance.Circles {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case HorizontalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case Key:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Path":
			res = inferedInstance.Path
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case Parameter:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "N":
			res = fmt.Sprintf("%d", inferedInstance.N)
		case "M":
			res = fmt.Sprintf("%d", inferedInstance.M)
		case "Z":
			res = fmt.Sprintf("%d", inferedInstance.Z)
		case "InsideAngle":
			res = fmt.Sprintf("%f", inferedInstance.InsideAngle)
		case "SideLength":
			res = fmt.Sprintf("%f", inferedInstance.SideLength)
		case "InitialRhombus":
			if inferedInstance.InitialRhombus != nil {
				res = inferedInstance.InitialRhombus.Name
			}
		case "InitialCircle":
			if inferedInstance.InitialCircle != nil {
				res = inferedInstance.InitialCircle.Name
			}
		case "InitialRhombusGrid":
			if inferedInstance.InitialRhombusGrid != nil {
				res = inferedInstance.InitialRhombusGrid.Name
			}
		case "InitialCircleGrid":
			if inferedInstance.InitialCircleGrid != nil {
				res = inferedInstance.InitialCircleGrid.Name
			}
		case "InitialAxis":
			if inferedInstance.InitialAxis != nil {
				res = inferedInstance.InitialAxis.Name
			}
		case "RotatedAxis":
			if inferedInstance.RotatedAxis != nil {
				res = inferedInstance.RotatedAxis.Name
			}
		case "RotatedRhombus":
			if inferedInstance.RotatedRhombus != nil {
				res = inferedInstance.RotatedRhombus.Name
			}
		case "RotatedRhombusGrid":
			if inferedInstance.RotatedRhombusGrid != nil {
				res = inferedInstance.RotatedRhombusGrid.Name
			}
		case "RotatedCircleGrid":
			if inferedInstance.RotatedCircleGrid != nil {
				res = inferedInstance.RotatedCircleGrid.Name
			}
		case "NextRhombus":
			if inferedInstance.NextRhombus != nil {
				res = inferedInstance.NextRhombus.Name
			}
		case "NextCircle":
			if inferedInstance.NextCircle != nil {
				res = inferedInstance.NextCircle.Name
			}
		case "GrowingRhombusGridSeed":
			if inferedInstance.GrowingRhombusGridSeed != nil {
				res = inferedInstance.GrowingRhombusGridSeed.Name
			}
		case "GrowingRhombusGrid":
			if inferedInstance.GrowingRhombusGrid != nil {
				res = inferedInstance.GrowingRhombusGrid.Name
			}
		case "GrowingCircleGridSeed":
			if inferedInstance.GrowingCircleGridSeed != nil {
				res = inferedInstance.GrowingCircleGridSeed.Name
			}
		case "GrowingCircleGrid":
			if inferedInstance.GrowingCircleGrid != nil {
				res = inferedInstance.GrowingCircleGrid.Name
			}
		case "GrowingCircleGridLeftSeed":
			if inferedInstance.GrowingCircleGridLeftSeed != nil {
				res = inferedInstance.GrowingCircleGridLeftSeed.Name
			}
		case "GrowingCircleGridLeft":
			if inferedInstance.GrowingCircleGridLeft != nil {
				res = inferedInstance.GrowingCircleGridLeft.Name
			}
		case "ConstructionAxis":
			if inferedInstance.ConstructionAxis != nil {
				res = inferedInstance.ConstructionAxis.Name
			}
		case "ConstructionAxisGrid":
			if inferedInstance.ConstructionAxisGrid != nil {
				res = inferedInstance.ConstructionAxisGrid.Name
			}
		case "ConstructionCircle":
			if inferedInstance.ConstructionCircle != nil {
				res = inferedInstance.ConstructionCircle.Name
			}
		case "ConstructionCircleGrid":
			if inferedInstance.ConstructionCircleGrid != nil {
				res = inferedInstance.ConstructionCircleGrid.Name
			}
		case "GrowthCurveSegment":
			if inferedInstance.GrowthCurveSegment != nil {
				res = inferedInstance.GrowthCurveSegment.Name
			}
		case "GrowthCurve":
			if inferedInstance.GrowthCurve != nil {
				res = inferedInstance.GrowthCurve.Name
			}
		case "GrowthCurveShiftedRightSeed":
			if inferedInstance.GrowthCurveShiftedRightSeed != nil {
				res = inferedInstance.GrowthCurveShiftedRightSeed.Name
			}
		case "GrowthCurveShiftedRight":
			if inferedInstance.GrowthCurveShiftedRight != nil {
				res = inferedInstance.GrowthCurveShiftedRight.Name
			}
		case "GrowthCurveNextSeed":
			if inferedInstance.GrowthCurveNextSeed != nil {
				res = inferedInstance.GrowthCurveNextSeed.Name
			}
		case "GrowthCurveNext":
			if inferedInstance.GrowthCurveNext != nil {
				res = inferedInstance.GrowthCurveNext.Name
			}
		case "GrowthCurveNextShiftedRightSeed":
			if inferedInstance.GrowthCurveNextShiftedRightSeed != nil {
				res = inferedInstance.GrowthCurveNextShiftedRightSeed.Name
			}
		case "GrowthCurveNextShiftedRight":
			if inferedInstance.GrowthCurveNextShiftedRight != nil {
				res = inferedInstance.GrowthCurveNextShiftedRight.Name
			}
		case "GrowthCurveStack":
			if inferedInstance.GrowthCurveStack != nil {
				res = inferedInstance.GrowthCurveStack.Name
			}
		case "StackWidth":
			res = fmt.Sprintf("%d", inferedInstance.StackWidth)
		case "NbShitRight":
			res = fmt.Sprintf("%d", inferedInstance.NbShitRight)
		case "StackHeight":
			res = fmt.Sprintf("%d", inferedInstance.StackHeight)
		case "BezierControlLengthRatio":
			res = fmt.Sprintf("%f", inferedInstance.BezierControlLengthRatio)
		case "Fkey":
			if inferedInstance.Fkey != nil {
				res = inferedInstance.Fkey.Name
			}
		case "FkeySizeRatio":
			res = fmt.Sprintf("%f", inferedInstance.FkeySizeRatio)
		case "FkeyOriginRelativeX":
			res = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeX)
		case "FkeyOriginRelativeY":
			res = fmt.Sprintf("%f", inferedInstance.FkeyOriginRelativeY)
		case "PitchLines":
			if inferedInstance.PitchLines != nil {
				res = inferedInstance.PitchLines.Name
			}
		case "PitchLinesHeightRatio":
			res = fmt.Sprintf("%f", inferedInstance.PitchLinesHeightRatio)
		case "NbPitchLines":
			res = fmt.Sprintf("%d", inferedInstance.NbPitchLines)
		case "MeasureLines":
			if inferedInstance.MeasureLines != nil {
				res = inferedInstance.MeasureLines.Name
			}
		case "MeasureLinesHeightRatio":
			res = fmt.Sprintf("%f", inferedInstance.MeasureLinesHeightRatio)
		case "NbMeasureLines":
			res = fmt.Sprintf("%d", inferedInstance.NbMeasureLines)
		case "NbMeasureLinesPerCurve":
			res = fmt.Sprintf("%d", inferedInstance.NbMeasureLinesPerCurve)
		case "SecondVoice":
			if inferedInstance.SecondVoice != nil {
				res = inferedInstance.SecondVoice.Name
			}
		case "PitchDifference":
			res = fmt.Sprintf("%d", inferedInstance.PitchDifference)
		case "OriginX":
			res = fmt.Sprintf("%f", inferedInstance.OriginX)
		case "OriginY":
			res = fmt.Sprintf("%f", inferedInstance.OriginY)
		case "HorizontalAxis":
			if inferedInstance.HorizontalAxis != nil {
				res = inferedInstance.HorizontalAxis.Name
			}
		case "VerticalAxis":
			if inferedInstance.VerticalAxis != nil {
				res = inferedInstance.VerticalAxis.Name
			}
		}
	case Rhombus:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "CenterX":
			res = fmt.Sprintf("%f", inferedInstance.CenterX)
		case "CenterY":
			res = fmt.Sprintf("%f", inferedInstance.CenterY)
		case "SideLength":
			res = fmt.Sprintf("%f", inferedInstance.SideLength)
		case "Angle":
			res = fmt.Sprintf("%f", inferedInstance.Angle)
		case "InsideAngle":
			res = fmt.Sprintf("%f", inferedInstance.InsideAngle)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	case RhombusGrid:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Reference":
			if inferedInstance.Reference != nil {
				res = inferedInstance.Reference.Name
			}
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "Rhombuses":
			for idx, __instance__ := range inferedInstance.Rhombuses {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case ShapeCategory:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsExpanded":
			res = fmt.Sprintf("%t", inferedInstance.IsExpanded)
		}
	case VerticalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "ShapeCategory":
			if inferedInstance.ShapeCategory != nil {
				res = inferedInstance.ShapeCategory.Name
			}
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "Color":
			res = inferedInstance.Color
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeOpacity":
			res = fmt.Sprintf("%f", inferedInstance.StrokeOpacity)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "StrokeDashArray":
			res = inferedInstance.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res = inferedInstance.StrokeDashArrayWhenSelected
		case "Transform":
			res = inferedInstance.Transform
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
