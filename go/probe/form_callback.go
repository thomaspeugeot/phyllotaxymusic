// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
)

// code to avoid error when generated code does not need to import packages
const __dummmy__time = time.Nanosecond

var _ = __dummmy__time

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

var _ = __dummmy__letters

const __dummy__log = log.Ldate

var _ = __dummy__log

// end of code to avoid error when generated code does not need to import packages

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

	// log.Println("AxisFormCallback, OnSave")

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
		case "AngleDegree":
			FormDivBasicFieldToField(&(axis_.AngleDegree), formDiv)
		case "Length":
			FormDivBasicFieldToField(&(axis_.Length), formDiv)
		case "CenterX":
			FormDivBasicFieldToField(&(axis_.CenterX), formDiv)
		case "CenterY":
			FormDivBasicFieldToField(&(axis_.CenterY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(axis_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(axis_.EndY), formDiv)
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
			// WARNING : this form deals with the N-N association "AxisGrid.Axiss []*Axis" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Axis). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.AxisGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "AxisGrid"
				rf.Fieldname = "Axiss"
				formerAssociationSource := models.GetReverseFieldOwner(
					axisFormCallback.probe.stageOfInterest,
					axis_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.AxisGrid)
					if !ok {
						log.Fatalln("Source of AxisGrid.Axiss []*Axis, is not an AxisGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.AxisGrid
			for _axisgrid := range *models.GetGongstructInstancesSet[models.AxisGrid](axisFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _axisgrid.GetName() == newSourceName.GetName() {
					newSource = _axisgrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of AxisGrid.Axiss []*Axis, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Axiss = append(newSource.Axiss, axis_)
		}
	}

	// manage the suppress operation
	if axisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axis_.Unstage(axisFormCallback.probe.stageOfInterest)
	}

	axisFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Axis](
		axisFormCallback.probe,
	)
	axisFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if axisFormCallback.CreationMode || axisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axisFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(axisFormCallback.probe)
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

	// log.Println("AxisGridFormCallback, OnSave")

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
	updateAndCommitTable[models.AxisGrid](
		axisgridFormCallback.probe,
	)
	axisgridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if axisgridFormCallback.CreationMode || axisgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axisgridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(axisgridFormCallback.probe)
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

	// log.Println("BezierFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "BezierGrid.Beziers []*Bezier" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Bezier). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.BezierGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "BezierGrid"
				rf.Fieldname = "Beziers"
				formerAssociationSource := models.GetReverseFieldOwner(
					bezierFormCallback.probe.stageOfInterest,
					bezier_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.BezierGrid)
					if !ok {
						log.Fatalln("Source of BezierGrid.Beziers []*Bezier, is not an BezierGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.BezierGrid
			for _beziergrid := range *models.GetGongstructInstancesSet[models.BezierGrid](bezierFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _beziergrid.GetName() == newSourceName.GetName() {
					newSource = _beziergrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of BezierGrid.Beziers []*Bezier, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Beziers = append(newSource.Beziers, bezier_)
		}
	}

	// manage the suppress operation
	if bezierFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bezier_.Unstage(bezierFormCallback.probe.stageOfInterest)
	}

	bezierFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Bezier](
		bezierFormCallback.probe,
	)
	bezierFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bezierFormCallback.CreationMode || bezierFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bezierFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(bezierFormCallback.probe)
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

	// log.Println("BezierGridFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "BezierGridStack.BezierGrids []*BezierGrid" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of BezierGrid). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.BezierGridStack
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "BezierGridStack"
				rf.Fieldname = "BezierGrids"
				formerAssociationSource := models.GetReverseFieldOwner(
					beziergridFormCallback.probe.stageOfInterest,
					beziergrid_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.BezierGridStack)
					if !ok {
						log.Fatalln("Source of BezierGridStack.BezierGrids []*BezierGrid, is not an BezierGridStack instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.BezierGridStack
			for _beziergridstack := range *models.GetGongstructInstancesSet[models.BezierGridStack](beziergridFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _beziergridstack.GetName() == newSourceName.GetName() {
					newSource = _beziergridstack // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of BezierGridStack.BezierGrids []*BezierGrid, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.BezierGrids = append(newSource.BezierGrids, beziergrid_)
		}
	}

	// manage the suppress operation
	if beziergridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziergrid_.Unstage(beziergridFormCallback.probe.stageOfInterest)
	}

	beziergridFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.BezierGrid](
		beziergridFormCallback.probe,
	)
	beziergridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beziergridFormCallback.CreationMode || beziergridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziergridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(beziergridFormCallback.probe)
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

	// log.Println("BezierGridStackFormCallback, OnSave")

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
	updateAndCommitTable[models.BezierGridStack](
		beziergridstackFormCallback.probe,
	)
	beziergridstackFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beziergridstackFormCallback.CreationMode || beziergridstackFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beziergridstackFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(beziergridstackFormCallback.probe)
}
func __gong__New__ChapterFormCallback(
	chapter *models.Chapter,
	probe *Probe,
	formGroup *table.FormGroup,
) (chapterFormCallback *ChapterFormCallback) {
	chapterFormCallback = new(ChapterFormCallback)
	chapterFormCallback.probe = probe
	chapterFormCallback.chapter = chapter
	chapterFormCallback.formGroup = formGroup

	chapterFormCallback.CreationMode = (chapter == nil)

	return
}

type ChapterFormCallback struct {
	chapter *models.Chapter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (chapterFormCallback *ChapterFormCallback) OnSave() {

	// log.Println("ChapterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	chapterFormCallback.probe.formStage.Checkout()

	if chapterFormCallback.chapter == nil {
		chapterFormCallback.chapter = new(models.Chapter).Stage(chapterFormCallback.probe.stageOfInterest)
	}
	chapter_ := chapterFormCallback.chapter
	_ = chapter_

	for _, formDiv := range chapterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(chapter_.Name), formDiv)
		case "MardownContent":
			FormDivBasicFieldToField(&(chapter_.MardownContent), formDiv)
		case "Content:Chapters":
			// WARNING : this form deals with the N-N association "Content.Chapters []*Chapter" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Chapter). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Content
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Content"
				rf.Fieldname = "Chapters"
				formerAssociationSource := models.GetReverseFieldOwner(
					chapterFormCallback.probe.stageOfInterest,
					chapter_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Content)
					if !ok {
						log.Fatalln("Source of Content.Chapters []*Chapter, is not an Content instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Content
			for _content := range *models.GetGongstructInstancesSet[models.Content](chapterFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _content.GetName() == newSourceName.GetName() {
					newSource = _content // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Content.Chapters []*Chapter, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Chapters = append(newSource.Chapters, chapter_)
		}
	}

	// manage the suppress operation
	if chapterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		chapter_.Unstage(chapterFormCallback.probe.stageOfInterest)
	}

	chapterFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Chapter](
		chapterFormCallback.probe,
	)
	chapterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if chapterFormCallback.CreationMode || chapterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		chapterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(chapterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ChapterFormCallback(
			nil,
			chapterFormCallback.probe,
			newFormGroup,
		)
		chapter := new(models.Chapter)
		FillUpForm(chapter, newFormGroup, chapterFormCallback.probe)
		chapterFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(chapterFormCallback.probe)
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

	// log.Println("CircleFormCallback, OnSave")

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
		case "Pitch":
			FormDivBasicFieldToField(&(circle_.Pitch), formDiv)
		case "ShowName":
			FormDivBasicFieldToField(&(circle_.ShowName), formDiv)
		case "BeatNb":
			FormDivBasicFieldToField(&(circle_.BeatNb), formDiv)
		case "CircleGrid:Circles":
			// WARNING : this form deals with the N-N association "CircleGrid.Circles []*Circle" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Circle). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.CircleGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "CircleGrid"
				rf.Fieldname = "Circles"
				formerAssociationSource := models.GetReverseFieldOwner(
					circleFormCallback.probe.stageOfInterest,
					circle_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.CircleGrid)
					if !ok {
						log.Fatalln("Source of CircleGrid.Circles []*Circle, is not an CircleGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.CircleGrid
			for _circlegrid := range *models.GetGongstructInstancesSet[models.CircleGrid](circleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _circlegrid.GetName() == newSourceName.GetName() {
					newSource = _circlegrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of CircleGrid.Circles []*Circle, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Circles = append(newSource.Circles, circle_)
		}
	}

	// manage the suppress operation
	if circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circle_.Unstage(circleFormCallback.probe.stageOfInterest)
	}

	circleFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Circle](
		circleFormCallback.probe,
	)
	circleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode || circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(circleFormCallback.probe)
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

	// log.Println("CircleGridFormCallback, OnSave")

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
	updateAndCommitTable[models.CircleGrid](
		circlegridFormCallback.probe,
	)
	circlegridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if circlegridFormCallback.CreationMode || circlegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circlegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(circlegridFormCallback.probe)
}
func __gong__New__ContentFormCallback(
	content *models.Content,
	probe *Probe,
	formGroup *table.FormGroup,
) (contentFormCallback *ContentFormCallback) {
	contentFormCallback = new(ContentFormCallback)
	contentFormCallback.probe = probe
	contentFormCallback.content = content
	contentFormCallback.formGroup = formGroup

	contentFormCallback.CreationMode = (content == nil)

	return
}

type ContentFormCallback struct {
	content *models.Content

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (contentFormCallback *ContentFormCallback) OnSave() {

	// log.Println("ContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	contentFormCallback.probe.formStage.Checkout()

	if contentFormCallback.content == nil {
		contentFormCallback.content = new(models.Content).Stage(contentFormCallback.probe.stageOfInterest)
	}
	content_ := contentFormCallback.content
	_ = content_

	for _, formDiv := range contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(content_.Name), formDiv)
		case "MardownContent":
			FormDivBasicFieldToField(&(content_.MardownContent), formDiv)
		case "ContentPath":
			FormDivBasicFieldToField(&(content_.ContentPath), formDiv)
		case "OutputPath":
			FormDivBasicFieldToField(&(content_.OutputPath), formDiv)
		case "LayoutPath":
			FormDivBasicFieldToField(&(content_.LayoutPath), formDiv)
		case "StaticPath":
			FormDivBasicFieldToField(&(content_.StaticPath), formDiv)
		case "Target":
			FormDivEnumStringFieldToField(&(content_.Target), formDiv)
		}
	}

	// manage the suppress operation
	if contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		content_.Unstage(contentFormCallback.probe.stageOfInterest)
	}

	contentFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Content](
		contentFormCallback.probe,
	)
	contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if contentFormCallback.CreationMode || contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ContentFormCallback(
			nil,
			contentFormCallback.probe,
			newFormGroup,
		)
		content := new(models.Content)
		FillUpForm(content, newFormGroup, contentFormCallback.probe)
		contentFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(contentFormCallback.probe)
}
func __gong__New__ExportToMusicxmlFormCallback(
	exporttomusicxml *models.ExportToMusicxml,
	probe *Probe,
	formGroup *table.FormGroup,
) (exporttomusicxmlFormCallback *ExportToMusicxmlFormCallback) {
	exporttomusicxmlFormCallback = new(ExportToMusicxmlFormCallback)
	exporttomusicxmlFormCallback.probe = probe
	exporttomusicxmlFormCallback.exporttomusicxml = exporttomusicxml
	exporttomusicxmlFormCallback.formGroup = formGroup

	exporttomusicxmlFormCallback.CreationMode = (exporttomusicxml == nil)

	return
}

