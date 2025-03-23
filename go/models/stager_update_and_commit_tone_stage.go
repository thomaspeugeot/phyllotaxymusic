package models

import (
	"fmt"

	gongtone_models "github.com/fullstack-lang/gong/lib/tone/go/models"
)

func (stager *Stager) UpdateAndCommitToneStage() {

	stager.toneStage.Reset()
	parameter := stager.parameter

	// we consider that the scale start at C3
	keyboard := gongtone_models.GeneratePianoNotes()

	keyboard = keyboard[12:]

	map_Freqs := make(map[string]*gongtone_models.Freqency)

	// log.Println("speed", stager.Speed)

	// note.Info = fmt.Sprintf("%40d", i)
	if parameter.FirstVoiceNotes.IsDisplayed {
		stager.generateNotesFromCircleGrid(keyboard, map_Freqs, parameter.FirstVoiceNotes, stager.toneStage)
	}
	if parameter.FirstVoiceNotesShiftedRight.IsDisplayed {
		stager.generateNotesFromCircleGrid(keyboard, map_Freqs, parameter.FirstVoiceNotesShiftedRight, stager.toneStage)
	}
	if parameter.SecondVoiceNotes.IsDisplayed {
		stager.generateNotesFromCircleGrid(keyboard, map_Freqs, parameter.SecondVoiceNotes, stager.toneStage)
	}
	if parameter.SecondVoiceNotesShiftedRight.IsDisplayed {
		stager.generateNotesFromCircleGrid(keyboard, map_Freqs, parameter.SecondVoiceNotesShiftedRight, stager.toneStage)
	}

	player := new(gongtone_models.Player).Stage(stager.toneStage)
	player.OnDI = func(player *gongtone_models.Player) error {

		// notify the cursor
		value := player.Status == gongtone_models.PLAYING
		stager.cursor.PlayCursor(stager.cursorStage, value)
		return nil
	}

	stager.toneStage.Commit()

}

// generateNotesFromCircleGrid processes a set of circles to create corresponding notes,
// assigning frequencies and computing their durations. It uses information from the given
// keyboard mapping, frequency map, and circle grid. Newly created Freqency and Note
// objects are staged in the provided gongtoneStage.
func (stager *Stager) generateNotesFromCircleGrid(
	keyboard []string,
	map_Freqs map[string]*gongtone_models.Freqency,
	circleGrid *CircleGrid,
	gongtoneStage *gongtone_models.StageStruct) {

	parameter := stager.parameter

	beatLength := parameter.RotatedAxis.Length / float64(parameter.NbOfBeatsInTheme)

	for idx, c := range circleGrid.Circles {

		if !c.isKept {
			continue
		}

		freqNotation := keyboard[c.Pitch]

		freq, ok := map_Freqs[freqNotation]

		if !ok {
			freq = new(gongtone_models.Freqency).Stage(gongtoneStage)
			freq.Name = freqNotation
		}

		note := new(gongtone_models.Note).Stage(gongtoneStage)
		note.Name = fmt.Sprintf("Note %2d", idx)
		c.note = note
		note.Frequencies = append(note.Frequencies, freq)

		note.Start = (c.CenterX / beatLength) / parameter.BeatsPerSecond

		note.Velocity = parameter.Level
	}

	// compute duration according to skipped notes
	for i, c := range circleGrid.Circles {

		if c.note == nil {
			continue
		}

		c.note.Duration = 1 / parameter.BeatsPerSecond

		remainingCircles := circleGrid.Circles[i+1:]

		for _, _c := range remainingCircles {
			if _c.isKept {
				break
			}
			c.note.Duration += 1 / parameter.BeatsPerSecond
		}
	}
}
