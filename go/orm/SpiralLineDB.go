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
var dummy_SpiralLine_sql sql.NullBool
var dummy_SpiralLine_time time.Duration
var dummy_SpiralLine_sort sort.Float64Slice

// SpiralLineAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model spirallineAPI
type SpiralLineAPI struct {
	gorm.Model

	models.SpiralLine_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SpiralLinePointersEncoding SpiralLinePointersEncoding
}

// SpiralLinePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SpiralLinePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ShapeCategory is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ShapeCategoryID sql.NullInt64
}

// SpiralLineDB describes a spiralline in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model spirallineDB
type SpiralLineDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field spirallineDB.Name
	Name_Data sql.NullString

	// Declation for basic field spirallineDB.IsDisplayed
	// provide the sql storage for the boolan
	IsDisplayed_Data sql.NullBool

	// Declation for basic field spirallineDB.StartX
	StartX_Data sql.NullFloat64

	// Declation for basic field spirallineDB.EndX
	EndX_Data sql.NullFloat64

	// Declation for basic field spirallineDB.StartY
	StartY_Data sql.NullFloat64

	// Declation for basic field spirallineDB.EndY
	EndY_Data sql.NullFloat64

	// Declation for basic field spirallineDB.Color
	Color_Data sql.NullString

	// Declation for basic field spirallineDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field spirallineDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field spirallineDB.StrokeOpacity
	StrokeOpacity_Data sql.NullFloat64

	// Declation for basic field spirallineDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field spirallineDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field spirallineDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field spirallineDB.Transform
	Transform_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SpiralLinePointersEncoding
}

// SpiralLineDBs arrays spirallineDBs
// swagger:response spirallineDBsResponse
type SpiralLineDBs []SpiralLineDB

// SpiralLineDBResponse provides response
// swagger:response spirallineDBResponse
type SpiralLineDBResponse struct {
	SpiralLineDB
}

// SpiralLineWOP is a SpiralLine without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SpiralLineWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsDisplayed bool `xlsx:"2"`

	StartX float64 `xlsx:"3"`

	EndX float64 `xlsx:"4"`

	StartY float64 `xlsx:"5"`

	EndY float64 `xlsx:"6"`

	Color string `xlsx:"7"`

	FillOpacity float64 `xlsx:"8"`

	Stroke string `xlsx:"9"`

	StrokeOpacity float64 `xlsx:"10"`

	StrokeWidth float64 `xlsx:"11"`

	StrokeDashArray string `xlsx:"12"`

	StrokeDashArrayWhenSelected string `xlsx:"13"`

	Transform string `xlsx:"14"`
	// insertion for WOP pointer fields
}

var SpiralLine_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsDisplayed",
	"StartX",
	"EndX",
	"StartY",
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

type BackRepoSpiralLineStruct struct {
	// stores SpiralLineDB according to their gorm ID
	Map_SpiralLineDBID_SpiralLineDB map[uint]*SpiralLineDB

	// stores SpiralLineDB ID according to SpiralLine address
	Map_SpiralLinePtr_SpiralLineDBID map[*models.SpiralLine]uint

	// stores SpiralLine according to their gorm ID
	Map_SpiralLineDBID_SpiralLinePtr map[uint]*models.SpiralLine

	db db.DBInterface

	stage *models.Stage
}

func (backRepoSpiralLine *BackRepoSpiralLineStruct) GetStage() (stage *models.Stage) {
	stage = backRepoSpiralLine.stage
	return
}

func (backRepoSpiralLine *BackRepoSpiralLineStruct) GetDB() db.DBInterface {
	return backRepoSpiralLine.db
}

// GetSpiralLineDBFromSpiralLinePtr is a handy function to access the back repo instance from the stage instance
func (backRepoSpiralLine *BackRepoSpiralLineStruct) GetSpiralLineDBFromSpiralLinePtr(spiralline *models.SpiralLine) (spirallineDB *SpiralLineDB) {
	id := backRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline]
	spirallineDB = backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[id]
	return
}

