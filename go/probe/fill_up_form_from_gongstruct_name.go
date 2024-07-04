// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/phylotaxymusic/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Line":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LineFormCallback(
			nil,
			probe,
			formGroup,
		)
		line := new(models.Line)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(line, formGroup, probe)
	case "Parameter":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Parameter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterFormCallback(
			nil,
			probe,
			formGroup,
		)
		parameter := new(models.Parameter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parameter, formGroup, probe)
	}
	formStage.Commit()
}
