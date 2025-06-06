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

	"github.com/fullstack-lang/gong/lib/table/go/db"
	"github.com/fullstack-lang/gong/lib/table/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_CellIcon_sql sql.NullBool
var dummy_CellIcon_time time.Duration
var dummy_CellIcon_sort sort.Float64Slice

// CellIconAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model celliconAPI
type CellIconAPI struct {
	gorm.Model

	models.CellIcon_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	CellIconPointersEncoding CellIconPointersEncoding
}

// CellIconPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type CellIconPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// CellIconDB describes a cellicon in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model celliconDB
type CellIconDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field celliconDB.Name
	Name_Data sql.NullString

	// Declation for basic field celliconDB.Icon
	Icon_Data sql.NullString

	// Declation for basic field celliconDB.NeedsConfirmation
	// provide the sql storage for the boolan
	NeedsConfirmation_Data sql.NullBool

	// Declation for basic field celliconDB.ConfirmationMessage
	ConfirmationMessage_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	CellIconPointersEncoding
}

// CellIconDBs arrays celliconDBs
// swagger:response celliconDBsResponse
type CellIconDBs []CellIconDB

// CellIconDBResponse provides response
// swagger:response celliconDBResponse
type CellIconDBResponse struct {
	CellIconDB
}

// CellIconWOP is a CellIcon without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type CellIconWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Icon string `xlsx:"2"`

	NeedsConfirmation bool `xlsx:"3"`

	ConfirmationMessage string `xlsx:"4"`
	// insertion for WOP pointer fields
}

var CellIcon_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Icon",
	"NeedsConfirmation",
	"ConfirmationMessage",
}

type BackRepoCellIconStruct struct {
	// stores CellIconDB according to their gorm ID
	Map_CellIconDBID_CellIconDB map[uint]*CellIconDB

	// stores CellIconDB ID according to CellIcon address
	Map_CellIconPtr_CellIconDBID map[*models.CellIcon]uint

	// stores CellIcon according to their gorm ID
	Map_CellIconDBID_CellIconPtr map[uint]*models.CellIcon

	db db.DBInterface

	stage *models.Stage
}

func (backRepoCellIcon *BackRepoCellIconStruct) GetStage() (stage *models.Stage) {
	stage = backRepoCellIcon.stage
	return
}

func (backRepoCellIcon *BackRepoCellIconStruct) GetDB() db.DBInterface {
	return backRepoCellIcon.db
}

// GetCellIconDBFromCellIconPtr is a handy function to access the back repo instance from the stage instance
func (backRepoCellIcon *BackRepoCellIconStruct) GetCellIconDBFromCellIconPtr(cellicon *models.CellIcon) (celliconDB *CellIconDB) {
	id := backRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon]
	celliconDB = backRepoCellIcon.Map_CellIconDBID_CellIconDB[id]
	return
}

// BackRepoCellIcon.CommitPhaseOne commits all staged instances of CellIcon to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCellIcon *BackRepoCellIconStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var cellicons []*models.CellIcon
	for cellicon := range stage.CellIcons {
		cellicons = append(cellicons, cellicon)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(cellicons, func(i, j int) bool {
		return stage.CellIconMap_Staged_Order[cellicons[i]] < stage.CellIconMap_Staged_Order[cellicons[j]]
	})

	for _, cellicon := range cellicons {
		backRepoCellIcon.CommitPhaseOneInstance(cellicon)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, cellicon := range backRepoCellIcon.Map_CellIconDBID_CellIconPtr {
		if _, ok := stage.CellIcons[cellicon]; !ok {
			backRepoCellIcon.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoCellIcon.CommitDeleteInstance commits deletion of CellIcon to the BackRepo
func (backRepoCellIcon *BackRepoCellIconStruct) CommitDeleteInstance(id uint) (Error error) {

	cellicon := backRepoCellIcon.Map_CellIconDBID_CellIconPtr[id]

	// cellicon is not staged anymore, remove celliconDB
	celliconDB := backRepoCellIcon.Map_CellIconDBID_CellIconDB[id]
	db, _ := backRepoCellIcon.db.Unscoped()
	_, err := db.Delete(celliconDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoCellIcon.Map_CellIconPtr_CellIconDBID, cellicon)
	delete(backRepoCellIcon.Map_CellIconDBID_CellIconPtr, id)
	delete(backRepoCellIcon.Map_CellIconDBID_CellIconDB, id)

	return
}

// BackRepoCellIcon.CommitPhaseOneInstance commits cellicon staged instances of CellIcon to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCellIcon *BackRepoCellIconStruct) CommitPhaseOneInstance(cellicon *models.CellIcon) (Error error) {

	// check if the cellicon is not commited yet
	if _, ok := backRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon]; ok {
		return
	}

	// initiate cellicon
	var celliconDB CellIconDB
	celliconDB.CopyBasicFieldsFromCellIcon(cellicon)

	_, err := backRepoCellIcon.db.Create(&celliconDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon] = celliconDB.ID
	backRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID] = cellicon
	backRepoCellIcon.Map_CellIconDBID_CellIconDB[celliconDB.ID] = &celliconDB

	return
}

// BackRepoCellIcon.CommitPhaseTwo commits all staged instances of CellIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellIcon *BackRepoCellIconStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, cellicon := range backRepoCellIcon.Map_CellIconDBID_CellIconPtr {
		backRepoCellIcon.CommitPhaseTwoInstance(backRepo, idx, cellicon)
	}

	return
}

