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

	"github.com/fullstack-lang/gong/lib/svg/go/db"
	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Line_sql sql.NullBool
var dummy_Line_time time.Duration
var dummy_Line_sort sort.Float64Slice

// LineAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model lineAPI
type LineAPI struct {
	gorm.Model

	models.Line_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	LinePointersEncoding LinePointersEncoding
}

// LinePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LinePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Animates is a slice of pointers to another Struct (optional or 0..1)
	Animates IntSlice `gorm:"type:TEXT"`
}

// LineDB describes a line in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model lineDB
type LineDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field lineDB.Name
	Name_Data sql.NullString

	// Declation for basic field lineDB.X1
	X1_Data sql.NullFloat64

	// Declation for basic field lineDB.Y1
	Y1_Data sql.NullFloat64

	// Declation for basic field lineDB.X2
	X2_Data sql.NullFloat64

	// Declation for basic field lineDB.Y2
	Y2_Data sql.NullFloat64

	// Declation for basic field lineDB.Color
	Color_Data sql.NullString

	// Declation for basic field lineDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field lineDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field lineDB.StrokeOpacity
	StrokeOpacity_Data sql.NullFloat64

	// Declation for basic field lineDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field lineDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field lineDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field lineDB.Transform
	Transform_Data sql.NullString

	// Declation for basic field lineDB.MouseClickX
	MouseClickX_Data sql.NullFloat64

	// Declation for basic field lineDB.MouseClickY
	MouseClickY_Data sql.NullFloat64

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	LinePointersEncoding
}

// LineDBs arrays lineDBs
// swagger:response lineDBsResponse
type LineDBs []LineDB

// LineDBResponse provides response
// swagger:response lineDBResponse
type LineDBResponse struct {
	LineDB
}

// LineWOP is a Line without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LineWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	X1 float64 `xlsx:"2"`

	Y1 float64 `xlsx:"3"`

	X2 float64 `xlsx:"4"`

	Y2 float64 `xlsx:"5"`

	Color string `xlsx:"6"`

	FillOpacity float64 `xlsx:"7"`

	Stroke string `xlsx:"8"`

	StrokeOpacity float64 `xlsx:"9"`

	StrokeWidth float64 `xlsx:"10"`

	StrokeDashArray string `xlsx:"11"`

	StrokeDashArrayWhenSelected string `xlsx:"12"`

	Transform string `xlsx:"13"`

	MouseClickX float64 `xlsx:"14"`

	MouseClickY float64 `xlsx:"15"`
	// insertion for WOP pointer fields
}

var Line_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"X1",
	"Y1",
	"X2",
	"Y2",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeOpacity",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
	"MouseClickX",
	"MouseClickY",
}

type BackRepoLineStruct struct {
	// stores LineDB according to their gorm ID
	Map_LineDBID_LineDB map[uint]*LineDB

	// stores LineDB ID according to Line address
	Map_LinePtr_LineDBID map[*models.Line]uint

	// stores Line according to their gorm ID
	Map_LineDBID_LinePtr map[uint]*models.Line

	db db.DBInterface

	stage *models.Stage
}

func (backRepoLine *BackRepoLineStruct) GetStage() (stage *models.Stage) {
	stage = backRepoLine.stage
	return
}

func (backRepoLine *BackRepoLineStruct) GetDB() db.DBInterface {
	return backRepoLine.db
}

// GetLineDBFromLinePtr is a handy function to access the back repo instance from the stage instance
func (backRepoLine *BackRepoLineStruct) GetLineDBFromLinePtr(line *models.Line) (lineDB *LineDB) {
	id := backRepoLine.Map_LinePtr_LineDBID[line]
	lineDB = backRepoLine.Map_LineDBID_LineDB[id]
	return
}

