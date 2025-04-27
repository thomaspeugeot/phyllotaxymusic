package models

// ExportToMusicxml is for catching user interactions with the export button
// gong will forward any update to an instance of the object
//
// it is necessary to have one instance in the persistance
type ExportToMusicxml struct {
	Name string

	Parameter *Parameter
}

func (exportToMusicxml *ExportToMusicxml) OnAfterUpdate(
	phyllotaxyStage *Stage,
	stagedExportToMusicxml, backRepoExportToMusicxml *ExportToMusicxml) {

	exportToMusicxml.Parameter.stager.UpdatePhyllotaxyStage()
	exportToMusicxml.Parameter.GenerateMusicXMLFile(DOWNLOAD)
}
