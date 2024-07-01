// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
	"github.com/thomaspeugeot/phylotaxymusic/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *table.FormGroup,
) (diagramFormCallback *DiagramFormCallback) {
	diagramFormCallback = new(DiagramFormCallback)
	diagramFormCallback.probe = probe
	diagramFormCallback.diagram = diagram
	diagramFormCallback.formGroup = formGroup

	diagramFormCallback.CreationMode = (diagram == nil)

	return
}

type DiagramFormCallback struct {
	diagram *models.Diagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {

	log.Println("DiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramFormCallback.probe.formStage.Checkout()

	if diagramFormCallback.diagram == nil {
		diagramFormCallback.diagram = new(models.Diagram).Stage(diagramFormCallback.probe.stageOfInterest)
	}
	diagram_ := diagramFormCallback.diagram
	_ = diagram_

	for _, formDiv := range diagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagram_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagram_.Unstage(diagramFormCallback.probe.stageOfInterest)
	}

	diagramFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Diagram](
		diagramFormCallback.probe,
	)
	diagramFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if diagramFormCallback.CreationMode || diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(diagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			diagramFormCallback.probe,
			newFormGroup,
		)
		diagram := new(models.Diagram)
		FillUpForm(diagram, newFormGroup, diagramFormCallback.probe)
		diagramFormCallback.probe.formStage.Commit()
	}

	fillUpTree(diagramFormCallback.probe)
}