type ExportToMusicxmlFormCallback struct {
	exporttomusicxml *models.ExportToMusicxml

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (exporttomusicxmlFormCallback *ExportToMusicxmlFormCallback) OnSave() {

	// log.Println("ExportToMusicxmlFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	exporttomusicxmlFormCallback.probe.formStage.Checkout()

	if exporttomusicxmlFormCallback.exporttomusicxml == nil {
		exporttomusicxmlFormCallback.exporttomusicxml = new(models.ExportToMusicxml).Stage(exporttomusicxmlFormCallback.probe.stageOfInterest)
	}
	exporttomusicxml_ := exporttomusicxmlFormCallback.exporttomusicxml
	_ = exporttomusicxml_

	for _, formDiv := range exporttomusicxmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(exporttomusicxml_.Name), formDiv)
		case "Parameter":
			FormDivSelectFieldToField(&(exporttomusicxml_.Parameter), exporttomusicxmlFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if exporttomusicxmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		exporttomusicxml_.Unstage(exporttomusicxmlFormCallback.probe.stageOfInterest)
	}

	exporttomusicxmlFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.ExportToMusicxml](
		exporttomusicxmlFormCallback.probe,
	)
	exporttomusicxmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if exporttomusicxmlFormCallback.CreationMode || exporttomusicxmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		exporttomusicxmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(exporttomusicxmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExportToMusicxmlFormCallback(
			nil,
			exporttomusicxmlFormCallback.probe,
			newFormGroup,
		)
		exporttomusicxml := new(models.ExportToMusicxml)
		FillUpForm(exporttomusicxml, newFormGroup, exporttomusicxmlFormCallback.probe)
		exporttomusicxmlFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(exporttomusicxmlFormCallback.probe)
}
func __gong__New__FrontCurveFormCallback(
	frontcurve *models.FrontCurve,
	probe *Probe,
	formGroup *table.FormGroup,
) (frontcurveFormCallback *FrontCurveFormCallback) {
	frontcurveFormCallback = new(FrontCurveFormCallback)
	frontcurveFormCallback.probe = probe
	frontcurveFormCallback.frontcurve = frontcurve
	frontcurveFormCallback.formGroup = formGroup

	frontcurveFormCallback.CreationMode = (frontcurve == nil)

	return
}

