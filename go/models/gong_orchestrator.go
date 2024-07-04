// generated code - do not edit
package models

// insertion point
// DiagramOrchestrator
type DiagramOrchestrator struct {
}

func (orchestrator *DiagramOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedDiagram, backRepoDiagram *Diagram) {

	stagedDiagram.OnAfterUpdate(gongsvgStage, stagedDiagram, backRepoDiagram)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Diagram:
		stage.OnAfterDiagramUpdateCallback = new(DiagramOrchestrator)

	}

}
