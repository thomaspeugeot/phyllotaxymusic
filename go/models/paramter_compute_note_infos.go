package models

func (p *Parameter) ComputeNoteInfos(stage *StageStruct) {

	if len(p.NoteInfos) == p.NbMeasureLinesPerCurve {
		return
	}

	// remove extra note infos
	if len(p.NoteInfos) > p.NbMeasureLinesPerCurve {
		for _, ni := range p.NoteInfos[p.NbMeasureLinesPerCurve:] {
			ni.Unstage(stage)
		}
		p.NoteInfos = p.NoteInfos[:p.NbMeasureLinesPerCurve]
	}
}
