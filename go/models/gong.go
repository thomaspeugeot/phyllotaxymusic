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
	HorizontalAxiss           map[*HorizontalAxis]any
	HorizontalAxiss_mapString map[string]*HorizontalAxis

	// insertion point for slice of pointers maps

	OnAfterHorizontalAxisCreateCallback OnAfterCreateInterface[HorizontalAxis]
	OnAfterHorizontalAxisUpdateCallback OnAfterUpdateInterface[HorizontalAxis]
	OnAfterHorizontalAxisDeleteCallback OnAfterDeleteInterface[HorizontalAxis]
	OnAfterHorizontalAxisReadCallback   OnAfterReadInterface[HorizontalAxis]

	InitialAxiss           map[*InitialAxis]any
	InitialAxiss_mapString map[string]*InitialAxis

	// insertion point for slice of pointers maps

	OnAfterInitialAxisCreateCallback OnAfterCreateInterface[InitialAxis]
	OnAfterInitialAxisUpdateCallback OnAfterUpdateInterface[InitialAxis]
	OnAfterInitialAxisDeleteCallback OnAfterDeleteInterface[InitialAxis]
	OnAfterInitialAxisReadCallback   OnAfterReadInterface[InitialAxis]

	Lines           map[*Line]any
	Lines_mapString map[string]*Line

	// insertion point for slice of pointers maps

	OnAfterLineCreateCallback OnAfterCreateInterface[Line]
	OnAfterLineUpdateCallback OnAfterUpdateInterface[Line]
	OnAfterLineDeleteCallback OnAfterDeleteInterface[Line]
	OnAfterLineReadCallback   OnAfterReadInterface[Line]

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
	CommitHorizontalAxis(horizontalaxis *HorizontalAxis)
	CheckoutHorizontalAxis(horizontalaxis *HorizontalAxis)
	CommitInitialAxis(initialaxis *InitialAxis)
	CheckoutInitialAxis(initialaxis *InitialAxis)
	CommitLine(line *Line)
	CheckoutLine(line *Line)
	CommitParameter(parameter *Parameter)
	CheckoutParameter(parameter *Parameter)
	CommitRhombus(rhombus *Rhombus)
	CheckoutRhombus(rhombus *Rhombus)
	CommitRhombusGrid(rhombusgrid *RhombusGrid)
	CheckoutRhombusGrid(rhombusgrid *RhombusGrid)
	CommitVerticalAxis(verticalaxis *VerticalAxis)
	CheckoutVerticalAxis(verticalaxis *VerticalAxis)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		HorizontalAxiss:           make(map[*HorizontalAxis]any),
		HorizontalAxiss_mapString: make(map[string]*HorizontalAxis),

		InitialAxiss:           make(map[*InitialAxis]any),
		InitialAxiss_mapString: make(map[string]*InitialAxis),

		Lines:           make(map[*Line]any),
		Lines_mapString: make(map[string]*Line),

		Parameters:           make(map[*Parameter]any),
		Parameters_mapString: make(map[string]*Parameter),

		Rhombuss:           make(map[*Rhombus]any),
		Rhombuss_mapString: make(map[string]*Rhombus),

		RhombusGrids:           make(map[*RhombusGrid]any),
		RhombusGrids_mapString: make(map[string]*RhombusGrid),

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
	stage.Map_GongStructName_InstancesNb["HorizontalAxis"] = len(stage.HorizontalAxiss)
	stage.Map_GongStructName_InstancesNb["InitialAxis"] = len(stage.InitialAxiss)
	stage.Map_GongStructName_InstancesNb["Line"] = len(stage.Lines)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)
	stage.Map_GongStructName_InstancesNb["Rhombus"] = len(stage.Rhombuss)
	stage.Map_GongStructName_InstancesNb["RhombusGrid"] = len(stage.RhombusGrids)
	stage.Map_GongStructName_InstancesNb["VerticalAxis"] = len(stage.VerticalAxiss)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["HorizontalAxis"] = len(stage.HorizontalAxiss)
	stage.Map_GongStructName_InstancesNb["InitialAxis"] = len(stage.InitialAxiss)
	stage.Map_GongStructName_InstancesNb["Line"] = len(stage.Lines)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)
	stage.Map_GongStructName_InstancesNb["Rhombus"] = len(stage.Rhombuss)
	stage.Map_GongStructName_InstancesNb["RhombusGrid"] = len(stage.RhombusGrids)
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

