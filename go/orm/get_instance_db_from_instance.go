// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Axis:
		axisInstance := any(concreteInstance).(*models.Axis)
		ret2 := backRepo.BackRepoAxis.GetAxisDBFromAxisPtr(axisInstance)
		ret = any(ret2).(*T2)
	case *models.AxisGrid:
		axisgridInstance := any(concreteInstance).(*models.AxisGrid)
		ret2 := backRepo.BackRepoAxisGrid.GetAxisGridDBFromAxisGridPtr(axisgridInstance)
		ret = any(ret2).(*T2)
	case *models.Bezier:
		bezierInstance := any(concreteInstance).(*models.Bezier)
		ret2 := backRepo.BackRepoBezier.GetBezierDBFromBezierPtr(bezierInstance)
		ret = any(ret2).(*T2)
	case *models.BezierGrid:
		beziergridInstance := any(concreteInstance).(*models.BezierGrid)
		ret2 := backRepo.BackRepoBezierGrid.GetBezierGridDBFromBezierGridPtr(beziergridInstance)
		ret = any(ret2).(*T2)
	case *models.BezierGridStack:
		beziergridstackInstance := any(concreteInstance).(*models.BezierGridStack)
		ret2 := backRepo.BackRepoBezierGridStack.GetBezierGridStackDBFromBezierGridStackPtr(beziergridstackInstance)
		ret = any(ret2).(*T2)
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
	case *models.Key:
		keyInstance := any(concreteInstance).(*models.Key)
		ret2 := backRepo.BackRepoKey.GetKeyDBFromKeyPtr(keyInstance)
		ret = any(ret2).(*T2)
	case *models.NoteInfo:
		noteinfoInstance := any(concreteInstance).(*models.NoteInfo)
		ret2 := backRepo.BackRepoNoteInfo.GetNoteInfoDBFromNoteInfoPtr(noteinfoInstance)
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
	case *models.ShapeCategory:
		shapecategoryInstance := any(concreteInstance).(*models.ShapeCategory)
		ret2 := backRepo.BackRepoShapeCategory.GetShapeCategoryDBFromShapeCategoryPtr(shapecategoryInstance)
		ret = any(ret2).(*T2)
	case *models.SpiralRhombus:
		spiralrhombusInstance := any(concreteInstance).(*models.SpiralRhombus)
		ret2 := backRepo.BackRepoSpiralRhombus.GetSpiralRhombusDBFromSpiralRhombusPtr(spiralrhombusInstance)
		ret = any(ret2).(*T2)
	case *models.SpiralRhombusGrid:
		spiralrhombusgridInstance := any(concreteInstance).(*models.SpiralRhombusGrid)
		ret2 := backRepo.BackRepoSpiralRhombusGrid.GetSpiralRhombusGridDBFromSpiralRhombusGridPtr(spiralrhombusgridInstance)
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
	case *models.Axis:
		tmp := GetInstanceDBFromInstance[models.Axis, AxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AxisGrid:
		tmp := GetInstanceDBFromInstance[models.AxisGrid, AxisGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bezier:
		tmp := GetInstanceDBFromInstance[models.Bezier, BezierDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BezierGrid:
		tmp := GetInstanceDBFromInstance[models.BezierGrid, BezierGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BezierGridStack:
		tmp := GetInstanceDBFromInstance[models.BezierGridStack, BezierGridStackDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
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
	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.NoteInfo:
		tmp := GetInstanceDBFromInstance[models.NoteInfo, NoteInfoDB](
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
	case *models.ShapeCategory:
		tmp := GetInstanceDBFromInstance[models.ShapeCategory, ShapeCategoryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SpiralRhombus:
		tmp := GetInstanceDBFromInstance[models.SpiralRhombus, SpiralRhombusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SpiralRhombusGrid:
		tmp := GetInstanceDBFromInstance[models.SpiralRhombusGrid, SpiralRhombusGridDB](
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
	case *models.Axis:
		tmp := GetInstanceDBFromInstance[models.Axis, AxisDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AxisGrid:
		tmp := GetInstanceDBFromInstance[models.AxisGrid, AxisGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bezier:
		tmp := GetInstanceDBFromInstance[models.Bezier, BezierDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BezierGrid:
		tmp := GetInstanceDBFromInstance[models.BezierGrid, BezierGridDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BezierGridStack:
		tmp := GetInstanceDBFromInstance[models.BezierGridStack, BezierGridStackDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
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
	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.NoteInfo:
		tmp := GetInstanceDBFromInstance[models.NoteInfo, NoteInfoDB](
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
	case *models.ShapeCategory:
		tmp := GetInstanceDBFromInstance[models.ShapeCategory, ShapeCategoryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SpiralRhombus:
		tmp := GetInstanceDBFromInstance[models.SpiralRhombus, SpiralRhombusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SpiralRhombusGrid:
		tmp := GetInstanceDBFromInstance[models.SpiralRhombusGrid, SpiralRhombusGridDB](
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
