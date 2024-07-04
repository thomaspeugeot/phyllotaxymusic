// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	LineAPIs []*LineAPI

	ParameterAPIs []*ParameterAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, lineDB := range backRepo.BackRepoLine.Map_LineDBID_LineDB {

		var lineAPI LineAPI
		lineAPI.ID = lineDB.ID
		lineAPI.LinePointersEncoding = lineDB.LinePointersEncoding
		lineDB.CopyBasicFieldsToLine_WOP(&lineAPI.Line_WOP)

		backRepoData.LineAPIs = append(backRepoData.LineAPIs, &lineAPI)
	}

	for _, parameterDB := range backRepo.BackRepoParameter.Map_ParameterDBID_ParameterDB {

		var parameterAPI ParameterAPI
		parameterAPI.ID = parameterDB.ID
		parameterAPI.ParameterPointersEncoding = parameterDB.ParameterPointersEncoding
		parameterDB.CopyBasicFieldsToParameter_WOP(&parameterAPI.Parameter_WOP)

		backRepoData.ParameterAPIs = append(backRepoData.ParameterAPIs, &parameterAPI)
	}

}
