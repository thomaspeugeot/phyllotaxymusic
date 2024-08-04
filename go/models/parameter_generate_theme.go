package models

import (
	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
)

func (p *Parameter) GenerateTone(gongtoneStage *gongtone_models.StageStruct) {

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

	gongtoneStage.Commit()

}

func (p *Parameter) generateNotesFromCircleGrid(
	keyboard []string,
	map_Freqs map[string]*gongtone_models.Freqency,
	circleGrid *CircleGrid,
	gongtoneStage *gongtone_models.StageStruct) {
	for _, c := range circleGrid.Circles {
		freqNotation := keyboard[c.Pitch]

		freq, ok := map_Freqs[freqNotation]

		if !ok {
			freq = new(gongtone_models.Freqency).Stage(gongtoneStage)
			freq.Name = freqNotation
		}
		unitMeasureLength := p.RotatedAxis.Length / float64(p.NbMeasureLinesPerCurve)
		note := new(gongtone_models.Note).Stage(gongtoneStage)
		note.Frequencies = append(note.Frequencies, freq)
		note.Start = c.CenterX / unitMeasureLength / p.Speed
		note.Duration = 1 / p.Speed
		note.Velocity = p.Level

	}
}