// BackRepoSpiralLine.CommitPhaseOne commits all staged instances of SpiralLine to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var spirallines []*models.SpiralLine
	for spiralline := range stage.SpiralLines {
		spirallines = append(spirallines, spiralline)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(spirallines, func(i, j int) bool {
		return stage.SpiralLineMap_Staged_Order[spirallines[i]] < stage.SpiralLineMap_Staged_Order[spirallines[j]]
	})

	for _, spiralline := range spirallines {
		backRepoSpiralLine.CommitPhaseOneInstance(spiralline)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, spiralline := range backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr {
		if _, ok := stage.SpiralLines[spiralline]; !ok {
			backRepoSpiralLine.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSpiralLine.CommitDeleteInstance commits deletion of SpiralLine to the BackRepo
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CommitDeleteInstance(id uint) (Error error) {

	spiralline := backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[id]

	// spiralline is not staged anymore, remove spirallineDB
	spirallineDB := backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[id]
	db, _ := backRepoSpiralLine.db.Unscoped()
	_, err := db.Delete(spirallineDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID, spiralline)
	delete(backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr, id)
	delete(backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB, id)

	return
}

// BackRepoSpiralLine.CommitPhaseOneInstance commits spiralline staged instances of SpiralLine to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CommitPhaseOneInstance(spiralline *models.SpiralLine) (Error error) {

	// check if the spiralline is not commited yet
	if _, ok := backRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline]; ok {
		return
	}

	// initiate spiralline
	var spirallineDB SpiralLineDB
	spirallineDB.CopyBasicFieldsFromSpiralLine(spiralline)

	_, err := backRepoSpiralLine.db.Create(&spirallineDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline] = spirallineDB.ID
	backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID] = spiralline
	backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[spirallineDB.ID] = &spirallineDB

	return
}

// BackRepoSpiralLine.CommitPhaseTwo commits all staged instances of SpiralLine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralline := range backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr {
		backRepoSpiralLine.CommitPhaseTwoInstance(backRepo, idx, spiralline)
	}

	return
}

// BackRepoSpiralLine.CommitPhaseTwoInstance commits {{structname }} of models.SpiralLine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, spiralline *models.SpiralLine) (Error error) {

	// fetch matching spirallineDB
	if spirallineDB, ok := backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[idx]; ok {

		spirallineDB.CopyBasicFieldsFromSpiralLine(spiralline)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value spiralline.ShapeCategory translates to updating the spiralline.ShapeCategoryID
		spirallineDB.ShapeCategoryID.Valid = true // allow for a 0 value (nil association)
		if spiralline.ShapeCategory != nil {
			if ShapeCategoryId, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[spiralline.ShapeCategory]; ok {
				spirallineDB.ShapeCategoryID.Int64 = int64(ShapeCategoryId)
				spirallineDB.ShapeCategoryID.Valid = true
			}
		} else {
			spirallineDB.ShapeCategoryID.Int64 = 0
			spirallineDB.ShapeCategoryID.Valid = true
		}

		_, err := backRepoSpiralLine.db.Save(spirallineDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SpiralLine intance %s", spiralline.Name))
		return err
	}

	return
}

// BackRepoSpiralLine.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CheckoutPhaseOne() (Error error) {

	spirallineDBArray := make([]SpiralLineDB, 0)
	_, err := backRepoSpiralLine.db.Find(&spirallineDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	spirallineInstancesToBeRemovedFromTheStage := make(map[*models.SpiralLine]any)
	for key, value := range backRepoSpiralLine.stage.SpiralLines {
		spirallineInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, spirallineDB := range spirallineDBArray {
		backRepoSpiralLine.CheckoutPhaseOneInstance(&spirallineDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		spiralline, ok := backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID]
		if ok {
			delete(spirallineInstancesToBeRemovedFromTheStage, spiralline)
		}
	}

	// remove from stage and back repo's 3 maps all spirallines that are not in the checkout
	for spiralline := range spirallineInstancesToBeRemovedFromTheStage {
		spiralline.Unstage(backRepoSpiralLine.GetStage())

		// remove instance from the back repo 3 maps
		spirallineID := backRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline]
		delete(backRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID, spiralline)
		delete(backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB, spirallineID)
		delete(backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr, spirallineID)
	}

	return
}

