// do not modify, generated file
package fullstack

import (
	"github.com/thomaspeugeot/phyllotaxymusic/go/controllers"
	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
	"github.com/thomaspeugeot/phyllotaxymusic/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/thomaspeugeot/phyllotaxymusic/ng-github.com-thomaspeugeot-phyllotaxymusic"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.Stage,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	controllers.GetController().AddBackRepo(backRepo, stackPath)

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.Axis](stage)
	models.SetOrchestratorOnAfterUpdate[models.AxisGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bezier](stage)
	models.SetOrchestratorOnAfterUpdate[models.BezierGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.BezierGridStack](stage)
	models.SetOrchestratorOnAfterUpdate[models.Chapter](stage)
	models.SetOrchestratorOnAfterUpdate[models.Circle](stage)
	models.SetOrchestratorOnAfterUpdate[models.CircleGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.Content](stage)
	models.SetOrchestratorOnAfterUpdate[models.ExportToMusicxml](stage)
	models.SetOrchestratorOnAfterUpdate[models.FrontCurve](stage)
	models.SetOrchestratorOnAfterUpdate[models.FrontCurveStack](stage)
	models.SetOrchestratorOnAfterUpdate[models.HorizontalAxis](stage)
	models.SetOrchestratorOnAfterUpdate[models.Key](stage)
	models.SetOrchestratorOnAfterUpdate[models.Parameter](stage)
	models.SetOrchestratorOnAfterUpdate[models.Rhombus](stage)
	models.SetOrchestratorOnAfterUpdate[models.RhombusGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.ShapeCategory](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralBezier](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralBezierGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralCircle](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralCircleGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralLine](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralLineGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralOrigin](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralRhombus](stage)
	models.SetOrchestratorOnAfterUpdate[models.SpiralRhombusGrid](stage)
	models.SetOrchestratorOnAfterUpdate[models.VerticalAxis](stage)

	return
}
