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

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_SpiralRhombusGrid_sql sql.NullBool
var dummy_SpiralRhombusGrid_time time.Duration
var dummy_SpiralRhombusGrid_sort sort.Float64Slice

// SpiralRhombusGridAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model spiralrhombusgridAPI
type SpiralRhombusGridAPI struct {
	gorm.Model

	models.SpiralRhombusGrid_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SpiralRhombusGridPointersEncoding SpiralRhombusGridPointersEncoding
}

// SpiralRhombusGridPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SpiralRhombusGridPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ShapeCategory is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ShapeCategoryID sql.NullInt64

	// field RhombusGrid is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	RhombusGridID sql.NullInt64
}

// SpiralRhombusGridDB describes a spiralrhombusgrid in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model spiralrhombusgridDB
type SpiralRhombusGridDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field spiralrhombusgridDB.Name
	Name_Data sql.NullString

	// Declation for basic field spiralrhombusgridDB.IsDisplayed
	// provide the sql storage for the boolan
	IsDisplayed_Data sql.NullBool
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SpiralRhombusGridPointersEncoding
}

// SpiralRhombusGridDBs arrays spiralrhombusgridDBs
// swagger:response spiralrhombusgridDBsResponse
type SpiralRhombusGridDBs []SpiralRhombusGridDB

// SpiralRhombusGridDBResponse provides response
// swagger:response spiralrhombusgridDBResponse
type SpiralRhombusGridDBResponse struct {
	SpiralRhombusGridDB
}

// SpiralRhombusGridWOP is a SpiralRhombusGrid without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SpiralRhombusGridWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsDisplayed bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var SpiralRhombusGrid_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsDisplayed",
}

type BackRepoSpiralRhombusGridStruct struct {
	// stores SpiralRhombusGridDB according to their gorm ID
	Map_SpiralRhombusGridDBID_SpiralRhombusGridDB map[uint]*SpiralRhombusGridDB

	// stores SpiralRhombusGridDB ID according to SpiralRhombusGrid address
	Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID map[*models.SpiralRhombusGrid]uint

	// stores SpiralRhombusGrid according to their gorm ID
	Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr map[uint]*models.SpiralRhombusGrid

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSpiralRhombusGrid.stage
	return
}

func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) GetDB() *gorm.DB {
	return backRepoSpiralRhombusGrid.db
}

// GetSpiralRhombusGridDBFromSpiralRhombusGridPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) GetSpiralRhombusGridDBFromSpiralRhombusGridPtr(spiralrhombusgrid *models.SpiralRhombusGrid) (spiralrhombusgridDB *SpiralRhombusGridDB) {
	id := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid]
	spiralrhombusgridDB = backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[id]
	return
}

// BackRepoSpiralRhombusGrid.CommitPhaseOne commits all staged instances of SpiralRhombusGrid to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for spiralrhombusgrid := range stage.SpiralRhombusGrids {
		backRepoSpiralRhombusGrid.CommitPhaseOneInstance(spiralrhombusgrid)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, spiralrhombusgrid := range backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr {
		if _, ok := stage.SpiralRhombusGrids[spiralrhombusgrid]; !ok {
			backRepoSpiralRhombusGrid.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSpiralRhombusGrid.CommitDeleteInstance commits deletion of SpiralRhombusGrid to the BackRepo
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CommitDeleteInstance(id uint) (Error error) {

	spiralrhombusgrid := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[id]

	// spiralrhombusgrid is not staged anymore, remove spiralrhombusgridDB
	spiralrhombusgridDB := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[id]
	query := backRepoSpiralRhombusGrid.db.Unscoped().Delete(&spiralrhombusgridDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID, spiralrhombusgrid)
	delete(backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr, id)
	delete(backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB, id)

	return
}

// BackRepoSpiralRhombusGrid.CommitPhaseOneInstance commits spiralrhombusgrid staged instances of SpiralRhombusGrid to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CommitPhaseOneInstance(spiralrhombusgrid *models.SpiralRhombusGrid) (Error error) {

	// check if the spiralrhombusgrid is not commited yet
	if _, ok := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid]; ok {
		return
	}

	// initiate spiralrhombusgrid
	var spiralrhombusgridDB SpiralRhombusGridDB
	spiralrhombusgridDB.CopyBasicFieldsFromSpiralRhombusGrid(spiralrhombusgrid)

	query := backRepoSpiralRhombusGrid.db.Create(&spiralrhombusgridDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid] = spiralrhombusgridDB.ID
	backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID] = spiralrhombusgrid
	backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[spiralrhombusgridDB.ID] = &spiralrhombusgridDB

	return
}

