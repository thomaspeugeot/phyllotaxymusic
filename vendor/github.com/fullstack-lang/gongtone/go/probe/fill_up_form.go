// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtone/go/models"
	"github.com/fullstack-lang/gongtone/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Freqency:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Frequencies"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Frequencies",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note, *models.Freqency](
					nil,
					"Frequencies",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Frequencies", instanceWithInferedType, &instanceWithInferedType.Frequencies, formGroup, probe)
		BasicFieldtoForm("Start", instanceWithInferedType.Start, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Velocity", instanceWithInferedType.Velocity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Info", instanceWithInferedType.Info, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}