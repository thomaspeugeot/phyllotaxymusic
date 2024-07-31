package models

import (
	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
)

func (parameter *Parameter) GenerateTone(gongtoneStage *gongtone_models.StageStruct) {

	gongtoneStage.Reset()

	// we consider that the scale start at C3
	keyboard := gongtone_models.GeneratePianoNotes()

	keyboard = keyboard[12:]

	map_Freqs := make(map[string]*gongtone_models.Freqency)

	for i, c := range parameter.FirstVoiceNotes.Circles {
		freqNotation := keyboard[c.Pitch]

		freq, ok := map_Freqs[freqNotation]

		if !ok {
			freq = new(gongtone_models.Freqency).Stage(gongtoneStage)
			freq.Name = freqNotation
		}
		note := new(gongtone_models.Note).Stage(gongtoneStage)
		note.Frequencies = append(note.Frequencies, freq)
		note.Start = float64(i)
		note.Duration = 1
		note.Velocity = 1
	}

	gongtoneStage.Commit()

}
