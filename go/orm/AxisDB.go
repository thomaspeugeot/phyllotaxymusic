// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/thomaspeugeot/phyllotaxymusic/go/db"
	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Axis_sql sql.NullBool
var dummy_Axis_time time.Duration
var dummy_Axis_sort sort.Float64Slice

// AxisAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model axisAPI
type AxisAPI struct {
	gorm.Model

	models.Axis_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	AxisPointersEncoding AxisPointersEncoding
}

// AxisPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type AxisPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ShapeCategory is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ShapeCategoryID sql.NullInt64
}

// AxisDB describes a axis in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model axisDB
type AxisDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field axisDB.Name
	Name_Data sql.NullString

	// Declation for basic field axisDB.IsDisplayed
	// provide the sql storage for the boolan
	IsDisplayed_Data sql.NullBool

	// Declation for basic field axisDB.AngleDegree
	AngleDegree_Data sql.NullFloat64

	// Declation for basic field axisDB.Length
	Length_Data sql.NullFloat64

	// Declation for basic field axisDB.CenterX
	CenterX_Data sql.NullFloat64

	// Declation for basic field axisDB.CenterY
	CenterY_Data sql.NullFloat64

	// Declation for basic field axisDB.EndX
	EndX_Data sql.NullFloat64

	// Declation for basic field axisDB.EndY
	EndY_Data sql.NullFloat64

	// Declation for basic field axisDB.Color
	Color_Data sql.NullString

	// Declation for basic field axisDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field axisDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field axisDB.StrokeOpacity
	StrokeOpacity_Data sql.NullFloat64

	// Declation for basic field axisDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field axisDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field axisDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field axisDB.Transform
	Transform_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	AxisPointersEncoding
}

// AxisDBs arrays axisDBs
// swagger:response axisDBsResponse
type AxisDBs []AxisDB

// AxisDBResponse provides response
// swagger:response axisDBResponse
type AxisDBResponse struct {
	AxisDB
}

// AxisWOP is a Axis without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type AxisWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsDisplayed bool `xlsx:"2"`

	AngleDegree float64 `xlsx:"3"`

	Length float64 `xlsx:"4"`

	CenterX float64 `xlsx:"5"`

	CenterY float64 `xlsx:"6"`

	EndX float64 `xlsx:"7"`

	EndY float64 `xlsx:"8"`

	Color string `xlsx:"9"`

	FillOpacity float64 `xlsx:"10"`

	Stroke string `xlsx:"11"`

	StrokeOpacity float64 `xlsx:"12"`

	StrokeWidth float64 `xlsx:"13"`

	StrokeDashArray string `xlsx:"14"`

	StrokeDashArrayWhenSelected string `xlsx:"15"`

	Transform string `xlsx:"16"`
	// insertion for WOP pointer fields
}

var Axis_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsDisplayed",
	"AngleDegree",
	"Length",
	"CenterX",
	"CenterY",
	"EndX",
	"EndY",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeOpacity",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
}

type BackRepoAxisStruct struct {
	// stores AxisDB according to their gorm ID
	Map_AxisDBID_AxisDB map[uint]*AxisDB

	// stores AxisDB ID according to Axis address
	Map_AxisPtr_AxisDBID map[*models.Axis]uint

	// stores Axis according to their gorm ID
	Map_AxisDBID_AxisPtr map[uint]*models.Axis

	db db.DBInterface

	stage *models.Stage
}

func (backRepoAxis *BackRepoAxisStruct) GetStage() (stage *models.Stage) {
	stage = backRepoAxis.stage
	return
}

func (backRepoAxis *BackRepoAxisStruct) GetDB() db.DBInterface {
	return backRepoAxis.db
}

// GetAxisDBFromAxisPtr is a handy function to access the back repo instance from the stage instance
func (backRepoAxis *BackRepoAxisStruct) GetAxisDBFromAxisPtr(axis *models.Axis) (axisDB *AxisDB) {
	id := backRepoAxis.Map_AxisPtr_AxisDBID[axis]
	axisDB = backRepoAxis.Map_AxisDBID_AxisDB[id]
	return
}

