// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.HorizontalAxis:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "HorizontalAxis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HorizontalAxisFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.InitialAxis:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "InitialAxis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InitialAxisFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Line:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Parameter:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Parameter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rhombus:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Rhombus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RhombusGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RhombusGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.VerticalAxis:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "VerticalAxis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VerticalAxisFormCallback(
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