// BackRepoLine.CommitPhaseOne commits all staged instances of Line to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLine *BackRepoLineStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var lines []*models.Line
	for line := range stage.Lines {
		lines = append(lines, line)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(lines, func(i, j int) bool {
		return stage.LineMap_Staged_Order[lines[i]] < stage.LineMap_Staged_Order[lines[j]]
	})

	for _, line := range lines {
		backRepoLine.CommitPhaseOneInstance(line)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, line := range backRepoLine.Map_LineDBID_LinePtr {
		if _, ok := stage.Lines[line]; !ok {
			backRepoLine.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLine.CommitDeleteInstance commits deletion of Line to the BackRepo
func (backRepoLine *BackRepoLineStruct) CommitDeleteInstance(id uint) (Error error) {

	line := backRepoLine.Map_LineDBID_LinePtr[id]

	// line is not staged anymore, remove lineDB
	lineDB := backRepoLine.Map_LineDBID_LineDB[id]
	db, _ := backRepoLine.db.Unscoped()
	_, err := db.Delete(lineDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoLine.Map_LinePtr_LineDBID, line)
	delete(backRepoLine.Map_LineDBID_LinePtr, id)
	delete(backRepoLine.Map_LineDBID_LineDB, id)

	return
}

// BackRepoLine.CommitPhaseOneInstance commits line staged instances of Line to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLine *BackRepoLineStruct) CommitPhaseOneInstance(line *models.Line) (Error error) {

	// check if the line is not commited yet
	if _, ok := backRepoLine.Map_LinePtr_LineDBID[line]; ok {
		return
	}

	// initiate line
	var lineDB LineDB
	lineDB.CopyBasicFieldsFromLine(line)

	_, err := backRepoLine.db.Create(&lineDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoLine.Map_LinePtr_LineDBID[line] = lineDB.ID
	backRepoLine.Map_LineDBID_LinePtr[lineDB.ID] = line
	backRepoLine.Map_LineDBID_LineDB[lineDB.ID] = &lineDB

	return
}

// BackRepoLine.CommitPhaseTwo commits all staged instances of Line to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLine *BackRepoLineStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, line := range backRepoLine.Map_LineDBID_LinePtr {
		backRepoLine.CommitPhaseTwoInstance(backRepo, idx, line)
	}

	return
}

// BackRepoLine.CommitPhaseTwoInstance commits {{structname }} of models.Line to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLine *BackRepoLineStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, line *models.Line) (Error error) {

	// fetch matching lineDB
	if lineDB, ok := backRepoLine.Map_LineDBID_LineDB[idx]; ok {

		lineDB.CopyBasicFieldsFromLine(line)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		lineDB.LinePointersEncoding.Animates = make([]int, 0)
		// 2. encode
		for _, animateAssocEnd := range line.Animates {
			animateAssocEnd_DB :=
				backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateAssocEnd)
			
			// the stage might be inconsistant, meaning that the animateAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if animateAssocEnd_DB == nil {
				continue
			}
			
			lineDB.LinePointersEncoding.Animates =
				append(lineDB.LinePointersEncoding.Animates, int(animateAssocEnd_DB.ID))
		}

		_, err := backRepoLine.db.Save(lineDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Line intance %s", line.Name))
		return err
	}

	return
}

