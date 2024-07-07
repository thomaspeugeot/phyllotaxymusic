// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	HorizontalAxisAPIs []*HorizontalAxisAPI

	InitialAxisAPIs []*InitialAxisAPI

	LineAPIs []*LineAPI

	ParameterAPIs []*ParameterAPI

	RhombusAPIs []*RhombusAPI

	RhombusGridAPIs []*RhombusGridAPI

	VerticalAxisAPIs []*VerticalAxisAPI
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

	for _, initialaxisDB := range backRepo.BackRepoInitialAxis.Map_InitialAxisDBID_InitialAxisDB {

		var initialaxisAPI InitialAxisAPI
		initialaxisAPI.ID = initialaxisDB.ID
		initialaxisAPI.InitialAxisPointersEncoding = initialaxisDB.InitialAxisPointersEncoding
		initialaxisDB.CopyBasicFieldsToInitialAxis_WOP(&initialaxisAPI.InitialAxis_WOP)

		backRepoData.InitialAxisAPIs = append(backRepoData.InitialAxisAPIs, &initialaxisAPI)
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

	for _, rhombusgridDB := range backRepo.BackRepoRhombusGrid.Map_RhombusGridDBID_RhombusGridDB {

		var rhombusgridAPI RhombusGridAPI
		rhombusgridAPI.ID = rhombusgridDB.ID
		rhombusgridAPI.RhombusGridPointersEncoding = rhombusgridDB.RhombusGridPointersEncoding
		rhombusgridDB.CopyBasicFieldsToRhombusGrid_WOP(&rhombusgridAPI.RhombusGrid_WOP)

		backRepoData.RhombusGridAPIs = append(backRepoData.RhombusGridAPIs, &rhombusgridAPI)
	}

	for _, verticalaxisDB := range backRepo.BackRepoVerticalAxis.Map_VerticalAxisDBID_VerticalAxisDB {

		var verticalaxisAPI VerticalAxisAPI
		verticalaxisAPI.ID = verticalaxisDB.ID
		verticalaxisAPI.VerticalAxisPointersEncoding = verticalaxisDB.VerticalAxisPointersEncoding
		verticalaxisDB.CopyBasicFieldsToVerticalAxis_WOP(&verticalaxisAPI.VerticalAxis_WOP)

		backRepoData.VerticalAxisAPIs = append(backRepoData.VerticalAxisAPIs, &verticalaxisAPI)
	}

}
