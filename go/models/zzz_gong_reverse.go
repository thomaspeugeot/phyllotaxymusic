// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Axis:
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

	case *AxisGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Bezier:
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

	case *BezierGrid:
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

	case *BezierGridStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Circle:
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

	case *CircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ExportToMusicxml:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FrontCurve:
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

	case *FrontCurveStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *HorizontalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Parameter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Rhombus:
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

	case *RhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ShapeCategory:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralBezier:
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

	case *SpiralBezierGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralCircle:
		switch reverseField.GongstructName {
		// insertion point
		case "FrontCurveStack":
			switch reverseField.Fieldname {
			case "SpiralCircles":
				if _frontcurvestack, ok := stage.FrontCurveStack_SpiralCircles_reverseMap[inst]; ok {
					res = _frontcurvestack.Name
				}
			}
		case "SpiralCircleGrid":
			switch reverseField.Fieldname {
			case "SpiralCircles":
				if _spiralcirclegrid, ok := stage.SpiralCircleGrid_SpiralCircles_reverseMap[inst]; ok {
					res = _spiralcirclegrid.Name
				}
			}
		}

	case *SpiralCircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralLine:
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

	case *SpiralLineGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralOrigin:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralRhombus:
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

	case *SpiralRhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *VerticalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *Axis:
		switch reverseField.GongstructName {
		// insertion point
		case "AxisGrid":
			switch reverseField.Fieldname {
			case "Axiss":
				res = stage.AxisGrid_Axiss_reverseMap[inst]
			}
		}

	case *AxisGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Bezier:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierGrid":
			switch reverseField.Fieldname {
			case "Beziers":
				res = stage.BezierGrid_Beziers_reverseMap[inst]
			}
		}

	case *BezierGrid:
		switch reverseField.GongstructName {
		// insertion point
		case "BezierGridStack":
			switch reverseField.Fieldname {
			case "BezierGrids":
				res = stage.BezierGridStack_BezierGrids_reverseMap[inst]
			}
		}

	case *BezierGridStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Circle:
		switch reverseField.GongstructName {
		// insertion point
		case "CircleGrid":
			switch reverseField.Fieldname {
			case "Circles":
				res = stage.CircleGrid_Circles_reverseMap[inst]
			}
		}

	case *CircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ExportToMusicxml:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FrontCurve:
		switch reverseField.GongstructName {
		// insertion point
		case "FrontCurveStack":
			switch reverseField.Fieldname {
			case "FrontCurves":
				res = stage.FrontCurveStack_FrontCurves_reverseMap[inst]
			}
		}

	case *FrontCurveStack:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *HorizontalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Parameter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Rhombus:
		switch reverseField.GongstructName {
		// insertion point
		case "RhombusGrid":
			switch reverseField.Fieldname {
			case "Rhombuses":
				res = stage.RhombusGrid_Rhombuses_reverseMap[inst]
			}
		}

	case *RhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ShapeCategory:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralBezier:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralBezierGrid":
			switch reverseField.Fieldname {
			case "SpiralBeziers":
				res = stage.SpiralBezierGrid_SpiralBeziers_reverseMap[inst]
			}
		}

	case *SpiralBezierGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralCircle:
		switch reverseField.GongstructName {
		// insertion point
		case "FrontCurveStack":
			switch reverseField.Fieldname {
			case "SpiralCircles":
				res = stage.FrontCurveStack_SpiralCircles_reverseMap[inst]
			}
		case "SpiralCircleGrid":
			switch reverseField.Fieldname {
			case "SpiralCircles":
				res = stage.SpiralCircleGrid_SpiralCircles_reverseMap[inst]
			}
		}

	case *SpiralCircleGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralLine:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralLineGrid":
			switch reverseField.Fieldname {
			case "SpiralLines":
				res = stage.SpiralLineGrid_SpiralLines_reverseMap[inst]
			}
		}

	case *SpiralLineGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralOrigin:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SpiralRhombus:
		switch reverseField.GongstructName {
		// insertion point
		case "SpiralRhombusGrid":
			switch reverseField.Fieldname {
			case "SpiralRhombuses":
				res = stage.SpiralRhombusGrid_SpiralRhombuses_reverseMap[inst]
			}
		}

	case *SpiralRhombusGrid:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *VerticalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
