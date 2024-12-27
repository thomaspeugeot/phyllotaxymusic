package models

import "log"

func (p *Parameter) computeThemeNotesShapes(bezierGrid *BezierGrid, g *CircleGrid) {

	g.Circles = g.Circles[:0]

	beatWidth := p.RotatedAxis.Length / float64(p.NbOfBeatsInTheme)

	//
	// nb of measures to jump before the first bezier
	//
	nbMeasureToJump := int(bezierGrid.Beziers[0].StartX/beatWidth + 0.5)

	for beatNb := range p.NbOfBeatsInTheme {

		c := new(Circle)
		c.BeatNb = beatNb
		*c = *g.Reference

		c.isKept = p.IsNotePlayed(beatNb)

		g.Circles = append(g.Circles, c)

		c.CenterX = float64(beatNb+nbMeasureToJump) * beatWidth

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

		pitchAdjusment := 0

		delta := ratio - float64(c.Pitch)
		if delta > 0 {
			pitchAdjusment = 1
		} else {
			pitchAdjusment = -1
		}

		//
		// set pitch on minor or major
		//
		relativePitch := c.Pitch % 12
		if p.IsMinor {
			switch relativePitch {
			case 1, 4, 6:
				c.Pitch += pitchAdjusment
			case 9:
				c.Pitch -= 1
			case 10:
				c.Pitch += 1
			}
		} else {
			switch relativePitch {
			case 1, 3, 6, 8, 10:
				c.Pitch += pitchAdjusment
			}
		}

		c.CenterY = float64(c.Pitch) * pitchHeight
	}
}
