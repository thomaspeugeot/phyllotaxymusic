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
var dummy_SpiralCircle_sql sql.NullBool
var dummy_SpiralCircle_time time.Duration
var dummy_SpiralCircle_sort sort.Float64Slice

// SpiralCircleAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model spiralcircleAPI
type SpiralCircleAPI struct {
	gorm.Model

	models.SpiralCircle_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SpiralCirclePointersEncoding SpiralCirclePointersEncoding
}

// SpiralCirclePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SpiralCirclePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ShapeCategory is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ShapeCategoryID sql.NullInt64
}

// SpiralCircleDB describes a spiralcircle in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model spiralcircleDB
type SpiralCircleDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field spiralcircleDB.Name
	Name_Data sql.NullString

	// Declation for basic field spiralcircleDB.IsDisplayed
	// provide the sql storage for the boolan
	IsDisplayed_Data sql.NullBool

	// Declation for basic field spiralcircleDB.CenterX
	CenterX_Data sql.NullFloat64

	// Declation for basic field spiralcircleDB.CenterY
	CenterY_Data sql.NullFloat64

	// Declation for basic field spiralcircleDB.HasBespokeRadius
	// provide the sql storage for the boolan
	HasBespokeRadius_Data sql.NullBool

	// Declation for basic field spiralcircleDB.BespopkeRadius
	BespopkeRadius_Data sql.NullFloat64

	// Declation for basic field spiralcircleDB.Color
	Color_Data sql.NullString

	// Declation for basic field spiralcircleDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field spiralcircleDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field spiralcircleDB.StrokeOpacity
	StrokeOpacity_Data sql.NullFloat64

	// Declation for basic field spiralcircleDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field spiralcircleDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field spiralcircleDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field spiralcircleDB.Transform
	Transform_Data sql.NullString

	// Declation for basic field spiralcircleDB.Pitch
	Pitch_Data sql.NullInt64

	// Declation for basic field spiralcircleDB.ShowName
	// provide the sql storage for the boolan
	ShowName_Data sql.NullBool

	// Declation for basic field spiralcircleDB.BeatNb
	BeatNb_Data sql.NullInt64

	// Declation for basic field spiralcircleDB.Path
	Path_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SpiralCirclePointersEncoding
}

// SpiralCircleDBs arrays spiralcircleDBs
// swagger:response spiralcircleDBsResponse
type SpiralCircleDBs []SpiralCircleDB

// SpiralCircleDBResponse provides response
// swagger:response spiralcircleDBResponse
type SpiralCircleDBResponse struct {
	SpiralCircleDB
}

// SpiralCircleWOP is a SpiralCircle without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SpiralCircleWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsDisplayed bool `xlsx:"2"`

	CenterX float64 `xlsx:"3"`

	CenterY float64 `xlsx:"4"`

	HasBespokeRadius bool `xlsx:"5"`

	BespopkeRadius float64 `xlsx:"6"`

	Color string `xlsx:"7"`

	FillOpacity float64 `xlsx:"8"`

	Stroke string `xlsx:"9"`

	StrokeOpacity float64 `xlsx:"10"`

	StrokeWidth float64 `xlsx:"11"`

	StrokeDashArray string `xlsx:"12"`

	StrokeDashArrayWhenSelected string `xlsx:"13"`

	Transform string `xlsx:"14"`

	Pitch int `xlsx:"15"`

	ShowName bool `xlsx:"16"`

	BeatNb int `xlsx:"17"`

	Path string `xlsx:"18"`
	// insertion for WOP pointer fields
}

var SpiralCircle_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsDisplayed",
	"CenterX",
	"CenterY",
	"HasBespokeRadius",
	"BespopkeRadius",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeOpacity",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
	"Pitch",
	"ShowName",
	"BeatNb",
	"Path",
}