// BackRepoCellIcon.CommitPhaseTwoInstance commits {{structname }} of models.CellIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellIcon *BackRepoCellIconStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, cellicon *models.CellIcon) (Error error) {

	// fetch matching celliconDB
	if celliconDB, ok := backRepoCellIcon.Map_CellIconDBID_CellIconDB[idx]; ok {

		celliconDB.CopyBasicFieldsFromCellIcon(cellicon)

		// insertion point for translating pointers encodings into actual pointers
		_, err := backRepoCellIcon.db.Save(celliconDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown CellIcon intance %s", cellicon.Name))
		return err
	}

	return
}

// BackRepoCellIcon.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoCellIcon *BackRepoCellIconStruct) CheckoutPhaseOne() (Error error) {

	celliconDBArray := make([]CellIconDB, 0)
	_, err := backRepoCellIcon.db.Find(&celliconDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	celliconInstancesToBeRemovedFromTheStage := make(map[*models.CellIcon]any)
	for key, value := range backRepoCellIcon.stage.CellIcons {
		celliconInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, celliconDB := range celliconDBArray {
		backRepoCellIcon.CheckoutPhaseOneInstance(&celliconDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		cellicon, ok := backRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID]
		if ok {
			delete(celliconInstancesToBeRemovedFromTheStage, cellicon)
		}
	}

	// remove from stage and back repo's 3 maps all cellicons that are not in the checkout
	for cellicon := range celliconInstancesToBeRemovedFromTheStage {
		cellicon.Unstage(backRepoCellIcon.GetStage())

		// remove instance from the back repo 3 maps
		celliconID := backRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon]
		delete(backRepoCellIcon.Map_CellIconPtr_CellIconDBID, cellicon)
		delete(backRepoCellIcon.Map_CellIconDBID_CellIconDB, celliconID)
		delete(backRepoCellIcon.Map_CellIconDBID_CellIconPtr, celliconID)
	}

	return
}

// CheckoutPhaseOneInstance takes a celliconDB that has been found in the DB, updates the backRepo and stages the
// models version of the celliconDB
func (backRepoCellIcon *BackRepoCellIconStruct) CheckoutPhaseOneInstance(celliconDB *CellIconDB) (Error error) {

	cellicon, ok := backRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID]
	if !ok {
		cellicon = new(models.CellIcon)

		backRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID] = cellicon
		backRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon] = celliconDB.ID

		// append model store with the new element
		cellicon.Name = celliconDB.Name_Data.String
		cellicon.Stage(backRepoCellIcon.GetStage())
	}
	celliconDB.CopyBasicFieldsToCellIcon(cellicon)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	cellicon.Stage(backRepoCellIcon.GetStage())

	// preserve pointer to celliconDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_CellIconDBID_CellIconDB)[celliconDB hold variable pointers
	celliconDB_Data := *celliconDB
	preservedPtrToCellIcon := &celliconDB_Data
	backRepoCellIcon.Map_CellIconDBID_CellIconDB[celliconDB.ID] = preservedPtrToCellIcon

	return
}

// BackRepoCellIcon.CheckoutPhaseTwo Checkouts all staged instances of CellIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellIcon *BackRepoCellIconStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, celliconDB := range backRepoCellIcon.Map_CellIconDBID_CellIconDB {
		backRepoCellIcon.CheckoutPhaseTwoInstance(backRepo, celliconDB)
	}
	return
}

// BackRepoCellIcon.CheckoutPhaseTwoInstance Checkouts staged instances of CellIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellIcon *BackRepoCellIconStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, celliconDB *CellIconDB) (Error error) {

	cellicon := backRepoCellIcon.Map_CellIconDBID_CellIconPtr[celliconDB.ID]

	celliconDB.DecodePointers(backRepo, cellicon)

	return
}

