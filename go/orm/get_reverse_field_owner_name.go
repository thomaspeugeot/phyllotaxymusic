// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Axis:
		switch reverseField.GongstructName {
		// insertion point
		case "AxisGrid":
			switch reverseField.Fieldname {
			case "Axiss":
				if _axisgrid, ok := stage.AxisGrid_Axiss_reverseMap[inst]; ok {
					res = _axisgrid.Name
				}
			}
		}

	case *models.AxisGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bezier:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierGrid":
			switch reverseField.Fieldname {
			case "Beziers":
				if _beziergrid, ok := stage.BezierGrid_Beziers_reverseMap[inst]; ok {
					res = _beziergrid.Name
				}
			}
		}

	case *models.BezierGrid:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierGridStack":
			switch reverseField.Fieldname {
			case "BezierGrids":
				if _beziergridstack, ok := stage.BezierGridStack_BezierGrids_reverseMap[inst]; ok {
					res = _beziergridstack.Name
				}
			}
		}

	case *models.BezierGridStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Circle:
		switch reverseField.GongstructName {
		// insertion point
		case "CircleGrid":
			switch reverseField.Fieldname {
			case "Circles":
				if _circlegrid, ok := stage.CircleGrid_Circles_reverseMap[inst]; ok {
					res = _circlegrid.Name
				}
			}
		}

	case *models.CircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FrontCurve:
		switch reverseField.GongstructName {
		// insertion point
		case "FrontCurveStack":
			switch reverseField.Fieldname {
			case "FrontCurves":
				if _frontcurvestack, ok := stage.FrontCurveStack_FrontCurves_reverseMap[inst]; ok {
					res = _frontcurvestack.Name
				}
			}
		}

	case *models.FrontCurveStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.HorizontalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.NoteInfo:
		switch reverseField.GongstructName {
		// insertion point
		case "Parameter":
			switch reverseField.Fieldname {
			case "NoteInfos":
				if _parameter, ok := stage.Parameter_NoteInfos_reverseMap[inst]; ok {
					res = _parameter.Name
				}
			}
		}

	case *models.Parameter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Rhombus:
		switch reverseField.GongstructName {
		// insertion point
		case "RhombusGrid":
			switch reverseField.Fieldname {
			case "Rhombuses":
				if _rhombusgrid, ok := stage.RhombusGrid_Rhombuses_reverseMap[inst]; ok {
					res = _rhombusgrid.Name
				}
			}
		}

	case *models.RhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ShapeCategory:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralBezier:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralBezierGrid":
			switch reverseField.Fieldname {
			case "SpiralBeziers":
				if _spiralbeziergrid, ok := stage.SpiralBezierGrid_SpiralBeziers_reverseMap[inst]; ok {
					res = _spiralbeziergrid.Name
				}
			}
		}

	case *models.SpiralBezierGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralCircle:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralCircleGrid":
			switch reverseField.Fieldname {
			case "SpiralCircles":
				if _spiralcirclegrid, ok := stage.SpiralCircleGrid_SpiralCircles_reverseMap[inst]; ok {
					res = _spiralcirclegrid.Name
				}
			}
		}

	case *models.SpiralCircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralLine:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralLineGrid":
			switch reverseField.Fieldname {
			case "SpiralLines":
				if _spirallinegrid, ok := stage.SpiralLineGrid_SpiralLines_reverseMap[inst]; ok {
					res = _spirallinegrid.Name
				}
			}
		}

	case *models.SpiralLineGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralOrigin:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralRhombus:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralRhombusGrid":
			switch reverseField.Fieldname {
			case "SpiralRhombuses":
				if _spiralrhombusgrid, ok := stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap[inst]; ok {
					res = _spiralrhombusgrid.Name
				}
			}
		}

	case *models.SpiralRhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VerticalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Axis:
		switch reverseField.GongstructName {
		// insertion point
		case "AxisGrid":
			switch reverseField.Fieldname {
			case "Axiss":
				res = stage.AxisGrid_Axiss_reverseMap[inst]
			}
		}

	case *models.AxisGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bezier:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierGrid":
			switch reverseField.Fieldname {
			case "Beziers":
				res = stage.BezierGrid_Beziers_reverseMap[inst]
			}
		}

	case *models.BezierGrid:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierGridStack":
			switch reverseField.Fieldname {
			case "BezierGrids":
				res = stage.BezierGridStack_BezierGrids_reverseMap[inst]
			}
		}

	case *models.BezierGridStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Circle:
		switch reverseField.GongstructName {
		// insertion point
		case "CircleGrid":
			switch reverseField.Fieldname {
			case "Circles":
				res = stage.CircleGrid_Circles_reverseMap[inst]
			}
		}

	case *models.CircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FrontCurve:
		switch reverseField.GongstructName {
		// insertion point
		case "FrontCurveStack":
			switch reverseField.Fieldname {
			case "FrontCurves":
				res = stage.FrontCurveStack_FrontCurves_reverseMap[inst]
			}
		}

	case *models.FrontCurveStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.HorizontalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.NoteInfo:
		switch reverseField.GongstructName {
		// insertion point
		case "Parameter":
			switch reverseField.Fieldname {
			case "NoteInfos":
				res = stage.Parameter_NoteInfos_reverseMap[inst]
			}
		}

	case *models.Parameter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Rhombus:
		switch reverseField.GongstructName {
		// insertion point
		case "RhombusGrid":
			switch reverseField.Fieldname {
			case "Rhombuses":
				res = stage.RhombusGrid_Rhombuses_reverseMap[inst]
			}
		}

	case *models.RhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ShapeCategory:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralBezier:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralBezierGrid":
			switch reverseField.Fieldname {
			case "SpiralBeziers":
				res = stage.SpiralBezierGrid_SpiralBeziers_reverseMap[inst]
			}
		}

	case *models.SpiralBezierGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralCircle:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralCircleGrid":
			switch reverseField.Fieldname {
			case "SpiralCircles":
				res = stage.SpiralCircleGrid_SpiralCircles_reverseMap[inst]
			}
		}

	case *models.SpiralCircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralLine:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralLineGrid":
			switch reverseField.Fieldname {
			case "SpiralLines":
				res = stage.SpiralLineGrid_SpiralLines_reverseMap[inst]
			}
		}

	case *models.SpiralLineGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralOrigin:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SpiralRhombus:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralRhombusGrid":
			switch reverseField.Fieldname {
			case "SpiralRhombuses":
				res = stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap[inst]
			}
		}

	case *models.SpiralRhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VerticalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
