// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/cursor/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Cursor":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Cursor Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CursorFormCallback(
			nil,
			probe,
			formGroup,
		)
		cursor := new(models.Cursor)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cursor, formGroup, probe)
	}
	formStage.Commit()
}
