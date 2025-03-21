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
	name string

	// insertion point for definition of arrays registering instances
	Cursors           map[*Cursor]any
	Cursors_mapString map[string]*Cursor

	// insertion point for slice of pointers maps
	OnAfterCursorCreateCallback OnAfterCreateInterface[Cursor]
	OnAfterCursorUpdateCallback OnAfterUpdateInterface[Cursor]
	OnAfterCursorDeleteCallback OnAfterDeleteInterface[Cursor]
	OnAfterCursorReadCallback   OnAfterReadInterface[Cursor]

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
	Order            uint
	Map_Staged_Order map[any]uint
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gong/lib/cursor/go/models"
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
	CommitCursor(cursor *Cursor)
	CheckoutCursor(cursor *Cursor)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Cursors:           make(map[*Cursor]any),
		Cursors_mapString: make(map[string]*Cursor),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		Map_Staged_Order: make(map[any]uint),
	}

	return
}

func (stage *StageStruct) GetName() string {
	return stage.name
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
	stage.Map_GongStructName_InstancesNb["Cursor"] = len(stage.Cursors)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Cursor"] = len(stage.Cursors)

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
// Stage puts cursor to the model stage
func (cursor *Cursor) Stage(stage *StageStruct) *Cursor {

	if _, ok := stage.Cursors[cursor]; !ok {
		stage.Cursors[cursor] = __member
		stage.Map_Staged_Order[cursor] = stage.Order
		stage.Order++
	}
	stage.Cursors_mapString[cursor.Name] = cursor

	return cursor
}

// Unstage removes cursor off the model stage
func (cursor *Cursor) Unstage(stage *StageStruct) *Cursor {
	delete(stage.Cursors, cursor)
	delete(stage.Cursors_mapString, cursor.Name)
	return cursor
}

// UnstageVoid removes cursor off the model stage
func (cursor *Cursor) UnstageVoid(stage *StageStruct) {
	delete(stage.Cursors, cursor)
	delete(stage.Cursors_mapString, cursor.Name)
}

// commit cursor to the back repo (if it is already staged)
func (cursor *Cursor) Commit(stage *StageStruct) *Cursor {
	if _, ok := stage.Cursors[cursor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCursor(cursor)
		}
	}
	return cursor
}

func (cursor *Cursor) CommitVoid(stage *StageStruct) {
	cursor.Commit(stage)
}

// Checkout cursor to the back repo (if it is already staged)
func (cursor *Cursor) Checkout(stage *StageStruct) *Cursor {
	if _, ok := stage.Cursors[cursor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCursor(cursor)
		}
	}
	return cursor
}

// for satisfaction of GongStruct interface
func (cursor *Cursor) GetName() (res string) {
	return cursor.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCursor(Cursor *Cursor)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCursor(Cursor *Cursor)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Cursors = make(map[*Cursor]any)
	stage.Cursors_mapString = make(map[string]*Cursor)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Cursors = nil
	stage.Cursors_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for cursor := range stage.Cursors {
		cursor.Unstage(stage)
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
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
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

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

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
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Cursor]any:
		return any(&stage.Cursors).(*Type)
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
	case map[string]*Cursor:
		return any(&stage.Cursors_mapString).(*Type)
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
	case Cursor:
		return any(&stage.Cursors).(*map[*Type]any)
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
	case *Cursor:
		return any(&stage.Cursors).(*map[Type]any)
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
	case Cursor:
		return any(&stage.Cursors_mapString).(*map[string]*Type)
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
	case Cursor:
		return any(&Cursor{
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
	// reverse maps of direct associations of Cursor
	case Cursor:
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
	// reverse maps of direct associations of Cursor
	case Cursor:
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
	case Cursor:
		res = "Cursor"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Cursor:
		res = "Cursor"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Cursor:
		res = []string{"Name", "StartX", "EndX", "Y1", "Y2", "DurationSeconds", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "IsPlaying"}
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
	case Cursor:
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
	case *Cursor:
		res = []string{"Name", "StartX", "EndX", "Y1", "Y2", "DurationSeconds", "Color", "FillOpacity", "Stroke", "StrokeOpacity", "StrokeWidth", "StrokeDashArray", "StrokeDashArrayWhenSelected", "Transform", "IsPlaying"}
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
	case *Cursor:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y1)
			res.valueFloat = inferedInstance.Y1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y2)
			res.valueFloat = inferedInstance.Y2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "DurationSeconds":
			res.valueString = fmt.Sprintf("%f", inferedInstance.DurationSeconds)
			res.valueFloat = inferedInstance.DurationSeconds
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
		case "IsPlaying":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsPlaying)
			res.valueBool = inferedInstance.IsPlaying
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Cursor:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StartX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartX)
			res.valueFloat = inferedInstance.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndX)
			res.valueFloat = inferedInstance.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y1":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y1)
			res.valueFloat = inferedInstance.Y1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y2)
			res.valueFloat = inferedInstance.Y2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "DurationSeconds":
			res.valueString = fmt.Sprintf("%f", inferedInstance.DurationSeconds)
			res.valueFloat = inferedInstance.DurationSeconds
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
		case "IsPlaying":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsPlaying)
			res.valueBool = inferedInstance.IsPlaying
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