// BackRepoSpiralRhombusGrid.CommitPhaseTwo commits all staged instances of SpiralRhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralrhombusgrid := range backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr {
		backRepoSpiralRhombusGrid.CommitPhaseTwoInstance(backRepo, idx, spiralrhombusgrid)
	}

	return
}

// BackRepoSpiralRhombusGrid.CommitPhaseTwoInstance commits {{structname }} of models.SpiralRhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, spiralrhombusgrid *models.SpiralRhombusGrid) (Error error) {

	// fetch matching spiralrhombusgridDB
	if spiralrhombusgridDB, ok := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[idx]; ok {

		spiralrhombusgridDB.CopyBasicFieldsFromSpiralRhombusGrid(spiralrhombusgrid)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value spiralrhombusgrid.ShapeCategory translates to updating the spiralrhombusgrid.ShapeCategoryID
		spiralrhombusgridDB.ShapeCategoryID.Valid = true // allow for a 0 value (nil association)
		if spiralrhombusgrid.ShapeCategory != nil {
			if ShapeCategoryId, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[spiralrhombusgrid.ShapeCategory]; ok {
				spiralrhombusgridDB.ShapeCategoryID.Int64 = int64(ShapeCategoryId)
				spiralrhombusgridDB.ShapeCategoryID.Valid = true
			}
		} else {
			spiralrhombusgridDB.ShapeCategoryID.Int64 = 0
			spiralrhombusgridDB.ShapeCategoryID.Valid = true
		}

		// commit pointer value spiralrhombusgrid.RhombusGrid translates to updating the spiralrhombusgrid.RhombusGridID
		spiralrhombusgridDB.RhombusGridID.Valid = true // allow for a 0 value (nil association)
		if spiralrhombusgrid.RhombusGrid != nil {
			if RhombusGridId, ok := backRepo.BackRepoRhombusGrid.Map_RhombusGridPtr_RhombusGridDBID[spiralrhombusgrid.RhombusGrid]; ok {
				spiralrhombusgridDB.RhombusGridID.Int64 = int64(RhombusGridId)
				spiralrhombusgridDB.RhombusGridID.Valid = true
			}
		} else {
			spiralrhombusgridDB.RhombusGridID.Int64 = 0
			spiralrhombusgridDB.RhombusGridID.Valid = true
		}

		query := backRepoSpiralRhombusGrid.db.Save(&spiralrhombusgridDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SpiralRhombusGrid intance %s", spiralrhombusgrid.Name))
		return err
	}

	return
}

// BackRepoSpiralRhombusGrid.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CheckoutPhaseOne() (Error error) {

	spiralrhombusgridDBArray := make([]SpiralRhombusGridDB, 0)
	query := backRepoSpiralRhombusGrid.db.Find(&spiralrhombusgridDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	spiralrhombusgridInstancesToBeRemovedFromTheStage := make(map[*models.SpiralRhombusGrid]any)
	for key, value := range backRepoSpiralRhombusGrid.stage.SpiralRhombusGrids {
		spiralrhombusgridInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, spiralrhombusgridDB := range spiralrhombusgridDBArray {
		backRepoSpiralRhombusGrid.CheckoutPhaseOneInstance(&spiralrhombusgridDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		spiralrhombusgrid, ok := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID]
		if ok {
			delete(spiralrhombusgridInstancesToBeRemovedFromTheStage, spiralrhombusgrid)
		}
	}

	// remove from stage and back repo's 3 maps all spiralrhombusgrids that are not in the checkout
	for spiralrhombusgrid := range spiralrhombusgridInstancesToBeRemovedFromTheStage {
		spiralrhombusgrid.Unstage(backRepoSpiralRhombusGrid.GetStage())

		// remove instance from the back repo 3 maps
		spiralrhombusgridID := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid]
		delete(backRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID, spiralrhombusgrid)
		delete(backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB, spiralrhombusgridID)
		delete(backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr, spiralrhombusgridID)
	}

	return
}

