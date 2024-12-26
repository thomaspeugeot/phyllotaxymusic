package models

import "log"

func (p *Parameter) computeCursorSVGCoords() {

	duration := float64(p.NbOfBeatsInTheme) / p.BeatsPerSecond

	firstCircleOfFirstVoice := p.FirstVoiceNotes.Circles[0]
	beatLines := p.PitchLines.Axiss

	firstCircleOfSecondVoice := p.SecondVoiceNotes.Circles[0]
	xShitToSecondVoice := firstCircleOfSecondVoice.CenterX - firstCircleOfFirstVoice.CenterX

	if firstCircleOfFirstVoice == nil {
		log.Fatalln("no notes !")
	}

	if p.cursor == nil {
		log.Fatal("no cursor")
	}

	p.cursor.StartX = p.OriginX + firstCircleOfFirstVoice.CenterX
	p.cursor.Y1 = p.OriginY - beatLines[0].CenterY
	p.cursor.Y2 = p.OriginY - beatLines[len(p.PitchLines.Axiss)-1].CenterY
	if p.SecondVoiceNotesShiftedRight.IsDisplayed {
		p.cursor.EndX = p.cursor.StartX + p.RotatedAxis.Length*2 + xShitToSecondVoice
		p.cursor.DurationSeconds = 2.0*float64(p.NbOfBeatsInTheme)/p.BeatsPerSecond +
			((firstCircleOfSecondVoice.CenterX-firstCircleOfFirstVoice.CenterX)/p.RotatedAxis.Length)*duration
	} else {
		if p.FirstVoiceNotesShiftedRight.IsDisplayed {
			p.cursor.EndX = p.cursor.StartX + p.RotatedAxis.Length*2
			p.cursor.DurationSeconds = 2.0 * duration
		} else {
			p.cursor.EndX = p.cursor.StartX + p.RotatedAxis.Length
			p.cursor.DurationSeconds = duration
		}
	}

	log.Printf("%+v", p.cursor)
}
