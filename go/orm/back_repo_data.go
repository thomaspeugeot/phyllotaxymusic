// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	HorizontalAxisAPIs []*HorizontalAxisAPI

	LineAPIs []*LineAPI

	ParameterAPIs []*ParameterAPI

	RhombusAPIs []*RhombusAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, horizontalaxisDB := range backRepo.BackRepoHorizontalAxis.Map_HorizontalAxisDBID_HorizontalAxisDB {

		var horizontalaxisAPI HorizontalAxisAPI
		horizontalaxisAPI.ID = horizontalaxisDB.ID
		horizontalaxisAPI.HorizontalAxisPointersEncoding = horizontalaxisDB.HorizontalAxisPointersEncoding
		horizontalaxisDB.CopyBasicFieldsToHorizontalAxis_WOP(&horizontalaxisAPI.HorizontalAxis_WOP)

		backRepoData.HorizontalAxisAPIs = append(backRepoData.HorizontalAxisAPIs, &horizontalaxisAPI)
	}

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

	for _, rhombusDB := range backRepo.BackRepoRhombus.Map_RhombusDBID_RhombusDB {

		var rhombusAPI RhombusAPI
		rhombusAPI.ID = rhombusDB.ID
		rhombusAPI.RhombusPointersEncoding = rhombusDB.RhombusPointersEncoding
		rhombusDB.CopyBasicFieldsToRhombus_WOP(&rhombusAPI.Rhombus_WOP)

		backRepoData.RhombusAPIs = append(backRepoData.RhombusAPIs, &rhombusAPI)
	}

}