// Stage puts initialaxis to the model stage
func (initialaxis *InitialAxis) Stage(stage *StageStruct) *InitialAxis {
	stage.InitialAxiss[initialaxis] = __member
	stage.InitialAxiss_mapString[initialaxis.Name] = initialaxis

	return initialaxis
}

// Unstage removes initialaxis off the model stage
func (initialaxis *InitialAxis) Unstage(stage *StageStruct) *InitialAxis {
	delete(stage.InitialAxiss, initialaxis)
	delete(stage.InitialAxiss_mapString, initialaxis.Name)
	return initialaxis
}

// UnstageVoid removes initialaxis off the model stage
func (initialaxis *InitialAxis) UnstageVoid(stage *StageStruct) {
	delete(stage.InitialAxiss, initialaxis)
	delete(stage.InitialAxiss_mapString, initialaxis.Name)
}

// commit initialaxis to the back repo (if it is already staged)
func (initialaxis *InitialAxis) Commit(stage *StageStruct) *InitialAxis {
	if _, ok := stage.InitialAxiss[initialaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitInitialAxis(initialaxis)
		}
	}
	return initialaxis
}

func (initialaxis *InitialAxis) CommitVoid(stage *StageStruct) {
	initialaxis.Commit(stage)
}

// Checkout initialaxis to the back repo (if it is already staged)
func (initialaxis *InitialAxis) Checkout(stage *StageStruct) *InitialAxis {
	if _, ok := stage.InitialAxiss[initialaxis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutInitialAxis(initialaxis)
		}
	}
	return initialaxis
}

// for satisfaction of GongStruct interface
func (initialaxis *InitialAxis) GetName() (res string) {
	return initialaxis.Name
}

// Stage puts line to the model stage
func (line *Line) Stage(stage *StageStruct) *Line {
	stage.Lines[line] = __member
	stage.Lines_mapString[line.Name] = line

	return line
}

// Unstage removes line off the model stage
func (line *Line) Unstage(stage *StageStruct) *Line {
	delete(stage.Lines, line)
	delete(stage.Lines_mapString, line.Name)
	return line
}

// UnstageVoid removes line off the model stage
func (line *Line) UnstageVoid(stage *StageStruct) {
	delete(stage.Lines, line)
	delete(stage.Lines_mapString, line.Name)
}

// commit line to the back repo (if it is already staged)
func (line *Line) Commit(stage *StageStruct) *Line {
	if _, ok := stage.Lines[line]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLine(line)
		}
	}
	return line
}

func (line *Line) CommitVoid(stage *StageStruct) {
	line.Commit(stage)
}

// Checkout line to the back repo (if it is already staged)
func (line *Line) Checkout(stage *StageStruct) *Line {
	if _, ok := stage.Lines[line]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLine(line)
		}
	}
	return line
}

