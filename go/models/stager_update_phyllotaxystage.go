package models

func (stager *Stager) UpdatePhyllotaxyStage() {

	p := stager.parameter

	if p.Z < p.M+p.N+1 {
		p.Z = p.M + p.N + 1
	}

	p.Shapes = p.Shapes[:0]

	p.Shapes = append(p.Shapes, p.HorizontalAxis)
	p.Shapes = append(p.Shapes, p.VerticalAxis)
	p.Shapes = append(p.Shapes, p.SpiralOrigin)

	p.ComputeInitialRhombus()
	p.Shapes = append(p.Shapes, p.InitialRhombus)

	p.ComputeInitialCircle()
	p.Shapes = append(p.Shapes, p.InitialCircle)

	p.ComputeInitialRhombusGrid(stager.phyllotaxymusicStage)
	p.Shapes = append(p.Shapes, p.InitialRhombusGrid)

	p.ComputeInitialCircleGrid(stager.phyllotaxymusicStage)
	p.Shapes = append(p.Shapes, p.InitialCircleGrid)

	p.ComputeInitialAxis()
	p.Shapes = append(p.Shapes, p.InitialAxis)

	p.computeRotatedAxis()
	p.Shapes = append(p.Shapes, p.RotatedAxis)

	p.computeRotatedRhombus()
	p.Shapes = append(p.Shapes, p.RotatedRhombus)

	p.computeRotatedRhombusGrid(stager.phyllotaxymusicStage)
	p.Shapes = append(p.Shapes, p.RotatedRhombusGrid)

	p.computeRotatedCircleGrid(stager.phyllotaxymusicStage)
	p.Shapes = append(p.Shapes, p.RotatedCircleGrid)

	p.ComputeNextRhombus()
	p.Shapes = append(p.Shapes, p.NextRhombus)

	p.ComputeNextCircle()
	p.Shapes = append(p.Shapes, p.NextCircle)

	p.ComputeGrowingRhombusGrid()
	p.Shapes = append(p.Shapes, p.GrowingRhombusGridSeed)
	p.Shapes = append(p.Shapes, p.GrowingRhombusGrid)

	p.ComputeGrowingCircleGrid()
	p.Shapes = append(p.Shapes, p.GrowingCircleGridSeed)
	p.Shapes = append(p.Shapes, p.GrowingCircleGrid)

	p.ComputeGrowingCircleGridLeft()
	p.Shapes = append(p.Shapes, p.GrowingCircleGridLeft)

	p.computeConstructionAxis()
	p.Shapes = append(p.Shapes, p.ConstructionAxis)

	p.computeConstructionCircle()
	p.Shapes = append(p.Shapes, p.ConstructionCircle)

	p.computeConstructionAxisGrid()
	p.Shapes = append(p.Shapes, p.ConstructionAxisGrid)

	p.computeConstructionCircleGrid()
	p.Shapes = append(p.Shapes, p.ConstructionCircleGrid)

	p.ComputeGrowthCurveSegment()
	p.Shapes = append(p.Shapes, p.GrowthCurveSeed)

	p.ComputeGrowthCurve()
	p.Shapes = append(p.Shapes, p.GrowthCurve)

	p.ComputeSpiralRhombusSeed()
	p.Shapes = append(p.Shapes, p.SpiralRhombusGridSeed)

	p.ComputeSpiralRhombusGrid()
	p.Shapes = append(p.Shapes, p.SpiralRhombusGrid)

	p.GrowthCurveShiftedRight.Move(p.GrowthCurveShiftedRightSeed, p.GrowthCurve,
		p.RotatedAxis.Length, 0)
	p.Shapes = append(p.Shapes, p.GrowthCurveShiftedRight)

	p.GrowthCurveNext.Move(p.GrowthCurveNextSeed, p.GrowthCurve,
		p.NextCircle.CenterX, p.NextCircle.CenterY)
	p.Shapes = append(p.Shapes, p.GrowthCurveNext)

	p.GrowthCurveNextShiftedRight.Move(p.GrowthCurveNextShiftedRightSeed, p.GrowthCurve,
		p.NextCircle.CenterX+p.RotatedAxis.Length, p.NextCircle.CenterY)
	p.Shapes = append(p.Shapes, p.GrowthCurveNextShiftedRight)

	p.GrowthCurveStack.BezierGrids = p.GrowthCurveStack.BezierGrids[:0]
	for s := range p.StackHeight {

		for i := range p.StackWidth {
			g := new(BezierGrid)
			p.GrowthCurveStack.BezierGrids = append(p.GrowthCurveStack.BezierGrids, g)

			shift := float64(s) * p.NextCircle.CenterX
			for shift > 0 {
				shift -= p.RotatedAxis.Length
			}
			g.Move(p.GrowthCurveNextSeed, p.GrowthCurve,
				shift+float64(i)*p.RotatedAxis.Length,
				float64(s)*p.NextCircle.CenterY,
			)
		}
	}
	p.Shapes = append(p.Shapes, p.GrowthCurveStack)

	p.ComputeSpiralCircleSeed()
	p.Shapes = append(p.Shapes, p.SpiralCircleSeed)
	p.ComputeSpiralCircleGrid()
	p.Shapes = append(p.Shapes, p.SpiralCircleGrid)
	p.ComputeSpiralCircleFullGrid()
	p.Shapes = append(p.Shapes, p.SpiralCircleFullGrid)
	p.computeSpiralConstructionOuterLineSeed()
	p.Shapes = append(p.Shapes, p.SpiralConstructionOuterLineSeed)
	p.computeSpiralConstructionInnerLineSeed()
	p.Shapes = append(p.Shapes, p.SpiralConstructionInnerLineSeed)

	p.computeSpiralConstructionOuterLineGrid()
	p.Shapes = append(p.Shapes, p.SpiralConstructionOuterLineGrid)
	p.computeSpiralConstructionInnerLineGrid()
	p.Shapes = append(p.Shapes, p.SpiralConstructionInnerLineGrid)

	p.computeSpiralConstructionOuterLineFullGrid()
	p.Shapes = append(p.Shapes, p.SpiralConstructionOuterLineFullGrid)

	p.ComputeSpiralConstructionCircleGrid()
	p.Shapes = append(p.Shapes, p.SpiralConstructionCircleGrid)
	p.ComputeSpiralBezierSeed()
	p.Shapes = append(p.Shapes, p.SpiralBezierSeed)
	p.ComputeSpiralBezierGrid()
	p.Shapes = append(p.Shapes, p.SpiralBezierGrid)
	p.ComputeSpiralBezierFullGrid()
	p.Shapes = append(p.Shapes, p.SpiralBezierFullGrid)

	p.ComputeFrontCurveStacks(stager.phyllotaxymusicStage)
	p.Shapes = append(p.Shapes, p.FrontCurveStack)

	//
	// MUSIC MAESTRO
	//

	p.ComputeFKey()
	p.Shapes = append(p.Shapes, p.Fkey)

	p.ComputePitchLines()
	p.Shapes = append(p.Shapes, p.PitchLines)

	p.ComputeBeatLines()
	p.Shapes = append(p.Shapes, p.BeatLines)

	p.FirstVoice.Move(p.FirstVoice.Reference, p.GrowthCurve,
		float64(p.FirstVoiceShiftX)*p.SideLength,
		float64(p.FirstVoiceShiftY)*p.SideLength)
	p.Shapes = append(p.Shapes, p.FirstVoice)

	p.FirstVoiceShiftedRigth.Move(p.FirstVoiceShiftedRigth.Reference, p.FirstVoice,
		p.RotatedAxis.Length, 0)
	p.Shapes = append(p.Shapes, p.FirstVoiceShiftedRigth)

	p.SecondVoice.Move(p.SecondVoice.Reference, p.FirstVoice,
		p.NextCircle.CenterX,
		p.NextCircle.CenterY+float64(p.PitchDifference)*p.PitchHeight*p.SideLength)
	p.Shapes = append(p.Shapes, p.SecondVoice)

	p.SecondVoiceShiftedRight.Move(p.SecondVoiceShiftedRight.Reference, p.SecondVoice,
		p.RotatedAxis.Length, 0)
	p.Shapes = append(p.Shapes, p.SecondVoiceShiftedRight)

	nbMeasureToJumpForFirstVoice := p.computeThemeNotesShapes(p.FirstVoice, p.FirstVoiceNotes)
	p.Shapes = append(p.Shapes, p.FirstVoiceNotes)

	p.computeThemeNotesShapes(p.FirstVoiceShiftedRigth, p.FirstVoiceNotesShiftedRight)
	p.Shapes = append(p.Shapes, p.FirstVoiceNotesShiftedRight)

	nbMeasureToJumpForSecondVoice := p.computeThemeNotesShapes(p.SecondVoice, p.SecondVoiceNotes)
	p.ActualBeatsTemporalShift = nbMeasureToJumpForSecondVoice - nbMeasureToJumpForFirstVoice
	p.Shapes = append(p.Shapes, p.SecondVoiceNotes)

	p.computeThemeNotesShapes(p.SecondVoiceShiftedRight, p.SecondVoiceNotesShiftedRight)
	p.Shapes = append(p.Shapes, p.SecondVoiceNotesShiftedRight)

}
