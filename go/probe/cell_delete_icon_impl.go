// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Axis:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.AxisGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Bezier:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.BezierGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.BezierGridStack:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Circle:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CircleGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FrontCurve:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FrontCurveStack:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.HorizontalAxis:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Key:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.MovingLine:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.NoteInfo:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Parameter:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Rhombus:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.RhombusGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ShapeCategory:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralBezier:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralBezierGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralCircle:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralCircleGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralLine:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralLineGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralOrigin:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralRhombus:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SpiralRhombusGrid:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.VerticalAxis:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