// BackRepoAxis.CommitPhaseOne commits all staged instances of Axis to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAxis *BackRepoAxisStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var axiss []*models.Axis
	for axis := range stage.Axiss {
		axiss = append(axiss, axis)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(axiss, func(i, j int) bool {
		return stage.AxisMap_Staged_Order[axiss[i]] < stage.AxisMap_Staged_Order[axiss[j]]
	})

	for _, axis := range axiss {
		backRepoAxis.CommitPhaseOneInstance(axis)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, axis := range backRepoAxis.Map_AxisDBID_AxisPtr {
		if _, ok := stage.Axiss[axis]; !ok {
			backRepoAxis.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoAxis.CommitDeleteInstance commits deletion of Axis to the BackRepo
func (backRepoAxis *BackRepoAxisStruct) CommitDeleteInstance(id uint) (Error error) {

	axis := backRepoAxis.Map_AxisDBID_AxisPtr[id]

	// axis is not staged anymore, remove axisDB
	axisDB := backRepoAxis.Map_AxisDBID_AxisDB[id]
	db, _ := backRepoAxis.db.Unscoped()
	_, err := db.Delete(axisDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoAxis.Map_AxisPtr_AxisDBID, axis)
	delete(backRepoAxis.Map_AxisDBID_AxisPtr, id)
	delete(backRepoAxis.Map_AxisDBID_AxisDB, id)

	return
}

// BackRepoAxis.CommitPhaseOneInstance commits axis staged instances of Axis to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAxis *BackRepoAxisStruct) CommitPhaseOneInstance(axis *models.Axis) (Error error) {

	// check if the axis is not commited yet
	if _, ok := backRepoAxis.Map_AxisPtr_AxisDBID[axis]; ok {
		return
	}

	// initiate axis
	var axisDB AxisDB
	axisDB.CopyBasicFieldsFromAxis(axis)

	_, err := backRepoAxis.db.Create(&axisDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoAxis.Map_AxisPtr_AxisDBID[axis] = axisDB.ID
	backRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID] = axis
	backRepoAxis.Map_AxisDBID_AxisDB[axisDB.ID] = &axisDB

	return
}

// BackRepoAxis.CommitPhaseTwo commits all staged instances of Axis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAxis *BackRepoAxisStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, axis := range backRepoAxis.Map_AxisDBID_AxisPtr {
		backRepoAxis.CommitPhaseTwoInstance(backRepo, idx, axis)
	}

	return
}

// BackRepoAxis.CommitPhaseTwoInstance commits {{structname }} of models.Axis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAxis *BackRepoAxisStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, axis *models.Axis) (Error error) {

	// fetch matching axisDB
	if axisDB, ok := backRepoAxis.Map_AxisDBID_AxisDB[idx]; ok {

		axisDB.CopyBasicFieldsFromAxis(axis)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value axis.ShapeCategory translates to updating the axis.ShapeCategoryID
		axisDB.ShapeCategoryID.Valid = true // allow for a 0 value (nil association)
		if axis.ShapeCategory != nil {
			if ShapeCategoryId, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[axis.ShapeCategory]; ok {
				axisDB.ShapeCategoryID.Int64 = int64(ShapeCategoryId)
				axisDB.ShapeCategoryID.Valid = true
			}
		} else {
			axisDB.ShapeCategoryID.Int64 = 0
			axisDB.ShapeCategoryID.Valid = true
		}

		_, err := backRepoAxis.db.Save(axisDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Axis intance %s", axis.Name))
		return err
	}

	return
}

// BackRepoAxis.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoAxis *BackRepoAxisStruct) CheckoutPhaseOne() (Error error) {

	axisDBArray := make([]AxisDB, 0)
	_, err := backRepoAxis.db.Find(&axisDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	axisInstancesToBeRemovedFromTheStage := make(map[*models.Axis]any)
	for key, value := range backRepoAxis.stage.Axiss {
		axisInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, axisDB := range axisDBArray {
		backRepoAxis.CheckoutPhaseOneInstance(&axisDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		axis, ok := backRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID]
		if ok {
			delete(axisInstancesToBeRemovedFromTheStage, axis)
		}
	}

	// remove from stage and back repo's 3 maps all axiss that are not in the checkout
	for axis := range axisInstancesToBeRemovedFromTheStage {
		axis.Unstage(backRepoAxis.GetStage())

		// remove instance from the back repo 3 maps
		axisID := backRepoAxis.Map_AxisPtr_AxisDBID[axis]
		delete(backRepoAxis.Map_AxisPtr_AxisDBID, axis)
		delete(backRepoAxis.Map_AxisDBID_AxisDB, axisID)
		delete(backRepoAxis.Map_AxisDBID_AxisPtr, axisID)
	}

	return
}

// CheckoutPhaseOneInstance takes a axisDB that has been found in the DB, updates the backRepo and stages the
// models version of the axisDB
func (backRepoAxis *BackRepoAxisStruct) CheckoutPhaseOneInstance(axisDB *AxisDB) (Error error) {

	axis, ok := backRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID]
	if !ok {
		axis = new(models.Axis)

		backRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID] = axis
		backRepoAxis.Map_AxisPtr_AxisDBID[axis] = axisDB.ID

		// append model store with the new element
		axis.Name = axisDB.Name_Data.String
		axis.Stage(backRepoAxis.GetStage())
	}
	axisDB.CopyBasicFieldsToAxis(axis)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	axis.Stage(backRepoAxis.GetStage())

	// preserve pointer to axisDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_AxisDBID_AxisDB)[axisDB hold variable pointers
	axisDB_Data := *axisDB
	preservedPtrToAxis := &axisDB_Data
	backRepoAxis.Map_AxisDBID_AxisDB[axisDB.ID] = preservedPtrToAxis

	return
}

