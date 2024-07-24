// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AxisAPIs []*AxisAPI

	AxisGridAPIs []*AxisGridAPI

	CircleAPIs []*CircleAPI

	CircleGridAPIs []*CircleGridAPI

	HorizontalAxisAPIs []*HorizontalAxisAPI

	ParameterAPIs []*ParameterAPI

	RhombusAPIs []*RhombusAPI

	RhombusGridAPIs []*RhombusGridAPI

	VerticalAxisAPIs []*VerticalAxisAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, axisDB := range backRepo.BackRepoAxis.Map_AxisDBID_AxisDB {

		var axisAPI AxisAPI
		axisAPI.ID = axisDB.ID
		axisAPI.AxisPointersEncoding = axisDB.AxisPointersEncoding
		axisDB.CopyBasicFieldsToAxis_WOP(&axisAPI.Axis_WOP)

		backRepoData.AxisAPIs = append(backRepoData.AxisAPIs, &axisAPI)
	}

	for _, axisgridDB := range backRepo.BackRepoAxisGrid.Map_AxisGridDBID_AxisGridDB {

		var axisgridAPI AxisGridAPI
		axisgridAPI.ID = axisgridDB.ID
		axisgridAPI.AxisGridPointersEncoding = axisgridDB.AxisGridPointersEncoding
		axisgridDB.CopyBasicFieldsToAxisGrid_WOP(&axisgridAPI.AxisGrid_WOP)

		backRepoData.AxisGridAPIs = append(backRepoData.AxisGridAPIs, &axisgridAPI)
	}

	for _, circleDB := range backRepo.BackRepoCircle.Map_CircleDBID_CircleDB {

		var circleAPI CircleAPI
		circleAPI.ID = circleDB.ID
		circleAPI.CirclePointersEncoding = circleDB.CirclePointersEncoding
		circleDB.CopyBasicFieldsToCircle_WOP(&circleAPI.Circle_WOP)

		backRepoData.CircleAPIs = append(backRepoData.CircleAPIs, &circleAPI)
	}

	for _, circlegridDB := range backRepo.BackRepoCircleGrid.Map_CircleGridDBID_CircleGridDB {

		var circlegridAPI CircleGridAPI
		circlegridAPI.ID = circlegridDB.ID
		circlegridAPI.CircleGridPointersEncoding = circlegridDB.CircleGridPointersEncoding
		circlegridDB.CopyBasicFieldsToCircleGrid_WOP(&circlegridAPI.CircleGrid_WOP)

		backRepoData.CircleGridAPIs = append(backRepoData.CircleGridAPIs, &circlegridAPI)
	}

	for _, horizontalaxisDB := range backRepo.BackRepoHorizontalAxis.Map_HorizontalAxisDBID_HorizontalAxisDB {

		var horizontalaxisAPI HorizontalAxisAPI
		horizontalaxisAPI.ID = horizontalaxisDB.ID
		horizontalaxisAPI.HorizontalAxisPointersEncoding = horizontalaxisDB.HorizontalAxisPointersEncoding
		horizontalaxisDB.CopyBasicFieldsToHorizontalAxis_WOP(&horizontalaxisAPI.HorizontalAxis_WOP)

		backRepoData.HorizontalAxisAPIs = append(backRepoData.HorizontalAxisAPIs, &horizontalaxisAPI)
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
