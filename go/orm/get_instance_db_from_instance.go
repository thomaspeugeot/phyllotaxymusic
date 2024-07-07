// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | CircleDB | CircleGridDB | HorizontalAxisDB | InitialAxisDB | LineDB | ParameterDB | RhombusDB | RhombusGridDB | VerticalAxisDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Circle:
		circleInstance := any(concreteInstance).(*models.Circle)
		ret2 := backRepo.BackRepoCircle.GetCircleDBFromCirclePtr(circleInstance)
		ret = any(ret2).(*T2)
	case *models.CircleGrid:
		circlegridInstance := any(concreteInstance).(*models.CircleGrid)
		ret2 := backRepo.BackRepoCircleGrid.GetCircleGridDBFromCircleGridPtr(circlegridInstance)
		ret = any(ret2).(*T2)
	case *models.HorizontalAxis:
		horizontalaxisInstance := any(concreteInstance).(*models.HorizontalAxis)
		ret2 := backRepo.BackRepoHorizontalAxis.GetHorizontalAxisDBFromHorizontalAxisPtr(horizontalaxisInstance)
		ret = any(ret2).(*T2)
	case *models.InitialAxis:
		initialaxisInstance := any(concreteInstance).(*models.InitialAxis)
		ret2 := backRepo.BackRepoInitialAxis.GetInitialAxisDBFromInitialAxisPtr(initialaxisInstance)
		ret = any(ret2).(*T2)
	case *models.Line:
		lineInstance := any(concreteInstance).(*models.Line)
		ret2 := backRepo.BackRepoLine.GetLineDBFromLinePtr(lineInstance)
		ret = any(ret2).(*T2)
	case *models.Parameter:
		parameterInstance := any(concreteInstance).(*models.Parameter)
		ret2 := backRepo.BackRepoParameter.GetParameterDBFromParameterPtr(parameterInstance)
		ret = any(ret2).(*T2)
	case *models.Rhombus:
		rhombusInstance := any(concreteInstance).(*models.Rhombus)
		ret2 := backRepo.BackRepoRhombus.GetRhombusDBFromRhombusPtr(rhombusInstance)
		ret = any(ret2).(*T2)
	case *models.RhombusGrid:
		rhombusgridInstance := any(concreteInstance).(*models.RhombusGrid)
		ret2 := backRepo.BackRepoRhombusGrid.GetRhombusGridDBFromRhombusGridPtr(rhombusgridInstance)
		ret = any(ret2).(*T2)
	case *models.VerticalAxis:
		verticalaxisInstance := any(concreteInstance).(*models.VerticalAxis)
		ret2 := backRepo.BackRepoVerticalAxis.GetVerticalAxisDBFromVerticalAxisPtr(verticalaxisInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.CircleGrid:
		tmp := GetInstanceDBFromInstance[models.CircleGrid, CircleGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.HorizontalAxis:
		tmp := GetInstanceDBFromInstance[models.HorizontalAxis, HorizontalAxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.InitialAxis:
		tmp := GetInstanceDBFromInstance[models.InitialAxis, InitialAxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Parameter:
		tmp := GetInstanceDBFromInstance[models.Parameter, ParameterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rhombus:
		tmp := GetInstanceDBFromInstance[models.Rhombus, RhombusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RhombusGrid:
		tmp := GetInstanceDBFromInstance[models.RhombusGrid, RhombusGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VerticalAxis:
		tmp := GetInstanceDBFromInstance[models.VerticalAxis, VerticalAxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.CircleGrid:
		tmp := GetInstanceDBFromInstance[models.CircleGrid, CircleGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.HorizontalAxis:
		tmp := GetInstanceDBFromInstance[models.HorizontalAxis, HorizontalAxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.InitialAxis:
		tmp := GetInstanceDBFromInstance[models.InitialAxis, InitialAxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Parameter:
		tmp := GetInstanceDBFromInstance[models.Parameter, ParameterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rhombus:
		tmp := GetInstanceDBFromInstance[models.Rhombus, RhombusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RhombusGrid:
		tmp := GetInstanceDBFromInstance[models.RhombusGrid, RhombusGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VerticalAxis:
		tmp := GetInstanceDBFromInstance[models.VerticalAxis, VerticalAxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