type BackRepoSpiralCircleStruct struct {
	// stores SpiralCircleDB according to their gorm ID
	Map_SpiralCircleDBID_SpiralCircleDB map[uint]*SpiralCircleDB

	// stores SpiralCircleDB ID according to SpiralCircle address
	Map_SpiralCirclePtr_SpiralCircleDBID map[*models.SpiralCircle]uint

	// stores SpiralCircle according to their gorm ID
	Map_SpiralCircleDBID_SpiralCirclePtr map[uint]*models.SpiralCircle

	db db.DBInterface

	stage *models.Stage
}

func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) GetStage() (stage *models.Stage) {
	stage = backRepoSpiralCircle.stage
	return
}

func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) GetDB() db.DBInterface {
	return backRepoSpiralCircle.db
}

// GetSpiralCircleDBFromSpiralCirclePtr is a handy function to access the back repo instance from the stage instance
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) GetSpiralCircleDBFromSpiralCirclePtr(spiralcircle *models.SpiralCircle) (spiralcircleDB *SpiralCircleDB) {
	id := backRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle]
	spiralcircleDB = backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[id]
	return
}

// BackRepoSpiralCircle.CommitPhaseOne commits all staged instances of SpiralCircle to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var spiralcircles []*models.SpiralCircle
	for spiralcircle := range stage.SpiralCircles {
		spiralcircles = append(spiralcircles, spiralcircle)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(spiralcircles, func(i, j int) bool {
		return stage.SpiralCircleMap_Staged_Order[spiralcircles[i]] < stage.SpiralCircleMap_Staged_Order[spiralcircles[j]]
	})

	for _, spiralcircle := range spiralcircles {
		backRepoSpiralCircle.CommitPhaseOneInstance(spiralcircle)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, spiralcircle := range backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr {
		if _, ok := stage.SpiralCircles[spiralcircle]; !ok {
			backRepoSpiralCircle.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSpiralCircle.CommitDeleteInstance commits deletion of SpiralCircle to the BackRepo
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CommitDeleteInstance(id uint) (Error error) {

	spiralcircle := backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[id]

	// spiralcircle is not staged anymore, remove spiralcircleDB
	spiralcircleDB := backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[id]
	db, _ := backRepoSpiralCircle.db.Unscoped()
	_, err := db.Delete(spiralcircleDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID, spiralcircle)
	delete(backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr, id)
	delete(backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB, id)

	return
}

// BackRepoSpiralCircle.CommitPhaseOneInstance commits spiralcircle staged instances of SpiralCircle to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CommitPhaseOneInstance(spiralcircle *models.SpiralCircle) (Error error) {

	// check if the spiralcircle is not commited yet
	if _, ok := backRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle]; ok {
		return
	}

	// initiate spiralcircle
	var spiralcircleDB SpiralCircleDB
	spiralcircleDB.CopyBasicFieldsFromSpiralCircle(spiralcircle)

	_, err := backRepoSpiralCircle.db.Create(&spiralcircleDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle] = spiralcircleDB.ID
	backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID] = spiralcircle
	backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[spiralcircleDB.ID] = &spiralcircleDB

	return
}

// BackRepoSpiralCircle.CommitPhaseTwo commits all staged instances of SpiralCircle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralcircle := range backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr {
		backRepoSpiralCircle.CommitPhaseTwoInstance(backRepo, idx, spiralcircle)
	}

	return
}

// BackRepoSpiralCircle.CommitPhaseTwoInstance commits {{structname }} of models.SpiralCircle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, spiralcircle *models.SpiralCircle) (Error error) {

	// fetch matching spiralcircleDB
	if spiralcircleDB, ok := backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[idx]; ok {

		spiralcircleDB.CopyBasicFieldsFromSpiralCircle(spiralcircle)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value spiralcircle.ShapeCategory translates to updating the spiralcircle.ShapeCategoryID
		spiralcircleDB.ShapeCategoryID.Valid = true // allow for a 0 value (nil association)
		if spiralcircle.ShapeCategory != nil {
			if ShapeCategoryId, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[spiralcircle.ShapeCategory]; ok {
				spiralcircleDB.ShapeCategoryID.Int64 = int64(ShapeCategoryId)
				spiralcircleDB.ShapeCategoryID.Valid = true
			}
		} else {
			spiralcircleDB.ShapeCategoryID.Int64 = 0
			spiralcircleDB.ShapeCategoryID.Valid = true
		}

		_, err := backRepoSpiralCircle.db.Save(spiralcircleDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SpiralCircle intance %s", spiralcircle.Name))
		return err
	}

	return
}

