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
	CommitLine(line *Line)
	CheckoutLine(line *Line)
	CommitParameter(parameter *Parameter)
	CheckoutParameter(parameter *Parameter)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Lines:           make(map[*Line]any),
		Lines_mapString: make(map[string]*Line),

		Parameters:           make(map[*Parameter]any),
		Parameters_mapString: make(map[string]*Parameter),

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
	stage.Map_GongStructName_InstancesNb["Line"] = len(stage.Lines)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Line"] = len(stage.Lines)
	stage.Map_GongStructName_InstancesNb["Parameter"] = len(stage.Parameters)

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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMLine(Line *Line)
	CreateORMParameter(Parameter *Parameter)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMLine(Line *Line)
	DeleteORMParameter(Parameter *Parameter)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Lines = make(map[*Line]any)
	stage.Lines_mapString = make(map[string]*Line)

	stage.Parameters = make(map[*Parameter]any)
	stage.Parameters_mapString = make(map[string]*Parameter)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Lines = nil
	stage.Lines_mapString = nil

	stage.Parameters = nil
	stage.Parameters_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for line := range stage.Lines {
		line.Unstage(stage)
	}

	for parameter := range stage.Parameters {
		parameter.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Line | Parameter
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
	*Line | *Parameter
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
		map[*Line]any |
		map[*Parameter]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Line |
		map[string]*Parameter |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Line]any:
		return any(&stage.Lines).(*Type)
	case map[*Parameter]any:
		return any(&stage.Parameters).(*Type)
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
	case map[string]*Line:
		return any(&stage.Lines_mapString).(*Type)
	case map[string]*Parameter:
		return any(&stage.Parameters_mapString).(*Type)
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
	case Line:
		return any(&stage.Lines).(*map[*Type]any)
	case Parameter:
		return any(&stage.Parameters).(*map[*Type]any)
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
	case *Line:
		return any(&stage.Lines).(*map[Type]any)
	case *Parameter:
		return any(&stage.Parameters).(*map[Type]any)
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
	case Line:
		return any(&stage.Lines_mapString).(*map[string]*Type)
	case Parameter:
		return any(&stage.Parameters_mapString).(*map[string]*Type)
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
	case Line:
		return any(&Line{
			// Initialisation of associations
		}).(*Type)
	case Parameter:
		return any(&Parameter{
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
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Line:
		res = "Line"
	case Parameter:
		res = "Parameter"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Line:
		res = "Line"
	case *Parameter:
		res = "Parameter"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Line:
		res = []string{"Name", "X1", "Y1", "X2", "Y2"}
	case Parameter:
		res = []string{"Name", "N", "M", "DiamondAngle", "OriginX", "OriginY", "DiamondSideLenght"}
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
	case Line:
		var rf ReverseField
		_ = rf
	case Parameter:
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
	case *Line:
		res = []string{"Name", "X1", "Y1", "X2", "Y2"}
	case *Parameter:
		res = []string{"Name", "N", "M", "DiamondAngle", "OriginX", "OriginY", "DiamondSideLenght"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
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
		case "DiamondAngle":
			res = fmt.Sprintf("%f", inferedInstance.DiamondAngle)
		case "OriginX":
			res = fmt.Sprintf("%f", inferedInstance.OriginX)
		case "OriginY":
			res = fmt.Sprintf("%f", inferedInstance.OriginY)
		case "DiamondSideLenght":
			res = fmt.Sprintf("%f", inferedInstance.DiamondSideLenght)
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
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
		case "DiamondAngle":
			res = fmt.Sprintf("%f", inferedInstance.DiamondAngle)
		case "OriginX":
			res = fmt.Sprintf("%f", inferedInstance.OriginX)
		case "OriginY":
			res = fmt.Sprintf("%f", inferedInstance.OriginY)
		case "DiamondSideLenght":
			res = fmt.Sprintf("%f", inferedInstance.DiamondSideLenght)
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