// BackRepoLine.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoLine *BackRepoLineStruct) CheckoutPhaseOne() (Error error) {

	lineDBArray := make([]LineDB, 0)
	_, err := backRepoLine.db.Find(&lineDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	lineInstancesToBeRemovedFromTheStage := make(map[*models.Line]any)
	for key, value := range backRepoLine.stage.Lines {
		lineInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, lineDB := range lineDBArray {
		backRepoLine.CheckoutPhaseOneInstance(&lineDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		line, ok := backRepoLine.Map_LineDBID_LinePtr[lineDB.ID]
		if ok {
			delete(lineInstancesToBeRemovedFromTheStage, line)
		}
	}

	// remove from stage and back repo's 3 maps all lines that are not in the checkout
	for line := range lineInstancesToBeRemovedFromTheStage {
		line.Unstage(backRepoLine.GetStage())

		// remove instance from the back repo 3 maps
		lineID := backRepoLine.Map_LinePtr_LineDBID[line]
		delete(backRepoLine.Map_LinePtr_LineDBID, line)
		delete(backRepoLine.Map_LineDBID_LineDB, lineID)
		delete(backRepoLine.Map_LineDBID_LinePtr, lineID)
	}

	return
}

// CheckoutPhaseOneInstance takes a lineDB that has been found in the DB, updates the backRepo and stages the
// models version of the lineDB
func (backRepoLine *BackRepoLineStruct) CheckoutPhaseOneInstance(lineDB *LineDB) (Error error) {

	line, ok := backRepoLine.Map_LineDBID_LinePtr[lineDB.ID]
	if !ok {
		line = new(models.Line)

		backRepoLine.Map_LineDBID_LinePtr[lineDB.ID] = line
		backRepoLine.Map_LinePtr_LineDBID[line] = lineDB.ID

		// append model store with the new element
		line.Name = lineDB.Name_Data.String
		line.Stage(backRepoLine.GetStage())
	}
	lineDB.CopyBasicFieldsToLine(line)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	line.Stage(backRepoLine.GetStage())

	// preserve pointer to lineDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LineDBID_LineDB)[lineDB hold variable pointers
	lineDB_Data := *lineDB
	preservedPtrToLine := &lineDB_Data
	backRepoLine.Map_LineDBID_LineDB[lineDB.ID] = preservedPtrToLine

	return
}

// BackRepoLine.CheckoutPhaseTwo Checkouts all staged instances of Line to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLine *BackRepoLineStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, lineDB := range backRepoLine.Map_LineDBID_LineDB {
		backRepoLine.CheckoutPhaseTwoInstance(backRepo, lineDB)
	}
	return
}

// BackRepoLine.CheckoutPhaseTwoInstance Checkouts staged instances of Line to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLine *BackRepoLineStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, lineDB *LineDB) (Error error) {

	line := backRepoLine.Map_LineDBID_LinePtr[lineDB.ID]

	lineDB.DecodePointers(backRepo, line)

	return
}

func (lineDB *LineDB) DecodePointers(backRepo *BackRepoStruct, line *models.Line) {

	// insertion point for checkout of pointer encoding
	// This loop redeem line.Animates in the stage from the encode in the back repo
	// It parses all AnimateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	line.Animates = line.Animates[:0]
	for _, _Animateid := range lineDB.LinePointersEncoding.Animates {
		line.Animates = append(line.Animates, backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[uint(_Animateid)])
	}

	return
}

