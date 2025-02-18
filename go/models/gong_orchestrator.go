// generated code - do not edit
package models

// insertion point
// ExportToMusicxmlOrchestrator
type ExportToMusicxmlOrchestrator struct {
}

func (orchestrator *ExportToMusicxmlOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedExportToMusicxml, backRepoExportToMusicxml *ExportToMusicxml) {

	stagedExportToMusicxml.OnAfterUpdate(gongsvgStage, stagedExportToMusicxml, backRepoExportToMusicxml)
}
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
	case ExportToMusicxml:
		stage.OnAfterExportToMusicxmlUpdateCallback = new(ExportToMusicxmlOrchestrator)
	case Parameter:
		stage.OnAfterParameterUpdateCallback = new(ParameterOrchestrator)

	}

}
