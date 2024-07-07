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
	case *models.Circle:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.HorizontalAxis:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.InitialAxis:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Line:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Parameter:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Rhombus:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.RhombusGrid:
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