type FrontCurveFormCallback struct {
	frontcurve *models.FrontCurve

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (frontcurveFormCallback *FrontCurveFormCallback) OnSave() {

	// log.Println("FrontCurveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	frontcurveFormCallback.probe.formStage.Checkout()

	if frontcurveFormCallback.frontcurve == nil {
		frontcurveFormCallback.frontcurve = new(models.FrontCurve).Stage(frontcurveFormCallback.probe.stageOfInterest)
	}
	frontcurve_ := frontcurveFormCallback.frontcurve
	_ = frontcurve_

	for _, formDiv := range frontcurveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(frontcurve_.Name), formDiv)
		case "Path":
			FormDivBasicFieldToField(&(frontcurve_.Path), formDiv)
		case "FrontCurveStack:FrontCurves":
			// WARNING : this form deals with the N-N association "FrontCurveStack.FrontCurves []*FrontCurve" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of FrontCurve). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.FrontCurveStack
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "FrontCurveStack"
				rf.Fieldname = "FrontCurves"
				formerAssociationSource := models.GetReverseFieldOwner(
					frontcurveFormCallback.probe.stageOfInterest,
					frontcurve_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.FrontCurveStack)
					if !ok {
						log.Fatalln("Source of FrontCurveStack.FrontCurves []*FrontCurve, is not an FrontCurveStack instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.FrontCurveStack
			for _frontcurvestack := range *models.GetGongstructInstancesSet[models.FrontCurveStack](frontcurveFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _frontcurvestack.GetName() == newSourceName.GetName() {
					newSource = _frontcurvestack // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of FrontCurveStack.FrontCurves []*FrontCurve, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.FrontCurves = append(newSource.FrontCurves, frontcurve_)
		}
	}

	// manage the suppress operation
	if frontcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frontcurve_.Unstage(frontcurveFormCallback.probe.stageOfInterest)
	}

	frontcurveFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FrontCurve](
		frontcurveFormCallback.probe,
	)
	frontcurveFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if frontcurveFormCallback.CreationMode || frontcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frontcurveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(frontcurveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FrontCurveFormCallback(
			nil,
			frontcurveFormCallback.probe,
			newFormGroup,
		)
		frontcurve := new(models.FrontCurve)
		FillUpForm(frontcurve, newFormGroup, frontcurveFormCallback.probe)
		frontcurveFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(frontcurveFormCallback.probe)
}
func __gong__New__FrontCurveStackFormCallback(
	frontcurvestack *models.FrontCurveStack,
	probe *Probe,
	formGroup *table.FormGroup,
) (frontcurvestackFormCallback *FrontCurveStackFormCallback) {
	frontcurvestackFormCallback = new(FrontCurveStackFormCallback)
	frontcurvestackFormCallback.probe = probe
	frontcurvestackFormCallback.frontcurvestack = frontcurvestack
	frontcurvestackFormCallback.formGroup = formGroup

	frontcurvestackFormCallback.CreationMode = (frontcurvestack == nil)

	return
}

type FrontCurveStackFormCallback struct {
	frontcurvestack *models.FrontCurveStack

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (frontcurvestackFormCallback *FrontCurveStackFormCallback) OnSave() {

	// log.Println("FrontCurveStackFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	frontcurvestackFormCallback.probe.formStage.Checkout()

	if frontcurvestackFormCallback.frontcurvestack == nil {
		frontcurvestackFormCallback.frontcurvestack = new(models.FrontCurveStack).Stage(frontcurvestackFormCallback.probe.stageOfInterest)
	}
	frontcurvestack_ := frontcurvestackFormCallback.frontcurvestack
	_ = frontcurvestack_

	for _, formDiv := range frontcurvestackFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(frontcurvestack_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(frontcurvestack_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(frontcurvestack_.ShapeCategory), frontcurvestackFormCallback.probe.stageOfInterest, formDiv)
		case "Color":
			FormDivBasicFieldToField(&(frontcurvestack_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(frontcurvestack_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(frontcurvestack_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(frontcurvestack_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(frontcurvestack_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(frontcurvestack_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(frontcurvestack_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(frontcurvestack_.Transform), formDiv)
		}
	}

	// manage the suppress operation
	if frontcurvestackFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frontcurvestack_.Unstage(frontcurvestackFormCallback.probe.stageOfInterest)
	}

	frontcurvestackFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FrontCurveStack](
		frontcurvestackFormCallback.probe,
	)
	frontcurvestackFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if frontcurvestackFormCallback.CreationMode || frontcurvestackFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frontcurvestackFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(frontcurvestackFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FrontCurveStackFormCallback(
			nil,
			frontcurvestackFormCallback.probe,
			newFormGroup,
		)
		frontcurvestack := new(models.FrontCurveStack)
		FillUpForm(frontcurvestack, newFormGroup, frontcurvestackFormCallback.probe)
		frontcurvestackFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(frontcurvestackFormCallback.probe)
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

	// log.Println("HorizontalAxisFormCallback, OnSave")

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
	updateAndCommitTable[models.HorizontalAxis](
		horizontalaxisFormCallback.probe,
	)
	horizontalaxisFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if horizontalaxisFormCallback.CreationMode || horizontalaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		horizontalaxisFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(horizontalaxisFormCallback.probe)
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

	// log.Println("KeyFormCallback, OnSave")

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
	updateAndCommitTable[models.Key](
		keyFormCallback.probe,
	)
	keyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if keyFormCallback.CreationMode || keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		keyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(keyFormCallback.probe)
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

	// log.Println("ParameterFormCallback, OnSave")

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
		case "BackendColor":
			FormDivBasicFieldToField(&(parameter_.BackendColor), formDiv)
		case "MinuteColor":
			FormDivBasicFieldToField(&(parameter_.MinuteColor), formDiv)
		case "HourColor":
			FormDivBasicFieldToField(&(parameter_.HourColor), formDiv)
		case "N":
			FormDivBasicFieldToField(&(parameter_.N), formDiv)
		case "M":
			FormDivBasicFieldToField(&(parameter_.M), formDiv)
		case "Z":
			FormDivBasicFieldToField(&(parameter_.Z), formDiv)
		case "InsideAngle":
			FormDivBasicFieldToField(&(parameter_.InsideAngle), formDiv)
		case "ShiftToNearestCircle":
			FormDivBasicFieldToField(&(parameter_.ShiftToNearestCircle), formDiv)
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
		case "GrowthCurveSeed":
			FormDivSelectFieldToField(&(parameter_.GrowthCurveSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
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
		case "SpiralRhombusGridSeed":
			FormDivSelectFieldToField(&(parameter_.SpiralRhombusGridSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralRhombusGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralRhombusGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralCircleSeed":
			FormDivSelectFieldToField(&(parameter_.SpiralCircleSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralCircleGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralCircleGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralCircleFullGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralCircleFullGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralConstructionOuterLineSeed":
			FormDivSelectFieldToField(&(parameter_.SpiralConstructionOuterLineSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralConstructionInnerLineSeed":
			FormDivSelectFieldToField(&(parameter_.SpiralConstructionInnerLineSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralConstructionOuterLineGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralConstructionOuterLineGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralConstructionInnerLineGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralConstructionInnerLineGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralConstructionCircleGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralConstructionCircleGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralConstructionOuterLineFullGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralConstructionOuterLineFullGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralBezierSeed":
			FormDivSelectFieldToField(&(parameter_.SpiralBezierSeed), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralBezierGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralBezierGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralBezierFullGrid":
			FormDivSelectFieldToField(&(parameter_.SpiralBezierFullGrid), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralBezierStrength":
			FormDivBasicFieldToField(&(parameter_.SpiralBezierStrength), formDiv)
		case "FrontCurveStack":
			FormDivSelectFieldToField(&(parameter_.FrontCurveStack), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "NbInterpolationPoints":
			FormDivBasicFieldToField(&(parameter_.NbInterpolationPoints), formDiv)
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
		case "BeatLines":
			FormDivSelectFieldToField(&(parameter_.BeatLines), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "BeatLinesHeightRatio":
			FormDivBasicFieldToField(&(parameter_.BeatLinesHeightRatio), formDiv)
		case "NbBeatLines":
			FormDivBasicFieldToField(&(parameter_.NbBeatLines), formDiv)
		case "NbOfBeatsInTheme":
			FormDivBasicFieldToField(&(parameter_.NbOfBeatsInTheme), formDiv)
		case "FirstVoice":
			FormDivSelectFieldToField(&(parameter_.FirstVoice), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "FirstVoiceShiftedRigth":
			FormDivSelectFieldToField(&(parameter_.FirstVoiceShiftedRigth), parameterFormCallback.probe.stageOfInterest, formDiv)
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
		case "BeatsPerSecond":
			FormDivBasicFieldToField(&(parameter_.BeatsPerSecond), formDiv)
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
		case "ThemeBinaryEncoding":
			FormDivBasicFieldToField(&(parameter_.ThemeBinaryEncoding), formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&(parameter_.OriginX), formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&(parameter_.OriginY), formDiv)
		case "HorizontalAxis":
			FormDivSelectFieldToField(&(parameter_.HorizontalAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "VerticalAxis":
			FormDivSelectFieldToField(&(parameter_.VerticalAxis), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralOrigin":
			FormDivSelectFieldToField(&(parameter_.SpiralOrigin), parameterFormCallback.probe.stageOfInterest, formDiv)
		case "SpiralOriginX":
			FormDivBasicFieldToField(&(parameter_.SpiralOriginX), formDiv)
		case "SpiralOriginY":
			FormDivBasicFieldToField(&(parameter_.SpiralOriginY), formDiv)
		case "OriginCrossWidth":
			FormDivBasicFieldToField(&(parameter_.OriginCrossWidth), formDiv)
		case "SpiralRadiusRatio":
			FormDivBasicFieldToField(&(parameter_.SpiralRadiusRatio), formDiv)
		case "ShowSpiralBezierConstruct":
			FormDivBasicFieldToField(&(parameter_.ShowSpiralBezierConstruct), formDiv)
		case "ShowInterpolationPoints":
			FormDivBasicFieldToField(&(parameter_.ShowInterpolationPoints), formDiv)
		case "ActualBeatsTemporalShift":
			FormDivBasicFieldToField(&(parameter_.ActualBeatsTemporalShift), formDiv)
		case "PathToStaticFiles":
			FormDivBasicFieldToField(&(parameter_.PathToStaticFiles), formDiv)
		case "PathToGeneratedSVG":
			FormDivBasicFieldToField(&(parameter_.PathToGeneratedSVG), formDiv)
		case "PathToGeneratedScore":
			FormDivBasicFieldToField(&(parameter_.PathToGeneratedScore), formDiv)
		}
	}

	// manage the suppress operation
	if parameterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parameter_.Unstage(parameterFormCallback.probe.stageOfInterest)
	}

	parameterFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Parameter](
		parameterFormCallback.probe,
	)
	parameterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if parameterFormCallback.CreationMode || parameterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parameterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(parameterFormCallback.probe)
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

	// log.Println("RhombusFormCallback, OnSave")

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
		case "AngleDegree":
			FormDivBasicFieldToField(&(rhombus_.AngleDegree), formDiv)
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
			// WARNING : this form deals with the N-N association "RhombusGrid.Rhombuses []*Rhombus" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Rhombus). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.RhombusGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "RhombusGrid"
				rf.Fieldname = "Rhombuses"
				formerAssociationSource := models.GetReverseFieldOwner(
					rhombusFormCallback.probe.stageOfInterest,
					rhombus_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.RhombusGrid)
					if !ok {
						log.Fatalln("Source of RhombusGrid.Rhombuses []*Rhombus, is not an RhombusGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.RhombusGrid
			for _rhombusgrid := range *models.GetGongstructInstancesSet[models.RhombusGrid](rhombusFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _rhombusgrid.GetName() == newSourceName.GetName() {
					newSource = _rhombusgrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of RhombusGrid.Rhombuses []*Rhombus, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Rhombuses = append(newSource.Rhombuses, rhombus_)
		}
	}

	// manage the suppress operation
	if rhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombus_.Unstage(rhombusFormCallback.probe.stageOfInterest)
	}

	rhombusFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Rhombus](
		rhombusFormCallback.probe,
	)
	rhombusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rhombusFormCallback.CreationMode || rhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(rhombusFormCallback.probe)
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

	// log.Println("RhombusGridFormCallback, OnSave")

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
	updateAndCommitTable[models.RhombusGrid](
		rhombusgridFormCallback.probe,
	)
	rhombusgridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rhombusgridFormCallback.CreationMode || rhombusgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusgridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(rhombusgridFormCallback.probe)
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

	// log.Println("ShapeCategoryFormCallback, OnSave")

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
	updateAndCommitTable[models.ShapeCategory](
		shapecategoryFormCallback.probe,
	)
	shapecategoryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if shapecategoryFormCallback.CreationMode || shapecategoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shapecategoryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(shapecategoryFormCallback.probe)
}
func __gong__New__SpiralBezierFormCallback(
	spiralbezier *models.SpiralBezier,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiralbezierFormCallback *SpiralBezierFormCallback) {
	spiralbezierFormCallback = new(SpiralBezierFormCallback)
	spiralbezierFormCallback.probe = probe
	spiralbezierFormCallback.spiralbezier = spiralbezier
	spiralbezierFormCallback.formGroup = formGroup

	spiralbezierFormCallback.CreationMode = (spiralbezier == nil)

	return
}

type SpiralBezierFormCallback struct {
	spiralbezier *models.SpiralBezier

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiralbezierFormCallback *SpiralBezierFormCallback) OnSave() {

	// log.Println("SpiralBezierFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiralbezierFormCallback.probe.formStage.Checkout()

	if spiralbezierFormCallback.spiralbezier == nil {
		spiralbezierFormCallback.spiralbezier = new(models.SpiralBezier).Stage(spiralbezierFormCallback.probe.stageOfInterest)
	}
	spiralbezier_ := spiralbezierFormCallback.spiralbezier
	_ = spiralbezier_

	for _, formDiv := range spiralbezierFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralbezier_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralbezier_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralbezier_.ShapeCategory), spiralbezierFormCallback.probe.stageOfInterest, formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(spiralbezier_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(spiralbezier_.StartY), formDiv)
		case "ControlPointStartX":
			FormDivBasicFieldToField(&(spiralbezier_.ControlPointStartX), formDiv)
		case "ControlPointStartY":
			FormDivBasicFieldToField(&(spiralbezier_.ControlPointStartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(spiralbezier_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(spiralbezier_.EndY), formDiv)
		case "ControlPointEndX":
			FormDivBasicFieldToField(&(spiralbezier_.ControlPointEndX), formDiv)
		case "ControlPointEndY":
			FormDivBasicFieldToField(&(spiralbezier_.ControlPointEndY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(spiralbezier_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(spiralbezier_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(spiralbezier_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(spiralbezier_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(spiralbezier_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(spiralbezier_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(spiralbezier_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(spiralbezier_.Transform), formDiv)
		case "SpiralBezierGrid:SpiralBeziers":
			// WARNING : this form deals with the N-N association "SpiralBezierGrid.SpiralBeziers []*SpiralBezier" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SpiralBezier). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.SpiralBezierGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "SpiralBezierGrid"
				rf.Fieldname = "SpiralBeziers"
				formerAssociationSource := models.GetReverseFieldOwner(
					spiralbezierFormCallback.probe.stageOfInterest,
					spiralbezier_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.SpiralBezierGrid)
					if !ok {
						log.Fatalln("Source of SpiralBezierGrid.SpiralBeziers []*SpiralBezier, is not an SpiralBezierGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.SpiralBezierGrid
			for _spiralbeziergrid := range *models.GetGongstructInstancesSet[models.SpiralBezierGrid](spiralbezierFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _spiralbeziergrid.GetName() == newSourceName.GetName() {
					newSource = _spiralbeziergrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of SpiralBezierGrid.SpiralBeziers []*SpiralBezier, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.SpiralBeziers = append(newSource.SpiralBeziers, spiralbezier_)
		}
	}

	// manage the suppress operation
	if spiralbezierFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralbezier_.Unstage(spiralbezierFormCallback.probe.stageOfInterest)
	}

	spiralbezierFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralBezier](
		spiralbezierFormCallback.probe,
	)
	spiralbezierFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralbezierFormCallback.CreationMode || spiralbezierFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralbezierFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spiralbezierFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralBezierFormCallback(
			nil,
			spiralbezierFormCallback.probe,
			newFormGroup,
		)
		spiralbezier := new(models.SpiralBezier)
		FillUpForm(spiralbezier, newFormGroup, spiralbezierFormCallback.probe)
		spiralbezierFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spiralbezierFormCallback.probe)
}
func __gong__New__SpiralBezierGridFormCallback(
	spiralbeziergrid *models.SpiralBezierGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiralbeziergridFormCallback *SpiralBezierGridFormCallback) {
	spiralbeziergridFormCallback = new(SpiralBezierGridFormCallback)
	spiralbeziergridFormCallback.probe = probe
	spiralbeziergridFormCallback.spiralbeziergrid = spiralbeziergrid
	spiralbeziergridFormCallback.formGroup = formGroup

	spiralbeziergridFormCallback.CreationMode = (spiralbeziergrid == nil)

	return
}

type SpiralBezierGridFormCallback struct {
	spiralbeziergrid *models.SpiralBezierGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiralbeziergridFormCallback *SpiralBezierGridFormCallback) OnSave() {

	// log.Println("SpiralBezierGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiralbeziergridFormCallback.probe.formStage.Checkout()

	if spiralbeziergridFormCallback.spiralbeziergrid == nil {
		spiralbeziergridFormCallback.spiralbeziergrid = new(models.SpiralBezierGrid).Stage(spiralbeziergridFormCallback.probe.stageOfInterest)
	}
	spiralbeziergrid_ := spiralbeziergridFormCallback.spiralbeziergrid
	_ = spiralbeziergrid_

	for _, formDiv := range spiralbeziergridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralbeziergrid_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralbeziergrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralbeziergrid_.ShapeCategory), spiralbeziergridFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if spiralbeziergridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralbeziergrid_.Unstage(spiralbeziergridFormCallback.probe.stageOfInterest)
	}

	spiralbeziergridFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralBezierGrid](
		spiralbeziergridFormCallback.probe,
	)
	spiralbeziergridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralbeziergridFormCallback.CreationMode || spiralbeziergridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralbeziergridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spiralbeziergridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralBezierGridFormCallback(
			nil,
			spiralbeziergridFormCallback.probe,
			newFormGroup,
		)
		spiralbeziergrid := new(models.SpiralBezierGrid)
		FillUpForm(spiralbeziergrid, newFormGroup, spiralbeziergridFormCallback.probe)
		spiralbeziergridFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spiralbeziergridFormCallback.probe)
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

	// log.Println("SpiralCircleFormCallback, OnSave")

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
		case "Pitch":
			FormDivBasicFieldToField(&(spiralcircle_.Pitch), formDiv)
		case "ShowName":
			FormDivBasicFieldToField(&(spiralcircle_.ShowName), formDiv)
		case "BeatNb":
			FormDivBasicFieldToField(&(spiralcircle_.BeatNb), formDiv)
		case "Path":
			FormDivBasicFieldToField(&(spiralcircle_.Path), formDiv)
		case "FrontCurveStack:SpiralCircles":
			// WARNING : this form deals with the N-N association "FrontCurveStack.SpiralCircles []*SpiralCircle" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SpiralCircle). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.FrontCurveStack
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "FrontCurveStack"
				rf.Fieldname = "SpiralCircles"
				formerAssociationSource := models.GetReverseFieldOwner(
					spiralcircleFormCallback.probe.stageOfInterest,
					spiralcircle_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.FrontCurveStack)
					if !ok {
						log.Fatalln("Source of FrontCurveStack.SpiralCircles []*SpiralCircle, is not an FrontCurveStack instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.FrontCurveStack
			for _frontcurvestack := range *models.GetGongstructInstancesSet[models.FrontCurveStack](spiralcircleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _frontcurvestack.GetName() == newSourceName.GetName() {
					newSource = _frontcurvestack // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of FrontCurveStack.SpiralCircles []*SpiralCircle, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.SpiralCircles = append(newSource.SpiralCircles, spiralcircle_)
		case "SpiralCircleGrid:SpiralCircles":
			// WARNING : this form deals with the N-N association "SpiralCircleGrid.SpiralCircles []*SpiralCircle" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SpiralCircle). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.SpiralCircleGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "SpiralCircleGrid"
				rf.Fieldname = "SpiralCircles"
				formerAssociationSource := models.GetReverseFieldOwner(
					spiralcircleFormCallback.probe.stageOfInterest,
					spiralcircle_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.SpiralCircleGrid)
					if !ok {
						log.Fatalln("Source of SpiralCircleGrid.SpiralCircles []*SpiralCircle, is not an SpiralCircleGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.SpiralCircleGrid
			for _spiralcirclegrid := range *models.GetGongstructInstancesSet[models.SpiralCircleGrid](spiralcircleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _spiralcirclegrid.GetName() == newSourceName.GetName() {
					newSource = _spiralcirclegrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of SpiralCircleGrid.SpiralCircles []*SpiralCircle, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.SpiralCircles = append(newSource.SpiralCircles, spiralcircle_)
		}
	}

	// manage the suppress operation
	if spiralcircleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralcircle_.Unstage(spiralcircleFormCallback.probe.stageOfInterest)
	}

	spiralcircleFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralCircle](
		spiralcircleFormCallback.probe,
	)
	spiralcircleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralcircleFormCallback.CreationMode || spiralcircleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralcircleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(spiralcircleFormCallback.probe)
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

	// log.Println("SpiralCircleGridFormCallback, OnSave")

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
	updateAndCommitTable[models.SpiralCircleGrid](
		spiralcirclegridFormCallback.probe,
	)
	spiralcirclegridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralcirclegridFormCallback.CreationMode || spiralcirclegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralcirclegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(spiralcirclegridFormCallback.probe)
}
func __gong__New__SpiralLineFormCallback(
	spiralline *models.SpiralLine,
	probe *Probe,
	formGroup *table.FormGroup,
) (spirallineFormCallback *SpiralLineFormCallback) {
	spirallineFormCallback = new(SpiralLineFormCallback)
	spirallineFormCallback.probe = probe
	spirallineFormCallback.spiralline = spiralline
	spirallineFormCallback.formGroup = formGroup

	spirallineFormCallback.CreationMode = (spiralline == nil)

	return
}

type SpiralLineFormCallback struct {
	spiralline *models.SpiralLine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spirallineFormCallback *SpiralLineFormCallback) OnSave() {

	// log.Println("SpiralLineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spirallineFormCallback.probe.formStage.Checkout()

	if spirallineFormCallback.spiralline == nil {
		spirallineFormCallback.spiralline = new(models.SpiralLine).Stage(spirallineFormCallback.probe.stageOfInterest)
	}
	spiralline_ := spirallineFormCallback.spiralline
	_ = spiralline_

	for _, formDiv := range spirallineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralline_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralline_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralline_.ShapeCategory), spirallineFormCallback.probe.stageOfInterest, formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(spiralline_.StartX), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(spiralline_.EndX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(spiralline_.StartY), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(spiralline_.EndY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(spiralline_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(spiralline_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(spiralline_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(spiralline_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(spiralline_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(spiralline_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(spiralline_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(spiralline_.Transform), formDiv)
		case "SpiralLineGrid:SpiralLines":
			// WARNING : this form deals with the N-N association "SpiralLineGrid.SpiralLines []*SpiralLine" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SpiralLine). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.SpiralLineGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "SpiralLineGrid"
				rf.Fieldname = "SpiralLines"
				formerAssociationSource := models.GetReverseFieldOwner(
					spirallineFormCallback.probe.stageOfInterest,
					spiralline_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.SpiralLineGrid)
					if !ok {
						log.Fatalln("Source of SpiralLineGrid.SpiralLines []*SpiralLine, is not an SpiralLineGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.SpiralLineGrid
			for _spirallinegrid := range *models.GetGongstructInstancesSet[models.SpiralLineGrid](spirallineFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _spirallinegrid.GetName() == newSourceName.GetName() {
					newSource = _spirallinegrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of SpiralLineGrid.SpiralLines []*SpiralLine, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.SpiralLines = append(newSource.SpiralLines, spiralline_)
		}
	}

	// manage the suppress operation
	if spirallineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralline_.Unstage(spirallineFormCallback.probe.stageOfInterest)
	}

	spirallineFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralLine](
		spirallineFormCallback.probe,
	)
	spirallineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spirallineFormCallback.CreationMode || spirallineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spirallineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spirallineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralLineFormCallback(
			nil,
			spirallineFormCallback.probe,
			newFormGroup,
		)
		spiralline := new(models.SpiralLine)
		FillUpForm(spiralline, newFormGroup, spirallineFormCallback.probe)
		spirallineFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spirallineFormCallback.probe)
}
func __gong__New__SpiralLineGridFormCallback(
	spirallinegrid *models.SpiralLineGrid,
	probe *Probe,
	formGroup *table.FormGroup,
) (spirallinegridFormCallback *SpiralLineGridFormCallback) {
	spirallinegridFormCallback = new(SpiralLineGridFormCallback)
	spirallinegridFormCallback.probe = probe
	spirallinegridFormCallback.spirallinegrid = spirallinegrid
	spirallinegridFormCallback.formGroup = formGroup

	spirallinegridFormCallback.CreationMode = (spirallinegrid == nil)

	return
}

type SpiralLineGridFormCallback struct {
	spirallinegrid *models.SpiralLineGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spirallinegridFormCallback *SpiralLineGridFormCallback) OnSave() {

	// log.Println("SpiralLineGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spirallinegridFormCallback.probe.formStage.Checkout()

	if spirallinegridFormCallback.spirallinegrid == nil {
		spirallinegridFormCallback.spirallinegrid = new(models.SpiralLineGrid).Stage(spirallinegridFormCallback.probe.stageOfInterest)
	}
	spirallinegrid_ := spirallinegridFormCallback.spirallinegrid
	_ = spirallinegrid_

	for _, formDiv := range spirallinegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spirallinegrid_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spirallinegrid_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spirallinegrid_.ShapeCategory), spirallinegridFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if spirallinegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spirallinegrid_.Unstage(spirallinegridFormCallback.probe.stageOfInterest)
	}

	spirallinegridFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralLineGrid](
		spirallinegridFormCallback.probe,
	)
	spirallinegridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spirallinegridFormCallback.CreationMode || spirallinegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spirallinegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spirallinegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralLineGridFormCallback(
			nil,
			spirallinegridFormCallback.probe,
			newFormGroup,
		)
		spirallinegrid := new(models.SpiralLineGrid)
		FillUpForm(spirallinegrid, newFormGroup, spirallinegridFormCallback.probe)
		spirallinegridFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spirallinegridFormCallback.probe)
}
func __gong__New__SpiralOriginFormCallback(
	spiralorigin *models.SpiralOrigin,
	probe *Probe,
	formGroup *table.FormGroup,
) (spiraloriginFormCallback *SpiralOriginFormCallback) {
	spiraloriginFormCallback = new(SpiralOriginFormCallback)
	spiraloriginFormCallback.probe = probe
	spiraloriginFormCallback.spiralorigin = spiralorigin
	spiraloriginFormCallback.formGroup = formGroup

	spiraloriginFormCallback.CreationMode = (spiralorigin == nil)

	return
}

type SpiralOriginFormCallback struct {
	spiralorigin *models.SpiralOrigin

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spiraloriginFormCallback *SpiralOriginFormCallback) OnSave() {

	// log.Println("SpiralOriginFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spiraloriginFormCallback.probe.formStage.Checkout()

	if spiraloriginFormCallback.spiralorigin == nil {
		spiraloriginFormCallback.spiralorigin = new(models.SpiralOrigin).Stage(spiraloriginFormCallback.probe.stageOfInterest)
	}
	spiralorigin_ := spiraloriginFormCallback.spiralorigin
	_ = spiralorigin_

	for _, formDiv := range spiraloriginFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spiralorigin_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(spiralorigin_.IsDisplayed), formDiv)
		case "ShapeCategory":
			FormDivSelectFieldToField(&(spiralorigin_.ShapeCategory), spiraloriginFormCallback.probe.stageOfInterest, formDiv)
		case "Color":
			FormDivBasicFieldToField(&(spiralorigin_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(spiralorigin_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(spiralorigin_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(spiralorigin_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(spiralorigin_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(spiralorigin_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(spiralorigin_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(spiralorigin_.Transform), formDiv)
		}
	}

	// manage the suppress operation
	if spiraloriginFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralorigin_.Unstage(spiraloriginFormCallback.probe.stageOfInterest)
	}

	spiraloriginFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralOrigin](
		spiraloriginFormCallback.probe,
	)
	spiraloriginFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiraloriginFormCallback.CreationMode || spiraloriginFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiraloriginFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spiraloriginFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SpiralOriginFormCallback(
			nil,
			spiraloriginFormCallback.probe,
			newFormGroup,
		)
		spiralorigin := new(models.SpiralOrigin)
		FillUpForm(spiralorigin, newFormGroup, spiraloriginFormCallback.probe)
		spiraloriginFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spiraloriginFormCallback.probe)
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

	// log.Println("SpiralRhombusFormCallback, OnSave")

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
		case "X_r0":
			FormDivBasicFieldToField(&(spiralrhombus_.X_r0), formDiv)
		case "Y_r0":
			FormDivBasicFieldToField(&(spiralrhombus_.Y_r0), formDiv)
		case "X_r1":
			FormDivBasicFieldToField(&(spiralrhombus_.X_r1), formDiv)
		case "Y_r1":
			FormDivBasicFieldToField(&(spiralrhombus_.Y_r1), formDiv)
		case "X_r2":
			FormDivBasicFieldToField(&(spiralrhombus_.X_r2), formDiv)
		case "Y_r2":
			FormDivBasicFieldToField(&(spiralrhombus_.Y_r2), formDiv)
		case "X_r3":
			FormDivBasicFieldToField(&(spiralrhombus_.X_r3), formDiv)
		case "Y_r3":
			FormDivBasicFieldToField(&(spiralrhombus_.Y_r3), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(spiralrhombus_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(spiralrhombus_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(spiralrhombus_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(spiralrhombus_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(spiralrhombus_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(spiralrhombus_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(spiralrhombus_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(spiralrhombus_.Transform), formDiv)
		case "SpiralRhombusGrid:SpiralRhombuses":
			// WARNING : this form deals with the N-N association "SpiralRhombusGrid.SpiralRhombuses []*SpiralRhombus" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SpiralRhombus). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.SpiralRhombusGrid
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "SpiralRhombusGrid"
				rf.Fieldname = "SpiralRhombuses"
				formerAssociationSource := models.GetReverseFieldOwner(
					spiralrhombusFormCallback.probe.stageOfInterest,
					spiralrhombus_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.SpiralRhombusGrid)
					if !ok {
						log.Fatalln("Source of SpiralRhombusGrid.SpiralRhombuses []*SpiralRhombus, is not an SpiralRhombusGrid instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.SpiralRhombusGrid
			for _spiralrhombusgrid := range *models.GetGongstructInstancesSet[models.SpiralRhombusGrid](spiralrhombusFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _spiralrhombusgrid.GetName() == newSourceName.GetName() {
					newSource = _spiralrhombusgrid // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of SpiralRhombusGrid.SpiralRhombuses []*SpiralRhombus, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.SpiralRhombuses = append(newSource.SpiralRhombuses, spiralrhombus_)
		}
	}

	// manage the suppress operation
	if spiralrhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombus_.Unstage(spiralrhombusFormCallback.probe.stageOfInterest)
	}

	spiralrhombusFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralRhombus](
		spiralrhombusFormCallback.probe,
	)
	spiralrhombusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralrhombusFormCallback.CreationMode || spiralrhombusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(spiralrhombusFormCallback.probe)
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

	// log.Println("SpiralRhombusGridFormCallback, OnSave")

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
		}
	}

	// manage the suppress operation
	if spiralrhombusgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombusgrid_.Unstage(spiralrhombusgridFormCallback.probe.stageOfInterest)
	}

	spiralrhombusgridFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SpiralRhombusGrid](
		spiralrhombusgridFormCallback.probe,
	)
	spiralrhombusgridFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spiralrhombusgridFormCallback.CreationMode || spiralrhombusgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spiralrhombusgridFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(spiralrhombusgridFormCallback.probe)
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

	// log.Println("VerticalAxisFormCallback, OnSave")

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
	updateAndCommitTable[models.VerticalAxis](
		verticalaxisFormCallback.probe,
	)
	verticalaxisFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if verticalaxisFormCallback.CreationMode || verticalaxisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		verticalaxisFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(verticalaxisFormCallback.probe)
}
