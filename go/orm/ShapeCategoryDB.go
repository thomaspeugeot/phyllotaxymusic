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
var dummy_ShapeCategory_sql sql.NullBool
var dummy_ShapeCategory_time time.Duration
var dummy_ShapeCategory_sort sort.Float64Slice

// ShapeCategoryAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model shapecategoryAPI
type ShapeCategoryAPI struct {
	gorm.Model

	models.ShapeCategory_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ShapeCategoryPointersEncoding ShapeCategoryPointersEncoding
}

// ShapeCategoryPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ShapeCategoryPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// ShapeCategoryDB describes a shapecategory in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model shapecategoryDB
type ShapeCategoryDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field shapecategoryDB.Name
	Name_Data sql.NullString

	// Declation for basic field shapecategoryDB.IsExpanded
	// provide the sql storage for the boolan
	IsExpanded_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ShapeCategoryPointersEncoding
}

// ShapeCategoryDBs arrays shapecategoryDBs
// swagger:response shapecategoryDBsResponse
type ShapeCategoryDBs []ShapeCategoryDB

// ShapeCategoryDBResponse provides response
// swagger:response shapecategoryDBResponse
type ShapeCategoryDBResponse struct {
	ShapeCategoryDB
}

// ShapeCategoryWOP is a ShapeCategory without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ShapeCategoryWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	IsExpanded bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var ShapeCategory_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"IsExpanded",
}

type BackRepoShapeCategoryStruct struct {
	// stores ShapeCategoryDB according to their gorm ID
	Map_ShapeCategoryDBID_ShapeCategoryDB map[uint]*ShapeCategoryDB

	// stores ShapeCategoryDB ID according to ShapeCategory address
	Map_ShapeCategoryPtr_ShapeCategoryDBID map[*models.ShapeCategory]uint

	// stores ShapeCategory according to their gorm ID
	Map_ShapeCategoryDBID_ShapeCategoryPtr map[uint]*models.ShapeCategory

	db db.DBInterface

	stage *models.Stage
}

func (backRepoShapeCategory *BackRepoShapeCategoryStruct) GetStage() (stage *models.Stage) {
	stage = backRepoShapeCategory.stage
	return
}

func (backRepoShapeCategory *BackRepoShapeCategoryStruct) GetDB() db.DBInterface {
	return backRepoShapeCategory.db
}

// GetShapeCategoryDBFromShapeCategoryPtr is a handy function to access the back repo instance from the stage instance
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) GetShapeCategoryDBFromShapeCategoryPtr(shapecategory *models.ShapeCategory) (shapecategoryDB *ShapeCategoryDB) {
	id := backRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory]
	shapecategoryDB = backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[id]
	return
}

// BackRepoShapeCategory.CommitPhaseOne commits all staged instances of ShapeCategory to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var shapecategorys []*models.ShapeCategory
	for shapecategory := range stage.ShapeCategorys {
		shapecategorys = append(shapecategorys, shapecategory)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(shapecategorys, func(i, j int) bool {
		return stage.ShapeCategoryMap_Staged_Order[shapecategorys[i]] < stage.ShapeCategoryMap_Staged_Order[shapecategorys[j]]
	})

	for _, shapecategory := range shapecategorys {
		backRepoShapeCategory.CommitPhaseOneInstance(shapecategory)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, shapecategory := range backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr {
		if _, ok := stage.ShapeCategorys[shapecategory]; !ok {
			backRepoShapeCategory.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoShapeCategory.CommitDeleteInstance commits deletion of ShapeCategory to the BackRepo
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CommitDeleteInstance(id uint) (Error error) {

	shapecategory := backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[id]

	// shapecategory is not staged anymore, remove shapecategoryDB
	shapecategoryDB := backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[id]
	db, _ := backRepoShapeCategory.db.Unscoped()
	_, err := db.Delete(shapecategoryDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID, shapecategory)
	delete(backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr, id)
	delete(backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB, id)

	return
}

// BackRepoShapeCategory.CommitPhaseOneInstance commits shapecategory staged instances of ShapeCategory to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CommitPhaseOneInstance(shapecategory *models.ShapeCategory) (Error error) {

	// check if the shapecategory is not commited yet
	if _, ok := backRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory]; ok {
		return
	}

	// initiate shapecategory
	var shapecategoryDB ShapeCategoryDB
	shapecategoryDB.CopyBasicFieldsFromShapeCategory(shapecategory)

	_, err := backRepoShapeCategory.db.Create(&shapecategoryDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory] = shapecategoryDB.ID
	backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID] = shapecategory
	backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[shapecategoryDB.ID] = &shapecategoryDB

	return
}

// BackRepoShapeCategory.CommitPhaseTwo commits all staged instances of ShapeCategory to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, shapecategory := range backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr {
		backRepoShapeCategory.CommitPhaseTwoInstance(backRepo, idx, shapecategory)
	}

	return
}

