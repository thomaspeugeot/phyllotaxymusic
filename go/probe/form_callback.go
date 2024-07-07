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
		case "AxisHandleBorderLength":
			FormDivBasicFieldToField(&(horizontalaxis_.AxisHandleBorderLength), formDiv)
		case "Axis_Length":
			FormDivBasicFieldToField(&(horizontalaxis_.Axis_Length), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(horizontalaxis_.StrokeWidth), formDiv)
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
func __gong__New__LineFormCallback(
	line *models.Line,
	probe *Probe,
	formGroup *table.FormGroup,
) (lineFormCallback *LineFormCallback) {
	lineFormCallback = new(LineFormCallback)
	lineFormCallback.probe = probe
	lineFormCallback.line = line
	lineFormCallback.formGroup = formGroup

	lineFormCallback.CreationMode = (line == nil)

	return
}

type LineFormCallback struct {
	line *models.Line

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lineFormCallback *LineFormCallback) OnSave() {

	log.Println("LineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lineFormCallback.probe.formStage.Checkout()

	if lineFormCallback.line == nil {
		lineFormCallback.line = new(models.Line).Stage(lineFormCallback.probe.stageOfInterest)
	}
	line_ := lineFormCallback.line
	_ = line_

	for _, formDiv := range lineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(line_.Name), formDiv)
		case "X1":
			FormDivBasicFieldToField(&(line_.X1), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(line_.Y1), formDiv)
		case "X2":
			FormDivBasicFieldToField(&(line_.X2), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(line_.Y2), formDiv)
		}
	}

	// manage the suppress operation
	if lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		line_.Unstage(lineFormCallback.probe.stageOfInterest)
	}

	lineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Line](
		lineFormCallback.probe,
	)
	lineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lineFormCallback.CreationMode || lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LineFormCallback(
			nil,
			lineFormCallback.probe,
			newFormGroup,
		)
		line := new(models.Line)
		FillUpForm(line, newFormGroup, lineFormCallback.probe)
		lineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lineFormCallback.probe)
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
		case "InsideAngle":
			FormDivBasicFieldToField(&(parameter_.InsideAngle), formDiv)
		case "DiamondSideLenght":
			FormDivBasicFieldToField(&(parameter_.DiamondSideLenght), formDiv)
		case "InitialRhombus":
			FormDivSelectFieldToField(&(parameter_.InitialRhombus), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "InitialRhombusGrid":
			FormDivSelectFieldToField(&(parameter_.InitialRhombusGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&(parameter_.OriginX), formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&(parameter_.OriginY), formDiv)
		case "HorizontalAxis":
			FormDivSelectFieldToField(&(parameter_.HorizontalAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "VerticalAxis":
			FormDivSelectFieldToField(&(parameter_.VerticalAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
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
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rhombus_.StrokeWidth), formDiv)
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
		case "N":
			FormDivBasicFieldToField(&(rhombusgrid_.N), formDiv)
		case "M":
			FormDivBasicFieldToField(&(rhombusgrid_.M), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(rhombusgrid_.IsDisplayed), formDiv)
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
		case "AxisHandleBorderLength":
			FormDivBasicFieldToField(&(verticalaxis_.AxisHandleBorderLength), formDiv)
		case "Axis_Length":
			FormDivBasicFieldToField(&(verticalaxis_.Axis_Length), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(verticalaxis_.StrokeWidth), formDiv)
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
