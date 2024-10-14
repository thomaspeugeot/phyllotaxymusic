// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AxisAPIs []*AxisAPI

	AxisGridAPIs []*AxisGridAPI

	BezierAPIs []*BezierAPI

	BezierGridAPIs []*BezierGridAPI

	BezierGridStackAPIs []*BezierGridStackAPI

	CircleAPIs []*CircleAPI

	CircleGridAPIs []*CircleGridAPI

	FrontCurveAPIs []*FrontCurveAPI

	FrontCurveStackAPIs []*FrontCurveStackAPI

	HorizontalAxisAPIs []*HorizontalAxisAPI

	KeyAPIs []*KeyAPI

	NoteInfoAPIs []*NoteInfoAPI

	ParameterAPIs []*ParameterAPI

	RhombusAPIs []*RhombusAPI

	RhombusGridAPIs []*RhombusGridAPI

	ShapeCategoryAPIs []*ShapeCategoryAPI

	SpiralBezierAPIs []*SpiralBezierAPI

	SpiralBezierGridAPIs []*SpiralBezierGridAPI

	SpiralCircleAPIs []*SpiralCircleAPI

	SpiralCircleGridAPIs []*SpiralCircleGridAPI

	SpiralLineAPIs []*SpiralLineAPI

	SpiralLineGridAPIs []*SpiralLineGridAPI

	SpiralOriginAPIs []*SpiralOriginAPI

	SpiralRhombusAPIs []*SpiralRhombusAPI

	SpiralRhombusGridAPIs []*SpiralRhombusGridAPI

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

	for _, bezierDB := range backRepo.BackRepoBezier.Map_BezierDBID_BezierDB {

		var bezierAPI BezierAPI
		bezierAPI.ID = bezierDB.ID
		bezierAPI.BezierPointersEncoding = bezierDB.BezierPointersEncoding
		bezierDB.CopyBasicFieldsToBezier_WOP(&bezierAPI.Bezier_WOP)

		backRepoData.BezierAPIs = append(backRepoData.BezierAPIs, &bezierAPI)
	}

	for _, beziergridDB := range backRepo.BackRepoBezierGrid.Map_BezierGridDBID_BezierGridDB {

		var beziergridAPI BezierGridAPI
		beziergridAPI.ID = beziergridDB.ID
		beziergridAPI.BezierGridPointersEncoding = beziergridDB.BezierGridPointersEncoding
		beziergridDB.CopyBasicFieldsToBezierGrid_WOP(&beziergridAPI.BezierGrid_WOP)

		backRepoData.BezierGridAPIs = append(backRepoData.BezierGridAPIs, &beziergridAPI)
	}

	for _, beziergridstackDB := range backRepo.BackRepoBezierGridStack.Map_BezierGridStackDBID_BezierGridStackDB {

		var beziergridstackAPI BezierGridStackAPI
		beziergridstackAPI.ID = beziergridstackDB.ID
		beziergridstackAPI.BezierGridStackPointersEncoding = beziergridstackDB.BezierGridStackPointersEncoding
		beziergridstackDB.CopyBasicFieldsToBezierGridStack_WOP(&beziergridstackAPI.BezierGridStack_WOP)

		backRepoData.BezierGridStackAPIs = append(backRepoData.BezierGridStackAPIs, &beziergridstackAPI)
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

	for _, frontcurveDB := range backRepo.BackRepoFrontCurve.Map_FrontCurveDBID_FrontCurveDB {

		var frontcurveAPI FrontCurveAPI
		frontcurveAPI.ID = frontcurveDB.ID
		frontcurveAPI.FrontCurvePointersEncoding = frontcurveDB.FrontCurvePointersEncoding
		frontcurveDB.CopyBasicFieldsToFrontCurve_WOP(&frontcurveAPI.FrontCurve_WOP)

		backRepoData.FrontCurveAPIs = append(backRepoData.FrontCurveAPIs, &frontcurveAPI)
	}

	for _, frontcurvestackDB := range backRepo.BackRepoFrontCurveStack.Map_FrontCurveStackDBID_FrontCurveStackDB {

		var frontcurvestackAPI FrontCurveStackAPI
		frontcurvestackAPI.ID = frontcurvestackDB.ID
		frontcurvestackAPI.FrontCurveStackPointersEncoding = frontcurvestackDB.FrontCurveStackPointersEncoding
		frontcurvestackDB.CopyBasicFieldsToFrontCurveStack_WOP(&frontcurvestackAPI.FrontCurveStack_WOP)

		backRepoData.FrontCurveStackAPIs = append(backRepoData.FrontCurveStackAPIs, &frontcurvestackAPI)
	}

	for _, horizontalaxisDB := range backRepo.BackRepoHorizontalAxis.Map_HorizontalAxisDBID_HorizontalAxisDB {

		var horizontalaxisAPI HorizontalAxisAPI
		horizontalaxisAPI.ID = horizontalaxisDB.ID
		horizontalaxisAPI.HorizontalAxisPointersEncoding = horizontalaxisDB.HorizontalAxisPointersEncoding
		horizontalaxisDB.CopyBasicFieldsToHorizontalAxis_WOP(&horizontalaxisAPI.HorizontalAxis_WOP)

		backRepoData.HorizontalAxisAPIs = append(backRepoData.HorizontalAxisAPIs, &horizontalaxisAPI)
	}

	for _, keyDB := range backRepo.BackRepoKey.Map_KeyDBID_KeyDB {

		var keyAPI KeyAPI
		keyAPI.ID = keyDB.ID
		keyAPI.KeyPointersEncoding = keyDB.KeyPointersEncoding
		keyDB.CopyBasicFieldsToKey_WOP(&keyAPI.Key_WOP)

		backRepoData.KeyAPIs = append(backRepoData.KeyAPIs, &keyAPI)
	}

	for _, noteinfoDB := range backRepo.BackRepoNoteInfo.Map_NoteInfoDBID_NoteInfoDB {

		var noteinfoAPI NoteInfoAPI
		noteinfoAPI.ID = noteinfoDB.ID
		noteinfoAPI.NoteInfoPointersEncoding = noteinfoDB.NoteInfoPointersEncoding
		noteinfoDB.CopyBasicFieldsToNoteInfo_WOP(&noteinfoAPI.NoteInfo_WOP)

		backRepoData.NoteInfoAPIs = append(backRepoData.NoteInfoAPIs, &noteinfoAPI)
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

	for _, shapecategoryDB := range backRepo.BackRepoShapeCategory.Map_ShapeCategoryDBID_ShapeCategoryDB {

		var shapecategoryAPI ShapeCategoryAPI
		shapecategoryAPI.ID = shapecategoryDB.ID
		shapecategoryAPI.ShapeCategoryPointersEncoding = shapecategoryDB.ShapeCategoryPointersEncoding
		shapecategoryDB.CopyBasicFieldsToShapeCategory_WOP(&shapecategoryAPI.ShapeCategory_WOP)

		backRepoData.ShapeCategoryAPIs = append(backRepoData.ShapeCategoryAPIs, &shapecategoryAPI)
	}

	for _, spiralbezierDB := range backRepo.BackRepoSpiralBezier.Map_SpiralBezierDBID_SpiralBezierDB {

		var spiralbezierAPI SpiralBezierAPI
		spiralbezierAPI.ID = spiralbezierDB.ID
		spiralbezierAPI.SpiralBezierPointersEncoding = spiralbezierDB.SpiralBezierPointersEncoding
		spiralbezierDB.CopyBasicFieldsToSpiralBezier_WOP(&spiralbezierAPI.SpiralBezier_WOP)

		backRepoData.SpiralBezierAPIs = append(backRepoData.SpiralBezierAPIs, &spiralbezierAPI)
	}

	for _, spiralbeziergridDB := range backRepo.BackRepoSpiralBezierGrid.Map_SpiralBezierGridDBID_SpiralBezierGridDB {

		var spiralbeziergridAPI SpiralBezierGridAPI
		spiralbeziergridAPI.ID = spiralbeziergridDB.ID
		spiralbeziergridAPI.SpiralBezierGridPointersEncoding = spiralbeziergridDB.SpiralBezierGridPointersEncoding
		spiralbeziergridDB.CopyBasicFieldsToSpiralBezierGrid_WOP(&spiralbeziergridAPI.SpiralBezierGrid_WOP)

		backRepoData.SpiralBezierGridAPIs = append(backRepoData.SpiralBezierGridAPIs, &spiralbeziergridAPI)
	}

	for _, spiralcircleDB := range backRepo.BackRepoSpiralCircle.Map_SpiralCircleDBID_SpiralCircleDB {

		var spiralcircleAPI SpiralCircleAPI
		spiralcircleAPI.ID = spiralcircleDB.ID
		spiralcircleAPI.SpiralCirclePointersEncoding = spiralcircleDB.SpiralCirclePointersEncoding
		spiralcircleDB.CopyBasicFieldsToSpiralCircle_WOP(&spiralcircleAPI.SpiralCircle_WOP)

		backRepoData.SpiralCircleAPIs = append(backRepoData.SpiralCircleAPIs, &spiralcircleAPI)
	}

	for _, spiralcirclegridDB := range backRepo.BackRepoSpiralCircleGrid.Map_SpiralCircleGridDBID_SpiralCircleGridDB {

		var spiralcirclegridAPI SpiralCircleGridAPI
		spiralcirclegridAPI.ID = spiralcirclegridDB.ID
		spiralcirclegridAPI.SpiralCircleGridPointersEncoding = spiralcirclegridDB.SpiralCircleGridPointersEncoding
		spiralcirclegridDB.CopyBasicFieldsToSpiralCircleGrid_WOP(&spiralcirclegridAPI.SpiralCircleGrid_WOP)

		backRepoData.SpiralCircleGridAPIs = append(backRepoData.SpiralCircleGridAPIs, &spiralcirclegridAPI)
	}

	for _, spirallineDB := range backRepo.BackRepoSpiralLine.Map_SpiralLineDBID_SpiralLineDB {

		var spirallineAPI SpiralLineAPI
		spirallineAPI.ID = spirallineDB.ID
		spirallineAPI.SpiralLinePointersEncoding = spirallineDB.SpiralLinePointersEncoding
		spirallineDB.CopyBasicFieldsToSpiralLine_WOP(&spirallineAPI.SpiralLine_WOP)

		backRepoData.SpiralLineAPIs = append(backRepoData.SpiralLineAPIs, &spirallineAPI)
	}

	for _, spirallinegridDB := range backRepo.BackRepoSpiralLineGrid.Map_SpiralLineGridDBID_SpiralLineGridDB {

		var spirallinegridAPI SpiralLineGridAPI
		spirallinegridAPI.ID = spirallinegridDB.ID
		spirallinegridAPI.SpiralLineGridPointersEncoding = spirallinegridDB.SpiralLineGridPointersEncoding
		spirallinegridDB.CopyBasicFieldsToSpiralLineGrid_WOP(&spirallinegridAPI.SpiralLineGrid_WOP)

		backRepoData.SpiralLineGridAPIs = append(backRepoData.SpiralLineGridAPIs, &spirallinegridAPI)
	}

	for _, spiraloriginDB := range backRepo.BackRepoSpiralOrigin.Map_SpiralOriginDBID_SpiralOriginDB {

		var spiraloriginAPI SpiralOriginAPI
		spiraloriginAPI.ID = spiraloriginDB.ID
		spiraloriginAPI.SpiralOriginPointersEncoding = spiraloriginDB.SpiralOriginPointersEncoding
		spiraloriginDB.CopyBasicFieldsToSpiralOrigin_WOP(&spiraloriginAPI.SpiralOrigin_WOP)

		backRepoData.SpiralOriginAPIs = append(backRepoData.SpiralOriginAPIs, &spiraloriginAPI)
	}

	for _, spiralrhombusDB := range backRepo.BackRepoSpiralRhombus.Map_SpiralRhombusDBID_SpiralRhombusDB {

		var spiralrhombusAPI SpiralRhombusAPI
		spiralrhombusAPI.ID = spiralrhombusDB.ID
		spiralrhombusAPI.SpiralRhombusPointersEncoding = spiralrhombusDB.SpiralRhombusPointersEncoding
		spiralrhombusDB.CopyBasicFieldsToSpiralRhombus_WOP(&spiralrhombusAPI.SpiralRhombus_WOP)

		backRepoData.SpiralRhombusAPIs = append(backRepoData.SpiralRhombusAPIs, &spiralrhombusAPI)
	}

	for _, spiralrhombusgridDB := range backRepo.BackRepoSpiralRhombusGrid.Map_SpiralRhombusGridDBID_SpiralRhombusGridDB {

		var spiralrhombusgridAPI SpiralRhombusGridAPI
		spiralrhombusgridAPI.ID = spiralrhombusgridDB.ID
		spiralrhombusgridAPI.SpiralRhombusGridPointersEncoding = spiralrhombusgridDB.SpiralRhombusGridPointersEncoding
		spiralrhombusgridDB.CopyBasicFieldsToSpiralRhombusGrid_WOP(&spiralrhombusgridAPI.SpiralRhombusGrid_WOP)

		backRepoData.SpiralRhombusGridAPIs = append(backRepoData.SpiralRhombusGridAPIs, &spiralrhombusgridAPI)
	}

	for _, verticalaxisDB := range backRepo.BackRepoVerticalAxis.Map_VerticalAxisDBID_VerticalAxisDB {

		var verticalaxisAPI VerticalAxisAPI
		verticalaxisAPI.ID = verticalaxisDB.ID
		verticalaxisAPI.VerticalAxisPointersEncoding = verticalaxisDB.VerticalAxisPointersEncoding
		verticalaxisDB.CopyBasicFieldsToVerticalAxis_WOP(&verticalaxisAPI.VerticalAxis_WOP)

		backRepoData.VerticalAxisAPIs = append(backRepoData.VerticalAxisAPIs, &verticalaxisAPI)
	}

}