// BackRepoSpiralCircle.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CheckoutPhaseOne() (Error error) {

	spiralcircleDBArray := make([]SpiralCircleDB, 0)
	_, err := backRepoSpiralCircle.db.Find(&spiralcircleDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	spiralcircleInstancesToBeRemovedFromTheStage := make(map[*models.SpiralCircle]any)
	for key, value := range backRepoSpiralCircle.stage.SpiralCircles {
		spiralcircleInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, spiralcircleDB := range spiralcircleDBArray {
		backRepoSpiralCircle.CheckoutPhaseOneInstance(&spiralcircleDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		spiralcircle, ok := backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID]
		if ok {
			delete(spiralcircleInstancesToBeRemovedFromTheStage, spiralcircle)
		}
	}

	// remove from stage and back repo's 3 maps all spiralcircles that are not in the checkout
	for spiralcircle := range spiralcircleInstancesToBeRemovedFromTheStage {
		spiralcircle.Unstage(backRepoSpiralCircle.GetStage())

		// remove instance from the back repo 3 maps
		spiralcircleID := backRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle]
		delete(backRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID, spiralcircle)
		delete(backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB, spiralcircleID)
		delete(backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr, spiralcircleID)
	}

	return
}

// CheckoutPhaseOneInstance takes a spiralcircleDB that has been found in the DB, updates the backRepo and stages the
// models version of the spiralcircleDB
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CheckoutPhaseOneInstance(spiralcircleDB *SpiralCircleDB) (Error error) {

	spiralcircle, ok := backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID]
	if !ok {
		spiralcircle = new(models.SpiralCircle)

		backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID] = spiralcircle
		backRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle] = spiralcircleDB.ID

		// append model store with the new element
		spiralcircle.Name = spiralcircleDB.Name_Data.String
		spiralcircle.Stage(backRepoSpiralCircle.GetStage())
	}
	spiralcircleDB.CopyBasicFieldsToSpiralCircle(spiralcircle)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	spiralcircle.Stage(backRepoSpiralCircle.GetStage())

	// preserve pointer to spiralcircleDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SpiralCircleDBID_SpiralCircleDB)[spiralcircleDB hold variable pointers
	spiralcircleDB_Data := *spiralcircleDB
	preservedPtrToSpiralCircle := &spiralcircleDB_Data
	backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[spiralcircleDB.ID] = preservedPtrToSpiralCircle

	return
}

// BackRepoSpiralCircle.CheckoutPhaseTwo Checkouts all staged instances of SpiralCircle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, spiralcircleDB := range backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB {
		backRepoSpiralCircle.CheckoutPhaseTwoInstance(backRepo, spiralcircleDB)
	}
	return
}

// BackRepoSpiralCircle.CheckoutPhaseTwoInstance Checkouts staged instances of SpiralCircle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, spiralcircleDB *SpiralCircleDB) (Error error) {

	spiralcircle := backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr[spiralcircleDB.ID]

	spiralcircleDB.DecodePointers(backRepo, spiralcircle)

	return
}

func (spiralcircleDB *SpiralCircleDB) DecodePointers(backRepo *BackRepoStruct, spiralcircle *models.SpiralCircle) {

	// insertion point for checkout of pointer encoding
	// ShapeCategory field	
	{
		id := spiralcircleDB.ShapeCategoryID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: spiralcircle.ShapeCategory, unknown pointer id", id)
				spiralcircle.ShapeCategory = nil
			} else {
				// updates only if field has changed
				if spiralcircle.ShapeCategory == nil || spiralcircle.ShapeCategory != tmp {
					spiralcircle.ShapeCategory = tmp
				}
			}
		} else {
			spiralcircle.ShapeCategory = nil
		}
	}
	
	return
}