// CheckoutPhaseOneInstance takes a spiralrhombusgridDB that has been found in the DB, updates the backRepo and stages the
// models version of the spiralrhombusgridDB
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CheckoutPhaseOneInstance(spiralrhombusgridDB *SpiralRhombusGridDB) (Error error) {

	spiralrhombusgrid, ok := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID]
	if !ok {
		spiralrhombusgrid = new(models.SpiralRhombusGrid)

		backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID] = spiralrhombusgrid
		backRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid] = spiralrhombusgridDB.ID

		// append model store with the new element
		spiralrhombusgrid.Name = spiralrhombusgridDB.Name_Data.String
		spiralrhombusgrid.Stage(backRepoSpiralRhombusGrid.GetStage())
	}
	spiralrhombusgridDB.CopyBasicFieldsToSpiralRhombusGrid(spiralrhombusgrid)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	spiralrhombusgrid.Stage(backRepoSpiralRhombusGrid.GetStage())

	// preserve pointer to spiralrhombusgridDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SpiralRhombusGridDBID_SpiralRhombusGridDB)[spiralrhombusgridDB hold variable pointers
	spiralrhombusgridDB_Data := *spiralrhombusgridDB
	preservedPtrToSpiralRhombusGrid := &spiralrhombusgridDB_Data
	backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[spiralrhombusgridDB.ID] = preservedPtrToSpiralRhombusGrid

	return
}

// BackRepoSpiralRhombusGrid.CheckoutPhaseTwo Checkouts all staged instances of SpiralRhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, spiralrhombusgridDB := range backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB {
		backRepoSpiralRhombusGrid.CheckoutPhaseTwoInstance(backRepo, spiralrhombusgridDB)
	}
	return
}

// BackRepoSpiralRhombusGrid.CheckoutPhaseTwoInstance Checkouts staged instances of SpiralRhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, spiralrhombusgridDB *SpiralRhombusGridDB) (Error error) {

	spiralrhombusgrid := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr[spiralrhombusgridDB.ID]

	spiralrhombusgridDB.DecodePointers(backRepo, spiralrhombusgrid)

	return
}

func (spiralrhombusgridDB *SpiralRhombusGridDB) DecodePointers(backRepo *BackRepoStruct, spiralrhombusgrid *models.SpiralRhombusGrid) {

	// insertion point for checkout of pointer encoding
	// ShapeCategory field
	spiralrhombusgrid.ShapeCategory = nil
	if spiralrhombusgridDB.ShapeCategoryID.Int64 != 0 {
		spiralrhombusgrid.ShapeCategory = backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[uint(spiralrhombusgridDB.ShapeCategoryID.Int64)]
	}
	// RhombusGrid field
	spiralrhombusgrid.RhombusGrid = nil
	if spiralrhombusgridDB.RhombusGridID.Int64 != 0 {
		spiralrhombusgrid.RhombusGrid = backRepo.BackRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridPtr[uint(spiralrhombusgridDB.RhombusGridID.Int64)]
	}
	return
}

