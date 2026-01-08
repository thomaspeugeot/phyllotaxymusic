// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
	"github.com/thomaspeugeot/phyllotaxymusic/go/probe"

	phyllotaxymusic_go "github.com/thomaspeugeot/phyllotaxymusic/go"

	"github.com/gin-gonic/gin"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.Stage) {

	// the ".go" is not provided
	filename := impl.marshallOnCommit
	if !strings.HasSuffix(filename, ".go") {
		filename = filename + ".go"
	}

	file, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	packageName := impl.packageName
	if packageName == "" {
		packageName = "main"
	}

	stage.Marshall(file, "github.com/thomaspeugeot/phyllotaxymusic/go/models", packageName)
}

type Level1Stack struct {
	Stage *models.Stage
	Probe *probe.Probe
	R     *gin.Engine
}

func NewLevel1Stack(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
) (level1Stack *Level1Stack) {

	level1Stack = new(Level1Stack)
	stage := models.NewStage(stackPath)
	level1Stack.Stage = stage

	if unmarshallFromCode != "" {
		err := models.ParseAstFile(stage, unmarshallFromCode, true)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
		stage.ComputeReference()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		hook.marshallOnCommit = marshallOnCommit
		stage.OnInitCommitCallback = hook
	}

	level1Stack.R = split_static.ServeStaticFiles(false)
	if withProbe {
		// if the application edits the diagrams via the probe, it is surmised
		// that the application is launched from "go/cmd/<appl>/". Therefore, to reach
		// "go/diagrams/diagrams.go", the path is "../../diagrams/diagrams.go"
		level1Stack.Probe = probe.NewProbe(
			level1Stack.R,
			phyllotaxymusic_go.GoModelsDir,
			phyllotaxymusic_go.GoDiagramsDir,
			embeddedDiagrams,
			stage,
		)

		stage.SetProbeIF(level1Stack.Probe)
	}

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
