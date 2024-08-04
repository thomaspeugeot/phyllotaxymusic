package models

import "log"

func (p *Parameter) computeFirstVoiceNotes() {

	g := p.FirstVoiceNotes
	g.Circles = g.Circles[:0]

	for i := range p.NbMeasureLinesPerCurve {
		c := new(Circle)
		*c = *g.Reference

		g.Circles = append(g.Circles, c)

		c.CenterX = float64(i) * p.RotatedAxis.Length / float64(p.NbMeasureLinesPerCurve)

		//
		// compute which bezier is concerned
		//
		for _, b := range p.FirstVoice.Beziers {
			if b.EndX > c.CenterX {
				var err error
				c.CenterY, err = b.ComputeYFromX(c.CenterX)
				if err != nil {
					log.Panicf("Problem with bezier y from x")
				}
				break
			}
		}

		// interpolate to the nearest pitch
		pitchHeight := p.PitchHeight * p.SideLength
		c.Pitch = int((c.CenterY - 0.5) / pitchHeight)
		c.CenterY = float64(c.Pitch) * pitchHeight
	}
}