// CheckoutPhaseOneInstance takes a spirallineDB that has been found in the DB, updates the backRepo and stages the
// models version of the spirallineDB
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CheckoutPhaseOneInstance(spirallineDB *SpiralLineDB) (Error error) {

	spiralline, ok := backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID]
	if !ok {
		spiralline = new(models.SpiralLine)

		backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID] = spiralline
		backRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline] = spirallineDB.ID

		// append model store with the new element
		spiralline.Name = spirallineDB.Name_Data.String
		spiralline.Stage(backRepoSpiralLine.GetStage())
	}
	spirallineDB.CopyBasicFieldsToSpiralLine(spiralline)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	spiralline.Stage(backRepoSpiralLine.GetStage())

	// preserve pointer to spirallineDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SpiralLineDBID_SpiralLineDB)[spirallineDB hold variable pointers
	spirallineDB_Data := *spirallineDB
	preservedPtrToSpiralLine := &spirallineDB_Data
	backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[spirallineDB.ID] = preservedPtrToSpiralLine

	return
}

// BackRepoSpiralLine.CheckoutPhaseTwo Checkouts all staged instances of SpiralLine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, spirallineDB := range backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB {
		backRepoSpiralLine.CheckoutPhaseTwoInstance(backRepo, spirallineDB)
	}
	return
}

// BackRepoSpiralLine.CheckoutPhaseTwoInstance Checkouts staged instances of SpiralLine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralLine *BackRepoSpiralLineStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, spirallineDB *SpiralLineDB) (Error error) {

	spiralline := backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr[spirallineDB.ID]

	spirallineDB.DecodePointers(backRepo, spiralline)

	return
}

func (spirallineDB *SpiralLineDB) DecodePointers(backRepo *BackRepoStruct, spiralline *models.SpiralLine) {

	// insertion point for checkout of pointer encoding
	// ShapeCategory field	
	{
		id := spirallineDB.ShapeCategoryID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: spiralline.ShapeCategory, unknown pointer id", id)
				spiralline.ShapeCategory = nil
			} else {
				// updates only if field has changed
				if spiralline.ShapeCategory == nil || spiralline.ShapeCategory != tmp {
					spiralline.ShapeCategory = tmp
				}
			}
		} else {
			spiralline.ShapeCategory = nil
		}
	}
	
	return
}

