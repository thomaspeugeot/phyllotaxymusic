// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"log"

	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type Stager struct {
	stage      *StageStruct
	splitStage *split.StageStruct
	parameter  *Parameter
}

func NewStager(r *gin.Engine, stage *StageStruct) (stager *Stager) {

	stager = new(Stager)

	// get the only diagram
	parameters := GetGongstructInstancesMap[Parameter](stage)
	parameter, ok := (*parameters)["Reference"]
	if !ok {
		log.Fatal("no Reference parameter on stage")
	}
	stager.parameter = parameter

	// temporary
	parameter.stager = stager

	stager.stage = stage
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	(&split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stage.GetName() + ProbeSplitSuffix,
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
		},
	}).Stage(stager.splitStage)

	stager.splitStage.Commit()

	return
}
