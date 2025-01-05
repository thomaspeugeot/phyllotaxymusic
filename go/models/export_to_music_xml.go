package models

import "log"

// ExportToMusicxml is for catching user interactions with the export button
// gong will forward any update to an instance of the object
//
// it is necessary to have one instance in the persistance
type ExportToMusicxml struct {
	Name string
}

func (exporttomusicxml *ExportToMusicxml) OnAfterUpdate(phyllotaxyStage *StageStruct, stagedExportToMusicxml, backRepoExportToMusicxml *ExportToMusicxml) {

	log.Println("Export to Music xml requested")
}
