// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/buttons/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Button:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Button Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Layout:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
