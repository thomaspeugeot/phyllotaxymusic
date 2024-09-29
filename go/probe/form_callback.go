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
func __gong__New__AxisFormCallback(
	axis *models.Axis,
	probe *Probe,
	formGroup *table.FormGroup,
) (axisFormCallback *AxisFormCallback) {
	axisFormCallback = new(AxisFormCallback)
	axisFormCallback.probe = probe
	axisFormCallback.axis = axis
	axisFormCallback.formGroup = formGroup

	axisFormCallback.CreationMode = (axis == nil)

	return
}

type AxisFormCallback struct {
	axis *models.Axis

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (axisFormCallback *AxisFormCallback) OnSave() {

	log.Println("AxisFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	axisFormCallback.probe.formStage.Checkout()

	if axisFormCallback.axis == nil {
		axisFormCallback.axis = new(models.Axis).Stage(axisFormCallback.probe.stageOfInterest)
	}
	axis_ := axisFormCallback.axis
	_ = axis_

	for _, formDiv := range axisFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(axis_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(axis_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(axis_.ShapeCategory), axisFormCallback.probe.stageOfInterest, formDiv)
		case "Angle":
			FormDivBasicFieldToField(&(axis_.Angle), formDiv)
		case "Length":
			FormDivBasicFieldToField(&(axis_.Length), formDiv)
		case "CenterX":
			FormDivBasicFieldToField(&(axis_.CenterX), formDiv)
		case "CenterY":
			FormDivBasicFieldToField(&(axis_.CenterY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(axis_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(axis_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(axis_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(axis_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(axis_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(axis_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(axis_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(axis_.Transform), formDiv)
		case "AxisGrid:Axiss":
			// we need to retrieve the field owner before the change
			var pastAxisGridOwner *models.AxisGrid
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AxisGrid"
			rf.Fieldname = "Axiss"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				axisFormCallback.probe.stageOfInterest,
				axisFormCallback.probe.backRepoOfInterest,
				axis_,
				&rf)

			if reverseFieldOwner != nil {
				pastAxisGridOwner = reverseFieldOwner.(*models.AxisGrid)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAxisGridOwner != nil {
					idx := slices.Index(pastAxisGridOwner.Axiss, axis_)
					pastAxisGridOwner.Axiss = slices.Delete(pastAxisGridOwner.Axiss, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _axisgrid := range *models.GetGongstructInstancesSet[models.AxisGrid](axisFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _axisgrid.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAxisGridOwner := _axisgrid // we have a match
						if pastAxisGridOwner != nil {
							if newAxisGridOwner != pastAxisGridOwner {
								idx := slices.Index(pastAxisGridOwner.Axiss, axis_)
								pastAxisGridOwner.Axiss = slices.Delete(pastAxisGridOwner.Axiss, idx, idx+1)
								newAxisGridOwner.Axiss = append(newAxisGridOwner.Axiss, axis_)
							}
						} else {
							newAxisGridOwner.Axiss = append(newAxisGridOwner.Axiss, axis_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if axisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axis_.Unstage(axisFormCallback.probe.stageOfInterest)
	}

	axisFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Axis](
		axisFormCallback.probe,
	)
	axisFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if axisFormCallback.CreationMode || axisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axisFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(axisFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AxisFormCallback(
			nil,
			axisFormCallback.probe,
			newFormGroup,
		)
		axis := new(models.Axis)
		FillUpForm(axis, newFormGroup, axisFormCallback.probe)
		axisFormCallback.probe.formStage.Commit()
	}

	fillUpTree(axisFormCallback.probe)
}
func __gong__New__AxisGridFormCallback(
	axisgrid *models.AxisGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (axisgridFormCallback *AxisGridFormCallback) {
	axisgridFormCallback = new(AxisGridFormCallback)
	axisgridFormCallback.probe = probe
	axisgridFormCallback.axisgrid = axisgrid
	axisgridFormCallback.formGroup = formGroup

	axisgridFormCallback.CreationMode = (axisgrid == nil)

	return
}

type AxisGridFormCallback struct {
	axisgrid *models.AxisGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (axisgridFormCallback *AxisGridFormCallback) OnSave() {

	log.Println("AxisGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	axisgridFormCallback.probe.formStage.Checkout()

	if axisgridFormCallback.axisgrid == nil {
		axisgridFormCallback.axisgrid = new(models.AxisGrid).Stage(axisgridFormCallback.probe.stageOfInterest)
	}
	axisgrid_ := axisgridFormCallback.axisgrid
	_ = axisgrid_

	for _, formDiv := range axisgridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(axisgrid_.Name), formDiv)
		case "Reference":
			FormDivSelectFieldToField(&(axisgrid_.Reference), axisgridFormCallback.probe.stageOfInterest, formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(axisgrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(axisgrid_.ShapeCategory), axisgridFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if axisgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axisgrid_.Unstage(axisgridFormCallback.probe.stageOfInterest)
	}

	axisgridFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AxisGrid](
		axisgridFormCallback.probe,
	)
	axisgridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if axisgridFormCallback.CreationMode || axisgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axisgridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(axisgridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AxisGridFormCallback(
			nil,
			axisgridFormCallback.probe,
			newFormGroup,
		)
		axisgrid := new(models.AxisGrid)
		FillUpForm(axisgrid, newFormGroup, axisgridFormCallback.probe)
		axisgridFormCallback.probe.formStage.Commit()
	}

	fillUpTree(axisgridFormCallback.probe)
}
func __gong__New__BezierFormCallback(
	bezier *models.Bezier,
	probe *Probe,
	formGroup *table.FormGroup,
) (bezierFormCallback *BezierFormCallback) {
	bezierFormCallback = new(BezierFormCallback)
	bezierFormCallback.probe = probe
	bezierFormCallback.bezier = bezier
	bezierFormCallback.formGroup = formGroup

	bezierFormCallback.CreationMode = (bezier == nil)

	return
}

type BezierFormCallback struct {
	bezier *models.Bezier

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bezierFormCallback *BezierFormCallback) OnSave() {

	log.Println("BezierFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bezierFormCallback.probe.formStage.Checkout()

	if bezierFormCallback.bezier == nil {
		bezierFormCallback.bezier = new(models.Bezier).Stage(bezierFormCallback.probe.stageOfInterest)
	}
	bezier_ := bezierFormCallback.bezier
	_ = bezier_

	for _, formDiv := range bezierFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bezier_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(bezier_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(bezier_.ShapeCategory), bezierFormCallback.probe.stageOfInterest, formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bezier_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bezier_.StartY), formDiv)
		case "ControlPointStartX":
			FormDivBasicFieldToField(&(bezier_.ControlPointStartX), formDiv)
		case "ControlPointStartY":
			FormDivBasicFieldToField(&(bezier_.ControlPointStartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bezier_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bezier_.EndY), formDiv)
		case "ControlPointEndX":
			FormDivBasicFieldToField(&(bezier_.ControlPointEndX), formDiv)
		case "ControlPointEndY":
			FormDivBasicFieldToField(&(bezier_.ControlPointEndY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(bezier_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(bezier_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(bezier_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(bezier_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(bezier_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(bezier_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(bezier_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(bezier_.Transform), formDiv)
		case "BezierGrid:Beziers":
			// we need to retrieve the field owner before the change
			var pastBezierGridOwner *models.BezierGrid
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BezierGrid"
			rf.Fieldname = "Beziers"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				bezierFormCallback.probe.stageOfInterest,
				bezierFormCallback.probe.backRepoOfInterest,
				bezier_,
				&rf)

			if reverseFieldOwner != nil {
				pastBezierGridOwner = reverseFieldOwner.(*models.BezierGrid)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBezierGridOwner != nil {
					idx := slices.Index(pastBezierGridOwner.Beziers, bezier_)
					pastBezierGridOwner.Beziers = slices.Delete(pastBezierGridOwner.Beziers, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _beziergrid := range *models.GetGongstructInstancesSet[models.BezierGrid](bezierFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _beziergrid.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBezierGridOwner := _beziergrid // we have a match
						if pastBezierGridOwner != nil {
							if newBezierGridOwner != pastBezierGridOwner {
								idx := slices.Index(pastBezierGridOwner.Beziers, bezier_)
								pastBezierGridOwner.Beziers = slices.Delete(pastBezierGridOwner.Beziers, idx, idx+1)
								newBezierGridOwner.Beziers = append(newBezierGridOwner.Beziers, bezier_)
							}
						} else {
							newBezierGridOwner.Beziers = append(newBezierGridOwner.Beziers, bezier_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bezierFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bezier_.Unstage(bezierFormCallback.probe.stageOfInterest)
	}

	bezierFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bezier](
		bezierFormCallback.probe,
	)
	bezierFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bezierFormCallback.CreationMode || bezierFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bezierFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bezierFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BezierFormCallback(
			nil,
			bezierFormCallback.probe,
			newFormGroup,
		)
		bezier := new(models.Bezier)
		FillUpForm(bezier, newFormGroup, bezierFormCallback.probe)
		bezierFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bezierFormCallback.probe)
}
func __gong__New__BezierGridFormCallback(
	beziergrid *models.BezierGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (beziergridFormCallback *BezierGridFormCallback) {
	beziergridFormCallback = new(BezierGridFormCallback)
	beziergridFormCallback.probe = probe
	beziergridFormCallback.beziergrid = beziergrid
	beziergridFormCallback.formGroup = formGroup

	beziergridFormCallback.CreationMode = (beziergrid == nil)

	return
}

type BezierGridFormCallback struct {
	beziergrid *models.BezierGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beziergridFormCallback *BezierGridFormCallback) OnSave() {

	log.Println("BezierGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beziergridFormCallback.probe.formStage.Checkout()

	if beziergridFormCallback.beziergrid == nil {
		beziergridFormCallback.beziergrid = new(models.BezierGrid).Stage(beziergridFormCallback.probe.stageOfInterest)
	}
	beziergrid_ := beziergridFormCallback.beziergrid
	_ = beziergrid_

	for _, formDiv := range beziergridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beziergrid_.Name), formDiv)
		case "Reference":
			FormDivSelectFieldToField(&(beziergrid_.Reference), beziergridFormCallback.probe.stageOfInterest, formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(beziergrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(beziergrid_.ShapeCategory), beziergridFormCallback.probe.stageOfInterest, formDiv)
		case "BezierGridStack:BezierGrids":
			// we need to retrieve the field owner before the change
			var pastBezierGridStackOwner *models.BezierGridStack
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BezierGridStack"
			rf.Fieldname = "BezierGrids"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				beziergridFormCallback.probe.stageOfInterest,
				beziergridFormCallback.probe.backRepoOfInterest,
				beziergrid_,
				&rf)

			if reverseFieldOwner != nil {
				pastBezierGridStackOwner = reverseFieldOwner.(*models.BezierGridStack)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBezierGridStackOwner != nil {
					idx := slices.Index(pastBezierGridStackOwner.BezierGrids, beziergrid_)
					pastBezierGridStackOwner.BezierGrids = slices.Delete(pastBezierGridStackOwner.BezierGrids, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _beziergridstack := range *models.GetGongstructInstancesSet[models.BezierGridStack](beziergridFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _beziergridstack.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBezierGridStackOwner := _beziergridstack // we have a match
						if pastBezierGridStackOwner != nil {
							if newBezierGridStackOwner != pastBezierGridStackOwner {
								idx := slices.Index(pastBezierGridStackOwner.BezierGrids, beziergrid_)
								pastBezierGridStackOwner.BezierGrids = slices.Delete(pastBezierGridStackOwner.BezierGrids, idx, idx+1)
								newBezierGridStackOwner.BezierGrids = append(newBezierGridStackOwner.BezierGrids, beziergrid_)
							}
						} else {
							newBezierGridStackOwner.BezierGrids = append(newBezierGridStackOwner.BezierGrids, beziergrid_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if beziergridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziergrid_.Unstage(beziergridFormCallback.probe.stageOfInterest)
	}

	beziergridFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BezierGrid](
		beziergridFormCallback.probe,
	)
	beziergridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beziergridFormCallback.CreationMode || beziergridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziergridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beziergridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BezierGridFormCallback(
			nil,
			beziergridFormCallback.probe,
			newFormGroup,
		)
		beziergrid := new(models.BezierGrid)
		FillUpForm(beziergrid, newFormGroup, beziergridFormCallback.probe)
		beziergridFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beziergridFormCallback.probe)
}
func __gong__New__BezierGridStackFormCallback(
	beziergridstack *models.BezierGridStack,
	probe *Probe,
	formGroup *table.FormGroup,
) (beziergridstackFormCallback *BezierGridStackFormCallback) {
	beziergridstackFormCallback = new(BezierGridStackFormCallback)
	beziergridstackFormCallback.probe = probe
	beziergridstackFormCallback.beziergridstack = beziergridstack
	beziergridstackFormCallback.formGroup = formGroup

	beziergridstackFormCallback.CreationMode = (beziergridstack == nil)

	return
}

type BezierGridStackFormCallback struct {
	beziergridstack *models.BezierGridStack

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beziergridstackFormCallback *BezierGridStackFormCallback) OnSave() {

	log.Println("BezierGridStackFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beziergridstackFormCallback.probe.formStage.Checkout()

	if beziergridstackFormCallback.beziergridstack == nil {
		beziergridstackFormCallback.beziergridstack = new(models.BezierGridStack).Stage(beziergridstackFormCallback.probe.stageOfInterest)
	}
	beziergridstack_ := beziergridstackFormCallback.beziergridstack
	_ = beziergridstack_

	for _, formDiv := range beziergridstackFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beziergridstack_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(beziergridstack_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(beziergridstack_.ShapeCategory), beziergridstackFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if beziergridstackFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziergridstack_.Unstage(beziergridstackFormCallback.probe.stageOfInterest)
	}

	beziergridstackFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BezierGridStack](
		beziergridstackFormCallback.probe,
	)
	beziergridstackFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beziergridstackFormCallback.CreationMode || beziergridstackFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziergridstackFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beziergridstackFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BezierGridStackFormCallback(
			nil,
			beziergridstackFormCallback.probe,
			newFormGroup,
		)
		beziergridstack := new(models.BezierGridStack)
		FillUpForm(beziergridstack, newFormGroup, beziergridstackFormCallback.probe)
		beziergridstackFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beziergridstackFormCallback.probe)
}
func __gong__New__CircleFormCallback(
	circle *models.Circle,
	probe *Probe,
	formGroup *table.FormGroup,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.probe = probe
	circleFormCallback.circle = circle
	circleFormCallback.formGroup = formGroup

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (circleFormCallback *CircleFormCallback) OnSave() {

	log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.probe.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.probe.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	for _, formDiv := range circleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(circle_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(circle_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(circle_.ShapeCategory), circleFormCallback.probe.stageOfInterest, formDiv)
		case "CenterX":
			FormDivBasicFieldToField(&(circle_.CenterX), formDiv)
		case "CenterY":
			FormDivBasicFieldToField(&(circle_.CenterY), formDiv)
		case "HasBespokeRadius":
			FormDivBasicFieldToField(&(circle_.HasBespokeRadius), formDiv)
		case "BespopkeRadius":
			FormDivBasicFieldToField(&(circle_.BespopkeRadius), formDiv)
		case "Pitch":
			FormDivBasicFieldToField(&(circle_.Pitch), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(circle_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(circle_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(circle_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(circle_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(circle_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(circle_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(circle_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(circle_.Transform), formDiv)
		case "CircleGrid:Circles":
			// we need to retrieve the field owner before the change
			var pastCircleGridOwner *models.CircleGrid
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "CircleGrid"
			rf.Fieldname = "Circles"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				circleFormCallback.probe.stageOfInterest,
				circleFormCallback.probe.backRepoOfInterest,
				circle_,
				&rf)

			if reverseFieldOwner != nil {
				pastCircleGridOwner = reverseFieldOwner.(*models.CircleGrid)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCircleGridOwner != nil {
					idx := slices.Index(pastCircleGridOwner.Circles, circle_)
					pastCircleGridOwner.Circles = slices.Delete(pastCircleGridOwner.Circles, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _circlegrid := range *models.GetGongstructInstancesSet[models.CircleGrid](circleFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _circlegrid.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCircleGridOwner := _circlegrid // we have a match
						if pastCircleGridOwner != nil {
							if newCircleGridOwner != pastCircleGridOwner {
								idx := slices.Index(pastCircleGridOwner.Circles, circle_)
								pastCircleGridOwner.Circles = slices.Delete(pastCircleGridOwner.Circles, idx, idx+1)
								newCircleGridOwner.Circles = append(newCircleGridOwner.Circles, circle_)
							}
						} else {
							newCircleGridOwner.Circles = append(newCircleGridOwner.Circles, circle_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circle_.Unstage(circleFormCallback.probe.stageOfInterest)
	}

	circleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Circle](
		circleFormCallback.probe,
	)
	circleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode || circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(circleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CircleFormCallback(
			nil,
			circleFormCallback.probe,
			newFormGroup,
		)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.probe)
		circleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(circleFormCallback.probe)
}
func __gong__New__CircleGridFormCallback(
	circlegrid *models.CircleGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (circlegridFormCallback *CircleGridFormCallback) {
	circlegridFormCallback = new(CircleGridFormCallback)
	circlegridFormCallback.probe = probe
	circlegridFormCallback.circlegrid = circlegrid
	circlegridFormCallback.formGroup = formGroup

	circlegridFormCallback.CreationMode = (circlegrid == nil)

	return
}

type CircleGridFormCallback struct {
	circlegrid *models.CircleGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (circlegridFormCallback *CircleGridFormCallback) OnSave() {

	log.Println("CircleGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circlegridFormCallback.probe.formStage.Checkout()

	if circlegridFormCallback.circlegrid == nil {
		circlegridFormCallback.circlegrid = new(models.CircleGrid).Stage(circlegridFormCallback.probe.stageOfInterest)
	}
	circlegrid_ := circlegridFormCallback.circlegrid
	_ = circlegrid_

	for _, formDiv := range circlegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(circlegrid_.Name), formDiv)
		case "Reference":
			FormDivSelectFieldToField(&(circlegrid_.Reference), circlegridFormCallback.probe.stageOfInterest, formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(circlegrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(circlegrid_.ShapeCategory), circlegridFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if circlegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circlegrid_.Unstage(circlegridFormCallback.probe.stageOfInterest)
	}

	circlegridFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CircleGrid](
		circlegridFormCallback.probe,
	)
	circlegridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if circlegridFormCallback.CreationMode || circlegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circlegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(circlegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CircleGridFormCallback(
			nil,
			circlegridFormCallback.probe,
			newFormGroup,
		)
		circlegrid := new(models.CircleGrid)
		FillUpForm(circlegrid, newFormGroup, circlegridFormCallback.probe)
		circlegridFormCallback.probe.formStage.Commit()
	}

	fillUpTree(circlegridFormCallback.probe)
}
func __gong__New__HorizontalAxisFormCallback(
	horizontalaxis *models.HorizontalAxis,
	probe *Probe,
	formGroup *table.FormGroup,
) (horizontalaxisFormCallback *HorizontalAxisFormCallback) {
	horizontalaxisFormCallback = new(HorizontalAxisFormCallback)
	horizontalaxisFormCallback.probe = probe
	horizontalaxisFormCallback.horizontalaxis = horizontalaxis
	horizontalaxisFormCallback.formGroup = formGroup

	horizontalaxisFormCallback.CreationMode = (horizontalaxis == nil)

	return
}

type HorizontalAxisFormCallback struct {
	horizontalaxis *models.HorizontalAxis

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (horizontalaxisFormCallback *HorizontalAxisFormCallback) OnSave() {

	log.Println("HorizontalAxisFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	horizontalaxisFormCallback.probe.formStage.Checkout()

	if horizontalaxisFormCallback.horizontalaxis == nil {
		horizontalaxisFormCallback.horizontalaxis = new(models.HorizontalAxis).Stage(horizontalaxisFormCallback.probe.stageOfInterest)
	}
	horizontalaxis_ := horizontalaxisFormCallback.horizontalaxis
	_ = horizontalaxis_

	for _, formDiv := range horizontalaxisFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(horizontalaxis_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(horizontalaxis_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(horizontalaxis_.ShapeCategory), horizontalaxisFormCallback.probe.stageOfInterest, formDiv)
		case "AxisHandleBorderLength":
			FormDivBasicFieldToField(&(horizontalaxis_.AxisHandleBorderLength), formDiv)
		case "Axis_Length":
			FormDivBasicFieldToField(&(horizontalaxis_.Axis_Length), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(horizontalaxis_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(horizontalaxis_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(horizontalaxis_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(horizontalaxis_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(horizontalaxis_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(horizontalaxis_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(horizontalaxis_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(horizontalaxis_.Transform), formDiv)
		}
	}

	// manage the suppress operation
	if horizontalaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		horizontalaxis_.Unstage(horizontalaxisFormCallback.probe.stageOfInterest)
	}

	horizontalaxisFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.HorizontalAxis](
		horizontalaxisFormCallback.probe,
	)
	horizontalaxisFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if horizontalaxisFormCallback.CreationMode || horizontalaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		horizontalaxisFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(horizontalaxisFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__HorizontalAxisFormCallback(
			nil,
			horizontalaxisFormCallback.probe,
			newFormGroup,
		)
		horizontalaxis := new(models.HorizontalAxis)
		FillUpForm(horizontalaxis, newFormGroup, horizontalaxisFormCallback.probe)
		horizontalaxisFormCallback.probe.formStage.Commit()
	}

	fillUpTree(horizontalaxisFormCallback.probe)
}
func __gong__New__KeyFormCallback(
	key *models.Key,
	probe *Probe,
	formGroup *table.FormGroup,
) (keyFormCallback *KeyFormCallback) {
	keyFormCallback = new(KeyFormCallback)
	keyFormCallback.probe = probe
	keyFormCallback.key = key
	keyFormCallback.formGroup = formGroup

	keyFormCallback.CreationMode = (key == nil)

	return
}

type KeyFormCallback struct {
	key *models.Key

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (keyFormCallback *KeyFormCallback) OnSave() {

	log.Println("KeyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	keyFormCallback.probe.formStage.Checkout()

	if keyFormCallback.key == nil {
		keyFormCallback.key = new(models.Key).Stage(keyFormCallback.probe.stageOfInterest)
	}
	key_ := keyFormCallback.key
	_ = key_

	for _, formDiv := range keyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(key_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(key_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(key_.ShapeCategory), keyFormCallback.probe.stageOfInterest, formDiv)
		case "Path":
			FormDivBasicFieldToField(&(key_.Path), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(key_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(key_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(key_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(key_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(key_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(key_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(key_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(key_.Transform), formDiv)
		}
	}

	// manage the suppress operation
	if keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		key_.Unstage(keyFormCallback.probe.stageOfInterest)
	}

	keyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Key](
		keyFormCallback.probe,
	)
	keyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if keyFormCallback.CreationMode || keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		keyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(keyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KeyFormCallback(
			nil,
			keyFormCallback.probe,
			newFormGroup,
		)
		key := new(models.Key)
		FillUpForm(key, newFormGroup, keyFormCallback.probe)
		keyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(keyFormCallback.probe)
}
func __gong__New__NoteInfoFormCallback(
	noteinfo *models.NoteInfo,
	probe *Probe,
	formGroup *table.FormGroup,
) (noteinfoFormCallback *NoteInfoFormCallback) {
	noteinfoFormCallback = new(NoteInfoFormCallback)
	noteinfoFormCallback.probe = probe
	noteinfoFormCallback.noteinfo = noteinfo
	noteinfoFormCallback.formGroup = formGroup

	noteinfoFormCallback.CreationMode = (noteinfo == nil)

	return
}

type NoteInfoFormCallback struct {
	noteinfo *models.NoteInfo

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (noteinfoFormCallback *NoteInfoFormCallback) OnSave() {

	log.Println("NoteInfoFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteinfoFormCallback.probe.formStage.Checkout()

	if noteinfoFormCallback.noteinfo == nil {
		noteinfoFormCallback.noteinfo = new(models.NoteInfo).Stage(noteinfoFormCallback.probe.stageOfInterest)
	}
	noteinfo_ := noteinfoFormCallback.noteinfo
	_ = noteinfo_

	for _, formDiv := range noteinfoFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteinfo_.Name), formDiv)
		case "IsKept":
			FormDivBasicFieldToField(&(noteinfo_.IsKept), formDiv)
		case "Parameter:NoteInfos":
			// we need to retrieve the field owner before the change
			var pastParameterOwner *models.Parameter
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Parameter"
			rf.Fieldname = "NoteInfos"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				noteinfoFormCallback.probe.stageOfInterest,
				noteinfoFormCallback.probe.backRepoOfInterest,
				noteinfo_,
				&rf)

			if reverseFieldOwner != nil {
				pastParameterOwner = reverseFieldOwner.(*models.Parameter)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastParameterOwner != nil {
					idx := slices.Index(pastParameterOwner.NoteInfos, noteinfo_)
					pastParameterOwner.NoteInfos = slices.Delete(pastParameterOwner.NoteInfos, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _parameter := range *models.GetGongstructInstancesSet[models.Parameter](noteinfoFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _parameter.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newParameterOwner := _parameter // we have a match
						if pastParameterOwner != nil {
							if newParameterOwner != pastParameterOwner {
								idx := slices.Index(pastParameterOwner.NoteInfos, noteinfo_)
								pastParameterOwner.NoteInfos = slices.Delete(pastParameterOwner.NoteInfos, idx, idx+1)
								newParameterOwner.NoteInfos = append(newParameterOwner.NoteInfos, noteinfo_)
							}
						} else {
							newParameterOwner.NoteInfos = append(newParameterOwner.NoteInfos, noteinfo_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteinfoFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteinfo_.Unstage(noteinfoFormCallback.probe.stageOfInterest)
	}

	noteinfoFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.NoteInfo](
		noteinfoFormCallback.probe,
	)
	noteinfoFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteinfoFormCallback.CreationMode || noteinfoFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteinfoFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(noteinfoFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteInfoFormCallback(
			nil,
			noteinfoFormCallback.probe,
			newFormGroup,
		)
		noteinfo := new(models.NoteInfo)
		FillUpForm(noteinfo, newFormGroup, noteinfoFormCallback.probe)
		noteinfoFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteinfoFormCallback.probe)
}
func __gong__New__ParameterFormCallback(
	parameter *models.Parameter,
	probe *Probe,
	formGroup *table.FormGroup,
) (parameterFormCallback *ParameterFormCallback) {
	parameterFormCallback = new(ParameterFormCallback)
	parameterFormCallback.probe = probe
	parameterFormCallback.parameter = parameter
	parameterFormCallback.formGroup = formGroup

	parameterFormCallback.CreationMode = (parameter == nil)

	return
}

type ParameterFormCallback struct {
	parameter *models.Parameter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (parameterFormCallback *ParameterFormCallback) OnSave() {

	log.Println("ParameterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	parameterFormCallback.probe.formStage.Checkout()

	if parameterFormCallback.parameter == nil {
		parameterFormCallback.parameter = new(models.Parameter).Stage(parameterFormCallback.probe.stageOfInterest)
	}
	parameter_ := parameterFormCallback.parameter
	_ = parameter_

	for _, formDiv := range parameterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(parameter_.Name), formDiv)
		case "N":
			FormDivBasicFieldToField(&(parameter_.N), formDiv)
		case "M":
			FormDivBasicFieldToField(&(parameter_.M), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(parameter_.Z), formDiv)
		case "InsideAngle":
			FormDivBasicFieldToField(&(parameter_.InsideAngle), formDiv)
		case "SideLength":
			FormDivBasicFieldToField(&(parameter_.SideLength), formDiv)
		case "InitialRhombus":
			FormDivSelectFieldToField(&(parameter_.InitialRhombus), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "InitialCircle":
			FormDivSelectFieldToField(&(parameter_.InitialCircle), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "InitialRhombusGrid":
			FormDivSelectFieldToField(&(parameter_.InitialRhombusGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "InitialCircleGrid":
			FormDivSelectFieldToField(&(parameter_.InitialCircleGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "InitialAxis":
			FormDivSelectFieldToField(&(parameter_.InitialAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedAxis":
			FormDivSelectFieldToField(&(parameter_.RotatedAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedRhombus":
			FormDivSelectFieldToField(&(parameter_.RotatedRhombus), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedRhombusGrid":
			FormDivSelectFieldToField(&(parameter_.RotatedRhombusGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedCircleGrid":
			FormDivSelectFieldToField(&(parameter_.RotatedCircleGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "NextRhombus":
			FormDivSelectFieldToField(&(parameter_.NextRhombus), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "NextCircle":
			FormDivSelectFieldToField(&(parameter_.NextCircle), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowingRhombusGridSeed":
			FormDivSelectFieldToField(&(parameter_.GrowingRhombusGridSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowingRhombusGrid":
			FormDivSelectFieldToField(&(parameter_.GrowingRhombusGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowingCircleGridSeed":
			FormDivSelectFieldToField(&(parameter_.GrowingCircleGridSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowingCircleGrid":
			FormDivSelectFieldToField(&(parameter_.GrowingCircleGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowingCircleGridLeftSeed":
			FormDivSelectFieldToField(&(parameter_.GrowingCircleGridLeftSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowingCircleGridLeft":
			FormDivSelectFieldToField(&(parameter_.GrowingCircleGridLeft), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "ConstructionAxis":
			FormDivSelectFieldToField(&(parameter_.ConstructionAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "ConstructionAxisGrid":
			FormDivSelectFieldToField(&(parameter_.ConstructionAxisGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "ConstructionCircle":
			FormDivSelectFieldToField(&(parameter_.ConstructionCircle), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "ConstructionCircleGrid":
			FormDivSelectFieldToField(&(parameter_.ConstructionCircleGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveSegment":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveSegment), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurve":
			FormDivSelectFieldToField(&(parameter_.GrowthCurve), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveShiftedRightSeed":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveShiftedRightSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveShiftedRight":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveShiftedRight), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveNextSeed":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveNextSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveNext":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveNext), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveNextShiftedRightSeed":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveNextShiftedRightSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveNextShiftedRight":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveNextShiftedRight), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveStack":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveStack), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "StackWidth":
			FormDivBasicFieldToField(&(parameter_.StackWidth), formDiv)
		case "NbShitRight":
			FormDivBasicFieldToField(&(parameter_.NbShitRight), formDiv)
		case "StackHeight":
			FormDivBasicFieldToField(&(parameter_.StackHeight), formDiv)
		case "BezierControlLengthRatio":
			FormDivBasicFieldToField(&(parameter_.BezierControlLengthRatio), formDiv)
		case "SpiralRhombus":
			FormDivSelectFieldToField(&(parameter_.SpiralRhombus), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralRhombusGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralRhombusGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralCircleSeed":
			FormDivSelectFieldToField(&(parameter_.SpiralCircleSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralCircleGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralCircleGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralConstructionAxis":
			FormDivSelectFieldToField(&(parameter_.SpiralConstructionAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "Fkey":
			FormDivSelectFieldToField(&(parameter_.Fkey), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "FkeySizeRatio":
			FormDivBasicFieldToField(&(parameter_.FkeySizeRatio), formDiv)
		case "FkeyOriginRelativeX":
			FormDivBasicFieldToField(&(parameter_.FkeyOriginRelativeX), formDiv)
		case "FkeyOriginRelativeY":
			FormDivBasicFieldToField(&(parameter_.FkeyOriginRelativeY), formDiv)
		case "PitchLines":
			FormDivSelectFieldToField(&(parameter_.PitchLines), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "PitchHeight":
			FormDivBasicFieldToField(&(parameter_.PitchHeight), formDiv)
		case "NbPitchLines":
			FormDivBasicFieldToField(&(parameter_.NbPitchLines), formDiv)
		case "MeasureLines":
			FormDivSelectFieldToField(&(parameter_.MeasureLines), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "MeasureLinesHeightRatio":
			FormDivBasicFieldToField(&(parameter_.MeasureLinesHeightRatio), formDiv)
		case "NbMeasureLines":
			FormDivBasicFieldToField(&(parameter_.NbMeasureLines), formDiv)
		case "NbMeasureLinesPerCurve":
			FormDivBasicFieldToField(&(parameter_.NbMeasureLinesPerCurve), formDiv)
		case "FirstVoice":
			FormDivSelectFieldToField(&(parameter_.FirstVoice), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "FirstVoiceShiftRigth":
			FormDivSelectFieldToField(&(parameter_.FirstVoiceShiftRigth), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "FirstVoiceShiftX":
			FormDivBasicFieldToField(&(parameter_.FirstVoiceShiftX), formDiv)
		case "FirstVoiceShiftY":
			FormDivBasicFieldToField(&(parameter_.FirstVoiceShiftY), formDiv)
		case "SecondVoice":
			FormDivSelectFieldToField(&(parameter_.SecondVoice), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SecondVoiceShiftedRight":
			FormDivSelectFieldToField(&(parameter_.SecondVoiceShiftedRight), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "PitchDifference":
			FormDivBasicFieldToField(&(parameter_.PitchDifference), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(parameter_.Speed), formDiv)
		case "Level":
			FormDivBasicFieldToField(&(parameter_.Level), formDiv)
		case "FirstVoiceNotes":
			FormDivSelectFieldToField(&(parameter_.FirstVoiceNotes), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "FirstVoiceNotesShiftedRight":
			FormDivSelectFieldToField(&(parameter_.FirstVoiceNotesShiftedRight), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SecondVoiceNotes":
			FormDivSelectFieldToField(&(parameter_.SecondVoiceNotes), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SecondVoiceNotesShiftedRight":
			FormDivSelectFieldToField(&(parameter_.SecondVoiceNotesShiftedRight), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "IsMinor":
			FormDivBasicFieldToField(&(parameter_.IsMinor), formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&(parameter_.OriginX), formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&(parameter_.OriginY), formDiv)
		case "HorizontalAxis":
			FormDivSelectFieldToField(&(parameter_.HorizontalAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "VerticalAxis":
			FormDivSelectFieldToField(&(parameter_.VerticalAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralOriginX":
			FormDivBasicFieldToField(&(parameter_.SpiralOriginX), formDiv)
		case "SpiralOriginY":
			FormDivBasicFieldToField(&(parameter_.SpiralOriginY), formDiv)
		case "SpiralInitialRadius":
			FormDivBasicFieldToField(&(parameter_.SpiralInitialRadius), formDiv)
		}
	}

	// manage the suppress operation
	if parameterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parameter_.Unstage(parameterFormCallback.probe.stageOfInterest)
	}

	parameterFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Parameter](
		parameterFormCallback.probe,
	)
	parameterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if parameterFormCallback.CreationMode || parameterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parameterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(parameterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParameterFormCallback(
			nil,
			parameterFormCallback.probe,
			newFormGroup,
		)
		parameter := new(models.Parameter)
		FillUpForm(parameter, newFormGroup, parameterFormCallback.probe)
		parameterFormCallback.probe.formStage.Commit()
	}

	fillUpTree(parameterFormCallback.probe)
}
func __gong__New__RhombusFormCallback(
	rhombus *models.Rhombus,
	probe *Probe,
	formGroup *table.FormGroup,
) (rhombusFormCallback *RhombusFormCallback) {
	rhombusFormCallback = new(RhombusFormCallback)
	rhombusFormCallback.probe = probe
	rhombusFormCallback.rhombus = rhombus
	rhombusFormCallback.formGroup = formGroup

	rhombusFormCallback.CreationMode = (rhombus == nil)

	return
}

type RhombusFormCallback struct {
	rhombus *models.Rhombus

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rhombusFormCallback *RhombusFormCallback) OnSave() {

	log.Println("RhombusFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rhombusFormCallback.probe.formStage.Checkout()

	if rhombusFormCallback.rhombus == nil {
		rhombusFormCallback.rhombus = new(models.Rhombus).Stage(rhombusFormCallback.probe.stageOfInterest)
	}
	rhombus_ := rhombusFormCallback.rhombus
	_ = rhombus_

	for _, formDiv := range rhombusFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rhombus_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(rhombus_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(rhombus_.ShapeCategory), rhombusFormCallback.probe.stageOfInterest, formDiv)
		case "CenterX":
			FormDivBasicFieldToField(&(rhombus_.CenterX), formDiv)
		case "CenterY":
			FormDivBasicFieldToField(&(rhombus_.CenterY), formDiv)
		case "SideLength":
			FormDivBasicFieldToField(&(rhombus_.SideLength), formDiv)
		case "Angle":
			FormDivBasicFieldToField(&(rhombus_.Angle), formDiv)
		case "InsideAngle":
			FormDivBasicFieldToField(&(rhombus_.InsideAngle), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rhombus_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rhombus_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rhombus_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(rhombus_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rhombus_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rhombus_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rhombus_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rhombus_.Transform), formDiv)
		case "RhombusGrid:Rhombuses":
			// we need to retrieve the field owner before the change
			var pastRhombusGridOwner *models.RhombusGrid
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RhombusGrid"
			rf.Fieldname = "Rhombuses"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				rhombusFormCallback.probe.stageOfInterest,
				rhombusFormCallback.probe.backRepoOfInterest,
				rhombus_,
				&rf)

			if reverseFieldOwner != nil {
				pastRhombusGridOwner = reverseFieldOwner.(*models.RhombusGrid)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRhombusGridOwner != nil {
					idx := slices.Index(pastRhombusGridOwner.Rhombuses, rhombus_)
					pastRhombusGridOwner.Rhombuses = slices.Delete(pastRhombusGridOwner.Rhombuses, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _rhombusgrid := range *models.GetGongstructInstancesSet[models.RhombusGrid](rhombusFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _rhombusgrid.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRhombusGridOwner := _rhombusgrid // we have a match
						if pastRhombusGridOwner != nil {
							if newRhombusGridOwner != pastRhombusGridOwner {
								idx := slices.Index(pastRhombusGridOwner.Rhombuses, rhombus_)
								pastRhombusGridOwner.Rhombuses = slices.Delete(pastRhombusGridOwner.Rhombuses, idx, idx+1)
								newRhombusGridOwner.Rhombuses = append(newRhombusGridOwner.Rhombuses, rhombus_)
							}
						} else {
							newRhombusGridOwner.Rhombuses = append(newRhombusGridOwner.Rhombuses, rhombus_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombus_.Unstage(rhombusFormCallback.probe.stageOfInterest)
	}

	rhombusFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Rhombus](
		rhombusFormCallback.probe,
	)
	rhombusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rhombusFormCallback.CreationMode || rhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(rhombusFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RhombusFormCallback(
			nil,
			rhombusFormCallback.probe,
			newFormGroup,
		)
		rhombus := new(models.Rhombus)
		FillUpForm(rhombus, newFormGroup, rhombusFormCallback.probe)
		rhombusFormCallback.probe.formStage.Commit()
	}

	fillUpTree(rhombusFormCallback.probe)
}
func __gong__New__RhombusGridFormCallback(
	rhombusgrid *models.RhombusGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (rhombusgridFormCallback *RhombusGridFormCallback) {
	rhombusgridFormCallback = new(RhombusGridFormCallback)
	rhombusgridFormCallback.probe = probe
	rhombusgridFormCallback.rhombusgrid = rhombusgrid
	rhombusgridFormCallback.formGroup = formGroup

	rhombusgridFormCallback.CreationMode = (rhombusgrid == nil)

	return
}

type RhombusGridFormCallback struct {
	rhombusgrid *models.RhombusGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rhombusgridFormCallback *RhombusGridFormCallback) OnSave() {

	log.Println("RhombusGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rhombusgridFormCallback.probe.formStage.Checkout()

	if rhombusgridFormCallback.rhombusgrid == nil {
		rhombusgridFormCallback.rhombusgrid = new(models.RhombusGrid).Stage(rhombusgridFormCallback.probe.stageOfInterest)
	}
	rhombusgrid_ := rhombusgridFormCallback.rhombusgrid
	_ = rhombusgrid_

	for _, formDiv := range rhombusgridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rhombusgrid_.Name), formDiv)
		case "Reference":
			FormDivSelectFieldToField(&(rhombusgrid_.Reference), rhombusgridFormCallback.probe.stageOfInterest, formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(rhombusgrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(rhombusgrid_.ShapeCategory), rhombusgridFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if rhombusgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusgrid_.Unstage(rhombusgridFormCallback.probe.stageOfInterest)
	}

	rhombusgridFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RhombusGrid](
		rhombusgridFormCallback.probe,
	)
	rhombusgridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rhombusgridFormCallback.CreationMode || rhombusgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusgridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(rhombusgridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RhombusGridFormCallback(
			nil,
			rhombusgridFormCallback.probe,
			newFormGroup,
		)
		rhombusgrid := new(models.RhombusGrid)
		FillUpForm(rhombusgrid, newFormGroup, rhombusgridFormCallback.probe)
		rhombusgridFormCallback.probe.formStage.Commit()
	}

	fillUpTree(rhombusgridFormCallback.probe)
}
func __gong__New__ShapeCategoryFormCallback(
	shapecategory *models.ShapeCategory,
	probe *Probe,
	formGroup *table.FormGroup,
) (shapecategoryFormCallback *ShapeCategoryFormCallback) {
	shapecategoryFormCallback = new(ShapeCategoryFormCallback)
	shapecategoryFormCallback.probe = probe
	shapecategoryFormCallback.shapecategory = shapecategory
	shapecategoryFormCallback.formGroup = formGroup

	shapecategoryFormCallback.CreationMode = (shapecategory == nil)

	return
}

type ShapeCategoryFormCallback struct {
	shapecategory *models.ShapeCategory

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (shapecategoryFormCallback *ShapeCategoryFormCallback) OnSave() {

	log.Println("ShapeCategoryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shapecategoryFormCallback.probe.formStage.Checkout()

	if shapecategoryFormCallback.shapecategory == nil {
		shapecategoryFormCallback.shapecategory = new(models.ShapeCategory).Stage(shapecategoryFormCallback.probe.stageOfInterest)
	}
	shapecategory_ := shapecategoryFormCallback.shapecategory
	_ = shapecategory_

	for _, formDiv := range shapecategoryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shapecategory_.Name), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(shapecategory_.IsExpanded), formDiv)
		}
	}

	// manage the suppress operation
	if shapecategoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shapecategory_.Unstage(shapecategoryFormCallback.probe.stageOfInterest)
	}

	shapecategoryFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ShapeCategory](
		shapecategoryFormCallback.probe,
	)
	shapecategoryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if shapecategoryFormCallback.CreationMode || shapecategoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shapecategoryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(shapecategoryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShapeCategoryFormCallback(
			nil,
			shapecategoryFormCallback.probe,
			newFormGroup,
		)
		shapecategory := new(models.ShapeCategory)
		FillUpForm(shapecategory, newFormGroup, shapecategoryFormCallback.probe)
		shapecategoryFormCallback.probe.formStage.Commit()
	}

	fillUpTree(shapecategoryFormCallback.probe)
}
func __gong__New__SpiralAxisFormCallback(
	spiralaxis *models.SpiralAxis,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiralaxisFormCallback *SpiralAxisFormCallback) {
	spiralaxisFormCallback = new(SpiralAxisFormCallback)
	spiralaxisFormCallback.probe = probe
	spiralaxisFormCallback.spiralaxis = spiralaxis
	spiralaxisFormCallback.formGroup = formGroup

	spiralaxisFormCallback.CreationMode = (spiralaxis == nil)

	return
}

type SpiralAxisFormCallback struct {
	spiralaxis *models.SpiralAxis

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiralaxisFormCallback *SpiralAxisFormCallback) OnSave() {

	log.Println("SpiralAxisFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiralaxisFormCallback.probe.formStage.Checkout()

	if spiralaxisFormCallback.spiralaxis == nil {
		spiralaxisFormCallback.spiralaxis = new(models.SpiralAxis).Stage(spiralaxisFormCallback.probe.stageOfInterest)
	}
	spiralaxis_ := spiralaxisFormCallback.spiralaxis
	_ = spiralaxis_

	for _, formDiv := range spiralaxisFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralaxis_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralaxis_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralaxis_.ShapeCategory), spiralaxisFormCallback.probe.stageOfInterest, formDiv)
		case "Angle":
			FormDivBasicFieldToField(&(spiralaxis_.Angle), formDiv)
		case "Length":
			FormDivBasicFieldToField(&(spiralaxis_.Length), formDiv)
		case "CenterX":
			FormDivBasicFieldToField(&(spiralaxis_.CenterX), formDiv)
		case "CenterY":
			FormDivBasicFieldToField(&(spiralaxis_.CenterY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(spiralaxis_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(spiralaxis_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(spiralaxis_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(spiralaxis_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(spiralaxis_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(spiralaxis_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(spiralaxis_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(spiralaxis_.Transform), formDiv)
		}
	}

	// manage the suppress operation
	if spiralaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralaxis_.Unstage(spiralaxisFormCallback.probe.stageOfInterest)
	}

	spiralaxisFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SpiralAxis](
		spiralaxisFormCallback.probe,
	)
	spiralaxisFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralaxisFormCallback.CreationMode || spiralaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralaxisFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spiralaxisFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralAxisFormCallback(
			nil,
			spiralaxisFormCallback.probe,
			newFormGroup,
		)
		spiralaxis := new(models.SpiralAxis)
		FillUpForm(spiralaxis, newFormGroup, spiralaxisFormCallback.probe)
		spiralaxisFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spiralaxisFormCallback.probe)
}
func __gong__New__SpiralCircleFormCallback(
	spiralcircle *models.SpiralCircle,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiralcircleFormCallback *SpiralCircleFormCallback) {
	spiralcircleFormCallback = new(SpiralCircleFormCallback)
	spiralcircleFormCallback.probe = probe
	spiralcircleFormCallback.spiralcircle = spiralcircle
	spiralcircleFormCallback.formGroup = formGroup

	spiralcircleFormCallback.CreationMode = (spiralcircle == nil)

	return
}

type SpiralCircleFormCallback struct {
	spiralcircle *models.SpiralCircle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiralcircleFormCallback *SpiralCircleFormCallback) OnSave() {

	log.Println("SpiralCircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiralcircleFormCallback.probe.formStage.Checkout()

	if spiralcircleFormCallback.spiralcircle == nil {
		spiralcircleFormCallback.spiralcircle = new(models.SpiralCircle).Stage(spiralcircleFormCallback.probe.stageOfInterest)
	}
	spiralcircle_ := spiralcircleFormCallback.spiralcircle
	_ = spiralcircle_

	for _, formDiv := range spiralcircleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralcircle_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralcircle_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralcircle_.ShapeCategory), spiralcircleFormCallback.probe.stageOfInterest, formDiv)
		case "CenterX":
			FormDivBasicFieldToField(&(spiralcircle_.CenterX), formDiv)
		case "CenterY":
			FormDivBasicFieldToField(&(spiralcircle_.CenterY), formDiv)
		case "HasBespokeRadius":
			FormDivBasicFieldToField(&(spiralcircle_.HasBespokeRadius), formDiv)
		case "BespopkeRadius":
			FormDivBasicFieldToField(&(spiralcircle_.BespopkeRadius), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(spiralcircle_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(spiralcircle_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(spiralcircle_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(spiralcircle_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(spiralcircle_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(spiralcircle_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(spiralcircle_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(spiralcircle_.Transform), formDiv)
		}
	}

	// manage the suppress operation
	if spiralcircleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralcircle_.Unstage(spiralcircleFormCallback.probe.stageOfInterest)
	}

	spiralcircleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SpiralCircle](
		spiralcircleFormCallback.probe,
	)
	spiralcircleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralcircleFormCallback.CreationMode || spiralcircleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralcircleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spiralcircleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralCircleFormCallback(
			nil,
			spiralcircleFormCallback.probe,
			newFormGroup,
		)
		spiralcircle := new(models.SpiralCircle)
		FillUpForm(spiralcircle, newFormGroup, spiralcircleFormCallback.probe)
		spiralcircleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spiralcircleFormCallback.probe)
}
func __gong__New__SpiralCircleGridFormCallback(
	spiralcirclegrid *models.SpiralCircleGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiralcirclegridFormCallback *SpiralCircleGridFormCallback) {
	spiralcirclegridFormCallback = new(SpiralCircleGridFormCallback)
	spiralcirclegridFormCallback.probe = probe
	spiralcirclegridFormCallback.spiralcirclegrid = spiralcirclegrid
	spiralcirclegridFormCallback.formGroup = formGroup

	spiralcirclegridFormCallback.CreationMode = (spiralcirclegrid == nil)

	return
}

type SpiralCircleGridFormCallback struct {
	spiralcirclegrid *models.SpiralCircleGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiralcirclegridFormCallback *SpiralCircleGridFormCallback) OnSave() {

	log.Println("SpiralCircleGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiralcirclegridFormCallback.probe.formStage.Checkout()

	if spiralcirclegridFormCallback.spiralcirclegrid == nil {
		spiralcirclegridFormCallback.spiralcirclegrid = new(models.SpiralCircleGrid).Stage(spiralcirclegridFormCallback.probe.stageOfInterest)
	}
	spiralcirclegrid_ := spiralcirclegridFormCallback.spiralcirclegrid
	_ = spiralcirclegrid_

	for _, formDiv := range spiralcirclegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralcirclegrid_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralcirclegrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralcirclegrid_.ShapeCategory), spiralcirclegridFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralRhombusGrid":
			FormDivSelectFieldToField(&(spiralcirclegrid_.SpiralRhombusGrid), spiralcirclegridFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if spiralcirclegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralcirclegrid_.Unstage(spiralcirclegridFormCallback.probe.stageOfInterest)
	}

	spiralcirclegridFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SpiralCircleGrid](
		spiralcirclegridFormCallback.probe,
	)
	spiralcirclegridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralcirclegridFormCallback.CreationMode || spiralcirclegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralcirclegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spiralcirclegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralCircleGridFormCallback(
			nil,
			spiralcirclegridFormCallback.probe,
			newFormGroup,
		)
		spiralcirclegrid := new(models.SpiralCircleGrid)
		FillUpForm(spiralcirclegrid, newFormGroup, spiralcirclegridFormCallback.probe)
		spiralcirclegridFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spiralcirclegridFormCallback.probe)
}
func __gong__New__SpiralRhombusFormCallback(
	spiralrhombus *models.SpiralRhombus,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiralrhombusFormCallback *SpiralRhombusFormCallback) {
	spiralrhombusFormCallback = new(SpiralRhombusFormCallback)
	spiralrhombusFormCallback.probe = probe
	spiralrhombusFormCallback.spiralrhombus = spiralrhombus
	spiralrhombusFormCallback.formGroup = formGroup

	spiralrhombusFormCallback.CreationMode = (spiralrhombus == nil)

	return
}

type SpiralRhombusFormCallback struct {
	spiralrhombus *models.SpiralRhombus

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiralrhombusFormCallback *SpiralRhombusFormCallback) OnSave() {

	log.Println("SpiralRhombusFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiralrhombusFormCallback.probe.formStage.Checkout()

	if spiralrhombusFormCallback.spiralrhombus == nil {
		spiralrhombusFormCallback.spiralrhombus = new(models.SpiralRhombus).Stage(spiralrhombusFormCallback.probe.stageOfInterest)
	}
	spiralrhombus_ := spiralrhombusFormCallback.spiralrhombus
	_ = spiralrhombus_

	for _, formDiv := range spiralrhombusFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralrhombus_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralrhombus_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralrhombus_.ShapeCategory), spiralrhombusFormCallback.probe.stageOfInterest, formDiv)
		case "Rhombus":
			FormDivSelectFieldToField(&(spiralrhombus_.Rhombus), spiralrhombusFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if spiralrhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombus_.Unstage(spiralrhombusFormCallback.probe.stageOfInterest)
	}

	spiralrhombusFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SpiralRhombus](
		spiralrhombusFormCallback.probe,
	)
	spiralrhombusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralrhombusFormCallback.CreationMode || spiralrhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spiralrhombusFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralRhombusFormCallback(
			nil,
			spiralrhombusFormCallback.probe,
			newFormGroup,
		)
		spiralrhombus := new(models.SpiralRhombus)
		FillUpForm(spiralrhombus, newFormGroup, spiralrhombusFormCallback.probe)
		spiralrhombusFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spiralrhombusFormCallback.probe)
}
func __gong__New__SpiralRhombusGridFormCallback(
	spiralrhombusgrid *models.SpiralRhombusGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiralrhombusgridFormCallback *SpiralRhombusGridFormCallback) {
	spiralrhombusgridFormCallback = new(SpiralRhombusGridFormCallback)
	spiralrhombusgridFormCallback.probe = probe
	spiralrhombusgridFormCallback.spiralrhombusgrid = spiralrhombusgrid
	spiralrhombusgridFormCallback.formGroup = formGroup

	spiralrhombusgridFormCallback.CreationMode = (spiralrhombusgrid == nil)

	return
}

type SpiralRhombusGridFormCallback struct {
	spiralrhombusgrid *models.SpiralRhombusGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiralrhombusgridFormCallback *SpiralRhombusGridFormCallback) OnSave() {

	log.Println("SpiralRhombusGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiralrhombusgridFormCallback.probe.formStage.Checkout()

	if spiralrhombusgridFormCallback.spiralrhombusgrid == nil {
		spiralrhombusgridFormCallback.spiralrhombusgrid = new(models.SpiralRhombusGrid).Stage(spiralrhombusgridFormCallback.probe.stageOfInterest)
	}
	spiralrhombusgrid_ := spiralrhombusgridFormCallback.spiralrhombusgrid
	_ = spiralrhombusgrid_

	for _, formDiv := range spiralrhombusgridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralrhombusgrid_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralrhombusgrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralrhombusgrid_.ShapeCategory), spiralrhombusgridFormCallback.probe.stageOfInterest, formDiv)
		case "RhombusGrid":
			FormDivSelectFieldToField(&(spiralrhombusgrid_.RhombusGrid), spiralrhombusgridFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if spiralrhombusgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombusgrid_.Unstage(spiralrhombusgridFormCallback.probe.stageOfInterest)
	}

	spiralrhombusgridFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SpiralRhombusGrid](
		spiralrhombusgridFormCallback.probe,
	)
	spiralrhombusgridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralrhombusgridFormCallback.CreationMode || spiralrhombusgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombusgridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spiralrhombusgridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralRhombusGridFormCallback(
			nil,
			spiralrhombusgridFormCallback.probe,
			newFormGroup,
		)
		spiralrhombusgrid := new(models.SpiralRhombusGrid)
		FillUpForm(spiralrhombusgrid, newFormGroup, spiralrhombusgridFormCallback.probe)
		spiralrhombusgridFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spiralrhombusgridFormCallback.probe)
}
func __gong__New__VerticalAxisFormCallback(
	verticalaxis *models.VerticalAxis,
	probe *Probe,
	formGroup *table.FormGroup,
) (verticalaxisFormCallback *VerticalAxisFormCallback) {
	verticalaxisFormCallback = new(VerticalAxisFormCallback)
	verticalaxisFormCallback.probe = probe
	verticalaxisFormCallback.verticalaxis = verticalaxis
	verticalaxisFormCallback.formGroup = formGroup

	verticalaxisFormCallback.CreationMode = (verticalaxis == nil)

	return
}

type VerticalAxisFormCallback struct {
	verticalaxis *models.VerticalAxis

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (verticalaxisFormCallback *VerticalAxisFormCallback) OnSave() {

	log.Println("VerticalAxisFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	verticalaxisFormCallback.probe.formStage.Checkout()

	if verticalaxisFormCallback.verticalaxis == nil {
		verticalaxisFormCallback.verticalaxis = new(models.VerticalAxis).Stage(verticalaxisFormCallback.probe.stageOfInterest)
	}
	verticalaxis_ := verticalaxisFormCallback.verticalaxis
	_ = verticalaxis_

	for _, formDiv := range verticalaxisFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(verticalaxis_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(verticalaxis_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(verticalaxis_.ShapeCategory), verticalaxisFormCallback.probe.stageOfInterest, formDiv)
		case "AxisHandleBorderLength":
			FormDivBasicFieldToField(&(verticalaxis_.AxisHandleBorderLength), formDiv)
		case "Axis_Length":
			FormDivBasicFieldToField(&(verticalaxis_.Axis_Length), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(verticalaxis_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(verticalaxis_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(verticalaxis_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(verticalaxis_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(verticalaxis_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(verticalaxis_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(verticalaxis_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(verticalaxis_.Transform), formDiv)
		}
	}

	// manage the suppress operation
	if verticalaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		verticalaxis_.Unstage(verticalaxisFormCallback.probe.stageOfInterest)
	}

	verticalaxisFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.VerticalAxis](
		verticalaxisFormCallback.probe,
	)
	verticalaxisFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if verticalaxisFormCallback.CreationMode || verticalaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		verticalaxisFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(verticalaxisFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__VerticalAxisFormCallback(
			nil,
			verticalaxisFormCallback.probe,
			newFormGroup,
		)
		verticalaxis := new(models.VerticalAxis)
		FillUpForm(verticalaxis, newFormGroup, verticalaxisFormCallback.probe)
		verticalaxisFormCallback.probe.formStage.Commit()
	}

	fillUpTree(verticalaxisFormCallback.probe)
}