// CommitSpiralCircle allows commit of a single spiralcircle (if already staged)
func (backRepo *BackRepoStruct) CommitSpiralCircle(spiralcircle *models.SpiralCircle) {
	backRepo.BackRepoSpiralCircle.CommitPhaseOneInstance(spiralcircle)
	if id, ok := backRepo.BackRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle]; ok {
		backRepo.BackRepoSpiralCircle.CommitPhaseTwoInstance(backRepo, id, spiralcircle)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSpiralCircle allows checkout of a single spiralcircle (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSpiralCircle(spiralcircle *models.SpiralCircle) {
	// check if the spiralcircle is staged
	if _, ok := backRepo.BackRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle]; ok {

		if id, ok := backRepo.BackRepoSpiralCircle.Map_SpiralCirclePtr_SpiralCircleDBID[spiralcircle]; ok {
			var spiralcircleDB SpiralCircleDB
			spiralcircleDB.ID = id

			if _, err := backRepo.BackRepoSpiralCircle.db.First(&spiralcircleDB, id); err != nil {
				log.Fatalln("CheckoutSpiralCircle : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSpiralCircle.CheckoutPhaseOneInstance(&spiralcircleDB)
			backRepo.BackRepoSpiralCircle.CheckoutPhaseTwoInstance(backRepo, &spiralcircleDB)
		}
	}
}

// CopyBasicFieldsFromSpiralCircle
func (spiralcircleDB *SpiralCircleDB) CopyBasicFieldsFromSpiralCircle(spiralcircle *models.SpiralCircle) {
	// insertion point for fields commit

	spiralcircleDB.Name_Data.String = spiralcircle.Name
	spiralcircleDB.Name_Data.Valid = true

	spiralcircleDB.IsDisplayed_Data.Bool = spiralcircle.IsDisplayed
	spiralcircleDB.IsDisplayed_Data.Valid = true

	spiralcircleDB.CenterX_Data.Float64 = spiralcircle.CenterX
	spiralcircleDB.CenterX_Data.Valid = true

	spiralcircleDB.CenterY_Data.Float64 = spiralcircle.CenterY
	spiralcircleDB.CenterY_Data.Valid = true

	spiralcircleDB.HasBespokeRadius_Data.Bool = spiralcircle.HasBespokeRadius
	spiralcircleDB.HasBespokeRadius_Data.Valid = true

	spiralcircleDB.BespopkeRadius_Data.Float64 = spiralcircle.BespopkeRadius
	spiralcircleDB.BespopkeRadius_Data.Valid = true

	spiralcircleDB.Color_Data.String = spiralcircle.Color
	spiralcircleDB.Color_Data.Valid = true

	spiralcircleDB.FillOpacity_Data.Float64 = spiralcircle.FillOpacity
	spiralcircleDB.FillOpacity_Data.Valid = true

	spiralcircleDB.Stroke_Data.String = spiralcircle.Stroke
	spiralcircleDB.Stroke_Data.Valid = true

	spiralcircleDB.StrokeOpacity_Data.Float64 = spiralcircle.StrokeOpacity
	spiralcircleDB.StrokeOpacity_Data.Valid = true

	spiralcircleDB.StrokeWidth_Data.Float64 = spiralcircle.StrokeWidth
	spiralcircleDB.StrokeWidth_Data.Valid = true

	spiralcircleDB.StrokeDashArray_Data.String = spiralcircle.StrokeDashArray
	spiralcircleDB.StrokeDashArray_Data.Valid = true

	spiralcircleDB.StrokeDashArrayWhenSelected_Data.String = spiralcircle.StrokeDashArrayWhenSelected
	spiralcircleDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spiralcircleDB.Transform_Data.String = spiralcircle.Transform
	spiralcircleDB.Transform_Data.Valid = true

	spiralcircleDB.Pitch_Data.Int64 = int64(spiralcircle.Pitch)
	spiralcircleDB.Pitch_Data.Valid = true

	spiralcircleDB.ShowName_Data.Bool = spiralcircle.ShowName
	spiralcircleDB.ShowName_Data.Valid = true

	spiralcircleDB.BeatNb_Data.Int64 = int64(spiralcircle.BeatNb)
	spiralcircleDB.BeatNb_Data.Valid = true

	spiralcircleDB.Path_Data.String = spiralcircle.Path
	spiralcircleDB.Path_Data.Valid = true
}

// CopyBasicFieldsFromSpiralCircle_WOP
func (spiralcircleDB *SpiralCircleDB) CopyBasicFieldsFromSpiralCircle_WOP(spiralcircle *models.SpiralCircle_WOP) {
	// insertion point for fields commit

	spiralcircleDB.Name_Data.String = spiralcircle.Name
	spiralcircleDB.Name_Data.Valid = true

	spiralcircleDB.IsDisplayed_Data.Bool = spiralcircle.IsDisplayed
	spiralcircleDB.IsDisplayed_Data.Valid = true

	spiralcircleDB.CenterX_Data.Float64 = spiralcircle.CenterX
	spiralcircleDB.CenterX_Data.Valid = true

	spiralcircleDB.CenterY_Data.Float64 = spiralcircle.CenterY
	spiralcircleDB.CenterY_Data.Valid = true

	spiralcircleDB.HasBespokeRadius_Data.Bool = spiralcircle.HasBespokeRadius
	spiralcircleDB.HasBespokeRadius_Data.Valid = true

	spiralcircleDB.BespopkeRadius_Data.Float64 = spiralcircle.BespopkeRadius
	spiralcircleDB.BespopkeRadius_Data.Valid = true

	spiralcircleDB.Color_Data.String = spiralcircle.Color
	spiralcircleDB.Color_Data.Valid = true

	spiralcircleDB.FillOpacity_Data.Float64 = spiralcircle.FillOpacity
	spiralcircleDB.FillOpacity_Data.Valid = true

	spiralcircleDB.Stroke_Data.String = spiralcircle.Stroke
	spiralcircleDB.Stroke_Data.Valid = true

	spiralcircleDB.StrokeOpacity_Data.Float64 = spiralcircle.StrokeOpacity
	spiralcircleDB.StrokeOpacity_Data.Valid = true

	spiralcircleDB.StrokeWidth_Data.Float64 = spiralcircle.StrokeWidth
	spiralcircleDB.StrokeWidth_Data.Valid = true

	spiralcircleDB.StrokeDashArray_Data.String = spiralcircle.StrokeDashArray
	spiralcircleDB.StrokeDashArray_Data.Valid = true

	spiralcircleDB.StrokeDashArrayWhenSelected_Data.String = spiralcircle.StrokeDashArrayWhenSelected
	spiralcircleDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spiralcircleDB.Transform_Data.String = spiralcircle.Transform
	spiralcircleDB.Transform_Data.Valid = true

	spiralcircleDB.Pitch_Data.Int64 = int64(spiralcircle.Pitch)
	spiralcircleDB.Pitch_Data.Valid = true

	spiralcircleDB.ShowName_Data.Bool = spiralcircle.ShowName
	spiralcircleDB.ShowName_Data.Valid = true

	spiralcircleDB.BeatNb_Data.Int64 = int64(spiralcircle.BeatNb)
	spiralcircleDB.BeatNb_Data.Valid = true

	spiralcircleDB.Path_Data.String = spiralcircle.Path
	spiralcircleDB.Path_Data.Valid = true
}

// CopyBasicFieldsFromSpiralCircleWOP
func (spiralcircleDB *SpiralCircleDB) CopyBasicFieldsFromSpiralCircleWOP(spiralcircle *SpiralCircleWOP) {
	// insertion point for fields commit

	spiralcircleDB.Name_Data.String = spiralcircle.Name
	spiralcircleDB.Name_Data.Valid = true

	spiralcircleDB.IsDisplayed_Data.Bool = spiralcircle.IsDisplayed
	spiralcircleDB.IsDisplayed_Data.Valid = true

	spiralcircleDB.CenterX_Data.Float64 = spiralcircle.CenterX
	spiralcircleDB.CenterX_Data.Valid = true

	spiralcircleDB.CenterY_Data.Float64 = spiralcircle.CenterY
	spiralcircleDB.CenterY_Data.Valid = true

	spiralcircleDB.HasBespokeRadius_Data.Bool = spiralcircle.HasBespokeRadius
	spiralcircleDB.HasBespokeRadius_Data.Valid = true

	spiralcircleDB.BespopkeRadius_Data.Float64 = spiralcircle.BespopkeRadius
	spiralcircleDB.BespopkeRadius_Data.Valid = true

	spiralcircleDB.Color_Data.String = spiralcircle.Color
	spiralcircleDB.Color_Data.Valid = true

	spiralcircleDB.FillOpacity_Data.Float64 = spiralcircle.FillOpacity
	spiralcircleDB.FillOpacity_Data.Valid = true

	spiralcircleDB.Stroke_Data.String = spiralcircle.Stroke
	spiralcircleDB.Stroke_Data.Valid = true

	spiralcircleDB.StrokeOpacity_Data.Float64 = spiralcircle.StrokeOpacity
	spiralcircleDB.StrokeOpacity_Data.Valid = true

	spiralcircleDB.StrokeWidth_Data.Float64 = spiralcircle.StrokeWidth
	spiralcircleDB.StrokeWidth_Data.Valid = true

	spiralcircleDB.StrokeDashArray_Data.String = spiralcircle.StrokeDashArray
	spiralcircleDB.StrokeDashArray_Data.Valid = true

	spiralcircleDB.StrokeDashArrayWhenSelected_Data.String = spiralcircle.StrokeDashArrayWhenSelected
	spiralcircleDB.StrokeDashArrayWhenSelected_Data.Valid = true

	spiralcircleDB.Transform_Data.String = spiralcircle.Transform
	spiralcircleDB.Transform_Data.Valid = true

	spiralcircleDB.Pitch_Data.Int64 = int64(spiralcircle.Pitch)
	spiralcircleDB.Pitch_Data.Valid = true

	spiralcircleDB.ShowName_Data.Bool = spiralcircle.ShowName
	spiralcircleDB.ShowName_Data.Valid = true

	spiralcircleDB.BeatNb_Data.Int64 = int64(spiralcircle.BeatNb)
	spiralcircleDB.BeatNb_Data.Valid = true

	spiralcircleDB.Path_Data.String = spiralcircle.Path
	spiralcircleDB.Path_Data.Valid = true
}

// CopyBasicFieldsToSpiralCircle
func (spiralcircleDB *SpiralCircleDB) CopyBasicFieldsToSpiralCircle(spiralcircle *models.SpiralCircle) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralcircle.Name = spiralcircleDB.Name_Data.String
	spiralcircle.IsDisplayed = spiralcircleDB.IsDisplayed_Data.Bool
	spiralcircle.CenterX = spiralcircleDB.CenterX_Data.Float64
	spiralcircle.CenterY = spiralcircleDB.CenterY_Data.Float64
	spiralcircle.HasBespokeRadius = spiralcircleDB.HasBespokeRadius_Data.Bool
	spiralcircle.BespopkeRadius = spiralcircleDB.BespopkeRadius_Data.Float64
	spiralcircle.Color = spiralcircleDB.Color_Data.String
	spiralcircle.FillOpacity = spiralcircleDB.FillOpacity_Data.Float64
	spiralcircle.Stroke = spiralcircleDB.Stroke_Data.String
	spiralcircle.StrokeOpacity = spiralcircleDB.StrokeOpacity_Data.Float64
	spiralcircle.StrokeWidth = spiralcircleDB.StrokeWidth_Data.Float64
	spiralcircle.StrokeDashArray = spiralcircleDB.StrokeDashArray_Data.String
	spiralcircle.StrokeDashArrayWhenSelected = spiralcircleDB.StrokeDashArrayWhenSelected_Data.String
	spiralcircle.Transform = spiralcircleDB.Transform_Data.String
	spiralcircle.Pitch = int(spiralcircleDB.Pitch_Data.Int64)
	spiralcircle.ShowName = spiralcircleDB.ShowName_Data.Bool
	spiralcircle.BeatNb = int(spiralcircleDB.BeatNb_Data.Int64)
	spiralcircle.Path = spiralcircleDB.Path_Data.String
}

// CopyBasicFieldsToSpiralCircle_WOP
func (spiralcircleDB *SpiralCircleDB) CopyBasicFieldsToSpiralCircle_WOP(spiralcircle *models.SpiralCircle_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralcircle.Name = spiralcircleDB.Name_Data.String
	spiralcircle.IsDisplayed = spiralcircleDB.IsDisplayed_Data.Bool
	spiralcircle.CenterX = spiralcircleDB.CenterX_Data.Float64
	spiralcircle.CenterY = spiralcircleDB.CenterY_Data.Float64
	spiralcircle.HasBespokeRadius = spiralcircleDB.HasBespokeRadius_Data.Bool
	spiralcircle.BespopkeRadius = spiralcircleDB.BespopkeRadius_Data.Float64
	spiralcircle.Color = spiralcircleDB.Color_Data.String
	spiralcircle.FillOpacity = spiralcircleDB.FillOpacity_Data.Float64
	spiralcircle.Stroke = spiralcircleDB.Stroke_Data.String
	spiralcircle.StrokeOpacity = spiralcircleDB.StrokeOpacity_Data.Float64
	spiralcircle.StrokeWidth = spiralcircleDB.StrokeWidth_Data.Float64
	spiralcircle.StrokeDashArray = spiralcircleDB.StrokeDashArray_Data.String
	spiralcircle.StrokeDashArrayWhenSelected = spiralcircleDB.StrokeDashArrayWhenSelected_Data.String
	spiralcircle.Transform = spiralcircleDB.Transform_Data.String
	spiralcircle.Pitch = int(spiralcircleDB.Pitch_Data.Int64)
	spiralcircle.ShowName = spiralcircleDB.ShowName_Data.Bool
	spiralcircle.BeatNb = int(spiralcircleDB.BeatNb_Data.Int64)
	spiralcircle.Path = spiralcircleDB.Path_Data.String
}

// CopyBasicFieldsToSpiralCircleWOP
func (spiralcircleDB *SpiralCircleDB) CopyBasicFieldsToSpiralCircleWOP(spiralcircle *SpiralCircleWOP) {
	spiralcircle.ID = int(spiralcircleDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	spiralcircle.Name = spiralcircleDB.Name_Data.String
	spiralcircle.IsDisplayed = spiralcircleDB.IsDisplayed_Data.Bool
	spiralcircle.CenterX = spiralcircleDB.CenterX_Data.Float64
	spiralcircle.CenterY = spiralcircleDB.CenterY_Data.Float64
	spiralcircle.HasBespokeRadius = spiralcircleDB.HasBespokeRadius_Data.Bool
	spiralcircle.BespopkeRadius = spiralcircleDB.BespopkeRadius_Data.Float64
	spiralcircle.Color = spiralcircleDB.Color_Data.String
	spiralcircle.FillOpacity = spiralcircleDB.FillOpacity_Data.Float64
	spiralcircle.Stroke = spiralcircleDB.Stroke_Data.String
	spiralcircle.StrokeOpacity = spiralcircleDB.StrokeOpacity_Data.Float64
	spiralcircle.StrokeWidth = spiralcircleDB.StrokeWidth_Data.Float64
	spiralcircle.StrokeDashArray = spiralcircleDB.StrokeDashArray_Data.String
	spiralcircle.StrokeDashArrayWhenSelected = spiralcircleDB.StrokeDashArrayWhenSelected_Data.String
	spiralcircle.Transform = spiralcircleDB.Transform_Data.String
	spiralcircle.Pitch = int(spiralcircleDB.Pitch_Data.Int64)
	spiralcircle.ShowName = spiralcircleDB.ShowName_Data.Bool
	spiralcircle.BeatNb = int(spiralcircleDB.BeatNb_Data.Int64)
	spiralcircle.Path = spiralcircleDB.Path_Data.String
}

// Backup generates a json file from a slice of all SpiralCircleDB instances in the backrepo
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SpiralCircleDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralCircleDB, 0)
	for _, spiralcircleDB := range backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB {
		forBackup = append(forBackup, spiralcircleDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SpiralCircle ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SpiralCircle file", err.Error())
	}
}

// Backup generates a json file from a slice of all SpiralCircleDB instances in the backrepo
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralCircleDB, 0)
	for _, spiralcircleDB := range backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB {
		forBackup = append(forBackup, spiralcircleDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SpiralCircle")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SpiralCircle_Fields, -1)
	for _, spiralcircleDB := range forBackup {

		var spiralcircleWOP SpiralCircleWOP
		spiralcircleDB.CopyBasicFieldsToSpiralCircleWOP(&spiralcircleWOP)

		row := sh.AddRow()
		row.WriteStruct(&spiralcircleWOP, -1)
	}
}

// RestoreXL from the "SpiralCircle" sheet all SpiralCircleDB instances
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSpiralCircleid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SpiralCircle"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSpiralCircle.rowVisitorSpiralCircle)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) rowVisitorSpiralCircle(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var spiralcircleWOP SpiralCircleWOP
		row.ReadStruct(&spiralcircleWOP)

		// add the unmarshalled struct to the stage
		spiralcircleDB := new(SpiralCircleDB)
		spiralcircleDB.CopyBasicFieldsFromSpiralCircleWOP(&spiralcircleWOP)

		spiralcircleDB_ID_atBackupTime := spiralcircleDB.ID
		spiralcircleDB.ID = 0
		_, err := backRepoSpiralCircle.db.Create(spiralcircleDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[spiralcircleDB.ID] = spiralcircleDB
		BackRepoSpiralCircleid_atBckpTime_newID[spiralcircleDB_ID_atBackupTime] = spiralcircleDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SpiralCircleDB.json" in dirPath that stores an array
// of SpiralCircleDB and stores it in the database
// the map BackRepoSpiralCircleid_atBckpTime_newID is updated accordingly
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSpiralCircleid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SpiralCircleDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SpiralCircle file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SpiralCircleDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SpiralCircleDBID_SpiralCircleDB
	for _, spiralcircleDB := range forRestore {

		spiralcircleDB_ID_atBackupTime := spiralcircleDB.ID
		spiralcircleDB.ID = 0
		_, err := backRepoSpiralCircle.db.Create(spiralcircleDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[spiralcircleDB.ID] = spiralcircleDB
		BackRepoSpiralCircleid_atBckpTime_newID[spiralcircleDB_ID_atBackupTime] = spiralcircleDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SpiralCircle file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SpiralCircle>id_atBckpTime_newID
// to compute new index
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) RestorePhaseTwo() {

	for _, spiralcircleDB := range backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB {

		// next line of code is to avert unused variable compilation error
		_ = spiralcircleDB

		// insertion point for reindexing pointers encoding
		// reindexing ShapeCategory field
		if spiralcircleDB.ShapeCategoryID.Int64 != 0 {
			spiralcircleDB.ShapeCategoryID.Int64 = int64(BackRepoShapeCategoryid_atBckpTime_newID[uint(spiralcircleDB.ShapeCategoryID.Int64)])
			spiralcircleDB.ShapeCategoryID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoSpiralCircle.db.Model(spiralcircleDB)
		_, err := db.Updates(*spiralcircleDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoSpiralCircle.ResetReversePointers commits all staged instances of SpiralCircle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralcircle := range backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCirclePtr {
		backRepoSpiralCircle.ResetReversePointersInstance(backRepo, idx, spiralcircle)
	}

	return
}

func (backRepoSpiralCircle *BackRepoSpiralCircleStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, spiralcircle *models.SpiralCircle) (Error error) {

	// fetch matching spiralcircleDB
	if spiralcircleDB, ok := backRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB[idx]; ok {
		_ = spiralcircleDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSpiralCircleid_atBckpTime_newID map[uint]uint