func (celliconDB *CellIconDB) DecodePointers(backRepo *BackRepoStruct, cellicon *models.CellIcon) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitCellIcon allows commit of a single cellicon (if already staged)
func (backRepo *BackRepoStruct) CommitCellIcon(cellicon *models.CellIcon) {
	backRepo.BackRepoCellIcon.CommitPhaseOneInstance(cellicon)
	if id, ok := backRepo.BackRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon]; ok {
		backRepo.BackRepoCellIcon.CommitPhaseTwoInstance(backRepo, id, cellicon)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitCellIcon allows checkout of a single cellicon (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutCellIcon(cellicon *models.CellIcon) {
	// check if the cellicon is staged
	if _, ok := backRepo.BackRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon]; ok {

		if id, ok := backRepo.BackRepoCellIcon.Map_CellIconPtr_CellIconDBID[cellicon]; ok {
			var celliconDB CellIconDB
			celliconDB.ID = id

			if _, err := backRepo.BackRepoCellIcon.db.First(&celliconDB, id); err != nil {
				log.Fatalln("CheckoutCellIcon : Problem with getting object with id:", id)
			}
			backRepo.BackRepoCellIcon.CheckoutPhaseOneInstance(&celliconDB)
			backRepo.BackRepoCellIcon.CheckoutPhaseTwoInstance(backRepo, &celliconDB)
		}
	}
}

// CopyBasicFieldsFromCellIcon
func (celliconDB *CellIconDB) CopyBasicFieldsFromCellIcon(cellicon *models.CellIcon) {
	// insertion point for fields commit

	celliconDB.Name_Data.String = cellicon.Name
	celliconDB.Name_Data.Valid = true

	celliconDB.Icon_Data.String = cellicon.Icon
	celliconDB.Icon_Data.Valid = true

	celliconDB.NeedsConfirmation_Data.Bool = cellicon.NeedsConfirmation
	celliconDB.NeedsConfirmation_Data.Valid = true

	celliconDB.ConfirmationMessage_Data.String = cellicon.ConfirmationMessage
	celliconDB.ConfirmationMessage_Data.Valid = true
}

// CopyBasicFieldsFromCellIcon_WOP
func (celliconDB *CellIconDB) CopyBasicFieldsFromCellIcon_WOP(cellicon *models.CellIcon_WOP) {
	// insertion point for fields commit

	celliconDB.Name_Data.String = cellicon.Name
	celliconDB.Name_Data.Valid = true

	celliconDB.Icon_Data.String = cellicon.Icon
	celliconDB.Icon_Data.Valid = true

	celliconDB.NeedsConfirmation_Data.Bool = cellicon.NeedsConfirmation
	celliconDB.NeedsConfirmation_Data.Valid = true

	celliconDB.ConfirmationMessage_Data.String = cellicon.ConfirmationMessage
	celliconDB.ConfirmationMessage_Data.Valid = true
}

// CopyBasicFieldsFromCellIconWOP
func (celliconDB *CellIconDB) CopyBasicFieldsFromCellIconWOP(cellicon *CellIconWOP) {
	// insertion point for fields commit

	celliconDB.Name_Data.String = cellicon.Name
	celliconDB.Name_Data.Valid = true

	celliconDB.Icon_Data.String = cellicon.Icon
	celliconDB.Icon_Data.Valid = true

	celliconDB.NeedsConfirmation_Data.Bool = cellicon.NeedsConfirmation
	celliconDB.NeedsConfirmation_Data.Valid = true

	celliconDB.ConfirmationMessage_Data.String = cellicon.ConfirmationMessage
	celliconDB.ConfirmationMessage_Data.Valid = true
}

// CopyBasicFieldsToCellIcon
func (celliconDB *CellIconDB) CopyBasicFieldsToCellIcon(cellicon *models.CellIcon) {
	// insertion point for checkout of basic fields (back repo to stage)
	cellicon.Name = celliconDB.Name_Data.String
	cellicon.Icon = celliconDB.Icon_Data.String
	cellicon.NeedsConfirmation = celliconDB.NeedsConfirmation_Data.Bool
	cellicon.ConfirmationMessage = celliconDB.ConfirmationMessage_Data.String
}

// CopyBasicFieldsToCellIcon_WOP
func (celliconDB *CellIconDB) CopyBasicFieldsToCellIcon_WOP(cellicon *models.CellIcon_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	cellicon.Name = celliconDB.Name_Data.String
	cellicon.Icon = celliconDB.Icon_Data.String
	cellicon.NeedsConfirmation = celliconDB.NeedsConfirmation_Data.Bool
	cellicon.ConfirmationMessage = celliconDB.ConfirmationMessage_Data.String
}

// CopyBasicFieldsToCellIconWOP
func (celliconDB *CellIconDB) CopyBasicFieldsToCellIconWOP(cellicon *CellIconWOP) {
	cellicon.ID = int(celliconDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	cellicon.Name = celliconDB.Name_Data.String
	cellicon.Icon = celliconDB.Icon_Data.String
	cellicon.NeedsConfirmation = celliconDB.NeedsConfirmation_Data.Bool
	cellicon.ConfirmationMessage = celliconDB.ConfirmationMessage_Data.String
}

// Backup generates a json file from a slice of all CellIconDB instances in the backrepo
func (backRepoCellIcon *BackRepoCellIconStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "CellIconDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CellIconDB, 0)
	for _, celliconDB := range backRepoCellIcon.Map_CellIconDBID_CellIconDB {
		forBackup = append(forBackup, celliconDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json CellIcon ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json CellIcon file", err.Error())
	}
}

// Backup generates a json file from a slice of all CellIconDB instances in the backrepo
func (backRepoCellIcon *BackRepoCellIconStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CellIconDB, 0)
	for _, celliconDB := range backRepoCellIcon.Map_CellIconDBID_CellIconDB {
		forBackup = append(forBackup, celliconDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("CellIcon")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&CellIcon_Fields, -1)
	for _, celliconDB := range forBackup {

		var celliconWOP CellIconWOP
		celliconDB.CopyBasicFieldsToCellIconWOP(&celliconWOP)

		row := sh.AddRow()
		row.WriteStruct(&celliconWOP, -1)
	}
}

// RestoreXL from the "CellIcon" sheet all CellIconDB instances
func (backRepoCellIcon *BackRepoCellIconStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoCellIconid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["CellIcon"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoCellIcon.rowVisitorCellIcon)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoCellIcon *BackRepoCellIconStruct) rowVisitorCellIcon(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var celliconWOP CellIconWOP
		row.ReadStruct(&celliconWOP)

		// add the unmarshalled struct to the stage
		celliconDB := new(CellIconDB)
		celliconDB.CopyBasicFieldsFromCellIconWOP(&celliconWOP)

		celliconDB_ID_atBackupTime := celliconDB.ID
		celliconDB.ID = 0
		_, err := backRepoCellIcon.db.Create(celliconDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoCellIcon.Map_CellIconDBID_CellIconDB[celliconDB.ID] = celliconDB
		BackRepoCellIconid_atBckpTime_newID[celliconDB_ID_atBackupTime] = celliconDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "CellIconDB.json" in dirPath that stores an array
// of CellIconDB and stores it in the database
// the map BackRepoCellIconid_atBckpTime_newID is updated accordingly
func (backRepoCellIcon *BackRepoCellIconStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoCellIconid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "CellIconDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json CellIcon file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*CellIconDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_CellIconDBID_CellIconDB
	for _, celliconDB := range forRestore {

		celliconDB_ID_atBackupTime := celliconDB.ID
		celliconDB.ID = 0
		_, err := backRepoCellIcon.db.Create(celliconDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoCellIcon.Map_CellIconDBID_CellIconDB[celliconDB.ID] = celliconDB
		BackRepoCellIconid_atBckpTime_newID[celliconDB_ID_atBackupTime] = celliconDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json CellIcon file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<CellIcon>id_atBckpTime_newID
// to compute new index
func (backRepoCellIcon *BackRepoCellIconStruct) RestorePhaseTwo() {

	for _, celliconDB := range backRepoCellIcon.Map_CellIconDBID_CellIconDB {

		// next line of code is to avert unused variable compilation error
		_ = celliconDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoCellIcon.db.Model(celliconDB)
		_, err := db.Updates(*celliconDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoCellIcon.ResetReversePointers commits all staged instances of CellIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellIcon *BackRepoCellIconStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, cellicon := range backRepoCellIcon.Map_CellIconDBID_CellIconPtr {
		backRepoCellIcon.ResetReversePointersInstance(backRepo, idx, cellicon)
	}

	return
}

func (backRepoCellIcon *BackRepoCellIconStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, cellicon *models.CellIcon) (Error error) {

	// fetch matching celliconDB
	if celliconDB, ok := backRepoCellIcon.Map_CellIconDBID_CellIconDB[idx]; ok {
		_ = celliconDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoCellIconid_atBckpTime_newID map[uint]uint