// CommitLine allows commit of a single line (if already staged)
func (backRepo *BackRepoStruct) CommitLine(line *models.Line) {
	backRepo.BackRepoLine.CommitPhaseOneInstance(line)
	if id, ok := backRepo.BackRepoLine.Map_LinePtr_LineDBID[line]; ok {
		backRepo.BackRepoLine.CommitPhaseTwoInstance(backRepo, id, line)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitLine allows checkout of a single line (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLine(line *models.Line) {
	// check if the line is staged
	if _, ok := backRepo.BackRepoLine.Map_LinePtr_LineDBID[line]; ok {

		if id, ok := backRepo.BackRepoLine.Map_LinePtr_LineDBID[line]; ok {
			var lineDB LineDB
			lineDB.ID = id

			if _, err := backRepo.BackRepoLine.db.First(&lineDB, id); err != nil {
				log.Fatalln("CheckoutLine : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLine.CheckoutPhaseOneInstance(&lineDB)
			backRepo.BackRepoLine.CheckoutPhaseTwoInstance(backRepo, &lineDB)
		}
	}
}

// CopyBasicFieldsFromLine
func (lineDB *LineDB) CopyBasicFieldsFromLine(line *models.Line) {
	// insertion point for fields commit

	lineDB.Name_Data.String = line.Name
	lineDB.Name_Data.Valid = true

	lineDB.X1_Data.Float64 = line.X1
	lineDB.X1_Data.Valid = true

	lineDB.Y1_Data.Float64 = line.Y1
	lineDB.Y1_Data.Valid = true

	lineDB.X2_Data.Float64 = line.X2
	lineDB.X2_Data.Valid = true

	lineDB.Y2_Data.Float64 = line.Y2
	lineDB.Y2_Data.Valid = true

	lineDB.Color_Data.String = line.Color
	lineDB.Color_Data.Valid = true

	lineDB.FillOpacity_Data.Float64 = line.FillOpacity
	lineDB.FillOpacity_Data.Valid = true

	lineDB.Stroke_Data.String = line.Stroke
	lineDB.Stroke_Data.Valid = true

	lineDB.StrokeOpacity_Data.Float64 = line.StrokeOpacity
	lineDB.StrokeOpacity_Data.Valid = true

	lineDB.StrokeWidth_Data.Float64 = line.StrokeWidth
	lineDB.StrokeWidth_Data.Valid = true

	lineDB.StrokeDashArray_Data.String = line.StrokeDashArray
	lineDB.StrokeDashArray_Data.Valid = true

	lineDB.StrokeDashArrayWhenSelected_Data.String = line.StrokeDashArrayWhenSelected
	lineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	lineDB.Transform_Data.String = line.Transform
	lineDB.Transform_Data.Valid = true

	lineDB.MouseClickX_Data.Float64 = line.MouseClickX
	lineDB.MouseClickX_Data.Valid = true

	lineDB.MouseClickY_Data.Float64 = line.MouseClickY
	lineDB.MouseClickY_Data.Valid = true
}

// CopyBasicFieldsFromLine_WOP
func (lineDB *LineDB) CopyBasicFieldsFromLine_WOP(line *models.Line_WOP) {
	// insertion point for fields commit

	lineDB.Name_Data.String = line.Name
	lineDB.Name_Data.Valid = true

	lineDB.X1_Data.Float64 = line.X1
	lineDB.X1_Data.Valid = true

	lineDB.Y1_Data.Float64 = line.Y1
	lineDB.Y1_Data.Valid = true

	lineDB.X2_Data.Float64 = line.X2
	lineDB.X2_Data.Valid = true

	lineDB.Y2_Data.Float64 = line.Y2
	lineDB.Y2_Data.Valid = true

	lineDB.Color_Data.String = line.Color
	lineDB.Color_Data.Valid = true

	lineDB.FillOpacity_Data.Float64 = line.FillOpacity
	lineDB.FillOpacity_Data.Valid = true

	lineDB.Stroke_Data.String = line.Stroke
	lineDB.Stroke_Data.Valid = true

	lineDB.StrokeOpacity_Data.Float64 = line.StrokeOpacity
	lineDB.StrokeOpacity_Data.Valid = true

	lineDB.StrokeWidth_Data.Float64 = line.StrokeWidth
	lineDB.StrokeWidth_Data.Valid = true

	lineDB.StrokeDashArray_Data.String = line.StrokeDashArray
	lineDB.StrokeDashArray_Data.Valid = true

	lineDB.StrokeDashArrayWhenSelected_Data.String = line.StrokeDashArrayWhenSelected
	lineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	lineDB.Transform_Data.String = line.Transform
	lineDB.Transform_Data.Valid = true

	lineDB.MouseClickX_Data.Float64 = line.MouseClickX
	lineDB.MouseClickX_Data.Valid = true

	lineDB.MouseClickY_Data.Float64 = line.MouseClickY
	lineDB.MouseClickY_Data.Valid = true
}

// CopyBasicFieldsFromLineWOP
func (lineDB *LineDB) CopyBasicFieldsFromLineWOP(line *LineWOP) {
	// insertion point for fields commit

	lineDB.Name_Data.String = line.Name
	lineDB.Name_Data.Valid = true

	lineDB.X1_Data.Float64 = line.X1
	lineDB.X1_Data.Valid = true

	lineDB.Y1_Data.Float64 = line.Y1
	lineDB.Y1_Data.Valid = true

	lineDB.X2_Data.Float64 = line.X2
	lineDB.X2_Data.Valid = true

	lineDB.Y2_Data.Float64 = line.Y2
	lineDB.Y2_Data.Valid = true

	lineDB.Color_Data.String = line.Color
	lineDB.Color_Data.Valid = true

	lineDB.FillOpacity_Data.Float64 = line.FillOpacity
	lineDB.FillOpacity_Data.Valid = true

	lineDB.Stroke_Data.String = line.Stroke
	lineDB.Stroke_Data.Valid = true

	lineDB.StrokeOpacity_Data.Float64 = line.StrokeOpacity
	lineDB.StrokeOpacity_Data.Valid = true

	lineDB.StrokeWidth_Data.Float64 = line.StrokeWidth
	lineDB.StrokeWidth_Data.Valid = true

	lineDB.StrokeDashArray_Data.String = line.StrokeDashArray
	lineDB.StrokeDashArray_Data.Valid = true

	lineDB.StrokeDashArrayWhenSelected_Data.String = line.StrokeDashArrayWhenSelected
	lineDB.StrokeDashArrayWhenSelected_Data.Valid = true

	lineDB.Transform_Data.String = line.Transform
	lineDB.Transform_Data.Valid = true

	lineDB.MouseClickX_Data.Float64 = line.MouseClickX
	lineDB.MouseClickX_Data.Valid = true

	lineDB.MouseClickY_Data.Float64 = line.MouseClickY
	lineDB.MouseClickY_Data.Valid = true
}

// CopyBasicFieldsToLine
func (lineDB *LineDB) CopyBasicFieldsToLine(line *models.Line) {
	// insertion point for checkout of basic fields (back repo to stage)
	line.Name = lineDB.Name_Data.String
	line.X1 = lineDB.X1_Data.Float64
	line.Y1 = lineDB.Y1_Data.Float64
	line.X2 = lineDB.X2_Data.Float64
	line.Y2 = lineDB.Y2_Data.Float64
	line.Color = lineDB.Color_Data.String
	line.FillOpacity = lineDB.FillOpacity_Data.Float64
	line.Stroke = lineDB.Stroke_Data.String
	line.StrokeOpacity = lineDB.StrokeOpacity_Data.Float64
	line.StrokeWidth = lineDB.StrokeWidth_Data.Float64
	line.StrokeDashArray = lineDB.StrokeDashArray_Data.String
	line.StrokeDashArrayWhenSelected = lineDB.StrokeDashArrayWhenSelected_Data.String
	line.Transform = lineDB.Transform_Data.String
	line.MouseClickX = lineDB.MouseClickX_Data.Float64
	line.MouseClickY = lineDB.MouseClickY_Data.Float64
}

// CopyBasicFieldsToLine_WOP
func (lineDB *LineDB) CopyBasicFieldsToLine_WOP(line *models.Line_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	line.Name = lineDB.Name_Data.String
	line.X1 = lineDB.X1_Data.Float64
	line.Y1 = lineDB.Y1_Data.Float64
	line.X2 = lineDB.X2_Data.Float64
	line.Y2 = lineDB.Y2_Data.Float64
	line.Color = lineDB.Color_Data.String
	line.FillOpacity = lineDB.FillOpacity_Data.Float64
	line.Stroke = lineDB.Stroke_Data.String
	line.StrokeOpacity = lineDB.StrokeOpacity_Data.Float64
	line.StrokeWidth = lineDB.StrokeWidth_Data.Float64
	line.StrokeDashArray = lineDB.StrokeDashArray_Data.String
	line.StrokeDashArrayWhenSelected = lineDB.StrokeDashArrayWhenSelected_Data.String
	line.Transform = lineDB.Transform_Data.String
	line.MouseClickX = lineDB.MouseClickX_Data.Float64
	line.MouseClickY = lineDB.MouseClickY_Data.Float64
}

// CopyBasicFieldsToLineWOP
func (lineDB *LineDB) CopyBasicFieldsToLineWOP(line *LineWOP) {
	line.ID = int(lineDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	line.Name = lineDB.Name_Data.String
	line.X1 = lineDB.X1_Data.Float64
	line.Y1 = lineDB.Y1_Data.Float64
	line.X2 = lineDB.X2_Data.Float64
	line.Y2 = lineDB.Y2_Data.Float64
	line.Color = lineDB.Color_Data.String
	line.FillOpacity = lineDB.FillOpacity_Data.Float64
	line.Stroke = lineDB.Stroke_Data.String
	line.StrokeOpacity = lineDB.StrokeOpacity_Data.Float64
	line.StrokeWidth = lineDB.StrokeWidth_Data.Float64
	line.StrokeDashArray = lineDB.StrokeDashArray_Data.String
	line.StrokeDashArrayWhenSelected = lineDB.StrokeDashArrayWhenSelected_Data.String
	line.Transform = lineDB.Transform_Data.String
	line.MouseClickX = lineDB.MouseClickX_Data.Float64
	line.MouseClickY = lineDB.MouseClickY_Data.Float64
}

// Backup generates a json file from a slice of all LineDB instances in the backrepo
func (backRepoLine *BackRepoLineStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LineDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LineDB, 0)
	for _, lineDB := range backRepoLine.Map_LineDBID_LineDB {
		forBackup = append(forBackup, lineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Line ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Line file", err.Error())
	}
}

// Backup generates a json file from a slice of all LineDB instances in the backrepo
func (backRepoLine *BackRepoLineStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LineDB, 0)
	for _, lineDB := range backRepoLine.Map_LineDBID_LineDB {
		forBackup = append(forBackup, lineDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Line")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Line_Fields, -1)
	for _, lineDB := range forBackup {

		var lineWOP LineWOP
		lineDB.CopyBasicFieldsToLineWOP(&lineWOP)

		row := sh.AddRow()
		row.WriteStruct(&lineWOP, -1)
	}
}

// RestoreXL from the "Line" sheet all LineDB instances
func (backRepoLine *BackRepoLineStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLineid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Line"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLine.rowVisitorLine)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoLine *BackRepoLineStruct) rowVisitorLine(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var lineWOP LineWOP
		row.ReadStruct(&lineWOP)

		// add the unmarshalled struct to the stage
		lineDB := new(LineDB)
		lineDB.CopyBasicFieldsFromLineWOP(&lineWOP)

		lineDB_ID_atBackupTime := lineDB.ID
		lineDB.ID = 0
		_, err := backRepoLine.db.Create(lineDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoLine.Map_LineDBID_LineDB[lineDB.ID] = lineDB
		BackRepoLineid_atBckpTime_newID[lineDB_ID_atBackupTime] = lineDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LineDB.json" in dirPath that stores an array
// of LineDB and stores it in the database
// the map BackRepoLineid_atBckpTime_newID is updated accordingly
func (backRepoLine *BackRepoLineStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLineid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LineDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Line file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LineDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LineDBID_LineDB
	for _, lineDB := range forRestore {

		lineDB_ID_atBackupTime := lineDB.ID
		lineDB.ID = 0
		_, err := backRepoLine.db.Create(lineDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoLine.Map_LineDBID_LineDB[lineDB.ID] = lineDB
		BackRepoLineid_atBckpTime_newID[lineDB_ID_atBackupTime] = lineDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Line file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Line>id_atBckpTime_newID
// to compute new index
func (backRepoLine *BackRepoLineStruct) RestorePhaseTwo() {

	for _, lineDB := range backRepoLine.Map_LineDBID_LineDB {

		// next line of code is to avert unused variable compilation error
		_ = lineDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoLine.db.Model(lineDB)
		_, err := db.Updates(*lineDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoLine.ResetReversePointers commits all staged instances of Line to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLine *BackRepoLineStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, line := range backRepoLine.Map_LineDBID_LinePtr {
		backRepoLine.ResetReversePointersInstance(backRepo, idx, line)
	}

	return
}

func (backRepoLine *BackRepoLineStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, line *models.Line) (Error error) {

	// fetch matching lineDB
	if lineDB, ok := backRepoLine.Map_LineDBID_LineDB[idx]; ok {
		_ = lineDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLineid_atBckpTime_newID map[uint]uint