// for satisfaction of GongStruct interface
func (line *Line) GetName() (res string) {
	return line.Name
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
	CreateORMHorizontalAxis(HorizontalAxis *HorizontalAxis)
	CreateORMInitialAxis(InitialAxis *InitialAxis)
	CreateORMLine(Line *Line)
	CreateORMParameter(Parameter *Parameter)
	CreateORMRhombus(Rhombus *Rhombus)
	CreateORMRhombusGrid(RhombusGrid *RhombusGrid)
	CreateORMVerticalAxis(VerticalAxis *VerticalAxis)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMHorizontalAxis(HorizontalAxis *HorizontalAxis)
	DeleteORMInitialAxis(InitialAxis *InitialAxis)
	DeleteORMLine(Line *Line)
	DeleteORMParameter(Parameter *Parameter)
	DeleteORMRhombus(Rhombus *Rhombus)
	DeleteORMRhombusGrid(RhombusGrid *RhombusGrid)
	DeleteORMVerticalAxis(VerticalAxis *VerticalAxis)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.HorizontalAxiss = make(map[*HorizontalAxis]any)
	stage.HorizontalAxiss_mapString = make(map[string]*HorizontalAxis)

	stage.InitialAxiss = make(map[*InitialAxis]any)
	stage.InitialAxiss_mapString = make(map[string]*InitialAxis)

	stage.Lines = make(map[*Line]any)
	stage.Lines_mapString = make(map[string]*Line)

	stage.Parameters = make(map[*Parameter]any)
	stage.Parameters_mapString = make(map[string]*Parameter)

	stage.Rhombuss = make(map[*Rhombus]any)
	stage.Rhombuss_mapString = make(map[string]*Rhombus)

	stage.RhombusGrids = make(map[*RhombusGrid]any)
	stage.RhombusGrids_mapString = make(map[string]*RhombusGrid)

	stage.VerticalAxiss = make(map[*VerticalAxis]any)
	stage.VerticalAxiss_mapString = make(map[string]*VerticalAxis)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.HorizontalAxiss = nil
	stage.HorizontalAxiss_mapString = nil

	stage.InitialAxiss = nil
	stage.InitialAxiss_mapString = nil

	stage.Lines = nil
	stage.Lines_mapString = nil

	stage.Parameters = nil
	stage.Parameters_mapString = nil

	stage.Rhombuss = nil
	stage.Rhombuss_mapString = nil

	stage.RhombusGrids = nil
	stage.RhombusGrids_mapString = nil

	stage.VerticalAxiss = nil
	stage.VerticalAxiss_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for horizontalaxis := range stage.HorizontalAxiss {
		horizontalaxis.Unstage(stage)
	}

	for initialaxis := range stage.InitialAxiss {
		initialaxis.Unstage(stage)
	}

	for line := range stage.Lines {
		line.Unstage(stage)
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
	HorizontalAxis | InitialAxis | Line | Parameter | Rhombus | RhombusGrid | VerticalAxis
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
	*HorizontalAxis | *InitialAxis | *Line | *Parameter | *Rhombus | *RhombusGrid | *VerticalAxis
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
		map[*HorizontalAxis]any |
		map[*InitialAxis]any |
		map[*Line]any |
		map[*Parameter]any |
		map[*Rhombus]any |
		map[*RhombusGrid]any |
		map[*VerticalAxis]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*HorizontalAxis |
		map[string]*InitialAxis |
		map[string]*Line |
		map[string]*Parameter |
		map[string]*Rhombus |
		map[string]*RhombusGrid |
		map[string]*VerticalAxis |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*HorizontalAxis]any:
		return any(&stage.HorizontalAxiss).(*Type)
	case map[*InitialAxis]any:
		return any(&stage.InitialAxiss).(*Type)
	case map[*Line]any:
		return any(&stage.Lines).(*Type)
	case map[*Parameter]any:
		return any(&stage.Parameters).(*Type)
	case map[*Rhombus]any:
		return any(&stage.Rhombuss).(*Type)
	case map[*RhombusGrid]any:
		return any(&stage.RhombusGrids).(*Type)
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
	case map[string]*HorizontalAxis:
		return any(&stage.HorizontalAxiss_mapString).(*Type)
	case map[string]*InitialAxis:
		return any(&stage.InitialAxiss_mapString).(*Type)
	case map[string]*Line:
		return any(&stage.Lines_mapString).(*Type)
	case map[string]*Parameter:
		return any(&stage.Parameters_mapString).(*Type)
	case map[string]*Rhombus:
		return any(&stage.Rhombuss_mapString).(*Type)
	case map[string]*RhombusGrid:
		return any(&stage.RhombusGrids_mapString).(*Type)
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
	case HorizontalAxis:
		return any(&stage.HorizontalAxiss).(*map[*Type]any)
	case InitialAxis:
		return any(&stage.InitialAxiss).(*map[*Type]any)
	case Line:
		return any(&stage.Lines).(*map[*Type]any)
	case Parameter:
		return any(&stage.Parameters).(*map[*Type]any)
	case Rhombus:
		return any(&stage.Rhombuss).(*map[*Type]any)
	case RhombusGrid:
		return any(&stage.RhombusGrids).(*map[*Type]any)
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
	case *HorizontalAxis:
		return any(&stage.HorizontalAxiss).(*map[Type]any)
	case *InitialAxis:
		return any(&stage.InitialAxiss).(*map[Type]any)
	case *Line:
		return any(&stage.Lines).(*map[Type]any)
	case *Parameter:
		return any(&stage.Parameters).(*map[Type]any)
	case *Rhombus:
		return any(&stage.Rhombuss).(*map[Type]any)
	case *RhombusGrid:
		return any(&stage.RhombusGrids).(*map[Type]any)
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
	case HorizontalAxis:
		return any(&stage.HorizontalAxiss_mapString).(*map[string]*Type)
	case InitialAxis:
		return any(&stage.InitialAxiss_mapString).(*map[string]*Type)
	case Line:
		return any(&stage.Lines_mapString).(*map[string]*Type)
	case Parameter:
		return any(&stage.Parameters_mapString).(*map[string]*Type)
	case Rhombus:
		return any(&stage.Rhombuss_mapString).(*map[string]*Type)
	case RhombusGrid:
		return any(&stage.RhombusGrids_mapString).(*map[string]*Type)
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
	case HorizontalAxis:
		return any(&HorizontalAxis{
			// Initialisation of associations
		}).(*Type)
	case InitialAxis:
		return any(&InitialAxis{
			// Initialisation of associations
		}).(*Type)
	case Line:
		return any(&Line{
			// Initialisation of associations
		}).(*Type)
	case Parameter:
		return any(&Parameter{
			// Initialisation of associations
			// field is initialized with an instance of Rhombus with the name of the field
			InitialRhombus: &Rhombus{Name: "InitialRhombus"},
			// field is initialized with an instance of RhombusGrid with the name of the field
			InitialRhombusGrid: &RhombusGrid{Name: "InitialRhombusGrid"},
			// field is initialized with an instance of InitialAxis with the name of the field
			InitialAxis: &InitialAxis{Name: "InitialAxis"},
			// field is initialized with an instance of HorizontalAxis with the name of the field
			HorizontalAxis: &HorizontalAxis{Name: "HorizontalAxis"},
			// field is initialized with an instance of VerticalAxis with the name of the field
			VerticalAxis: &VerticalAxis{Name: "VerticalAxis"},
		}).(*Type)
	case Rhombus:
		return any(&Rhombus{
			// Initialisation of associations
		}).(*Type)
	case RhombusGrid:
		return any(&RhombusGrid{
			// Initialisation of associations
			// field is initialized with an instance of Rhombus with the name of the field
			Reference: &Rhombus{Name: "Reference"},
			// field is initialized with an instance of Rhombus with the name of the field
			Rhombuses: []*Rhombus{{Name: "Rhombuses"}},
		}).(*Type)
	case VerticalAxis:
		return any(&VerticalAxis{
			// Initialisation of associations
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
	// reverse maps of direct associations of HorizontalAxis
	case HorizontalAxis:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InitialAxis
	case InitialAxis:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Line
	case Line:
		switch fieldname {
		// insertion point for per direct association field
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
		case "InitialAxis":
			res := make(map[*InitialAxis][]*Parameter)
			for parameter := range stage.Parameters {
				if parameter.InitialAxis != nil {
					initialaxis_ := parameter.InitialAxis
					var parameters []*Parameter
					_, ok := res[initialaxis_]
					if ok {
						parameters = res[initialaxis_]
					} else {
						parameters = make([]*Parameter, 0)
					}
					parameters = append(parameters, parameter)
					res[initialaxis_] = parameters
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
		}
	// reverse maps of direct associations of VerticalAxis
	case VerticalAxis:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of HorizontalAxis
	case HorizontalAxis:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InitialAxis
	case InitialAxis:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Line
	case Line:
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
	case HorizontalAxis:
		res = "HorizontalAxis"
	case InitialAxis:
		res = "InitialAxis"
	case Line:
		res = "Line"
	case Parameter:
		res = "Parameter"
	case Rhombus:
		res = "Rhombus"
	case RhombusGrid:
		res = "RhombusGrid"
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
	case *HorizontalAxis:
		res = "HorizontalAxis"
	case *InitialAxis:
		res = "InitialAxis"
	case *Line:
		res = "Line"
	case *Parameter:
		res = "Parameter"
	case *Rhombus:
		res = "Rhombus"
	case *RhombusGrid:
		res = "RhombusGrid"
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
	case HorizontalAxis:
		res = []string{"Name", "IsDisplayed", "AxisHandleBorderLength", "Axis_Length", "StrokeWidth"}
	case InitialAxis:
		res = []string{"Name", "IsDisplayed", "Angle", "Length", "StrokeWidth"}
	case Line:
		res = []string{"Name", "X1", "Y1", "X2", "Y2"}
	case Parameter:
		res = []string{"Name", "N", "M", "InsideAngle", "DiamondSideLenght", "InitialRhombus", "InitialRhombusGrid", "InitialAxis", "OriginX", "OriginY", "HorizontalAxis", "VerticalAxis"}
	case Rhombus:
		res = []string{"Name", "IsDisplayed", "CenterX", "CenterY", "SideLength", "Angle", "InsideAngle", "StrokeWidth"}
	case RhombusGrid:
		res = []string{"Name", "Reference", "N", "M", "IsDisplayed", "Rhombuses"}
	case VerticalAxis:
		res = []string{"Name", "IsDisplayed", "AxisHandleBorderLength", "Axis_Length", "StrokeWidth"}
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
	case HorizontalAxis:
		var rf ReverseField
		_ = rf
	case InitialAxis:
		var rf ReverseField
		_ = rf
	case Line:
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
	case *HorizontalAxis:
		res = []string{"Name", "IsDisplayed", "AxisHandleBorderLength", "Axis_Length", "StrokeWidth"}
	case *InitialAxis:
		res = []string{"Name", "IsDisplayed", "Angle", "Length", "StrokeWidth"}
	case *Line:
		res = []string{"Name", "X1", "Y1", "X2", "Y2"}
	case *Parameter:
		res = []string{"Name", "N", "M", "InsideAngle", "DiamondSideLenght", "InitialRhombus", "InitialRhombusGrid", "InitialAxis", "OriginX", "OriginY", "HorizontalAxis", "VerticalAxis"}
	case *Rhombus:
		res = []string{"Name", "IsDisplayed", "CenterX", "CenterY", "SideLength", "Angle", "InsideAngle", "StrokeWidth"}
	case *RhombusGrid:
		res = []string{"Name", "Reference", "N", "M", "IsDisplayed", "Rhombuses"}
	case *VerticalAxis:
		res = []string{"Name", "IsDisplayed", "AxisHandleBorderLength", "Axis_Length", "StrokeWidth"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *HorizontalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		}
	case *InitialAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "Angle":
			res = fmt.Sprintf("%f", inferedInstance.Angle)
		case "Length":
			res = fmt.Sprintf("%f", inferedInstance.Length)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		}
	case *Line:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X1":
			res = fmt.Sprintf("%f", inferedInstance.X1)
		case "Y1":
			res = fmt.Sprintf("%f", inferedInstance.Y1)
		case "X2":
			res = fmt.Sprintf("%f", inferedInstance.X2)
		case "Y2":
			res = fmt.Sprintf("%f", inferedInstance.Y2)
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
		case "InsideAngle":
			res = fmt.Sprintf("%f", inferedInstance.InsideAngle)
		case "DiamondSideLenght":
			res = fmt.Sprintf("%f", inferedInstance.DiamondSideLenght)
		case "InitialRhombus":
			if inferedInstance.InitialRhombus != nil {
				res = inferedInstance.InitialRhombus.Name
			}
		case "InitialRhombusGrid":
			if inferedInstance.InitialRhombusGrid != nil {
				res = inferedInstance.InitialRhombusGrid.Name
			}
		case "InitialAxis":
			if inferedInstance.InitialAxis != nil {
				res = inferedInstance.InitialAxis.Name
			}
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
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
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
		case "N":
			res = fmt.Sprintf("%d", inferedInstance.N)
		case "M":
			res = fmt.Sprintf("%d", inferedInstance.M)
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "Rhombuses":
			for idx, __instance__ := range inferedInstance.Rhombuses {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *VerticalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case HorizontalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		}
	case InitialAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "Angle":
			res = fmt.Sprintf("%f", inferedInstance.Angle)
		case "Length":
			res = fmt.Sprintf("%f", inferedInstance.Length)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		}
	case Line:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X1":
			res = fmt.Sprintf("%f", inferedInstance.X1)
		case "Y1":
			res = fmt.Sprintf("%f", inferedInstance.Y1)
		case "X2":
			res = fmt.Sprintf("%f", inferedInstance.X2)
		case "Y2":
			res = fmt.Sprintf("%f", inferedInstance.Y2)
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
		case "InsideAngle":
			res = fmt.Sprintf("%f", inferedInstance.InsideAngle)
		case "DiamondSideLenght":
			res = fmt.Sprintf("%f", inferedInstance.DiamondSideLenght)
		case "InitialRhombus":
			if inferedInstance.InitialRhombus != nil {
				res = inferedInstance.InitialRhombus.Name
			}
		case "InitialRhombusGrid":
			if inferedInstance.InitialRhombusGrid != nil {
				res = inferedInstance.InitialRhombusGrid.Name
			}
		case "InitialAxis":
			if inferedInstance.InitialAxis != nil {
				res = inferedInstance.InitialAxis.Name
			}
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
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
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
		case "N":
			res = fmt.Sprintf("%d", inferedInstance.N)
		case "M":
			res = fmt.Sprintf("%d", inferedInstance.M)
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "Rhombuses":
			for idx, __instance__ := range inferedInstance.Rhombuses {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case VerticalAxis:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "AxisHandleBorderLength":
			res = fmt.Sprintf("%f", inferedInstance.AxisHandleBorderLength)
		case "Axis_Length":
			res = fmt.Sprintf("%f", inferedInstance.Axis_Length)
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
