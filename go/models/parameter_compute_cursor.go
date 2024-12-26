package models

import "log"

func (p *Parameter) computeCursorSVGCoords() {

	firstCircle := p.FirstVoiceNotes.Circles[0]
	beatLines := p.PitchLines.Axiss

	p.cursor.StartX = p.OriginX + firstCircle.CenterX
	p.cursor.Y1 = p.OriginY - beatLines[0].CenterY
	p.cursor.Y2 = p.OriginY - beatLines[len(p.PitchLines.Axiss)-1].CenterY
	p.cursor.EndX = p.cursor.StartX + p.SideLength*2
	p.cursor.DurationSeconds = float64(p.NbOfBeatsInTheme) / p.BeatsPerSecond

	log.Printf("%+v", p.cursor)
}
