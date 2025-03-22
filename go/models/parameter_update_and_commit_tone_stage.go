package models

import (
	"fmt"

	gongtone_models "github.com/fullstack-lang/gong/lib/tone/go/models"
)

func (p *Parameter) UpdateAndCommitToneStage() {

	p.toneStage.Reset()

	// we consider that the scale start at C3
	keyboard := gongtone_models.GeneratePianoNotes()

	keyboard = keyboard[12:]

	map_Freqs := make(map[string]*gongtone_models.Freqency)

	// log.Println("speed", parameter.Speed)

	// note.Info = fmt.Sprintf("%40d", i)
	if p.FirstVoiceNotes.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.FirstVoiceNotes, p.toneStage)
	}
	if p.FirstVoiceNotesShiftedRight.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.FirstVoiceNotesShiftedRight, p.toneStage)
	}
	if p.SecondVoiceNotes.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.SecondVoiceNotes, p.toneStage)
	}
	if p.SecondVoiceNotesShiftedRight.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.SecondVoiceNotesShiftedRight, p.toneStage)
	}

	player := new(gongtone_models.Player).Stage(p.toneStage)
	player.OnDI = func(player *gongtone_models.Player) error {

		// notify the cursor
		value := player.Status == gongtone_models.PLAYING
		p.cursor.PlayCursor(p.cursorStage, value)
		return nil
	}

	p.toneStage.Commit()

}

// generateNotesFromCircleGrid processes a set of circles to create corresponding notes,
// assigning frequencies and computing their durations. It uses information from the given
// keyboard mapping, frequency map, and circle grid. Newly created Freqency and Note
// objects are staged in the provided gongtoneStage.
func (p *Parameter) generateNotesFromCircleGrid(
	keyboard []string,
	map_Freqs map[string]*gongtone_models.Freqency,
	circleGrid *CircleGrid,
	gongtoneStage *gongtone_models.StageStruct) {

	beatLength := p.RotatedAxis.Length / float64(p.NbOfBeatsInTheme)

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

		note.Start = (c.CenterX / beatLength) / p.BeatsPerSecond

		note.Velocity = p.Level
	}

	// compute duration according to skipped notes
	for i, c := range circleGrid.Circles {

		if c.note == nil {
			continue
		}

		c.note.Duration = 1 / p.BeatsPerSecond

		remainingCircles := circleGrid.Circles[i+1:]

		for _, _c := range remainingCircles {
			if _c.isKept {
				break
			}
			c.note.Duration += 1 / p.BeatsPerSecond
		}
	}
}