// CommitSpiralLine allows commit of a single spiralline (if already staged)
func (backRepo *BackRepoStruct) CommitSpiralLine(spiralline *models.SpiralLine) {
	backRepo.BackRepoSpiralLine.CommitPhaseOneInstance(spiralline)
	if id, ok := backRepo.BackRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline]; ok {
		backRepo.BackRepoSpiralLine.CommitPhaseTwoInstance(backRepo, id, spiralline)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSpiralLine allows checkout of a single spiralline (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSpiralLine(spiralline *models.SpiralLine) {
	// check if the spiralline is staged
	if _, ok := backRepo.BackRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline]; ok {

		if id, ok := backRepo.BackRepoSpiralLine.Map_SpiralLinePtr_SpiralLineDBID[spiralline]; ok {
			var spirallineDB SpiralLineDB
			spirallineDB.ID = id

			if _, err := backRepo.BackRepoSpiralLine.db.First(&spirallineDB, id); err != nil {
				log.Fatalln("CheckoutSpiralLine : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSpiralLine.CheckoutPhaseOneInstance(&spirallineDB)
			backRepo.BackRepoSpiralLine.CheckoutPhaseTwoInstance(backRepo, &spirallineDB)
		}
	}
}

// CopyBasicFieldsFromSpiralLine
func (spirallineDB *SpiralLineDB) CopyBasicFieldsFromSpiralLine(spiralline *models.SpiralLine) {
	// insertion point for fields commit

	spirallineDB.Name_Data.String = spiralline.Name
	spirallineDB.Name_Data.Valid = true

	spirallineDB.IsDisplayed_Data.Bool = spiralline.IsDisplayed
	spirallineDB.IsDisplayed_Data.Valid = true

	spirallineDB.StartX_Data.Float64 = spiralline.StartX
	spirallineDB.StartX_Data.Valid = true

	spirallineDB.EndX_Data.Float64 = spiralline.EndX
	spirallineDB.EndX_Data.Valid = true

	spirallineDB.StartY_Data.Float64 = spiralline.StartY
	spirallineDB.StartY_Data.Valid = true

	spirallineDB.EndY_Data.Float64 = spiralline.EndY
	spirallineDB.EndY_Data.Valid = true

	spirallineDB.Color_Data.String = spiralline.Color
	spirallineDB.Color_Data.Valid = true

	spirallineDB.FillOpacity_Data.Float64 = spiralline.FillOpacity
	spirallineDB.FillOpacity_Data.Valid = true

	spirallineDB.Stroke_Data.String = spiralline.Stroke
	spirallineDB.Stroke_Data.Valid = true

	spirallineDB.StrokeOpacity_Data.Float64 = spiralline.StrokeOpacity
	spirallineDB.StrokeOpacity_Data.Valid = true

	spirallineDB.StrokeWidth_Data.Float64 = spiralline.StrokeWidth
	spirallineDB.StrokeWidth_Data.Valid = true

	spirallineDB.StrokeDashArray_Data.String = spiralline.StrokeDashArray
	spirallineDB.StrokeDashArray_Data.Valid = true

	spirallineDB.StrokeDashArrayWhenSelected_Data.String = spiralline.StrokeDashArrayWhenSelected
	spirallineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spirallineDB.Transform_Data.String = spiralline.Transform
	spirallineDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromSpiralLine_WOP
func (spirallineDB *SpiralLineDB) CopyBasicFieldsFromSpiralLine_WOP(spiralline *models.SpiralLine_WOP) {
	// insertion point for fields commit

	spirallineDB.Name_Data.String = spiralline.Name
	spirallineDB.Name_Data.Valid = true

	spirallineDB.IsDisplayed_Data.Bool = spiralline.IsDisplayed
	spirallineDB.IsDisplayed_Data.Valid = true

	spirallineDB.StartX_Data.Float64 = spiralline.StartX
	spirallineDB.StartX_Data.Valid = true

	spirallineDB.EndX_Data.Float64 = spiralline.EndX
	spirallineDB.EndX_Data.Valid = true

	spirallineDB.StartY_Data.Float64 = spiralline.StartY
	spirallineDB.StartY_Data.Valid = true

	spirallineDB.EndY_Data.Float64 = spiralline.EndY
	spirallineDB.EndY_Data.Valid = true

	spirallineDB.Color_Data.String = spiralline.Color
	spirallineDB.Color_Data.Valid = true

	spirallineDB.FillOpacity_Data.Float64 = spiralline.FillOpacity
	spirallineDB.FillOpacity_Data.Valid = true

	spirallineDB.Stroke_Data.String = spiralline.Stroke
	spirallineDB.Stroke_Data.Valid = true

	spirallineDB.StrokeOpacity_Data.Float64 = spiralline.StrokeOpacity
	spirallineDB.StrokeOpacity_Data.Valid = true

	spirallineDB.StrokeWidth_Data.Float64 = spiralline.StrokeWidth
	spirallineDB.StrokeWidth_Data.Valid = true

	spirallineDB.StrokeDashArray_Data.String = spiralline.StrokeDashArray
	spirallineDB.StrokeDashArray_Data.Valid = true

	spirallineDB.StrokeDashArrayWhenSelected_Data.String = spiralline.StrokeDashArrayWhenSelected
	spirallineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spirallineDB.Transform_Data.String = spiralline.Transform
	spirallineDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromSpiralLineWOP
func (spirallineDB *SpiralLineDB) CopyBasicFieldsFromSpiralLineWOP(spiralline *SpiralLineWOP) {
	// insertion point for fields commit

	spirallineDB.Name_Data.String = spiralline.Name
	spirallineDB.Name_Data.Valid = true

	spirallineDB.IsDisplayed_Data.Bool = spiralline.IsDisplayed
	spirallineDB.IsDisplayed_Data.Valid = true

	spirallineDB.StartX_Data.Float64 = spiralline.StartX
	spirallineDB.StartX_Data.Valid = true

	spirallineDB.EndX_Data.Float64 = spiralline.EndX
	spirallineDB.EndX_Data.Valid = true

	spirallineDB.StartY_Data.Float64 = spiralline.StartY
	spirallineDB.StartY_Data.Valid = true

	spirallineDB.EndY_Data.Float64 = spiralline.EndY
	spirallineDB.EndY_Data.Valid = true

	spirallineDB.Color_Data.String = spiralline.Color
	spirallineDB.Color_Data.Valid = true

	spirallineDB.FillOpacity_Data.Float64 = spiralline.FillOpacity
	spirallineDB.FillOpacity_Data.Valid = true

	spirallineDB.Stroke_Data.String = spiralline.Stroke
	spirallineDB.Stroke_Data.Valid = true

	spirallineDB.StrokeOpacity_Data.Float64 = spiralline.StrokeOpacity
	spirallineDB.StrokeOpacity_Data.Valid = true

	spirallineDB.StrokeWidth_Data.Float64 = spiralline.StrokeWidth
	spirallineDB.StrokeWidth_Data.Valid = true

	spirallineDB.StrokeDashArray_Data.String = spiralline.StrokeDashArray
	spirallineDB.StrokeDashArray_Data.Valid = true

	spirallineDB.StrokeDashArrayWhenSelected_Data.String = spiralline.StrokeDashArrayWhenSelected
	spirallineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spirallineDB.Transform_Data.String = spiralline.Transform
	spirallineDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToSpiralLine
func (spirallineDB *SpiralLineDB) CopyBasicFieldsToSpiralLine(spiralline *models.SpiralLine) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralline.Name = spirallineDB.Name_Data.String
	spiralline.IsDisplayed = spirallineDB.IsDisplayed_Data.Bool
	spiralline.StartX = spirallineDB.StartX_Data.Float64
	spiralline.EndX = spirallineDB.EndX_Data.Float64
	spiralline.StartY = spirallineDB.StartY_Data.Float64
	spiralline.EndY = spirallineDB.EndY_Data.Float64
	spiralline.Color = spirallineDB.Color_Data.String
	spiralline.FillOpacity = spirallineDB.FillOpacity_Data.Float64
	spiralline.Stroke = spirallineDB.Stroke_Data.String
	spiralline.StrokeOpacity = spirallineDB.StrokeOpacity_Data.Float64
	spiralline.StrokeWidth = spirallineDB.StrokeWidth_Data.Float64
	spiralline.StrokeDashArray = spirallineDB.StrokeDashArray_Data.String
	spiralline.StrokeDashArrayWhenSelected = spirallineDB.StrokeDashArrayWhenSelected_Data.String
	spiralline.Transform = spirallineDB.Transform_Data.String
}

// CopyBasicFieldsToSpiralLine_WOP
func (spirallineDB *SpiralLineDB) CopyBasicFieldsToSpiralLine_WOP(spiralline *models.SpiralLine_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralline.Name = spirallineDB.Name_Data.String
	spiralline.IsDisplayed = spirallineDB.IsDisplayed_Data.Bool
	spiralline.StartX = spirallineDB.StartX_Data.Float64
	spiralline.EndX = spirallineDB.EndX_Data.Float64
	spiralline.StartY = spirallineDB.StartY_Data.Float64
	spiralline.EndY = spirallineDB.EndY_Data.Float64
	spiralline.Color = spirallineDB.Color_Data.String
	spiralline.FillOpacity = spirallineDB.FillOpacity_Data.Float64
	spiralline.Stroke = spirallineDB.Stroke_Data.String
	spiralline.StrokeOpacity = spirallineDB.StrokeOpacity_Data.Float64
	spiralline.StrokeWidth = spirallineDB.StrokeWidth_Data.Float64
	spiralline.StrokeDashArray = spirallineDB.StrokeDashArray_Data.String
	spiralline.StrokeDashArrayWhenSelected = spirallineDB.StrokeDashArrayWhenSelected_Data.String
	spiralline.Transform = spirallineDB.Transform_Data.String
}

// CopyBasicFieldsToSpiralLineWOP
func (spirallineDB *SpiralLineDB) CopyBasicFieldsToSpiralLineWOP(spiralline *SpiralLineWOP) {
	spiralline.ID = int(spirallineDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	spiralline.Name = spirallineDB.Name_Data.String
	spiralline.IsDisplayed = spirallineDB.IsDisplayed_Data.Bool
	spiralline.StartX = spirallineDB.StartX_Data.Float64
	spiralline.EndX = spirallineDB.EndX_Data.Float64
	spiralline.StartY = spirallineDB.StartY_Data.Float64
	spiralline.EndY = spirallineDB.EndY_Data.Float64
	spiralline.Color = spirallineDB.Color_Data.String
	spiralline.FillOpacity = spirallineDB.FillOpacity_Data.Float64
	spiralline.Stroke = spirallineDB.Stroke_Data.String
	spiralline.StrokeOpacity = spirallineDB.StrokeOpacity_Data.Float64
	spiralline.StrokeWidth = spirallineDB.StrokeWidth_Data.Float64
	spiralline.StrokeDashArray = spirallineDB.StrokeDashArray_Data.String
	spiralline.StrokeDashArrayWhenSelected = spirallineDB.StrokeDashArrayWhenSelected_Data.String
	spiralline.Transform = spirallineDB.Transform_Data.String
}

// Backup generates a json file from a slice of all SpiralLineDB instances in the backrepo
func (backRepoSpiralLine *BackRepoSpiralLineStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SpiralLineDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralLineDB, 0)
	for _, spirallineDB := range backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB {
		forBackup = append(forBackup, spirallineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SpiralLine ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SpiralLine file", err.Error())
	}
}

// Backup generates a json file from a slice of all SpiralLineDB instances in the backrepo
func (backRepoSpiralLine *BackRepoSpiralLineStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralLineDB, 0)
	for _, spirallineDB := range backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB {
		forBackup = append(forBackup, spirallineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SpiralLine")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SpiralLine_Fields, -1)
	for _, spirallineDB := range forBackup {

		var spirallineWOP SpiralLineWOP
		spirallineDB.CopyBasicFieldsToSpiralLineWOP(&spirallineWOP)

		row := sh.AddRow()
		row.WriteStruct(&spirallineWOP, -1)
	}
}

// RestoreXL from the "SpiralLine" sheet all SpiralLineDB instances
func (backRepoSpiralLine *BackRepoSpiralLineStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSpiralLineid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SpiralLine"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSpiralLine.rowVisitorSpiralLine)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSpiralLine *BackRepoSpiralLineStruct) rowVisitorSpiralLine(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var spirallineWOP SpiralLineWOP
		row.ReadStruct(&spirallineWOP)

		// add the unmarshalled struct to the stage
		spirallineDB := new(SpiralLineDB)
		spirallineDB.CopyBasicFieldsFromSpiralLineWOP(&spirallineWOP)

		spirallineDB_ID_atBackupTime := spirallineDB.ID
		spirallineDB.ID = 0
		_, err := backRepoSpiralLine.db.Create(spirallineDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[spirallineDB.ID] = spirallineDB
		BackRepoSpiralLineid_atBckpTime_newID[spirallineDB_ID_atBackupTime] = spirallineDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SpiralLineDB.json" in dirPath that stores an array
// of SpiralLineDB and stores it in the database
// the map BackRepoSpiralLineid_atBckpTime_newID is updated accordingly
func (backRepoSpiralLine *BackRepoSpiralLineStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSpiralLineid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SpiralLineDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SpiralLine file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SpiralLineDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SpiralLineDBID_SpiralLineDB
	for _, spirallineDB := range forRestore {

		spirallineDB_ID_atBackupTime := spirallineDB.ID
		spirallineDB.ID = 0
		_, err := backRepoSpiralLine.db.Create(spirallineDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[spirallineDB.ID] = spirallineDB
		BackRepoSpiralLineid_atBckpTime_newID[spirallineDB_ID_atBackupTime] = spirallineDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SpiralLine file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SpiralLine>id_atBckpTime_newID
// to compute new index
func (backRepoSpiralLine *BackRepoSpiralLineStruct) RestorePhaseTwo() {

	for _, spirallineDB := range backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB {

		// next line of code is to avert unused variable compilation error
		_ = spirallineDB

		// insertion point for reindexing pointers encoding
		// reindexing ShapeCategory field
		if spirallineDB.ShapeCategoryID.Int64 != 0 {
			spirallineDB.ShapeCategoryID.Int64 = int64(BackRepoShapeCategoryid_atBckpTime_newID[uint(spirallineDB.ShapeCategoryID.Int64)])
			spirallineDB.ShapeCategoryID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoSpiralLine.db.Model(spirallineDB)
		_, err := db.Updates(*spirallineDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoSpiralLine.ResetReversePointers commits all staged instances of SpiralLine to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralLine *BackRepoSpiralLineStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralline := range backRepoSpiralLine.Map_SpiralLineDBID_SpiralLinePtr {
		backRepoSpiralLine.ResetReversePointersInstance(backRepo, idx, spiralline)
	}

	return
}

func (backRepoSpiralLine *BackRepoSpiralLineStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, spiralline *models.SpiralLine) (Error error) {

	// fetch matching spirallineDB
	if spirallineDB, ok := backRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB[idx]; ok {
		_ = spirallineDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSpiralLineid_atBckpTime_newID map[uint]uint