// BackRepoShapeCategory.CommitPhaseTwoInstance commits {{structname }} of models.ShapeCategory to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, shapecategory *models.ShapeCategory) (Error error) {

	// fetch matching shapecategoryDB
	if shapecategoryDB, ok := backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[idx]; ok {

		shapecategoryDB.CopyBasicFieldsFromShapeCategory(shapecategory)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoShapeCategory.db.Save(shapecategoryDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ShapeCategory intance %s", shapecategory.Name))
		return err
	}

	return
}

// BackRepoShapeCategory.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CheckoutPhaseOne() (Error error) {

	shapecategoryDBArray := make([]ShapeCategoryDB, 0)
	_, err := backRepoShapeCategory.db.Find(&shapecategoryDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	shapecategoryInstancesToBeRemovedFromTheStage := make(map[*models.ShapeCategory]any)
	for key, value := range backRepoShapeCategory.stage.ShapeCategorys {
		shapecategoryInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, shapecategoryDB := range shapecategoryDBArray {
		backRepoShapeCategory.CheckoutPhaseOneInstance(&shapecategoryDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		shapecategory, ok := backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID]
		if ok {
			delete(shapecategoryInstancesToBeRemovedFromTheStage, shapecategory)
		}
	}

	// remove from stage and back repo's 3 maps all shapecategorys that are not in the checkout
	for shapecategory := range shapecategoryInstancesToBeRemovedFromTheStage {
		shapecategory.Unstage(backRepoShapeCategory.GetStage())

		// remove instance from the back repo 3 maps
		shapecategoryID := backRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory]
		delete(backRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID, shapecategory)
		delete(backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB, shapecategoryID)
		delete(backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr, shapecategoryID)
	}

	return
}

// CheckoutPhaseOneInstance takes a shapecategoryDB that has been found in the DB, updates the backRepo and stages the
// models version of the shapecategoryDB
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CheckoutPhaseOneInstance(shapecategoryDB *ShapeCategoryDB) (Error error) {

	shapecategory, ok := backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID]
	if !ok {
		shapecategory = new(models.ShapeCategory)

		backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID] = shapecategory
		backRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory] = shapecategoryDB.ID

		// append model store with the new element
		shapecategory.Name = shapecategoryDB.Name_Data.String
		shapecategory.Stage(backRepoShapeCategory.GetStage())
	}
	shapecategoryDB.CopyBasicFieldsToShapeCategory(shapecategory)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	shapecategory.Stage(backRepoShapeCategory.GetStage())

	// preserve pointer to shapecategoryDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ShapeCategoryDBID_ShapeCategoryDB)[shapecategoryDB hold variable pointers
	shapecategoryDB_Data := *shapecategoryDB
	preservedPtrToShapeCategory := &shapecategoryDB_Data
	backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[shapecategoryDB.ID] = preservedPtrToShapeCategory

	return
}

// BackRepoShapeCategory.CheckoutPhaseTwo Checkouts all staged instances of ShapeCategory to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, shapecategoryDB := range backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB {
		backRepoShapeCategory.CheckoutPhaseTwoInstance(backRepo, shapecategoryDB)
	}
	return
}

