package models

func (p *Parameter) ComputeNoteInfos() {

	if len(p.NoteInfos) == p.NbMeasureLinesPerCurve {
		return
	}

}
