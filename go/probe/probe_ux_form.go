// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AxisFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Axis", true)
			} else {
				FillUpFormFromGongstruct(onSave.axis, probe)
			}
		case *AxisGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AxisGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.axisgrid, probe)
			}
		case *BezierFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Bezier", true)
			} else {
				FillUpFormFromGongstruct(onSave.bezier, probe)
			}
		case *BezierGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "BezierGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.beziergrid, probe)
			}
		case *BezierGridStackFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "BezierGridStack", true)
			} else {
				FillUpFormFromGongstruct(onSave.beziergridstack, probe)
			}
		case *ChapterFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Chapter", true)
			} else {
				FillUpFormFromGongstruct(onSave.chapter, probe)
			}
		case *CircleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Circle", true)
			} else {
				FillUpFormFromGongstruct(onSave.circle, probe)
			}
		case *CircleGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CircleGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.circlegrid, probe)
			}
		case *ContentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Content", true)
			} else {
				FillUpFormFromGongstruct(onSave.content, probe)
			}
		case *ExportToMusicxmlFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ExportToMusicxml", true)
			} else {
				FillUpFormFromGongstruct(onSave.exporttomusicxml, probe)
			}
		case *FrontCurveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "FrontCurve", true)
			} else {
				FillUpFormFromGongstruct(onSave.frontcurve, probe)
			}
		case *FrontCurveStackFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "FrontCurveStack", true)
			} else {
				FillUpFormFromGongstruct(onSave.frontcurvestack, probe)
			}
		case *HorizontalAxisFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "HorizontalAxis", true)
			} else {
				FillUpFormFromGongstruct(onSave.horizontalaxis, probe)
			}
		case *KeyFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Key", true)
			} else {
				FillUpFormFromGongstruct(onSave.key, probe)
			}
		case *ParameterFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Parameter", true)
			} else {
				FillUpFormFromGongstruct(onSave.parameter, probe)
			}
		case *RhombusFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Rhombus", true)
			} else {
				FillUpFormFromGongstruct(onSave.rhombus, probe)
			}
		case *RhombusGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RhombusGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.rhombusgrid, probe)
			}
		case *ShapeCategoryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ShapeCategory", true)
			} else {
				FillUpFormFromGongstruct(onSave.shapecategory, probe)
			}
		case *SpiralBezierFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralBezier", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralbezier, probe)
			}
		case *SpiralBezierGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralBezierGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralbeziergrid, probe)
			}
		case *SpiralCircleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralCircle", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralcircle, probe)
			}
		case *SpiralCircleGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralCircleGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralcirclegrid, probe)
			}
		case *SpiralLineFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralLine", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralline, probe)
			}
		case *SpiralLineGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralLineGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.spirallinegrid, probe)
			}
		case *SpiralOriginFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralOrigin", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralorigin, probe)
			}
		case *SpiralRhombusFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralRhombus", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralrhombus, probe)
			}
		case *SpiralRhombusGridFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SpiralRhombusGrid", true)
			} else {
				FillUpFormFromGongstruct(onSave.spiralrhombusgrid, probe)
			}
		case *VerticalAxisFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "VerticalAxis", true)
			} else {
				FillUpFormFromGongstruct(onSave.verticalaxis, probe)
			}
		}
	}
}

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
	case "Axis":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Axis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AxisFormCallback(
			nil,
			probe,
			formGroup,
		)
		axis := new(models.Axis)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(axis, formGroup, probe)
	case "AxisGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AxisGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AxisGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		axisgrid := new(models.AxisGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(axisgrid, formGroup, probe)
	case "Bezier":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bezier Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierFormCallback(
			nil,
			probe,
			formGroup,
		)
		bezier := new(models.Bezier)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bezier, formGroup, probe)
	case "BezierGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BezierGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		beziergrid := new(models.BezierGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beziergrid, formGroup, probe)
	case "BezierGridStack":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BezierGridStack Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BezierGridStackFormCallback(
			nil,
			probe,
			formGroup,
		)
		beziergridstack := new(models.BezierGridStack)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beziergridstack, formGroup, probe)
	case "Chapter":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Chapter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChapterFormCallback(
			nil,
			probe,
			formGroup,
		)
		chapter := new(models.Chapter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(chapter, formGroup, probe)
	case "Circle":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Circle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleFormCallback(
			nil,
			probe,
			formGroup,
		)
		circle := new(models.Circle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(circle, formGroup, probe)
	case "CircleGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CircleGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		circlegrid := new(models.CircleGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(circlegrid, formGroup, probe)
	case "Content":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Content Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		content := new(models.Content)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(content, formGroup, probe)
	case "ExportToMusicxml":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ExportToMusicxml Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExportToMusicxmlFormCallback(
			nil,
			probe,
			formGroup,
		)
		exporttomusicxml := new(models.ExportToMusicxml)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(exporttomusicxml, formGroup, probe)
	case "FrontCurve":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FrontCurve Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrontCurveFormCallback(
			nil,
			probe,
			formGroup,
		)
		frontcurve := new(models.FrontCurve)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(frontcurve, formGroup, probe)
	case "FrontCurveStack":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FrontCurveStack Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrontCurveStackFormCallback(
			nil,
			probe,
			formGroup,
		)
		frontcurvestack := new(models.FrontCurveStack)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(frontcurvestack, formGroup, probe)
	case "HorizontalAxis":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "HorizontalAxis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HorizontalAxisFormCallback(
			nil,
			probe,
			formGroup,
		)
		horizontalaxis := new(models.HorizontalAxis)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(horizontalaxis, formGroup, probe)
	case "Key":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyFormCallback(
			nil,
			probe,
			formGroup,
		)
		key := new(models.Key)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key, formGroup, probe)
	case "Parameter":
		formGroup := (&form.FormGroup{
			Name:  FormName,
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
	case "Rhombus":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Rhombus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusFormCallback(
			nil,
			probe,
			formGroup,
		)
		rhombus := new(models.Rhombus)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rhombus, formGroup, probe)
	case "RhombusGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RhombusGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RhombusGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		rhombusgrid := new(models.RhombusGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rhombusgrid, formGroup, probe)
	case "ShapeCategory":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ShapeCategory Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ShapeCategoryFormCallback(
			nil,
			probe,
			formGroup,
		)
		shapecategory := new(models.ShapeCategory)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(shapecategory, formGroup, probe)
	case "SpiralBezier":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralBezier Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralBezierFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralbezier := new(models.SpiralBezier)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralbezier, formGroup, probe)
	case "SpiralBezierGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralBezierGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralBezierGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralbeziergrid := new(models.SpiralBezierGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralbeziergrid, formGroup, probe)
	case "SpiralCircle":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralCircle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralCircleFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralcircle := new(models.SpiralCircle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralcircle, formGroup, probe)
	case "SpiralCircleGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralCircleGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralCircleGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralcirclegrid := new(models.SpiralCircleGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralcirclegrid, formGroup, probe)
	case "SpiralLine":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralLine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralLineFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralline := new(models.SpiralLine)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralline, formGroup, probe)
	case "SpiralLineGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralLineGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralLineGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		spirallinegrid := new(models.SpiralLineGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spirallinegrid, formGroup, probe)
	case "SpiralOrigin":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralOrigin Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralOriginFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralorigin := new(models.SpiralOrigin)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralorigin, formGroup, probe)
	case "SpiralRhombus":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralRhombus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralRhombusFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralrhombus := new(models.SpiralRhombus)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralrhombus, formGroup, probe)
	case "SpiralRhombusGrid":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SpiralRhombusGrid Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SpiralRhombusGridFormCallback(
			nil,
			probe,
			formGroup,
		)
		spiralrhombusgrid := new(models.SpiralRhombusGrid)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spiralrhombusgrid, formGroup, probe)
	case "VerticalAxis":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "VerticalAxis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VerticalAxisFormCallback(
			nil,
			probe,
			formGroup,
		)
		verticalaxis := new(models.VerticalAxis)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(verticalaxis, formGroup, probe)
	}
	formStage.Commit()
}
