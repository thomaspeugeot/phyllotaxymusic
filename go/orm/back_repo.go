// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/thomaspeugeot/phylotaxymusic/go/db"
	"github.com/thomaspeugeot/phylotaxymusic/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/thomaspeugeot/phylotaxymusic/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAxis BackRepoAxisStruct

	BackRepoAxisGrid BackRepoAxisGridStruct

	BackRepoBezier BackRepoBezierStruct

	BackRepoBezierGrid BackRepoBezierGridStruct

	BackRepoBezierGridStack BackRepoBezierGridStackStruct

	BackRepoCircle BackRepoCircleStruct

	BackRepoCircleGrid BackRepoCircleGridStruct

	BackRepoFrontCurve BackRepoFrontCurveStruct

	BackRepoFrontCurveStack BackRepoFrontCurveStackStruct

	BackRepoHorizontalAxis BackRepoHorizontalAxisStruct

	BackRepoKey BackRepoKeyStruct

	BackRepoMovingLine BackRepoMovingLineStruct

	BackRepoNoteInfo BackRepoNoteInfoStruct

	BackRepoParameter BackRepoParameterStruct

	BackRepoRhombus BackRepoRhombusStruct

	BackRepoRhombusGrid BackRepoRhombusGridStruct

	BackRepoShapeCategory BackRepoShapeCategoryStruct

	BackRepoSpiralBezier BackRepoSpiralBezierStruct

	BackRepoSpiralBezierGrid BackRepoSpiralBezierGridStruct

	BackRepoSpiralCircle BackRepoSpiralCircleStruct

	BackRepoSpiralCircleGrid BackRepoSpiralCircleGridStruct

	BackRepoSpiralLine BackRepoSpiralLineStruct

	BackRepoSpiralLineGrid BackRepoSpiralLineGridStruct

	BackRepoSpiralOrigin BackRepoSpiralOriginStruct

	BackRepoSpiralRhombus BackRepoSpiralRhombusStruct

	BackRepoSpiralRhombusGrid BackRepoSpiralRhombusGridStruct

	BackRepoVerticalAxis BackRepoVerticalAxisStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_thomaspeugeot_phylotaxymusic_go",
		&AxisDB{},
		&AxisGridDB{},
		&BezierDB{},
		&BezierGridDB{},
		&BezierGridStackDB{},
		&CircleDB{},
		&CircleGridDB{},
		&FrontCurveDB{},
		&FrontCurveStackDB{},
		&HorizontalAxisDB{},
		&KeyDB{},
		&MovingLineDB{},
		&NoteInfoDB{},
		&ParameterDB{},
		&RhombusDB{},
		&RhombusGridDB{},
		&ShapeCategoryDB{},
		&SpiralBezierDB{},
		&SpiralBezierGridDB{},
		&SpiralCircleDB{},
		&SpiralCircleGridDB{},
		&SpiralLineDB{},
		&SpiralLineGridDB{},
		&SpiralOriginDB{},
		&SpiralRhombusDB{},
		&SpiralRhombusGridDB{},
		&VerticalAxisDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

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
	backRepo.BackRepoBezierGrid = BackRepoBezierGridStruct{
		Map_BezierGridDBID_BezierGridPtr: make(map[uint]*models.BezierGrid, 0),
		Map_BezierGridDBID_BezierGridDB:  make(map[uint]*BezierGridDB, 0),
		Map_BezierGridPtr_BezierGridDBID: make(map[*models.BezierGrid]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBezierGridStack = BackRepoBezierGridStackStruct{
		Map_BezierGridStackDBID_BezierGridStackPtr: make(map[uint]*models.BezierGridStack, 0),
		Map_BezierGridStackDBID_BezierGridStackDB:  make(map[uint]*BezierGridStackDB, 0),
		Map_BezierGridStackPtr_BezierGridStackDBID: make(map[*models.BezierGridStack]uint, 0),

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
	backRepo.BackRepoFrontCurve = BackRepoFrontCurveStruct{
		Map_FrontCurveDBID_FrontCurvePtr: make(map[uint]*models.FrontCurve, 0),
		Map_FrontCurveDBID_FrontCurveDB:  make(map[uint]*FrontCurveDB, 0),
		Map_FrontCurvePtr_FrontCurveDBID: make(map[*models.FrontCurve]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFrontCurveStack = BackRepoFrontCurveStackStruct{
		Map_FrontCurveStackDBID_FrontCurveStackPtr: make(map[uint]*models.FrontCurveStack, 0),
		Map_FrontCurveStackDBID_FrontCurveStackDB:  make(map[uint]*FrontCurveStackDB, 0),
		Map_FrontCurveStackPtr_FrontCurveStackDBID: make(map[*models.FrontCurveStack]uint, 0),

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
	backRepo.BackRepoKey = BackRepoKeyStruct{
		Map_KeyDBID_KeyPtr: make(map[uint]*models.Key, 0),
		Map_KeyDBID_KeyDB:  make(map[uint]*KeyDB, 0),
		Map_KeyPtr_KeyDBID: make(map[*models.Key]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMovingLine = BackRepoMovingLineStruct{
		Map_MovingLineDBID_MovingLinePtr: make(map[uint]*models.MovingLine, 0),
		Map_MovingLineDBID_MovingLineDB:  make(map[uint]*MovingLineDB, 0),
		Map_MovingLinePtr_MovingLineDBID: make(map[*models.MovingLine]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNoteInfo = BackRepoNoteInfoStruct{
		Map_NoteInfoDBID_NoteInfoPtr: make(map[uint]*models.NoteInfo, 0),
		Map_NoteInfoDBID_NoteInfoDB:  make(map[uint]*NoteInfoDB, 0),
		Map_NoteInfoPtr_NoteInfoDBID: make(map[*models.NoteInfo]uint, 0),

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
	backRepo.BackRepoShapeCategory = BackRepoShapeCategoryStruct{
		Map_ShapeCategoryDBID_ShapeCategoryPtr: make(map[uint]*models.ShapeCategory, 0),
		Map_ShapeCategoryDBID_ShapeCategoryDB:  make(map[uint]*ShapeCategoryDB, 0),
		Map_ShapeCategoryPtr_ShapeCategoryDBID: make(map[*models.ShapeCategory]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralBezier = BackRepoSpiralBezierStruct{
		Map_SpiralBezierDBID_SpiralBezierPtr: make(map[uint]*models.SpiralBezier, 0),
		Map_SpiralBezierDBID_SpiralBezierDB:  make(map[uint]*SpiralBezierDB, 0),
		Map_SpiralBezierPtr_SpiralBezierDBID: make(map[*models.SpiralBezier]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralBezierGrid = BackRepoSpiralBezierGridStruct{
		Map_SpiralBezierGridDBID_SpiralBezierGridPtr: make(map[uint]*models.SpiralBezierGrid, 0),
		Map_SpiralBezierGridDBID_SpiralBezierGridDB:  make(map[uint]*SpiralBezierGridDB, 0),
		Map_SpiralBezierGridPtr_SpiralBezierGridDBID: make(map[*models.SpiralBezierGrid]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralCircle = BackRepoSpiralCircleStruct{
		Map_SpiralCircleDBID_SpiralCirclePtr: make(map[uint]*models.SpiralCircle, 0),
		Map_SpiralCircleDBID_SpiralCircleDB:  make(map[uint]*SpiralCircleDB, 0),
		Map_SpiralCirclePtr_SpiralCircleDBID: make(map[*models.SpiralCircle]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralCircleGrid = BackRepoSpiralCircleGridStruct{
		Map_SpiralCircleGridDBID_SpiralCircleGridPtr: make(map[uint]*models.SpiralCircleGrid, 0),
		Map_SpiralCircleGridDBID_SpiralCircleGridDB:  make(map[uint]*SpiralCircleGridDB, 0),
		Map_SpiralCircleGridPtr_SpiralCircleGridDBID: make(map[*models.SpiralCircleGrid]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralLine = BackRepoSpiralLineStruct{
		Map_SpiralLineDBID_SpiralLinePtr: make(map[uint]*models.SpiralLine, 0),
		Map_SpiralLineDBID_SpiralLineDB:  make(map[uint]*SpiralLineDB, 0),
		Map_SpiralLinePtr_SpiralLineDBID: make(map[*models.SpiralLine]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralLineGrid = BackRepoSpiralLineGridStruct{
		Map_SpiralLineGridDBID_SpiralLineGridPtr: make(map[uint]*models.SpiralLineGrid, 0),
		Map_SpiralLineGridDBID_SpiralLineGridDB:  make(map[uint]*SpiralLineGridDB, 0),
		Map_SpiralLineGridPtr_SpiralLineGridDBID: make(map[*models.SpiralLineGrid]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralOrigin = BackRepoSpiralOriginStruct{
		Map_SpiralOriginDBID_SpiralOriginPtr: make(map[uint]*models.SpiralOrigin, 0),
		Map_SpiralOriginDBID_SpiralOriginDB:  make(map[uint]*SpiralOriginDB, 0),
		Map_SpiralOriginPtr_SpiralOriginDBID: make(map[*models.SpiralOrigin]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralRhombus = BackRepoSpiralRhombusStruct{
		Map_SpiralRhombusDBID_SpiralRhombusPtr: make(map[uint]*models.SpiralRhombus, 0),
		Map_SpiralRhombusDBID_SpiralRhombusDB:  make(map[uint]*SpiralRhombusDB, 0),
		Map_SpiralRhombusPtr_SpiralRhombusDBID: make(map[*models.SpiralRhombus]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSpiralRhombusGrid = BackRepoSpiralRhombusGridStruct{
		Map_SpiralRhombusGridDBID_SpiralRhombusGridPtr: make(map[uint]*models.SpiralRhombusGrid, 0),
		Map_SpiralRhombusGridDBID_SpiralRhombusGridDB:  make(map[uint]*SpiralRhombusGridDB, 0),
		Map_SpiralRhombusGridPtr_SpiralRhombusGridDBID: make(map[*models.SpiralRhombusGrid]uint, 0),

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

	// forbid read of back repo during commit
	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAxis.CommitPhaseOne(stage)
	backRepo.BackRepoAxisGrid.CommitPhaseOne(stage)
	backRepo.BackRepoBezier.CommitPhaseOne(stage)
	backRepo.BackRepoBezierGrid.CommitPhaseOne(stage)
	backRepo.BackRepoBezierGridStack.CommitPhaseOne(stage)
	backRepo.BackRepoCircle.CommitPhaseOne(stage)
	backRepo.BackRepoCircleGrid.CommitPhaseOne(stage)
	backRepo.BackRepoFrontCurve.CommitPhaseOne(stage)
	backRepo.BackRepoFrontCurveStack.CommitPhaseOne(stage)
	backRepo.BackRepoHorizontalAxis.CommitPhaseOne(stage)
	backRepo.BackRepoKey.CommitPhaseOne(stage)
	backRepo.BackRepoMovingLine.CommitPhaseOne(stage)
	backRepo.BackRepoNoteInfo.CommitPhaseOne(stage)
	backRepo.BackRepoParameter.CommitPhaseOne(stage)
	backRepo.BackRepoRhombus.CommitPhaseOne(stage)
	backRepo.BackRepoRhombusGrid.CommitPhaseOne(stage)
	backRepo.BackRepoShapeCategory.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralBezier.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralBezierGrid.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralCircle.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralCircleGrid.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralLine.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralLineGrid.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralOrigin.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralRhombus.CommitPhaseOne(stage)
	backRepo.BackRepoSpiralRhombusGrid.CommitPhaseOne(stage)
	backRepo.BackRepoVerticalAxis.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAxis.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAxisGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBezier.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBezierGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBezierGridStack.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCircle.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCircleGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFrontCurve.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFrontCurveStack.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHorizontalAxis.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKey.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMovingLine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNoteInfo.CommitPhaseTwo(backRepo)
	backRepo.BackRepoParameter.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRhombus.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRhombusGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoShapeCategory.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralBezier.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralBezierGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralCircle.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralCircleGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralLine.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralLineGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralOrigin.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralRhombus.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSpiralRhombusGrid.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVerticalAxis.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAxis.CheckoutPhaseOne()
	backRepo.BackRepoAxisGrid.CheckoutPhaseOne()
	backRepo.BackRepoBezier.CheckoutPhaseOne()
	backRepo.BackRepoBezierGrid.CheckoutPhaseOne()
	backRepo.BackRepoBezierGridStack.CheckoutPhaseOne()
	backRepo.BackRepoCircle.CheckoutPhaseOne()
	backRepo.BackRepoCircleGrid.CheckoutPhaseOne()
	backRepo.BackRepoFrontCurve.CheckoutPhaseOne()
	backRepo.BackRepoFrontCurveStack.CheckoutPhaseOne()
	backRepo.BackRepoHorizontalAxis.CheckoutPhaseOne()
	backRepo.BackRepoKey.CheckoutPhaseOne()
	backRepo.BackRepoMovingLine.CheckoutPhaseOne()
	backRepo.BackRepoNoteInfo.CheckoutPhaseOne()
	backRepo.BackRepoParameter.CheckoutPhaseOne()
	backRepo.BackRepoRhombus.CheckoutPhaseOne()
	backRepo.BackRepoRhombusGrid.CheckoutPhaseOne()
	backRepo.BackRepoShapeCategory.CheckoutPhaseOne()
	backRepo.BackRepoSpiralBezier.CheckoutPhaseOne()
	backRepo.BackRepoSpiralBezierGrid.CheckoutPhaseOne()
	backRepo.BackRepoSpiralCircle.CheckoutPhaseOne()
	backRepo.BackRepoSpiralCircleGrid.CheckoutPhaseOne()
	backRepo.BackRepoSpiralLine.CheckoutPhaseOne()
	backRepo.BackRepoSpiralLineGrid.CheckoutPhaseOne()
	backRepo.BackRepoSpiralOrigin.CheckoutPhaseOne()
	backRepo.BackRepoSpiralRhombus.CheckoutPhaseOne()
	backRepo.BackRepoSpiralRhombusGrid.CheckoutPhaseOne()
	backRepo.BackRepoVerticalAxis.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAxis.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAxisGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBezier.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBezierGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBezierGridStack.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCircle.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCircleGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFrontCurve.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFrontCurveStack.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHorizontalAxis.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKey.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMovingLine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNoteInfo.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoParameter.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRhombus.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRhombusGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoShapeCategory.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralBezier.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralBezierGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralCircle.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralCircleGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralLine.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralLineGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralOrigin.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralRhombus.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSpiralRhombusGrid.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVerticalAxis.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAxis.Backup(dirPath)
	backRepo.BackRepoAxisGrid.Backup(dirPath)
	backRepo.BackRepoBezier.Backup(dirPath)
	backRepo.BackRepoBezierGrid.Backup(dirPath)
	backRepo.BackRepoBezierGridStack.Backup(dirPath)
	backRepo.BackRepoCircle.Backup(dirPath)
	backRepo.BackRepoCircleGrid.Backup(dirPath)
	backRepo.BackRepoFrontCurve.Backup(dirPath)
	backRepo.BackRepoFrontCurveStack.Backup(dirPath)
	backRepo.BackRepoHorizontalAxis.Backup(dirPath)
	backRepo.BackRepoKey.Backup(dirPath)
	backRepo.BackRepoMovingLine.Backup(dirPath)
	backRepo.BackRepoNoteInfo.Backup(dirPath)
	backRepo.BackRepoParameter.Backup(dirPath)
	backRepo.BackRepoRhombus.Backup(dirPath)
	backRepo.BackRepoRhombusGrid.Backup(dirPath)
	backRepo.BackRepoShapeCategory.Backup(dirPath)
	backRepo.BackRepoSpiralBezier.Backup(dirPath)
	backRepo.BackRepoSpiralBezierGrid.Backup(dirPath)
	backRepo.BackRepoSpiralCircle.Backup(dirPath)
	backRepo.BackRepoSpiralCircleGrid.Backup(dirPath)
	backRepo.BackRepoSpiralLine.Backup(dirPath)
	backRepo.BackRepoSpiralLineGrid.Backup(dirPath)
	backRepo.BackRepoSpiralOrigin.Backup(dirPath)
	backRepo.BackRepoSpiralRhombus.Backup(dirPath)
	backRepo.BackRepoSpiralRhombusGrid.Backup(dirPath)
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
	backRepo.BackRepoBezierGrid.BackupXL(file)
	backRepo.BackRepoBezierGridStack.BackupXL(file)
	backRepo.BackRepoCircle.BackupXL(file)
	backRepo.BackRepoCircleGrid.BackupXL(file)
	backRepo.BackRepoFrontCurve.BackupXL(file)
	backRepo.BackRepoFrontCurveStack.BackupXL(file)
	backRepo.BackRepoHorizontalAxis.BackupXL(file)
	backRepo.BackRepoKey.BackupXL(file)
	backRepo.BackRepoMovingLine.BackupXL(file)
	backRepo.BackRepoNoteInfo.BackupXL(file)
	backRepo.BackRepoParameter.BackupXL(file)
	backRepo.BackRepoRhombus.BackupXL(file)
	backRepo.BackRepoRhombusGrid.BackupXL(file)
	backRepo.BackRepoShapeCategory.BackupXL(file)
	backRepo.BackRepoSpiralBezier.BackupXL(file)
	backRepo.BackRepoSpiralBezierGrid.BackupXL(file)
	backRepo.BackRepoSpiralCircle.BackupXL(file)
	backRepo.BackRepoSpiralCircleGrid.BackupXL(file)
	backRepo.BackRepoSpiralLine.BackupXL(file)
	backRepo.BackRepoSpiralLineGrid.BackupXL(file)
	backRepo.BackRepoSpiralOrigin.BackupXL(file)
	backRepo.BackRepoSpiralRhombus.BackupXL(file)
	backRepo.BackRepoSpiralRhombusGrid.BackupXL(file)
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
	backRepo.BackRepoBezierGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoBezierGridStack.RestorePhaseOne(dirPath)
	backRepo.BackRepoCircle.RestorePhaseOne(dirPath)
	backRepo.BackRepoCircleGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoFrontCurve.RestorePhaseOne(dirPath)
	backRepo.BackRepoFrontCurveStack.RestorePhaseOne(dirPath)
	backRepo.BackRepoHorizontalAxis.RestorePhaseOne(dirPath)
	backRepo.BackRepoKey.RestorePhaseOne(dirPath)
	backRepo.BackRepoMovingLine.RestorePhaseOne(dirPath)
	backRepo.BackRepoNoteInfo.RestorePhaseOne(dirPath)
	backRepo.BackRepoParameter.RestorePhaseOne(dirPath)
	backRepo.BackRepoRhombus.RestorePhaseOne(dirPath)
	backRepo.BackRepoRhombusGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoShapeCategory.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralBezier.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralBezierGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralCircle.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralCircleGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralLine.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralLineGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralOrigin.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralRhombus.RestorePhaseOne(dirPath)
	backRepo.BackRepoSpiralRhombusGrid.RestorePhaseOne(dirPath)
	backRepo.BackRepoVerticalAxis.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAxis.RestorePhaseTwo()
	backRepo.BackRepoAxisGrid.RestorePhaseTwo()
	backRepo.BackRepoBezier.RestorePhaseTwo()
	backRepo.BackRepoBezierGrid.RestorePhaseTwo()
	backRepo.BackRepoBezierGridStack.RestorePhaseTwo()
	backRepo.BackRepoCircle.RestorePhaseTwo()
	backRepo.BackRepoCircleGrid.RestorePhaseTwo()
	backRepo.BackRepoFrontCurve.RestorePhaseTwo()
	backRepo.BackRepoFrontCurveStack.RestorePhaseTwo()
	backRepo.BackRepoHorizontalAxis.RestorePhaseTwo()
	backRepo.BackRepoKey.RestorePhaseTwo()
	backRepo.BackRepoMovingLine.RestorePhaseTwo()
	backRepo.BackRepoNoteInfo.RestorePhaseTwo()
	backRepo.BackRepoParameter.RestorePhaseTwo()
	backRepo.BackRepoRhombus.RestorePhaseTwo()
	backRepo.BackRepoRhombusGrid.RestorePhaseTwo()
	backRepo.BackRepoShapeCategory.RestorePhaseTwo()
	backRepo.BackRepoSpiralBezier.RestorePhaseTwo()
	backRepo.BackRepoSpiralBezierGrid.RestorePhaseTwo()
	backRepo.BackRepoSpiralCircle.RestorePhaseTwo()
	backRepo.BackRepoSpiralCircleGrid.RestorePhaseTwo()
	backRepo.BackRepoSpiralLine.RestorePhaseTwo()
	backRepo.BackRepoSpiralLineGrid.RestorePhaseTwo()
	backRepo.BackRepoSpiralOrigin.RestorePhaseTwo()
	backRepo.BackRepoSpiralRhombus.RestorePhaseTwo()
	backRepo.BackRepoSpiralRhombusGrid.RestorePhaseTwo()
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
	backRepo.BackRepoBezierGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoBezierGridStack.RestoreXLPhaseOne(file)
	backRepo.BackRepoCircle.RestoreXLPhaseOne(file)
	backRepo.BackRepoCircleGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoFrontCurve.RestoreXLPhaseOne(file)
	backRepo.BackRepoFrontCurveStack.RestoreXLPhaseOne(file)
	backRepo.BackRepoHorizontalAxis.RestoreXLPhaseOne(file)
	backRepo.BackRepoKey.RestoreXLPhaseOne(file)
	backRepo.BackRepoMovingLine.RestoreXLPhaseOne(file)
	backRepo.BackRepoNoteInfo.RestoreXLPhaseOne(file)
	backRepo.BackRepoParameter.RestoreXLPhaseOne(file)
	backRepo.BackRepoRhombus.RestoreXLPhaseOne(file)
	backRepo.BackRepoRhombusGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoShapeCategory.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralBezier.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralBezierGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralCircle.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralCircleGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralLine.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralLineGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralOrigin.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralRhombus.RestoreXLPhaseOne(file)
	backRepo.BackRepoSpiralRhombusGrid.RestoreXLPhaseOne(file)
	backRepo.BackRepoVerticalAxis.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.subscribersRwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.subscribersRwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.subscribersRwMutex.Lock()
	defer backRepoStruct.subscribersRwMutex.Unlock()
	for i, subscriber := range backRepoStruct.subscribers {
		if subscriber == ch {
			backRepoStruct.subscribers =
				append(backRepoStruct.subscribers[:i],
					backRepoStruct.subscribers[i+1:]...)
			close(ch) // Close the channel to signal completion
			break
		}
	}
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.subscribersRwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.subscribersRwMutex.RUnlock()

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}
