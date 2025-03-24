package models

import (
	"log"
	"math"
)

// UpdateAndCommitCursorStage computes cursor
// StartX, EndX and DurationSeconds according to
// parameter displayed voices notes amon
//
// FirstVoiceNotes,
// SecondVoiceNotes,
// FirstVoiceNotesShiftedRight
// SecondVoiceNotesShiftedRight
func (stager *Stager) UpdateAndCommitCursorStage() {

	if stager.cursor == nil {
		log.Fatal("no cursor")
	}
	parameter := stager.parameter

	themeDuration := float64(parameter.NbOfBeatsInTheme) / parameter.BeatsPerSecond
	themVisualLenght := parameter.RotatedAxis.Length

	var voices []*CircleGrid
	voices = append(voices,
		parameter.FirstVoiceNotes,
		parameter.FirstVoiceNotesShiftedRight,
		parameter.SecondVoiceNotes,
		parameter.SecondVoiceNotesShiftedRight)

	pitchLines := parameter.PitchLines.Axiss
	stager.cursor.Y1 = parameter.OriginY - pitchLines[0].CenterY
	stager.cursor.Y2 = parameter.OriginY - pitchLines[len(pitchLines)-1].CenterY

	_, x2 := computeStartEnd(voices)

	stager.cursor.StartX,
		stager.cursor.EndX =
		parameter.OriginX,
		x2+parameter.OriginX+(themVisualLenght)/float64(parameter.NbOfBeatsInTheme)
	stager.cursor.DurationSeconds = ((stager.cursor.EndX - stager.cursor.StartX) / themVisualLenght) * themeDuration

	stager.cursorStage.Commit()
}

// computeStartEnd inspects each displayed CircleGrid and returns two values:
//   - startX: the smallest CenterX among the *first circles* of all displayed grids
//   - endX:   the largest CenterX among the *last circles* of all displayed grids
//
// If no displayed CircleGrid contains circles, both startX and endX are returned as 0.
func computeStartEnd(circleGrids []*CircleGrid) (startX, endX float64) {
	// Initialize to extreme values so they can be correctly updated.
	startX = math.Inf(1) // positive infinity
	endX = math.Inf(-1)  // negative infinity

	// Loop over all CircleGrid pointers.
	for _, grid := range circleGrids {
		// Skip grids that are not displayed or have no circles.
		if !grid.IsDisplayed || len(grid.Circles) == 0 {
			continue
		}

		// Identify the first and last circles in the current grid.
		firstCircle := grid.Circles[0]
		lastCircle := grid.Circles[len(grid.Circles)-1]

		// Update the smallest CenterX found so far (startX).
		if firstCircle.CenterX < startX {
			startX = firstCircle.CenterX
		}

		// Update the largest CenterX found so far (endX).
		if lastCircle.CenterX > endX {
			endX = lastCircle.CenterX
		}
	}

	// If startX/endX were never updated, it means no valid circles were found.
	if startX == math.Inf(1) {
		startX = 0
	}
	if endX == math.Inf(-1) {
		endX = 0
	}

	return startX, endX
}
