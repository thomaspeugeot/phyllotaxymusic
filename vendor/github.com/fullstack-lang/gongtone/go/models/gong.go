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
	Freqencys           map[*Freqency]any
	Freqencys_mapString map[string]*Freqency

	// insertion point for slice of pointers maps

	OnAfterFreqencyCreateCallback OnAfterCreateInterface[Freqency]
	OnAfterFreqencyUpdateCallback OnAfterUpdateInterface[Freqency]
	OnAfterFreqencyDeleteCallback OnAfterDeleteInterface[Freqency]
	OnAfterFreqencyReadCallback   OnAfterReadInterface[Freqency]

	Notes           map[*Note]any
	Notes_mapString map[string]*Note

	// insertion point for slice of pointers maps
	Note_Frequencies_reverseMap map[*Freqency]*Note

	OnAfterNoteCreateCallback OnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback OnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback OnAfterDeleteInterface[Note]
	OnAfterNoteReadCallback   OnAfterReadInterface[Note]

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
	return "github.com/fullstack-lang/gongtone/go/models"
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
	CommitFreqency(freqency *Freqency)
	CheckoutFreqency(freqency *Freqency)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Freqencys:           make(map[*Freqency]any),
		Freqencys_mapString: make(map[string]*Freqency),

		Notes:           make(map[*Note]any),
		Notes_mapString: make(map[string]*Note),

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
	stage.Map_GongStructName_InstancesNb["Freqency"] = len(stage.Freqencys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Freqency"] = len(stage.Freqencys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)

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
// Stage puts freqency to the model stage
func (freqency *Freqency) Stage(stage *StageStruct) *Freqency {
	stage.Freqencys[freqency] = __member
	stage.Freqencys_mapString[freqency.Name] = freqency

	return freqency
}

// Unstage removes freqency off the model stage
func (freqency *Freqency) Unstage(stage *StageStruct) *Freqency {
	delete(stage.Freqencys, freqency)
	delete(stage.Freqencys_mapString, freqency.Name)
	return freqency
}

// UnstageVoid removes freqency off the model stage
func (freqency *Freqency) UnstageVoid(stage *StageStruct) {
	delete(stage.Freqencys, freqency)
	delete(stage.Freqencys_mapString, freqency.Name)
}

// commit freqency to the back repo (if it is already staged)
func (freqency *Freqency) Commit(stage *StageStruct) *Freqency {
	if _, ok := stage.Freqencys[freqency]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFreqency(freqency)
		}
	}
	return freqency
}

func (freqency *Freqency) CommitVoid(stage *StageStruct) {
	freqency.Commit(stage)
}

// Checkout freqency to the back repo (if it is already staged)
func (freqency *Freqency) Checkout(stage *StageStruct) *Freqency {
	if _, ok := stage.Freqencys[freqency]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFreqency(freqency)
		}
	}
	return freqency
}

// for satisfaction of GongStruct interface
func (freqency *Freqency) GetName() (res string) {
	return freqency.Name
}

// Stage puts note to the model stage
func (note *Note) Stage(stage *StageStruct) *Note {
	stage.Notes[note] = __member
	stage.Notes_mapString[note.Name] = note

	return note
}

// Unstage removes note off the model stage
func (note *Note) Unstage(stage *StageStruct) *Note {
	delete(stage.Notes, note)
	delete(stage.Notes_mapString, note.Name)
	return note
}

// UnstageVoid removes note off the model stage
func (note *Note) UnstageVoid(stage *StageStruct) {
	delete(stage.Notes, note)
	delete(stage.Notes_mapString, note.Name)
}

// commit note to the back repo (if it is already staged)
func (note *Note) Commit(stage *StageStruct) *Note {
	if _, ok := stage.Notes[note]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNote(note)
		}
	}
	return note
}

func (note *Note) CommitVoid(stage *StageStruct) {
	note.Commit(stage)
}

// Checkout note to the back repo (if it is already staged)
func (note *Note) Checkout(stage *StageStruct) *Note {
	if _, ok := stage.Notes[note]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNote(note)
		}
	}
	return note
}

// for satisfaction of GongStruct interface
func (note *Note) GetName() (res string) {
	return note.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMFreqency(Freqency *Freqency)
	CreateORMNote(Note *Note)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMFreqency(Freqency *Freqency)
	DeleteORMNote(Note *Note)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Freqencys = make(map[*Freqency]any)
	stage.Freqencys_mapString = make(map[string]*Freqency)

	stage.Notes = make(map[*Note]any)
	stage.Notes_mapString = make(map[string]*Note)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Freqencys = nil
	stage.Freqencys_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for freqency := range stage.Freqencys {
		freqency.Unstage(stage)
	}

	for note := range stage.Notes {
		note.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Freqency | Note
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
	*Freqency | *Note
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
		map[*Freqency]any |
		map[*Note]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Freqency |
		map[string]*Note |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Freqency]any:
		return any(&stage.Freqencys).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
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
	case map[string]*Freqency:
		return any(&stage.Freqencys_mapString).(*Type)
	case map[string]*Note:
		return any(&stage.Notes_mapString).(*Type)
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
	case Freqency:
		return any(&stage.Freqencys).(*map[*Type]any)
	case Note:
		return any(&stage.Notes).(*map[*Type]any)
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
	case *Freqency:
		return any(&stage.Freqencys).(*map[Type]any)
	case *Note:
		return any(&stage.Notes).(*map[Type]any)
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
	case Freqency:
		return any(&stage.Freqencys_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
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
	case Freqency:
		return any(&Freqency{
			// Initialisation of associations
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of Freqency with the name of the field
			Frequencies: []*Freqency{{Name: "Frequencies"}},
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
	// reverse maps of direct associations of Freqency
	case Freqency:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
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
	// reverse maps of direct associations of Freqency
	case Freqency:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Frequencies":
			res := make(map[*Freqency]*Note)
			for note := range stage.Notes {
				for _, freqency_ := range note.Frequencies {
					res[freqency_] = note
				}
			}
			return any(res).(map[*End]*Start)
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
	case Freqency:
		res = "Freqency"
	case Note:
		res = "Note"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Freqency:
		res = "Freqency"
	case *Note:
		res = "Note"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Freqency:
		res = []string{"Name"}
	case Note:
		res = []string{"Name", "Frequencies", "Start", "Duration", "Velocity", "Info"}
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
	case Freqency:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Note"
		rf.Fieldname = "Frequencies"
		res = append(res, rf)
	case Note:
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
	case *Freqency:
		res = []string{"Name"}
	case *Note:
		res = []string{"Name", "Frequencies", "Start", "Duration", "Velocity", "Info"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Freqency:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Note:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Frequencies":
			for idx, __instance__ := range inferedInstance.Frequencies {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Start":
			res = fmt.Sprintf("%f", inferedInstance.Start)
		case "Duration":
			res = fmt.Sprintf("%f", inferedInstance.Duration)
		case "Velocity":
			res = fmt.Sprintf("%f", inferedInstance.Velocity)
		case "Info":
			res = inferedInstance.Info
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Freqency:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Note:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Frequencies":
			for idx, __instance__ := range inferedInstance.Frequencies {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Start":
			res = fmt.Sprintf("%f", inferedInstance.Start)
		case "Duration":
			res = fmt.Sprintf("%f", inferedInstance.Duration)
		case "Velocity":
			res = fmt.Sprintf("%f", inferedInstance.Velocity)
		case "Info":
			res = inferedInstance.Info
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template