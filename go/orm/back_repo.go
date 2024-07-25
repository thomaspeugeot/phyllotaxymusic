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
	BackRepoAxis BackRepoAxisStruct

	BackRepoAxisGrid BackRepoAxisGridStruct

	BackRepoBezier BackRepoBezierStruct

	BackRepoCircle BackRepoCircleStruct

	BackRepoCircleGrid BackRepoCircleGridStruct

	BackRepoHorizontalAxis BackRepoHorizontalAxisStruct

	BackRepoParameter BackRepoParameterStruct

	BackRepoRhombus BackRepoRhombusStruct

	BackRepoRhombusGrid BackRepoRhombusGridStruct

	BackRepoVerticalAxis BackRepoVerticalAxisStruct

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
		&AxisDB{},
		&AxisGridDB{},
		&BezierDB{},
		&CircleDB{},
		&CircleGridDB{},
		&HorizontalAxisDB{},
		&ParameterDB{},
		&RhombusDB{},
		&RhombusGridDB{},
		&VerticalAxisDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAxis = BackRepoAxisStruct{
		Map_AxisDBID_AxisPtr: make(map[uint]*models.Axis, 0),
		Map_AxisDBID_AxisDB:  make(map[uint]*AxisDB, 0),
		Map_AxisPtr_AxisDBID: make(map[*models.Axis]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAxisGrid = BackRepoAxisGridStruct{
		Map_AxisGridDBID_AxisGridPtr: make(map[uint]*models.AxisGrid, 0),
		Map_AxisGridDBID_AxisGridDB:  make(map[uint]*AxisGridDB, 0),
		Map_AxisGridPtr_AxisGridDBID: make(map[*models.AxisGrid]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBezier = BackRepoBezierStruct{
		Map_BezierDBID_BezierPtr: make(map[uint]*models.Bezier, 0),
		Map_BezierDBID_BezierDB:  make(map[uint]*BezierDB, 0),
		Map_BezierPtr_BezierDBID: make(map[*models.Bezier]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCircle = BackRepoCircleStruct{
		Map_CircleDBID_CirclePtr: make(map[uint]*models.Circle, 0),
		Map_CircleDBID_CircleDB:  make(map[uint]*CircleDB, 0),
		Map_CirclePtr_CircleDBID: make(map[*models.Circle]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCircleGrid = BackRepoCircleGridStruct{
		Map_CircleGridDBID_CircleGridPtr: make(map[uint]*models.CircleGrid, 0),
		Map_CircleGridDBID_CircleGridDB:  make(map[uint]*CircleGridDB, 0),
		Map_CircleGridPtr_CircleGridDBID: make(map[*models.CircleGrid]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHorizontalAxis = BackRepoHorizontalAxisStruct{
		Map_HorizontalAxisDBID_HorizontalAxisPtr: make(map[uint]*models.HorizontalAxis, 0),
		Map_HorizontalAxisDBID_HorizontalAxisDB:  make(map[uint]*HorizontalAxisDB, 0),
		Map_HorizontalAxisPtr_HorizontalAxisDBID: make(map[*models.HorizontalAxis]uint, 0),

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
	backRepo.BackRepoRhombusGrid = BackRepoRhombusGridStruct{
		Map_RhombusGridDBID_RhombusGridPtr: make(map[uint]*models.RhombusGrid, 0),
		Map_RhombusGridDBID_RhombusGridDB:  make(map[uint]*RhombusGridDB, 0),
		Map_RhombusGridPtr_RhombusGridDBID: make(map[*models.RhombusGrid]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVerticalAxis = BackRepoVerticalAxisStruct{
		Map_VerticalAxisDBID_VerticalAxisPtr: make(map[uint]*models.VerticalAxis, 0),
		Map_VerticalAxisDBID_VerticalAxisDB:  make(map[uint]*VerticalAxisDB, 0),
		Map_VerticalAxisPtr_VerticalAxisDBID: make(map[*models.VerticalAxis]uint, 0),

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
	backRepo.BackRepoAxis.CommitPhaseOne(stage)
	backRepo.BackRepoAxisGrid.CommitPhaseOne(stage)
	backRepo.BackRepoBezier.CommitPhaseOne(stage)
	backRepo.BackRepoCircle.CommitPhaseOne(stage)
	backRepo.BackRepoCircleGrid.CommitPhaseOne(stage)
	backRepo.BackRepoHorizontalAxis.CommitPhaseOne(stage)
	backRepo.BackRepoParameter.CommitPhaseOne(stage)
	backRepo.BackRepoRhombus.CommitPhaseOne(stage)
	backRepo.BackRepoRhombusGrid.CommitPhaseOne(stage)
	backRepo.BackRepoVerticalAxis.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAxis.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAxisGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBezier.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCircle.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCircleGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHorizontalAxis.CommitPhaseTwo(backRepo)
	backRepo.BackRepoParameter.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRhombus.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRhombusGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVerticalAxis.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAxis.CheckoutPhaseOne()
	backRepo.BackRepoAxisGrid.CheckoutPhaseOne()
	backRepo.BackRepoBezier.CheckoutPhaseOne()
	backRepo.BackRepoCircle.CheckoutPhaseOne()
	backRepo.BackRepoCircleGrid.CheckoutPhaseOne()
	backRepo.BackRepoHorizontalAxis.CheckoutPhaseOne()
	backRepo.BackRepoParameter.CheckoutPhaseOne()
	backRepo.BackRepoRhombus.CheckoutPhaseOne()
	backRepo.BackRepoRhombusGrid.CheckoutPhaseOne()
	backRepo.BackRepoVerticalAxis.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAxis.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAxisGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBezier.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCircle.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCircleGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHorizontalAxis.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoParameter.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRhombus.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRhombusGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVerticalAxis.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAxis.Backup(dirPath)
	backRepo.BackRepoAxisGrid.Backup(dirPath)
	backRepo.BackRepoBezier.Backup(dirPath)
	backRepo.BackRepoCircle.Backup(dirPath)
	backRepo.BackRepoCircleGrid.Backup(dirPath)
	backRepo.BackRepoHorizontalAxis.Backup(dirPath)
	backRepo.BackRepoParameter.Backup(dirPath)
	backRepo.BackRepoRhombus.Backup(dirPath)
	backRepo.BackRepoRhombusGrid.Backup(dirPath)
	backRepo.BackRepoVerticalAxis.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAxis.BackupXL(file)
	backRepo.BackRepoAxisGrid.BackupXL(file)
	backRepo.BackRepoBezier.BackupXL(file)
	backRepo.BackRepoCircle.BackupXL(file)
	backRepo.BackRepoCircleGrid.BackupXL(file)
	backRepo.BackRepoHorizontalAxis.BackupXL(file)
	backRepo.BackRepoParameter.BackupXL(file)
	backRepo.BackRepoRhombus.BackupXL(file)
	backRepo.BackRepoRhombusGrid.BackupXL(file)
	backRepo.BackRepoVerticalAxis.BackupXL(file)

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
	backRepo.BackRepoAxis.RestorePhaseOne(dirPath)
	backRepo.BackRepoAxisGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoBezier.RestorePhaseOne(dirPath)
	backRepo.BackRepoCircle.RestorePhaseOne(dirPath)
	backRepo.BackRepoCircleGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoHorizontalAxis.RestorePhaseOne(dirPath)
	backRepo.BackRepoParameter.RestorePhaseOne(dirPath)
	backRepo.BackRepoRhombus.RestorePhaseOne(dirPath)
	backRepo.BackRepoRhombusGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoVerticalAxis.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAxis.RestorePhaseTwo()
	backRepo.BackRepoAxisGrid.RestorePhaseTwo()
	backRepo.BackRepoBezier.RestorePhaseTwo()
	backRepo.BackRepoCircle.RestorePhaseTwo()
	backRepo.BackRepoCircleGrid.RestorePhaseTwo()
	backRepo.BackRepoHorizontalAxis.RestorePhaseTwo()
	backRepo.BackRepoParameter.RestorePhaseTwo()
	backRepo.BackRepoRhombus.RestorePhaseTwo()
	backRepo.BackRepoRhombusGrid.RestorePhaseTwo()
	backRepo.BackRepoVerticalAxis.RestorePhaseTwo()

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
	backRepo.BackRepoAxis.RestoreXLPhaseOne(file)
	backRepo.BackRepoAxisGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoBezier.RestoreXLPhaseOne(file)
	backRepo.BackRepoCircle.RestoreXLPhaseOne(file)
	backRepo.BackRepoCircleGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoHorizontalAxis.RestoreXLPhaseOne(file)
	backRepo.BackRepoParameter.RestoreXLPhaseOne(file)
	backRepo.BackRepoRhombus.RestoreXLPhaseOne(file)
	backRepo.BackRepoRhombusGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoVerticalAxis.RestoreXLPhaseOne(file)

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
