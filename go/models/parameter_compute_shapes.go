package models

func (p *Parameter) ComputeShapes(stage *StageStruct) {

	if p.Z < p.M+p.N+1 {
		p.Z = p.M + p.N + 1
	}

	p.Shapes = p.Shapes[:0]

	p.Shapes = append(p.Shapes, p.HorizontalAxis)
	p.Shapes = append(p.Shapes, p.VerticalAxis)

	p.ComputeInitialRhombus()
	p.Shapes = append(p.Shapes, p.InitialRhombus)

	p.ComputeInitialCircle()
	p.Shapes = append(p.Shapes, p.InitialCircle)

	p.ComputeInitialRhombusGrid(stage)
	p.Shapes = append(p.Shapes, p.InitialRhombusGrid)

	p.ComputeInitialCircleGrid(stage)
	p.Shapes = append(p.Shapes, p.InitialCircleGrid)

	p.ComputeInitialAxis()
	p.Shapes = append(p.Shapes, p.InitialAxis)

	p.computeRotatedAxis()
	p.Shapes = append(p.Shapes, p.RotatedAxis)

	p.computeRotatedRhombus()
	p.Shapes = append(p.Shapes, p.RotatedRhombus)

	p.computeRotatedRhombusGrid(stage)
	p.Shapes = append(p.Shapes, p.RotatedRhombusGrid)

	p.computeRotatedCircleGrid(stage)
	p.Shapes = append(p.Shapes, p.RotatedCircleGrid)

	p.ComputeNextRhombus()
	p.Shapes = append(p.Shapes, p.NextRhombus)

	p.ComputeNextCircle()
	p.Shapes = append(p.Shapes, p.NextCircle)

	p.ComputeGrowingRhombusGrid()
	p.Shapes = append(p.Shapes, p.GrowingRhombusGrid)

	p.ComputeGrowingCircleGrid()
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
	p.Shapes = append(p.Shapes, p.GrowthCurveSegment)

	p.ComputeGrowthCurve()
	p.Shapes = append(p.Shapes, p.GrowthCurve)

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

	p.ComputeFKey()
	p.Shapes = append(p.Shapes, p.Fkey)

	p.ComputePitchLines()
	p.Shapes = append(p.Shapes, p.PitchLines)

	p.ComputeMeasureLines()
	p.Shapes = append(p.Shapes, p.MeasureLines)

	p.FirstVoice.Move(p.FirstVoice.Reference, p.GrowthCurve,
		float64(p.FirstVoiceShiftX)*p.SideLength,
		float64(p.FirstVoiceShiftY)*p.SideLength)
	p.Shapes = append(p.Shapes, p.FirstVoice)

	p.FirstVoiceShiftRigth.Move(p.FirstVoiceShiftRigth.Reference, p.FirstVoice,
		p.RotatedAxis.Length, 0)
	p.Shapes = append(p.Shapes, p.FirstVoiceShiftRigth)

	p.SecondVoice.Move(p.SecondVoice.Reference, p.FirstVoice,
		p.NextCircle.CenterX,
		p.NextCircle.CenterY+float64(p.PitchDifference)*p.PitchHeight*p.SideLength)
	p.Shapes = append(p.Shapes, p.SecondVoice)

	p.SecondVoiceShiftedRight.Move(p.SecondVoiceShiftedRight.Reference, p.SecondVoice,
		p.RotatedAxis.Length, 0)
	p.Shapes = append(p.Shapes, p.SecondVoiceShiftedRight)

	p.computeVoiceNotes(p.FirstVoice, p.FirstVoiceNotes)
	p.Shapes = append(p.Shapes, p.FirstVoiceNotes)

	p.FirstVoiceNotesShiftedRight.Move(p.FirstVoiceNotes.Reference, p.FirstVoiceNotes,
		p.RotatedAxis.Length, 0)
	p.Shapes = append(p.Shapes, p.FirstVoiceNotesShiftedRight)

	p.computeVoiceNotes(p.SecondVoice, p.SecondVoiceNotes)
	p.Shapes = append(p.Shapes, p.SecondVoiceNotes)

	p.SecondVoiceNotesShiftedRight.Move(p.SecondVoiceNotesShiftedRight.Reference, p.SecondVoiceNotes,
		p.RotatedAxis.Length, 0)
	p.Shapes = append(p.Shapes, p.SecondVoiceNotesShiftedRight)

	p.ComputeNoteInfos()
}
