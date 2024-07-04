// generated code - do not edit
package models

// insertion point
// ParameterOrchestrator
type ParameterOrchestrator struct {
}

func (orchestrator *ParameterOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedParameter, backRepoParameter *Parameter) {

	stagedParameter.OnAfterUpdate(gongsvgStage, stagedParameter, backRepoParameter)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Parameter:
		stage.OnAfterParameterUpdateCallback = new(ParameterOrchestrator)

	}

}
