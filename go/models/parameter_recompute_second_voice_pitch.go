package models

func (parameter *Parameter) recomputeSecondVoicePitch(firstVoice []*Circle, secondVoice []*Circle) {

	shift := parameter.ActualBeatsTemporalShift
	_ = shift
	pitchDiff := parameter.PitchDifference
	_ = pitchDiff

	for i, firstVoiceNode := range firstVoice {

		// to find the index of the second voice note, compared to the first note index
		// you have to substract the shift.
		secondVoiceNote := secondVoice[(i-shift+len(firstVoice))%len(firstVoice)]
		secondVoiceNote.Pitch = firstVoiceNode.Pitch + pitchDiff

	}
}
