package models

import (
	"fmt"

	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
)

func (p *Parameter) GenerateNotes(gongtoneStage *gongtone_models.StageStruct) {

	gongtoneStage.Reset()

	// we consider that the scale start at C3
	keyboard := gongtone_models.GeneratePianoNotes()

	keyboard = keyboard[12:]

	map_Freqs := make(map[string]*gongtone_models.Freqency)

	// log.Println("speed", parameter.Speed)

	// note.Info = fmt.Sprintf("%40d", i)
	if p.FirstVoiceNotes.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.FirstVoiceNotes, gongtoneStage)
	}
	if p.FirstVoiceNotesShiftedRight.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.FirstVoiceNotesShiftedRight, gongtoneStage)
	}
	if p.SecondVoiceNotes.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.SecondVoiceNotes, gongtoneStage)
	}
	if p.SecondVoiceNotesShiftedRight.IsDisplayed {
		p.generateNotesFromCircleGrid(keyboard, map_Freqs, p.SecondVoiceNotesShiftedRight, gongtoneStage)
	}

	player := new(gongtone_models.Player).Stage(gongtoneStage)
	player.OnDI = func(p *gongtone_models.Player) error {
		fmt.Printf("Injecting dependencies for player: %s\n", p.Name)
		// Perform any dependency injection logic here
		return nil
	}

	gongtoneStage.Commit()

}

func (p *Parameter) generateNotesFromCircleGrid(
	keyboard []string,
	map_Freqs map[string]*gongtone_models.Freqency,
	circleGrid *CircleGrid,
	gongtoneStage *gongtone_models.StageStruct) {

	unitMeasureLength := p.RotatedAxis.Length / float64(p.NbBeatLinesPerCurve)

	for _, c := range circleGrid.Circles {

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
		c.note = note
		note.Frequencies = append(note.Frequencies, freq)

		note.Start = (c.CenterX / unitMeasureLength) / p.Speed

		note.Velocity = p.Level
	}

	// compute duration according to skipped notes
	for i, c := range circleGrid.Circles {

		if c.note == nil {
			continue
		}

		c.note.Duration = 1 / p.Speed

		for _, _c := range circleGrid.Circles[i:] {
			if !_c.isKept {
				c.note.Duration += 1 / p.Speed
			}
		}
	}
}
