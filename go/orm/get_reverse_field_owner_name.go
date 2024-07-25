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

	case *models.HorizontalAxis:
		switch reverseField.GongstructName {
		// insertion point
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

	case *models.HorizontalAxis:
		switch reverseField.GongstructName {
		// insertion point
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

	case *models.VerticalAxis:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