// BackRepoShapeCategory.CheckoutPhaseTwoInstance Checkouts staged instances of ShapeCategory to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, shapecategoryDB *ShapeCategoryDB) (Error error) {

	shapecategory := backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr[shapecategoryDB.ID]

	shapecategoryDB.DecodePointers(backRepo, shapecategory)

	return
}

func (shapecategoryDB *ShapeCategoryDB) DecodePointers(backRepo *BackRepoStruct, shapecategory *models.ShapeCategory) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitShapeCategory allows commit of a single shapecategory (if already staged)
func (backRepo *BackRepoStruct) CommitShapeCategory(shapecategory *models.ShapeCategory) {
	backRepo.BackRepoShapeCategory.CommitPhaseOneInstance(shapecategory)
	if id, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory]; ok {
		backRepo.BackRepoShapeCategory.CommitPhaseTwoInstance(backRepo, id, shapecategory)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitShapeCategory allows checkout of a single shapecategory (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutShapeCategory(shapecategory *models.ShapeCategory) {
	// check if the shapecategory is staged
	if _, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory]; ok {

		if id, ok := backRepo.BackRepoShapeCategory.Map_ShapeCategoryPtr_ShapeCategoryDBID[shapecategory]; ok {
			var shapecategoryDB ShapeCategoryDB
			shapecategoryDB.ID = id

			if _, err := backRepo.BackRepoShapeCategory.db.First(&shapecategoryDB, id); err != nil {
				log.Fatalln("CheckoutShapeCategory : Problem with getting object with id:", id)
			}
			backRepo.BackRepoShapeCategory.CheckoutPhaseOneInstance(&shapecategoryDB)
			backRepo.BackRepoShapeCategory.CheckoutPhaseTwoInstance(backRepo, &shapecategoryDB)
		}
	}
}

// CopyBasicFieldsFromShapeCategory
func (shapecategoryDB *ShapeCategoryDB) CopyBasicFieldsFromShapeCategory(shapecategory *models.ShapeCategory) {
	// insertion point for fields commit

	shapecategoryDB.Name_Data.String = shapecategory.Name
	shapecategoryDB.Name_Data.Valid = true

	shapecategoryDB.IsExpanded_Data.Bool = shapecategory.IsExpanded
	shapecategoryDB.IsExpanded_Data.Valid = true
}

// CopyBasicFieldsFromShapeCategory_WOP
func (shapecategoryDB *ShapeCategoryDB) CopyBasicFieldsFromShapeCategory_WOP(shapecategory *models.ShapeCategory_WOP) {
	// insertion point for fields commit

	shapecategoryDB.Name_Data.String = shapecategory.Name
	shapecategoryDB.Name_Data.Valid = true

	shapecategoryDB.IsExpanded_Data.Bool = shapecategory.IsExpanded
	shapecategoryDB.IsExpanded_Data.Valid = true
}

// CopyBasicFieldsFromShapeCategoryWOP
func (shapecategoryDB *ShapeCategoryDB) CopyBasicFieldsFromShapeCategoryWOP(shapecategory *ShapeCategoryWOP) {
	// insertion point for fields commit

	shapecategoryDB.Name_Data.String = shapecategory.Name
	shapecategoryDB.Name_Data.Valid = true

	shapecategoryDB.IsExpanded_Data.Bool = shapecategory.IsExpanded
	shapecategoryDB.IsExpanded_Data.Valid = true
}

// CopyBasicFieldsToShapeCategory
func (shapecategoryDB *ShapeCategoryDB) CopyBasicFieldsToShapeCategory(shapecategory *models.ShapeCategory) {
	// insertion point for checkout of basic fields (back repo to stage)
	shapecategory.Name = shapecategoryDB.Name_Data.String
	shapecategory.IsExpanded = shapecategoryDB.IsExpanded_Data.Bool
}

// CopyBasicFieldsToShapeCategory_WOP
func (shapecategoryDB *ShapeCategoryDB) CopyBasicFieldsToShapeCategory_WOP(shapecategory *models.ShapeCategory_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	shapecategory.Name = shapecategoryDB.Name_Data.String
	shapecategory.IsExpanded = shapecategoryDB.IsExpanded_Data.Bool
}

// CopyBasicFieldsToShapeCategoryWOP
func (shapecategoryDB *ShapeCategoryDB) CopyBasicFieldsToShapeCategoryWOP(shapecategory *ShapeCategoryWOP) {
	shapecategory.ID = int(shapecategoryDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	shapecategory.Name = shapecategoryDB.Name_Data.String
	shapecategory.IsExpanded = shapecategoryDB.IsExpanded_Data.Bool
}

// Backup generates a json file from a slice of all ShapeCategoryDB instances in the backrepo
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ShapeCategoryDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ShapeCategoryDB, 0)
	for _, shapecategoryDB := range backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB {
		forBackup = append(forBackup, shapecategoryDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ShapeCategory ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ShapeCategory file", err.Error())
	}
}

// Backup generates a json file from a slice of all ShapeCategoryDB instances in the backrepo
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ShapeCategoryDB, 0)
	for _, shapecategoryDB := range backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB {
		forBackup = append(forBackup, shapecategoryDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ShapeCategory")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ShapeCategory_Fields, -1)
	for _, shapecategoryDB := range forBackup {

		var shapecategoryWOP ShapeCategoryWOP
		shapecategoryDB.CopyBasicFieldsToShapeCategoryWOP(&shapecategoryWOP)

		row := sh.AddRow()
		row.WriteStruct(&shapecategoryWOP, -1)
	}
}

// RestoreXL from the "ShapeCategory" sheet all ShapeCategoryDB instances
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoShapeCategoryid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ShapeCategory"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoShapeCategory.rowVisitorShapeCategory)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoShapeCategory *BackRepoShapeCategoryStruct) rowVisitorShapeCategory(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var shapecategoryWOP ShapeCategoryWOP
		row.ReadStruct(&shapecategoryWOP)

		// add the unmarshalled struct to the stage
		shapecategoryDB := new(ShapeCategoryDB)
		shapecategoryDB.CopyBasicFieldsFromShapeCategoryWOP(&shapecategoryWOP)

		shapecategoryDB_ID_atBackupTime := shapecategoryDB.ID
		shapecategoryDB.ID = 0
		_, err := backRepoShapeCategory.db.Create(shapecategoryDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[shapecategoryDB.ID] = shapecategoryDB
		BackRepoShapeCategoryid_atBckpTime_newID[shapecategoryDB_ID_atBackupTime] = shapecategoryDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ShapeCategoryDB.json" in dirPath that stores an array
// of ShapeCategoryDB and stores it in the database
// the map BackRepoShapeCategoryid_atBckpTime_newID is updated accordingly
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoShapeCategoryid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ShapeCategoryDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ShapeCategory file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ShapeCategoryDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ShapeCategoryDBID_ShapeCategoryDB
	for _, shapecategoryDB := range forRestore {

		shapecategoryDB_ID_atBackupTime := shapecategoryDB.ID
		shapecategoryDB.ID = 0
		_, err := backRepoShapeCategory.db.Create(shapecategoryDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[shapecategoryDB.ID] = shapecategoryDB
		BackRepoShapeCategoryid_atBckpTime_newID[shapecategoryDB_ID_atBackupTime] = shapecategoryDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ShapeCategory file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ShapeCategory>id_atBckpTime_newID
// to compute new index
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) RestorePhaseTwo() {

	for _, shapecategoryDB := range backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB {

		// next line of code is to avert unused variable compilation error
		_ = shapecategoryDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoShapeCategory.db.Model(shapecategoryDB)
		_, err := db.Updates(*shapecategoryDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoShapeCategory.ResetReversePointers commits all staged instances of ShapeCategory to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoShapeCategory *BackRepoShapeCategoryStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, shapecategory := range backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryPtr {
		backRepoShapeCategory.ResetReversePointersInstance(backRepo, idx, shapecategory)
	}

	return
}

func (backRepoShapeCategory *BackRepoShapeCategoryStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, shapecategory *models.ShapeCategory) (Error error) {

	// fetch matching shapecategoryDB
	if shapecategoryDB, ok := backRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB[idx]; ok {
		_ = shapecategoryDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoShapeCategoryid_atBckpTime_newID map[uint]uint
