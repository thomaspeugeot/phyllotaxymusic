// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
	"github.com/thomaspeugeot/phyllotaxymusic/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Axis:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("AngleDegree", instanceWithInferedType.AngleDegree, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Length", instanceWithInferedType.Length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CenterX", instanceWithInferedType.CenterX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CenterY", instanceWithInferedType.CenterY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AxisGrid"
			rf.Fieldname = "Axiss"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.AxisGrid),
					"Axiss",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.AxisGrid, *models.Axis](
					nil,
					"Axiss",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.AxisGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Reference", instanceWithInferedType.Reference, formGroup, probe)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("Axiss", instanceWithInferedType, &instanceWithInferedType.Axiss, formGroup, probe)

	case *models.Bezier:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointStartX", instanceWithInferedType.ControlPointStartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointStartY", instanceWithInferedType.ControlPointStartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointEndX", instanceWithInferedType.ControlPointEndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointEndY", instanceWithInferedType.ControlPointEndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BezierGrid"
			rf.Fieldname = "Beziers"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.BezierGrid),
					"Beziers",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.BezierGrid, *models.Bezier](
					nil,
					"Beziers",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.BezierGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Reference", instanceWithInferedType.Reference, formGroup, probe)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("Beziers", instanceWithInferedType, &instanceWithInferedType.Beziers, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BezierGridStack"
			rf.Fieldname = "BezierGrids"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.BezierGridStack),
					"BezierGrids",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.BezierGridStack, *models.BezierGrid](
					nil,
					"BezierGrids",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.BezierGridStack:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("BezierGrids", instanceWithInferedType, &instanceWithInferedType.BezierGrids, formGroup, probe)

	case *models.Circle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("CenterX", instanceWithInferedType.CenterX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CenterY", instanceWithInferedType.CenterY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasBespokeRadius", instanceWithInferedType.HasBespokeRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespopkeRadius", instanceWithInferedType.BespopkeRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		BasicFieldtoForm("Pitch", instanceWithInferedType.Pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowName", instanceWithInferedType.ShowName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BeatNb", instanceWithInferedType.BeatNb, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "CircleGrid"
			rf.Fieldname = "Circles"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.CircleGrid),
					"Circles",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.CircleGrid, *models.Circle](
					nil,
					"Circles",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.CircleGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Reference", instanceWithInferedType.Reference, formGroup, probe)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("Circles", instanceWithInferedType, &instanceWithInferedType.Circles, formGroup, probe)

	case *models.ExportToMusicxml:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Parameter", instanceWithInferedType.Parameter, formGroup, probe)

	case *models.FrontCurve:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FrontCurveStack"
			rf.Fieldname = "FrontCurves"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FrontCurveStack),
					"FrontCurves",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.FrontCurveStack, *models.FrontCurve](
					nil,
					"FrontCurves",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.FrontCurveStack:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("FrontCurves", instanceWithInferedType, &instanceWithInferedType.FrontCurves, formGroup, probe)
		AssociationSliceToForm("SpiralCircles", instanceWithInferedType, &instanceWithInferedType.SpiralCircles, formGroup, probe)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	case *models.HorizontalAxis:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("AxisHandleBorderLength", instanceWithInferedType.AxisHandleBorderLength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Axis_Length", instanceWithInferedType.Axis_Length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	case *models.Key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	case *models.Parameter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BackendColor", instanceWithInferedType.BackendColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinuteColor", instanceWithInferedType.MinuteColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HourColor", instanceWithInferedType.HourColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("N", instanceWithInferedType.N, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("M", instanceWithInferedType.M, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Z", instanceWithInferedType.Z, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InsideAngle", instanceWithInferedType.InsideAngle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShiftToNearestCircle", instanceWithInferedType.ShiftToNearestCircle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SideLength", instanceWithInferedType.SideLength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("InitialRhombus", instanceWithInferedType.InitialRhombus, formGroup, probe)
		AssociationFieldToForm("InitialCircle", instanceWithInferedType.InitialCircle, formGroup, probe)
		AssociationFieldToForm("InitialRhombusGrid", instanceWithInferedType.InitialRhombusGrid, formGroup, probe)
		AssociationFieldToForm("InitialCircleGrid", instanceWithInferedType.InitialCircleGrid, formGroup, probe)
		AssociationFieldToForm("InitialAxis", instanceWithInferedType.InitialAxis, formGroup, probe)
		AssociationFieldToForm("RotatedAxis", instanceWithInferedType.RotatedAxis, formGroup, probe)
		AssociationFieldToForm("RotatedRhombus", instanceWithInferedType.RotatedRhombus, formGroup, probe)
		AssociationFieldToForm("RotatedRhombusGrid", instanceWithInferedType.RotatedRhombusGrid, formGroup, probe)
		AssociationFieldToForm("RotatedCircleGrid", instanceWithInferedType.RotatedCircleGrid, formGroup, probe)
		AssociationFieldToForm("NextRhombus", instanceWithInferedType.NextRhombus, formGroup, probe)
		AssociationFieldToForm("NextCircle", instanceWithInferedType.NextCircle, formGroup, probe)
		AssociationFieldToForm("GrowingRhombusGridSeed", instanceWithInferedType.GrowingRhombusGridSeed, formGroup, probe)
		AssociationFieldToForm("GrowingRhombusGrid", instanceWithInferedType.GrowingRhombusGrid, formGroup, probe)
		AssociationFieldToForm("GrowingCircleGridSeed", instanceWithInferedType.GrowingCircleGridSeed, formGroup, probe)
		AssociationFieldToForm("GrowingCircleGrid", instanceWithInferedType.GrowingCircleGrid, formGroup, probe)
		AssociationFieldToForm("GrowingCircleGridLeftSeed", instanceWithInferedType.GrowingCircleGridLeftSeed, formGroup, probe)
		AssociationFieldToForm("GrowingCircleGridLeft", instanceWithInferedType.GrowingCircleGridLeft, formGroup, probe)
		AssociationFieldToForm("ConstructionAxis", instanceWithInferedType.ConstructionAxis, formGroup, probe)
		AssociationFieldToForm("ConstructionAxisGrid", instanceWithInferedType.ConstructionAxisGrid, formGroup, probe)
		AssociationFieldToForm("ConstructionCircle", instanceWithInferedType.ConstructionCircle, formGroup, probe)
		AssociationFieldToForm("ConstructionCircleGrid", instanceWithInferedType.ConstructionCircleGrid, formGroup, probe)
		AssociationFieldToForm("GrowthCurveSeed", instanceWithInferedType.GrowthCurveSeed, formGroup, probe)
		AssociationFieldToForm("GrowthCurve", instanceWithInferedType.GrowthCurve, formGroup, probe)
		AssociationFieldToForm("GrowthCurveShiftedRightSeed", instanceWithInferedType.GrowthCurveShiftedRightSeed, formGroup, probe)
		AssociationFieldToForm("GrowthCurveShiftedRight", instanceWithInferedType.GrowthCurveShiftedRight, formGroup, probe)
		AssociationFieldToForm("GrowthCurveNextSeed", instanceWithInferedType.GrowthCurveNextSeed, formGroup, probe)
		AssociationFieldToForm("GrowthCurveNext", instanceWithInferedType.GrowthCurveNext, formGroup, probe)
		AssociationFieldToForm("GrowthCurveNextShiftedRightSeed", instanceWithInferedType.GrowthCurveNextShiftedRightSeed, formGroup, probe)
		AssociationFieldToForm("GrowthCurveNextShiftedRight", instanceWithInferedType.GrowthCurveNextShiftedRight, formGroup, probe)
		AssociationFieldToForm("GrowthCurveStack", instanceWithInferedType.GrowthCurveStack, formGroup, probe)
		BasicFieldtoForm("StackWidth", instanceWithInferedType.StackWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NbShitRight", instanceWithInferedType.NbShitRight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackHeight", instanceWithInferedType.StackHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BezierControlLengthRatio", instanceWithInferedType.BezierControlLengthRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SpiralRhombusGridSeed", instanceWithInferedType.SpiralRhombusGridSeed, formGroup, probe)
		AssociationFieldToForm("SpiralRhombusGrid", instanceWithInferedType.SpiralRhombusGrid, formGroup, probe)
		AssociationFieldToForm("SpiralCircleSeed", instanceWithInferedType.SpiralCircleSeed, formGroup, probe)
		AssociationFieldToForm("SpiralCircleGrid", instanceWithInferedType.SpiralCircleGrid, formGroup, probe)
		AssociationFieldToForm("SpiralCircleFullGrid", instanceWithInferedType.SpiralCircleFullGrid, formGroup, probe)
		AssociationFieldToForm("SpiralConstructionOuterLineSeed", instanceWithInferedType.SpiralConstructionOuterLineSeed, formGroup, probe)
		AssociationFieldToForm("SpiralConstructionInnerLineSeed", instanceWithInferedType.SpiralConstructionInnerLineSeed, formGroup, probe)
		AssociationFieldToForm("SpiralConstructionOuterLineGrid", instanceWithInferedType.SpiralConstructionOuterLineGrid, formGroup, probe)
		AssociationFieldToForm("SpiralConstructionInnerLineGrid", instanceWithInferedType.SpiralConstructionInnerLineGrid, formGroup, probe)
		AssociationFieldToForm("SpiralConstructionCircleGrid", instanceWithInferedType.SpiralConstructionCircleGrid, formGroup, probe)
		AssociationFieldToForm("SpiralConstructionOuterLineFullGrid", instanceWithInferedType.SpiralConstructionOuterLineFullGrid, formGroup, probe)
		AssociationFieldToForm("SpiralBezierSeed", instanceWithInferedType.SpiralBezierSeed, formGroup, probe)
		AssociationFieldToForm("SpiralBezierGrid", instanceWithInferedType.SpiralBezierGrid, formGroup, probe)
		AssociationFieldToForm("SpiralBezierFullGrid", instanceWithInferedType.SpiralBezierFullGrid, formGroup, probe)
		BasicFieldtoForm("SpiralBezierStrength", instanceWithInferedType.SpiralBezierStrength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("FrontCurveStack", instanceWithInferedType.FrontCurveStack, formGroup, probe)
		BasicFieldtoForm("NbInterpolationPoints", instanceWithInferedType.NbInterpolationPoints, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Fkey", instanceWithInferedType.Fkey, formGroup, probe)
		BasicFieldtoForm("FkeySizeRatio", instanceWithInferedType.FkeySizeRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FkeyOriginRelativeX", instanceWithInferedType.FkeyOriginRelativeX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FkeyOriginRelativeY", instanceWithInferedType.FkeyOriginRelativeY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("PitchLines", instanceWithInferedType.PitchLines, formGroup, probe)
		BasicFieldtoForm("PitchHeight", instanceWithInferedType.PitchHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NbPitchLines", instanceWithInferedType.NbPitchLines, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("BeatLines", instanceWithInferedType.BeatLines, formGroup, probe)
		BasicFieldtoForm("BeatLinesHeightRatio", instanceWithInferedType.BeatLinesHeightRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NbBeatLines", instanceWithInferedType.NbBeatLines, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NbOfBeatsInTheme", instanceWithInferedType.NbOfBeatsInTheme, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("FirstVoice", instanceWithInferedType.FirstVoice, formGroup, probe)
		AssociationFieldToForm("FirstVoiceShiftedRigth", instanceWithInferedType.FirstVoiceShiftedRigth, formGroup, probe)
		BasicFieldtoForm("FirstVoiceShiftX", instanceWithInferedType.FirstVoiceShiftX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FirstVoiceShiftY", instanceWithInferedType.FirstVoiceShiftY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SecondVoice", instanceWithInferedType.SecondVoice, formGroup, probe)
		AssociationFieldToForm("SecondVoiceShiftedRight", instanceWithInferedType.SecondVoiceShiftedRight, formGroup, probe)
		BasicFieldtoForm("PitchDifference", instanceWithInferedType.PitchDifference, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BeatsPerSecond", instanceWithInferedType.BeatsPerSecond, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Level", instanceWithInferedType.Level, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("FirstVoiceNotes", instanceWithInferedType.FirstVoiceNotes, formGroup, probe)
		AssociationFieldToForm("FirstVoiceNotesShiftedRight", instanceWithInferedType.FirstVoiceNotesShiftedRight, formGroup, probe)
		AssociationFieldToForm("SecondVoiceNotes", instanceWithInferedType.SecondVoiceNotes, formGroup, probe)
		AssociationFieldToForm("SecondVoiceNotesShiftedRight", instanceWithInferedType.SecondVoiceNotesShiftedRight, formGroup, probe)
		BasicFieldtoForm("IsMinor", instanceWithInferedType.IsMinor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ThemeBinaryEncoding", instanceWithInferedType.ThemeBinaryEncoding, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OriginX", instanceWithInferedType.OriginX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OriginY", instanceWithInferedType.OriginY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("HorizontalAxis", instanceWithInferedType.HorizontalAxis, formGroup, probe)
		AssociationFieldToForm("VerticalAxis", instanceWithInferedType.VerticalAxis, formGroup, probe)
		AssociationFieldToForm("SpiralOrigin", instanceWithInferedType.SpiralOrigin, formGroup, probe)
		BasicFieldtoForm("SpiralOriginX", instanceWithInferedType.SpiralOriginX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SpiralOriginY", instanceWithInferedType.SpiralOriginY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OriginCrossWidth", instanceWithInferedType.OriginCrossWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SpiralRadiusRatio", instanceWithInferedType.SpiralRadiusRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowSpiralBezierConstruct", instanceWithInferedType.ShowSpiralBezierConstruct, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowInterpolationPoints", instanceWithInferedType.ShowInterpolationPoints, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ActualBeatsTemporalShift", instanceWithInferedType.ActualBeatsTemporalShift, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Rhombus:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("CenterX", instanceWithInferedType.CenterX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CenterY", instanceWithInferedType.CenterY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SideLength", instanceWithInferedType.SideLength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AngleDegree", instanceWithInferedType.AngleDegree, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InsideAngle", instanceWithInferedType.InsideAngle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RhombusGrid"
			rf.Fieldname = "Rhombuses"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RhombusGrid),
					"Rhombuses",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RhombusGrid, *models.Rhombus](
					nil,
					"Rhombuses",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.RhombusGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Reference", instanceWithInferedType.Reference, formGroup, probe)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("Rhombuses", instanceWithInferedType, &instanceWithInferedType.Rhombuses, formGroup, probe)

	case *models.ShapeCategory:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SpiralBezier:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointStartX", instanceWithInferedType.ControlPointStartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointStartY", instanceWithInferedType.ControlPointStartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointEndX", instanceWithInferedType.ControlPointEndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControlPointEndY", instanceWithInferedType.ControlPointEndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SpiralBezierGrid"
			rf.Fieldname = "SpiralBeziers"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SpiralBezierGrid),
					"SpiralBeziers",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SpiralBezierGrid, *models.SpiralBezier](
					nil,
					"SpiralBeziers",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SpiralBezierGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("SpiralBeziers", instanceWithInferedType, &instanceWithInferedType.SpiralBeziers, formGroup, probe)

	case *models.SpiralCircle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("CenterX", instanceWithInferedType.CenterX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CenterY", instanceWithInferedType.CenterY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasBespokeRadius", instanceWithInferedType.HasBespokeRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespopkeRadius", instanceWithInferedType.BespopkeRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		BasicFieldtoForm("Pitch", instanceWithInferedType.Pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowName", instanceWithInferedType.ShowName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BeatNb", instanceWithInferedType.BeatNb, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FrontCurveStack"
			rf.Fieldname = "SpiralCircles"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FrontCurveStack),
					"SpiralCircles",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.FrontCurveStack, *models.SpiralCircle](
					nil,
					"SpiralCircles",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SpiralCircleGrid"
			rf.Fieldname = "SpiralCircles"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SpiralCircleGrid),
					"SpiralCircles",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SpiralCircleGrid, *models.SpiralCircle](
					nil,
					"SpiralCircles",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SpiralCircleGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationFieldToForm("SpiralRhombusGrid", instanceWithInferedType.SpiralRhombusGrid, formGroup, probe)
		AssociationSliceToForm("SpiralCircles", instanceWithInferedType, &instanceWithInferedType.SpiralCircles, formGroup, probe)

	case *models.SpiralLine:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SpiralLineGrid"
			rf.Fieldname = "SpiralLines"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SpiralLineGrid),
					"SpiralLines",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SpiralLineGrid, *models.SpiralLine](
					nil,
					"SpiralLines",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SpiralLineGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("SpiralLines", instanceWithInferedType, &instanceWithInferedType.SpiralLines, formGroup, probe)

	case *models.SpiralOrigin:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	case *models.SpiralRhombus:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("X_r0", instanceWithInferedType.X_r0, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_r0", instanceWithInferedType.Y_r0, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_r1", instanceWithInferedType.X_r1, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_r1", instanceWithInferedType.Y_r1, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_r2", instanceWithInferedType.X_r2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_r2", instanceWithInferedType.Y_r2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_r3", instanceWithInferedType.X_r3, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_r3", instanceWithInferedType.Y_r3, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SpiralRhombusGrid"
			rf.Fieldname = "SpiralRhombuses"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SpiralRhombusGrid),
					"SpiralRhombuses",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SpiralRhombusGrid, *models.SpiralRhombus](
					nil,
					"SpiralRhombuses",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SpiralRhombusGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		AssociationSliceToForm("SpiralRhombuses", instanceWithInferedType, &instanceWithInferedType.SpiralRhombuses, formGroup, probe)

	case *models.VerticalAxis:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDisplayed", instanceWithInferedType.IsDisplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ShapeCategory", instanceWithInferedType.ShapeCategory, formGroup, probe)
		BasicFieldtoForm("AxisHandleBorderLength", instanceWithInferedType.AxisHandleBorderLength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Axis_Length", instanceWithInferedType.Axis_Length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	default:
		_ = instanceWithInferedType
	}
}
