// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoHorizontalAxis BackRepoHorizontalAxisStruct

	BackRepoLine BackRepoLineStruct

	BackRepoParameter BackRepoParameterStruct

	BackRepoRhombus BackRepoRhombusStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&HorizontalAxisDB{},
		&LineDB{},
		&ParameterDB{},
		&RhombusDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoHorizontalAxis = BackRepoHorizontalAxisStruct{
		Map_HorizontalAxisDBID_HorizontalAxisPtr: make(map[uint]*models.HorizontalAxis, 0),
		Map_HorizontalAxisDBID_HorizontalAxisDB:  make(map[uint]*HorizontalAxisDB, 0),
		Map_HorizontalAxisPtr_HorizontalAxisDBID: make(map[*models.HorizontalAxis]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLine = BackRepoLineStruct{
		Map_LineDBID_LinePtr: make(map[uint]*models.Line, 0),
		Map_LineDBID_LineDB:  make(map[uint]*LineDB, 0),
		Map_LinePtr_LineDBID: make(map[*models.Line]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoParameter = BackRepoParameterStruct{
		Map_ParameterDBID_ParameterPtr: make(map[uint]*models.Parameter, 0),
		Map_ParameterDBID_ParameterDB:  make(map[uint]*ParameterDB, 0),
		Map_ParameterPtr_ParameterDBID: make(map[*models.Parameter]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRhombus = BackRepoRhombusStruct{
		Map_RhombusDBID_RhombusPtr: make(map[uint]*models.Rhombus, 0),
		Map_RhombusDBID_RhombusDB:  make(map[uint]*RhombusDB, 0),
		Map_RhombusPtr_RhombusDBID: make(map[*models.Rhombus]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()
	
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoHorizontalAxis.CommitPhaseOne(stage)
	backRepo.BackRepoLine.CommitPhaseOne(stage)
	backRepo.BackRepoParameter.CommitPhaseOne(stage)
	backRepo.BackRepoRhombus.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoHorizontalAxis.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoParameter.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRhombus.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoHorizontalAxis.CheckoutPhaseOne()
	backRepo.BackRepoLine.CheckoutPhaseOne()
	backRepo.BackRepoParameter.CheckoutPhaseOne()
	backRepo.BackRepoRhombus.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoHorizontalAxis.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoParameter.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRhombus.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoHorizontalAxis.Backup(dirPath)
	backRepo.BackRepoLine.Backup(dirPath)
	backRepo.BackRepoParameter.Backup(dirPath)
	backRepo.BackRepoRhombus.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoHorizontalAxis.BackupXL(file)
	backRepo.BackRepoLine.BackupXL(file)
	backRepo.BackRepoParameter.BackupXL(file)
	backRepo.BackRepoRhombus.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoHorizontalAxis.RestorePhaseOne(dirPath)
	backRepo.BackRepoLine.RestorePhaseOne(dirPath)
	backRepo.BackRepoParameter.RestorePhaseOne(dirPath)
	backRepo.BackRepoRhombus.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoHorizontalAxis.RestorePhaseTwo()
	backRepo.BackRepoLine.RestorePhaseTwo()
	backRepo.BackRepoParameter.RestorePhaseTwo()
	backRepo.BackRepoRhombus.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoHorizontalAxis.RestoreXLPhaseOne(file)
	backRepo.BackRepoLine.RestoreXLPhaseOne(file)
	backRepo.BackRepoParameter.RestoreXLPhaseOne(file)
	backRepo.BackRepoRhombus.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	activeChannels := make([]chan int, 0)

	for _, ch := range backRepoStruct.subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			activeChannels = append(activeChannels, ch)
		default:
			// Assume channel is no longer active; don't add to activeChannels
			log.Println("Channel no longer active")
			close(ch)
		}
	}
	backRepoStruct.subscribers = activeChannels
}