// BackRepoAxis.CheckoutPhaseTwo Checkouts all staged instances of Axis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAxis *BackRepoAxisStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, axisDB := range backRepoAxis.Map_AxisDBID_AxisDB {
		backRepoAxis.CheckoutPhaseTwoInstance(backRepo, axisDB)
	}
	return
}

// BackRepoAxis.CheckoutPhaseTwoInstance Checkouts staged instances of Axis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAxis *BackRepoAxisStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, axisDB *AxisDB) (Error error) {

	axis := backRepoAxis.Map_AxisDBID_AxisPtr[axisDB.ID]

	axisDB.DecodePointers(backRepo, axis)

	return
}

func (axisDB *AxisDB) DecodePointers(backRepo *BackRepoStruct, axis *models.Axis) {

	// insertion point for checkout of pointer encoding
	// ShapeCategory field	
	{
		id := axisDB.ShapeCategoryID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: axis.ShapeCategory, unknown pointer id", id)
				axis.ShapeCategory = nil
			} else {
				// updates only if field has changed
				if axis.ShapeCategory == nil || axis.ShapeCategory != tmp {
					axis.ShapeCategory = tmp
				}
			}
		} else {
			axis.ShapeCategory = nil
		}
	}
	
	return
}

// CommitAxis allows commit of a single axis (if already staged)
func (backRepo *BackRepoStruct) CommitAxis(axis *models.Axis) {
	backRepo.BackRepoAxis.CommitPhaseOneInstance(axis)
	if id, ok := backRepo.BackRepoAxis.Map_AxisPtr_AxisDBID[axis]; ok {
		backRepo.BackRepoAxis.CommitPhaseTwoInstance(backRepo, id, axis)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitAxis allows checkout of a single axis (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutAxis(axis *models.Axis) {
	// check if the axis is staged
	if _, ok := backRepo.BackRepoAxis.Map_AxisPtr_AxisDBID[axis]; ok {

		if id, ok := backRepo.BackRepoAxis.Map_AxisPtr_AxisDBID[axis]; ok {
			var axisDB AxisDB
			axisDB.ID = id

			if _, err := backRepo.BackRepoAxis.db.First(&axisDB, id); err != nil {
				log.Fatalln("CheckoutAxis : Problem with getting object with id:", id)
			}
			backRepo.BackRepoAxis.CheckoutPhaseOneInstance(&axisDB)
			backRepo.BackRepoAxis.CheckoutPhaseTwoInstance(backRepo, &axisDB)
		}
	}
}

// CopyBasicFieldsFromAxis
func (axisDB *AxisDB) CopyBasicFieldsFromAxis(axis *models.Axis) {
	// insertion point for fields commit

	axisDB.Name_Data.String = axis.Name
	axisDB.Name_Data.Valid = true

	axisDB.IsDisplayed_Data.Bool = axis.IsDisplayed
	axisDB.IsDisplayed_Data.Valid = true

	axisDB.AngleDegree_Data.Float64 = axis.AngleDegree
	axisDB.AngleDegree_Data.Valid = true

	axisDB.Length_Data.Float64 = axis.Length
	axisDB.Length_Data.Valid = true

	axisDB.CenterX_Data.Float64 = axis.CenterX
	axisDB.CenterX_Data.Valid = true

	axisDB.CenterY_Data.Float64 = axis.CenterY
	axisDB.CenterY_Data.Valid = true

	axisDB.EndX_Data.Float64 = axis.EndX
	axisDB.EndX_Data.Valid = true

	axisDB.EndY_Data.Float64 = axis.EndY
	axisDB.EndY_Data.Valid = true

	axisDB.Color_Data.String = axis.Color
	axisDB.Color_Data.Valid = true

	axisDB.FillOpacity_Data.Float64 = axis.FillOpacity
	axisDB.FillOpacity_Data.Valid = true

	axisDB.Stroke_Data.String = axis.Stroke
	axisDB.Stroke_Data.Valid = true

	axisDB.StrokeOpacity_Data.Float64 = axis.StrokeOpacity
	axisDB.StrokeOpacity_Data.Valid = true

	axisDB.StrokeWidth_Data.Float64 = axis.StrokeWidth
	axisDB.StrokeWidth_Data.Valid = true

	axisDB.StrokeDashArray_Data.String = axis.StrokeDashArray
	axisDB.StrokeDashArray_Data.Valid = true

	axisDB.StrokeDashArrayWhenSelected_Data.String = axis.StrokeDashArrayWhenSelected
	axisDB.StrokeDashArrayWhenSelected_Data.Valid = true

	axisDB.Transform_Data.String = axis.Transform
	axisDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromAxis_WOP
func (axisDB *AxisDB) CopyBasicFieldsFromAxis_WOP(axis *models.Axis_WOP) {
	// insertion point for fields commit

	axisDB.Name_Data.String = axis.Name
	axisDB.Name_Data.Valid = true

	axisDB.IsDisplayed_Data.Bool = axis.IsDisplayed
	axisDB.IsDisplayed_Data.Valid = true

	axisDB.AngleDegree_Data.Float64 = axis.AngleDegree
	axisDB.AngleDegree_Data.Valid = true

	axisDB.Length_Data.Float64 = axis.Length
	axisDB.Length_Data.Valid = true

	axisDB.CenterX_Data.Float64 = axis.CenterX
	axisDB.CenterX_Data.Valid = true

	axisDB.CenterY_Data.Float64 = axis.CenterY
	axisDB.CenterY_Data.Valid = true

	axisDB.EndX_Data.Float64 = axis.EndX
	axisDB.EndX_Data.Valid = true

	axisDB.EndY_Data.Float64 = axis.EndY
	axisDB.EndY_Data.Valid = true

	axisDB.Color_Data.String = axis.Color
	axisDB.Color_Data.Valid = true

	axisDB.FillOpacity_Data.Float64 = axis.FillOpacity
	axisDB.FillOpacity_Data.Valid = true

	axisDB.Stroke_Data.String = axis.Stroke
	axisDB.Stroke_Data.Valid = true

	axisDB.StrokeOpacity_Data.Float64 = axis.StrokeOpacity
	axisDB.StrokeOpacity_Data.Valid = true

	axisDB.StrokeWidth_Data.Float64 = axis.StrokeWidth
	axisDB.StrokeWidth_Data.Valid = true

	axisDB.StrokeDashArray_Data.String = axis.StrokeDashArray
	axisDB.StrokeDashArray_Data.Valid = true

	axisDB.StrokeDashArrayWhenSelected_Data.String = axis.StrokeDashArrayWhenSelected
	axisDB.StrokeDashArrayWhenSelected_Data.Valid = true

	axisDB.Transform_Data.String = axis.Transform
	axisDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromAxisWOP
func (axisDB *AxisDB) CopyBasicFieldsFromAxisWOP(axis *AxisWOP) {
	// insertion point for fields commit

	axisDB.Name_Data.String = axis.Name
	axisDB.Name_Data.Valid = true

	axisDB.IsDisplayed_Data.Bool = axis.IsDisplayed
	axisDB.IsDisplayed_Data.Valid = true

	axisDB.AngleDegree_Data.Float64 = axis.AngleDegree
	axisDB.AngleDegree_Data.Valid = true

	axisDB.Length_Data.Float64 = axis.Length
	axisDB.Length_Data.Valid = true

	axisDB.CenterX_Data.Float64 = axis.CenterX
	axisDB.CenterX_Data.Valid = true

	axisDB.CenterY_Data.Float64 = axis.CenterY
	axisDB.CenterY_Data.Valid = true

	axisDB.EndX_Data.Float64 = axis.EndX
	axisDB.EndX_Data.Valid = true

	axisDB.EndY_Data.Float64 = axis.EndY
	axisDB.EndY_Data.Valid = true

	axisDB.Color_Data.String = axis.Color
	axisDB.Color_Data.Valid = true

	axisDB.FillOpacity_Data.Float64 = axis.FillOpacity
	axisDB.FillOpacity_Data.Valid = true

	axisDB.Stroke_Data.String = axis.Stroke
	axisDB.Stroke_Data.Valid = true

	axisDB.StrokeOpacity_Data.Float64 = axis.StrokeOpacity
	axisDB.StrokeOpacity_Data.Valid = true

	axisDB.StrokeWidth_Data.Float64 = axis.StrokeWidth
	axisDB.StrokeWidth_Data.Valid = true

	axisDB.StrokeDashArray_Data.String = axis.StrokeDashArray
	axisDB.StrokeDashArray_Data.Valid = true

	axisDB.StrokeDashArrayWhenSelected_Data.String = axis.StrokeDashArrayWhenSelected
	axisDB.StrokeDashArrayWhenSelected_Data.Valid = true

	axisDB.Transform_Data.String = axis.Transform
	axisDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToAxis
func (axisDB *AxisDB) CopyBasicFieldsToAxis(axis *models.Axis) {
	// insertion point for checkout of basic fields (back repo to stage)
	axis.Name = axisDB.Name_Data.String
	axis.IsDisplayed = axisDB.IsDisplayed_Data.Bool
	axis.AngleDegree = axisDB.AngleDegree_Data.Float64
	axis.Length = axisDB.Length_Data.Float64
	axis.CenterX = axisDB.CenterX_Data.Float64
	axis.CenterY = axisDB.CenterY_Data.Float64
	axis.EndX = axisDB.EndX_Data.Float64
	axis.EndY = axisDB.EndY_Data.Float64
	axis.Color = axisDB.Color_Data.String
	axis.FillOpacity = axisDB.FillOpacity_Data.Float64
	axis.Stroke = axisDB.Stroke_Data.String
	axis.StrokeOpacity = axisDB.StrokeOpacity_Data.Float64
	axis.StrokeWidth = axisDB.StrokeWidth_Data.Float64
	axis.StrokeDashArray = axisDB.StrokeDashArray_Data.String
	axis.StrokeDashArrayWhenSelected = axisDB.StrokeDashArrayWhenSelected_Data.String
	axis.Transform = axisDB.Transform_Data.String
}

// CopyBasicFieldsToAxis_WOP
func (axisDB *AxisDB) CopyBasicFieldsToAxis_WOP(axis *models.Axis_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	axis.Name = axisDB.Name_Data.String
	axis.IsDisplayed = axisDB.IsDisplayed_Data.Bool
	axis.AngleDegree = axisDB.AngleDegree_Data.Float64
	axis.Length = axisDB.Length_Data.Float64
	axis.CenterX = axisDB.CenterX_Data.Float64
	axis.CenterY = axisDB.CenterY_Data.Float64
	axis.EndX = axisDB.EndX_Data.Float64
	axis.EndY = axisDB.EndY_Data.Float64
	axis.Color = axisDB.Color_Data.String
	axis.FillOpacity = axisDB.FillOpacity_Data.Float64
	axis.Stroke = axisDB.Stroke_Data.String
	axis.StrokeOpacity = axisDB.StrokeOpacity_Data.Float64
	axis.StrokeWidth = axisDB.StrokeWidth_Data.Float64
	axis.StrokeDashArray = axisDB.StrokeDashArray_Data.String
	axis.StrokeDashArrayWhenSelected = axisDB.StrokeDashArrayWhenSelected_Data.String
	axis.Transform = axisDB.Transform_Data.String
}

// CopyBasicFieldsToAxisWOP
func (axisDB *AxisDB) CopyBasicFieldsToAxisWOP(axis *AxisWOP) {
	axis.ID = int(axisDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	axis.Name = axisDB.Name_Data.String
	axis.IsDisplayed = axisDB.IsDisplayed_Data.Bool
	axis.AngleDegree = axisDB.AngleDegree_Data.Float64
	axis.Length = axisDB.Length_Data.Float64
	axis.CenterX = axisDB.CenterX_Data.Float64
	axis.CenterY = axisDB.CenterY_Data.Float64
	axis.EndX = axisDB.EndX_Data.Float64
	axis.EndY = axisDB.EndY_Data.Float64
	axis.Color = axisDB.Color_Data.String
	axis.FillOpacity = axisDB.FillOpacity_Data.Float64
	axis.Stroke = axisDB.Stroke_Data.String
	axis.StrokeOpacity = axisDB.StrokeOpacity_Data.Float64
	axis.StrokeWidth = axisDB.StrokeWidth_Data.Float64
	axis.StrokeDashArray = axisDB.StrokeDashArray_Data.String
	axis.StrokeDashArrayWhenSelected = axisDB.StrokeDashArrayWhenSelected_Data.String
	axis.Transform = axisDB.Transform_Data.String
}

// Backup generates a json file from a slice of all AxisDB instances in the backrepo
func (backRepoAxis *BackRepoAxisStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "AxisDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AxisDB, 0)
	for _, axisDB := range backRepoAxis.Map_AxisDBID_AxisDB {
		forBackup = append(forBackup, axisDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Axis ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Axis file", err.Error())
	}
}

// Backup generates a json file from a slice of all AxisDB instances in the backrepo
func (backRepoAxis *BackRepoAxisStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AxisDB, 0)
	for _, axisDB := range backRepoAxis.Map_AxisDBID_AxisDB {
		forBackup = append(forBackup, axisDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Axis")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Axis_Fields, -1)
	for _, axisDB := range forBackup {

		var axisWOP AxisWOP
		axisDB.CopyBasicFieldsToAxisWOP(&axisWOP)

		row := sh.AddRow()
		row.WriteStruct(&axisWOP, -1)
	}
}

// RestoreXL from the "Axis" sheet all AxisDB instances
func (backRepoAxis *BackRepoAxisStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoAxisid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Axis"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoAxis.rowVisitorAxis)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoAxis *BackRepoAxisStruct) rowVisitorAxis(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var axisWOP AxisWOP
		row.ReadStruct(&axisWOP)

		// add the unmarshalled struct to the stage
		axisDB := new(AxisDB)
		axisDB.CopyBasicFieldsFromAxisWOP(&axisWOP)

		axisDB_ID_atBackupTime := axisDB.ID
		axisDB.ID = 0
		_, err := backRepoAxis.db.Create(axisDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoAxis.Map_AxisDBID_AxisDB[axisDB.ID] = axisDB
		BackRepoAxisid_atBckpTime_newID[axisDB_ID_atBackupTime] = axisDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "AxisDB.json" in dirPath that stores an array
// of AxisDB and stores it in the database
// the map BackRepoAxisid_atBckpTime_newID is updated accordingly
func (backRepoAxis *BackRepoAxisStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoAxisid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "AxisDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Axis file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*AxisDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_AxisDBID_AxisDB
	for _, axisDB := range forRestore {

		axisDB_ID_atBackupTime := axisDB.ID
		axisDB.ID = 0
		_, err := backRepoAxis.db.Create(axisDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoAxis.Map_AxisDBID_AxisDB[axisDB.ID] = axisDB
		BackRepoAxisid_atBckpTime_newID[axisDB_ID_atBackupTime] = axisDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Axis file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Axis>id_atBckpTime_newID
// to compute new index
func (backRepoAxis *BackRepoAxisStruct) RestorePhaseTwo() {

	for _, axisDB := range backRepoAxis.Map_AxisDBID_AxisDB {

		// next line of code is to avert unused variable compilation error
		_ = axisDB

		// insertion point for reindexing pointers encoding
		// reindexing ShapeCategory field
		if axisDB.ShapeCategoryID.Int64 != 0 {
			axisDB.ShapeCategoryID.Int64 = int64(BackRepoShapeCategoryid_atBckpTime_newID[uint(axisDB.ShapeCategoryID.Int64)])
			axisDB.ShapeCategoryID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoAxis.db.Model(axisDB)
		_, err := db.Updates(*axisDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoAxis.ResetReversePointers commits all staged instances of Axis to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAxis *BackRepoAxisStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, axis := range backRepoAxis.Map_AxisDBID_AxisPtr {
		backRepoAxis.ResetReversePointersInstance(backRepo, idx, axis)
	}

	return
}

func (backRepoAxis *BackRepoAxisStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, axis *models.Axis) (Error error) {

	// fetch matching axisDB
	if axisDB, ok := backRepoAxis.Map_AxisDBID_AxisDB[idx]; ok {
		_ = axisDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoAxisid_atBckpTime_newID map[uint]uint
