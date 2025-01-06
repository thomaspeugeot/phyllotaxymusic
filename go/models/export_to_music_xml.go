package models

import "github.com/thomaspeugeot/phyllotaxymusic/go/models/musicxml"

// ExportToMusicxml is for catching user interactions with the export button
// gong will forward any update to an instance of the object
//
// it is necessary to have one instance in the persistance
type ExportToMusicxml struct {
	Name string

	Parameter *Parameter
}

func (exportToMusicxml *ExportToMusicxml) OnAfterUpdate(
	phyllotaxyStage *StageStruct,
	stagedExportToMusicxml, backRepoExportToMusicxml *ExportToMusicxml) {

	shouldReturn := musicxml.GenerateMusicXMLFile()
	if shouldReturn {
		return
	}
}
