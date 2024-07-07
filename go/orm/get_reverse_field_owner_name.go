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
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			case "Circles":
				if _circlegrid, ok := stage.CircleGrid_Circles_reverseMap[inst]; ok {
					res = _circlegrid.Name
				}
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.CircleGrid:
		tmp := GetInstanceDBFromInstance[models.CircleGrid, CircleGridDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.HorizontalAxis:
		tmp := GetInstanceDBFromInstance[models.HorizontalAxis, HorizontalAxisDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.InitialAxis:
		tmp := GetInstanceDBFromInstance[models.InitialAxis, InitialAxisDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.Parameter:
		tmp := GetInstanceDBFromInstance[models.Parameter, ParameterDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.Rhombus:
		tmp := GetInstanceDBFromInstance[models.Rhombus, RhombusDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			case "Rhombuses":
				if _rhombusgrid, ok := stage.RhombusGrid_Rhombuses_reverseMap[inst]; ok {
					res = _rhombusgrid.Name
				}
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.RhombusGrid:
		tmp := GetInstanceDBFromInstance[models.RhombusGrid, RhombusGridDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.VerticalAxis:
		tmp := GetInstanceDBFromInstance[models.VerticalAxis, VerticalAxisDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
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
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			case "Circles":
				res = stage.CircleGrid_Circles_reverseMap[inst]
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.CircleGrid:
		tmp := GetInstanceDBFromInstance[models.CircleGrid, CircleGridDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.HorizontalAxis:
		tmp := GetInstanceDBFromInstance[models.HorizontalAxis, HorizontalAxisDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.InitialAxis:
		tmp := GetInstanceDBFromInstance[models.InitialAxis, InitialAxisDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.Parameter:
		tmp := GetInstanceDBFromInstance[models.Parameter, ParameterDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.Rhombus:
		tmp := GetInstanceDBFromInstance[models.Rhombus, RhombusDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			case "Rhombuses":
				res = stage.RhombusGrid_Rhombuses_reverseMap[inst]
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.RhombusGrid:
		tmp := GetInstanceDBFromInstance[models.RhombusGrid, RhombusGridDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	case *models.VerticalAxis:
		tmp := GetInstanceDBFromInstance[models.VerticalAxis, VerticalAxisDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "CircleGrid":
			switch reverseField.Fieldname {
			}
		case "HorizontalAxis":
			switch reverseField.Fieldname {
			}
		case "InitialAxis":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Parameter":
			switch reverseField.Fieldname {
			}
		case "Rhombus":
			switch reverseField.Fieldname {
			}
		case "RhombusGrid":
			switch reverseField.Fieldname {
			}
		case "VerticalAxis":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
