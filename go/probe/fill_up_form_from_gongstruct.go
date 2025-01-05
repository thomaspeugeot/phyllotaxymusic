// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
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
	case *models.Axis:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Axis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AxisFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AxisGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AxisGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AxisGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bezier:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bezier Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BezierGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BezierGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BezierGridStack:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BezierGridStack Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierGridStackFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Circle:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Circle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CircleGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "CircleGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ExportToMusicxml:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ExportToMusicxml Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExportToMusicxmlFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FrontCurve:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FrontCurve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrontCurveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FrontCurveStack:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FrontCurveStack Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrontCurveStackFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
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
	case *models.Key:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyFormCallback(
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
	case *models.ShapeCategory:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ShapeCategory Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShapeCategoryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralBezier:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralBezier Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralBezierFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralBezierGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralBezierGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralBezierGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralCircle:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralCircle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralCircleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralCircleGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralCircleGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralCircleGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralLine:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralLine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralLineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralLineGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralLineGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralLineGridFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralOrigin:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralOrigin Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralOriginFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralRhombus:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralRhombus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralRhombusFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SpiralRhombusGrid:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SpiralRhombusGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralRhombusGridFormCallback(
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
