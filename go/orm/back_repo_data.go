// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	DiagramAPIs []*DiagramAPI

	LineAPIs []*LineAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, diagramDB := range backRepo.BackRepoDiagram.Map_DiagramDBID_DiagramDB {

		var diagramAPI DiagramAPI
		diagramAPI.ID = diagramDB.ID
		diagramAPI.DiagramPointersEncoding = diagramDB.DiagramPointersEncoding
		diagramDB.CopyBasicFieldsToDiagram_WOP(&diagramAPI.Diagram_WOP)

		backRepoData.DiagramAPIs = append(backRepoData.DiagramAPIs, &diagramAPI)
	}

	for _, lineDB := range backRepo.BackRepoLine.Map_LineDBID_LineDB {

		var lineAPI LineAPI
		lineAPI.ID = lineDB.ID
		lineAPI.LinePointersEncoding = lineDB.LinePointersEncoding
		lineDB.CopyBasicFieldsToLine_WOP(&lineAPI.Line_WOP)

		backRepoData.LineAPIs = append(backRepoData.LineAPIs, &lineAPI)
	}

}
