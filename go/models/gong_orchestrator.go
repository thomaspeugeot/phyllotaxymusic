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

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case ExportToMusicxml:
		stage.OnAfterExportToMusicxmlUpdateCallback = new(ExportToMusicxmlOrchestrator)

	}

}
