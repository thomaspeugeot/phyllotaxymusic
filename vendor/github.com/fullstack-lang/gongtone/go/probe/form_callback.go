// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtone/go/models"
	"github.com/fullstack-lang/gongtone/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__FreqencyFormCallback(
	freqency *models.Freqency,
	probe *Probe,
	formGroup *table.FormGroup,
) (freqencyFormCallback *FreqencyFormCallback) {
	freqencyFormCallback = new(FreqencyFormCallback)
	freqencyFormCallback.probe = probe
	freqencyFormCallback.freqency = freqency
	freqencyFormCallback.formGroup = formGroup

	freqencyFormCallback.CreationMode = (freqency == nil)

	return
}

type FreqencyFormCallback struct {
	freqency *models.Freqency

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (freqencyFormCallback *FreqencyFormCallback) OnSave() {

	log.Println("FreqencyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	freqencyFormCallback.probe.formStage.Checkout()

	if freqencyFormCallback.freqency == nil {
		freqencyFormCallback.freqency = new(models.Freqency).Stage(freqencyFormCallback.probe.stageOfInterest)
	}
	freqency_ := freqencyFormCallback.freqency
	_ = freqency_

	for _, formDiv := range freqencyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(freqency_.Name), formDiv)
		case "Note:Frequencies":
			// we need to retrieve the field owner before the change
			var pastNoteOwner *models.Note
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Frequencies"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				freqencyFormCallback.probe.stageOfInterest,
				freqencyFormCallback.probe.backRepoOfInterest,
				freqency_,
				&rf)

			if reverseFieldOwner != nil {
				pastNoteOwner = reverseFieldOwner.(*models.Note)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNoteOwner != nil {
					idx := slices.Index(pastNoteOwner.Frequencies, freqency_)
					pastNoteOwner.Frequencies = slices.Delete(pastNoteOwner.Frequencies, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _note := range *models.GetGongstructInstancesSet[models.Note](freqencyFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _note.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNoteOwner := _note // we have a match
						if pastNoteOwner != nil {
							if newNoteOwner != pastNoteOwner {
								idx := slices.Index(pastNoteOwner.Frequencies, freqency_)
								pastNoteOwner.Frequencies = slices.Delete(pastNoteOwner.Frequencies, idx, idx+1)
								newNoteOwner.Frequencies = append(newNoteOwner.Frequencies, freqency_)
							}
						} else {
							newNoteOwner.Frequencies = append(newNoteOwner.Frequencies, freqency_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if freqencyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		freqency_.Unstage(freqencyFormCallback.probe.stageOfInterest)
	}

	freqencyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Freqency](
		freqencyFormCallback.probe,
	)
	freqencyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if freqencyFormCallback.CreationMode || freqencyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		freqencyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(freqencyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FreqencyFormCallback(
			nil,
			freqencyFormCallback.probe,
			newFormGroup,
		)
		freqency := new(models.Freqency)
		FillUpForm(freqency, newFormGroup, freqencyFormCallback.probe)
		freqencyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(freqencyFormCallback.probe)
}
func __gong__New__NoteFormCallback(
	note *models.Note,
	probe *Probe,
	formGroup *table.FormGroup,
) (noteFormCallback *NoteFormCallback) {
	noteFormCallback = new(NoteFormCallback)
	noteFormCallback.probe = probe
	noteFormCallback.note = note
	noteFormCallback.formGroup = formGroup

	noteFormCallback.CreationMode = (note == nil)

	return
}

type NoteFormCallback struct {
	note *models.Note

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (noteFormCallback *NoteFormCallback) OnSave() {

	log.Println("NoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteFormCallback.probe.formStage.Checkout()

	if noteFormCallback.note == nil {
		noteFormCallback.note = new(models.Note).Stage(noteFormCallback.probe.stageOfInterest)
	}
	note_ := noteFormCallback.note
	_ = note_

	for _, formDiv := range noteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(note_.Name), formDiv)
		case "Start":
			FormDivBasicFieldToField(&(note_.Start), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(note_.Duration), formDiv)
		case "Velocity":
			FormDivBasicFieldToField(&(note_.Velocity), formDiv)
		case "Info":
			FormDivBasicFieldToField(&(note_.Info), formDiv)
		}
	}

	// manage the suppress operation
	if noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_.Unstage(noteFormCallback.probe.stageOfInterest)
	}

	noteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Note](
		noteFormCallback.probe,
	)
	noteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteFormCallback.CreationMode || noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(noteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			noteFormCallback.probe,
			newFormGroup,
		)
		note := new(models.Note)
		FillUpForm(note, newFormGroup, noteFormCallback.probe)
		noteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteFormCallback.probe)
}