// CommitSpiralRhombusGrid allows commit of a single spiralrhombusgrid (if already staged)
func (backRepo *BackRepoStruct) CommitSpiralRhombusGrid(spiralrhombusgrid *models.SpiralRhombusGrid) {
	backRepo.BackRepoSpiralRhombusGrid.CommitPhaseOneInstance(spiralrhombusgrid)
	if id, ok := backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid]; ok {
		backRepo.BackRepoSpiralRhombusGrid.CommitPhaseTwoInstance(backRepo, id, spiralrhombusgrid)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSpiralRhombusGrid allows checkout of a single spiralrhombusgrid (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSpiralRhombusGrid(spiralrhombusgrid *models.SpiralRhombusGrid) {
	// check if the spiralrhombusgrid is staged
	if _, ok := backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid]; ok {

		if id, ok := backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID[spiralrhombusgrid]; ok {
			var spiralrhombusgridDB SpiralRhombusGridDB
			spiralrhombusgridDB.ID = id

			if err := backRepo.BackRepoSpiralRhombusGrid.db.First(&spiralrhombusgridDB, id).Error; err != nil {
				log.Fatalln("CheckoutSpiralRhombusGrid : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSpiralRhombusGrid.CheckoutPhaseOneInstance(&spiralrhombusgridDB)
			backRepo.BackRepoSpiralRhombusGrid.CheckoutPhaseTwoInstance(backRepo, &spiralrhombusgridDB)
		}
	}
}

// CopyBasicFieldsFromSpiralRhombusGrid
func (spiralrhombusgridDB *SpiralRhombusGridDB) CopyBasicFieldsFromSpiralRhombusGrid(spiralrhombusgrid *models.SpiralRhombusGrid) {
	// insertion point for fields commit

	spiralrhombusgridDB.Name_Data.String = spiralrhombusgrid.Name
	spiralrhombusgridDB.Name_Data.Valid = true

	spiralrhombusgridDB.IsDisplayed_Data.Bool = spiralrhombusgrid.IsDisplayed
	spiralrhombusgridDB.IsDisplayed_Data.Valid = true
}

// CopyBasicFieldsFromSpiralRhombusGrid_WOP
func (spiralrhombusgridDB *SpiralRhombusGridDB) CopyBasicFieldsFromSpiralRhombusGrid_WOP(spiralrhombusgrid *models.SpiralRhombusGrid_WOP) {
	// insertion point for fields commit

	spiralrhombusgridDB.Name_Data.String = spiralrhombusgrid.Name
	spiralrhombusgridDB.Name_Data.Valid = true

	spiralrhombusgridDB.IsDisplayed_Data.Bool = spiralrhombusgrid.IsDisplayed
	spiralrhombusgridDB.IsDisplayed_Data.Valid = true
}

// CopyBasicFieldsFromSpiralRhombusGridWOP
func (spiralrhombusgridDB *SpiralRhombusGridDB) CopyBasicFieldsFromSpiralRhombusGridWOP(spiralrhombusgrid *SpiralRhombusGridWOP) {
	// insertion point for fields commit

	spiralrhombusgridDB.Name_Data.String = spiralrhombusgrid.Name
	spiralrhombusgridDB.Name_Data.Valid = true

	spiralrhombusgridDB.IsDisplayed_Data.Bool = spiralrhombusgrid.IsDisplayed
	spiralrhombusgridDB.IsDisplayed_Data.Valid = true
}

// CopyBasicFieldsToSpiralRhombusGrid
func (spiralrhombusgridDB *SpiralRhombusGridDB) CopyBasicFieldsToSpiralRhombusGrid(spiralrhombusgrid *models.SpiralRhombusGrid) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralrhombusgrid.Name = spiralrhombusgridDB.Name_Data.String
	spiralrhombusgrid.IsDisplayed = spiralrhombusgridDB.IsDisplayed_Data.Bool
}

// CopyBasicFieldsToSpiralRhombusGrid_WOP
func (spiralrhombusgridDB *SpiralRhombusGridDB) CopyBasicFieldsToSpiralRhombusGrid_WOP(spiralrhombusgrid *models.SpiralRhombusGrid_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	spiralrhombusgrid.Name = spiralrhombusgridDB.Name_Data.String
	spiralrhombusgrid.IsDisplayed = spiralrhombusgridDB.IsDisplayed_Data.Bool
}

// CopyBasicFieldsToSpiralRhombusGridWOP
func (spiralrhombusgridDB *SpiralRhombusGridDB) CopyBasicFieldsToSpiralRhombusGridWOP(spiralrhombusgrid *SpiralRhombusGridWOP) {
	spiralrhombusgrid.ID = int(spiralrhombusgridDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	spiralrhombusgrid.Name = spiralrhombusgridDB.Name_Data.String
	spiralrhombusgrid.IsDisplayed = spiralrhombusgridDB.IsDisplayed_Data.Bool
}

// Backup generates a json file from a slice of all SpiralRhombusGridDB instances in the backrepo
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SpiralRhombusGridDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralRhombusGridDB, 0)
	for _, spiralrhombusgridDB := range backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB {
		forBackup = append(forBackup, spiralrhombusgridDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SpiralRhombusGrid ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SpiralRhombusGrid file", err.Error())
	}
}

// Backup generates a json file from a slice of all SpiralRhombusGridDB instances in the backrepo
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SpiralRhombusGridDB, 0)
	for _, spiralrhombusgridDB := range backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB {
		forBackup = append(forBackup, spiralrhombusgridDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SpiralRhombusGrid")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SpiralRhombusGrid_Fields, -1)
	for _, spiralrhombusgridDB := range forBackup {

		var spiralrhombusgridWOP SpiralRhombusGridWOP
		spiralrhombusgridDB.CopyBasicFieldsToSpiralRhombusGridWOP(&spiralrhombusgridWOP)

		row := sh.AddRow()
		row.WriteStruct(&spiralrhombusgridWOP, -1)
	}
}

// RestoreXL from the "SpiralRhombusGrid" sheet all SpiralRhombusGridDB instances
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSpiralRhombusGridid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SpiralRhombusGrid"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSpiralRhombusGrid.rowVisitorSpiralRhombusGrid)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) rowVisitorSpiralRhombusGrid(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var spiralrhombusgridWOP SpiralRhombusGridWOP
		row.ReadStruct(&spiralrhombusgridWOP)

		// add the unmarshalled struct to the stage
		spiralrhombusgridDB := new(SpiralRhombusGridDB)
		spiralrhombusgridDB.CopyBasicFieldsFromSpiralRhombusGridWOP(&spiralrhombusgridWOP)

		spiralrhombusgridDB_ID_atBackupTime := spiralrhombusgridDB.ID
		spiralrhombusgridDB.ID = 0
		query := backRepoSpiralRhombusGrid.db.Create(spiralrhombusgridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[spiralrhombusgridDB.ID] = spiralrhombusgridDB
		BackRepoSpiralRhombusGridid_atBckpTime_newID[spiralrhombusgridDB_ID_atBackupTime] = spiralrhombusgridDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SpiralRhombusGridDB.json" in dirPath that stores an array
// of SpiralRhombusGridDB and stores it in the database
// the map BackRepoSpiralRhombusGridid_atBckpTime_newID is updated accordingly
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSpiralRhombusGridid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SpiralRhombusGridDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SpiralRhombusGrid file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SpiralRhombusGridDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SpiralRhombusGridDBID_SpiralRhombusGridDB
	for _, spiralrhombusgridDB := range forRestore {

		spiralrhombusgridDB_ID_atBackupTime := spiralrhombusgridDB.ID
		spiralrhombusgridDB.ID = 0
		query := backRepoSpiralRhombusGrid.db.Create(spiralrhombusgridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[spiralrhombusgridDB.ID] = spiralrhombusgridDB
		BackRepoSpiralRhombusGridid_atBckpTime_newID[spiralrhombusgridDB_ID_atBackupTime] = spiralrhombusgridDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SpiralRhombusGrid file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SpiralRhombusGrid>id_atBckpTime_newID
// to compute new index
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) RestorePhaseTwo() {

	for _, spiralrhombusgridDB := range backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB {

		// next line of code is to avert unused variable compilation error
		_ = spiralrhombusgridDB

		// insertion point for reindexing pointers encoding
		// reindexing ShapeCategory field
		if spiralrhombusgridDB.ShapeCategoryID.Int64 != 0 {
			spiralrhombusgridDB.ShapeCategoryID.Int64 = int64(BackRepoShapeCategoryid_atBckpTime_newID[uint(spiralrhombusgridDB.ShapeCategoryID.Int64)])
			spiralrhombusgridDB.ShapeCategoryID.Valid = true
		}

		// reindexing RhombusGrid field
		if spiralrhombusgridDB.RhombusGridID.Int64 != 0 {
			spiralrhombusgridDB.RhombusGridID.Int64 = int64(BackRepoRhombusGridid_atBckpTime_newID[uint(spiralrhombusgridDB.RhombusGridID.Int64)])
			spiralrhombusgridDB.RhombusGridID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoSpiralRhombusGrid.db.Model(spiralrhombusgridDB).Updates(*spiralrhombusgridDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSpiralRhombusGrid.ResetReversePointers commits all staged instances of SpiralRhombusGrid to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, spiralrhombusgrid := range backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr {
		backRepoSpiralRhombusGrid.ResetReversePointersInstance(backRepo, idx, spiralrhombusgrid)
	}

	return
}

func (backRepoSpiralRhombusGrid *BackRepoSpiralRhombusGridStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, spiralrhombusgrid *models.SpiralRhombusGrid) (Error error) {

	// fetch matching spiralrhombusgridDB
	if spiralrhombusgridDB, ok := backRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB[idx]; ok {
		_ = spiralrhombusgridDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSpiralRhombusGridid_atBckpTime_newID map[uint]uint