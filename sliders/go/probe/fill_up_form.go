// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
	"github.com/thomaspeugeot/phyllotaxymusic/sliders/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Checkbox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Checkboxes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"Checkboxes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.Checkbox](
					nil,
					"Checkboxes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Sliders", instanceWithInferedType, &instanceWithInferedType.Sliders, formGroup, probe)
		AssociationSliceToForm("Checkboxes", instanceWithInferedType, &instanceWithInferedType.Checkboxes, formGroup, probe)

	case *models.Slider:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Sliders"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"Sliders",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.Slider](
					nil,
					"Sliders",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
