package models

import "log"

func (p *Parameter) computeVoiceNotes(bezierGrid *BezierGrid, g *CircleGrid) {

	g.Circles = g.Circles[:0]

	measureLength := p.RotatedAxis.Length / float64(p.NbMeasureLinesPerCurve)

	//
	// nb of measures to jump before the first bezier
	//
	nbMeasureToJump := int(bezierGrid.Beziers[0].StartX/measureLength + 0.5)

	for i := range p.NbMeasureLinesPerCurve {
		c := new(Circle)
		*c = *g.Reference

		g.Circles = append(g.Circles, c)

		c.CenterX = float64(i+nbMeasureToJump) * measureLength

		//
		// compute which bezier is concerned
		//
		for _, b := range bezierGrid.Beziers {
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
		ratio := c.CenterY / pitchHeight
		c.Pitch = int(ratio + 0.5)
		c.CenterY = float64(c.Pitch) * pitchHeight
	}
}
